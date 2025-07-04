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

// checks if the V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner{}

// V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner struct for V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner
type V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner struct {
	Context *string `json:"context,omitempty"`
	GroupContextMatch *string `json:"groupContextMatch,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ReadView *string `json:"readView,omitempty"`
	SecurityLevel *string `json:"securityLevel,omitempty"`
	WriteView *string `json:"writeView,omitempty"`
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner{}
	return &this
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) SetContext(v string) {
	o.Context = &v
}

// GetGroupContextMatch returns the GroupContextMatch field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetGroupContextMatch() string {
	if o == nil || IsNil(o.GroupContextMatch) {
		var ret string
		return ret
	}
	return *o.GroupContextMatch
}

// GetGroupContextMatchOk returns a tuple with the GroupContextMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetGroupContextMatchOk() (*string, bool) {
	if o == nil || IsNil(o.GroupContextMatch) {
		return nil, false
	}
	return o.GroupContextMatch, true
}

// HasGroupContextMatch returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) HasGroupContextMatch() bool {
	if o != nil && !IsNil(o.GroupContextMatch) {
		return true
	}

	return false
}

// SetGroupContextMatch gets a reference to the given string and assigns it to the GroupContextMatch field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) SetGroupContextMatch(v string) {
	o.GroupContextMatch = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) SetName(v string) {
	o.Name = &v
}

// GetReadView returns the ReadView field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetReadView() string {
	if o == nil || IsNil(o.ReadView) {
		var ret string
		return ret
	}
	return *o.ReadView
}

// GetReadViewOk returns a tuple with the ReadView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetReadViewOk() (*string, bool) {
	if o == nil || IsNil(o.ReadView) {
		return nil, false
	}
	return o.ReadView, true
}

// HasReadView returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) HasReadView() bool {
	if o != nil && !IsNil(o.ReadView) {
		return true
	}

	return false
}

// SetReadView gets a reference to the given string and assigns it to the ReadView field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) SetReadView(v string) {
	o.ReadView = &v
}

// GetSecurityLevel returns the SecurityLevel field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetSecurityLevel() string {
	if o == nil || IsNil(o.SecurityLevel) {
		var ret string
		return ret
	}
	return *o.SecurityLevel
}

// GetSecurityLevelOk returns a tuple with the SecurityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetSecurityLevelOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityLevel) {
		return nil, false
	}
	return o.SecurityLevel, true
}

// HasSecurityLevel returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) HasSecurityLevel() bool {
	if o != nil && !IsNil(o.SecurityLevel) {
		return true
	}

	return false
}

// SetSecurityLevel gets a reference to the given string and assigns it to the SecurityLevel field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) SetSecurityLevel(v string) {
	o.SecurityLevel = &v
}

// GetWriteView returns the WriteView field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetWriteView() string {
	if o == nil || IsNil(o.WriteView) {
		var ret string
		return ret
	}
	return *o.WriteView
}

// GetWriteViewOk returns a tuple with the WriteView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) GetWriteViewOk() (*string, bool) {
	if o == nil || IsNil(o.WriteView) {
		return nil, false
	}
	return o.WriteView, true
}

// HasWriteView returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) HasWriteView() bool {
	if o != nil && !IsNil(o.WriteView) {
		return true
	}

	return false
}

// SetWriteView gets a reference to the given string and assigns it to the WriteView field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) SetWriteView(v string) {
	o.WriteView = &v
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.GroupContextMatch) {
		toSerialize["groupContextMatch"] = o.GroupContextMatch
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ReadView) {
		toSerialize["readView"] = o.ReadView
	}
	if !IsNil(o.SecurityLevel) {
		toSerialize["securityLevel"] = o.SecurityLevel
	}
	if !IsNil(o.WriteView) {
		toSerialize["writeView"] = o.WriteView
	}
	return toSerialize, nil
}

type NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner struct {
	value *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner
	isSet bool
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) Get() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner {
	return v.value
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) Set(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner {
	return &NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner{value: val, isSet: true}
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInnerAccessesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


