# V1LogsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorRef** | Pointer to **string** |  | [optional] 
**CustomerView** | Pointer to **bool** |  | [optional] 
**DeviceIds** | Pointer to **[]int64** |  | [optional] 
**HistogramBucketSizeSec** | Pointer to **int32** |  | [optional] 
**NumLogs** | Pointer to **int32** |  | [optional] 
**OldTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**RecentTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Selectors** | Pointer to [**[]V1LogsPostRequestSelectorsInner**](V1LogsPostRequestSelectorsInner.md) |  | [optional] 

## Methods

### NewV1LogsPostRequest

`func NewV1LogsPostRequest() *V1LogsPostRequest`

NewV1LogsPostRequest instantiates a new V1LogsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LogsPostRequestWithDefaults

`func NewV1LogsPostRequestWithDefaults() *V1LogsPostRequest`

NewV1LogsPostRequestWithDefaults instantiates a new V1LogsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorRef

`func (o *V1LogsPostRequest) GetCursorRef() string`

GetCursorRef returns the CursorRef field if non-nil, zero value otherwise.

### GetCursorRefOk

`func (o *V1LogsPostRequest) GetCursorRefOk() (*string, bool)`

GetCursorRefOk returns a tuple with the CursorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorRef

`func (o *V1LogsPostRequest) SetCursorRef(v string)`

SetCursorRef sets CursorRef field to given value.

### HasCursorRef

`func (o *V1LogsPostRequest) HasCursorRef() bool`

HasCursorRef returns a boolean if a field has been set.

### GetCustomerView

`func (o *V1LogsPostRequest) GetCustomerView() bool`

GetCustomerView returns the CustomerView field if non-nil, zero value otherwise.

### GetCustomerViewOk

`func (o *V1LogsPostRequest) GetCustomerViewOk() (*bool, bool)`

GetCustomerViewOk returns a tuple with the CustomerView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerView

`func (o *V1LogsPostRequest) SetCustomerView(v bool)`

SetCustomerView sets CustomerView field to given value.

### HasCustomerView

`func (o *V1LogsPostRequest) HasCustomerView() bool`

HasCustomerView returns a boolean if a field has been set.

### GetDeviceIds

`func (o *V1LogsPostRequest) GetDeviceIds() []int64`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *V1LogsPostRequest) GetDeviceIdsOk() (*[]int64, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *V1LogsPostRequest) SetDeviceIds(v []int64)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *V1LogsPostRequest) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetHistogramBucketSizeSec

`func (o *V1LogsPostRequest) GetHistogramBucketSizeSec() int32`

GetHistogramBucketSizeSec returns the HistogramBucketSizeSec field if non-nil, zero value otherwise.

### GetHistogramBucketSizeSecOk

`func (o *V1LogsPostRequest) GetHistogramBucketSizeSecOk() (*int32, bool)`

GetHistogramBucketSizeSecOk returns a tuple with the HistogramBucketSizeSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramBucketSizeSec

`func (o *V1LogsPostRequest) SetHistogramBucketSizeSec(v int32)`

SetHistogramBucketSizeSec sets HistogramBucketSizeSec field to given value.

### HasHistogramBucketSizeSec

`func (o *V1LogsPostRequest) HasHistogramBucketSizeSec() bool`

HasHistogramBucketSizeSec returns a boolean if a field has been set.

### GetNumLogs

`func (o *V1LogsPostRequest) GetNumLogs() int32`

GetNumLogs returns the NumLogs field if non-nil, zero value otherwise.

### GetNumLogsOk

`func (o *V1LogsPostRequest) GetNumLogsOk() (*int32, bool)`

GetNumLogsOk returns a tuple with the NumLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLogs

`func (o *V1LogsPostRequest) SetNumLogs(v int32)`

SetNumLogs sets NumLogs field to given value.

### HasNumLogs

`func (o *V1LogsPostRequest) HasNumLogs() bool`

HasNumLogs returns a boolean if a field has been set.

### GetOldTs

`func (o *V1LogsPostRequest) GetOldTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetOldTs returns the OldTs field if non-nil, zero value otherwise.

### GetOldTsOk

`func (o *V1LogsPostRequest) GetOldTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetOldTsOk returns a tuple with the OldTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTs

`func (o *V1LogsPostRequest) SetOldTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetOldTs sets OldTs field to given value.

### HasOldTs

`func (o *V1LogsPostRequest) HasOldTs() bool`

HasOldTs returns a boolean if a field has been set.

### GetRecentTs

`func (o *V1LogsPostRequest) GetRecentTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetRecentTs returns the RecentTs field if non-nil, zero value otherwise.

### GetRecentTsOk

`func (o *V1LogsPostRequest) GetRecentTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetRecentTsOk returns a tuple with the RecentTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTs

`func (o *V1LogsPostRequest) SetRecentTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetRecentTs sets RecentTs field to given value.

### HasRecentTs

`func (o *V1LogsPostRequest) HasRecentTs() bool`

HasRecentTs returns a boolean if a field has been set.

### GetSelectors

`func (o *V1LogsPostRequest) GetSelectors() []V1LogsPostRequestSelectorsInner`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *V1LogsPostRequest) GetSelectorsOk() (*[]V1LogsPostRequestSelectorsInner, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *V1LogsPostRequest) SetSelectors(v []V1LogsPostRequestSelectorsInner)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *V1LogsPostRequest) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


