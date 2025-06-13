# V1AuditLogsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorRef** | Pointer to **string** |  | [optional] 
**HistogramBucketSizeSec** | Pointer to **int32** |  | [optional] 
**NumLogs** | Pointer to **int32** |  | [optional] 
**OldTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**RecentTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Selectors** | Pointer to [**[]V1AuditLogsPostRequestSelectorsInner**](V1AuditLogsPostRequestSelectorsInner.md) |  | [optional] 

## Methods

### NewV1AuditLogsPostRequest

`func NewV1AuditLogsPostRequest() *V1AuditLogsPostRequest`

NewV1AuditLogsPostRequest instantiates a new V1AuditLogsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AuditLogsPostRequestWithDefaults

`func NewV1AuditLogsPostRequestWithDefaults() *V1AuditLogsPostRequest`

NewV1AuditLogsPostRequestWithDefaults instantiates a new V1AuditLogsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorRef

`func (o *V1AuditLogsPostRequest) GetCursorRef() string`

GetCursorRef returns the CursorRef field if non-nil, zero value otherwise.

### GetCursorRefOk

`func (o *V1AuditLogsPostRequest) GetCursorRefOk() (*string, bool)`

GetCursorRefOk returns a tuple with the CursorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorRef

`func (o *V1AuditLogsPostRequest) SetCursorRef(v string)`

SetCursorRef sets CursorRef field to given value.

### HasCursorRef

`func (o *V1AuditLogsPostRequest) HasCursorRef() bool`

HasCursorRef returns a boolean if a field has been set.

### GetHistogramBucketSizeSec

`func (o *V1AuditLogsPostRequest) GetHistogramBucketSizeSec() int32`

GetHistogramBucketSizeSec returns the HistogramBucketSizeSec field if non-nil, zero value otherwise.

### GetHistogramBucketSizeSecOk

`func (o *V1AuditLogsPostRequest) GetHistogramBucketSizeSecOk() (*int32, bool)`

GetHistogramBucketSizeSecOk returns a tuple with the HistogramBucketSizeSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramBucketSizeSec

`func (o *V1AuditLogsPostRequest) SetHistogramBucketSizeSec(v int32)`

SetHistogramBucketSizeSec sets HistogramBucketSizeSec field to given value.

### HasHistogramBucketSizeSec

`func (o *V1AuditLogsPostRequest) HasHistogramBucketSizeSec() bool`

HasHistogramBucketSizeSec returns a boolean if a field has been set.

### GetNumLogs

`func (o *V1AuditLogsPostRequest) GetNumLogs() int32`

GetNumLogs returns the NumLogs field if non-nil, zero value otherwise.

### GetNumLogsOk

`func (o *V1AuditLogsPostRequest) GetNumLogsOk() (*int32, bool)`

GetNumLogsOk returns a tuple with the NumLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLogs

`func (o *V1AuditLogsPostRequest) SetNumLogs(v int32)`

SetNumLogs sets NumLogs field to given value.

### HasNumLogs

`func (o *V1AuditLogsPostRequest) HasNumLogs() bool`

HasNumLogs returns a boolean if a field has been set.

### GetOldTs

`func (o *V1AuditLogsPostRequest) GetOldTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetOldTs returns the OldTs field if non-nil, zero value otherwise.

### GetOldTsOk

`func (o *V1AuditLogsPostRequest) GetOldTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetOldTsOk returns a tuple with the OldTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTs

`func (o *V1AuditLogsPostRequest) SetOldTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetOldTs sets OldTs field to given value.

### HasOldTs

`func (o *V1AuditLogsPostRequest) HasOldTs() bool`

HasOldTs returns a boolean if a field has been set.

### GetRecentTs

`func (o *V1AuditLogsPostRequest) GetRecentTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetRecentTs returns the RecentTs field if non-nil, zero value otherwise.

### GetRecentTsOk

`func (o *V1AuditLogsPostRequest) GetRecentTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetRecentTsOk returns a tuple with the RecentTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTs

`func (o *V1AuditLogsPostRequest) SetRecentTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetRecentTs sets RecentTs field to given value.

### HasRecentTs

`func (o *V1AuditLogsPostRequest) HasRecentTs() bool`

HasRecentTs returns a boolean if a field has been set.

### GetSelectors

`func (o *V1AuditLogsPostRequest) GetSelectors() []V1AuditLogsPostRequestSelectorsInner`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *V1AuditLogsPostRequest) GetSelectorsOk() (*[]V1AuditLogsPostRequestSelectorsInner, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *V1AuditLogsPostRequest) SetSelectors(v []V1AuditLogsPostRequestSelectorsInner)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *V1AuditLogsPostRequest) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


