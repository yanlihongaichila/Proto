// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: goods.proto

package goods

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

const (
	Goods_GetGoods_FullMethodName     = "/goods.Goods/GetGoods"
	Goods_GetGoodses_FullMethodName   = "/goods.Goods/GetGoodses"
	Goods_CreatedGoods_FullMethodName = "/goods.Goods/CreatedGoods"
	Goods_UpdatedGoods_FullMethodName = "/goods.Goods/UpdatedGoods"
	Goods_DeletedGoods_FullMethodName = "/goods.Goods/DeletedGoods"
)

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsClient interface {
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error)
	GetGoodses(ctx context.Context, in *GetGoodsesRequest, opts ...grpc.CallOption) (*GetGoodsesResponse, error)
	CreatedGoods(ctx context.Context, in *CreatedGoodsRequest, opts ...grpc.CallOption) (*CreatedGoodsResponse, error)
	UpdatedGoods(ctx context.Context, in *UpdatedGoodsRequest, opts ...grpc.CallOption) (*UpdatedGoodsResponse, error)
	DeletedGoods(ctx context.Context, in *DeletedGoodsRequest, opts ...grpc.CallOption) (*DeletedGoodsResponse, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error) {
	out := new(GetGoodsResponse)
	err := c.cc.Invoke(ctx, Goods_GetGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoodses(ctx context.Context, in *GetGoodsesRequest, opts ...grpc.CallOption) (*GetGoodsesResponse, error) {
	out := new(GetGoodsesResponse)
	err := c.cc.Invoke(ctx, Goods_GetGoodses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreatedGoods(ctx context.Context, in *CreatedGoodsRequest, opts ...grpc.CallOption) (*CreatedGoodsResponse, error) {
	out := new(CreatedGoodsResponse)
	err := c.cc.Invoke(ctx, Goods_CreatedGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdatedGoods(ctx context.Context, in *UpdatedGoodsRequest, opts ...grpc.CallOption) (*UpdatedGoodsResponse, error) {
	out := new(UpdatedGoodsResponse)
	err := c.cc.Invoke(ctx, Goods_UpdatedGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeletedGoods(ctx context.Context, in *DeletedGoodsRequest, opts ...grpc.CallOption) (*DeletedGoodsResponse, error) {
	out := new(DeletedGoodsResponse)
	err := c.cc.Invoke(ctx, Goods_DeletedGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility
type GoodsServer interface {
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error)
	GetGoodses(context.Context, *GetGoodsesRequest) (*GetGoodsesResponse, error)
	CreatedGoods(context.Context, *CreatedGoodsRequest) (*CreatedGoodsResponse, error)
	UpdatedGoods(context.Context, *UpdatedGoodsRequest) (*UpdatedGoodsResponse, error)
	DeletedGoods(context.Context, *DeletedGoodsRequest) (*DeletedGoodsResponse, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (UnimplementedGoodsServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedGoodsServer) GetGoodses(context.Context, *GetGoodsesRequest) (*GetGoodsesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodses not implemented")
}
func (UnimplementedGoodsServer) CreatedGoods(context.Context, *CreatedGoodsRequest) (*CreatedGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedGoods not implemented")
}
func (UnimplementedGoodsServer) UpdatedGoods(context.Context, *UpdatedGoodsRequest) (*UpdatedGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatedGoods not implemented")
}
func (UnimplementedGoodsServer) DeletedGoods(context.Context, *DeletedGoodsRequest) (*DeletedGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletedGoods not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoodses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetGoodses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodses(ctx, req.(*GetGoodsesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreatedGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreatedGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_CreatedGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreatedGoods(ctx, req.(*CreatedGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdatedGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatedGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdatedGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_UpdatedGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdatedGoods(ctx, req.(*UpdatedGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeletedGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletedGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeletedGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_DeletedGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeletedGoods(ctx, req.(*DeletedGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoods",
			Handler:    _Goods_GetGoods_Handler,
		},
		{
			MethodName: "GetGoodses",
			Handler:    _Goods_GetGoodses_Handler,
		},
		{
			MethodName: "CreatedGoods",
			Handler:    _Goods_CreatedGoods_Handler,
		},
		{
			MethodName: "UpdatedGoods",
			Handler:    _Goods_UpdatedGoods_Handler,
		},
		{
			MethodName: "DeletedGoods",
			Handler:    _Goods_DeletedGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods.proto",
}
