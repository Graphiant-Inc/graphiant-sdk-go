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

// checks if the V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue{}

// V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue struct for V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue
type V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue struct {
	Neighbor *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor `json:"neighbor,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue {
	this := V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue {
	this := V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue{}
	return &this
}

// GetNeighbor returns the Neighbor field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) GetNeighbor() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor {
	if o == nil || IsNil(o.Neighbor) {
		var ret V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor
		return ret
	}
	return *o.Neighbor
}

// GetNeighborOk returns a tuple with the Neighbor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) GetNeighborOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor, bool) {
	if o == nil || IsNil(o.Neighbor) {
		return nil, false
	}
	return o.Neighbor, true
}

// HasNeighbor returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) HasNeighbor() bool {
	if o != nil && !IsNil(o.Neighbor) {
		return true
	}

	return false
}

// SetNeighbor gets a reference to the given V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor and assigns it to the Neighbor field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) SetNeighbor(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) {
	o.Neighbor = &v
}

func (o V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Neighbor) {
		toSerialize["neighbor"] = o.Neighbor
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue struct {
	value *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) Get() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) Set(val *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue(val *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue {
	return &NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


