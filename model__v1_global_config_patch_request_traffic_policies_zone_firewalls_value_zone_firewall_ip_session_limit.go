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

// checks if the V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit{}

// V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit struct for V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit
type V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit struct {
	Limit *int32 `json:"limit,omitempty"`
}

// NewV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit() *V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit {
	this := V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit{}
	return &this
}

// NewV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimitWithDefaults instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimitWithDefaults() *V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit {
	this := V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) SetLimit(v int32) {
	o.Limit = &v
}

func (o V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit struct {
	value *V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) Get() *V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) Set(val *V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit(val *V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) *NullableV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit {
	return &NullableV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValueZoneFirewallIpSessionLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


