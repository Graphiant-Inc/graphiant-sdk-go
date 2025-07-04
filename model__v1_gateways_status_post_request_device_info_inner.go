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

// checks if the V1GatewaysStatusPostRequestDeviceInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GatewaysStatusPostRequestDeviceInfoInner{}

// V1GatewaysStatusPostRequestDeviceInfoInner struct for V1GatewaysStatusPostRequestDeviceInfoInner
type V1GatewaysStatusPostRequestDeviceInfoInner struct {
	DeviceId *int64 `json:"deviceId,omitempty"`
	InterfaceId *int64 `json:"interfaceId,omitempty"`
}

// NewV1GatewaysStatusPostRequestDeviceInfoInner instantiates a new V1GatewaysStatusPostRequestDeviceInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GatewaysStatusPostRequestDeviceInfoInner() *V1GatewaysStatusPostRequestDeviceInfoInner {
	this := V1GatewaysStatusPostRequestDeviceInfoInner{}
	return &this
}

// NewV1GatewaysStatusPostRequestDeviceInfoInnerWithDefaults instantiates a new V1GatewaysStatusPostRequestDeviceInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GatewaysStatusPostRequestDeviceInfoInnerWithDefaults() *V1GatewaysStatusPostRequestDeviceInfoInner {
	this := V1GatewaysStatusPostRequestDeviceInfoInner{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *V1GatewaysStatusPostRequestDeviceInfoInner) GetDeviceId() int64 {
	if o == nil || IsNil(o.DeviceId) {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GatewaysStatusPostRequestDeviceInfoInner) GetDeviceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *V1GatewaysStatusPostRequestDeviceInfoInner) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *V1GatewaysStatusPostRequestDeviceInfoInner) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *V1GatewaysStatusPostRequestDeviceInfoInner) GetInterfaceId() int64 {
	if o == nil || IsNil(o.InterfaceId) {
		var ret int64
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GatewaysStatusPostRequestDeviceInfoInner) GetInterfaceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.InterfaceId) {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *V1GatewaysStatusPostRequestDeviceInfoInner) HasInterfaceId() bool {
	if o != nil && !IsNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given int64 and assigns it to the InterfaceId field.
func (o *V1GatewaysStatusPostRequestDeviceInfoInner) SetInterfaceId(v int64) {
	o.InterfaceId = &v
}

func (o V1GatewaysStatusPostRequestDeviceInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GatewaysStatusPostRequestDeviceInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.InterfaceId) {
		toSerialize["interfaceId"] = o.InterfaceId
	}
	return toSerialize, nil
}

type NullableV1GatewaysStatusPostRequestDeviceInfoInner struct {
	value *V1GatewaysStatusPostRequestDeviceInfoInner
	isSet bool
}

func (v NullableV1GatewaysStatusPostRequestDeviceInfoInner) Get() *V1GatewaysStatusPostRequestDeviceInfoInner {
	return v.value
}

func (v *NullableV1GatewaysStatusPostRequestDeviceInfoInner) Set(val *V1GatewaysStatusPostRequestDeviceInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GatewaysStatusPostRequestDeviceInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GatewaysStatusPostRequestDeviceInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GatewaysStatusPostRequestDeviceInfoInner(val *V1GatewaysStatusPostRequestDeviceInfoInner) *NullableV1GatewaysStatusPostRequestDeviceInfoInner {
	return &NullableV1GatewaysStatusPostRequestDeviceInfoInner{value: val, isSet: true}
}

func (v NullableV1GatewaysStatusPostRequestDeviceInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GatewaysStatusPostRequestDeviceInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


