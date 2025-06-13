# V1AppsAppSummaryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**IsDia** | Pointer to **bool** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV1AppsAppSummaryPostRequest

`func NewV1AppsAppSummaryPostRequest() *V1AppsAppSummaryPostRequest`

NewV1AppsAppSummaryPostRequest instantiates a new V1AppsAppSummaryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AppsAppSummaryPostRequestWithDefaults

`func NewV1AppsAppSummaryPostRequestWithDefaults() *V1AppsAppSummaryPostRequest`

NewV1AppsAppSummaryPostRequestWithDefaults instantiates a new V1AppsAppSummaryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V1AppsAppSummaryPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1AppsAppSummaryPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1AppsAppSummaryPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1AppsAppSummaryPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetIsDia

`func (o *V1AppsAppSummaryPostRequest) GetIsDia() bool`

GetIsDia returns the IsDia field if non-nil, zero value otherwise.

### GetIsDiaOk

`func (o *V1AppsAppSummaryPostRequest) GetIsDiaOk() (*bool, bool)`

GetIsDiaOk returns a tuple with the IsDia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDia

`func (o *V1AppsAppSummaryPostRequest) SetIsDia(v bool)`

SetIsDia sets IsDia field to given value.

### HasIsDia

`func (o *V1AppsAppSummaryPostRequest) HasIsDia() bool`

HasIsDia returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V1AppsAppSummaryPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V1AppsAppSummaryPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V1AppsAppSummaryPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V1AppsAppSummaryPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


