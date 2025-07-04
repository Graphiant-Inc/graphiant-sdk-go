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

// checks if the V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos{}

// V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos struct for V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos
type V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos struct {
	FlexAlgoNames []string `json:"flexAlgoNames,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos {
	this := V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgosWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgosWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos {
	this := V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos{}
	return &this
}

// GetFlexAlgoNames returns the FlexAlgoNames field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) GetFlexAlgoNames() []string {
	if o == nil || IsNil(o.FlexAlgoNames) {
		var ret []string
		return ret
	}
	return o.FlexAlgoNames
}

// GetFlexAlgoNamesOk returns a tuple with the FlexAlgoNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) GetFlexAlgoNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.FlexAlgoNames) {
		return nil, false
	}
	return o.FlexAlgoNames, true
}

// HasFlexAlgoNames returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) HasFlexAlgoNames() bool {
	if o != nil && !IsNil(o.FlexAlgoNames) {
		return true
	}

	return false
}

// SetFlexAlgoNames gets a reference to the given []string and assigns it to the FlexAlgoNames field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) SetFlexAlgoNames(v []string) {
	o.FlexAlgoNames = v
}

func (o V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlexAlgoNames) {
		toSerialize["flexAlgoNames"] = o.FlexAlgoNames
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos struct {
	value *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) Get() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) Set(val *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos(val *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos {
	return &NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


