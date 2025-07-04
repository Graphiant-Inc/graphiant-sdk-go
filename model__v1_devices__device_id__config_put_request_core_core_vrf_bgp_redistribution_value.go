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

// checks if the V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue{}

// V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue struct for V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue
type V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue struct {
	Protocol *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue `json:"protocol,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue {
	this := V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValueWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValueWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue {
	this := V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) GetProtocol() V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue {
	if o == nil || IsNil(o.Protocol) {
		var ret V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) GetProtocolOk() (*V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue and assigns it to the Protocol field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) SetProtocol(v V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue) {
	o.Protocol = &v
}

func (o V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue struct {
	value *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) Get() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) Set(val *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue(val *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue {
	return &NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


