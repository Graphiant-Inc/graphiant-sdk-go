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

// checks if the V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner{}

// V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner struct for V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner
type V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner struct {
	Jitter *float32 `json:"jitter,omitempty"`
	Latency *float32 `json:"latency,omitempty"`
	Loss *float32 `json:"loss,omitempty"`
}

// NewV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner instantiates a new V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner() *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner {
	this := V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner{}
	return &this
}

// NewV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInnerWithDefaults instantiates a new V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInnerWithDefaults() *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner {
	this := V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner{}
	return &this
}

// GetJitter returns the Jitter field value if set, zero value otherwise.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) GetJitter() float32 {
	if o == nil || IsNil(o.Jitter) {
		var ret float32
		return ret
	}
	return *o.Jitter
}

// GetJitterOk returns a tuple with the Jitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) GetJitterOk() (*float32, bool) {
	if o == nil || IsNil(o.Jitter) {
		return nil, false
	}
	return o.Jitter, true
}

// HasJitter returns a boolean if a field has been set.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) HasJitter() bool {
	if o != nil && !IsNil(o.Jitter) {
		return true
	}

	return false
}

// SetJitter gets a reference to the given float32 and assigns it to the Jitter field.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) SetJitter(v float32) {
	o.Jitter = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) GetLatency() float32 {
	if o == nil || IsNil(o.Latency) {
		var ret float32
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) GetLatencyOk() (*float32, bool) {
	if o == nil || IsNil(o.Latency) {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) HasLatency() bool {
	if o != nil && !IsNil(o.Latency) {
		return true
	}

	return false
}

// SetLatency gets a reference to the given float32 and assigns it to the Latency field.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) SetLatency(v float32) {
	o.Latency = &v
}

// GetLoss returns the Loss field value if set, zero value otherwise.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) GetLoss() float32 {
	if o == nil || IsNil(o.Loss) {
		var ret float32
		return ret
	}
	return *o.Loss
}

// GetLossOk returns a tuple with the Loss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) GetLossOk() (*float32, bool) {
	if o == nil || IsNil(o.Loss) {
		return nil, false
	}
	return o.Loss, true
}

// HasLoss returns a boolean if a field has been set.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) HasLoss() bool {
	if o != nil && !IsNil(o.Loss) {
		return true
	}

	return false
}

// SetLoss gets a reference to the given float32 and assigns it to the Loss field.
func (o *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) SetLoss(v float32) {
	o.Loss = &v
}

func (o V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Jitter) {
		toSerialize["jitter"] = o.Jitter
	}
	if !IsNil(o.Latency) {
		toSerialize["latency"] = o.Latency
	}
	if !IsNil(o.Loss) {
		toSerialize["loss"] = o.Loss
	}
	return toSerialize, nil
}

type NullableV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner struct {
	value *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner
	isSet bool
}

func (v NullableV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) Get() *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner {
	return v.value
}

func (v *NullableV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) Set(val *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner(val *V2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) *NullableV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner {
	return &NullableV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner{value: val, isSet: true}
}

func (v NullableV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssuranceTopologyOverviewPost200ResponseTopologyEdgesInnerPerformanceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


