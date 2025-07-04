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

// checks if the V1DeviceStatusHistoryPost200ResponseMappingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DeviceStatusHistoryPost200ResponseMappingsInner{}

// V1DeviceStatusHistoryPost200ResponseMappingsInner struct for V1DeviceStatusHistoryPost200ResponseMappingsInner
type V1DeviceStatusHistoryPost200ResponseMappingsInner struct {
	DeviceId *int64 `json:"deviceId,omitempty"`
	EnterpriseId *int64 `json:"enterpriseId,omitempty"`
	Event *string `json:"event,omitempty"`
	EventTime *V1AlarmHistoryGet200ResponseHistoryInnerTime `json:"eventTime,omitempty"`
	Ipaddress *string `json:"ipaddress,omitempty"`
	TtIdentity *string `json:"ttIdentity,omitempty"`
}

// NewV1DeviceStatusHistoryPost200ResponseMappingsInner instantiates a new V1DeviceStatusHistoryPost200ResponseMappingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DeviceStatusHistoryPost200ResponseMappingsInner() *V1DeviceStatusHistoryPost200ResponseMappingsInner {
	this := V1DeviceStatusHistoryPost200ResponseMappingsInner{}
	return &this
}

// NewV1DeviceStatusHistoryPost200ResponseMappingsInnerWithDefaults instantiates a new V1DeviceStatusHistoryPost200ResponseMappingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DeviceStatusHistoryPost200ResponseMappingsInnerWithDefaults() *V1DeviceStatusHistoryPost200ResponseMappingsInner {
	this := V1DeviceStatusHistoryPost200ResponseMappingsInner{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetDeviceId() int64 {
	if o == nil || IsNil(o.DeviceId) {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetDeviceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetEnterpriseId() int64 {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret int64
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetEnterpriseIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given int64 and assigns it to the EnterpriseId field.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) SetEnterpriseId(v int64) {
	o.EnterpriseId = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) SetEvent(v string) {
	o.Event = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetEventTime() V1AlarmHistoryGet200ResponseHistoryInnerTime {
	if o == nil || IsNil(o.EventTime) {
		var ret V1AlarmHistoryGet200ResponseHistoryInnerTime
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetEventTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given V1AlarmHistoryGet200ResponseHistoryInnerTime and assigns it to the EventTime field.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) SetEventTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime) {
	o.EventTime = &v
}

// GetIpaddress returns the Ipaddress field value if set, zero value otherwise.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetIpaddress() string {
	if o == nil || IsNil(o.Ipaddress) {
		var ret string
		return ret
	}
	return *o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetIpaddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipaddress) {
		return nil, false
	}
	return o.Ipaddress, true
}

// HasIpaddress returns a boolean if a field has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) HasIpaddress() bool {
	if o != nil && !IsNil(o.Ipaddress) {
		return true
	}

	return false
}

// SetIpaddress gets a reference to the given string and assigns it to the Ipaddress field.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) SetIpaddress(v string) {
	o.Ipaddress = &v
}

// GetTtIdentity returns the TtIdentity field value if set, zero value otherwise.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetTtIdentity() string {
	if o == nil || IsNil(o.TtIdentity) {
		var ret string
		return ret
	}
	return *o.TtIdentity
}

// GetTtIdentityOk returns a tuple with the TtIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) GetTtIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.TtIdentity) {
		return nil, false
	}
	return o.TtIdentity, true
}

// HasTtIdentity returns a boolean if a field has been set.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) HasTtIdentity() bool {
	if o != nil && !IsNil(o.TtIdentity) {
		return true
	}

	return false
}

// SetTtIdentity gets a reference to the given string and assigns it to the TtIdentity field.
func (o *V1DeviceStatusHistoryPost200ResponseMappingsInner) SetTtIdentity(v string) {
	o.TtIdentity = &v
}

func (o V1DeviceStatusHistoryPost200ResponseMappingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DeviceStatusHistoryPost200ResponseMappingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterpriseId"] = o.EnterpriseId
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.EventTime) {
		toSerialize["eventTime"] = o.EventTime
	}
	if !IsNil(o.Ipaddress) {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if !IsNil(o.TtIdentity) {
		toSerialize["ttIdentity"] = o.TtIdentity
	}
	return toSerialize, nil
}

type NullableV1DeviceStatusHistoryPost200ResponseMappingsInner struct {
	value *V1DeviceStatusHistoryPost200ResponseMappingsInner
	isSet bool
}

func (v NullableV1DeviceStatusHistoryPost200ResponseMappingsInner) Get() *V1DeviceStatusHistoryPost200ResponseMappingsInner {
	return v.value
}

func (v *NullableV1DeviceStatusHistoryPost200ResponseMappingsInner) Set(val *V1DeviceStatusHistoryPost200ResponseMappingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DeviceStatusHistoryPost200ResponseMappingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DeviceStatusHistoryPost200ResponseMappingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DeviceStatusHistoryPost200ResponseMappingsInner(val *V1DeviceStatusHistoryPost200ResponseMappingsInner) *NullableV1DeviceStatusHistoryPost200ResponseMappingsInner {
	return &NullableV1DeviceStatusHistoryPost200ResponseMappingsInner{value: val, isSet: true}
}

func (v NullableV1DeviceStatusHistoryPost200ResponseMappingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DeviceStatusHistoryPost200ResponseMappingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


