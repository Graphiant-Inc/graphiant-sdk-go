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

// checks if the V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue{}

// V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue struct for V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue
type V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue struct {
	RuleGroup *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValueRuleGroup `json:"ruleGroup,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue instantiates a new V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue() *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue {
	this := V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValueWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValueWithDefaults() *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue {
	this := V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue{}
	return &this
}

// GetRuleGroup returns the RuleGroup field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) GetRuleGroup() V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValueRuleGroup {
	if o == nil || IsNil(o.RuleGroup) {
		var ret V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValueRuleGroup
		return ret
	}
	return *o.RuleGroup
}

// GetRuleGroupOk returns a tuple with the RuleGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) GetRuleGroupOk() (*V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValueRuleGroup, bool) {
	if o == nil || IsNil(o.RuleGroup) {
		return nil, false
	}
	return o.RuleGroup, true
}

// HasRuleGroup returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) HasRuleGroup() bool {
	if o != nil && !IsNil(o.RuleGroup) {
		return true
	}

	return false
}

// SetRuleGroup gets a reference to the given V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValueRuleGroup and assigns it to the RuleGroup field.
func (o *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) SetRuleGroup(v V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValueRuleGroup) {
	o.RuleGroup = &v
}

func (o V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleGroup) {
		toSerialize["ruleGroup"] = o.RuleGroup
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue struct {
	value *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) Get() *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) Set(val *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue(val *V1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) *NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue {
	return &NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusRuleGroupsValueRuleGroupRulesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


