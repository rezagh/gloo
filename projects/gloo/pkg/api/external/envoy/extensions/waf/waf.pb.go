// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/waf/waf.proto

package waf

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ModSecurity struct {
	// Disable all rules on the current route
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Global rule sets for the current http connection manager
	RuleSets             []*RuleSet `protobuf:"bytes,2,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ModSecurity) Reset()         { *m = ModSecurity{} }
func (m *ModSecurity) String() string { return proto.CompactTextString(m) }
func (*ModSecurity) ProtoMessage()    {}
func (*ModSecurity) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf8967e7dbe03c2, []int{0}
}
func (m *ModSecurity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModSecurity.Unmarshal(m, b)
}
func (m *ModSecurity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModSecurity.Marshal(b, m, deterministic)
}
func (m *ModSecurity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModSecurity.Merge(m, src)
}
func (m *ModSecurity) XXX_Size() int {
	return xxx_messageInfo_ModSecurity.Size(m)
}
func (m *ModSecurity) XXX_DiscardUnknown() {
	xxx_messageInfo_ModSecurity.DiscardUnknown(m)
}

var xxx_messageInfo_ModSecurity proto.InternalMessageInfo

func (m *ModSecurity) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *ModSecurity) GetRuleSets() []*RuleSet {
	if m != nil {
		return m.RuleSets
	}
	return nil
}

//
//String options are not recommended unless they are relatively short as they will be sent over the wire quite often.
//
//Any files referenced by this proto should be mounted into the relevant envoy pod prior to use or
//the filter will fail to initialize and the configuration will be rejected
type RuleSet struct {
	// string of rules which are added directly
	RuleStr string `protobuf:"bytes,1,opt,name=rule_str,json=ruleStr,proto3" json:"rule_str,omitempty"`
	// array of files to include
	Files                []string `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuleSet) Reset()         { *m = RuleSet{} }
func (m *RuleSet) String() string { return proto.CompactTextString(m) }
func (*RuleSet) ProtoMessage()    {}
func (*RuleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf8967e7dbe03c2, []int{1}
}
func (m *RuleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleSet.Unmarshal(m, b)
}
func (m *RuleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleSet.Marshal(b, m, deterministic)
}
func (m *RuleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleSet.Merge(m, src)
}
func (m *RuleSet) XXX_Size() int {
	return xxx_messageInfo_RuleSet.Size(m)
}
func (m *RuleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleSet.DiscardUnknown(m)
}

var xxx_messageInfo_RuleSet proto.InternalMessageInfo

func (m *RuleSet) GetRuleStr() string {
	if m != nil {
		return m.RuleStr
	}
	return ""
}

func (m *RuleSet) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

type ModSecurityPerRoute struct {
	// Disable all rules on the current route
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Overwite the global rules on this route
	RuleSets             []*RuleSet `protobuf:"bytes,2,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ModSecurityPerRoute) Reset()         { *m = ModSecurityPerRoute{} }
func (m *ModSecurityPerRoute) String() string { return proto.CompactTextString(m) }
func (*ModSecurityPerRoute) ProtoMessage()    {}
func (*ModSecurityPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf8967e7dbe03c2, []int{2}
}
func (m *ModSecurityPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModSecurityPerRoute.Unmarshal(m, b)
}
func (m *ModSecurityPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModSecurityPerRoute.Marshal(b, m, deterministic)
}
func (m *ModSecurityPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModSecurityPerRoute.Merge(m, src)
}
func (m *ModSecurityPerRoute) XXX_Size() int {
	return xxx_messageInfo_ModSecurityPerRoute.Size(m)
}
func (m *ModSecurityPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ModSecurityPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ModSecurityPerRoute proto.InternalMessageInfo

func (m *ModSecurityPerRoute) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *ModSecurityPerRoute) GetRuleSets() []*RuleSet {
	if m != nil {
		return m.RuleSets
	}
	return nil
}

func init() {
	proto.RegisterType((*ModSecurity)(nil), "envoy.config.filter.http.modsecurity.v2.ModSecurity")
	proto.RegisterType((*RuleSet)(nil), "envoy.config.filter.http.modsecurity.v2.RuleSet")
	proto.RegisterType((*ModSecurityPerRoute)(nil), "envoy.config.filter.http.modsecurity.v2.ModSecurityPerRoute")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/waf/waf.proto", fileDescriptor_aaf8967e7dbe03c2)
}

var fileDescriptor_aaf8967e7dbe03c2 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x95, 0xaf, 0xfa, 0x68, 0xe2, 0x6e, 0xa1, 0x43, 0xe8, 0x80, 0xa2, 0x2c, 0x64, 0xc1,
	0x46, 0x65, 0x63, 0x64, 0x45, 0x95, 0xa8, 0xbb, 0xb1, 0xa0, 0xfc, 0xb9, 0x71, 0x0d, 0x6e, 0x6e,
	0x64, 0xdf, 0x94, 0x76, 0xe3, 0x71, 0x78, 0x2e, 0x9e, 0x04, 0xd5, 0xa9, 0x10, 0x1b, 0x9d, 0x18,
	0x2c, 0xdd, 0x9f, 0x7d, 0x8f, 0xcf, 0x91, 0x0e, 0x5b, 0x2a, 0x4d, 0xeb, 0xbe, 0xe4, 0x15, 0x6e,
	0x84, 0x43, 0x83, 0xd7, 0x1a, 0x85, 0x32, 0x88, 0xa2, 0xb3, 0xf8, 0x02, 0x15, 0xb9, 0x81, 0x8a,
	0x4e, 0x0b, 0xd8, 0x11, 0xd8, 0xb6, 0x30, 0x02, 0xda, 0x2d, 0xee, 0x3d, 0xb6, 0x4e, 0x63, 0xeb,
	0xc4, 0x5b, 0xd1, 0x1c, 0x0e, 0xef, 0x2c, 0x12, 0xc6, 0x57, 0xfe, 0x9d, 0x57, 0xd8, 0x36, 0x5a,
	0xf1, 0x46, 0x1b, 0x02, 0xcb, 0xd7, 0x44, 0x1d, 0xdf, 0x60, 0xed, 0xa0, 0xea, 0xad, 0xa6, 0x3d,
	0xdf, 0xce, 0x67, 0x53, 0x85, 0x0a, 0xbd, 0x46, 0x1c, 0xa6, 0x41, 0x9e, 0xed, 0xd8, 0x64, 0x81,
	0xf5, 0xea, 0xb8, 0x17, 0xcf, 0x58, 0x58, 0x6b, 0x57, 0x94, 0x06, 0xea, 0x24, 0x48, 0x83, 0x3c,
	0x94, 0xdf, 0x1c, 0x2f, 0x58, 0x64, 0x7b, 0x03, 0xcf, 0x0e, 0xc8, 0x25, 0xff, 0xd2, 0x51, 0x3e,
	0x99, 0xdf, 0xf0, 0x13, 0xdd, 0xb9, 0xec, 0x0d, 0xac, 0x80, 0x64, 0x68, 0x87, 0xc1, 0x65, 0x77,
	0x6c, 0x7c, 0xbc, 0x8c, 0x2f, 0x58, 0x38, 0xfc, 0x4c, 0xd6, 0xbb, 0x46, 0x72, 0xec, 0xd7, 0xc8,
	0xc6, 0x53, 0xf6, 0xbf, 0xd1, 0x06, 0x5c, 0x32, 0x4a, 0x47, 0x79, 0x24, 0x07, 0xc8, 0xde, 0x03,
	0x76, 0xfe, 0x23, 0xf6, 0x23, 0x58, 0x89, 0x3d, 0xc1, 0x1f, 0xc6, 0xbf, 0x5f, 0x7e, 0x7c, 0x5e,
	0x06, 0x4f, 0x0f, 0xa7, 0x15, 0xda, 0xbd, 0xaa, 0xdf, 0x4b, 0x2d, 0xcf, 0x7c, 0x25, 0xb7, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x9c, 0xbc, 0xae, 0x26, 0x02, 0x00, 0x00,
}

func (this *ModSecurity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ModSecurity)
	if !ok {
		that2, ok := that.(ModSecurity)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Disabled != that1.Disabled {
		return false
	}
	if len(this.RuleSets) != len(that1.RuleSets) {
		return false
	}
	for i := range this.RuleSets {
		if !this.RuleSets[i].Equal(that1.RuleSets[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RuleSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RuleSet)
	if !ok {
		that2, ok := that.(RuleSet)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.RuleStr != that1.RuleStr {
		return false
	}
	if len(this.Files) != len(that1.Files) {
		return false
	}
	for i := range this.Files {
		if this.Files[i] != that1.Files[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ModSecurityPerRoute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ModSecurityPerRoute)
	if !ok {
		that2, ok := that.(ModSecurityPerRoute)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Disabled != that1.Disabled {
		return false
	}
	if len(this.RuleSets) != len(that1.RuleSets) {
		return false
	}
	for i := range this.RuleSets {
		if !this.RuleSets[i].Equal(that1.RuleSets[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}