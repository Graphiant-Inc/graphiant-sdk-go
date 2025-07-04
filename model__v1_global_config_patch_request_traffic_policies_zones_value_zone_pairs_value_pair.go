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

// checks if the V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair{}

// V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair struct for V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair
type V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair struct {
	Outside *string `json:"outside,omitempty"`
	Ruleset *string `json:"ruleset,omitempty"`
	TcpProtection *bool `json:"tcpProtection,omitempty"`
}

// NewV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair() *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair {
	this := V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair{}
	return &this
}

// NewV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePairWithDefaults instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePairWithDefaults() *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair {
	this := V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair{}
	return &this
}

// GetOutside returns the Outside field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) GetOutside() string {
	if o == nil || IsNil(o.Outside) {
		var ret string
		return ret
	}
	return *o.Outside
}

// GetOutsideOk returns a tuple with the Outside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) GetOutsideOk() (*string, bool) {
	if o == nil || IsNil(o.Outside) {
		return nil, false
	}
	return o.Outside, true
}

// HasOutside returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) HasOutside() bool {
	if o != nil && !IsNil(o.Outside) {
		return true
	}

	return false
}

// SetOutside gets a reference to the given string and assigns it to the Outside field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) SetOutside(v string) {
	o.Outside = &v
}

// GetRuleset returns the Ruleset field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) GetRuleset() string {
	if o == nil || IsNil(o.Ruleset) {
		var ret string
		return ret
	}
	return *o.Ruleset
}

// GetRulesetOk returns a tuple with the Ruleset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) GetRulesetOk() (*string, bool) {
	if o == nil || IsNil(o.Ruleset) {
		return nil, false
	}
	return o.Ruleset, true
}

// HasRuleset returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) HasRuleset() bool {
	if o != nil && !IsNil(o.Ruleset) {
		return true
	}

	return false
}

// SetRuleset gets a reference to the given string and assigns it to the Ruleset field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) SetRuleset(v string) {
	o.Ruleset = &v
}

// GetTcpProtection returns the TcpProtection field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) GetTcpProtection() bool {
	if o == nil || IsNil(o.TcpProtection) {
		var ret bool
		return ret
	}
	return *o.TcpProtection
}

// GetTcpProtectionOk returns a tuple with the TcpProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) GetTcpProtectionOk() (*bool, bool) {
	if o == nil || IsNil(o.TcpProtection) {
		return nil, false
	}
	return o.TcpProtection, true
}

// HasTcpProtection returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) HasTcpProtection() bool {
	if o != nil && !IsNil(o.TcpProtection) {
		return true
	}

	return false
}

// SetTcpProtection gets a reference to the given bool and assigns it to the TcpProtection field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) SetTcpProtection(v bool) {
	o.TcpProtection = &v
}

func (o V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Outside) {
		toSerialize["outside"] = o.Outside
	}
	if !IsNil(o.Ruleset) {
		toSerialize["ruleset"] = o.Ruleset
	}
	if !IsNil(o.TcpProtection) {
		toSerialize["tcpProtection"] = o.TcpProtection
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair struct {
	value *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) Get() *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) Set(val *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair(val *V1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) *NullableV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair {
	return &NullableV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesZonesValueZonePairsValuePair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


