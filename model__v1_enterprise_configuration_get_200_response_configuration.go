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

// checks if the V1EnterpriseConfigurationGet200ResponseConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EnterpriseConfigurationGet200ResponseConfiguration{}

// V1EnterpriseConfigurationGet200ResponseConfiguration struct for V1EnterpriseConfigurationGet200ResponseConfiguration
type V1EnterpriseConfigurationGet200ResponseConfiguration struct {
	Encryption *string `json:"encryption,omitempty"`
	UseEncryption *bool `json:"useEncryption,omitempty"`
}

// NewV1EnterpriseConfigurationGet200ResponseConfiguration instantiates a new V1EnterpriseConfigurationGet200ResponseConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EnterpriseConfigurationGet200ResponseConfiguration() *V1EnterpriseConfigurationGet200ResponseConfiguration {
	this := V1EnterpriseConfigurationGet200ResponseConfiguration{}
	return &this
}

// NewV1EnterpriseConfigurationGet200ResponseConfigurationWithDefaults instantiates a new V1EnterpriseConfigurationGet200ResponseConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EnterpriseConfigurationGet200ResponseConfigurationWithDefaults() *V1EnterpriseConfigurationGet200ResponseConfiguration {
	this := V1EnterpriseConfigurationGet200ResponseConfiguration{}
	return &this
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *V1EnterpriseConfigurationGet200ResponseConfiguration) GetEncryption() string {
	if o == nil || IsNil(o.Encryption) {
		var ret string
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnterpriseConfigurationGet200ResponseConfiguration) GetEncryptionOk() (*string, bool) {
	if o == nil || IsNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *V1EnterpriseConfigurationGet200ResponseConfiguration) HasEncryption() bool {
	if o != nil && !IsNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given string and assigns it to the Encryption field.
func (o *V1EnterpriseConfigurationGet200ResponseConfiguration) SetEncryption(v string) {
	o.Encryption = &v
}

// GetUseEncryption returns the UseEncryption field value if set, zero value otherwise.
func (o *V1EnterpriseConfigurationGet200ResponseConfiguration) GetUseEncryption() bool {
	if o == nil || IsNil(o.UseEncryption) {
		var ret bool
		return ret
	}
	return *o.UseEncryption
}

// GetUseEncryptionOk returns a tuple with the UseEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnterpriseConfigurationGet200ResponseConfiguration) GetUseEncryptionOk() (*bool, bool) {
	if o == nil || IsNil(o.UseEncryption) {
		return nil, false
	}
	return o.UseEncryption, true
}

// HasUseEncryption returns a boolean if a field has been set.
func (o *V1EnterpriseConfigurationGet200ResponseConfiguration) HasUseEncryption() bool {
	if o != nil && !IsNil(o.UseEncryption) {
		return true
	}

	return false
}

// SetUseEncryption gets a reference to the given bool and assigns it to the UseEncryption field.
func (o *V1EnterpriseConfigurationGet200ResponseConfiguration) SetUseEncryption(v bool) {
	o.UseEncryption = &v
}

func (o V1EnterpriseConfigurationGet200ResponseConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EnterpriseConfigurationGet200ResponseConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Encryption) {
		toSerialize["encryption"] = o.Encryption
	}
	if !IsNil(o.UseEncryption) {
		toSerialize["useEncryption"] = o.UseEncryption
	}
	return toSerialize, nil
}

type NullableV1EnterpriseConfigurationGet200ResponseConfiguration struct {
	value *V1EnterpriseConfigurationGet200ResponseConfiguration
	isSet bool
}

func (v NullableV1EnterpriseConfigurationGet200ResponseConfiguration) Get() *V1EnterpriseConfigurationGet200ResponseConfiguration {
	return v.value
}

func (v *NullableV1EnterpriseConfigurationGet200ResponseConfiguration) Set(val *V1EnterpriseConfigurationGet200ResponseConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EnterpriseConfigurationGet200ResponseConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EnterpriseConfigurationGet200ResponseConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EnterpriseConfigurationGet200ResponseConfiguration(val *V1EnterpriseConfigurationGet200ResponseConfiguration) *NullableV1EnterpriseConfigurationGet200ResponseConfiguration {
	return &NullableV1EnterpriseConfigurationGet200ResponseConfiguration{value: val, isSet: true}
}

func (v NullableV1EnterpriseConfigurationGet200ResponseConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EnterpriseConfigurationGet200ResponseConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


