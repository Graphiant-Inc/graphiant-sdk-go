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

// checks if the V2MonitoringCircuitsUtilizationPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2MonitoringCircuitsUtilizationPost200Response{}

// V2MonitoringCircuitsUtilizationPost200Response struct for V2MonitoringCircuitsUtilizationPost200Response
type V2MonitoringCircuitsUtilizationPost200Response struct {
	Data []V2MonitoringCircuitsUtilizationPost200ResponseDataInner `json:"data,omitempty"`
}

// NewV2MonitoringCircuitsUtilizationPost200Response instantiates a new V2MonitoringCircuitsUtilizationPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2MonitoringCircuitsUtilizationPost200Response() *V2MonitoringCircuitsUtilizationPost200Response {
	this := V2MonitoringCircuitsUtilizationPost200Response{}
	return &this
}

// NewV2MonitoringCircuitsUtilizationPost200ResponseWithDefaults instantiates a new V2MonitoringCircuitsUtilizationPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2MonitoringCircuitsUtilizationPost200ResponseWithDefaults() *V2MonitoringCircuitsUtilizationPost200Response {
	this := V2MonitoringCircuitsUtilizationPost200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V2MonitoringCircuitsUtilizationPost200Response) GetData() []V2MonitoringCircuitsUtilizationPost200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []V2MonitoringCircuitsUtilizationPost200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MonitoringCircuitsUtilizationPost200Response) GetDataOk() ([]V2MonitoringCircuitsUtilizationPost200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V2MonitoringCircuitsUtilizationPost200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []V2MonitoringCircuitsUtilizationPost200ResponseDataInner and assigns it to the Data field.
func (o *V2MonitoringCircuitsUtilizationPost200Response) SetData(v []V2MonitoringCircuitsUtilizationPost200ResponseDataInner) {
	o.Data = v
}

func (o V2MonitoringCircuitsUtilizationPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2MonitoringCircuitsUtilizationPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableV2MonitoringCircuitsUtilizationPost200Response struct {
	value *V2MonitoringCircuitsUtilizationPost200Response
	isSet bool
}

func (v NullableV2MonitoringCircuitsUtilizationPost200Response) Get() *V2MonitoringCircuitsUtilizationPost200Response {
	return v.value
}

func (v *NullableV2MonitoringCircuitsUtilizationPost200Response) Set(val *V2MonitoringCircuitsUtilizationPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2MonitoringCircuitsUtilizationPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2MonitoringCircuitsUtilizationPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2MonitoringCircuitsUtilizationPost200Response(val *V2MonitoringCircuitsUtilizationPost200Response) *NullableV2MonitoringCircuitsUtilizationPost200Response {
	return &NullableV2MonitoringCircuitsUtilizationPost200Response{value: val, isSet: true}
}

func (v NullableV2MonitoringCircuitsUtilizationPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2MonitoringCircuitsUtilizationPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


