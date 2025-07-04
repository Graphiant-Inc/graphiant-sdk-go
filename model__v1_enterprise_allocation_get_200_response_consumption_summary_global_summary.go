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

// checks if the V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary{}

// V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary struct for V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary
type V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary struct {
	AllocatedMonthlyCredits *float32 `json:"allocatedMonthlyCredits,omitempty"`
	ConsumedMonthlyCredits *float32 `json:"consumedMonthlyCredits,omitempty"`
	RecommendedMonthlyCredits *float32 `json:"recommendedMonthlyCredits,omitempty"`
}

// NewV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary instantiates a new V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary() *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary {
	this := V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary{}
	return &this
}

// NewV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummaryWithDefaults instantiates a new V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummaryWithDefaults() *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary {
	this := V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary{}
	return &this
}

// GetAllocatedMonthlyCredits returns the AllocatedMonthlyCredits field value if set, zero value otherwise.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) GetAllocatedMonthlyCredits() float32 {
	if o == nil || IsNil(o.AllocatedMonthlyCredits) {
		var ret float32
		return ret
	}
	return *o.AllocatedMonthlyCredits
}

// GetAllocatedMonthlyCreditsOk returns a tuple with the AllocatedMonthlyCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) GetAllocatedMonthlyCreditsOk() (*float32, bool) {
	if o == nil || IsNil(o.AllocatedMonthlyCredits) {
		return nil, false
	}
	return o.AllocatedMonthlyCredits, true
}

// HasAllocatedMonthlyCredits returns a boolean if a field has been set.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) HasAllocatedMonthlyCredits() bool {
	if o != nil && !IsNil(o.AllocatedMonthlyCredits) {
		return true
	}

	return false
}

// SetAllocatedMonthlyCredits gets a reference to the given float32 and assigns it to the AllocatedMonthlyCredits field.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) SetAllocatedMonthlyCredits(v float32) {
	o.AllocatedMonthlyCredits = &v
}

// GetConsumedMonthlyCredits returns the ConsumedMonthlyCredits field value if set, zero value otherwise.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) GetConsumedMonthlyCredits() float32 {
	if o == nil || IsNil(o.ConsumedMonthlyCredits) {
		var ret float32
		return ret
	}
	return *o.ConsumedMonthlyCredits
}

// GetConsumedMonthlyCreditsOk returns a tuple with the ConsumedMonthlyCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) GetConsumedMonthlyCreditsOk() (*float32, bool) {
	if o == nil || IsNil(o.ConsumedMonthlyCredits) {
		return nil, false
	}
	return o.ConsumedMonthlyCredits, true
}

// HasConsumedMonthlyCredits returns a boolean if a field has been set.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) HasConsumedMonthlyCredits() bool {
	if o != nil && !IsNil(o.ConsumedMonthlyCredits) {
		return true
	}

	return false
}

// SetConsumedMonthlyCredits gets a reference to the given float32 and assigns it to the ConsumedMonthlyCredits field.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) SetConsumedMonthlyCredits(v float32) {
	o.ConsumedMonthlyCredits = &v
}

// GetRecommendedMonthlyCredits returns the RecommendedMonthlyCredits field value if set, zero value otherwise.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) GetRecommendedMonthlyCredits() float32 {
	if o == nil || IsNil(o.RecommendedMonthlyCredits) {
		var ret float32
		return ret
	}
	return *o.RecommendedMonthlyCredits
}

// GetRecommendedMonthlyCreditsOk returns a tuple with the RecommendedMonthlyCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) GetRecommendedMonthlyCreditsOk() (*float32, bool) {
	if o == nil || IsNil(o.RecommendedMonthlyCredits) {
		return nil, false
	}
	return o.RecommendedMonthlyCredits, true
}

// HasRecommendedMonthlyCredits returns a boolean if a field has been set.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) HasRecommendedMonthlyCredits() bool {
	if o != nil && !IsNil(o.RecommendedMonthlyCredits) {
		return true
	}

	return false
}

// SetRecommendedMonthlyCredits gets a reference to the given float32 and assigns it to the RecommendedMonthlyCredits field.
func (o *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) SetRecommendedMonthlyCredits(v float32) {
	o.RecommendedMonthlyCredits = &v
}

func (o V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllocatedMonthlyCredits) {
		toSerialize["allocatedMonthlyCredits"] = o.AllocatedMonthlyCredits
	}
	if !IsNil(o.ConsumedMonthlyCredits) {
		toSerialize["consumedMonthlyCredits"] = o.ConsumedMonthlyCredits
	}
	if !IsNil(o.RecommendedMonthlyCredits) {
		toSerialize["recommendedMonthlyCredits"] = o.RecommendedMonthlyCredits
	}
	return toSerialize, nil
}

type NullableV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary struct {
	value *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary
	isSet bool
}

func (v NullableV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) Get() *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary {
	return v.value
}

func (v *NullableV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) Set(val *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary(val *V1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) *NullableV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary {
	return &NullableV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary{value: val, isSet: true}
}

func (v NullableV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EnterpriseAllocationGet200ResponseConsumptionSummaryGlobalSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


