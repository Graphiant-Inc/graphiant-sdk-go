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

// checks if the V1LogsPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1LogsPost200Response{}

// V1LogsPost200Response struct for V1LogsPost200Response
type V1LogsPost200Response struct {
	CursorRef *string `json:"cursorRef,omitempty"`
	Histogram []V1AuditLogsPost200ResponseHistogramInner `json:"histogram,omitempty"`
	Logs []V1LogsPost200ResponseLogsInner `json:"logs,omitempty"`
	TotalLogs *int64 `json:"totalLogs,omitempty"`
}

// NewV1LogsPost200Response instantiates a new V1LogsPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1LogsPost200Response() *V1LogsPost200Response {
	this := V1LogsPost200Response{}
	return &this
}

// NewV1LogsPost200ResponseWithDefaults instantiates a new V1LogsPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1LogsPost200ResponseWithDefaults() *V1LogsPost200Response {
	this := V1LogsPost200Response{}
	return &this
}

// GetCursorRef returns the CursorRef field value if set, zero value otherwise.
func (o *V1LogsPost200Response) GetCursorRef() string {
	if o == nil || IsNil(o.CursorRef) {
		var ret string
		return ret
	}
	return *o.CursorRef
}

// GetCursorRefOk returns a tuple with the CursorRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LogsPost200Response) GetCursorRefOk() (*string, bool) {
	if o == nil || IsNil(o.CursorRef) {
		return nil, false
	}
	return o.CursorRef, true
}

// HasCursorRef returns a boolean if a field has been set.
func (o *V1LogsPost200Response) HasCursorRef() bool {
	if o != nil && !IsNil(o.CursorRef) {
		return true
	}

	return false
}

// SetCursorRef gets a reference to the given string and assigns it to the CursorRef field.
func (o *V1LogsPost200Response) SetCursorRef(v string) {
	o.CursorRef = &v
}

// GetHistogram returns the Histogram field value if set, zero value otherwise.
func (o *V1LogsPost200Response) GetHistogram() []V1AuditLogsPost200ResponseHistogramInner {
	if o == nil || IsNil(o.Histogram) {
		var ret []V1AuditLogsPost200ResponseHistogramInner
		return ret
	}
	return o.Histogram
}

// GetHistogramOk returns a tuple with the Histogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LogsPost200Response) GetHistogramOk() ([]V1AuditLogsPost200ResponseHistogramInner, bool) {
	if o == nil || IsNil(o.Histogram) {
		return nil, false
	}
	return o.Histogram, true
}

// HasHistogram returns a boolean if a field has been set.
func (o *V1LogsPost200Response) HasHistogram() bool {
	if o != nil && !IsNil(o.Histogram) {
		return true
	}

	return false
}

// SetHistogram gets a reference to the given []V1AuditLogsPost200ResponseHistogramInner and assigns it to the Histogram field.
func (o *V1LogsPost200Response) SetHistogram(v []V1AuditLogsPost200ResponseHistogramInner) {
	o.Histogram = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *V1LogsPost200Response) GetLogs() []V1LogsPost200ResponseLogsInner {
	if o == nil || IsNil(o.Logs) {
		var ret []V1LogsPost200ResponseLogsInner
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LogsPost200Response) GetLogsOk() ([]V1LogsPost200ResponseLogsInner, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *V1LogsPost200Response) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []V1LogsPost200ResponseLogsInner and assigns it to the Logs field.
func (o *V1LogsPost200Response) SetLogs(v []V1LogsPost200ResponseLogsInner) {
	o.Logs = v
}

// GetTotalLogs returns the TotalLogs field value if set, zero value otherwise.
func (o *V1LogsPost200Response) GetTotalLogs() int64 {
	if o == nil || IsNil(o.TotalLogs) {
		var ret int64
		return ret
	}
	return *o.TotalLogs
}

// GetTotalLogsOk returns a tuple with the TotalLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LogsPost200Response) GetTotalLogsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalLogs) {
		return nil, false
	}
	return o.TotalLogs, true
}

// HasTotalLogs returns a boolean if a field has been set.
func (o *V1LogsPost200Response) HasTotalLogs() bool {
	if o != nil && !IsNil(o.TotalLogs) {
		return true
	}

	return false
}

// SetTotalLogs gets a reference to the given int64 and assigns it to the TotalLogs field.
func (o *V1LogsPost200Response) SetTotalLogs(v int64) {
	o.TotalLogs = &v
}

func (o V1LogsPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1LogsPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CursorRef) {
		toSerialize["cursorRef"] = o.CursorRef
	}
	if !IsNil(o.Histogram) {
		toSerialize["histogram"] = o.Histogram
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.TotalLogs) {
		toSerialize["totalLogs"] = o.TotalLogs
	}
	return toSerialize, nil
}

type NullableV1LogsPost200Response struct {
	value *V1LogsPost200Response
	isSet bool
}

func (v NullableV1LogsPost200Response) Get() *V1LogsPost200Response {
	return v.value
}

func (v *NullableV1LogsPost200Response) Set(val *V1LogsPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1LogsPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1LogsPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1LogsPost200Response(val *V1LogsPost200Response) *NullableV1LogsPost200Response {
	return &NullableV1LogsPost200Response{value: val, isSet: true}
}

func (v NullableV1LogsPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1LogsPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


