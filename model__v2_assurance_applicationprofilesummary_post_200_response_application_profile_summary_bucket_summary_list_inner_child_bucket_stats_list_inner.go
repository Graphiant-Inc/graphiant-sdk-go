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

// checks if the V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner{}

// V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner struct for V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner
type V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner struct {
	AssuranceBucket *string `json:"assuranceBucket,omitempty"`
	BucketNameToDisplay *string `json:"bucketNameToDisplay,omitempty"`
	PrevUniqueAppCount *int64 `json:"prevUniqueAppCount,omitempty"`
	UniqueAppCount *int64 `json:"uniqueAppCount,omitempty"`
}

// NewV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner instantiates a new V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner() *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner {
	this := V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner{}
	return &this
}

// NewV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInnerWithDefaults instantiates a new V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInnerWithDefaults() *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner {
	this := V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner{}
	return &this
}

// GetAssuranceBucket returns the AssuranceBucket field value if set, zero value otherwise.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) GetAssuranceBucket() string {
	if o == nil || IsNil(o.AssuranceBucket) {
		var ret string
		return ret
	}
	return *o.AssuranceBucket
}

// GetAssuranceBucketOk returns a tuple with the AssuranceBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) GetAssuranceBucketOk() (*string, bool) {
	if o == nil || IsNil(o.AssuranceBucket) {
		return nil, false
	}
	return o.AssuranceBucket, true
}

// HasAssuranceBucket returns a boolean if a field has been set.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) HasAssuranceBucket() bool {
	if o != nil && !IsNil(o.AssuranceBucket) {
		return true
	}

	return false
}

// SetAssuranceBucket gets a reference to the given string and assigns it to the AssuranceBucket field.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) SetAssuranceBucket(v string) {
	o.AssuranceBucket = &v
}

// GetBucketNameToDisplay returns the BucketNameToDisplay field value if set, zero value otherwise.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) GetBucketNameToDisplay() string {
	if o == nil || IsNil(o.BucketNameToDisplay) {
		var ret string
		return ret
	}
	return *o.BucketNameToDisplay
}

// GetBucketNameToDisplayOk returns a tuple with the BucketNameToDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) GetBucketNameToDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.BucketNameToDisplay) {
		return nil, false
	}
	return o.BucketNameToDisplay, true
}

// HasBucketNameToDisplay returns a boolean if a field has been set.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) HasBucketNameToDisplay() bool {
	if o != nil && !IsNil(o.BucketNameToDisplay) {
		return true
	}

	return false
}

// SetBucketNameToDisplay gets a reference to the given string and assigns it to the BucketNameToDisplay field.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) SetBucketNameToDisplay(v string) {
	o.BucketNameToDisplay = &v
}

// GetPrevUniqueAppCount returns the PrevUniqueAppCount field value if set, zero value otherwise.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) GetPrevUniqueAppCount() int64 {
	if o == nil || IsNil(o.PrevUniqueAppCount) {
		var ret int64
		return ret
	}
	return *o.PrevUniqueAppCount
}

// GetPrevUniqueAppCountOk returns a tuple with the PrevUniqueAppCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) GetPrevUniqueAppCountOk() (*int64, bool) {
	if o == nil || IsNil(o.PrevUniqueAppCount) {
		return nil, false
	}
	return o.PrevUniqueAppCount, true
}

// HasPrevUniqueAppCount returns a boolean if a field has been set.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) HasPrevUniqueAppCount() bool {
	if o != nil && !IsNil(o.PrevUniqueAppCount) {
		return true
	}

	return false
}

// SetPrevUniqueAppCount gets a reference to the given int64 and assigns it to the PrevUniqueAppCount field.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) SetPrevUniqueAppCount(v int64) {
	o.PrevUniqueAppCount = &v
}

// GetUniqueAppCount returns the UniqueAppCount field value if set, zero value otherwise.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) GetUniqueAppCount() int64 {
	if o == nil || IsNil(o.UniqueAppCount) {
		var ret int64
		return ret
	}
	return *o.UniqueAppCount
}

// GetUniqueAppCountOk returns a tuple with the UniqueAppCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) GetUniqueAppCountOk() (*int64, bool) {
	if o == nil || IsNil(o.UniqueAppCount) {
		return nil, false
	}
	return o.UniqueAppCount, true
}

// HasUniqueAppCount returns a boolean if a field has been set.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) HasUniqueAppCount() bool {
	if o != nil && !IsNil(o.UniqueAppCount) {
		return true
	}

	return false
}

// SetUniqueAppCount gets a reference to the given int64 and assigns it to the UniqueAppCount field.
func (o *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) SetUniqueAppCount(v int64) {
	o.UniqueAppCount = &v
}

func (o V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssuranceBucket) {
		toSerialize["assuranceBucket"] = o.AssuranceBucket
	}
	if !IsNil(o.BucketNameToDisplay) {
		toSerialize["bucketNameToDisplay"] = o.BucketNameToDisplay
	}
	if !IsNil(o.PrevUniqueAppCount) {
		toSerialize["prevUniqueAppCount"] = o.PrevUniqueAppCount
	}
	if !IsNil(o.UniqueAppCount) {
		toSerialize["uniqueAppCount"] = o.UniqueAppCount
	}
	return toSerialize, nil
}

type NullableV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner struct {
	value *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner
	isSet bool
}

func (v NullableV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) Get() *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner {
	return v.value
}

func (v *NullableV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) Set(val *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner(val *V2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) *NullableV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner {
	return &NullableV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner{value: val, isSet: true}
}

func (v NullableV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssuranceApplicationprofilesummaryPost200ResponseApplicationProfileSummaryBucketSummaryListInnerChildBucketStatsListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


