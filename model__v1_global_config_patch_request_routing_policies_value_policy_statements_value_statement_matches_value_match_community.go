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

// checks if the V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity{}

// V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity struct for V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity
type V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity struct {
	Community *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunityCommunity `json:"community,omitempty"`
}

// NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity {
	this := V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity{}
	return &this
}

// NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunityWithDefaults instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunityWithDefaults() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity {
	this := V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity{}
	return &this
}

// GetCommunity returns the Community field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) GetCommunity() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunityCommunity {
	if o == nil || IsNil(o.Community) {
		var ret V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunityCommunity
		return ret
	}
	return *o.Community
}

// GetCommunityOk returns a tuple with the Community field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) GetCommunityOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunityCommunity, bool) {
	if o == nil || IsNil(o.Community) {
		return nil, false
	}
	return o.Community, true
}

// HasCommunity returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) HasCommunity() bool {
	if o != nil && !IsNil(o.Community) {
		return true
	}

	return false
}

// SetCommunity gets a reference to the given V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunityCommunity and assigns it to the Community field.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) SetCommunity(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunityCommunity) {
	o.Community = &v
}

func (o V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Community) {
		toSerialize["community"] = o.Community
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity struct {
	value *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) Get() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) Set(val *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity(val *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity {
	return &NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


