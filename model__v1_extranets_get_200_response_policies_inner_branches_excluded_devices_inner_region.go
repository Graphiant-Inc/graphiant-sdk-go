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

// checks if the V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion{}

// V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion struct for V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion
type V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion{}
	return &this
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegionWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegionWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) SetName(v string) {
	o.Name = &v
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion struct {
	value *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion
	isSet bool
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) Get() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion {
	return v.value
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) Set(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion {
	return &NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion{value: val, isSet: true}
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


