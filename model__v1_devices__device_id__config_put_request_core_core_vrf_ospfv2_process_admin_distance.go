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

// checks if the V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance{}

// V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance struct for V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance
type V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance struct {
	AdminDistance *int32 `json:"adminDistance,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance {
	this := V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistanceWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistanceWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance {
	this := V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance{}
	return &this
}

// GetAdminDistance returns the AdminDistance field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) GetAdminDistance() int32 {
	if o == nil || IsNil(o.AdminDistance) {
		var ret int32
		return ret
	}
	return *o.AdminDistance
}

// GetAdminDistanceOk returns a tuple with the AdminDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) GetAdminDistanceOk() (*int32, bool) {
	if o == nil || IsNil(o.AdminDistance) {
		return nil, false
	}
	return o.AdminDistance, true
}

// HasAdminDistance returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) HasAdminDistance() bool {
	if o != nil && !IsNil(o.AdminDistance) {
		return true
	}

	return false
}

// SetAdminDistance gets a reference to the given int32 and assigns it to the AdminDistance field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) SetAdminDistance(v int32) {
	o.AdminDistance = &v
}

func (o V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminDistance) {
		toSerialize["adminDistance"] = o.AdminDistance
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance struct {
	value *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) Get() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) Set(val *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance(val *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance {
	return &NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


