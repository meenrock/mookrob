// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: user_input.proto

package servicecalculator

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
	Calculator_BMI_FullMethodName = "/protocalculator.Calculator/BMI"
	Calculator_BMR_FullMethodName = "/protocalculator.Calculator/BMR"
)

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	BMI(ctx context.Context, in *GetUserBMIRequest, opts ...grpc.CallOption) (*BMIResponse, error)
	BMR(ctx context.Context, in *GetUserBMRRequest, opts ...grpc.CallOption) (*BMRResponse, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) BMI(ctx context.Context, in *GetUserBMIRequest, opts ...grpc.CallOption) (*BMIResponse, error) {
	out := new(BMIResponse)
	err := c.cc.Invoke(ctx, Calculator_BMI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) BMR(ctx context.Context, in *GetUserBMRRequest, opts ...grpc.CallOption) (*BMRResponse, error) {
	out := new(BMRResponse)
	err := c.cc.Invoke(ctx, Calculator_BMR_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
// All implementations must embed UnimplementedCalculatorServer
// for forward compatibility
type CalculatorServer interface {
	BMI(context.Context, *GetUserBMIRequest) (*BMIResponse, error)
	BMR(context.Context, *GetUserBMRRequest) (*BMRResponse, error)
	mustEmbedUnimplementedCalculatorServer()
}

// UnimplementedCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (UnimplementedCalculatorServer) BMI(context.Context, *GetUserBMIRequest) (*BMIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BMI not implemented")
}
func (UnimplementedCalculatorServer) BMR(context.Context, *GetUserBMRRequest) (*BMRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BMR not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_BMI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).BMI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_BMI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).BMI(ctx, req.(*GetUserBMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_BMR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBMRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).BMR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_BMR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).BMR(ctx, req.(*GetUserBMRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocalculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BMI",
			Handler:    _Calculator_BMI_Handler,
		},
		{
			MethodName: "BMR",
			Handler:    _Calculator_BMR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_input.proto",
}
