// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/tradeshield/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the tradeshield module's genesis state.
type GenesisState struct {
	Params                     Params           `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PendingSpotOrderList       []SpotOrder      `protobuf:"bytes,2,rep,name=pendingSpotOrderList,proto3" json:"pendingSpotOrderList"`
	PendingSpotOrderCount      uint64           `protobuf:"varint,3,opt,name=pendingSpotOrderCount,proto3" json:"pendingSpotOrderCount,omitempty"`
	PendingPerpetualOrderList  []PerpetualOrder `protobuf:"bytes,4,rep,name=pendingPerpetualOrderList,proto3" json:"pendingPerpetualOrderList"`
	PendingPerpetualOrderCount uint64           `protobuf:"varint,5,opt,name=pendingPerpetualOrderCount,proto3" json:"pendingPerpetualOrderCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_da92b96d2e7e2f05, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPendingSpotOrderList() []SpotOrder {
	if m != nil {
		return m.PendingSpotOrderList
	}
	return nil
}

func (m *GenesisState) GetPendingSpotOrderCount() uint64 {
	if m != nil {
		return m.PendingSpotOrderCount
	}
	return 0
}

func (m *GenesisState) GetPendingPerpetualOrderList() []PerpetualOrder {
	if m != nil {
		return m.PendingPerpetualOrderList
	}
	return nil
}

func (m *GenesisState) GetPendingPerpetualOrderCount() uint64 {
	if m != nil {
		return m.PendingPerpetualOrderCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "elys.tradeshield.GenesisState")
}

func init() { proto.RegisterFile("elys/tradeshield/genesis.proto", fileDescriptor_da92b96d2e7e2f05) }

var fileDescriptor_da92b96d2e7e2f05 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcd, 0xa9, 0x2c,
	0xd6, 0x2f, 0x29, 0x4a, 0x4c, 0x49, 0x2d, 0xce, 0xc8, 0x4c, 0xcd, 0x49, 0xd1, 0x4f, 0x4f, 0xcd,
	0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x00, 0xc9, 0xeb, 0x21,
	0xc9, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x25, 0xf5, 0x41, 0x2c, 0x88, 0x3a, 0x29, 0x59,
	0x0c, 0x73, 0x0a, 0x12, 0x8b, 0x12, 0x73, 0xa1, 0xc6, 0x48, 0xc9, 0x60, 0x48, 0x97, 0x54, 0x16,
	0xa4, 0x42, 0x65, 0x95, 0x7e, 0x32, 0x71, 0xf1, 0xb8, 0x43, 0xac, 0x0d, 0x2e, 0x49, 0x2c, 0x49,
	0x15, 0x32, 0xe3, 0x62, 0x83, 0x68, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd0, 0x43,
	0x77, 0x86, 0x5e, 0x00, 0x58, 0xde, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0xa8, 0x6a, 0xa1,
	0x50, 0x2e, 0x91, 0x82, 0xd4, 0xbc, 0x94, 0xcc, 0xbc, 0xf4, 0xe0, 0x82, 0xfc, 0x12, 0xff, 0xa2,
	0x94, 0xd4, 0x22, 0x9f, 0xcc, 0xe2, 0x12, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x69, 0x4c,
	0x53, 0xe0, 0xca, 0xa0, 0x06, 0x61, 0xd5, 0x2e, 0x64, 0xc2, 0x25, 0x8a, 0x2e, 0xee, 0x9c, 0x5f,
	0x9a, 0x57, 0x22, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x12, 0x84, 0x5d, 0x52, 0x28, 0x85, 0x4b, 0x12,
	0x2a, 0x11, 0x90, 0x5a, 0x54, 0x90, 0x5a, 0x52, 0x9a, 0x98, 0x83, 0x70, 0x11, 0x0b, 0xd8, 0x45,
	0x0a, 0x58, 0xfc, 0x85, 0xa2, 0x16, 0xea, 0x2c, 0xdc, 0x06, 0x09, 0xd9, 0x71, 0x49, 0x61, 0x95,
	0x84, 0x38, 0x90, 0x15, 0xec, 0x40, 0x3c, 0x2a, 0x9c, 0xbc, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8,
	0xf1, 0x58, 0x8e, 0x21, 0xca, 0x20, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x1f, 0xe4, 0x4c, 0xdd, 0xbc, 0xd4, 0x92, 0xf2, 0xfc, 0xa2, 0x6c, 0x30, 0x47, 0xbf, 0x02, 0x33,
	0x36, 0x93, 0xd8, 0xc0, 0xd1, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x77, 0x68, 0xf6,
	0x55, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PendingPerpetualOrderCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PendingPerpetualOrderCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PendingPerpetualOrderList) > 0 {
		for iNdEx := len(m.PendingPerpetualOrderList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingPerpetualOrderList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.PendingSpotOrderCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PendingSpotOrderCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PendingSpotOrderList) > 0 {
		for iNdEx := len(m.PendingSpotOrderList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingSpotOrderList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.PendingSpotOrderList) > 0 {
		for _, e := range m.PendingSpotOrderList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PendingSpotOrderCount != 0 {
		n += 1 + sovGenesis(uint64(m.PendingSpotOrderCount))
	}
	if len(m.PendingPerpetualOrderList) > 0 {
		for _, e := range m.PendingPerpetualOrderList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PendingPerpetualOrderCount != 0 {
		n += 1 + sovGenesis(uint64(m.PendingPerpetualOrderCount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingSpotOrderList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingSpotOrderList = append(m.PendingSpotOrderList, SpotOrder{})
			if err := m.PendingSpotOrderList[len(m.PendingSpotOrderList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingSpotOrderCount", wireType)
			}
			m.PendingSpotOrderCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PendingSpotOrderCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingPerpetualOrderList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingPerpetualOrderList = append(m.PendingPerpetualOrderList, PerpetualOrder{})
			if err := m.PendingPerpetualOrderList[len(m.PendingPerpetualOrderList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingPerpetualOrderCount", wireType)
			}
			m.PendingPerpetualOrderCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PendingPerpetualOrderCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)