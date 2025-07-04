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

// checks if the V1AuthLoginPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1AuthLoginPost200Response{}

// V1AuthLoginPost200Response struct for V1AuthLoginPost200Response
type V1AuthLoginPost200Response struct {
	AccountType *string `json:"accountType,omitempty"`
	Auth *bool `json:"auth,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewV1AuthLoginPost200Response instantiates a new V1AuthLoginPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AuthLoginPost200Response() *V1AuthLoginPost200Response {
	this := V1AuthLoginPost200Response{}
	return &this
}

// NewV1AuthLoginPost200ResponseWithDefaults instantiates a new V1AuthLoginPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AuthLoginPost200ResponseWithDefaults() *V1AuthLoginPost200Response {
	this := V1AuthLoginPost200Response{}
	return &this
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *V1AuthLoginPost200Response) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthLoginPost200Response) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *V1AuthLoginPost200Response) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *V1AuthLoginPost200Response) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *V1AuthLoginPost200Response) GetAuth() bool {
	if o == nil || IsNil(o.Auth) {
		var ret bool
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthLoginPost200Response) GetAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.Auth) {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *V1AuthLoginPost200Response) HasAuth() bool {
	if o != nil && !IsNil(o.Auth) {
		return true
	}

	return false
}

// SetAuth gets a reference to the given bool and assigns it to the Auth field.
func (o *V1AuthLoginPost200Response) SetAuth(v bool) {
	o.Auth = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *V1AuthLoginPost200Response) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthLoginPost200Response) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *V1AuthLoginPost200Response) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *V1AuthLoginPost200Response) SetToken(v string) {
	o.Token = &v
}

func (o V1AuthLoginPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1AuthLoginPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.Auth) {
		toSerialize["auth"] = o.Auth
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableV1AuthLoginPost200Response struct {
	value *V1AuthLoginPost200Response
	isSet bool
}

func (v NullableV1AuthLoginPost200Response) Get() *V1AuthLoginPost200Response {
	return v.value
}

func (v *NullableV1AuthLoginPost200Response) Set(val *V1AuthLoginPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AuthLoginPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AuthLoginPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AuthLoginPost200Response(val *V1AuthLoginPost200Response) *NullableV1AuthLoginPost200Response {
	return &NullableV1AuthLoginPost200Response{value: val, isSet: true}
}

func (v NullableV1AuthLoginPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AuthLoginPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


