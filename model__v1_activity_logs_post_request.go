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

// checks if the V1ActivityLogsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ActivityLogsPostRequest{}

// V1ActivityLogsPostRequest struct for V1ActivityLogsPostRequest
type V1ActivityLogsPostRequest struct {
	CursorRef *string `json:"cursorRef,omitempty"`
	NumLogs *int32 `json:"numLogs,omitempty"`
	OldTs *V1AlarmHistoryGet200ResponseHistoryInnerTime `json:"oldTs,omitempty"`
	RecentTs *V1AlarmHistoryGet200ResponseHistoryInnerTime `json:"recentTs,omitempty"`
	Selector *V1ActivityLogsPostRequestSelector `json:"selector,omitempty"`
	SelectorV2 *V1ActivityLogsPostRequestSelectorV2 `json:"selectorV2,omitempty"`
}

// NewV1ActivityLogsPostRequest instantiates a new V1ActivityLogsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ActivityLogsPostRequest() *V1ActivityLogsPostRequest {
	this := V1ActivityLogsPostRequest{}
	return &this
}

// NewV1ActivityLogsPostRequestWithDefaults instantiates a new V1ActivityLogsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ActivityLogsPostRequestWithDefaults() *V1ActivityLogsPostRequest {
	this := V1ActivityLogsPostRequest{}
	return &this
}

// GetCursorRef returns the CursorRef field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequest) GetCursorRef() string {
	if o == nil || IsNil(o.CursorRef) {
		var ret string
		return ret
	}
	return *o.CursorRef
}

// GetCursorRefOk returns a tuple with the CursorRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequest) GetCursorRefOk() (*string, bool) {
	if o == nil || IsNil(o.CursorRef) {
		return nil, false
	}
	return o.CursorRef, true
}

// HasCursorRef returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequest) HasCursorRef() bool {
	if o != nil && !IsNil(o.CursorRef) {
		return true
	}

	return false
}

// SetCursorRef gets a reference to the given string and assigns it to the CursorRef field.
func (o *V1ActivityLogsPostRequest) SetCursorRef(v string) {
	o.CursorRef = &v
}

// GetNumLogs returns the NumLogs field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequest) GetNumLogs() int32 {
	if o == nil || IsNil(o.NumLogs) {
		var ret int32
		return ret
	}
	return *o.NumLogs
}

// GetNumLogsOk returns a tuple with the NumLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequest) GetNumLogsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumLogs) {
		return nil, false
	}
	return o.NumLogs, true
}

// HasNumLogs returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequest) HasNumLogs() bool {
	if o != nil && !IsNil(o.NumLogs) {
		return true
	}

	return false
}

// SetNumLogs gets a reference to the given int32 and assigns it to the NumLogs field.
func (o *V1ActivityLogsPostRequest) SetNumLogs(v int32) {
	o.NumLogs = &v
}

// GetOldTs returns the OldTs field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequest) GetOldTs() V1AlarmHistoryGet200ResponseHistoryInnerTime {
	if o == nil || IsNil(o.OldTs) {
		var ret V1AlarmHistoryGet200ResponseHistoryInnerTime
		return ret
	}
	return *o.OldTs
}

// GetOldTsOk returns a tuple with the OldTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequest) GetOldTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool) {
	if o == nil || IsNil(o.OldTs) {
		return nil, false
	}
	return o.OldTs, true
}

// HasOldTs returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequest) HasOldTs() bool {
	if o != nil && !IsNil(o.OldTs) {
		return true
	}

	return false
}

// SetOldTs gets a reference to the given V1AlarmHistoryGet200ResponseHistoryInnerTime and assigns it to the OldTs field.
func (o *V1ActivityLogsPostRequest) SetOldTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime) {
	o.OldTs = &v
}

// GetRecentTs returns the RecentTs field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequest) GetRecentTs() V1AlarmHistoryGet200ResponseHistoryInnerTime {
	if o == nil || IsNil(o.RecentTs) {
		var ret V1AlarmHistoryGet200ResponseHistoryInnerTime
		return ret
	}
	return *o.RecentTs
}

// GetRecentTsOk returns a tuple with the RecentTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequest) GetRecentTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool) {
	if o == nil || IsNil(o.RecentTs) {
		return nil, false
	}
	return o.RecentTs, true
}

// HasRecentTs returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequest) HasRecentTs() bool {
	if o != nil && !IsNil(o.RecentTs) {
		return true
	}

	return false
}

// SetRecentTs gets a reference to the given V1AlarmHistoryGet200ResponseHistoryInnerTime and assigns it to the RecentTs field.
func (o *V1ActivityLogsPostRequest) SetRecentTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime) {
	o.RecentTs = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequest) GetSelector() V1ActivityLogsPostRequestSelector {
	if o == nil || IsNil(o.Selector) {
		var ret V1ActivityLogsPostRequestSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequest) GetSelectorOk() (*V1ActivityLogsPostRequestSelector, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequest) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given V1ActivityLogsPostRequestSelector and assigns it to the Selector field.
func (o *V1ActivityLogsPostRequest) SetSelector(v V1ActivityLogsPostRequestSelector) {
	o.Selector = &v
}

// GetSelectorV2 returns the SelectorV2 field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequest) GetSelectorV2() V1ActivityLogsPostRequestSelectorV2 {
	if o == nil || IsNil(o.SelectorV2) {
		var ret V1ActivityLogsPostRequestSelectorV2
		return ret
	}
	return *o.SelectorV2
}

// GetSelectorV2Ok returns a tuple with the SelectorV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequest) GetSelectorV2Ok() (*V1ActivityLogsPostRequestSelectorV2, bool) {
	if o == nil || IsNil(o.SelectorV2) {
		return nil, false
	}
	return o.SelectorV2, true
}

// HasSelectorV2 returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequest) HasSelectorV2() bool {
	if o != nil && !IsNil(o.SelectorV2) {
		return true
	}

	return false
}

// SetSelectorV2 gets a reference to the given V1ActivityLogsPostRequestSelectorV2 and assigns it to the SelectorV2 field.
func (o *V1ActivityLogsPostRequest) SetSelectorV2(v V1ActivityLogsPostRequestSelectorV2) {
	o.SelectorV2 = &v
}

func (o V1ActivityLogsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ActivityLogsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CursorRef) {
		toSerialize["cursorRef"] = o.CursorRef
	}
	if !IsNil(o.NumLogs) {
		toSerialize["numLogs"] = o.NumLogs
	}
	if !IsNil(o.OldTs) {
		toSerialize["oldTs"] = o.OldTs
	}
	if !IsNil(o.RecentTs) {
		toSerialize["recentTs"] = o.RecentTs
	}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !IsNil(o.SelectorV2) {
		toSerialize["selectorV2"] = o.SelectorV2
	}
	return toSerialize, nil
}

type NullableV1ActivityLogsPostRequest struct {
	value *V1ActivityLogsPostRequest
	isSet bool
}

func (v NullableV1ActivityLogsPostRequest) Get() *V1ActivityLogsPostRequest {
	return v.value
}

func (v *NullableV1ActivityLogsPostRequest) Set(val *V1ActivityLogsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ActivityLogsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ActivityLogsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ActivityLogsPostRequest(val *V1ActivityLogsPostRequest) *NullableV1ActivityLogsPostRequest {
	return &NullableV1ActivityLogsPostRequest{value: val, isSet: true}
}

func (v NullableV1ActivityLogsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ActivityLogsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


