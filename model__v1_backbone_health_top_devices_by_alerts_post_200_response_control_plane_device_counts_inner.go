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

// checks if the V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner{}

// V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner struct for V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner
type V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner struct {
	DeviceName *string `json:"deviceName,omitempty"`
	NumAlerts *int32 `json:"numAlerts,omitempty"`
}

// NewV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner instantiates a new V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner() *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner {
	this := V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner{}
	return &this
}

// NewV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInnerWithDefaults instantiates a new V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInnerWithDefaults() *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner {
	this := V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner{}
	return &this
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetNumAlerts returns the NumAlerts field value if set, zero value otherwise.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) GetNumAlerts() int32 {
	if o == nil || IsNil(o.NumAlerts) {
		var ret int32
		return ret
	}
	return *o.NumAlerts
}

// GetNumAlertsOk returns a tuple with the NumAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) GetNumAlertsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumAlerts) {
		return nil, false
	}
	return o.NumAlerts, true
}

// HasNumAlerts returns a boolean if a field has been set.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) HasNumAlerts() bool {
	if o != nil && !IsNil(o.NumAlerts) {
		return true
	}

	return false
}

// SetNumAlerts gets a reference to the given int32 and assigns it to the NumAlerts field.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) SetNumAlerts(v int32) {
	o.NumAlerts = &v
}

func (o V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !IsNil(o.NumAlerts) {
		toSerialize["numAlerts"] = o.NumAlerts
	}
	return toSerialize, nil
}

type NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner struct {
	value *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner
	isSet bool
}

func (v NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) Get() *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner {
	return v.value
}

func (v *NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) Set(val *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner(val *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) *NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner {
	return &NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner{value: val, isSet: true}
}

func (v NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


