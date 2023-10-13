# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [google/api/http.proto](#google_api_http-proto)
    - [CustomHttpPattern](#google-api-CustomHttpPattern)
    - [Http](#google-api-Http)
    - [HttpRule](#google-api-HttpRule)
  
- [google/api/annotations.proto](#google_api_annotations-proto)
    - [File-level Extensions](#google_api_annotations-proto-extensions)
  
- [google/api/launch_stage.proto](#google_api_launch_stage-proto)
    - [LaunchStage](#google-api-LaunchStage)
  
- [google/api/client.proto](#google_api_client-proto)
    - [ClientLibrarySettings](#google-api-ClientLibrarySettings)
    - [CommonLanguageSettings](#google-api-CommonLanguageSettings)
    - [CppSettings](#google-api-CppSettings)
    - [DotnetSettings](#google-api-DotnetSettings)
    - [DotnetSettings.RenamedResourcesEntry](#google-api-DotnetSettings-RenamedResourcesEntry)
    - [DotnetSettings.RenamedServicesEntry](#google-api-DotnetSettings-RenamedServicesEntry)
    - [GoSettings](#google-api-GoSettings)
    - [JavaSettings](#google-api-JavaSettings)
    - [JavaSettings.ServiceClassNamesEntry](#google-api-JavaSettings-ServiceClassNamesEntry)
    - [MethodSettings](#google-api-MethodSettings)
    - [MethodSettings.LongRunning](#google-api-MethodSettings-LongRunning)
    - [NodeSettings](#google-api-NodeSettings)
    - [PhpSettings](#google-api-PhpSettings)
    - [Publishing](#google-api-Publishing)
    - [PythonSettings](#google-api-PythonSettings)
    - [RubySettings](#google-api-RubySettings)
  
    - [ClientLibraryDestination](#google-api-ClientLibraryDestination)
    - [ClientLibraryOrganization](#google-api-ClientLibraryOrganization)
  
    - [File-level Extensions](#google_api_client-proto-extensions)
    - [File-level Extensions](#google_api_client-proto-extensions)
    - [File-level Extensions](#google_api_client-proto-extensions)
  
- [google/api/field_behavior.proto](#google_api_field_behavior-proto)
    - [FieldBehavior](#google-api-FieldBehavior)
  
    - [File-level Extensions](#google_api_field_behavior-proto-extensions)
  
- [google/api/auth.proto](#google_api_auth-proto)
    - [AuthProvider](#google-api-AuthProvider)
    - [AuthRequirement](#google-api-AuthRequirement)
    - [Authentication](#google-api-Authentication)
    - [AuthenticationRule](#google-api-AuthenticationRule)
    - [JwtLocation](#google-api-JwtLocation)
    - [OAuthRequirements](#google-api-OAuthRequirements)
  
- [google/api/backend.proto](#google_api_backend-proto)
    - [Backend](#google-api-Backend)
    - [BackendRule](#google-api-BackendRule)
    - [BackendRule.OverridesByRequestProtocolEntry](#google-api-BackendRule-OverridesByRequestProtocolEntry)
  
    - [BackendRule.PathTranslation](#google-api-BackendRule-PathTranslation)
  
- [google/api/billing.proto](#google_api_billing-proto)
    - [Billing](#google-api-Billing)
    - [Billing.BillingDestination](#google-api-Billing-BillingDestination)
  
- [google/api/config_change.proto](#google_api_config_change-proto)
    - [Advice](#google-api-Advice)
    - [ConfigChange](#google-api-ConfigChange)
  
    - [ChangeType](#google-api-ChangeType)
  
- [google/api/consumer.proto](#google_api_consumer-proto)
    - [ProjectProperties](#google-api-ProjectProperties)
    - [Property](#google-api-Property)
  
    - [Property.PropertyType](#google-api-Property-PropertyType)
  
- [google/api/context.proto](#google_api_context-proto)
    - [Context](#google-api-Context)
    - [ContextRule](#google-api-ContextRule)
  
- [google/api/policy.proto](#google_api_policy-proto)
    - [FieldPolicy](#google-api-FieldPolicy)
    - [MethodPolicy](#google-api-MethodPolicy)
  
    - [File-level Extensions](#google_api_policy-proto-extensions)
    - [File-level Extensions](#google_api_policy-proto-extensions)
  
- [google/api/control.proto](#google_api_control-proto)
    - [Control](#google-api-Control)
  
- [google/api/distribution.proto](#google_api_distribution-proto)
    - [Distribution](#google-api-Distribution)
    - [Distribution.BucketOptions](#google-api-Distribution-BucketOptions)
    - [Distribution.BucketOptions.Explicit](#google-api-Distribution-BucketOptions-Explicit)
    - [Distribution.BucketOptions.Exponential](#google-api-Distribution-BucketOptions-Exponential)
    - [Distribution.BucketOptions.Linear](#google-api-Distribution-BucketOptions-Linear)
    - [Distribution.Exemplar](#google-api-Distribution-Exemplar)
    - [Distribution.Range](#google-api-Distribution-Range)
  
- [google/api/documentation.proto](#google_api_documentation-proto)
    - [Documentation](#google-api-Documentation)
    - [DocumentationRule](#google-api-DocumentationRule)
    - [Page](#google-api-Page)
  
- [google/api/endpoint.proto](#google_api_endpoint-proto)
    - [Endpoint](#google-api-Endpoint)
  
- [google/api/error_reason.proto](#google_api_error_reason-proto)
    - [ErrorReason](#google-api-ErrorReason)
  
- [google/api/httpbody.proto](#google_api_httpbody-proto)
    - [HttpBody](#google-api-HttpBody)
  
- [google/api/label.proto](#google_api_label-proto)
    - [LabelDescriptor](#google-api-LabelDescriptor)
  
    - [LabelDescriptor.ValueType](#google-api-LabelDescriptor-ValueType)
  
- [google/api/log.proto](#google_api_log-proto)
    - [LogDescriptor](#google-api-LogDescriptor)
  
- [google/api/logging.proto](#google_api_logging-proto)
    - [Logging](#google-api-Logging)
    - [Logging.LoggingDestination](#google-api-Logging-LoggingDestination)
  
- [google/api/metric.proto](#google_api_metric-proto)
    - [Metric](#google-api-Metric)
    - [Metric.LabelsEntry](#google-api-Metric-LabelsEntry)
    - [MetricDescriptor](#google-api-MetricDescriptor)
    - [MetricDescriptor.MetricDescriptorMetadata](#google-api-MetricDescriptor-MetricDescriptorMetadata)
  
    - [MetricDescriptor.MetricKind](#google-api-MetricDescriptor-MetricKind)
    - [MetricDescriptor.ValueType](#google-api-MetricDescriptor-ValueType)
  
- [google/api/monitored_resource.proto](#google_api_monitored_resource-proto)
    - [MonitoredResource](#google-api-MonitoredResource)
    - [MonitoredResource.LabelsEntry](#google-api-MonitoredResource-LabelsEntry)
    - [MonitoredResourceDescriptor](#google-api-MonitoredResourceDescriptor)
    - [MonitoredResourceMetadata](#google-api-MonitoredResourceMetadata)
    - [MonitoredResourceMetadata.UserLabelsEntry](#google-api-MonitoredResourceMetadata-UserLabelsEntry)
  
- [google/api/monitoring.proto](#google_api_monitoring-proto)
    - [Monitoring](#google-api-Monitoring)
    - [Monitoring.MonitoringDestination](#google-api-Monitoring-MonitoringDestination)
  
- [google/api/quota.proto](#google_api_quota-proto)
    - [MetricRule](#google-api-MetricRule)
    - [MetricRule.MetricCostsEntry](#google-api-MetricRule-MetricCostsEntry)
    - [Quota](#google-api-Quota)
    - [QuotaLimit](#google-api-QuotaLimit)
    - [QuotaLimit.ValuesEntry](#google-api-QuotaLimit-ValuesEntry)
  
- [google/api/resource.proto](#google_api_resource-proto)
    - [ResourceDescriptor](#google-api-ResourceDescriptor)
    - [ResourceReference](#google-api-ResourceReference)
  
    - [ResourceDescriptor.History](#google-api-ResourceDescriptor-History)
    - [ResourceDescriptor.Style](#google-api-ResourceDescriptor-Style)
  
    - [File-level Extensions](#google_api_resource-proto-extensions)
    - [File-level Extensions](#google_api_resource-proto-extensions)
    - [File-level Extensions](#google_api_resource-proto-extensions)
  
- [google/api/routing.proto](#google_api_routing-proto)
    - [RoutingParameter](#google-api-RoutingParameter)
    - [RoutingRule](#google-api-RoutingRule)
  
    - [File-level Extensions](#google_api_routing-proto-extensions)
  
- [google/api/source_info.proto](#google_api_source_info-proto)
    - [SourceInfo](#google-api-SourceInfo)
  
- [google/api/system_parameter.proto](#google_api_system_parameter-proto)
    - [SystemParameter](#google-api-SystemParameter)
    - [SystemParameterRule](#google-api-SystemParameterRule)
    - [SystemParameters](#google-api-SystemParameters)
  
- [google/api/usage.proto](#google_api_usage-proto)
    - [Usage](#google-api-Usage)
    - [UsageRule](#google-api-UsageRule)
  
- [google/api/service.proto](#google_api_service-proto)
    - [Service](#google-api-Service)
  
- [google/api/visibility.proto](#google_api_visibility-proto)
    - [Visibility](#google-api-Visibility)
    - [VisibilityRule](#google-api-VisibilityRule)
  
    - [File-level Extensions](#google_api_visibility-proto-extensions)
    - [File-level Extensions](#google_api_visibility-proto-extensions)
    - [File-level Extensions](#google_api_visibility-proto-extensions)
    - [File-level Extensions](#google_api_visibility-proto-extensions)
    - [File-level Extensions](#google_api_visibility-proto-extensions)
    - [File-level Extensions](#google_api_visibility-proto-extensions)
  
- [Scalar Value Types](#scalar-value-types)



<a name="google_api_http-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/http.proto



<a name="google-api-CustomHttpPattern"></a>

### CustomHttpPattern
A custom pattern is used for defining custom HTTP verb.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  | The name of this custom HTTP verb. |
| path | [string](#string) |  | The path matched by this custom verb. |






<a name="google-api-Http"></a>

### Http
Defines the HTTP configuration for an API service. It contains a list of
[HttpRule][google.api.HttpRule], each specifying the mapping of an RPC method
to one or more HTTP REST API methods.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rules | [HttpRule](#google-api-HttpRule) | repeated | A list of HTTP configuration rules that apply to individual API methods.

**NOTE:** All service configuration rules follow &#34;last one wins&#34; order. |
| fully_decode_reserved_expansion | [bool](#bool) |  | When set to true, URL path parameters will be fully URI-decoded except in cases of single segment matches in reserved expansion, where &#34;%2F&#34; will be left encoded.

The default behavior is to not decode RFC 6570 reserved characters in multi segment matches. |






<a name="google-api-HttpRule"></a>

### HttpRule
# gRPC Transcoding

gRPC Transcoding is a feature for mapping between a gRPC method and one or
more HTTP REST endpoints. It allows developers to build a single API service
that supports both gRPC APIs and REST APIs. Many systems, including [Google
APIs](https://github.com/googleapis/googleapis),
[Cloud Endpoints](https://cloud.google.com/endpoints), [gRPC
Gateway](https://github.com/grpc-ecosystem/grpc-gateway),
and [Envoy](https://github.com/envoyproxy/envoy) proxy support this feature
and use it for large scale production services.

`HttpRule` defines the schema of the gRPC/REST mapping. The mapping specifies
how different portions of the gRPC request message are mapped to the URL
path, URL query parameters, and HTTP request body. It also controls how the
gRPC response message is mapped to the HTTP response body. `HttpRule` is
typically specified as an `google.api.http` annotation on the gRPC method.

Each mapping specifies a URL path template and an HTTP method. The path
template may refer to one or more fields in the gRPC request message, as long
as each field is a non-repeated field with a primitive (non-message) type.
The path template controls how fields of the request message are mapped to
the URL path.

Example:

    service Messaging {
      rpc GetMessage(GetMessageRequest) returns (Message) {
        option (google.api.http) = {
            get: &#34;/v1/{name=messages/*}&#34;
        };
      }
    }
    message GetMessageRequest {
      string name = 1; // Mapped to URL path.
    }
    message Message {
      string text = 1; // The resource content.
    }

This enables an HTTP REST to gRPC mapping as below:

HTTP | gRPC
-----|-----
`GET /v1/messages/123456`  | `GetMessage(name: &#34;messages/123456&#34;)`

Any fields in the request message which are not bound by the path template
automatically become HTTP query parameters if there is no HTTP request body.
For example:

    service Messaging {
      rpc GetMessage(GetMessageRequest) returns (Message) {
        option (google.api.http) = {
            get:&#34;/v1/messages/{message_id}&#34;
        };
      }
    }
    message GetMessageRequest {
      message SubMessage {
        string subfield = 1;
      }
      string message_id = 1; // Mapped to URL path.
      int64 revision = 2;    // Mapped to URL query parameter `revision`.
      SubMessage sub = 3;    // Mapped to URL query parameter `sub.subfield`.
    }

This enables a HTTP JSON to RPC mapping as below:

HTTP | gRPC
-----|-----
`GET /v1/messages/123456?revision=2&amp;sub.subfield=foo` |
`GetMessage(message_id: &#34;123456&#34; revision: 2 sub: SubMessage(subfield:
&#34;foo&#34;))`

Note that fields which are mapped to URL query parameters must have a
primitive type or a repeated primitive type or a non-repeated message type.
In the case of a repeated type, the parameter can be repeated in the URL
as `...?param=A&amp;param=B`. In the case of a message type, each field of the
message is mapped to a separate parameter, such as
`...?foo.a=A&amp;foo.b=B&amp;foo.c=C`.

For HTTP methods that allow a request body, the `body` field
specifies the mapping. Consider a REST update method on the
message resource collection:

    service Messaging {
      rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
        option (google.api.http) = {
          patch: &#34;/v1/messages/{message_id}&#34;
          body: &#34;message&#34;
        };
      }
    }
    message UpdateMessageRequest {
      string message_id = 1; // mapped to the URL
      Message message = 2;   // mapped to the body
    }

The following HTTP JSON to RPC mapping is enabled, where the
representation of the JSON in the request body is determined by
protos JSON encoding:

HTTP | gRPC
-----|-----
`PATCH /v1/messages/123456 { &#34;text&#34;: &#34;Hi!&#34; }` | `UpdateMessage(message_id:
&#34;123456&#34; message { text: &#34;Hi!&#34; })`

The special name `*` can be used in the body mapping to define that
every field not bound by the path template should be mapped to the
request body.  This enables the following alternative definition of
the update method:

    service Messaging {
      rpc UpdateMessage(Message) returns (Message) {
        option (google.api.http) = {
          patch: &#34;/v1/messages/{message_id}&#34;
          body: &#34;*&#34;
        };
      }
    }
    message Message {
      string message_id = 1;
      string text = 2;
    }


The following HTTP JSON to RPC mapping is enabled:

HTTP | gRPC
-----|-----
`PATCH /v1/messages/123456 { &#34;text&#34;: &#34;Hi!&#34; }` | `UpdateMessage(message_id:
&#34;123456&#34; text: &#34;Hi!&#34;)`

Note that when using `*` in the body mapping, it is not possible to
have HTTP parameters, as all fields not bound by the path end in
the body. This makes this option more rarely used in practice when
defining REST APIs. The common usage of `*` is in custom methods
which don&#39;t use the URL at all for transferring data.

It is possible to define multiple HTTP methods for one RPC by using
the `additional_bindings` option. Example:

    service Messaging {
      rpc GetMessage(GetMessageRequest) returns (Message) {
        option (google.api.http) = {
          get: &#34;/v1/messages/{message_id}&#34;
          additional_bindings {
            get: &#34;/v1/users/{user_id}/messages/{message_id}&#34;
          }
        };
      }
    }
    message GetMessageRequest {
      string message_id = 1;
      string user_id = 2;
    }

This enables the following two alternative HTTP JSON to RPC mappings:

HTTP | gRPC
-----|-----
`GET /v1/messages/123456` | `GetMessage(message_id: &#34;123456&#34;)`
`GET /v1/users/me/messages/123456` | `GetMessage(user_id: &#34;me&#34; message_id:
&#34;123456&#34;)`

## Rules for HTTP mapping

1. Leaf request fields (recursive expansion nested messages in the request
   message) are classified into three categories:
   - Fields referred by the path template. They are passed via the URL path.
   - Fields referred by the [HttpRule.body][google.api.HttpRule.body]. They
   are passed via the HTTP
     request body.
   - All other fields are passed via the URL query parameters, and the
     parameter name is the field path in the request message. A repeated
     field can be represented as multiple query parameters under the same
     name.
 2. If [HttpRule.body][google.api.HttpRule.body] is &#34;*&#34;, there is no URL
 query parameter, all fields
    are passed via URL path and HTTP request body.
 3. If [HttpRule.body][google.api.HttpRule.body] is omitted, there is no HTTP
 request body, all
    fields are passed via URL path and URL query parameters.

### Path template syntax

    Template = &#34;/&#34; Segments [ Verb ] ;
    Segments = Segment { &#34;/&#34; Segment } ;
    Segment  = &#34;*&#34; | &#34;**&#34; | LITERAL | Variable ;
    Variable = &#34;{&#34; FieldPath [ &#34;=&#34; Segments ] &#34;}&#34; ;
    FieldPath = IDENT { &#34;.&#34; IDENT } ;
    Verb     = &#34;:&#34; LITERAL ;

The syntax `*` matches a single URL path segment. The syntax `**` matches
zero or more URL path segments, which must be the last part of the URL path
except the `Verb`.

The syntax `Variable` matches part of the URL path as specified by its
template. A variable template must not contain other variables. If a variable
matches a single path segment, its template may be omitted, e.g. `{var}`
is equivalent to `{var=*}`.

The syntax `LITERAL` matches literal text in the URL path. If the `LITERAL`
contains any reserved character, such characters should be percent-encoded
before the matching.

If a variable contains exactly one path segment, such as `&#34;{var}&#34;` or
`&#34;{var=*}&#34;`, when such a variable is expanded into a URL path on the client
side, all characters except `[-_.~0-9a-zA-Z]` are percent-encoded. The
server side does the reverse decoding. Such variables show up in the
[Discovery
Document](https://developers.google.com/discovery/v1/reference/apis) as
`{var}`.

If a variable contains multiple path segments, such as `&#34;{var=foo/*}&#34;`
or `&#34;{var=**}&#34;`, when such a variable is expanded into a URL path on the
client side, all characters except `[-_.~/0-9a-zA-Z]` are percent-encoded.
The server side does the reverse decoding, except &#34;%2F&#34; and &#34;%2f&#34; are left
unchanged. Such variables show up in the
[Discovery
Document](https://developers.google.com/discovery/v1/reference/apis) as
`{&#43;var}`.

## Using gRPC API Service Configuration

gRPC API Service Configuration (service config) is a configuration language
for configuring a gRPC service to become a user-facing product. The
service config is simply the YAML representation of the `google.api.Service`
proto message.

As an alternative to annotating your proto file, you can configure gRPC
transcoding in your service config YAML files. You do this by specifying a
`HttpRule` that maps the gRPC method to a REST endpoint, achieving the same
effect as the proto annotation. This can be particularly useful if you
have a proto that is reused in multiple services. Note that any transcoding
specified in the service config will override any matching transcoding
configuration in the proto.

Example:

    http:
      rules:
        # Selects a gRPC method and applies HttpRule to it.
        - selector: example.v1.Messaging.GetMessage
          get: /v1/messages/{message_id}/{sub.subfield}

## Special notes

When gRPC Transcoding is used to map a gRPC to JSON REST endpoints, the
proto to JSON conversion must follow the [proto3
specification](https://developers.google.com/protocol-buffers/docs/proto3#json).

While the single segment variable follows the semantics of
[RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2 Simple String
Expansion, the multi segment variable **does not** follow RFC 6570 Section
3.2.3 Reserved Expansion. The reason is that the Reserved Expansion
does not expand special characters like `?` and `#`, which would lead
to invalid URLs. As the result, gRPC Transcoding uses a custom encoding
for multi segment variables.

The path variables **must not** refer to any repeated or mapped field,
because client libraries are not capable of handling such variable expansion.

The path variables **must not** capture the leading &#34;/&#34; character. The reason
is that the most common use case &#34;{var}&#34; does not capture the leading &#34;/&#34;
character. For consistency, all path variables must share the same behavior.

Repeated message fields must not be mapped to URL query parameters, because
no client library can support such complicated mapping.

If an API needs to use a JSON array for request or response body, it can map
the request or response body to a repeated field. However, some gRPC
Transcoding implementations may not support this feature.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | Selects a method to which this rule applies.

Refer to [selector][google.api.DocumentationRule.selector] for syntax details. |
| get | [string](#string) |  | Maps to HTTP GET. Used for listing and getting information about resources. |
| put | [string](#string) |  | Maps to HTTP PUT. Used for replacing a resource. |
| post | [string](#string) |  | Maps to HTTP POST. Used for creating a resource or performing an action. |
| delete | [string](#string) |  | Maps to HTTP DELETE. Used for deleting a resource. |
| patch | [string](#string) |  | Maps to HTTP PATCH. Used for updating a resource. |
| custom | [CustomHttpPattern](#google-api-CustomHttpPattern) |  | The custom pattern is used for specifying an HTTP method that is not included in the `pattern` field, such as HEAD, or &#34;*&#34; to leave the HTTP method unspecified for this rule. The wild-card rule is useful for services that provide content to Web (HTML) clients. |
| body | [string](#string) |  | The name of the request field whose value is mapped to the HTTP request body, or `*` for mapping all request fields not captured by the path pattern to the HTTP body, or omitted for not having any HTTP request body.

NOTE: the referred field must be present at the top-level of the request message type. |
| response_body | [string](#string) |  | Optional. The name of the response field whose value is mapped to the HTTP response body. When omitted, the entire response message will be used as the HTTP response body.

NOTE: The referred field must be present at the top-level of the response message type. |
| additional_bindings | [HttpRule](#google-api-HttpRule) | repeated | Additional HTTP bindings for the selector. Nested bindings must not contain an `additional_bindings` field themselves (that is, the nesting may only be one level deep). |





 

 

 

 



<a name="google_api_annotations-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/annotations.proto


 

 


<a name="google_api_annotations-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| http | HttpRule | .google.protobuf.MethodOptions | 72295728 | See `HttpRule`. |

 

 



<a name="google_api_launch_stage-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/launch_stage.proto


 


<a name="google-api-LaunchStage"></a>

### LaunchStage
The launch stage as defined by [Google Cloud Platform
Launch Stages](https://cloud.google.com/terms/launch-stages).

| Name | Number | Description |
| ---- | ------ | ----------- |
| LAUNCH_STAGE_UNSPECIFIED | 0 | Do not use this default value. |
| UNIMPLEMENTED | 6 | The feature is not yet implemented. Users can not use it. |
| PRELAUNCH | 7 | Prelaunch features are hidden from users and are only visible internally. |
| EARLY_ACCESS | 1 | Early Access features are limited to a closed group of testers. To use these features, you must sign up in advance and sign a Trusted Tester agreement (which includes confidentiality provisions). These features may be unstable, changed in backward-incompatible ways, and are not guaranteed to be released. |
| ALPHA | 2 | Alpha is a limited availability test for releases before they are cleared for widespread use. By Alpha, all significant design issues are resolved and we are in the process of verifying functionality. Alpha customers need to apply for access, agree to applicable terms, and have their projects allowlisted. Alpha releases don&#39;t have to be feature complete, no SLAs are provided, and there are no technical support obligations, but they will be far enough along that customers can actually use them in test environments or for limited-use tests -- just like they would in normal production cases. |
| BETA | 3 | Beta is the point at which we are ready to open a release for any customer to use. There are no SLA or technical support obligations in a Beta release. Products will be complete from a feature perspective, but may have some open outstanding issues. Beta releases are suitable for limited production use cases. |
| GA | 4 | GA features are open to all developers and are considered stable and fully qualified for production use. |
| DEPRECATED | 5 | Deprecated features are scheduled to be shut down and removed. For more information, see the &#34;Deprecation Policy&#34; section of our [Terms of Service](https://cloud.google.com/terms/) and the [Google Cloud Platform Subject to the Deprecation Policy](https://cloud.google.com/terms/deprecation) documentation. |


 

 

 



<a name="google_api_client-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/client.proto



<a name="google-api-ClientLibrarySettings"></a>

### ClientLibrarySettings
Details about how and where to publish client libraries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  | Version of the API to apply these settings to. This is the full protobuf package for the API, ending in the version element. Examples: &#34;google.cloud.speech.v1&#34; and &#34;google.spanner.admin.database.v1&#34;. |
| launch_stage | [LaunchStage](#google-api-LaunchStage) |  | Launch stage of this version of the API. |
| rest_numeric_enums | [bool](#bool) |  | When using transport=rest, the client request will encode enums as numbers rather than strings. |
| java_settings | [JavaSettings](#google-api-JavaSettings) |  | Settings for legacy Java features, supported in the Service YAML. |
| cpp_settings | [CppSettings](#google-api-CppSettings) |  | Settings for C&#43;&#43; client libraries. |
| php_settings | [PhpSettings](#google-api-PhpSettings) |  | Settings for PHP client libraries. |
| python_settings | [PythonSettings](#google-api-PythonSettings) |  | Settings for Python client libraries. |
| node_settings | [NodeSettings](#google-api-NodeSettings) |  | Settings for Node client libraries. |
| dotnet_settings | [DotnetSettings](#google-api-DotnetSettings) |  | Settings for .NET client libraries. |
| ruby_settings | [RubySettings](#google-api-RubySettings) |  | Settings for Ruby client libraries. |
| go_settings | [GoSettings](#google-api-GoSettings) |  | Settings for Go client libraries. |






<a name="google-api-CommonLanguageSettings"></a>

### CommonLanguageSettings
Required information for every language.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reference_docs_uri | [string](#string) |  | **Deprecated.** Link to automatically generated reference documentation. Example: https://cloud.google.com/nodejs/docs/reference/asset/latest |
| destinations | [ClientLibraryDestination](#google-api-ClientLibraryDestination) | repeated | The destination where API teams want this client library to be published. |






<a name="google-api-CppSettings"></a>

### CppSettings
Settings for C&#43;&#43; client libraries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| common | [CommonLanguageSettings](#google-api-CommonLanguageSettings) |  | Some settings. |






<a name="google-api-DotnetSettings"></a>

### DotnetSettings
Settings for Dotnet client libraries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| common | [CommonLanguageSettings](#google-api-CommonLanguageSettings) |  | Some settings. |
| renamed_services | [DotnetSettings.RenamedServicesEntry](#google-api-DotnetSettings-RenamedServicesEntry) | repeated | Map from original service names to renamed versions. This is used when the default generated types would cause a naming conflict. (Neither name is fully-qualified.) Example: Subscriber to SubscriberServiceApi. |
| renamed_resources | [DotnetSettings.RenamedResourcesEntry](#google-api-DotnetSettings-RenamedResourcesEntry) | repeated | Map from full resource types to the effective short name for the resource. This is used when otherwise resource named from different services would cause naming collisions. Example entry: &#34;datalabeling.googleapis.com/Dataset&#34;: &#34;DataLabelingDataset&#34; |
| ignored_resources | [string](#string) | repeated | List of full resource types to ignore during generation. This is typically used for API-specific Location resources, which should be handled by the generator as if they were actually the common Location resources. Example entry: &#34;documentai.googleapis.com/Location&#34; |
| forced_namespace_aliases | [string](#string) | repeated | Namespaces which must be aliased in snippets due to a known (but non-generator-predictable) naming collision |
| handwritten_signatures | [string](#string) | repeated | Method signatures (in the form &#34;service.method(signature)&#34;) which are provided separately, so shouldn&#39;t be generated. Snippets *calling* these methods are still generated, however. |






<a name="google-api-DotnetSettings-RenamedResourcesEntry"></a>

### DotnetSettings.RenamedResourcesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="google-api-DotnetSettings-RenamedServicesEntry"></a>

### DotnetSettings.RenamedServicesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="google-api-GoSettings"></a>

### GoSettings
Settings for Go client libraries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| common | [CommonLanguageSettings](#google-api-CommonLanguageSettings) |  | Some settings. |






<a name="google-api-JavaSettings"></a>

### JavaSettings
Settings for Java client libraries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| library_package | [string](#string) |  | The package name to use in Java. Clobbers the java_package option set in the protobuf. This should be used **only** by APIs who have already set the language_settings.java.package_name&#34; field in gapic.yaml. API teams should use the protobuf java_package option where possible.

Example of a YAML configuration::

 publishing: java_settings: library_package: com.google.cloud.pubsub.v1 |
| service_class_names | [JavaSettings.ServiceClassNamesEntry](#google-api-JavaSettings-ServiceClassNamesEntry) | repeated | Configure the Java class name to use instead of the service&#39;s for its corresponding generated GAPIC client. Keys are fully-qualified service names as they appear in the protobuf (including the full the language_settings.java.interface_names&#34; field in gapic.yaml. API teams should otherwise use the service name as it appears in the protobuf.

Example of a YAML configuration::

 publishing: java_settings: service_class_names: - google.pubsub.v1.Publisher: TopicAdmin - google.pubsub.v1.Subscriber: SubscriptionAdmin |
| common | [CommonLanguageSettings](#google-api-CommonLanguageSettings) |  | Some settings. |






<a name="google-api-JavaSettings-ServiceClassNamesEntry"></a>

### JavaSettings.ServiceClassNamesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="google-api-MethodSettings"></a>

### MethodSettings
Describes the generator configuration for a method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | The fully qualified name of the method, for which the options below apply. This is used to find the method to apply the options. |
| long_running | [MethodSettings.LongRunning](#google-api-MethodSettings-LongRunning) |  | Describes settings to use for long-running operations when generating API methods for RPCs. Complements RPCs that use the annotations in google/longrunning/operations.proto.

Example of a YAML configuration::

 publishing: method_settings: - selector: google.cloud.speech.v2.Speech.BatchRecognize long_running: initial_poll_delay: seconds: 60 # 1 minute poll_delay_multiplier: 1.5 max_poll_delay: seconds: 360 # 6 minutes total_poll_timeout: seconds: 54000 # 90 minutes |






<a name="google-api-MethodSettings-LongRunning"></a>

### MethodSettings.LongRunning
Describes settings to use when generating API methods that use the
long-running operation pattern.
All default values below are from those used in the client library
generators (e.g.
[Java](https://github.com/googleapis/gapic-generator-java/blob/04c2faa191a9b5a10b92392fe8482279c4404803/src/main/java/com/google/api/generator/gapic/composer/common/RetrySettingsComposer.java)).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| initial_poll_delay | [google.protobuf.Duration](#google-protobuf-Duration) |  | Initial delay after which the first poll request will be made. Default value: 5 seconds. |
| poll_delay_multiplier | [float](#float) |  | Multiplier to gradually increase delay between subsequent polls until it reaches max_poll_delay. Default value: 1.5. |
| max_poll_delay | [google.protobuf.Duration](#google-protobuf-Duration) |  | Maximum time between two subsequent poll requests. Default value: 45 seconds. |
| total_poll_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  | Total polling timeout. Default value: 5 minutes. |






<a name="google-api-NodeSettings"></a>

### NodeSettings
Settings for Node client libraries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| common | [CommonLanguageSettings](#google-api-CommonLanguageSettings) |  | Some settings. |






<a name="google-api-PhpSettings"></a>

### PhpSettings
Settings for Php client libraries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| common | [CommonLanguageSettings](#google-api-CommonLanguageSettings) |  | Some settings. |






<a name="google-api-Publishing"></a>

### Publishing
This message configures the settings for publishing [Google Cloud Client
libraries](https://cloud.google.com/apis/docs/cloud-client-libraries)
generated from the service config.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| method_settings | [MethodSettings](#google-api-MethodSettings) | repeated | A list of API method settings, e.g. the behavior for methods that use the long-running operation pattern. |
| new_issue_uri | [string](#string) |  | Link to a *public* URI where users can report issues. Example: https://issuetracker.google.com/issues/new?component=190865&amp;template=1161103 |
| documentation_uri | [string](#string) |  | Link to product home page. Example: https://cloud.google.com/asset-inventory/docs/overview |
| api_short_name | [string](#string) |  | Used as a tracking tag when collecting data about the APIs developer relations artifacts like docs, packages delivered to package managers, etc. Example: &#34;speech&#34;. |
| github_label | [string](#string) |  | GitHub label to apply to issues and pull requests opened for this API. |
| codeowner_github_teams | [string](#string) | repeated | GitHub teams to be added to CODEOWNERS in the directory in GitHub containing source code for the client libraries for this API. |
| doc_tag_prefix | [string](#string) |  | A prefix used in sample code when demarking regions to be included in documentation. |
| organization | [ClientLibraryOrganization](#google-api-ClientLibraryOrganization) |  | For whom the client library is being published. |
| library_settings | [ClientLibrarySettings](#google-api-ClientLibrarySettings) | repeated | Client library settings. If the same version string appears multiple times in this list, then the last one wins. Settings from earlier settings with the same version string are discarded. |
| proto_reference_documentation_uri | [string](#string) |  | Optional link to proto reference documentation. Example: https://cloud.google.com/pubsub/lite/docs/reference/rpc |






<a name="google-api-PythonSettings"></a>

### PythonSettings
Settings for Python client libraries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| common | [CommonLanguageSettings](#google-api-CommonLanguageSettings) |  | Some settings. |






<a name="google-api-RubySettings"></a>

### RubySettings
Settings for Ruby client libraries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| common | [CommonLanguageSettings](#google-api-CommonLanguageSettings) |  | Some settings. |





 


<a name="google-api-ClientLibraryDestination"></a>

### ClientLibraryDestination
To where should client libraries be published?

| Name | Number | Description |
| ---- | ------ | ----------- |
| CLIENT_LIBRARY_DESTINATION_UNSPECIFIED | 0 | Client libraries will neither be generated nor published to package managers. |
| GITHUB | 10 | Generate the client library in a repo under github.com/googleapis, but don&#39;t publish it to package managers. |
| PACKAGE_MANAGER | 20 | Publish the library to package managers like nuget.org and npmjs.com. |



<a name="google-api-ClientLibraryOrganization"></a>

### ClientLibraryOrganization
The organization for which the client libraries are being published.
Affects the url where generated docs are published, etc.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CLIENT_LIBRARY_ORGANIZATION_UNSPECIFIED | 0 | Not useful. |
| CLOUD | 1 | Google Cloud Platform Org. |
| ADS | 2 | Ads (Advertising) Org. |
| PHOTOS | 3 | Photos Org. |
| STREET_VIEW | 4 | Street View Org. |
| SHOPPING | 5 | Shopping Org. |
| GEO | 6 | Geo Org. |
| GENERATIVE_AI | 7 | Generative AI - https://developers.generativeai.google |


 


<a name="google_api_client-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| method_signature | string | .google.protobuf.MethodOptions | 1051 | A definition of a client library method signature.

In client libraries, each proto RPC corresponds to one or more methods which the end user is able to call, and calls the underlying RPC. Normally, this method receives a single argument (a struct or instance corresponding to the RPC request object). Defining this field will add one or more overloads providing flattened or simpler method signatures in some languages.

The fields on the method signature are provided as a comma-separated string.

For example, the proto RPC and annotation:

 rpc CreateSubscription(CreateSubscriptionRequest) returns (Subscription) { option (google.api.method_signature) = &#34;name,topic&#34;; }

Would add the following Java overload (in addition to the method accepting the request object):

 public final Subscription createSubscription(String name, String topic)

The following backwards-compatibility guidelines apply:

 * Adding this annotation to an unannotated method is backwards compatible. * Adding this annotation to a method which already has existing method signature annotations is backwards compatible if and only if the new method signature annotation is last in the sequence. * Modifying or removing an existing method signature annotation is a breaking change. * Re-ordering existing method signature annotations is a breaking change. |
| default_host | string | .google.protobuf.ServiceOptions | 1049 | The hostname for this service. This should be specified with no prefix or protocol.

Example:

 service Foo { option (google.api.default_host) = &#34;foo.googleapi.com&#34;; ... } |
| oauth_scopes | string | .google.protobuf.ServiceOptions | 1050 | OAuth scopes needed for the client.

Example:

 service Foo { option (google.api.oauth_scopes) = \ &#34;https://www.googleapis.com/auth/cloud-platform&#34;; ... }

If there is more than one scope, use a comma-separated string:

Example:

 service Foo { option (google.api.oauth_scopes) = \ &#34;https://www.googleapis.com/auth/cloud-platform,&#34; &#34;https://www.googleapis.com/auth/monitoring&#34;; ... } |

 

 



<a name="google_api_field_behavior-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/field_behavior.proto


 


<a name="google-api-FieldBehavior"></a>

### FieldBehavior
An indicator of the behavior of a given field (for example, that a field
is required in requests, or given as output but ignored as input).
This **does not** change the behavior in protocol buffers itself; it only
denotes the behavior and may affect how API tooling handles the field.

Note: This enum **may** receive new values in the future.

| Name | Number | Description |
| ---- | ------ | ----------- |
| FIELD_BEHAVIOR_UNSPECIFIED | 0 | Conventional default for enums. Do not use this. |
| OPTIONAL | 1 | Specifically denotes a field as optional. While all fields in protocol buffers are optional, this may be specified for emphasis if appropriate. |
| REQUIRED | 2 | Denotes a field as required. This indicates that the field **must** be provided as part of the request, and failure to do so will cause an error (usually `INVALID_ARGUMENT`). |
| OUTPUT_ONLY | 3 | Denotes a field as output only. This indicates that the field is provided in responses, but including the field in a request does nothing (the server *must* ignore it and *must not* throw an error as a result of the field&#39;s presence). |
| INPUT_ONLY | 4 | Denotes a field as input only. This indicates that the field is provided in requests, and the corresponding field is not included in output. |
| IMMUTABLE | 5 | Denotes a field as immutable. This indicates that the field may be set once in a request to create a resource, but may not be changed thereafter. |
| UNORDERED_LIST | 6 | Denotes that a (repeated) field is an unordered list. This indicates that the service may provide the elements of the list in any arbitrary order, rather than the order the user originally provided. Additionally, the list&#39;s order may or may not be stable. |
| NON_EMPTY_DEFAULT | 7 | Denotes that this field returns a non-empty default value if not set. This indicates that if the user provides the empty value in a request, a non-empty value will be returned. The user will not be aware of what non-empty value to expect. |
| IDENTIFIER | 8 | Denotes that the field in a resource (a message annotated with google.api.resource) is used in the resource name to uniquely identify the resource. For AIP-compliant APIs, this should only be applied to the `name` field on the resource.

This behavior should not be applied to references to other resources within the message.

The identifier field of resources often have different field behavior depending on the request it is embedded in (e.g. for Create methods name is optional and unused, while for Update methods it is required). Instead of method-specific annotations, only `IDENTIFIER` is required. |


 


<a name="google_api_field_behavior-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| field_behavior | FieldBehavior | .google.protobuf.FieldOptions | 1052 | A designation of a specific field behavior (required, output only, etc.) in protobuf messages.

Examples:

 string name = 1 [(google.api.field_behavior) = REQUIRED]; State state = 1 [(google.api.field_behavior) = OUTPUT_ONLY]; google.protobuf.Duration ttl = 1 [(google.api.field_behavior) = INPUT_ONLY]; google.protobuf.Timestamp expire_time = 1 [(google.api.field_behavior) = OUTPUT_ONLY, (google.api.field_behavior) = IMMUTABLE]; |

 

 



<a name="google_api_auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/auth.proto



<a name="google-api-AuthProvider"></a>

### AuthProvider
Configuration for an authentication provider, including support for
[JSON Web Token
(JWT)](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | The unique identifier of the auth provider. It will be referred to by `AuthRequirement.provider_id`.

Example: &#34;bookstore_auth&#34;. |
| issuer | [string](#string) |  | Identifies the principal that issued the JWT. See https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.1 Usually a URL or an email address.

Example: https://securetoken.google.com Example: 1234567-compute@developer.gserviceaccount.com |
| jwks_uri | [string](#string) |  | URL of the provider&#39;s public key set to validate signature of the JWT. See [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata). Optional if the key set document: - can be retrieved from [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html) of the issuer. - can be inferred from the email domain of the issuer (e.g. a Google service account).

Example: https://www.googleapis.com/oauth2/v1/certs |
| audiences | [string](#string) |  | The list of JWT [audiences](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.3). that are allowed to access. A JWT containing any of these audiences will be accepted. When this setting is absent, JWTs with audiences: - &#34;https://[service.name]/[google.protobuf.Api.name]&#34; - &#34;https://[service.name]/&#34; will be accepted. For example, if no audiences are in the setting, LibraryService API will accept JWTs with the following audiences: - https://library-example.googleapis.com/google.example.library.v1.LibraryService - https://library-example.googleapis.com/

Example:

 audiences: bookstore_android.apps.googleusercontent.com, bookstore_web.apps.googleusercontent.com |
| authorization_url | [string](#string) |  | Redirect URL if JWT token is required but not present or is expired. Implement authorizationUrl of securityDefinitions in OpenAPI spec. |
| jwt_locations | [JwtLocation](#google-api-JwtLocation) | repeated | Defines the locations to extract the JWT. For now it is only used by the Cloud Endpoints to store the OpenAPI extension [x-google-jwt-locations] (https://cloud.google.com/endpoints/docs/openapi/openapi-extensions#x-google-jwt-locations)

JWT locations can be one of HTTP headers, URL query parameters or cookies. The rule is that the first match wins.

If not specified, default to use following 3 locations: 1) Authorization: Bearer 2) x-goog-iap-jwt-assertion 3) access_token query parameter

Default locations can be specified as followings: jwt_locations: - header: Authorization value_prefix: &#34;Bearer &#34; - header: x-goog-iap-jwt-assertion - query: access_token |






<a name="google-api-AuthRequirement"></a>

### AuthRequirement
User-defined authentication requirements, including support for
[JSON Web Token
(JWT)](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider_id | [string](#string) |  | [id][google.api.AuthProvider.id] from authentication provider.

Example:

 provider_id: bookstore_auth |
| audiences | [string](#string) |  | NOTE: This will be deprecated soon, once AuthProvider.audiences is implemented and accepted in all the runtime components.

The list of JWT [audiences](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.3). that are allowed to access. A JWT containing any of these audiences will be accepted. When this setting is absent, only JWTs with audience &#34;https://[Service_name][google.api.Service.name]/[API_name][google.protobuf.Api.name]&#34; will be accepted. For example, if no audiences are in the setting, LibraryService API will only accept JWTs with the following audience &#34;https://library-example.googleapis.com/google.example.library.v1.LibraryService&#34;.

Example:

 audiences: bookstore_android.apps.googleusercontent.com, bookstore_web.apps.googleusercontent.com |






<a name="google-api-Authentication"></a>

### Authentication
`Authentication` defines the authentication configuration for API methods
provided by an API service.

Example:

    name: calendar.googleapis.com
    authentication:
      providers:
      - id: google_calendar_auth
        jwks_uri: https://www.googleapis.com/oauth2/v1/certs
        issuer: https://securetoken.google.com
      rules:
      - selector: &#34;*&#34;
        requirements:
          provider_id: google_calendar_auth
      - selector: google.calendar.Delegate
        oauth:
          canonical_scopes: https://www.googleapis.com/auth/calendar.read


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rules | [AuthenticationRule](#google-api-AuthenticationRule) | repeated | A list of authentication rules that apply to individual API methods.

**NOTE:** All service configuration rules follow &#34;last one wins&#34; order. |
| providers | [AuthProvider](#google-api-AuthProvider) | repeated | Defines a set of authentication providers that a service supports. |






<a name="google-api-AuthenticationRule"></a>

### AuthenticationRule
Authentication rules for the service.

By default, if a method has any authentication requirements, every request
must include a valid credential matching one of the requirements.
It&#39;s an error to include more than one kind of credential in a single
request.

If a method doesn&#39;t have any auth requirements, request credentials will be
ignored.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | Selects the methods to which this rule applies.

Refer to [selector][google.api.DocumentationRule.selector] for syntax details. |
| oauth | [OAuthRequirements](#google-api-OAuthRequirements) |  | The requirements for OAuth credentials. |
| allow_without_credential | [bool](#bool) |  | If true, the service accepts API keys without any other credential. This flag only applies to HTTP and gRPC requests. |
| requirements | [AuthRequirement](#google-api-AuthRequirement) | repeated | Requirements for additional authentication providers. |






<a name="google-api-JwtLocation"></a>

### JwtLocation
Specifies a location to extract JWT from an API request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [string](#string) |  | Specifies HTTP header name to extract JWT token. |
| query | [string](#string) |  | Specifies URL query parameter name to extract JWT token. |
| cookie | [string](#string) |  | Specifies cookie name to extract JWT token. |
| value_prefix | [string](#string) |  | The value prefix. The value format is &#34;value_prefix{token}&#34; Only applies to &#34;in&#34; header type. Must be empty for &#34;in&#34; query type. If not empty, the header value has to match (case sensitive) this prefix. If not matched, JWT will not be extracted. If matched, JWT will be extracted after the prefix is removed.

For example, for &#34;Authorization: Bearer {JWT}&#34;, value_prefix=&#34;Bearer &#34; with a space at the end. |






<a name="google-api-OAuthRequirements"></a>

### OAuthRequirements
OAuth scopes are a way to define data and permissions on data. For example,
there are scopes defined for &#34;Read-only access to Google Calendar&#34; and
&#34;Access to Cloud Platform&#34;. Users can consent to a scope for an application,
giving it permission to access that data on their behalf.

OAuth scope specifications should be fairly coarse grained; a user will need
to see and understand the text description of what your scope means.

In most cases: use one or at most two OAuth scopes for an entire family of
products. If your product has multiple APIs, you should probably be sharing
the OAuth scope across all of those APIs.

When you need finer grained OAuth consent screens: talk with your product
management about how developers will use them in practice.

Please note that even though each of the canonical scopes is enough for a
request to be accepted and passed to the backend, a request can still fail
due to the backend requiring additional scopes or permissions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| canonical_scopes | [string](#string) |  | The list of publicly documented OAuth scopes that are allowed access. An OAuth token containing any of these scopes will be accepted.

Example:

 canonical_scopes: https://www.googleapis.com/auth/calendar, https://www.googleapis.com/auth/calendar.read |





 

 

 

 



<a name="google_api_backend-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/backend.proto



<a name="google-api-Backend"></a>

### Backend
`Backend` defines the backend configuration for a service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rules | [BackendRule](#google-api-BackendRule) | repeated | A list of API backend rules that apply to individual API methods.

**NOTE:** All service configuration rules follow &#34;last one wins&#34; order. |






<a name="google-api-BackendRule"></a>

### BackendRule
A backend rule provides configuration for an individual API element.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | Selects the methods to which this rule applies.

Refer to [selector][google.api.DocumentationRule.selector] for syntax details. |
| address | [string](#string) |  | The address of the API backend.

The scheme is used to determine the backend protocol and security. The following schemes are accepted:

 SCHEME PROTOCOL SECURITY http:// HTTP None https:// HTTP TLS grpc:// gRPC None grpcs:// gRPC TLS

It is recommended to explicitly include a scheme. Leaving out the scheme may cause constrasting behaviors across platforms.

If the port is unspecified, the default is: - 80 for schemes without TLS - 443 for schemes with TLS

For HTTP backends, use [protocol][google.api.BackendRule.protocol] to specify the protocol version. |
| deadline | [double](#double) |  | The number of seconds to wait for a response from a request. The default varies based on the request protocol and deployment environment. |
| min_deadline | [double](#double) |  | **Deprecated.** Deprecated, do not use. |
| operation_deadline | [double](#double) |  | The number of seconds to wait for the completion of a long running operation. The default is no deadline. |
| path_translation | [BackendRule.PathTranslation](#google-api-BackendRule-PathTranslation) |  |  |
| jwt_audience | [string](#string) |  | The JWT audience is used when generating a JWT ID token for the backend. This ID token will be added in the HTTP &#34;authorization&#34; header, and sent to the backend. |
| disable_auth | [bool](#bool) |  | When disable_auth is true, a JWT ID token won&#39;t be generated and the original &#34;Authorization&#34; HTTP header will be preserved. If the header is used to carry the original token and is expected by the backend, this field must be set to true to preserve the header. |
| protocol | [string](#string) |  | The protocol used for sending a request to the backend. The supported values are &#34;http/1.1&#34; and &#34;h2&#34;.

The default value is inferred from the scheme in the [address][google.api.BackendRule.address] field:

 SCHEME PROTOCOL http:// http/1.1 https:// http/1.1 grpc:// h2 grpcs:// h2

For secure HTTP backends (https://) that support HTTP/2, set this field to &#34;h2&#34; for improved performance.

Configuring this field to non-default values is only supported for secure HTTP backends. This field will be ignored for all other backends.

See https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids for more details on the supported values. |
| overrides_by_request_protocol | [BackendRule.OverridesByRequestProtocolEntry](#google-api-BackendRule-OverridesByRequestProtocolEntry) | repeated | The map between request protocol and the backend address. |






<a name="google-api-BackendRule-OverridesByRequestProtocolEntry"></a>

### BackendRule.OverridesByRequestProtocolEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [BackendRule](#google-api-BackendRule) |  |  |





 


<a name="google-api-BackendRule-PathTranslation"></a>

### BackendRule.PathTranslation
Path Translation specifies how to combine the backend address with the
request path in order to produce the appropriate forwarding URL for the
request.

Path Translation is applicable only to HTTP-based backends. Backends which
do not accept requests over HTTP/HTTPS should leave `path_translation`
unspecified.

| Name | Number | Description |
| ---- | ------ | ----------- |
| PATH_TRANSLATION_UNSPECIFIED | 0 |  |
| CONSTANT_ADDRESS | 1 | Use the backend address as-is, with no modification to the path. If the URL pattern contains variables, the variable names and values will be appended to the query string. If a query string parameter and a URL pattern variable have the same name, this may result in duplicate keys in the query string.

# Examples

Given the following operation config:

 Method path: /api/company/{cid}/user/{uid} Backend address: https://example.cloudfunctions.net/getUser

Requests to the following request paths will call the backend at the translated path:

 Request path: /api/company/widgetworks/user/johndoe Translated: https://example.cloudfunctions.net/getUser?cid=widgetworks&amp;uid=johndoe

 Request path: /api/company/widgetworks/user/johndoe?timezone=EST Translated: https://example.cloudfunctions.net/getUser?timezone=EST&amp;cid=widgetworks&amp;uid=johndoe |
| APPEND_PATH_TO_ADDRESS | 2 | The request path will be appended to the backend address.

# Examples

Given the following operation config:

 Method path: /api/company/{cid}/user/{uid} Backend address: https://example.appspot.com

Requests to the following request paths will call the backend at the translated path:

 Request path: /api/company/widgetworks/user/johndoe Translated: https://example.appspot.com/api/company/widgetworks/user/johndoe

 Request path: /api/company/widgetworks/user/johndoe?timezone=EST Translated: https://example.appspot.com/api/company/widgetworks/user/johndoe?timezone=EST |


 

 

 



<a name="google_api_billing-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/billing.proto



<a name="google-api-Billing"></a>

### Billing
Billing related configuration of the service.

The following example shows how to configure monitored resources and metrics
for billing, `consumer_destinations` is the only supported destination and
the monitored resources need at least one label key
`cloud.googleapis.com/location` to indicate the location of the billing
usage, using different monitored resources between monitoring and billing is
recommended so they can be evolved independently:


    monitored_resources:
    - type: library.googleapis.com/billing_branch
      labels:
      - key: cloud.googleapis.com/location
        description: |
          Predefined label to support billing location restriction.
      - key: city
        description: |
          Custom label to define the city where the library branch is located
          in.
      - key: name
        description: Custom label to define the name of the library branch.
    metrics:
    - name: library.googleapis.com/book/borrowed_count
      metric_kind: DELTA
      value_type: INT64
      unit: &#34;1&#34;
    billing:
      consumer_destinations:
      - monitored_resource: library.googleapis.com/billing_branch
        metrics:
        - library.googleapis.com/book/borrowed_count


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| consumer_destinations | [Billing.BillingDestination](#google-api-Billing-BillingDestination) | repeated | Billing configurations for sending metrics to the consumer project. There can be multiple consumer destinations per service, each one must have a different monitored resource type. A metric can be used in at most one consumer destination. |






<a name="google-api-Billing-BillingDestination"></a>

### Billing.BillingDestination
Configuration of a specific billing destination (Currently only support
bill against consumer project).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| monitored_resource | [string](#string) |  | The monitored resource type. The type must be defined in [Service.monitored_resources][google.api.Service.monitored_resources] section. |
| metrics | [string](#string) | repeated | Names of the metrics to report to this billing destination. Each name must be defined in [Service.metrics][google.api.Service.metrics] section. |





 

 

 

 



<a name="google_api_config_change-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/config_change.proto



<a name="google-api-Advice"></a>

### Advice
Generated advice about this change, used for providing more
information about how a change will affect the existing service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  | Useful description for why this advice was applied and what actions should be taken to mitigate any implied risks. |






<a name="google-api-ConfigChange"></a>

### ConfigChange
Output generated from semantically comparing two versions of a service
configuration.

Includes detailed information about a field that have changed with
applicable advice about potential consequences for the change, such as
backwards-incompatibility.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| element | [string](#string) |  | Object hierarchy path to the change, with levels separated by a &#39;.&#39; character. For repeated fields, an applicable unique identifier field is used for the index (usually selector, name, or id). For maps, the term &#39;key&#39; is used. If the field has no unique identifier, the numeric index is used. Examples: - visibility.rules[selector==&#34;google.LibraryService.ListBooks&#34;].restriction - quota.metric_rules[selector==&#34;google&#34;].metric_costs[key==&#34;reads&#34;].value - logging.producer_destinations[0] |
| old_value | [string](#string) |  | Value of the changed object in the old Service configuration, in JSON format. This field will not be populated if ChangeType == ADDED. |
| new_value | [string](#string) |  | Value of the changed object in the new Service configuration, in JSON format. This field will not be populated if ChangeType == REMOVED. |
| change_type | [ChangeType](#google-api-ChangeType) |  | The type for this change, either ADDED, REMOVED, or MODIFIED. |
| advices | [Advice](#google-api-Advice) | repeated | Collection of advice provided for this change, useful for determining the possible impact of this change. |





 


<a name="google-api-ChangeType"></a>

### ChangeType
Classifies set of possible modifications to an object in the service
configuration.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CHANGE_TYPE_UNSPECIFIED | 0 | No value was provided. |
| ADDED | 1 | The changed object exists in the &#39;new&#39; service configuration, but not in the &#39;old&#39; service configuration. |
| REMOVED | 2 | The changed object exists in the &#39;old&#39; service configuration, but not in the &#39;new&#39; service configuration. |
| MODIFIED | 3 | The changed object exists in both service configurations, but its value is different. |


 

 

 



<a name="google_api_consumer-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/consumer.proto



<a name="google-api-ProjectProperties"></a>

### ProjectProperties
A descriptor for defining project properties for a service. One service may
have many consumer projects, and the service may want to behave differently
depending on some properties on the project. For example, a project may be
associated with a school, or a business, or a government agency, a business
type property on the project may affect how a service responds to the client.
This descriptor defines which properties are allowed to be set on a project.

Example:

   project_properties:
     properties:
     - name: NO_WATERMARK
       type: BOOL
       description: Allows usage of the API without watermarks.
     - name: EXTENDED_TILE_CACHE_PERIOD
       type: INT64


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| properties | [Property](#google-api-Property) | repeated | List of per consumer project-specific properties. |






<a name="google-api-Property"></a>

### Property
Defines project properties.

API services can define properties that can be assigned to consumer projects
so that backends can perform response customization without having to make
additional calls or maintain additional storage. For example, Maps API
defines properties that controls map tile cache period, or whether to embed a
watermark in a result.

These values can be set via API producer console. Only API providers can
define and set these properties.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the property (a.k.a key). |
| type | [Property.PropertyType](#google-api-Property-PropertyType) |  | The type of this property. |
| description | [string](#string) |  | The description of the property |





 


<a name="google-api-Property-PropertyType"></a>

### Property.PropertyType
Supported data type of the property values

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 | The type is unspecified, and will result in an error. |
| INT64 | 1 | The type is `int64`. |
| BOOL | 2 | The type is `bool`. |
| STRING | 3 | The type is `string`. |
| DOUBLE | 4 | The type is &#39;double&#39;. |


 

 

 



<a name="google_api_context-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/context.proto



<a name="google-api-Context"></a>

### Context
`Context` defines which contexts an API requests.

Example:

    context:
      rules:
      - selector: &#34;*&#34;
        requested:
        - google.rpc.context.ProjectContext
        - google.rpc.context.OriginContext

The above specifies that all methods in the API request
`google.rpc.context.ProjectContext` and
`google.rpc.context.OriginContext`.

Available context types are defined in package
`google.rpc.context`.

This also provides mechanism to allowlist any protobuf message extension that
can be sent in grpc metadata using “x-goog-ext-&lt;extension_id&gt;-bin” and
“x-goog-ext-&lt;extension_id&gt;-jspb” format. For example, list any service
specific protobuf types that can appear in grpc metadata as follows in your
yaml file:

Example:

    context:
      rules:
       - selector: &#34;google.example.library.v1.LibraryService.CreateBook&#34;
         allowed_request_extensions:
         - google.foo.v1.NewExtension
         allowed_response_extensions:
         - google.foo.v1.NewExtension

You can also specify extension ID instead of fully qualified extension name
here.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rules | [ContextRule](#google-api-ContextRule) | repeated | A list of RPC context rules that apply to individual API methods.

**NOTE:** All service configuration rules follow &#34;last one wins&#34; order. |






<a name="google-api-ContextRule"></a>

### ContextRule
A context rule provides information about the context for an individual API
element.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | Selects the methods to which this rule applies.

Refer to [selector][google.api.DocumentationRule.selector] for syntax details. |
| requested | [string](#string) | repeated | A list of full type names of requested contexts. |
| provided | [string](#string) | repeated | A list of full type names of provided contexts. |
| allowed_request_extensions | [string](#string) | repeated | A list of full type names or extension IDs of extensions allowed in grpc side channel from client to backend. |
| allowed_response_extensions | [string](#string) | repeated | A list of full type names or extension IDs of extensions allowed in grpc side channel from backend to client. |





 

 

 

 



<a name="google_api_policy-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/policy.proto



<a name="google-api-FieldPolicy"></a>

### FieldPolicy
Google API Policy Annotation

This message defines a simple API policy annotation that can be used to
annotate API request and response message fields with applicable policies.
One field may have multiple applicable policies that must all be satisfied
before a request can be processed. This policy annotation is used to
generate the overall policy that will be used for automatic runtime
policy enforcement and documentation generation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | Selects one or more request or response message fields to apply this `FieldPolicy`.

When a `FieldPolicy` is used in proto annotation, the selector must be left as empty. The service config generator will automatically fill the correct value.

When a `FieldPolicy` is used in service config, the selector must be a comma-separated string with valid request or response field paths, such as &#34;foo.bar&#34; or &#34;foo.bar,foo.baz&#34;. |
| resource_permission | [string](#string) |  | Specifies the required permission(s) for the resource referred to by the field. It requires the field contains a valid resource reference, and the request must pass the permission checks to proceed. For example, &#34;resourcemanager.projects.get&#34;. |
| resource_type | [string](#string) |  | Specifies the resource type for the resource referred to by the field. |






<a name="google-api-MethodPolicy"></a>

### MethodPolicy
Defines policies applying to an RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | Selects a method to which these policies should be enforced, for example, &#34;google.pubsub.v1.Subscriber.CreateSubscription&#34;.

Refer to [selector][google.api.DocumentationRule.selector] for syntax details.

NOTE: This field must not be set in the proto annotation. It will be automatically filled by the service config compiler . |
| request_policies | [FieldPolicy](#google-api-FieldPolicy) | repeated | Policies that are applicable to the request message. |





 

 


<a name="google_api_policy-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| field_policy | FieldPolicy | .google.protobuf.FieldOptions | 158361448 | See [FieldPolicy][]. |
| method_policy | MethodPolicy | .google.protobuf.MethodOptions | 161893301 | See [MethodPolicy][]. |

 

 



<a name="google_api_control-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/control.proto



<a name="google-api-Control"></a>

### Control
Selects and configures the service controller used by the service.

Example:

    control:
      environment: servicecontrol.googleapis.com


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| environment | [string](#string) |  | The service controller environment to use. If empty, no control plane feature (like quota and billing) will be enabled. The recommended value for most services is servicecontrol.googleapis.com |
| method_policies | [MethodPolicy](#google-api-MethodPolicy) | repeated | Defines policies applying to the API methods of the service. |





 

 

 

 



<a name="google_api_distribution-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/distribution.proto



<a name="google-api-Distribution"></a>

### Distribution
`Distribution` contains summary statistics for a population of values. It
optionally contains a histogram representing the distribution of those values
across a set of buckets.

The summary statistics are the count, mean, sum of the squared deviation from
the mean, the minimum, and the maximum of the set of population of values.
The histogram is based on a sequence of buckets and gives a count of values
that fall into each bucket. The boundaries of the buckets are given either
explicitly or by formulas for buckets of fixed or exponentially increasing
widths.

Although it is not forbidden, it is generally a bad idea to include
non-finite values (infinities or NaNs) in the population of values, as this
will render the `mean` and `sum_of_squared_deviation` fields meaningless.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | The number of values in the population. Must be non-negative. This value must equal the sum of the values in `bucket_counts` if a histogram is provided. |
| mean | [double](#double) |  | The arithmetic mean of the values in the population. If `count` is zero then this field must be zero. |
| sum_of_squared_deviation | [double](#double) |  | The sum of squared deviations from the mean of the values in the population. For values x_i this is:

 Sum[i=1..n]((x_i - mean)^2)

Knuth, &#34;The Art of Computer Programming&#34;, Vol. 2, page 232, 3rd edition describes Welford&#39;s method for accumulating this sum in one pass.

If `count` is zero then this field must be zero. |
| range | [Distribution.Range](#google-api-Distribution-Range) |  | If specified, contains the range of the population values. The field must not be present if the `count` is zero. |
| bucket_options | [Distribution.BucketOptions](#google-api-Distribution-BucketOptions) |  | Defines the histogram bucket boundaries. If the distribution does not contain a histogram, then omit this field. |
| bucket_counts | [int64](#int64) | repeated | The number of values in each bucket of the histogram, as described in `bucket_options`. If the distribution does not have a histogram, then omit this field. If there is a histogram, then the sum of the values in `bucket_counts` must equal the value in the `count` field of the distribution.

If present, `bucket_counts` should contain N values, where N is the number of buckets specified in `bucket_options`. If you supply fewer than N values, the remaining values are assumed to be 0.

The order of the values in `bucket_counts` follows the bucket numbering schemes described for the three bucket types. The first value must be the count for the underflow bucket (number 0). The next N-2 values are the counts for the finite buckets (number 1 through N-2). The N&#39;th value in `bucket_counts` is the count for the overflow bucket (number N-1). |
| exemplars | [Distribution.Exemplar](#google-api-Distribution-Exemplar) | repeated | Must be in increasing order of `value` field. |






<a name="google-api-Distribution-BucketOptions"></a>

### Distribution.BucketOptions
`BucketOptions` describes the bucket boundaries used to create a histogram
for the distribution. The buckets can be in a linear sequence, an
exponential sequence, or each bucket can be specified explicitly.
`BucketOptions` does not include the number of values in each bucket.

A bucket has an inclusive lower bound and exclusive upper bound for the
values that are counted for that bucket. The upper bound of a bucket must
be strictly greater than the lower bound. The sequence of N buckets for a
distribution consists of an underflow bucket (number 0), zero or more
finite buckets (number 1 through N - 2) and an overflow bucket (number N -
1). The buckets are contiguous: the lower bound of bucket i (i &gt; 0) is the
same as the upper bound of bucket i - 1. The buckets span the whole range
of finite values: lower bound of the underflow bucket is -infinity and the
upper bound of the overflow bucket is &#43;infinity. The finite buckets are
so-called because both bounds are finite.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| linear_buckets | [Distribution.BucketOptions.Linear](#google-api-Distribution-BucketOptions-Linear) |  | The linear bucket. |
| exponential_buckets | [Distribution.BucketOptions.Exponential](#google-api-Distribution-BucketOptions-Exponential) |  | The exponential buckets. |
| explicit_buckets | [Distribution.BucketOptions.Explicit](#google-api-Distribution-BucketOptions-Explicit) |  | The explicit buckets. |






<a name="google-api-Distribution-BucketOptions-Explicit"></a>

### Distribution.BucketOptions.Explicit
Specifies a set of buckets with arbitrary widths.

There are `size(bounds) &#43; 1` (= N) buckets. Bucket `i` has the following
boundaries:

   Upper bound (0 &lt;= i &lt; N-1):     bounds[i]
   Lower bound (1 &lt;= i &lt; N);       bounds[i - 1]

The `bounds` field must contain at least one element. If `bounds` has
only one element, then there are no finite buckets, and that single
element is the common boundary of the overflow and underflow buckets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bounds | [double](#double) | repeated | The values must be monotonically increasing. |






<a name="google-api-Distribution-BucketOptions-Exponential"></a>

### Distribution.BucketOptions.Exponential
Specifies an exponential sequence of buckets that have a width that is
proportional to the value of the lower bound. Each bucket represents a
constant relative uncertainty on a specific value in the bucket.

There are `num_finite_buckets &#43; 2` (= N) buckets. Bucket `i` has the
following boundaries:

   Upper bound (0 &lt;= i &lt; N-1):     scale * (growth_factor ^ i).

   Lower bound (1 &lt;= i &lt; N):       scale * (growth_factor ^ (i - 1)).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num_finite_buckets | [int32](#int32) |  | Must be greater than 0. |
| growth_factor | [double](#double) |  | Must be greater than 1. |
| scale | [double](#double) |  | Must be greater than 0. |






<a name="google-api-Distribution-BucketOptions-Linear"></a>

### Distribution.BucketOptions.Linear
Specifies a linear sequence of buckets that all have the same width
(except overflow and underflow). Each bucket represents a constant
absolute uncertainty on the specific value in the bucket.

There are `num_finite_buckets &#43; 2` (= N) buckets. Bucket `i` has the
following boundaries:

   Upper bound (0 &lt;= i &lt; N-1):     offset &#43; (width * i).

   Lower bound (1 &lt;= i &lt; N):       offset &#43; (width * (i - 1)).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num_finite_buckets | [int32](#int32) |  | Must be greater than 0. |
| width | [double](#double) |  | Must be greater than 0. |
| offset | [double](#double) |  | Lower bound of the first bucket. |






<a name="google-api-Distribution-Exemplar"></a>

### Distribution.Exemplar
Exemplars are example points that may be used to annotate aggregated
distribution values. They are metadata that gives information about a
particular value added to a Distribution bucket, such as a trace ID that
was active when a value was added. They may contain further information,
such as a example values and timestamps, origin, etc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  | Value of the exemplar point. This value determines to which bucket the exemplar belongs. |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The observation (sampling) time of the above value. |
| attachments | [google.protobuf.Any](#google-protobuf-Any) | repeated | Contextual information about the example value. Examples are:

 Trace: type.googleapis.com/google.monitoring.v3.SpanContext

 Literal string: type.googleapis.com/google.protobuf.StringValue

 Labels dropped during aggregation: type.googleapis.com/google.monitoring.v3.DroppedLabels

There may be only a single attachment of any given message type in a single exemplar, and this is enforced by the system. |






<a name="google-api-Distribution-Range"></a>

### Distribution.Range
The range of the population values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min | [double](#double) |  | The minimum of the population values. |
| max | [double](#double) |  | The maximum of the population values. |





 

 

 

 



<a name="google_api_documentation-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/documentation.proto



<a name="google-api-Documentation"></a>

### Documentation
`Documentation` provides the information for describing a service.

Example:
&lt;pre&gt;&lt;code&gt;documentation:
  summary: &gt;
    The Google Calendar API gives access
    to most calendar features.
  pages:
  - name: Overview
    content: &amp;#40;== include google/foo/overview.md ==&amp;#41;
  - name: Tutorial
    content: &amp;#40;== include google/foo/tutorial.md ==&amp;#41;
    subpages;
    - name: Java
      content: &amp;#40;== include google/foo/tutorial_java.md ==&amp;#41;
  rules:
  - selector: google.calendar.Calendar.Get
    description: &gt;
      ...
  - selector: google.calendar.Calendar.Put
    description: &gt;
      ...
&lt;/code&gt;&lt;/pre&gt;
Documentation is provided in markdown syntax. In addition to
standard markdown features, definition lists, tables and fenced
code blocks are supported. Section headers can be provided and are
interpreted relative to the section nesting of the context where
a documentation fragment is embedded.

Documentation from the IDL is merged with documentation defined
via the config at normalization time, where documentation provided
by config rules overrides IDL provided.

A number of constructs specific to the API platform are supported
in documentation text.

In order to reference a proto element, the following
notation can be used:
&lt;pre&gt;&lt;code&gt;&amp;#91;fully.qualified.proto.name]&amp;#91;]&lt;/code&gt;&lt;/pre&gt;
To override the display text used for the link, this can be used:
&lt;pre&gt;&lt;code&gt;&amp;#91;display text]&amp;#91;fully.qualified.proto.name]&lt;/code&gt;&lt;/pre&gt;
Text can be excluded from doc using the following notation:
&lt;pre&gt;&lt;code&gt;&amp;#40;-- internal comment --&amp;#41;&lt;/code&gt;&lt;/pre&gt;

A few directives are available in documentation. Note that
directives must appear on a single line to be properly
identified. The `include` directive includes a markdown file from
an external source:
&lt;pre&gt;&lt;code&gt;&amp;#40;== include path/to/file ==&amp;#41;&lt;/code&gt;&lt;/pre&gt;
The `resource_for` directive marks a message to be the resource of
a collection in REST view. If it is not specified, tools attempt
to infer the resource from the operations in a collection:
&lt;pre&gt;&lt;code&gt;&amp;#40;== resource_for v1.shelves.books ==&amp;#41;&lt;/code&gt;&lt;/pre&gt;
The directive `suppress_warning` does not directly affect documentation
and is documented together with service config validation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| summary | [string](#string) |  | A short description of what the service does. The summary must be plain text. It becomes the overview of the service displayed in Google Cloud Console. NOTE: This field is equivalent to the standard field `description`. |
| pages | [Page](#google-api-Page) | repeated | The top level pages for the documentation set. |
| rules | [DocumentationRule](#google-api-DocumentationRule) | repeated | A list of documentation rules that apply to individual API elements.

**NOTE:** All service configuration rules follow &#34;last one wins&#34; order. |
| documentation_root_url | [string](#string) |  | The URL to the root of documentation. |
| service_root_url | [string](#string) |  | Specifies the service root url if the default one (the service name from the yaml file) is not suitable. This can be seen in any fully specified service urls as well as sections that show a base that other urls are relative to. |
| overview | [string](#string) |  | Declares a single overview page. For example: &lt;pre&gt;&lt;code&gt;documentation: summary: ... overview: &amp;#40;== include overview.md ==&amp;#41; &lt;/code&gt;&lt;/pre&gt; This is a shortcut for the following declaration (using pages style): &lt;pre&gt;&lt;code&gt;documentation: summary: ... pages: - name: Overview content: &amp;#40;== include overview.md ==&amp;#41; &lt;/code&gt;&lt;/pre&gt; Note: you cannot specify both `overview` field and `pages` field. |






<a name="google-api-DocumentationRule"></a>

### DocumentationRule
A documentation rule provides information about individual API elements.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | The selector is a comma-separated list of patterns for any element such as a method, a field, an enum value. Each pattern is a qualified name of the element which may end in &#34;*&#34;, indicating a wildcard. Wildcards are only allowed at the end and for a whole component of the qualified name, i.e. &#34;foo.*&#34; is ok, but not &#34;foo.b*&#34; or &#34;foo.*.bar&#34;. A wildcard will match one or more components. To specify a default for all applicable elements, the whole pattern &#34;*&#34; is used. |
| description | [string](#string) |  | Description of the selected proto element (e.g. a message, a method, a &#39;service&#39; definition, or a field). Defaults to leading &amp; trailing comments taken from the proto source definition of the proto element. |
| deprecation_description | [string](#string) |  | Deprecation description of the selected element(s). It can be provided if an element is marked as `deprecated`. |






<a name="google-api-Page"></a>

### Page
Represents a documentation page. A page can contain subpages to represent
nested documentation set structure.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the page. It will be used as an identity of the page to generate URI of the page, text of the link to this page in navigation, etc. The full page name (start from the root page name to this page concatenated with `.`) can be used as reference to the page in your documentation. For example: &lt;pre&gt;&lt;code&gt;pages: - name: Tutorial content: &amp;#40;== include tutorial.md ==&amp;#41; subpages: - name: Java content: &amp;#40;== include tutorial_java.md ==&amp;#41; &lt;/code&gt;&lt;/pre&gt; You can reference `Java` page using Markdown reference link syntax: `[Java][Tutorial.Java]`. |
| content | [string](#string) |  | The Markdown content of the page. You can use &lt;code&gt;&amp;#40;== include {path} ==&amp;#41;&lt;/code&gt; to include content from a Markdown file. The content can be used to produce the documentation page such as HTML format page. |
| subpages | [Page](#google-api-Page) | repeated | Subpages of this page. The order of subpages specified here will be honored in the generated docset. |





 

 

 

 



<a name="google_api_endpoint-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/endpoint.proto



<a name="google-api-Endpoint"></a>

### Endpoint
`Endpoint` describes a network address of a service that serves a set of
APIs. It is commonly known as a service endpoint. A service may expose
any number of service endpoints, and all service endpoints share the same
service definition, such as quota limits and monitoring metrics.

Example:

    type: google.api.Service
    name: library-example.googleapis.com
    endpoints:
      # Declares network address `https://library-example.googleapis.com`
      # for service `library-example.googleapis.com`. The `https` scheme
      # is implicit for all service endpoints. Other schemes may be
      # supported in the future.
    - name: library-example.googleapis.com
      allow_cors: false
    - name: content-staging-library-example.googleapis.com
      # Allows HTTP OPTIONS calls to be passed to the API frontend, for it
      # to decide whether the subsequent cross-origin request is allowed
      # to proceed.
      allow_cors: true


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The canonical name of this endpoint. |
| aliases | [string](#string) | repeated | **Deprecated.** Unimplemented. Dot not use.

DEPRECATED: This field is no longer supported. Instead of using aliases, please specify multiple [google.api.Endpoint][google.api.Endpoint] for each of the intended aliases.

Additional names that this endpoint will be hosted on. |
| target | [string](#string) |  | The specification of an Internet routable address of API frontend that will handle requests to this [API Endpoint](https://cloud.google.com/apis/design/glossary). It should be either a valid IPv4 address or a fully-qualified domain name. For example, &#34;8.8.8.8&#34; or &#34;myservice.appspot.com&#34;. |
| allow_cors | [bool](#bool) |  | Allowing [CORS](https://en.wikipedia.org/wiki/Cross-origin_resource_sharing), aka cross-domain traffic, would allow the backends served from this endpoint to receive and respond to HTTP OPTIONS requests. The response will be used by the browser to determine whether the subsequent cross-origin request is allowed to proceed. |





 

 

 

 



<a name="google_api_error_reason-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/error_reason.proto


 


<a name="google-api-ErrorReason"></a>

### ErrorReason
Defines the supported values for `google.rpc.ErrorInfo.reason` for the
`googleapis.com` error domain. This error domain is reserved for [Service
Infrastructure](https://cloud.google.com/service-infrastructure/docs/overview).
For each error info of this domain, the metadata key &#34;service&#34; refers to the
logical identifier of an API service, such as &#34;pubsub.googleapis.com&#34;. The
&#34;consumer&#34; refers to the entity that consumes an API Service. It typically is
a Google project that owns the client application or the server resource,
such as &#34;projects/123&#34;. Other metadata keys are specific to each error
reason. For more information, see the definition of the specific error
reason.

| Name | Number | Description |
| ---- | ------ | ----------- |
| ERROR_REASON_UNSPECIFIED | 0 | Do not use this default value. |
| SERVICE_DISABLED | 1 | The request is calling a disabled service for a consumer.

Example of an ErrorInfo when the consumer &#34;projects/123&#34; contacting &#34;pubsub.googleapis.com&#34; service which is disabled:

 { &#34;reason&#34;: &#34;SERVICE_DISABLED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;pubsub.googleapis.com&#34; } }

This response indicates the &#34;pubsub.googleapis.com&#34; has been disabled in &#34;projects/123&#34;. |
| BILLING_DISABLED | 2 | The request whose associated billing account is disabled.

Example of an ErrorInfo when the consumer &#34;projects/123&#34; fails to contact &#34;pubsub.googleapis.com&#34; service because the associated billing account is disabled:

 { &#34;reason&#34;: &#34;BILLING_DISABLED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;pubsub.googleapis.com&#34; } }

This response indicates the billing account associated has been disabled. |
| API_KEY_INVALID | 3 | The request is denied because the provided [API key](https://cloud.google.com/docs/authentication/api-keys) is invalid. It may be in a bad format, cannot be found, or has been expired).

Example of an ErrorInfo when the request is contacting &#34;storage.googleapis.com&#34; service with an invalid API key:

 { &#34;reason&#34;: &#34;API_KEY_INVALID&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;service&#34;: &#34;storage.googleapis.com&#34;, } } |
| API_KEY_SERVICE_BLOCKED | 4 | The request is denied because it violates [API key API restrictions](https://cloud.google.com/docs/authentication/api-keys#adding_api_restrictions).

Example of an ErrorInfo when the consumer &#34;projects/123&#34; fails to call the &#34;storage.googleapis.com&#34; service because this service is restricted in the API key:

 { &#34;reason&#34;: &#34;API_KEY_SERVICE_BLOCKED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34; } } |
| API_KEY_HTTP_REFERRER_BLOCKED | 7 | The request is denied because it violates [API key HTTP restrictions](https://cloud.google.com/docs/authentication/api-keys#adding_http_restrictions).

Example of an ErrorInfo when the consumer &#34;projects/123&#34; fails to call &#34;storage.googleapis.com&#34; service because the http referrer of the request violates API key HTTP restrictions:

 { &#34;reason&#34;: &#34;API_KEY_HTTP_REFERRER_BLOCKED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34;, } } |
| API_KEY_IP_ADDRESS_BLOCKED | 8 | The request is denied because it violates [API key IP address restrictions](https://cloud.google.com/docs/authentication/api-keys#adding_application_restrictions).

Example of an ErrorInfo when the consumer &#34;projects/123&#34; fails to call &#34;storage.googleapis.com&#34; service because the caller IP of the request violates API key IP address restrictions:

 { &#34;reason&#34;: &#34;API_KEY_IP_ADDRESS_BLOCKED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34;, } } |
| API_KEY_ANDROID_APP_BLOCKED | 9 | The request is denied because it violates [API key Android application restrictions](https://cloud.google.com/docs/authentication/api-keys#adding_application_restrictions).

Example of an ErrorInfo when the consumer &#34;projects/123&#34; fails to call &#34;storage.googleapis.com&#34; service because the request from the Android apps violates the API key Android application restrictions:

 { &#34;reason&#34;: &#34;API_KEY_ANDROID_APP_BLOCKED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34; } } |
| API_KEY_IOS_APP_BLOCKED | 13 | The request is denied because it violates [API key iOS application restrictions](https://cloud.google.com/docs/authentication/api-keys#adding_application_restrictions).

Example of an ErrorInfo when the consumer &#34;projects/123&#34; fails to call &#34;storage.googleapis.com&#34; service because the request from the iOS apps violates the API key iOS application restrictions:

 { &#34;reason&#34;: &#34;API_KEY_IOS_APP_BLOCKED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34; } } |
| RATE_LIMIT_EXCEEDED | 5 | The request is denied because there is not enough rate quota for the consumer.

Example of an ErrorInfo when the consumer &#34;projects/123&#34; fails to contact &#34;pubsub.googleapis.com&#34; service because consumer&#39;s rate quota usage has reached the maximum value set for the quota limit &#34;ReadsPerMinutePerProject&#34; on the quota metric &#34;pubsub.googleapis.com/read_requests&#34;:

 { &#34;reason&#34;: &#34;RATE_LIMIT_EXCEEDED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;pubsub.googleapis.com&#34;, &#34;quota_metric&#34;: &#34;pubsub.googleapis.com/read_requests&#34;, &#34;quota_limit&#34;: &#34;ReadsPerMinutePerProject&#34; } }

Example of an ErrorInfo when the consumer &#34;projects/123&#34; checks quota on the service &#34;dataflow.googleapis.com&#34; and hits the organization quota limit &#34;DefaultRequestsPerMinutePerOrganization&#34; on the metric &#34;dataflow.googleapis.com/default_requests&#34;.

 { &#34;reason&#34;: &#34;RATE_LIMIT_EXCEEDED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;dataflow.googleapis.com&#34;, &#34;quota_metric&#34;: &#34;dataflow.googleapis.com/default_requests&#34;, &#34;quota_limit&#34;: &#34;DefaultRequestsPerMinutePerOrganization&#34; } } |
| RESOURCE_QUOTA_EXCEEDED | 6 | The request is denied because there is not enough resource quota for the consumer.

Example of an ErrorInfo when the consumer &#34;projects/123&#34; fails to contact &#34;compute.googleapis.com&#34; service because consumer&#39;s resource quota usage has reached the maximum value set for the quota limit &#34;VMsPerProject&#34; on the quota metric &#34;compute.googleapis.com/vms&#34;:

 { &#34;reason&#34;: &#34;RESOURCE_QUOTA_EXCEEDED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;compute.googleapis.com&#34;, &#34;quota_metric&#34;: &#34;compute.googleapis.com/vms&#34;, &#34;quota_limit&#34;: &#34;VMsPerProject&#34; } }

Example of an ErrorInfo when the consumer &#34;projects/123&#34; checks resource quota on the service &#34;dataflow.googleapis.com&#34; and hits the organization quota limit &#34;jobs-per-organization&#34; on the metric &#34;dataflow.googleapis.com/job_count&#34;.

 { &#34;reason&#34;: &#34;RESOURCE_QUOTA_EXCEEDED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;dataflow.googleapis.com&#34;, &#34;quota_metric&#34;: &#34;dataflow.googleapis.com/job_count&#34;, &#34;quota_limit&#34;: &#34;jobs-per-organization&#34; } } |
| LOCATION_TAX_POLICY_VIOLATED | 10 | The request whose associated billing account address is in a tax restricted location, violates the local tax restrictions when creating resources in the restricted region.

Example of an ErrorInfo when creating the Cloud Storage Bucket in the container &#34;projects/123&#34; under a tax restricted region &#34;locations/asia-northeast3&#34;:

 { &#34;reason&#34;: &#34;LOCATION_TAX_POLICY_VIOLATED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34;, &#34;location&#34;: &#34;locations/asia-northeast3&#34; } }

This response indicates creating the Cloud Storage Bucket in &#34;locations/asia-northeast3&#34; violates the location tax restriction. |
| USER_PROJECT_DENIED | 11 | The request is denied because the caller does not have required permission on the user project &#34;projects/123&#34; or the user project is invalid. For more information, check the [userProject System Parameters](https://cloud.google.com/apis/docs/system-parameters).

Example of an ErrorInfo when the caller is calling Cloud Storage service with insufficient permissions on the user project:

 { &#34;reason&#34;: &#34;USER_PROJECT_DENIED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34; } } |
| CONSUMER_SUSPENDED | 12 | The request is denied because the consumer &#34;projects/123&#34; is suspended due to Terms of Service(Tos) violations. Check [Project suspension guidelines](https://cloud.google.com/resource-manager/docs/project-suspension-guidelines) for more information.

Example of an ErrorInfo when calling Cloud Storage service with the suspended consumer &#34;projects/123&#34;:

 { &#34;reason&#34;: &#34;CONSUMER_SUSPENDED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34; } } |
| CONSUMER_INVALID | 14 | The request is denied because the associated consumer is invalid. It may be in a bad format, cannot be found, or have been deleted.

Example of an ErrorInfo when calling Cloud Storage service with the invalid consumer &#34;projects/123&#34;:

 { &#34;reason&#34;: &#34;CONSUMER_INVALID&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34; } } |
| SECURITY_POLICY_VIOLATED | 15 | The request is denied because it violates [VPC Service Controls](https://cloud.google.com/vpc-service-controls/docs/overview). The &#39;uid&#39; field is a random generated identifier that customer can use it to search the audit log for a request rejected by VPC Service Controls. For more information, please refer [VPC Service Controls Troubleshooting](https://cloud.google.com/vpc-service-controls/docs/troubleshooting#unique-id)

Example of an ErrorInfo when the consumer &#34;projects/123&#34; fails to call Cloud Storage service because the request is prohibited by the VPC Service Controls.

 { &#34;reason&#34;: &#34;SECURITY_POLICY_VIOLATED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;uid&#34;: &#34;123456789abcde&#34;, &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34; } } |
| ACCESS_TOKEN_EXPIRED | 16 | The request is denied because the provided access token has expired.

Example of an ErrorInfo when the request is calling Cloud Storage service with an expired access token:

 { &#34;reason&#34;: &#34;ACCESS_TOKEN_EXPIRED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;service&#34;: &#34;storage.googleapis.com&#34;, &#34;method&#34;: &#34;google.storage.v1.Storage.GetObject&#34; } } |
| ACCESS_TOKEN_SCOPE_INSUFFICIENT | 17 | The request is denied because the provided access token doesn&#39;t have at least one of the acceptable scopes required for the API. Please check [OAuth 2.0 Scopes for Google APIs](https://developers.google.com/identity/protocols/oauth2/scopes) for the list of the OAuth 2.0 scopes that you might need to request to access the API.

Example of an ErrorInfo when the request is calling Cloud Storage service with an access token that is missing required scopes:

 { &#34;reason&#34;: &#34;ACCESS_TOKEN_SCOPE_INSUFFICIENT&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;service&#34;: &#34;storage.googleapis.com&#34;, &#34;method&#34;: &#34;google.storage.v1.Storage.GetObject&#34; } } |
| ACCOUNT_STATE_INVALID | 18 | The request is denied because the account associated with the provided access token is in an invalid state, such as disabled or deleted. For more information, see https://cloud.google.com/docs/authentication.

Warning: For privacy reasons, the server may not be able to disclose the email address for some accounts. The client MUST NOT depend on the availability of the `email` attribute.

Example of an ErrorInfo when the request is to the Cloud Storage API with an access token that is associated with a disabled or deleted [service account](http://cloud/iam/docs/service-accounts):

 { &#34;reason&#34;: &#34;ACCOUNT_STATE_INVALID&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;service&#34;: &#34;storage.googleapis.com&#34;, &#34;method&#34;: &#34;google.storage.v1.Storage.GetObject&#34;, &#34;email&#34;: &#34;user@123.iam.gserviceaccount.com&#34; } } |
| ACCESS_TOKEN_TYPE_UNSUPPORTED | 19 | The request is denied because the type of the provided access token is not supported by the API being called.

Example of an ErrorInfo when the request is to the Cloud Storage API with an unsupported token type.

 { &#34;reason&#34;: &#34;ACCESS_TOKEN_TYPE_UNSUPPORTED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;service&#34;: &#34;storage.googleapis.com&#34;, &#34;method&#34;: &#34;google.storage.v1.Storage.GetObject&#34; } } |
| CREDENTIALS_MISSING | 20 | The request is denied because the request doesn&#39;t have any authentication credentials. For more information regarding the supported authentication strategies for Google Cloud APIs, see https://cloud.google.com/docs/authentication.

Example of an ErrorInfo when the request is to the Cloud Storage API without any authentication credentials.

 { &#34;reason&#34;: &#34;CREDENTIALS_MISSING&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;service&#34;: &#34;storage.googleapis.com&#34;, &#34;method&#34;: &#34;google.storage.v1.Storage.GetObject&#34; } } |
| RESOURCE_PROJECT_INVALID | 21 | The request is denied because the provided project owning the resource which acts as the [API consumer](https://cloud.google.com/apis/design/glossary#api_consumer) is invalid. It may be in a bad format or empty.

Example of an ErrorInfo when the request is to the Cloud Functions API, but the offered resource project in the request in a bad format which can&#39;t perform the ListFunctions method.

 { &#34;reason&#34;: &#34;RESOURCE_PROJECT_INVALID&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;service&#34;: &#34;cloudfunctions.googleapis.com&#34;, &#34;method&#34;: &#34;google.cloud.functions.v1.CloudFunctionsService.ListFunctions&#34; } } |
| SESSION_COOKIE_INVALID | 23 | The request is denied because the provided session cookie is missing, invalid or failed to decode.

Example of an ErrorInfo when the request is calling Cloud Storage service with a SID cookie which can&#39;t be decoded.

 { &#34;reason&#34;: &#34;SESSION_COOKIE_INVALID&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;service&#34;: &#34;storage.googleapis.com&#34;, &#34;method&#34;: &#34;google.storage.v1.Storage.GetObject&#34;, &#34;cookie&#34;: &#34;SID&#34; } } |
| USER_BLOCKED_BY_ADMIN | 24 | The request is denied because the user is from a Google Workspace customer that blocks their users from accessing a particular service.

Example scenario: https://support.google.com/a/answer/9197205?hl=en

Example of an ErrorInfo when access to Google Cloud Storage service is blocked by the Google Workspace administrator:

 { &#34;reason&#34;: &#34;USER_BLOCKED_BY_ADMIN&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;service&#34;: &#34;storage.googleapis.com&#34;, &#34;method&#34;: &#34;google.storage.v1.Storage.GetObject&#34;, } } |
| RESOURCE_USAGE_RESTRICTION_VIOLATED | 25 | The request is denied because the resource service usage is restricted by administrators according to the organization policy constraint. For more information see https://cloud.google.com/resource-manager/docs/organization-policy/restricting-services.

Example of an ErrorInfo when access to Google Cloud Storage service is restricted by Resource Usage Restriction policy:

 { &#34;reason&#34;: &#34;RESOURCE_USAGE_RESTRICTION_VIOLATED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/project-123&#34;, &#34;service&#34;: &#34;storage.googleapis.com&#34; } } |
| SYSTEM_PARAMETER_UNSUPPORTED | 26 | Unimplemented. Do not use.

The request is denied because it contains unsupported system parameters in URL query parameters or HTTP headers. For more information, see https://cloud.google.com/apis/docs/system-parameters

Example of an ErrorInfo when access &#34;pubsub.googleapis.com&#34; service with a request header of &#34;x-goog-user-ip&#34;:

 { &#34;reason&#34;: &#34;SYSTEM_PARAMETER_UNSUPPORTED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;service&#34;: &#34;pubsub.googleapis.com&#34; &#34;parameter&#34;: &#34;x-goog-user-ip&#34; } } |
| ORG_RESTRICTION_VIOLATION | 27 | The request is denied because it violates Org Restriction: the requested resource does not belong to allowed organizations specified in &#34;X-Goog-Allowed-Resources&#34; header.

Example of an ErrorInfo when accessing a GCP resource that is restricted by Org Restriction for &#34;pubsub.googleapis.com&#34; service.

{ reason: &#34;ORG_RESTRICTION_VIOLATION&#34; domain: &#34;googleapis.com&#34; metadata { &#34;consumer&#34;:&#34;projects/123456&#34; &#34;service&#34;: &#34;pubsub.googleapis.com&#34; } } |
| ORG_RESTRICTION_HEADER_INVALID | 28 | The request is denied because &#34;X-Goog-Allowed-Resources&#34; header is in a bad format.

Example of an ErrorInfo when accessing &#34;pubsub.googleapis.com&#34; service with an invalid &#34;X-Goog-Allowed-Resources&#34; request header.

{ reason: &#34;ORG_RESTRICTION_HEADER_INVALID&#34; domain: &#34;googleapis.com&#34; metadata { &#34;consumer&#34;:&#34;projects/123456&#34; &#34;service&#34;: &#34;pubsub.googleapis.com&#34; } } |
| SERVICE_NOT_VISIBLE | 29 | Unimplemented. Do not use.

The request is calling a service that is not visible to the consumer.

Example of an ErrorInfo when the consumer &#34;projects/123&#34; contacting &#34;pubsub.googleapis.com&#34; service which is not visible to the consumer.

 { &#34;reason&#34;: &#34;SERVICE_NOT_VISIBLE&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;pubsub.googleapis.com&#34; } }

This response indicates the &#34;pubsub.googleapis.com&#34; is not visible to &#34;projects/123&#34; (or it may not exist). |
| GCP_SUSPENDED | 30 | The request is related to a project for which GCP access is suspended.

Example of an ErrorInfo when the consumer &#34;projects/123&#34; fails to contact &#34;pubsub.googleapis.com&#34; service because GCP access is suspended:

 { &#34;reason&#34;: &#34;GCP_SUSPENDED&#34;, &#34;domain&#34;: &#34;googleapis.com&#34;, &#34;metadata&#34;: { &#34;consumer&#34;: &#34;projects/123&#34;, &#34;service&#34;: &#34;pubsub.googleapis.com&#34; } }

This response indicates the associated GCP account has been suspended. |


 

 

 



<a name="google_api_httpbody-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/httpbody.proto



<a name="google-api-HttpBody"></a>

### HttpBody
Message that represents an arbitrary HTTP body. It should only be used for
payload formats that can&#39;t be represented as JSON, such as raw binary or
an HTML page.


This message can be used both in streaming and non-streaming API methods in
the request as well as the response.

It can be used as a top-level request field, which is convenient if one
wants to extract parameters from either the URL or HTTP template into the
request fields and also want access to the raw HTTP body.

Example:

    message GetResourceRequest {
      // A unique request id.
      string request_id = 1;

      // The raw HTTP body is bound to this field.
      google.api.HttpBody http_body = 2;

    }

    service ResourceService {
      rpc GetResource(GetResourceRequest)
        returns (google.api.HttpBody);
      rpc UpdateResource(google.api.HttpBody)
        returns (google.protobuf.Empty);

    }

Example with streaming methods:

    service CaldavService {
      rpc GetCalendar(stream google.api.HttpBody)
        returns (stream google.api.HttpBody);
      rpc UpdateCalendar(stream google.api.HttpBody)
        returns (stream google.api.HttpBody);

    }

Use of this type only changes how the request and response bodies are
handled, all other features will continue to work unchanged.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content_type | [string](#string) |  | The HTTP Content-Type header value specifying the content type of the body. |
| data | [bytes](#bytes) |  | The HTTP request/response body as raw binary. |
| extensions | [google.protobuf.Any](#google-protobuf-Any) | repeated | Application specific response metadata. Must be set in the first response for streaming APIs. |





 

 

 

 



<a name="google_api_label-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/label.proto



<a name="google-api-LabelDescriptor"></a>

### LabelDescriptor
A description of a label.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | The label key. |
| value_type | [LabelDescriptor.ValueType](#google-api-LabelDescriptor-ValueType) |  | The type of data that can be assigned to the label. |
| description | [string](#string) |  | A human-readable description for the label. |





 


<a name="google-api-LabelDescriptor-ValueType"></a>

### LabelDescriptor.ValueType
Value types that can be used as label values.

| Name | Number | Description |
| ---- | ------ | ----------- |
| STRING | 0 | A variable-length string. This is the default. |
| BOOL | 1 | Boolean; true or false. |
| INT64 | 2 | A 64-bit signed integer. |


 

 

 



<a name="google_api_log-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/log.proto



<a name="google-api-LogDescriptor"></a>

### LogDescriptor
A description of a log type. Example in YAML format:

    - name: library.googleapis.com/activity_history
      description: The history of borrowing and returning library items.
      display_name: Activity
      labels:
      - key: /customer_id
        description: Identifier of a library customer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the log. It must be less than 512 characters long and can include the following characters: upper- and lower-case alphanumeric characters [A-Za-z0-9], and punctuation characters including slash, underscore, hyphen, period [/_-.]. |
| labels | [LabelDescriptor](#google-api-LabelDescriptor) | repeated | The set of labels that are available to describe a specific log entry. Runtime requests that contain labels not specified here are considered invalid. |
| description | [string](#string) |  | A human-readable description of this log. This information appears in the documentation and can contain details. |
| display_name | [string](#string) |  | The human-readable name for this log. This information appears on the user interface and should be concise. |





 

 

 

 



<a name="google_api_logging-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/logging.proto



<a name="google-api-Logging"></a>

### Logging
Logging configuration of the service.

The following example shows how to configure logs to be sent to the
producer and consumer projects. In the example, the `activity_history`
log is sent to both the producer and consumer projects, whereas the
`purchase_history` log is only sent to the producer project.

    monitored_resources:
    - type: library.googleapis.com/branch
      labels:
      - key: /city
        description: The city where the library branch is located in.
      - key: /name
        description: The name of the branch.
    logs:
    - name: activity_history
      labels:
      - key: /customer_id
    - name: purchase_history
    logging:
      producer_destinations:
      - monitored_resource: library.googleapis.com/branch
        logs:
        - activity_history
        - purchase_history
      consumer_destinations:
      - monitored_resource: library.googleapis.com/branch
        logs:
        - activity_history


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| producer_destinations | [Logging.LoggingDestination](#google-api-Logging-LoggingDestination) | repeated | Logging configurations for sending logs to the producer project. There can be multiple producer destinations, each one must have a different monitored resource type. A log can be used in at most one producer destination. |
| consumer_destinations | [Logging.LoggingDestination](#google-api-Logging-LoggingDestination) | repeated | Logging configurations for sending logs to the consumer project. There can be multiple consumer destinations, each one must have a different monitored resource type. A log can be used in at most one consumer destination. |






<a name="google-api-Logging-LoggingDestination"></a>

### Logging.LoggingDestination
Configuration of a specific logging destination (the producer project
or the consumer project).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| monitored_resource | [string](#string) |  | The monitored resource type. The type must be defined in the [Service.monitored_resources][google.api.Service.monitored_resources] section. |
| logs | [string](#string) | repeated | Names of the logs to be sent to this destination. Each name must be defined in the [Service.logs][google.api.Service.logs] section. If the log name is not a domain scoped name, it will be automatically prefixed with the service name followed by &#34;/&#34;. |





 

 

 

 



<a name="google_api_metric-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/metric.proto



<a name="google-api-Metric"></a>

### Metric
A specific metric, identified by specifying values for all of the
labels of a [`MetricDescriptor`][google.api.MetricDescriptor].


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | An existing metric type, see [google.api.MetricDescriptor][google.api.MetricDescriptor]. For example, `custom.googleapis.com/invoice/paid/amount`. |
| labels | [Metric.LabelsEntry](#google-api-Metric-LabelsEntry) | repeated | The set of label values that uniquely identify this metric. All labels listed in the `MetricDescriptor` must be assigned values. |






<a name="google-api-Metric-LabelsEntry"></a>

### Metric.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="google-api-MetricDescriptor"></a>

### MetricDescriptor
Defines a metric type and its schema. Once a metric descriptor is created,
deleting or altering it stops data collection and makes the metric type&#39;s
existing data unusable.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The resource name of the metric descriptor. |
| type | [string](#string) |  | The metric type, including its DNS name prefix. The type is not URL-encoded. All user-defined metric types have the DNS name `custom.googleapis.com` or `external.googleapis.com`. Metric types should use a natural hierarchical grouping. For example:

 &#34;custom.googleapis.com/invoice/paid/amount&#34; &#34;external.googleapis.com/prometheus/up&#34; &#34;appengine.googleapis.com/http/server/response_latencies&#34; |
| labels | [LabelDescriptor](#google-api-LabelDescriptor) | repeated | The set of labels that can be used to describe a specific instance of this metric type. For example, the `appengine.googleapis.com/http/server/response_latencies` metric type has a label for the HTTP response code, `response_code`, so you can look at latencies for successful responses or just for responses that failed. |
| metric_kind | [MetricDescriptor.MetricKind](#google-api-MetricDescriptor-MetricKind) |  | Whether the metric records instantaneous values, changes to a value, etc. Some combinations of `metric_kind` and `value_type` might not be supported. |
| value_type | [MetricDescriptor.ValueType](#google-api-MetricDescriptor-ValueType) |  | Whether the measurement is an integer, a floating-point number, etc. Some combinations of `metric_kind` and `value_type` might not be supported. |
| unit | [string](#string) |  | The units in which the metric value is reported. It is only applicable if the `value_type` is `INT64`, `DOUBLE`, or `DISTRIBUTION`. The `unit` defines the representation of the stored metric values.

Different systems might scale the values to be more easily displayed (so a value of `0.02kBy` _might_ be displayed as `20By`, and a value of `3523kBy` _might_ be displayed as `3.5MBy`). However, if the `unit` is `kBy`, then the value of the metric is always in thousands of bytes, no matter how it might be displayed.

If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an `INT64 CUMULATIVE` metric whose `unit` is `s{CPU}` (or equivalently `1s{CPU}` or just `s`). If the job uses 12,005 CPU-seconds, then the value is written as `12005`.

Alternatively, if you want a custom metric to record data in a more granular way, you can create a `DOUBLE CUMULATIVE` metric whose `unit` is `ks{CPU}`, and then write the value `12.005` (which is `12005/1000`), or use `Kis{CPU}` and write `11.723` (which is `12005/1024`).

The supported units are a subset of [The Unified Code for Units of Measure](https://unitsofmeasure.org/ucum.html) standard:

**Basic units (UNIT)**

* `bit` bit * `By` byte * `s` second * `min` minute * `h` hour * `d` day * `1` dimensionless

**Prefixes (PREFIX)**

* `k` kilo (10^3) * `M` mega (10^6) * `G` giga (10^9) * `T` tera (10^12) * `P` peta (10^15) * `E` exa (10^18) * `Z` zetta (10^21) * `Y` yotta (10^24)

* `m` milli (10^-3) * `u` micro (10^-6) * `n` nano (10^-9) * `p` pico (10^-12) * `f` femto (10^-15) * `a` atto (10^-18) * `z` zepto (10^-21) * `y` yocto (10^-24)

* `Ki` kibi (2^10) * `Mi` mebi (2^20) * `Gi` gibi (2^30) * `Ti` tebi (2^40) * `Pi` pebi (2^50)

**Grammar**

The grammar also includes these connectors:

* `/` division or ratio (as an infix operator). For examples, `kBy/{email}` or `MiBy/10ms` (although you should almost never have `/s` in a metric `unit`; rates should always be computed at query time from the underlying cumulative or delta value). * `.` multiplication or composition (as an infix operator). For examples, `GBy.d` or `k{watt}.h`.

The grammar for a unit is as follows:

 Expression = Component { &#34;.&#34; Component } { &#34;/&#34; Component } ;

 Component = ( [ PREFIX ] UNIT | &#34;%&#34; ) [ Annotation ] | Annotation | &#34;1&#34; ;

 Annotation = &#34;{&#34; NAME &#34;}&#34; ;

Notes:

* `Annotation` is just a comment if it follows a `UNIT`. If the annotation is used alone, then the unit is equivalent to `1`. For examples, `{request}/s == 1/s`, `By{transmitted}/s == By/s`. * `NAME` is a sequence of non-blank printable ASCII characters not containing `{` or `}`. * `1` represents a unitary [dimensionless unit](https://en.wikipedia.org/wiki/Dimensionless_quantity) of 1, such as in `1/s`. It is typically used when none of the basic units are appropriate. For example, &#34;new users per day&#34; can be represented as `1/d` or `{new-users}/d` (and a metric value `5` would mean &#34;5 new users). Alternatively, &#34;thousands of page views per day&#34; would be represented as `1000/d` or `k1/d` or `k{page_views}/d` (and a metric value of `5.3` would mean &#34;5300 page views per day&#34;). * `%` represents dimensionless value of 1/100, and annotates values giving a percentage (so the metric values are typically in the range of 0..100, and a metric value `3` means &#34;3 percent&#34;). * `10^2.%` indicates a metric contains a ratio, typically in the range 0..1, that will be multiplied by 100 and displayed as a percentage (so a metric value `0.03` means &#34;3 percent&#34;). |
| description | [string](#string) |  | A detailed description of the metric, which can be used in documentation. |
| display_name | [string](#string) |  | A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example &#34;Request count&#34;. This field is optional but it is recommended to be set for any metrics associated with user-visible concepts, such as Quota. |
| metadata | [MetricDescriptor.MetricDescriptorMetadata](#google-api-MetricDescriptor-MetricDescriptorMetadata) |  | Optional. Metadata which can be used to guide usage of the metric. |
| launch_stage | [LaunchStage](#google-api-LaunchStage) |  | Optional. The launch stage of the metric definition. |
| monitored_resource_types | [string](#string) | repeated | Read-only. If present, then a [time series][google.monitoring.v3.TimeSeries], which is identified partially by a metric type and a [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor], that is associated with this metric type can only be associated with one of the monitored resource types listed here. |






<a name="google-api-MetricDescriptor-MetricDescriptorMetadata"></a>

### MetricDescriptor.MetricDescriptorMetadata
Additional annotations that can be used to guide the usage of a metric.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| launch_stage | [LaunchStage](#google-api-LaunchStage) |  | **Deprecated.** Deprecated. Must use the [MetricDescriptor.launch_stage][google.api.MetricDescriptor.launch_stage] instead. |
| sample_period | [google.protobuf.Duration](#google-protobuf-Duration) |  | The sampling period of metric data points. For metrics which are written periodically, consecutive data points are stored at this time interval, excluding data loss due to errors. Metrics with a higher granularity have a smaller sampling period. |
| ingest_delay | [google.protobuf.Duration](#google-protobuf-Duration) |  | The delay of data points caused by ingestion. Data points older than this age are guaranteed to be ingested and available to be read, excluding data loss due to errors. |





 


<a name="google-api-MetricDescriptor-MetricKind"></a>

### MetricDescriptor.MetricKind
The kind of measurement. It describes how the data is reported.
For information on setting the start time and end time based on
the MetricKind, see [TimeInterval][google.monitoring.v3.TimeInterval].

| Name | Number | Description |
| ---- | ------ | ----------- |
| METRIC_KIND_UNSPECIFIED | 0 | Do not use this default value. |
| GAUGE | 1 | An instantaneous measurement of a value. |
| DELTA | 2 | The change in a value during a time interval. |
| CUMULATIVE | 3 | A value accumulated over a time interval. Cumulative measurements in a time series should have the same start time and increasing end times, until an event resets the cumulative value to zero and sets a new start time for the following points. |



<a name="google-api-MetricDescriptor-ValueType"></a>

### MetricDescriptor.ValueType
The value type of a metric.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VALUE_TYPE_UNSPECIFIED | 0 | Do not use this default value. |
| BOOL | 1 | The value is a boolean. This value type can be used only if the metric kind is `GAUGE`. |
| INT64 | 2 | The value is a signed 64-bit integer. |
| DOUBLE | 3 | The value is a double precision floating point number. |
| STRING | 4 | The value is a text string. This value type can be used only if the metric kind is `GAUGE`. |
| DISTRIBUTION | 5 | The value is a [`Distribution`][google.api.Distribution]. |
| MONEY | 6 | The value is money. |


 

 

 



<a name="google_api_monitored_resource-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/monitored_resource.proto



<a name="google-api-MonitoredResource"></a>

### MonitoredResource
An object representing a resource that can be used for monitoring, logging,
billing, or other purposes. Examples include virtual machine instances,
databases, and storage devices such as disks. The `type` field identifies a
[MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor] object
that describes the resource&#39;s schema. Information in the `labels` field
identifies the actual resource and its attributes according to the schema.
For example, a particular Compute Engine VM instance could be represented by
the following object, because the
[MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor] for
`&#34;gce_instance&#34;` has labels
`&#34;project_id&#34;`, `&#34;instance_id&#34;` and `&#34;zone&#34;`:

    { &#34;type&#34;: &#34;gce_instance&#34;,
      &#34;labels&#34;: { &#34;project_id&#34;: &#34;my-project&#34;,
                  &#34;instance_id&#34;: &#34;12345678901234&#34;,
                  &#34;zone&#34;: &#34;us-central1-a&#34; }}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Required. The monitored resource type. This field must match the `type` field of a [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor] object. For example, the type of a Compute Engine VM instance is `gce_instance`. Some descriptors include the service name in the type; for example, the type of a Datastream stream is `datastream.googleapis.com/Stream`. |
| labels | [MonitoredResource.LabelsEntry](#google-api-MonitoredResource-LabelsEntry) | repeated | Required. Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels `&#34;project_id&#34;`, `&#34;instance_id&#34;`, and `&#34;zone&#34;`. |






<a name="google-api-MonitoredResource-LabelsEntry"></a>

### MonitoredResource.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="google-api-MonitoredResourceDescriptor"></a>

### MonitoredResourceDescriptor
An object that describes the schema of a
[MonitoredResource][google.api.MonitoredResource] object using a type name
and a set of labels.  For example, the monitored resource descriptor for
Google Compute Engine VM instances has a type of
`&#34;gce_instance&#34;` and specifies the use of the labels `&#34;instance_id&#34;` and
`&#34;zone&#34;` to identify particular VM instances.

Different APIs can support different monitored resource types. APIs generally
provide a `list` method that returns the monitored resource descriptors used
by the API.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Optional. The resource name of the monitored resource descriptor: `&#34;projects/{project_id}/monitoredResourceDescriptors/{type}&#34;` where {type} is the value of the `type` field in this object and {project_id} is a project ID that provides API-specific context for accessing the type. APIs that do not use project information can use the resource name format `&#34;monitoredResourceDescriptors/{type}&#34;`. |
| type | [string](#string) |  | Required. The monitored resource type. For example, the type `&#34;cloudsql_database&#34;` represents databases in Google Cloud SQL. For a list of types, see [Monitoring resource types](https://cloud.google.com/monitoring/api/resources) and [Logging resource types](https://cloud.google.com/logging/docs/api/v2/resource-list). |
| display_name | [string](#string) |  | Optional. A concise name for the monitored resource type that might be displayed in user interfaces. It should be a Title Cased Noun Phrase, without any article or other determiners. For example, `&#34;Google Cloud SQL Database&#34;`. |
| description | [string](#string) |  | Optional. A detailed description of the monitored resource type that might be used in documentation. |
| labels | [LabelDescriptor](#google-api-LabelDescriptor) | repeated | Required. A set of labels used to describe instances of this monitored resource type. For example, an individual Google Cloud SQL database is identified by values for the labels `&#34;database_id&#34;` and `&#34;zone&#34;`. |
| launch_stage | [LaunchStage](#google-api-LaunchStage) |  | Optional. The launch stage of the monitored resource definition. |






<a name="google-api-MonitoredResourceMetadata"></a>

### MonitoredResourceMetadata
Auxiliary metadata for a [MonitoredResource][google.api.MonitoredResource]
object. [MonitoredResource][google.api.MonitoredResource] objects contain the
minimum set of information to uniquely identify a monitored resource
instance. There is some other useful auxiliary metadata. Monitoring and
Logging use an ingestion pipeline to extract metadata for cloud resources of
all types, and store the metadata in this message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| system_labels | [google.protobuf.Struct](#google-protobuf-Struct) |  | Output only. Values for predefined system metadata labels. System labels are a kind of metadata extracted by Google, including &#34;machine_image&#34;, &#34;vpc&#34;, &#34;subnet_id&#34;, &#34;security_group&#34;, &#34;name&#34;, etc. System label values can be only strings, Boolean values, or a list of strings. For example:

 { &#34;name&#34;: &#34;my-test-instance&#34;, &#34;security_group&#34;: [&#34;a&#34;, &#34;b&#34;, &#34;c&#34;], &#34;spot_instance&#34;: false } |
| user_labels | [MonitoredResourceMetadata.UserLabelsEntry](#google-api-MonitoredResourceMetadata-UserLabelsEntry) | repeated | Output only. A map of user-defined metadata labels. |






<a name="google-api-MonitoredResourceMetadata-UserLabelsEntry"></a>

### MonitoredResourceMetadata.UserLabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



<a name="google_api_monitoring-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/monitoring.proto



<a name="google-api-Monitoring"></a>

### Monitoring
Monitoring configuration of the service.

The example below shows how to configure monitored resources and metrics
for monitoring. In the example, a monitored resource and two metrics are
defined. The `library.googleapis.com/book/returned_count` metric is sent
to both producer and consumer projects, whereas the
`library.googleapis.com/book/num_overdue` metric is only sent to the
consumer project.

    monitored_resources:
    - type: library.googleapis.com/Branch
      display_name: &#34;Library Branch&#34;
      description: &#34;A branch of a library.&#34;
      launch_stage: GA
      labels:
      - key: resource_container
        description: &#34;The Cloud container (ie. project id) for the Branch.&#34;
      - key: location
        description: &#34;The location of the library branch.&#34;
      - key: branch_id
        description: &#34;The id of the branch.&#34;
    metrics:
    - name: library.googleapis.com/book/returned_count
      display_name: &#34;Books Returned&#34;
      description: &#34;The count of books that have been returned.&#34;
      launch_stage: GA
      metric_kind: DELTA
      value_type: INT64
      unit: &#34;1&#34;
      labels:
      - key: customer_id
        description: &#34;The id of the customer.&#34;
    - name: library.googleapis.com/book/num_overdue
      display_name: &#34;Books Overdue&#34;
      description: &#34;The current number of overdue books.&#34;
      launch_stage: GA
      metric_kind: GAUGE
      value_type: INT64
      unit: &#34;1&#34;
      labels:
      - key: customer_id
        description: &#34;The id of the customer.&#34;
    monitoring:
      producer_destinations:
      - monitored_resource: library.googleapis.com/Branch
        metrics:
        - library.googleapis.com/book/returned_count
      consumer_destinations:
      - monitored_resource: library.googleapis.com/Branch
        metrics:
        - library.googleapis.com/book/returned_count
        - library.googleapis.com/book/num_overdue


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| producer_destinations | [Monitoring.MonitoringDestination](#google-api-Monitoring-MonitoringDestination) | repeated | Monitoring configurations for sending metrics to the producer project. There can be multiple producer destinations. A monitored resource type may appear in multiple monitoring destinations if different aggregations are needed for different sets of metrics associated with that monitored resource type. A monitored resource and metric pair may only be used once in the Monitoring configuration. |
| consumer_destinations | [Monitoring.MonitoringDestination](#google-api-Monitoring-MonitoringDestination) | repeated | Monitoring configurations for sending metrics to the consumer project. There can be multiple consumer destinations. A monitored resource type may appear in multiple monitoring destinations if different aggregations are needed for different sets of metrics associated with that monitored resource type. A monitored resource and metric pair may only be used once in the Monitoring configuration. |






<a name="google-api-Monitoring-MonitoringDestination"></a>

### Monitoring.MonitoringDestination
Configuration of a specific monitoring destination (the producer project
or the consumer project).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| monitored_resource | [string](#string) |  | The monitored resource type. The type must be defined in [Service.monitored_resources][google.api.Service.monitored_resources] section. |
| metrics | [string](#string) | repeated | Types of the metrics to report to this monitoring destination. Each type must be defined in [Service.metrics][google.api.Service.metrics] section. |





 

 

 

 



<a name="google_api_quota-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/quota.proto



<a name="google-api-MetricRule"></a>

### MetricRule
Bind API methods to metrics. Binding a method to a metric causes that
metric&#39;s configured quota behaviors to apply to the method call.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | Selects the methods to which this rule applies.

Refer to [selector][google.api.DocumentationRule.selector] for syntax details. |
| metric_costs | [MetricRule.MetricCostsEntry](#google-api-MetricRule-MetricCostsEntry) | repeated | Metrics to update when the selected methods are called, and the associated cost applied to each metric.

The key of the map is the metric name, and the values are the amount increased for the metric against which the quota limits are defined. The value must not be negative. |






<a name="google-api-MetricRule-MetricCostsEntry"></a>

### MetricRule.MetricCostsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [int64](#int64) |  |  |






<a name="google-api-Quota"></a>

### Quota
Quota configuration helps to achieve fairness and budgeting in service
usage.

The metric based quota configuration works this way:
- The service configuration defines a set of metrics.
- For API calls, the quota.metric_rules maps methods to metrics with
  corresponding costs.
- The quota.limits defines limits on the metrics, which will be used for
  quota checks at runtime.

An example quota configuration in yaml format:

   quota:
     limits:

     - name: apiWriteQpsPerProject
       metric: library.googleapis.com/write_calls
       unit: &#34;1/min/{project}&#34;  # rate limit for consumer projects
       values:
         STANDARD: 10000


     (The metric rules bind all methods to the read_calls metric,
      except for the UpdateBook and DeleteBook methods. These two methods
      are mapped to the write_calls metric, with the UpdateBook method
      consuming at twice rate as the DeleteBook method.)
     metric_rules:
     - selector: &#34;*&#34;
       metric_costs:
         library.googleapis.com/read_calls: 1
     - selector: google.example.library.v1.LibraryService.UpdateBook
       metric_costs:
         library.googleapis.com/write_calls: 2
     - selector: google.example.library.v1.LibraryService.DeleteBook
       metric_costs:
         library.googleapis.com/write_calls: 1

 Corresponding Metric definition:

     metrics:
     - name: library.googleapis.com/read_calls
       display_name: Read requests
       metric_kind: DELTA
       value_type: INT64

     - name: library.googleapis.com/write_calls
       display_name: Write requests
       metric_kind: DELTA
       value_type: INT64


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limits | [QuotaLimit](#google-api-QuotaLimit) | repeated | List of QuotaLimit definitions for the service. |
| metric_rules | [MetricRule](#google-api-MetricRule) | repeated | List of MetricRule definitions, each one mapping a selected method to one or more metrics. |






<a name="google-api-QuotaLimit"></a>

### QuotaLimit
`QuotaLimit` defines a specific limit that applies over a specified duration
for a limit type. There can be at most one limit for a duration and limit
type combination defined within a `QuotaGroup`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the quota limit.

The name must be provided, and it must be unique within the service. The name can only include alphanumeric characters as well as &#39;-&#39;.

The maximum length of the limit name is 64 characters. |
| description | [string](#string) |  | Optional. User-visible, extended description for this quota limit. Should be used only when more context is needed to understand this limit than provided by the limit&#39;s display name (see: `display_name`). |
| default_limit | [int64](#int64) |  | Default number of tokens that can be consumed during the specified duration. This is the number of tokens assigned when a client application developer activates the service for his/her project.

Specifying a value of 0 will block all requests. This can be used if you are provisioning quota to selected consumers and blocking others. Similarly, a value of -1 will indicate an unlimited quota. No other negative values are allowed.

Used by group-based quotas only. |
| max_limit | [int64](#int64) |  | Maximum number of tokens that can be consumed during the specified duration. Client application developers can override the default limit up to this maximum. If specified, this value cannot be set to a value less than the default limit. If not specified, it is set to the default limit.

To allow clients to apply overrides with no upper bound, set this to -1, indicating unlimited maximum quota.

Used by group-based quotas only. |
| free_tier | [int64](#int64) |  | Free tier value displayed in the Developers Console for this limit. The free tier is the number of tokens that will be subtracted from the billed amount when billing is enabled. This field can only be set on a limit with duration &#34;1d&#34;, in a billable group; it is invalid on any other limit. If this field is not set, it defaults to 0, indicating that there is no free tier for this service.

Used by group-based quotas only. |
| duration | [string](#string) |  | Duration of this limit in textual notation. Must be &#34;100s&#34; or &#34;1d&#34;.

Used by group-based quotas only. |
| metric | [string](#string) |  | The name of the metric this quota limit applies to. The quota limits with the same metric will be checked together during runtime. The metric must be defined within the service config. |
| unit | [string](#string) |  | Specify the unit of the quota limit. It uses the same syntax as [Metric.unit][]. The supported unit kinds are determined by the quota backend system.

Here are some examples: * &#34;1/min/{project}&#34; for quota per minute per project.

Note: the order of unit components is insignificant. The &#34;1&#34; at the beginning is required to follow the metric unit syntax. |
| values | [QuotaLimit.ValuesEntry](#google-api-QuotaLimit-ValuesEntry) | repeated | Tiered limit values. You must specify this as a key:value pair, with an integer value that is the maximum number of requests allowed for the specified unit. Currently only STANDARD is supported. |
| display_name | [string](#string) |  | User-visible display name for this limit. Optional. If not set, the UI will provide a default display name based on the quota configuration. This field can be used to override the default display name generated from the configuration. |






<a name="google-api-QuotaLimit-ValuesEntry"></a>

### QuotaLimit.ValuesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [int64](#int64) |  |  |





 

 

 

 



<a name="google_api_resource-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/resource.proto



<a name="google-api-ResourceDescriptor"></a>

### ResourceDescriptor
A simple descriptor of a resource type.

ResourceDescriptor annotates a resource message (either by means of a
protobuf annotation or use in the service config), and associates the
resource&#39;s schema, the resource type, and the pattern of the resource name.

Example:

    message Topic {
      // Indicates this message defines a resource schema.
      // Declares the resource type in the format of {service}/{kind}.
      // For Kubernetes resources, the format is {api group}/{kind}.
      option (google.api.resource) = {
        type: &#34;pubsub.googleapis.com/Topic&#34;
        pattern: &#34;projects/{project}/topics/{topic}&#34;
      };
    }

The ResourceDescriptor Yaml config will look like:

    resources:
    - type: &#34;pubsub.googleapis.com/Topic&#34;
      pattern: &#34;projects/{project}/topics/{topic}&#34;

Sometimes, resources have multiple patterns, typically because they can
live under multiple parents.

Example:

    message LogEntry {
      option (google.api.resource) = {
        type: &#34;logging.googleapis.com/LogEntry&#34;
        pattern: &#34;projects/{project}/logs/{log}&#34;
        pattern: &#34;folders/{folder}/logs/{log}&#34;
        pattern: &#34;organizations/{organization}/logs/{log}&#34;
        pattern: &#34;billingAccounts/{billing_account}/logs/{log}&#34;
      };
    }

The ResourceDescriptor Yaml config will look like:

    resources:
    - type: &#39;logging.googleapis.com/LogEntry&#39;
      pattern: &#34;projects/{project}/logs/{log}&#34;
      pattern: &#34;folders/{folder}/logs/{log}&#34;
      pattern: &#34;organizations/{organization}/logs/{log}&#34;
      pattern: &#34;billingAccounts/{billing_account}/logs/{log}&#34;


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | The resource type. It must be in the format of {service_name}/{resource_type_kind}. The `resource_type_kind` must be singular and must not include version numbers.

Example: `storage.googleapis.com/Bucket`

The value of the resource_type_kind must follow the regular expression /[A-Za-z][a-zA-Z0-9]&#43;/. It should start with an upper case character and should use PascalCase (UpperCamelCase). The maximum number of characters allowed for the `resource_type_kind` is 100. |
| pattern | [string](#string) | repeated | Optional. The relative resource name pattern associated with this resource type. The DNS prefix of the full resource name shouldn&#39;t be specified here.

The path pattern must follow the syntax, which aligns with HTTP binding syntax:

 Template = Segment { &#34;/&#34; Segment } ; Segment = LITERAL | Variable ; Variable = &#34;{&#34; LITERAL &#34;}&#34; ;

Examples:

 - &#34;projects/{project}/topics/{topic}&#34; - &#34;projects/{project}/knowledgeBases/{knowledge_base}&#34;

The components in braces correspond to the IDs for each resource in the hierarchy. It is expected that, if multiple patterns are provided, the same component name (e.g. &#34;project&#34;) refers to IDs of the same type of resource. |
| name_field | [string](#string) |  | Optional. The field on the resource that designates the resource name field. If omitted, this is assumed to be &#34;name&#34;. |
| history | [ResourceDescriptor.History](#google-api-ResourceDescriptor-History) |  | Optional. The historical or future-looking state of the resource pattern.

Example:

 // The InspectTemplate message originally only supported resource // names with organization, and project was added later. message InspectTemplate { option (google.api.resource) = { type: &#34;dlp.googleapis.com/InspectTemplate&#34; pattern: &#34;organizations/{organization}/inspectTemplates/{inspect_template}&#34; pattern: &#34;projects/{project}/inspectTemplates/{inspect_template}&#34; history: ORIGINALLY_SINGLE_PATTERN }; } |
| plural | [string](#string) |  | The plural name used in the resource name and permission names, such as &#39;projects&#39; for the resource name of &#39;projects/{project}&#39; and the permission name of &#39;cloudresourcemanager.googleapis.com/projects.get&#39;. It is the same concept of the `plural` field in k8s CRD spec https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/

Note: The plural form is required even for singleton resources. See https://aip.dev/156 |
| singular | [string](#string) |  | The same concept of the `singular` field in k8s CRD spec https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/ Such as &#34;project&#34; for the `resourcemanager.googleapis.com/Project` type. |
| style | [ResourceDescriptor.Style](#google-api-ResourceDescriptor-Style) | repeated | Style flag(s) for this resource. These indicate that a resource is expected to conform to a given style. See the specific style flags for additional information. |






<a name="google-api-ResourceReference"></a>

### ResourceReference
Defines a proto annotation that describes a string field that refers to
an API resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | The resource type that the annotated field references.

Example:

 message Subscription { string topic = 2 [(google.api.resource_reference) = { type: &#34;pubsub.googleapis.com/Topic&#34; }]; }

Occasionally, a field may reference an arbitrary resource. In this case, APIs use the special value * in their resource reference.

Example:

 message GetIamPolicyRequest { string resource = 2 [(google.api.resource_reference) = { type: &#34;*&#34; }]; } |
| child_type | [string](#string) |  | The resource type of a child collection that the annotated field references. This is useful for annotating the `parent` field that doesn&#39;t have a fixed resource type.

Example:

 message ListLogEntriesRequest { string parent = 1 [(google.api.resource_reference) = { child_type: &#34;logging.googleapis.com/LogEntry&#34; }; } |





 


<a name="google-api-ResourceDescriptor-History"></a>

### ResourceDescriptor.History
A description of the historical or future-looking state of the
resource pattern.

| Name | Number | Description |
| ---- | ------ | ----------- |
| HISTORY_UNSPECIFIED | 0 | The &#34;unset&#34; value. |
| ORIGINALLY_SINGLE_PATTERN | 1 | The resource originally had one pattern and launched as such, and additional patterns were added later. |
| FUTURE_MULTI_PATTERN | 2 | The resource has one pattern, but the API owner expects to add more later. (This is the inverse of ORIGINALLY_SINGLE_PATTERN, and prevents that from being necessary once there are multiple patterns.) |



<a name="google-api-ResourceDescriptor-Style"></a>

### ResourceDescriptor.Style
A flag representing a specific style that a resource claims to conform to.

| Name | Number | Description |
| ---- | ------ | ----------- |
| STYLE_UNSPECIFIED | 0 | The unspecified value. Do not use. |
| DECLARATIVE_FRIENDLY | 1 | This resource is intended to be &#34;declarative-friendly&#34;.

Declarative-friendly resources must be more strictly consistent, and setting this to true communicates to tools that this resource should adhere to declarative-friendly expectations.

Note: This is used by the API linter (linter.aip.dev) to enable additional checks. |


 


<a name="google_api_resource-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| resource_reference | ResourceReference | .google.protobuf.FieldOptions | 1055 | An annotation that describes a resource reference, see [ResourceReference][]. |
| resource_definition | ResourceDescriptor | .google.protobuf.FileOptions | 1053 | An annotation that describes a resource definition without a corresponding message; see [ResourceDescriptor][]. |
| resource | ResourceDescriptor | .google.protobuf.MessageOptions | 1053 | An annotation that describes a resource definition, see [ResourceDescriptor][]. |

 

 



<a name="google_api_routing-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/routing.proto



<a name="google-api-RoutingParameter"></a>

### RoutingParameter
A projection from an input message to the GRPC or REST header.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [string](#string) |  | A request field to extract the header key-value pair from. |
| path_template | [string](#string) |  | A pattern matching the key-value field. Optional. If not specified, the whole field specified in the `field` field will be taken as value, and its name used as key. If specified, it MUST contain exactly one named segment (along with any number of unnamed segments) The pattern will be matched over the field specified in the `field` field, then if the match is successful: - the name of the single named segment will be used as a header name, - the match value of the segment will be used as a header value; if the match is NOT successful, nothing will be sent.

Example:

 -- This is a field in the request message | that the header value will be extracted from. | | -- This is the key name in the | | routing header. V | field: &#34;table_name&#34; v path_template: &#34;projects/*/{table_location=instances/*}/tables/*&#34; ^ ^ | | In the {} brackets is the pattern that -- | specifies what to extract from the | field as a value to be sent. | | The string in the field must match the whole pattern -- before brackets, inside brackets, after brackets.

When looking at this specific example, we can see that: - A key-value pair with the key `table_location` and the value matching `instances/*` should be added to the x-goog-request-params routing header. - The value is extracted from the request message&#39;s `table_name` field if it matches the full pattern specified: `projects/*/instances/*/tables/*`.

**NB:** If the `path_template` field is not provided, the key name is equal to the field name, and the whole field should be sent as a value. This makes the pattern for the field and the value functionally equivalent to `**`, and the configuration

 { field: &#34;table_name&#34; }

is a functionally equivalent shorthand to:

 { field: &#34;table_name&#34; path_template: &#34;{table_name=**}&#34; }

See Example 1 for more details. |






<a name="google-api-RoutingRule"></a>

### RoutingRule
Specifies the routing information that should be sent along with the request
in the form of routing header.
**NOTE:** All service configuration rules follow the &#34;last one wins&#34; order.

The examples below will apply to an RPC which has the following request type:

Message Definition:

    message Request {
      // The name of the Table
      // Values can be of the following formats:
      // - `projects/&lt;project&gt;/tables/&lt;table&gt;`
      // - `projects/&lt;project&gt;/instances/&lt;instance&gt;/tables/&lt;table&gt;`
      // - `region/&lt;region&gt;/zones/&lt;zone&gt;/tables/&lt;table&gt;`
      string table_name = 1;

      // This value specifies routing for replication.
      // It can be in the following formats:
      // - `profiles/&lt;profile_id&gt;`
      // - a legacy `profile_id` that can be any string
      string app_profile_id = 2;
    }

Example message:

    {
      table_name: projects/proj_foo/instances/instance_bar/table/table_baz,
      app_profile_id: profiles/prof_qux
    }

The routing header consists of one or multiple key-value pairs. Every key
and value must be percent-encoded, and joined together in the format of
`key1=value1&amp;key2=value2`.
In the examples below I am skipping the percent-encoding for readablity.

Example 1

Extracting a field from the request to put into the routing header
unchanged, with the key equal to the field name.

annotation:

    option (google.api.routing) = {
      // Take the `app_profile_id`.
      routing_parameters {
        field: &#34;app_profile_id&#34;
      }
    };

result:

    x-goog-request-params: app_profile_id=profiles/prof_qux

Example 2

Extracting a field from the request to put into the routing header
unchanged, with the key different from the field name.

annotation:

    option (google.api.routing) = {
      // Take the `app_profile_id`, but name it `routing_id` in the header.
      routing_parameters {
        field: &#34;app_profile_id&#34;
        path_template: &#34;{routing_id=**}&#34;
      }
    };

result:

    x-goog-request-params: routing_id=profiles/prof_qux

Example 3

Extracting a field from the request to put into the routing
header, while matching a path template syntax on the field&#39;s value.

NB: it is more useful to send nothing than to send garbage for the purpose
of dynamic routing, since garbage pollutes cache. Thus the matching.

Sub-example 3a

The field matches the template.

annotation:

    option (google.api.routing) = {
      // Take the `table_name`, if it&#39;s well-formed (with project-based
      // syntax).
      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{table_name=projects/*/instances/*/**}&#34;
      }
    };

result:

    x-goog-request-params:
    table_name=projects/proj_foo/instances/instance_bar/table/table_baz

Sub-example 3b

The field does not match the template.

annotation:

    option (google.api.routing) = {
      // Take the `table_name`, if it&#39;s well-formed (with region-based
      // syntax).
      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{table_name=regions/*/zones/*/**}&#34;
      }
    };

result:

    &lt;no routing header will be sent&gt;

Sub-example 3c

Multiple alternative conflictingly named path templates are
specified. The one that matches is used to construct the header.

annotation:

    option (google.api.routing) = {
      // Take the `table_name`, if it&#39;s well-formed, whether
      // using the region- or projects-based syntax.

      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{table_name=regions/*/zones/*/**}&#34;
      }
      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{table_name=projects/*/instances/*/**}&#34;
      }
    };

result:

    x-goog-request-params:
    table_name=projects/proj_foo/instances/instance_bar/table/table_baz

Example 4

Extracting a single routing header key-value pair by matching a
template syntax on (a part of) a single request field.

annotation:

    option (google.api.routing) = {
      // Take just the project id from the `table_name` field.
      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{routing_id=projects/*}/**&#34;
      }
    };

result:

    x-goog-request-params: routing_id=projects/proj_foo

Example 5

Extracting a single routing header key-value pair by matching
several conflictingly named path templates on (parts of) a single request
field. The last template to match &#34;wins&#34; the conflict.

annotation:

    option (google.api.routing) = {
      // If the `table_name` does not have instances information,
      // take just the project id for routing.
      // Otherwise take project &#43; instance.

      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{routing_id=projects/*}/**&#34;
      }
      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{routing_id=projects/*/instances/*}/**&#34;
      }
    };

result:

    x-goog-request-params:
    routing_id=projects/proj_foo/instances/instance_bar

Example 6

Extracting multiple routing header key-value pairs by matching
several non-conflicting path templates on (parts of) a single request field.

Sub-example 6a

Make the templates strict, so that if the `table_name` does not
have an instance information, nothing is sent.

annotation:

    option (google.api.routing) = {
      // The routing code needs two keys instead of one composite
      // but works only for the tables with the &#34;project-instance&#34; name
      // syntax.

      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{project_id=projects/*}/instances/*/**&#34;
      }
      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;projects/*/{instance_id=instances/*}/**&#34;
      }
    };

result:

    x-goog-request-params:
    project_id=projects/proj_foo&amp;instance_id=instances/instance_bar

Sub-example 6b

Make the templates loose, so that if the `table_name` does not
have an instance information, just the project id part is sent.

annotation:

    option (google.api.routing) = {
      // The routing code wants two keys instead of one composite
      // but will work with just the `project_id` for tables without
      // an instance in the `table_name`.

      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{project_id=projects/*}/**&#34;
      }
      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;projects/*/{instance_id=instances/*}/**&#34;
      }
    };

result (is the same as 6a for our example message because it has the instance
information):

    x-goog-request-params:
    project_id=projects/proj_foo&amp;instance_id=instances/instance_bar

Example 7

Extracting multiple routing header key-value pairs by matching
several path templates on multiple request fields.

NB: note that here there is no way to specify sending nothing if one of the
fields does not match its template. E.g. if the `table_name` is in the wrong
format, the `project_id` will not be sent, but the `routing_id` will be.
The backend routing code has to be aware of that and be prepared to not
receive a full complement of keys if it expects multiple.

annotation:

    option (google.api.routing) = {
      // The routing needs both `project_id` and `routing_id`
      // (from the `app_profile_id` field) for routing.

      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{project_id=projects/*}/**&#34;
      }
      routing_parameters {
        field: &#34;app_profile_id&#34;
        path_template: &#34;{routing_id=**}&#34;
      }
    };

result:

    x-goog-request-params:
    project_id=projects/proj_foo&amp;routing_id=profiles/prof_qux

Example 8

Extracting a single routing header key-value pair by matching
several conflictingly named path templates on several request fields. The
last template to match &#34;wins&#34; the conflict.

annotation:

    option (google.api.routing) = {
      // The `routing_id` can be a project id or a region id depending on
      // the table name format, but only if the `app_profile_id` is not set.
      // If `app_profile_id` is set it should be used instead.

      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{routing_id=projects/*}/**&#34;
      }
      routing_parameters {
         field: &#34;table_name&#34;
         path_template: &#34;{routing_id=regions/*}/**&#34;
      }
      routing_parameters {
        field: &#34;app_profile_id&#34;
        path_template: &#34;{routing_id=**}&#34;
      }
    };

result:

    x-goog-request-params: routing_id=profiles/prof_qux

Example 9

Bringing it all together.

annotation:

    option (google.api.routing) = {
      // For routing both `table_location` and a `routing_id` are needed.
      //
      // table_location can be either an instance id or a region&#43;zone id.
      //
      // For `routing_id`, take the value of `app_profile_id`
      // - If it&#39;s in the format `profiles/&lt;profile_id&gt;`, send
      // just the `&lt;profile_id&gt;` part.
      // - If it&#39;s any other literal, send it as is.
      // If the `app_profile_id` is empty, and the `table_name` starts with
      // the project_id, send that instead.

      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;projects/*/{table_location=instances/*}/tables/*&#34;
      }
      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{table_location=regions/*/zones/*}/tables/*&#34;
      }
      routing_parameters {
        field: &#34;table_name&#34;
        path_template: &#34;{routing_id=projects/*}/**&#34;
      }
      routing_parameters {
        field: &#34;app_profile_id&#34;
        path_template: &#34;{routing_id=**}&#34;
      }
      routing_parameters {
        field: &#34;app_profile_id&#34;
        path_template: &#34;profiles/{routing_id=*}&#34;
      }
    };

result:

    x-goog-request-params:
    table_location=instances/instance_bar&amp;routing_id=prof_qux


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| routing_parameters | [RoutingParameter](#google-api-RoutingParameter) | repeated | A collection of Routing Parameter specifications. **NOTE:** If multiple Routing Parameters describe the same key (via the `path_template` field or via the `field` field when `path_template` is not provided), &#34;last one wins&#34; rule determines which Parameter gets used. See the examples for more details. |





 

 


<a name="google_api_routing-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| routing | RoutingRule | .google.protobuf.MethodOptions | 72295729 | See RoutingRule. |

 

 



<a name="google_api_source_info-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/source_info.proto



<a name="google-api-SourceInfo"></a>

### SourceInfo
Source information used to create a Service Config


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_files | [google.protobuf.Any](#google-protobuf-Any) | repeated | All files used during config generation. |





 

 

 

 



<a name="google_api_system_parameter-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/system_parameter.proto



<a name="google-api-SystemParameter"></a>

### SystemParameter
Define a parameter&#39;s name and location. The parameter may be passed as either
an HTTP header or a URL query parameter, and if both are passed the behavior
is implementation-dependent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Define the name of the parameter, such as &#34;api_key&#34; . It is case sensitive. |
| http_header | [string](#string) |  | Define the HTTP header name to use for the parameter. It is case insensitive. |
| url_query_parameter | [string](#string) |  | Define the URL query parameter name to use for the parameter. It is case sensitive. |






<a name="google-api-SystemParameterRule"></a>

### SystemParameterRule
Define a system parameter rule mapping system parameter definitions to
methods.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | Selects the methods to which this rule applies. Use &#39;*&#39; to indicate all methods in all APIs.

Refer to [selector][google.api.DocumentationRule.selector] for syntax details. |
| parameters | [SystemParameter](#google-api-SystemParameter) | repeated | Define parameters. Multiple names may be defined for a parameter. For a given method call, only one of them should be used. If multiple names are used the behavior is implementation-dependent. If none of the specified names are present the behavior is parameter-dependent. |






<a name="google-api-SystemParameters"></a>

### SystemParameters
### System parameter configuration

A system parameter is a special kind of parameter defined by the API
system, not by an individual API. It is typically mapped to an HTTP header
and/or a URL query parameter. This configuration specifies which methods
change the names of the system parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rules | [SystemParameterRule](#google-api-SystemParameterRule) | repeated | Define system parameters.

The parameters defined here will override the default parameters implemented by the system. If this field is missing from the service config, default system parameters will be used. Default system parameters and names is implementation-dependent.

Example: define api key for all methods

 system_parameters rules: - selector: &#34;*&#34; parameters: - name: api_key url_query_parameter: api_key

Example: define 2 api key names for a specific method.

 system_parameters rules: - selector: &#34;/ListShelves&#34; parameters: - name: api_key http_header: Api-Key1 - name: api_key http_header: Api-Key2

**NOTE:** All service configuration rules follow &#34;last one wins&#34; order. |





 

 

 

 



<a name="google_api_usage-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/usage.proto



<a name="google-api-Usage"></a>

### Usage
Configuration controlling usage of a service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requirements | [string](#string) | repeated | Requirements that must be satisfied before a consumer project can use the service. Each requirement is of the form &lt;service.name&gt;/&lt;requirement-id&gt;; for example &#39;serviceusage.googleapis.com/billing-enabled&#39;.

For Google APIs, a Terms of Service requirement must be included here. Google Cloud APIs must include &#34;serviceusage.googleapis.com/tos/cloud&#34;. Other Google APIs should include &#34;serviceusage.googleapis.com/tos/universal&#34;. Additional ToS can be included based on the business needs. |
| rules | [UsageRule](#google-api-UsageRule) | repeated | A list of usage rules that apply to individual API methods.

**NOTE:** All service configuration rules follow &#34;last one wins&#34; order. |
| producer_notification_channel | [string](#string) |  | The full resource name of a channel used for sending notifications to the service producer.

Google Service Management currently only supports [Google Cloud Pub/Sub](https://cloud.google.com/pubsub) as a notification channel. To use Google Cloud Pub/Sub as the channel, this must be the name of a Cloud Pub/Sub topic that uses the Cloud Pub/Sub topic name format documented in https://cloud.google.com/pubsub/docs/overview. |






<a name="google-api-UsageRule"></a>

### UsageRule
Usage configuration rules for the service.

NOTE: Under development.


Use this rule to configure unregistered calls for the service. Unregistered
calls are calls that do not contain consumer project identity.
(Example: calls that do not contain an API key).
By default, API methods do not allow unregistered calls, and each method call
must be identified by a consumer project identity. Use this rule to
allow/disallow unregistered calls.

Example of an API that wants to allow unregistered calls for entire service.

    usage:
      rules:
      - selector: &#34;*&#34;
        allow_unregistered_calls: true

Example of a method that wants to allow unregistered calls.

    usage:
      rules:
      - selector: &#34;google.example.library.v1.LibraryService.CreateBook&#34;
        allow_unregistered_calls: true


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | Selects the methods to which this rule applies. Use &#39;*&#39; to indicate all methods in all APIs.

Refer to [selector][google.api.DocumentationRule.selector] for syntax details. |
| allow_unregistered_calls | [bool](#bool) |  | If true, the selected method allows unregistered calls, e.g. calls that don&#39;t identify any user or application. |
| skip_service_control | [bool](#bool) |  | If true, the selected method should skip service control and the control plane features, such as quota and billing, will not be available. This flag is used by Google Cloud Endpoints to bypass checks for internal methods, such as service health check methods. |





 

 

 

 



<a name="google_api_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/service.proto



<a name="google-api-Service"></a>

### Service
`Service` is the root object of Google API service configuration (service
config). It describes the basic information about a logical service,
such as the service name and the user-facing title, and delegates other
aspects to sub-sections. Each sub-section is either a proto message or a
repeated proto message that configures a specific aspect, such as auth.
For more information, see each proto message definition.

Example:

    type: google.api.Service
    name: calendar.googleapis.com
    title: Google Calendar API
    apis:
    - name: google.calendar.v3.Calendar

    visibility:
      rules:
      - selector: &#34;google.calendar.v3.*&#34;
        restriction: PREVIEW
    backend:
      rules:
      - selector: &#34;google.calendar.v3.*&#34;
        address: calendar.example.com

    authentication:
      providers:
      - id: google_calendar_auth
        jwks_uri: https://www.googleapis.com/oauth2/v1/certs
        issuer: https://securetoken.google.com
      rules:
      - selector: &#34;*&#34;
        requirements:
          provider_id: google_calendar_auth


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The service name, which is a DNS-like logical identifier for the service, such as `calendar.googleapis.com`. The service name typically goes through DNS verification to make sure the owner of the service also owns the DNS name. |
| title | [string](#string) |  | The product title for this service, it is the name displayed in Google Cloud Console. |
| producer_project_id | [string](#string) |  | The Google project that owns this service. |
| id | [string](#string) |  | A unique ID for a specific instance of this message, typically assigned by the client for tracking purpose. Must be no longer than 63 characters and only lower case letters, digits, &#39;.&#39;, &#39;_&#39; and &#39;-&#39; are allowed. If empty, the server may choose to generate one instead. |
| apis | [google.protobuf.Api](#google-protobuf-Api) | repeated | A list of API interfaces exported by this service. Only the `name` field of the [google.protobuf.Api][google.protobuf.Api] needs to be provided by the configuration author, as the remaining fields will be derived from the IDL during the normalization process. It is an error to specify an API interface here which cannot be resolved against the associated IDL files. |
| types | [google.protobuf.Type](#google-protobuf-Type) | repeated | A list of all proto message types included in this API service. Types referenced directly or indirectly by the `apis` are automatically included. Messages which are not referenced but shall be included, such as types used by the `google.protobuf.Any` type, should be listed here by name by the configuration author. Example:

 types: - name: google.protobuf.Int32 |
| enums | [google.protobuf.Enum](#google-protobuf-Enum) | repeated | A list of all enum types included in this API service. Enums referenced directly or indirectly by the `apis` are automatically included. Enums which are not referenced but shall be included should be listed here by name by the configuration author. Example:

 enums: - name: google.someapi.v1.SomeEnum |
| documentation | [Documentation](#google-api-Documentation) |  | Additional API documentation. |
| backend | [Backend](#google-api-Backend) |  | API backend configuration. |
| http | [Http](#google-api-Http) |  | HTTP configuration. |
| quota | [Quota](#google-api-Quota) |  | Quota configuration. |
| authentication | [Authentication](#google-api-Authentication) |  | Auth configuration. |
| context | [Context](#google-api-Context) |  | Context configuration. |
| usage | [Usage](#google-api-Usage) |  | Configuration controlling usage of this service. |
| endpoints | [Endpoint](#google-api-Endpoint) | repeated | Configuration for network endpoints. If this is empty, then an endpoint with the same name as the service is automatically generated to service all defined APIs. |
| control | [Control](#google-api-Control) |  | Configuration for the service control plane. |
| logs | [LogDescriptor](#google-api-LogDescriptor) | repeated | Defines the logs used by this service. |
| metrics | [MetricDescriptor](#google-api-MetricDescriptor) | repeated | Defines the metrics used by this service. |
| monitored_resources | [MonitoredResourceDescriptor](#google-api-MonitoredResourceDescriptor) | repeated | Defines the monitored resources used by this service. This is required by the [Service.monitoring][google.api.Service.monitoring] and [Service.logging][google.api.Service.logging] configurations. |
| billing | [Billing](#google-api-Billing) |  | Billing configuration. |
| logging | [Logging](#google-api-Logging) |  | Logging configuration. |
| monitoring | [Monitoring](#google-api-Monitoring) |  | Monitoring configuration. |
| system_parameters | [SystemParameters](#google-api-SystemParameters) |  | System parameter configuration. |
| source_info | [SourceInfo](#google-api-SourceInfo) |  | Output only. The source information for this configuration if available. |
| publishing | [Publishing](#google-api-Publishing) |  | Settings for [Google Cloud Client libraries](https://cloud.google.com/apis/docs/cloud-client-libraries) generated from APIs defined as protocol buffers. |
| config_version | [google.protobuf.UInt32Value](#google-protobuf-UInt32Value) |  | Obsolete. Do not use.

This field has no semantic meaning. The service config compiler always sets this field to `3`. |





 

 

 

 



<a name="google_api_visibility-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google/api/visibility.proto



<a name="google-api-Visibility"></a>

### Visibility
`Visibility` restricts service consumer&#39;s access to service elements,
such as whether an application can call a visibility-restricted method.
The restriction is expressed by applying visibility labels on service
elements. The visibility labels are elsewhere linked to service consumers.

A service can define multiple visibility labels, but a service consumer
should be granted at most one visibility label. Multiple visibility
labels for a single service consumer are not supported.

If an element and all its parents have no visibility label, its visibility
is unconditionally granted.

Example:

    visibility:
      rules:
      - selector: google.calendar.Calendar.EnhancedSearch
        restriction: PREVIEW
      - selector: google.calendar.Calendar.Delegate
        restriction: INTERNAL

Here, all methods are publicly visible except for the restricted methods
EnhancedSearch and Delegate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rules | [VisibilityRule](#google-api-VisibilityRule) | repeated | A list of visibility rules that apply to individual API elements.

**NOTE:** All service configuration rules follow &#34;last one wins&#34; order. |






<a name="google-api-VisibilityRule"></a>

### VisibilityRule
A visibility rule provides visibility configuration for an individual API
element.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [string](#string) |  | Selects methods, messages, fields, enums, etc. to which this rule applies.

Refer to [selector][google.api.DocumentationRule.selector] for syntax details. |
| restriction | [string](#string) |  | A comma-separated list of visibility labels that apply to the `selector`. Any of the listed labels can be used to grant the visibility.

If a rule has multiple labels, removing one of the labels but not all of them can break clients.

Example:

 visibility: rules: - selector: google.calendar.Calendar.EnhancedSearch restriction: INTERNAL, PREVIEW

Removing INTERNAL from this restriction will break clients that rely on this method and only had access to it through INTERNAL. |





 

 


<a name="google_api_visibility-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| enum_visibility | VisibilityRule | .google.protobuf.EnumOptions | 72295727 | See `VisibilityRule`. |
| value_visibility | VisibilityRule | .google.protobuf.EnumValueOptions | 72295727 | See `VisibilityRule`. |
| field_visibility | VisibilityRule | .google.protobuf.FieldOptions | 72295727 | See `VisibilityRule`. |
| message_visibility | VisibilityRule | .google.protobuf.MessageOptions | 72295727 | See `VisibilityRule`. |
| method_visibility | VisibilityRule | .google.protobuf.MethodOptions | 72295727 | See `VisibilityRule`. |
| api_visibility | VisibilityRule | .google.protobuf.ServiceOptions | 72295727 | See `VisibilityRule`. |

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

