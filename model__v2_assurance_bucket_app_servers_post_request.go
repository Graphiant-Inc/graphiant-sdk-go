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

// checks if the V2AssuranceBucketAppServersPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2AssuranceBucketAppServersPostRequest{}

// V2AssuranceBucketAppServersPostRequest struct for V2AssuranceBucketAppServersPostRequest
type V2AssuranceBucketAppServersPostRequest struct {
	AppName *string `json:"appName,omitempty"`
	BucketId *string `json:"bucketId,omitempty"`
	FlexAlgoId *int64 `json:"flexAlgoId,omitempty"`
	TimeWindow *V2NotificationlistPostRequestTimeWindow `json:"timeWindow,omitempty"`
}

// NewV2AssuranceBucketAppServersPostRequest instantiates a new V2AssuranceBucketAppServersPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssuranceBucketAppServersPostRequest() *V2AssuranceBucketAppServersPostRequest {
	this := V2AssuranceBucketAppServersPostRequest{}
	return &this
}

// NewV2AssuranceBucketAppServersPostRequestWithDefaults instantiates a new V2AssuranceBucketAppServersPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssuranceBucketAppServersPostRequestWithDefaults() *V2AssuranceBucketAppServersPostRequest {
	this := V2AssuranceBucketAppServersPostRequest{}
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *V2AssuranceBucketAppServersPostRequest) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceBucketAppServersPostRequest) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *V2AssuranceBucketAppServersPostRequest) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *V2AssuranceBucketAppServersPostRequest) SetAppName(v string) {
	o.AppName = &v
}

// GetBucketId returns the BucketId field value if set, zero value otherwise.
func (o *V2AssuranceBucketAppServersPostRequest) GetBucketId() string {
	if o == nil || IsNil(o.BucketId) {
		var ret string
		return ret
	}
	return *o.BucketId
}

// GetBucketIdOk returns a tuple with the BucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceBucketAppServersPostRequest) GetBucketIdOk() (*string, bool) {
	if o == nil || IsNil(o.BucketId) {
		return nil, false
	}
	return o.BucketId, true
}

// HasBucketId returns a boolean if a field has been set.
func (o *V2AssuranceBucketAppServersPostRequest) HasBucketId() bool {
	if o != nil && !IsNil(o.BucketId) {
		return true
	}

	return false
}

// SetBucketId gets a reference to the given string and assigns it to the BucketId field.
func (o *V2AssuranceBucketAppServersPostRequest) SetBucketId(v string) {
	o.BucketId = &v
}

// GetFlexAlgoId returns the FlexAlgoId field value if set, zero value otherwise.
func (o *V2AssuranceBucketAppServersPostRequest) GetFlexAlgoId() int64 {
	if o == nil || IsNil(o.FlexAlgoId) {
		var ret int64
		return ret
	}
	return *o.FlexAlgoId
}

// GetFlexAlgoIdOk returns a tuple with the FlexAlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceBucketAppServersPostRequest) GetFlexAlgoIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FlexAlgoId) {
		return nil, false
	}
	return o.FlexAlgoId, true
}

// HasFlexAlgoId returns a boolean if a field has been set.
func (o *V2AssuranceBucketAppServersPostRequest) HasFlexAlgoId() bool {
	if o != nil && !IsNil(o.FlexAlgoId) {
		return true
	}

	return false
}

// SetFlexAlgoId gets a reference to the given int64 and assigns it to the FlexAlgoId field.
func (o *V2AssuranceBucketAppServersPostRequest) SetFlexAlgoId(v int64) {
	o.FlexAlgoId = &v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *V2AssuranceBucketAppServersPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow {
	if o == nil || IsNil(o.TimeWindow) {
		var ret V2NotificationlistPostRequestTimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceBucketAppServersPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool) {
	if o == nil || IsNil(o.TimeWindow) {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *V2AssuranceBucketAppServersPostRequest) HasTimeWindow() bool {
	if o != nil && !IsNil(o.TimeWindow) {
		return true
	}

	return false
}

// SetTimeWindow gets a reference to the given V2NotificationlistPostRequestTimeWindow and assigns it to the TimeWindow field.
func (o *V2AssuranceBucketAppServersPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow) {
	o.TimeWindow = &v
}

func (o V2AssuranceBucketAppServersPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2AssuranceBucketAppServersPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppName) {
		toSerialize["appName"] = o.AppName
	}
	if !IsNil(o.BucketId) {
		toSerialize["bucketId"] = o.BucketId
	}
	if !IsNil(o.FlexAlgoId) {
		toSerialize["flexAlgoId"] = o.FlexAlgoId
	}
	if !IsNil(o.TimeWindow) {
		toSerialize["timeWindow"] = o.TimeWindow
	}
	return toSerialize, nil
}

type NullableV2AssuranceBucketAppServersPostRequest struct {
	value *V2AssuranceBucketAppServersPostRequest
	isSet bool
}

func (v NullableV2AssuranceBucketAppServersPostRequest) Get() *V2AssuranceBucketAppServersPostRequest {
	return v.value
}

func (v *NullableV2AssuranceBucketAppServersPostRequest) Set(val *V2AssuranceBucketAppServersPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssuranceBucketAppServersPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssuranceBucketAppServersPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssuranceBucketAppServersPostRequest(val *V2AssuranceBucketAppServersPostRequest) *NullableV2AssuranceBucketAppServersPostRequest {
	return &NullableV2AssuranceBucketAppServersPostRequest{value: val, isSet: true}
}

func (v NullableV2AssuranceBucketAppServersPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssuranceBucketAppServersPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


