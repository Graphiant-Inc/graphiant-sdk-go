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

// checks if the V1EdgesSummaryGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EdgesSummaryGet200Response{}

// V1EdgesSummaryGet200Response struct for V1EdgesSummaryGet200Response
type V1EdgesSummaryGet200Response struct {
	EdgesSummary []V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner `json:"edgesSummary,omitempty"`
	IsNewCount *int32 `json:"isNewCount,omitempty"`
}

// NewV1EdgesSummaryGet200Response instantiates a new V1EdgesSummaryGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EdgesSummaryGet200Response() *V1EdgesSummaryGet200Response {
	this := V1EdgesSummaryGet200Response{}
	return &this
}

// NewV1EdgesSummaryGet200ResponseWithDefaults instantiates a new V1EdgesSummaryGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EdgesSummaryGet200ResponseWithDefaults() *V1EdgesSummaryGet200Response {
	this := V1EdgesSummaryGet200Response{}
	return &this
}

// GetEdgesSummary returns the EdgesSummary field value if set, zero value otherwise.
func (o *V1EdgesSummaryGet200Response) GetEdgesSummary() []V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner {
	if o == nil || IsNil(o.EdgesSummary) {
		var ret []V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner
		return ret
	}
	return o.EdgesSummary
}

// GetEdgesSummaryOk returns a tuple with the EdgesSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EdgesSummaryGet200Response) GetEdgesSummaryOk() ([]V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner, bool) {
	if o == nil || IsNil(o.EdgesSummary) {
		return nil, false
	}
	return o.EdgesSummary, true
}

// HasEdgesSummary returns a boolean if a field has been set.
func (o *V1EdgesSummaryGet200Response) HasEdgesSummary() bool {
	if o != nil && !IsNil(o.EdgesSummary) {
		return true
	}

	return false
}

// SetEdgesSummary gets a reference to the given []V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner and assigns it to the EdgesSummary field.
func (o *V1EdgesSummaryGet200Response) SetEdgesSummary(v []V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) {
	o.EdgesSummary = v
}

// GetIsNewCount returns the IsNewCount field value if set, zero value otherwise.
func (o *V1EdgesSummaryGet200Response) GetIsNewCount() int32 {
	if o == nil || IsNil(o.IsNewCount) {
		var ret int32
		return ret
	}
	return *o.IsNewCount
}

// GetIsNewCountOk returns a tuple with the IsNewCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EdgesSummaryGet200Response) GetIsNewCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IsNewCount) {
		return nil, false
	}
	return o.IsNewCount, true
}

// HasIsNewCount returns a boolean if a field has been set.
func (o *V1EdgesSummaryGet200Response) HasIsNewCount() bool {
	if o != nil && !IsNil(o.IsNewCount) {
		return true
	}

	return false
}

// SetIsNewCount gets a reference to the given int32 and assigns it to the IsNewCount field.
func (o *V1EdgesSummaryGet200Response) SetIsNewCount(v int32) {
	o.IsNewCount = &v
}

func (o V1EdgesSummaryGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EdgesSummaryGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EdgesSummary) {
		toSerialize["edgesSummary"] = o.EdgesSummary
	}
	if !IsNil(o.IsNewCount) {
		toSerialize["isNewCount"] = o.IsNewCount
	}
	return toSerialize, nil
}

type NullableV1EdgesSummaryGet200Response struct {
	value *V1EdgesSummaryGet200Response
	isSet bool
}

func (v NullableV1EdgesSummaryGet200Response) Get() *V1EdgesSummaryGet200Response {
	return v.value
}

func (v *NullableV1EdgesSummaryGet200Response) Set(val *V1EdgesSummaryGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EdgesSummaryGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EdgesSummaryGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EdgesSummaryGet200Response(val *V1EdgesSummaryGet200Response) *NullableV1EdgesSummaryGet200Response {
	return &NullableV1EdgesSummaryGet200Response{value: val, isSet: true}
}

func (v NullableV1EdgesSummaryGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EdgesSummaryGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


