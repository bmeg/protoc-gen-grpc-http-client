// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: simple.proto

package simple

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GraphInterfaceClient is the client API for GraphInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GraphInterfaceClient interface {
	AddVertex(ctx context.Context, in *GraphElement, opts ...grpc.CallOption) (*EditResult, error)
	GetVertex(ctx context.Context, in *ElementID, opts ...grpc.CallOption) (*Vertex, error)
	BulkAdd(ctx context.Context, opts ...grpc.CallOption) (GraphInterface_BulkAddClient, error)
	AddGraph(ctx context.Context, in *GraphID, opts ...grpc.CallOption) (*EditResult, error)
	DeleteGraph(ctx context.Context, in *GraphID, opts ...grpc.CallOption) (*EditResult, error)
	DeleteVertex(ctx context.Context, in *ElementID, opts ...grpc.CallOption) (*EditResult, error)
	DeleteIndex(ctx context.Context, in *IndexID, opts ...grpc.CallOption) (*EditResult, error)
	AddSchema(ctx context.Context, in *Graph, opts ...grpc.CallOption) (*EditResult, error)
}

type graphInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewGraphInterfaceClient(cc grpc.ClientConnInterface) GraphInterfaceClient {
	return &graphInterfaceClient{cc}
}

func (c *graphInterfaceClient) AddVertex(ctx context.Context, in *GraphElement, opts ...grpc.CallOption) (*EditResult, error) {
	out := new(EditResult)
	err := c.cc.Invoke(ctx, "/simple.GraphInterface/AddVertex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphInterfaceClient) GetVertex(ctx context.Context, in *ElementID, opts ...grpc.CallOption) (*Vertex, error) {
	out := new(Vertex)
	err := c.cc.Invoke(ctx, "/simple.GraphInterface/GetVertex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphInterfaceClient) BulkAdd(ctx context.Context, opts ...grpc.CallOption) (GraphInterface_BulkAddClient, error) {
	stream, err := c.cc.NewStream(ctx, &GraphInterface_ServiceDesc.Streams[0], "/simple.GraphInterface/BulkAdd", opts...)
	if err != nil {
		return nil, err
	}
	x := &graphInterfaceBulkAddClient{stream}
	return x, nil
}

type GraphInterface_BulkAddClient interface {
	Send(*GraphElement) error
	CloseAndRecv() (*BulkEditResult, error)
	grpc.ClientStream
}

type graphInterfaceBulkAddClient struct {
	grpc.ClientStream
}

func (x *graphInterfaceBulkAddClient) Send(m *GraphElement) error {
	return x.ClientStream.SendMsg(m)
}

func (x *graphInterfaceBulkAddClient) CloseAndRecv() (*BulkEditResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BulkEditResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *graphInterfaceClient) AddGraph(ctx context.Context, in *GraphID, opts ...grpc.CallOption) (*EditResult, error) {
	out := new(EditResult)
	err := c.cc.Invoke(ctx, "/simple.GraphInterface/AddGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphInterfaceClient) DeleteGraph(ctx context.Context, in *GraphID, opts ...grpc.CallOption) (*EditResult, error) {
	out := new(EditResult)
	err := c.cc.Invoke(ctx, "/simple.GraphInterface/DeleteGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphInterfaceClient) DeleteVertex(ctx context.Context, in *ElementID, opts ...grpc.CallOption) (*EditResult, error) {
	out := new(EditResult)
	err := c.cc.Invoke(ctx, "/simple.GraphInterface/DeleteVertex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphInterfaceClient) DeleteIndex(ctx context.Context, in *IndexID, opts ...grpc.CallOption) (*EditResult, error) {
	out := new(EditResult)
	err := c.cc.Invoke(ctx, "/simple.GraphInterface/DeleteIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphInterfaceClient) AddSchema(ctx context.Context, in *Graph, opts ...grpc.CallOption) (*EditResult, error) {
	out := new(EditResult)
	err := c.cc.Invoke(ctx, "/simple.GraphInterface/AddSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GraphInterfaceServer is the server API for GraphInterface service.
// All implementations must embed UnimplementedGraphInterfaceServer
// for forward compatibility
type GraphInterfaceServer interface {
	AddVertex(context.Context, *GraphElement) (*EditResult, error)
	GetVertex(context.Context, *ElementID) (*Vertex, error)
	BulkAdd(GraphInterface_BulkAddServer) error
	AddGraph(context.Context, *GraphID) (*EditResult, error)
	DeleteGraph(context.Context, *GraphID) (*EditResult, error)
	DeleteVertex(context.Context, *ElementID) (*EditResult, error)
	DeleteIndex(context.Context, *IndexID) (*EditResult, error)
	AddSchema(context.Context, *Graph) (*EditResult, error)
	mustEmbedUnimplementedGraphInterfaceServer()
}

// UnimplementedGraphInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedGraphInterfaceServer struct {
}

func (UnimplementedGraphInterfaceServer) AddVertex(context.Context, *GraphElement) (*EditResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVertex not implemented")
}
func (UnimplementedGraphInterfaceServer) GetVertex(context.Context, *ElementID) (*Vertex, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVertex not implemented")
}
func (UnimplementedGraphInterfaceServer) BulkAdd(GraphInterface_BulkAddServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkAdd not implemented")
}
func (UnimplementedGraphInterfaceServer) AddGraph(context.Context, *GraphID) (*EditResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGraph not implemented")
}
func (UnimplementedGraphInterfaceServer) DeleteGraph(context.Context, *GraphID) (*EditResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGraph not implemented")
}
func (UnimplementedGraphInterfaceServer) DeleteVertex(context.Context, *ElementID) (*EditResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVertex not implemented")
}
func (UnimplementedGraphInterfaceServer) DeleteIndex(context.Context, *IndexID) (*EditResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndex not implemented")
}
func (UnimplementedGraphInterfaceServer) AddSchema(context.Context, *Graph) (*EditResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSchema not implemented")
}
func (UnimplementedGraphInterfaceServer) mustEmbedUnimplementedGraphInterfaceServer() {}

// UnsafeGraphInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GraphInterfaceServer will
// result in compilation errors.
type UnsafeGraphInterfaceServer interface {
	mustEmbedUnimplementedGraphInterfaceServer()
}

func RegisterGraphInterfaceServer(s grpc.ServiceRegistrar, srv GraphInterfaceServer) {
	s.RegisterService(&GraphInterface_ServiceDesc, srv)
}

func _GraphInterface_AddVertex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphElement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphInterfaceServer).AddVertex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.GraphInterface/AddVertex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphInterfaceServer).AddVertex(ctx, req.(*GraphElement))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraphInterface_GetVertex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElementID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphInterfaceServer).GetVertex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.GraphInterface/GetVertex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphInterfaceServer).GetVertex(ctx, req.(*ElementID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraphInterface_BulkAdd_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GraphInterfaceServer).BulkAdd(&graphInterfaceBulkAddServer{stream})
}

type GraphInterface_BulkAddServer interface {
	SendAndClose(*BulkEditResult) error
	Recv() (*GraphElement, error)
	grpc.ServerStream
}

type graphInterfaceBulkAddServer struct {
	grpc.ServerStream
}

func (x *graphInterfaceBulkAddServer) SendAndClose(m *BulkEditResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *graphInterfaceBulkAddServer) Recv() (*GraphElement, error) {
	m := new(GraphElement)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GraphInterface_AddGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphInterfaceServer).AddGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.GraphInterface/AddGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphInterfaceServer).AddGraph(ctx, req.(*GraphID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraphInterface_DeleteGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphInterfaceServer).DeleteGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.GraphInterface/DeleteGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphInterfaceServer).DeleteGraph(ctx, req.(*GraphID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraphInterface_DeleteVertex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElementID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphInterfaceServer).DeleteVertex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.GraphInterface/DeleteVertex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphInterfaceServer).DeleteVertex(ctx, req.(*ElementID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraphInterface_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphInterfaceServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.GraphInterface/DeleteIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphInterfaceServer).DeleteIndex(ctx, req.(*IndexID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraphInterface_AddSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Graph)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphInterfaceServer).AddSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.GraphInterface/AddSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphInterfaceServer).AddSchema(ctx, req.(*Graph))
	}
	return interceptor(ctx, in, info, handler)
}

// GraphInterface_ServiceDesc is the grpc.ServiceDesc for GraphInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GraphInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simple.GraphInterface",
	HandlerType: (*GraphInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddVertex",
			Handler:    _GraphInterface_AddVertex_Handler,
		},
		{
			MethodName: "GetVertex",
			Handler:    _GraphInterface_GetVertex_Handler,
		},
		{
			MethodName: "AddGraph",
			Handler:    _GraphInterface_AddGraph_Handler,
		},
		{
			MethodName: "DeleteGraph",
			Handler:    _GraphInterface_DeleteGraph_Handler,
		},
		{
			MethodName: "DeleteVertex",
			Handler:    _GraphInterface_DeleteVertex_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _GraphInterface_DeleteIndex_Handler,
		},
		{
			MethodName: "AddSchema",
			Handler:    _GraphInterface_AddSchema_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkAdd",
			Handler:       _GraphInterface_BulkAdd_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "simple.proto",
}
