# V1DiagnosticPingPauseResumePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**Params** | Pointer to [**V1DiagnosticPingPauseResumePostRequestParams**](V1DiagnosticPingPauseResumePostRequestParams.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TransportType** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DiagnosticPingPauseResumePostRequest

`func NewV1DiagnosticPingPauseResumePostRequest() *V1DiagnosticPingPauseResumePostRequest`

NewV1DiagnosticPingPauseResumePostRequest instantiates a new V1DiagnosticPingPauseResumePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DiagnosticPingPauseResumePostRequestWithDefaults

`func NewV1DiagnosticPingPauseResumePostRequestWithDefaults() *V1DiagnosticPingPauseResumePostRequest`

NewV1DiagnosticPingPauseResumePostRequestWithDefaults instantiates a new V1DiagnosticPingPauseResumePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V1DiagnosticPingPauseResumePostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1DiagnosticPingPauseResumePostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1DiagnosticPingPauseResumePostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1DiagnosticPingPauseResumePostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetParams

`func (o *V1DiagnosticPingPauseResumePostRequest) GetParams() V1DiagnosticPingPauseResumePostRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *V1DiagnosticPingPauseResumePostRequest) GetParamsOk() (*V1DiagnosticPingPauseResumePostRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *V1DiagnosticPingPauseResumePostRequest) SetParams(v V1DiagnosticPingPauseResumePostRequestParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *V1DiagnosticPingPauseResumePostRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetToken

`func (o *V1DiagnosticPingPauseResumePostRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *V1DiagnosticPingPauseResumePostRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *V1DiagnosticPingPauseResumePostRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *V1DiagnosticPingPauseResumePostRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTransportType

`func (o *V1DiagnosticPingPauseResumePostRequest) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *V1DiagnosticPingPauseResumePostRequest) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *V1DiagnosticPingPauseResumePostRequest) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.

### HasTransportType

`func (o *V1DiagnosticPingPauseResumePostRequest) HasTransportType() bool`

HasTransportType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


