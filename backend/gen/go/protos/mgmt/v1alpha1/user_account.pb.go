// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: mgmt/v1alpha1/user_account.proto

package mgmtv1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserAccountType int32

const (
	UserAccountType_USER_ACCOUNT_TYPE_UNSPECIFIED UserAccountType = 0
	UserAccountType_USER_ACCOUNT_TYPE_PERSONAL    UserAccountType = 1
	UserAccountType_USER_ACCOUNT_TYPE_TEAM        UserAccountType = 2
)

// Enum value maps for UserAccountType.
var (
	UserAccountType_name = map[int32]string{
		0: "USER_ACCOUNT_TYPE_UNSPECIFIED",
		1: "USER_ACCOUNT_TYPE_PERSONAL",
		2: "USER_ACCOUNT_TYPE_TEAM",
	}
	UserAccountType_value = map[string]int32{
		"USER_ACCOUNT_TYPE_UNSPECIFIED": 0,
		"USER_ACCOUNT_TYPE_PERSONAL":    1,
		"USER_ACCOUNT_TYPE_TEAM":        2,
	}
)

func (x UserAccountType) Enum() *UserAccountType {
	p := new(UserAccountType)
	*p = x
	return p
}

func (x UserAccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserAccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_mgmt_v1alpha1_user_account_proto_enumTypes[0].Descriptor()
}

func (UserAccountType) Type() protoreflect.EnumType {
	return &file_mgmt_v1alpha1_user_account_proto_enumTypes[0]
}

func (x UserAccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserAccountType.Descriptor instead.
func (UserAccountType) EnumDescriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{0}
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{0}
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetUserRequest) Reset() {
	*x = SetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserRequest) ProtoMessage() {}

func (x *SetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserRequest.ProtoReflect.Descriptor instead.
func (*SetUserRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{2}
}

type SetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SetUserResponse) Reset() {
	*x = SetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserResponse) ProtoMessage() {}

func (x *SetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserResponse.ProtoReflect.Descriptor instead.
func (*SetUserResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{3}
}

func (x *SetUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserAccountsRequest) Reset() {
	*x = GetUserAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAccountsRequest) ProtoMessage() {}

func (x *GetUserAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAccountsRequest.ProtoReflect.Descriptor instead.
func (*GetUserAccountsRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{4}
}

type GetUserAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*UserAccount `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *GetUserAccountsResponse) Reset() {
	*x = GetUserAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAccountsResponse) ProtoMessage() {}

func (x *GetUserAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAccountsResponse.ProtoReflect.Descriptor instead.
func (*GetUserAccountsResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserAccountsResponse) GetAccounts() []*UserAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type UserAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type UserAccountType `protobuf:"varint,3,opt,name=type,proto3,enum=mgmt.v1alpha1.UserAccountType" json:"type,omitempty"`
}

func (x *UserAccount) Reset() {
	*x = UserAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccount) ProtoMessage() {}

func (x *UserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAccount.ProtoReflect.Descriptor instead.
func (*UserAccount) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{6}
}

func (x *UserAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserAccount) GetType() UserAccountType {
	if x != nil {
		return x.Type
	}
	return UserAccountType_USER_ACCOUNT_TYPE_UNSPECIFIED
}

type ConvertPersonalToTeamAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConvertPersonalToTeamAccountRequest) Reset() {
	*x = ConvertPersonalToTeamAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertPersonalToTeamAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertPersonalToTeamAccountRequest) ProtoMessage() {}

func (x *ConvertPersonalToTeamAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertPersonalToTeamAccountRequest.ProtoReflect.Descriptor instead.
func (*ConvertPersonalToTeamAccountRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{7}
}

type ConvertPersonalToTeamAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConvertPersonalToTeamAccountResponse) Reset() {
	*x = ConvertPersonalToTeamAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertPersonalToTeamAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertPersonalToTeamAccountResponse) ProtoMessage() {}

func (x *ConvertPersonalToTeamAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertPersonalToTeamAccountResponse.ProtoReflect.Descriptor instead.
func (*ConvertPersonalToTeamAccountResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{8}
}

type SetPersonalAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetPersonalAccountRequest) Reset() {
	*x = SetPersonalAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPersonalAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPersonalAccountRequest) ProtoMessage() {}

func (x *SetPersonalAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPersonalAccountRequest.ProtoReflect.Descriptor instead.
func (*SetPersonalAccountRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{9}
}

type SetPersonalAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *SetPersonalAccountResponse) Reset() {
	*x = SetPersonalAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPersonalAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPersonalAccountResponse) ProtoMessage() {}

func (x *SetPersonalAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPersonalAccountResponse.ProtoReflect.Descriptor instead.
func (*SetPersonalAccountResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{10}
}

func (x *SetPersonalAccountResponse) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type IsUserInAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *IsUserInAccountRequest) Reset() {
	*x = IsUserInAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsUserInAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserInAccountRequest) ProtoMessage() {}

func (x *IsUserInAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserInAccountRequest.ProtoReflect.Descriptor instead.
func (*IsUserInAccountRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{11}
}

func (x *IsUserInAccountRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type IsUserInAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *IsUserInAccountResponse) Reset() {
	*x = IsUserInAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsUserInAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserInAccountResponse) ProtoMessage() {}

func (x *IsUserInAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_user_account_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserInAccountResponse.ProtoReflect.Descriptor instead.
func (*IsUserInAccountResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_user_account_proto_rawDescGZIP(), []int{12}
}

func (x *IsUserInAccountResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_mgmt_v1alpha1_user_account_proto protoreflect.FileDescriptor

var file_mgmt_v1alpha1_user_account_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x2a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e,
	0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a,
	0x0a, 0x0f, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x65, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x25,
	0x0a, 0x23, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x54, 0x6f, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x24, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x0a,
	0x19, 0x53, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x1a, 0x53, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x16, 0x49, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x17, 0x49, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x2a, 0x70, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x02, 0x32, 0xed, 0x04, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07, 0x53, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x12, 0x53, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x28, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x53, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x89, 0x01, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x54, 0x65, 0x61,
	0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x54, 0x65,
	0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0f, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xcc, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75,
	0x63, 0x6c, 0x65, 0x75, 0x73, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65, 0x6f, 0x73, 0x79,
	0x6e, 0x63, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6d, 0x67, 0x6d, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x67, 0x6d, 0x74,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x0d, 0x4d, 0x67, 0x6d, 0x74,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x19, 0x4d, 0x67, 0x6d, 0x74,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d, 0x67, 0x6d, 0x74, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mgmt_v1alpha1_user_account_proto_rawDescOnce sync.Once
	file_mgmt_v1alpha1_user_account_proto_rawDescData = file_mgmt_v1alpha1_user_account_proto_rawDesc
)

func file_mgmt_v1alpha1_user_account_proto_rawDescGZIP() []byte {
	file_mgmt_v1alpha1_user_account_proto_rawDescOnce.Do(func() {
		file_mgmt_v1alpha1_user_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_mgmt_v1alpha1_user_account_proto_rawDescData)
	})
	return file_mgmt_v1alpha1_user_account_proto_rawDescData
}

var file_mgmt_v1alpha1_user_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mgmt_v1alpha1_user_account_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_mgmt_v1alpha1_user_account_proto_goTypes = []interface{}{
	(UserAccountType)(0),                         // 0: mgmt.v1alpha1.UserAccountType
	(*GetUserRequest)(nil),                       // 1: mgmt.v1alpha1.GetUserRequest
	(*GetUserResponse)(nil),                      // 2: mgmt.v1alpha1.GetUserResponse
	(*SetUserRequest)(nil),                       // 3: mgmt.v1alpha1.SetUserRequest
	(*SetUserResponse)(nil),                      // 4: mgmt.v1alpha1.SetUserResponse
	(*GetUserAccountsRequest)(nil),               // 5: mgmt.v1alpha1.GetUserAccountsRequest
	(*GetUserAccountsResponse)(nil),              // 6: mgmt.v1alpha1.GetUserAccountsResponse
	(*UserAccount)(nil),                          // 7: mgmt.v1alpha1.UserAccount
	(*ConvertPersonalToTeamAccountRequest)(nil),  // 8: mgmt.v1alpha1.ConvertPersonalToTeamAccountRequest
	(*ConvertPersonalToTeamAccountResponse)(nil), // 9: mgmt.v1alpha1.ConvertPersonalToTeamAccountResponse
	(*SetPersonalAccountRequest)(nil),            // 10: mgmt.v1alpha1.SetPersonalAccountRequest
	(*SetPersonalAccountResponse)(nil),           // 11: mgmt.v1alpha1.SetPersonalAccountResponse
	(*IsUserInAccountRequest)(nil),               // 12: mgmt.v1alpha1.IsUserInAccountRequest
	(*IsUserInAccountResponse)(nil),              // 13: mgmt.v1alpha1.IsUserInAccountResponse
}
var file_mgmt_v1alpha1_user_account_proto_depIdxs = []int32{
	7,  // 0: mgmt.v1alpha1.GetUserAccountsResponse.accounts:type_name -> mgmt.v1alpha1.UserAccount
	0,  // 1: mgmt.v1alpha1.UserAccount.type:type_name -> mgmt.v1alpha1.UserAccountType
	1,  // 2: mgmt.v1alpha1.UserAccountService.GetUser:input_type -> mgmt.v1alpha1.GetUserRequest
	3,  // 3: mgmt.v1alpha1.UserAccountService.SetUser:input_type -> mgmt.v1alpha1.SetUserRequest
	5,  // 4: mgmt.v1alpha1.UserAccountService.GetUserAccounts:input_type -> mgmt.v1alpha1.GetUserAccountsRequest
	10, // 5: mgmt.v1alpha1.UserAccountService.SetPersonalAccount:input_type -> mgmt.v1alpha1.SetPersonalAccountRequest
	8,  // 6: mgmt.v1alpha1.UserAccountService.ConvertPersonalToTeamAccount:input_type -> mgmt.v1alpha1.ConvertPersonalToTeamAccountRequest
	12, // 7: mgmt.v1alpha1.UserAccountService.IsUserInAccount:input_type -> mgmt.v1alpha1.IsUserInAccountRequest
	2,  // 8: mgmt.v1alpha1.UserAccountService.GetUser:output_type -> mgmt.v1alpha1.GetUserResponse
	4,  // 9: mgmt.v1alpha1.UserAccountService.SetUser:output_type -> mgmt.v1alpha1.SetUserResponse
	6,  // 10: mgmt.v1alpha1.UserAccountService.GetUserAccounts:output_type -> mgmt.v1alpha1.GetUserAccountsResponse
	11, // 11: mgmt.v1alpha1.UserAccountService.SetPersonalAccount:output_type -> mgmt.v1alpha1.SetPersonalAccountResponse
	9,  // 12: mgmt.v1alpha1.UserAccountService.ConvertPersonalToTeamAccount:output_type -> mgmt.v1alpha1.ConvertPersonalToTeamAccountResponse
	13, // 13: mgmt.v1alpha1.UserAccountService.IsUserInAccount:output_type -> mgmt.v1alpha1.IsUserInAccountResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_mgmt_v1alpha1_user_account_proto_init() }
func file_mgmt_v1alpha1_user_account_proto_init() {
	if File_mgmt_v1alpha1_user_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mgmt_v1alpha1_user_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAccountsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAccountsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertPersonalToTeamAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertPersonalToTeamAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPersonalAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPersonalAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsUserInAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mgmt_v1alpha1_user_account_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsUserInAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mgmt_v1alpha1_user_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mgmt_v1alpha1_user_account_proto_goTypes,
		DependencyIndexes: file_mgmt_v1alpha1_user_account_proto_depIdxs,
		EnumInfos:         file_mgmt_v1alpha1_user_account_proto_enumTypes,
		MessageInfos:      file_mgmt_v1alpha1_user_account_proto_msgTypes,
	}.Build()
	File_mgmt_v1alpha1_user_account_proto = out.File
	file_mgmt_v1alpha1_user_account_proto_rawDesc = nil
	file_mgmt_v1alpha1_user_account_proto_goTypes = nil
	file_mgmt_v1alpha1_user_account_proto_depIdxs = nil
}
