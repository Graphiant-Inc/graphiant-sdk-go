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

// checks if the V1FlowsFlowTablePostRequestSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1FlowsFlowTablePostRequestSelector{}

// V1FlowsFlowTablePostRequestSelector struct for V1FlowsFlowTablePostRequestSelector
type V1FlowsFlowTablePostRequestSelector struct {
	CircuitName []string `json:"circuitName,omitempty"`
	SlaClass []string `json:"slaClass,omitempty"`
}

// NewV1FlowsFlowTablePostRequestSelector instantiates a new V1FlowsFlowTablePostRequestSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1FlowsFlowTablePostRequestSelector() *V1FlowsFlowTablePostRequestSelector {
	this := V1FlowsFlowTablePostRequestSelector{}
	return &this
}

// NewV1FlowsFlowTablePostRequestSelectorWithDefaults instantiates a new V1FlowsFlowTablePostRequestSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1FlowsFlowTablePostRequestSelectorWithDefaults() *V1FlowsFlowTablePostRequestSelector {
	this := V1FlowsFlowTablePostRequestSelector{}
	return &this
}

// GetCircuitName returns the CircuitName field value if set, zero value otherwise.
func (o *V1FlowsFlowTablePostRequestSelector) GetCircuitName() []string {
	if o == nil || IsNil(o.CircuitName) {
		var ret []string
		return ret
	}
	return o.CircuitName
}

// GetCircuitNameOk returns a tuple with the CircuitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1FlowsFlowTablePostRequestSelector) GetCircuitNameOk() ([]string, bool) {
	if o == nil || IsNil(o.CircuitName) {
		return nil, false
	}
	return o.CircuitName, true
}

// HasCircuitName returns a boolean if a field has been set.
func (o *V1FlowsFlowTablePostRequestSelector) HasCircuitName() bool {
	if o != nil && !IsNil(o.CircuitName) {
		return true
	}

	return false
}

// SetCircuitName gets a reference to the given []string and assigns it to the CircuitName field.
func (o *V1FlowsFlowTablePostRequestSelector) SetCircuitName(v []string) {
	o.CircuitName = v
}

// GetSlaClass returns the SlaClass field value if set, zero value otherwise.
func (o *V1FlowsFlowTablePostRequestSelector) GetSlaClass() []string {
	if o == nil || IsNil(o.SlaClass) {
		var ret []string
		return ret
	}
	return o.SlaClass
}

// GetSlaClassOk returns a tuple with the SlaClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1FlowsFlowTablePostRequestSelector) GetSlaClassOk() ([]string, bool) {
	if o == nil || IsNil(o.SlaClass) {
		return nil, false
	}
	return o.SlaClass, true
}

// HasSlaClass returns a boolean if a field has been set.
func (o *V1FlowsFlowTablePostRequestSelector) HasSlaClass() bool {
	if o != nil && !IsNil(o.SlaClass) {
		return true
	}

	return false
}

// SetSlaClass gets a reference to the given []string and assigns it to the SlaClass field.
func (o *V1FlowsFlowTablePostRequestSelector) SetSlaClass(v []string) {
	o.SlaClass = v
}

func (o V1FlowsFlowTablePostRequestSelector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1FlowsFlowTablePostRequestSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CircuitName) {
		toSerialize["circuitName"] = o.CircuitName
	}
	if !IsNil(o.SlaClass) {
		toSerialize["slaClass"] = o.SlaClass
	}
	return toSerialize, nil
}

type NullableV1FlowsFlowTablePostRequestSelector struct {
	value *V1FlowsFlowTablePostRequestSelector
	isSet bool
}

func (v NullableV1FlowsFlowTablePostRequestSelector) Get() *V1FlowsFlowTablePostRequestSelector {
	return v.value
}

func (v *NullableV1FlowsFlowTablePostRequestSelector) Set(val *V1FlowsFlowTablePostRequestSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableV1FlowsFlowTablePostRequestSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableV1FlowsFlowTablePostRequestSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1FlowsFlowTablePostRequestSelector(val *V1FlowsFlowTablePostRequestSelector) *NullableV1FlowsFlowTablePostRequestSelector {
	return &NullableV1FlowsFlowTablePostRequestSelector{value: val, isSet: true}
}

func (v NullableV1FlowsFlowTablePostRequestSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1FlowsFlowTablePostRequestSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


