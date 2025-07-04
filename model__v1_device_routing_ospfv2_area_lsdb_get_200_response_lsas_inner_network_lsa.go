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

// checks if the V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa{}

// V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa struct for V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa
type V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa struct {
	AttachedRouters []string `json:"attachedRouters,omitempty"`
	NetworkMask *int32 `json:"networkMask,omitempty"`
}

// NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa instantiates a new V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa() *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa {
	this := V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa{}
	return &this
}

// NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsaWithDefaults instantiates a new V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsaWithDefaults() *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa {
	this := V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa{}
	return &this
}

// GetAttachedRouters returns the AttachedRouters field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) GetAttachedRouters() []string {
	if o == nil || IsNil(o.AttachedRouters) {
		var ret []string
		return ret
	}
	return o.AttachedRouters
}

// GetAttachedRoutersOk returns a tuple with the AttachedRouters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) GetAttachedRoutersOk() ([]string, bool) {
	if o == nil || IsNil(o.AttachedRouters) {
		return nil, false
	}
	return o.AttachedRouters, true
}

// HasAttachedRouters returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) HasAttachedRouters() bool {
	if o != nil && !IsNil(o.AttachedRouters) {
		return true
	}

	return false
}

// SetAttachedRouters gets a reference to the given []string and assigns it to the AttachedRouters field.
func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) SetAttachedRouters(v []string) {
	o.AttachedRouters = v
}

// GetNetworkMask returns the NetworkMask field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) GetNetworkMask() int32 {
	if o == nil || IsNil(o.NetworkMask) {
		var ret int32
		return ret
	}
	return *o.NetworkMask
}

// GetNetworkMaskOk returns a tuple with the NetworkMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) GetNetworkMaskOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkMask) {
		return nil, false
	}
	return o.NetworkMask, true
}

// HasNetworkMask returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) HasNetworkMask() bool {
	if o != nil && !IsNil(o.NetworkMask) {
		return true
	}

	return false
}

// SetNetworkMask gets a reference to the given int32 and assigns it to the NetworkMask field.
func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) SetNetworkMask(v int32) {
	o.NetworkMask = &v
}

func (o V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttachedRouters) {
		toSerialize["attachedRouters"] = o.AttachedRouters
	}
	if !IsNil(o.NetworkMask) {
		toSerialize["networkMask"] = o.NetworkMask
	}
	return toSerialize, nil
}

type NullableV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa struct {
	value *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa
	isSet bool
}

func (v NullableV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) Get() *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa {
	return v.value
}

func (v *NullableV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) Set(val *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa(val *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) *NullableV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa {
	return &NullableV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa{value: val, isSet: true}
}

func (v NullableV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


