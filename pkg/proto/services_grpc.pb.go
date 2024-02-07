// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: pkg/proto/services.proto

package proto

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

// KeyValueServiceClient is the client API for KeyValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyValueServiceClient interface {
	SetKeyValue(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValueResponse, error)
	GetValue(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyValueResponse, error)
}

type keyValueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyValueServiceClient(cc grpc.ClientConnInterface) KeyValueServiceClient {
	return &keyValueServiceClient{cc}
}

func (c *keyValueServiceClient) SetKeyValue(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValueResponse, error) {
	out := new(KeyValueResponse)
	err := c.cc.Invoke(ctx, "/proto.KeyValueService/SetKeyValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) GetValue(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyValueResponse, error) {
	out := new(KeyValueResponse)
	err := c.cc.Invoke(ctx, "/proto.KeyValueService/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueServiceServer is the server API for KeyValueService service.
// All implementations must embed UnimplementedKeyValueServiceServer
// for forward compatibility
type KeyValueServiceServer interface {
	SetKeyValue(context.Context, *KeyValue) (*KeyValueResponse, error)
	GetValue(context.Context, *Key) (*KeyValueResponse, error)
	mustEmbedUnimplementedKeyValueServiceServer()
}

// UnimplementedKeyValueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyValueServiceServer struct {
}

func (UnimplementedKeyValueServiceServer) SetKeyValue(context.Context, *KeyValue) (*KeyValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetKeyValue not implemented")
}
func (UnimplementedKeyValueServiceServer) GetValue(context.Context, *Key) (*KeyValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedKeyValueServiceServer) mustEmbedUnimplementedKeyValueServiceServer() {}

// UnsafeKeyValueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyValueServiceServer will
// result in compilation errors.
type UnsafeKeyValueServiceServer interface {
	mustEmbedUnimplementedKeyValueServiceServer()
}

func RegisterKeyValueServiceServer(s grpc.ServiceRegistrar, srv KeyValueServiceServer) {
	s.RegisterService(&KeyValueService_ServiceDesc, srv)
}

func _KeyValueService_SetKeyValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).SetKeyValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KeyValueService/SetKeyValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).SetKeyValue(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KeyValueService/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).GetValue(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyValueService_ServiceDesc is the grpc.ServiceDesc for KeyValueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyValueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.KeyValueService",
	HandlerType: (*KeyValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetKeyValue",
			Handler:    _KeyValueService_SetKeyValue_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _KeyValueService_GetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/services.proto",
}