// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: api/calendar/calendar.proto

package calendar

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type AddEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime   *timestamp.Timestamp `protobuf:"bytes,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime     *timestamp.Timestamp `protobuf:"bytes,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddEventRequest) Reset() {
	*x = AddEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_calendar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventRequest) ProtoMessage() {}

func (x *AddEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_calendar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventRequest.ProtoReflect.Descriptor instead.
func (*AddEventRequest) Descriptor() ([]byte, []int) {
	return file_api_calendar_calendar_proto_rawDescGZIP(), []int{0}
}

func (x *AddEventRequest) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *AddEventRequest) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *AddEventRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AddEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddEventResponse) Reset() {
	*x = AddEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_calendar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventResponse) ProtoMessage() {}

func (x *AddEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_calendar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventResponse.ProtoReflect.Descriptor instead.
func (*AddEventResponse) Descriptor() ([]byte, []int) {
	return file_api_calendar_calendar_proto_rawDescGZIP(), []int{1}
}

func (x *AddEventResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetEventByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEventByIDRequest) Reset() {
	*x = GetEventByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_calendar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventByIDRequest) ProtoMessage() {}

func (x *GetEventByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_calendar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventByIDRequest.ProtoReflect.Descriptor instead.
func (*GetEventByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_calendar_calendar_proto_rawDescGZIP(), []int{2}
}

func (x *GetEventByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetEventByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime   *timestamp.Timestamp `protobuf:"bytes,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime     *timestamp.Timestamp `protobuf:"bytes,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetEventByIDResponse) Reset() {
	*x = GetEventByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_calendar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventByIDResponse) ProtoMessage() {}

func (x *GetEventByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_calendar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventByIDResponse.ProtoReflect.Descriptor instead.
func (*GetEventByIDResponse) Descriptor() ([]byte, []int) {
	return file_api_calendar_calendar_proto_rawDescGZIP(), []int{3}
}

func (x *GetEventByIDResponse) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *GetEventByIDResponse) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *GetEventByIDResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_api_calendar_calendar_proto protoreflect.FileDescriptor

var file_api_calendar_calendar_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x41, 0x72, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x0f, 0x41, 0x64, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0x9e, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x12, 0x81, 0x01, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x39,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x41, 0x72, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x41, 0x72, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x2e, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x41, 0x72, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x2e, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x41, 0x72, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_calendar_calendar_proto_rawDescOnce sync.Once
	file_api_calendar_calendar_proto_rawDescData = file_api_calendar_calendar_proto_rawDesc
)

func file_api_calendar_calendar_proto_rawDescGZIP() []byte {
	file_api_calendar_calendar_proto_rawDescOnce.Do(func() {
		file_api_calendar_calendar_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_calendar_calendar_proto_rawDescData)
	})
	return file_api_calendar_calendar_proto_rawDescData
}

var file_api_calendar_calendar_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_calendar_calendar_proto_goTypes = []interface{}{
	(*AddEventRequest)(nil),      // 0: github.com.Artenso.calendar.api.calendar.AddEventRequest
	(*AddEventResponse)(nil),     // 1: github.com.Artenso.calendar.api.calendar.AddEventResponse
	(*GetEventByIDRequest)(nil),  // 2: github.com.Artenso.calendar.api.calendar.GetEventByIDRequest
	(*GetEventByIDResponse)(nil), // 3: github.com.Artenso.calendar.api.calendar.GetEventByIDResponse
	(*timestamp.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_api_calendar_calendar_proto_depIdxs = []int32{
	4, // 0: github.com.Artenso.calendar.api.calendar.AddEventRequest.startTime:type_name -> google.protobuf.Timestamp
	4, // 1: github.com.Artenso.calendar.api.calendar.AddEventRequest.endTime:type_name -> google.protobuf.Timestamp
	4, // 2: github.com.Artenso.calendar.api.calendar.GetEventByIDResponse.startTime:type_name -> google.protobuf.Timestamp
	4, // 3: github.com.Artenso.calendar.api.calendar.GetEventByIDResponse.endTime:type_name -> google.protobuf.Timestamp
	0, // 4: github.com.Artenso.calendar.api.calendar.Calendar.AddEvent:input_type -> github.com.Artenso.calendar.api.calendar.AddEventRequest
	2, // 5: github.com.Artenso.calendar.api.calendar.Calendar.GetEventByID:input_type -> github.com.Artenso.calendar.api.calendar.GetEventByIDRequest
	1, // 6: github.com.Artenso.calendar.api.calendar.Calendar.AddEvent:output_type -> github.com.Artenso.calendar.api.calendar.AddEventResponse
	3, // 7: github.com.Artenso.calendar.api.calendar.Calendar.GetEventByID:output_type -> github.com.Artenso.calendar.api.calendar.GetEventByIDResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_calendar_calendar_proto_init() }
func file_api_calendar_calendar_proto_init() {
	if File_api_calendar_calendar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_calendar_calendar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventRequest); i {
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
		file_api_calendar_calendar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventResponse); i {
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
		file_api_calendar_calendar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventByIDRequest); i {
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
		file_api_calendar_calendar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventByIDResponse); i {
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
			RawDescriptor: file_api_calendar_calendar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_calendar_calendar_proto_goTypes,
		DependencyIndexes: file_api_calendar_calendar_proto_depIdxs,
		MessageInfos:      file_api_calendar_calendar_proto_msgTypes,
	}.Build()
	File_api_calendar_calendar_proto = out.File
	file_api_calendar_calendar_proto_rawDesc = nil
	file_api_calendar_calendar_proto_goTypes = nil
	file_api_calendar_calendar_proto_depIdxs = nil
}
