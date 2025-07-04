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

// checks if the V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process{}

// V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process struct for V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process
type V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process struct {
	AddressFamilies []string `json:"addressFamilies,omitempty"`
	AdminDistance *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance `json:"adminDistance,omitempty"`
	Areas *map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue `json:"areas,omitempty"`
	Auto *bool `json:"auto,omitempty"`
	DefaultOriginate *string `json:"defaultOriginate,omitempty"`
	Manual *string `json:"manual,omitempty"`
	Name *string `json:"name,omitempty"`
	Redistribution *map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue `json:"redistribution,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process {
	this := V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process {
	this := V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process{}
	return &this
}

// GetAddressFamilies returns the AddressFamilies field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAddressFamilies() []string {
	if o == nil || IsNil(o.AddressFamilies) {
		var ret []string
		return ret
	}
	return o.AddressFamilies
}

// GetAddressFamiliesOk returns a tuple with the AddressFamilies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAddressFamiliesOk() ([]string, bool) {
	if o == nil || IsNil(o.AddressFamilies) {
		return nil, false
	}
	return o.AddressFamilies, true
}

// HasAddressFamilies returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasAddressFamilies() bool {
	if o != nil && !IsNil(o.AddressFamilies) {
		return true
	}

	return false
}

// SetAddressFamilies gets a reference to the given []string and assigns it to the AddressFamilies field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetAddressFamilies(v []string) {
	o.AddressFamilies = v
}

// GetAdminDistance returns the AdminDistance field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAdminDistance() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance {
	if o == nil || IsNil(o.AdminDistance) {
		var ret V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance
		return ret
	}
	return *o.AdminDistance
}

// GetAdminDistanceOk returns a tuple with the AdminDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAdminDistanceOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance, bool) {
	if o == nil || IsNil(o.AdminDistance) {
		return nil, false
	}
	return o.AdminDistance, true
}

// HasAdminDistance returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasAdminDistance() bool {
	if o != nil && !IsNil(o.AdminDistance) {
		return true
	}

	return false
}

// SetAdminDistance gets a reference to the given V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance and assigns it to the AdminDistance field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetAdminDistance(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance) {
	o.AdminDistance = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAreas() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue {
	if o == nil || IsNil(o.Areas) {
		var ret map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue
		return ret
	}
	return *o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAreasOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue, bool) {
	if o == nil || IsNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue and assigns it to the Areas field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetAreas(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue) {
	o.Areas = &v
}

// GetAuto returns the Auto field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAuto() bool {
	if o == nil || IsNil(o.Auto) {
		var ret bool
		return ret
	}
	return *o.Auto
}

// GetAutoOk returns a tuple with the Auto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAutoOk() (*bool, bool) {
	if o == nil || IsNil(o.Auto) {
		return nil, false
	}
	return o.Auto, true
}

// HasAuto returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasAuto() bool {
	if o != nil && !IsNil(o.Auto) {
		return true
	}

	return false
}

// SetAuto gets a reference to the given bool and assigns it to the Auto field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetAuto(v bool) {
	o.Auto = &v
}

// GetDefaultOriginate returns the DefaultOriginate field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetDefaultOriginate() string {
	if o == nil || IsNil(o.DefaultOriginate) {
		var ret string
		return ret
	}
	return *o.DefaultOriginate
}

// GetDefaultOriginateOk returns a tuple with the DefaultOriginate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetDefaultOriginateOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultOriginate) {
		return nil, false
	}
	return o.DefaultOriginate, true
}

// HasDefaultOriginate returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasDefaultOriginate() bool {
	if o != nil && !IsNil(o.DefaultOriginate) {
		return true
	}

	return false
}

// SetDefaultOriginate gets a reference to the given string and assigns it to the DefaultOriginate field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetDefaultOriginate(v string) {
	o.DefaultOriginate = &v
}

// GetManual returns the Manual field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetManual() string {
	if o == nil || IsNil(o.Manual) {
		var ret string
		return ret
	}
	return *o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetManualOk() (*string, bool) {
	if o == nil || IsNil(o.Manual) {
		return nil, false
	}
	return o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasManual() bool {
	if o != nil && !IsNil(o.Manual) {
		return true
	}

	return false
}

// SetManual gets a reference to the given string and assigns it to the Manual field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetManual(v string) {
	o.Manual = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetName(v string) {
	o.Name = &v
}

// GetRedistribution returns the Redistribution field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetRedistribution() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue {
	if o == nil || IsNil(o.Redistribution) {
		var ret map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue
		return ret
	}
	return *o.Redistribution
}

// GetRedistributionOk returns a tuple with the Redistribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetRedistributionOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue, bool) {
	if o == nil || IsNil(o.Redistribution) {
		return nil, false
	}
	return o.Redistribution, true
}

// HasRedistribution returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasRedistribution() bool {
	if o != nil && !IsNil(o.Redistribution) {
		return true
	}

	return false
}

// SetRedistribution gets a reference to the given map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue and assigns it to the Redistribution field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetRedistribution(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue) {
	o.Redistribution = &v
}

func (o V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressFamilies) {
		toSerialize["addressFamilies"] = o.AddressFamilies
	}
	if !IsNil(o.AdminDistance) {
		toSerialize["adminDistance"] = o.AdminDistance
	}
	if !IsNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	if !IsNil(o.Auto) {
		toSerialize["auto"] = o.Auto
	}
	if !IsNil(o.DefaultOriginate) {
		toSerialize["defaultOriginate"] = o.DefaultOriginate
	}
	if !IsNil(o.Manual) {
		toSerialize["manual"] = o.Manual
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Redistribution) {
		toSerialize["redistribution"] = o.Redistribution
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process struct {
	value *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) Get() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) Set(val *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process(val *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process {
	return &NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


