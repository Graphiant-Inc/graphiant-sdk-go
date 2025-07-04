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

// checks if the V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner{}

// V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner struct for V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner
type V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner struct {
	DeviceCount *int32 `json:"deviceCount,omitempty"`
	LevelOne *int64 `json:"levelOne,omitempty"`
	LevelOneTag *string `json:"levelOneTag,omitempty"`
	LevelTwo *int64 `json:"levelTwo,omitempty"`
	LevelTwoTag *string `json:"levelTwoTag,omitempty"`
	LevelZero *int64 `json:"levelZero,omitempty"`
	LevelZeroTag *string `json:"levelZeroTag,omitempty"`
}

// NewV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner instantiates a new V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner() *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner {
	this := V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner{}
	return &this
}

// NewV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInnerWithDefaults instantiates a new V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInnerWithDefaults() *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner {
	this := V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner{}
	return &this
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetDeviceCount() int32 {
	if o == nil || IsNil(o.DeviceCount) {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetDeviceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceCount) {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) HasDeviceCount() bool {
	if o != nil && !IsNil(o.DeviceCount) {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

// GetLevelOne returns the LevelOne field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelOne() int64 {
	if o == nil || IsNil(o.LevelOne) {
		var ret int64
		return ret
	}
	return *o.LevelOne
}

// GetLevelOneOk returns a tuple with the LevelOne field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelOneOk() (*int64, bool) {
	if o == nil || IsNil(o.LevelOne) {
		return nil, false
	}
	return o.LevelOne, true
}

// HasLevelOne returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) HasLevelOne() bool {
	if o != nil && !IsNil(o.LevelOne) {
		return true
	}

	return false
}

// SetLevelOne gets a reference to the given int64 and assigns it to the LevelOne field.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) SetLevelOne(v int64) {
	o.LevelOne = &v
}

// GetLevelOneTag returns the LevelOneTag field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelOneTag() string {
	if o == nil || IsNil(o.LevelOneTag) {
		var ret string
		return ret
	}
	return *o.LevelOneTag
}

// GetLevelOneTagOk returns a tuple with the LevelOneTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelOneTagOk() (*string, bool) {
	if o == nil || IsNil(o.LevelOneTag) {
		return nil, false
	}
	return o.LevelOneTag, true
}

// HasLevelOneTag returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) HasLevelOneTag() bool {
	if o != nil && !IsNil(o.LevelOneTag) {
		return true
	}

	return false
}

// SetLevelOneTag gets a reference to the given string and assigns it to the LevelOneTag field.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) SetLevelOneTag(v string) {
	o.LevelOneTag = &v
}

// GetLevelTwo returns the LevelTwo field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelTwo() int64 {
	if o == nil || IsNil(o.LevelTwo) {
		var ret int64
		return ret
	}
	return *o.LevelTwo
}

// GetLevelTwoOk returns a tuple with the LevelTwo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelTwoOk() (*int64, bool) {
	if o == nil || IsNil(o.LevelTwo) {
		return nil, false
	}
	return o.LevelTwo, true
}

// HasLevelTwo returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) HasLevelTwo() bool {
	if o != nil && !IsNil(o.LevelTwo) {
		return true
	}

	return false
}

// SetLevelTwo gets a reference to the given int64 and assigns it to the LevelTwo field.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) SetLevelTwo(v int64) {
	o.LevelTwo = &v
}

// GetLevelTwoTag returns the LevelTwoTag field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelTwoTag() string {
	if o == nil || IsNil(o.LevelTwoTag) {
		var ret string
		return ret
	}
	return *o.LevelTwoTag
}

// GetLevelTwoTagOk returns a tuple with the LevelTwoTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelTwoTagOk() (*string, bool) {
	if o == nil || IsNil(o.LevelTwoTag) {
		return nil, false
	}
	return o.LevelTwoTag, true
}

// HasLevelTwoTag returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) HasLevelTwoTag() bool {
	if o != nil && !IsNil(o.LevelTwoTag) {
		return true
	}

	return false
}

// SetLevelTwoTag gets a reference to the given string and assigns it to the LevelTwoTag field.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) SetLevelTwoTag(v string) {
	o.LevelTwoTag = &v
}

// GetLevelZero returns the LevelZero field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelZero() int64 {
	if o == nil || IsNil(o.LevelZero) {
		var ret int64
		return ret
	}
	return *o.LevelZero
}

// GetLevelZeroOk returns a tuple with the LevelZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelZeroOk() (*int64, bool) {
	if o == nil || IsNil(o.LevelZero) {
		return nil, false
	}
	return o.LevelZero, true
}

// HasLevelZero returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) HasLevelZero() bool {
	if o != nil && !IsNil(o.LevelZero) {
		return true
	}

	return false
}

// SetLevelZero gets a reference to the given int64 and assigns it to the LevelZero field.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) SetLevelZero(v int64) {
	o.LevelZero = &v
}

// GetLevelZeroTag returns the LevelZeroTag field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelZeroTag() string {
	if o == nil || IsNil(o.LevelZeroTag) {
		var ret string
		return ret
	}
	return *o.LevelZeroTag
}

// GetLevelZeroTagOk returns a tuple with the LevelZeroTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) GetLevelZeroTagOk() (*string, bool) {
	if o == nil || IsNil(o.LevelZeroTag) {
		return nil, false
	}
	return o.LevelZeroTag, true
}

// HasLevelZeroTag returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) HasLevelZeroTag() bool {
	if o != nil && !IsNil(o.LevelZeroTag) {
		return true
	}

	return false
}

// SetLevelZeroTag gets a reference to the given string and assigns it to the LevelZeroTag field.
func (o *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) SetLevelZeroTag(v string) {
	o.LevelZeroTag = &v
}

func (o V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceCount) {
		toSerialize["deviceCount"] = o.DeviceCount
	}
	if !IsNil(o.LevelOne) {
		toSerialize["levelOne"] = o.LevelOne
	}
	if !IsNil(o.LevelOneTag) {
		toSerialize["levelOneTag"] = o.LevelOneTag
	}
	if !IsNil(o.LevelTwo) {
		toSerialize["levelTwo"] = o.LevelTwo
	}
	if !IsNil(o.LevelTwoTag) {
		toSerialize["levelTwoTag"] = o.LevelTwoTag
	}
	if !IsNil(o.LevelZero) {
		toSerialize["levelZero"] = o.LevelZero
	}
	if !IsNil(o.LevelZeroTag) {
		toSerialize["levelZeroTag"] = o.LevelZeroTag
	}
	return toSerialize, nil
}

type NullableV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner struct {
	value *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner
	isSet bool
}

func (v NullableV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) Get() *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner {
	return v.value
}

func (v *NullableV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) Set(val *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner(val *V1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) *NullableV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner {
	return &NullableV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner{value: val, isSet: true}
}

func (v NullableV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PolicyRouteTagSetsTagsSummaryGet200ResponseTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


