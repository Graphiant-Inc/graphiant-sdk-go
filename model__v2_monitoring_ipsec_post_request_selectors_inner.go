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

// checks if the V2MonitoringIpsecPostRequestSelectorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2MonitoringIpsecPostRequestSelectorsInner{}

// V2MonitoringIpsecPostRequestSelectorsInner struct for V2MonitoringIpsecPostRequestSelectorsInner
type V2MonitoringIpsecPostRequestSelectorsInner struct {
	PeerGdi *string `json:"peerGdi,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewV2MonitoringIpsecPostRequestSelectorsInner instantiates a new V2MonitoringIpsecPostRequestSelectorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2MonitoringIpsecPostRequestSelectorsInner() *V2MonitoringIpsecPostRequestSelectorsInner {
	this := V2MonitoringIpsecPostRequestSelectorsInner{}
	return &this
}

// NewV2MonitoringIpsecPostRequestSelectorsInnerWithDefaults instantiates a new V2MonitoringIpsecPostRequestSelectorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2MonitoringIpsecPostRequestSelectorsInnerWithDefaults() *V2MonitoringIpsecPostRequestSelectorsInner {
	this := V2MonitoringIpsecPostRequestSelectorsInner{}
	return &this
}

// GetPeerGdi returns the PeerGdi field value if set, zero value otherwise.
func (o *V2MonitoringIpsecPostRequestSelectorsInner) GetPeerGdi() string {
	if o == nil || IsNil(o.PeerGdi) {
		var ret string
		return ret
	}
	return *o.PeerGdi
}

// GetPeerGdiOk returns a tuple with the PeerGdi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringIpsecPostRequestSelectorsInner) GetPeerGdiOk() (*string, bool) {
	if o == nil || IsNil(o.PeerGdi) {
		return nil, false
	}
	return o.PeerGdi, true
}

// HasPeerGdi returns a boolean if a field has been set.
func (o *V2MonitoringIpsecPostRequestSelectorsInner) HasPeerGdi() bool {
	if o != nil && !IsNil(o.PeerGdi) {
		return true
	}

	return false
}

// SetPeerGdi gets a reference to the given string and assigns it to the PeerGdi field.
func (o *V2MonitoringIpsecPostRequestSelectorsInner) SetPeerGdi(v string) {
	o.PeerGdi = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V2MonitoringIpsecPostRequestSelectorsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringIpsecPostRequestSelectorsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V2MonitoringIpsecPostRequestSelectorsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V2MonitoringIpsecPostRequestSelectorsInner) SetType(v string) {
	o.Type = &v
}

func (o V2MonitoringIpsecPostRequestSelectorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2MonitoringIpsecPostRequestSelectorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PeerGdi) {
		toSerialize["peerGdi"] = o.PeerGdi
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableV2MonitoringIpsecPostRequestSelectorsInner struct {
	value *V2MonitoringIpsecPostRequestSelectorsInner
	isSet bool
}

func (v NullableV2MonitoringIpsecPostRequestSelectorsInner) Get() *V2MonitoringIpsecPostRequestSelectorsInner {
	return v.value
}

func (v *NullableV2MonitoringIpsecPostRequestSelectorsInner) Set(val *V2MonitoringIpsecPostRequestSelectorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2MonitoringIpsecPostRequestSelectorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2MonitoringIpsecPostRequestSelectorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2MonitoringIpsecPostRequestSelectorsInner(val *V2MonitoringIpsecPostRequestSelectorsInner) *NullableV2MonitoringIpsecPostRequestSelectorsInner {
	return &NullableV2MonitoringIpsecPostRequestSelectorsInner{value: val, isSet: true}
}

func (v NullableV2MonitoringIpsecPostRequestSelectorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2MonitoringIpsecPostRequestSelectorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


