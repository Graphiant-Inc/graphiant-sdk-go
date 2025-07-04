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

// checks if the V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue{}

// V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue struct for V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue
type V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue struct {
	SysdbMonitor *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValueSysdbMonitor `json:"sysdbMonitor,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue instantiates a new V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue() *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue {
	this := V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValueWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValueWithDefaults() *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue {
	this := V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue{}
	return &this
}

// GetSysdbMonitor returns the SysdbMonitor field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) GetSysdbMonitor() V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValueSysdbMonitor {
	if o == nil || IsNil(o.SysdbMonitor) {
		var ret V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValueSysdbMonitor
		return ret
	}
	return *o.SysdbMonitor
}

// GetSysdbMonitorOk returns a tuple with the SysdbMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) GetSysdbMonitorOk() (*V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValueSysdbMonitor, bool) {
	if o == nil || IsNil(o.SysdbMonitor) {
		return nil, false
	}
	return o.SysdbMonitor, true
}

// HasSysdbMonitor returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) HasSysdbMonitor() bool {
	if o != nil && !IsNil(o.SysdbMonitor) {
		return true
	}

	return false
}

// SetSysdbMonitor gets a reference to the given V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValueSysdbMonitor and assigns it to the SysdbMonitor field.
func (o *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) SetSysdbMonitor(v V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValueSysdbMonitor) {
	o.SysdbMonitor = &v
}

func (o V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SysdbMonitor) {
		toSerialize["sysdbMonitor"] = o.SysdbMonitor
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue struct {
	value *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) Get() *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) Set(val *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue(val *V1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) *NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue {
	return &NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCorePrometheusSysdbMonitorsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


