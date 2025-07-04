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

// checks if the V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface{}

// V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface struct for V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface
type V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface struct {
	CoreNeighbor *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor `json:"coreNeighbor,omitempty"`
	Default map[string]interface{} `json:"default,omitempty"`
	GatewayNeighbor *V1AccountMfaGet200Response `json:"gatewayNeighbor,omitempty"`
	Interface *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface `json:"interface,omitempty"`
	InterfaceType *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType `json:"interfaceType,omitempty"`
	VlanId *int32 `json:"vlanId,omitempty"`
	Wan *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan `json:"wan,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface {
	this := V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface {
	this := V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface{}
	return &this
}

// GetCoreNeighbor returns the CoreNeighbor field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetCoreNeighbor() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor {
	if o == nil || IsNil(o.CoreNeighbor) {
		var ret V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor
		return ret
	}
	return *o.CoreNeighbor
}

// GetCoreNeighborOk returns a tuple with the CoreNeighbor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetCoreNeighborOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor, bool) {
	if o == nil || IsNil(o.CoreNeighbor) {
		return nil, false
	}
	return o.CoreNeighbor, true
}

// HasCoreNeighbor returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasCoreNeighbor() bool {
	if o != nil && !IsNil(o.CoreNeighbor) {
		return true
	}

	return false
}

// SetCoreNeighbor gets a reference to the given V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor and assigns it to the CoreNeighbor field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetCoreNeighbor(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor) {
	o.CoreNeighbor = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetDefault() map[string]interface{} {
	if o == nil || IsNil(o.Default) {
		var ret map[string]interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetDefaultOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Default) {
		return map[string]interface{}{}, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given map[string]interface{} and assigns it to the Default field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetDefault(v map[string]interface{}) {
	o.Default = v
}

// GetGatewayNeighbor returns the GatewayNeighbor field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetGatewayNeighbor() V1AccountMfaGet200Response {
	if o == nil || IsNil(o.GatewayNeighbor) {
		var ret V1AccountMfaGet200Response
		return ret
	}
	return *o.GatewayNeighbor
}

// GetGatewayNeighborOk returns a tuple with the GatewayNeighbor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetGatewayNeighborOk() (*V1AccountMfaGet200Response, bool) {
	if o == nil || IsNil(o.GatewayNeighbor) {
		return nil, false
	}
	return o.GatewayNeighbor, true
}

// HasGatewayNeighbor returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasGatewayNeighbor() bool {
	if o != nil && !IsNil(o.GatewayNeighbor) {
		return true
	}

	return false
}

// SetGatewayNeighbor gets a reference to the given V1AccountMfaGet200Response and assigns it to the GatewayNeighbor field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetGatewayNeighbor(v V1AccountMfaGet200Response) {
	o.GatewayNeighbor = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetInterface() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface {
	if o == nil || IsNil(o.Interface) {
		var ret V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetInterfaceOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface and assigns it to the Interface field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetInterface(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface) {
	o.Interface = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetInterfaceType() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType {
	if o == nil || IsNil(o.InterfaceType) {
		var ret V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetInterfaceTypeOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType, bool) {
	if o == nil || IsNil(o.InterfaceType) {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasInterfaceType() bool {
	if o != nil && !IsNil(o.InterfaceType) {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType and assigns it to the InterfaceType field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetInterfaceType(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType) {
	o.InterfaceType = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetWan returns the Wan field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetWan() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan {
	if o == nil || IsNil(o.Wan) {
		var ret V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan
		return ret
	}
	return *o.Wan
}

// GetWanOk returns a tuple with the Wan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetWanOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan, bool) {
	if o == nil || IsNil(o.Wan) {
		return nil, false
	}
	return o.Wan, true
}

// HasWan returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasWan() bool {
	if o != nil && !IsNil(o.Wan) {
		return true
	}

	return false
}

// SetWan gets a reference to the given V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan and assigns it to the Wan field.
func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetWan(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan) {
	o.Wan = &v
}

func (o V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CoreNeighbor) {
		toSerialize["coreNeighbor"] = o.CoreNeighbor
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.GatewayNeighbor) {
		toSerialize["gatewayNeighbor"] = o.GatewayNeighbor
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.InterfaceType) {
		toSerialize["interfaceType"] = o.InterfaceType
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !IsNil(o.Wan) {
		toSerialize["wan"] = o.Wan
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface struct {
	value *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) Get() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) Set(val *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface(val *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface {
	return &NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


