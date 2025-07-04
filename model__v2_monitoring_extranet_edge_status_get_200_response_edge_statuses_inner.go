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

// checks if the V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner{}

// V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner struct for V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner
type V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner struct {
	CreatedAt *V1AlarmHistoryGet200ResponseHistoryInnerTime `json:"createdAt,omitempty"`
	DisconnectedReason *string `json:"disconnectedReason,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Id *int64 `json:"id,omitempty"`
	SiteName *string `json:"siteName,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner instantiates a new V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner() *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner {
	this := V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner{}
	return &this
}

// NewV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInnerWithDefaults instantiates a new V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInnerWithDefaults() *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner {
	this := V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetCreatedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret V1AlarmHistoryGet200ResponseHistoryInnerTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetCreatedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given V1AlarmHistoryGet200ResponseHistoryInnerTime and assigns it to the CreatedAt field.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetCreatedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime) {
	o.CreatedAt = &v
}

// GetDisconnectedReason returns the DisconnectedReason field value if set, zero value otherwise.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetDisconnectedReason() string {
	if o == nil || IsNil(o.DisconnectedReason) {
		var ret string
		return ret
	}
	return *o.DisconnectedReason
}

// GetDisconnectedReasonOk returns a tuple with the DisconnectedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetDisconnectedReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DisconnectedReason) {
		return nil, false
	}
	return o.DisconnectedReason, true
}

// HasDisconnectedReason returns a boolean if a field has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasDisconnectedReason() bool {
	if o != nil && !IsNil(o.DisconnectedReason) {
		return true
	}

	return false
}

// SetDisconnectedReason gets a reference to the given string and assigns it to the DisconnectedReason field.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetDisconnectedReason(v string) {
	o.DisconnectedReason = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetHostname(v string) {
	o.Hostname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetId(v int64) {
	o.Id = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetSiteName(v string) {
	o.SiteName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetStatus(v string) {
	o.Status = &v
}

func (o V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DisconnectedReason) {
		toSerialize["disconnectedReason"] = o.DisconnectedReason
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SiteName) {
		toSerialize["siteName"] = o.SiteName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner struct {
	value *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner
	isSet bool
}

func (v NullableV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) Get() *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner {
	return v.value
}

func (v *NullableV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) Set(val *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner(val *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) *NullableV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner {
	return &NullableV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner{value: val, isSet: true}
}

func (v NullableV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


