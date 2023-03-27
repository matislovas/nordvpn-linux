// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: fileshare.proto

package pb

import (
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

// ServiceErrorCode defines a set of error codes whose handling
// does not depend on any specific command used.
type ServiceErrorCode int32

const (
	ServiceErrorCode_NOT_LOGGED_IN    ServiceErrorCode = 0
	ServiceErrorCode_MESH_NOT_ENABLED ServiceErrorCode = 1
	ServiceErrorCode_INTERNAL_FAILURE ServiceErrorCode = 2
)

// Enum value maps for ServiceErrorCode.
var (
	ServiceErrorCode_name = map[int32]string{
		0: "NOT_LOGGED_IN",
		1: "MESH_NOT_ENABLED",
		2: "INTERNAL_FAILURE",
	}
	ServiceErrorCode_value = map[string]int32{
		"NOT_LOGGED_IN":    0,
		"MESH_NOT_ENABLED": 1,
		"INTERNAL_FAILURE": 2,
	}
)

func (x ServiceErrorCode) Enum() *ServiceErrorCode {
	p := new(ServiceErrorCode)
	*p = x
	return p
}

func (x ServiceErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_fileshare_proto_enumTypes[0].Descriptor()
}

func (ServiceErrorCode) Type() protoreflect.EnumType {
	return &file_fileshare_proto_enumTypes[0]
}

func (x ServiceErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceErrorCode.Descriptor instead.
func (ServiceErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_fileshare_proto_rawDescGZIP(), []int{0}
}

// FileshareErrorCode defines a set of fileshare specific error codes.
type FileshareErrorCode int32

const (
	FileshareErrorCode_LIB_FAILURE                   FileshareErrorCode = 0
	FileshareErrorCode_TRANSFER_NOT_FOUND            FileshareErrorCode = 1
	FileshareErrorCode_INVALID_PEER                  FileshareErrorCode = 2
	FileshareErrorCode_FILE_NOT_FOUND                FileshareErrorCode = 3
	FileshareErrorCode_ACCEPT_ALL_FILES_FAILED       FileshareErrorCode = 5 // Accept failed for all files
	FileshareErrorCode_ACCEPT_OUTGOING               FileshareErrorCode = 6 // Can't accept outgoing transfers
	FileshareErrorCode_ALREADY_ACCEPTED              FileshareErrorCode = 7
	FileshareErrorCode_FILE_INVALIDATED              FileshareErrorCode = 8
	FileshareErrorCode_TRANSFER_INVALIDATED          FileshareErrorCode = 9
	FileshareErrorCode_TOO_MANY_FILES                FileshareErrorCode = 10
	FileshareErrorCode_DIRECTORY_TOO_DEEP            FileshareErrorCode = 11
	FileshareErrorCode_SENDING_NOT_ALLOWED           FileshareErrorCode = 12
	FileshareErrorCode_PEER_DISCONNECTED             FileshareErrorCode = 13
	FileshareErrorCode_FILE_NOT_IN_PROGRESS          FileshareErrorCode = 14 // Returned when user tries to cancel a file that is not in flight
	FileshareErrorCode_TRANSFER_NOT_CREATED          FileshareErrorCode = 15 // When libdrop doesn't return transfer ID, most likely permission issue
	FileshareErrorCode_NOT_ENOUGH_SPACE              FileshareErrorCode = 16 // Transfer larger than available hard drive space
	FileshareErrorCode_ACCEPT_DIR_NOT_FOUND          FileshareErrorCode = 17
	FileshareErrorCode_ACCEPT_DIR_IS_A_SYMLINK       FileshareErrorCode = 18
	FileshareErrorCode_ACCEPT_DIR_IS_NOT_A_DIRECTORY FileshareErrorCode = 19
	FileshareErrorCode_NO_FILES                      FileshareErrorCode = 20
)

// Enum value maps for FileshareErrorCode.
var (
	FileshareErrorCode_name = map[int32]string{
		0:  "LIB_FAILURE",
		1:  "TRANSFER_NOT_FOUND",
		2:  "INVALID_PEER",
		3:  "FILE_NOT_FOUND",
		5:  "ACCEPT_ALL_FILES_FAILED",
		6:  "ACCEPT_OUTGOING",
		7:  "ALREADY_ACCEPTED",
		8:  "FILE_INVALIDATED",
		9:  "TRANSFER_INVALIDATED",
		10: "TOO_MANY_FILES",
		11: "DIRECTORY_TOO_DEEP",
		12: "SENDING_NOT_ALLOWED",
		13: "PEER_DISCONNECTED",
		14: "FILE_NOT_IN_PROGRESS",
		15: "TRANSFER_NOT_CREATED",
		16: "NOT_ENOUGH_SPACE",
		17: "ACCEPT_DIR_NOT_FOUND",
		18: "ACCEPT_DIR_IS_A_SYMLINK",
		19: "ACCEPT_DIR_IS_NOT_A_DIRECTORY",
		20: "NO_FILES",
	}
	FileshareErrorCode_value = map[string]int32{
		"LIB_FAILURE":                   0,
		"TRANSFER_NOT_FOUND":            1,
		"INVALID_PEER":                  2,
		"FILE_NOT_FOUND":                3,
		"ACCEPT_ALL_FILES_FAILED":       5,
		"ACCEPT_OUTGOING":               6,
		"ALREADY_ACCEPTED":              7,
		"FILE_INVALIDATED":              8,
		"TRANSFER_INVALIDATED":          9,
		"TOO_MANY_FILES":                10,
		"DIRECTORY_TOO_DEEP":            11,
		"SENDING_NOT_ALLOWED":           12,
		"PEER_DISCONNECTED":             13,
		"FILE_NOT_IN_PROGRESS":          14,
		"TRANSFER_NOT_CREATED":          15,
		"NOT_ENOUGH_SPACE":              16,
		"ACCEPT_DIR_NOT_FOUND":          17,
		"ACCEPT_DIR_IS_A_SYMLINK":       18,
		"ACCEPT_DIR_IS_NOT_A_DIRECTORY": 19,
		"NO_FILES":                      20,
	}
)

func (x FileshareErrorCode) Enum() *FileshareErrorCode {
	p := new(FileshareErrorCode)
	*p = x
	return p
}

func (x FileshareErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileshareErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_fileshare_proto_enumTypes[1].Descriptor()
}

func (FileshareErrorCode) Type() protoreflect.EnumType {
	return &file_fileshare_proto_enumTypes[1]
}

func (x FileshareErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileshareErrorCode.Descriptor instead.
func (FileshareErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_fileshare_proto_rawDescGZIP(), []int{1}
}

// Used when there is no error or there is no data to be sent
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileshare_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_fileshare_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_fileshare_proto_rawDescGZIP(), []int{0}
}

// Generic error to be used through all responses. If empty then no error occured.
// If there's no data to be returned then this can be used as a response type,
// otherwise it should be included as a field in the response.
// Response handlers should always firstly check whether error is Empty (like Go err != nil check)
//
// Previously (in meshnet) we have used oneof to either return data or an error. But the problem
// with oneof is that when it is used the same messages are returned as different types
// (SendResponse_FileshareResponse and ReceiveResponse_FileshareResponse for example). Because of that
// we couldn't DRY their handling and that resulted in lots of almost duplicate code.
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*Error_Empty
	//	*Error_ServiceError
	//	*Error_FileshareError
	Response isError_Response `protobuf_oneof:"response"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileshare_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_fileshare_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_fileshare_proto_rawDescGZIP(), []int{1}
}

func (m *Error) GetResponse() isError_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *Error) GetEmpty() *Empty {
	if x, ok := x.GetResponse().(*Error_Empty); ok {
		return x.Empty
	}
	return nil
}

func (x *Error) GetServiceError() ServiceErrorCode {
	if x, ok := x.GetResponse().(*Error_ServiceError); ok {
		return x.ServiceError
	}
	return ServiceErrorCode_NOT_LOGGED_IN
}

func (x *Error) GetFileshareError() FileshareErrorCode {
	if x, ok := x.GetResponse().(*Error_FileshareError); ok {
		return x.FileshareError
	}
	return FileshareErrorCode_LIB_FAILURE
}

type isError_Response interface {
	isError_Response()
}

type Error_Empty struct {
	Empty *Empty `protobuf:"bytes,1,opt,name=empty,proto3,oneof"`
}

type Error_ServiceError struct {
	ServiceError ServiceErrorCode `protobuf:"varint,2,opt,name=service_error,json=serviceError,proto3,enum=filesharepb.ServiceErrorCode,oneof"`
}

type Error_FileshareError struct {
	FileshareError FileshareErrorCode `protobuf:"varint,3,opt,name=fileshare_error,json=fileshareError,proto3,enum=filesharepb.FileshareErrorCode,oneof"`
}

func (*Error_Empty) isError_Response() {}

func (*Error_ServiceError) isError_Response() {}

func (*Error_FileshareError) isError_Response() {}

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer   string   `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`      // IP to which the request will be sent
	Paths  []string `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`    // Absolute path of the file or dir to be sent
	Silent bool     `protobuf:"varint,3,opt,name=silent,proto3" json:"silent,omitempty"` // Do transfer in background (true) or Report progress info back (false)
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileshare_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileshare_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_fileshare_proto_rawDescGZIP(), []int{2}
}

func (x *SendRequest) GetPeer() string {
	if x != nil {
		return x.Peer
	}
	return ""
}

func (x *SendRequest) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *SendRequest) GetSilent() bool {
	if x != nil {
		return x.Silent
	}
	return false
}

type AcceptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId string   `protobuf:"bytes,1,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"` // ID taken from TransferRequested libdrop event
	DstPath    string   `protobuf:"bytes,2,opt,name=dst_path,json=dstPath,proto3" json:"dst_path,omitempty"`          // Directory to store the received files
	Silent     bool     `protobuf:"varint,3,opt,name=silent,proto3" json:"silent,omitempty"`                          // Do transfer in background (true) or Report progress info back (false)
	Files      []string `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty"`                             // A list of specific files to be accepted
}

func (x *AcceptRequest) Reset() {
	*x = AcceptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileshare_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptRequest) ProtoMessage() {}

func (x *AcceptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileshare_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptRequest.ProtoReflect.Descriptor instead.
func (*AcceptRequest) Descriptor() ([]byte, []int) {
	return file_fileshare_proto_rawDescGZIP(), []int{3}
}

func (x *AcceptRequest) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *AcceptRequest) GetDstPath() string {
	if x != nil {
		return x.DstPath
	}
	return ""
}

func (x *AcceptRequest) GetSilent() bool {
	if x != nil {
		return x.Silent
	}
	return false
}

func (x *AcceptRequest) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	TransferId string `protobuf:"bytes,2,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"` // Newly created transfer's ID
	Progress   uint32 `protobuf:"varint,3,opt,name=progress,proto3" json:"progress,omitempty"`                      // Transfer progress percent
	Status     Status `protobuf:"varint,4,opt,name=status,proto3,enum=filesharepb.Status" json:"status,omitempty"`  // Transfer status
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileshare_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileshare_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_fileshare_proto_rawDescGZIP(), []int{4}
}

func (x *StatusResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *StatusResponse) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *StatusResponse) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *StatusResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

type CancelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId string `protobuf:"bytes,1,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"` // ID taken from TransferRequested libdrop event
}

func (x *CancelRequest) Reset() {
	*x = CancelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileshare_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelRequest) ProtoMessage() {}

func (x *CancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileshare_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelRequest.ProtoReflect.Descriptor instead.
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return file_fileshare_proto_rawDescGZIP(), []int{5}
}

func (x *CancelRequest) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// Transfers are sorted by creation date from oldest to newest
	Transfers []*Transfer `protobuf:"bytes,2,rep,name=transfers,proto3" json:"transfers,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileshare_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileshare_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_fileshare_proto_rawDescGZIP(), []int{6}
}

func (x *ListResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ListResponse) GetTransfers() []*Transfer {
	if x != nil {
		return x.Transfers
	}
	return nil
}

type CancelFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId string `protobuf:"bytes,1,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"` // ID taken from TransferRequested libdrop event
	FileId     string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`             // ID taken from TransferRequested libdrop event
}

func (x *CancelFileRequest) Reset() {
	*x = CancelFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileshare_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelFileRequest) ProtoMessage() {}

func (x *CancelFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileshare_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelFileRequest.ProtoReflect.Descriptor instead.
func (*CancelFileRequest) Descriptor() ([]byte, []int) {
	return file_fileshare_proto_rawDescGZIP(), []int{7}
}

func (x *CancelFileRequest) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *CancelFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

var File_fileshare_proto protoreflect.FileDescriptor

var file_fileshare_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x70, 0x62, 0x1a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xd1, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44, 0x0a,
	0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x4a, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52,
	0x0e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42,
	0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x0a, 0x0b, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x74, 0x22, 0x79, 0x0a, 0x0d,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6c,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x69, 0x6c, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30,
	0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x6d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x22,
	0x4d, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x2a, 0x51,
	0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x44,
	0x5f, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x53, 0x48, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10,
	0x02, 0x2a, 0xe9, 0x03, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x49, 0x42, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x45, 0x45,
	0x52, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x43, 0x43, 0x45, 0x50,
	0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x5f, 0x4f,
	0x55, 0x54, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4c, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x07, 0x12,
	0x14, 0x0a, 0x10, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45,
	0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x09, 0x12,
	0x12, 0x0a, 0x0e, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x53, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x59,
	0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x44, 0x45, 0x45, 0x50, 0x10, 0x0b, 0x12, 0x17, 0x0a, 0x13, 0x53,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x45, 0x44, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x53,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x18, 0x0a, 0x14, 0x46,
	0x49, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x0e, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x0f, 0x12,
	0x14, 0x0a, 0x10, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x53, 0x50,
	0x41, 0x43, 0x45, 0x10, 0x10, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x5f,
	0x44, 0x49, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x11, 0x12,
	0x1b, 0x0a, 0x17, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x5f, 0x49, 0x53,
	0x5f, 0x41, 0x5f, 0x53, 0x59, 0x4d, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x12, 0x12, 0x21, 0x0a, 0x1d,
	0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x41, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x13, 0x12,
	0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x14, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x72, 0x64,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x6e, 0x6f, 0x72, 0x64, 0x76, 0x70, 0x6e,
	0x2d, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fileshare_proto_rawDescOnce sync.Once
	file_fileshare_proto_rawDescData = file_fileshare_proto_rawDesc
)

func file_fileshare_proto_rawDescGZIP() []byte {
	file_fileshare_proto_rawDescOnce.Do(func() {
		file_fileshare_proto_rawDescData = protoimpl.X.CompressGZIP(file_fileshare_proto_rawDescData)
	})
	return file_fileshare_proto_rawDescData
}

var file_fileshare_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_fileshare_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_fileshare_proto_goTypes = []interface{}{
	(ServiceErrorCode)(0),     // 0: filesharepb.ServiceErrorCode
	(FileshareErrorCode)(0),   // 1: filesharepb.FileshareErrorCode
	(*Empty)(nil),             // 2: filesharepb.Empty
	(*Error)(nil),             // 3: filesharepb.Error
	(*SendRequest)(nil),       // 4: filesharepb.SendRequest
	(*AcceptRequest)(nil),     // 5: filesharepb.AcceptRequest
	(*StatusResponse)(nil),    // 6: filesharepb.StatusResponse
	(*CancelRequest)(nil),     // 7: filesharepb.CancelRequest
	(*ListResponse)(nil),      // 8: filesharepb.ListResponse
	(*CancelFileRequest)(nil), // 9: filesharepb.CancelFileRequest
	(Status)(0),               // 10: filesharepb.Status
	(*Transfer)(nil),          // 11: filesharepb.Transfer
}
var file_fileshare_proto_depIdxs = []int32{
	2,  // 0: filesharepb.Error.empty:type_name -> filesharepb.Empty
	0,  // 1: filesharepb.Error.service_error:type_name -> filesharepb.ServiceErrorCode
	1,  // 2: filesharepb.Error.fileshare_error:type_name -> filesharepb.FileshareErrorCode
	3,  // 3: filesharepb.StatusResponse.error:type_name -> filesharepb.Error
	10, // 4: filesharepb.StatusResponse.status:type_name -> filesharepb.Status
	3,  // 5: filesharepb.ListResponse.error:type_name -> filesharepb.Error
	11, // 6: filesharepb.ListResponse.transfers:type_name -> filesharepb.Transfer
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_fileshare_proto_init() }
func file_fileshare_proto_init() {
	if File_fileshare_proto != nil {
		return
	}
	file_transfer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fileshare_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_fileshare_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_fileshare_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_fileshare_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptRequest); i {
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
		file_fileshare_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_fileshare_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelRequest); i {
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
		file_fileshare_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_fileshare_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelFileRequest); i {
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
	file_fileshare_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Error_Empty)(nil),
		(*Error_ServiceError)(nil),
		(*Error_FileshareError)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fileshare_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fileshare_proto_goTypes,
		DependencyIndexes: file_fileshare_proto_depIdxs,
		EnumInfos:         file_fileshare_proto_enumTypes,
		MessageInfos:      file_fileshare_proto_msgTypes,
	}.Build()
	File_fileshare_proto = out.File
	file_fileshare_proto_rawDesc = nil
	file_fileshare_proto_goTypes = nil
	file_fileshare_proto_depIdxs = nil
}
