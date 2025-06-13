# V2MonitoringQueuePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**Selectors** | Pointer to [**[]V2MonitoringQueuePostRequestSelectorsInner**](V2MonitoringQueuePostRequestSelectorsInner.md) |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV2MonitoringQueuePostRequest

`func NewV2MonitoringQueuePostRequest() *V2MonitoringQueuePostRequest`

NewV2MonitoringQueuePostRequest instantiates a new V2MonitoringQueuePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringQueuePostRequestWithDefaults

`func NewV2MonitoringQueuePostRequestWithDefaults() *V2MonitoringQueuePostRequest`

NewV2MonitoringQueuePostRequestWithDefaults instantiates a new V2MonitoringQueuePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V2MonitoringQueuePostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V2MonitoringQueuePostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V2MonitoringQueuePostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V2MonitoringQueuePostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetSelectors

`func (o *V2MonitoringQueuePostRequest) GetSelectors() []V2MonitoringQueuePostRequestSelectorsInner`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *V2MonitoringQueuePostRequest) GetSelectorsOk() (*[]V2MonitoringQueuePostRequestSelectorsInner, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *V2MonitoringQueuePostRequest) SetSelectors(v []V2MonitoringQueuePostRequestSelectorsInner)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *V2MonitoringQueuePostRequest) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2MonitoringQueuePostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2MonitoringQueuePostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2MonitoringQueuePostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2MonitoringQueuePostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


