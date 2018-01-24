// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2beta1/session_entity_type.proto

package dialogflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf3 "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The types of modifications for a session entity type.
type SessionEntityType_EntityOverrideMode int32

const (
	// Not specified. This value should be never used.
	SessionEntityType_ENTITY_OVERRIDE_MODE_UNSPECIFIED SessionEntityType_EntityOverrideMode = 0
	// The collection of session entities overrides the collection of entities
	// in the corresponding developer entity type.
	SessionEntityType_ENTITY_OVERRIDE_MODE_OVERRIDE SessionEntityType_EntityOverrideMode = 1
	// The collection of session entities extends the collection of entities in
	// the corresponding developer entity type.
	// Calls to `ListSessionEntityTypes`, `GetSessionEntityType`,
	// `CreateSessionEntityType` and `UpdateSessionEntityType` return the full
	// collection of entities from the developer entity type in the agent's
	// default language and the session entity type.
	SessionEntityType_ENTITY_OVERRIDE_MODE_SUPPLEMENT SessionEntityType_EntityOverrideMode = 2
)

var SessionEntityType_EntityOverrideMode_name = map[int32]string{
	0: "ENTITY_OVERRIDE_MODE_UNSPECIFIED",
	1: "ENTITY_OVERRIDE_MODE_OVERRIDE",
	2: "ENTITY_OVERRIDE_MODE_SUPPLEMENT",
}
var SessionEntityType_EntityOverrideMode_value = map[string]int32{
	"ENTITY_OVERRIDE_MODE_UNSPECIFIED": 0,
	"ENTITY_OVERRIDE_MODE_OVERRIDE":    1,
	"ENTITY_OVERRIDE_MODE_SUPPLEMENT":  2,
}

func (x SessionEntityType_EntityOverrideMode) String() string {
	return proto.EnumName(SessionEntityType_EntityOverrideMode_name, int32(x))
}
func (SessionEntityType_EntityOverrideMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 0}
}

// Represents a session entity type.
//
// Extends or replaces a developer entity type at the user session level (we
// refer to the entity types defined at the agent level as "developer entity
// types").
//
// Note: session entity types apply to all queries, regardless of the language.
type SessionEntityType struct {
	// Required. The unique identifier of this session entity type. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Required. Indicates whether the additional data should override or
	// supplement the developer entity type definition.
	EntityOverrideMode SessionEntityType_EntityOverrideMode `protobuf:"varint,2,opt,name=entity_override_mode,json=entityOverrideMode,enum=google.cloud.dialogflow.v2beta1.SessionEntityType_EntityOverrideMode" json:"entity_override_mode,omitempty"`
	// Required. The collection of entities associated with this session entity
	// type.
	Entities []*EntityType_Entity `protobuf:"bytes,3,rep,name=entities" json:"entities,omitempty"`
}

func (m *SessionEntityType) Reset()                    { *m = SessionEntityType{} }
func (m *SessionEntityType) String() string            { return proto.CompactTextString(m) }
func (*SessionEntityType) ProtoMessage()               {}
func (*SessionEntityType) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *SessionEntityType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SessionEntityType) GetEntityOverrideMode() SessionEntityType_EntityOverrideMode {
	if m != nil {
		return m.EntityOverrideMode
	}
	return SessionEntityType_ENTITY_OVERRIDE_MODE_UNSPECIFIED
}

func (m *SessionEntityType) GetEntities() []*EntityType_Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

// The request message for
// [SessionEntityTypes.ListSessionEntityTypes][google.cloud.dialogflow.v2beta1.SessionEntityTypes.ListSessionEntityTypes].
type ListSessionEntityTypesRequest struct {
	// Required. The session to list all session entity types from.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. The maximum number of items to return in a single page. By
	// default 100 and at most 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListSessionEntityTypesRequest) Reset()                    { *m = ListSessionEntityTypesRequest{} }
func (m *ListSessionEntityTypesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSessionEntityTypesRequest) ProtoMessage()               {}
func (*ListSessionEntityTypesRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *ListSessionEntityTypesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListSessionEntityTypesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListSessionEntityTypesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for
// [SessionEntityTypes.ListSessionEntityTypes][google.cloud.dialogflow.v2beta1.SessionEntityTypes.ListSessionEntityTypes].
type ListSessionEntityTypesResponse struct {
	// The list of session entity types. There will be a maximum number of items
	// returned based on the page_size field in the request.
	SessionEntityTypes []*SessionEntityType `protobuf:"bytes,1,rep,name=session_entity_types,json=sessionEntityTypes" json:"session_entity_types,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListSessionEntityTypesResponse) Reset()                    { *m = ListSessionEntityTypesResponse{} }
func (m *ListSessionEntityTypesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSessionEntityTypesResponse) ProtoMessage()               {}
func (*ListSessionEntityTypesResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *ListSessionEntityTypesResponse) GetSessionEntityTypes() []*SessionEntityType {
	if m != nil {
		return m.SessionEntityTypes
	}
	return nil
}

func (m *ListSessionEntityTypesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for
// [SessionEntityTypes.GetSessionEntityType][google.cloud.dialogflow.v2beta1.SessionEntityTypes.GetSessionEntityType].
type GetSessionEntityTypeRequest struct {
	// Required. The name of the session entity type. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetSessionEntityTypeRequest) Reset()                    { *m = GetSessionEntityTypeRequest{} }
func (m *GetSessionEntityTypeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSessionEntityTypeRequest) ProtoMessage()               {}
func (*GetSessionEntityTypeRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *GetSessionEntityTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for
// [SessionEntityTypes.CreateSessionEntityType][google.cloud.dialogflow.v2beta1.SessionEntityTypes.CreateSessionEntityType].
type CreateSessionEntityTypeRequest struct {
	// Required. The session to create a session entity type for.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Required. The session entity type to create.
	SessionEntityType *SessionEntityType `protobuf:"bytes,2,opt,name=session_entity_type,json=sessionEntityType" json:"session_entity_type,omitempty"`
}

func (m *CreateSessionEntityTypeRequest) Reset()                    { *m = CreateSessionEntityTypeRequest{} }
func (m *CreateSessionEntityTypeRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSessionEntityTypeRequest) ProtoMessage()               {}
func (*CreateSessionEntityTypeRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *CreateSessionEntityTypeRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateSessionEntityTypeRequest) GetSessionEntityType() *SessionEntityType {
	if m != nil {
		return m.SessionEntityType
	}
	return nil
}

// The request message for
// [SessionEntityTypes.UpdateSessionEntityType][google.cloud.dialogflow.v2beta1.SessionEntityTypes.UpdateSessionEntityType].
type UpdateSessionEntityTypeRequest struct {
	// Required. The entity type to update. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	SessionEntityType *SessionEntityType `protobuf:"bytes,1,opt,name=session_entity_type,json=sessionEntityType" json:"session_entity_type,omitempty"`
	// Optional. The mask to control which fields get updated.
	UpdateMask *google_protobuf3.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateSessionEntityTypeRequest) Reset()                    { *m = UpdateSessionEntityTypeRequest{} }
func (m *UpdateSessionEntityTypeRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateSessionEntityTypeRequest) ProtoMessage()               {}
func (*UpdateSessionEntityTypeRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *UpdateSessionEntityTypeRequest) GetSessionEntityType() *SessionEntityType {
	if m != nil {
		return m.SessionEntityType
	}
	return nil
}

func (m *UpdateSessionEntityTypeRequest) GetUpdateMask() *google_protobuf3.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The request message for
// [SessionEntityTypes.DeleteSessionEntityType][google.cloud.dialogflow.v2beta1.SessionEntityTypes.DeleteSessionEntityType].
type DeleteSessionEntityTypeRequest struct {
	// Required. The name of the entity type to delete. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteSessionEntityTypeRequest) Reset()                    { *m = DeleteSessionEntityTypeRequest{} }
func (m *DeleteSessionEntityTypeRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSessionEntityTypeRequest) ProtoMessage()               {}
func (*DeleteSessionEntityTypeRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *DeleteSessionEntityTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*SessionEntityType)(nil), "google.cloud.dialogflow.v2beta1.SessionEntityType")
	proto.RegisterType((*ListSessionEntityTypesRequest)(nil), "google.cloud.dialogflow.v2beta1.ListSessionEntityTypesRequest")
	proto.RegisterType((*ListSessionEntityTypesResponse)(nil), "google.cloud.dialogflow.v2beta1.ListSessionEntityTypesResponse")
	proto.RegisterType((*GetSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2beta1.GetSessionEntityTypeRequest")
	proto.RegisterType((*CreateSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2beta1.CreateSessionEntityTypeRequest")
	proto.RegisterType((*UpdateSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2beta1.UpdateSessionEntityTypeRequest")
	proto.RegisterType((*DeleteSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2beta1.DeleteSessionEntityTypeRequest")
	proto.RegisterEnum("google.cloud.dialogflow.v2beta1.SessionEntityType_EntityOverrideMode", SessionEntityType_EntityOverrideMode_name, SessionEntityType_EntityOverrideMode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SessionEntityTypes service

type SessionEntityTypesClient interface {
	// Returns the list of all session entity types in the specified session.
	ListSessionEntityTypes(ctx context.Context, in *ListSessionEntityTypesRequest, opts ...grpc.CallOption) (*ListSessionEntityTypesResponse, error)
	// Retrieves the specified session entity type.
	GetSessionEntityType(ctx context.Context, in *GetSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error)
	// Creates a session entity type.
	CreateSessionEntityType(ctx context.Context, in *CreateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error)
	// Updates the specified session entity type.
	UpdateSessionEntityType(ctx context.Context, in *UpdateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error)
	// Deletes the specified session entity type.
	DeleteSessionEntityType(ctx context.Context, in *DeleteSessionEntityTypeRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type sessionEntityTypesClient struct {
	cc *grpc.ClientConn
}

func NewSessionEntityTypesClient(cc *grpc.ClientConn) SessionEntityTypesClient {
	return &sessionEntityTypesClient{cc}
}

func (c *sessionEntityTypesClient) ListSessionEntityTypes(ctx context.Context, in *ListSessionEntityTypesRequest, opts ...grpc.CallOption) (*ListSessionEntityTypesResponse, error) {
	out := new(ListSessionEntityTypesResponse)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.SessionEntityTypes/ListSessionEntityTypes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) GetSessionEntityType(ctx context.Context, in *GetSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error) {
	out := new(SessionEntityType)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.SessionEntityTypes/GetSessionEntityType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) CreateSessionEntityType(ctx context.Context, in *CreateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error) {
	out := new(SessionEntityType)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.SessionEntityTypes/CreateSessionEntityType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) UpdateSessionEntityType(ctx context.Context, in *UpdateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error) {
	out := new(SessionEntityType)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.SessionEntityTypes/UpdateSessionEntityType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) DeleteSessionEntityType(ctx context.Context, in *DeleteSessionEntityTypeRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.SessionEntityTypes/DeleteSessionEntityType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SessionEntityTypes service

type SessionEntityTypesServer interface {
	// Returns the list of all session entity types in the specified session.
	ListSessionEntityTypes(context.Context, *ListSessionEntityTypesRequest) (*ListSessionEntityTypesResponse, error)
	// Retrieves the specified session entity type.
	GetSessionEntityType(context.Context, *GetSessionEntityTypeRequest) (*SessionEntityType, error)
	// Creates a session entity type.
	CreateSessionEntityType(context.Context, *CreateSessionEntityTypeRequest) (*SessionEntityType, error)
	// Updates the specified session entity type.
	UpdateSessionEntityType(context.Context, *UpdateSessionEntityTypeRequest) (*SessionEntityType, error)
	// Deletes the specified session entity type.
	DeleteSessionEntityType(context.Context, *DeleteSessionEntityTypeRequest) (*google_protobuf2.Empty, error)
}

func RegisterSessionEntityTypesServer(s *grpc.Server, srv SessionEntityTypesServer) {
	s.RegisterService(&_SessionEntityTypes_serviceDesc, srv)
}

func _SessionEntityTypes_ListSessionEntityTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSessionEntityTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).ListSessionEntityTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.SessionEntityTypes/ListSessionEntityTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).ListSessionEntityTypes(ctx, req.(*ListSessionEntityTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_GetSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).GetSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.SessionEntityTypes/GetSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).GetSessionEntityType(ctx, req.(*GetSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_CreateSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).CreateSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.SessionEntityTypes/CreateSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).CreateSessionEntityType(ctx, req.(*CreateSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_UpdateSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).UpdateSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.SessionEntityTypes/UpdateSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).UpdateSessionEntityType(ctx, req.(*UpdateSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_DeleteSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).DeleteSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.SessionEntityTypes/DeleteSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).DeleteSessionEntityType(ctx, req.(*DeleteSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionEntityTypes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.v2beta1.SessionEntityTypes",
	HandlerType: (*SessionEntityTypesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSessionEntityTypes",
			Handler:    _SessionEntityTypes_ListSessionEntityTypes_Handler,
		},
		{
			MethodName: "GetSessionEntityType",
			Handler:    _SessionEntityTypes_GetSessionEntityType_Handler,
		},
		{
			MethodName: "CreateSessionEntityType",
			Handler:    _SessionEntityTypes_CreateSessionEntityType_Handler,
		},
		{
			MethodName: "UpdateSessionEntityType",
			Handler:    _SessionEntityTypes_UpdateSessionEntityType_Handler,
		},
		{
			MethodName: "DeleteSessionEntityType",
			Handler:    _SessionEntityTypes_DeleteSessionEntityType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/v2beta1/session_entity_type.proto",
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2beta1/session_entity_type.proto", fileDescriptor5)
}

var fileDescriptor5 = []byte{
	// 804 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x4f, 0xdb, 0x48,
	0x14, 0xdf, 0x31, 0xbb, 0x08, 0x06, 0xed, 0x2e, 0xcc, 0xa2, 0x10, 0x85, 0x25, 0x64, 0xcd, 0x6a,
	0x85, 0x72, 0xb0, 0x45, 0x76, 0x2f, 0x2c, 0xfd, 0xa3, 0x42, 0x0c, 0x8a, 0x44, 0x42, 0xe4, 0x04,
	0xa4, 0xf6, 0x62, 0x39, 0xf8, 0x61, 0x19, 0x12, 0x8f, 0x9b, 0x99, 0x40, 0x43, 0xc5, 0x85, 0xaf,
	0xd0, 0x6b, 0x4f, 0x3d, 0xf6, 0xd0, 0x53, 0xfb, 0x01, 0x7a, 0xee, 0xb1, 0xd7, 0x4a, 0xbd, 0x54,
	0xfd, 0x06, 0x95, 0xaa, 0x9e, 0x2a, 0x8f, 0x1d, 0x82, 0xea, 0x3f, 0xa1, 0x51, 0x6f, 0x9e, 0x37,
	0xf3, 0x7b, 0xef, 0xf7, 0x7b, 0xf3, 0xde, 0xf3, 0xe0, 0x75, 0x9b, 0x52, 0xbb, 0x0d, 0xea, 0x61,
	0x9b, 0xf6, 0x2c, 0xd5, 0x72, 0xcc, 0x36, 0xb5, 0x8f, 0xda, 0xf4, 0x4c, 0x3d, 0x2d, 0xb5, 0x80,
	0x9b, 0x6b, 0x2a, 0x03, 0xc6, 0x1c, 0xea, 0x1a, 0xe0, 0x72, 0x87, 0xf7, 0x0d, 0xde, 0xf7, 0x40,
	0xf1, 0xba, 0x94, 0x53, 0xb2, 0x1c, 0x40, 0x15, 0x01, 0x55, 0x86, 0x50, 0x25, 0x84, 0xe6, 0xfe,
	0x0c, 0x7d, 0x9b, 0x9e, 0xa3, 0x9a, 0xae, 0x4b, 0xb9, 0xc9, 0x1d, 0xea, 0xb2, 0x00, 0x9e, 0x5b,
	0x1b, 0x15, 0x39, 0x12, 0x31, 0xb7, 0x18, 0x42, 0xc4, 0xaa, 0xd5, 0x3b, 0x52, 0xa1, 0xe3, 0xf1,
	0x7e, 0xb8, 0x59, 0xf8, 0x76, 0xf3, 0xc8, 0x81, 0xb6, 0x65, 0x74, 0x4c, 0x76, 0x12, 0x9c, 0x90,
	0x3f, 0x49, 0x78, 0xae, 0x11, 0xc8, 0xd1, 0x84, 0xef, 0x66, 0xdf, 0x03, 0x42, 0xf0, 0xcf, 0xae,
	0xd9, 0x81, 0x2c, 0x2a, 0xa0, 0xd5, 0x69, 0x5d, 0x7c, 0x93, 0x33, 0x3c, 0x1f, 0x46, 0xa7, 0xa7,
	0xd0, 0xed, 0x3a, 0x16, 0x18, 0x1d, 0x6a, 0x41, 0x56, 0x2a, 0xa0, 0xd5, 0xdf, 0x4a, 0x9a, 0x32,
	0x42, 0xb9, 0x12, 0x89, 0xa2, 0x04, 0x9f, 0x7b, 0xa1, 0xb7, 0x2a, 0xb5, 0x40, 0x27, 0x10, 0xb1,
	0x91, 0x1a, 0x9e, 0x12, 0x56, 0x07, 0x58, 0x76, 0xa2, 0x30, 0xb1, 0x3a, 0x53, 0x2a, 0x8d, 0x0c,
	0x16, 0x89, 0xa2, 0x5f, 0xf9, 0x90, 0x2f, 0x11, 0x26, 0xd1, 0xd0, 0xe4, 0x6f, 0x5c, 0xd0, 0x6a,
	0xcd, 0x4a, 0xf3, 0xbe, 0xb1, 0x77, 0xa0, 0xe9, 0x7a, 0xa5, 0xac, 0x19, 0xd5, 0xbd, 0xb2, 0x66,
	0xec, 0xd7, 0x1a, 0x75, 0x6d, 0xab, 0xb2, 0x5d, 0xd1, 0xca, 0xb3, 0x3f, 0x91, 0xbf, 0xf0, 0x52,
	0xec, 0xa9, 0xc1, 0x6a, 0x16, 0x91, 0x15, 0xbc, 0x1c, 0x7b, 0xa4, 0xb1, 0x5f, 0xaf, 0xef, 0x6a,
	0x55, 0xad, 0xd6, 0x9c, 0x95, 0x64, 0x86, 0x97, 0x76, 0x1d, 0xc6, 0x23, 0x49, 0x61, 0x3a, 0x3c,
	0xec, 0x01, 0xe3, 0x24, 0x83, 0x27, 0x3d, 0xb3, 0x0b, 0x2e, 0x0f, 0x2f, 0x21, 0x5c, 0x91, 0x45,
	0x3c, 0xed, 0x99, 0x36, 0x18, 0xcc, 0x39, 0x0f, 0x72, 0xff, 0x8b, 0x3e, 0xe5, 0x1b, 0x1a, 0xce,
	0x39, 0x90, 0x25, 0x8c, 0xc5, 0x26, 0xa7, 0x27, 0xe0, 0x66, 0x27, 0x04, 0x50, 0x1c, 0x6f, 0xfa,
	0x06, 0xf9, 0x05, 0xc2, 0xf9, 0xa4, 0xa8, 0xcc, 0xa3, 0x2e, 0x03, 0x62, 0xe1, 0xf9, 0x98, 0xea,
	0x66, 0x59, 0x74, 0xc3, 0xc4, 0x47, 0x5c, 0xeb, 0x84, 0x45, 0xa2, 0x91, 0x7f, 0xf0, 0xef, 0x2e,
	0x3c, 0xe2, 0xc6, 0x35, 0xb2, 0x92, 0x20, 0xfb, 0xab, 0x6f, 0xae, 0x5f, 0x11, 0x5e, 0xc3, 0x8b,
	0x3b, 0x10, 0xa5, 0x3b, 0xc8, 0x51, 0x4c, 0x99, 0xca, 0x4f, 0x11, 0xce, 0x6f, 0x75, 0xc1, 0xe4,
	0x90, 0x08, 0x4b, 0x4a, 0x6d, 0x0b, 0xff, 0x11, 0xa3, 0x5d, 0x30, 0x1b, 0x4f, 0xfa, 0x5c, 0x44,
	0xba, 0xfc, 0x1a, 0xe1, 0xfc, 0xbe, 0x67, 0xa5, 0xd1, 0x4b, 0xa0, 0x81, 0x7e, 0x20, 0x0d, 0xb2,
	0x81, 0x67, 0x7a, 0x82, 0x85, 0x98, 0x05, 0xa1, 0xc4, 0xdc, 0xc0, 0xf7, 0x60, 0x5c, 0x28, 0xdb,
	0xfe, 0xb8, 0xa8, 0x9a, 0xec, 0x44, 0xc7, 0xc1, 0x71, 0xff, 0x5b, 0xfe, 0x0f, 0xe7, 0xcb, 0xd0,
	0x86, 0x14, 0x09, 0x31, 0x17, 0x53, 0x7a, 0x3f, 0x85, 0x49, 0xb4, 0xf0, 0xc8, 0x3b, 0x84, 0x33,
	0xf1, 0x35, 0x49, 0xee, 0x8c, 0xd4, 0x9a, 0xda, 0x42, 0xb9, 0xbb, 0x63, 0xe3, 0x83, 0x66, 0x90,
	0xef, 0x5d, 0xbe, 0xfd, 0xf0, 0x44, 0xda, 0x20, 0xeb, 0x57, 0xf3, 0xf7, 0x71, 0x50, 0x2a, 0xb7,
	0xbd, 0x2e, 0x3d, 0x86, 0x43, 0xce, 0xd4, 0xa2, 0x6a, 0xda, 0xe0, 0xf2, 0xc1, 0x2f, 0x81, 0xa9,
	0xc5, 0x8b, 0x70, 0x48, 0x07, 0x1a, 0xde, 0x20, 0x3c, 0x1f, 0x57, 0xc2, 0xe4, 0xd6, 0x48, 0x72,
	0x29, 0x95, 0x9f, 0x1b, 0xa3, 0x0c, 0xe2, 0xd4, 0xf8, 0x17, 0x93, 0xa6, 0xe5, 0xba, 0x14, 0xb5,
	0x78, 0x41, 0x3e, 0x22, 0xbc, 0x90, 0xd0, 0x5c, 0x64, 0x74, 0xb6, 0xd3, 0xdb, 0x72, 0x2c, 0x4d,
	0x07, 0x42, 0x53, 0x5d, 0x1e, 0xff, 0x86, 0xfe, 0x8f, 0x6b, 0x36, 0xf2, 0x05, 0xe1, 0x85, 0x84,
	0x36, 0xbd, 0x81, 0xd0, 0xf4, 0x06, 0x1f, 0x4b, 0xe8, 0xb1, 0x10, 0x6a, 0x95, 0xaa, 0x43, 0xa1,
	0x71, 0xaf, 0x90, 0xef, 0xbc, 0xd0, 0x78, 0xf1, 0xaf, 0x10, 0x5e, 0x48, 0x68, 0xf0, 0x1b, 0x88,
	0x4f, 0x1f, 0x0d, 0xb9, 0x4c, 0x64, 0xc8, 0x68, 0xfe, 0x83, 0x65, 0x50, 0x9d, 0xc5, 0xf1, 0xab,
	0x73, 0xf3, 0x25, 0xc2, 0x2b, 0x87, 0xb4, 0x33, 0x8a, 0xe1, 0x66, 0x26, 0x42, 0xae, 0xee, 0x73,
	0xa9, 0xa3, 0x07, 0x95, 0x10, 0x6a, 0xd3, 0xb6, 0xe9, 0xda, 0x0a, 0xed, 0xda, 0xaa, 0x0d, 0xae,
	0x60, 0xaa, 0x06, 0x5b, 0xa6, 0xe7, 0xb0, 0xc4, 0xe7, 0xd9, 0xc6, 0xd0, 0xf4, 0x19, 0xa1, 0x67,
	0x92, 0x54, 0xde, 0x7e, 0x2e, 0x2d, 0xef, 0x04, 0x3e, 0xb7, 0x04, 0x9d, 0xf2, 0x90, 0xce, 0x41,
	0x00, 0x6a, 0x4d, 0x0a, 0xff, 0xff, 0x7e, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x87, 0xb0, 0xb6,
	0x77, 0x0a, 0x00, 0x00,
}
