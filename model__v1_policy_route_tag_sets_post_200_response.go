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

// checks if the V1PolicyRouteTagSetsPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PolicyRouteTagSetsPost200Response{}

// V1PolicyRouteTagSetsPost200Response struct for V1PolicyRouteTagSetsPost200Response
type V1PolicyRouteTagSetsPost200Response struct {
	Id *int64 `json:"id,omitempty"`
}

// NewV1PolicyRouteTagSetsPost200Response instantiates a new V1PolicyRouteTagSetsPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PolicyRouteTagSetsPost200Response() *V1PolicyRouteTagSetsPost200Response {
	this := V1PolicyRouteTagSetsPost200Response{}
	return &this
}

// NewV1PolicyRouteTagSetsPost200ResponseWithDefaults instantiates a new V1PolicyRouteTagSetsPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PolicyRouteTagSetsPost200ResponseWithDefaults() *V1PolicyRouteTagSetsPost200Response {
	this := V1PolicyRouteTagSetsPost200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsPost200Response) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsPost200Response) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsPost200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *V1PolicyRouteTagSetsPost200Response) SetId(v int64) {
	o.Id = &v
}

func (o V1PolicyRouteTagSetsPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PolicyRouteTagSetsPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableV1PolicyRouteTagSetsPost200Response struct {
	value *V1PolicyRouteTagSetsPost200Response
	isSet bool
}

func (v NullableV1PolicyRouteTagSetsPost200Response) Get() *V1PolicyRouteTagSetsPost200Response {
	return v.value
}

func (v *NullableV1PolicyRouteTagSetsPost200Response) Set(val *V1PolicyRouteTagSetsPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PolicyRouteTagSetsPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PolicyRouteTagSetsPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PolicyRouteTagSetsPost200Response(val *V1PolicyRouteTagSetsPost200Response) *NullableV1PolicyRouteTagSetsPost200Response {
	return &NullableV1PolicyRouteTagSetsPost200Response{value: val, isSet: true}
}

func (v NullableV1PolicyRouteTagSetsPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PolicyRouteTagSetsPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


