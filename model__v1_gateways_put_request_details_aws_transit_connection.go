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

// checks if the V1GatewaysPutRequestDetailsAwsTransitConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GatewaysPutRequestDetailsAwsTransitConnection{}

// V1GatewaysPutRequestDetailsAwsTransitConnection struct for V1GatewaysPutRequestDetailsAwsTransitConnection
type V1GatewaysPutRequestDetailsAwsTransitConnection struct {
	Credentials *V1GatewaysPutRequestDetailsAwsTransitConnectionCredentials `json:"credentials,omitempty"`
	CustomerAsn *int32 `json:"customerAsn,omitempty"`
	Gateway *V1GatewaysPutRequestDetailsAwsTransitConnectionGateway `json:"gateway,omitempty"`
	Region *string `json:"region,omitempty"`
}

// NewV1GatewaysPutRequestDetailsAwsTransitConnection instantiates a new V1GatewaysPutRequestDetailsAwsTransitConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GatewaysPutRequestDetailsAwsTransitConnection() *V1GatewaysPutRequestDetailsAwsTransitConnection {
	this := V1GatewaysPutRequestDetailsAwsTransitConnection{}
	return &this
}

// NewV1GatewaysPutRequestDetailsAwsTransitConnectionWithDefaults instantiates a new V1GatewaysPutRequestDetailsAwsTransitConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GatewaysPutRequestDetailsAwsTransitConnectionWithDefaults() *V1GatewaysPutRequestDetailsAwsTransitConnection {
	this := V1GatewaysPutRequestDetailsAwsTransitConnection{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) GetCredentials() V1GatewaysPutRequestDetailsAwsTransitConnectionCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret V1GatewaysPutRequestDetailsAwsTransitConnectionCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) GetCredentialsOk() (*V1GatewaysPutRequestDetailsAwsTransitConnectionCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given V1GatewaysPutRequestDetailsAwsTransitConnectionCredentials and assigns it to the Credentials field.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) SetCredentials(v V1GatewaysPutRequestDetailsAwsTransitConnectionCredentials) {
	o.Credentials = &v
}

// GetCustomerAsn returns the CustomerAsn field value if set, zero value otherwise.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) GetCustomerAsn() int32 {
	if o == nil || IsNil(o.CustomerAsn) {
		var ret int32
		return ret
	}
	return *o.CustomerAsn
}

// GetCustomerAsnOk returns a tuple with the CustomerAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) GetCustomerAsnOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomerAsn) {
		return nil, false
	}
	return o.CustomerAsn, true
}

// HasCustomerAsn returns a boolean if a field has been set.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) HasCustomerAsn() bool {
	if o != nil && !IsNil(o.CustomerAsn) {
		return true
	}

	return false
}

// SetCustomerAsn gets a reference to the given int32 and assigns it to the CustomerAsn field.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) SetCustomerAsn(v int32) {
	o.CustomerAsn = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) GetGateway() V1GatewaysPutRequestDetailsAwsTransitConnectionGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret V1GatewaysPutRequestDetailsAwsTransitConnectionGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) GetGatewayOk() (*V1GatewaysPutRequestDetailsAwsTransitConnectionGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given V1GatewaysPutRequestDetailsAwsTransitConnectionGateway and assigns it to the Gateway field.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) SetGateway(v V1GatewaysPutRequestDetailsAwsTransitConnectionGateway) {
	o.Gateway = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *V1GatewaysPutRequestDetailsAwsTransitConnection) SetRegion(v string) {
	o.Region = &v
}

func (o V1GatewaysPutRequestDetailsAwsTransitConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GatewaysPutRequestDetailsAwsTransitConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.CustomerAsn) {
		toSerialize["customerAsn"] = o.CustomerAsn
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableV1GatewaysPutRequestDetailsAwsTransitConnection struct {
	value *V1GatewaysPutRequestDetailsAwsTransitConnection
	isSet bool
}

func (v NullableV1GatewaysPutRequestDetailsAwsTransitConnection) Get() *V1GatewaysPutRequestDetailsAwsTransitConnection {
	return v.value
}

func (v *NullableV1GatewaysPutRequestDetailsAwsTransitConnection) Set(val *V1GatewaysPutRequestDetailsAwsTransitConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GatewaysPutRequestDetailsAwsTransitConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GatewaysPutRequestDetailsAwsTransitConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GatewaysPutRequestDetailsAwsTransitConnection(val *V1GatewaysPutRequestDetailsAwsTransitConnection) *NullableV1GatewaysPutRequestDetailsAwsTransitConnection {
	return &NullableV1GatewaysPutRequestDetailsAwsTransitConnection{value: val, isSet: true}
}

func (v NullableV1GatewaysPutRequestDetailsAwsTransitConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GatewaysPutRequestDetailsAwsTransitConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


