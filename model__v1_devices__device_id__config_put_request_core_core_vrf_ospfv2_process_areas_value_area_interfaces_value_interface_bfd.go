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

// checks if the V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd{}

// V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd struct for V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd
type V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd struct {
	Bfd map[string]interface{} `json:"bfd,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd {
	this := V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfdWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfdWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd {
	this := V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd{}
	return &this
}

// GetBfd returns the Bfd field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) GetBfd() map[string]interface{} {
	if o == nil || IsNil(o.Bfd) {
		var ret map[string]interface{}
		return ret
	}
	return o.Bfd
}

// GetBfdOk returns a tuple with the Bfd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) GetBfdOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Bfd) {
		return map[string]interface{}{}, false
	}
	return o.Bfd, true
}

// HasBfd returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) HasBfd() bool {
	if o != nil && !IsNil(o.Bfd) {
		return true
	}

	return false
}

// SetBfd gets a reference to the given map[string]interface{} and assigns it to the Bfd field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) SetBfd(v map[string]interface{}) {
	o.Bfd = v
}

func (o V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bfd) {
		toSerialize["bfd"] = o.Bfd
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd struct {
	value *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) Get() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) Set(val *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd(val *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd {
	return &NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceBfd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


