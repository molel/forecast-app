// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: predict.proto

package predict

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
	PredictService_MakePredict_FullMethodName = "/PredictService/MakePredict"
	PredictService_GetPredict_FullMethodName  = "/PredictService/GetPredict"
	PredictService_GetPredicts_FullMethodName = "/PredictService/GetPredicts"
)

// PredictServiceClient is the client API for PredictService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PredictServiceClient interface {
	MakePredict(ctx context.Context, in *MakePredictRequest, opts ...grpc.CallOption) (*Empty, error)
	GetPredict(ctx context.Context, in *GetPredictRequest, opts ...grpc.CallOption) (*GetPredictResponse, error)
	GetPredicts(ctx context.Context, in *GetPredictsRequest, opts ...grpc.CallOption) (*GetPredictsResponse, error)
}

type predictServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictServiceClient(cc grpc.ClientConnInterface) PredictServiceClient {
	return &predictServiceClient{cc}
}

func (c *predictServiceClient) MakePredict(ctx context.Context, in *MakePredictRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, PredictService_MakePredict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictServiceClient) GetPredict(ctx context.Context, in *GetPredictRequest, opts ...grpc.CallOption) (*GetPredictResponse, error) {
	out := new(GetPredictResponse)
	err := c.cc.Invoke(ctx, PredictService_GetPredict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictServiceClient) GetPredicts(ctx context.Context, in *GetPredictsRequest, opts ...grpc.CallOption) (*GetPredictsResponse, error) {
	out := new(GetPredictsResponse)
	err := c.cc.Invoke(ctx, PredictService_GetPredicts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictServiceServer is the server API for PredictService service.
// All implementations must embed UnimplementedPredictServiceServer
// for forward compatibility
type PredictServiceServer interface {
	MakePredict(context.Context, *MakePredictRequest) (*Empty, error)
	GetPredict(context.Context, *GetPredictRequest) (*GetPredictResponse, error)
	GetPredicts(context.Context, *GetPredictsRequest) (*GetPredictsResponse, error)
	mustEmbedUnimplementedPredictServiceServer()
}

// UnimplementedPredictServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPredictServiceServer struct {
}

func (UnimplementedPredictServiceServer) MakePredict(context.Context, *MakePredictRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePredict not implemented")
}
func (UnimplementedPredictServiceServer) GetPredict(context.Context, *GetPredictRequest) (*GetPredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredict not implemented")
}
func (UnimplementedPredictServiceServer) GetPredicts(context.Context, *GetPredictsRequest) (*GetPredictsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredicts not implemented")
}
func (UnimplementedPredictServiceServer) mustEmbedUnimplementedPredictServiceServer() {}

// UnsafePredictServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PredictServiceServer will
// result in compilation errors.
type UnsafePredictServiceServer interface {
	mustEmbedUnimplementedPredictServiceServer()
}

func RegisterPredictServiceServer(s grpc.ServiceRegistrar, srv PredictServiceServer) {
	s.RegisterService(&PredictService_ServiceDesc, srv)
}

func _PredictService_MakePredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakePredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictServiceServer).MakePredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PredictService_MakePredict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictServiceServer).MakePredict(ctx, req.(*MakePredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictService_GetPredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictServiceServer).GetPredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PredictService_GetPredict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictServiceServer).GetPredict(ctx, req.(*GetPredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictService_GetPredicts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPredictsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictServiceServer).GetPredicts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PredictService_GetPredicts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictServiceServer).GetPredicts(ctx, req.(*GetPredictsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PredictService_ServiceDesc is the grpc.ServiceDesc for PredictService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PredictService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PredictService",
	HandlerType: (*PredictServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakePredict",
			Handler:    _PredictService_MakePredict_Handler,
		},
		{
			MethodName: "GetPredict",
			Handler:    _PredictService_GetPredict_Handler,
		},
		{
			MethodName: "GetPredicts",
			Handler:    _PredictService_GetPredicts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "predict.proto",
}
