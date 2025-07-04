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

// checks if the V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner{}

// V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner struct for V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner
type V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner struct {
	DataframeDictionaryMap *map[string]string `json:"dataframeDictionaryMap,omitempty"`
	XAxis *string `json:"xAxis,omitempty"`
	YAxis *string `json:"yAxis,omitempty"`
}

// NewV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner instantiates a new V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner() *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner {
	this := V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner{}
	return &this
}

// NewV2AssistantAddToConversationPost200ResponseDataframeDictionaryInnerWithDefaults instantiates a new V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssistantAddToConversationPost200ResponseDataframeDictionaryInnerWithDefaults() *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner {
	this := V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner{}
	return &this
}

// GetDataframeDictionaryMap returns the DataframeDictionaryMap field value if set, zero value otherwise.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) GetDataframeDictionaryMap() map[string]string {
	if o == nil || IsNil(o.DataframeDictionaryMap) {
		var ret map[string]string
		return ret
	}
	return *o.DataframeDictionaryMap
}

// GetDataframeDictionaryMapOk returns a tuple with the DataframeDictionaryMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) GetDataframeDictionaryMapOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DataframeDictionaryMap) {
		return nil, false
	}
	return o.DataframeDictionaryMap, true
}

// HasDataframeDictionaryMap returns a boolean if a field has been set.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) HasDataframeDictionaryMap() bool {
	if o != nil && !IsNil(o.DataframeDictionaryMap) {
		return true
	}

	return false
}

// SetDataframeDictionaryMap gets a reference to the given map[string]string and assigns it to the DataframeDictionaryMap field.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) SetDataframeDictionaryMap(v map[string]string) {
	o.DataframeDictionaryMap = &v
}

// GetXAxis returns the XAxis field value if set, zero value otherwise.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) GetXAxis() string {
	if o == nil || IsNil(o.XAxis) {
		var ret string
		return ret
	}
	return *o.XAxis
}

// GetXAxisOk returns a tuple with the XAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) GetXAxisOk() (*string, bool) {
	if o == nil || IsNil(o.XAxis) {
		return nil, false
	}
	return o.XAxis, true
}

// HasXAxis returns a boolean if a field has been set.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) HasXAxis() bool {
	if o != nil && !IsNil(o.XAxis) {
		return true
	}

	return false
}

// SetXAxis gets a reference to the given string and assigns it to the XAxis field.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) SetXAxis(v string) {
	o.XAxis = &v
}

// GetYAxis returns the YAxis field value if set, zero value otherwise.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) GetYAxis() string {
	if o == nil || IsNil(o.YAxis) {
		var ret string
		return ret
	}
	return *o.YAxis
}

// GetYAxisOk returns a tuple with the YAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) GetYAxisOk() (*string, bool) {
	if o == nil || IsNil(o.YAxis) {
		return nil, false
	}
	return o.YAxis, true
}

// HasYAxis returns a boolean if a field has been set.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) HasYAxis() bool {
	if o != nil && !IsNil(o.YAxis) {
		return true
	}

	return false
}

// SetYAxis gets a reference to the given string and assigns it to the YAxis field.
func (o *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) SetYAxis(v string) {
	o.YAxis = &v
}

func (o V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataframeDictionaryMap) {
		toSerialize["dataframeDictionaryMap"] = o.DataframeDictionaryMap
	}
	if !IsNil(o.XAxis) {
		toSerialize["xAxis"] = o.XAxis
	}
	if !IsNil(o.YAxis) {
		toSerialize["yAxis"] = o.YAxis
	}
	return toSerialize, nil
}

type NullableV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner struct {
	value *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner
	isSet bool
}

func (v NullableV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) Get() *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner {
	return v.value
}

func (v *NullableV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) Set(val *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner(val *V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) *NullableV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner {
	return &NullableV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner{value: val, isSet: true}
}

func (v NullableV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssistantAddToConversationPost200ResponseDataframeDictionaryInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


