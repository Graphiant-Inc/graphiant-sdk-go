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

// checks if the V1DiagnosticResetIpsecSessionDeviceIdPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DiagnosticResetIpsecSessionDeviceIdPutRequest{}

// V1DiagnosticResetIpsecSessionDeviceIdPutRequest struct for V1DiagnosticResetIpsecSessionDeviceIdPutRequest
type V1DiagnosticResetIpsecSessionDeviceIdPutRequest struct {
	All3RdParty *bool `json:"all3RdParty,omitempty"`
	AllControllers *bool `json:"allControllers,omitempty"`
	AllE2E *bool `json:"allE2E,omitempty"`
	Vrf []string `json:"vrf,omitempty"`
}

// NewV1DiagnosticResetIpsecSessionDeviceIdPutRequest instantiates a new V1DiagnosticResetIpsecSessionDeviceIdPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DiagnosticResetIpsecSessionDeviceIdPutRequest() *V1DiagnosticResetIpsecSessionDeviceIdPutRequest {
	this := V1DiagnosticResetIpsecSessionDeviceIdPutRequest{}
	return &this
}

// NewV1DiagnosticResetIpsecSessionDeviceIdPutRequestWithDefaults instantiates a new V1DiagnosticResetIpsecSessionDeviceIdPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DiagnosticResetIpsecSessionDeviceIdPutRequestWithDefaults() *V1DiagnosticResetIpsecSessionDeviceIdPutRequest {
	this := V1DiagnosticResetIpsecSessionDeviceIdPutRequest{}
	return &this
}

// GetAll3RdParty returns the All3RdParty field value if set, zero value otherwise.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) GetAll3RdParty() bool {
	if o == nil || IsNil(o.All3RdParty) {
		var ret bool
		return ret
	}
	return *o.All3RdParty
}

// GetAll3RdPartyOk returns a tuple with the All3RdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) GetAll3RdPartyOk() (*bool, bool) {
	if o == nil || IsNil(o.All3RdParty) {
		return nil, false
	}
	return o.All3RdParty, true
}

// HasAll3RdParty returns a boolean if a field has been set.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) HasAll3RdParty() bool {
	if o != nil && !IsNil(o.All3RdParty) {
		return true
	}

	return false
}

// SetAll3RdParty gets a reference to the given bool and assigns it to the All3RdParty field.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) SetAll3RdParty(v bool) {
	o.All3RdParty = &v
}

// GetAllControllers returns the AllControllers field value if set, zero value otherwise.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) GetAllControllers() bool {
	if o == nil || IsNil(o.AllControllers) {
		var ret bool
		return ret
	}
	return *o.AllControllers
}

// GetAllControllersOk returns a tuple with the AllControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) GetAllControllersOk() (*bool, bool) {
	if o == nil || IsNil(o.AllControllers) {
		return nil, false
	}
	return o.AllControllers, true
}

// HasAllControllers returns a boolean if a field has been set.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) HasAllControllers() bool {
	if o != nil && !IsNil(o.AllControllers) {
		return true
	}

	return false
}

// SetAllControllers gets a reference to the given bool and assigns it to the AllControllers field.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) SetAllControllers(v bool) {
	o.AllControllers = &v
}

// GetAllE2E returns the AllE2E field value if set, zero value otherwise.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) GetAllE2E() bool {
	if o == nil || IsNil(o.AllE2E) {
		var ret bool
		return ret
	}
	return *o.AllE2E
}

// GetAllE2EOk returns a tuple with the AllE2E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) GetAllE2EOk() (*bool, bool) {
	if o == nil || IsNil(o.AllE2E) {
		return nil, false
	}
	return o.AllE2E, true
}

// HasAllE2E returns a boolean if a field has been set.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) HasAllE2E() bool {
	if o != nil && !IsNil(o.AllE2E) {
		return true
	}

	return false
}

// SetAllE2E gets a reference to the given bool and assigns it to the AllE2E field.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) SetAllE2E(v bool) {
	o.AllE2E = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) GetVrf() []string {
	if o == nil || IsNil(o.Vrf) {
		var ret []string
		return ret
	}
	return o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) GetVrfOk() ([]string, bool) {
	if o == nil || IsNil(o.Vrf) {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) HasVrf() bool {
	if o != nil && !IsNil(o.Vrf) {
		return true
	}

	return false
}

// SetVrf gets a reference to the given []string and assigns it to the Vrf field.
func (o *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) SetVrf(v []string) {
	o.Vrf = v
}

func (o V1DiagnosticResetIpsecSessionDeviceIdPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DiagnosticResetIpsecSessionDeviceIdPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.All3RdParty) {
		toSerialize["all3RdParty"] = o.All3RdParty
	}
	if !IsNil(o.AllControllers) {
		toSerialize["allControllers"] = o.AllControllers
	}
	if !IsNil(o.AllE2E) {
		toSerialize["allE2E"] = o.AllE2E
	}
	if !IsNil(o.Vrf) {
		toSerialize["vrf"] = o.Vrf
	}
	return toSerialize, nil
}

type NullableV1DiagnosticResetIpsecSessionDeviceIdPutRequest struct {
	value *V1DiagnosticResetIpsecSessionDeviceIdPutRequest
	isSet bool
}

func (v NullableV1DiagnosticResetIpsecSessionDeviceIdPutRequest) Get() *V1DiagnosticResetIpsecSessionDeviceIdPutRequest {
	return v.value
}

func (v *NullableV1DiagnosticResetIpsecSessionDeviceIdPutRequest) Set(val *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DiagnosticResetIpsecSessionDeviceIdPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DiagnosticResetIpsecSessionDeviceIdPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DiagnosticResetIpsecSessionDeviceIdPutRequest(val *V1DiagnosticResetIpsecSessionDeviceIdPutRequest) *NullableV1DiagnosticResetIpsecSessionDeviceIdPutRequest {
	return &NullableV1DiagnosticResetIpsecSessionDeviceIdPutRequest{value: val, isSet: true}
}

func (v NullableV1DiagnosticResetIpsecSessionDeviceIdPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DiagnosticResetIpsecSessionDeviceIdPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


