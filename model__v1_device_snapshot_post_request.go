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

// checks if the V1DeviceSnapshotPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DeviceSnapshotPostRequest{}

// V1DeviceSnapshotPostRequest struct for V1DeviceSnapshotPostRequest
type V1DeviceSnapshotPostRequest struct {
	Category *string `json:"category,omitempty"`
	DeviceId *int64 `json:"deviceId,omitempty"`
	Golden *bool `json:"golden,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewV1DeviceSnapshotPostRequest instantiates a new V1DeviceSnapshotPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DeviceSnapshotPostRequest() *V1DeviceSnapshotPostRequest {
	this := V1DeviceSnapshotPostRequest{}
	return &this
}

// NewV1DeviceSnapshotPostRequestWithDefaults instantiates a new V1DeviceSnapshotPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DeviceSnapshotPostRequestWithDefaults() *V1DeviceSnapshotPostRequest {
	this := V1DeviceSnapshotPostRequest{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *V1DeviceSnapshotPostRequest) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceSnapshotPostRequest) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *V1DeviceSnapshotPostRequest) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *V1DeviceSnapshotPostRequest) SetCategory(v string) {
	o.Category = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *V1DeviceSnapshotPostRequest) GetDeviceId() int64 {
	if o == nil || IsNil(o.DeviceId) {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceSnapshotPostRequest) GetDeviceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *V1DeviceSnapshotPostRequest) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *V1DeviceSnapshotPostRequest) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetGolden returns the Golden field value if set, zero value otherwise.
func (o *V1DeviceSnapshotPostRequest) GetGolden() bool {
	if o == nil || IsNil(o.Golden) {
		var ret bool
		return ret
	}
	return *o.Golden
}

// GetGoldenOk returns a tuple with the Golden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceSnapshotPostRequest) GetGoldenOk() (*bool, bool) {
	if o == nil || IsNil(o.Golden) {
		return nil, false
	}
	return o.Golden, true
}

// HasGolden returns a boolean if a field has been set.
func (o *V1DeviceSnapshotPostRequest) HasGolden() bool {
	if o != nil && !IsNil(o.Golden) {
		return true
	}

	return false
}

// SetGolden gets a reference to the given bool and assigns it to the Golden field.
func (o *V1DeviceSnapshotPostRequest) SetGolden(v bool) {
	o.Golden = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1DeviceSnapshotPostRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceSnapshotPostRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1DeviceSnapshotPostRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1DeviceSnapshotPostRequest) SetName(v string) {
	o.Name = &v
}

func (o V1DeviceSnapshotPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DeviceSnapshotPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.Golden) {
		toSerialize["golden"] = o.Golden
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV1DeviceSnapshotPostRequest struct {
	value *V1DeviceSnapshotPostRequest
	isSet bool
}

func (v NullableV1DeviceSnapshotPostRequest) Get() *V1DeviceSnapshotPostRequest {
	return v.value
}

func (v *NullableV1DeviceSnapshotPostRequest) Set(val *V1DeviceSnapshotPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DeviceSnapshotPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DeviceSnapshotPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DeviceSnapshotPostRequest(val *V1DeviceSnapshotPostRequest) *NullableV1DeviceSnapshotPostRequest {
	return &NullableV1DeviceSnapshotPostRequest{value: val, isSet: true}
}

func (v NullableV1DeviceSnapshotPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DeviceSnapshotPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


