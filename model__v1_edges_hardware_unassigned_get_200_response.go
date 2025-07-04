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

// checks if the V1EdgesHardwareUnassignedGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EdgesHardwareUnassignedGet200Response{}

// V1EdgesHardwareUnassignedGet200Response struct for V1EdgesHardwareUnassignedGet200Response
type V1EdgesHardwareUnassignedGet200Response struct {
	Inventory []V1DevicesInventoryGet200ResponseInventoryInner `json:"inventory,omitempty"`
	IsNewCount *int32 `json:"isNewCount,omitempty"`
}

// NewV1EdgesHardwareUnassignedGet200Response instantiates a new V1EdgesHardwareUnassignedGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EdgesHardwareUnassignedGet200Response() *V1EdgesHardwareUnassignedGet200Response {
	this := V1EdgesHardwareUnassignedGet200Response{}
	return &this
}

// NewV1EdgesHardwareUnassignedGet200ResponseWithDefaults instantiates a new V1EdgesHardwareUnassignedGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EdgesHardwareUnassignedGet200ResponseWithDefaults() *V1EdgesHardwareUnassignedGet200Response {
	this := V1EdgesHardwareUnassignedGet200Response{}
	return &this
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *V1EdgesHardwareUnassignedGet200Response) GetInventory() []V1DevicesInventoryGet200ResponseInventoryInner {
	if o == nil || IsNil(o.Inventory) {
		var ret []V1DevicesInventoryGet200ResponseInventoryInner
		return ret
	}
	return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EdgesHardwareUnassignedGet200Response) GetInventoryOk() ([]V1DevicesInventoryGet200ResponseInventoryInner, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *V1EdgesHardwareUnassignedGet200Response) HasInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given []V1DevicesInventoryGet200ResponseInventoryInner and assigns it to the Inventory field.
func (o *V1EdgesHardwareUnassignedGet200Response) SetInventory(v []V1DevicesInventoryGet200ResponseInventoryInner) {
	o.Inventory = v
}

// GetIsNewCount returns the IsNewCount field value if set, zero value otherwise.
func (o *V1EdgesHardwareUnassignedGet200Response) GetIsNewCount() int32 {
	if o == nil || IsNil(o.IsNewCount) {
		var ret int32
		return ret
	}
	return *o.IsNewCount
}

// GetIsNewCountOk returns a tuple with the IsNewCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EdgesHardwareUnassignedGet200Response) GetIsNewCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IsNewCount) {
		return nil, false
	}
	return o.IsNewCount, true
}

// HasIsNewCount returns a boolean if a field has been set.
func (o *V1EdgesHardwareUnassignedGet200Response) HasIsNewCount() bool {
	if o != nil && !IsNil(o.IsNewCount) {
		return true
	}

	return false
}

// SetIsNewCount gets a reference to the given int32 and assigns it to the IsNewCount field.
func (o *V1EdgesHardwareUnassignedGet200Response) SetIsNewCount(v int32) {
	o.IsNewCount = &v
}

func (o V1EdgesHardwareUnassignedGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EdgesHardwareUnassignedGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.IsNewCount) {
		toSerialize["isNewCount"] = o.IsNewCount
	}
	return toSerialize, nil
}

type NullableV1EdgesHardwareUnassignedGet200Response struct {
	value *V1EdgesHardwareUnassignedGet200Response
	isSet bool
}

func (v NullableV1EdgesHardwareUnassignedGet200Response) Get() *V1EdgesHardwareUnassignedGet200Response {
	return v.value
}

func (v *NullableV1EdgesHardwareUnassignedGet200Response) Set(val *V1EdgesHardwareUnassignedGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EdgesHardwareUnassignedGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EdgesHardwareUnassignedGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EdgesHardwareUnassignedGet200Response(val *V1EdgesHardwareUnassignedGet200Response) *NullableV1EdgesHardwareUnassignedGet200Response {
	return &NullableV1EdgesHardwareUnassignedGet200Response{value: val, isSet: true}
}

func (v NullableV1EdgesHardwareUnassignedGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EdgesHardwareUnassignedGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


