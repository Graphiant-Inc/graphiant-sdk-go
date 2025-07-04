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

// checks if the V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue{}

// V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue struct for V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue
type V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue struct {
	Interface *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface `json:"interface,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue {
	this := V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue {
	this := V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) GetInterface() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface {
	if o == nil || IsNil(o.Interface) {
		var ret V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) GetInterfaceOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface and assigns it to the Interface field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) SetInterface(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) {
	o.Interface = &v
}

func (o V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue struct {
	value *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) Get() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) Set(val *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue(val *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue {
	return &NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


