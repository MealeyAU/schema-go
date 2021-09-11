// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// IdServiceClient is the client API for IdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdServiceClient interface {
	// Generate will create and store a new id within the system
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResult, error)
	// Verify will return the status of a given id
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResult, error)
	// Delete will mark the id as deleted, typically signifying the resource associated
	// being deleted
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResult, error)
}

type idServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdServiceClient(cc grpc.ClientConnInterface) IdServiceClient {
	return &idServiceClient{cc}
}

func (c *idServiceClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResult, error) {
	out := new(GenerateResult)
	err := c.cc.Invoke(ctx, "/mealey.api.ids.v1.IdService/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idServiceClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResult, error) {
	out := new(VerifyResult)
	err := c.cc.Invoke(ctx, "/mealey.api.ids.v1.IdService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResult, error) {
	out := new(DeleteResult)
	err := c.cc.Invoke(ctx, "/mealey.api.ids.v1.IdService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdServiceServer is the server API for IdService service.
// All implementations must embed UnimplementedIdServiceServer
// for forward compatibility
type IdServiceServer interface {
	// Generate will create and store a new id within the system
	Generate(context.Context, *GenerateRequest) (*GenerateResult, error)
	// Verify will return the status of a given id
	Verify(context.Context, *VerifyRequest) (*VerifyResult, error)
	// Delete will mark the id as deleted, typically signifying the resource associated
	// being deleted
	Delete(context.Context, *DeleteRequest) (*DeleteResult, error)
	mustEmbedUnimplementedIdServiceServer()
}

// UnimplementedIdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIdServiceServer struct {
}

func (UnimplementedIdServiceServer) Generate(context.Context, *GenerateRequest) (*GenerateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedIdServiceServer) Verify(context.Context, *VerifyRequest) (*VerifyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedIdServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedIdServiceServer) mustEmbedUnimplementedIdServiceServer() {}

// UnsafeIdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdServiceServer will
// result in compilation errors.
type UnsafeIdServiceServer interface {
	mustEmbedUnimplementedIdServiceServer()
}

func RegisterIdServiceServer(s grpc.ServiceRegistrar, srv IdServiceServer) {
	s.RegisterService(&IdService_ServiceDesc, srv)
}

func _IdService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mealey.api.ids.v1.IdService/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServiceServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mealey.api.ids.v1.IdService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServiceServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mealey.api.ids.v1.IdService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdService_ServiceDesc is the grpc.ServiceDesc for IdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mealey.api.ids.v1.IdService",
	HandlerType: (*IdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _IdService_Generate_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _IdService_Verify_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _IdService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ids/v1/ids.proto",
}
