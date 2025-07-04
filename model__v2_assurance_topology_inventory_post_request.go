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

// checks if the V2AssuranceTopologyInventoryPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2AssuranceTopologyInventoryPostRequest{}

// V2AssuranceTopologyInventoryPostRequest struct for V2AssuranceTopologyInventoryPostRequest
type V2AssuranceTopologyInventoryPostRequest struct {
	BucketId *string `json:"bucketId,omitempty"`
	TimeWindow *V2NotificationlistPostRequestTimeWindow `json:"timeWindow,omitempty"`
	TopologyType *string `json:"topologyType,omitempty"`
}

// NewV2AssuranceTopologyInventoryPostRequest instantiates a new V2AssuranceTopologyInventoryPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssuranceTopologyInventoryPostRequest() *V2AssuranceTopologyInventoryPostRequest {
	this := V2AssuranceTopologyInventoryPostRequest{}
	return &this
}

// NewV2AssuranceTopologyInventoryPostRequestWithDefaults instantiates a new V2AssuranceTopologyInventoryPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssuranceTopologyInventoryPostRequestWithDefaults() *V2AssuranceTopologyInventoryPostRequest {
	this := V2AssuranceTopologyInventoryPostRequest{}
	return &this
}

// GetBucketId returns the BucketId field value if set, zero value otherwise.
func (o *V2AssuranceTopologyInventoryPostRequest) GetBucketId() string {
	if o == nil || IsNil(o.BucketId) {
		var ret string
		return ret
	}
	return *o.BucketId
}

// GetBucketIdOk returns a tuple with the BucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyInventoryPostRequest) GetBucketIdOk() (*string, bool) {
	if o == nil || IsNil(o.BucketId) {
		return nil, false
	}
	return o.BucketId, true
}

// HasBucketId returns a boolean if a field has been set.
func (o *V2AssuranceTopologyInventoryPostRequest) HasBucketId() bool {
	if o != nil && !IsNil(o.BucketId) {
		return true
	}

	return false
}

// SetBucketId gets a reference to the given string and assigns it to the BucketId field.
func (o *V2AssuranceTopologyInventoryPostRequest) SetBucketId(v string) {
	o.BucketId = &v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *V2AssuranceTopologyInventoryPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow {
	if o == nil || IsNil(o.TimeWindow) {
		var ret V2NotificationlistPostRequestTimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyInventoryPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool) {
	if o == nil || IsNil(o.TimeWindow) {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *V2AssuranceTopologyInventoryPostRequest) HasTimeWindow() bool {
	if o != nil && !IsNil(o.TimeWindow) {
		return true
	}

	return false
}

// SetTimeWindow gets a reference to the given V2NotificationlistPostRequestTimeWindow and assigns it to the TimeWindow field.
func (o *V2AssuranceTopologyInventoryPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow) {
	o.TimeWindow = &v
}

// GetTopologyType returns the TopologyType field value if set, zero value otherwise.
func (o *V2AssuranceTopologyInventoryPostRequest) GetTopologyType() string {
	if o == nil || IsNil(o.TopologyType) {
		var ret string
		return ret
	}
	return *o.TopologyType
}

// GetTopologyTypeOk returns a tuple with the TopologyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyInventoryPostRequest) GetTopologyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TopologyType) {
		return nil, false
	}
	return o.TopologyType, true
}

// HasTopologyType returns a boolean if a field has been set.
func (o *V2AssuranceTopologyInventoryPostRequest) HasTopologyType() bool {
	if o != nil && !IsNil(o.TopologyType) {
		return true
	}

	return false
}

// SetTopologyType gets a reference to the given string and assigns it to the TopologyType field.
func (o *V2AssuranceTopologyInventoryPostRequest) SetTopologyType(v string) {
	o.TopologyType = &v
}

func (o V2AssuranceTopologyInventoryPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2AssuranceTopologyInventoryPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BucketId) {
		toSerialize["bucketId"] = o.BucketId
	}
	if !IsNil(o.TimeWindow) {
		toSerialize["timeWindow"] = o.TimeWindow
	}
	if !IsNil(o.TopologyType) {
		toSerialize["topologyType"] = o.TopologyType
	}
	return toSerialize, nil
}

type NullableV2AssuranceTopologyInventoryPostRequest struct {
	value *V2AssuranceTopologyInventoryPostRequest
	isSet bool
}

func (v NullableV2AssuranceTopologyInventoryPostRequest) Get() *V2AssuranceTopologyInventoryPostRequest {
	return v.value
}

func (v *NullableV2AssuranceTopologyInventoryPostRequest) Set(val *V2AssuranceTopologyInventoryPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssuranceTopologyInventoryPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssuranceTopologyInventoryPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssuranceTopologyInventoryPostRequest(val *V2AssuranceTopologyInventoryPostRequest) *NullableV2AssuranceTopologyInventoryPostRequest {
	return &NullableV2AssuranceTopologyInventoryPostRequest{value: val, isSet: true}
}

func (v NullableV2AssuranceTopologyInventoryPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssuranceTopologyInventoryPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


