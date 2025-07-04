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

// checks if the V1BackboneHealthOverviewPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BackboneHealthOverviewPostRequest{}

// V1BackboneHealthOverviewPostRequest struct for V1BackboneHealthOverviewPostRequest
type V1BackboneHealthOverviewPostRequest struct {
	Dimensions *V1BackboneHealthOverviewPostRequestDimensions `json:"dimensions,omitempty"`
	Filter *V1BackboneHealthOverviewPostRequestFilter `json:"filter,omitempty"`
}

// NewV1BackboneHealthOverviewPostRequest instantiates a new V1BackboneHealthOverviewPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BackboneHealthOverviewPostRequest() *V1BackboneHealthOverviewPostRequest {
	this := V1BackboneHealthOverviewPostRequest{}
	return &this
}

// NewV1BackboneHealthOverviewPostRequestWithDefaults instantiates a new V1BackboneHealthOverviewPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BackboneHealthOverviewPostRequestWithDefaults() *V1BackboneHealthOverviewPostRequest {
	this := V1BackboneHealthOverviewPostRequest{}
	return &this
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPostRequest) GetDimensions() V1BackboneHealthOverviewPostRequestDimensions {
	if o == nil || IsNil(o.Dimensions) {
		var ret V1BackboneHealthOverviewPostRequestDimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPostRequest) GetDimensionsOk() (*V1BackboneHealthOverviewPostRequestDimensions, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPostRequest) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given V1BackboneHealthOverviewPostRequestDimensions and assigns it to the Dimensions field.
func (o *V1BackboneHealthOverviewPostRequest) SetDimensions(v V1BackboneHealthOverviewPostRequestDimensions) {
	o.Dimensions = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPostRequest) GetFilter() V1BackboneHealthOverviewPostRequestFilter {
	if o == nil || IsNil(o.Filter) {
		var ret V1BackboneHealthOverviewPostRequestFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPostRequest) GetFilterOk() (*V1BackboneHealthOverviewPostRequestFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPostRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given V1BackboneHealthOverviewPostRequestFilter and assigns it to the Filter field.
func (o *V1BackboneHealthOverviewPostRequest) SetFilter(v V1BackboneHealthOverviewPostRequestFilter) {
	o.Filter = &v
}

func (o V1BackboneHealthOverviewPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BackboneHealthOverviewPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	return toSerialize, nil
}

type NullableV1BackboneHealthOverviewPostRequest struct {
	value *V1BackboneHealthOverviewPostRequest
	isSet bool
}

func (v NullableV1BackboneHealthOverviewPostRequest) Get() *V1BackboneHealthOverviewPostRequest {
	return v.value
}

func (v *NullableV1BackboneHealthOverviewPostRequest) Set(val *V1BackboneHealthOverviewPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BackboneHealthOverviewPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BackboneHealthOverviewPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BackboneHealthOverviewPostRequest(val *V1BackboneHealthOverviewPostRequest) *NullableV1BackboneHealthOverviewPostRequest {
	return &NullableV1BackboneHealthOverviewPostRequest{value: val, isSet: true}
}

func (v NullableV1BackboneHealthOverviewPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BackboneHealthOverviewPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


