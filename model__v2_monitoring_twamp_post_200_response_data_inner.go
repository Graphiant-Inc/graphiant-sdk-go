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

// checks if the V2MonitoringTwampPost200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2MonitoringTwampPost200ResponseDataInner{}

// V2MonitoringTwampPost200ResponseDataInner struct for V2MonitoringTwampPost200ResponseDataInner
type V2MonitoringTwampPost200ResponseDataInner struct {
	Samples []V2MonitoringBfdPost200ResponseDataInnerSamplesInner `json:"samples,omitempty"`
	Selector *V2MonitoringSiteTwampSiteIdPostRequestSelectorsInner `json:"selector,omitempty"`
}

// NewV2MonitoringTwampPost200ResponseDataInner instantiates a new V2MonitoringTwampPost200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2MonitoringTwampPost200ResponseDataInner() *V2MonitoringTwampPost200ResponseDataInner {
	this := V2MonitoringTwampPost200ResponseDataInner{}
	return &this
}

// NewV2MonitoringTwampPost200ResponseDataInnerWithDefaults instantiates a new V2MonitoringTwampPost200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2MonitoringTwampPost200ResponseDataInnerWithDefaults() *V2MonitoringTwampPost200ResponseDataInner {
	this := V2MonitoringTwampPost200ResponseDataInner{}
	return &this
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *V2MonitoringTwampPost200ResponseDataInner) GetSamples() []V2MonitoringBfdPost200ResponseDataInnerSamplesInner {
	if o == nil || IsNil(o.Samples) {
		var ret []V2MonitoringBfdPost200ResponseDataInnerSamplesInner
		return ret
	}
	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringTwampPost200ResponseDataInner) GetSamplesOk() ([]V2MonitoringBfdPost200ResponseDataInnerSamplesInner, bool) {
	if o == nil || IsNil(o.Samples) {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *V2MonitoringTwampPost200ResponseDataInner) HasSamples() bool {
	if o != nil && !IsNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []V2MonitoringBfdPost200ResponseDataInnerSamplesInner and assigns it to the Samples field.
func (o *V2MonitoringTwampPost200ResponseDataInner) SetSamples(v []V2MonitoringBfdPost200ResponseDataInnerSamplesInner) {
	o.Samples = v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *V2MonitoringTwampPost200ResponseDataInner) GetSelector() V2MonitoringSiteTwampSiteIdPostRequestSelectorsInner {
	if o == nil || IsNil(o.Selector) {
		var ret V2MonitoringSiteTwampSiteIdPostRequestSelectorsInner
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringTwampPost200ResponseDataInner) GetSelectorOk() (*V2MonitoringSiteTwampSiteIdPostRequestSelectorsInner, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *V2MonitoringTwampPost200ResponseDataInner) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given V2MonitoringSiteTwampSiteIdPostRequestSelectorsInner and assigns it to the Selector field.
func (o *V2MonitoringTwampPost200ResponseDataInner) SetSelector(v V2MonitoringSiteTwampSiteIdPostRequestSelectorsInner) {
	o.Selector = &v
}

func (o V2MonitoringTwampPost200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2MonitoringTwampPost200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	return toSerialize, nil
}

type NullableV2MonitoringTwampPost200ResponseDataInner struct {
	value *V2MonitoringTwampPost200ResponseDataInner
	isSet bool
}

func (v NullableV2MonitoringTwampPost200ResponseDataInner) Get() *V2MonitoringTwampPost200ResponseDataInner {
	return v.value
}

func (v *NullableV2MonitoringTwampPost200ResponseDataInner) Set(val *V2MonitoringTwampPost200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2MonitoringTwampPost200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2MonitoringTwampPost200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2MonitoringTwampPost200ResponseDataInner(val *V2MonitoringTwampPost200ResponseDataInner) *NullableV2MonitoringTwampPost200ResponseDataInner {
	return &NullableV2MonitoringTwampPost200ResponseDataInner{value: val, isSet: true}
}

func (v NullableV2MonitoringTwampPost200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2MonitoringTwampPost200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


