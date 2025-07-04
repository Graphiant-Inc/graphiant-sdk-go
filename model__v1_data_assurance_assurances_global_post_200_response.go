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

// checks if the V1DataAssuranceAssurancesGlobalPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DataAssuranceAssurancesGlobalPost200Response{}

// V1DataAssuranceAssurancesGlobalPost200Response struct for V1DataAssuranceAssurancesGlobalPost200Response
type V1DataAssuranceAssurancesGlobalPost200Response struct {
	AssuranceId *int64 `json:"assuranceId,omitempty"`
	UnsyncedDeviceNames []string `json:"unsyncedDeviceNames,omitempty"`
}

// NewV1DataAssuranceAssurancesGlobalPost200Response instantiates a new V1DataAssuranceAssurancesGlobalPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DataAssuranceAssurancesGlobalPost200Response() *V1DataAssuranceAssurancesGlobalPost200Response {
	this := V1DataAssuranceAssurancesGlobalPost200Response{}
	return &this
}

// NewV1DataAssuranceAssurancesGlobalPost200ResponseWithDefaults instantiates a new V1DataAssuranceAssurancesGlobalPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DataAssuranceAssurancesGlobalPost200ResponseWithDefaults() *V1DataAssuranceAssurancesGlobalPost200Response {
	this := V1DataAssuranceAssurancesGlobalPost200Response{}
	return &this
}

// GetAssuranceId returns the AssuranceId field value if set, zero value otherwise.
func (o *V1DataAssuranceAssurancesGlobalPost200Response) GetAssuranceId() int64 {
	if o == nil || IsNil(o.AssuranceId) {
		var ret int64
		return ret
	}
	return *o.AssuranceId
}

// GetAssuranceIdOk returns a tuple with the AssuranceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DataAssuranceAssurancesGlobalPost200Response) GetAssuranceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AssuranceId) {
		return nil, false
	}
	return o.AssuranceId, true
}

// HasAssuranceId returns a boolean if a field has been set.
func (o *V1DataAssuranceAssurancesGlobalPost200Response) HasAssuranceId() bool {
	if o != nil && !IsNil(o.AssuranceId) {
		return true
	}

	return false
}

// SetAssuranceId gets a reference to the given int64 and assigns it to the AssuranceId field.
func (o *V1DataAssuranceAssurancesGlobalPost200Response) SetAssuranceId(v int64) {
	o.AssuranceId = &v
}

// GetUnsyncedDeviceNames returns the UnsyncedDeviceNames field value if set, zero value otherwise.
func (o *V1DataAssuranceAssurancesGlobalPost200Response) GetUnsyncedDeviceNames() []string {
	if o == nil || IsNil(o.UnsyncedDeviceNames) {
		var ret []string
		return ret
	}
	return o.UnsyncedDeviceNames
}

// GetUnsyncedDeviceNamesOk returns a tuple with the UnsyncedDeviceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DataAssuranceAssurancesGlobalPost200Response) GetUnsyncedDeviceNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.UnsyncedDeviceNames) {
		return nil, false
	}
	return o.UnsyncedDeviceNames, true
}

// HasUnsyncedDeviceNames returns a boolean if a field has been set.
func (o *V1DataAssuranceAssurancesGlobalPost200Response) HasUnsyncedDeviceNames() bool {
	if o != nil && !IsNil(o.UnsyncedDeviceNames) {
		return true
	}

	return false
}

// SetUnsyncedDeviceNames gets a reference to the given []string and assigns it to the UnsyncedDeviceNames field.
func (o *V1DataAssuranceAssurancesGlobalPost200Response) SetUnsyncedDeviceNames(v []string) {
	o.UnsyncedDeviceNames = v
}

func (o V1DataAssuranceAssurancesGlobalPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DataAssuranceAssurancesGlobalPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssuranceId) {
		toSerialize["assuranceId"] = o.AssuranceId
	}
	if !IsNil(o.UnsyncedDeviceNames) {
		toSerialize["unsyncedDeviceNames"] = o.UnsyncedDeviceNames
	}
	return toSerialize, nil
}

type NullableV1DataAssuranceAssurancesGlobalPost200Response struct {
	value *V1DataAssuranceAssurancesGlobalPost200Response
	isSet bool
}

func (v NullableV1DataAssuranceAssurancesGlobalPost200Response) Get() *V1DataAssuranceAssurancesGlobalPost200Response {
	return v.value
}

func (v *NullableV1DataAssuranceAssurancesGlobalPost200Response) Set(val *V1DataAssuranceAssurancesGlobalPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DataAssuranceAssurancesGlobalPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DataAssuranceAssurancesGlobalPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DataAssuranceAssurancesGlobalPost200Response(val *V1DataAssuranceAssurancesGlobalPost200Response) *NullableV1DataAssuranceAssurancesGlobalPost200Response {
	return &NullableV1DataAssuranceAssurancesGlobalPost200Response{value: val, isSet: true}
}

func (v NullableV1DataAssuranceAssurancesGlobalPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DataAssuranceAssurancesGlobalPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


