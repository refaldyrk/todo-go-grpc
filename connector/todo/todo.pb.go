// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/todo.proto

package todo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//Model
type Todo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsDone               bool     `protobuf:"varint,3,opt,name=is_done,json=isDone,proto3" json:"is_done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Todo) GetIsDone() bool {
	if m != nil {
		return m.IsDone
	}
	return false
}

//Request
type AddTodoRequest struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTodoRequest) Reset()         { *m = AddTodoRequest{} }
func (m *AddTodoRequest) String() string { return proto.CompactTextString(m) }
func (*AddTodoRequest) ProtoMessage()    {}
func (*AddTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{1}
}

func (m *AddTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTodoRequest.Unmarshal(m, b)
}
func (m *AddTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTodoRequest.Marshal(b, m, deterministic)
}
func (m *AddTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTodoRequest.Merge(m, src)
}
func (m *AddTodoRequest) XXX_Size() int {
	return xxx_messageInfo_AddTodoRequest.Size(m)
}
func (m *AddTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddTodoRequest proto.InternalMessageInfo

func (m *AddTodoRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type UpdateTodoRequest struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoRequest) Reset()         { *m = UpdateTodoRequest{} }
func (m *UpdateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoRequest) ProtoMessage()    {}
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{2}
}

func (m *UpdateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoRequest.Unmarshal(m, b)
}
func (m *UpdateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoRequest.Merge(m, src)
}
func (m *UpdateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoRequest.Size(m)
}
func (m *UpdateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoRequest proto.InternalMessageInfo

func (m *UpdateTodoRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type DeleteTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRequest) Reset()         { *m = DeleteTodoRequest{} }
func (m *DeleteTodoRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRequest) ProtoMessage()    {}
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{3}
}

func (m *DeleteTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRequest.Unmarshal(m, b)
}
func (m *DeleteTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRequest.Merge(m, src)
}
func (m *DeleteTodoRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRequest.Size(m)
}
func (m *DeleteTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRequest proto.InternalMessageInfo

func (m *DeleteTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//Response
type TodoResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *Todo    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Error                bool     `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoResponse) Reset()         { *m = TodoResponse{} }
func (m *TodoResponse) String() string { return proto.CompactTextString(m) }
func (*TodoResponse) ProtoMessage()    {}
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{4}
}

func (m *TodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoResponse.Unmarshal(m, b)
}
func (m *TodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoResponse.Marshal(b, m, deterministic)
}
func (m *TodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoResponse.Merge(m, src)
}
func (m *TodoResponse) XXX_Size() int {
	return xxx_messageInfo_TodoResponse.Size(m)
}
func (m *TodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TodoResponse proto.InternalMessageInfo

func (m *TodoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TodoResponse) GetData() *Todo {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TodoResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func init() {
	proto.RegisterType((*Todo)(nil), "connector.Todo")
	proto.RegisterType((*AddTodoRequest)(nil), "connector.AddTodoRequest")
	proto.RegisterType((*UpdateTodoRequest)(nil), "connector.UpdateTodoRequest")
	proto.RegisterType((*DeleteTodoRequest)(nil), "connector.DeleteTodoRequest")
	proto.RegisterType((*TodoResponse)(nil), "connector.TodoResponse")
}

func init() {
	proto.RegisterFile("proto/todo.proto", fileDescriptor_d28b1ccc7160a78f)
}

var fileDescriptor_d28b1ccc7160a78f = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x25, 0xf9, 0xe5, 0xd7, 0xd8, 0xa9, 0xc4, 0x76, 0x10, 0x1a, 0xc5, 0x43, 0x49, 0x2f, 0x3d,
	0x45, 0xa8, 0x08, 0x5e, 0x3c, 0x68, 0xfb, 0x09, 0xa2, 0x5e, 0xbc, 0x48, 0xcc, 0x0e, 0xb2, 0x60,
	0x77, 0xe2, 0xee, 0xea, 0x47, 0xf5, 0xf3, 0xc8, 0x6e, 0xd3, 0x3f, 0x69, 0x55, 0xf0, 0x36, 0xb3,
	0xef, 0xcd, 0x7b, 0x3b, 0x8f, 0x81, 0x7e, 0xad, 0xd9, 0xf2, 0xb9, 0x65, 0xc1, 0xb9, 0x2f, 0xb1,
	0x5b, 0xb1, 0x52, 0x54, 0x59, 0xd6, 0xd9, 0x0c, 0xa2, 0x7b, 0x16, 0x8c, 0x09, 0x84, 0x52, 0xa4,
	0xc1, 0x28, 0x98, 0x74, 0x8b, 0x50, 0x0a, 0x44, 0x88, 0x54, 0xb9, 0xa0, 0x34, 0xf4, 0x2f, 0xbe,
	0xc6, 0x21, 0xc4, 0xd2, 0x3c, 0x09, 0x56, 0x94, 0xfe, 0x1b, 0x05, 0x93, 0x83, 0xa2, 0x23, 0xcd,
	0x9c, 0x15, 0x65, 0x97, 0x90, 0xdc, 0x08, 0xe1, 0x74, 0x0a, 0x7a, 0x7b, 0x27, 0x63, 0x71, 0x0c,
	0x91, 0xf3, 0xf3, 0x82, 0xbd, 0xe9, 0x51, 0xbe, 0x36, 0xcc, 0x3d, 0xcb, 0x83, 0xd9, 0x15, 0x0c,
	0x1e, 0x6a, 0x51, 0x5a, 0xfa, 0xf3, 0xe4, 0x18, 0x06, 0x73, 0x7a, 0xa5, 0xf6, 0xe4, 0xce, 0x0a,
	0x59, 0x05, 0x87, 0x4b, 0xd8, 0xd4, 0xac, 0x0c, 0x61, 0x0a, 0xf1, 0x82, 0x8c, 0x29, 0x5f, 0xa8,
	0x21, 0xad, 0x5a, 0xe7, 0x29, 0x4a, 0x5b, 0xfa, 0x65, 0xbf, 0xf3, 0x74, 0x20, 0x1e, 0xc3, 0x7f,
	0xd2, 0x9a, 0x75, 0xb3, 0xfb, 0xb2, 0x99, 0x7e, 0x06, 0xd0, 0x73, 0xa4, 0x3b, 0xd2, 0x1f, 0xb2,
	0x22, 0xbc, 0x86, 0xb8, 0x89, 0x02, 0x4f, 0xb6, 0x74, 0xda, 0xf1, 0x9c, 0x0e, 0x77, 0x2d, 0x56,
	0x7f, 0x9c, 0x01, 0x6c, 0x22, 0xc1, 0xb3, 0x2d, 0xda, 0x5e, 0x52, 0xbf, 0x8a, 0x6c, 0xd2, 0x69,
	0x89, 0xec, 0x85, 0xf6, 0xa3, 0xc8, 0x6d, 0xff, 0x31, 0x59, 0x23, 0xfe, 0x76, 0x9e, 0x3b, 0xfe,
	0x78, 0x2e, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x24, 0x06, 0x37, 0x44, 0x50, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	AddTodo(ctx context.Context, in *AddTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) AddTodo(ctx context.Context, in *AddTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, "/connector.TodoService/AddTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, "/connector.TodoService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, "/connector.TodoService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	AddTodo(context.Context, *AddTodoRequest) (*TodoResponse, error)
	UpdateTodo(context.Context, *UpdateTodoRequest) (*TodoResponse, error)
	DeleteTodo(context.Context, *DeleteTodoRequest) (*TodoResponse, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) AddTodo(ctx context.Context, req *AddTodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (*UnimplementedTodoServiceServer) UpdateTodo(ctx context.Context, req *UpdateTodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (*UnimplementedTodoServiceServer) DeleteTodo(ctx context.Context, req *DeleteTodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connector.TodoService/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).AddTodo(ctx, req.(*AddTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connector.TodoService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*UpdateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connector.TodoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connector.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTodo",
			Handler:    _TodoService_AddTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/todo.proto",
}
