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

// checks if the V2MonitoringQueuePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2MonitoringQueuePostRequest{}

// V2MonitoringQueuePostRequest struct for V2MonitoringQueuePostRequest
type V2MonitoringQueuePostRequest struct {
	DeviceId *int64 `json:"deviceId,omitempty"`
	Selectors []V2MonitoringQueuePostRequestSelectorsInner `json:"selectors,omitempty"`
	TimeWindow *V2NotificationlistPostRequestTimeWindow `json:"timeWindow,omitempty"`
}

// NewV2MonitoringQueuePostRequest instantiates a new V2MonitoringQueuePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2MonitoringQueuePostRequest() *V2MonitoringQueuePostRequest {
	this := V2MonitoringQueuePostRequest{}
	return &this
}

// NewV2MonitoringQueuePostRequestWithDefaults instantiates a new V2MonitoringQueuePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2MonitoringQueuePostRequestWithDefaults() *V2MonitoringQueuePostRequest {
	this := V2MonitoringQueuePostRequest{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *V2MonitoringQueuePostRequest) GetDeviceId() int64 {
	if o == nil || IsNil(o.DeviceId) {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringQueuePostRequest) GetDeviceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *V2MonitoringQueuePostRequest) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *V2MonitoringQueuePostRequest) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *V2MonitoringQueuePostRequest) GetSelectors() []V2MonitoringQueuePostRequestSelectorsInner {
	if o == nil || IsNil(o.Selectors) {
		var ret []V2MonitoringQueuePostRequestSelectorsInner
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringQueuePostRequest) GetSelectorsOk() ([]V2MonitoringQueuePostRequestSelectorsInner, bool) {
	if o == nil || IsNil(o.Selectors) {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *V2MonitoringQueuePostRequest) HasSelectors() bool {
	if o != nil && !IsNil(o.Selectors) {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given []V2MonitoringQueuePostRequestSelectorsInner and assigns it to the Selectors field.
func (o *V2MonitoringQueuePostRequest) SetSelectors(v []V2MonitoringQueuePostRequestSelectorsInner) {
	o.Selectors = v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *V2MonitoringQueuePostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow {
	if o == nil || IsNil(o.TimeWindow) {
		var ret V2NotificationlistPostRequestTimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringQueuePostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool) {
	if o == nil || IsNil(o.TimeWindow) {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *V2MonitoringQueuePostRequest) HasTimeWindow() bool {
	if o != nil && !IsNil(o.TimeWindow) {
		return true
	}

	return false
}

// SetTimeWindow gets a reference to the given V2NotificationlistPostRequestTimeWindow and assigns it to the TimeWindow field.
func (o *V2MonitoringQueuePostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow) {
	o.TimeWindow = &v
}

func (o V2MonitoringQueuePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2MonitoringQueuePostRequest) ToMap() (map[string]interface{}, error) {
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

type NullableV2MonitoringQueuePostRequest struct {
	value *V2MonitoringQueuePostRequest
	isSet bool
}

func (v NullableV2MonitoringQueuePostRequest) Get() *V2MonitoringQueuePostRequest {
	return v.value
}

func (v *NullableV2MonitoringQueuePostRequest) Set(val *V2MonitoringQueuePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2MonitoringQueuePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2MonitoringQueuePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2MonitoringQueuePostRequest(val *V2MonitoringQueuePostRequest) *NullableV2MonitoringQueuePostRequest {
	return &NullableV2MonitoringQueuePostRequest{value: val, isSet: true}
}

func (v NullableV2MonitoringQueuePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2MonitoringQueuePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


