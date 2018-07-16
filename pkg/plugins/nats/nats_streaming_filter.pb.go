// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nats_streaming_filter.proto

/*
Package nats is a generated protocol buffer package.

It is generated from these files:
	nats_streaming_filter.proto

It has these top-level messages:
	NatsStreaming
*/
package nats

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// [#proto-status: experimental]
type NatsStreaming struct {
	Cluster        string         `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	MaxConnections uint32         `protobuf:"varint,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	OpTimeout      *time.Duration `protobuf:"bytes,3,opt,name=op_timeout,json=opTimeout,stdduration" json:"op_timeout,omitempty"`
}

func (m *NatsStreaming) Reset()                    { *m = NatsStreaming{} }
func (m *NatsStreaming) String() string            { return proto.CompactTextString(m) }
func (*NatsStreaming) ProtoMessage()               {}
func (*NatsStreaming) Descriptor() ([]byte, []int) { return fileDescriptorNatsStreamingFilter, []int{0} }

func (m *NatsStreaming) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *NatsStreaming) GetMaxConnections() uint32 {
	if m != nil {
		return m.MaxConnections
	}
	return 0
}

func (m *NatsStreaming) GetOpTimeout() *time.Duration {
	if m != nil {
		return m.OpTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*NatsStreaming)(nil), "nats.NatsStreaming")
}

func init() { proto.RegisterFile("nats_streaming_filter.proto", fileDescriptorNatsStreamingFilter) }

var fileDescriptorNatsStreamingFilter = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x4b, 0x2c, 0x29,
	0x8e, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xcd, 0xcc, 0x4b, 0x8f, 0x4f, 0xcb, 0xcc, 0x29, 0x49,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x49, 0x4a, 0xc9, 0xa5, 0xe7, 0xe7,
	0xa7, 0xe7, 0xa4, 0xea, 0x83, 0xc5, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x4a, 0x8b, 0x12, 0x4b, 0x32,
	0xf3, 0xf3, 0x20, 0xaa, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4c, 0x7d, 0x10, 0x0b, 0x22,
	0xaa, 0x34, 0x89, 0x91, 0x8b, 0xd7, 0x2f, 0xb1, 0xa4, 0x38, 0x18, 0x66, 0xb4, 0x90, 0x04, 0x17,
	0x7b, 0x72, 0x4e, 0x69, 0x71, 0x49, 0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c,
	0x2b, 0xa4, 0xce, 0xc5, 0x9f, 0x9b, 0x58, 0x11, 0x9f, 0x9c, 0x9f, 0x97, 0x97, 0x9a, 0x0c, 0x32,
	0xb9, 0x58, 0x82, 0x49, 0x81, 0x51, 0x83, 0x37, 0x88, 0x2f, 0x37, 0xb1, 0xc2, 0x19, 0x21, 0x2a,
	0x64, 0xc7, 0xc5, 0x95, 0x5f, 0x10, 0x5f, 0x92, 0x99, 0x9b, 0x9a, 0x5f, 0x5a, 0x22, 0xc1, 0xac,
	0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa9, 0x07, 0x71, 0x9f, 0x1e, 0xcc, 0x7d, 0x7a, 0x2e, 0x50, 0xf7,
	0x39, 0xb1, 0xcc, 0xb8, 0x2f, 0xcf, 0x18, 0xc4, 0x99, 0x5f, 0x10, 0x02, 0xd1, 0x91, 0xc4, 0x06,
	0x56, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xd6, 0x8a, 0x22, 0xf6, 0x00, 0x00, 0x00,
}