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

// checks if the V2AssuranceTopologyInventoryPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2AssuranceTopologyInventoryPost200Response{}

// V2AssuranceTopologyInventoryPost200Response struct for V2AssuranceTopologyInventoryPost200Response
type V2AssuranceTopologyInventoryPost200Response struct {
	AppNames []string `json:"appNames,omitempty"`
	ClientSites []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite `json:"clientSites,omitempty"`
	Regions []V2AssuranceTopologyInventoryPost200ResponseRegionsInner `json:"regions,omitempty"`
	ServerSites []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite `json:"serverSites,omitempty"`
	TopologyChangeTs []V1AlarmHistoryGet200ResponseHistoryInnerTime `json:"topologyChangeTs,omitempty"`
}

// NewV2AssuranceTopologyInventoryPost200Response instantiates a new V2AssuranceTopologyInventoryPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssuranceTopologyInventoryPost200Response() *V2AssuranceTopologyInventoryPost200Response {
	this := V2AssuranceTopologyInventoryPost200Response{}
	return &this
}

// NewV2AssuranceTopologyInventoryPost200ResponseWithDefaults instantiates a new V2AssuranceTopologyInventoryPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssuranceTopologyInventoryPost200ResponseWithDefaults() *V2AssuranceTopologyInventoryPost200Response {
	this := V2AssuranceTopologyInventoryPost200Response{}
	return &this
}

// GetAppNames returns the AppNames field value if set, zero value otherwise.
func (o *V2AssuranceTopologyInventoryPost200Response) GetAppNames() []string {
	if o == nil || IsNil(o.AppNames) {
		var ret []string
		return ret
	}
	return o.AppNames
}

// GetAppNamesOk returns a tuple with the AppNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyInventoryPost200Response) GetAppNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AppNames) {
		return nil, false
	}
	return o.AppNames, true
}

// HasAppNames returns a boolean if a field has been set.
func (o *V2AssuranceTopologyInventoryPost200Response) HasAppNames() bool {
	if o != nil && !IsNil(o.AppNames) {
		return true
	}

	return false
}

// SetAppNames gets a reference to the given []string and assigns it to the AppNames field.
func (o *V2AssuranceTopologyInventoryPost200Response) SetAppNames(v []string) {
	o.AppNames = v
}

// GetClientSites returns the ClientSites field value if set, zero value otherwise.
func (o *V2AssuranceTopologyInventoryPost200Response) GetClientSites() []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite {
	if o == nil || IsNil(o.ClientSites) {
		var ret []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite
		return ret
	}
	return o.ClientSites
}

// GetClientSitesOk returns a tuple with the ClientSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyInventoryPost200Response) GetClientSitesOk() ([]V2AssuranceFlowSummaryPost200ResponseClientEndpointSite, bool) {
	if o == nil || IsNil(o.ClientSites) {
		return nil, false
	}
	return o.ClientSites, true
}

// HasClientSites returns a boolean if a field has been set.
func (o *V2AssuranceTopologyInventoryPost200Response) HasClientSites() bool {
	if o != nil && !IsNil(o.ClientSites) {
		return true
	}

	return false
}

// SetClientSites gets a reference to the given []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite and assigns it to the ClientSites field.
func (o *V2AssuranceTopologyInventoryPost200Response) SetClientSites(v []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite) {
	o.ClientSites = v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *V2AssuranceTopologyInventoryPost200Response) GetRegions() []V2AssuranceTopologyInventoryPost200ResponseRegionsInner {
	if o == nil || IsNil(o.Regions) {
		var ret []V2AssuranceTopologyInventoryPost200ResponseRegionsInner
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyInventoryPost200Response) GetRegionsOk() ([]V2AssuranceTopologyInventoryPost200ResponseRegionsInner, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *V2AssuranceTopologyInventoryPost200Response) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []V2AssuranceTopologyInventoryPost200ResponseRegionsInner and assigns it to the Regions field.
func (o *V2AssuranceTopologyInventoryPost200Response) SetRegions(v []V2AssuranceTopologyInventoryPost200ResponseRegionsInner) {
	o.Regions = v
}

// GetServerSites returns the ServerSites field value if set, zero value otherwise.
func (o *V2AssuranceTopologyInventoryPost200Response) GetServerSites() []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite {
	if o == nil || IsNil(o.ServerSites) {
		var ret []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite
		return ret
	}
	return o.ServerSites
}

// GetServerSitesOk returns a tuple with the ServerSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyInventoryPost200Response) GetServerSitesOk() ([]V2AssuranceFlowSummaryPost200ResponseClientEndpointSite, bool) {
	if o == nil || IsNil(o.ServerSites) {
		return nil, false
	}
	return o.ServerSites, true
}

// HasServerSites returns a boolean if a field has been set.
func (o *V2AssuranceTopologyInventoryPost200Response) HasServerSites() bool {
	if o != nil && !IsNil(o.ServerSites) {
		return true
	}

	return false
}

// SetServerSites gets a reference to the given []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite and assigns it to the ServerSites field.
func (o *V2AssuranceTopologyInventoryPost200Response) SetServerSites(v []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite) {
	o.ServerSites = v
}

// GetTopologyChangeTs returns the TopologyChangeTs field value if set, zero value otherwise.
func (o *V2AssuranceTopologyInventoryPost200Response) GetTopologyChangeTs() []V1AlarmHistoryGet200ResponseHistoryInnerTime {
	if o == nil || IsNil(o.TopologyChangeTs) {
		var ret []V1AlarmHistoryGet200ResponseHistoryInnerTime
		return ret
	}
	return o.TopologyChangeTs
}

// GetTopologyChangeTsOk returns a tuple with the TopologyChangeTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssuranceTopologyInventoryPost200Response) GetTopologyChangeTsOk() ([]V1AlarmHistoryGet200ResponseHistoryInnerTime, bool) {
	if o == nil || IsNil(o.TopologyChangeTs) {
		return nil, false
	}
	return o.TopologyChangeTs, true
}

// HasTopologyChangeTs returns a boolean if a field has been set.
func (o *V2AssuranceTopologyInventoryPost200Response) HasTopologyChangeTs() bool {
	if o != nil && !IsNil(o.TopologyChangeTs) {
		return true
	}

	return false
}

// SetTopologyChangeTs gets a reference to the given []V1AlarmHistoryGet200ResponseHistoryInnerTime and assigns it to the TopologyChangeTs field.
func (o *V2AssuranceTopologyInventoryPost200Response) SetTopologyChangeTs(v []V1AlarmHistoryGet200ResponseHistoryInnerTime) {
	o.TopologyChangeTs = v
}

func (o V2AssuranceTopologyInventoryPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2AssuranceTopologyInventoryPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppNames) {
		toSerialize["appNames"] = o.AppNames
	}
	if !IsNil(o.ClientSites) {
		toSerialize["clientSites"] = o.ClientSites
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !IsNil(o.ServerSites) {
		toSerialize["serverSites"] = o.ServerSites
	}
	if !IsNil(o.TopologyChangeTs) {
		toSerialize["topologyChangeTs"] = o.TopologyChangeTs
	}
	return toSerialize, nil
}

type NullableV2AssuranceTopologyInventoryPost200Response struct {
	value *V2AssuranceTopologyInventoryPost200Response
	isSet bool
}

func (v NullableV2AssuranceTopologyInventoryPost200Response) Get() *V2AssuranceTopologyInventoryPost200Response {
	return v.value
}

func (v *NullableV2AssuranceTopologyInventoryPost200Response) Set(val *V2AssuranceTopologyInventoryPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssuranceTopologyInventoryPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssuranceTopologyInventoryPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssuranceTopologyInventoryPost200Response(val *V2AssuranceTopologyInventoryPost200Response) *NullableV2AssuranceTopologyInventoryPost200Response {
	return &NullableV2AssuranceTopologyInventoryPost200Response{value: val, isSet: true}
}

func (v NullableV2AssuranceTopologyInventoryPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssuranceTopologyInventoryPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


