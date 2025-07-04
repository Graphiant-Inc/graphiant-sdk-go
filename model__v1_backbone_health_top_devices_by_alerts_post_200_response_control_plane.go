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

// checks if the V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane{}

// V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane struct for V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane
type V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane struct {
	DeviceCounts []V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner `json:"deviceCounts,omitempty"`
}

// NewV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane instantiates a new V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane() *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane {
	this := V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane{}
	return &this
}

// NewV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneWithDefaults instantiates a new V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneWithDefaults() *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane {
	this := V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane{}
	return &this
}

// GetDeviceCounts returns the DeviceCounts field value if set, zero value otherwise.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) GetDeviceCounts() []V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner {
	if o == nil || IsNil(o.DeviceCounts) {
		var ret []V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner
		return ret
	}
	return o.DeviceCounts
}

// GetDeviceCountsOk returns a tuple with the DeviceCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) GetDeviceCountsOk() ([]V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner, bool) {
	if o == nil || IsNil(o.DeviceCounts) {
		return nil, false
	}
	return o.DeviceCounts, true
}

// HasDeviceCounts returns a boolean if a field has been set.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) HasDeviceCounts() bool {
	if o != nil && !IsNil(o.DeviceCounts) {
		return true
	}

	return false
}

// SetDeviceCounts gets a reference to the given []V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner and assigns it to the DeviceCounts field.
func (o *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) SetDeviceCounts(v []V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlaneDeviceCountsInner) {
	o.DeviceCounts = v
}

func (o V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceCounts) {
		toSerialize["deviceCounts"] = o.DeviceCounts
	}
	return toSerialize, nil
}

type NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane struct {
	value *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane
	isSet bool
}

func (v NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) Get() *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane {
	return v.value
}

func (v *NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) Set(val *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane(val *V1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) *NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane {
	return &NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane{value: val, isSet: true}
}

func (v NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BackboneHealthTopDevicesByAlertsPost200ResponseControlPlane) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


