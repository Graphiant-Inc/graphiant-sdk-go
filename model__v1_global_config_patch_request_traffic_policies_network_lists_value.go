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

// checks if the V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue{}

// V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue struct for V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue
type V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue struct {
	List *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks `json:"list,omitempty"`
}

// NewV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue() *V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue {
	this := V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue{}
	return &this
}

// NewV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValueWithDefaults instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValueWithDefaults() *V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue {
	this := V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) GetList() V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks {
	if o == nil || IsNil(o.List) {
		var ret V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) GetListOk() (*V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks and assigns it to the List field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) SetList(v V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks) {
	o.List = &v
}

func (o V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue struct {
	value *V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) Get() *V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) Set(val *V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue(val *V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) *NullableV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue {
	return &NullableV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


