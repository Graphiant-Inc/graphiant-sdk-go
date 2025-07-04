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

// checks if the V1GlobalAppsAppListsGet200ResponseEntriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalAppsAppListsGet200ResponseEntriesInner{}

// V1GlobalAppsAppListsGet200ResponseEntriesInner struct for V1GlobalAppsAppListsGet200ResponseEntriesInner
type V1GlobalAppsAppListsGet200ResponseEntriesInner struct {
	AppCount *int32 `json:"appCount,omitempty"`
	AppList *V1GlobalAppsAppListsGet200ResponseEntriesInnerAppList `json:"appList,omitempty"`
	PolicyReferenceCount *int32 `json:"policyReferenceCount,omitempty"`
}

// NewV1GlobalAppsAppListsGet200ResponseEntriesInner instantiates a new V1GlobalAppsAppListsGet200ResponseEntriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalAppsAppListsGet200ResponseEntriesInner() *V1GlobalAppsAppListsGet200ResponseEntriesInner {
	this := V1GlobalAppsAppListsGet200ResponseEntriesInner{}
	return &this
}

// NewV1GlobalAppsAppListsGet200ResponseEntriesInnerWithDefaults instantiates a new V1GlobalAppsAppListsGet200ResponseEntriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalAppsAppListsGet200ResponseEntriesInnerWithDefaults() *V1GlobalAppsAppListsGet200ResponseEntriesInner {
	this := V1GlobalAppsAppListsGet200ResponseEntriesInner{}
	return &this
}

// GetAppCount returns the AppCount field value if set, zero value otherwise.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) GetAppCount() int32 {
	if o == nil || IsNil(o.AppCount) {
		var ret int32
		return ret
	}
	return *o.AppCount
}

// GetAppCountOk returns a tuple with the AppCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) GetAppCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AppCount) {
		return nil, false
	}
	return o.AppCount, true
}

// HasAppCount returns a boolean if a field has been set.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) HasAppCount() bool {
	if o != nil && !IsNil(o.AppCount) {
		return true
	}

	return false
}

// SetAppCount gets a reference to the given int32 and assigns it to the AppCount field.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) SetAppCount(v int32) {
	o.AppCount = &v
}

// GetAppList returns the AppList field value if set, zero value otherwise.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) GetAppList() V1GlobalAppsAppListsGet200ResponseEntriesInnerAppList {
	if o == nil || IsNil(o.AppList) {
		var ret V1GlobalAppsAppListsGet200ResponseEntriesInnerAppList
		return ret
	}
	return *o.AppList
}

// GetAppListOk returns a tuple with the AppList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) GetAppListOk() (*V1GlobalAppsAppListsGet200ResponseEntriesInnerAppList, bool) {
	if o == nil || IsNil(o.AppList) {
		return nil, false
	}
	return o.AppList, true
}

// HasAppList returns a boolean if a field has been set.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) HasAppList() bool {
	if o != nil && !IsNil(o.AppList) {
		return true
	}

	return false
}

// SetAppList gets a reference to the given V1GlobalAppsAppListsGet200ResponseEntriesInnerAppList and assigns it to the AppList field.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) SetAppList(v V1GlobalAppsAppListsGet200ResponseEntriesInnerAppList) {
	o.AppList = &v
}

// GetPolicyReferenceCount returns the PolicyReferenceCount field value if set, zero value otherwise.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) GetPolicyReferenceCount() int32 {
	if o == nil || IsNil(o.PolicyReferenceCount) {
		var ret int32
		return ret
	}
	return *o.PolicyReferenceCount
}

// GetPolicyReferenceCountOk returns a tuple with the PolicyReferenceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) GetPolicyReferenceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PolicyReferenceCount) {
		return nil, false
	}
	return o.PolicyReferenceCount, true
}

// HasPolicyReferenceCount returns a boolean if a field has been set.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) HasPolicyReferenceCount() bool {
	if o != nil && !IsNil(o.PolicyReferenceCount) {
		return true
	}

	return false
}

// SetPolicyReferenceCount gets a reference to the given int32 and assigns it to the PolicyReferenceCount field.
func (o *V1GlobalAppsAppListsGet200ResponseEntriesInner) SetPolicyReferenceCount(v int32) {
	o.PolicyReferenceCount = &v
}

func (o V1GlobalAppsAppListsGet200ResponseEntriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalAppsAppListsGet200ResponseEntriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppCount) {
		toSerialize["appCount"] = o.AppCount
	}
	if !IsNil(o.AppList) {
		toSerialize["appList"] = o.AppList
	}
	if !IsNil(o.PolicyReferenceCount) {
		toSerialize["policyReferenceCount"] = o.PolicyReferenceCount
	}
	return toSerialize, nil
}

type NullableV1GlobalAppsAppListsGet200ResponseEntriesInner struct {
	value *V1GlobalAppsAppListsGet200ResponseEntriesInner
	isSet bool
}

func (v NullableV1GlobalAppsAppListsGet200ResponseEntriesInner) Get() *V1GlobalAppsAppListsGet200ResponseEntriesInner {
	return v.value
}

func (v *NullableV1GlobalAppsAppListsGet200ResponseEntriesInner) Set(val *V1GlobalAppsAppListsGet200ResponseEntriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalAppsAppListsGet200ResponseEntriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalAppsAppListsGet200ResponseEntriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalAppsAppListsGet200ResponseEntriesInner(val *V1GlobalAppsAppListsGet200ResponseEntriesInner) *NullableV1GlobalAppsAppListsGet200ResponseEntriesInner {
	return &NullableV1GlobalAppsAppListsGet200ResponseEntriesInner{value: val, isSet: true}
}

func (v NullableV1GlobalAppsAppListsGet200ResponseEntriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalAppsAppListsGet200ResponseEntriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


