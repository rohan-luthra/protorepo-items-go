// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

package item

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Type int32

const (
	Type_CHECKBOX     Type = 0
	Type_RADIO_BUTTON Type = 1
	Type_SLIDER       Type = 2
)

var Type_name = map[int32]string{
	0: "CHECKBOX",
	1: "RADIO_BUTTON",
	2: "SLIDER",
}

var Type_value = map[string]int32{
	"CHECKBOX":     0,
	"RADIO_BUTTON": 1,
	"SLIDER":       2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{0}
}

type Timings int32

const (
	Timings_BREAKFAST Timings = 0
	Timings_LUNCH     Timings = 1
	Timings_SNACKS    Timings = 2
	Timings_DINNER    Timings = 3
)

var Timings_name = map[int32]string{
	0: "BREAKFAST",
	1: "LUNCH",
	2: "SNACKS",
	3: "DINNER",
}

var Timings_value = map[string]int32{
	"BREAKFAST": 0,
	"LUNCH":     1,
	"SNACKS":    2,
	"DINNER":    3,
}

func (x Timings) String() string {
	return proto.EnumName(Timings_name, int32(x))
}

func (Timings) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{1}
}

type SpecialCategory int32

const (
	SpecialCategory_CHEFS_SPECIAL  SpecialCategory = 0
	SpecialCategory_TODAYS_SPECIAL SpecialCategory = 1
)

var SpecialCategory_name = map[int32]string{
	0: "CHEFS_SPECIAL",
	1: "TODAYS_SPECIAL",
}

var SpecialCategory_value = map[string]int32{
	"CHEFS_SPECIAL":  0,
	"TODAYS_SPECIAL": 1,
}

func (x SpecialCategory) String() string {
	return proto.EnumName(SpecialCategory_name, int32(x))
}

func (SpecialCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{2}
}

type AddonItem struct {
	// TODO: add multiple language to name & description
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                float32  `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Available            bool     `protobuf:"varint,4,opt,name=available,proto3" json:"available,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddonItem) Reset()         { *m = AddonItem{} }
func (m *AddonItem) String() string { return proto.CompactTextString(m) }
func (*AddonItem) ProtoMessage()    {}
func (*AddonItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{0}
}

func (m *AddonItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddonItem.Unmarshal(m, b)
}
func (m *AddonItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddonItem.Marshal(b, m, deterministic)
}
func (m *AddonItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddonItem.Merge(m, src)
}
func (m *AddonItem) XXX_Size() int {
	return xxx_messageInfo_AddonItem.Size(m)
}
func (m *AddonItem) XXX_DiscardUnknown() {
	xxx_messageInfo_AddonItem.DiscardUnknown(m)
}

var xxx_messageInfo_AddonItem proto.InternalMessageInfo

func (m *AddonItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddonItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddonItem) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *AddonItem) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *AddonItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Addon struct {
	// TODO: add multiple language to name & description
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	RestId               string       `protobuf:"bytes,4,opt,name=rest_id,json=restId,proto3" json:"rest_id,omitempty"`
	AddonItem            []*AddonItem `protobuf:"bytes,5,rep,name=addon_item,json=addonItem,proto3" json:"addon_item,omitempty"`
	XType                Type         `protobuf:"varint,6,opt,name=__type,json=Type,proto3,enum=item.Type" json:"__type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Addon) Reset()         { *m = Addon{} }
func (m *Addon) String() string { return proto.CompactTextString(m) }
func (*Addon) ProtoMessage()    {}
func (*Addon) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{1}
}

func (m *Addon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addon.Unmarshal(m, b)
}
func (m *Addon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addon.Marshal(b, m, deterministic)
}
func (m *Addon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addon.Merge(m, src)
}
func (m *Addon) XXX_Size() int {
	return xxx_messageInfo_Addon.Size(m)
}
func (m *Addon) XXX_DiscardUnknown() {
	xxx_messageInfo_Addon.DiscardUnknown(m)
}

var xxx_messageInfo_Addon proto.InternalMessageInfo

func (m *Addon) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Addon) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Addon) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Addon) GetRestId() string {
	if m != nil {
		return m.RestId
	}
	return ""
}

func (m *Addon) GetAddonItem() []*AddonItem {
	if m != nil {
		return m.AddonItem
	}
	return nil
}

func (m *Addon) GetXType() Type {
	if m != nil {
		return m.XType
	}
	return Type_CHECKBOX
}

type Item struct {
	// TODO: add multiple language to name & description
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price                float32              `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Ratings              float32              `protobuf:"fixed32,5,opt,name=ratings,proto3" json:"ratings,omitempty"`
	RatingsCount         int32                `protobuf:"varint,6,opt,name=ratings_count,json=ratingsCount,proto3" json:"ratings_count,omitempty"`
	OrderCount           int32                `protobuf:"varint,7,opt,name=order_count,json=orderCount,proto3" json:"order_count,omitempty"`
	CategoryId           string               `protobuf:"bytes,8,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	SubCategoryId        string               `protobuf:"bytes,9,opt,name=sub_category_id,json=subCategoryId,proto3" json:"sub_category_id,omitempty"`
	SplCategories        []SpecialCategory    `protobuf:"varint,10,rep,packed,name=spl_categories,json=splCategories,proto3,enum=item.SpecialCategory" json:"spl_categories,omitempty"`
	Tags                 []string             `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	Images               []string             `protobuf:"bytes,12,rep,name=images,proto3" json:"images,omitempty"`
	Timings              []Timings            `protobuf:"varint,13,rep,packed,name=timings,proto3,enum=item.Timings" json:"timings,omitempty"`
	Priority             int32                `protobuf:"varint,14,opt,name=priority,proto3" json:"priority,omitempty"`
	Available            bool                 `protobuf:"varint,15,opt,name=available,proto3" json:"available,omitempty"`
	Veg                  bool                 `protobuf:"varint,16,opt,name=veg,proto3" json:"veg,omitempty"`
	Addons               []*Addon             `protobuf:"bytes,17,rep,name=addons,proto3" json:"addons,omitempty"`
	RestId               string               `protobuf:"bytes,18,opt,name=rest_id,json=restId,proto3" json:"rest_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{2}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Item) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetRatings() float32 {
	if m != nil {
		return m.Ratings
	}
	return 0
}

func (m *Item) GetRatingsCount() int32 {
	if m != nil {
		return m.RatingsCount
	}
	return 0
}

func (m *Item) GetOrderCount() int32 {
	if m != nil {
		return m.OrderCount
	}
	return 0
}

func (m *Item) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

func (m *Item) GetSubCategoryId() string {
	if m != nil {
		return m.SubCategoryId
	}
	return ""
}

func (m *Item) GetSplCategories() []SpecialCategory {
	if m != nil {
		return m.SplCategories
	}
	return nil
}

func (m *Item) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Item) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Item) GetTimings() []Timings {
	if m != nil {
		return m.Timings
	}
	return nil
}

func (m *Item) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Item) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Item) GetVeg() bool {
	if m != nil {
		return m.Veg
	}
	return false
}

func (m *Item) GetAddons() []*Addon {
	if m != nil {
		return m.Addons
	}
	return nil
}

func (m *Item) GetRestId() string {
	if m != nil {
		return m.RestId
	}
	return ""
}

func (m *Item) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Item) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Id struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{3}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RestId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestId) Reset()         { *m = RestId{} }
func (m *RestId) String() string { return proto.CompactTextString(m) }
func (*RestId) ProtoMessage()    {}
func (*RestId) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{4}
}

func (m *RestId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestId.Unmarshal(m, b)
}
func (m *RestId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestId.Marshal(b, m, deterministic)
}
func (m *RestId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestId.Merge(m, src)
}
func (m *RestId) XXX_Size() int {
	return xxx_messageInfo_RestId.Size(m)
}
func (m *RestId) XXX_DiscardUnknown() {
	xxx_messageInfo_RestId.DiscardUnknown(m)
}

var xxx_messageInfo_RestId proto.InternalMessageInfo

func (m *RestId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("item.Type", Type_name, Type_value)
	proto.RegisterEnum("item.Timings", Timings_name, Timings_value)
	proto.RegisterEnum("item.SpecialCategory", SpecialCategory_name, SpecialCategory_value)
	proto.RegisterType((*AddonItem)(nil), "item.AddonItem")
	proto.RegisterType((*Addon)(nil), "item.Addon")
	proto.RegisterType((*Item)(nil), "item.Item")
	proto.RegisterType((*Id)(nil), "item.Id")
	proto.RegisterType((*RestId)(nil), "item.RestId")
}

func init() {
	proto.RegisterFile("item.proto", fileDescriptor_6007f868cf6553df)
}

var fileDescriptor_6007f868cf6553df = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xda, 0x4a,
	0x14, 0x8d, 0x0d, 0x18, 0x7c, 0xf9, 0x72, 0xe6, 0xe5, 0xbd, 0x37, 0x42, 0x4f, 0x8a, 0x45, 0x9e,
	0x5a, 0xc4, 0x82, 0x54, 0x74, 0xd3, 0xaa, 0xdd, 0x38, 0x86, 0x24, 0x56, 0x22, 0xa8, 0x06, 0x22,
	0xb5, 0x2b, 0xcb, 0xe0, 0x29, 0x1a, 0x09, 0x63, 0xcb, 0x1e, 0x22, 0xf1, 0x03, 0xfa, 0x37, 0xfa,
	0x3b, 0xba, 0xef, 0x1f, 0xab, 0x66, 0x6c, 0x83, 0x43, 0x16, 0xc9, 0x06, 0x9d, 0x7b, 0xe6, 0xdc,
	0x99, 0xcb, 0xf1, 0xb9, 0x00, 0x8c, 0xd3, 0x60, 0x10, 0xc5, 0x21, 0x0f, 0x51, 0x59, 0xe0, 0xce,
	0xf9, 0x2a, 0x0c, 0x57, 0x6b, 0x7a, 0x29, 0xb9, 0xc5, 0xf6, 0xfb, 0x25, 0x67, 0x01, 0x4d, 0xb8,
	0x17, 0x44, 0xa9, 0xac, 0xfb, 0x43, 0x01, 0xdd, 0xf2, 0xfd, 0x70, 0xe3, 0x70, 0x1a, 0xa0, 0x16,
	0xa8, 0xcc, 0xc7, 0x8a, 0xa9, 0xf4, 0x74, 0xa2, 0x32, 0x1f, 0x21, 0x28, 0x6f, 0xbc, 0x80, 0x62,
	0x55, 0x32, 0x12, 0xa3, 0x33, 0xa8, 0x44, 0x31, 0x5b, 0x52, 0x5c, 0x32, 0x95, 0x9e, 0x4a, 0xd2,
	0x02, 0xfd, 0x07, 0xba, 0xf7, 0xe8, 0xb1, 0xb5, 0xb7, 0x58, 0x53, 0x5c, 0x36, 0x95, 0x5e, 0x8d,
	0x1c, 0x08, 0x64, 0x42, 0xdd, 0xa7, 0xc9, 0x32, 0x66, 0x11, 0x67, 0xe1, 0x06, 0x57, 0xe4, 0x75,
	0x45, 0xaa, 0xfb, 0x4b, 0x81, 0x8a, 0x9c, 0xe3, 0x55, 0x33, 0x1c, 0xdd, 0x57, 0x7a, 0x76, 0x1f,
	0xfa, 0x17, 0xaa, 0x31, 0x4d, 0xb8, 0xcb, 0x7c, 0x39, 0x8d, 0x4e, 0x34, 0x51, 0x3a, 0x3e, 0x1a,
	0x00, 0x78, 0xe2, 0x1d, 0x57, 0xf8, 0x83, 0x2b, 0x66, 0xa9, 0x57, 0x1f, 0xb6, 0x07, 0xd2, 0xb8,
	0xbd, 0x0f, 0x44, 0xf7, 0xf6, 0x96, 0x98, 0xa0, 0xb9, 0x2e, 0xdf, 0x45, 0x14, 0x6b, 0xa6, 0xd2,
	0x6b, 0x0d, 0x21, 0xd5, 0xce, 0x77, 0x11, 0x25, 0x65, 0xf1, 0xdb, 0xfd, 0x59, 0x81, 0xf2, 0xab,
	0xdd, 0x7b, 0x79, 0xf2, 0xbd, 0xbf, 0xe5, 0xa2, 0xbf, 0x18, 0xaa, 0xb1, 0xc7, 0xd9, 0x66, 0x95,
	0x48, 0xf7, 0x54, 0x92, 0x97, 0xe8, 0x02, 0x9a, 0x19, 0x74, 0x97, 0xe1, 0x76, 0xc3, 0xe5, 0x9c,
	0x15, 0xd2, 0xc8, 0x48, 0x5b, 0x70, 0xe8, 0x1c, 0xea, 0x61, 0xec, 0xd3, 0x38, 0x93, 0x54, 0xa5,
	0x04, 0x24, 0xb5, 0x17, 0x2c, 0x3d, 0x4e, 0x57, 0x61, 0xbc, 0x13, 0x9e, 0xd5, 0xe4, 0x5c, 0x90,
	0x53, 0x8e, 0x8f, 0xde, 0x40, 0x3b, 0xd9, 0x2e, 0xdc, 0xa2, 0x48, 0x97, 0xa2, 0x66, 0xb2, 0x5d,
	0xd8, 0x07, 0xdd, 0x67, 0x68, 0x25, 0xd1, 0x3a, 0xd7, 0x31, 0x9a, 0x60, 0x30, 0x4b, 0xbd, 0xd6,
	0xf0, 0xef, 0xd4, 0xb7, 0x59, 0x44, 0x97, 0xcc, 0x5b, 0xe7, 0x0d, 0xa4, 0x99, 0x44, 0x79, 0xc1,
	0x68, 0x22, 0x2c, 0xe3, 0xde, 0x2a, 0xc1, 0x75, 0xb3, 0x24, 0x2c, 0x13, 0x18, 0xfd, 0x03, 0x1a,
	0x0b, 0xbc, 0x15, 0x4d, 0x70, 0x43, 0xb2, 0x59, 0x85, 0xde, 0x42, 0x95, 0xb3, 0x40, 0x5a, 0xd2,
	0x94, 0x4f, 0x34, 0xb3, 0x4f, 0x93, 0x92, 0x24, 0x3f, 0x45, 0x1d, 0xa8, 0x45, 0x31, 0x0b, 0x63,
	0xc6, 0x77, 0xb8, 0x25, 0xff, 0xf9, 0xbe, 0x7e, 0x9a, 0xdb, 0xf6, 0x71, 0x6e, 0x0d, 0x28, 0x3d,
	0xd2, 0x15, 0x36, 0x24, 0x2f, 0x20, 0xba, 0x00, 0x4d, 0x66, 0x23, 0xc1, 0xa7, 0x32, 0x3a, 0xf5,
	0x42, 0x74, 0x48, 0x76, 0x54, 0x0c, 0x1f, 0x7a, 0x12, 0xbe, 0x8f, 0x00, 0xcb, 0x98, 0x7a, 0x9c,
	0xfa, 0xae, 0xc7, 0xf1, 0x5f, 0xa6, 0xd2, 0xab, 0x0f, 0x3b, 0x83, 0x74, 0x47, 0x07, 0xf9, 0x8e,
	0x8a, 0x3f, 0x90, 0xee, 0x28, 0xd1, 0x33, 0xb5, 0xc5, 0x45, 0xeb, 0x36, 0xf2, 0xf3, 0xd6, 0xb3,
	0x97, 0x5b, 0x33, 0xb5, 0xc5, 0xbb, 0x67, 0xa0, 0x3a, 0xfe, 0x71, 0x3a, 0xbb, 0x18, 0x34, 0x92,
	0x4e, 0x75, 0x74, 0xd2, 0x1f, 0x82, 0x0c, 0x36, 0x6a, 0x40, 0xcd, 0xbe, 0x1d, 0xdb, 0x77, 0x57,
	0xd3, 0xaf, 0xc6, 0x09, 0x32, 0xa0, 0x41, 0xac, 0x91, 0x33, 0x75, 0xaf, 0x1e, 0xe6, 0xf3, 0xe9,
	0xc4, 0x50, 0x10, 0x80, 0x36, 0xbb, 0x77, 0x46, 0x63, 0x62, 0xa8, 0xfd, 0x4f, 0x50, 0xcd, 0x7c,
	0x47, 0x4d, 0xd0, 0xaf, 0xc8, 0xd8, 0xba, 0xbb, 0xb6, 0x66, 0x73, 0xe3, 0x04, 0xe9, 0x50, 0xb9,
	0x7f, 0x98, 0xd8, 0xb7, 0x59, 0xc3, 0xc4, 0xb2, 0xef, 0x66, 0x86, 0x2a, 0xf0, 0xc8, 0x99, 0x4c,
	0xc6, 0xc4, 0x28, 0xf5, 0x3f, 0x40, 0xfb, 0x28, 0x17, 0xe8, 0x14, 0x9a, 0xf6, 0xed, 0xf8, 0x7a,
	0xe6, 0xce, 0xbe, 0x8c, 0x6d, 0xc7, 0xba, 0x37, 0x4e, 0x10, 0x82, 0xd6, 0x7c, 0x3a, 0xb2, 0xbe,
	0x1d, 0x38, 0x65, 0xf8, 0x5b, 0x81, 0xba, 0xd8, 0xbd, 0x19, 0x8d, 0x1f, 0xc5, 0x9a, 0x9c, 0x43,
	0xd5, 0xf2, 0x7d, 0xb9, 0x8d, 0xd9, 0xa2, 0x0a, 0xdc, 0xa9, 0x65, 0xd8, 0x17, 0x82, 0x1b, 0xca,
	0xa5, 0x60, 0x4f, 0x76, 0x0a, 0x52, 0xd4, 0x87, 0xc6, 0x0d, 0xe5, 0xd2, 0x19, 0x4e, 0x83, 0x04,
	0x35, 0xd2, 0xb3, 0xd4, 0xaa, 0xa2, 0xf2, 0x9d, 0x82, 0xfe, 0x07, 0x78, 0x90, 0x2e, 0x3f, 0x7b,
	0xb0, 0x78, 0xa3, 0x09, 0x30, 0xa2, 0x6b, 0x9a, 0xa9, 0x0e, 0xaf, 0xee, 0xd1, 0x42, 0x93, 0xdf,
	0xef, 0xfd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xa6, 0x0b, 0x72, 0xc0, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemServiceClient interface {
	AddItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Id, error)
	GetItem(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Item, error)
	GetRestItems(ctx context.Context, in *RestId, opts ...grpc.CallOption) (ItemService_GetRestItemsClient, error)
	UpdateItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	DeleteItem(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error)
}

type itemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemServiceClient(cc grpc.ClientConnInterface) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) AddItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/item.ItemService/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItem(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetRestItems(ctx context.Context, in *RestId, opts ...grpc.CallOption) (ItemService_GetRestItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ItemService_serviceDesc.Streams[0], "/item.ItemService/GetRestItems", opts...)
	if err != nil {
		return nil, err
	}
	x := &itemServiceGetRestItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemService_GetRestItemsClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type itemServiceGetRestItemsClient struct {
	grpc.ClientStream
}

func (x *itemServiceGetRestItemsClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemServiceClient) UpdateItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/item.ItemService/UpdateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) DeleteItem(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/item.ItemService/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServiceServer is the server API for ItemService service.
type ItemServiceServer interface {
	AddItem(context.Context, *Item) (*Id, error)
	GetItem(context.Context, *Id) (*Item, error)
	GetRestItems(*RestId, ItemService_GetRestItemsServer) error
	UpdateItem(context.Context, *Item) (*Item, error)
	DeleteItem(context.Context, *Id) (*Id, error)
}

// UnimplementedItemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedItemServiceServer struct {
}

func (*UnimplementedItemServiceServer) AddItem(ctx context.Context, req *Item) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (*UnimplementedItemServiceServer) GetItem(ctx context.Context, req *Id) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (*UnimplementedItemServiceServer) GetRestItems(req *RestId, srv ItemService_GetRestItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRestItems not implemented")
}
func (*UnimplementedItemServiceServer) UpdateItem(ctx context.Context, req *Item) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (*UnimplementedItemServiceServer) DeleteItem(ctx context.Context, req *Id) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}

func RegisterItemServiceServer(s *grpc.Server, srv ItemServiceServer) {
	s.RegisterService(&_ItemService_serviceDesc, srv)
}

func _ItemService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).AddItem(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItem(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetRestItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RestId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemServiceServer).GetRestItems(m, &itemServiceGetRestItemsServer{stream})
}

type ItemService_GetRestItemsServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type itemServiceGetRestItemsServer struct {
	grpc.ServerStream
}

func (x *itemServiceGetRestItemsServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

func _ItemService_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).UpdateItem(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).DeleteItem(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _ItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "item.ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _ItemService_AddItem_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _ItemService_GetItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _ItemService_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _ItemService_DeleteItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRestItems",
			Handler:       _ItemService_GetRestItems_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "item.proto",
}
