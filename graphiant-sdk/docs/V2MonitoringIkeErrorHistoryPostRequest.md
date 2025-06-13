# V2MonitoringIkeErrorHistoryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 
**TunnelName** | Pointer to **string** |  | [optional] 

## Methods

### NewV2MonitoringIkeErrorHistoryPostRequest

`func NewV2MonitoringIkeErrorHistoryPostRequest() *V2MonitoringIkeErrorHistoryPostRequest`

NewV2MonitoringIkeErrorHistoryPostRequest instantiates a new V2MonitoringIkeErrorHistoryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringIkeErrorHistoryPostRequestWithDefaults

`func NewV2MonitoringIkeErrorHistoryPostRequestWithDefaults() *V2MonitoringIkeErrorHistoryPostRequest`

NewV2MonitoringIkeErrorHistoryPostRequestWithDefaults instantiates a new V2MonitoringIkeErrorHistoryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V2MonitoringIkeErrorHistoryPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V2MonitoringIkeErrorHistoryPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V2MonitoringIkeErrorHistoryPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V2MonitoringIkeErrorHistoryPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2MonitoringIkeErrorHistoryPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2MonitoringIkeErrorHistoryPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2MonitoringIkeErrorHistoryPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2MonitoringIkeErrorHistoryPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetTunnelName

`func (o *V2MonitoringIkeErrorHistoryPostRequest) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *V2MonitoringIkeErrorHistoryPostRequest) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *V2MonitoringIkeErrorHistoryPostRequest) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.

### HasTunnelName

`func (o *V2MonitoringIkeErrorHistoryPostRequest) HasTunnelName() bool`

HasTunnelName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


