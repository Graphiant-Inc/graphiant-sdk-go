# V2DeviceDeviceIdTopologyPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotTime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV2DeviceDeviceIdTopologyPostRequest

`func NewV2DeviceDeviceIdTopologyPostRequest() *V2DeviceDeviceIdTopologyPostRequest`

NewV2DeviceDeviceIdTopologyPostRequest instantiates a new V2DeviceDeviceIdTopologyPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2DeviceDeviceIdTopologyPostRequestWithDefaults

`func NewV2DeviceDeviceIdTopologyPostRequestWithDefaults() *V2DeviceDeviceIdTopologyPostRequest`

NewV2DeviceDeviceIdTopologyPostRequestWithDefaults instantiates a new V2DeviceDeviceIdTopologyPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotTime

`func (o *V2DeviceDeviceIdTopologyPostRequest) GetSnapshotTime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetSnapshotTime returns the SnapshotTime field if non-nil, zero value otherwise.

### GetSnapshotTimeOk

`func (o *V2DeviceDeviceIdTopologyPostRequest) GetSnapshotTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetSnapshotTimeOk returns a tuple with the SnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTime

`func (o *V2DeviceDeviceIdTopologyPostRequest) SetSnapshotTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetSnapshotTime sets SnapshotTime field to given value.

### HasSnapshotTime

`func (o *V2DeviceDeviceIdTopologyPostRequest) HasSnapshotTime() bool`

HasSnapshotTime returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2DeviceDeviceIdTopologyPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2DeviceDeviceIdTopologyPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2DeviceDeviceIdTopologyPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2DeviceDeviceIdTopologyPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


