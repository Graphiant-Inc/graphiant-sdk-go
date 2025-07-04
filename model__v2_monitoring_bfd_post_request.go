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

// checks if the V2MonitoringBfdPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2MonitoringBfdPostRequest{}

// V2MonitoringBfdPostRequest struct for V2MonitoringBfdPostRequest
type V2MonitoringBfdPostRequest struct {
	DeviceId *int64 `json:"deviceId,omitempty"`
	Selectors []V2MonitoringBfdPostRequestSelectorsInner `json:"selectors,omitempty"`
	TimeWindow *V2NotificationlistPostRequestTimeWindow `json:"timeWindow,omitempty"`
}

// NewV2MonitoringBfdPostRequest instantiates a new V2MonitoringBfdPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2MonitoringBfdPostRequest() *V2MonitoringBfdPostRequest {
	this := V2MonitoringBfdPostRequest{}
	return &this
}

// NewV2MonitoringBfdPostRequestWithDefaults instantiates a new V2MonitoringBfdPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2MonitoringBfdPostRequestWithDefaults() *V2MonitoringBfdPostRequest {
	this := V2MonitoringBfdPostRequest{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *V2MonitoringBfdPostRequest) GetDeviceId() int64 {
	if o == nil || IsNil(o.DeviceId) {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringBfdPostRequest) GetDeviceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *V2MonitoringBfdPostRequest) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *V2MonitoringBfdPostRequest) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *V2MonitoringBfdPostRequest) GetSelectors() []V2MonitoringBfdPostRequestSelectorsInner {
	if o == nil || IsNil(o.Selectors) {
		var ret []V2MonitoringBfdPostRequestSelectorsInner
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringBfdPostRequest) GetSelectorsOk() ([]V2MonitoringBfdPostRequestSelectorsInner, bool) {
	if o == nil || IsNil(o.Selectors) {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *V2MonitoringBfdPostRequest) HasSelectors() bool {
	if o != nil && !IsNil(o.Selectors) {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given []V2MonitoringBfdPostRequestSelectorsInner and assigns it to the Selectors field.
func (o *V2MonitoringBfdPostRequest) SetSelectors(v []V2MonitoringBfdPostRequestSelectorsInner) {
	o.Selectors = v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *V2MonitoringBfdPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow {
	if o == nil || IsNil(o.TimeWindow) {
		var ret V2NotificationlistPostRequestTimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringBfdPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool) {
	if o == nil || IsNil(o.TimeWindow) {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *V2MonitoringBfdPostRequest) HasTimeWindow() bool {
	if o != nil && !IsNil(o.TimeWindow) {
		return true
	}

	return false
}

// SetTimeWindow gets a reference to the given V2NotificationlistPostRequestTimeWindow and assigns it to the TimeWindow field.
func (o *V2MonitoringBfdPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow) {
	o.TimeWindow = &v
}

func (o V2MonitoringBfdPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2MonitoringBfdPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.Selectors) {
		toSerialize["selectors"] = o.Selectors
	}
	if !IsNil(o.TimeWindow) {
		toSerialize["timeWindow"] = o.TimeWindow
	}
	return toSerialize, nil
}

type NullableV2MonitoringBfdPostRequest struct {
	value *V2MonitoringBfdPostRequest
	isSet bool
}

func (v NullableV2MonitoringBfdPostRequest) Get() *V2MonitoringBfdPostRequest {
	return v.value
}

func (v *NullableV2MonitoringBfdPostRequest) Set(val *V2MonitoringBfdPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2MonitoringBfdPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2MonitoringBfdPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2MonitoringBfdPostRequest(val *V2MonitoringBfdPostRequest) *NullableV2MonitoringBfdPostRequest {
	return &NullableV2MonitoringBfdPostRequest{value: val, isSet: true}
}

func (v NullableV2MonitoringBfdPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2MonitoringBfdPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


