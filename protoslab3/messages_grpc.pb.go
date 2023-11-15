// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protoslab3/messages.proto

package messages

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

// Brocker_ServiceClient is the client API for Brocker_Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Brocker_ServiceClient interface {
	// Informante quiere hacer un cambio
	MandarCambio(ctx context.Context, in *Cambio, opts ...grpc.CallOption) (*Direccion, error)
	// Vanguardia pide soldados
	PedirSoldados(ctx context.Context, in *Soldados, opts ...grpc.CallOption) (*Numero, error)
	// Informante/Vanguardia declara inconsistencia
	Inconsistencia(ctx context.Context, in *Sector, opts ...grpc.CallOption) (*Confirmar, error)
}

type brocker_ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrocker_ServiceClient(cc grpc.ClientConnInterface) Brocker_ServiceClient {
	return &brocker_ServiceClient{cc}
}

func (c *brocker_ServiceClient) MandarCambio(ctx context.Context, in *Cambio, opts ...grpc.CallOption) (*Direccion, error) {
	out := new(Direccion)
	err := c.cc.Invoke(ctx, "/Brocker_Service/MandarCambio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brocker_ServiceClient) PedirSoldados(ctx context.Context, in *Soldados, opts ...grpc.CallOption) (*Numero, error) {
	out := new(Numero)
	err := c.cc.Invoke(ctx, "/Brocker_Service/PedirSoldados", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brocker_ServiceClient) Inconsistencia(ctx context.Context, in *Sector, opts ...grpc.CallOption) (*Confirmar, error) {
	out := new(Confirmar)
	err := c.cc.Invoke(ctx, "/Brocker_Service/Inconsistencia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Brocker_ServiceServer is the server API for Brocker_Service service.
// All implementations must embed UnimplementedBrocker_ServiceServer
// for forward compatibility
type Brocker_ServiceServer interface {
	// Informante quiere hacer un cambio
	MandarCambio(context.Context, *Cambio) (*Direccion, error)
	// Vanguardia pide soldados
	PedirSoldados(context.Context, *Soldados) (*Numero, error)
	// Informante/Vanguardia declara inconsistencia
	Inconsistencia(context.Context, *Sector) (*Confirmar, error)
	mustEmbedUnimplementedBrocker_ServiceServer()
}

// UnimplementedBrocker_ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBrocker_ServiceServer struct {
}

func (UnimplementedBrocker_ServiceServer) MandarCambio(context.Context, *Cambio) (*Direccion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MandarCambio not implemented")
}
func (UnimplementedBrocker_ServiceServer) PedirSoldados(context.Context, *Soldados) (*Numero, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirSoldados not implemented")
}
func (UnimplementedBrocker_ServiceServer) Inconsistencia(context.Context, *Sector) (*Confirmar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inconsistencia not implemented")
}
func (UnimplementedBrocker_ServiceServer) mustEmbedUnimplementedBrocker_ServiceServer() {}

// UnsafeBrocker_ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Brocker_ServiceServer will
// result in compilation errors.
type UnsafeBrocker_ServiceServer interface {
	mustEmbedUnimplementedBrocker_ServiceServer()
}

func RegisterBrocker_ServiceServer(s grpc.ServiceRegistrar, srv Brocker_ServiceServer) {
	s.RegisterService(&Brocker_Service_ServiceDesc, srv)
}

func _Brocker_Service_MandarCambio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cambio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Brocker_ServiceServer).MandarCambio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Brocker_Service/MandarCambio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Brocker_ServiceServer).MandarCambio(ctx, req.(*Cambio))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brocker_Service_PedirSoldados_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Soldados)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Brocker_ServiceServer).PedirSoldados(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Brocker_Service/PedirSoldados",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Brocker_ServiceServer).PedirSoldados(ctx, req.(*Soldados))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brocker_Service_Inconsistencia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Brocker_ServiceServer).Inconsistencia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Brocker_Service/Inconsistencia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Brocker_ServiceServer).Inconsistencia(ctx, req.(*Sector))
	}
	return interceptor(ctx, in, info, handler)
}

// Brocker_Service_ServiceDesc is the grpc.ServiceDesc for Brocker_Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Brocker_Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Brocker_Service",
	HandlerType: (*Brocker_ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MandarCambio",
			Handler:    _Brocker_Service_MandarCambio_Handler,
		},
		{
			MethodName: "PedirSoldados",
			Handler:    _Brocker_Service_PedirSoldados_Handler,
		},
		{
			MethodName: "Inconsistencia",
			Handler:    _Brocker_Service_Inconsistencia_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoslab3/messages.proto",
}

// Fulcrum_ServiceClient is the client API for Fulcrum_Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Fulcrum_ServiceClient interface {
	// Informante quiere hacer un cambio
	MandarCambio(ctx context.Context, in *Cambio, opts ...grpc.CallOption) (*Vector, error)
	// Brocker declara inconsistencia
	Inconsistencia(ctx context.Context, in *Sector, opts ...grpc.CallOption) (*Confirmar, error)
	// Fulcrum pide Vector
	PedirVector(ctx context.Context, in *Sector, opts ...grpc.CallOption) (*Vector, error)
	// Fulcrum manda Vector nuevo
	MandarVector(ctx context.Context, in *SectorVector, opts ...grpc.CallOption) (*Confirmar, error)
	// Fulcrum pide Base y Valor
	PedirBase(ctx context.Context, in *Base, opts ...grpc.CallOption) (*Cambio, error)
	// Fulcrum quiere hacer un cambio
	MandarCambioFulcrum(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*Confirmar, error)
}

type fulcrum_ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFulcrum_ServiceClient(cc grpc.ClientConnInterface) Fulcrum_ServiceClient {
	return &fulcrum_ServiceClient{cc}
}

func (c *fulcrum_ServiceClient) MandarCambio(ctx context.Context, in *Cambio, opts ...grpc.CallOption) (*Vector, error) {
	out := new(Vector)
	err := c.cc.Invoke(ctx, "/Fulcrum_Service/MandarCambio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrum_ServiceClient) Inconsistencia(ctx context.Context, in *Sector, opts ...grpc.CallOption) (*Confirmar, error) {
	out := new(Confirmar)
	err := c.cc.Invoke(ctx, "/Fulcrum_Service/Inconsistencia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrum_ServiceClient) PedirVector(ctx context.Context, in *Sector, opts ...grpc.CallOption) (*Vector, error) {
	out := new(Vector)
	err := c.cc.Invoke(ctx, "/Fulcrum_Service/PedirVector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrum_ServiceClient) MandarVector(ctx context.Context, in *SectorVector, opts ...grpc.CallOption) (*Confirmar, error) {
	out := new(Confirmar)
	err := c.cc.Invoke(ctx, "/Fulcrum_Service/MandarVector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrum_ServiceClient) PedirBase(ctx context.Context, in *Base, opts ...grpc.CallOption) (*Cambio, error) {
	out := new(Cambio)
	err := c.cc.Invoke(ctx, "/Fulcrum_Service/PedirBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrum_ServiceClient) MandarCambioFulcrum(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*Confirmar, error) {
	out := new(Confirmar)
	err := c.cc.Invoke(ctx, "/Fulcrum_Service/MandarCambioFulcrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Fulcrum_ServiceServer is the server API for Fulcrum_Service service.
// All implementations must embed UnimplementedFulcrum_ServiceServer
// for forward compatibility
type Fulcrum_ServiceServer interface {
	// Informante quiere hacer un cambio
	MandarCambio(context.Context, *Cambio) (*Vector, error)
	// Brocker declara inconsistencia
	Inconsistencia(context.Context, *Sector) (*Confirmar, error)
	// Fulcrum pide Vector
	PedirVector(context.Context, *Sector) (*Vector, error)
	// Fulcrum manda Vector nuevo
	MandarVector(context.Context, *SectorVector) (*Confirmar, error)
	// Fulcrum pide Base y Valor
	PedirBase(context.Context, *Base) (*Cambio, error)
	// Fulcrum quiere hacer un cambio
	MandarCambioFulcrum(context.Context, *Comando) (*Confirmar, error)
	mustEmbedUnimplementedFulcrum_ServiceServer()
}

// UnimplementedFulcrum_ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFulcrum_ServiceServer struct {
}

func (UnimplementedFulcrum_ServiceServer) MandarCambio(context.Context, *Cambio) (*Vector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MandarCambio not implemented")
}
func (UnimplementedFulcrum_ServiceServer) Inconsistencia(context.Context, *Sector) (*Confirmar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inconsistencia not implemented")
}
func (UnimplementedFulcrum_ServiceServer) PedirVector(context.Context, *Sector) (*Vector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirVector not implemented")
}
func (UnimplementedFulcrum_ServiceServer) MandarVector(context.Context, *SectorVector) (*Confirmar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MandarVector not implemented")
}
func (UnimplementedFulcrum_ServiceServer) PedirBase(context.Context, *Base) (*Cambio, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirBase not implemented")
}
func (UnimplementedFulcrum_ServiceServer) MandarCambioFulcrum(context.Context, *Comando) (*Confirmar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MandarCambioFulcrum not implemented")
}
func (UnimplementedFulcrum_ServiceServer) mustEmbedUnimplementedFulcrum_ServiceServer() {}

// UnsafeFulcrum_ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Fulcrum_ServiceServer will
// result in compilation errors.
type UnsafeFulcrum_ServiceServer interface {
	mustEmbedUnimplementedFulcrum_ServiceServer()
}

func RegisterFulcrum_ServiceServer(s grpc.ServiceRegistrar, srv Fulcrum_ServiceServer) {
	s.RegisterService(&Fulcrum_Service_ServiceDesc, srv)
}

func _Fulcrum_Service_MandarCambio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cambio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Fulcrum_ServiceServer).MandarCambio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fulcrum_Service/MandarCambio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Fulcrum_ServiceServer).MandarCambio(ctx, req.(*Cambio))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_Service_Inconsistencia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Fulcrum_ServiceServer).Inconsistencia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fulcrum_Service/Inconsistencia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Fulcrum_ServiceServer).Inconsistencia(ctx, req.(*Sector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_Service_PedirVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Fulcrum_ServiceServer).PedirVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fulcrum_Service/PedirVector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Fulcrum_ServiceServer).PedirVector(ctx, req.(*Sector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_Service_MandarVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SectorVector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Fulcrum_ServiceServer).MandarVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fulcrum_Service/MandarVector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Fulcrum_ServiceServer).MandarVector(ctx, req.(*SectorVector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_Service_PedirBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Base)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Fulcrum_ServiceServer).PedirBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fulcrum_Service/PedirBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Fulcrum_ServiceServer).PedirBase(ctx, req.(*Base))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_Service_MandarCambioFulcrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Fulcrum_ServiceServer).MandarCambioFulcrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fulcrum_Service/MandarCambioFulcrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Fulcrum_ServiceServer).MandarCambioFulcrum(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

// Fulcrum_Service_ServiceDesc is the grpc.ServiceDesc for Fulcrum_Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fulcrum_Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Fulcrum_Service",
	HandlerType: (*Fulcrum_ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MandarCambio",
			Handler:    _Fulcrum_Service_MandarCambio_Handler,
		},
		{
			MethodName: "Inconsistencia",
			Handler:    _Fulcrum_Service_Inconsistencia_Handler,
		},
		{
			MethodName: "PedirVector",
			Handler:    _Fulcrum_Service_PedirVector_Handler,
		},
		{
			MethodName: "MandarVector",
			Handler:    _Fulcrum_Service_MandarVector_Handler,
		},
		{
			MethodName: "PedirBase",
			Handler:    _Fulcrum_Service_PedirBase_Handler,
		},
		{
			MethodName: "MandarCambioFulcrum",
			Handler:    _Fulcrum_Service_MandarCambioFulcrum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoslab3/messages.proto",
}