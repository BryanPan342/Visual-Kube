// Code generated by protoc-gen-go.
// source: perf.proto
// DO NOT EDIT!

/*
Package codec_perf is a generated protocol buffer package.

It is generated from these files:
	perf.proto

It has these top-level messages:
	Buffer
*/
package codec_perf

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

// Buffer is a message that contains a body of bytes that is used to exercise
// encoding and decoding overheads.
type Buffer struct {
	Body             []byte `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Buffer) Reset()                    { *m = Buffer{} }
func (m *Buffer) String() string            { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()               {}
func (*Buffer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Buffer) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Buffer)(nil), "codec.perf.Buffer")
}

func init() { proto.RegisterFile("perf.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 73 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x48, 0x2d, 0x4a,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0xce, 0x4f, 0x49, 0x4d, 0xd6, 0x03, 0x89,
	0x28, 0xc9, 0x70, 0xb1, 0x39, 0x95, 0xa6, 0xa5, 0xa5, 0x16, 0x09, 0x09, 0x71, 0xb1, 0x24, 0xe5,
	0xa7, 0x54, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x3a, 0x58, 0x92, 0x53, 0x36, 0x00, 0x00, 0x00,
}
