// Code generated by protoc-gen-go.
// source: tool.proto
// DO NOT EDIT!

/*
Package tool is a generated protocol buffer package.

It is generated from these files:
	tool.proto

It has these top-level messages:
	ExportHeader
	FieldMeta
*/
package tool

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// 电子表格1,1单元格的表头格式描述
type ExportHeader struct {
	ProtoTypeName string `protobuf:"bytes,1,opt,name=ProtoTypeName" json:"ProtoTypeName,omitempty"`
	RowFieldName  string `protobuf:"bytes,2,opt,name=RowFieldName" json:"RowFieldName,omitempty"`
	ProtoFile     string `protobuf:"bytes,3,opt,name=ProtoFile" json:"ProtoFile,omitempty"`
}

func (m *ExportHeader) Reset()                    { *m = ExportHeader{} }
func (m *ExportHeader) String() string            { return proto.CompactTextString(m) }
func (*ExportHeader) ProtoMessage()               {}
func (*ExportHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// proto字段后部注释中的meta扩展描述
type FieldMeta struct {
	String2ListSpliter string `protobuf:"bytes,2,opt,name=String2ListSpliter" json:"String2ListSpliter,omitempty"`
	RepeatCheck        bool   `protobuf:"varint,3,opt,name=RepeatCheck" json:"RepeatCheck,omitempty"`
	Alias              string `protobuf:"bytes,4,opt,name=Alias" json:"Alias,omitempty"`
	String2Struct      bool   `protobuf:"varint,5,opt,name=String2Struct" json:"String2Struct,omitempty"`
}

func (m *FieldMeta) Reset()                    { *m = FieldMeta{} }
func (m *FieldMeta) String() string            { return proto.CompactTextString(m) }
func (*FieldMeta) ProtoMessage()               {}
func (*FieldMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ExportHeader)(nil), "tool.ExportHeader")
	proto.RegisterType((*FieldMeta)(nil), "tool.FieldMeta")
}

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8f, 0xcd, 0x4e, 0xc5, 0x20,
	0x10, 0x85, 0x53, 0xbd, 0x35, 0x32, 0xb7, 0xdd, 0x4c, 0x5c, 0xb0, 0x70, 0x71, 0xd3, 0xb8, 0x70,
	0x75, 0x17, 0xfa, 0x04, 0xc6, 0xd8, 0xb8, 0x50, 0x63, 0xa8, 0x2f, 0x80, 0xed, 0x44, 0x89, 0x28,
	0x04, 0xc6, 0xbf, 0x47, 0xf1, 0x6d, 0x95, 0x49, 0x8d, 0x4d, 0x5c, 0x01, 0xdf, 0x39, 0x7c, 0x30,
	0x00, 0x1c, 0x82, 0xdf, 0xc6, 0x14, 0x38, 0xe0, 0xaa, 0xec, 0xbb, 0x37, 0x68, 0x2e, 0x3e, 0x62,
	0x48, 0x7c, 0x49, 0x76, 0xa2, 0x84, 0x47, 0xd0, 0xde, 0x96, 0xf8, 0xee, 0x33, 0xd2, 0x8d, 0x7d,
	0x26, 0x5d, 0x6d, 0xaa, 0x63, 0x65, 0xda, 0xb8, 0x84, 0xd8, 0x41, 0x63, 0xc2, 0x7b, 0xef, 0xc8,
	0x4f, 0x52, 0xda, 0x91, 0x52, 0x93, 0x16, 0x0c, 0x0f, 0x41, 0x89, 0xa9, 0x77, 0x9e, 0xf4, 0xae,
	0x14, 0x54, 0xfc, 0x05, 0xdd, 0x57, 0x05, 0x4a, 0xba, 0xd7, 0xc4, 0x16, 0xb7, 0x80, 0x03, 0x27,
	0xf7, 0xf2, 0x70, 0x72, 0xe5, 0x32, 0x0f, 0xd1, 0x3b, 0xa6, 0x34, 0x5b, 0x31, 0xff, 0x4b, 0x70,
	0x03, 0x6b, 0x43, 0x91, 0x2c, 0x9f, 0x3f, 0xd2, 0xf8, 0x24, 0xf6, 0x7d, 0xb3, 0x4e, 0x7f, 0x08,
	0x0f, 0xa0, 0x3e, 0xf3, 0xce, 0x66, 0xbd, 0x12, 0x49, 0x6d, 0xcb, 0xa1, 0x4c, 0x37, 0xbf, 0xf3,
	0xb3, 0xbc, 0x8e, 0xac, 0x6b, 0xb9, 0xd9, 0xe6, 0x25, 0xbc, 0xdf, 0x93, 0x6f, 0x9e, 0x7e, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x0c, 0x9f, 0x10, 0xa8, 0x2e, 0x01, 0x00, 0x00,
}
