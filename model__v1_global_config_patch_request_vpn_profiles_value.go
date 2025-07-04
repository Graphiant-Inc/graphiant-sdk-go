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

// checks if the V1GlobalConfigPatchRequestVpnProfilesValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestVpnProfilesValue{}

// V1GlobalConfigPatchRequestVpnProfilesValue struct for V1GlobalConfigPatchRequestVpnProfilesValue
type V1GlobalConfigPatchRequestVpnProfilesValue struct {
	VpnProfile *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile `json:"vpnProfile,omitempty"`
}

// NewV1GlobalConfigPatchRequestVpnProfilesValue instantiates a new V1GlobalConfigPatchRequestVpnProfilesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestVpnProfilesValue() *V1GlobalConfigPatchRequestVpnProfilesValue {
	this := V1GlobalConfigPatchRequestVpnProfilesValue{}
	return &this
}

// NewV1GlobalConfigPatchRequestVpnProfilesValueWithDefaults instantiates a new V1GlobalConfigPatchRequestVpnProfilesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestVpnProfilesValueWithDefaults() *V1GlobalConfigPatchRequestVpnProfilesValue {
	this := V1GlobalConfigPatchRequestVpnProfilesValue{}
	return &this
}

// GetVpnProfile returns the VpnProfile field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestVpnProfilesValue) GetVpnProfile() V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile {
	if o == nil || IsNil(o.VpnProfile) {
		var ret V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile
		return ret
	}
	return *o.VpnProfile
}

// GetVpnProfileOk returns a tuple with the VpnProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestVpnProfilesValue) GetVpnProfileOk() (*V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile, bool) {
	if o == nil || IsNil(o.VpnProfile) {
		return nil, false
	}
	return o.VpnProfile, true
}

// HasVpnProfile returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestVpnProfilesValue) HasVpnProfile() bool {
	if o != nil && !IsNil(o.VpnProfile) {
		return true
	}

	return false
}

// SetVpnProfile gets a reference to the given V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile and assigns it to the VpnProfile field.
func (o *V1GlobalConfigPatchRequestVpnProfilesValue) SetVpnProfile(v V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) {
	o.VpnProfile = &v
}

func (o V1GlobalConfigPatchRequestVpnProfilesValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestVpnProfilesValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VpnProfile) {
		toSerialize["vpnProfile"] = o.VpnProfile
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestVpnProfilesValue struct {
	value *V1GlobalConfigPatchRequestVpnProfilesValue
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestVpnProfilesValue) Get() *V1GlobalConfigPatchRequestVpnProfilesValue {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestVpnProfilesValue) Set(val *V1GlobalConfigPatchRequestVpnProfilesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestVpnProfilesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestVpnProfilesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestVpnProfilesValue(val *V1GlobalConfigPatchRequestVpnProfilesValue) *NullableV1GlobalConfigPatchRequestVpnProfilesValue {
	return &NullableV1GlobalConfigPatchRequestVpnProfilesValue{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestVpnProfilesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestVpnProfilesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


