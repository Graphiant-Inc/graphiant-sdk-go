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

// checks if the V1DevicesOauthRedirectGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesOauthRedirectGet200Response{}

// V1DevicesOauthRedirectGet200Response struct for V1DevicesOauthRedirectGet200Response
type V1DevicesOauthRedirectGet200Response struct {
	OnboardingRedirection *V1DevicesOauthRedirectGet200ResponseOnboardingRedirection `json:"onboardingRedirection,omitempty"`
}

// NewV1DevicesOauthRedirectGet200Response instantiates a new V1DevicesOauthRedirectGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesOauthRedirectGet200Response() *V1DevicesOauthRedirectGet200Response {
	this := V1DevicesOauthRedirectGet200Response{}
	return &this
}

// NewV1DevicesOauthRedirectGet200ResponseWithDefaults instantiates a new V1DevicesOauthRedirectGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesOauthRedirectGet200ResponseWithDefaults() *V1DevicesOauthRedirectGet200Response {
	this := V1DevicesOauthRedirectGet200Response{}
	return &this
}

// GetOnboardingRedirection returns the OnboardingRedirection field value if set, zero value otherwise.
func (o *V1DevicesOauthRedirectGet200Response) GetOnboardingRedirection() V1DevicesOauthRedirectGet200ResponseOnboardingRedirection {
	if o == nil || IsNil(o.OnboardingRedirection) {
		var ret V1DevicesOauthRedirectGet200ResponseOnboardingRedirection
		return ret
	}
	return *o.OnboardingRedirection
}

// GetOnboardingRedirectionOk returns a tuple with the OnboardingRedirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesOauthRedirectGet200Response) GetOnboardingRedirectionOk() (*V1DevicesOauthRedirectGet200ResponseOnboardingRedirection, bool) {
	if o == nil || IsNil(o.OnboardingRedirection) {
		return nil, false
	}
	return o.OnboardingRedirection, true
}

// HasOnboardingRedirection returns a boolean if a field has been set.
func (o *V1DevicesOauthRedirectGet200Response) HasOnboardingRedirection() bool {
	if o != nil && !IsNil(o.OnboardingRedirection) {
		return true
	}

	return false
}

// SetOnboardingRedirection gets a reference to the given V1DevicesOauthRedirectGet200ResponseOnboardingRedirection and assigns it to the OnboardingRedirection field.
func (o *V1DevicesOauthRedirectGet200Response) SetOnboardingRedirection(v V1DevicesOauthRedirectGet200ResponseOnboardingRedirection) {
	o.OnboardingRedirection = &v
}

func (o V1DevicesOauthRedirectGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesOauthRedirectGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnboardingRedirection) {
		toSerialize["onboardingRedirection"] = o.OnboardingRedirection
	}
	return toSerialize, nil
}

type NullableV1DevicesOauthRedirectGet200Response struct {
	value *V1DevicesOauthRedirectGet200Response
	isSet bool
}

func (v NullableV1DevicesOauthRedirectGet200Response) Get() *V1DevicesOauthRedirectGet200Response {
	return v.value
}

func (v *NullableV1DevicesOauthRedirectGet200Response) Set(val *V1DevicesOauthRedirectGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesOauthRedirectGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesOauthRedirectGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesOauthRedirectGet200Response(val *V1DevicesOauthRedirectGet200Response) *NullableV1DevicesOauthRedirectGet200Response {
	return &NullableV1DevicesOauthRedirectGet200Response{value: val, isSet: true}
}

func (v NullableV1DevicesOauthRedirectGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesOauthRedirectGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


