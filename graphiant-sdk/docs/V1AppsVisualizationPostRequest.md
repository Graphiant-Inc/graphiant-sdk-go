# V1AppsVisualizationPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitName** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**IsDia** | Pointer to **bool** |  | [optional] 
**SlaClass** | Pointer to **string** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV1AppsVisualizationPostRequest

`func NewV1AppsVisualizationPostRequest() *V1AppsVisualizationPostRequest`

NewV1AppsVisualizationPostRequest instantiates a new V1AppsVisualizationPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AppsVisualizationPostRequestWithDefaults

`func NewV1AppsVisualizationPostRequestWithDefaults() *V1AppsVisualizationPostRequest`

NewV1AppsVisualizationPostRequestWithDefaults instantiates a new V1AppsVisualizationPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitName

`func (o *V1AppsVisualizationPostRequest) GetCircuitName() string`

GetCircuitName returns the CircuitName field if non-nil, zero value otherwise.

### GetCircuitNameOk

`func (o *V1AppsVisualizationPostRequest) GetCircuitNameOk() (*string, bool)`

GetCircuitNameOk returns a tuple with the CircuitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitName

`func (o *V1AppsVisualizationPostRequest) SetCircuitName(v string)`

SetCircuitName sets CircuitName field to given value.

### HasCircuitName

`func (o *V1AppsVisualizationPostRequest) HasCircuitName() bool`

HasCircuitName returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1AppsVisualizationPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1AppsVisualizationPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1AppsVisualizationPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1AppsVisualizationPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetIsDia

`func (o *V1AppsVisualizationPostRequest) GetIsDia() bool`

GetIsDia returns the IsDia field if non-nil, zero value otherwise.

### GetIsDiaOk

`func (o *V1AppsVisualizationPostRequest) GetIsDiaOk() (*bool, bool)`

GetIsDiaOk returns a tuple with the IsDia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDia

`func (o *V1AppsVisualizationPostRequest) SetIsDia(v bool)`

SetIsDia sets IsDia field to given value.

### HasIsDia

`func (o *V1AppsVisualizationPostRequest) HasIsDia() bool`

HasIsDia returns a boolean if a field has been set.

### GetSlaClass

`func (o *V1AppsVisualizationPostRequest) GetSlaClass() string`

GetSlaClass returns the SlaClass field if non-nil, zero value otherwise.

### GetSlaClassOk

`func (o *V1AppsVisualizationPostRequest) GetSlaClassOk() (*string, bool)`

GetSlaClassOk returns a tuple with the SlaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaClass

`func (o *V1AppsVisualizationPostRequest) SetSlaClass(v string)`

SetSlaClass sets SlaClass field to given value.

### HasSlaClass

`func (o *V1AppsVisualizationPostRequest) HasSlaClass() bool`

HasSlaClass returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V1AppsVisualizationPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V1AppsVisualizationPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V1AppsVisualizationPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V1AppsVisualizationPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


