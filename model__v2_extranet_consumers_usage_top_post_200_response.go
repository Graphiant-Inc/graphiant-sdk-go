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

// checks if the V2ExtranetConsumersUsageTopPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ExtranetConsumersUsageTopPost200Response{}

// V2ExtranetConsumersUsageTopPost200Response struct for V2ExtranetConsumersUsageTopPost200Response
type V2ExtranetConsumersUsageTopPost200Response struct {
	TopConsumers []V2ExtranetConsumersUsageTopPost200ResponseTopConsumersInner `json:"topConsumers,omitempty"`
}

// NewV2ExtranetConsumersUsageTopPost200Response instantiates a new V2ExtranetConsumersUsageTopPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ExtranetConsumersUsageTopPost200Response() *V2ExtranetConsumersUsageTopPost200Response {
	this := V2ExtranetConsumersUsageTopPost200Response{}
	return &this
}

// NewV2ExtranetConsumersUsageTopPost200ResponseWithDefaults instantiates a new V2ExtranetConsumersUsageTopPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ExtranetConsumersUsageTopPost200ResponseWithDefaults() *V2ExtranetConsumersUsageTopPost200Response {
	this := V2ExtranetConsumersUsageTopPost200Response{}
	return &this
}

// GetTopConsumers returns the TopConsumers field value if set, zero value otherwise.
func (o *V2ExtranetConsumersUsageTopPost200Response) GetTopConsumers() []V2ExtranetConsumersUsageTopPost200ResponseTopConsumersInner {
	if o == nil || IsNil(o.TopConsumers) {
		var ret []V2ExtranetConsumersUsageTopPost200ResponseTopConsumersInner
		return ret
	}
	return o.TopConsumers
}

// GetTopConsumersOk returns a tuple with the TopConsumers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ExtranetConsumersUsageTopPost200Response) GetTopConsumersOk() ([]V2ExtranetConsumersUsageTopPost200ResponseTopConsumersInner, bool) {
	if o == nil || IsNil(o.TopConsumers) {
		return nil, false
	}
	return o.TopConsumers, true
}

// HasTopConsumers returns a boolean if a field has been set.
func (o *V2ExtranetConsumersUsageTopPost200Response) HasTopConsumers() bool {
	if o != nil && !IsNil(o.TopConsumers) {
		return true
	}

	return false
}

// SetTopConsumers gets a reference to the given []V2ExtranetConsumersUsageTopPost200ResponseTopConsumersInner and assigns it to the TopConsumers field.
func (o *V2ExtranetConsumersUsageTopPost200Response) SetTopConsumers(v []V2ExtranetConsumersUsageTopPost200ResponseTopConsumersInner) {
	o.TopConsumers = v
}

func (o V2ExtranetConsumersUsageTopPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ExtranetConsumersUsageTopPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TopConsumers) {
		toSerialize["topConsumers"] = o.TopConsumers
	}
	return toSerialize, nil
}

type NullableV2ExtranetConsumersUsageTopPost200Response struct {
	value *V2ExtranetConsumersUsageTopPost200Response
	isSet bool
}

func (v NullableV2ExtranetConsumersUsageTopPost200Response) Get() *V2ExtranetConsumersUsageTopPost200Response {
	return v.value
}

func (v *NullableV2ExtranetConsumersUsageTopPost200Response) Set(val *V2ExtranetConsumersUsageTopPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ExtranetConsumersUsageTopPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ExtranetConsumersUsageTopPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ExtranetConsumersUsageTopPost200Response(val *V2ExtranetConsumersUsageTopPost200Response) *NullableV2ExtranetConsumersUsageTopPost200Response {
	return &NullableV2ExtranetConsumersUsageTopPost200Response{value: val, isSet: true}
}

func (v NullableV2ExtranetConsumersUsageTopPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ExtranetConsumersUsageTopPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


