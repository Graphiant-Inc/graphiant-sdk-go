# V1AppsBandwidthPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **int32** |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**DlCircuitName** | Pointer to **string** |  | [optional] 
**IsDia** | Pointer to **bool** |  | [optional] 
**SlaClass** | Pointer to **string** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 
**UlCircuitName** | Pointer to **string** |  | [optional] 

## Methods

### NewV1AppsBandwidthPostRequest

`func NewV1AppsBandwidthPostRequest() *V1AppsBandwidthPostRequest`

NewV1AppsBandwidthPostRequest instantiates a new V1AppsBandwidthPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AppsBandwidthPostRequestWithDefaults

`func NewV1AppsBandwidthPostRequestWithDefaults() *V1AppsBandwidthPostRequest`

NewV1AppsBandwidthPostRequestWithDefaults instantiates a new V1AppsBandwidthPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *V1AppsBandwidthPostRequest) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *V1AppsBandwidthPostRequest) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *V1AppsBandwidthPostRequest) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *V1AppsBandwidthPostRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1AppsBandwidthPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1AppsBandwidthPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1AppsBandwidthPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1AppsBandwidthPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDlCircuitName

`func (o *V1AppsBandwidthPostRequest) GetDlCircuitName() string`

GetDlCircuitName returns the DlCircuitName field if non-nil, zero value otherwise.

### GetDlCircuitNameOk

`func (o *V1AppsBandwidthPostRequest) GetDlCircuitNameOk() (*string, bool)`

GetDlCircuitNameOk returns a tuple with the DlCircuitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlCircuitName

`func (o *V1AppsBandwidthPostRequest) SetDlCircuitName(v string)`

SetDlCircuitName sets DlCircuitName field to given value.

### HasDlCircuitName

`func (o *V1AppsBandwidthPostRequest) HasDlCircuitName() bool`

HasDlCircuitName returns a boolean if a field has been set.

### GetIsDia

`func (o *V1AppsBandwidthPostRequest) GetIsDia() bool`

GetIsDia returns the IsDia field if non-nil, zero value otherwise.

### GetIsDiaOk

`func (o *V1AppsBandwidthPostRequest) GetIsDiaOk() (*bool, bool)`

GetIsDiaOk returns a tuple with the IsDia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDia

`func (o *V1AppsBandwidthPostRequest) SetIsDia(v bool)`

SetIsDia sets IsDia field to given value.

### HasIsDia

`func (o *V1AppsBandwidthPostRequest) HasIsDia() bool`

HasIsDia returns a boolean if a field has been set.

### GetSlaClass

`func (o *V1AppsBandwidthPostRequest) GetSlaClass() string`

GetSlaClass returns the SlaClass field if non-nil, zero value otherwise.

### GetSlaClassOk

`func (o *V1AppsBandwidthPostRequest) GetSlaClassOk() (*string, bool)`

GetSlaClassOk returns a tuple with the SlaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaClass

`func (o *V1AppsBandwidthPostRequest) SetSlaClass(v string)`

SetSlaClass sets SlaClass field to given value.

### HasSlaClass

`func (o *V1AppsBandwidthPostRequest) HasSlaClass() bool`

HasSlaClass returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V1AppsBandwidthPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V1AppsBandwidthPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V1AppsBandwidthPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V1AppsBandwidthPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetUlCircuitName

`func (o *V1AppsBandwidthPostRequest) GetUlCircuitName() string`

GetUlCircuitName returns the UlCircuitName field if non-nil, zero value otherwise.

### GetUlCircuitNameOk

`func (o *V1AppsBandwidthPostRequest) GetUlCircuitNameOk() (*string, bool)`

GetUlCircuitNameOk returns a tuple with the UlCircuitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlCircuitName

`func (o *V1AppsBandwidthPostRequest) SetUlCircuitName(v string)`

SetUlCircuitName sets UlCircuitName field to given value.

### HasUlCircuitName

`func (o *V1AppsBandwidthPostRequest) HasUlCircuitName() bool`

HasUlCircuitName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


