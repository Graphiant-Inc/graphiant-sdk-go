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

// checks if the V1DiagnosticInterfaceResetDeviceIdPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DiagnosticInterfaceResetDeviceIdPutRequest{}

// V1DiagnosticInterfaceResetDeviceIdPutRequest struct for V1DiagnosticInterfaceResetDeviceIdPutRequest
type V1DiagnosticInterfaceResetDeviceIdPutRequest struct {
	Interface *string `json:"interface,omitempty"`
}

// NewV1DiagnosticInterfaceResetDeviceIdPutRequest instantiates a new V1DiagnosticInterfaceResetDeviceIdPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DiagnosticInterfaceResetDeviceIdPutRequest() *V1DiagnosticInterfaceResetDeviceIdPutRequest {
	this := V1DiagnosticInterfaceResetDeviceIdPutRequest{}
	return &this
}

// NewV1DiagnosticInterfaceResetDeviceIdPutRequestWithDefaults instantiates a new V1DiagnosticInterfaceResetDeviceIdPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DiagnosticInterfaceResetDeviceIdPutRequestWithDefaults() *V1DiagnosticInterfaceResetDeviceIdPutRequest {
	this := V1DiagnosticInterfaceResetDeviceIdPutRequest{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *V1DiagnosticInterfaceResetDeviceIdPutRequest) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DiagnosticInterfaceResetDeviceIdPutRequest) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *V1DiagnosticInterfaceResetDeviceIdPutRequest) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *V1DiagnosticInterfaceResetDeviceIdPutRequest) SetInterface(v string) {
	o.Interface = &v
}

func (o V1DiagnosticInterfaceResetDeviceIdPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DiagnosticInterfaceResetDeviceIdPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	return toSerialize, nil
}

type NullableV1DiagnosticInterfaceResetDeviceIdPutRequest struct {
	value *V1DiagnosticInterfaceResetDeviceIdPutRequest
	isSet bool
}

func (v NullableV1DiagnosticInterfaceResetDeviceIdPutRequest) Get() *V1DiagnosticInterfaceResetDeviceIdPutRequest {
	return v.value
}

func (v *NullableV1DiagnosticInterfaceResetDeviceIdPutRequest) Set(val *V1DiagnosticInterfaceResetDeviceIdPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DiagnosticInterfaceResetDeviceIdPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DiagnosticInterfaceResetDeviceIdPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DiagnosticInterfaceResetDeviceIdPutRequest(val *V1DiagnosticInterfaceResetDeviceIdPutRequest) *NullableV1DiagnosticInterfaceResetDeviceIdPutRequest {
	return &NullableV1DiagnosticInterfaceResetDeviceIdPutRequest{value: val, isSet: true}
}

func (v NullableV1DiagnosticInterfaceResetDeviceIdPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DiagnosticInterfaceResetDeviceIdPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


