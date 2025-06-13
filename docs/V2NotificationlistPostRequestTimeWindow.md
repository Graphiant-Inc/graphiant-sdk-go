# V2NotificationlistPostRequestTimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketSizeSec** | Pointer to **int32** |  | [optional] 
**OldTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**RecentTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV2NotificationlistPostRequestTimeWindow

`func NewV2NotificationlistPostRequestTimeWindow() *V2NotificationlistPostRequestTimeWindow`

NewV2NotificationlistPostRequestTimeWindow instantiates a new V2NotificationlistPostRequestTimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2NotificationlistPostRequestTimeWindowWithDefaults

`func NewV2NotificationlistPostRequestTimeWindowWithDefaults() *V2NotificationlistPostRequestTimeWindow`

NewV2NotificationlistPostRequestTimeWindowWithDefaults instantiates a new V2NotificationlistPostRequestTimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketSizeSec

`func (o *V2NotificationlistPostRequestTimeWindow) GetBucketSizeSec() int32`

GetBucketSizeSec returns the BucketSizeSec field if non-nil, zero value otherwise.

### GetBucketSizeSecOk

`func (o *V2NotificationlistPostRequestTimeWindow) GetBucketSizeSecOk() (*int32, bool)`

GetBucketSizeSecOk returns a tuple with the BucketSizeSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSizeSec

`func (o *V2NotificationlistPostRequestTimeWindow) SetBucketSizeSec(v int32)`

SetBucketSizeSec sets BucketSizeSec field to given value.

### HasBucketSizeSec

`func (o *V2NotificationlistPostRequestTimeWindow) HasBucketSizeSec() bool`

HasBucketSizeSec returns a boolean if a field has been set.

### GetOldTs

`func (o *V2NotificationlistPostRequestTimeWindow) GetOldTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetOldTs returns the OldTs field if non-nil, zero value otherwise.

### GetOldTsOk

`func (o *V2NotificationlistPostRequestTimeWindow) GetOldTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetOldTsOk returns a tuple with the OldTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTs

`func (o *V2NotificationlistPostRequestTimeWindow) SetOldTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetOldTs sets OldTs field to given value.

### HasOldTs

`func (o *V2NotificationlistPostRequestTimeWindow) HasOldTs() bool`

HasOldTs returns a boolean if a field has been set.

### GetRecentTs

`func (o *V2NotificationlistPostRequestTimeWindow) GetRecentTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetRecentTs returns the RecentTs field if non-nil, zero value otherwise.

### GetRecentTsOk

`func (o *V2NotificationlistPostRequestTimeWindow) GetRecentTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetRecentTsOk returns a tuple with the RecentTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTs

`func (o *V2NotificationlistPostRequestTimeWindow) SetRecentTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetRecentTs sets RecentTs field to given value.

### HasRecentTs

`func (o *V2NotificationlistPostRequestTimeWindow) HasRecentTs() bool`

HasRecentTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


