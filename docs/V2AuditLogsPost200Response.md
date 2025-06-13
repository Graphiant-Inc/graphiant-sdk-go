# V2AuditLogsPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorRef** | Pointer to **string** |  | [optional] 
**Logs** | Pointer to [**[]V2AuditLogsPost200ResponseLogsInner**](V2AuditLogsPost200ResponseLogsInner.md) |  | [optional] 
**TotalLogs** | Pointer to **int64** |  | [optional] 

## Methods

### NewV2AuditLogsPost200Response

`func NewV2AuditLogsPost200Response() *V2AuditLogsPost200Response`

NewV2AuditLogsPost200Response instantiates a new V2AuditLogsPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AuditLogsPost200ResponseWithDefaults

`func NewV2AuditLogsPost200ResponseWithDefaults() *V2AuditLogsPost200Response`

NewV2AuditLogsPost200ResponseWithDefaults instantiates a new V2AuditLogsPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorRef

`func (o *V2AuditLogsPost200Response) GetCursorRef() string`

GetCursorRef returns the CursorRef field if non-nil, zero value otherwise.

### GetCursorRefOk

`func (o *V2AuditLogsPost200Response) GetCursorRefOk() (*string, bool)`

GetCursorRefOk returns a tuple with the CursorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorRef

`func (o *V2AuditLogsPost200Response) SetCursorRef(v string)`

SetCursorRef sets CursorRef field to given value.

### HasCursorRef

`func (o *V2AuditLogsPost200Response) HasCursorRef() bool`

HasCursorRef returns a boolean if a field has been set.

### GetLogs

`func (o *V2AuditLogsPost200Response) GetLogs() []V2AuditLogsPost200ResponseLogsInner`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *V2AuditLogsPost200Response) GetLogsOk() (*[]V2AuditLogsPost200ResponseLogsInner, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *V2AuditLogsPost200Response) SetLogs(v []V2AuditLogsPost200ResponseLogsInner)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *V2AuditLogsPost200Response) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetTotalLogs

`func (o *V2AuditLogsPost200Response) GetTotalLogs() int64`

GetTotalLogs returns the TotalLogs field if non-nil, zero value otherwise.

### GetTotalLogsOk

`func (o *V2AuditLogsPost200Response) GetTotalLogsOk() (*int64, bool)`

GetTotalLogsOk returns a tuple with the TotalLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogs

`func (o *V2AuditLogsPost200Response) SetTotalLogs(v int64)`

SetTotalLogs sets TotalLogs field to given value.

### HasTotalLogs

`func (o *V2AuditLogsPost200Response) HasTotalLogs() bool`

HasTotalLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


