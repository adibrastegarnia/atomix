# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [runtime/map/v1/maps.proto](#runtime_map_v1_maps-proto)
    - [CloseRequest](#atomix-runtime-map-v1-CloseRequest)
    - [CloseResponse](#atomix-runtime-map-v1-CloseResponse)
    - [CreateRequest](#atomix-runtime-map-v1-CreateRequest)
    - [CreateResponse](#atomix-runtime-map-v1-CreateResponse)
    - [MapCacheConfig](#atomix-runtime-map-v1-MapCacheConfig)
    - [MapConfig](#atomix-runtime-map-v1-MapConfig)
  
    - [Maps](#atomix-runtime-map-v1-Maps)
  
- [Scalar Value Types](#scalar-value-types)



<a name="runtime_map_v1_maps-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## runtime/map/v1/maps.proto



<a name="atomix-runtime-map-v1-CloseRequest"></a>

### CloseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [atomix.runtime.v1.PrimitiveID](#atomix-runtime-v1-PrimitiveID) |  |  |






<a name="atomix-runtime-map-v1-CloseResponse"></a>

### CloseResponse







<a name="atomix-runtime-map-v1-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [atomix.runtime.v1.PrimitiveID](#atomix-runtime-v1-PrimitiveID) |  |  |
| tags | [string](#string) | repeated |  |






<a name="atomix-runtime-map-v1-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [MapConfig](#atomix-runtime-map-v1-MapConfig) |  |  |






<a name="atomix-runtime-map-v1-MapCacheConfig"></a>

### MapCacheConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |
| size | [uint64](#uint64) |  |  |






<a name="atomix-runtime-map-v1-MapConfig"></a>

### MapConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cache | [MapCacheConfig](#atomix-runtime-map-v1-MapCacheConfig) |  |  |





 

 

 


<a name="atomix-runtime-map-v1-Maps"></a>

### Maps
Maps is a service for managing map primitives

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix-runtime-map-v1-CreateRequest) | [CreateResponse](#atomix-runtime-map-v1-CreateResponse) | Create creates the map |
| Close | [CloseRequest](#atomix-runtime-map-v1-CloseRequest) | [CloseResponse](#atomix-runtime-map-v1-CloseResponse) | Close closes the map |

 



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

