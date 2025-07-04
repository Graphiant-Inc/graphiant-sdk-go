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

// checks if the V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response{}

// V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response struct for V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response
type V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response struct {
	VrfAssociations *map[string]V1DevicesDeviceIdCircuitsVrfAssociationsGet200ResponseVrfAssociationsValue `json:"vrfAssociations,omitempty"`
}

// NewV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response instantiates a new V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response() *V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response {
	this := V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response{}
	return &this
}

// NewV1DevicesDeviceIdCircuitsVrfAssociationsGet200ResponseWithDefaults instantiates a new V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdCircuitsVrfAssociationsGet200ResponseWithDefaults() *V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response {
	this := V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response{}
	return &this
}

// GetVrfAssociations returns the VrfAssociations field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) GetVrfAssociations() map[string]V1DevicesDeviceIdCircuitsVrfAssociationsGet200ResponseVrfAssociationsValue {
	if o == nil || IsNil(o.VrfAssociations) {
		var ret map[string]V1DevicesDeviceIdCircuitsVrfAssociationsGet200ResponseVrfAssociationsValue
		return ret
	}
	return *o.VrfAssociations
}

// GetVrfAssociationsOk returns a tuple with the VrfAssociations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) GetVrfAssociationsOk() (*map[string]V1DevicesDeviceIdCircuitsVrfAssociationsGet200ResponseVrfAssociationsValue, bool) {
	if o == nil || IsNil(o.VrfAssociations) {
		return nil, false
	}
	return o.VrfAssociations, true
}

// HasVrfAssociations returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) HasVrfAssociations() bool {
	if o != nil && !IsNil(o.VrfAssociations) {
		return true
	}

	return false
}

// SetVrfAssociations gets a reference to the given map[string]V1DevicesDeviceIdCircuitsVrfAssociationsGet200ResponseVrfAssociationsValue and assigns it to the VrfAssociations field.
func (o *V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) SetVrfAssociations(v map[string]V1DevicesDeviceIdCircuitsVrfAssociationsGet200ResponseVrfAssociationsValue) {
	o.VrfAssociations = &v
}

func (o V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VrfAssociations) {
		toSerialize["vrfAssociations"] = o.VrfAssociations
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response struct {
	value *V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response
	isSet bool
}

func (v NullableV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) Get() *V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response {
	return v.value
}

func (v *NullableV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) Set(val *V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response(val *V1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) *NullableV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response {
	return &NullableV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdCircuitsVrfAssociationsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


