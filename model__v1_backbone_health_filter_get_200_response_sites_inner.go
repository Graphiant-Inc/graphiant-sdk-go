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

// checks if the V1BackboneHealthFilterGet200ResponseSitesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BackboneHealthFilterGet200ResponseSitesInner{}

// V1BackboneHealthFilterGet200ResponseSitesInner struct for V1BackboneHealthFilterGet200ResponseSitesInner
type V1BackboneHealthFilterGet200ResponseSitesInner struct {
	SiteId *int64 `json:"siteId,omitempty"`
	SiteName *string `json:"siteName,omitempty"`
}

// NewV1BackboneHealthFilterGet200ResponseSitesInner instantiates a new V1BackboneHealthFilterGet200ResponseSitesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BackboneHealthFilterGet200ResponseSitesInner() *V1BackboneHealthFilterGet200ResponseSitesInner {
	this := V1BackboneHealthFilterGet200ResponseSitesInner{}
	return &this
}

// NewV1BackboneHealthFilterGet200ResponseSitesInnerWithDefaults instantiates a new V1BackboneHealthFilterGet200ResponseSitesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BackboneHealthFilterGet200ResponseSitesInnerWithDefaults() *V1BackboneHealthFilterGet200ResponseSitesInner {
	this := V1BackboneHealthFilterGet200ResponseSitesInner{}
	return &this
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *V1BackboneHealthFilterGet200ResponseSitesInner) GetSiteId() int64 {
	if o == nil || IsNil(o.SiteId) {
		var ret int64
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthFilterGet200ResponseSitesInner) GetSiteIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *V1BackboneHealthFilterGet200ResponseSitesInner) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given int64 and assigns it to the SiteId field.
func (o *V1BackboneHealthFilterGet200ResponseSitesInner) SetSiteId(v int64) {
	o.SiteId = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *V1BackboneHealthFilterGet200ResponseSitesInner) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BackboneHealthFilterGet200ResponseSitesInner) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *V1BackboneHealthFilterGet200ResponseSitesInner) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *V1BackboneHealthFilterGet200ResponseSitesInner) SetSiteName(v string) {
	o.SiteName = &v
}

func (o V1BackboneHealthFilterGet200ResponseSitesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BackboneHealthFilterGet200ResponseSitesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	if !IsNil(o.SiteName) {
		toSerialize["siteName"] = o.SiteName
	}
	return toSerialize, nil
}

type NullableV1BackboneHealthFilterGet200ResponseSitesInner struct {
	value *V1BackboneHealthFilterGet200ResponseSitesInner
	isSet bool
}

func (v NullableV1BackboneHealthFilterGet200ResponseSitesInner) Get() *V1BackboneHealthFilterGet200ResponseSitesInner {
	return v.value
}

func (v *NullableV1BackboneHealthFilterGet200ResponseSitesInner) Set(val *V1BackboneHealthFilterGet200ResponseSitesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BackboneHealthFilterGet200ResponseSitesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BackboneHealthFilterGet200ResponseSitesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BackboneHealthFilterGet200ResponseSitesInner(val *V1BackboneHealthFilterGet200ResponseSitesInner) *NullableV1BackboneHealthFilterGet200ResponseSitesInner {
	return &NullableV1BackboneHealthFilterGet200ResponseSitesInner{value: val, isSet: true}
}

func (v NullableV1BackboneHealthFilterGet200ResponseSitesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BackboneHealthFilterGet200ResponseSitesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


