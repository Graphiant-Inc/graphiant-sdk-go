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

// checks if the V1DeviceRoutingOspfv2RibGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DeviceRoutingOspfv2RibGet200Response{}

// V1DeviceRoutingOspfv2RibGet200Response struct for V1DeviceRoutingOspfv2RibGet200Response
type V1DeviceRoutingOspfv2RibGet200Response struct {
	PageInfo *V1NatEntriesDeviceIdGet200ResponsePageInfo `json:"pageInfo,omitempty"`
	Routes []V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner `json:"routes,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewV1DeviceRoutingOspfv2RibGet200Response instantiates a new V1DeviceRoutingOspfv2RibGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DeviceRoutingOspfv2RibGet200Response() *V1DeviceRoutingOspfv2RibGet200Response {
	this := V1DeviceRoutingOspfv2RibGet200Response{}
	return &this
}

// NewV1DeviceRoutingOspfv2RibGet200ResponseWithDefaults instantiates a new V1DeviceRoutingOspfv2RibGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DeviceRoutingOspfv2RibGet200ResponseWithDefaults() *V1DeviceRoutingOspfv2RibGet200Response {
	this := V1DeviceRoutingOspfv2RibGet200Response{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2RibGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo {
	if o == nil || IsNil(o.PageInfo) {
		var ret V1NatEntriesDeviceIdGet200ResponsePageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2RibGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool) {
	if o == nil || IsNil(o.PageInfo) {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2RibGet200Response) HasPageInfo() bool {
	if o != nil && !IsNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given V1NatEntriesDeviceIdGet200ResponsePageInfo and assigns it to the PageInfo field.
func (o *V1DeviceRoutingOspfv2RibGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo) {
	o.PageInfo = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2RibGet200Response) GetRoutes() []V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner {
	if o == nil || IsNil(o.Routes) {
		var ret []V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2RibGet200Response) GetRoutesOk() ([]V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2RibGet200Response) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner and assigns it to the Routes field.
func (o *V1DeviceRoutingOspfv2RibGet200Response) SetRoutes(v []V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner) {
	o.Routes = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2RibGet200Response) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2RibGet200Response) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2RibGet200Response) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *V1DeviceRoutingOspfv2RibGet200Response) SetToken(v string) {
	o.Token = &v
}

func (o V1DeviceRoutingOspfv2RibGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DeviceRoutingOspfv2RibGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageInfo) {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableV1DeviceRoutingOspfv2RibGet200Response struct {
	value *V1DeviceRoutingOspfv2RibGet200Response
	isSet bool
}

func (v NullableV1DeviceRoutingOspfv2RibGet200Response) Get() *V1DeviceRoutingOspfv2RibGet200Response {
	return v.value
}

func (v *NullableV1DeviceRoutingOspfv2RibGet200Response) Set(val *V1DeviceRoutingOspfv2RibGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DeviceRoutingOspfv2RibGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DeviceRoutingOspfv2RibGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DeviceRoutingOspfv2RibGet200Response(val *V1DeviceRoutingOspfv2RibGet200Response) *NullableV1DeviceRoutingOspfv2RibGet200Response {
	return &NullableV1DeviceRoutingOspfv2RibGet200Response{value: val, isSet: true}
}

func (v NullableV1DeviceRoutingOspfv2RibGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DeviceRoutingOspfv2RibGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


