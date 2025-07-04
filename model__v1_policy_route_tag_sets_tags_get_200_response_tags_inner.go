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

// checks if the V1PolicyRouteTagSetsTagsGet200ResponseTagsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PolicyRouteTagSetsTagsGet200ResponseTagsInner{}

// V1PolicyRouteTagSetsTagsGet200ResponseTagsInner struct for V1PolicyRouteTagSetsTagsGet200ResponseTagsInner
type V1PolicyRouteTagSetsTagsGet200ResponseTagsInner struct {
	Id *int64 `json:"id,omitempty"`
	NextSet []V1PolicyRouteTagSetsTagsGet200ResponseTagsInnerNextSetInner `json:"nextSet,omitempty"`
	Tag *string `json:"tag,omitempty"`
	TagValue *int32 `json:"tagValue,omitempty"`
}

// NewV1PolicyRouteTagSetsTagsGet200ResponseTagsInner instantiates a new V1PolicyRouteTagSetsTagsGet200ResponseTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PolicyRouteTagSetsTagsGet200ResponseTagsInner() *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner {
	this := V1PolicyRouteTagSetsTagsGet200ResponseTagsInner{}
	return &this
}

// NewV1PolicyRouteTagSetsTagsGet200ResponseTagsInnerWithDefaults instantiates a new V1PolicyRouteTagSetsTagsGet200ResponseTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PolicyRouteTagSetsTagsGet200ResponseTagsInnerWithDefaults() *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner {
	this := V1PolicyRouteTagSetsTagsGet200ResponseTagsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) SetId(v int64) {
	o.Id = &v
}

// GetNextSet returns the NextSet field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) GetNextSet() []V1PolicyRouteTagSetsTagsGet200ResponseTagsInnerNextSetInner {
	if o == nil || IsNil(o.NextSet) {
		var ret []V1PolicyRouteTagSetsTagsGet200ResponseTagsInnerNextSetInner
		return ret
	}
	return o.NextSet
}

// GetNextSetOk returns a tuple with the NextSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) GetNextSetOk() ([]V1PolicyRouteTagSetsTagsGet200ResponseTagsInnerNextSetInner, bool) {
	if o == nil || IsNil(o.NextSet) {
		return nil, false
	}
	return o.NextSet, true
}

// HasNextSet returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) HasNextSet() bool {
	if o != nil && !IsNil(o.NextSet) {
		return true
	}

	return false
}

// SetNextSet gets a reference to the given []V1PolicyRouteTagSetsTagsGet200ResponseTagsInnerNextSetInner and assigns it to the NextSet field.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) SetNextSet(v []V1PolicyRouteTagSetsTagsGet200ResponseTagsInnerNextSetInner) {
	o.NextSet = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) SetTag(v string) {
	o.Tag = &v
}

// GetTagValue returns the TagValue field value if set, zero value otherwise.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) GetTagValue() int32 {
	if o == nil || IsNil(o.TagValue) {
		var ret int32
		return ret
	}
	return *o.TagValue
}

// GetTagValueOk returns a tuple with the TagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) GetTagValueOk() (*int32, bool) {
	if o == nil || IsNil(o.TagValue) {
		return nil, false
	}
	return o.TagValue, true
}

// HasTagValue returns a boolean if a field has been set.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) HasTagValue() bool {
	if o != nil && !IsNil(o.TagValue) {
		return true
	}

	return false
}

// SetTagValue gets a reference to the given int32 and assigns it to the TagValue field.
func (o *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) SetTagValue(v int32) {
	o.TagValue = &v
}

func (o V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NextSet) {
		toSerialize["nextSet"] = o.NextSet
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TagValue) {
		toSerialize["tagValue"] = o.TagValue
	}
	return toSerialize, nil
}

type NullableV1PolicyRouteTagSetsTagsGet200ResponseTagsInner struct {
	value *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner
	isSet bool
}

func (v NullableV1PolicyRouteTagSetsTagsGet200ResponseTagsInner) Get() *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner {
	return v.value
}

func (v *NullableV1PolicyRouteTagSetsTagsGet200ResponseTagsInner) Set(val *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PolicyRouteTagSetsTagsGet200ResponseTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PolicyRouteTagSetsTagsGet200ResponseTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PolicyRouteTagSetsTagsGet200ResponseTagsInner(val *V1PolicyRouteTagSetsTagsGet200ResponseTagsInner) *NullableV1PolicyRouteTagSetsTagsGet200ResponseTagsInner {
	return &NullableV1PolicyRouteTagSetsTagsGet200ResponseTagsInner{value: val, isSet: true}
}

func (v NullableV1PolicyRouteTagSetsTagsGet200ResponseTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PolicyRouteTagSetsTagsGet200ResponseTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


