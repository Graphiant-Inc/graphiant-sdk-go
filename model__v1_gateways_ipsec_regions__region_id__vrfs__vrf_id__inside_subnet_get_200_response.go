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

// checks if the V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response{}

// V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response struct for V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response
type V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response struct {
	Ipv4Subnet *string `json:"ipv4Subnet,omitempty"`
	Ipv6Subnet *string `json:"ipv6Subnet,omitempty"`
}

// NewV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response instantiates a new V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response() *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response {
	this := V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response{}
	return &this
}

// NewV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200ResponseWithDefaults instantiates a new V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200ResponseWithDefaults() *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response {
	this := V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response{}
	return &this
}

// GetIpv4Subnet returns the Ipv4Subnet field value if set, zero value otherwise.
func (o *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) GetIpv4Subnet() string {
	if o == nil || IsNil(o.Ipv4Subnet) {
		var ret string
		return ret
	}
	return *o.Ipv4Subnet
}

// GetIpv4SubnetOk returns a tuple with the Ipv4Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) GetIpv4SubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Subnet) {
		return nil, false
	}
	return o.Ipv4Subnet, true
}

// HasIpv4Subnet returns a boolean if a field has been set.
func (o *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) HasIpv4Subnet() bool {
	if o != nil && !IsNil(o.Ipv4Subnet) {
		return true
	}

	return false
}

// SetIpv4Subnet gets a reference to the given string and assigns it to the Ipv4Subnet field.
func (o *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) SetIpv4Subnet(v string) {
	o.Ipv4Subnet = &v
}

// GetIpv6Subnet returns the Ipv6Subnet field value if set, zero value otherwise.
func (o *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) GetIpv6Subnet() string {
	if o == nil || IsNil(o.Ipv6Subnet) {
		var ret string
		return ret
	}
	return *o.Ipv6Subnet
}

// GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) GetIpv6SubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6Subnet) {
		return nil, false
	}
	return o.Ipv6Subnet, true
}

// HasIpv6Subnet returns a boolean if a field has been set.
func (o *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) HasIpv6Subnet() bool {
	if o != nil && !IsNil(o.Ipv6Subnet) {
		return true
	}

	return false
}

// SetIpv6Subnet gets a reference to the given string and assigns it to the Ipv6Subnet field.
func (o *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) SetIpv6Subnet(v string) {
	o.Ipv6Subnet = &v
}

func (o V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4Subnet) {
		toSerialize["ipv4Subnet"] = o.Ipv4Subnet
	}
	if !IsNil(o.Ipv6Subnet) {
		toSerialize["ipv6Subnet"] = o.Ipv6Subnet
	}
	return toSerialize, nil
}

type NullableV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response struct {
	value *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response
	isSet bool
}

func (v NullableV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) Get() *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response {
	return v.value
}

func (v *NullableV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) Set(val *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response(val *V1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) *NullableV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response {
	return &NullableV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response{value: val, isSet: true}
}

func (v NullableV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GatewaysIpsecRegionsRegionIdVrfsVrfIdInsideSubnetGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


