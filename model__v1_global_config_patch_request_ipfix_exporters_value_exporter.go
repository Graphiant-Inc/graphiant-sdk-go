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

// checks if the V1GlobalConfigPatchRequestIpfixExportersValueExporter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestIpfixExportersValueExporter{}

// V1GlobalConfigPatchRequestIpfixExportersValueExporter struct for V1GlobalConfigPatchRequestIpfixExportersValueExporter
type V1GlobalConfigPatchRequestIpfixExportersValueExporter struct {
	DestinationAddress *string `json:"destinationAddress,omitempty"`
	DestinationPort *int32 `json:"destinationPort,omitempty"`
	GlobalId *int64 `json:"globalId,omitempty"`
	IsGlobalSync *bool `json:"isGlobalSync,omitempty"`
	MonitoredSegments []string `json:"monitoredSegments,omitempty"`
	Name *string `json:"name,omitempty"`
	SampleMode *string `json:"sampleMode,omitempty"`
	SampleRate *int32 `json:"sampleRate,omitempty"`
	SourceInterfaceName *string `json:"sourceInterfaceName,omitempty"`
	VrfId *int64 `json:"vrfId,omitempty"`
}

// NewV1GlobalConfigPatchRequestIpfixExportersValueExporter instantiates a new V1GlobalConfigPatchRequestIpfixExportersValueExporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestIpfixExportersValueExporter() *V1GlobalConfigPatchRequestIpfixExportersValueExporter {
	this := V1GlobalConfigPatchRequestIpfixExportersValueExporter{}
	return &this
}

// NewV1GlobalConfigPatchRequestIpfixExportersValueExporterWithDefaults instantiates a new V1GlobalConfigPatchRequestIpfixExportersValueExporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestIpfixExportersValueExporterWithDefaults() *V1GlobalConfigPatchRequestIpfixExportersValueExporter {
	this := V1GlobalConfigPatchRequestIpfixExportersValueExporter{}
	return &this
}

// GetDestinationAddress returns the DestinationAddress field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetDestinationAddress() string {
	if o == nil || IsNil(o.DestinationAddress) {
		var ret string
		return ret
	}
	return *o.DestinationAddress
}

// GetDestinationAddressOk returns a tuple with the DestinationAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetDestinationAddressOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationAddress) {
		return nil, false
	}
	return o.DestinationAddress, true
}

// HasDestinationAddress returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasDestinationAddress() bool {
	if o != nil && !IsNil(o.DestinationAddress) {
		return true
	}

	return false
}

// SetDestinationAddress gets a reference to the given string and assigns it to the DestinationAddress field.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetDestinationAddress(v string) {
	o.DestinationAddress = &v
}

// GetDestinationPort returns the DestinationPort field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetDestinationPort() int32 {
	if o == nil || IsNil(o.DestinationPort) {
		var ret int32
		return ret
	}
	return *o.DestinationPort
}

// GetDestinationPortOk returns a tuple with the DestinationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetDestinationPortOk() (*int32, bool) {
	if o == nil || IsNil(o.DestinationPort) {
		return nil, false
	}
	return o.DestinationPort, true
}

// HasDestinationPort returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasDestinationPort() bool {
	if o != nil && !IsNil(o.DestinationPort) {
		return true
	}

	return false
}

// SetDestinationPort gets a reference to the given int32 and assigns it to the DestinationPort field.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetDestinationPort(v int32) {
	o.DestinationPort = &v
}

// GetGlobalId returns the GlobalId field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetGlobalId() int64 {
	if o == nil || IsNil(o.GlobalId) {
		var ret int64
		return ret
	}
	return *o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetGlobalIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GlobalId) {
		return nil, false
	}
	return o.GlobalId, true
}

// HasGlobalId returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasGlobalId() bool {
	if o != nil && !IsNil(o.GlobalId) {
		return true
	}

	return false
}

// SetGlobalId gets a reference to the given int64 and assigns it to the GlobalId field.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetGlobalId(v int64) {
	o.GlobalId = &v
}

// GetIsGlobalSync returns the IsGlobalSync field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetIsGlobalSync() bool {
	if o == nil || IsNil(o.IsGlobalSync) {
		var ret bool
		return ret
	}
	return *o.IsGlobalSync
}

// GetIsGlobalSyncOk returns a tuple with the IsGlobalSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetIsGlobalSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobalSync) {
		return nil, false
	}
	return o.IsGlobalSync, true
}

// HasIsGlobalSync returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasIsGlobalSync() bool {
	if o != nil && !IsNil(o.IsGlobalSync) {
		return true
	}

	return false
}

// SetIsGlobalSync gets a reference to the given bool and assigns it to the IsGlobalSync field.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetIsGlobalSync(v bool) {
	o.IsGlobalSync = &v
}

// GetMonitoredSegments returns the MonitoredSegments field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetMonitoredSegments() []string {
	if o == nil || IsNil(o.MonitoredSegments) {
		var ret []string
		return ret
	}
	return o.MonitoredSegments
}

// GetMonitoredSegmentsOk returns a tuple with the MonitoredSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetMonitoredSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.MonitoredSegments) {
		return nil, false
	}
	return o.MonitoredSegments, true
}

// HasMonitoredSegments returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasMonitoredSegments() bool {
	if o != nil && !IsNil(o.MonitoredSegments) {
		return true
	}

	return false
}

// SetMonitoredSegments gets a reference to the given []string and assigns it to the MonitoredSegments field.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetMonitoredSegments(v []string) {
	o.MonitoredSegments = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetName(v string) {
	o.Name = &v
}

// GetSampleMode returns the SampleMode field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSampleMode() string {
	if o == nil || IsNil(o.SampleMode) {
		var ret string
		return ret
	}
	return *o.SampleMode
}

// GetSampleModeOk returns a tuple with the SampleMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSampleModeOk() (*string, bool) {
	if o == nil || IsNil(o.SampleMode) {
		return nil, false
	}
	return o.SampleMode, true
}

// HasSampleMode returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasSampleMode() bool {
	if o != nil && !IsNil(o.SampleMode) {
		return true
	}

	return false
}

// SetSampleMode gets a reference to the given string and assigns it to the SampleMode field.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetSampleMode(v string) {
	o.SampleMode = &v
}

// GetSampleRate returns the SampleRate field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSampleRate() int32 {
	if o == nil || IsNil(o.SampleRate) {
		var ret int32
		return ret
	}
	return *o.SampleRate
}

// GetSampleRateOk returns a tuple with the SampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSampleRateOk() (*int32, bool) {
	if o == nil || IsNil(o.SampleRate) {
		return nil, false
	}
	return o.SampleRate, true
}

// HasSampleRate returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasSampleRate() bool {
	if o != nil && !IsNil(o.SampleRate) {
		return true
	}

	return false
}

// SetSampleRate gets a reference to the given int32 and assigns it to the SampleRate field.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetSampleRate(v int32) {
	o.SampleRate = &v
}

// GetSourceInterfaceName returns the SourceInterfaceName field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSourceInterfaceName() string {
	if o == nil || IsNil(o.SourceInterfaceName) {
		var ret string
		return ret
	}
	return *o.SourceInterfaceName
}

// GetSourceInterfaceNameOk returns a tuple with the SourceInterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSourceInterfaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceInterfaceName) {
		return nil, false
	}
	return o.SourceInterfaceName, true
}

// HasSourceInterfaceName returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasSourceInterfaceName() bool {
	if o != nil && !IsNil(o.SourceInterfaceName) {
		return true
	}

	return false
}

// SetSourceInterfaceName gets a reference to the given string and assigns it to the SourceInterfaceName field.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetSourceInterfaceName(v string) {
	o.SourceInterfaceName = &v
}

// GetVrfId returns the VrfId field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetVrfId() int64 {
	if o == nil || IsNil(o.VrfId) {
		var ret int64
		return ret
	}
	return *o.VrfId
}

// GetVrfIdOk returns a tuple with the VrfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetVrfIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VrfId) {
		return nil, false
	}
	return o.VrfId, true
}

// HasVrfId returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasVrfId() bool {
	if o != nil && !IsNil(o.VrfId) {
		return true
	}

	return false
}

// SetVrfId gets a reference to the given int64 and assigns it to the VrfId field.
func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetVrfId(v int64) {
	o.VrfId = &v
}

func (o V1GlobalConfigPatchRequestIpfixExportersValueExporter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestIpfixExportersValueExporter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DestinationAddress) {
		toSerialize["destinationAddress"] = o.DestinationAddress
	}
	if !IsNil(o.DestinationPort) {
		toSerialize["destinationPort"] = o.DestinationPort
	}
	if !IsNil(o.GlobalId) {
		toSerialize["globalId"] = o.GlobalId
	}
	if !IsNil(o.IsGlobalSync) {
		toSerialize["isGlobalSync"] = o.IsGlobalSync
	}
	if !IsNil(o.MonitoredSegments) {
		toSerialize["monitoredSegments"] = o.MonitoredSegments
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SampleMode) {
		toSerialize["sampleMode"] = o.SampleMode
	}
	if !IsNil(o.SampleRate) {
		toSerialize["sampleRate"] = o.SampleRate
	}
	if !IsNil(o.SourceInterfaceName) {
		toSerialize["sourceInterfaceName"] = o.SourceInterfaceName
	}
	if !IsNil(o.VrfId) {
		toSerialize["vrfId"] = o.VrfId
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestIpfixExportersValueExporter struct {
	value *V1GlobalConfigPatchRequestIpfixExportersValueExporter
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestIpfixExportersValueExporter) Get() *V1GlobalConfigPatchRequestIpfixExportersValueExporter {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestIpfixExportersValueExporter) Set(val *V1GlobalConfigPatchRequestIpfixExportersValueExporter) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestIpfixExportersValueExporter) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestIpfixExportersValueExporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestIpfixExportersValueExporter(val *V1GlobalConfigPatchRequestIpfixExportersValueExporter) *NullableV1GlobalConfigPatchRequestIpfixExportersValueExporter {
	return &NullableV1GlobalConfigPatchRequestIpfixExportersValueExporter{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestIpfixExportersValueExporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestIpfixExportersValueExporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


