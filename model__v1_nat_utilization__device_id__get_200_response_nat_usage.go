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

// checks if the V1NatUtilizationDeviceIdGet200ResponseNatUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1NatUtilizationDeviceIdGet200ResponseNatUsage{}

// V1NatUtilizationDeviceIdGet200ResponseNatUsage struct for V1NatUtilizationDeviceIdGet200ResponseNatUsage
type V1NatUtilizationDeviceIdGet200ResponseNatUsage struct {
	CurrentCount *int64 `json:"currentCount,omitempty"`
	CurrentCountExtranet *int64 `json:"currentCountExtranet,omitempty"`
	CurrentCountPat *int64 `json:"currentCountPat,omitempty"`
	CurrentCountStatic *int64 `json:"currentCountStatic,omitempty"`
	MaxCount *int64 `json:"maxCount,omitempty"`
}

// NewV1NatUtilizationDeviceIdGet200ResponseNatUsage instantiates a new V1NatUtilizationDeviceIdGet200ResponseNatUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NatUtilizationDeviceIdGet200ResponseNatUsage() *V1NatUtilizationDeviceIdGet200ResponseNatUsage {
	this := V1NatUtilizationDeviceIdGet200ResponseNatUsage{}
	return &this
}

// NewV1NatUtilizationDeviceIdGet200ResponseNatUsageWithDefaults instantiates a new V1NatUtilizationDeviceIdGet200ResponseNatUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NatUtilizationDeviceIdGet200ResponseNatUsageWithDefaults() *V1NatUtilizationDeviceIdGet200ResponseNatUsage {
	this := V1NatUtilizationDeviceIdGet200ResponseNatUsage{}
	return &this
}

// GetCurrentCount returns the CurrentCount field value if set, zero value otherwise.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) GetCurrentCount() int64 {
	if o == nil || IsNil(o.CurrentCount) {
		var ret int64
		return ret
	}
	return *o.CurrentCount
}

// GetCurrentCountOk returns a tuple with the CurrentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) GetCurrentCountOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentCount) {
		return nil, false
	}
	return o.CurrentCount, true
}

// HasCurrentCount returns a boolean if a field has been set.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) HasCurrentCount() bool {
	if o != nil && !IsNil(o.CurrentCount) {
		return true
	}

	return false
}

// SetCurrentCount gets a reference to the given int64 and assigns it to the CurrentCount field.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) SetCurrentCount(v int64) {
	o.CurrentCount = &v
}

// GetCurrentCountExtranet returns the CurrentCountExtranet field value if set, zero value otherwise.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) GetCurrentCountExtranet() int64 {
	if o == nil || IsNil(o.CurrentCountExtranet) {
		var ret int64
		return ret
	}
	return *o.CurrentCountExtranet
}

// GetCurrentCountExtranetOk returns a tuple with the CurrentCountExtranet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) GetCurrentCountExtranetOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentCountExtranet) {
		return nil, false
	}
	return o.CurrentCountExtranet, true
}

// HasCurrentCountExtranet returns a boolean if a field has been set.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) HasCurrentCountExtranet() bool {
	if o != nil && !IsNil(o.CurrentCountExtranet) {
		return true
	}

	return false
}

// SetCurrentCountExtranet gets a reference to the given int64 and assigns it to the CurrentCountExtranet field.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) SetCurrentCountExtranet(v int64) {
	o.CurrentCountExtranet = &v
}

// GetCurrentCountPat returns the CurrentCountPat field value if set, zero value otherwise.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) GetCurrentCountPat() int64 {
	if o == nil || IsNil(o.CurrentCountPat) {
		var ret int64
		return ret
	}
	return *o.CurrentCountPat
}

// GetCurrentCountPatOk returns a tuple with the CurrentCountPat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) GetCurrentCountPatOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentCountPat) {
		return nil, false
	}
	return o.CurrentCountPat, true
}

// HasCurrentCountPat returns a boolean if a field has been set.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) HasCurrentCountPat() bool {
	if o != nil && !IsNil(o.CurrentCountPat) {
		return true
	}

	return false
}

// SetCurrentCountPat gets a reference to the given int64 and assigns it to the CurrentCountPat field.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) SetCurrentCountPat(v int64) {
	o.CurrentCountPat = &v
}

// GetCurrentCountStatic returns the CurrentCountStatic field value if set, zero value otherwise.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) GetCurrentCountStatic() int64 {
	if o == nil || IsNil(o.CurrentCountStatic) {
		var ret int64
		return ret
	}
	return *o.CurrentCountStatic
}

// GetCurrentCountStaticOk returns a tuple with the CurrentCountStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) GetCurrentCountStaticOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentCountStatic) {
		return nil, false
	}
	return o.CurrentCountStatic, true
}

// HasCurrentCountStatic returns a boolean if a field has been set.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) HasCurrentCountStatic() bool {
	if o != nil && !IsNil(o.CurrentCountStatic) {
		return true
	}

	return false
}

// SetCurrentCountStatic gets a reference to the given int64 and assigns it to the CurrentCountStatic field.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) SetCurrentCountStatic(v int64) {
	o.CurrentCountStatic = &v
}

// GetMaxCount returns the MaxCount field value if set, zero value otherwise.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) GetMaxCount() int64 {
	if o == nil || IsNil(o.MaxCount) {
		var ret int64
		return ret
	}
	return *o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) GetMaxCountOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxCount) {
		return nil, false
	}
	return o.MaxCount, true
}

// HasMaxCount returns a boolean if a field has been set.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) HasMaxCount() bool {
	if o != nil && !IsNil(o.MaxCount) {
		return true
	}

	return false
}

// SetMaxCount gets a reference to the given int64 and assigns it to the MaxCount field.
func (o *V1NatUtilizationDeviceIdGet200ResponseNatUsage) SetMaxCount(v int64) {
	o.MaxCount = &v
}

func (o V1NatUtilizationDeviceIdGet200ResponseNatUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1NatUtilizationDeviceIdGet200ResponseNatUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentCount) {
		toSerialize["currentCount"] = o.CurrentCount
	}
	if !IsNil(o.CurrentCountExtranet) {
		toSerialize["currentCountExtranet"] = o.CurrentCountExtranet
	}
	if !IsNil(o.CurrentCountPat) {
		toSerialize["currentCountPat"] = o.CurrentCountPat
	}
	if !IsNil(o.CurrentCountStatic) {
		toSerialize["currentCountStatic"] = o.CurrentCountStatic
	}
	if !IsNil(o.MaxCount) {
		toSerialize["maxCount"] = o.MaxCount
	}
	return toSerialize, nil
}

type NullableV1NatUtilizationDeviceIdGet200ResponseNatUsage struct {
	value *V1NatUtilizationDeviceIdGet200ResponseNatUsage
	isSet bool
}

func (v NullableV1NatUtilizationDeviceIdGet200ResponseNatUsage) Get() *V1NatUtilizationDeviceIdGet200ResponseNatUsage {
	return v.value
}

func (v *NullableV1NatUtilizationDeviceIdGet200ResponseNatUsage) Set(val *V1NatUtilizationDeviceIdGet200ResponseNatUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NatUtilizationDeviceIdGet200ResponseNatUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NatUtilizationDeviceIdGet200ResponseNatUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NatUtilizationDeviceIdGet200ResponseNatUsage(val *V1NatUtilizationDeviceIdGet200ResponseNatUsage) *NullableV1NatUtilizationDeviceIdGet200ResponseNatUsage {
	return &NullableV1NatUtilizationDeviceIdGet200ResponseNatUsage{value: val, isSet: true}
}

func (v NullableV1NatUtilizationDeviceIdGet200ResponseNatUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NatUtilizationDeviceIdGet200ResponseNatUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


