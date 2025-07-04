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

// checks if the V1AccountPasswordPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1AccountPasswordPatchRequest{}

// V1AccountPasswordPatchRequest struct for V1AccountPasswordPatchRequest
type V1AccountPasswordPatchRequest struct {
	OldPassword *string `json:"oldPassword,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewV1AccountPasswordPatchRequest instantiates a new V1AccountPasswordPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AccountPasswordPatchRequest() *V1AccountPasswordPatchRequest {
	this := V1AccountPasswordPatchRequest{}
	return &this
}

// NewV1AccountPasswordPatchRequestWithDefaults instantiates a new V1AccountPasswordPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AccountPasswordPatchRequestWithDefaults() *V1AccountPasswordPatchRequest {
	this := V1AccountPasswordPatchRequest{}
	return &this
}

// GetOldPassword returns the OldPassword field value if set, zero value otherwise.
func (o *V1AccountPasswordPatchRequest) GetOldPassword() string {
	if o == nil || IsNil(o.OldPassword) {
		var ret string
		return ret
	}
	return *o.OldPassword
}

// GetOldPasswordOk returns a tuple with the OldPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AccountPasswordPatchRequest) GetOldPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.OldPassword) {
		return nil, false
	}
	return o.OldPassword, true
}

// HasOldPassword returns a boolean if a field has been set.
func (o *V1AccountPasswordPatchRequest) HasOldPassword() bool {
	if o != nil && !IsNil(o.OldPassword) {
		return true
	}

	return false
}

// SetOldPassword gets a reference to the given string and assigns it to the OldPassword field.
func (o *V1AccountPasswordPatchRequest) SetOldPassword(v string) {
	o.OldPassword = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *V1AccountPasswordPatchRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AccountPasswordPatchRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *V1AccountPasswordPatchRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *V1AccountPasswordPatchRequest) SetPassword(v string) {
	o.Password = &v
}

func (o V1AccountPasswordPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1AccountPasswordPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OldPassword) {
		toSerialize["oldPassword"] = o.OldPassword
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableV1AccountPasswordPatchRequest struct {
	value *V1AccountPasswordPatchRequest
	isSet bool
}

func (v NullableV1AccountPasswordPatchRequest) Get() *V1AccountPasswordPatchRequest {
	return v.value
}

func (v *NullableV1AccountPasswordPatchRequest) Set(val *V1AccountPasswordPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AccountPasswordPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AccountPasswordPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AccountPasswordPatchRequest(val *V1AccountPasswordPatchRequest) *NullableV1AccountPasswordPatchRequest {
	return &NullableV1AccountPasswordPatchRequest{value: val, isSet: true}
}

func (v NullableV1AccountPasswordPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AccountPasswordPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


