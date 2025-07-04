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

// checks if the V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue{}

// V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue struct for V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue
type V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue struct {
	DynamicServers []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsCloudflareServersInner `json:"dynamicServers,omitempty"`
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue{}
	return &this
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValueWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValueWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue{}
	return &this
}

// GetDynamicServers returns the DynamicServers field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) GetDynamicServers() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsCloudflareServersInner {
	if o == nil || IsNil(o.DynamicServers) {
		var ret []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsCloudflareServersInner
		return ret
	}
	return o.DynamicServers
}

// GetDynamicServersOk returns a tuple with the DynamicServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) GetDynamicServersOk() ([]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsCloudflareServersInner, bool) {
	if o == nil || IsNil(o.DynamicServers) {
		return nil, false
	}
	return o.DynamicServers, true
}

// HasDynamicServers returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) HasDynamicServers() bool {
	if o != nil && !IsNil(o.DynamicServers) {
		return true
	}

	return false
}

// SetDynamicServers gets a reference to the given []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsCloudflareServersInner and assigns it to the DynamicServers field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) SetDynamicServers(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsCloudflareServersInner) {
	o.DynamicServers = v
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DynamicServers) {
		toSerialize["dynamicServers"] = o.DynamicServers
	}
	return toSerialize, nil
}

type NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue struct {
	value *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue
	isSet bool
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) Get() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue {
	return v.value
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) Set(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue {
	return &NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue{value: val, isSet: true}
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


