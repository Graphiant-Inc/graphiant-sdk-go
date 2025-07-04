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

// checks if the V2IntegrationPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2IntegrationPostRequest{}

// V2IntegrationPostRequest struct for V2IntegrationPostRequest
type V2IntegrationPostRequest struct {
	IntegrationBody *V2IntegrationPostRequestIntegrationBody `json:"integrationBody,omitempty"`
}

// NewV2IntegrationPostRequest instantiates a new V2IntegrationPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2IntegrationPostRequest() *V2IntegrationPostRequest {
	this := V2IntegrationPostRequest{}
	return &this
}

// NewV2IntegrationPostRequestWithDefaults instantiates a new V2IntegrationPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2IntegrationPostRequestWithDefaults() *V2IntegrationPostRequest {
	this := V2IntegrationPostRequest{}
	return &this
}

// GetIntegrationBody returns the IntegrationBody field value if set, zero value otherwise.
func (o *V2IntegrationPostRequest) GetIntegrationBody() V2IntegrationPostRequestIntegrationBody {
	if o == nil || IsNil(o.IntegrationBody) {
		var ret V2IntegrationPostRequestIntegrationBody
		return ret
	}
	return *o.IntegrationBody
}

// GetIntegrationBodyOk returns a tuple with the IntegrationBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2IntegrationPostRequest) GetIntegrationBodyOk() (*V2IntegrationPostRequestIntegrationBody, bool) {
	if o == nil || IsNil(o.IntegrationBody) {
		return nil, false
	}
	return o.IntegrationBody, true
}

// HasIntegrationBody returns a boolean if a field has been set.
func (o *V2IntegrationPostRequest) HasIntegrationBody() bool {
	if o != nil && !IsNil(o.IntegrationBody) {
		return true
	}

	return false
}

// SetIntegrationBody gets a reference to the given V2IntegrationPostRequestIntegrationBody and assigns it to the IntegrationBody field.
func (o *V2IntegrationPostRequest) SetIntegrationBody(v V2IntegrationPostRequestIntegrationBody) {
	o.IntegrationBody = &v
}

func (o V2IntegrationPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2IntegrationPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntegrationBody) {
		toSerialize["integrationBody"] = o.IntegrationBody
	}
	return toSerialize, nil
}

type NullableV2IntegrationPostRequest struct {
	value *V2IntegrationPostRequest
	isSet bool
}

func (v NullableV2IntegrationPostRequest) Get() *V2IntegrationPostRequest {
	return v.value
}

func (v *NullableV2IntegrationPostRequest) Set(val *V2IntegrationPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2IntegrationPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2IntegrationPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2IntegrationPostRequest(val *V2IntegrationPostRequest) *NullableV2IntegrationPostRequest {
	return &NullableV2IntegrationPostRequest{value: val, isSet: true}
}

func (v NullableV2IntegrationPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2IntegrationPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


