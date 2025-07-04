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

// checks if the V2SiteSiteIdDetailPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2SiteSiteIdDetailPost200Response{}

// V2SiteSiteIdDetailPost200Response struct for V2SiteSiteIdDetailPost200Response
type V2SiteSiteIdDetailPost200Response struct {
	Site *V2SiteSiteIdDetailPost200ResponseSite `json:"site,omitempty"`
	Snapshots []V1AlarmHistoryGet200ResponseHistoryInnerTime `json:"snapshots,omitempty"`
}

// NewV2SiteSiteIdDetailPost200Response instantiates a new V2SiteSiteIdDetailPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2SiteSiteIdDetailPost200Response() *V2SiteSiteIdDetailPost200Response {
	this := V2SiteSiteIdDetailPost200Response{}
	return &this
}

// NewV2SiteSiteIdDetailPost200ResponseWithDefaults instantiates a new V2SiteSiteIdDetailPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2SiteSiteIdDetailPost200ResponseWithDefaults() *V2SiteSiteIdDetailPost200Response {
	this := V2SiteSiteIdDetailPost200Response{}
	return &this
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *V2SiteSiteIdDetailPost200Response) GetSite() V2SiteSiteIdDetailPost200ResponseSite {
	if o == nil || IsNil(o.Site) {
		var ret V2SiteSiteIdDetailPost200ResponseSite
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2SiteSiteIdDetailPost200Response) GetSiteOk() (*V2SiteSiteIdDetailPost200ResponseSite, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *V2SiteSiteIdDetailPost200Response) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given V2SiteSiteIdDetailPost200ResponseSite and assigns it to the Site field.
func (o *V2SiteSiteIdDetailPost200Response) SetSite(v V2SiteSiteIdDetailPost200ResponseSite) {
	o.Site = &v
}

// GetSnapshots returns the Snapshots field value if set, zero value otherwise.
func (o *V2SiteSiteIdDetailPost200Response) GetSnapshots() []V1AlarmHistoryGet200ResponseHistoryInnerTime {
	if o == nil || IsNil(o.Snapshots) {
		var ret []V1AlarmHistoryGet200ResponseHistoryInnerTime
		return ret
	}
	return o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2SiteSiteIdDetailPost200Response) GetSnapshotsOk() ([]V1AlarmHistoryGet200ResponseHistoryInnerTime, bool) {
	if o == nil || IsNil(o.Snapshots) {
		return nil, false
	}
	return o.Snapshots, true
}

// HasSnapshots returns a boolean if a field has been set.
func (o *V2SiteSiteIdDetailPost200Response) HasSnapshots() bool {
	if o != nil && !IsNil(o.Snapshots) {
		return true
	}

	return false
}

// SetSnapshots gets a reference to the given []V1AlarmHistoryGet200ResponseHistoryInnerTime and assigns it to the Snapshots field.
func (o *V2SiteSiteIdDetailPost200Response) SetSnapshots(v []V1AlarmHistoryGet200ResponseHistoryInnerTime) {
	o.Snapshots = v
}

func (o V2SiteSiteIdDetailPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2SiteSiteIdDetailPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.Snapshots) {
		toSerialize["snapshots"] = o.Snapshots
	}
	return toSerialize, nil
}

type NullableV2SiteSiteIdDetailPost200Response struct {
	value *V2SiteSiteIdDetailPost200Response
	isSet bool
}

func (v NullableV2SiteSiteIdDetailPost200Response) Get() *V2SiteSiteIdDetailPost200Response {
	return v.value
}

func (v *NullableV2SiteSiteIdDetailPost200Response) Set(val *V2SiteSiteIdDetailPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2SiteSiteIdDetailPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2SiteSiteIdDetailPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2SiteSiteIdDetailPost200Response(val *V2SiteSiteIdDetailPost200Response) *NullableV2SiteSiteIdDetailPost200Response {
	return &NullableV2SiteSiteIdDetailPost200Response{value: val, isSet: true}
}

func (v NullableV2SiteSiteIdDetailPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2SiteSiteIdDetailPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


