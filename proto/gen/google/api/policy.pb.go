// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: google/api/policy.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Google API Policy Annotation
//
// This message defines a simple API policy annotation that can be used to
// annotate API request and response message fields with applicable policies.
// One field may have multiple applicable policies that must all be satisfied
// before a request can be processed. This policy annotation is used to
// generate the overall policy that will be used for automatic runtime
// policy enforcement and documentation generation.
type FieldPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selects one or more request or response message fields to apply this
	// `FieldPolicy`.
	//
	// When a `FieldPolicy` is used in proto annotation, the selector must
	// be left as empty. The service config generator will automatically fill
	// the correct value.
	//
	// When a `FieldPolicy` is used in service config, the selector must be a
	// comma-separated string with valid request or response field paths,
	// such as "foo.bar" or "foo.bar,foo.baz".
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Specifies the required permission(s) for the resource referred to by the
	// field. It requires the field contains a valid resource reference, and
	// the request must pass the permission checks to proceed. For example,
	// "resourcemanager.projects.get".
	ResourcePermission string `protobuf:"bytes,2,opt,name=resource_permission,json=resourcePermission,proto3" json:"resource_permission,omitempty"`
	// Specifies the resource type for the resource referred to by the field.
	ResourceType string `protobuf:"bytes,3,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
}

func (x *FieldPolicy) Reset() {
	*x = FieldPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldPolicy) ProtoMessage() {}

func (x *FieldPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldPolicy.ProtoReflect.Descriptor instead.
func (*FieldPolicy) Descriptor() ([]byte, []int) {
	return file_google_api_policy_proto_rawDescGZIP(), []int{0}
}

func (x *FieldPolicy) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *FieldPolicy) GetResourcePermission() string {
	if x != nil {
		return x.ResourcePermission
	}
	return ""
}

func (x *FieldPolicy) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

// Defines policies applying to an RPC method.
type MethodPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selects a method to which these policies should be enforced, for example,
	// "google.pubsub.v1.Subscriber.CreateSubscription".
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax
	// details.
	//
	// NOTE: This field must not be set in the proto annotation. It will be
	// automatically filled by the service config compiler .
	Selector string `protobuf:"bytes,9,opt,name=selector,proto3" json:"selector,omitempty"`
	// Policies that are applicable to the request message.
	RequestPolicies []*FieldPolicy `protobuf:"bytes,2,rep,name=request_policies,json=requestPolicies,proto3" json:"request_policies,omitempty"`
}

func (x *MethodPolicy) Reset() {
	*x = MethodPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodPolicy) ProtoMessage() {}

func (x *MethodPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodPolicy.ProtoReflect.Descriptor instead.
func (*MethodPolicy) Descriptor() ([]byte, []int) {
	return file_google_api_policy_proto_rawDescGZIP(), []int{1}
}

func (x *MethodPolicy) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *MethodPolicy) GetRequestPolicies() []*FieldPolicy {
	if x != nil {
		return x.RequestPolicies
	}
	return nil
}

var file_google_api_policy_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldPolicy)(nil),
		Field:         158361448,
		Name:          "google.api.field_policy",
		Tag:           "bytes,158361448,opt,name=field_policy",
		Filename:      "google/api/policy.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodPolicy)(nil),
		Field:         161893301,
		Name:          "google.api.method_policy",
		Tag:           "bytes,161893301,opt,name=method_policy",
		Filename:      "google/api/policy.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// See [FieldPolicy][].
	//
	// optional google.api.FieldPolicy field_policy = 158361448;
	E_FieldPolicy = &file_google_api_policy_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// See [MethodPolicy][].
	//
	// optional google.api.MethodPolicy method_policy = 161893301;
	E_MethodPolicy = &file_google_api_policy_proto_extTypes[1]
)

var File_google_api_policy_proto protoreflect.FileDescriptor

var file_google_api_policy_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6e, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x3a, 0x5c, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe8, 0xce, 0xc1, 0x4b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x60, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb5, 0x97, 0x99, 0x4d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x9f, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0b, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x7a,
	0x7a, 0x30, 0x30, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x58, 0xaa, 0x02, 0x0a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x70, 0x69, 0xca, 0x02, 0x0a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x70, 0x69, 0xe2, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_api_policy_proto_rawDescOnce sync.Once
	file_google_api_policy_proto_rawDescData = file_google_api_policy_proto_rawDesc
)

func file_google_api_policy_proto_rawDescGZIP() []byte {
	file_google_api_policy_proto_rawDescOnce.Do(func() {
		file_google_api_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_policy_proto_rawDescData)
	})
	return file_google_api_policy_proto_rawDescData
}

var file_google_api_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_api_policy_proto_goTypes = []interface{}{
	(*FieldPolicy)(nil),                // 0: google.api.FieldPolicy
	(*MethodPolicy)(nil),               // 1: google.api.MethodPolicy
	(*descriptorpb.FieldOptions)(nil),  // 2: google.protobuf.FieldOptions
	(*descriptorpb.MethodOptions)(nil), // 3: google.protobuf.MethodOptions
}
var file_google_api_policy_proto_depIdxs = []int32{
	0, // 0: google.api.MethodPolicy.request_policies:type_name -> google.api.FieldPolicy
	2, // 1: google.api.field_policy:extendee -> google.protobuf.FieldOptions
	3, // 2: google.api.method_policy:extendee -> google.protobuf.MethodOptions
	0, // 3: google.api.field_policy:type_name -> google.api.FieldPolicy
	1, // 4: google.api.method_policy:type_name -> google.api.MethodPolicy
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_api_policy_proto_init() }
func file_google_api_policy_proto_init() {
	if File_google_api_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldPolicy); i {
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
		file_google_api_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodPolicy); i {
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
			RawDescriptor: file_google_api_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_google_api_policy_proto_goTypes,
		DependencyIndexes: file_google_api_policy_proto_depIdxs,
		MessageInfos:      file_google_api_policy_proto_msgTypes,
		ExtensionInfos:    file_google_api_policy_proto_extTypes,
	}.Build()
	File_google_api_policy_proto = out.File
	file_google_api_policy_proto_rawDesc = nil
	file_google_api_policy_proto_goTypes = nil
	file_google_api_policy_proto_depIdxs = nil
}
