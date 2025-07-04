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

// checks if the V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2{}

// V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 struct for V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2
type V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 struct {
	Circuit *string `json:"circuit,omitempty"`
	Servers *map[string]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue `json:"servers,omitempty"`
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2{}
	return &this
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2WithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2WithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2{}
	return &this
}

// GetCircuit returns the Circuit field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) GetCircuit() string {
	if o == nil || IsNil(o.Circuit) {
		var ret string
		return ret
	}
	return *o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) GetCircuitOk() (*string, bool) {
	if o == nil || IsNil(o.Circuit) {
		return nil, false
	}
	return o.Circuit, true
}

// HasCircuit returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) HasCircuit() bool {
	if o != nil && !IsNil(o.Circuit) {
		return true
	}

	return false
}

// SetCircuit gets a reference to the given string and assigns it to the Circuit field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) SetCircuit(v string) {
	o.Circuit = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) GetServers() map[string]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue {
	if o == nil || IsNil(o.Servers) {
		var ret map[string]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) GetServersOk() (*map[string]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given map[string]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue and assigns it to the Servers field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) SetServers(v map[string]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2ServersValue) {
	o.Servers = &v
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Circuit) {
		toSerialize["circuit"] = o.Circuit
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return toSerialize, nil
}

type NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 struct {
	value *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2
	isSet bool
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) Get() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 {
	return v.value
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) Set(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2 {
	return &NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2{value: val, isSet: true}
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerDnsDynamicServersV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


