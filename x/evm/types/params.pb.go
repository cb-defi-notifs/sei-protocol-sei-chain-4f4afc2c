// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Params defines the parameters for the module
type Params struct {
	// string base_denom = 1 [
	//   (gogoproto.moretags)   = "yaml:\"base_denom\"",
	//   (gogoproto.jsontag) = "base_denom"
	// ];
	PriorityNormalizer github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=priority_normalizer,json=priorityNormalizer,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"priority_normalizer" yaml:"priority_normalizer"`
	BaseFeePerGas      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=base_fee_per_gas,json=baseFeePerGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_fee_per_gas" yaml:"base_fee_per_gas"`
	MinimumFeePerGas   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=minimum_fee_per_gas,json=minimumFeePerGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"minimum_fee_per_gas" yaml:"minimum_fee_per_gas"`
	// ChainConfig chain_config = 5 [(gogoproto.moretags) = "yaml:\"chain_config\"", (gogoproto.nullable) = false];
	//   string chain_id = 6 [
	//   (gogoproto.moretags)   = "yaml:\"chain_id\"",
	//   (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
	//   (gogoproto.nullable)   = false,
	//   (gogoproto.jsontag) = "chain_id"
	// ];
	// repeated string whitelisted_codehashes_bank_send = 7 [
	//   (gogoproto.moretags)   = "yaml:\"whitelisted_codehashes_bank_send\"",
	//   (gogoproto.jsontag) = "whitelisted_codehashes_bank_send"
	// ];
	WhitelistedCwCodeHashesForDelegateCall [][]byte                               `protobuf:"bytes,8,rep,name=whitelisted_cw_code_hashes_for_delegate_call,json=whitelistedCwCodeHashesForDelegateCall,proto3" json:"whitelisted_cw_code_hashes_for_delegate_call" yaml:"whitelisted_cw_code_hashes_for_delegate_call"`
	DeliverTxHookWasmGasLimit              uint64                                 `protobuf:"varint,9,opt,name=deliver_tx_hook_wasm_gas_limit,json=deliverTxHookWasmGasLimit,proto3" json:"deliver_tx_hook_wasm_gas_limit,omitempty"`
	MaxDynamicBaseFeeUpwardAdjustment      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=max_dynamic_base_fee_upward_adjustment,json=maxDynamicBaseFeeUpwardAdjustment,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_dynamic_base_fee_upward_adjustment" yaml:"max_dynamic_base_fee_upward_adjustment"`
	MaxDynamicBaseFeeDownwardAdjustment    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=max_dynamic_base_fee_downward_adjustment,json=maxDynamicBaseFeeDownwardAdjustment,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_dynamic_base_fee_downward_adjustment" yaml:"max_dynamic_base_fee_downward_adjustment"`
	TargetGasUsedPerBlock                  uint64                                 `protobuf:"varint,12,opt,name=target_gas_used_per_block,json=targetGasUsedPerBlock,proto3" json:"target_gas_used_per_block,omitempty"`
	MaximumFeePerGas                       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=maximum_fee_per_gas,json=maximumFeePerGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"maximum_fee_per_gas" yaml:"maximum_fee_per_gas"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9272f3679901ea94, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetWhitelistedCwCodeHashesForDelegateCall() [][]byte {
	if m != nil {
		return m.WhitelistedCwCodeHashesForDelegateCall
	}
	return nil
}

func (m *Params) GetDeliverTxHookWasmGasLimit() uint64 {
	if m != nil {
		return m.DeliverTxHookWasmGasLimit
	}
	return 0
}

func (m *Params) GetTargetGasUsedPerBlock() uint64 {
	if m != nil {
		return m.TargetGasUsedPerBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "seiprotocol.seichain.evm.Params")
}

func init() { proto.RegisterFile("evm/params.proto", fileDescriptor_9272f3679901ea94) }

var fileDescriptor_9272f3679901ea94 = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x41, 0x6b, 0xdb, 0x4a,
	0x10, 0xb6, 0x5e, 0xf2, 0xc2, 0xcb, 0xbe, 0x04, 0x82, 0xf2, 0x1e, 0x55, 0x72, 0x90, 0x52, 0x15,
	0x82, 0x0f, 0xb5, 0x55, 0x08, 0x85, 0x90, 0x5b, 0x1c, 0x13, 0xe7, 0x50, 0x8a, 0x11, 0x4d, 0x0b,
	0x85, 0xb2, 0xac, 0xa5, 0x89, 0xbc, 0xb5, 0xd6, 0x6b, 0x76, 0xd7, 0xb1, 0xdd, 0x1f, 0x50, 0xe8,
	0xa1, 0x50, 0x42, 0x0f, 0x3d, 0xf6, 0xcf, 0x14, 0x72, 0xcc, 0xb1, 0x14, 0x2a, 0x4a, 0x42, 0x2e,
	0x39, 0xfa, 0x17, 0x14, 0xad, 0x1c, 0xc7, 0x8d, 0x4d, 0x89, 0x73, 0xe9, 0x25, 0x27, 0xad, 0x66,
	0xbe, 0x9d, 0x9d, 0x6f, 0x66, 0xbf, 0x61, 0xd1, 0x12, 0x1c, 0x32, 0xaf, 0x45, 0x04, 0x61, 0xb2,
	0xd8, 0x12, 0x5c, 0x71, 0xd3, 0x92, 0x40, 0xf5, 0x2a, 0xe0, 0x71, 0x51, 0x02, 0x0d, 0xea, 0x84,
	0x36, 0x8b, 0x70, 0xc8, 0x56, 0xff, 0x8b, 0x78, 0xc4, 0xb5, 0xcb, 0x4b, 0x57, 0x19, 0xde, 0x3d,
	0x47, 0x68, 0xae, 0xaa, 0x03, 0x98, 0x1f, 0x0d, 0xb4, 0xdc, 0x12, 0x94, 0x0b, 0xaa, 0x7a, 0xb8,
	0xc9, 0x05, 0x23, 0x31, 0x7d, 0x03, 0xc2, 0xfa, 0x6b, 0xcd, 0xc8, 0xcf, 0x97, 0x82, 0xe3, 0xc4,
	0xc9, 0x7d, 0x4b, 0x9c, 0xf5, 0x88, 0xaa, 0x7a, 0xbb, 0x56, 0x0c, 0x38, 0xf3, 0x02, 0x2e, 0x19,
	0x97, 0x83, 0x4f, 0x41, 0x86, 0x0d, 0x4f, 0xf5, 0x5a, 0x20, 0x8b, 0x65, 0x08, 0x2e, 0x12, 0x67,
	0x52, 0xb0, 0x7e, 0xe2, 0xac, 0xf6, 0x08, 0x8b, 0xb7, 0xdc, 0x09, 0x4e, 0xd7, 0x37, 0x2f, 0xad,
	0x4f, 0x87, 0x46, 0xf3, 0xad, 0x81, 0x96, 0x6a, 0x44, 0x02, 0x3e, 0x00, 0xc0, 0x2d, 0x10, 0x38,
	0x22, 0xd2, 0x9a, 0xd1, 0x39, 0xbd, 0x9a, 0x3a, 0xa7, 0xb1, 0x48, 0xfd, 0xc4, 0xb9, 0x97, 0x25,
	0x74, 0xdd, 0xe3, 0xfa, 0x8b, 0xa9, 0x69, 0x17, 0xa0, 0x0a, 0xa2, 0x42, 0xa4, 0x79, 0x64, 0xa0,
	0x65, 0x46, 0x9b, 0x94, 0xb5, 0xd9, 0x2f, 0xb9, 0xcc, 0xde, 0xb6, 0x3e, 0x13, 0x82, 0x5d, 0xd5,
	0x67, 0x82, 0xd3, 0xf5, 0x97, 0x06, 0xd6, 0xab, 0xa4, 0xbe, 0x18, 0xe8, 0x61, 0xa7, 0x4e, 0x15,
	0xc4, 0x54, 0x2a, 0x08, 0x71, 0xd0, 0xc1, 0x01, 0x0f, 0x01, 0xd7, 0x89, 0xac, 0x83, 0xc4, 0x07,
	0x5c, 0xe0, 0x10, 0x62, 0x88, 0x88, 0x02, 0x1c, 0x90, 0x38, 0xb6, 0xfe, 0x59, 0x9b, 0xc9, 0x2f,
	0x94, 0xa2, 0x8b, 0xc4, 0x99, 0x6a, 0x5f, 0x3f, 0x71, 0x36, 0xb2, 0xc4, 0xa6, 0xd9, 0xe5, 0xfa,
	0xeb, 0x23, 0xf0, 0x9d, 0xce, 0x0e, 0x0f, 0x61, 0x4f, 0x63, 0x77, 0xb9, 0x28, 0x0f, 0x90, 0x3b,
	0x24, 0x8e, 0xcd, 0x6d, 0x64, 0x87, 0x10, 0xd3, 0x43, 0x10, 0x58, 0x75, 0x71, 0x9d, 0xf3, 0x06,
	0xee, 0x10, 0xc9, 0x52, 0xda, 0x38, 0xa6, 0x8c, 0x2a, 0x6b, 0x7e, 0xcd, 0xc8, 0xcf, 0xfa, 0x2b,
	0x03, 0xd4, 0xb3, 0xee, 0x1e, 0xe7, 0x8d, 0x17, 0x44, 0xb2, 0x0a, 0x91, 0x4f, 0x52, 0x80, 0xf9,
	0xdd, 0x40, 0xeb, 0x8c, 0x74, 0x71, 0xd8, 0x6b, 0x12, 0x46, 0x03, 0x3c, 0x6c, 0x68, 0xbb, 0xd5,
	0x21, 0x22, 0xc4, 0x24, 0x7c, 0xdd, 0x96, 0x8a, 0x41, 0x53, 0x59, 0x48, 0xb7, 0xec, 0x9d, 0x31,
	0x75, 0xcf, 0x6e, 0x78, 0x40, 0x3f, 0x71, 0x0a, 0x83, 0x36, 0xde, 0x08, 0xef, 0xfa, 0xf7, 0x19,
	0xe9, 0x96, 0x33, 0x5c, 0x29, 0xbb, 0x75, 0xfb, 0x1a, 0xb4, 0x3d, 0xc4, 0x98, 0xe7, 0x06, 0xca,
	0x4f, 0x0c, 0x17, 0xf2, 0x4e, 0xf3, 0x3a, 0xc3, 0x7f, 0x35, 0xc3, 0xf7, 0xd3, 0x33, 0xbc, 0xf1,
	0x11, 0xfd, 0xc4, 0xf1, 0x7e, 0xc3, 0x71, 0xc2, 0x0e, 0xd7, 0x7f, 0x30, 0xc6, 0xb2, 0x3c, 0x80,
	0x8d, 0xf0, 0xdc, 0x44, 0x2b, 0x8a, 0x88, 0x08, 0x94, 0x6e, 0x7e, 0x5b, 0x42, 0xa8, 0x05, 0x50,
	0x8b, 0x79, 0xd0, 0xb0, 0x16, 0xf4, 0x2d, 0xf8, 0x3f, 0x03, 0x54, 0x88, 0xdc, 0x97, 0x10, 0x56,
	0x41, 0x94, 0x52, 0x67, 0xa6, 0x50, 0xd2, 0x1d, 0x53, 0xe8, 0xe2, 0xad, 0x15, 0x3a, 0x1e, 0x6c,
	0x44, 0xa1, 0xe3, 0xce, 0x54, 0xa1, 0x99, 0x75, 0xa8, 0xd0, 0xad, 0xd9, 0x4f, 0x9f, 0x9d, 0x9c,
	0x7b, 0xf4, 0x37, 0x5a, 0xcc, 0xe6, 0x6c, 0x55, 0xc0, 0xf3, 0xc7, 0x9b, 0x8f, 0xee, 0xc6, 0xed,
	0xdd, 0xb8, 0xfd, 0x63, 0xe3, 0x36, 0xbb, 0x94, 0xa5, 0xca, 0xf1, 0xa9, 0x6d, 0x9c, 0x9c, 0xda,
	0xc6, 0x8f, 0x53, 0xdb, 0xf8, 0x70, 0x66, 0xe7, 0x4e, 0xce, 0xec, 0xdc, 0xd7, 0x33, 0x3b, 0xf7,
	0xb2, 0x30, 0x52, 0x56, 0x09, 0xb4, 0x70, 0xf9, 0xa4, 0xd0, 0x3f, 0xfa, 0x4d, 0xe1, 0x75, 0xbd,
	0xf4, 0xf1, 0xa1, 0x2b, 0x5c, 0x9b, 0xd3, 0xfe, 0x8d, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19,
	0x33, 0xf5, 0x29, 0x90, 0x08, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaximumFeePerGas.Size()
		i -= size
		if _, err := m.MaximumFeePerGas.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	if m.TargetGasUsedPerBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TargetGasUsedPerBlock))
		i--
		dAtA[i] = 0x60
	}
	{
		size := m.MaxDynamicBaseFeeDownwardAdjustment.Size()
		i -= size
		if _, err := m.MaxDynamicBaseFeeDownwardAdjustment.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.MaxDynamicBaseFeeUpwardAdjustment.Size()
		i -= size
		if _, err := m.MaxDynamicBaseFeeUpwardAdjustment.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.DeliverTxHookWasmGasLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DeliverTxHookWasmGasLimit))
		i--
		dAtA[i] = 0x48
	}
	if len(m.WhitelistedCwCodeHashesForDelegateCall) > 0 {
		for iNdEx := len(m.WhitelistedCwCodeHashesForDelegateCall) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WhitelistedCwCodeHashesForDelegateCall[iNdEx])
			copy(dAtA[i:], m.WhitelistedCwCodeHashesForDelegateCall[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.WhitelistedCwCodeHashesForDelegateCall[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	{
		size := m.MinimumFeePerGas.Size()
		i -= size
		if _, err := m.MinimumFeePerGas.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.BaseFeePerGas.Size()
		i -= size
		if _, err := m.BaseFeePerGas.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.PriorityNormalizer.Size()
		i -= size
		if _, err := m.PriorityNormalizer.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PriorityNormalizer.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.BaseFeePerGas.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinimumFeePerGas.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.WhitelistedCwCodeHashesForDelegateCall) > 0 {
		for _, b := range m.WhitelistedCwCodeHashesForDelegateCall {
			l = len(b)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.DeliverTxHookWasmGasLimit != 0 {
		n += 1 + sovParams(uint64(m.DeliverTxHookWasmGasLimit))
	}
	l = m.MaxDynamicBaseFeeUpwardAdjustment.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxDynamicBaseFeeDownwardAdjustment.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.TargetGasUsedPerBlock != 0 {
		n += 1 + sovParams(uint64(m.TargetGasUsedPerBlock))
	}
	l = m.MaximumFeePerGas.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriorityNormalizer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PriorityNormalizer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFeePerGas", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseFeePerGas.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumFeePerGas", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinimumFeePerGas.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedCwCodeHashesForDelegateCall", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistedCwCodeHashesForDelegateCall = append(m.WhitelistedCwCodeHashesForDelegateCall, make([]byte, postIndex-iNdEx))
			copy(m.WhitelistedCwCodeHashesForDelegateCall[len(m.WhitelistedCwCodeHashesForDelegateCall)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeliverTxHookWasmGasLimit", wireType)
			}
			m.DeliverTxHookWasmGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeliverTxHookWasmGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDynamicBaseFeeUpwardAdjustment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxDynamicBaseFeeUpwardAdjustment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDynamicBaseFeeDownwardAdjustment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxDynamicBaseFeeDownwardAdjustment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetGasUsedPerBlock", wireType)
			}
			m.TargetGasUsedPerBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetGasUsedPerBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaximumFeePerGas", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaximumFeePerGas.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
