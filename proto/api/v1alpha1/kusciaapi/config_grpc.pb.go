// Copyright 2024 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.8
// source: kuscia/proto/api/v1alpha1/kusciaapi/config.proto

package kusciaapi

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
	ConfigService_CreateConfig_FullMethodName     = "/kuscia.proto.api.v1alpha1.kusciaapi.ConfigService/CreateConfig"
	ConfigService_QueryConfig_FullMethodName      = "/kuscia.proto.api.v1alpha1.kusciaapi.ConfigService/QueryConfig"
	ConfigService_UpdateConfig_FullMethodName     = "/kuscia.proto.api.v1alpha1.kusciaapi.ConfigService/UpdateConfig"
	ConfigService_DeleteConfig_FullMethodName     = "/kuscia.proto.api.v1alpha1.kusciaapi.ConfigService/DeleteConfig"
	ConfigService_BatchQueryConfig_FullMethodName = "/kuscia.proto.api.v1alpha1.kusciaapi.ConfigService/BatchQueryConfig"
)

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigServiceClient interface {
	CreateConfig(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*CreateConfigResponse, error)
	QueryConfig(ctx context.Context, in *QueryConfigRequest, opts ...grpc.CallOption) (*QueryConfigResponse, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error)
	DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*DeleteConfigResponse, error)
	BatchQueryConfig(ctx context.Context, in *BatchQueryConfigRequest, opts ...grpc.CallOption) (*BatchQueryConfigResponse, error)
}

type configServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServiceClient(cc grpc.ClientConnInterface) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) CreateConfig(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*CreateConfigResponse, error) {
	out := new(CreateConfigResponse)
	err := c.cc.Invoke(ctx, ConfigService_CreateConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) QueryConfig(ctx context.Context, in *QueryConfigRequest, opts ...grpc.CallOption) (*QueryConfigResponse, error) {
	out := new(QueryConfigResponse)
	err := c.cc.Invoke(ctx, ConfigService_QueryConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error) {
	out := new(UpdateConfigResponse)
	err := c.cc.Invoke(ctx, ConfigService_UpdateConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*DeleteConfigResponse, error) {
	out := new(DeleteConfigResponse)
	err := c.cc.Invoke(ctx, ConfigService_DeleteConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) BatchQueryConfig(ctx context.Context, in *BatchQueryConfigRequest, opts ...grpc.CallOption) (*BatchQueryConfigResponse, error) {
	out := new(BatchQueryConfigResponse)
	err := c.cc.Invoke(ctx, ConfigService_BatchQueryConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
// All implementations must embed UnimplementedConfigServiceServer
// for forward compatibility
type ConfigServiceServer interface {
	CreateConfig(context.Context, *CreateConfigRequest) (*CreateConfigResponse, error)
	QueryConfig(context.Context, *QueryConfigRequest) (*QueryConfigResponse, error)
	UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error)
	DeleteConfig(context.Context, *DeleteConfigRequest) (*DeleteConfigResponse, error)
	BatchQueryConfig(context.Context, *BatchQueryConfigRequest) (*BatchQueryConfigResponse, error)
	mustEmbedUnimplementedConfigServiceServer()
}

// UnimplementedConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (UnimplementedConfigServiceServer) CreateConfig(context.Context, *CreateConfigRequest) (*CreateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConfig not implemented")
}
func (UnimplementedConfigServiceServer) QueryConfig(context.Context, *QueryConfigRequest) (*QueryConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryConfig not implemented")
}
func (UnimplementedConfigServiceServer) UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedConfigServiceServer) DeleteConfig(context.Context, *DeleteConfigRequest) (*DeleteConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}
func (UnimplementedConfigServiceServer) BatchQueryConfig(context.Context, *BatchQueryConfigRequest) (*BatchQueryConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchQueryConfig not implemented")
}
func (UnimplementedConfigServiceServer) mustEmbedUnimplementedConfigServiceServer() {}

// UnsafeConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServiceServer will
// result in compilation errors.
type UnsafeConfigServiceServer interface {
	mustEmbedUnimplementedConfigServiceServer()
}

func RegisterConfigServiceServer(s grpc.ServiceRegistrar, srv ConfigServiceServer) {
	s.RegisterService(&ConfigService_ServiceDesc, srv)
}

func _ConfigService_CreateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).CreateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_CreateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).CreateConfig(ctx, req.(*CreateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_QueryConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).QueryConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_QueryConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).QueryConfig(ctx, req.(*QueryConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_UpdateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_DeleteConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).DeleteConfig(ctx, req.(*DeleteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_BatchQueryConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchQueryConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).BatchQueryConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_BatchQueryConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).BatchQueryConfig(ctx, req.(*BatchQueryConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigService_ServiceDesc is the grpc.ServiceDesc for ConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kuscia.proto.api.v1alpha1.kusciaapi.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConfig",
			Handler:    _ConfigService_CreateConfig_Handler,
		},
		{
			MethodName: "QueryConfig",
			Handler:    _ConfigService_QueryConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _ConfigService_UpdateConfig_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _ConfigService_DeleteConfig_Handler,
		},
		{
			MethodName: "BatchQueryConfig",
			Handler:    _ConfigService_BatchQueryConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kuscia/proto/api/v1alpha1/kusciaapi/config.proto",
}
