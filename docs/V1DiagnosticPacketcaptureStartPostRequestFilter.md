# V1DiagnosticPacketcaptureStartPostRequestFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**V1DiagnosticPacketcaptureStartPostRequestFilterDestination**](V1DiagnosticPacketcaptureStartPostRequestFilterDestination.md) |  | [optional] 
**Dscp** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**V1DiagnosticPacketcaptureStartPostRequestFilterDestination**](V1DiagnosticPacketcaptureStartPostRequestFilterDestination.md) |  | [optional] 

## Methods

### NewV1DiagnosticPacketcaptureStartPostRequestFilter

`func NewV1DiagnosticPacketcaptureStartPostRequestFilter() *V1DiagnosticPacketcaptureStartPostRequestFilter`

NewV1DiagnosticPacketcaptureStartPostRequestFilter instantiates a new V1DiagnosticPacketcaptureStartPostRequestFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DiagnosticPacketcaptureStartPostRequestFilterWithDefaults

`func NewV1DiagnosticPacketcaptureStartPostRequestFilterWithDefaults() *V1DiagnosticPacketcaptureStartPostRequestFilter`

NewV1DiagnosticPacketcaptureStartPostRequestFilterWithDefaults instantiates a new V1DiagnosticPacketcaptureStartPostRequestFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) GetDestination() V1DiagnosticPacketcaptureStartPostRequestFilterDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) GetDestinationOk() (*V1DiagnosticPacketcaptureStartPostRequestFilterDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) SetDestination(v V1DiagnosticPacketcaptureStartPostRequestFilterDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDscp

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) GetDscp() int32`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) GetDscpOk() (*int32, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) SetDscp(v int32)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetProtocol

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) GetSource() V1DiagnosticPacketcaptureStartPostRequestFilterDestination`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) GetSourceOk() (*V1DiagnosticPacketcaptureStartPostRequestFilterDestination, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) SetSource(v V1DiagnosticPacketcaptureStartPostRequestFilterDestination)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1DiagnosticPacketcaptureStartPostRequestFilter) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


