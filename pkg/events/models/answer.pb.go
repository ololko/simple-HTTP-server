// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/events/models/answer.proto

package models

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Event struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Timestamp            int32    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e72c3bdb575315, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Request struct {
	From                 int32    `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   int32    `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e72c3bdb575315, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Request) GetTo() int32 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Answer struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e72c3bdb575315, []int{2}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Answer) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "models.Event")
	proto.RegisterType((*Request)(nil), "models.Request")
	proto.RegisterType((*Answer)(nil), "models.Answer")
}

func init() { proto.RegisterFile("pkg/events/models/answer.proto", fileDescriptor_f3e72c3bdb575315) }

var fileDescriptor_f3e72c3bdb575315 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0xa9, 0x49, 0xc9, 0xaa, 0x2d, 0x2c, 0x22, 0x25, 0x16, 0x29, 0x7b, 0x2a, 0x1e, 0x76,
	0xa1, 0xde, 0x04, 0x0f, 0x45, 0x0a, 0xde, 0x84, 0xfc, 0xc1, 0xd6, 0x4e, 0x43, 0xb0, 0xd9, 0x5d,
	0xb3, 0x13, 0xa5, 0x57, 0xff, 0x40, 0xfc, 0x34, 0x7f, 0xc1, 0x0f, 0x91, 0xce, 0x26, 0xf6, 0xea,
	0x6d, 0xe7, 0xcd, 0xdb, 0xf7, 0xe6, 0x3d, 0x76, 0xed, 0x5e, 0x4a, 0x05, 0x6f, 0x60, 0xd0, 0xab,
	0xda, 0x6e, 0x60, 0xe7, 0x95, 0x36, 0xfe, 0x1d, 0x1a, 0xe9, 0x1a, 0x8b, 0x96, 0xa7, 0x01, 0xcc,
	0xa7, 0xa5, 0xb5, 0xe5, 0x0e, 0x94, 0x76, 0x95, 0xd2, 0xc6, 0x58, 0xd4, 0x58, 0x59, 0xe3, 0x03,
	0x2b, 0xbf, 0xea, 0xb6, 0x34, 0xad, 0xdb, 0xad, 0x82, 0xda, 0xe1, 0x3e, 0x2c, 0xc5, 0x13, 0x4b,
	0x56, 0x07, 0x03, 0x7e, 0xc1, 0x92, 0x67, 0xdb, 0x1a, 0x9c, 0x44, 0xb3, 0x68, 0x9e, 0x14, 0x61,
	0xe0, 0x9c, 0x9d, 0xe0, 0xde, 0xc1, 0x24, 0x9e, 0x45, 0xf3, 0xac, 0xa0, 0x37, 0x9f, 0xb2, 0x0c,
	0xab, 0x1a, 0x3c, 0xea, 0xda, 0x4d, 0x06, 0xc4, 0x3e, 0x02, 0x62, 0xc9, 0x86, 0x05, 0xbc, 0xb6,
	0xe0, 0xe9, 0xf3, 0xb6, 0xb1, 0x75, 0xa7, 0x48, 0x6f, 0x3e, 0x62, 0x31, 0x5a, 0x92, 0x4b, 0x8a,
	0x18, 0xed, 0x9f, 0xc1, 0xe0, 0x68, 0x20, 0x16, 0x2c, 0x5d, 0x52, 0xcc, 0xff, 0x1f, 0xb5, 0xf8,
	0x8c, 0x58, 0x4a, 0x41, 0x3c, 0xbf, 0x67, 0x59, 0x01, 0x7a, 0x13, 0x62, 0x8d, 0x65, 0xe8, 0x48,
	0x76, 0x47, 0xe5, 0xa3, 0x1e, 0x08, 0x16, 0x62, 0xfc, 0xf1, 0xfd, 0xf3, 0x15, 0x67, 0x7c, 0xd8,
	0x15, 0xcd, 0x1f, 0xd9, 0xe9, 0x43, 0x03, 0x1a, 0x21, 0x08, 0x9c, 0xf7, 0x7c, 0x1a, 0xf3, 0x4b,
	0x19, 0xda, 0x94, 0x7d, 0x9b, 0x72, 0x75, 0x68, 0x53, 0x70, 0x92, 0x39, 0x13, 0xbd, 0xcc, 0x5d,
	0x74, 0xb3, 0x4e, 0x89, 0x73, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x45, 0x58, 0xfe, 0x8f, 0xc7,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	ReadEvent(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Answer, error)
	CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) ReadEvent(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/models.Events/ReadEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/models.Events/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	ReadEvent(context.Context, *Request) (*Answer, error)
	CreateEvent(context.Context, *Event) (*empty.Empty, error)
}

// UnimplementedEventsServer can be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (*UnimplementedEventsServer) ReadEvent(ctx context.Context, req *Request) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEvent not implemented")
}
func (*UnimplementedEventsServer) CreateEvent(ctx context.Context, req *Event) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_ReadEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).ReadEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Events/ReadEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).ReadEvent(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Events/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).CreateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadEvent",
			Handler:    _Events_ReadEvent_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Events_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/events/models/answer.proto",
}
