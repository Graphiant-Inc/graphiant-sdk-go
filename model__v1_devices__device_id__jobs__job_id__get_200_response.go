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

// checks if the V1DevicesDeviceIdJobsJobIdGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdJobsJobIdGet200Response{}

// V1DevicesDeviceIdJobsJobIdGet200Response struct for V1DevicesDeviceIdJobsJobIdGet200Response
type V1DevicesDeviceIdJobsJobIdGet200Response struct {
	DeviceId *int64 `json:"deviceId,omitempty"`
	JobStatus *V1DevicesDeviceIdJobsJobIdGet200ResponseJobStatus `json:"jobStatus,omitempty"`
}

// NewV1DevicesDeviceIdJobsJobIdGet200Response instantiates a new V1DevicesDeviceIdJobsJobIdGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdJobsJobIdGet200Response() *V1DevicesDeviceIdJobsJobIdGet200Response {
	this := V1DevicesDeviceIdJobsJobIdGet200Response{}
	return &this
}

// NewV1DevicesDeviceIdJobsJobIdGet200ResponseWithDefaults instantiates a new V1DevicesDeviceIdJobsJobIdGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdJobsJobIdGet200ResponseWithDefaults() *V1DevicesDeviceIdJobsJobIdGet200Response {
	this := V1DevicesDeviceIdJobsJobIdGet200Response{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdJobsJobIdGet200Response) GetDeviceId() int64 {
	if o == nil || IsNil(o.DeviceId) {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdJobsJobIdGet200Response) GetDeviceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdJobsJobIdGet200Response) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *V1DevicesDeviceIdJobsJobIdGet200Response) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetJobStatus returns the JobStatus field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdJobsJobIdGet200Response) GetJobStatus() V1DevicesDeviceIdJobsJobIdGet200ResponseJobStatus {
	if o == nil || IsNil(o.JobStatus) {
		var ret V1DevicesDeviceIdJobsJobIdGet200ResponseJobStatus
		return ret
	}
	return *o.JobStatus
}

// GetJobStatusOk returns a tuple with the JobStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdJobsJobIdGet200Response) GetJobStatusOk() (*V1DevicesDeviceIdJobsJobIdGet200ResponseJobStatus, bool) {
	if o == nil || IsNil(o.JobStatus) {
		return nil, false
	}
	return o.JobStatus, true
}

// HasJobStatus returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdJobsJobIdGet200Response) HasJobStatus() bool {
	if o != nil && !IsNil(o.JobStatus) {
		return true
	}

	return false
}

// SetJobStatus gets a reference to the given V1DevicesDeviceIdJobsJobIdGet200ResponseJobStatus and assigns it to the JobStatus field.
func (o *V1DevicesDeviceIdJobsJobIdGet200Response) SetJobStatus(v V1DevicesDeviceIdJobsJobIdGet200ResponseJobStatus) {
	o.JobStatus = &v
}

func (o V1DevicesDeviceIdJobsJobIdGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdJobsJobIdGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.JobStatus) {
		toSerialize["jobStatus"] = o.JobStatus
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdJobsJobIdGet200Response struct {
	value *V1DevicesDeviceIdJobsJobIdGet200Response
	isSet bool
}

func (v NullableV1DevicesDeviceIdJobsJobIdGet200Response) Get() *V1DevicesDeviceIdJobsJobIdGet200Response {
	return v.value
}

func (v *NullableV1DevicesDeviceIdJobsJobIdGet200Response) Set(val *V1DevicesDeviceIdJobsJobIdGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdJobsJobIdGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdJobsJobIdGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdJobsJobIdGet200Response(val *V1DevicesDeviceIdJobsJobIdGet200Response) *NullableV1DevicesDeviceIdJobsJobIdGet200Response {
	return &NullableV1DevicesDeviceIdJobsJobIdGet200Response{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdJobsJobIdGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdJobsJobIdGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


