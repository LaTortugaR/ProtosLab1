// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoslab3/messages.proto

package messages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Comando struct {
	Accion               string   `protobuf:"bytes,1,opt,name=accion,proto3" json:"accion,omitempty"`
	Sector               string   `protobuf:"bytes,2,opt,name=sector,proto3" json:"sector,omitempty"`
	Base                 string   `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	Valor                string   `protobuf:"bytes,4,opt,name=valor,proto3" json:"valor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comando) Reset()         { *m = Comando{} }
func (m *Comando) String() string { return proto.CompactTextString(m) }
func (*Comando) ProtoMessage()    {}
func (*Comando) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{0}
}

func (m *Comando) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comando.Unmarshal(m, b)
}
func (m *Comando) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comando.Marshal(b, m, deterministic)
}
func (m *Comando) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comando.Merge(m, src)
}
func (m *Comando) XXX_Size() int {
	return xxx_messageInfo_Comando.Size(m)
}
func (m *Comando) XXX_DiscardUnknown() {
	xxx_messageInfo_Comando.DiscardUnknown(m)
}

var xxx_messageInfo_Comando proto.InternalMessageInfo

func (m *Comando) GetAccion() string {
	if m != nil {
		return m.Accion
	}
	return ""
}

func (m *Comando) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

func (m *Comando) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *Comando) GetValor() string {
	if m != nil {
		return m.Valor
	}
	return ""
}

type Direccion struct {
	Direccion            string   `protobuf:"bytes,1,opt,name=direccion,proto3" json:"direccion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Direccion) Reset()         { *m = Direccion{} }
func (m *Direccion) String() string { return proto.CompactTextString(m) }
func (*Direccion) ProtoMessage()    {}
func (*Direccion) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{1}
}

func (m *Direccion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Direccion.Unmarshal(m, b)
}
func (m *Direccion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Direccion.Marshal(b, m, deterministic)
}
func (m *Direccion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Direccion.Merge(m, src)
}
func (m *Direccion) XXX_Size() int {
	return xxx_messageInfo_Direccion.Size(m)
}
func (m *Direccion) XXX_DiscardUnknown() {
	xxx_messageInfo_Direccion.DiscardUnknown(m)
}

var xxx_messageInfo_Direccion proto.InternalMessageInfo

func (m *Direccion) GetDireccion() string {
	if m != nil {
		return m.Direccion
	}
	return ""
}

type Vector struct {
	X                    string   `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    string   `protobuf:"bytes,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    string   `protobuf:"bytes,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector) Reset()         { *m = Vector{} }
func (m *Vector) String() string { return proto.CompactTextString(m) }
func (*Vector) ProtoMessage()    {}
func (*Vector) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{2}
}

func (m *Vector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector.Unmarshal(m, b)
}
func (m *Vector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector.Marshal(b, m, deterministic)
}
func (m *Vector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector.Merge(m, src)
}
func (m *Vector) XXX_Size() int {
	return xxx_messageInfo_Vector.Size(m)
}
func (m *Vector) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector.DiscardUnknown(m)
}

var xxx_messageInfo_Vector proto.InternalMessageInfo

func (m *Vector) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *Vector) GetY() string {
	if m != nil {
		return m.Y
	}
	return ""
}

func (m *Vector) GetZ() string {
	if m != nil {
		return m.Z
	}
	return ""
}

type Sector struct {
	Sector               string   `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sector) Reset()         { *m = Sector{} }
func (m *Sector) String() string { return proto.CompactTextString(m) }
func (*Sector) ProtoMessage()    {}
func (*Sector) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{3}
}

func (m *Sector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sector.Unmarshal(m, b)
}
func (m *Sector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sector.Marshal(b, m, deterministic)
}
func (m *Sector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sector.Merge(m, src)
}
func (m *Sector) XXX_Size() int {
	return xxx_messageInfo_Sector.Size(m)
}
func (m *Sector) XXX_DiscardUnknown() {
	xxx_messageInfo_Sector.DiscardUnknown(m)
}

var xxx_messageInfo_Sector proto.InternalMessageInfo

func (m *Sector) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

type Base struct {
	Base                 string   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Base) Reset()         { *m = Base{} }
func (m *Base) String() string { return proto.CompactTextString(m) }
func (*Base) ProtoMessage()    {}
func (*Base) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{4}
}

func (m *Base) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Base.Unmarshal(m, b)
}
func (m *Base) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Base.Marshal(b, m, deterministic)
}
func (m *Base) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Base.Merge(m, src)
}
func (m *Base) XXX_Size() int {
	return xxx_messageInfo_Base.Size(m)
}
func (m *Base) XXX_DiscardUnknown() {
	xxx_messageInfo_Base.DiscardUnknown(m)
}

var xxx_messageInfo_Base proto.InternalMessageInfo

func (m *Base) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

type Confirmar struct {
	Flag                 string   `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Confirmar) Reset()         { *m = Confirmar{} }
func (m *Confirmar) String() string { return proto.CompactTextString(m) }
func (*Confirmar) ProtoMessage()    {}
func (*Confirmar) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{5}
}

func (m *Confirmar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Confirmar.Unmarshal(m, b)
}
func (m *Confirmar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Confirmar.Marshal(b, m, deterministic)
}
func (m *Confirmar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Confirmar.Merge(m, src)
}
func (m *Confirmar) XXX_Size() int {
	return xxx_messageInfo_Confirmar.Size(m)
}
func (m *Confirmar) XXX_DiscardUnknown() {
	xxx_messageInfo_Confirmar.DiscardUnknown(m)
}

var xxx_messageInfo_Confirmar proto.InternalMessageInfo

func (m *Confirmar) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

type Soldados struct {
	Sector               string   `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base                 string   `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Soldados) Reset()         { *m = Soldados{} }
func (m *Soldados) String() string { return proto.CompactTextString(m) }
func (*Soldados) ProtoMessage()    {}
func (*Soldados) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{6}
}

func (m *Soldados) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Soldados.Unmarshal(m, b)
}
func (m *Soldados) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Soldados.Marshal(b, m, deterministic)
}
func (m *Soldados) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Soldados.Merge(m, src)
}
func (m *Soldados) XXX_Size() int {
	return xxx_messageInfo_Soldados.Size(m)
}
func (m *Soldados) XXX_DiscardUnknown() {
	xxx_messageInfo_Soldados.DiscardUnknown(m)
}

var xxx_messageInfo_Soldados proto.InternalMessageInfo

func (m *Soldados) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

func (m *Soldados) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

type Numero struct {
	Numero               string   `protobuf:"bytes,1,opt,name=numero,proto3" json:"numero,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Numero) Reset()         { *m = Numero{} }
func (m *Numero) String() string { return proto.CompactTextString(m) }
func (*Numero) ProtoMessage()    {}
func (*Numero) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{7}
}

func (m *Numero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Numero.Unmarshal(m, b)
}
func (m *Numero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Numero.Marshal(b, m, deterministic)
}
func (m *Numero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Numero.Merge(m, src)
}
func (m *Numero) XXX_Size() int {
	return xxx_messageInfo_Numero.Size(m)
}
func (m *Numero) XXX_DiscardUnknown() {
	xxx_messageInfo_Numero.DiscardUnknown(m)
}

var xxx_messageInfo_Numero proto.InternalMessageInfo

func (m *Numero) GetNumero() string {
	if m != nil {
		return m.Numero
	}
	return ""
}

type Cambio struct {
	Nombre               string   `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Tiempo1              string   `protobuf:"bytes,2,opt,name=tiempo1,proto3" json:"tiempo1,omitempty"`
	Valor                string   `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tiempo2              string   `protobuf:"bytes,4,opt,name=tiempo2,proto3" json:"tiempo2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cambio) Reset()         { *m = Cambio{} }
func (m *Cambio) String() string { return proto.CompactTextString(m) }
func (*Cambio) ProtoMessage()    {}
func (*Cambio) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{8}
}

func (m *Cambio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cambio.Unmarshal(m, b)
}
func (m *Cambio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cambio.Marshal(b, m, deterministic)
}
func (m *Cambio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cambio.Merge(m, src)
}
func (m *Cambio) XXX_Size() int {
	return xxx_messageInfo_Cambio.Size(m)
}
func (m *Cambio) XXX_DiscardUnknown() {
	xxx_messageInfo_Cambio.DiscardUnknown(m)
}

var xxx_messageInfo_Cambio proto.InternalMessageInfo

func (m *Cambio) GetNombre() string {
	if m != nil {
		return m.Nombre
	}
	return ""
}

func (m *Cambio) GetTiempo1() string {
	if m != nil {
		return m.Tiempo1
	}
	return ""
}

func (m *Cambio) GetValor() string {
	if m != nil {
		return m.Valor
	}
	return ""
}

func (m *Cambio) GetTiempo2() string {
	if m != nil {
		return m.Tiempo2
	}
	return ""
}

type SectorVector struct {
	Sector               string   `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	X                    string   `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    string   `protobuf:"bytes,3,opt,name=y,proto3" json:"y,omitempty"`
	Z                    string   `protobuf:"bytes,4,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SectorVector) Reset()         { *m = SectorVector{} }
func (m *SectorVector) String() string { return proto.CompactTextString(m) }
func (*SectorVector) ProtoMessage()    {}
func (*SectorVector) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{9}
}

func (m *SectorVector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SectorVector.Unmarshal(m, b)
}
func (m *SectorVector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SectorVector.Marshal(b, m, deterministic)
}
func (m *SectorVector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SectorVector.Merge(m, src)
}
func (m *SectorVector) XXX_Size() int {
	return xxx_messageInfo_SectorVector.Size(m)
}
func (m *SectorVector) XXX_DiscardUnknown() {
	xxx_messageInfo_SectorVector.DiscardUnknown(m)
}

var xxx_messageInfo_SectorVector proto.InternalMessageInfo

func (m *SectorVector) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

func (m *SectorVector) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *SectorVector) GetY() string {
	if m != nil {
		return m.Y
	}
	return ""
}

func (m *SectorVector) GetZ() string {
	if m != nil {
		return m.Z
	}
	return ""
}

func init() {
	proto.RegisterType((*Comando)(nil), "Comando")
	proto.RegisterType((*Direccion)(nil), "Direccion")
	proto.RegisterType((*Vector)(nil), "Vector")
	proto.RegisterType((*Sector)(nil), "Sector")
	proto.RegisterType((*Base)(nil), "Base")
	proto.RegisterType((*Confirmar)(nil), "Confirmar")
	proto.RegisterType((*Soldados)(nil), "Soldados")
	proto.RegisterType((*Numero)(nil), "Numero")
	proto.RegisterType((*Cambio)(nil), "Cambio")
	proto.RegisterType((*SectorVector)(nil), "SectorVector")
}

func init() { proto.RegisterFile("protoslab3/messages.proto", fileDescriptor_d2e02856b5ed9737) }

var fileDescriptor_d2e02856b5ed9737 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x8a, 0xd4, 0x30,
	0x14, 0xc6, 0x27, 0x33, 0xb3, 0xed, 0xf6, 0x38, 0xeb, 0x60, 0x14, 0xa9, 0x45, 0x71, 0x8c, 0xf8,
	0x0f, 0x31, 0xe2, 0x2c, 0xf8, 0x00, 0x33, 0x22, 0x78, 0xa1, 0x2c, 0x0e, 0xec, 0x85, 0x37, 0x92,
	0xa6, 0xd9, 0x25, 0xda, 0x36, 0x4b, 0xd2, 0x59, 0x76, 0xf7, 0x0d, 0x7c, 0x68, 0x41, 0x9a, 0xa4,
	0x69, 0x47, 0x18, 0xc1, 0xab, 0xc9, 0x77, 0xe6, 0xf4, 0x3b, 0xdf, 0xf9, 0x35, 0x85, 0x07, 0x17,
	0x5a, 0x35, 0xca, 0x94, 0x2c, 0x3f, 0x7e, 0x5b, 0x09, 0x63, 0xd8, 0xb9, 0x30, 0xd4, 0xd6, 0x08,
	0x87, 0x78, 0xad, 0x2a, 0x56, 0x17, 0x0a, 0xdf, 0x87, 0x88, 0x71, 0x2e, 0x55, 0x9d, 0xa2, 0x05,
	0x7a, 0x99, 0x7c, 0xf5, 0xaa, 0xad, 0x1b, 0xc1, 0x1b, 0xa5, 0xd3, 0xb1, 0xab, 0x3b, 0x85, 0x31,
	0x4c, 0x73, 0x66, 0x44, 0x3a, 0xb1, 0x55, 0x7b, 0xc6, 0xf7, 0xe0, 0xe0, 0x92, 0x95, 0x4a, 0xa7,
	0x53, 0x5b, 0x74, 0x82, 0xbc, 0x82, 0xe4, 0x83, 0xd4, 0xc2, 0xd9, 0x3d, 0x84, 0xa4, 0xe8, 0x84,
	0x9f, 0xd4, 0x17, 0xc8, 0x12, 0xa2, 0x53, 0x67, 0x3f, 0x03, 0x74, 0xe5, 0xff, 0x47, 0x57, 0xad,
	0xba, 0xf6, 0xf3, 0xd1, 0x75, 0xab, 0x6e, 0xfc, 0x5c, 0x74, 0x43, 0x16, 0x10, 0x6d, 0xdc, 0x33,
	0x7d, 0x54, 0x34, 0x8c, 0x4a, 0x32, 0x98, 0xae, 0xda, 0x78, 0x5d, 0x64, 0xd4, 0x47, 0x26, 0x8f,
	0x21, 0x59, 0xab, 0xfa, 0x4c, 0xea, 0x8a, 0xd9, 0x9d, 0xce, 0x4a, 0x76, 0xde, 0x35, 0xb4, 0x67,
	0xf2, 0x1e, 0x0e, 0x37, 0xaa, 0x2c, 0x58, 0xa1, 0xcc, 0xbe, 0x01, 0xc1, 0x78, 0x3c, 0x30, 0x5e,
	0x40, 0xf4, 0x65, 0x5b, 0x09, 0x6d, 0xc9, 0xd6, 0xf6, 0xd4, 0x3d, 0xe5, 0x14, 0xf9, 0x01, 0xd1,
	0x9a, 0x55, 0xb9, 0x74, 0x1d, 0xaa, 0xca, 0xb5, 0x08, 0x1d, 0x56, 0xe1, 0x14, 0xe2, 0x46, 0x8a,
	0xea, 0x42, 0xbd, 0xf3, 0xd6, 0x9d, 0xec, 0x49, 0x4f, 0x06, 0xa4, 0xfb, 0xfe, 0xa5, 0x7f, 0x03,
	0x9d, 0x24, 0x27, 0x30, 0x73, 0x90, 0x4e, 0xff, 0x89, 0xca, 0x61, 0x1f, 0xef, 0x60, 0x9f, 0xec,
	0x60, 0x9f, 0x7a, 0xec, 0xcb, 0x5f, 0x08, 0xe6, 0x2b, 0xad, 0xf8, 0x4f, 0xa1, 0xbf, 0x6f, 0x84,
	0xbe, 0x94, 0x5c, 0xe0, 0xe7, 0x30, 0xfb, 0xcc, 0xea, 0x82, 0x69, 0xbf, 0xd7, 0x21, 0xf5, 0xb7,
	0x2b, 0x03, 0x1a, 0xae, 0x00, 0x19, 0xe1, 0x67, 0x70, 0x74, 0x22, 0x0a, 0xa9, 0x03, 0xd8, 0x84,
	0x76, 0xc7, 0x2c, 0xa6, 0x0e, 0x1b, 0x19, 0xe1, 0x17, 0x70, 0xfb, 0x53, 0xcd, 0x55, 0x6d, 0xa4,
	0x69, 0x44, 0xcd, 0x25, 0xc3, 0x31, 0x75, 0x5b, 0x64, 0x40, 0xc3, 0x5b, 0x23, 0xa3, 0xe5, 0x6f,
	0x04, 0xf3, 0x8f, 0xdb, 0x92, 0xeb, 0x6d, 0x15, 0xb2, 0x3c, 0xdd, 0x9b, 0x25, 0xa6, 0x0e, 0xc2,
	0x7f, 0x4c, 0xc0, 0x4f, 0xe0, 0x96, 0x4d, 0xec, 0xf1, 0x85, 0xae, 0x81, 0xd7, 0xeb, 0x6e, 0xa0,
	0xef, 0x39, 0xa2, 0x43, 0xe2, 0x7f, 0xf9, 0x3d, 0x82, 0xc4, 0xfa, 0xd9, 0x7b, 0x79, 0x40, 0xdb,
	0x9f, 0x2c, 0xa6, 0x2e, 0x2a, 0x19, 0xe1, 0x37, 0x70, 0x77, 0x18, 0xde, 0xef, 0xb6, 0xc3, 0x73,
	0xe0, 0xb6, 0xba, 0xf3, 0x6d, 0xee, 0xbe, 0xf1, 0xf0, 0x7d, 0xe7, 0x91, 0x2d, 0x1c, 0xff, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x21, 0x6c, 0x76, 0xcf, 0xfd, 0x03, 0x00, 0x00,
}
