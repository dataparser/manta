// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dota_modifiers.proto

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

type DOTA_MODIFIER_ENTRY_TYPE int32

const (
	DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_ACTIVE  DOTA_MODIFIER_ENTRY_TYPE = 1
	DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_REMOVED DOTA_MODIFIER_ENTRY_TYPE = 2
)

var DOTA_MODIFIER_ENTRY_TYPE_name = map[int32]string{
	1: "DOTA_MODIFIER_ENTRY_TYPE_ACTIVE",
	2: "DOTA_MODIFIER_ENTRY_TYPE_REMOVED",
}

var DOTA_MODIFIER_ENTRY_TYPE_value = map[string]int32{
	"DOTA_MODIFIER_ENTRY_TYPE_ACTIVE":  1,
	"DOTA_MODIFIER_ENTRY_TYPE_REMOVED": 2,
}

func (x DOTA_MODIFIER_ENTRY_TYPE) Enum() *DOTA_MODIFIER_ENTRY_TYPE {
	p := new(DOTA_MODIFIER_ENTRY_TYPE)
	*p = x
	return p
}

func (x DOTA_MODIFIER_ENTRY_TYPE) String() string {
	return proto.EnumName(DOTA_MODIFIER_ENTRY_TYPE_name, int32(x))
}

func (x *DOTA_MODIFIER_ENTRY_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DOTA_MODIFIER_ENTRY_TYPE_value, data, "DOTA_MODIFIER_ENTRY_TYPE")
	if err != nil {
		return err
	}
	*x = DOTA_MODIFIER_ENTRY_TYPE(value)
	return nil
}

func (DOTA_MODIFIER_ENTRY_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1f582eabb1c63b44, []int{0}
}

type CDOTAModifierBuffTableEntry struct {
	EntryType            *DOTA_MODIFIER_ENTRY_TYPE `protobuf:"varint,1,req,name=entry_type,json=entryType,enum=dota.DOTA_MODIFIER_ENTRY_TYPE,def=1" json:"entry_type,omitempty"`
	Parent               *int32                    `protobuf:"varint,2,req,name=parent" json:"parent,omitempty"`
	Index                *int32                    `protobuf:"varint,3,req,name=index" json:"index,omitempty"`
	SerialNum            *int32                    `protobuf:"varint,4,req,name=serial_num,json=serialNum" json:"serial_num,omitempty"`
	ModifierClass        *int32                    `protobuf:"varint,5,opt,name=modifier_class,json=modifierClass" json:"modifier_class,omitempty"`
	AbilityLevel         *int32                    `protobuf:"varint,6,opt,name=ability_level,json=abilityLevel" json:"ability_level,omitempty"`
	StackCount           *int32                    `protobuf:"varint,7,opt,name=stack_count,json=stackCount" json:"stack_count,omitempty"`
	CreationTime         *float32                  `protobuf:"fixed32,8,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	Duration             *float32                  `protobuf:"fixed32,9,opt,name=duration,def=-1" json:"duration,omitempty"`
	Caster               *int32                    `protobuf:"varint,10,opt,name=caster" json:"caster,omitempty"`
	Ability              *int32                    `protobuf:"varint,11,opt,name=ability" json:"ability,omitempty"`
	Armor                *int32                    `protobuf:"varint,12,opt,name=armor" json:"armor,omitempty"`
	FadeTime             *float32                  `protobuf:"fixed32,13,opt,name=fade_time,json=fadeTime" json:"fade_time,omitempty"`
	Subtle               *bool                     `protobuf:"varint,14,opt,name=subtle" json:"subtle,omitempty"`
	ChannelTime          *float32                  `protobuf:"fixed32,15,opt,name=channel_time,json=channelTime" json:"channel_time,omitempty"`
	VStart               *CMsgVector               `protobuf:"bytes,16,opt,name=v_start,json=vStart" json:"v_start,omitempty"`
	VEnd                 *CMsgVector               `protobuf:"bytes,17,opt,name=v_end,json=vEnd" json:"v_end,omitempty"`
	PortalLoopAppear     *string                   `protobuf:"bytes,18,opt,name=portal_loop_appear,json=portalLoopAppear" json:"portal_loop_appear,omitempty"`
	PortalLoopDisappear  *string                   `protobuf:"bytes,19,opt,name=portal_loop_disappear,json=portalLoopDisappear" json:"portal_loop_disappear,omitempty"`
	HeroLoopAppear       *string                   `protobuf:"bytes,20,opt,name=hero_loop_appear,json=heroLoopAppear" json:"hero_loop_appear,omitempty"`
	HeroLoopDisappear    *string                   `protobuf:"bytes,21,opt,name=hero_loop_disappear,json=heroLoopDisappear" json:"hero_loop_disappear,omitempty"`
	MovementSpeed        *int32                    `protobuf:"varint,22,opt,name=movement_speed,json=movementSpeed" json:"movement_speed,omitempty"`
	Aura                 *bool                     `protobuf:"varint,23,opt,name=aura" json:"aura,omitempty"`
	Activity             *int32                    `protobuf:"varint,24,opt,name=activity" json:"activity,omitempty"`
	Damage               *int32                    `protobuf:"varint,25,opt,name=damage" json:"damage,omitempty"`
	Range                *int32                    `protobuf:"varint,26,opt,name=range" json:"range,omitempty"`
	DdModifierIndex      *int32                    `protobuf:"varint,27,opt,name=dd_modifier_index,json=ddModifierIndex" json:"dd_modifier_index,omitempty"`
	DdAbilityId          *int32                    `protobuf:"varint,28,opt,name=dd_ability_id,json=ddAbilityId" json:"dd_ability_id,omitempty"`
	IllusionLabel        *string                   `protobuf:"bytes,29,opt,name=illusion_label,json=illusionLabel" json:"illusion_label,omitempty"`
	Active               *bool                     `protobuf:"varint,30,opt,name=active" json:"active,omitempty"`
	PlayerIds            *string                   `protobuf:"bytes,31,opt,name=player_ids,json=playerIds" json:"player_ids,omitempty"`
	LuaName              *string                   `protobuf:"bytes,32,opt,name=lua_name,json=luaName" json:"lua_name,omitempty"`
	AttackSpeed          *int32                    `protobuf:"varint,33,opt,name=attack_speed,json=attackSpeed" json:"attack_speed,omitempty"`
	AuraOwner            *int32                    `protobuf:"varint,34,opt,name=aura_owner,json=auraOwner" json:"aura_owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CDOTAModifierBuffTableEntry) Reset()         { *m = CDOTAModifierBuffTableEntry{} }
func (m *CDOTAModifierBuffTableEntry) String() string { return proto.CompactTextString(m) }
func (*CDOTAModifierBuffTableEntry) ProtoMessage()    {}
func (*CDOTAModifierBuffTableEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f582eabb1c63b44, []int{0}
}

func (m *CDOTAModifierBuffTableEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDOTAModifierBuffTableEntry.Unmarshal(m, b)
}
func (m *CDOTAModifierBuffTableEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDOTAModifierBuffTableEntry.Marshal(b, m, deterministic)
}
func (m *CDOTAModifierBuffTableEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDOTAModifierBuffTableEntry.Merge(m, src)
}
func (m *CDOTAModifierBuffTableEntry) XXX_Size() int {
	return xxx_messageInfo_CDOTAModifierBuffTableEntry.Size(m)
}
func (m *CDOTAModifierBuffTableEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CDOTAModifierBuffTableEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CDOTAModifierBuffTableEntry proto.InternalMessageInfo

const Default_CDOTAModifierBuffTableEntry_EntryType DOTA_MODIFIER_ENTRY_TYPE = DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_ACTIVE
const Default_CDOTAModifierBuffTableEntry_Duration float32 = -1

func (m *CDOTAModifierBuffTableEntry) GetEntryType() DOTA_MODIFIER_ENTRY_TYPE {
	if m != nil && m.EntryType != nil {
		return *m.EntryType
	}
	return Default_CDOTAModifierBuffTableEntry_EntryType
}

func (m *CDOTAModifierBuffTableEntry) GetParent() int32 {
	if m != nil && m.Parent != nil {
		return *m.Parent
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetSerialNum() int32 {
	if m != nil && m.SerialNum != nil {
		return *m.SerialNum
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetModifierClass() int32 {
	if m != nil && m.ModifierClass != nil {
		return *m.ModifierClass
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetAbilityLevel() int32 {
	if m != nil && m.AbilityLevel != nil {
		return *m.AbilityLevel
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetStackCount() int32 {
	if m != nil && m.StackCount != nil {
		return *m.StackCount
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetCreationTime() float32 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDuration() float32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return Default_CDOTAModifierBuffTableEntry_Duration
}

func (m *CDOTAModifierBuffTableEntry) GetCaster() int32 {
	if m != nil && m.Caster != nil {
		return *m.Caster
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetAbility() int32 {
	if m != nil && m.Ability != nil {
		return *m.Ability
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetArmor() int32 {
	if m != nil && m.Armor != nil {
		return *m.Armor
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetFadeTime() float32 {
	if m != nil && m.FadeTime != nil {
		return *m.FadeTime
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetSubtle() bool {
	if m != nil && m.Subtle != nil {
		return *m.Subtle
	}
	return false
}

func (m *CDOTAModifierBuffTableEntry) GetChannelTime() float32 {
	if m != nil && m.ChannelTime != nil {
		return *m.ChannelTime
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetVStart() *CMsgVector {
	if m != nil {
		return m.VStart
	}
	return nil
}

func (m *CDOTAModifierBuffTableEntry) GetVEnd() *CMsgVector {
	if m != nil {
		return m.VEnd
	}
	return nil
}

func (m *CDOTAModifierBuffTableEntry) GetPortalLoopAppear() string {
	if m != nil && m.PortalLoopAppear != nil {
		return *m.PortalLoopAppear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetPortalLoopDisappear() string {
	if m != nil && m.PortalLoopDisappear != nil {
		return *m.PortalLoopDisappear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetHeroLoopAppear() string {
	if m != nil && m.HeroLoopAppear != nil {
		return *m.HeroLoopAppear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetHeroLoopDisappear() string {
	if m != nil && m.HeroLoopDisappear != nil {
		return *m.HeroLoopDisappear
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetMovementSpeed() int32 {
	if m != nil && m.MovementSpeed != nil {
		return *m.MovementSpeed
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetAura() bool {
	if m != nil && m.Aura != nil {
		return *m.Aura
	}
	return false
}

func (m *CDOTAModifierBuffTableEntry) GetActivity() int32 {
	if m != nil && m.Activity != nil {
		return *m.Activity
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDamage() int32 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetRange() int32 {
	if m != nil && m.Range != nil {
		return *m.Range
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDdModifierIndex() int32 {
	if m != nil && m.DdModifierIndex != nil {
		return *m.DdModifierIndex
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetDdAbilityId() int32 {
	if m != nil && m.DdAbilityId != nil {
		return *m.DdAbilityId
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetIllusionLabel() string {
	if m != nil && m.IllusionLabel != nil {
		return *m.IllusionLabel
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetActive() bool {
	if m != nil && m.Active != nil {
		return *m.Active
	}
	return false
}

func (m *CDOTAModifierBuffTableEntry) GetPlayerIds() string {
	if m != nil && m.PlayerIds != nil {
		return *m.PlayerIds
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetLuaName() string {
	if m != nil && m.LuaName != nil {
		return *m.LuaName
	}
	return ""
}

func (m *CDOTAModifierBuffTableEntry) GetAttackSpeed() int32 {
	if m != nil && m.AttackSpeed != nil {
		return *m.AttackSpeed
	}
	return 0
}

func (m *CDOTAModifierBuffTableEntry) GetAuraOwner() int32 {
	if m != nil && m.AuraOwner != nil {
		return *m.AuraOwner
	}
	return 0
}

type CDOTALuaModifierEntry struct {
	ModifierType         *int32   `protobuf:"varint,1,req,name=modifier_type,json=modifierType" json:"modifier_type,omitempty"`
	ModifierFilename     *string  `protobuf:"bytes,2,req,name=modifier_filename,json=modifierFilename" json:"modifier_filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CDOTALuaModifierEntry) Reset()         { *m = CDOTALuaModifierEntry{} }
func (m *CDOTALuaModifierEntry) String() string { return proto.CompactTextString(m) }
func (*CDOTALuaModifierEntry) ProtoMessage()    {}
func (*CDOTALuaModifierEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f582eabb1c63b44, []int{1}
}

func (m *CDOTALuaModifierEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDOTALuaModifierEntry.Unmarshal(m, b)
}
func (m *CDOTALuaModifierEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDOTALuaModifierEntry.Marshal(b, m, deterministic)
}
func (m *CDOTALuaModifierEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDOTALuaModifierEntry.Merge(m, src)
}
func (m *CDOTALuaModifierEntry) XXX_Size() int {
	return xxx_messageInfo_CDOTALuaModifierEntry.Size(m)
}
func (m *CDOTALuaModifierEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CDOTALuaModifierEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CDOTALuaModifierEntry proto.InternalMessageInfo

func (m *CDOTALuaModifierEntry) GetModifierType() int32 {
	if m != nil && m.ModifierType != nil {
		return *m.ModifierType
	}
	return 0
}

func (m *CDOTALuaModifierEntry) GetModifierFilename() string {
	if m != nil && m.ModifierFilename != nil {
		return *m.ModifierFilename
	}
	return ""
}

func init() {
	proto.RegisterEnum("dota.DOTA_MODIFIER_ENTRY_TYPE", DOTA_MODIFIER_ENTRY_TYPE_name, DOTA_MODIFIER_ENTRY_TYPE_value)
	proto.RegisterType((*CDOTAModifierBuffTableEntry)(nil), "dota.CDOTAModifierBuffTableEntry")
	proto.RegisterType((*CDOTALuaModifierEntry)(nil), "dota.CDOTALuaModifierEntry")
}

func init() { proto.RegisterFile("dota_modifiers.proto", fileDescriptor_1f582eabb1c63b44) }

var fileDescriptor_1f582eabb1c63b44 = []byte{
	// 797 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4b, 0x6f, 0x1b, 0x37,
	0x10, 0xc7, 0xbb, 0xaa, 0x65, 0x6b, 0xc7, 0x92, 0x23, 0xd3, 0x8f, 0x32, 0x76, 0x6d, 0x6f, 0xec,
	0x06, 0x50, 0xd3, 0xd6, 0x40, 0x7d, 0xcc, 0xcd, 0x91, 0x15, 0x54, 0x80, 0x1f, 0xc5, 0x46, 0x30,
	0x90, 0x4b, 0x09, 0x4a, 0x1c, 0x39, 0x8b, 0x70, 0x1f, 0xe0, 0x72, 0x95, 0xea, 0xd6, 0xcf, 0xd5,
	0x4f, 0x57, 0x70, 0xb8, 0x2b, 0xa5, 0x07, 0x23, 0xb7, 0x9d, 0xdf, 0xfc, 0x67, 0x76, 0x48, 0xce,
	0x0c, 0xec, 0xab, 0xdc, 0x4a, 0x91, 0xe6, 0x2a, 0x99, 0x27, 0x68, 0xca, 0xcb, 0xc2, 0xe4, 0x36,
	0x67, 0x1b, 0x8e, 0x1e, 0x1d, 0x66, 0x68, 0xbf, 0xe4, 0xe6, 0xf3, 0x54, 0x96, 0x68, 0x97, 0x05,
	0xd6, 0xde, 0xf3, 0x7f, 0x43, 0x38, 0x1e, 0xde, 0x3c, 0x4c, 0xae, 0xef, 0xea, 0xb0, 0x77, 0xd5,
	0x7c, 0x3e, 0x91, 0x53, 0x8d, 0xa3, 0xcc, 0x9a, 0x25, 0xfb, 0x0b, 0x00, 0xdd, 0x87, 0x70, 0x41,
	0x3c, 0x88, 0x5a, 0x83, 0x9d, 0xab, 0xd3, 0x4b, 0x97, 0xf2, 0xd2, 0x45, 0x89, 0xbb, 0x87, 0x9b,
	0xf1, 0xfb, 0xf1, 0x28, 0x16, 0xa3, 0xfb, 0x49, 0xfc, 0x51, 0x4c, 0x3e, 0xfe, 0x39, 0x7a, 0x7b,
	0xf6, 0x9c, 0x47, 0x5c, 0x0f, 0x27, 0xe3, 0xc7, 0x51, 0x1c, 0x52, 0xca, 0xc9, 0xb2, 0x40, 0x76,
	0x08, 0x9b, 0x85, 0x34, 0x98, 0x59, 0xde, 0x8a, 0x5a, 0x83, 0x76, 0x5c, 0x5b, 0x6c, 0x1f, 0xda,
	0x49, 0xa6, 0xf0, 0x6f, 0xfe, 0x3d, 0x61, 0x6f, 0xb0, 0x13, 0x80, 0x12, 0x4d, 0x22, 0xb5, 0xc8,
	0xaa, 0x94, 0x6f, 0x90, 0x2b, 0xf4, 0xe4, 0xbe, 0x4a, 0xd9, 0x6b, 0xd8, 0x69, 0x4e, 0x2f, 0x66,
	0x5a, 0x96, 0x25, 0x6f, 0x47, 0xc1, 0xa0, 0x1d, 0xf7, 0x1a, 0x3a, 0x74, 0x90, 0x5d, 0x40, 0x4f,
	0x4e, 0x13, 0x9d, 0xd8, 0xa5, 0xd0, 0xb8, 0x40, 0xcd, 0x37, 0x49, 0xd5, 0xad, 0xe1, 0xad, 0x63,
	0xec, 0x0c, 0xb6, 0x4b, 0x2b, 0x67, 0x9f, 0xc5, 0x2c, 0xaf, 0x32, 0xcb, 0xb7, 0x48, 0x02, 0x84,
	0x86, 0x8e, 0xb8, 0x2c, 0x33, 0x83, 0xd2, 0x26, 0x79, 0x26, 0x6c, 0x92, 0x22, 0xef, 0x44, 0xc1,
	0xa0, 0x15, 0x77, 0x1b, 0x38, 0x49, 0x52, 0x64, 0xa7, 0xd0, 0x51, 0x95, 0x21, 0x9b, 0x87, 0xce,
	0xff, 0xb6, 0xf5, 0xdb, 0xef, 0xf1, 0x8a, 0xb9, 0xe3, 0xcf, 0x64, 0x69, 0xd1, 0x70, 0xa0, 0x1f,
	0xd4, 0x16, 0xe3, 0xb0, 0x55, 0x57, 0xc3, 0xb7, 0xc9, 0xd1, 0x98, 0xee, 0x62, 0xa4, 0x49, 0x73,
	0xc3, 0xbb, 0xc4, 0xbd, 0xc1, 0x8e, 0x21, 0x9c, 0x4b, 0x85, 0xbe, 0x90, 0x1e, 0x15, 0xd2, 0x71,
	0x80, 0x8a, 0x38, 0x84, 0xcd, 0xb2, 0x9a, 0x5a, 0x8d, 0x7c, 0x27, 0x0a, 0x06, 0x9d, 0xb8, 0xb6,
	0xd8, 0x2b, 0xe8, 0xce, 0x3e, 0xc9, 0x2c, 0x43, 0xed, 0xe3, 0x5e, 0x50, 0xdc, 0x76, 0xcd, 0x28,
	0xf4, 0x67, 0xd8, 0x5a, 0x88, 0xd2, 0x4a, 0x63, 0x79, 0x3f, 0x0a, 0x06, 0xdb, 0x57, 0x7d, 0xff,
	0xf6, 0xc3, 0xbb, 0xf2, 0xe9, 0x11, 0x67, 0x36, 0x37, 0xf1, 0xe6, 0xe2, 0x83, 0xf3, 0xb3, 0xd7,
	0xd0, 0x5e, 0x08, 0xcc, 0x14, 0xdf, 0x7d, 0x46, 0xb8, 0xb1, 0x18, 0x65, 0x8a, 0xfd, 0x0a, 0xac,
	0xc8, 0x8d, 0x95, 0x5a, 0xe8, 0x3c, 0x2f, 0x84, 0x2c, 0x0a, 0x94, 0x86, 0xb3, 0x28, 0x18, 0x84,
	0x71, 0xdf, 0x7b, 0x6e, 0xf3, 0xbc, 0xb8, 0x26, 0xce, 0xae, 0xe0, 0xe0, 0x6b, 0xb5, 0x4a, 0xca,
	0x3a, 0x60, 0x8f, 0x02, 0xf6, 0xd6, 0x01, 0x37, 0x8d, 0x8b, 0x0d, 0xa0, 0xff, 0x09, 0x4d, 0xfe,
	0xbf, 0xfc, 0xfb, 0x24, 0xdf, 0x71, 0xfc, 0xab, 0xec, 0x97, 0xb0, 0xb7, 0x56, 0xae, 0x73, 0x1f,
	0x90, 0x78, 0xb7, 0x11, 0xaf, 0x33, 0x53, 0x7f, 0x2d, 0x30, 0xc5, 0xcc, 0x8a, 0xb2, 0x40, 0x54,
	0xfc, 0xb0, 0xe9, 0x2f, 0x4f, 0x3f, 0x38, 0xc8, 0x18, 0x6c, 0xc8, 0xca, 0x48, 0xfe, 0x03, 0xdd,
	0x36, 0x7d, 0xb3, 0x23, 0xe8, 0xc8, 0x99, 0x4d, 0x16, 0xee, 0x45, 0x39, 0x05, 0xad, 0x6c, 0xf7,
	0x3e, 0x4a, 0xa6, 0xf2, 0x09, 0xf9, 0x4b, 0xdf, 0x04, 0xde, 0x72, 0x4f, 0x6d, 0x64, 0xf6, 0x84,
	0xfc, 0xc8, 0x3f, 0x35, 0x19, 0xec, 0x0d, 0xec, 0x2a, 0xb5, 0x9a, 0x72, 0xe1, 0xa7, 0xe4, 0x98,
	0x14, 0x2f, 0x94, 0x6a, 0xc6, 0x78, 0x4c, 0xf3, 0x72, 0x0e, 0x3d, 0xa5, 0x44, 0xd3, 0xec, 0x89,
	0xe2, 0x3f, 0x92, 0x6e, 0x5b, 0xa9, 0x6b, 0xcf, 0xc6, 0xca, 0x1d, 0x2a, 0xd1, 0xba, 0x2a, 0x5d,
	0x1f, 0x6b, 0x39, 0x45, 0xcd, 0x4f, 0xe8, 0xfc, 0xbd, 0x86, 0xde, 0x3a, 0xe8, 0x8a, 0xa4, 0x82,
	0x91, 0x9f, 0xfa, 0x26, 0xf2, 0x96, 0x1b, 0xc9, 0x42, 0xcb, 0xa5, 0xab, 0x44, 0x95, 0xfc, 0x8c,
	0x42, 0x43, 0x4f, 0xc6, 0xaa, 0x64, 0x2f, 0xa1, 0xa3, 0x2b, 0x29, 0x32, 0x99, 0x22, 0x8f, 0xc8,
	0xb9, 0xa5, 0x2b, 0x79, 0x2f, 0x53, 0x6a, 0x3f, 0x69, 0x69, 0xc4, 0xfc, 0x5d, 0xbe, 0xf2, 0xb5,
	0x79, 0xe6, 0x6f, 0xf2, 0x04, 0xc0, 0xdd, 0x9e, 0xc8, 0xbf, 0x64, 0x68, 0xf8, 0x39, 0x09, 0x42,
	0x47, 0x1e, 0x1c, 0x38, 0x4f, 0xe0, 0x80, 0x76, 0xd7, 0x6d, 0x25, 0x9b, 0x73, 0xfb, 0xad, 0x75,
	0x01, 0xab, 0x91, 0x5f, 0x2f, 0xae, 0x76, 0xdc, 0x6d, 0x20, 0xad, 0x9e, 0x5f, 0x60, 0x77, 0x25,
	0x9a, 0x27, 0x1a, 0xa9, 0x46, 0xb7, 0x85, 0xc2, 0xb8, 0xdf, 0x38, 0xde, 0xd7, 0xfc, 0x0d, 0x02,
	0x7f, 0x6e, 0xab, 0xb1, 0x0b, 0xf8, 0xd6, 0xc6, 0xeb, 0x07, 0xec, 0x27, 0x88, 0x9e, 0x15, 0xc5,
	0xa3, 0xbb, 0x87, 0xc7, 0xd1, 0x4d, 0xbf, 0xf5, 0xae, 0xfd, 0x47, 0xf0, 0x4f, 0xf0, 0xdd, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0x93, 0x28, 0x25, 0xca, 0x05, 0x00, 0x00,
}
