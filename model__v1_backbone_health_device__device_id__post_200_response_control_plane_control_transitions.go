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

// checks if the V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions{}

// V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions struct for V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions
type V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions struct {
	Transitions []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner `json:"transitions,omitempty"`
}

// NewV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions() *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions {
	this := V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions{}
	return &this
}

// NewV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsWithDefaults instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsWithDefaults() *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions {
	this := V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions{}
	return &this
}

// GetTransitions returns the Transitions field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) GetTransitions() []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner {
	if o == nil || IsNil(o.Transitions) {
		var ret []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner
		return ret
	}
	return o.Transitions
}

// GetTransitionsOk returns a tuple with the Transitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) GetTransitionsOk() ([]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner, bool) {
	if o == nil || IsNil(o.Transitions) {
		return nil, false
	}
	return o.Transitions, true
}

// HasTransitions returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) HasTransitions() bool {
	if o != nil && !IsNil(o.Transitions) {
		return true
	}

	return false
}

// SetTransitions gets a reference to the given []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner and assigns it to the Transitions field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) SetTransitions(v []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) {
	o.Transitions = v
}

func (o V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transitions) {
		toSerialize["transitions"] = o.Transitions
	}
	return toSerialize, nil
}

type NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions struct {
	value *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions
	isSet bool
}

func (v NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) Get() *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions {
	return v.value
}

func (v *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) Set(val *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions(val *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions {
	return &NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions{value: val, isSet: true}
}

func (v NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


