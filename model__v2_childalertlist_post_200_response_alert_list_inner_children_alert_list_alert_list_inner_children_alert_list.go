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

// checks if the V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList{}

// V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList struct for V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList
type V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList struct {
	AlertList []V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertListAlertListInner `json:"alertList,omitempty"`
}

// NewV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList instantiates a new V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList() *V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList {
	this := V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList{}
	return &this
}

// NewV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertListWithDefaults instantiates a new V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertListWithDefaults() *V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList {
	this := V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList{}
	return &this
}

// GetAlertList returns the AlertList field value if set, zero value otherwise.
func (o *V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) GetAlertList() []V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertListAlertListInner {
	if o == nil || IsNil(o.AlertList) {
		var ret []V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertListAlertListInner
		return ret
	}
	return o.AlertList
}

// GetAlertListOk returns a tuple with the AlertList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) GetAlertListOk() ([]V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertListAlertListInner, bool) {
	if o == nil || IsNil(o.AlertList) {
		return nil, false
	}
	return o.AlertList, true
}

// HasAlertList returns a boolean if a field has been set.
func (o *V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) HasAlertList() bool {
	if o != nil && !IsNil(o.AlertList) {
		return true
	}

	return false
}

// SetAlertList gets a reference to the given []V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertListAlertListInner and assigns it to the AlertList field.
func (o *V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) SetAlertList(v []V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertListAlertListInner) {
	o.AlertList = v
}

func (o V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertList) {
		toSerialize["alertList"] = o.AlertList
	}
	return toSerialize, nil
}

type NullableV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList struct {
	value *V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList
	isSet bool
}

func (v NullableV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) Get() *V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList {
	return v.value
}

func (v *NullableV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) Set(val *V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList(val *V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) *NullableV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList {
	return &NullableV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList{value: val, isSet: true}
}

func (v NullableV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ChildalertlistPost200ResponseAlertListInnerChildrenAlertListAlertListInnerChildrenAlertList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


