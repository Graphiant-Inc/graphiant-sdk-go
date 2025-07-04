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

// checks if the V1ExtranetsB2bPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ExtranetsB2bPostRequest{}

// V1ExtranetsB2bPostRequest struct for V1ExtranetsB2bPostRequest
type V1ExtranetsB2bPostRequest struct {
	Policy *V1ExtranetsB2bPostRequestPolicy `json:"policy,omitempty"`
	ServiceName *string `json:"serviceName,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewV1ExtranetsB2bPostRequest instantiates a new V1ExtranetsB2bPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ExtranetsB2bPostRequest() *V1ExtranetsB2bPostRequest {
	this := V1ExtranetsB2bPostRequest{}
	return &this
}

// NewV1ExtranetsB2bPostRequestWithDefaults instantiates a new V1ExtranetsB2bPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ExtranetsB2bPostRequestWithDefaults() *V1ExtranetsB2bPostRequest {
	this := V1ExtranetsB2bPostRequest{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *V1ExtranetsB2bPostRequest) GetPolicy() V1ExtranetsB2bPostRequestPolicy {
	if o == nil || IsNil(o.Policy) {
		var ret V1ExtranetsB2bPostRequestPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsB2bPostRequest) GetPolicyOk() (*V1ExtranetsB2bPostRequestPolicy, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *V1ExtranetsB2bPostRequest) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given V1ExtranetsB2bPostRequestPolicy and assigns it to the Policy field.
func (o *V1ExtranetsB2bPostRequest) SetPolicy(v V1ExtranetsB2bPostRequestPolicy) {
	o.Policy = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *V1ExtranetsB2bPostRequest) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsB2bPostRequest) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *V1ExtranetsB2bPostRequest) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *V1ExtranetsB2bPostRequest) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1ExtranetsB2bPostRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsB2bPostRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1ExtranetsB2bPostRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1ExtranetsB2bPostRequest) SetType(v string) {
	o.Type = &v
}

func (o V1ExtranetsB2bPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ExtranetsB2bPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.ServiceName) {
		toSerialize["serviceName"] = o.ServiceName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableV1ExtranetsB2bPostRequest struct {
	value *V1ExtranetsB2bPostRequest
	isSet bool
}

func (v NullableV1ExtranetsB2bPostRequest) Get() *V1ExtranetsB2bPostRequest {
	return v.value
}

func (v *NullableV1ExtranetsB2bPostRequest) Set(val *V1ExtranetsB2bPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ExtranetsB2bPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ExtranetsB2bPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ExtranetsB2bPostRequest(val *V1ExtranetsB2bPostRequest) *NullableV1ExtranetsB2bPostRequest {
	return &NullableV1ExtranetsB2bPostRequest{value: val, isSet: true}
}

func (v NullableV1ExtranetsB2bPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ExtranetsB2bPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


