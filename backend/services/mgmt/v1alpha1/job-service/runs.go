package v1alpha1_jobservice

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	logger_interceptor "github.com/nucleuscloud/neosync/backend/internal/connect/interceptors/logger"
	"github.com/nucleuscloud/neosync/backend/internal/dtomaps"
	nucleuserrors "github.com/nucleuscloud/neosync/backend/internal/errors"
	"github.com/nucleuscloud/neosync/backend/internal/nucleusdb"
	"go.temporal.io/api/enums/v1"
	workflowpb "go.temporal.io/api/workflow/v1"
	"go.temporal.io/api/workflowservice/v1"
	temporalclient "go.temporal.io/sdk/client"
	"golang.org/x/sync/errgroup"
)

func (s *Service) GetJobRuns(
	ctx context.Context,
	req *connect.Request[mgmtv1alpha1.GetJobRunsRequest],
) (*connect.Response[mgmtv1alpha1.GetJobRunsResponse], error) {
	logger := logger_interceptor.GetLoggerFromContextOrDefault(ctx)

	var accountId string
	var workflows []*workflowpb.WorkflowExecutionInfo
	switch id := req.Msg.Id.(type) {
	case *mgmtv1alpha1.GetJobRunsRequest_JobId:
		jobUuid, err := nucleusdb.ToUuid(id.JobId)
		if err != nil {
			return nil, err
		}
		job, err := s.db.Q.GetJobById(ctx, jobUuid)
		if err != nil {
			return nil, err
		}
		accountId = nucleusdb.UUIDString(job.AccountID)
		workflows, err = getWorkflowExecutionsByJobIds(ctx, s.temporalClient, logger, s.cfg.TemporalNamespace, []string{id.JobId})
		if err != nil {
			return nil, err
		}
	case *mgmtv1alpha1.GetJobRunsRequest_AccountId:
		accountId = id.AccountId
		accountUuid, err := nucleusdb.ToUuid(accountId)
		if err != nil {
			return nil, err
		}
		jobs, err := s.db.Q.GetJobsByAccount(ctx, accountUuid)
		if err != nil {
			return nil, err
		}
		jobIds := []string{}
		for i := range jobs {
			job := jobs[i]
			jobIds = append(jobIds, nucleusdb.UUIDString(job.ID))
		}
		workflows, err = getWorkflowExecutionsByJobIds(ctx, s.temporalClient, logger, s.cfg.TemporalNamespace, jobIds)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("must provide jobId or accountId")
	}

	_, err := s.verifyUserInAccount(ctx, accountId)
	if err != nil {
		return nil, err
	}

	runs := make([]*mgmtv1alpha1.JobRun, len(workflows))
	errGrp, errCtx := errgroup.WithContext(ctx)
	for index, workflow := range workflows {
		index := index
		workflow := workflow
		errGrp.Go(func() error {
			res, err := s.temporalClient.DescribeWorkflowExecution(errCtx, workflow.Execution.WorkflowId, workflow.Execution.RunId)
			if err != nil {
				return err
			}
			runs[index] = dtomaps.ToJobRunDto(logger, res)
			return nil
		})
	}

	err = errGrp.Wait()
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&mgmtv1alpha1.GetJobRunsResponse{
		JobRuns: runs,
	}), nil
}

func (s *Service) GetJobRun(
	ctx context.Context,
	req *connect.Request[mgmtv1alpha1.GetJobRunRequest],
) (*connect.Response[mgmtv1alpha1.GetJobRunResponse], error) {
	logger := logger_interceptor.GetLoggerFromContextOrDefault(ctx)
	logger = logger.With("runId", req.Msg.Id)
	run, err := getWorkflowExecutionsByRunId(ctx, s.temporalClient, s.cfg.TemporalNamespace, req.Msg.Id)
	if err != nil {
		logger.Error(fmt.Errorf("unable to retrieve job run: %w", err).Error())
		return nil, err
	}
	jobId := dtomaps.GetJobIdFromWorkflow(logger, run.GetSearchAttributes())
	jobUuid, err := nucleusdb.ToUuid(jobId)
	if err != nil {
		return nil, err
	}
	job, err := s.db.Q.GetJobById(ctx, jobUuid)
	if err != nil {
		return nil, err
	}
	accountId := nucleusdb.UUIDString(job.AccountID)
	_, err = s.verifyUserInAccount(ctx, accountId)
	if err != nil {
		return nil, err
	}
	res, err := s.temporalClient.DescribeWorkflowExecution(ctx, run.Execution.WorkflowId, run.Execution.RunId)
	if err != nil {
		return nil, err
	}

	dto := dtomaps.ToJobRunDto(logger, res)
	return connect.NewResponse(&mgmtv1alpha1.GetJobRunResponse{
		JobRun: dto,
	}), nil
}

func (s *Service) GetJobRunEvents(
	ctx context.Context,
	req *connect.Request[mgmtv1alpha1.GetJobRunEventsRequest],
) (*connect.Response[mgmtv1alpha1.GetJobRunEventsResponse], error) {
	logger := logger_interceptor.GetLoggerFromContextOrDefault(ctx)
	logger = logger.With("runId", req.Msg.Id)
	run, err := getWorkflowExecutionsByRunId(ctx, s.temporalClient, s.cfg.TemporalNamespace, req.Msg.Id)
	if err != nil {
		logger.Error(fmt.Errorf("unable to retrieve job run: %w", err).Error())
		return nil, err
	}

	jobId := dtomaps.GetJobIdFromWorkflow(logger, run.GetSearchAttributes())
	jobUuid, err := nucleusdb.ToUuid(jobId)
	if err != nil {
		return nil, err
	}
	job, err := s.db.Q.GetJobById(ctx, jobUuid)
	if err != nil {
		return nil, err
	}
	accountId := nucleusdb.UUIDString(job.AccountID)
	_, err = s.verifyUserInAccount(ctx, accountId)
	if err != nil {
		return nil, err
	}

	activityNameMap := map[int64]string{}
	events := []*mgmtv1alpha1.JobRunEvent{}
	iter := s.temporalClient.GetWorkflowHistory(
		ctx,
		run.Execution.WorkflowId,
		run.Execution.RunId,
		false,
		enums.HISTORY_EVENT_FILTER_TYPE_ALL_EVENT,
	)
	for iter.HasNext() {
		event, err := iter.Next()
		if err != nil {
			return nil, err
		}

		// workflow events
		if event.GetWorkflowExecutionStartedEventAttributes() != nil {
			workflowEvent := event.GetWorkflowExecutionStartedEventAttributes()
			events = append(events, dtomaps.ToJobRunEventDto(event, workflowEvent.WorkflowType.Name, "Workflow Started"))
			activityNameMap[event.EventId] = workflowEvent.WorkflowType.Name
		}
		if event.GetWorkflowExecutionCompletedEventAttributes() != nil {
			workflowEvent := event.GetWorkflowExecutionCompletedEventAttributes()
			name := activityNameMap[workflowEvent.WorkflowTaskCompletedEventId]
			events = append(events, dtomaps.ToJobRunEventDto(event, name, "Workflow Completed"))
		}

		// activity events
		if event.GetActivityTaskScheduledEventAttributes() != nil {
			attributes := event.GetActivityTaskScheduledEventAttributes()
			events = append(events, dtomaps.ToJobRunEventDto(event, attributes.ActivityType.Name, "Activity Scheduled"))
			activityNameMap[event.EventId] = attributes.ActivityType.Name
		}

		if event.GetActivityTaskFailedEventAttributes() != nil {
			attributes := event.GetActivityTaskFailedEventAttributes()
			name := activityNameMap[attributes.ScheduledEventId]
			events = append(events, dtomaps.ToJobRunEventDto(event, name, "Activity Failed"))
		}

		if event.GetActivityTaskStartedEventAttributes() != nil {
			attributes := event.GetActivityTaskStartedEventAttributes()
			name := activityNameMap[attributes.ScheduledEventId]
			events = append(events, dtomaps.ToJobRunEventDto(event, name, "Activity Started"))
		}

		if event.GetActivityTaskCompletedEventAttributes() != nil {
			attributes := event.GetActivityTaskCompletedEventAttributes()
			name := activityNameMap[attributes.ScheduledEventId]
			events = append(events, dtomaps.ToJobRunEventDto(event, name, "Activity Completed"))
		}
	}

	return connect.NewResponse(&mgmtv1alpha1.GetJobRunEventsResponse{
		Events: events,
	}), nil
}

func (s *Service) CreateJobRun(
	ctx context.Context,
	req *connect.Request[mgmtv1alpha1.CreateJobRunRequest],
) (*connect.Response[mgmtv1alpha1.CreateJobRunResponse], error) {

	return connect.NewResponse(&mgmtv1alpha1.CreateJobRunResponse{}), nil
}

func (s *Service) CancelJobRun(
	ctx context.Context,
	req *connect.Request[mgmtv1alpha1.CancelJobRunRequest],
) (*connect.Response[mgmtv1alpha1.CancelJobRunResponse], error) {

	return connect.NewResponse(&mgmtv1alpha1.CancelJobRunResponse{}), nil
}

func (s *Service) DeleteJobRun(
	ctx context.Context,
	req *connect.Request[mgmtv1alpha1.DeleteJobRunRequest],
) (*connect.Response[mgmtv1alpha1.DeleteJobRunResponse], error) {
	return connect.NewResponse(&mgmtv1alpha1.DeleteJobRunResponse{}), nil
}

func getWorkflowExecutionsByRunId(
	ctx context.Context,
	tc temporalclient.Client,
	namespace string,
	runId string,
) (*workflowpb.WorkflowExecutionInfo, error) {
	query := fmt.Sprintf("WorkflowId = %q", runId)
	request := &workflowservice.ListWorkflowExecutionsRequest{Query: query, Namespace: namespace}
	resp, err := tc.ListWorkflow(ctx, request)
	if err != nil {
		return nil, err
	}
	if len(resp.Executions) == 0 {
		return nil, nucleuserrors.NewNotFound("job run not found")
	}
	if len(resp.Executions) > 1 {
		return nil, nucleuserrors.NewInternalError("found more than 1 job run")
	}
	return resp.Executions[0], nil
}
