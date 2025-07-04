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

// checks if the V2AssuranceTopologyFlowsPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2AssuranceTopologyFlowsPost200Response{}

// V2AssuranceTopologyFlowsPost200Response struct for V2AssuranceTopologyFlowsPost200Response
type V2AssuranceTopologyFlowsPost200Response struct {
	Flows []V2AssuranceTopologyFlowsPost200ResponseFlowsInner `json:"flows,omitempty"`
}

// NewV2AssuranceTopologyFlowsPost200Response instantiates a new V2AssuranceTopologyFlowsPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssuranceTopologyFlowsPost200Response() *V2AssuranceTopologyFlowsPost200Response {
	this := V2AssuranceTopologyFlowsPost200Response{}
	return &this
}

// NewV2AssuranceTopologyFlowsPost200ResponseWithDefaults instantiates a new V2AssuranceTopologyFlowsPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssuranceTopologyFlowsPost200ResponseWithDefaults() *V2AssuranceTopologyFlowsPost200Response {
	this := V2AssuranceTopologyFlowsPost200Response{}
	return &this
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *V2AssuranceTopologyFlowsPost200Response) GetFlows() []V2AssuranceTopologyFlowsPost200ResponseFlowsInner {
	if o == nil || IsNil(o.Flows) {
		var ret []V2AssuranceTopologyFlowsPost200ResponseFlowsInner
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyFlowsPost200Response) GetFlowsOk() ([]V2AssuranceTopologyFlowsPost200ResponseFlowsInner, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *V2AssuranceTopologyFlowsPost200Response) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []V2AssuranceTopologyFlowsPost200ResponseFlowsInner and assigns it to the Flows field.
func (o *V2AssuranceTopologyFlowsPost200Response) SetFlows(v []V2AssuranceTopologyFlowsPost200ResponseFlowsInner) {
	o.Flows = v
}

func (o V2AssuranceTopologyFlowsPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2AssuranceTopologyFlowsPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	return toSerialize, nil
}

type NullableV2AssuranceTopologyFlowsPost200Response struct {
	value *V2AssuranceTopologyFlowsPost200Response
	isSet bool
}

func (v NullableV2AssuranceTopologyFlowsPost200Response) Get() *V2AssuranceTopologyFlowsPost200Response {
	return v.value
}

func (v *NullableV2AssuranceTopologyFlowsPost200Response) Set(val *V2AssuranceTopologyFlowsPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssuranceTopologyFlowsPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssuranceTopologyFlowsPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssuranceTopologyFlowsPost200Response(val *V2AssuranceTopologyFlowsPost200Response) *NullableV2AssuranceTopologyFlowsPost200Response {
	return &NullableV2AssuranceTopologyFlowsPost200Response{value: val, isSet: true}
}

func (v NullableV2AssuranceTopologyFlowsPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssuranceTopologyFlowsPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


