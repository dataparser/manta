// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connectionless_netmessages.proto

package dota

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

type C2S_CONNECT_Message struct {
	HostVersion          *uint32                       `protobuf:"varint,1,opt,name=host_version,json=hostVersion" json:"host_version,omitempty"`
	AuthProtocol         *uint32                       `protobuf:"varint,2,opt,name=auth_protocol,json=authProtocol" json:"auth_protocol,omitempty"`
	ChallengeNumber      *uint32                       `protobuf:"varint,3,opt,name=challenge_number,json=challengeNumber" json:"challenge_number,omitempty"`
	ReservationCookie    *uint64                       `protobuf:"fixed64,4,opt,name=reservation_cookie,json=reservationCookie" json:"reservation_cookie,omitempty"`
	LowViolence          *bool                         `protobuf:"varint,5,opt,name=low_violence,json=lowViolence" json:"low_violence,omitempty"`
	EncryptedPassword    []byte                        `protobuf:"bytes,6,opt,name=encrypted_password,json=encryptedPassword" json:"encrypted_password,omitempty"`
	Splitplayers         []*CCLCMsg_SplitPlayerConnect `protobuf:"bytes,7,rep,name=splitplayers" json:"splitplayers,omitempty"`
	AuthSteam            []byte                        `protobuf:"bytes,8,opt,name=auth_steam,json=authSteam" json:"auth_steam,omitempty"`
	ChallengeContext     *string                       `protobuf:"bytes,9,opt,name=challenge_context,json=challengeContext" json:"challenge_context,omitempty"`
	UseSnp               *int32                        `protobuf:"zigzag32,10,opt,name=use_snp,json=useSnp" json:"use_snp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *C2S_CONNECT_Message) Reset()         { *m = C2S_CONNECT_Message{} }
func (m *C2S_CONNECT_Message) String() string { return proto.CompactTextString(m) }
func (*C2S_CONNECT_Message) ProtoMessage()    {}
func (*C2S_CONNECT_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5f1aef3c0e41f21, []int{0}
}

func (m *C2S_CONNECT_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S_CONNECT_Message.Unmarshal(m, b)
}
func (m *C2S_CONNECT_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S_CONNECT_Message.Marshal(b, m, deterministic)
}
func (m *C2S_CONNECT_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_CONNECT_Message.Merge(m, src)
}
func (m *C2S_CONNECT_Message) XXX_Size() int {
	return xxx_messageInfo_C2S_CONNECT_Message.Size(m)
}
func (m *C2S_CONNECT_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_CONNECT_Message.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_CONNECT_Message proto.InternalMessageInfo

func (m *C2S_CONNECT_Message) GetHostVersion() uint32 {
	if m != nil && m.HostVersion != nil {
		return *m.HostVersion
	}
	return 0
}

func (m *C2S_CONNECT_Message) GetAuthProtocol() uint32 {
	if m != nil && m.AuthProtocol != nil {
		return *m.AuthProtocol
	}
	return 0
}

func (m *C2S_CONNECT_Message) GetChallengeNumber() uint32 {
	if m != nil && m.ChallengeNumber != nil {
		return *m.ChallengeNumber
	}
	return 0
}

func (m *C2S_CONNECT_Message) GetReservationCookie() uint64 {
	if m != nil && m.ReservationCookie != nil {
		return *m.ReservationCookie
	}
	return 0
}

func (m *C2S_CONNECT_Message) GetLowViolence() bool {
	if m != nil && m.LowViolence != nil {
		return *m.LowViolence
	}
	return false
}

func (m *C2S_CONNECT_Message) GetEncryptedPassword() []byte {
	if m != nil {
		return m.EncryptedPassword
	}
	return nil
}

func (m *C2S_CONNECT_Message) GetSplitplayers() []*CCLCMsg_SplitPlayerConnect {
	if m != nil {
		return m.Splitplayers
	}
	return nil
}

func (m *C2S_CONNECT_Message) GetAuthSteam() []byte {
	if m != nil {
		return m.AuthSteam
	}
	return nil
}

func (m *C2S_CONNECT_Message) GetChallengeContext() string {
	if m != nil && m.ChallengeContext != nil {
		return *m.ChallengeContext
	}
	return ""
}

func (m *C2S_CONNECT_Message) GetUseSnp() int32 {
	if m != nil && m.UseSnp != nil {
		return *m.UseSnp
	}
	return 0
}

type C2S_CONNECTION_Message struct {
	AddonName            *string  `protobuf:"bytes,1,opt,name=addon_name,json=addonName" json:"addon_name,omitempty"`
	UseSnp               *bool    `protobuf:"varint,2,opt,name=use_snp,json=useSnp" json:"use_snp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S_CONNECTION_Message) Reset()         { *m = C2S_CONNECTION_Message{} }
func (m *C2S_CONNECTION_Message) String() string { return proto.CompactTextString(m) }
func (*C2S_CONNECTION_Message) ProtoMessage()    {}
func (*C2S_CONNECTION_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5f1aef3c0e41f21, []int{1}
}

func (m *C2S_CONNECTION_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S_CONNECTION_Message.Unmarshal(m, b)
}
func (m *C2S_CONNECTION_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S_CONNECTION_Message.Marshal(b, m, deterministic)
}
func (m *C2S_CONNECTION_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_CONNECTION_Message.Merge(m, src)
}
func (m *C2S_CONNECTION_Message) XXX_Size() int {
	return xxx_messageInfo_C2S_CONNECTION_Message.Size(m)
}
func (m *C2S_CONNECTION_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_CONNECTION_Message.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_CONNECTION_Message proto.InternalMessageInfo

func (m *C2S_CONNECTION_Message) GetAddonName() string {
	if m != nil && m.AddonName != nil {
		return *m.AddonName
	}
	return ""
}

func (m *C2S_CONNECTION_Message) GetUseSnp() bool {
	if m != nil && m.UseSnp != nil {
		return *m.UseSnp
	}
	return false
}

func init() {
	proto.RegisterType((*C2S_CONNECT_Message)(nil), "dota.C2S_CONNECT_Message")
	proto.RegisterType((*C2S_CONNECTION_Message)(nil), "dota.C2S_CONNECTION_Message")
}

func init() {
	proto.RegisterFile("connectionless_netmessages.proto", fileDescriptor_c5f1aef3c0e41f21)
}

var fileDescriptor_c5f1aef3c0e41f21 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x51, 0x6f, 0xd3, 0x30,
	0x14, 0x85, 0xc9, 0x3a, 0xba, 0xc5, 0xed, 0x04, 0x31, 0x12, 0x58, 0x48, 0x48, 0x66, 0xbc, 0x04,
	0x21, 0xfa, 0xb0, 0x9f, 0x80, 0xe1, 0x01, 0x89, 0x65, 0x55, 0x8a, 0xf6, 0x6a, 0x99, 0xe4, 0xaa,
	0x8d, 0x70, 0x7c, 0x23, 0x5f, 0xb7, 0x65, 0x6f, 0xfc, 0x28, 0x7e, 0x20, 0xb2, 0x1b, 0x75, 0x45,
	0xbc, 0x7e, 0xe7, 0xe4, 0xdc, 0xf8, 0x1c, 0x26, 0x1b, 0x74, 0x0e, 0x9a, 0xd0, 0xa1, 0xb3, 0x40,
	0xa4, 0x1d, 0x84, 0x1e, 0x88, 0xcc, 0x1a, 0x68, 0x31, 0x78, 0x0c, 0xc8, 0xcf, 0x5b, 0x0c, 0xe6,
	0x75, 0xf1, 0x9f, 0x70, 0xfd, 0x67, 0xc2, 0x5e, 0xa8, 0x9b, 0x95, 0x56, 0x77, 0x55, 0xf5, 0x45,
	0x7d, 0xd7, 0xb7, 0x07, 0x99, 0xbf, 0x65, 0xf3, 0x0d, 0x52, 0xd0, 0x3b, 0xf0, 0xd4, 0xa1, 0x13,
	0x99, 0xcc, 0xca, 0xab, 0x7a, 0x16, 0xd9, 0xfd, 0x01, 0xf1, 0x77, 0xec, 0xca, 0x6c, 0xc3, 0x46,
	0xa7, 0xa0, 0x06, 0xad, 0x38, 0x4b, 0x9e, 0x79, 0x84, 0xcb, 0x91, 0xf1, 0xf7, 0xec, 0x79, 0xb3,
	0x31, 0xd6, 0x82, 0x5b, 0x83, 0x76, 0xdb, 0xfe, 0x07, 0x78, 0x31, 0x49, 0xbe, 0x67, 0x47, 0x5e,
	0x25, 0xcc, 0x3f, 0x32, 0xee, 0x81, 0xc0, 0xef, 0x4c, 0x7c, 0x88, 0x6e, 0x10, 0x7f, 0x76, 0x20,
	0xce, 0x65, 0x56, 0x4e, 0xeb, 0xe2, 0x44, 0x51, 0x49, 0x88, 0x7f, 0x68, 0x71, 0xaf, 0x77, 0x1d,
	0x5a, 0x70, 0x0d, 0x88, 0xa7, 0x32, 0x2b, 0x2f, 0xeb, 0x99, 0xc5, 0xfd, 0xfd, 0x88, 0x62, 0x22,
	0xb8, 0xc6, 0x3f, 0x0c, 0x01, 0x5a, 0x3d, 0x18, 0xa2, 0x3d, 0xfa, 0x56, 0x4c, 0x65, 0x56, 0xce,
	0xeb, 0xe2, 0xa8, 0x2c, 0x47, 0x81, 0x7f, 0x66, 0x73, 0x1a, 0x6c, 0x17, 0x06, 0x6b, 0x1e, 0xc0,
	0x93, 0xb8, 0x90, 0x93, 0x72, 0x76, 0x23, 0x17, 0xb1, 0xbb, 0x85, 0x52, 0xdf, 0xd4, 0x2d, 0xad,
	0xf5, 0x2a, 0x3a, 0x96, 0xc9, 0xa1, 0x0e, 0xad, 0xd7, 0xff, 0x7c, 0xc5, 0xdf, 0x30, 0x96, 0x6a,
	0xa1, 0x00, 0xa6, 0x17, 0x97, 0xe9, 0x58, 0x1e, 0xc9, 0x2a, 0x02, 0xfe, 0x81, 0x15, 0x8f, 0x85,
	0x34, 0xe8, 0x02, 0xfc, 0x0a, 0x22, 0x97, 0x59, 0x99, 0xd7, 0x8f, 0x4d, 0xa9, 0x03, 0xe7, 0xaf,
	0xd8, 0xc5, 0x96, 0x40, 0x93, 0x1b, 0x04, 0x93, 0x59, 0x59, 0xd4, 0xd3, 0x2d, 0xc1, 0xca, 0x0d,
	0xd7, 0x4b, 0xf6, 0xf2, 0x64, 0xb5, 0xaf, 0x77, 0xd5, 0x71, 0xb8, 0x78, 0xbe, 0x6d, 0xd1, 0x69,
	0x67, 0x7a, 0x48, 0xb3, 0xe5, 0x75, 0x9e, 0x48, 0x65, 0x7a, 0x38, 0x4d, 0x3c, 0x4b, 0x85, 0x8d,
	0x89, 0x9f, 0x26, 0xbf, 0xb3, 0x27, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x6c, 0xcf, 0xae,
	0x49, 0x02, 0x00, 0x00,
}
