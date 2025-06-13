# V1DiagnosticPingPauseResumePostRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestAddress** | Pointer to **string** |  | [optional] 
**HopStatsCount** | Pointer to **int32** |  | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**PayloadSize** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**ProbeCount** | Pointer to **int32** |  | [optional] 
**SrcAddress** | Pointer to **string** |  | [optional] 
**Tos** | Pointer to **int32** |  | [optional] 
**VrfName** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DiagnosticPingPauseResumePostRequestParams

`func NewV1DiagnosticPingPauseResumePostRequestParams() *V1DiagnosticPingPauseResumePostRequestParams`

NewV1DiagnosticPingPauseResumePostRequestParams instantiates a new V1DiagnosticPingPauseResumePostRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DiagnosticPingPauseResumePostRequestParamsWithDefaults

`func NewV1DiagnosticPingPauseResumePostRequestParamsWithDefaults() *V1DiagnosticPingPauseResumePostRequestParams`

NewV1DiagnosticPingPauseResumePostRequestParamsWithDefaults instantiates a new V1DiagnosticPingPauseResumePostRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestAddress

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetDestAddress() string`

GetDestAddress returns the DestAddress field if non-nil, zero value otherwise.

### GetDestAddressOk

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetDestAddressOk() (*string, bool)`

GetDestAddressOk returns a tuple with the DestAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAddress

`func (o *V1DiagnosticPingPauseResumePostRequestParams) SetDestAddress(v string)`

SetDestAddress sets DestAddress field to given value.

### HasDestAddress

`func (o *V1DiagnosticPingPauseResumePostRequestParams) HasDestAddress() bool`

HasDestAddress returns a boolean if a field has been set.

### GetHopStatsCount

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetHopStatsCount() int32`

GetHopStatsCount returns the HopStatsCount field if non-nil, zero value otherwise.

### GetHopStatsCountOk

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetHopStatsCountOk() (*int32, bool)`

GetHopStatsCountOk returns a tuple with the HopStatsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopStatsCount

`func (o *V1DiagnosticPingPauseResumePostRequestParams) SetHopStatsCount(v int32)`

SetHopStatsCount sets HopStatsCount field to given value.

### HasHopStatsCount

`func (o *V1DiagnosticPingPauseResumePostRequestParams) HasHopStatsCount() bool`

HasHopStatsCount returns a boolean if a field has been set.

### GetInterface

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *V1DiagnosticPingPauseResumePostRequestParams) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *V1DiagnosticPingPauseResumePostRequestParams) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetPayloadSize

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetPayloadSize() int32`

GetPayloadSize returns the PayloadSize field if non-nil, zero value otherwise.

### GetPayloadSizeOk

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetPayloadSizeOk() (*int32, bool)`

GetPayloadSizeOk returns a tuple with the PayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadSize

`func (o *V1DiagnosticPingPauseResumePostRequestParams) SetPayloadSize(v int32)`

SetPayloadSize sets PayloadSize field to given value.

### HasPayloadSize

`func (o *V1DiagnosticPingPauseResumePostRequestParams) HasPayloadSize() bool`

HasPayloadSize returns a boolean if a field has been set.

### GetPort

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V1DiagnosticPingPauseResumePostRequestParams) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V1DiagnosticPingPauseResumePostRequestParams) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProbeCount

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetProbeCount() int32`

GetProbeCount returns the ProbeCount field if non-nil, zero value otherwise.

### GetProbeCountOk

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetProbeCountOk() (*int32, bool)`

GetProbeCountOk returns a tuple with the ProbeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeCount

`func (o *V1DiagnosticPingPauseResumePostRequestParams) SetProbeCount(v int32)`

SetProbeCount sets ProbeCount field to given value.

### HasProbeCount

`func (o *V1DiagnosticPingPauseResumePostRequestParams) HasProbeCount() bool`

HasProbeCount returns a boolean if a field has been set.

### GetSrcAddress

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetSrcAddress() string`

GetSrcAddress returns the SrcAddress field if non-nil, zero value otherwise.

### GetSrcAddressOk

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetSrcAddressOk() (*string, bool)`

GetSrcAddressOk returns a tuple with the SrcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcAddress

`func (o *V1DiagnosticPingPauseResumePostRequestParams) SetSrcAddress(v string)`

SetSrcAddress sets SrcAddress field to given value.

### HasSrcAddress

`func (o *V1DiagnosticPingPauseResumePostRequestParams) HasSrcAddress() bool`

HasSrcAddress returns a boolean if a field has been set.

### GetTos

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetTos() int32`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetTosOk() (*int32, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *V1DiagnosticPingPauseResumePostRequestParams) SetTos(v int32)`

SetTos sets Tos field to given value.

### HasTos

`func (o *V1DiagnosticPingPauseResumePostRequestParams) HasTos() bool`

HasTos returns a boolean if a field has been set.

### GetVrfName

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *V1DiagnosticPingPauseResumePostRequestParams) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *V1DiagnosticPingPauseResumePostRequestParams) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *V1DiagnosticPingPauseResumePostRequestParams) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


