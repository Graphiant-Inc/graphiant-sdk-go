# V2MonitoringInterfacePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**Selectors** | Pointer to [**[]V2MonitoringInterfacePostRequestSelectorsInner**](V2MonitoringInterfacePostRequestSelectorsInner.md) |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV2MonitoringInterfacePostRequest

`func NewV2MonitoringInterfacePostRequest() *V2MonitoringInterfacePostRequest`

NewV2MonitoringInterfacePostRequest instantiates a new V2MonitoringInterfacePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringInterfacePostRequestWithDefaults

`func NewV2MonitoringInterfacePostRequestWithDefaults() *V2MonitoringInterfacePostRequest`

NewV2MonitoringInterfacePostRequestWithDefaults instantiates a new V2MonitoringInterfacePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V2MonitoringInterfacePostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V2MonitoringInterfacePostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V2MonitoringInterfacePostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V2MonitoringInterfacePostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetSelectors

`func (o *V2MonitoringInterfacePostRequest) GetSelectors() []V2MonitoringInterfacePostRequestSelectorsInner`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *V2MonitoringInterfacePostRequest) GetSelectorsOk() (*[]V2MonitoringInterfacePostRequestSelectorsInner, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *V2MonitoringInterfacePostRequest) SetSelectors(v []V2MonitoringInterfacePostRequestSelectorsInner)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *V2MonitoringInterfacePostRequest) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2MonitoringInterfacePostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2MonitoringInterfacePostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2MonitoringInterfacePostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2MonitoringInterfacePostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


