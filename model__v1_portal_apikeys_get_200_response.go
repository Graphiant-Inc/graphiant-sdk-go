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

// checks if the V1PortalApikeysGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PortalApikeysGet200Response{}

// V1PortalApikeysGet200Response struct for V1PortalApikeysGet200Response
type V1PortalApikeysGet200Response struct {
	ApiKeyInfo []V1PortalApikeysGet200ResponseApiKeyInfoInner `json:"apiKeyInfo,omitempty"`
}

// NewV1PortalApikeysGet200Response instantiates a new V1PortalApikeysGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PortalApikeysGet200Response() *V1PortalApikeysGet200Response {
	this := V1PortalApikeysGet200Response{}
	return &this
}

// NewV1PortalApikeysGet200ResponseWithDefaults instantiates a new V1PortalApikeysGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PortalApikeysGet200ResponseWithDefaults() *V1PortalApikeysGet200Response {
	this := V1PortalApikeysGet200Response{}
	return &this
}

// GetApiKeyInfo returns the ApiKeyInfo field value if set, zero value otherwise.
func (o *V1PortalApikeysGet200Response) GetApiKeyInfo() []V1PortalApikeysGet200ResponseApiKeyInfoInner {
	if o == nil || IsNil(o.ApiKeyInfo) {
		var ret []V1PortalApikeysGet200ResponseApiKeyInfoInner
		return ret
	}
	return o.ApiKeyInfo
}

// GetApiKeyInfoOk returns a tuple with the ApiKeyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PortalApikeysGet200Response) GetApiKeyInfoOk() ([]V1PortalApikeysGet200ResponseApiKeyInfoInner, bool) {
	if o == nil || IsNil(o.ApiKeyInfo) {
		return nil, false
	}
	return o.ApiKeyInfo, true
}

// HasApiKeyInfo returns a boolean if a field has been set.
func (o *V1PortalApikeysGet200Response) HasApiKeyInfo() bool {
	if o != nil && !IsNil(o.ApiKeyInfo) {
		return true
	}

	return false
}

// SetApiKeyInfo gets a reference to the given []V1PortalApikeysGet200ResponseApiKeyInfoInner and assigns it to the ApiKeyInfo field.
func (o *V1PortalApikeysGet200Response) SetApiKeyInfo(v []V1PortalApikeysGet200ResponseApiKeyInfoInner) {
	o.ApiKeyInfo = v
}

func (o V1PortalApikeysGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PortalApikeysGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKeyInfo) {
		toSerialize["apiKeyInfo"] = o.ApiKeyInfo
	}
	return toSerialize, nil
}

type NullableV1PortalApikeysGet200Response struct {
	value *V1PortalApikeysGet200Response
	isSet bool
}

func (v NullableV1PortalApikeysGet200Response) Get() *V1PortalApikeysGet200Response {
	return v.value
}

func (v *NullableV1PortalApikeysGet200Response) Set(val *V1PortalApikeysGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PortalApikeysGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PortalApikeysGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PortalApikeysGet200Response(val *V1PortalApikeysGet200Response) *NullableV1PortalApikeysGet200Response {
	return &NullableV1PortalApikeysGet200Response{value: val, isSet: true}
}

func (v NullableV1PortalApikeysGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PortalApikeysGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


