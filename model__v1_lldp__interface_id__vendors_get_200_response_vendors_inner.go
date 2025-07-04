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

// checks if the V1LldpInterfaceIdVendorsGet200ResponseVendorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1LldpInterfaceIdVendorsGet200ResponseVendorsInner{}

// V1LldpInterfaceIdVendorsGet200ResponseVendorsInner struct for V1LldpInterfaceIdVendorsGet200ResponseVendorsInner
type V1LldpInterfaceIdVendorsGet200ResponseVendorsInner struct {
	Name *string `json:"name,omitempty"`
	NeighborCount *int64 `json:"neighborCount,omitempty"`
}

// NewV1LldpInterfaceIdVendorsGet200ResponseVendorsInner instantiates a new V1LldpInterfaceIdVendorsGet200ResponseVendorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1LldpInterfaceIdVendorsGet200ResponseVendorsInner() *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner {
	this := V1LldpInterfaceIdVendorsGet200ResponseVendorsInner{}
	return &this
}

// NewV1LldpInterfaceIdVendorsGet200ResponseVendorsInnerWithDefaults instantiates a new V1LldpInterfaceIdVendorsGet200ResponseVendorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1LldpInterfaceIdVendorsGet200ResponseVendorsInnerWithDefaults() *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner {
	this := V1LldpInterfaceIdVendorsGet200ResponseVendorsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) SetName(v string) {
	o.Name = &v
}

// GetNeighborCount returns the NeighborCount field value if set, zero value otherwise.
func (o *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) GetNeighborCount() int64 {
	if o == nil || IsNil(o.NeighborCount) {
		var ret int64
		return ret
	}
	return *o.NeighborCount
}

// GetNeighborCountOk returns a tuple with the NeighborCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) GetNeighborCountOk() (*int64, bool) {
	if o == nil || IsNil(o.NeighborCount) {
		return nil, false
	}
	return o.NeighborCount, true
}

// HasNeighborCount returns a boolean if a field has been set.
func (o *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) HasNeighborCount() bool {
	if o != nil && !IsNil(o.NeighborCount) {
		return true
	}

	return false
}

// SetNeighborCount gets a reference to the given int64 and assigns it to the NeighborCount field.
func (o *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) SetNeighborCount(v int64) {
	o.NeighborCount = &v
}

func (o V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NeighborCount) {
		toSerialize["neighborCount"] = o.NeighborCount
	}
	return toSerialize, nil
}

type NullableV1LldpInterfaceIdVendorsGet200ResponseVendorsInner struct {
	value *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner
	isSet bool
}

func (v NullableV1LldpInterfaceIdVendorsGet200ResponseVendorsInner) Get() *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner {
	return v.value
}

func (v *NullableV1LldpInterfaceIdVendorsGet200ResponseVendorsInner) Set(val *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1LldpInterfaceIdVendorsGet200ResponseVendorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1LldpInterfaceIdVendorsGet200ResponseVendorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1LldpInterfaceIdVendorsGet200ResponseVendorsInner(val *V1LldpInterfaceIdVendorsGet200ResponseVendorsInner) *NullableV1LldpInterfaceIdVendorsGet200ResponseVendorsInner {
	return &NullableV1LldpInterfaceIdVendorsGet200ResponseVendorsInner{value: val, isSet: true}
}

func (v NullableV1LldpInterfaceIdVendorsGet200ResponseVendorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1LldpInterfaceIdVendorsGet200ResponseVendorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


