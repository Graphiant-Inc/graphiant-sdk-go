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

// checks if the V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner{}

// V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner struct for V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner
type V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner struct {
	EgressInterface *string `json:"egressInterface,omitempty"`
	LastModified *V1AlarmHistoryGet200ResponseHistoryInnerTime `json:"lastModified,omitempty"`
	Metric *int32 `json:"metric,omitempty"`
	NextHop *string `json:"nextHop,omitempty"`
	Tag *int32 `json:"tag,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner instantiates a new V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner() *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner {
	this := V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner{}
	return &this
}

// NewV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInnerWithDefaults instantiates a new V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInnerWithDefaults() *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner {
	this := V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner{}
	return &this
}

// GetEgressInterface returns the EgressInterface field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetEgressInterface() string {
	if o == nil || IsNil(o.EgressInterface) {
		var ret string
		return ret
	}
	return *o.EgressInterface
}

// GetEgressInterfaceOk returns a tuple with the EgressInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetEgressInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.EgressInterface) {
		return nil, false
	}
	return o.EgressInterface, true
}

// HasEgressInterface returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) HasEgressInterface() bool {
	if o != nil && !IsNil(o.EgressInterface) {
		return true
	}

	return false
}

// SetEgressInterface gets a reference to the given string and assigns it to the EgressInterface field.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) SetEgressInterface(v string) {
	o.EgressInterface = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetLastModified() V1AlarmHistoryGet200ResponseHistoryInnerTime {
	if o == nil || IsNil(o.LastModified) {
		var ret V1AlarmHistoryGet200ResponseHistoryInnerTime
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetLastModifiedOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given V1AlarmHistoryGet200ResponseHistoryInnerTime and assigns it to the LastModified field.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) SetLastModified(v V1AlarmHistoryGet200ResponseHistoryInnerTime) {
	o.LastModified = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetMetric() int32 {
	if o == nil || IsNil(o.Metric) {
		var ret int32
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetMetricOk() (*int32, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given int32 and assigns it to the Metric field.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) SetMetric(v int32) {
	o.Metric = &v
}

// GetNextHop returns the NextHop field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetNextHop() string {
	if o == nil || IsNil(o.NextHop) {
		var ret string
		return ret
	}
	return *o.NextHop
}

// GetNextHopOk returns a tuple with the NextHop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetNextHopOk() (*string, bool) {
	if o == nil || IsNil(o.NextHop) {
		return nil, false
	}
	return o.NextHop, true
}

// HasNextHop returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) HasNextHop() bool {
	if o != nil && !IsNil(o.NextHop) {
		return true
	}

	return false
}

// SetNextHop gets a reference to the given string and assigns it to the NextHop field.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) SetNextHop(v string) {
	o.NextHop = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetTag() int32 {
	if o == nil || IsNil(o.Tag) {
		var ret int32
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetTagOk() (*int32, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given int32 and assigns it to the Tag field.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) SetTag(v int32) {
	o.Tag = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) SetType(v string) {
	o.Type = &v
}

func (o V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EgressInterface) {
		toSerialize["egressInterface"] = o.EgressInterface
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.NextHop) {
		toSerialize["nextHop"] = o.NextHop
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner struct {
	value *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner
	isSet bool
}

func (v NullableV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) Get() *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner {
	return v.value
}

func (v *NullableV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) Set(val *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner(val *V1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) *NullableV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner {
	return &NullableV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner{value: val, isSet: true}
}

func (v NullableV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DeviceRoutingOspfv2RibGet200ResponseRoutesInnerPathInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


