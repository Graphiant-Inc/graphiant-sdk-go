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

// checks if the V1BackboneHealthOverviewPostRequestFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BackboneHealthOverviewPostRequestFilter{}

// V1BackboneHealthOverviewPostRequestFilter struct for V1BackboneHealthOverviewPostRequestFilter
type V1BackboneHealthOverviewPostRequestFilter struct {
	CircuitNames []string `json:"circuitNames,omitempty"`
	DeviceIds []int64 `json:"deviceIds,omitempty"`
	LanSegments []string `json:"lanSegments,omitempty"`
	RegionIds []int64 `json:"regionIds,omitempty"`
	SiteIds []int64 `json:"siteIds,omitempty"`
}

// NewV1BackboneHealthOverviewPostRequestFilter instantiates a new V1BackboneHealthOverviewPostRequestFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BackboneHealthOverviewPostRequestFilter() *V1BackboneHealthOverviewPostRequestFilter {
	this := V1BackboneHealthOverviewPostRequestFilter{}
	return &this
}

// NewV1BackboneHealthOverviewPostRequestFilterWithDefaults instantiates a new V1BackboneHealthOverviewPostRequestFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BackboneHealthOverviewPostRequestFilterWithDefaults() *V1BackboneHealthOverviewPostRequestFilter {
	this := V1BackboneHealthOverviewPostRequestFilter{}
	return &this
}

// GetCircuitNames returns the CircuitNames field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPostRequestFilter) GetCircuitNames() []string {
	if o == nil || IsNil(o.CircuitNames) {
		var ret []string
		return ret
	}
	return o.CircuitNames
}

// GetCircuitNamesOk returns a tuple with the CircuitNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPostRequestFilter) GetCircuitNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CircuitNames) {
		return nil, false
	}
	return o.CircuitNames, true
}

// HasCircuitNames returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPostRequestFilter) HasCircuitNames() bool {
	if o != nil && !IsNil(o.CircuitNames) {
		return true
	}

	return false
}

// SetCircuitNames gets a reference to the given []string and assigns it to the CircuitNames field.
func (o *V1BackboneHealthOverviewPostRequestFilter) SetCircuitNames(v []string) {
	o.CircuitNames = v
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPostRequestFilter) GetDeviceIds() []int64 {
	if o == nil || IsNil(o.DeviceIds) {
		var ret []int64
		return ret
	}
	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPostRequestFilter) GetDeviceIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.DeviceIds) {
		return nil, false
	}
	return o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPostRequestFilter) HasDeviceIds() bool {
	if o != nil && !IsNil(o.DeviceIds) {
		return true
	}

	return false
}

// SetDeviceIds gets a reference to the given []int64 and assigns it to the DeviceIds field.
func (o *V1BackboneHealthOverviewPostRequestFilter) SetDeviceIds(v []int64) {
	o.DeviceIds = v
}

// GetLanSegments returns the LanSegments field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPostRequestFilter) GetLanSegments() []string {
	if o == nil || IsNil(o.LanSegments) {
		var ret []string
		return ret
	}
	return o.LanSegments
}

// GetLanSegmentsOk returns a tuple with the LanSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPostRequestFilter) GetLanSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.LanSegments) {
		return nil, false
	}
	return o.LanSegments, true
}

// HasLanSegments returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPostRequestFilter) HasLanSegments() bool {
	if o != nil && !IsNil(o.LanSegments) {
		return true
	}

	return false
}

// SetLanSegments gets a reference to the given []string and assigns it to the LanSegments field.
func (o *V1BackboneHealthOverviewPostRequestFilter) SetLanSegments(v []string) {
	o.LanSegments = v
}

// GetRegionIds returns the RegionIds field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPostRequestFilter) GetRegionIds() []int64 {
	if o == nil || IsNil(o.RegionIds) {
		var ret []int64
		return ret
	}
	return o.RegionIds
}

// GetRegionIdsOk returns a tuple with the RegionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPostRequestFilter) GetRegionIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.RegionIds) {
		return nil, false
	}
	return o.RegionIds, true
}

// HasRegionIds returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPostRequestFilter) HasRegionIds() bool {
	if o != nil && !IsNil(o.RegionIds) {
		return true
	}

	return false
}

// SetRegionIds gets a reference to the given []int64 and assigns it to the RegionIds field.
func (o *V1BackboneHealthOverviewPostRequestFilter) SetRegionIds(v []int64) {
	o.RegionIds = v
}

// GetSiteIds returns the SiteIds field value if set, zero value otherwise.
func (o *V1BackboneHealthOverviewPostRequestFilter) GetSiteIds() []int64 {
	if o == nil || IsNil(o.SiteIds) {
		var ret []int64
		return ret
	}
	return o.SiteIds
}

// GetSiteIdsOk returns a tuple with the SiteIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthOverviewPostRequestFilter) GetSiteIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.SiteIds) {
		return nil, false
	}
	return o.SiteIds, true
}

// HasSiteIds returns a boolean if a field has been set.
func (o *V1BackboneHealthOverviewPostRequestFilter) HasSiteIds() bool {
	if o != nil && !IsNil(o.SiteIds) {
		return true
	}

	return false
}

// SetSiteIds gets a reference to the given []int64 and assigns it to the SiteIds field.
func (o *V1BackboneHealthOverviewPostRequestFilter) SetSiteIds(v []int64) {
	o.SiteIds = v
}

func (o V1BackboneHealthOverviewPostRequestFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BackboneHealthOverviewPostRequestFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CircuitNames) {
		toSerialize["circuitNames"] = o.CircuitNames
	}
	if !IsNil(o.DeviceIds) {
		toSerialize["deviceIds"] = o.DeviceIds
	}
	if !IsNil(o.LanSegments) {
		toSerialize["lanSegments"] = o.LanSegments
	}
	if !IsNil(o.RegionIds) {
		toSerialize["regionIds"] = o.RegionIds
	}
	if !IsNil(o.SiteIds) {
		toSerialize["siteIds"] = o.SiteIds
	}
	return toSerialize, nil
}

type NullableV1BackboneHealthOverviewPostRequestFilter struct {
	value *V1BackboneHealthOverviewPostRequestFilter
	isSet bool
}

func (v NullableV1BackboneHealthOverviewPostRequestFilter) Get() *V1BackboneHealthOverviewPostRequestFilter {
	return v.value
}

func (v *NullableV1BackboneHealthOverviewPostRequestFilter) Set(val *V1BackboneHealthOverviewPostRequestFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BackboneHealthOverviewPostRequestFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BackboneHealthOverviewPostRequestFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BackboneHealthOverviewPostRequestFilter(val *V1BackboneHealthOverviewPostRequestFilter) *NullableV1BackboneHealthOverviewPostRequestFilter {
	return &NullableV1BackboneHealthOverviewPostRequestFilter{value: val, isSet: true}
}

func (v NullableV1BackboneHealthOverviewPostRequestFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BackboneHealthOverviewPostRequestFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


