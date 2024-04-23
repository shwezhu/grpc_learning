// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: pokemon.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Pokemon_Create_FullMethodName  = "/pokemon/Create"
	Pokemon_Read_FullMethodName    = "/pokemon/Read"
	Pokemon_ReadOne_FullMethodName = "/pokemon/ReadOne"
	Pokemon_Update_FullMethodName  = "/pokemon/Update"
	Pokemon_Delete_FullMethodName  = "/pokemon/Delete"
)

// PokemonClient is the client API for Pokemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokemonClient interface {
	Create(ctx context.Context, in *PokemonRequest, opts ...grpc.CallOption) (*Pokemon, error)
	Read(ctx context.Context, in *PokemonFilter, opts ...grpc.CallOption) (*PokemonListResponse, error)
	ReadOne(ctx context.Context, in *PokemonID, opts ...grpc.CallOption) (*Pokemon, error)
	Update(ctx context.Context, in *PokemonUpdateRequest, opts ...grpc.CallOption) (*Pokemon, error)
	Delete(ctx context.Context, in *PokemonID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pokemonClient struct {
	cc grpc.ClientConnInterface
}

func NewPokemonClient(cc grpc.ClientConnInterface) PokemonClient {
	return &pokemonClient{cc}
}

func (c *pokemonClient) Create(ctx context.Context, in *PokemonRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, Pokemon_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonClient) Read(ctx context.Context, in *PokemonFilter, opts ...grpc.CallOption) (*PokemonListResponse, error) {
	out := new(PokemonListResponse)
	err := c.cc.Invoke(ctx, Pokemon_Read_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonClient) ReadOne(ctx context.Context, in *PokemonID, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, Pokemon_ReadOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonClient) Update(ctx context.Context, in *PokemonUpdateRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, Pokemon_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonClient) Delete(ctx context.Context, in *PokemonID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Pokemon_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokemonServer is the server API for Pokemon service.
// All implementations must embed UnimplementedPokemonServer
// for forward compatibility
type PokemonServer interface {
	Create(context.Context, *PokemonRequest) (*Pokemon, error)
	Read(context.Context, *PokemonFilter) (*PokemonListResponse, error)
	ReadOne(context.Context, *PokemonID) (*Pokemon, error)
	Update(context.Context, *PokemonUpdateRequest) (*Pokemon, error)
	Delete(context.Context, *PokemonID) (*emptypb.Empty, error)
	mustEmbedUnimplementedPokemonServer()
}

// UnimplementedPokemonServer must be embedded to have forward compatible implementations.
type UnimplementedPokemonServer struct {
}

func (UnimplementedPokemonServer) Create(context.Context, *PokemonRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPokemonServer) Read(context.Context, *PokemonFilter) (*PokemonListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedPokemonServer) ReadOne(context.Context, *PokemonID) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOne not implemented")
}
func (UnimplementedPokemonServer) Update(context.Context, *PokemonUpdateRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPokemonServer) Delete(context.Context, *PokemonID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPokemonServer) mustEmbedUnimplementedPokemonServer() {}

// UnsafePokemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokemonServer will
// result in compilation errors.
type UnsafePokemonServer interface {
	mustEmbedUnimplementedPokemonServer()
}

func RegisterPokemonServer(s grpc.ServiceRegistrar, srv PokemonServer) {
	s.RegisterService(&Pokemon_ServiceDesc, srv)
}

func _Pokemon_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pokemon_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServer).Create(ctx, req.(*PokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokemon_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pokemon_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServer).Read(ctx, req.(*PokemonFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokemon_ReadOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServer).ReadOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pokemon_ReadOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServer).ReadOne(ctx, req.(*PokemonID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokemon_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pokemon_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServer).Update(ctx, req.(*PokemonUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokemon_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pokemon_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServer).Delete(ctx, req.(*PokemonID))
	}
	return interceptor(ctx, in, info, handler)
}

// Pokemon_ServiceDesc is the grpc.ServiceDesc for Pokemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pokemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon",
	HandlerType: (*PokemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Pokemon_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Pokemon_Read_Handler,
		},
		{
			MethodName: "ReadOne",
			Handler:    _Pokemon_ReadOne_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Pokemon_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Pokemon_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon.proto",
}