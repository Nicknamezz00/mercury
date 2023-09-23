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
// source: google/api/resource.proto

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

// A description of the historical or future-looking state of the
// resource pattern.
type ResourceDescriptor_History int32

const (
	// The "unset" value.
	ResourceDescriptor_HISTORY_UNSPECIFIED ResourceDescriptor_History = 0
	// The resource originally had one pattern and launched as such, and
	// additional patterns were added later.
	ResourceDescriptor_ORIGINALLY_SINGLE_PATTERN ResourceDescriptor_History = 1
	// The resource has one pattern, but the API owner expects to add more
	// later. (This is the inverse of ORIGINALLY_SINGLE_PATTERN, and prevents
	// that from being necessary once there are multiple patterns.)
	ResourceDescriptor_FUTURE_MULTI_PATTERN ResourceDescriptor_History = 2
)

// Enum value maps for ResourceDescriptor_History.
var (
	ResourceDescriptor_History_name = map[int32]string{
		0: "HISTORY_UNSPECIFIED",
		1: "ORIGINALLY_SINGLE_PATTERN",
		2: "FUTURE_MULTI_PATTERN",
	}
	ResourceDescriptor_History_value = map[string]int32{
		"HISTORY_UNSPECIFIED":       0,
		"ORIGINALLY_SINGLE_PATTERN": 1,
		"FUTURE_MULTI_PATTERN":      2,
	}
)

func (x ResourceDescriptor_History) Enum() *ResourceDescriptor_History {
	p := new(ResourceDescriptor_History)
	*p = x
	return p
}

func (x ResourceDescriptor_History) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceDescriptor_History) Descriptor() protoreflect.EnumDescriptor {
	return file_google_api_resource_proto_enumTypes[0].Descriptor()
}

func (ResourceDescriptor_History) Type() protoreflect.EnumType {
	return &file_google_api_resource_proto_enumTypes[0]
}

func (x ResourceDescriptor_History) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceDescriptor_History.Descriptor instead.
func (ResourceDescriptor_History) EnumDescriptor() ([]byte, []int) {
	return file_google_api_resource_proto_rawDescGZIP(), []int{0, 0}
}

// A flag representing a specific style that a resource claims to conform to.
type ResourceDescriptor_Style int32

const (
	// The unspecified value. Do not use.
	ResourceDescriptor_STYLE_UNSPECIFIED ResourceDescriptor_Style = 0
	// This resource is intended to be "declarative-friendly".
	//
	// Declarative-friendly resources must be more strictly consistent, and
	// setting this to true communicates to tools that this resource should
	// adhere to declarative-friendly expectations.
	//
	// Note: This is used by the API linter (linter.aip.dev) to enable
	// additional checks.
	ResourceDescriptor_DECLARATIVE_FRIENDLY ResourceDescriptor_Style = 1
)

// Enum value maps for ResourceDescriptor_Style.
var (
	ResourceDescriptor_Style_name = map[int32]string{
		0: "STYLE_UNSPECIFIED",
		1: "DECLARATIVE_FRIENDLY",
	}
	ResourceDescriptor_Style_value = map[string]int32{
		"STYLE_UNSPECIFIED":    0,
		"DECLARATIVE_FRIENDLY": 1,
	}
)

func (x ResourceDescriptor_Style) Enum() *ResourceDescriptor_Style {
	p := new(ResourceDescriptor_Style)
	*p = x
	return p
}

func (x ResourceDescriptor_Style) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceDescriptor_Style) Descriptor() protoreflect.EnumDescriptor {
	return file_google_api_resource_proto_enumTypes[1].Descriptor()
}

func (ResourceDescriptor_Style) Type() protoreflect.EnumType {
	return &file_google_api_resource_proto_enumTypes[1]
}

func (x ResourceDescriptor_Style) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceDescriptor_Style.Descriptor instead.
func (ResourceDescriptor_Style) EnumDescriptor() ([]byte, []int) {
	return file_google_api_resource_proto_rawDescGZIP(), []int{0, 1}
}

// A simple descriptor of a resource type.
//
// ResourceDescriptor annotates a resource message (either by means of a
// protobuf annotation or use in the service config), and associates the
// resource's schema, the resource type, and the pattern of the resource name.
//
// Example:
//
//	message Topic {
//	  // Indicates this message defines a resource schema.
//	  // Declares the resource type in the format of {service}/{kind}.
//	  // For Kubernetes resources, the format is {api group}/{kind}.
//	  option (google.api.resource) = {
//	    type: "pubsub.googleapis.com/Topic"
//	    pattern: "projects/{project}/topics/{topic}"
//	  };
//	}
//
// The ResourceDescriptor Yaml config will look like:
//
//	resources:
//	- type: "pubsub.googleapis.com/Topic"
//	  pattern: "projects/{project}/topics/{topic}"
//
// Sometimes, resources have multiple patterns, typically because they can
// live under multiple parents.
//
// Example:
//
//	message LogEntry {
//	  option (google.api.resource) = {
//	    type: "logging.googleapis.com/LogEntry"
//	    pattern: "projects/{project}/logs/{log}"
//	    pattern: "folders/{folder}/logs/{log}"
//	    pattern: "organizations/{organization}/logs/{log}"
//	    pattern: "billingAccounts/{billing_account}/logs/{log}"
//	  };
//	}
//
// The ResourceDescriptor Yaml config will look like:
//
//	resources:
//	- type: 'logging.googleapis.com/LogEntry'
//	  pattern: "projects/{project}/logs/{log}"
//	  pattern: "folders/{folder}/logs/{log}"
//	  pattern: "organizations/{organization}/logs/{log}"
//	  pattern: "billingAccounts/{billing_account}/logs/{log}"
type ResourceDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource type. It must be in the format of
	// {service_name}/{resource_type_kind}. The `resource_type_kind` must be
	// singular and must not include version numbers.
	//
	// Example: `storage.googleapis.com/Bucket`
	//
	// The value of the resource_type_kind must follow the regular expression
	// /[A-Za-z][a-zA-Z0-9]+/. It should start with an upper case character and
	// should use PascalCase (UpperCamelCase). The maximum number of
	// characters allowed for the `resource_type_kind` is 100.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Optional. The relative resource name pattern associated with this resource
	// type. The DNS prefix of the full resource name shouldn't be specified here.
	//
	// The path pattern must follow the syntax, which aligns with HTTP binding
	// syntax:
	//
	//	Template = Segment { "/" Segment } ;
	//	Segment = LITERAL | Variable ;
	//	Variable = "{" LITERAL "}" ;
	//
	// Examples:
	//
	//   - "projects/{project}/topics/{topic}"
	//   - "projects/{project}/knowledgeBases/{knowledge_base}"
	//
	// The components in braces correspond to the IDs for each resource in the
	// hierarchy. It is expected that, if multiple patterns are provided,
	// the same component name (e.g. "project") refers to IDs of the same
	// type of resource.
	Pattern []string `protobuf:"bytes,2,rep,name=pattern,proto3" json:"pattern,omitempty"`
	// Optional. The field on the resource that designates the resource name
	// field. If omitted, this is assumed to be "name".
	NameField string `protobuf:"bytes,3,opt,name=name_field,json=nameField,proto3" json:"name_field,omitempty"`
	// Optional. The historical or future-looking state of the resource pattern.
	//
	// Example:
	//
	//	// The InspectTemplate message originally only supported resource
	//	// names with organization, and project was added later.
	//	message InspectTemplate {
	//	  option (google.api.resource) = {
	//	    type: "dlp.googleapis.com/InspectTemplate"
	//	    pattern:
	//	    "organizations/{organization}/inspectTemplates/{inspect_template}"
	//	    pattern: "projects/{project}/inspectTemplates/{inspect_template}"
	//	    history: ORIGINALLY_SINGLE_PATTERN
	//	  };
	//	}
	History ResourceDescriptor_History `protobuf:"varint,4,opt,name=history,proto3,enum=google.api.ResourceDescriptor_History" json:"history,omitempty"`
	// The plural name used in the resource name and permission names, such as
	// 'projects' for the resource name of 'projects/{project}' and the permission
	// name of 'cloudresourcemanager.googleapis.com/projects.get'. It is the same
	// concept of the `plural` field in k8s CRD spec
	// https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/
	//
	// Note: The plural form is required even for singleton resources. See
	// https://aip.dev/156
	Plural string `protobuf:"bytes,5,opt,name=plural,proto3" json:"plural,omitempty"`
	// The same concept of the `singular` field in k8s CRD spec
	// https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/
	// Such as "project" for the `resourcemanager.googleapis.com/Project` type.
	Singular string `protobuf:"bytes,6,opt,name=singular,proto3" json:"singular,omitempty"`
	// Style flag(s) for this resource.
	// These indicate that a resource is expected to conform to a given
	// style. See the specific style flags for additional information.
	Style []ResourceDescriptor_Style `protobuf:"varint,10,rep,packed,name=style,proto3,enum=google.api.ResourceDescriptor_Style" json:"style,omitempty"`
}

func (x *ResourceDescriptor) Reset() {
	*x = ResourceDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceDescriptor) ProtoMessage() {}

func (x *ResourceDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceDescriptor.ProtoReflect.Descriptor instead.
func (*ResourceDescriptor) Descriptor() ([]byte, []int) {
	return file_google_api_resource_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceDescriptor) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ResourceDescriptor) GetPattern() []string {
	if x != nil {
		return x.Pattern
	}
	return nil
}

func (x *ResourceDescriptor) GetNameField() string {
	if x != nil {
		return x.NameField
	}
	return ""
}

func (x *ResourceDescriptor) GetHistory() ResourceDescriptor_History {
	if x != nil {
		return x.History
	}
	return ResourceDescriptor_HISTORY_UNSPECIFIED
}

func (x *ResourceDescriptor) GetPlural() string {
	if x != nil {
		return x.Plural
	}
	return ""
}

func (x *ResourceDescriptor) GetSingular() string {
	if x != nil {
		return x.Singular
	}
	return ""
}

func (x *ResourceDescriptor) GetStyle() []ResourceDescriptor_Style {
	if x != nil {
		return x.Style
	}
	return nil
}

// Defines a proto annotation that describes a string field that refers to
// an API resource.
type ResourceReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource type that the annotated field references.
	//
	// Example:
	//
	//	message Subscription {
	//	  string topic = 2 [(google.api.resource_reference) = {
	//	    type: "pubsub.googleapis.com/Topic"
	//	  }];
	//	}
	//
	// Occasionally, a field may reference an arbitrary resource. In this case,
	// APIs use the special value * in their resource reference.
	//
	// Example:
	//
	//	message GetIamPolicyRequest {
	//	  string resource = 2 [(google.api.resource_reference) = {
	//	    type: "*"
	//	  }];
	//	}
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The resource type of a child collection that the annotated field
	// references. This is useful for annotating the `parent` field that
	// doesn't have a fixed resource type.
	//
	// Example:
	//
	//	message ListLogEntriesRequest {
	//	  string parent = 1 [(google.api.resource_reference) = {
	//	    child_type: "logging.googleapis.com/LogEntry"
	//	  };
	//	}
	ChildType string `protobuf:"bytes,2,opt,name=child_type,json=childType,proto3" json:"child_type,omitempty"`
}

func (x *ResourceReference) Reset() {
	*x = ResourceReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceReference) ProtoMessage() {}

func (x *ResourceReference) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceReference.ProtoReflect.Descriptor instead.
func (*ResourceReference) Descriptor() ([]byte, []int) {
	return file_google_api_resource_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceReference) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ResourceReference) GetChildType() string {
	if x != nil {
		return x.ChildType
	}
	return ""
}

var file_google_api_resource_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*ResourceReference)(nil),
		Field:         1055,
		Name:          "google.api.resource_reference",
		Tag:           "bytes,1055,opt,name=resource_reference",
		Filename:      "google/api/resource.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: ([]*ResourceDescriptor)(nil),
		Field:         1053,
		Name:          "google.api.resource_definition",
		Tag:           "bytes,1053,rep,name=resource_definition",
		Filename:      "google/api/resource.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*ResourceDescriptor)(nil),
		Field:         1053,
		Name:          "google.api.resource",
		Tag:           "bytes,1053,opt,name=resource",
		Filename:      "google/api/resource.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// An annotation that describes a resource reference, see
	// [ResourceReference][].
	//
	// optional google.api.ResourceReference resource_reference = 1055;
	E_ResourceReference = &file_google_api_resource_proto_extTypes[0]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// An annotation that describes a resource definition without a corresponding
	// message; see [ResourceDescriptor][].
	//
	// repeated google.api.ResourceDescriptor resource_definition = 1053;
	E_ResourceDefinition = &file_google_api_resource_proto_extTypes[1]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// An annotation that describes a resource definition, see
	// [ResourceDescriptor][].
	//
	// optional google.api.ResourceDescriptor resource = 1053;
	E_Resource = &file_google_api_resource_proto_extTypes[2]
)

var File_google_api_resource_proto protoreflect.FileDescriptor

var file_google_api_resource_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x12, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x40, 0x0a,
	0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x72, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x6c, 0x75, 0x72, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x6e, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x6e, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x12, 0x3a, 0x0a, 0x05, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x05, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x22,
	0x5b, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x49,
	0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x4c,
	0x59, 0x5f, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x54, 0x54, 0x45, 0x52, 0x4e,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x4d, 0x55, 0x4c,
	0x54, 0x49, 0x5f, 0x50, 0x41, 0x54, 0x54, 0x45, 0x52, 0x4e, 0x10, 0x02, 0x22, 0x38, 0x0a, 0x05,
	0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x44, 0x45, 0x43, 0x4c, 0x41, 0x52, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x46, 0x52, 0x49, 0x45,
	0x4e, 0x44, 0x4c, 0x59, 0x10, 0x01, 0x22, 0x46, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x6c,
	0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x9f, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x3a, 0x6e, 0x0a, 0x13,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x9d, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x5c, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9d, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0xa1, 0x01, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0d, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x7a, 0x7a, 0x30, 0x30, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x58, 0xaa, 0x02, 0x0a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70, 0x69, 0xca, 0x02, 0x0a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x70, 0x69, 0xe2, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_resource_proto_rawDescOnce sync.Once
	file_google_api_resource_proto_rawDescData = file_google_api_resource_proto_rawDesc
)

func file_google_api_resource_proto_rawDescGZIP() []byte {
	file_google_api_resource_proto_rawDescOnce.Do(func() {
		file_google_api_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_resource_proto_rawDescData)
	})
	return file_google_api_resource_proto_rawDescData
}

var file_google_api_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_api_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_api_resource_proto_goTypes = []interface{}{
	(ResourceDescriptor_History)(0),     // 0: google.api.ResourceDescriptor.History
	(ResourceDescriptor_Style)(0),       // 1: google.api.ResourceDescriptor.Style
	(*ResourceDescriptor)(nil),          // 2: google.api.ResourceDescriptor
	(*ResourceReference)(nil),           // 3: google.api.ResourceReference
	(*descriptorpb.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
	(*descriptorpb.FileOptions)(nil),    // 5: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 6: google.protobuf.MessageOptions
}
var file_google_api_resource_proto_depIdxs = []int32{
	0, // 0: google.api.ResourceDescriptor.history:type_name -> google.api.ResourceDescriptor.History
	1, // 1: google.api.ResourceDescriptor.style:type_name -> google.api.ResourceDescriptor.Style
	4, // 2: google.api.resource_reference:extendee -> google.protobuf.FieldOptions
	5, // 3: google.api.resource_definition:extendee -> google.protobuf.FileOptions
	6, // 4: google.api.resource:extendee -> google.protobuf.MessageOptions
	3, // 5: google.api.resource_reference:type_name -> google.api.ResourceReference
	2, // 6: google.api.resource_definition:type_name -> google.api.ResourceDescriptor
	2, // 7: google.api.resource:type_name -> google.api.ResourceDescriptor
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	5, // [5:8] is the sub-list for extension type_name
	2, // [2:5] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_api_resource_proto_init() }
func file_google_api_resource_proto_init() {
	if File_google_api_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceDescriptor); i {
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
		file_google_api_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceReference); i {
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
			RawDescriptor: file_google_api_resource_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_google_api_resource_proto_goTypes,
		DependencyIndexes: file_google_api_resource_proto_depIdxs,
		EnumInfos:         file_google_api_resource_proto_enumTypes,
		MessageInfos:      file_google_api_resource_proto_msgTypes,
		ExtensionInfos:    file_google_api_resource_proto_extTypes,
	}.Build()
	File_google_api_resource_proto = out.File
	file_google_api_resource_proto_rawDesc = nil
	file_google_api_resource_proto_goTypes = nil
	file_google_api_resource_proto_depIdxs = nil
}
