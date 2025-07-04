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

// checks if the V1ExtranetsPostRequestPolicyBranchesPrefixSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ExtranetsPostRequestPolicyBranchesPrefixSet{}

// V1ExtranetsPostRequestPolicyBranchesPrefixSet struct for V1ExtranetsPostRequestPolicyBranchesPrefixSet
type V1ExtranetsPostRequestPolicyBranchesPrefixSet struct {
	Description *string `json:"description,omitempty"`
	Entries *map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry `json:"entries,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewV1ExtranetsPostRequestPolicyBranchesPrefixSet instantiates a new V1ExtranetsPostRequestPolicyBranchesPrefixSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ExtranetsPostRequestPolicyBranchesPrefixSet() *V1ExtranetsPostRequestPolicyBranchesPrefixSet {
	this := V1ExtranetsPostRequestPolicyBranchesPrefixSet{}
	return &this
}

// NewV1ExtranetsPostRequestPolicyBranchesPrefixSetWithDefaults instantiates a new V1ExtranetsPostRequestPolicyBranchesPrefixSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ExtranetsPostRequestPolicyBranchesPrefixSetWithDefaults() *V1ExtranetsPostRequestPolicyBranchesPrefixSet {
	this := V1ExtranetsPostRequestPolicyBranchesPrefixSet{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) SetDescription(v string) {
	o.Description = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) GetEntries() map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry {
	if o == nil || IsNil(o.Entries) {
		var ret map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry
		return ret
	}
	return *o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) GetEntriesOk() (*map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry and assigns it to the Entries field.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) SetEntries(v map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry) {
	o.Entries = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) SetId(v int64) {
	o.Id = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1ExtranetsPostRequestPolicyBranchesPrefixSet) SetName(v string) {
	o.Name = &v
}

func (o V1ExtranetsPostRequestPolicyBranchesPrefixSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ExtranetsPostRequestPolicyBranchesPrefixSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV1ExtranetsPostRequestPolicyBranchesPrefixSet struct {
	value *V1ExtranetsPostRequestPolicyBranchesPrefixSet
	isSet bool
}

func (v NullableV1ExtranetsPostRequestPolicyBranchesPrefixSet) Get() *V1ExtranetsPostRequestPolicyBranchesPrefixSet {
	return v.value
}

func (v *NullableV1ExtranetsPostRequestPolicyBranchesPrefixSet) Set(val *V1ExtranetsPostRequestPolicyBranchesPrefixSet) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ExtranetsPostRequestPolicyBranchesPrefixSet) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ExtranetsPostRequestPolicyBranchesPrefixSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ExtranetsPostRequestPolicyBranchesPrefixSet(val *V1ExtranetsPostRequestPolicyBranchesPrefixSet) *NullableV1ExtranetsPostRequestPolicyBranchesPrefixSet {
	return &NullableV1ExtranetsPostRequestPolicyBranchesPrefixSet{value: val, isSet: true}
}

func (v NullableV1ExtranetsPostRequestPolicyBranchesPrefixSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ExtranetsPostRequestPolicyBranchesPrefixSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


