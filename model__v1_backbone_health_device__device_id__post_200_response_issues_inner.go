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

// checks if the V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner{}

// V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner struct for V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner
type V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner struct {
	AlertId *string `json:"alertId,omitempty"`
	AllowListed *bool `json:"allowListed,omitempty"`
	Component *string `json:"component,omitempty"`
	EndTime *V1AlarmHistoryGet200ResponseHistoryInnerTime `json:"endTime,omitempty"`
	Entity *string `json:"entity,omitempty"`
	Issue *string `json:"issue,omitempty"`
	MuteListed *bool `json:"muteListed,omitempty"`
	NotificationCreated *bool `json:"notificationCreated,omitempty"`
	Plane *string `json:"plane,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Severity *string `json:"severity,omitempty"`
	StartTime *V1AlarmHistoryGet200ResponseHistoryInnerTime `json:"startTime,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner() *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner {
	this := V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner{}
	return &this
}

// NewV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInnerWithDefaults instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInnerWithDefaults() *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner {
	this := V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner{}
	return &this
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetAlertId() string {
	if o == nil || IsNil(o.AlertId) {
		var ret string
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetAlertIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlertId) {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasAlertId() bool {
	if o != nil && !IsNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given string and assigns it to the AlertId field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetAlertId(v string) {
	o.AlertId = &v
}

// GetAllowListed returns the AllowListed field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetAllowListed() bool {
	if o == nil || IsNil(o.AllowListed) {
		var ret bool
		return ret
	}
	return *o.AllowListed
}

// GetAllowListedOk returns a tuple with the AllowListed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetAllowListedOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowListed) {
		return nil, false
	}
	return o.AllowListed, true
}

// HasAllowListed returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasAllowListed() bool {
	if o != nil && !IsNil(o.AllowListed) {
		return true
	}

	return false
}

// SetAllowListed gets a reference to the given bool and assigns it to the AllowListed field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetAllowListed(v bool) {
	o.AllowListed = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetComponent(v string) {
	o.Component = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetEndTime() V1AlarmHistoryGet200ResponseHistoryInnerTime {
	if o == nil || IsNil(o.EndTime) {
		var ret V1AlarmHistoryGet200ResponseHistoryInnerTime
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetEndTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given V1AlarmHistoryGet200ResponseHistoryInnerTime and assigns it to the EndTime field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetEndTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime) {
	o.EndTime = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetEntity() string {
	if o == nil || IsNil(o.Entity) {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetEntityOk() (*string, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetEntity(v string) {
	o.Entity = &v
}

// GetIssue returns the Issue field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetIssue() string {
	if o == nil || IsNil(o.Issue) {
		var ret string
		return ret
	}
	return *o.Issue
}

// GetIssueOk returns a tuple with the Issue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetIssueOk() (*string, bool) {
	if o == nil || IsNil(o.Issue) {
		return nil, false
	}
	return o.Issue, true
}

// HasIssue returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasIssue() bool {
	if o != nil && !IsNil(o.Issue) {
		return true
	}

	return false
}

// SetIssue gets a reference to the given string and assigns it to the Issue field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetIssue(v string) {
	o.Issue = &v
}

// GetMuteListed returns the MuteListed field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetMuteListed() bool {
	if o == nil || IsNil(o.MuteListed) {
		var ret bool
		return ret
	}
	return *o.MuteListed
}

// GetMuteListedOk returns a tuple with the MuteListed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetMuteListedOk() (*bool, bool) {
	if o == nil || IsNil(o.MuteListed) {
		return nil, false
	}
	return o.MuteListed, true
}

// HasMuteListed returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasMuteListed() bool {
	if o != nil && !IsNil(o.MuteListed) {
		return true
	}

	return false
}

// SetMuteListed gets a reference to the given bool and assigns it to the MuteListed field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetMuteListed(v bool) {
	o.MuteListed = &v
}

// GetNotificationCreated returns the NotificationCreated field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetNotificationCreated() bool {
	if o == nil || IsNil(o.NotificationCreated) {
		var ret bool
		return ret
	}
	return *o.NotificationCreated
}

// GetNotificationCreatedOk returns a tuple with the NotificationCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetNotificationCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.NotificationCreated) {
		return nil, false
	}
	return o.NotificationCreated, true
}

// HasNotificationCreated returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasNotificationCreated() bool {
	if o != nil && !IsNil(o.NotificationCreated) {
		return true
	}

	return false
}

// SetNotificationCreated gets a reference to the given bool and assigns it to the NotificationCreated field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetNotificationCreated(v bool) {
	o.NotificationCreated = &v
}

// GetPlane returns the Plane field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetPlane() string {
	if o == nil || IsNil(o.Plane) {
		var ret string
		return ret
	}
	return *o.Plane
}

// GetPlaneOk returns a tuple with the Plane field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetPlaneOk() (*string, bool) {
	if o == nil || IsNil(o.Plane) {
		return nil, false
	}
	return o.Plane, true
}

// HasPlane returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasPlane() bool {
	if o != nil && !IsNil(o.Plane) {
		return true
	}

	return false
}

// SetPlane gets a reference to the given string and assigns it to the Plane field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetPlane(v string) {
	o.Plane = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetReason(v string) {
	o.Reason = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetSeverity(v string) {
	o.Severity = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetStartTime() V1AlarmHistoryGet200ResponseHistoryInnerTime {
	if o == nil || IsNil(o.StartTime) {
		var ret V1AlarmHistoryGet200ResponseHistoryInnerTime
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetStartTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given V1AlarmHistoryGet200ResponseHistoryInnerTime and assigns it to the StartTime field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetStartTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) SetStatus(v string) {
	o.Status = &v
}

func (o V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertId) {
		toSerialize["alertId"] = o.AlertId
	}
	if !IsNil(o.AllowListed) {
		toSerialize["allowListed"] = o.AllowListed
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Issue) {
		toSerialize["issue"] = o.Issue
	}
	if !IsNil(o.MuteListed) {
		toSerialize["muteListed"] = o.MuteListed
	}
	if !IsNil(o.NotificationCreated) {
		toSerialize["notificationCreated"] = o.NotificationCreated
	}
	if !IsNil(o.Plane) {
		toSerialize["plane"] = o.Plane
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner struct {
	value *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner
	isSet bool
}

func (v NullableV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) Get() *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner {
	return v.value
}

func (v *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) Set(val *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner(val *V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner {
	return &NullableV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner{value: val, isSet: true}
}

func (v NullableV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


