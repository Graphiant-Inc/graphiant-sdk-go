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

// checks if the V1DeviceRoutingBgpNbrStatsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DeviceRoutingBgpNbrStatsGet200Response{}

// V1DeviceRoutingBgpNbrStatsGet200Response struct for V1DeviceRoutingBgpNbrStatsGet200Response
type V1DeviceRoutingBgpNbrStatsGet200Response struct {
	PageInfo *V1NatEntriesDeviceIdGet200ResponsePageInfo `json:"pageInfo,omitempty"`
	Stats *V1DeviceRoutingBgpNbrStatsGet200ResponseStats `json:"stats,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewV1DeviceRoutingBgpNbrStatsGet200Response instantiates a new V1DeviceRoutingBgpNbrStatsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DeviceRoutingBgpNbrStatsGet200Response() *V1DeviceRoutingBgpNbrStatsGet200Response {
	this := V1DeviceRoutingBgpNbrStatsGet200Response{}
	return &this
}

// NewV1DeviceRoutingBgpNbrStatsGet200ResponseWithDefaults instantiates a new V1DeviceRoutingBgpNbrStatsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DeviceRoutingBgpNbrStatsGet200ResponseWithDefaults() *V1DeviceRoutingBgpNbrStatsGet200Response {
	this := V1DeviceRoutingBgpNbrStatsGet200Response{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo {
	if o == nil || IsNil(o.PageInfo) {
		var ret V1NatEntriesDeviceIdGet200ResponsePageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool) {
	if o == nil || IsNil(o.PageInfo) {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) HasPageInfo() bool {
	if o != nil && !IsNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given V1NatEntriesDeviceIdGet200ResponsePageInfo and assigns it to the PageInfo field.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo) {
	o.PageInfo = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetStats() V1DeviceRoutingBgpNbrStatsGet200ResponseStats {
	if o == nil || IsNil(o.Stats) {
		var ret V1DeviceRoutingBgpNbrStatsGet200ResponseStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetStatsOk() (*V1DeviceRoutingBgpNbrStatsGet200ResponseStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given V1DeviceRoutingBgpNbrStatsGet200ResponseStats and assigns it to the Stats field.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) SetStats(v V1DeviceRoutingBgpNbrStatsGet200ResponseStats) {
	o.Stats = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *V1DeviceRoutingBgpNbrStatsGet200Response) SetToken(v string) {
	o.Token = &v
}

func (o V1DeviceRoutingBgpNbrStatsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DeviceRoutingBgpNbrStatsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageInfo) {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableV1DeviceRoutingBgpNbrStatsGet200Response struct {
	value *V1DeviceRoutingBgpNbrStatsGet200Response
	isSet bool
}

func (v NullableV1DeviceRoutingBgpNbrStatsGet200Response) Get() *V1DeviceRoutingBgpNbrStatsGet200Response {
	return v.value
}

func (v *NullableV1DeviceRoutingBgpNbrStatsGet200Response) Set(val *V1DeviceRoutingBgpNbrStatsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DeviceRoutingBgpNbrStatsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DeviceRoutingBgpNbrStatsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DeviceRoutingBgpNbrStatsGet200Response(val *V1DeviceRoutingBgpNbrStatsGet200Response) *NullableV1DeviceRoutingBgpNbrStatsGet200Response {
	return &NullableV1DeviceRoutingBgpNbrStatsGet200Response{value: val, isSet: true}
}

func (v NullableV1DeviceRoutingBgpNbrStatsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DeviceRoutingBgpNbrStatsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


