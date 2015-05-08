// Code generated by protoc-gen-go.
// source: msg.proto
// DO NOT EDIT!

/*
Package streams is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	Response
	Request
	Message
*/
package streams

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Response struct {
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

type Request struct {
	Origin string `protobuf:"bytes,1,opt,name=origin" json:"origin,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

type Message struct {
	Msg     string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

func init() {
}

// Client API for Pipeline service

type PipelineClient interface {
	Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *Request, opts ...grpc.CallOption) (Pipeline_ListClient, error)
}

type pipelineClient struct {
	cc *grpc.ClientConn
}

func NewPipelineClient(cc *grpc.ClientConn) PipelineClient {
	return &pipelineClient{cc}
}

func (c *pipelineClient) Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/streams.Pipeline/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineClient) List(ctx context.Context, in *Request, opts ...grpc.CallOption) (Pipeline_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Pipeline_serviceDesc.Streams[0], c.cc, "/streams.Pipeline/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipelineListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pipeline_ListClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type pipelineListClient struct {
	grpc.ClientStream
}

func (x *pipelineListClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Pipeline service

type PipelineServer interface {
	Send(context.Context, *Message) (*Response, error)
	List(*Request, Pipeline_ListServer) error
}

func RegisterPipelineServer(s *grpc.Server, srv PipelineServer) {
	s.RegisterService(&_Pipeline_serviceDesc, srv)
}

func _Pipeline_Send_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Message)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineServer).Send(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Pipeline_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PipelineServer).List(m, &pipelineListServer{stream})
}

type Pipeline_ListServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type pipelineListServer struct {
	grpc.ServerStream
}

func (x *pipelineListServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _Pipeline_serviceDesc = grpc.ServiceDesc{
	ServiceName: "streams.Pipeline",
	HandlerType: (*PipelineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Pipeline_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Pipeline_List_Handler,
			ServerStreams: true,
		},
	},
}