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

// checks if the V1ActivityLogsPost200ResponseFilterEntitiesValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ActivityLogsPost200ResponseFilterEntitiesValue{}

// V1ActivityLogsPost200ResponseFilterEntitiesValue struct for V1ActivityLogsPost200ResponseFilterEntitiesValue
type V1ActivityLogsPost200ResponseFilterEntitiesValue struct {
	Items []V1ActivityLogsPostRequestSelectorJobEntity `json:"items,omitempty"`
}

// NewV1ActivityLogsPost200ResponseFilterEntitiesValue instantiates a new V1ActivityLogsPost200ResponseFilterEntitiesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ActivityLogsPost200ResponseFilterEntitiesValue() *V1ActivityLogsPost200ResponseFilterEntitiesValue {
	this := V1ActivityLogsPost200ResponseFilterEntitiesValue{}
	return &this
}

// NewV1ActivityLogsPost200ResponseFilterEntitiesValueWithDefaults instantiates a new V1ActivityLogsPost200ResponseFilterEntitiesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ActivityLogsPost200ResponseFilterEntitiesValueWithDefaults() *V1ActivityLogsPost200ResponseFilterEntitiesValue {
	this := V1ActivityLogsPost200ResponseFilterEntitiesValue{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *V1ActivityLogsPost200ResponseFilterEntitiesValue) GetItems() []V1ActivityLogsPostRequestSelectorJobEntity {
	if o == nil || IsNil(o.Items) {
		var ret []V1ActivityLogsPostRequestSelectorJobEntity
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPost200ResponseFilterEntitiesValue) GetItemsOk() ([]V1ActivityLogsPostRequestSelectorJobEntity, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *V1ActivityLogsPost200ResponseFilterEntitiesValue) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []V1ActivityLogsPostRequestSelectorJobEntity and assigns it to the Items field.
func (o *V1ActivityLogsPost200ResponseFilterEntitiesValue) SetItems(v []V1ActivityLogsPostRequestSelectorJobEntity) {
	o.Items = v
}

func (o V1ActivityLogsPost200ResponseFilterEntitiesValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ActivityLogsPost200ResponseFilterEntitiesValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableV1ActivityLogsPost200ResponseFilterEntitiesValue struct {
	value *V1ActivityLogsPost200ResponseFilterEntitiesValue
	isSet bool
}

func (v NullableV1ActivityLogsPost200ResponseFilterEntitiesValue) Get() *V1ActivityLogsPost200ResponseFilterEntitiesValue {
	return v.value
}

func (v *NullableV1ActivityLogsPost200ResponseFilterEntitiesValue) Set(val *V1ActivityLogsPost200ResponseFilterEntitiesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ActivityLogsPost200ResponseFilterEntitiesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ActivityLogsPost200ResponseFilterEntitiesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ActivityLogsPost200ResponseFilterEntitiesValue(val *V1ActivityLogsPost200ResponseFilterEntitiesValue) *NullableV1ActivityLogsPost200ResponseFilterEntitiesValue {
	return &NullableV1ActivityLogsPost200ResponseFilterEntitiesValue{value: val, isSet: true}
}

func (v NullableV1ActivityLogsPost200ResponseFilterEntitiesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ActivityLogsPost200ResponseFilterEntitiesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


