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

// checks if the V1PortalApikeysGet200ResponseApiKeyInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PortalApikeysGet200ResponseApiKeyInfoInner{}

// V1PortalApikeysGet200ResponseApiKeyInfoInner struct for V1PortalApikeysGet200ResponseApiKeyInfoInner
type V1PortalApikeysGet200ResponseApiKeyInfoInner struct {
	ExpiryTs *V1AlarmHistoryGet200ResponseHistoryInnerTime `json:"expiryTs,omitempty"`
	GcsName *string `json:"gcsName,omitempty"`
}

// NewV1PortalApikeysGet200ResponseApiKeyInfoInner instantiates a new V1PortalApikeysGet200ResponseApiKeyInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PortalApikeysGet200ResponseApiKeyInfoInner() *V1PortalApikeysGet200ResponseApiKeyInfoInner {
	this := V1PortalApikeysGet200ResponseApiKeyInfoInner{}
	return &this
}

// NewV1PortalApikeysGet200ResponseApiKeyInfoInnerWithDefaults instantiates a new V1PortalApikeysGet200ResponseApiKeyInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PortalApikeysGet200ResponseApiKeyInfoInnerWithDefaults() *V1PortalApikeysGet200ResponseApiKeyInfoInner {
	this := V1PortalApikeysGet200ResponseApiKeyInfoInner{}
	return &this
}

// GetExpiryTs returns the ExpiryTs field value if set, zero value otherwise.
func (o *V1PortalApikeysGet200ResponseApiKeyInfoInner) GetExpiryTs() V1AlarmHistoryGet200ResponseHistoryInnerTime {
	if o == nil || IsNil(o.ExpiryTs) {
		var ret V1AlarmHistoryGet200ResponseHistoryInnerTime
		return ret
	}
	return *o.ExpiryTs
}

// GetExpiryTsOk returns a tuple with the ExpiryTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PortalApikeysGet200ResponseApiKeyInfoInner) GetExpiryTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool) {
	if o == nil || IsNil(o.ExpiryTs) {
		return nil, false
	}
	return o.ExpiryTs, true
}

// HasExpiryTs returns a boolean if a field has been set.
func (o *V1PortalApikeysGet200ResponseApiKeyInfoInner) HasExpiryTs() bool {
	if o != nil && !IsNil(o.ExpiryTs) {
		return true
	}

	return false
}

// SetExpiryTs gets a reference to the given V1AlarmHistoryGet200ResponseHistoryInnerTime and assigns it to the ExpiryTs field.
func (o *V1PortalApikeysGet200ResponseApiKeyInfoInner) SetExpiryTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime) {
	o.ExpiryTs = &v
}

// GetGcsName returns the GcsName field value if set, zero value otherwise.
func (o *V1PortalApikeysGet200ResponseApiKeyInfoInner) GetGcsName() string {
	if o == nil || IsNil(o.GcsName) {
		var ret string
		return ret
	}
	return *o.GcsName
}

// GetGcsNameOk returns a tuple with the GcsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PortalApikeysGet200ResponseApiKeyInfoInner) GetGcsNameOk() (*string, bool) {
	if o == nil || IsNil(o.GcsName) {
		return nil, false
	}
	return o.GcsName, true
}

// HasGcsName returns a boolean if a field has been set.
func (o *V1PortalApikeysGet200ResponseApiKeyInfoInner) HasGcsName() bool {
	if o != nil && !IsNil(o.GcsName) {
		return true
	}

	return false
}

// SetGcsName gets a reference to the given string and assigns it to the GcsName field.
func (o *V1PortalApikeysGet200ResponseApiKeyInfoInner) SetGcsName(v string) {
	o.GcsName = &v
}

func (o V1PortalApikeysGet200ResponseApiKeyInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PortalApikeysGet200ResponseApiKeyInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiryTs) {
		toSerialize["expiryTs"] = o.ExpiryTs
	}
	if !IsNil(o.GcsName) {
		toSerialize["gcsName"] = o.GcsName
	}
	return toSerialize, nil
}

type NullableV1PortalApikeysGet200ResponseApiKeyInfoInner struct {
	value *V1PortalApikeysGet200ResponseApiKeyInfoInner
	isSet bool
}

func (v NullableV1PortalApikeysGet200ResponseApiKeyInfoInner) Get() *V1PortalApikeysGet200ResponseApiKeyInfoInner {
	return v.value
}

func (v *NullableV1PortalApikeysGet200ResponseApiKeyInfoInner) Set(val *V1PortalApikeysGet200ResponseApiKeyInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PortalApikeysGet200ResponseApiKeyInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PortalApikeysGet200ResponseApiKeyInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PortalApikeysGet200ResponseApiKeyInfoInner(val *V1PortalApikeysGet200ResponseApiKeyInfoInner) *NullableV1PortalApikeysGet200ResponseApiKeyInfoInner {
	return &NullableV1PortalApikeysGet200ResponseApiKeyInfoInner{value: val, isSet: true}
}

func (v NullableV1PortalApikeysGet200ResponseApiKeyInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PortalApikeysGet200ResponseApiKeyInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


