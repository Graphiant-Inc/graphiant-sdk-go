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

// checks if the V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue{}

// V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue struct for V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue
type V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue struct {
	Access *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValueAccess `json:"access,omitempty"`
}

// NewV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue() *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue {
	this := V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue{}
	return &this
}

// NewV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValueWithDefaults instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValueWithDefaults() *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue {
	this := V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) GetAccess() V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValueAccess {
	if o == nil || IsNil(o.Access) {
		var ret V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValueAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) GetAccessOk() (*V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValueAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValueAccess and assigns it to the Access field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) SetAccess(v V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValueAccess) {
	o.Access = &v
}

func (o V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue struct {
	value *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) Get() *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) Set(val *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue(val *V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) *NullableV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue {
	return &NullableV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValueGroupAccessesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


