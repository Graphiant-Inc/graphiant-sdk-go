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

// checks if the V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue{}

// V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue struct for V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue
type V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue struct {
	AdminEmail *string `json:"adminEmail,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	Counts *V1EnterprisesGet200ResponseEnterprisesInnerCounts `json:"counts,omitempty"`
	EnterpriseId *int64 `json:"enterpriseId,omitempty"`
	ImpersonationEnabled *bool `json:"impersonationEnabled,omitempty"`
}

// NewV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue instantiates a new V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue() *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue {
	this := V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue{}
	return &this
}

// NewV1EnterprisesGet200ResponseEnterprisesInnerCustomersValueWithDefaults instantiates a new V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EnterprisesGet200ResponseEnterprisesInnerCustomersValueWithDefaults() *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue {
	this := V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue{}
	return &this
}

// GetAdminEmail returns the AdminEmail field value if set, zero value otherwise.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) GetAdminEmail() string {
	if o == nil || IsNil(o.AdminEmail) {
		var ret string
		return ret
	}
	return *o.AdminEmail
}

// GetAdminEmailOk returns a tuple with the AdminEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) GetAdminEmailOk() (*string, bool) {
	if o == nil || IsNil(o.AdminEmail) {
		return nil, false
	}
	return o.AdminEmail, true
}

// HasAdminEmail returns a boolean if a field has been set.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) HasAdminEmail() bool {
	if o != nil && !IsNil(o.AdminEmail) {
		return true
	}

	return false
}

// SetAdminEmail gets a reference to the given string and assigns it to the AdminEmail field.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) SetAdminEmail(v string) {
	o.AdminEmail = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) GetCounts() V1EnterprisesGet200ResponseEnterprisesInnerCounts {
	if o == nil || IsNil(o.Counts) {
		var ret V1EnterprisesGet200ResponseEnterprisesInnerCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) GetCountsOk() (*V1EnterprisesGet200ResponseEnterprisesInnerCounts, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given V1EnterprisesGet200ResponseEnterprisesInnerCounts and assigns it to the Counts field.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) SetCounts(v V1EnterprisesGet200ResponseEnterprisesInnerCounts) {
	o.Counts = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) GetEnterpriseId() int64 {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret int64
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) GetEnterpriseIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given int64 and assigns it to the EnterpriseId field.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) SetEnterpriseId(v int64) {
	o.EnterpriseId = &v
}

// GetImpersonationEnabled returns the ImpersonationEnabled field value if set, zero value otherwise.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) GetImpersonationEnabled() bool {
	if o == nil || IsNil(o.ImpersonationEnabled) {
		var ret bool
		return ret
	}
	return *o.ImpersonationEnabled
}

// GetImpersonationEnabledOk returns a tuple with the ImpersonationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) GetImpersonationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ImpersonationEnabled) {
		return nil, false
	}
	return o.ImpersonationEnabled, true
}

// HasImpersonationEnabled returns a boolean if a field has been set.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) HasImpersonationEnabled() bool {
	if o != nil && !IsNil(o.ImpersonationEnabled) {
		return true
	}

	return false
}

// SetImpersonationEnabled gets a reference to the given bool and assigns it to the ImpersonationEnabled field.
func (o *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) SetImpersonationEnabled(v bool) {
	o.ImpersonationEnabled = &v
}

func (o V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminEmail) {
		toSerialize["adminEmail"] = o.AdminEmail
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterpriseId"] = o.EnterpriseId
	}
	if !IsNil(o.ImpersonationEnabled) {
		toSerialize["impersonationEnabled"] = o.ImpersonationEnabled
	}
	return toSerialize, nil
}

type NullableV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue struct {
	value *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue
	isSet bool
}

func (v NullableV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) Get() *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue {
	return v.value
}

func (v *NullableV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) Set(val *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue(val *V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) *NullableV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue {
	return &NullableV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue{value: val, isSet: true}
}

func (v NullableV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EnterprisesGet200ResponseEnterprisesInnerCustomersValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


