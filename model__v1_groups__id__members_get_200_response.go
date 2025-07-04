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

// checks if the V1GroupsIdMembersGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GroupsIdMembersGet200Response{}

// V1GroupsIdMembersGet200Response struct for V1GroupsIdMembersGet200Response
type V1GroupsIdMembersGet200Response struct {
	Users []V1GroupsIdMembersGet200ResponseUsersInner `json:"users,omitempty"`
}

// NewV1GroupsIdMembersGet200Response instantiates a new V1GroupsIdMembersGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GroupsIdMembersGet200Response() *V1GroupsIdMembersGet200Response {
	this := V1GroupsIdMembersGet200Response{}
	return &this
}

// NewV1GroupsIdMembersGet200ResponseWithDefaults instantiates a new V1GroupsIdMembersGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GroupsIdMembersGet200ResponseWithDefaults() *V1GroupsIdMembersGet200Response {
	this := V1GroupsIdMembersGet200Response{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *V1GroupsIdMembersGet200Response) GetUsers() []V1GroupsIdMembersGet200ResponseUsersInner {
	if o == nil || IsNil(o.Users) {
		var ret []V1GroupsIdMembersGet200ResponseUsersInner
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GroupsIdMembersGet200Response) GetUsersOk() ([]V1GroupsIdMembersGet200ResponseUsersInner, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *V1GroupsIdMembersGet200Response) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []V1GroupsIdMembersGet200ResponseUsersInner and assigns it to the Users field.
func (o *V1GroupsIdMembersGet200Response) SetUsers(v []V1GroupsIdMembersGet200ResponseUsersInner) {
	o.Users = v
}

func (o V1GroupsIdMembersGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GroupsIdMembersGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableV1GroupsIdMembersGet200Response struct {
	value *V1GroupsIdMembersGet200Response
	isSet bool
}

func (v NullableV1GroupsIdMembersGet200Response) Get() *V1GroupsIdMembersGet200Response {
	return v.value
}

func (v *NullableV1GroupsIdMembersGet200Response) Set(val *V1GroupsIdMembersGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GroupsIdMembersGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GroupsIdMembersGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GroupsIdMembersGet200Response(val *V1GroupsIdMembersGet200Response) *NullableV1GroupsIdMembersGet200Response {
	return &NullableV1GroupsIdMembersGet200Response{value: val, isSet: true}
}

func (v NullableV1GroupsIdMembersGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GroupsIdMembersGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


