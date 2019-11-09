// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Action int32

const (
	Action_NONE   Action = 0
	Action_ADD    Action = 1
	Action_REMOVE Action = 2
)

var Action_name = map[int32]string{
	0: "NONE",
	1: "ADD",
	2: "REMOVE",
}

var Action_value = map[string]int32{
	"NONE":   0,
	"ADD":    1,
	"REMOVE": 2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{0}
}

type Command struct {
	ProfileID            string   `protobuf:"bytes,1,opt,name=profileID,proto3" json:"profileID,omitempty"`
	SegmentIDs           []string `protobuf:"bytes,2,rep,name=segmentIDs,proto3" json:"segmentIDs,omitempty"`
	Action               Action   `protobuf:"varint,3,opt,name=action,proto3,enum=audience.v1.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetProfileID() string {
	if m != nil {
		return m.ProfileID
	}
	return ""
}

func (m *Command) GetSegmentIDs() []string {
	if m != nil {
		return m.SegmentIDs
	}
	return nil
}

func (m *Command) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_NONE
}

type BulkCommand struct {
	Commands             []*Command `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BulkCommand) Reset()         { *m = BulkCommand{} }
func (m *BulkCommand) String() string { return proto.CompactTextString(m) }
func (*BulkCommand) ProtoMessage()    {}
func (*BulkCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{1}
}

func (m *BulkCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkCommand.Unmarshal(m, b)
}
func (m *BulkCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkCommand.Marshal(b, m, deterministic)
}
func (m *BulkCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkCommand.Merge(m, src)
}
func (m *BulkCommand) XXX_Size() int {
	return xxx_messageInfo_BulkCommand.Size(m)
}
func (m *BulkCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkCommand.DiscardUnknown(m)
}

var xxx_messageInfo_BulkCommand proto.InternalMessageInfo

func (m *BulkCommand) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

type Confirmation struct {
	TxID                 string   `protobuf:"bytes,1,opt,name=txID,proto3" json:"txID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Confirmation) Reset()         { *m = Confirmation{} }
func (m *Confirmation) String() string { return proto.CompactTextString(m) }
func (*Confirmation) ProtoMessage()    {}
func (*Confirmation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{2}
}

func (m *Confirmation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Confirmation.Unmarshal(m, b)
}
func (m *Confirmation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Confirmation.Marshal(b, m, deterministic)
}
func (m *Confirmation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Confirmation.Merge(m, src)
}
func (m *Confirmation) XXX_Size() int {
	return xxx_messageInfo_Confirmation.Size(m)
}
func (m *Confirmation) XXX_DiscardUnknown() {
	xxx_messageInfo_Confirmation.DiscardUnknown(m)
}

var xxx_messageInfo_Confirmation proto.InternalMessageInfo

func (m *Confirmation) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func init() {
	proto.RegisterEnum("audience.v1.Action", Action_name, Action_value)
	proto.RegisterType((*Command)(nil), "audience.v1.Command")
	proto.RegisterType((*BulkCommand)(nil), "audience.v1.BulkCommand")
	proto.RegisterType((*Confirmation)(nil), "audience.v1.Confirmation")
}

func init() { proto.RegisterFile("commands.proto", fileDescriptor_0dff099eb2e3dfdb) }

var fileDescriptor_0dff099eb2e3dfdb = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x4f, 0xc2, 0x30,
	0x14, 0xb7, 0x8c, 0x0c, 0x78, 0x33, 0x84, 0x3c, 0x3d, 0x4c, 0x63, 0xcc, 0xb2, 0x8b, 0x8b, 0xc6,
	0x4d, 0xf1, 0xe0, 0x91, 0x00, 0xe3, 0xc0, 0x41, 0x20, 0x33, 0xf1, 0xe0, 0x6d, 0x8c, 0x32, 0x16,
	0x68, 0xbb, 0xac, 0x9d, 0xf1, 0x23, 0xfa, 0xb1, 0xcc, 0xc6, 0x7f, 0x63, 0x4c, 0x3c, 0xb5, 0xf9,
	0xbd, 0xdf, 0xbf, 0xbe, 0x14, 0x9a, 0x91, 0x60, 0x2c, 0xe4, 0x33, 0xe9, 0xa6, 0x99, 0x50, 0x02,
	0x8d, 0x30, 0x9f, 0x25, 0x94, 0x47, 0xd4, 0xfd, 0x78, 0xb4, 0x15, 0xd4, 0xfa, 0xeb, 0x31, 0x5e,
	0x41, 0x23, 0xcd, 0xc4, 0x3c, 0x59, 0xd1, 0xa1, 0x6f, 0x12, 0x8b, 0x38, 0x8d, 0x60, 0x0f, 0xe0,
	0x35, 0x80, 0xa4, 0x31, 0xa3, 0x5c, 0x0d, 0x7d, 0x69, 0x56, 0x2c, 0xcd, 0x69, 0x04, 0x07, 0x08,
	0xde, 0x81, 0x1e, 0x46, 0x2a, 0x11, 0xdc, 0xd4, 0x2c, 0xe2, 0x34, 0xdb, 0x67, 0xee, 0x41, 0x8c,
	0xdb, 0x2d, 0x47, 0xc1, 0x86, 0x62, 0x77, 0xc0, 0xe8, 0xe5, 0xab, 0xe5, 0x36, 0xf9, 0x01, 0xea,
	0xdb, 0x8e, 0x26, 0xb1, 0x34, 0xc7, 0x68, 0x9f, 0x1f, 0xa9, 0x37, 0xbc, 0x60, 0xc7, 0xb2, 0x6d,
	0x38, 0xed, 0x0b, 0x3e, 0x4f, 0x32, 0x16, 0x16, 0x86, 0x88, 0x50, 0x55, 0x9f, 0xbb, 0xda, 0xe5,
	0xfd, 0xf6, 0x06, 0xf4, 0x75, 0x2c, 0xd6, 0xa1, 0x3a, 0x1a, 0x8f, 0x06, 0xad, 0x13, 0xac, 0x81,
	0xd6, 0xf5, 0xfd, 0x16, 0x41, 0x00, 0x3d, 0x18, 0xbc, 0x8c, 0xdf, 0x06, 0xad, 0x4a, 0xfb, 0x8b,
	0x80, 0xb1, 0x89, 0x90, 0xdd, 0xc9, 0x10, 0x3b, 0x00, 0x93, 0x5c, 0x2e, 0x5e, 0x55, 0x46, 0x43,
	0x86, 0xbf, 0x56, 0xb9, 0xbc, 0xf8, 0x81, 0xee, 0xbb, 0x38, 0x04, 0x9f, 0xa1, 0x5a, 0x18, 0xfc,
	0x5b, 0x8a, 0x1d, 0xa8, 0x17, 0xc2, 0x62, 0x37, 0x68, 0x1e, 0xd1, 0x0e, 0xd6, 0xf5, 0x87, 0x41,
	0xcf, 0x7b, 0xbf, 0x8f, 0x13, 0xb5, 0xc8, 0xa7, 0x6e, 0x24, 0x98, 0xb7, 0x5c, 0x85, 0x0b, 0x29,
	0x29, 0xf7, 0x62, 0x91, 0xc6, 0x5e, 0x9c, 0xa5, 0x91, 0x27, 0xcb, 0x47, 0x25, 0x3c, 0xf6, 0xca,
	0xcf, 0x30, 0xd5, 0xcb, 0xe3, 0xe9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x97, 0x45, 0x43, 0x25,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommandsAPIClient is the client API for CommandsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandsAPIClient interface {
	PushStream(ctx context.Context, opts ...grpc.CallOption) (CommandsAPI_PushStreamClient, error)
	Push(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Confirmation, error)
	PushBulk(ctx context.Context, in *BulkCommand, opts ...grpc.CallOption) (*Confirmation, error)
}

type commandsAPIClient struct {
	cc *grpc.ClientConn
}

func NewCommandsAPIClient(cc *grpc.ClientConn) CommandsAPIClient {
	return &commandsAPIClient{cc}
}

func (c *commandsAPIClient) PushStream(ctx context.Context, opts ...grpc.CallOption) (CommandsAPI_PushStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CommandsAPI_serviceDesc.Streams[0], "/audience.v1.CommandsAPI/PushStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandsAPIPushStreamClient{stream}
	return x, nil
}

type CommandsAPI_PushStreamClient interface {
	Send(*Command) error
	CloseAndRecv() (*Confirmation, error)
	grpc.ClientStream
}

type commandsAPIPushStreamClient struct {
	grpc.ClientStream
}

func (x *commandsAPIPushStreamClient) Send(m *Command) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commandsAPIPushStreamClient) CloseAndRecv() (*Confirmation, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Confirmation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *commandsAPIClient) Push(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/audience.v1.CommandsAPI/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsAPIClient) PushBulk(ctx context.Context, in *BulkCommand, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/audience.v1.CommandsAPI/PushBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandsAPIServer is the server API for CommandsAPI service.
type CommandsAPIServer interface {
	PushStream(CommandsAPI_PushStreamServer) error
	Push(context.Context, *Command) (*Confirmation, error)
	PushBulk(context.Context, *BulkCommand) (*Confirmation, error)
}

func RegisterCommandsAPIServer(s *grpc.Server, srv CommandsAPIServer) {
	s.RegisterService(&_CommandsAPI_serviceDesc, srv)
}

func _CommandsAPI_PushStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommandsAPIServer).PushStream(&commandsAPIPushStreamServer{stream})
}

type CommandsAPI_PushStreamServer interface {
	SendAndClose(*Confirmation) error
	Recv() (*Command, error)
	grpc.ServerStream
}

type commandsAPIPushStreamServer struct {
	grpc.ServerStream
}

func (x *commandsAPIPushStreamServer) SendAndClose(m *Confirmation) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commandsAPIPushStreamServer) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CommandsAPI_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsAPIServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audience.v1.CommandsAPI/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsAPIServer).Push(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandsAPI_PushBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsAPIServer).PushBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audience.v1.CommandsAPI/PushBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsAPIServer).PushBulk(ctx, req.(*BulkCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommandsAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "audience.v1.CommandsAPI",
	HandlerType: (*CommandsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _CommandsAPI_Push_Handler,
		},
		{
			MethodName: "PushBulk",
			Handler:    _CommandsAPI_PushBulk_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PushStream",
			Handler:       _CommandsAPI_PushStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "commands.proto",
}