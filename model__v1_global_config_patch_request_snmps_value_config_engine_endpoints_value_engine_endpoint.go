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

// checks if the V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint{}

// V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint struct for V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint
type V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint struct {
	Addresses *map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue `json:"addresses,omitempty"`
	AutoIpv4 *bool `json:"autoIpv4,omitempty"`
	AutoIpv6 *bool `json:"autoIpv6,omitempty"`
	Interface *string `json:"interface,omitempty"`
	LanSegment *string `json:"lanSegment,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint() *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint {
	this := V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint{}
	return &this
}

// NewV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointWithDefaults instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointWithDefaults() *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint {
	this := V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAddresses() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue {
	if o == nil || IsNil(o.Addresses) {
		var ret map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAddressesOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue and assigns it to the Addresses field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetAddresses(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue) {
	o.Addresses = &v
}

// GetAutoIpv4 returns the AutoIpv4 field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAutoIpv4() bool {
	if o == nil || IsNil(o.AutoIpv4) {
		var ret bool
		return ret
	}
	return *o.AutoIpv4
}

// GetAutoIpv4Ok returns a tuple with the AutoIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAutoIpv4Ok() (*bool, bool) {
	if o == nil || IsNil(o.AutoIpv4) {
		return nil, false
	}
	return o.AutoIpv4, true
}

// HasAutoIpv4 returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasAutoIpv4() bool {
	if o != nil && !IsNil(o.AutoIpv4) {
		return true
	}

	return false
}

// SetAutoIpv4 gets a reference to the given bool and assigns it to the AutoIpv4 field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetAutoIpv4(v bool) {
	o.AutoIpv4 = &v
}

// GetAutoIpv6 returns the AutoIpv6 field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAutoIpv6() bool {
	if o == nil || IsNil(o.AutoIpv6) {
		var ret bool
		return ret
	}
	return *o.AutoIpv6
}

// GetAutoIpv6Ok returns a tuple with the AutoIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAutoIpv6Ok() (*bool, bool) {
	if o == nil || IsNil(o.AutoIpv6) {
		return nil, false
	}
	return o.AutoIpv6, true
}

// HasAutoIpv6 returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasAutoIpv6() bool {
	if o != nil && !IsNil(o.AutoIpv6) {
		return true
	}

	return false
}

// SetAutoIpv6 gets a reference to the given bool and assigns it to the AutoIpv6 field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetAutoIpv6(v bool) {
	o.AutoIpv6 = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetInterface(v string) {
	o.Interface = &v
}

// GetLanSegment returns the LanSegment field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetLanSegment() string {
	if o == nil || IsNil(o.LanSegment) {
		var ret string
		return ret
	}
	return *o.LanSegment
}

// GetLanSegmentOk returns a tuple with the LanSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetLanSegmentOk() (*string, bool) {
	if o == nil || IsNil(o.LanSegment) {
		return nil, false
	}
	return o.LanSegment, true
}

// HasLanSegment returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasLanSegment() bool {
	if o != nil && !IsNil(o.LanSegment) {
		return true
	}

	return false
}

// SetLanSegment gets a reference to the given string and assigns it to the LanSegment field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetLanSegment(v string) {
	o.LanSegment = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetName(v string) {
	o.Name = &v
}

func (o V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.AutoIpv4) {
		toSerialize["autoIpv4"] = o.AutoIpv4
	}
	if !IsNil(o.AutoIpv6) {
		toSerialize["autoIpv6"] = o.AutoIpv6
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.LanSegment) {
		toSerialize["lanSegment"] = o.LanSegment
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint struct {
	value *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) Get() *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) Set(val *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint(val *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) *NullableV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint {
	return &NullableV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


