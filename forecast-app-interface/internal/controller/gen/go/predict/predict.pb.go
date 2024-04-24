// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: predict.proto

package predict

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_predict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_predict_proto_rawDescGZIP(), []int{0}
}

type TimeSeriesItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts    int64   `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TimeSeriesItem) Reset() {
	*x = TimeSeriesItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeriesItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesItem) ProtoMessage() {}

func (x *TimeSeriesItem) ProtoReflect() protoreflect.Message {
	mi := &file_predict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesItem.ProtoReflect.Descriptor instead.
func (*TimeSeriesItem) Descriptor() ([]byte, []int) {
	return file_predict_proto_rawDescGZIP(), []int{1}
}

func (x *TimeSeriesItem) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *TimeSeriesItem) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type MakePredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string            `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name     string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Unit     string            `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	Period   int32             `protobuf:"varint,4,opt,name=period,proto3" json:"period,omitempty"`
	Items    []*TimeSeriesItem `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *MakePredictRequest) Reset() {
	*x = MakePredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predict_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakePredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakePredictRequest) ProtoMessage() {}

func (x *MakePredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_predict_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakePredictRequest.ProtoReflect.Descriptor instead.
func (*MakePredictRequest) Descriptor() ([]byte, []int) {
	return file_predict_proto_rawDescGZIP(), []int{2}
}

func (x *MakePredictRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MakePredictRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MakePredictRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *MakePredictRequest) GetPeriod() int32 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *MakePredictRequest) GetItems() []*TimeSeriesItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetPredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPredictRequest) Reset() {
	*x = GetPredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predict_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPredictRequest) ProtoMessage() {}

func (x *GetPredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_predict_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPredictRequest.ProtoReflect.Descriptor instead.
func (*GetPredictRequest) Descriptor() ([]byte, []int) {
	return file_predict_proto_rawDescGZIP(), []int{3}
}

func (x *GetPredictRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetPredictRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit      string            `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	Delimiter int64             `protobuf:"varint,2,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	Items     []*TimeSeriesItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetPredictResponse) Reset() {
	*x = GetPredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predict_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPredictResponse) ProtoMessage() {}

func (x *GetPredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_predict_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPredictResponse.ProtoReflect.Descriptor instead.
func (*GetPredictResponse) Descriptor() ([]byte, []int) {
	return file_predict_proto_rawDescGZIP(), []int{4}
}

func (x *GetPredictResponse) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *GetPredictResponse) GetDelimiter() int64 {
	if x != nil {
		return x.Delimiter
	}
	return 0
}

func (x *GetPredictResponse) GetItems() []*TimeSeriesItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_predict_proto protoreflect.FileDescriptor

var file_predict_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x36, 0x0a, 0x0e, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x97, 0x01, 0x0a, 0x12, 0x4d, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x43, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x6d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x73,
	0x0a, 0x0e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2a, 0x0a, 0x0b, 0x4d, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12,
	0x13, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_predict_proto_rawDescOnce sync.Once
	file_predict_proto_rawDescData = file_predict_proto_rawDesc
)

func file_predict_proto_rawDescGZIP() []byte {
	file_predict_proto_rawDescOnce.Do(func() {
		file_predict_proto_rawDescData = protoimpl.X.CompressGZIP(file_predict_proto_rawDescData)
	})
	return file_predict_proto_rawDescData
}

var file_predict_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_predict_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: Empty
	(*TimeSeriesItem)(nil),     // 1: TimeSeriesItem
	(*MakePredictRequest)(nil), // 2: MakePredictRequest
	(*GetPredictRequest)(nil),  // 3: GetPredictRequest
	(*GetPredictResponse)(nil), // 4: GetPredictResponse
}
var file_predict_proto_depIdxs = []int32{
	1, // 0: MakePredictRequest.items:type_name -> TimeSeriesItem
	1, // 1: GetPredictResponse.items:type_name -> TimeSeriesItem
	2, // 2: PredictService.MakePredict:input_type -> MakePredictRequest
	3, // 3: PredictService.GetPredict:input_type -> GetPredictRequest
	0, // 4: PredictService.MakePredict:output_type -> Empty
	4, // 5: PredictService.GetPredict:output_type -> GetPredictResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_predict_proto_init() }
func file_predict_proto_init() {
	if File_predict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_predict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_predict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeriesItem); i {
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
		file_predict_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakePredictRequest); i {
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
		file_predict_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPredictRequest); i {
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
		file_predict_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPredictResponse); i {
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
			RawDescriptor: file_predict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_predict_proto_goTypes,
		DependencyIndexes: file_predict_proto_depIdxs,
		MessageInfos:      file_predict_proto_msgTypes,
	}.Build()
	File_predict_proto = out.File
	file_predict_proto_rawDesc = nil
	file_predict_proto_goTypes = nil
	file_predict_proto_depIdxs = nil
}
