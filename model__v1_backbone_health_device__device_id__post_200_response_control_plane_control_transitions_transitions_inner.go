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

// checks if the V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner{}

// V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner struct for V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner
type V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner struct {
	Name *string `json:"name,omitempty"`
	Transitions []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner `json:"transitions,omitempty"`
}

// NewV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner() *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner {
	this := V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner{}
	return &this
}

// NewV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerWithDefaults instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerWithDefaults() *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner {
	this := V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) SetName(v string) {
	o.Name = &v
}

// GetTransitions returns the Transitions field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) GetTransitions() []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner {
	if o == nil || IsNil(o.Transitions) {
		var ret []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner
		return ret
	}
	return o.Transitions
}

// GetTransitionsOk returns a tuple with the Transitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) GetTransitionsOk() ([]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner, bool) {
	if o == nil || IsNil(o.Transitions) {
		return nil, false
	}
	return o.Transitions, true
}

// HasTransitions returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) HasTransitions() bool {
	if o != nil && !IsNil(o.Transitions) {
		return true
	}

	return false
}

// SetTransitions gets a reference to the given []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner and assigns it to the Transitions field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) SetTransitions(v []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner) {
	o.Transitions = v
}

func (o V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Transitions) {
		toSerialize["transitions"] = o.Transitions
	}
	return toSerialize, nil
}

type NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner struct {
	value *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner
	isSet bool
}

func (v NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) Get() *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner {
	return v.value
}

func (v *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) Set(val *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner(val *V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner {
	return &NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner{value: val, isSet: true}
}

func (v NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


