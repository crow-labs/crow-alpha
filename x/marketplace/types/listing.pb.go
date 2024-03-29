// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: marketplace/listing.proto

package types

import (
	fmt "fmt"
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

type Order struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ListingId uint64 `protobuf:"varint,3,opt,name=listingId,proto3" json:"listingId,omitempty"`
	MaxPrice  string `protobuf:"bytes,4,opt,name=maxPrice,proto3" json:"maxPrice,omitempty"`
	Status    string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba83ca96c8a90121, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Order) GetListingId() uint64 {
	if m != nil {
		return m.ListingId
	}
	return 0
}

func (m *Order) GetMaxPrice() string {
	if m != nil {
		return m.MaxPrice
	}
	return ""
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Listing struct {
	Id            uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProducerId    uint64  `protobuf:"varint,2,opt,name=producerId,proto3" json:"producerId,omitempty"`
	Title         string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description   string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	MinPrice      string  `protobuf:"bytes,5,opt,name=minPrice,proto3" json:"minPrice,omitempty"`
	Status        string  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	PendingOrders []Order `protobuf:"bytes,7,rep,name=pendingOrders,proto3" json:"pendingOrders"`
	PurchaseOrder Order   `protobuf:"bytes,8,opt,name=purchaseOrder,proto3" json:"purchaseOrder"`
}

func (m *Listing) Reset()         { *m = Listing{} }
func (m *Listing) String() string { return proto.CompactTextString(m) }
func (*Listing) ProtoMessage()    {}
func (*Listing) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba83ca96c8a90121, []int{1}
}
func (m *Listing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Listing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Listing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Listing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listing.Merge(m, src)
}
func (m *Listing) XXX_Size() int {
	return m.Size()
}
func (m *Listing) XXX_DiscardUnknown() {
	xxx_messageInfo_Listing.DiscardUnknown(m)
}

var xxx_messageInfo_Listing proto.InternalMessageInfo

func (m *Listing) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Listing) GetProducerId() uint64 {
	if m != nil {
		return m.ProducerId
	}
	return 0
}

func (m *Listing) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Listing) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Listing) GetMinPrice() string {
	if m != nil {
		return m.MinPrice
	}
	return ""
}

func (m *Listing) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Listing) GetPendingOrders() []Order {
	if m != nil {
		return m.PendingOrders
	}
	return nil
}

func (m *Listing) GetPurchaseOrder() Order {
	if m != nil {
		return m.PurchaseOrder
	}
	return Order{}
}

func init() {
	proto.RegisterType((*Order)(nil), "crowlabs.crow.marketplace.Order")
	proto.RegisterType((*Listing)(nil), "crowlabs.crow.marketplace.Listing")
}

func init() { proto.RegisterFile("marketplace/listing.proto", fileDescriptor_ba83ca96c8a90121) }

var fileDescriptor_ba83ca96c8a90121 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x6a, 0xeb, 0x30,
	0x14, 0xc6, 0x2d, 0x27, 0xce, 0x1f, 0x85, 0x7b, 0x07, 0x11, 0x2e, 0x4a, 0xb8, 0xe8, 0x9a, 0x4c,
	0x59, 0x62, 0xc3, 0xed, 0x1b, 0x64, 0x29, 0x81, 0x40, 0x8b, 0xc7, 0x6e, 0x8e, 0x24, 0x1c, 0x51,
	0xc7, 0x32, 0x92, 0x4c, 0xd3, 0xb1, 0x6f, 0xd0, 0x57, 0xea, 0x96, 0x31, 0x63, 0xa7, 0x52, 0x92,
	0x17, 0x29, 0x96, 0x4d, 0xab, 0x52, 0x3a, 0x74, 0x3b, 0xdf, 0xa7, 0x73, 0x3e, 0xfd, 0x38, 0x1c,
	0x38, 0xd9, 0xa5, 0xea, 0x96, 0x9b, 0x32, 0x4f, 0x29, 0x8f, 0x73, 0xa1, 0x8d, 0x28, 0xb2, 0xa8,
	0x54, 0xd2, 0x48, 0x34, 0xa1, 0x4a, 0xde, 0xe5, 0xe9, 0x46, 0x47, 0x75, 0x11, 0x39, 0x8d, 0xd3,
	0x71, 0x26, 0x33, 0x69, 0xbb, 0xe2, 0xba, 0x6a, 0x06, 0x66, 0x0f, 0x00, 0x06, 0x57, 0x8a, 0x71,
	0x85, 0x7e, 0x43, 0x5f, 0x30, 0x0c, 0x42, 0x30, 0xef, 0x26, 0xbe, 0x60, 0xe8, 0x0f, 0xec, 0x55,
	0x9a, 0xab, 0x15, 0xc3, 0xbe, 0xf5, 0x5a, 0x85, 0xfe, 0xc2, 0x61, 0xfb, 0xe7, 0x8a, 0xe1, 0x8e,
	0x7d, 0xfa, 0x30, 0xd0, 0x14, 0x0e, 0x76, 0xe9, 0xfe, 0x5a, 0x09, 0xca, 0x71, 0x37, 0x04, 0xf3,
	0x61, 0xf2, 0xae, 0xeb, 0x44, 0x6d, 0x52, 0x53, 0x69, 0x1c, 0xd8, 0x97, 0x56, 0xcd, 0x9e, 0x7c,
	0xd8, 0x5f, 0x37, 0x09, 0x5f, 0x28, 0x08, 0x84, 0xa5, 0x92, 0xac, 0xa2, 0x0e, 0x89, 0xe3, 0xa0,
	0x31, 0x0c, 0x8c, 0x30, 0x39, 0xb7, 0x24, 0xc3, 0xa4, 0x11, 0x28, 0x84, 0x23, 0xc6, 0x35, 0x55,
	0xa2, 0x34, 0x42, 0x16, 0x2d, 0x88, 0x6b, 0x59, 0x4e, 0x51, 0x34, 0x9c, 0x41, 0xcb, 0xd9, 0x6a,
	0x87, 0xb3, 0xe7, 0x72, 0xa2, 0x35, 0xfc, 0x55, 0xf2, 0x82, 0x89, 0x22, 0xb3, 0x1b, 0xd3, 0xb8,
	0x1f, 0x76, 0xe6, 0xa3, 0xff, 0x61, 0xf4, 0xed, 0xd2, 0x23, 0xdb, 0xb8, 0xec, 0x1e, 0x5e, 0xfe,
	0x79, 0xc9, 0xe7, 0x61, 0x9b, 0x56, 0x29, 0xba, 0x4d, 0x35, 0xb7, 0x0e, 0x1e, 0x84, 0xe0, 0x47,
	0x69, 0xee, 0xf0, 0xf2, 0xf2, 0x70, 0x22, 0xe0, 0x78, 0x22, 0xe0, 0xf5, 0x44, 0xc0, 0xe3, 0x99,
	0x78, 0xc7, 0x33, 0xf1, 0x9e, 0xcf, 0xc4, 0xbb, 0x59, 0x64, 0xc2, 0x6c, 0xab, 0x4d, 0x44, 0xe5,
	0x2e, 0xae, 0x13, 0x17, 0x75, 0xb6, 0xad, 0xe2, 0x7d, 0xec, 0x5e, 0x92, 0xb9, 0x2f, 0xb9, 0xde,
	0xf4, 0xec, 0x5d, 0x5c, 0xbc, 0x05, 0x00, 0x00, 0xff, 0xff, 0x31, 0x38, 0x9d, 0x90, 0x65, 0x02,
	0x00, 0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintListing(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MaxPrice) > 0 {
		i -= len(m.MaxPrice)
		copy(dAtA[i:], m.MaxPrice)
		i = encodeVarintListing(dAtA, i, uint64(len(m.MaxPrice)))
		i--
		dAtA[i] = 0x22
	}
	if m.ListingId != 0 {
		i = encodeVarintListing(dAtA, i, uint64(m.ListingId))
		i--
		dAtA[i] = 0x18
	}
	if m.UserId != 0 {
		i = encodeVarintListing(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintListing(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Listing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Listing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Listing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PurchaseOrder.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintListing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.PendingOrders) > 0 {
		for iNdEx := len(m.PendingOrders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingOrders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintListing(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintListing(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MinPrice) > 0 {
		i -= len(m.MinPrice)
		copy(dAtA[i:], m.MinPrice)
		i = encodeVarintListing(dAtA, i, uint64(len(m.MinPrice)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintListing(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintListing(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ProducerId != 0 {
		i = encodeVarintListing(dAtA, i, uint64(m.ProducerId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintListing(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintListing(dAtA []byte, offset int, v uint64) int {
	offset -= sovListing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovListing(uint64(m.Id))
	}
	if m.UserId != 0 {
		n += 1 + sovListing(uint64(m.UserId))
	}
	if m.ListingId != 0 {
		n += 1 + sovListing(uint64(m.ListingId))
	}
	l = len(m.MaxPrice)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	return n
}

func (m *Listing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovListing(uint64(m.Id))
	}
	if m.ProducerId != 0 {
		n += 1 + sovListing(uint64(m.ProducerId))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = len(m.MinPrice)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	if len(m.PendingOrders) > 0 {
		for _, e := range m.PendingOrders {
			l = e.Size()
			n += 1 + l + sovListing(uint64(l))
		}
	}
	l = m.PurchaseOrder.Size()
	n += 1 + l + sovListing(uint64(l))
	return n
}

func sovListing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListing(x uint64) (n int) {
	return sovListing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListing
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListingId", wireType)
			}
			m.ListingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ListingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListing
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
func (m *Listing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListing
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
			return fmt.Errorf("proto: Listing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Listing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducerId", wireType)
			}
			m.ProducerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProducerId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingOrders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingOrders = append(m.PendingOrders, Order{})
			if err := m.PendingOrders[len(m.PendingOrders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PurchaseOrder", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PurchaseOrder.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListing
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
func skipListing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListing
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
					return 0, ErrIntOverflowListing
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
					return 0, ErrIntOverflowListing
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
				return 0, ErrInvalidLengthListing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupListing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthListing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthListing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupListing = fmt.Errorf("proto: unexpected end of group")
)
