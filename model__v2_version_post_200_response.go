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

// checks if the V2VersionPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2VersionPost200Response{}

// V2VersionPost200Response struct for V2VersionPost200Response
type V2VersionPost200Response struct {
	VersionId *string `json:"versionId,omitempty"`
}

// NewV2VersionPost200Response instantiates a new V2VersionPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2VersionPost200Response() *V2VersionPost200Response {
	this := V2VersionPost200Response{}
	return &this
}

// NewV2VersionPost200ResponseWithDefaults instantiates a new V2VersionPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2VersionPost200ResponseWithDefaults() *V2VersionPost200Response {
	this := V2VersionPost200Response{}
	return &this
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *V2VersionPost200Response) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VersionPost200Response) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *V2VersionPost200Response) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *V2VersionPost200Response) SetVersionId(v string) {
	o.VersionId = &v
}

func (o V2VersionPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2VersionPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VersionId) {
		toSerialize["versionId"] = o.VersionId
	}
	return toSerialize, nil
}

type NullableV2VersionPost200Response struct {
	value *V2VersionPost200Response
	isSet bool
}

func (v NullableV2VersionPost200Response) Get() *V2VersionPost200Response {
	return v.value
}

func (v *NullableV2VersionPost200Response) Set(val *V2VersionPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2VersionPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2VersionPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2VersionPost200Response(val *V2VersionPost200Response) *NullableV2VersionPost200Response {
	return &NullableV2VersionPost200Response{value: val, isSet: true}
}

func (v NullableV2VersionPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2VersionPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


