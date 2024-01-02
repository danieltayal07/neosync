// @generated by protoc-gen-connect-es v1.0.0 with parameter "target=ts,import_extension=.js"
// @generated from file mgmt/v1alpha1/transformer.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateUserDefinedTransformerRequest, CreateUserDefinedTransformerResponse, DeleteUserDefinedTransformerRequest, DeleteUserDefinedTransformerResponse, GetSystemTransformersRequest, GetSystemTransformersResponse, GetUserDefinedTransformerByIdRequest, GetUserDefinedTransformerByIdResponse, GetUserDefinedTransformersRequest, GetUserDefinedTransformersResponse, IsTransformerNameAvailableRequest, IsTransformerNameAvailableResponse, UpdateUserDefinedTransformerRequest, UpdateUserDefinedTransformerResponse, ValidateUserJavascriptCodeRequest, ValidateUserJavascriptCodeResponse } from "./transformer_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service mgmt.v1alpha1.TransformersService
 */
export const TransformersService = {
  typeName: "mgmt.v1alpha1.TransformersService",
  methods: {
    /**
     * @generated from rpc mgmt.v1alpha1.TransformersService.GetSystemTransformers
     */
    getSystemTransformers: {
      name: "GetSystemTransformers",
      I: GetSystemTransformersRequest,
      O: GetSystemTransformersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.TransformersService.GetUserDefinedTransformers
     */
    getUserDefinedTransformers: {
      name: "GetUserDefinedTransformers",
      I: GetUserDefinedTransformersRequest,
      O: GetUserDefinedTransformersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.TransformersService.GetUserDefinedTransformerById
     */
    getUserDefinedTransformerById: {
      name: "GetUserDefinedTransformerById",
      I: GetUserDefinedTransformerByIdRequest,
      O: GetUserDefinedTransformerByIdResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.TransformersService.CreateUserDefinedTransformer
     */
    createUserDefinedTransformer: {
      name: "CreateUserDefinedTransformer",
      I: CreateUserDefinedTransformerRequest,
      O: CreateUserDefinedTransformerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.TransformersService.DeleteUserDefinedTransformer
     */
    deleteUserDefinedTransformer: {
      name: "DeleteUserDefinedTransformer",
      I: DeleteUserDefinedTransformerRequest,
      O: DeleteUserDefinedTransformerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.TransformersService.UpdateUserDefinedTransformer
     */
    updateUserDefinedTransformer: {
      name: "UpdateUserDefinedTransformer",
      I: UpdateUserDefinedTransformerRequest,
      O: UpdateUserDefinedTransformerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.TransformersService.IsTransformerNameAvailable
     */
    isTransformerNameAvailable: {
      name: "IsTransformerNameAvailable",
      I: IsTransformerNameAvailableRequest,
      O: IsTransformerNameAvailableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.TransformersService.ValidateUserJavascriptCode
     */
    validateUserJavascriptCode: {
      name: "ValidateUserJavascriptCode",
      I: ValidateUserJavascriptCodeRequest,
      O: ValidateUserJavascriptCodeResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

