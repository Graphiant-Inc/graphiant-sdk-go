# V1DiagnosticSpeedtestPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**Params** | Pointer to [**V1DiagnosticSpeedtestPostRequestParams**](V1DiagnosticSpeedtestPostRequestParams.md) |  | [optional] 

## Methods

### NewV1DiagnosticSpeedtestPostRequest

`func NewV1DiagnosticSpeedtestPostRequest() *V1DiagnosticSpeedtestPostRequest`

NewV1DiagnosticSpeedtestPostRequest instantiates a new V1DiagnosticSpeedtestPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DiagnosticSpeedtestPostRequestWithDefaults

`func NewV1DiagnosticSpeedtestPostRequestWithDefaults() *V1DiagnosticSpeedtestPostRequest`

NewV1DiagnosticSpeedtestPostRequestWithDefaults instantiates a new V1DiagnosticSpeedtestPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V1DiagnosticSpeedtestPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1DiagnosticSpeedtestPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1DiagnosticSpeedtestPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1DiagnosticSpeedtestPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetParams

`func (o *V1DiagnosticSpeedtestPostRequest) GetParams() V1DiagnosticSpeedtestPostRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *V1DiagnosticSpeedtestPostRequest) GetParamsOk() (*V1DiagnosticSpeedtestPostRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *V1DiagnosticSpeedtestPostRequest) SetParams(v V1DiagnosticSpeedtestPostRequestParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *V1DiagnosticSpeedtestPostRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


