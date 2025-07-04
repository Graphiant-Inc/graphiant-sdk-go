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

// checks if the V1GroupsIdMembersGet200ResponseUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GroupsIdMembersGet200ResponseUsersInner{}

// V1GroupsIdMembersGet200ResponseUsersInner struct for V1GroupsIdMembersGet200ResponseUsersInner
type V1GroupsIdMembersGet200ResponseUsersInner struct {
	Email *string `json:"email,omitempty"`
	EnterpriseId *int64 `json:"enterpriseId,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	MfaFactor *string `json:"mfaFactor,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	TimeZone *string `json:"timeZone,omitempty"`
	UserId *string `json:"userId,omitempty"`
	Verified *bool `json:"verified,omitempty"`
}

// NewV1GroupsIdMembersGet200ResponseUsersInner instantiates a new V1GroupsIdMembersGet200ResponseUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GroupsIdMembersGet200ResponseUsersInner() *V1GroupsIdMembersGet200ResponseUsersInner {
	this := V1GroupsIdMembersGet200ResponseUsersInner{}
	return &this
}

// NewV1GroupsIdMembersGet200ResponseUsersInnerWithDefaults instantiates a new V1GroupsIdMembersGet200ResponseUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GroupsIdMembersGet200ResponseUsersInnerWithDefaults() *V1GroupsIdMembersGet200ResponseUsersInner {
	this := V1GroupsIdMembersGet200ResponseUsersInner{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) SetEmail(v string) {
	o.Email = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetEnterpriseId() int64 {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret int64
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetEnterpriseIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given int64 and assigns it to the EnterpriseId field.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) SetEnterpriseId(v int64) {
	o.EnterpriseId = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) SetLastName(v string) {
	o.LastName = &v
}

// GetMfaFactor returns the MfaFactor field value if set, zero value otherwise.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetMfaFactor() string {
	if o == nil || IsNil(o.MfaFactor) {
		var ret string
		return ret
	}
	return *o.MfaFactor
}

// GetMfaFactorOk returns a tuple with the MfaFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetMfaFactorOk() (*string, bool) {
	if o == nil || IsNil(o.MfaFactor) {
		return nil, false
	}
	return o.MfaFactor, true
}

// HasMfaFactor returns a boolean if a field has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) HasMfaFactor() bool {
	if o != nil && !IsNil(o.MfaFactor) {
		return true
	}

	return false
}

// SetMfaFactor gets a reference to the given string and assigns it to the MfaFactor field.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) SetMfaFactor(v string) {
	o.MfaFactor = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) SetUserId(v string) {
	o.UserId = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *V1GroupsIdMembersGet200ResponseUsersInner) SetVerified(v bool) {
	o.Verified = &v
}

func (o V1GroupsIdMembersGet200ResponseUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GroupsIdMembersGet200ResponseUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterpriseId"] = o.EnterpriseId
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.MfaFactor) {
		toSerialize["mfaFactor"] = o.MfaFactor
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	return toSerialize, nil
}

type NullableV1GroupsIdMembersGet200ResponseUsersInner struct {
	value *V1GroupsIdMembersGet200ResponseUsersInner
	isSet bool
}

func (v NullableV1GroupsIdMembersGet200ResponseUsersInner) Get() *V1GroupsIdMembersGet200ResponseUsersInner {
	return v.value
}

func (v *NullableV1GroupsIdMembersGet200ResponseUsersInner) Set(val *V1GroupsIdMembersGet200ResponseUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GroupsIdMembersGet200ResponseUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GroupsIdMembersGet200ResponseUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GroupsIdMembersGet200ResponseUsersInner(val *V1GroupsIdMembersGet200ResponseUsersInner) *NullableV1GroupsIdMembersGet200ResponseUsersInner {
	return &NullableV1GroupsIdMembersGet200ResponseUsersInner{value: val, isSet: true}
}

func (v NullableV1GroupsIdMembersGet200ResponseUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GroupsIdMembersGet200ResponseUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


