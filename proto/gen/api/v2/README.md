# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v2/common.proto](#api_v2_common-proto)
    - [RowStatus](#mercury-api-v2-RowStatus)
  
- [api/v2/message_service.proto](#api_v2_message_service-proto)
    - [CreateMessageCommentRequest](#mercury-api-v2-CreateMessageCommentRequest)
    - [CreateMessageCommentResponse](#mercury-api-v2-CreateMessageCommentResponse)
    - [CreateMessageRequest](#mercury-api-v2-CreateMessageRequest)
    - [CreateMessageResponse](#mercury-api-v2-CreateMessageResponse)
    - [GetMessageRequest](#mercury-api-v2-GetMessageRequest)
    - [GetMessageResponse](#mercury-api-v2-GetMessageResponse)
    - [ListMessageCommentsRequest](#mercury-api-v2-ListMessageCommentsRequest)
    - [ListMessageCommentsResponse](#mercury-api-v2-ListMessageCommentsResponse)
    - [ListMessagesRequest](#mercury-api-v2-ListMessagesRequest)
    - [ListMessagesResponse](#mercury-api-v2-ListMessagesResponse)
    - [Message](#mercury-api-v2-Message)
  
    - [Visibility](#mercury-api-v2-Visibility)
  
    - [MessageService](#mercury-api-v2-MessageService)
  
- [api/v2/resource_service.proto](#api_v2_resource_service-proto)
    - [CreateResourceRequest](#mercury-api-v2-CreateResourceRequest)
    - [CreateResourceResponse](#mercury-api-v2-CreateResourceResponse)
    - [DeleteResourceRequest](#mercury-api-v2-DeleteResourceRequest)
    - [DeleteResourceResponse](#mercury-api-v2-DeleteResourceResponse)
    - [ListResourcesRequest](#mercury-api-v2-ListResourcesRequest)
    - [ListResourcesResponse](#mercury-api-v2-ListResourcesResponse)
    - [Resource](#mercury-api-v2-Resource)
    - [UpdateResourceRequest](#mercury-api-v2-UpdateResourceRequest)
    - [UpdateResourceResponse](#mercury-api-v2-UpdateResourceResponse)
  
    - [ResourceService](#mercury-api-v2-ResourceService)
  
- [api/v2/system_service.proto](#api_v2_system_service-proto)
    - [GetSystemInfoRequest](#mercury-api-v2-GetSystemInfoRequest)
    - [GetSystemInfoResponse](#mercury-api-v2-GetSystemInfoResponse)
    - [SystemInfo](#mercury-api-v2-SystemInfo)
    - [UpdateSystemInfoRequest](#mercury-api-v2-UpdateSystemInfoRequest)
    - [UpdateSystemInfoResponse](#mercury-api-v2-UpdateSystemInfoResponse)
  
    - [SystemService](#mercury-api-v2-SystemService)
  
- [api/v2/tag_service.proto](#api_v2_tag_service-proto)
    - [DeleteTagRequest](#mercury-api-v2-DeleteTagRequest)
    - [DeleteTagResponse](#mercury-api-v2-DeleteTagResponse)
    - [ListTagsRequest](#mercury-api-v2-ListTagsRequest)
    - [ListTagsResponse](#mercury-api-v2-ListTagsResponse)
    - [Tag](#mercury-api-v2-Tag)
    - [UpsertTagRequest](#mercury-api-v2-UpsertTagRequest)
    - [UpsertTagResponse](#mercury-api-v2-UpsertTagResponse)
  
    - [TagService](#mercury-api-v2-TagService)
  
- [api/v2/user_service.proto](#api_v2_user_service-proto)
    - [CreateUserAccessTokenRequest](#mercury-api-v2-CreateUserAccessTokenRequest)
    - [CreateUserAccessTokenResponse](#mercury-api-v2-CreateUserAccessTokenResponse)
    - [DeleteUserAccessTokenRequest](#mercury-api-v2-DeleteUserAccessTokenRequest)
    - [DeleteUserAccessTokenResponse](#mercury-api-v2-DeleteUserAccessTokenResponse)
    - [GetUserRequest](#mercury-api-v2-GetUserRequest)
    - [GetUserResponse](#mercury-api-v2-GetUserResponse)
    - [ListUserAccessTokensRequest](#mercury-api-v2-ListUserAccessTokensRequest)
    - [ListUserAccessTokensResponse](#mercury-api-v2-ListUserAccessTokensResponse)
    - [UpdateUserRequest](#mercury-api-v2-UpdateUserRequest)
    - [UpdateUserResponse](#mercury-api-v2-UpdateUserResponse)
    - [User](#mercury-api-v2-User)
    - [UserAccessToken](#mercury-api-v2-UserAccessToken)
  
    - [User.Role](#mercury-api-v2-User-Role)
  
    - [UserService](#mercury-api-v2-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v2_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/common.proto


 


<a name="mercury-api-v2-RowStatus"></a>

### RowStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| ROW_STATUS_UNSPECIFIED | 0 |  |
| ACTIVE | 1 |  |
| ARCHIVED | 2 |  |


 

 

 



<a name="api_v2_message_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/message_service.proto



<a name="mercury-api-v2-CreateMessageCommentRequest"></a>

### CreateMessageCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | id to create comment for. |
| create | [CreateMessageRequest](#mercury-api-v2-CreateMessageRequest) |  |  |






<a name="mercury-api-v2-CreateMessageCommentResponse"></a>

### CreateMessageCommentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [Message](#mercury-api-v2-Message) |  |  |






<a name="mercury-api-v2-CreateMessageRequest"></a>

### CreateMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [string](#string) |  |  |
| visibility | [Visibility](#mercury-api-v2-Visibility) |  |  |






<a name="mercury-api-v2-CreateMessageResponse"></a>

### CreateMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [Message](#mercury-api-v2-Message) |  |  |






<a name="mercury-api-v2-GetMessageRequest"></a>

### GetMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="mercury-api-v2-GetMessageResponse"></a>

### GetMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [Message](#mercury-api-v2-Message) |  |  |






<a name="mercury-api-v2-ListMessageCommentsRequest"></a>

### ListMessageCommentsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="mercury-api-v2-ListMessageCommentsResponse"></a>

### ListMessageCommentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| messages | [Message](#mercury-api-v2-Message) | repeated |  |






<a name="mercury-api-v2-ListMessagesRequest"></a>

### ListMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| page_size | [int32](#int32) |  |  |
| filter | [string](#string) |  |  |






<a name="mercury-api-v2-ListMessagesResponse"></a>

### ListMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| messages | [Message](#mercury-api-v2-Message) | repeated |  |






<a name="mercury-api-v2-Message"></a>

### Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| row_status | [RowStatus](#mercury-api-v2-RowStatus) |  |  |
| creator_id | [int32](#int32) |  |  |
| creator_timestamp | [int64](#int64) |  |  |
| updated_timestamp | [int64](#int64) |  |  |
| content | [string](#string) |  |  |
| visibility | [Visibility](#mercury-api-v2-Visibility) |  |  |
| pinned | [bool](#bool) |  |  |





 


<a name="mercury-api-v2-Visibility"></a>

### Visibility


| Name | Number | Description |
| ---- | ------ | ----------- |
| VISIBILITY_UNSPECIFIED | 0 |  |
| PRIVATE | 1 |  |
| PROTECTED | 2 |  |
| PUBLIC | 3 |  |


 

 


<a name="mercury-api-v2-MessageService"></a>

### MessageService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateMessage | [CreateMessageRequest](#mercury-api-v2-CreateMessageRequest) | [CreateMessageResponse](#mercury-api-v2-CreateMessageResponse) |  |
| ListMessages | [ListMessagesRequest](#mercury-api-v2-ListMessagesRequest) | [ListMessagesResponse](#mercury-api-v2-ListMessagesResponse) |  |
| GetMessage | [GetMessageRequest](#mercury-api-v2-GetMessageRequest) | [GetMessageResponse](#mercury-api-v2-GetMessageResponse) |  |
| CreateMessageComment | [CreateMessageCommentRequest](#mercury-api-v2-CreateMessageCommentRequest) | [CreateMessageCommentResponse](#mercury-api-v2-CreateMessageCommentResponse) |  |
| ListMessageComments | [ListMessageCommentsRequest](#mercury-api-v2-ListMessageCommentsRequest) | [ListMessageCommentsResponse](#mercury-api-v2-ListMessageCommentsResponse) |  |

 



<a name="api_v2_resource_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/resource_service.proto



<a name="mercury-api-v2-CreateResourceRequest"></a>

### CreateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filename | [string](#string) |  |  |
| external_link | [string](#string) |  |  |
| type | [string](#string) |  |  |
| message_id | [int32](#int32) | optional |  |






<a name="mercury-api-v2-CreateResourceResponse"></a>

### CreateResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [Resource](#mercury-api-v2-Resource) |  |  |






<a name="mercury-api-v2-DeleteResourceRequest"></a>

### DeleteResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="mercury-api-v2-DeleteResourceResponse"></a>

### DeleteResourceResponse







<a name="mercury-api-v2-ListResourcesRequest"></a>

### ListResourcesRequest







<a name="mercury-api-v2-ListResourcesResponse"></a>

### ListResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [Resource](#mercury-api-v2-Resource) | repeated |  |






<a name="mercury-api-v2-Resource"></a>

### Resource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| created_timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| filename | [string](#string) |  |  |
| external_link | [string](#string) |  |  |
| type | [string](#string) |  |  |
| size | [int64](#int64) |  |  |
| related_message_id | [int32](#int32) | optional |  |






<a name="mercury-api-v2-UpdateResourceRequest"></a>

### UpdateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| resource | [Resource](#mercury-api-v2-Resource) |  |  |
| update_mask | [string](#string) | repeated |  |






<a name="mercury-api-v2-UpdateResourceResponse"></a>

### UpdateResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [Resource](#mercury-api-v2-Resource) |  |  |





 

 

 


<a name="mercury-api-v2-ResourceService"></a>

### ResourceService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListResources | [ListResourcesRequest](#mercury-api-v2-ListResourcesRequest) | [ListResourcesResponse](#mercury-api-v2-ListResourcesResponse) |  |
| CreateResource | [CreateResourceRequest](#mercury-api-v2-CreateResourceRequest) | [CreateResourceResponse](#mercury-api-v2-CreateResourceResponse) |  |
| UpdateResource | [UpdateResourceRequest](#mercury-api-v2-UpdateResourceRequest) | [UpdateResourceResponse](#mercury-api-v2-UpdateResourceResponse) |  |
| DeleteResource | [DeleteResourceRequest](#mercury-api-v2-DeleteResourceRequest) | [DeleteResourceResponse](#mercury-api-v2-DeleteResourceResponse) |  |

 



<a name="api_v2_system_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/system_service.proto



<a name="mercury-api-v2-GetSystemInfoRequest"></a>

### GetSystemInfoRequest







<a name="mercury-api-v2-GetSystemInfoResponse"></a>

### GetSystemInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| system_info | [SystemInfo](#mercury-api-v2-SystemInfo) |  |  |






<a name="mercury-api-v2-SystemInfo"></a>

### SystemInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |
| mode | [string](#string) |  |  |
| allow_registration | [bool](#bool) |  |  |
| disable_password_login | [bool](#bool) |  |  |
| additional_script | [string](#string) |  |  |
| additional_style | [string](#string) |  |  |
| db_size | [int64](#int64) |  |  |






<a name="mercury-api-v2-UpdateSystemInfoRequest"></a>

### UpdateSystemInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| system_info | [SystemInfo](#mercury-api-v2-SystemInfo) |  |  |
| update_mask | [string](#string) | repeated |  |






<a name="mercury-api-v2-UpdateSystemInfoResponse"></a>

### UpdateSystemInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| system_info | [SystemInfo](#mercury-api-v2-SystemInfo) |  |  |





 

 

 


<a name="mercury-api-v2-SystemService"></a>

### SystemService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSystemInfo | [GetSystemInfoRequest](#mercury-api-v2-GetSystemInfoRequest) | [GetSystemInfoResponse](#mercury-api-v2-GetSystemInfoResponse) |  |
| UpdateSystemInfo | [UpdateSystemInfoRequest](#mercury-api-v2-UpdateSystemInfoRequest) | [UpdateSystemInfoResponse](#mercury-api-v2-UpdateSystemInfoResponse) |  |

 



<a name="api_v2_tag_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/tag_service.proto



<a name="mercury-api-v2-DeleteTagRequest"></a>

### DeleteTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [Tag](#mercury-api-v2-Tag) |  |  |






<a name="mercury-api-v2-DeleteTagResponse"></a>

### DeleteTagResponse







<a name="mercury-api-v2-ListTagsRequest"></a>

### ListTagsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_id | [int32](#int32) |  |  |






<a name="mercury-api-v2-ListTagsResponse"></a>

### ListTagsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tags | [Tag](#mercury-api-v2-Tag) | repeated |  |






<a name="mercury-api-v2-Tag"></a>

### Tag



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| creator_id | [int32](#int32) |  |  |






<a name="mercury-api-v2-UpsertTagRequest"></a>

### UpsertTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="mercury-api-v2-UpsertTagResponse"></a>

### UpsertTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [Tag](#mercury-api-v2-Tag) |  |  |





 

 

 


<a name="mercury-api-v2-TagService"></a>

### TagService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListTags | [ListTagsRequest](#mercury-api-v2-ListTagsRequest) | [ListTagsResponse](#mercury-api-v2-ListTagsResponse) |  |
| UpsertTag | [UpsertTagRequest](#mercury-api-v2-UpsertTagRequest) | [UpsertTagResponse](#mercury-api-v2-UpsertTagResponse) |  |
| DeleteTag | [DeleteTagRequest](#mercury-api-v2-DeleteTagRequest) | [DeleteTagResponse](#mercury-api-v2-DeleteTagResponse) |  |

 



<a name="api_v2_user_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/user_service.proto



<a name="mercury-api-v2-CreateUserAccessTokenRequest"></a>

### CreateUserAccessTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| description | [string](#string) |  |  |
| expires_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional |  |






<a name="mercury-api-v2-CreateUserAccessTokenResponse"></a>

### CreateUserAccessTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [UserAccessToken](#mercury-api-v2-UserAccessToken) |  |  |






<a name="mercury-api-v2-DeleteUserAccessTokenRequest"></a>

### DeleteUserAccessTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| access_token | [string](#string) |  |  |






<a name="mercury-api-v2-DeleteUserAccessTokenResponse"></a>

### DeleteUserAccessTokenResponse







<a name="mercury-api-v2-GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |






<a name="mercury-api-v2-GetUserResponse"></a>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#mercury-api-v2-User) |  |  |






<a name="mercury-api-v2-ListUserAccessTokensRequest"></a>

### ListUserAccessTokensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |






<a name="mercury-api-v2-ListUserAccessTokensResponse"></a>

### ListUserAccessTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_tokens | [UserAccessToken](#mercury-api-v2-UserAccessToken) | repeated |  |






<a name="mercury-api-v2-UpdateUserRequest"></a>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| user | [User](#mercury-api-v2-User) |  |  |
| update_mask | [string](#string) | repeated |  |






<a name="mercury-api-v2-UpdateUserResponse"></a>

### UpdateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#mercury-api-v2-User) |  |  |






<a name="mercury-api-v2-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| username | [string](#string) |  |  |
| role | [User.Role](#mercury-api-v2-User-Role) |  |  |
| email | [string](#string) |  |  |
| nickname | [string](#string) |  |  |
| avatar_url | [string](#string) |  |  |
| password | [string](#string) |  |  |
| row_status | [RowStatus](#mercury-api-v2-RowStatus) |  |  |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| update_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="mercury-api-v2-UserAccessToken"></a>

### UserAccessToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |
| description | [string](#string) |  |  |
| issued_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| expires_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 


<a name="mercury-api-v2-User-Role"></a>

### User.Role


| Name | Number | Description |
| ---- | ------ | ----------- |
| ROLE_UNSPECIFIED | 0 |  |
| HOST | 1 |  |
| ADMIN | 2 |  |
| USER | 3 |  |


 

 


<a name="mercury-api-v2-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUser | [GetUserRequest](#mercury-api-v2-GetUserRequest) | [GetUserResponse](#mercury-api-v2-GetUserResponse) |  |
| UpdateUser | [UpdateUserRequest](#mercury-api-v2-UpdateUserRequest) | [UpdateUserResponse](#mercury-api-v2-UpdateUserResponse) |  |
| ListUserAccessTokens | [ListUserAccessTokensRequest](#mercury-api-v2-ListUserAccessTokensRequest) | [ListUserAccessTokensResponse](#mercury-api-v2-ListUserAccessTokensResponse) |  |
| CreateUserAccessToken | [CreateUserAccessTokenRequest](#mercury-api-v2-CreateUserAccessTokenRequest) | [CreateUserAccessTokenResponse](#mercury-api-v2-CreateUserAccessTokenResponse) |  |
| DeleteUserAccessToken | [DeleteUserAccessTokenRequest](#mercury-api-v2-DeleteUserAccessTokenRequest) | [DeleteUserAccessTokenResponse](#mercury-api-v2-DeleteUserAccessTokenResponse) |  |

 



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

