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

// checks if the V2AssuranceEndpointIntelPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2AssuranceEndpointIntelPost200Response{}

// V2AssuranceEndpointIntelPost200Response struct for V2AssuranceEndpointIntelPost200Response
type V2AssuranceEndpointIntelPost200Response struct {
	AmCategories []string `json:"amCategories,omitempty"`
	AmRiskScore *float32 `json:"amRiskScore,omitempty"`
	GraphiantRisky *bool `json:"graphiantRisky,omitempty"`
}

// NewV2AssuranceEndpointIntelPost200Response instantiates a new V2AssuranceEndpointIntelPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssuranceEndpointIntelPost200Response() *V2AssuranceEndpointIntelPost200Response {
	this := V2AssuranceEndpointIntelPost200Response{}
	return &this
}

// NewV2AssuranceEndpointIntelPost200ResponseWithDefaults instantiates a new V2AssuranceEndpointIntelPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssuranceEndpointIntelPost200ResponseWithDefaults() *V2AssuranceEndpointIntelPost200Response {
	this := V2AssuranceEndpointIntelPost200Response{}
	return &this
}

// GetAmCategories returns the AmCategories field value if set, zero value otherwise.
func (o *V2AssuranceEndpointIntelPost200Response) GetAmCategories() []string {
	if o == nil || IsNil(o.AmCategories) {
		var ret []string
		return ret
	}
	return o.AmCategories
}

// GetAmCategoriesOk returns a tuple with the AmCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceEndpointIntelPost200Response) GetAmCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.AmCategories) {
		return nil, false
	}
	return o.AmCategories, true
}

// HasAmCategories returns a boolean if a field has been set.
func (o *V2AssuranceEndpointIntelPost200Response) HasAmCategories() bool {
	if o != nil && !IsNil(o.AmCategories) {
		return true
	}

	return false
}

// SetAmCategories gets a reference to the given []string and assigns it to the AmCategories field.
func (o *V2AssuranceEndpointIntelPost200Response) SetAmCategories(v []string) {
	o.AmCategories = v
}

// GetAmRiskScore returns the AmRiskScore field value if set, zero value otherwise.
func (o *V2AssuranceEndpointIntelPost200Response) GetAmRiskScore() float32 {
	if o == nil || IsNil(o.AmRiskScore) {
		var ret float32
		return ret
	}
	return *o.AmRiskScore
}

// GetAmRiskScoreOk returns a tuple with the AmRiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceEndpointIntelPost200Response) GetAmRiskScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.AmRiskScore) {
		return nil, false
	}
	return o.AmRiskScore, true
}

// HasAmRiskScore returns a boolean if a field has been set.
func (o *V2AssuranceEndpointIntelPost200Response) HasAmRiskScore() bool {
	if o != nil && !IsNil(o.AmRiskScore) {
		return true
	}

	return false
}

// SetAmRiskScore gets a reference to the given float32 and assigns it to the AmRiskScore field.
func (o *V2AssuranceEndpointIntelPost200Response) SetAmRiskScore(v float32) {
	o.AmRiskScore = &v
}

// GetGraphiantRisky returns the GraphiantRisky field value if set, zero value otherwise.
func (o *V2AssuranceEndpointIntelPost200Response) GetGraphiantRisky() bool {
	if o == nil || IsNil(o.GraphiantRisky) {
		var ret bool
		return ret
	}
	return *o.GraphiantRisky
}

// GetGraphiantRiskyOk returns a tuple with the GraphiantRisky field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceEndpointIntelPost200Response) GetGraphiantRiskyOk() (*bool, bool) {
	if o == nil || IsNil(o.GraphiantRisky) {
		return nil, false
	}
	return o.GraphiantRisky, true
}

// HasGraphiantRisky returns a boolean if a field has been set.
func (o *V2AssuranceEndpointIntelPost200Response) HasGraphiantRisky() bool {
	if o != nil && !IsNil(o.GraphiantRisky) {
		return true
	}

	return false
}

// SetGraphiantRisky gets a reference to the given bool and assigns it to the GraphiantRisky field.
func (o *V2AssuranceEndpointIntelPost200Response) SetGraphiantRisky(v bool) {
	o.GraphiantRisky = &v
}

func (o V2AssuranceEndpointIntelPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2AssuranceEndpointIntelPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmCategories) {
		toSerialize["amCategories"] = o.AmCategories
	}
	if !IsNil(o.AmRiskScore) {
		toSerialize["amRiskScore"] = o.AmRiskScore
	}
	if !IsNil(o.GraphiantRisky) {
		toSerialize["graphiantRisky"] = o.GraphiantRisky
	}
	return toSerialize, nil
}

type NullableV2AssuranceEndpointIntelPost200Response struct {
	value *V2AssuranceEndpointIntelPost200Response
	isSet bool
}

func (v NullableV2AssuranceEndpointIntelPost200Response) Get() *V2AssuranceEndpointIntelPost200Response {
	return v.value
}

func (v *NullableV2AssuranceEndpointIntelPost200Response) Set(val *V2AssuranceEndpointIntelPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssuranceEndpointIntelPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssuranceEndpointIntelPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssuranceEndpointIntelPost200Response(val *V2AssuranceEndpointIntelPost200Response) *NullableV2AssuranceEndpointIntelPost200Response {
	return &NullableV2AssuranceEndpointIntelPost200Response{value: val, isSet: true}
}

func (v NullableV2AssuranceEndpointIntelPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssuranceEndpointIntelPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


