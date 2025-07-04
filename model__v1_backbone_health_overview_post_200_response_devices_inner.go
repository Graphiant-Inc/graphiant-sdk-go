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

// checks if the V1BackboneHealthOverviewPost200ResponseDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BackboneHealthOverviewPost200ResponseDevicesInner{}

// V1BackboneHealthOverviewPost200ResponseDevicesInner struct for V1BackboneHealthOverviewPost200ResponseDevicesInner
type V1BackboneHealthOverviewPost200ResponseDevicesInner struct {
	ControlStatus *string `json:"controlStatus,omitempty"`
	DataStatus *string `json:"dataStatus,omitempty"`
	DeviceId *int64 `json:"deviceId,omitempty"`
	DeviceName *string `json:"deviceName,omitempty"`
	DeviceRole *string `json:"deviceRole,omitempty"`
	OverallStatus *string `json:"overallStatus,omitempty"`
	Region *V2AssuranceTopologyInventoryPost200ResponseRegionsInner `json:"region,omitempty"`
	SelectedStatus *string `json:"selectedStatus,omitempty"`
	SystemStatus *string `json:"systemStatus,omitempty"`
}

// NewV1BackboneHealthOverviewPost200ResponseDevicesInner instantiates a new V1BackboneHealthOverviewPost200ResponseDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BackboneHealthOverviewPost200ResponseDevicesInner() *V1BackboneHealthOverviewPost200ResponseDevicesInner {
	this := V1BackboneHealthOverviewPost200ResponseDevicesInner{}
	return &this
}

// NewV1BackboneHealthOverviewPost200ResponseDevicesInnerWithDefaults instantiates a new V1BackboneHealthOverviewPost200ResponseDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BackboneHealthOverviewPost200ResponseDevicesInnerWithDefaults() *V1BackboneHealthOverviewPost200ResponseDevicesInner {
	this := V1BackboneHealthOverviewPost200ResponseDevicesInner{}
	return &this
}

// GetControlStatus returns the ControlStatus field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetControlStatus() string {
	if o == nil || IsNil(o.ControlStatus) {
		var ret string
		return ret
	}
	return *o.ControlStatus
}

// GetControlStatusOk returns a tuple with the ControlStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetControlStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ControlStatus) {
		return nil, false
	}
	return o.ControlStatus, true
}

// HasControlStatus returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasControlStatus() bool {
	if o != nil && !IsNil(o.ControlStatus) {
		return true
	}

	return false
}

// SetControlStatus gets a reference to the given string and assigns it to the ControlStatus field.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetControlStatus(v string) {
	o.ControlStatus = &v
}

// GetDataStatus returns the DataStatus field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDataStatus() string {
	if o == nil || IsNil(o.DataStatus) {
		var ret string
		return ret
	}
	return *o.DataStatus
}

// GetDataStatusOk returns a tuple with the DataStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDataStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DataStatus) {
		return nil, false
	}
	return o.DataStatus, true
}

// HasDataStatus returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasDataStatus() bool {
	if o != nil && !IsNil(o.DataStatus) {
		return true
	}

	return false
}

// SetDataStatus gets a reference to the given string and assigns it to the DataStatus field.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetDataStatus(v string) {
	o.DataStatus = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceId() int64 {
	if o == nil || IsNil(o.DeviceId) {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceRole returns the DeviceRole field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceRole() string {
	if o == nil || IsNil(o.DeviceRole) {
		var ret string
		return ret
	}
	return *o.DeviceRole
}

// GetDeviceRoleOk returns a tuple with the DeviceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceRoleOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceRole) {
		return nil, false
	}
	return o.DeviceRole, true
}

// HasDeviceRole returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasDeviceRole() bool {
	if o != nil && !IsNil(o.DeviceRole) {
		return true
	}

	return false
}

// SetDeviceRole gets a reference to the given string and assigns it to the DeviceRole field.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetDeviceRole(v string) {
	o.DeviceRole = &v
}

// GetOverallStatus returns the OverallStatus field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetOverallStatus() string {
	if o == nil || IsNil(o.OverallStatus) {
		var ret string
		return ret
	}
	return *o.OverallStatus
}

// GetOverallStatusOk returns a tuple with the OverallStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetOverallStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OverallStatus) {
		return nil, false
	}
	return o.OverallStatus, true
}

// HasOverallStatus returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasOverallStatus() bool {
	if o != nil && !IsNil(o.OverallStatus) {
		return true
	}

	return false
}

// SetOverallStatus gets a reference to the given string and assigns it to the OverallStatus field.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetOverallStatus(v string) {
	o.OverallStatus = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetRegion() V2AssuranceTopologyInventoryPost200ResponseRegionsInner {
	if o == nil || IsNil(o.Region) {
		var ret V2AssuranceTopologyInventoryPost200ResponseRegionsInner
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetRegionOk() (*V2AssuranceTopologyInventoryPost200ResponseRegionsInner, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given V2AssuranceTopologyInventoryPost200ResponseRegionsInner and assigns it to the Region field.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetRegion(v V2AssuranceTopologyInventoryPost200ResponseRegionsInner) {
	o.Region = &v
}

// GetSelectedStatus returns the SelectedStatus field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetSelectedStatus() string {
	if o == nil || IsNil(o.SelectedStatus) {
		var ret string
		return ret
	}
	return *o.SelectedStatus
}

// GetSelectedStatusOk returns a tuple with the SelectedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetSelectedStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SelectedStatus) {
		return nil, false
	}
	return o.SelectedStatus, true
}

// HasSelectedStatus returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasSelectedStatus() bool {
	if o != nil && !IsNil(o.SelectedStatus) {
		return true
	}

	return false
}

// SetSelectedStatus gets a reference to the given string and assigns it to the SelectedStatus field.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetSelectedStatus(v string) {
	o.SelectedStatus = &v
}

// GetSystemStatus returns the SystemStatus field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetSystemStatus() string {
	if o == nil || IsNil(o.SystemStatus) {
		var ret string
		return ret
	}
	return *o.SystemStatus
}

// GetSystemStatusOk returns a tuple with the SystemStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetSystemStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SystemStatus) {
		return nil, false
	}
	return o.SystemStatus, true
}

// HasSystemStatus returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasSystemStatus() bool {
	if o != nil && !IsNil(o.SystemStatus) {
		return true
	}

	return false
}

// SetSystemStatus gets a reference to the given string and assigns it to the SystemStatus field.
func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetSystemStatus(v string) {
	o.SystemStatus = &v
}

func (o V1BackboneHealthOverviewPost200ResponseDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BackboneHealthOverviewPost200ResponseDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ControlStatus) {
		toSerialize["controlStatus"] = o.ControlStatus
	}
	if !IsNil(o.DataStatus) {
		toSerialize["dataStatus"] = o.DataStatus
	}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !IsNil(o.DeviceRole) {
		toSerialize["deviceRole"] = o.DeviceRole
	}
	if !IsNil(o.OverallStatus) {
		toSerialize["overallStatus"] = o.OverallStatus
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.SelectedStatus) {
		toSerialize["selectedStatus"] = o.SelectedStatus
	}
	if !IsNil(o.SystemStatus) {
		toSerialize["systemStatus"] = o.SystemStatus
	}
	return toSerialize, nil
}

type NullableV1BackboneHealthOverviewPost200ResponseDevicesInner struct {
	value *V1BackboneHealthOverviewPost200ResponseDevicesInner
	isSet bool
}

func (v NullableV1BackboneHealthOverviewPost200ResponseDevicesInner) Get() *V1BackboneHealthOverviewPost200ResponseDevicesInner {
	return v.value
}

func (v *NullableV1BackboneHealthOverviewPost200ResponseDevicesInner) Set(val *V1BackboneHealthOverviewPost200ResponseDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BackboneHealthOverviewPost200ResponseDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BackboneHealthOverviewPost200ResponseDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BackboneHealthOverviewPost200ResponseDevicesInner(val *V1BackboneHealthOverviewPost200ResponseDevicesInner) *NullableV1BackboneHealthOverviewPost200ResponseDevicesInner {
	return &NullableV1BackboneHealthOverviewPost200ResponseDevicesInner{value: val, isSet: true}
}

func (v NullableV1BackboneHealthOverviewPost200ResponseDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BackboneHealthOverviewPost200ResponseDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


