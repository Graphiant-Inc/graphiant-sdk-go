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

// checks if the V1ActivityLogsPostRequestSelectorV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ActivityLogsPostRequestSelectorV2{}

// V1ActivityLogsPostRequestSelectorV2 struct for V1ActivityLogsPostRequestSelectorV2
type V1ActivityLogsPostRequestSelectorV2 struct {
	DeviceIds []int64 `json:"deviceIds,omitempty"`
	Entities []V1ActivityLogsPostRequestSelectorJobEntity `json:"entities,omitempty"`
	SiteIds []int64 `json:"siteIds,omitempty"`
	Statuses []string `json:"statuses,omitempty"`
	Types []string `json:"types,omitempty"`
	UserIds []string `json:"userIds,omitempty"`
}

// NewV1ActivityLogsPostRequestSelectorV2 instantiates a new V1ActivityLogsPostRequestSelectorV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ActivityLogsPostRequestSelectorV2() *V1ActivityLogsPostRequestSelectorV2 {
	this := V1ActivityLogsPostRequestSelectorV2{}
	return &this
}

// NewV1ActivityLogsPostRequestSelectorV2WithDefaults instantiates a new V1ActivityLogsPostRequestSelectorV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ActivityLogsPostRequestSelectorV2WithDefaults() *V1ActivityLogsPostRequestSelectorV2 {
	this := V1ActivityLogsPostRequestSelectorV2{}
	return &this
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequestSelectorV2) GetDeviceIds() []int64 {
	if o == nil || IsNil(o.DeviceIds) {
		var ret []int64
		return ret
	}
	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) GetDeviceIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.DeviceIds) {
		return nil, false
	}
	return o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) HasDeviceIds() bool {
	if o != nil && !IsNil(o.DeviceIds) {
		return true
	}

	return false
}

// SetDeviceIds gets a reference to the given []int64 and assigns it to the DeviceIds field.
func (o *V1ActivityLogsPostRequestSelectorV2) SetDeviceIds(v []int64) {
	o.DeviceIds = v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequestSelectorV2) GetEntities() []V1ActivityLogsPostRequestSelectorJobEntity {
	if o == nil || IsNil(o.Entities) {
		var ret []V1ActivityLogsPostRequestSelectorJobEntity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) GetEntitiesOk() ([]V1ActivityLogsPostRequestSelectorJobEntity, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []V1ActivityLogsPostRequestSelectorJobEntity and assigns it to the Entities field.
func (o *V1ActivityLogsPostRequestSelectorV2) SetEntities(v []V1ActivityLogsPostRequestSelectorJobEntity) {
	o.Entities = v
}

// GetSiteIds returns the SiteIds field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequestSelectorV2) GetSiteIds() []int64 {
	if o == nil || IsNil(o.SiteIds) {
		var ret []int64
		return ret
	}
	return o.SiteIds
}

// GetSiteIdsOk returns a tuple with the SiteIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) GetSiteIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.SiteIds) {
		return nil, false
	}
	return o.SiteIds, true
}

// HasSiteIds returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) HasSiteIds() bool {
	if o != nil && !IsNil(o.SiteIds) {
		return true
	}

	return false
}

// SetSiteIds gets a reference to the given []int64 and assigns it to the SiteIds field.
func (o *V1ActivityLogsPostRequestSelectorV2) SetSiteIds(v []int64) {
	o.SiteIds = v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequestSelectorV2) GetStatuses() []string {
	if o == nil || IsNil(o.Statuses) {
		var ret []string
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) GetStatusesOk() ([]string, bool) {
	if o == nil || IsNil(o.Statuses) {
		return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) HasStatuses() bool {
	if o != nil && !IsNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given []string and assigns it to the Statuses field.
func (o *V1ActivityLogsPostRequestSelectorV2) SetStatuses(v []string) {
	o.Statuses = v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequestSelectorV2) GetTypes() []string {
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *V1ActivityLogsPostRequestSelectorV2) SetTypes(v []string) {
	o.Types = v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise.
func (o *V1ActivityLogsPostRequestSelectorV2) GetUserIds() []string {
	if o == nil || IsNil(o.UserIds) {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *V1ActivityLogsPostRequestSelectorV2) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *V1ActivityLogsPostRequestSelectorV2) SetUserIds(v []string) {
	o.UserIds = v
}

func (o V1ActivityLogsPostRequestSelectorV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ActivityLogsPostRequestSelectorV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceIds) {
		toSerialize["deviceIds"] = o.DeviceIds
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.SiteIds) {
		toSerialize["siteIds"] = o.SiteIds
	}
	if !IsNil(o.Statuses) {
		toSerialize["statuses"] = o.Statuses
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.UserIds) {
		toSerialize["userIds"] = o.UserIds
	}
	return toSerialize, nil
}

type NullableV1ActivityLogsPostRequestSelectorV2 struct {
	value *V1ActivityLogsPostRequestSelectorV2
	isSet bool
}

func (v NullableV1ActivityLogsPostRequestSelectorV2) Get() *V1ActivityLogsPostRequestSelectorV2 {
	return v.value
}

func (v *NullableV1ActivityLogsPostRequestSelectorV2) Set(val *V1ActivityLogsPostRequestSelectorV2) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ActivityLogsPostRequestSelectorV2) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ActivityLogsPostRequestSelectorV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ActivityLogsPostRequestSelectorV2(val *V1ActivityLogsPostRequestSelectorV2) *NullableV1ActivityLogsPostRequestSelectorV2 {
	return &NullableV1ActivityLogsPostRequestSelectorV2{value: val, isSet: true}
}

func (v NullableV1ActivityLogsPostRequestSelectorV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ActivityLogsPostRequestSelectorV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


