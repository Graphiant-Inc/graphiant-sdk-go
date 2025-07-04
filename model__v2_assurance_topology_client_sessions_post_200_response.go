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

// checks if the V2AssuranceTopologyClientSessionsPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2AssuranceTopologyClientSessionsPost200Response{}

// V2AssuranceTopologyClientSessionsPost200Response struct for V2AssuranceTopologyClientSessionsPost200Response
type V2AssuranceTopologyClientSessionsPost200Response struct {
	Sessions []V2AssuranceTopologyClientSessionDetailsPost200ResponseSession `json:"sessions,omitempty"`
}

// NewV2AssuranceTopologyClientSessionsPost200Response instantiates a new V2AssuranceTopologyClientSessionsPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssuranceTopologyClientSessionsPost200Response() *V2AssuranceTopologyClientSessionsPost200Response {
	this := V2AssuranceTopologyClientSessionsPost200Response{}
	return &this
}

// NewV2AssuranceTopologyClientSessionsPost200ResponseWithDefaults instantiates a new V2AssuranceTopologyClientSessionsPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssuranceTopologyClientSessionsPost200ResponseWithDefaults() *V2AssuranceTopologyClientSessionsPost200Response {
	this := V2AssuranceTopologyClientSessionsPost200Response{}
	return &this
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *V2AssuranceTopologyClientSessionsPost200Response) GetSessions() []V2AssuranceTopologyClientSessionDetailsPost200ResponseSession {
	if o == nil || IsNil(o.Sessions) {
		var ret []V2AssuranceTopologyClientSessionDetailsPost200ResponseSession
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyClientSessionsPost200Response) GetSessionsOk() ([]V2AssuranceTopologyClientSessionDetailsPost200ResponseSession, bool) {
	if o == nil || IsNil(o.Sessions) {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *V2AssuranceTopologyClientSessionsPost200Response) HasSessions() bool {
	if o != nil && !IsNil(o.Sessions) {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []V2AssuranceTopologyClientSessionDetailsPost200ResponseSession and assigns it to the Sessions field.
func (o *V2AssuranceTopologyClientSessionsPost200Response) SetSessions(v []V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) {
	o.Sessions = v
}

func (o V2AssuranceTopologyClientSessionsPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2AssuranceTopologyClientSessionsPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sessions) {
		toSerialize["sessions"] = o.Sessions
	}
	return toSerialize, nil
}

type NullableV2AssuranceTopologyClientSessionsPost200Response struct {
	value *V2AssuranceTopologyClientSessionsPost200Response
	isSet bool
}

func (v NullableV2AssuranceTopologyClientSessionsPost200Response) Get() *V2AssuranceTopologyClientSessionsPost200Response {
	return v.value
}

func (v *NullableV2AssuranceTopologyClientSessionsPost200Response) Set(val *V2AssuranceTopologyClientSessionsPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssuranceTopologyClientSessionsPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssuranceTopologyClientSessionsPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssuranceTopologyClientSessionsPost200Response(val *V2AssuranceTopologyClientSessionsPost200Response) *NullableV2AssuranceTopologyClientSessionsPost200Response {
	return &NullableV2AssuranceTopologyClientSessionsPost200Response{value: val, isSet: true}
}

func (v NullableV2AssuranceTopologyClientSessionsPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssuranceTopologyClientSessionsPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


