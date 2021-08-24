// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package policer contains generated bindings for API file policer.api.
//
// Contents:
//   8 messages
//
package policer

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/projectcalico/vpp-dataplane/vpplink/binapi/vppapi/interface_types"
	policer_types "github.com/projectcalico/vpp-dataplane/vpplink/binapi/vppapi/policer_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "policer"
	APIVersion = "2.0.0"
	VersionCrc = 0xf14848c
)

// PolicerAddDel defines message 'policer_add_del'.
type PolicerAddDel struct {
	IsAdd         bool                             `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Name          string                           `binapi:"string[64],name=name" json:"name,omitempty"`
	Cir           uint32                           `binapi:"u32,name=cir" json:"cir,omitempty"`
	Eir           uint32                           `binapi:"u32,name=eir" json:"eir,omitempty"`
	Cb            uint64                           `binapi:"u64,name=cb" json:"cb,omitempty"`
	Eb            uint64                           `binapi:"u64,name=eb" json:"eb,omitempty"`
	RateType      policer_types.Sse2QosRateType    `binapi:"sse2_qos_rate_type,name=rate_type" json:"rate_type,omitempty"`
	RoundType     policer_types.Sse2QosRoundType   `binapi:"sse2_qos_round_type,name=round_type" json:"round_type,omitempty"`
	Type          policer_types.Sse2QosPolicerType `binapi:"sse2_qos_policer_type,name=type" json:"type,omitempty"`
	ColorAware    bool                             `binapi:"bool,name=color_aware" json:"color_aware,omitempty"`
	ConformAction policer_types.Sse2QosAction      `binapi:"sse2_qos_action,name=conform_action" json:"conform_action,omitempty"`
	ExceedAction  policer_types.Sse2QosAction      `binapi:"sse2_qos_action,name=exceed_action" json:"exceed_action,omitempty"`
	ViolateAction policer_types.Sse2QosAction      `binapi:"sse2_qos_action,name=violate_action" json:"violate_action,omitempty"`
}

func (m *PolicerAddDel) Reset()               { *m = PolicerAddDel{} }
func (*PolicerAddDel) GetMessageName() string { return "policer_add_del" }
func (*PolicerAddDel) GetCrcString() string   { return "2b31dd38" }
func (*PolicerAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PolicerAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1  // m.IsAdd
	size += 64 // m.Name
	size += 4  // m.Cir
	size += 4  // m.Eir
	size += 8  // m.Cb
	size += 8  // m.Eb
	size += 1  // m.RateType
	size += 1  // m.RoundType
	size += 1  // m.Type
	size += 1  // m.ColorAware
	size += 1  // m.ConformAction.Type
	size += 1  // m.ConformAction.Dscp
	size += 1  // m.ExceedAction.Type
	size += 1  // m.ExceedAction.Dscp
	size += 1  // m.ViolateAction.Type
	size += 1  // m.ViolateAction.Dscp
	return size
}
func (m *PolicerAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeString(m.Name, 64)
	buf.EncodeUint32(m.Cir)
	buf.EncodeUint32(m.Eir)
	buf.EncodeUint64(m.Cb)
	buf.EncodeUint64(m.Eb)
	buf.EncodeUint8(uint8(m.RateType))
	buf.EncodeUint8(uint8(m.RoundType))
	buf.EncodeUint8(uint8(m.Type))
	buf.EncodeBool(m.ColorAware)
	buf.EncodeUint8(uint8(m.ConformAction.Type))
	buf.EncodeUint8(m.ConformAction.Dscp)
	buf.EncodeUint8(uint8(m.ExceedAction.Type))
	buf.EncodeUint8(m.ExceedAction.Dscp)
	buf.EncodeUint8(uint8(m.ViolateAction.Type))
	buf.EncodeUint8(m.ViolateAction.Dscp)
	return buf.Bytes(), nil
}
func (m *PolicerAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Name = buf.DecodeString(64)
	m.Cir = buf.DecodeUint32()
	m.Eir = buf.DecodeUint32()
	m.Cb = buf.DecodeUint64()
	m.Eb = buf.DecodeUint64()
	m.RateType = policer_types.Sse2QosRateType(buf.DecodeUint8())
	m.RoundType = policer_types.Sse2QosRoundType(buf.DecodeUint8())
	m.Type = policer_types.Sse2QosPolicerType(buf.DecodeUint8())
	m.ColorAware = buf.DecodeBool()
	m.ConformAction.Type = policer_types.Sse2QosActionType(buf.DecodeUint8())
	m.ConformAction.Dscp = buf.DecodeUint8()
	m.ExceedAction.Type = policer_types.Sse2QosActionType(buf.DecodeUint8())
	m.ExceedAction.Dscp = buf.DecodeUint8()
	m.ViolateAction.Type = policer_types.Sse2QosActionType(buf.DecodeUint8())
	m.ViolateAction.Dscp = buf.DecodeUint8()
	return nil
}

// PolicerAddDelReply defines message 'policer_add_del_reply'.
type PolicerAddDelReply struct {
	Retval       int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	PolicerIndex uint32 `binapi:"u32,name=policer_index" json:"policer_index,omitempty"`
}

func (m *PolicerAddDelReply) Reset()               { *m = PolicerAddDelReply{} }
func (*PolicerAddDelReply) GetMessageName() string { return "policer_add_del_reply" }
func (*PolicerAddDelReply) GetCrcString() string   { return "a177cef2" }
func (*PolicerAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PolicerAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.PolicerIndex
	return size
}
func (m *PolicerAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.PolicerIndex)
	return buf.Bytes(), nil
}
func (m *PolicerAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.PolicerIndex = buf.DecodeUint32()
	return nil
}

// PolicerBind defines message 'policer_bind'.
type PolicerBind struct {
	Name        string `binapi:"string[64],name=name" json:"name,omitempty"`
	WorkerIndex uint32 `binapi:"u32,name=worker_index" json:"worker_index,omitempty"`
	BindEnable  bool   `binapi:"bool,name=bind_enable" json:"bind_enable,omitempty"`
}

func (m *PolicerBind) Reset()               { *m = PolicerBind{} }
func (*PolicerBind) GetMessageName() string { return "policer_bind" }
func (*PolicerBind) GetCrcString() string   { return "dcf516f9" }
func (*PolicerBind) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PolicerBind) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.Name
	size += 4  // m.WorkerIndex
	size += 1  // m.BindEnable
	return size
}
func (m *PolicerBind) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.Name, 64)
	buf.EncodeUint32(m.WorkerIndex)
	buf.EncodeBool(m.BindEnable)
	return buf.Bytes(), nil
}
func (m *PolicerBind) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Name = buf.DecodeString(64)
	m.WorkerIndex = buf.DecodeUint32()
	m.BindEnable = buf.DecodeBool()
	return nil
}

// PolicerBindReply defines message 'policer_bind_reply'.
type PolicerBindReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PolicerBindReply) Reset()               { *m = PolicerBindReply{} }
func (*PolicerBindReply) GetMessageName() string { return "policer_bind_reply" }
func (*PolicerBindReply) GetCrcString() string   { return "e8d4e804" }
func (*PolicerBindReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PolicerBindReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PolicerBindReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PolicerBindReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PolicerDetails defines message 'policer_details'.
type PolicerDetails struct {
	Name               string                           `binapi:"string[64],name=name" json:"name,omitempty"`
	Cir                uint32                           `binapi:"u32,name=cir" json:"cir,omitempty"`
	Eir                uint32                           `binapi:"u32,name=eir" json:"eir,omitempty"`
	Cb                 uint64                           `binapi:"u64,name=cb" json:"cb,omitempty"`
	Eb                 uint64                           `binapi:"u64,name=eb" json:"eb,omitempty"`
	RateType           policer_types.Sse2QosRateType    `binapi:"sse2_qos_rate_type,name=rate_type" json:"rate_type,omitempty"`
	RoundType          policer_types.Sse2QosRoundType   `binapi:"sse2_qos_round_type,name=round_type" json:"round_type,omitempty"`
	Type               policer_types.Sse2QosPolicerType `binapi:"sse2_qos_policer_type,name=type" json:"type,omitempty"`
	ConformAction      policer_types.Sse2QosAction      `binapi:"sse2_qos_action,name=conform_action" json:"conform_action,omitempty"`
	ExceedAction       policer_types.Sse2QosAction      `binapi:"sse2_qos_action,name=exceed_action" json:"exceed_action,omitempty"`
	ViolateAction      policer_types.Sse2QosAction      `binapi:"sse2_qos_action,name=violate_action" json:"violate_action,omitempty"`
	SingleRate         bool                             `binapi:"bool,name=single_rate" json:"single_rate,omitempty"`
	ColorAware         bool                             `binapi:"bool,name=color_aware" json:"color_aware,omitempty"`
	Scale              uint32                           `binapi:"u32,name=scale" json:"scale,omitempty"`
	CirTokensPerPeriod uint32                           `binapi:"u32,name=cir_tokens_per_period" json:"cir_tokens_per_period,omitempty"`
	PirTokensPerPeriod uint32                           `binapi:"u32,name=pir_tokens_per_period" json:"pir_tokens_per_period,omitempty"`
	CurrentLimit       uint32                           `binapi:"u32,name=current_limit" json:"current_limit,omitempty"`
	CurrentBucket      uint32                           `binapi:"u32,name=current_bucket" json:"current_bucket,omitempty"`
	ExtendedLimit      uint32                           `binapi:"u32,name=extended_limit" json:"extended_limit,omitempty"`
	ExtendedBucket     uint32                           `binapi:"u32,name=extended_bucket" json:"extended_bucket,omitempty"`
	LastUpdateTime     uint64                           `binapi:"u64,name=last_update_time" json:"last_update_time,omitempty"`
}

func (m *PolicerDetails) Reset()               { *m = PolicerDetails{} }
func (*PolicerDetails) GetMessageName() string { return "policer_details" }
func (*PolicerDetails) GetCrcString() string   { return "72d0e248" }
func (*PolicerDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PolicerDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.Name
	size += 4  // m.Cir
	size += 4  // m.Eir
	size += 8  // m.Cb
	size += 8  // m.Eb
	size += 1  // m.RateType
	size += 1  // m.RoundType
	size += 1  // m.Type
	size += 1  // m.ConformAction.Type
	size += 1  // m.ConformAction.Dscp
	size += 1  // m.ExceedAction.Type
	size += 1  // m.ExceedAction.Dscp
	size += 1  // m.ViolateAction.Type
	size += 1  // m.ViolateAction.Dscp
	size += 1  // m.SingleRate
	size += 1  // m.ColorAware
	size += 4  // m.Scale
	size += 4  // m.CirTokensPerPeriod
	size += 4  // m.PirTokensPerPeriod
	size += 4  // m.CurrentLimit
	size += 4  // m.CurrentBucket
	size += 4  // m.ExtendedLimit
	size += 4  // m.ExtendedBucket
	size += 8  // m.LastUpdateTime
	return size
}
func (m *PolicerDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.Name, 64)
	buf.EncodeUint32(m.Cir)
	buf.EncodeUint32(m.Eir)
	buf.EncodeUint64(m.Cb)
	buf.EncodeUint64(m.Eb)
	buf.EncodeUint8(uint8(m.RateType))
	buf.EncodeUint8(uint8(m.RoundType))
	buf.EncodeUint8(uint8(m.Type))
	buf.EncodeUint8(uint8(m.ConformAction.Type))
	buf.EncodeUint8(m.ConformAction.Dscp)
	buf.EncodeUint8(uint8(m.ExceedAction.Type))
	buf.EncodeUint8(m.ExceedAction.Dscp)
	buf.EncodeUint8(uint8(m.ViolateAction.Type))
	buf.EncodeUint8(m.ViolateAction.Dscp)
	buf.EncodeBool(m.SingleRate)
	buf.EncodeBool(m.ColorAware)
	buf.EncodeUint32(m.Scale)
	buf.EncodeUint32(m.CirTokensPerPeriod)
	buf.EncodeUint32(m.PirTokensPerPeriod)
	buf.EncodeUint32(m.CurrentLimit)
	buf.EncodeUint32(m.CurrentBucket)
	buf.EncodeUint32(m.ExtendedLimit)
	buf.EncodeUint32(m.ExtendedBucket)
	buf.EncodeUint64(m.LastUpdateTime)
	return buf.Bytes(), nil
}
func (m *PolicerDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Name = buf.DecodeString(64)
	m.Cir = buf.DecodeUint32()
	m.Eir = buf.DecodeUint32()
	m.Cb = buf.DecodeUint64()
	m.Eb = buf.DecodeUint64()
	m.RateType = policer_types.Sse2QosRateType(buf.DecodeUint8())
	m.RoundType = policer_types.Sse2QosRoundType(buf.DecodeUint8())
	m.Type = policer_types.Sse2QosPolicerType(buf.DecodeUint8())
	m.ConformAction.Type = policer_types.Sse2QosActionType(buf.DecodeUint8())
	m.ConformAction.Dscp = buf.DecodeUint8()
	m.ExceedAction.Type = policer_types.Sse2QosActionType(buf.DecodeUint8())
	m.ExceedAction.Dscp = buf.DecodeUint8()
	m.ViolateAction.Type = policer_types.Sse2QosActionType(buf.DecodeUint8())
	m.ViolateAction.Dscp = buf.DecodeUint8()
	m.SingleRate = buf.DecodeBool()
	m.ColorAware = buf.DecodeBool()
	m.Scale = buf.DecodeUint32()
	m.CirTokensPerPeriod = buf.DecodeUint32()
	m.PirTokensPerPeriod = buf.DecodeUint32()
	m.CurrentLimit = buf.DecodeUint32()
	m.CurrentBucket = buf.DecodeUint32()
	m.ExtendedLimit = buf.DecodeUint32()
	m.ExtendedBucket = buf.DecodeUint32()
	m.LastUpdateTime = buf.DecodeUint64()
	return nil
}

// PolicerDump defines message 'policer_dump'.
type PolicerDump struct {
	MatchNameValid bool   `binapi:"bool,name=match_name_valid" json:"match_name_valid,omitempty"`
	MatchName      string `binapi:"string[64],name=match_name" json:"match_name,omitempty"`
}

func (m *PolicerDump) Reset()               { *m = PolicerDump{} }
func (*PolicerDump) GetMessageName() string { return "policer_dump" }
func (*PolicerDump) GetCrcString() string   { return "35f1ae0f" }
func (*PolicerDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PolicerDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1  // m.MatchNameValid
	size += 64 // m.MatchName
	return size
}
func (m *PolicerDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.MatchNameValid)
	buf.EncodeString(m.MatchName, 64)
	return buf.Bytes(), nil
}
func (m *PolicerDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.MatchNameValid = buf.DecodeBool()
	m.MatchName = buf.DecodeString(64)
	return nil
}

// PolicerInput defines message 'policer_input'.
type PolicerInput struct {
	Name      string                         `binapi:"string[64],name=name" json:"name,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Apply     bool                           `binapi:"bool,name=apply" json:"apply,omitempty"`
}

func (m *PolicerInput) Reset()               { *m = PolicerInput{} }
func (*PolicerInput) GetMessageName() string { return "policer_input" }
func (*PolicerInput) GetCrcString() string   { return "233f0ef5" }
func (*PolicerInput) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PolicerInput) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.Name
	size += 4  // m.SwIfIndex
	size += 1  // m.Apply
	return size
}
func (m *PolicerInput) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.Name, 64)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.Apply)
	return buf.Bytes(), nil
}
func (m *PolicerInput) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Name = buf.DecodeString(64)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Apply = buf.DecodeBool()
	return nil
}

// PolicerInputReply defines message 'policer_input_reply'.
type PolicerInputReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PolicerInputReply) Reset()               { *m = PolicerInputReply{} }
func (*PolicerInputReply) GetMessageName() string { return "policer_input_reply" }
func (*PolicerInputReply) GetCrcString() string   { return "e8d4e804" }
func (*PolicerInputReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PolicerInputReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PolicerInputReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PolicerInputReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_policer_binapi_init() }
func file_policer_binapi_init() {
	api.RegisterMessage((*PolicerAddDel)(nil), "policer_add_del_2b31dd38")
	api.RegisterMessage((*PolicerAddDelReply)(nil), "policer_add_del_reply_a177cef2")
	api.RegisterMessage((*PolicerBind)(nil), "policer_bind_dcf516f9")
	api.RegisterMessage((*PolicerBindReply)(nil), "policer_bind_reply_e8d4e804")
	api.RegisterMessage((*PolicerDetails)(nil), "policer_details_72d0e248")
	api.RegisterMessage((*PolicerDump)(nil), "policer_dump_35f1ae0f")
	api.RegisterMessage((*PolicerInput)(nil), "policer_input_233f0ef5")
	api.RegisterMessage((*PolicerInputReply)(nil), "policer_input_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PolicerAddDel)(nil),
		(*PolicerAddDelReply)(nil),
		(*PolicerBind)(nil),
		(*PolicerBindReply)(nil),
		(*PolicerDetails)(nil),
		(*PolicerDump)(nil),
		(*PolicerInput)(nil),
		(*PolicerInputReply)(nil),
	}
}
