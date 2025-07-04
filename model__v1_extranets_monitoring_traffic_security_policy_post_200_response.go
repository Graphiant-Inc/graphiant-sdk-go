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

// checks if the V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response{}

// V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response struct for V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response
type V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response struct {
	SecurityRules []V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseSecurityRulesInner `json:"securityRules,omitempty"`
	TrafficRules []V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseTrafficRulesInner `json:"trafficRules,omitempty"`
}

// NewV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response instantiates a new V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response() *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response {
	this := V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response{}
	return &this
}

// NewV1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseWithDefaults instantiates a new V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseWithDefaults() *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response {
	this := V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response{}
	return &this
}

// GetSecurityRules returns the SecurityRules field value if set, zero value otherwise.
func (o *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) GetSecurityRules() []V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseSecurityRulesInner {
	if o == nil || IsNil(o.SecurityRules) {
		var ret []V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseSecurityRulesInner
		return ret
	}
	return o.SecurityRules
}

// GetSecurityRulesOk returns a tuple with the SecurityRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) GetSecurityRulesOk() ([]V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseSecurityRulesInner, bool) {
	if o == nil || IsNil(o.SecurityRules) {
		return nil, false
	}
	return o.SecurityRules, true
}

// HasSecurityRules returns a boolean if a field has been set.
func (o *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) HasSecurityRules() bool {
	if o != nil && !IsNil(o.SecurityRules) {
		return true
	}

	return false
}

// SetSecurityRules gets a reference to the given []V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseSecurityRulesInner and assigns it to the SecurityRules field.
func (o *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) SetSecurityRules(v []V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseSecurityRulesInner) {
	o.SecurityRules = v
}

// GetTrafficRules returns the TrafficRules field value if set, zero value otherwise.
func (o *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) GetTrafficRules() []V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseTrafficRulesInner {
	if o == nil || IsNil(o.TrafficRules) {
		var ret []V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseTrafficRulesInner
		return ret
	}
	return o.TrafficRules
}

// GetTrafficRulesOk returns a tuple with the TrafficRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) GetTrafficRulesOk() ([]V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseTrafficRulesInner, bool) {
	if o == nil || IsNil(o.TrafficRules) {
		return nil, false
	}
	return o.TrafficRules, true
}

// HasTrafficRules returns a boolean if a field has been set.
func (o *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) HasTrafficRules() bool {
	if o != nil && !IsNil(o.TrafficRules) {
		return true
	}

	return false
}

// SetTrafficRules gets a reference to the given []V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseTrafficRulesInner and assigns it to the TrafficRules field.
func (o *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) SetTrafficRules(v []V1ExtranetsMonitoringTrafficSecurityPolicyPost200ResponseTrafficRulesInner) {
	o.TrafficRules = v
}

func (o V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecurityRules) {
		toSerialize["securityRules"] = o.SecurityRules
	}
	if !IsNil(o.TrafficRules) {
		toSerialize["trafficRules"] = o.TrafficRules
	}
	return toSerialize, nil
}

type NullableV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response struct {
	value *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response
	isSet bool
}

func (v NullableV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) Get() *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response {
	return v.value
}

func (v *NullableV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) Set(val *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response(val *V1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) *NullableV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response {
	return &NullableV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response{value: val, isSet: true}
}

func (v NullableV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ExtranetsMonitoringTrafficSecurityPolicyPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


