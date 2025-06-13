# V1ActivityLogsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorRef** | Pointer to **string** |  | [optional] 
**NumLogs** | Pointer to **int32** |  | [optional] 
**OldTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**RecentTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Selector** | Pointer to [**V1ActivityLogsPostRequestSelector**](V1ActivityLogsPostRequestSelector.md) |  | [optional] 
**SelectorV2** | Pointer to [**V1ActivityLogsPostRequestSelectorV2**](V1ActivityLogsPostRequestSelectorV2.md) |  | [optional] 

## Methods

### NewV1ActivityLogsPostRequest

`func NewV1ActivityLogsPostRequest() *V1ActivityLogsPostRequest`

NewV1ActivityLogsPostRequest instantiates a new V1ActivityLogsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ActivityLogsPostRequestWithDefaults

`func NewV1ActivityLogsPostRequestWithDefaults() *V1ActivityLogsPostRequest`

NewV1ActivityLogsPostRequestWithDefaults instantiates a new V1ActivityLogsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorRef

`func (o *V1ActivityLogsPostRequest) GetCursorRef() string`

GetCursorRef returns the CursorRef field if non-nil, zero value otherwise.

### GetCursorRefOk

`func (o *V1ActivityLogsPostRequest) GetCursorRefOk() (*string, bool)`

GetCursorRefOk returns a tuple with the CursorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorRef

`func (o *V1ActivityLogsPostRequest) SetCursorRef(v string)`

SetCursorRef sets CursorRef field to given value.

### HasCursorRef

`func (o *V1ActivityLogsPostRequest) HasCursorRef() bool`

HasCursorRef returns a boolean if a field has been set.

### GetNumLogs

`func (o *V1ActivityLogsPostRequest) GetNumLogs() int32`

GetNumLogs returns the NumLogs field if non-nil, zero value otherwise.

### GetNumLogsOk

`func (o *V1ActivityLogsPostRequest) GetNumLogsOk() (*int32, bool)`

GetNumLogsOk returns a tuple with the NumLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLogs

`func (o *V1ActivityLogsPostRequest) SetNumLogs(v int32)`

SetNumLogs sets NumLogs field to given value.

### HasNumLogs

`func (o *V1ActivityLogsPostRequest) HasNumLogs() bool`

HasNumLogs returns a boolean if a field has been set.

### GetOldTs

`func (o *V1ActivityLogsPostRequest) GetOldTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetOldTs returns the OldTs field if non-nil, zero value otherwise.

### GetOldTsOk

`func (o *V1ActivityLogsPostRequest) GetOldTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetOldTsOk returns a tuple with the OldTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTs

`func (o *V1ActivityLogsPostRequest) SetOldTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetOldTs sets OldTs field to given value.

### HasOldTs

`func (o *V1ActivityLogsPostRequest) HasOldTs() bool`

HasOldTs returns a boolean if a field has been set.

### GetRecentTs

`func (o *V1ActivityLogsPostRequest) GetRecentTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetRecentTs returns the RecentTs field if non-nil, zero value otherwise.

### GetRecentTsOk

`func (o *V1ActivityLogsPostRequest) GetRecentTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetRecentTsOk returns a tuple with the RecentTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTs

`func (o *V1ActivityLogsPostRequest) SetRecentTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetRecentTs sets RecentTs field to given value.

### HasRecentTs

`func (o *V1ActivityLogsPostRequest) HasRecentTs() bool`

HasRecentTs returns a boolean if a field has been set.

### GetSelector

`func (o *V1ActivityLogsPostRequest) GetSelector() V1ActivityLogsPostRequestSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1ActivityLogsPostRequest) GetSelectorOk() (*V1ActivityLogsPostRequestSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1ActivityLogsPostRequest) SetSelector(v V1ActivityLogsPostRequestSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1ActivityLogsPostRequest) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSelectorV2

`func (o *V1ActivityLogsPostRequest) GetSelectorV2() V1ActivityLogsPostRequestSelectorV2`

GetSelectorV2 returns the SelectorV2 field if non-nil, zero value otherwise.

### GetSelectorV2Ok

`func (o *V1ActivityLogsPostRequest) GetSelectorV2Ok() (*V1ActivityLogsPostRequestSelectorV2, bool)`

GetSelectorV2Ok returns a tuple with the SelectorV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorV2

`func (o *V1ActivityLogsPostRequest) SetSelectorV2(v V1ActivityLogsPostRequestSelectorV2)`

SetSelectorV2 sets SelectorV2 field to given value.

### HasSelectorV2

`func (o *V1ActivityLogsPostRequest) HasSelectorV2() bool`

HasSelectorV2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


