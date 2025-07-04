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

// checks if the V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType{}

// V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType struct for V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType
type V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType struct {
	RouteType *string `json:"routeType,omitempty"`
}

// NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType {
	this := V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType{}
	return &this
}

// NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteTypeWithDefaults instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteTypeWithDefaults() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType {
	this := V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType{}
	return &this
}

// GetRouteType returns the RouteType field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) GetRouteType() string {
	if o == nil || IsNil(o.RouteType) {
		var ret string
		return ret
	}
	return *o.RouteType
}

// GetRouteTypeOk returns a tuple with the RouteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) GetRouteTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RouteType) {
		return nil, false
	}
	return o.RouteType, true
}

// HasRouteType returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) HasRouteType() bool {
	if o != nil && !IsNil(o.RouteType) {
		return true
	}

	return false
}

// SetRouteType gets a reference to the given string and assigns it to the RouteType field.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) SetRouteType(v string) {
	o.RouteType = &v
}

func (o V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RouteType) {
		toSerialize["routeType"] = o.RouteType
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType struct {
	value *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) Get() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) Set(val *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType(val *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType {
	return &NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


