# V1MonitoringCircuitsVisualizationPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**Selectors** | Pointer to [**[]V1MonitoringCircuitsVisualizationPostRequestSelectorsInner**](V1MonitoringCircuitsVisualizationPostRequestSelectorsInner.md) |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV1MonitoringCircuitsVisualizationPostRequest

`func NewV1MonitoringCircuitsVisualizationPostRequest() *V1MonitoringCircuitsVisualizationPostRequest`

NewV1MonitoringCircuitsVisualizationPostRequest instantiates a new V1MonitoringCircuitsVisualizationPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MonitoringCircuitsVisualizationPostRequestWithDefaults

`func NewV1MonitoringCircuitsVisualizationPostRequestWithDefaults() *V1MonitoringCircuitsVisualizationPostRequest`

NewV1MonitoringCircuitsVisualizationPostRequestWithDefaults instantiates a new V1MonitoringCircuitsVisualizationPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V1MonitoringCircuitsVisualizationPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1MonitoringCircuitsVisualizationPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1MonitoringCircuitsVisualizationPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1MonitoringCircuitsVisualizationPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetSelectors

`func (o *V1MonitoringCircuitsVisualizationPostRequest) GetSelectors() []V1MonitoringCircuitsVisualizationPostRequestSelectorsInner`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *V1MonitoringCircuitsVisualizationPostRequest) GetSelectorsOk() (*[]V1MonitoringCircuitsVisualizationPostRequestSelectorsInner, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *V1MonitoringCircuitsVisualizationPostRequest) SetSelectors(v []V1MonitoringCircuitsVisualizationPostRequestSelectorsInner)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *V1MonitoringCircuitsVisualizationPostRequest) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V1MonitoringCircuitsVisualizationPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V1MonitoringCircuitsVisualizationPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V1MonitoringCircuitsVisualizationPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V1MonitoringCircuitsVisualizationPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


