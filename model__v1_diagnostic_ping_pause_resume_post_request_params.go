/*
Graphiant APIs

**To use the APIs:**  1) Login using `/api/v1/auth/login`   2) Copy the value of \"token\" in the response   3) Click the \"Authorize\" button   4) In the \"Value\" text field enter: `Bearer <your token>`   5) Click \"Authorize\"   6) All requests are now authorized.  **Token valid for 2 hours. If expired:**   - Login again, click \"Authorize\", paste new token.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graphiant_sdk

import (
	"encoding/json"
)

// checks if the V1DiagnosticPingPauseResumePostRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DiagnosticPingPauseResumePostRequestParams{}

// V1DiagnosticPingPauseResumePostRequestParams struct for V1DiagnosticPingPauseResumePostRequestParams
type V1DiagnosticPingPauseResumePostRequestParams struct {
	DestAddress *string `json:"destAddress,omitempty"`
	HopStatsCount *int32 `json:"hopStatsCount,omitempty"`
	Interface *string `json:"interface,omitempty"`
	PayloadSize *int32 `json:"payloadSize,omitempty"`
	Port *int32 `json:"port,omitempty"`
	ProbeCount *int32 `json:"probeCount,omitempty"`
	SrcAddress *string `json:"srcAddress,omitempty"`
	Tos *int32 `json:"tos,omitempty"`
	VrfName *string `json:"vrfName,omitempty"`
}

// NewV1DiagnosticPingPauseResumePostRequestParams instantiates a new V1DiagnosticPingPauseResumePostRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DiagnosticPingPauseResumePostRequestParams() *V1DiagnosticPingPauseResumePostRequestParams {
	this := V1DiagnosticPingPauseResumePostRequestParams{}
	return &this
}

// NewV1DiagnosticPingPauseResumePostRequestParamsWithDefaults instantiates a new V1DiagnosticPingPauseResumePostRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DiagnosticPingPauseResumePostRequestParamsWithDefaults() *V1DiagnosticPingPauseResumePostRequestParams {
	this := V1DiagnosticPingPauseResumePostRequestParams{}
	return &this
}

// GetDestAddress returns the DestAddress field value if set, zero value otherwise.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetDestAddress() string {
	if o == nil || IsNil(o.DestAddress) {
		var ret string
		return ret
	}
	return *o.DestAddress
}

// GetDestAddressOk returns a tuple with the DestAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetDestAddressOk() (*string, bool) {
	if o == nil || IsNil(o.DestAddress) {
		return nil, false
	}
	return o.DestAddress, true
}

// HasDestAddress returns a boolean if a field has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) HasDestAddress() bool {
	if o != nil && !IsNil(o.DestAddress) {
		return true
	}

	return false
}

// SetDestAddress gets a reference to the given string and assigns it to the DestAddress field.
func (o *V1DiagnosticPingPauseResumePostRequestParams) SetDestAddress(v string) {
	o.DestAddress = &v
}

// GetHopStatsCount returns the HopStatsCount field value if set, zero value otherwise.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetHopStatsCount() int32 {
	if o == nil || IsNil(o.HopStatsCount) {
		var ret int32
		return ret
	}
	return *o.HopStatsCount
}

// GetHopStatsCountOk returns a tuple with the HopStatsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetHopStatsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.HopStatsCount) {
		return nil, false
	}
	return o.HopStatsCount, true
}

// HasHopStatsCount returns a boolean if a field has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) HasHopStatsCount() bool {
	if o != nil && !IsNil(o.HopStatsCount) {
		return true
	}

	return false
}

// SetHopStatsCount gets a reference to the given int32 and assigns it to the HopStatsCount field.
func (o *V1DiagnosticPingPauseResumePostRequestParams) SetHopStatsCount(v int32) {
	o.HopStatsCount = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *V1DiagnosticPingPauseResumePostRequestParams) SetInterface(v string) {
	o.Interface = &v
}

// GetPayloadSize returns the PayloadSize field value if set, zero value otherwise.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetPayloadSize() int32 {
	if o == nil || IsNil(o.PayloadSize) {
		var ret int32
		return ret
	}
	return *o.PayloadSize
}

// GetPayloadSizeOk returns a tuple with the PayloadSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetPayloadSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PayloadSize) {
		return nil, false
	}
	return o.PayloadSize, true
}

// HasPayloadSize returns a boolean if a field has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) HasPayloadSize() bool {
	if o != nil && !IsNil(o.PayloadSize) {
		return true
	}

	return false
}

// SetPayloadSize gets a reference to the given int32 and assigns it to the PayloadSize field.
func (o *V1DiagnosticPingPauseResumePostRequestParams) SetPayloadSize(v int32) {
	o.PayloadSize = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *V1DiagnosticPingPauseResumePostRequestParams) SetPort(v int32) {
	o.Port = &v
}

// GetProbeCount returns the ProbeCount field value if set, zero value otherwise.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetProbeCount() int32 {
	if o == nil || IsNil(o.ProbeCount) {
		var ret int32
		return ret
	}
	return *o.ProbeCount
}

// GetProbeCountOk returns a tuple with the ProbeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetProbeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProbeCount) {
		return nil, false
	}
	return o.ProbeCount, true
}

// HasProbeCount returns a boolean if a field has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) HasProbeCount() bool {
	if o != nil && !IsNil(o.ProbeCount) {
		return true
	}

	return false
}

// SetProbeCount gets a reference to the given int32 and assigns it to the ProbeCount field.
func (o *V1DiagnosticPingPauseResumePostRequestParams) SetProbeCount(v int32) {
	o.ProbeCount = &v
}

// GetSrcAddress returns the SrcAddress field value if set, zero value otherwise.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetSrcAddress() string {
	if o == nil || IsNil(o.SrcAddress) {
		var ret string
		return ret
	}
	return *o.SrcAddress
}

// GetSrcAddressOk returns a tuple with the SrcAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetSrcAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SrcAddress) {
		return nil, false
	}
	return o.SrcAddress, true
}

// HasSrcAddress returns a boolean if a field has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) HasSrcAddress() bool {
	if o != nil && !IsNil(o.SrcAddress) {
		return true
	}

	return false
}

// SetSrcAddress gets a reference to the given string and assigns it to the SrcAddress field.
func (o *V1DiagnosticPingPauseResumePostRequestParams) SetSrcAddress(v string) {
	o.SrcAddress = &v
}

// GetTos returns the Tos field value if set, zero value otherwise.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetTos() int32 {
	if o == nil || IsNil(o.Tos) {
		var ret int32
		return ret
	}
	return *o.Tos
}

// GetTosOk returns a tuple with the Tos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetTosOk() (*int32, bool) {
	if o == nil || IsNil(o.Tos) {
		return nil, false
	}
	return o.Tos, true
}

// HasTos returns a boolean if a field has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) HasTos() bool {
	if o != nil && !IsNil(o.Tos) {
		return true
	}

	return false
}

// SetTos gets a reference to the given int32 and assigns it to the Tos field.
func (o *V1DiagnosticPingPauseResumePostRequestParams) SetTos(v int32) {
	o.Tos = &v
}

// GetVrfName returns the VrfName field value if set, zero value otherwise.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetVrfName() string {
	if o == nil || IsNil(o.VrfName) {
		var ret string
		return ret
	}
	return *o.VrfName
}

// GetVrfNameOk returns a tuple with the VrfName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) GetVrfNameOk() (*string, bool) {
	if o == nil || IsNil(o.VrfName) {
		return nil, false
	}
	return o.VrfName, true
}

// HasVrfName returns a boolean if a field has been set.
func (o *V1DiagnosticPingPauseResumePostRequestParams) HasVrfName() bool {
	if o != nil && !IsNil(o.VrfName) {
		return true
	}

	return false
}

// SetVrfName gets a reference to the given string and assigns it to the VrfName field.
func (o *V1DiagnosticPingPauseResumePostRequestParams) SetVrfName(v string) {
	o.VrfName = &v
}

func (o V1DiagnosticPingPauseResumePostRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DiagnosticPingPauseResumePostRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DestAddress) {
		toSerialize["destAddress"] = o.DestAddress
	}
	if !IsNil(o.HopStatsCount) {
		toSerialize["hopStatsCount"] = o.HopStatsCount
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.PayloadSize) {
		toSerialize["payloadSize"] = o.PayloadSize
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.ProbeCount) {
		toSerialize["probeCount"] = o.ProbeCount
	}
	if !IsNil(o.SrcAddress) {
		toSerialize["srcAddress"] = o.SrcAddress
	}
	if !IsNil(o.Tos) {
		toSerialize["tos"] = o.Tos
	}
	if !IsNil(o.VrfName) {
		toSerialize["vrfName"] = o.VrfName
	}
	return toSerialize, nil
}

type NullableV1DiagnosticPingPauseResumePostRequestParams struct {
	value *V1DiagnosticPingPauseResumePostRequestParams
	isSet bool
}

func (v NullableV1DiagnosticPingPauseResumePostRequestParams) Get() *V1DiagnosticPingPauseResumePostRequestParams {
	return v.value
}

func (v *NullableV1DiagnosticPingPauseResumePostRequestParams) Set(val *V1DiagnosticPingPauseResumePostRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DiagnosticPingPauseResumePostRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DiagnosticPingPauseResumePostRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DiagnosticPingPauseResumePostRequestParams(val *V1DiagnosticPingPauseResumePostRequestParams) *NullableV1DiagnosticPingPauseResumePostRequestParams {
	return &NullableV1DiagnosticPingPauseResumePostRequestParams{value: val, isSet: true}
}

func (v NullableV1DiagnosticPingPauseResumePostRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DiagnosticPingPauseResumePostRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


