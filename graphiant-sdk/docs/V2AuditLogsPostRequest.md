# V2AuditLogsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorRef** | Pointer to **string** |  | [optional] 
**NumLogs** | Pointer to **int32** |  | [optional] 
**OldTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**RecentTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Selector** | Pointer to [**V2AuditLogsPostRequestSelector**](V2AuditLogsPostRequestSelector.md) |  | [optional] 

## Methods

### NewV2AuditLogsPostRequest

`func NewV2AuditLogsPostRequest() *V2AuditLogsPostRequest`

NewV2AuditLogsPostRequest instantiates a new V2AuditLogsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AuditLogsPostRequestWithDefaults

`func NewV2AuditLogsPostRequestWithDefaults() *V2AuditLogsPostRequest`

NewV2AuditLogsPostRequestWithDefaults instantiates a new V2AuditLogsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorRef

`func (o *V2AuditLogsPostRequest) GetCursorRef() string`

GetCursorRef returns the CursorRef field if non-nil, zero value otherwise.

### GetCursorRefOk

`func (o *V2AuditLogsPostRequest) GetCursorRefOk() (*string, bool)`

GetCursorRefOk returns a tuple with the CursorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorRef

`func (o *V2AuditLogsPostRequest) SetCursorRef(v string)`

SetCursorRef sets CursorRef field to given value.

### HasCursorRef

`func (o *V2AuditLogsPostRequest) HasCursorRef() bool`

HasCursorRef returns a boolean if a field has been set.

### GetNumLogs

`func (o *V2AuditLogsPostRequest) GetNumLogs() int32`

GetNumLogs returns the NumLogs field if non-nil, zero value otherwise.

### GetNumLogsOk

`func (o *V2AuditLogsPostRequest) GetNumLogsOk() (*int32, bool)`

GetNumLogsOk returns a tuple with the NumLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLogs

`func (o *V2AuditLogsPostRequest) SetNumLogs(v int32)`

SetNumLogs sets NumLogs field to given value.

### HasNumLogs

`func (o *V2AuditLogsPostRequest) HasNumLogs() bool`

HasNumLogs returns a boolean if a field has been set.

### GetOldTs

`func (o *V2AuditLogsPostRequest) GetOldTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetOldTs returns the OldTs field if non-nil, zero value otherwise.

### GetOldTsOk

`func (o *V2AuditLogsPostRequest) GetOldTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetOldTsOk returns a tuple with the OldTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTs

`func (o *V2AuditLogsPostRequest) SetOldTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetOldTs sets OldTs field to given value.

### HasOldTs

`func (o *V2AuditLogsPostRequest) HasOldTs() bool`

HasOldTs returns a boolean if a field has been set.

### GetRecentTs

`func (o *V2AuditLogsPostRequest) GetRecentTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetRecentTs returns the RecentTs field if non-nil, zero value otherwise.

### GetRecentTsOk

`func (o *V2AuditLogsPostRequest) GetRecentTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetRecentTsOk returns a tuple with the RecentTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTs

`func (o *V2AuditLogsPostRequest) SetRecentTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetRecentTs sets RecentTs field to given value.

### HasRecentTs

`func (o *V2AuditLogsPostRequest) HasRecentTs() bool`

HasRecentTs returns a boolean if a field has been set.

### GetSelector

`func (o *V2AuditLogsPostRequest) GetSelector() V2AuditLogsPostRequestSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V2AuditLogsPostRequest) GetSelectorOk() (*V2AuditLogsPostRequestSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V2AuditLogsPostRequest) SetSelector(v V2AuditLogsPostRequestSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V2AuditLogsPostRequest) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


