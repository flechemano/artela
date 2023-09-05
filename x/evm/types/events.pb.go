// Code support by protoc-gen-gogo. DO NOT EDIT.
// source: artela/evm/v1/events.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this support file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EventEthereumTx defines the event for an Ethereum txs
type EventEthereumTx struct {
	// amount
	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// eth_hash is the Ethereum hash of the txs
	EthHash string `protobuf:"bytes,2,opt,name=eth_hash,json=ethHash,proto3" json:"eth_hash,omitempty"`
	// index of the txs in the block
	Index string `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
	// gas_used is the amount of gas used by the txs
	GasUsed string `protobuf:"bytes,4,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	// hash is the Tendermint hash of the txs
	Hash string `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	// recipient of the txs
	Recipient string `protobuf:"bytes,6,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// eth_tx_failed contains a VM error should it occur
	EthTxFailed string `protobuf:"bytes,7,opt,name=eth_tx_failed,json=ethTxFailed,proto3" json:"eth_tx_failed,omitempty"`
}

func (m *EventEthereumTx) Reset()         { *m = EventEthereumTx{} }
func (m *EventEthereumTx) String() string { return proto.CompactTextString(m) }
func (*EventEthereumTx) ProtoMessage()    {}
func (*EventEthereumTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f484ef7abe148ad, []int{0}
}
func (m *EventEthereumTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEthereumTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEthereumTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEthereumTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEthereumTx.Merge(m, src)
}
func (m *EventEthereumTx) XXX_Size() int {
	return m.Size()
}
func (m *EventEthereumTx) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEthereumTx.DiscardUnknown(m)
}

var xxx_messageInfo_EventEthereumTx proto.InternalMessageInfo

func (m *EventEthereumTx) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *EventEthereumTx) GetEthHash() string {
	if m != nil {
		return m.EthHash
	}
	return ""
}

func (m *EventEthereumTx) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *EventEthereumTx) GetGasUsed() string {
	if m != nil {
		return m.GasUsed
	}
	return ""
}

func (m *EventEthereumTx) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *EventEthereumTx) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *EventEthereumTx) GetEthTxFailed() string {
	if m != nil {
		return m.EthTxFailed
	}
	return ""
}

// EventTxLog defines the event for an Ethereum txs log
type EventTxLog struct {
	// tx_logs is an array of txs logs
	TxLogs []string `protobuf:"bytes,1,rep,name=tx_logs,json=txLogs,proto3" json:"tx_logs,omitempty"`
}

func (m *EventTxLog) Reset()         { *m = EventTxLog{} }
func (m *EventTxLog) String() string { return proto.CompactTextString(m) }
func (*EventTxLog) ProtoMessage()    {}
func (*EventTxLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f484ef7abe148ad, []int{1}
}
func (m *EventTxLog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventTxLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventTxLog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventTxLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTxLog.Merge(m, src)
}
func (m *EventTxLog) XXX_Size() int {
	return m.Size()
}
func (m *EventTxLog) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTxLog.DiscardUnknown(m)
}

var xxx_messageInfo_EventTxLog proto.InternalMessageInfo

func (m *EventTxLog) GetTxLogs() []string {
	if m != nil {
		return m.TxLogs
	}
	return nil
}

// EventMessage
type EventMessage struct {
	// module which emits the event
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// sender of the message
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	// tx_type is the type of the message
	TxType string `protobuf:"bytes,3,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
}

func (m *EventMessage) Reset()         { *m = EventMessage{} }
func (m *EventMessage) String() string { return proto.CompactTextString(m) }
func (*EventMessage) ProtoMessage()    {}
func (*EventMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f484ef7abe148ad, []int{2}
}
func (m *EventMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMessage.Merge(m, src)
}
func (m *EventMessage) XXX_Size() int {
	return m.Size()
}
func (m *EventMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EventMessage proto.InternalMessageInfo

func (m *EventMessage) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *EventMessage) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventMessage) GetTxType() string {
	if m != nil {
		return m.TxType
	}
	return ""
}

// EventBlockBloom defines an Ethereum block bloom filter event
type EventBlockBloom struct {
	// bloom is the bloom filter of the block
	Bloom string `protobuf:"bytes,1,opt,name=bloom,proto3" json:"bloom,omitempty"`
}

func (m *EventBlockBloom) Reset()         { *m = EventBlockBloom{} }
func (m *EventBlockBloom) String() string { return proto.CompactTextString(m) }
func (*EventBlockBloom) ProtoMessage()    {}
func (*EventBlockBloom) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f484ef7abe148ad, []int{3}
}
func (m *EventBlockBloom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBlockBloom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBlockBloom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBlockBloom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBlockBloom.Merge(m, src)
}
func (m *EventBlockBloom) XXX_Size() int {
	return m.Size()
}
func (m *EventBlockBloom) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBlockBloom.DiscardUnknown(m)
}

var xxx_messageInfo_EventBlockBloom proto.InternalMessageInfo

func (m *EventBlockBloom) GetBloom() string {
	if m != nil {
		return m.Bloom
	}
	return ""
}

func init() {
	proto.RegisterType((*EventEthereumTx)(nil), "artela.evm.v1.EventEthereumTx")
	proto.RegisterType((*EventTxLog)(nil), "artela.evm.v1.EventTxLog")
	proto.RegisterType((*EventMessage)(nil), "artela.evm.v1.EventMessage")
	proto.RegisterType((*EventBlockBloom)(nil), "artela.evm.v1.EventBlockBloom")
}

func init() { proto.RegisterFile("artela/evm/v1/events.proto", fileDescriptor_4f484ef7abe148ad) }

var fileDescriptor_4f484ef7abe148ad = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x51, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xb5, 0x6a, 0x5b, 0xae, 0xb7, 0x35, 0x85, 0xa5, 0xb4, 0x6a, 0x29, 0xc2, 0x08, 0x4a, 0x7b,
	0x89, 0x84, 0xc9, 0x1f, 0x18, 0x1c, 0x12, 0x48, 0x2e, 0x41, 0x21, 0x90, 0x8b, 0x58, 0x5b, 0x13,
	0xad, 0xb0, 0xa4, 0x15, 0xda, 0x91, 0xb2, 0xfe, 0x8b, 0x7c, 0x56, 0x20, 0x17, 0x1f, 0x73, 0x0c,
	0xf6, 0x8f, 0x84, 0x5d, 0x6d, 0xc8, 0x6d, 0xde, 0x7b, 0x33, 0xb3, 0xfb, 0xe6, 0x91, 0xdf, 0xac,
	0x41, 0x28, 0x58, 0x04, 0x5d, 0x19, 0x75, 0x8b, 0x08, 0x3a, 0xa8, 0x50, 0x86, 0x75, 0x23, 0x50,
	0xd0, 0x59, 0xaf, 0x85, 0xd0, 0x95, 0x61, 0xb7, 0x08, 0x9e, 0x1d, 0xf2, 0x6d, 0xa5, 0xf5, 0x15,
	0x72, 0x68, 0xa0, 0x2d, 0x63, 0x45, 0x7f, 0x10, 0x97, 0x95, 0xa2, 0xad, 0xd0, 0x73, 0xe6, 0xce,
	0xff, 0xe9, 0xb5, 0x45, 0xf4, 0x17, 0xf9, 0x0c, 0xc8, 0x13, 0xce, 0x24, 0xf7, 0x3e, 0x19, 0x65,
	0x02, 0xc8, 0xcf, 0x99, 0xe4, 0xf4, 0x3b, 0x19, 0xe7, 0x55, 0x0a, 0xca, 0x1b, 0x1a, 0xbe, 0x07,
	0x7a, 0x20, 0x63, 0x32, 0x69, 0x25, 0xa4, 0xde, 0xa8, 0x1f, 0xc8, 0x98, 0xbc, 0x91, 0x90, 0x52,
	0x4a, 0x46, 0x66, 0xcf, 0xd8, 0xd0, 0xa6, 0xa6, 0x7f, 0xc8, 0xb4, 0x81, 0x4d, 0x5e, 0xe7, 0x50,
	0xa1, 0xe7, 0x1a, 0xe1, 0x83, 0xa0, 0x01, 0x99, 0xe9, 0xd7, 0x51, 0x25, 0xf7, 0x2c, 0x2f, 0x20,
	0xf5, 0x26, 0xa6, 0xe3, 0x0b, 0x20, 0x8f, 0xd5, 0x99, 0xa1, 0x82, 0xbf, 0x84, 0x18, 0x33, 0xb1,
	0xba, 0x14, 0x19, 0xfd, 0x49, 0x26, 0xa8, 0x92, 0x42, 0x64, 0xd2, 0x73, 0xe6, 0x43, 0x6d, 0x04,
	0x35, 0x2f, 0x83, 0x5b, 0xf2, 0xd5, 0xb4, 0x5d, 0x81, 0x94, 0x2c, 0x03, 0x6d, 0xb8, 0x14, 0x69,
	0x5b, 0xc0, 0xbb, 0xe1, 0x1e, 0x69, 0x5e, 0x42, 0x95, 0x42, 0x63, 0xed, 0x5a, 0x64, 0x17, 0xe3,
	0xae, 0x06, 0xeb, 0xd7, 0x45, 0x15, 0xef, 0x6a, 0x08, 0xfe, 0xd9, 0x63, 0x2e, 0x0b, 0xb1, 0xd9,
	0x2e, 0x0b, 0x21, 0x4a, 0x7d, 0x99, 0xb5, 0x2e, 0xec, 0xea, 0x1e, 0x2c, 0x2f, 0x9e, 0x0e, 0xbe,
	0xb3, 0x3f, 0xf8, 0xce, 0xeb, 0xc1, 0x77, 0x1e, 0x8f, 0xfe, 0x60, 0x7f, 0xf4, 0x07, 0x2f, 0x47,
	0x7f, 0x70, 0x17, 0x65, 0x39, 0xf2, 0x76, 0x1d, 0x6e, 0x44, 0x19, 0xf5, 0x51, 0x9d, 0x54, 0x80,
	0x0f, 0xa2, 0xd9, 0x5a, 0xa8, 0x13, 0x55, 0x26, 0x5a, 0xfd, 0x01, 0xb9, 0x76, 0x4d, 0xae, 0xa7,
	0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x95, 0x62, 0x1c, 0xf5, 0x01, 0x00, 0x00,
}

func (m *EventEthereumTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEthereumTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEthereumTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EthTxFailed) > 0 {
		i -= len(m.EthTxFailed)
		copy(dAtA[i:], m.EthTxFailed)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.EthTxFailed)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GasUsed) > 0 {
		i -= len(m.GasUsed)
		copy(dAtA[i:], m.GasUsed)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.GasUsed)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EthHash) > 0 {
		i -= len(m.EthHash)
		copy(dAtA[i:], m.EthHash)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.EthHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventTxLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventTxLog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventTxLog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxLogs) > 0 {
		for iNdEx := len(m.TxLogs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TxLogs[iNdEx])
			copy(dAtA[i:], m.TxLogs[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.TxLogs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *EventMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxType) > 0 {
		i -= len(m.TxType)
		copy(dAtA[i:], m.TxType)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.TxType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBlockBloom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBlockBloom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBlockBloom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bloom) > 0 {
		i -= len(m.Bloom)
		copy(dAtA[i:], m.Bloom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Bloom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventEthereumTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.EthHash)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.GasUsed)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.EthTxFailed)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventTxLog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TxLogs) > 0 {
		for _, s := range m.TxLogs {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.TxType)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventBlockBloom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bloom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventEthereumTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventEthereumTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEthereumTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasUsed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthTxFailed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthTxFailed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventTxLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventTxLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventTxLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxLogs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxLogs = append(m.TxLogs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventBlockBloom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventBlockBloom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBlockBloom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bloom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
