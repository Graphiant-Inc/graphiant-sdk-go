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

// checks if the V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp{}

// V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp struct for V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp
type V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp struct {
	AppName *string `json:"appName,omitempty"`
	BucketId *string `json:"bucketId,omitempty"`
}

// NewV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp instantiates a new V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp() *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp {
	this := V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp{}
	return &this
}

// NewV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerAppWithDefaults instantiates a new V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerAppWithDefaults() *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp {
	this := V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp{}
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) SetAppName(v string) {
	o.AppName = &v
}

// GetBucketId returns the BucketId field value if set, zero value otherwise.
func (o *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) GetBucketId() string {
	if o == nil || IsNil(o.BucketId) {
		var ret string
		return ret
	}
	return *o.BucketId
}

// GetBucketIdOk returns a tuple with the BucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) GetBucketIdOk() (*string, bool) {
	if o == nil || IsNil(o.BucketId) {
		return nil, false
	}
	return o.BucketId, true
}

// HasBucketId returns a boolean if a field has been set.
func (o *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) HasBucketId() bool {
	if o != nil && !IsNil(o.BucketId) {
		return true
	}

	return false
}

// SetBucketId gets a reference to the given string and assigns it to the BucketId field.
func (o *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) SetBucketId(v string) {
	o.BucketId = &v
}

func (o V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppName) {
		toSerialize["appName"] = o.AppName
	}
	if !IsNil(o.BucketId) {
		toSerialize["bucketId"] = o.BucketId
	}
	return toSerialize, nil
}

type NullableV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp struct {
	value *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp
	isSet bool
}

func (v NullableV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) Get() *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp {
	return v.value
}

func (v *NullableV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) Set(val *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp(val *V2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) *NullableV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp {
	return &NullableV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp{value: val, isSet: true}
}

func (v NullableV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssuranceBucketAppServersAllGet200ResponseAppServerChangesInnerApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


