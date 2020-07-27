// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: orders/orderspb/orders.proto

package orderspb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OrdersCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrdersCountRequest) Reset() {
	*x = OrdersCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orderspb_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdersCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersCountRequest) ProtoMessage() {}

func (x *OrdersCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orderspb_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersCountRequest.ProtoReflect.Descriptor instead.
func (*OrdersCountRequest) Descriptor() ([]byte, []int) {
	return file_orders_orderspb_orders_proto_rawDescGZIP(), []int{0}
}

type OrdersCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count string `protobuf:"bytes,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *OrdersCountResponse) Reset() {
	*x = OrdersCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orderspb_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdersCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersCountResponse) ProtoMessage() {}

func (x *OrdersCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orderspb_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersCountResponse.ProtoReflect.Descriptor instead.
func (*OrdersCountResponse) Descriptor() ([]byte, []int) {
	return file_orders_orderspb_orders_proto_rawDescGZIP(), []int{1}
}

func (x *OrdersCountResponse) GetCount() string {
	if x != nil {
		return x.Count
	}
	return ""
}

type OrderDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderNumber string `protobuf:"bytes,1,opt,name=OrderNumber,proto3" json:"OrderNumber,omitempty"`
}

func (x *OrderDetailRequest) Reset() {
	*x = OrderDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orderspb_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetailRequest) ProtoMessage() {}

func (x *OrderDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orderspb_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetailRequest.ProtoReflect.Descriptor instead.
func (*OrderDetailRequest) Descriptor() ([]byte, []int) {
	return file_orders_orderspb_orders_proto_rawDescGZIP(), []int{2}
}

func (x *OrderDetailRequest) GetOrderNumber() string {
	if x != nil {
		return x.OrderNumber
	}
	return ""
}

type OrderDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderDetail string `protobuf:"bytes,1,opt,name=orderDetail,proto3" json:"orderDetail,omitempty"`
}

func (x *OrderDetailResponse) Reset() {
	*x = OrderDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orderspb_orders_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetailResponse) ProtoMessage() {}

func (x *OrderDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orderspb_orders_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetailResponse.ProtoReflect.Descriptor instead.
func (*OrderDetailResponse) Descriptor() ([]byte, []int) {
	return file_orders_orderspb_orders_proto_rawDescGZIP(), []int{3}
}

func (x *OrderDetailResponse) GetOrderDetail() string {
	if x != nil {
		return x.OrderDetail
	}
	return ""
}

type PopularDishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DishName string `protobuf:"bytes,1,opt,name=DishName,proto3" json:"DishName,omitempty"`
}

func (x *PopularDishResponse) Reset() {
	*x = PopularDishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orderspb_orders_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopularDishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopularDishResponse) ProtoMessage() {}

func (x *PopularDishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orderspb_orders_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopularDishResponse.ProtoReflect.Descriptor instead.
func (*PopularDishResponse) Descriptor() ([]byte, []int) {
	return file_orders_orderspb_orders_proto_rawDescGZIP(), []int{4}
}

func (x *PopularDishResponse) GetDishName() string {
	if x != nil {
		return x.DishName
	}
	return ""
}

type PopularDishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityName string `protobuf:"bytes,1,opt,name=CityName,proto3" json:"CityName,omitempty"`
}

func (x *PopularDishRequest) Reset() {
	*x = PopularDishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orderspb_orders_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopularDishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopularDishRequest) ProtoMessage() {}

func (x *PopularDishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orderspb_orders_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopularDishRequest.ProtoReflect.Descriptor instead.
func (*PopularDishRequest) Descriptor() ([]byte, []int) {
	return file_orders_orderspb_orders_proto_rawDescGZIP(), []int{5}
}

func (x *PopularDishRequest) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

type UpdateDishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     int64  `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	UpdatedDish string `protobuf:"bytes,2,opt,name=updatedDish,proto3" json:"updatedDish,omitempty"`
}

func (x *UpdateDishRequest) Reset() {
	*x = UpdateDishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orderspb_orders_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDishRequest) ProtoMessage() {}

func (x *UpdateDishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orderspb_orders_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDishRequest.ProtoReflect.Descriptor instead.
func (*UpdateDishRequest) Descriptor() ([]byte, []int) {
	return file_orders_orderspb_orders_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDishRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *UpdateDishRequest) GetUpdatedDish() string {
	if x != nil {
		return x.UpdatedDish
	}
	return ""
}

type UpdateDishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateDishResponse) Reset() {
	*x = UpdateDishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orderspb_orders_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDishResponse) ProtoMessage() {}

func (x *UpdateDishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orderspb_orders_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDishResponse.ProtoReflect.Descriptor instead.
func (*UpdateDishResponse) Descriptor() ([]byte, []int) {
	return file_orders_orderspb_orders_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateDishResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_orders_orderspb_orders_proto protoreflect.FileDescriptor

var file_orders_orderspb_orders_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x70,
	0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x13,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x36, 0x0a, 0x12, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x37, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x31, 0x0a, 0x13, 0x50, 0x6f,
	0x70, 0x75, 0x6c, 0x61, 0x72, 0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a,
	0x12, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x4f, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x44, 0x69, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x44, 0x69, 0x73, 0x68,
	0x22, 0x2c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xbd,
	0x02, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x44, 0x69, 0x73, 0x68, 0x12, 0x1a, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x44, 0x69, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x2e, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x73, 0x68, 0x12, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11,
	0x5a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_orderspb_orders_proto_rawDescOnce sync.Once
	file_orders_orderspb_orders_proto_rawDescData = file_orders_orderspb_orders_proto_rawDesc
)

func file_orders_orderspb_orders_proto_rawDescGZIP() []byte {
	file_orders_orderspb_orders_proto_rawDescOnce.Do(func() {
		file_orders_orderspb_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_orderspb_orders_proto_rawDescData)
	})
	return file_orders_orderspb_orders_proto_rawDescData
}

var file_orders_orderspb_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_orders_orderspb_orders_proto_goTypes = []interface{}{
	(*OrdersCountRequest)(nil),  // 0: orders.OrdersCountRequest
	(*OrdersCountResponse)(nil), // 1: orders.OrdersCountResponse
	(*OrderDetailRequest)(nil),  // 2: orders.OrderDetailRequest
	(*OrderDetailResponse)(nil), // 3: orders.OrderDetailResponse
	(*PopularDishResponse)(nil), // 4: orders.PopularDishResponse
	(*PopularDishRequest)(nil),  // 5: orders.PopularDishRequest
	(*UpdateDishRequest)(nil),   // 6: orders.UpdateDishRequest
	(*UpdateDishResponse)(nil),  // 7: orders.UpdateDishResponse
}
var file_orders_orderspb_orders_proto_depIdxs = []int32{
	0, // 0: orders.OrdersService.GetOrdersCount:input_type -> orders.OrdersCountRequest
	2, // 1: orders.OrdersService.GetOrderDetail:input_type -> orders.OrderDetailRequest
	5, // 2: orders.OrdersService.GetPopularDish:input_type -> orders.PopularDishRequest
	6, // 3: orders.OrdersService.UpdateDish:input_type -> orders.UpdateDishRequest
	1, // 4: orders.OrdersService.GetOrdersCount:output_type -> orders.OrdersCountResponse
	3, // 5: orders.OrdersService.GetOrderDetail:output_type -> orders.OrderDetailResponse
	4, // 6: orders.OrdersService.GetPopularDish:output_type -> orders.PopularDishResponse
	7, // 7: orders.OrdersService.UpdateDish:output_type -> orders.UpdateDishResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orders_orderspb_orders_proto_init() }
func file_orders_orderspb_orders_proto_init() {
	if File_orders_orderspb_orders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_orderspb_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdersCountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orders_orderspb_orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdersCountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orders_orderspb_orders_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orders_orderspb_orders_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orders_orderspb_orders_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopularDishResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orders_orderspb_orders_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopularDishRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orders_orderspb_orders_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDishRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orders_orderspb_orders_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDishResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orders_orderspb_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_orderspb_orders_proto_goTypes,
		DependencyIndexes: file_orders_orderspb_orders_proto_depIdxs,
		MessageInfos:      file_orders_orderspb_orders_proto_msgTypes,
	}.Build()
	File_orders_orderspb_orders_proto = out.File
	file_orders_orderspb_orders_proto_rawDesc = nil
	file_orders_orderspb_orders_proto_goTypes = nil
	file_orders_orderspb_orders_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrdersServiceClient is the client API for OrdersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrdersServiceClient interface {
	GetOrdersCount(ctx context.Context, in *OrdersCountRequest, opts ...grpc.CallOption) (*OrdersCountResponse, error)
	GetOrderDetail(ctx context.Context, in *OrderDetailRequest, opts ...grpc.CallOption) (*OrderDetailResponse, error)
	GetPopularDish(ctx context.Context, in *PopularDishRequest, opts ...grpc.CallOption) (*PopularDishResponse, error)
	UpdateDish(ctx context.Context, in *UpdateDishRequest, opts ...grpc.CallOption) (*UpdateDishResponse, error)
}

type ordersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersServiceClient(cc grpc.ClientConnInterface) OrdersServiceClient {
	return &ordersServiceClient{cc}
}

func (c *ordersServiceClient) GetOrdersCount(ctx context.Context, in *OrdersCountRequest, opts ...grpc.CallOption) (*OrdersCountResponse, error) {
	out := new(OrdersCountResponse)
	err := c.cc.Invoke(ctx, "/orders.OrdersService/GetOrdersCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) GetOrderDetail(ctx context.Context, in *OrderDetailRequest, opts ...grpc.CallOption) (*OrderDetailResponse, error) {
	out := new(OrderDetailResponse)
	err := c.cc.Invoke(ctx, "/orders.OrdersService/GetOrderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) GetPopularDish(ctx context.Context, in *PopularDishRequest, opts ...grpc.CallOption) (*PopularDishResponse, error) {
	out := new(PopularDishResponse)
	err := c.cc.Invoke(ctx, "/orders.OrdersService/GetPopularDish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) UpdateDish(ctx context.Context, in *UpdateDishRequest, opts ...grpc.CallOption) (*UpdateDishResponse, error) {
	out := new(UpdateDishResponse)
	err := c.cc.Invoke(ctx, "/orders.OrdersService/UpdateDish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServiceServer is the server API for OrdersService service.
type OrdersServiceServer interface {
	GetOrdersCount(context.Context, *OrdersCountRequest) (*OrdersCountResponse, error)
	GetOrderDetail(context.Context, *OrderDetailRequest) (*OrderDetailResponse, error)
	GetPopularDish(context.Context, *PopularDishRequest) (*PopularDishResponse, error)
	UpdateDish(context.Context, *UpdateDishRequest) (*UpdateDishResponse, error)
}

// UnimplementedOrdersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrdersServiceServer struct {
}

func (*UnimplementedOrdersServiceServer) GetOrdersCount(context.Context, *OrdersCountRequest) (*OrdersCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersCount not implemented")
}
func (*UnimplementedOrdersServiceServer) GetOrderDetail(context.Context, *OrderDetailRequest) (*OrderDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderDetail not implemented")
}
func (*UnimplementedOrdersServiceServer) GetPopularDish(context.Context, *PopularDishRequest) (*PopularDishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPopularDish not implemented")
}
func (*UnimplementedOrdersServiceServer) UpdateDish(context.Context, *UpdateDishRequest) (*UpdateDishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDish not implemented")
}

func RegisterOrdersServiceServer(s *grpc.Server, srv OrdersServiceServer) {
	s.RegisterService(&_OrdersService_serviceDesc, srv)
}

func _OrdersService_GetOrdersCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).GetOrdersCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrdersService/GetOrdersCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).GetOrdersCount(ctx, req.(*OrdersCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_GetOrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).GetOrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrdersService/GetOrderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).GetOrderDetail(ctx, req.(*OrderDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_GetPopularDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopularDishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).GetPopularDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrdersService/GetPopularDish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).GetPopularDish(ctx, req.(*PopularDishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_UpdateDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).UpdateDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrdersService/UpdateDish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).UpdateDish(ctx, req.(*UpdateDishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrdersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orders.OrdersService",
	HandlerType: (*OrdersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrdersCount",
			Handler:    _OrdersService_GetOrdersCount_Handler,
		},
		{
			MethodName: "GetOrderDetail",
			Handler:    _OrdersService_GetOrderDetail_Handler,
		},
		{
			MethodName: "GetPopularDish",
			Handler:    _OrdersService_GetPopularDish_Handler,
		},
		{
			MethodName: "UpdateDish",
			Handler:    _OrdersService_UpdateDish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/orderspb/orders.proto",
}
