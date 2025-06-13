# V2MonitoringQueueInstantPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**IsDelta** | Pointer to **bool** |  | [optional] 
**IsTotal** | Pointer to **bool** |  | [optional] 
**Selectors** | Pointer to [**[]V2MonitoringQueueInstantPostRequestSelectorsInner**](V2MonitoringQueueInstantPostRequestSelectorsInner.md) |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV2MonitoringQueueInstantPostRequest

`func NewV2MonitoringQueueInstantPostRequest() *V2MonitoringQueueInstantPostRequest`

NewV2MonitoringQueueInstantPostRequest instantiates a new V2MonitoringQueueInstantPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringQueueInstantPostRequestWithDefaults

`func NewV2MonitoringQueueInstantPostRequestWithDefaults() *V2MonitoringQueueInstantPostRequest`

NewV2MonitoringQueueInstantPostRequestWithDefaults instantiates a new V2MonitoringQueueInstantPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V2MonitoringQueueInstantPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V2MonitoringQueueInstantPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V2MonitoringQueueInstantPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V2MonitoringQueueInstantPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetIsDelta

`func (o *V2MonitoringQueueInstantPostRequest) GetIsDelta() bool`

GetIsDelta returns the IsDelta field if non-nil, zero value otherwise.

### GetIsDeltaOk

`func (o *V2MonitoringQueueInstantPostRequest) GetIsDeltaOk() (*bool, bool)`

GetIsDeltaOk returns a tuple with the IsDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelta

`func (o *V2MonitoringQueueInstantPostRequest) SetIsDelta(v bool)`

SetIsDelta sets IsDelta field to given value.

### HasIsDelta

`func (o *V2MonitoringQueueInstantPostRequest) HasIsDelta() bool`

HasIsDelta returns a boolean if a field has been set.

### GetIsTotal

`func (o *V2MonitoringQueueInstantPostRequest) GetIsTotal() bool`

GetIsTotal returns the IsTotal field if non-nil, zero value otherwise.

### GetIsTotalOk

`func (o *V2MonitoringQueueInstantPostRequest) GetIsTotalOk() (*bool, bool)`

GetIsTotalOk returns a tuple with the IsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTotal

`func (o *V2MonitoringQueueInstantPostRequest) SetIsTotal(v bool)`

SetIsTotal sets IsTotal field to given value.

### HasIsTotal

`func (o *V2MonitoringQueueInstantPostRequest) HasIsTotal() bool`

HasIsTotal returns a boolean if a field has been set.

### GetSelectors

`func (o *V2MonitoringQueueInstantPostRequest) GetSelectors() []V2MonitoringQueueInstantPostRequestSelectorsInner`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *V2MonitoringQueueInstantPostRequest) GetSelectorsOk() (*[]V2MonitoringQueueInstantPostRequestSelectorsInner, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *V2MonitoringQueueInstantPostRequest) SetSelectors(v []V2MonitoringQueueInstantPostRequestSelectorsInner)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *V2MonitoringQueueInstantPostRequest) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2MonitoringQueueInstantPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2MonitoringQueueInstantPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2MonitoringQueueInstantPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2MonitoringQueueInstantPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


