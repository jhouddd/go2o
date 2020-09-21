// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/product.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeleteProductRequest struct {
	SellerId             int64    `protobuf:"varint,1,opt,name=SellerId,proto3" json:"SellerId,omitempty"`
	ProductId            int64    `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProductRequest) Reset()         { *m = DeleteProductRequest{} }
func (m *DeleteProductRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProductRequest) ProtoMessage()    {}
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_60fa8c6d4c2def2f, []int{0}
}
func (m *DeleteProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProductRequest.Unmarshal(m, b)
}
func (m *DeleteProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProductRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProductRequest.Merge(dst, src)
}
func (m *DeleteProductRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProductRequest.Size(m)
}
func (m *DeleteProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProductRequest proto.InternalMessageInfo

func (m *DeleteProductRequest) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *DeleteProductRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func init() {
	proto.RegisterType((*DeleteProductRequest)(nil), "DeleteProductRequest")
}

func init() { proto.RegisterFile("message/product.proto", fileDescriptor_product_60fa8c6d4c2def2f) }

var fileDescriptor_product_60fa8c6d4c2def2f = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2f, 0x28, 0xca, 0x4f, 0x29, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x57, 0x0a, 0xe0, 0x12, 0x71, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0x0d, 0x80, 0x08, 0x07, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x71, 0x71, 0x04, 0xa7, 0xe6, 0xe4, 0xa4, 0x16, 0x79,
	0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xf9, 0x42, 0x32, 0x5c, 0x9c, 0x50, 0xd5,
	0x9e, 0x29, 0x12, 0x4c, 0x60, 0x49, 0x84, 0x80, 0x13, 0x67, 0x14, 0xbb, 0x9e, 0x35, 0xd8, 0xf0,
	0x24, 0x36, 0x30, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xe4, 0x52, 0x19, 0x7c, 0x00,
	0x00, 0x00,
}
