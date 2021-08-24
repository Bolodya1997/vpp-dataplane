// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package session contains generated bindings for API file session.api.
//
// Contents:
//   2 enums
//  22 messages
//
package session

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/projectcalico/vpp-dataplane/vpplink/binapi/vppapi/interface_types"
	ip_types "github.com/projectcalico/vpp-dataplane/vpplink/binapi/vppapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "session"
	APIVersion = "4.0.0"
	VersionCrc = 0x32fe1c77
)

// SessionRuleScope defines enum 'session_rule_scope'.
type SessionRuleScope uint32

const (
	SESSION_RULE_SCOPE_API_GLOBAL SessionRuleScope = 0
	SESSION_RULE_SCOPE_API_LOCAL  SessionRuleScope = 1
	SESSION_RULE_SCOPE_API_BOTH   SessionRuleScope = 2
)

var (
	SessionRuleScope_name = map[uint32]string{
		0: "SESSION_RULE_SCOPE_API_GLOBAL",
		1: "SESSION_RULE_SCOPE_API_LOCAL",
		2: "SESSION_RULE_SCOPE_API_BOTH",
	}
	SessionRuleScope_value = map[string]uint32{
		"SESSION_RULE_SCOPE_API_GLOBAL": 0,
		"SESSION_RULE_SCOPE_API_LOCAL":  1,
		"SESSION_RULE_SCOPE_API_BOTH":   2,
	}
)

func (x SessionRuleScope) String() string {
	s, ok := SessionRuleScope_name[uint32(x)]
	if ok {
		return s
	}
	return "SessionRuleScope(" + strconv.Itoa(int(x)) + ")"
}

// TransportProto defines enum 'transport_proto'.
type TransportProto uint8

const (
	TRANSPORT_PROTO_API_TCP  TransportProto = 1
	TRANSPORT_PROTO_API_UDP  TransportProto = 2
	TRANSPORT_PROTO_API_NONE TransportProto = 3
	TRANSPORT_PROTO_API_TLS  TransportProto = 4
	TRANSPORT_PROTO_API_QUIC TransportProto = 5
)

var (
	TransportProto_name = map[uint8]string{
		1: "TRANSPORT_PROTO_API_TCP",
		2: "TRANSPORT_PROTO_API_UDP",
		3: "TRANSPORT_PROTO_API_NONE",
		4: "TRANSPORT_PROTO_API_TLS",
		5: "TRANSPORT_PROTO_API_QUIC",
	}
	TransportProto_value = map[string]uint8{
		"TRANSPORT_PROTO_API_TCP":  1,
		"TRANSPORT_PROTO_API_UDP":  2,
		"TRANSPORT_PROTO_API_NONE": 3,
		"TRANSPORT_PROTO_API_TLS":  4,
		"TRANSPORT_PROTO_API_QUIC": 5,
	}
)

func (x TransportProto) String() string {
	s, ok := TransportProto_name[uint8(x)]
	if ok {
		return s
	}
	return "TransportProto(" + strconv.Itoa(int(x)) + ")"
}

// AppAddCertKeyPair defines message 'app_add_cert_key_pair'.
type AppAddCertKeyPair struct {
	CertLen    uint16 `binapi:"u16,name=cert_len" json:"cert_len,omitempty"`
	CertkeyLen uint16 `binapi:"u16,name=certkey_len" json:"-"`
	Certkey    []byte `binapi:"u8[certkey_len],name=certkey" json:"certkey,omitempty"`
}

func (m *AppAddCertKeyPair) Reset()               { *m = AppAddCertKeyPair{} }
func (*AppAddCertKeyPair) GetMessageName() string { return "app_add_cert_key_pair" }
func (*AppAddCertKeyPair) GetCrcString() string   { return "02eb8016" }
func (*AppAddCertKeyPair) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AppAddCertKeyPair) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 2                  // m.CertLen
	size += 2                  // m.CertkeyLen
	size += 1 * len(m.Certkey) // m.Certkey
	return size
}
func (m *AppAddCertKeyPair) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint16(m.CertLen)
	buf.EncodeUint16(uint16(len(m.Certkey)))
	buf.EncodeBytes(m.Certkey, 0)
	return buf.Bytes(), nil
}
func (m *AppAddCertKeyPair) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.CertLen = buf.DecodeUint16()
	m.CertkeyLen = buf.DecodeUint16()
	m.Certkey = make([]byte, m.CertkeyLen)
	copy(m.Certkey, buf.DecodeBytes(len(m.Certkey)))
	return nil
}

// AppAddCertKeyPairReply defines message 'app_add_cert_key_pair_reply'.
type AppAddCertKeyPairReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Index  uint32 `binapi:"u32,name=index" json:"index,omitempty"`
}

func (m *AppAddCertKeyPairReply) Reset()               { *m = AppAddCertKeyPairReply{} }
func (*AppAddCertKeyPairReply) GetMessageName() string { return "app_add_cert_key_pair_reply" }
func (*AppAddCertKeyPairReply) GetCrcString() string   { return "b42958d0" }
func (*AppAddCertKeyPairReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AppAddCertKeyPairReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Index
	return size
}
func (m *AppAddCertKeyPairReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.Index)
	return buf.Bytes(), nil
}
func (m *AppAddCertKeyPairReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Index = buf.DecodeUint32()
	return nil
}

// AppAttach defines message 'app_attach'.
type AppAttach struct {
	Options     []uint64 `binapi:"u64[18],name=options" json:"options,omitempty"`
	NamespaceID string   `binapi:"string[],name=namespace_id" json:"namespace_id,omitempty"`
}

func (m *AppAttach) Reset()               { *m = AppAttach{} }
func (*AppAttach) GetMessageName() string { return "app_attach" }
func (*AppAttach) GetCrcString() string   { return "5f4a260d" }
func (*AppAttach) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AppAttach) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8 * 18                 // m.Options
	size += 4 + len(m.NamespaceID) // m.NamespaceID
	return size
}
func (m *AppAttach) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	for i := 0; i < 18; i++ {
		var x uint64
		if i < len(m.Options) {
			x = uint64(m.Options[i])
		}
		buf.EncodeUint64(x)
	}
	buf.EncodeString(m.NamespaceID, 0)
	return buf.Bytes(), nil
}
func (m *AppAttach) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Options = make([]uint64, 18)
	for i := 0; i < len(m.Options); i++ {
		m.Options[i] = buf.DecodeUint64()
	}
	m.NamespaceID = buf.DecodeString(0)
	return nil
}

// AppAttachReply defines message 'app_attach_reply'.
type AppAttachReply struct {
	Retval          int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	AppMq           uint64 `binapi:"u64,name=app_mq" json:"app_mq,omitempty"`
	VppCtrlMq       uint64 `binapi:"u64,name=vpp_ctrl_mq" json:"vpp_ctrl_mq,omitempty"`
	VppCtrlMqThread uint8  `binapi:"u8,name=vpp_ctrl_mq_thread" json:"vpp_ctrl_mq_thread,omitempty"`
	AppIndex        uint32 `binapi:"u32,name=app_index" json:"app_index,omitempty"`
	NFds            uint8  `binapi:"u8,name=n_fds" json:"n_fds,omitempty"`
	FdFlags         uint8  `binapi:"u8,name=fd_flags" json:"fd_flags,omitempty"`
	SegmentSize     uint32 `binapi:"u32,name=segment_size" json:"segment_size,omitempty"`
	SegmentHandle   uint64 `binapi:"u64,name=segment_handle" json:"segment_handle,omitempty"`
	SegmentName     string `binapi:"string[],name=segment_name" json:"segment_name,omitempty"`
}

func (m *AppAttachReply) Reset()               { *m = AppAttachReply{} }
func (*AppAttachReply) GetMessageName() string { return "app_attach_reply" }
func (*AppAttachReply) GetCrcString() string   { return "5c89c3b0" }
func (*AppAttachReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AppAttachReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                      // m.Retval
	size += 8                      // m.AppMq
	size += 8                      // m.VppCtrlMq
	size += 1                      // m.VppCtrlMqThread
	size += 4                      // m.AppIndex
	size += 1                      // m.NFds
	size += 1                      // m.FdFlags
	size += 4                      // m.SegmentSize
	size += 8                      // m.SegmentHandle
	size += 4 + len(m.SegmentName) // m.SegmentName
	return size
}
func (m *AppAttachReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint64(m.AppMq)
	buf.EncodeUint64(m.VppCtrlMq)
	buf.EncodeUint8(m.VppCtrlMqThread)
	buf.EncodeUint32(m.AppIndex)
	buf.EncodeUint8(m.NFds)
	buf.EncodeUint8(m.FdFlags)
	buf.EncodeUint32(m.SegmentSize)
	buf.EncodeUint64(m.SegmentHandle)
	buf.EncodeString(m.SegmentName, 0)
	return buf.Bytes(), nil
}
func (m *AppAttachReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.AppMq = buf.DecodeUint64()
	m.VppCtrlMq = buf.DecodeUint64()
	m.VppCtrlMqThread = buf.DecodeUint8()
	m.AppIndex = buf.DecodeUint32()
	m.NFds = buf.DecodeUint8()
	m.FdFlags = buf.DecodeUint8()
	m.SegmentSize = buf.DecodeUint32()
	m.SegmentHandle = buf.DecodeUint64()
	m.SegmentName = buf.DecodeString(0)
	return nil
}

// AppDelCertKeyPair defines message 'app_del_cert_key_pair'.
type AppDelCertKeyPair struct {
	Index uint32 `binapi:"u32,name=index" json:"index,omitempty"`
}

func (m *AppDelCertKeyPair) Reset()               { *m = AppDelCertKeyPair{} }
func (*AppDelCertKeyPair) GetMessageName() string { return "app_del_cert_key_pair" }
func (*AppDelCertKeyPair) GetCrcString() string   { return "8ac76db6" }
func (*AppDelCertKeyPair) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AppDelCertKeyPair) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Index
	return size
}
func (m *AppDelCertKeyPair) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Index)
	return buf.Bytes(), nil
}
func (m *AppDelCertKeyPair) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Index = buf.DecodeUint32()
	return nil
}

// AppDelCertKeyPairReply defines message 'app_del_cert_key_pair_reply'.
type AppDelCertKeyPairReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AppDelCertKeyPairReply) Reset()               { *m = AppDelCertKeyPairReply{} }
func (*AppDelCertKeyPairReply) GetMessageName() string { return "app_del_cert_key_pair_reply" }
func (*AppDelCertKeyPairReply) GetCrcString() string   { return "e8d4e804" }
func (*AppDelCertKeyPairReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AppDelCertKeyPairReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AppDelCertKeyPairReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AppDelCertKeyPairReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// AppNamespaceAddDel defines message 'app_namespace_add_del'.
type AppNamespaceAddDel struct {
	Secret      uint64                         `binapi:"u64,name=secret" json:"secret,omitempty"`
	SwIfIndex   interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
	IP4FibID    uint32                         `binapi:"u32,name=ip4_fib_id" json:"ip4_fib_id,omitempty"`
	IP6FibID    uint32                         `binapi:"u32,name=ip6_fib_id" json:"ip6_fib_id,omitempty"`
	NamespaceID string                         `binapi:"string[],name=namespace_id" json:"namespace_id,omitempty"`
}

func (m *AppNamespaceAddDel) Reset()               { *m = AppNamespaceAddDel{} }
func (*AppNamespaceAddDel) GetMessageName() string { return "app_namespace_add_del" }
func (*AppNamespaceAddDel) GetCrcString() string   { return "6306aecb" }
func (*AppNamespaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AppNamespaceAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8                      // m.Secret
	size += 4                      // m.SwIfIndex
	size += 4                      // m.IP4FibID
	size += 4                      // m.IP6FibID
	size += 4 + len(m.NamespaceID) // m.NamespaceID
	return size
}
func (m *AppNamespaceAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint64(m.Secret)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(m.IP4FibID)
	buf.EncodeUint32(m.IP6FibID)
	buf.EncodeString(m.NamespaceID, 0)
	return buf.Bytes(), nil
}
func (m *AppNamespaceAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Secret = buf.DecodeUint64()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IP4FibID = buf.DecodeUint32()
	m.IP6FibID = buf.DecodeUint32()
	m.NamespaceID = buf.DecodeString(0)
	return nil
}

// AppNamespaceAddDelReply defines message 'app_namespace_add_del_reply'.
type AppNamespaceAddDelReply struct {
	Retval     int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	AppnsIndex uint32 `binapi:"u32,name=appns_index" json:"appns_index,omitempty"`
}

func (m *AppNamespaceAddDelReply) Reset()               { *m = AppNamespaceAddDelReply{} }
func (*AppNamespaceAddDelReply) GetMessageName() string { return "app_namespace_add_del_reply" }
func (*AppNamespaceAddDelReply) GetCrcString() string   { return "85137120" }
func (*AppNamespaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AppNamespaceAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.AppnsIndex
	return size
}
func (m *AppNamespaceAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.AppnsIndex)
	return buf.Bytes(), nil
}
func (m *AppNamespaceAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.AppnsIndex = buf.DecodeUint32()
	return nil
}

// AppWorkerAddDel defines message 'app_worker_add_del'.
type AppWorkerAddDel struct {
	AppIndex uint32 `binapi:"u32,name=app_index" json:"app_index,omitempty"`
	WrkIndex uint32 `binapi:"u32,name=wrk_index" json:"wrk_index,omitempty"`
	IsAdd    bool   `binapi:"bool,name=is_add,default=true" json:"is_add,omitempty"`
}

func (m *AppWorkerAddDel) Reset()               { *m = AppWorkerAddDel{} }
func (*AppWorkerAddDel) GetMessageName() string { return "app_worker_add_del" }
func (*AppWorkerAddDel) GetCrcString() string   { return "753253dc" }
func (*AppWorkerAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AppWorkerAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.AppIndex
	size += 4 // m.WrkIndex
	size += 1 // m.IsAdd
	return size
}
func (m *AppWorkerAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.AppIndex)
	buf.EncodeUint32(m.WrkIndex)
	buf.EncodeBool(m.IsAdd)
	return buf.Bytes(), nil
}
func (m *AppWorkerAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.AppIndex = buf.DecodeUint32()
	m.WrkIndex = buf.DecodeUint32()
	m.IsAdd = buf.DecodeBool()
	return nil
}

// AppWorkerAddDelReply defines message 'app_worker_add_del_reply'.
type AppWorkerAddDelReply struct {
	Retval               int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	WrkIndex             uint32 `binapi:"u32,name=wrk_index" json:"wrk_index,omitempty"`
	AppEventQueueAddress uint64 `binapi:"u64,name=app_event_queue_address" json:"app_event_queue_address,omitempty"`
	NFds                 uint8  `binapi:"u8,name=n_fds" json:"n_fds,omitempty"`
	FdFlags              uint8  `binapi:"u8,name=fd_flags" json:"fd_flags,omitempty"`
	SegmentHandle        uint64 `binapi:"u64,name=segment_handle" json:"segment_handle,omitempty"`
	IsAdd                bool   `binapi:"bool,name=is_add,default=true" json:"is_add,omitempty"`
	SegmentName          string `binapi:"string[],name=segment_name" json:"segment_name,omitempty"`
}

func (m *AppWorkerAddDelReply) Reset()               { *m = AppWorkerAddDelReply{} }
func (*AppWorkerAddDelReply) GetMessageName() string { return "app_worker_add_del_reply" }
func (*AppWorkerAddDelReply) GetCrcString() string   { return "5735ffe7" }
func (*AppWorkerAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AppWorkerAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                      // m.Retval
	size += 4                      // m.WrkIndex
	size += 8                      // m.AppEventQueueAddress
	size += 1                      // m.NFds
	size += 1                      // m.FdFlags
	size += 8                      // m.SegmentHandle
	size += 1                      // m.IsAdd
	size += 4 + len(m.SegmentName) // m.SegmentName
	return size
}
func (m *AppWorkerAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.WrkIndex)
	buf.EncodeUint64(m.AppEventQueueAddress)
	buf.EncodeUint8(m.NFds)
	buf.EncodeUint8(m.FdFlags)
	buf.EncodeUint64(m.SegmentHandle)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeString(m.SegmentName, 0)
	return buf.Bytes(), nil
}
func (m *AppWorkerAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.WrkIndex = buf.DecodeUint32()
	m.AppEventQueueAddress = buf.DecodeUint64()
	m.NFds = buf.DecodeUint8()
	m.FdFlags = buf.DecodeUint8()
	m.SegmentHandle = buf.DecodeUint64()
	m.IsAdd = buf.DecodeBool()
	m.SegmentName = buf.DecodeString(0)
	return nil
}

// ApplicationDetach defines message 'application_detach'.
type ApplicationDetach struct{}

func (m *ApplicationDetach) Reset()               { *m = ApplicationDetach{} }
func (*ApplicationDetach) GetMessageName() string { return "application_detach" }
func (*ApplicationDetach) GetCrcString() string   { return "51077d14" }
func (*ApplicationDetach) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ApplicationDetach) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *ApplicationDetach) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *ApplicationDetach) Unmarshal(b []byte) error {
	return nil
}

// ApplicationDetachReply defines message 'application_detach_reply'.
type ApplicationDetachReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *ApplicationDetachReply) Reset()               { *m = ApplicationDetachReply{} }
func (*ApplicationDetachReply) GetMessageName() string { return "application_detach_reply" }
func (*ApplicationDetachReply) GetCrcString() string   { return "e8d4e804" }
func (*ApplicationDetachReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ApplicationDetachReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *ApplicationDetachReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *ApplicationDetachReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// ApplicationTLSCertAdd defines message 'application_tls_cert_add'.
type ApplicationTLSCertAdd struct {
	AppIndex uint32 `binapi:"u32,name=app_index" json:"app_index,omitempty"`
	CertLen  uint16 `binapi:"u16,name=cert_len" json:"-"`
	Cert     []byte `binapi:"u8[cert_len],name=cert" json:"cert,omitempty"`
}

func (m *ApplicationTLSCertAdd) Reset()               { *m = ApplicationTLSCertAdd{} }
func (*ApplicationTLSCertAdd) GetMessageName() string { return "application_tls_cert_add" }
func (*ApplicationTLSCertAdd) GetCrcString() string   { return "3f5cfe45" }
func (*ApplicationTLSCertAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ApplicationTLSCertAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4               // m.AppIndex
	size += 2               // m.CertLen
	size += 1 * len(m.Cert) // m.Cert
	return size
}
func (m *ApplicationTLSCertAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.AppIndex)
	buf.EncodeUint16(uint16(len(m.Cert)))
	buf.EncodeBytes(m.Cert, 0)
	return buf.Bytes(), nil
}
func (m *ApplicationTLSCertAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.AppIndex = buf.DecodeUint32()
	m.CertLen = buf.DecodeUint16()
	m.Cert = make([]byte, m.CertLen)
	copy(m.Cert, buf.DecodeBytes(len(m.Cert)))
	return nil
}

// ApplicationTLSCertAddReply defines message 'application_tls_cert_add_reply'.
type ApplicationTLSCertAddReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *ApplicationTLSCertAddReply) Reset()               { *m = ApplicationTLSCertAddReply{} }
func (*ApplicationTLSCertAddReply) GetMessageName() string { return "application_tls_cert_add_reply" }
func (*ApplicationTLSCertAddReply) GetCrcString() string   { return "e8d4e804" }
func (*ApplicationTLSCertAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ApplicationTLSCertAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *ApplicationTLSCertAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *ApplicationTLSCertAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// ApplicationTLSKeyAdd defines message 'application_tls_key_add'.
type ApplicationTLSKeyAdd struct {
	AppIndex uint32 `binapi:"u32,name=app_index" json:"app_index,omitempty"`
	KeyLen   uint16 `binapi:"u16,name=key_len" json:"-"`
	Key      []byte `binapi:"u8[key_len],name=key" json:"key,omitempty"`
}

func (m *ApplicationTLSKeyAdd) Reset()               { *m = ApplicationTLSKeyAdd{} }
func (*ApplicationTLSKeyAdd) GetMessageName() string { return "application_tls_key_add" }
func (*ApplicationTLSKeyAdd) GetCrcString() string   { return "5eaf70cd" }
func (*ApplicationTLSKeyAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ApplicationTLSKeyAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4              // m.AppIndex
	size += 2              // m.KeyLen
	size += 1 * len(m.Key) // m.Key
	return size
}
func (m *ApplicationTLSKeyAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.AppIndex)
	buf.EncodeUint16(uint16(len(m.Key)))
	buf.EncodeBytes(m.Key, 0)
	return buf.Bytes(), nil
}
func (m *ApplicationTLSKeyAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.AppIndex = buf.DecodeUint32()
	m.KeyLen = buf.DecodeUint16()
	m.Key = make([]byte, m.KeyLen)
	copy(m.Key, buf.DecodeBytes(len(m.Key)))
	return nil
}

// ApplicationTLSKeyAddReply defines message 'application_tls_key_add_reply'.
type ApplicationTLSKeyAddReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *ApplicationTLSKeyAddReply) Reset()               { *m = ApplicationTLSKeyAddReply{} }
func (*ApplicationTLSKeyAddReply) GetMessageName() string { return "application_tls_key_add_reply" }
func (*ApplicationTLSKeyAddReply) GetCrcString() string   { return "e8d4e804" }
func (*ApplicationTLSKeyAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ApplicationTLSKeyAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *ApplicationTLSKeyAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *ApplicationTLSKeyAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SessionEnableDisable defines message 'session_enable_disable'.
type SessionEnableDisable struct {
	IsEnable bool `binapi:"bool,name=is_enable,default=true" json:"is_enable,omitempty"`
}

func (m *SessionEnableDisable) Reset()               { *m = SessionEnableDisable{} }
func (*SessionEnableDisable) GetMessageName() string { return "session_enable_disable" }
func (*SessionEnableDisable) GetCrcString() string   { return "c264d7bf" }
func (*SessionEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SessionEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsEnable
	return size
}
func (m *SessionEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsEnable)
	return buf.Bytes(), nil
}
func (m *SessionEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsEnable = buf.DecodeBool()
	return nil
}

// SessionEnableDisableReply defines message 'session_enable_disable_reply'.
type SessionEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SessionEnableDisableReply) Reset()               { *m = SessionEnableDisableReply{} }
func (*SessionEnableDisableReply) GetMessageName() string { return "session_enable_disable_reply" }
func (*SessionEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*SessionEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SessionEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SessionEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SessionEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SessionRuleAddDel defines message 'session_rule_add_del'.
type SessionRuleAddDel struct {
	TransportProto TransportProto   `binapi:"transport_proto,name=transport_proto" json:"transport_proto,omitempty"`
	Lcl            ip_types.Prefix  `binapi:"prefix,name=lcl" json:"lcl,omitempty"`
	Rmt            ip_types.Prefix  `binapi:"prefix,name=rmt" json:"rmt,omitempty"`
	LclPort        uint16           `binapi:"u16,name=lcl_port" json:"lcl_port,omitempty"`
	RmtPort        uint16           `binapi:"u16,name=rmt_port" json:"rmt_port,omitempty"`
	ActionIndex    uint32           `binapi:"u32,name=action_index" json:"action_index,omitempty"`
	IsAdd          bool             `binapi:"bool,name=is_add,default=true" json:"is_add,omitempty"`
	AppnsIndex     uint32           `binapi:"u32,name=appns_index" json:"appns_index,omitempty"`
	Scope          SessionRuleScope `binapi:"session_rule_scope,name=scope" json:"scope,omitempty"`
	Tag            string           `binapi:"string[64],name=tag" json:"tag,omitempty"`
}

func (m *SessionRuleAddDel) Reset()               { *m = SessionRuleAddDel{} }
func (*SessionRuleAddDel) GetMessageName() string { return "session_rule_add_del" }
func (*SessionRuleAddDel) GetCrcString() string   { return "e4895422" }
func (*SessionRuleAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SessionRuleAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.TransportProto
	size += 1      // m.Lcl.Address.Af
	size += 1 * 16 // m.Lcl.Address.Un
	size += 1      // m.Lcl.Len
	size += 1      // m.Rmt.Address.Af
	size += 1 * 16 // m.Rmt.Address.Un
	size += 1      // m.Rmt.Len
	size += 2      // m.LclPort
	size += 2      // m.RmtPort
	size += 4      // m.ActionIndex
	size += 1      // m.IsAdd
	size += 4      // m.AppnsIndex
	size += 4      // m.Scope
	size += 64     // m.Tag
	return size
}
func (m *SessionRuleAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.TransportProto))
	buf.EncodeUint8(uint8(m.Lcl.Address.Af))
	buf.EncodeBytes(m.Lcl.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Lcl.Len)
	buf.EncodeUint8(uint8(m.Rmt.Address.Af))
	buf.EncodeBytes(m.Rmt.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Rmt.Len)
	buf.EncodeUint16(m.LclPort)
	buf.EncodeUint16(m.RmtPort)
	buf.EncodeUint32(m.ActionIndex)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.AppnsIndex)
	buf.EncodeUint32(uint32(m.Scope))
	buf.EncodeString(m.Tag, 64)
	return buf.Bytes(), nil
}
func (m *SessionRuleAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TransportProto = TransportProto(buf.DecodeUint8())
	m.Lcl.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Lcl.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Lcl.Len = buf.DecodeUint8()
	m.Rmt.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Rmt.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Rmt.Len = buf.DecodeUint8()
	m.LclPort = buf.DecodeUint16()
	m.RmtPort = buf.DecodeUint16()
	m.ActionIndex = buf.DecodeUint32()
	m.IsAdd = buf.DecodeBool()
	m.AppnsIndex = buf.DecodeUint32()
	m.Scope = SessionRuleScope(buf.DecodeUint32())
	m.Tag = buf.DecodeString(64)
	return nil
}

// SessionRuleAddDelReply defines message 'session_rule_add_del_reply'.
type SessionRuleAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SessionRuleAddDelReply) Reset()               { *m = SessionRuleAddDelReply{} }
func (*SessionRuleAddDelReply) GetMessageName() string { return "session_rule_add_del_reply" }
func (*SessionRuleAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SessionRuleAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SessionRuleAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SessionRuleAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SessionRuleAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SessionRulesDetails defines message 'session_rules_details'.
type SessionRulesDetails struct {
	TransportProto TransportProto   `binapi:"transport_proto,name=transport_proto" json:"transport_proto,omitempty"`
	Lcl            ip_types.Prefix  `binapi:"prefix,name=lcl" json:"lcl,omitempty"`
	Rmt            ip_types.Prefix  `binapi:"prefix,name=rmt" json:"rmt,omitempty"`
	LclPort        uint16           `binapi:"u16,name=lcl_port" json:"lcl_port,omitempty"`
	RmtPort        uint16           `binapi:"u16,name=rmt_port" json:"rmt_port,omitempty"`
	ActionIndex    uint32           `binapi:"u32,name=action_index" json:"action_index,omitempty"`
	AppnsIndex     uint32           `binapi:"u32,name=appns_index" json:"appns_index,omitempty"`
	Scope          SessionRuleScope `binapi:"session_rule_scope,name=scope" json:"scope,omitempty"`
	Tag            string           `binapi:"string[64],name=tag" json:"tag,omitempty"`
}

func (m *SessionRulesDetails) Reset()               { *m = SessionRulesDetails{} }
func (*SessionRulesDetails) GetMessageName() string { return "session_rules_details" }
func (*SessionRulesDetails) GetCrcString() string   { return "28d71830" }
func (*SessionRulesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SessionRulesDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.TransportProto
	size += 1      // m.Lcl.Address.Af
	size += 1 * 16 // m.Lcl.Address.Un
	size += 1      // m.Lcl.Len
	size += 1      // m.Rmt.Address.Af
	size += 1 * 16 // m.Rmt.Address.Un
	size += 1      // m.Rmt.Len
	size += 2      // m.LclPort
	size += 2      // m.RmtPort
	size += 4      // m.ActionIndex
	size += 4      // m.AppnsIndex
	size += 4      // m.Scope
	size += 64     // m.Tag
	return size
}
func (m *SessionRulesDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.TransportProto))
	buf.EncodeUint8(uint8(m.Lcl.Address.Af))
	buf.EncodeBytes(m.Lcl.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Lcl.Len)
	buf.EncodeUint8(uint8(m.Rmt.Address.Af))
	buf.EncodeBytes(m.Rmt.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Rmt.Len)
	buf.EncodeUint16(m.LclPort)
	buf.EncodeUint16(m.RmtPort)
	buf.EncodeUint32(m.ActionIndex)
	buf.EncodeUint32(m.AppnsIndex)
	buf.EncodeUint32(uint32(m.Scope))
	buf.EncodeString(m.Tag, 64)
	return buf.Bytes(), nil
}
func (m *SessionRulesDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TransportProto = TransportProto(buf.DecodeUint8())
	m.Lcl.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Lcl.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Lcl.Len = buf.DecodeUint8()
	m.Rmt.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Rmt.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Rmt.Len = buf.DecodeUint8()
	m.LclPort = buf.DecodeUint16()
	m.RmtPort = buf.DecodeUint16()
	m.ActionIndex = buf.DecodeUint32()
	m.AppnsIndex = buf.DecodeUint32()
	m.Scope = SessionRuleScope(buf.DecodeUint32())
	m.Tag = buf.DecodeString(64)
	return nil
}

// SessionRulesDump defines message 'session_rules_dump'.
type SessionRulesDump struct{}

func (m *SessionRulesDump) Reset()               { *m = SessionRulesDump{} }
func (*SessionRulesDump) GetMessageName() string { return "session_rules_dump" }
func (*SessionRulesDump) GetCrcString() string   { return "51077d14" }
func (*SessionRulesDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SessionRulesDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SessionRulesDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SessionRulesDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_session_binapi_init() }
func file_session_binapi_init() {
	api.RegisterMessage((*AppAddCertKeyPair)(nil), "app_add_cert_key_pair_02eb8016")
	api.RegisterMessage((*AppAddCertKeyPairReply)(nil), "app_add_cert_key_pair_reply_b42958d0")
	api.RegisterMessage((*AppAttach)(nil), "app_attach_5f4a260d")
	api.RegisterMessage((*AppAttachReply)(nil), "app_attach_reply_5c89c3b0")
	api.RegisterMessage((*AppDelCertKeyPair)(nil), "app_del_cert_key_pair_8ac76db6")
	api.RegisterMessage((*AppDelCertKeyPairReply)(nil), "app_del_cert_key_pair_reply_e8d4e804")
	api.RegisterMessage((*AppNamespaceAddDel)(nil), "app_namespace_add_del_6306aecb")
	api.RegisterMessage((*AppNamespaceAddDelReply)(nil), "app_namespace_add_del_reply_85137120")
	api.RegisterMessage((*AppWorkerAddDel)(nil), "app_worker_add_del_753253dc")
	api.RegisterMessage((*AppWorkerAddDelReply)(nil), "app_worker_add_del_reply_5735ffe7")
	api.RegisterMessage((*ApplicationDetach)(nil), "application_detach_51077d14")
	api.RegisterMessage((*ApplicationDetachReply)(nil), "application_detach_reply_e8d4e804")
	api.RegisterMessage((*ApplicationTLSCertAdd)(nil), "application_tls_cert_add_3f5cfe45")
	api.RegisterMessage((*ApplicationTLSCertAddReply)(nil), "application_tls_cert_add_reply_e8d4e804")
	api.RegisterMessage((*ApplicationTLSKeyAdd)(nil), "application_tls_key_add_5eaf70cd")
	api.RegisterMessage((*ApplicationTLSKeyAddReply)(nil), "application_tls_key_add_reply_e8d4e804")
	api.RegisterMessage((*SessionEnableDisable)(nil), "session_enable_disable_c264d7bf")
	api.RegisterMessage((*SessionEnableDisableReply)(nil), "session_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*SessionRuleAddDel)(nil), "session_rule_add_del_e4895422")
	api.RegisterMessage((*SessionRuleAddDelReply)(nil), "session_rule_add_del_reply_e8d4e804")
	api.RegisterMessage((*SessionRulesDetails)(nil), "session_rules_details_28d71830")
	api.RegisterMessage((*SessionRulesDump)(nil), "session_rules_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AppAddCertKeyPair)(nil),
		(*AppAddCertKeyPairReply)(nil),
		(*AppAttach)(nil),
		(*AppAttachReply)(nil),
		(*AppDelCertKeyPair)(nil),
		(*AppDelCertKeyPairReply)(nil),
		(*AppNamespaceAddDel)(nil),
		(*AppNamespaceAddDelReply)(nil),
		(*AppWorkerAddDel)(nil),
		(*AppWorkerAddDelReply)(nil),
		(*ApplicationDetach)(nil),
		(*ApplicationDetachReply)(nil),
		(*ApplicationTLSCertAdd)(nil),
		(*ApplicationTLSCertAddReply)(nil),
		(*ApplicationTLSKeyAdd)(nil),
		(*ApplicationTLSKeyAddReply)(nil),
		(*SessionEnableDisable)(nil),
		(*SessionEnableDisableReply)(nil),
		(*SessionRuleAddDel)(nil),
		(*SessionRuleAddDelReply)(nil),
		(*SessionRulesDetails)(nil),
		(*SessionRulesDump)(nil),
	}
}
