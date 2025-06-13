# V1FlowsTopologyPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppSelector** | Pointer to [**V1FlowsTopologyPostRequestAppSelector**](V1FlowsTopologyPostRequestAppSelector.md) |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV1FlowsTopologyPostRequest

`func NewV1FlowsTopologyPostRequest() *V1FlowsTopologyPostRequest`

NewV1FlowsTopologyPostRequest instantiates a new V1FlowsTopologyPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FlowsTopologyPostRequestWithDefaults

`func NewV1FlowsTopologyPostRequestWithDefaults() *V1FlowsTopologyPostRequest`

NewV1FlowsTopologyPostRequestWithDefaults instantiates a new V1FlowsTopologyPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppSelector

`func (o *V1FlowsTopologyPostRequest) GetAppSelector() V1FlowsTopologyPostRequestAppSelector`

GetAppSelector returns the AppSelector field if non-nil, zero value otherwise.

### GetAppSelectorOk

`func (o *V1FlowsTopologyPostRequest) GetAppSelectorOk() (*V1FlowsTopologyPostRequestAppSelector, bool)`

GetAppSelectorOk returns a tuple with the AppSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSelector

`func (o *V1FlowsTopologyPostRequest) SetAppSelector(v V1FlowsTopologyPostRequestAppSelector)`

SetAppSelector sets AppSelector field to given value.

### HasAppSelector

`func (o *V1FlowsTopologyPostRequest) HasAppSelector() bool`

HasAppSelector returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1FlowsTopologyPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1FlowsTopologyPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1FlowsTopologyPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1FlowsTopologyPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V1FlowsTopologyPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V1FlowsTopologyPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V1FlowsTopologyPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V1FlowsTopologyPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


