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

// checks if the V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner{}

// V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner struct for V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner
type V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner struct {
	Accesses []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner `json:"accesses,omitempty"`
	GroupMembers []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerGroupMembersInner `json:"groupMembers,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Views []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerViewsInner `json:"views,omitempty"`
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner{}
	return &this
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner{}
	return &this
}

// GetAccesses returns the Accesses field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) GetAccesses() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner {
	if o == nil || IsNil(o.Accesses) {
		var ret []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner
		return ret
	}
	return o.Accesses
}

// GetAccessesOk returns a tuple with the Accesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) GetAccessesOk() ([]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner, bool) {
	if o == nil || IsNil(o.Accesses) {
		return nil, false
	}
	return o.Accesses, true
}

// HasAccesses returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) HasAccesses() bool {
	if o != nil && !IsNil(o.Accesses) {
		return true
	}

	return false
}

// SetAccesses gets a reference to the given []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner and assigns it to the Accesses field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) SetAccesses(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) {
	o.Accesses = v
}

// GetGroupMembers returns the GroupMembers field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) GetGroupMembers() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerGroupMembersInner {
	if o == nil || IsNil(o.GroupMembers) {
		var ret []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerGroupMembersInner
		return ret
	}
	return o.GroupMembers
}

// GetGroupMembersOk returns a tuple with the GroupMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) GetGroupMembersOk() ([]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerGroupMembersInner, bool) {
	if o == nil || IsNil(o.GroupMembers) {
		return nil, false
	}
	return o.GroupMembers, true
}

// HasGroupMembers returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) HasGroupMembers() bool {
	if o != nil && !IsNil(o.GroupMembers) {
		return true
	}

	return false
}

// SetGroupMembers gets a reference to the given []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerGroupMembersInner and assigns it to the GroupMembers field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) SetGroupMembers(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerGroupMembersInner) {
	o.GroupMembers = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) SetName(v string) {
	o.Name = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) GetViews() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerViewsInner {
	if o == nil || IsNil(o.Views) {
		var ret []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerViewsInner
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) GetViewsOk() ([]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerViewsInner, bool) {
	if o == nil || IsNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) HasViews() bool {
	if o != nil && !IsNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerViewsInner and assigns it to the Views field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) SetViews(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerViewsInner) {
	o.Views = v
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accesses) {
		toSerialize["accesses"] = o.Accesses
	}
	if !IsNil(o.GroupMembers) {
		toSerialize["groupMembers"] = o.GroupMembers
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Views) {
		toSerialize["views"] = o.Views
	}
	return toSerialize, nil
}

type NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner struct {
	value *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner
	isSet bool
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) Get() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner {
	return v.value
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) Set(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner {
	return &NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner{value: val, isSet: true}
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


