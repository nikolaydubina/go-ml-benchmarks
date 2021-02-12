// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: proto/predictor.proto

package proto

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

type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Survived    int32   `protobuf:"varint,1,opt,name=Survived,proto3" json:"Survived,omitempty"`
	PassengerId int32   `protobuf:"varint,2,opt,name=PassengerId,proto3" json:"PassengerId,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Pclass      float64 `protobuf:"fixed64,4,opt,name=Pclass,proto3" json:"Pclass,omitempty"`
	Sex         string  `protobuf:"bytes,5,opt,name=Sex,proto3" json:"Sex,omitempty"`
	Age         float64 `protobuf:"fixed64,6,opt,name=Age,proto3" json:"Age,omitempty"`
	SibSp       float64 `protobuf:"fixed64,7,opt,name=SibSp,proto3" json:"SibSp,omitempty"`
	Parch       float64 `protobuf:"fixed64,8,opt,name=Parch,proto3" json:"Parch,omitempty"`
	Ticket      string  `protobuf:"bytes,9,opt,name=Ticket,proto3" json:"Ticket,omitempty"`
	Fare        float64 `protobuf:"fixed64,10,opt,name=Fare,proto3" json:"Fare,omitempty"`
	Cabin       string  `protobuf:"bytes,11,opt,name=Cabin,proto3" json:"Cabin,omitempty"`
	Embarked    string  `protobuf:"bytes,12,opt,name=Embarked,proto3" json:"Embarked,omitempty"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_predictor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_predictor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_proto_predictor_proto_rawDescGZIP(), []int{0}
}

func (x *PredictRequest) GetSurvived() int32 {
	if x != nil {
		return x.Survived
	}
	return 0
}

func (x *PredictRequest) GetPassengerId() int32 {
	if x != nil {
		return x.PassengerId
	}
	return 0
}

func (x *PredictRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PredictRequest) GetPclass() float64 {
	if x != nil {
		return x.Pclass
	}
	return 0
}

func (x *PredictRequest) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *PredictRequest) GetAge() float64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *PredictRequest) GetSibSp() float64 {
	if x != nil {
		return x.SibSp
	}
	return 0
}

func (x *PredictRequest) GetParch() float64 {
	if x != nil {
		return x.Parch
	}
	return 0
}

func (x *PredictRequest) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

func (x *PredictRequest) GetFare() float64 {
	if x != nil {
		return x.Fare
	}
	return 0
}

func (x *PredictRequest) GetCabin() string {
	if x != nil {
		return x.Cabin
	}
	return ""
}

func (x *PredictRequest) GetEmbarked() string {
	if x != nil {
		return x.Embarked
	}
	return ""
}

type PredictProcessedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features []float64 `protobuf:"fixed64,1,rep,packed,name=Features,proto3" json:"Features,omitempty"`
}

func (x *PredictProcessedRequest) Reset() {
	*x = PredictProcessedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_predictor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictProcessedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictProcessedRequest) ProtoMessage() {}

func (x *PredictProcessedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_predictor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictProcessedRequest.ProtoReflect.Descriptor instead.
func (*PredictProcessedRequest) Descriptor() ([]byte, []int) {
	return file_proto_predictor_proto_rawDescGZIP(), []int{1}
}

func (x *PredictProcessedRequest) GetFeatures() []float64 {
	if x != nil {
		return x.Features
	}
	return nil
}

type PredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prediction float64 `protobuf:"fixed64,1,opt,name=Prediction,proto3" json:"Prediction,omitempty"`
}

func (x *PredictResponse) Reset() {
	*x = PredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_predictor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse) ProtoMessage() {}

func (x *PredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_predictor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse.ProtoReflect.Descriptor instead.
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return file_proto_predictor_proto_rawDescGZIP(), []int{2}
}

func (x *PredictResponse) GetPrediction() float64 {
	if x != nil {
		return x.Prediction
	}
	return 0
}

var File_proto_predictor_proto protoreflect.FileDescriptor

var file_proto_predictor_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x6f, 0x72, 0x22, 0xa8, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76, 0x65,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x50, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x65,
	0x78, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x41, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x69, 0x62, 0x53, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x53, 0x69, 0x62, 0x53, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x61, 0x72, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x46, 0x61, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43,
	0x61, 0x62, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x61, 0x62, 0x69,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6d, 0x62, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x6d, 0x62, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x22, 0x35, 0x0a,
	0x17, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x08, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xa5, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x10, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x22, 0x2e,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69,
	0x6b, 0x6f, 0x6c, 0x61, 0x79, 0x64, 0x75, 0x62, 0x69, 0x6e, 0x61, 0x2f, 0x67, 0x6f, 0x2d, 0x6d,
	0x6c, 0x2d, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_predictor_proto_rawDescOnce sync.Once
	file_proto_predictor_proto_rawDescData = file_proto_predictor_proto_rawDesc
)

func file_proto_predictor_proto_rawDescGZIP() []byte {
	file_proto_predictor_proto_rawDescOnce.Do(func() {
		file_proto_predictor_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_predictor_proto_rawDescData)
	})
	return file_proto_predictor_proto_rawDescData
}

var file_proto_predictor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_predictor_proto_goTypes = []interface{}{
	(*PredictRequest)(nil),          // 0: predictor.PredictRequest
	(*PredictProcessedRequest)(nil), // 1: predictor.PredictProcessedRequest
	(*PredictResponse)(nil),         // 2: predictor.PredictResponse
}
var file_proto_predictor_proto_depIdxs = []int32{
	0, // 0: predictor.Predictor.Predict:input_type -> predictor.PredictRequest
	1, // 1: predictor.Predictor.PredictProcessed:input_type -> predictor.PredictProcessedRequest
	2, // 2: predictor.Predictor.Predict:output_type -> predictor.PredictResponse
	2, // 3: predictor.Predictor.PredictProcessed:output_type -> predictor.PredictResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_predictor_proto_init() }
func file_proto_predictor_proto_init() {
	if File_proto_predictor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_predictor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictRequest); i {
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
		file_proto_predictor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictProcessedRequest); i {
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
		file_proto_predictor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictResponse); i {
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
			RawDescriptor: file_proto_predictor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_predictor_proto_goTypes,
		DependencyIndexes: file_proto_predictor_proto_depIdxs,
		MessageInfos:      file_proto_predictor_proto_msgTypes,
	}.Build()
	File_proto_predictor_proto = out.File
	file_proto_predictor_proto_rawDesc = nil
	file_proto_predictor_proto_goTypes = nil
	file_proto_predictor_proto_depIdxs = nil
}