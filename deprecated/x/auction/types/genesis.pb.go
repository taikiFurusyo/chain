// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// GenesisState defines the auction module's genesis state.
type GenesisState struct {
	NextAuctionId uint64       `protobuf:"varint,1,opt,name=next_auction_id,json=nextAuctionId,proto3" json:"next_auction_id,omitempty" yaml:"next_auction_id"`
	Params        Params       `protobuf:"bytes,2,opt,name=params,proto3" json:"params" yaml:"params"`
	Auctions      []*types.Any `protobuf:"bytes,3,rep,name=auctions,proto3" json:"auctions,omitempty" yaml:"auctions"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec0c841528c842a6, []int{0}
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

func (m *GenesisState) GetNextAuctionId() uint64 {
	if m != nil {
		return m.NextAuctionId
	}
	return 0
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetAuctions() []*types.Any {
	if m != nil {
		return m.Auctions
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ununifi.auction.GenesisState")
}

func init() { proto.RegisterFile("auction/genesis.proto", fileDescriptor_ec0c841528c842a6) }

var fileDescriptor_ec0c841528c842a6 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x6e, 0xc2, 0x30,
	0x14, 0x86, 0xe3, 0x52, 0xa1, 0x2a, 0x14, 0x21, 0xa5, 0xd0, 0x52, 0x06, 0x07, 0x65, 0xca, 0x64,
	0x0b, 0xba, 0x75, 0x23, 0x52, 0xa9, 0xba, 0x55, 0x54, 0x2c, 0x5d, 0x90, 0x13, 0x4c, 0xb0, 0x44,
	0xec, 0x08, 0x3b, 0x15, 0xb9, 0x45, 0x8f, 0xc5, 0xc8, 0xd8, 0x29, 0x6a, 0xe1, 0x06, 0x9c, 0xa0,
	0x22, 0x76, 0x18, 0x98, 0xec, 0xf7, 0xbe, 0xff, 0x7f, 0xcf, 0xfe, 0xed, 0x0e, 0xc9, 0x22, 0xc5,
	0x04, 0xc7, 0x31, 0xe5, 0x54, 0x32, 0x89, 0xd2, 0xb5, 0x50, 0xc2, 0x69, 0x65, 0x3c, 0xe3, 0x6c,
	0xc1, 0x90, 0xc1, 0xbd, 0x76, 0x2c, 0x62, 0x51, 0x32, 0x7c, 0xba, 0x69, 0x59, 0xef, 0x31, 0x16,
	0x22, 0x5e, 0x51, 0x5c, 0x56, 0x61, 0xb6, 0xc0, 0x84, 0xe7, 0x06, 0xb9, 0x97, 0x48, 0xb1, 0x84,
	0x4a, 0x45, 0x92, 0xd4, 0x08, 0x60, 0x24, 0x64, 0x22, 0x24, 0x0e, 0x89, 0xa4, 0xf8, 0x6b, 0x10,
	0x52, 0x45, 0x06, 0x38, 0x12, 0x8c, 0x57, 0xb3, 0x35, 0x9f, 0xe9, 0xa5, 0xba, 0x30, 0xe8, 0xfc,
	0x68, 0x73, 0xea, 0xb6, 0xf7, 0x07, 0xec, 0xdb, 0x57, 0xfd, 0x8d, 0x0f, 0x45, 0x14, 0x75, 0x02,
	0xbb, 0xc5, 0xe9, 0x46, 0xcd, 0x8c, 0x6c, 0xc6, 0xe6, 0x5d, 0xd0, 0x07, 0xfe, 0x75, 0xd0, 0x3b,
	0x16, 0xee, 0x7d, 0x4e, 0x92, 0xd5, 0xb3, 0x77, 0x21, 0xf0, 0x26, 0xcd, 0x53, 0x67, 0xa4, 0x1b,
	0x6f, 0x73, 0x67, 0x6c, 0xd7, 0x53, 0xb2, 0x26, 0x89, 0xec, 0x5e, 0xf5, 0x81, 0xdf, 0x18, 0x3e,
	0xa0, 0x8b, 0x68, 0xd0, 0x7b, 0x89, 0x83, 0xce, 0xb6, 0x70, 0xad, 0x63, 0xe1, 0x36, 0xf5, 0x5c,
	0x6d, 0xf2, 0x26, 0xc6, 0xed, 0xbc, 0xd8, 0x37, 0xc6, 0x20, 0xbb, 0xb5, 0x7e, 0xcd, 0x6f, 0x0c,
	0xdb, 0x48, 0x47, 0x84, 0xaa, 0x88, 0xd0, 0x88, 0xe7, 0xc1, 0xdd, 0xb1, 0x70, 0x5b, 0x7a, 0x44,
	0xa5, 0xf7, 0x26, 0x67, 0x6b, 0x10, 0x6c, 0xf7, 0x10, 0xec, 0xf6, 0x10, 0xfc, 0xee, 0x21, 0xf8,
	0x3e, 0x40, 0x6b, 0x77, 0x80, 0xd6, 0xcf, 0x01, 0x5a, 0x9f, 0x7e, 0xcc, 0xd4, 0x32, 0x0b, 0x51,
	0x24, 0x12, 0x3c, 0xe5, 0x53, 0xce, 0xc6, 0x0c, 0x47, 0x4b, 0xc2, 0x38, 0xde, 0x54, 0x39, 0x61,
	0x95, 0xa7, 0x54, 0x86, 0xf5, 0x72, 0xe1, 0xd3, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0x51,
	0xeb, 0x87, 0xfc, 0x01, 0x00, 0x00,
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
	if len(m.Auctions) > 0 {
		for iNdEx := len(m.Auctions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Auctions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
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
	dAtA[i] = 0x12
	if m.NextAuctionId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextAuctionId))
		i--
		dAtA[i] = 0x8
	}
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
	if m.NextAuctionId != 0 {
		n += 1 + sovGenesis(uint64(m.NextAuctionId))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Auctions) > 0 {
		for _, e := range m.Auctions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextAuctionId", wireType)
			}
			m.NextAuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextAuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auctions", wireType)
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
			m.Auctions = append(m.Auctions, &types.Any{})
			if err := m.Auctions[len(m.Auctions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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