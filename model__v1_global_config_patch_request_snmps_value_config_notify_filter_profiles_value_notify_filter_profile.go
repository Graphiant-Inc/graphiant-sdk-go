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

// checks if the V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile{}

// V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile struct for V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile
type V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile struct {
	IncludeExcludeList *map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue `json:"includeExcludeList,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile() *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile {
	this := V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile{}
	return &this
}

// NewV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileWithDefaults instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileWithDefaults() *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile {
	this := V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile{}
	return &this
}

// GetIncludeExcludeList returns the IncludeExcludeList field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) GetIncludeExcludeList() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue {
	if o == nil || IsNil(o.IncludeExcludeList) {
		var ret map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue
		return ret
	}
	return *o.IncludeExcludeList
}

// GetIncludeExcludeListOk returns a tuple with the IncludeExcludeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) GetIncludeExcludeListOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue, bool) {
	if o == nil || IsNil(o.IncludeExcludeList) {
		return nil, false
	}
	return o.IncludeExcludeList, true
}

// HasIncludeExcludeList returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) HasIncludeExcludeList() bool {
	if o != nil && !IsNil(o.IncludeExcludeList) {
		return true
	}

	return false
}

// SetIncludeExcludeList gets a reference to the given map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue and assigns it to the IncludeExcludeList field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) SetIncludeExcludeList(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfileIncludeExcludeListValue) {
	o.IncludeExcludeList = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) SetName(v string) {
	o.Name = &v
}

func (o V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludeExcludeList) {
		toSerialize["includeExcludeList"] = o.IncludeExcludeList
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile struct {
	value *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) Get() *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) Set(val *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile(val *V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) *NullableV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile {
	return &NullableV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValueNotifyFilterProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


