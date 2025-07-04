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

// checks if the V1DevicesDeviceIdConfigPutRequestEdge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DevicesDeviceIdConfigPutRequestEdge{}

// V1DevicesDeviceIdConfigPutRequestEdge struct for V1DevicesDeviceIdConfigPutRequestEdge
type V1DevicesDeviceIdConfigPutRequestEdge struct {
	BgpEnabled *bool `json:"bgpEnabled,omitempty"`
	BgpInstance *V1DevicesDeviceIdConfigPutRequestCoreBgpInstance `json:"bgpInstance,omitempty"`
	Circuits *map[string]V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue `json:"circuits,omitempty"`
	DhcpServerEnabled *bool `json:"dhcpServerEnabled,omitempty"`
	Dns *V1DevicesDeviceIdConfigPutRequestEdgeDns `json:"dns,omitempty"`
	Interfaces *map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue `json:"interfaces,omitempty"`
	IpfixEnabled *bool `json:"ipfixEnabled,omitempty"`
	IpfixExporters *map[string]V1GlobalConfigPatchRequestIpfixExportersValue `json:"ipfixExporters,omitempty"`
	LldpEnabled *bool `json:"lldpEnabled,omitempty"`
	LocalRouteTag *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag `json:"localRouteTag,omitempty"`
	LocalWebServerPassword *string `json:"localWebServerPassword,omitempty"`
	Location *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation `json:"location,omitempty"`
	MaintenanceMode *bool `json:"maintenanceMode,omitempty"`
	Name *string `json:"name,omitempty"`
	NatPolicy *V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy `json:"natPolicy,omitempty"`
	Ospfv2Enabled *bool `json:"ospfv2Enabled,omitempty"`
	Ospfv3Enabled *bool `json:"ospfv3Enabled,omitempty"`
	PrefixSets *map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue `json:"prefixSets,omitempty"`
	Region *string `json:"region,omitempty"`
	RegionName *string `json:"regionName,omitempty"`
	RoutePolicies *map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue `json:"routePolicies,omitempty"`
	Segments *map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf `json:"segments,omitempty"`
	Site *V1SitesPostRequestSite `json:"site,omitempty"`
	SiteToSiteVpn *map[string]V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue `json:"siteToSiteVpn,omitempty"`
	Snmp *V1GlobalConfigPatchRequestSnmpsValue `json:"snmp,omitempty"`
	SnmpGlobalObject *map[string]V1GlobalConfigPatchRequestSnmpsValue `json:"snmpGlobalObject,omitempty"`
	StaticRoutesEnabled *bool `json:"staticRoutesEnabled,omitempty"`
	TrafficPolicy *V1GlobalConfigPatchRequestTrafficPolicies `json:"trafficPolicy,omitempty"`
	VrrpEnabled *bool `json:"vrrpEnabled,omitempty"`
}

// NewV1DevicesDeviceIdConfigPutRequestEdge instantiates a new V1DevicesDeviceIdConfigPutRequestEdge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DevicesDeviceIdConfigPutRequestEdge() *V1DevicesDeviceIdConfigPutRequestEdge {
	this := V1DevicesDeviceIdConfigPutRequestEdge{}
	return &this
}

// NewV1DevicesDeviceIdConfigPutRequestEdgeWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestEdge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DevicesDeviceIdConfigPutRequestEdgeWithDefaults() *V1DevicesDeviceIdConfigPutRequestEdge {
	this := V1DevicesDeviceIdConfigPutRequestEdge{}
	return &this
}

// GetBgpEnabled returns the BgpEnabled field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetBgpEnabled() bool {
	if o == nil || IsNil(o.BgpEnabled) {
		var ret bool
		return ret
	}
	return *o.BgpEnabled
}

// GetBgpEnabledOk returns a tuple with the BgpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetBgpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BgpEnabled) {
		return nil, false
	}
	return o.BgpEnabled, true
}

// HasBgpEnabled returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasBgpEnabled() bool {
	if o != nil && !IsNil(o.BgpEnabled) {
		return true
	}

	return false
}

// SetBgpEnabled gets a reference to the given bool and assigns it to the BgpEnabled field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetBgpEnabled(v bool) {
	o.BgpEnabled = &v
}

// GetBgpInstance returns the BgpInstance field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetBgpInstance() V1DevicesDeviceIdConfigPutRequestCoreBgpInstance {
	if o == nil || IsNil(o.BgpInstance) {
		var ret V1DevicesDeviceIdConfigPutRequestCoreBgpInstance
		return ret
	}
	return *o.BgpInstance
}

// GetBgpInstanceOk returns a tuple with the BgpInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetBgpInstanceOk() (*V1DevicesDeviceIdConfigPutRequestCoreBgpInstance, bool) {
	if o == nil || IsNil(o.BgpInstance) {
		return nil, false
	}
	return o.BgpInstance, true
}

// HasBgpInstance returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasBgpInstance() bool {
	if o != nil && !IsNil(o.BgpInstance) {
		return true
	}

	return false
}

// SetBgpInstance gets a reference to the given V1DevicesDeviceIdConfigPutRequestCoreBgpInstance and assigns it to the BgpInstance field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetBgpInstance(v V1DevicesDeviceIdConfigPutRequestCoreBgpInstance) {
	o.BgpInstance = &v
}

// GetCircuits returns the Circuits field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetCircuits() map[string]V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue {
	if o == nil || IsNil(o.Circuits) {
		var ret map[string]V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue
		return ret
	}
	return *o.Circuits
}

// GetCircuitsOk returns a tuple with the Circuits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetCircuitsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue, bool) {
	if o == nil || IsNil(o.Circuits) {
		return nil, false
	}
	return o.Circuits, true
}

// HasCircuits returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasCircuits() bool {
	if o != nil && !IsNil(o.Circuits) {
		return true
	}

	return false
}

// SetCircuits gets a reference to the given map[string]V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue and assigns it to the Circuits field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetCircuits(v map[string]V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) {
	o.Circuits = &v
}

// GetDhcpServerEnabled returns the DhcpServerEnabled field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetDhcpServerEnabled() bool {
	if o == nil || IsNil(o.DhcpServerEnabled) {
		var ret bool
		return ret
	}
	return *o.DhcpServerEnabled
}

// GetDhcpServerEnabledOk returns a tuple with the DhcpServerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetDhcpServerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DhcpServerEnabled) {
		return nil, false
	}
	return o.DhcpServerEnabled, true
}

// HasDhcpServerEnabled returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasDhcpServerEnabled() bool {
	if o != nil && !IsNil(o.DhcpServerEnabled) {
		return true
	}

	return false
}

// SetDhcpServerEnabled gets a reference to the given bool and assigns it to the DhcpServerEnabled field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetDhcpServerEnabled(v bool) {
	o.DhcpServerEnabled = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetDns() V1DevicesDeviceIdConfigPutRequestEdgeDns {
	if o == nil || IsNil(o.Dns) {
		var ret V1DevicesDeviceIdConfigPutRequestEdgeDns
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetDnsOk() (*V1DevicesDeviceIdConfigPutRequestEdgeDns, bool) {
	if o == nil || IsNil(o.Dns) {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasDns() bool {
	if o != nil && !IsNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given V1DevicesDeviceIdConfigPutRequestEdgeDns and assigns it to the Dns field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetDns(v V1DevicesDeviceIdConfigPutRequestEdgeDns) {
	o.Dns = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetInterfaces() map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue {
	if o == nil || IsNil(o.Interfaces) {
		var ret map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue
		return ret
	}
	return *o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetInterfacesOk() (*map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue, bool) {
	if o == nil || IsNil(o.Interfaces) {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasInterfaces() bool {
	if o != nil && !IsNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue and assigns it to the Interfaces field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetInterfaces(v map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue) {
	o.Interfaces = &v
}

// GetIpfixEnabled returns the IpfixEnabled field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetIpfixEnabled() bool {
	if o == nil || IsNil(o.IpfixEnabled) {
		var ret bool
		return ret
	}
	return *o.IpfixEnabled
}

// GetIpfixEnabledOk returns a tuple with the IpfixEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetIpfixEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IpfixEnabled) {
		return nil, false
	}
	return o.IpfixEnabled, true
}

// HasIpfixEnabled returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasIpfixEnabled() bool {
	if o != nil && !IsNil(o.IpfixEnabled) {
		return true
	}

	return false
}

// SetIpfixEnabled gets a reference to the given bool and assigns it to the IpfixEnabled field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetIpfixEnabled(v bool) {
	o.IpfixEnabled = &v
}

// GetIpfixExporters returns the IpfixExporters field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetIpfixExporters() map[string]V1GlobalConfigPatchRequestIpfixExportersValue {
	if o == nil || IsNil(o.IpfixExporters) {
		var ret map[string]V1GlobalConfigPatchRequestIpfixExportersValue
		return ret
	}
	return *o.IpfixExporters
}

// GetIpfixExportersOk returns a tuple with the IpfixExporters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetIpfixExportersOk() (*map[string]V1GlobalConfigPatchRequestIpfixExportersValue, bool) {
	if o == nil || IsNil(o.IpfixExporters) {
		return nil, false
	}
	return o.IpfixExporters, true
}

// HasIpfixExporters returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasIpfixExporters() bool {
	if o != nil && !IsNil(o.IpfixExporters) {
		return true
	}

	return false
}

// SetIpfixExporters gets a reference to the given map[string]V1GlobalConfigPatchRequestIpfixExportersValue and assigns it to the IpfixExporters field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetIpfixExporters(v map[string]V1GlobalConfigPatchRequestIpfixExportersValue) {
	o.IpfixExporters = &v
}

// GetLldpEnabled returns the LldpEnabled field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLldpEnabled() bool {
	if o == nil || IsNil(o.LldpEnabled) {
		var ret bool
		return ret
	}
	return *o.LldpEnabled
}

// GetLldpEnabledOk returns a tuple with the LldpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLldpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LldpEnabled) {
		return nil, false
	}
	return o.LldpEnabled, true
}

// HasLldpEnabled returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasLldpEnabled() bool {
	if o != nil && !IsNil(o.LldpEnabled) {
		return true
	}

	return false
}

// SetLldpEnabled gets a reference to the given bool and assigns it to the LldpEnabled field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetLldpEnabled(v bool) {
	o.LldpEnabled = &v
}

// GetLocalRouteTag returns the LocalRouteTag field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocalRouteTag() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag {
	if o == nil || IsNil(o.LocalRouteTag) {
		var ret V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag
		return ret
	}
	return *o.LocalRouteTag
}

// GetLocalRouteTagOk returns a tuple with the LocalRouteTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocalRouteTagOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag, bool) {
	if o == nil || IsNil(o.LocalRouteTag) {
		return nil, false
	}
	return o.LocalRouteTag, true
}

// HasLocalRouteTag returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasLocalRouteTag() bool {
	if o != nil && !IsNil(o.LocalRouteTag) {
		return true
	}

	return false
}

// SetLocalRouteTag gets a reference to the given V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag and assigns it to the LocalRouteTag field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetLocalRouteTag(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag) {
	o.LocalRouteTag = &v
}

// GetLocalWebServerPassword returns the LocalWebServerPassword field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocalWebServerPassword() string {
	if o == nil || IsNil(o.LocalWebServerPassword) {
		var ret string
		return ret
	}
	return *o.LocalWebServerPassword
}

// GetLocalWebServerPasswordOk returns a tuple with the LocalWebServerPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocalWebServerPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.LocalWebServerPassword) {
		return nil, false
	}
	return o.LocalWebServerPassword, true
}

// HasLocalWebServerPassword returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasLocalWebServerPassword() bool {
	if o != nil && !IsNil(o.LocalWebServerPassword) {
		return true
	}

	return false
}

// SetLocalWebServerPassword gets a reference to the given string and assigns it to the LocalWebServerPassword field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetLocalWebServerPassword(v string) {
	o.LocalWebServerPassword = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocation() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation {
	if o == nil || IsNil(o.Location) {
		var ret V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocationOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation and assigns it to the Location field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetLocation(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation) {
	o.Location = &v
}

// GetMaintenanceMode returns the MaintenanceMode field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetMaintenanceMode() bool {
	if o == nil || IsNil(o.MaintenanceMode) {
		var ret bool
		return ret
	}
	return *o.MaintenanceMode
}

// GetMaintenanceModeOk returns a tuple with the MaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetMaintenanceModeOk() (*bool, bool) {
	if o == nil || IsNil(o.MaintenanceMode) {
		return nil, false
	}
	return o.MaintenanceMode, true
}

// HasMaintenanceMode returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasMaintenanceMode() bool {
	if o != nil && !IsNil(o.MaintenanceMode) {
		return true
	}

	return false
}

// SetMaintenanceMode gets a reference to the given bool and assigns it to the MaintenanceMode field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetMaintenanceMode(v bool) {
	o.MaintenanceMode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetName(v string) {
	o.Name = &v
}

// GetNatPolicy returns the NatPolicy field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetNatPolicy() V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy {
	if o == nil || IsNil(o.NatPolicy) {
		var ret V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy
		return ret
	}
	return *o.NatPolicy
}

// GetNatPolicyOk returns a tuple with the NatPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetNatPolicyOk() (*V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy, bool) {
	if o == nil || IsNil(o.NatPolicy) {
		return nil, false
	}
	return o.NatPolicy, true
}

// HasNatPolicy returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasNatPolicy() bool {
	if o != nil && !IsNil(o.NatPolicy) {
		return true
	}

	return false
}

// SetNatPolicy gets a reference to the given V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy and assigns it to the NatPolicy field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetNatPolicy(v V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy) {
	o.NatPolicy = &v
}

// GetOspfv2Enabled returns the Ospfv2Enabled field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetOspfv2Enabled() bool {
	if o == nil || IsNil(o.Ospfv2Enabled) {
		var ret bool
		return ret
	}
	return *o.Ospfv2Enabled
}

// GetOspfv2EnabledOk returns a tuple with the Ospfv2Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetOspfv2EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Ospfv2Enabled) {
		return nil, false
	}
	return o.Ospfv2Enabled, true
}

// HasOspfv2Enabled returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasOspfv2Enabled() bool {
	if o != nil && !IsNil(o.Ospfv2Enabled) {
		return true
	}

	return false
}

// SetOspfv2Enabled gets a reference to the given bool and assigns it to the Ospfv2Enabled field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetOspfv2Enabled(v bool) {
	o.Ospfv2Enabled = &v
}

// GetOspfv3Enabled returns the Ospfv3Enabled field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetOspfv3Enabled() bool {
	if o == nil || IsNil(o.Ospfv3Enabled) {
		var ret bool
		return ret
	}
	return *o.Ospfv3Enabled
}

// GetOspfv3EnabledOk returns a tuple with the Ospfv3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetOspfv3EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Ospfv3Enabled) {
		return nil, false
	}
	return o.Ospfv3Enabled, true
}

// HasOspfv3Enabled returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasOspfv3Enabled() bool {
	if o != nil && !IsNil(o.Ospfv3Enabled) {
		return true
	}

	return false
}

// SetOspfv3Enabled gets a reference to the given bool and assigns it to the Ospfv3Enabled field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetOspfv3Enabled(v bool) {
	o.Ospfv3Enabled = &v
}

// GetPrefixSets returns the PrefixSets field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetPrefixSets() map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue {
	if o == nil || IsNil(o.PrefixSets) {
		var ret map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue
		return ret
	}
	return *o.PrefixSets
}

// GetPrefixSetsOk returns a tuple with the PrefixSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetPrefixSetsOk() (*map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue, bool) {
	if o == nil || IsNil(o.PrefixSets) {
		return nil, false
	}
	return o.PrefixSets, true
}

// HasPrefixSets returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasPrefixSets() bool {
	if o != nil && !IsNil(o.PrefixSets) {
		return true
	}

	return false
}

// SetPrefixSets gets a reference to the given map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue and assigns it to the PrefixSets field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetPrefixSets(v map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue) {
	o.PrefixSets = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetRegion(v string) {
	o.Region = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetRegionName(v string) {
	o.RegionName = &v
}

// GetRoutePolicies returns the RoutePolicies field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRoutePolicies() map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue {
	if o == nil || IsNil(o.RoutePolicies) {
		var ret map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue
		return ret
	}
	return *o.RoutePolicies
}

// GetRoutePoliciesOk returns a tuple with the RoutePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRoutePoliciesOk() (*map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue, bool) {
	if o == nil || IsNil(o.RoutePolicies) {
		return nil, false
	}
	return o.RoutePolicies, true
}

// HasRoutePolicies returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasRoutePolicies() bool {
	if o != nil && !IsNil(o.RoutePolicies) {
		return true
	}

	return false
}

// SetRoutePolicies gets a reference to the given map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue and assigns it to the RoutePolicies field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetRoutePolicies(v map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue) {
	o.RoutePolicies = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSegments() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf {
	if o == nil || IsNil(o.Segments) {
		var ret map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf
		return ret
	}
	return *o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSegmentsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf and assigns it to the Segments field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetSegments(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) {
	o.Segments = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSite() V1SitesPostRequestSite {
	if o == nil || IsNil(o.Site) {
		var ret V1SitesPostRequestSite
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSiteOk() (*V1SitesPostRequestSite, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given V1SitesPostRequestSite and assigns it to the Site field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetSite(v V1SitesPostRequestSite) {
	o.Site = &v
}

// GetSiteToSiteVpn returns the SiteToSiteVpn field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSiteToSiteVpn() map[string]V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue {
	if o == nil || IsNil(o.SiteToSiteVpn) {
		var ret map[string]V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue
		return ret
	}
	return *o.SiteToSiteVpn
}

// GetSiteToSiteVpnOk returns a tuple with the SiteToSiteVpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSiteToSiteVpnOk() (*map[string]V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue, bool) {
	if o == nil || IsNil(o.SiteToSiteVpn) {
		return nil, false
	}
	return o.SiteToSiteVpn, true
}

// HasSiteToSiteVpn returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasSiteToSiteVpn() bool {
	if o != nil && !IsNil(o.SiteToSiteVpn) {
		return true
	}

	return false
}

// SetSiteToSiteVpn gets a reference to the given map[string]V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue and assigns it to the SiteToSiteVpn field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetSiteToSiteVpn(v map[string]V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue) {
	o.SiteToSiteVpn = &v
}

// GetSnmp returns the Snmp field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSnmp() V1GlobalConfigPatchRequestSnmpsValue {
	if o == nil || IsNil(o.Snmp) {
		var ret V1GlobalConfigPatchRequestSnmpsValue
		return ret
	}
	return *o.Snmp
}

// GetSnmpOk returns a tuple with the Snmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSnmpOk() (*V1GlobalConfigPatchRequestSnmpsValue, bool) {
	if o == nil || IsNil(o.Snmp) {
		return nil, false
	}
	return o.Snmp, true
}

// HasSnmp returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasSnmp() bool {
	if o != nil && !IsNil(o.Snmp) {
		return true
	}

	return false
}

// SetSnmp gets a reference to the given V1GlobalConfigPatchRequestSnmpsValue and assigns it to the Snmp field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetSnmp(v V1GlobalConfigPatchRequestSnmpsValue) {
	o.Snmp = &v
}

// GetSnmpGlobalObject returns the SnmpGlobalObject field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSnmpGlobalObject() map[string]V1GlobalConfigPatchRequestSnmpsValue {
	if o == nil || IsNil(o.SnmpGlobalObject) {
		var ret map[string]V1GlobalConfigPatchRequestSnmpsValue
		return ret
	}
	return *o.SnmpGlobalObject
}

// GetSnmpGlobalObjectOk returns a tuple with the SnmpGlobalObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSnmpGlobalObjectOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValue, bool) {
	if o == nil || IsNil(o.SnmpGlobalObject) {
		return nil, false
	}
	return o.SnmpGlobalObject, true
}

// HasSnmpGlobalObject returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasSnmpGlobalObject() bool {
	if o != nil && !IsNil(o.SnmpGlobalObject) {
		return true
	}

	return false
}

// SetSnmpGlobalObject gets a reference to the given map[string]V1GlobalConfigPatchRequestSnmpsValue and assigns it to the SnmpGlobalObject field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetSnmpGlobalObject(v map[string]V1GlobalConfigPatchRequestSnmpsValue) {
	o.SnmpGlobalObject = &v
}

// GetStaticRoutesEnabled returns the StaticRoutesEnabled field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetStaticRoutesEnabled() bool {
	if o == nil || IsNil(o.StaticRoutesEnabled) {
		var ret bool
		return ret
	}
	return *o.StaticRoutesEnabled
}

// GetStaticRoutesEnabledOk returns a tuple with the StaticRoutesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetStaticRoutesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StaticRoutesEnabled) {
		return nil, false
	}
	return o.StaticRoutesEnabled, true
}

// HasStaticRoutesEnabled returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasStaticRoutesEnabled() bool {
	if o != nil && !IsNil(o.StaticRoutesEnabled) {
		return true
	}

	return false
}

// SetStaticRoutesEnabled gets a reference to the given bool and assigns it to the StaticRoutesEnabled field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetStaticRoutesEnabled(v bool) {
	o.StaticRoutesEnabled = &v
}

// GetTrafficPolicy returns the TrafficPolicy field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetTrafficPolicy() V1GlobalConfigPatchRequestTrafficPolicies {
	if o == nil || IsNil(o.TrafficPolicy) {
		var ret V1GlobalConfigPatchRequestTrafficPolicies
		return ret
	}
	return *o.TrafficPolicy
}

// GetTrafficPolicyOk returns a tuple with the TrafficPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetTrafficPolicyOk() (*V1GlobalConfigPatchRequestTrafficPolicies, bool) {
	if o == nil || IsNil(o.TrafficPolicy) {
		return nil, false
	}
	return o.TrafficPolicy, true
}

// HasTrafficPolicy returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasTrafficPolicy() bool {
	if o != nil && !IsNil(o.TrafficPolicy) {
		return true
	}

	return false
}

// SetTrafficPolicy gets a reference to the given V1GlobalConfigPatchRequestTrafficPolicies and assigns it to the TrafficPolicy field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetTrafficPolicy(v V1GlobalConfigPatchRequestTrafficPolicies) {
	o.TrafficPolicy = &v
}

// GetVrrpEnabled returns the VrrpEnabled field value if set, zero value otherwise.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetVrrpEnabled() bool {
	if o == nil || IsNil(o.VrrpEnabled) {
		var ret bool
		return ret
	}
	return *o.VrrpEnabled
}

// GetVrrpEnabledOk returns a tuple with the VrrpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetVrrpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.VrrpEnabled) {
		return nil, false
	}
	return o.VrrpEnabled, true
}

// HasVrrpEnabled returns a boolean if a field has been set.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasVrrpEnabled() bool {
	if o != nil && !IsNil(o.VrrpEnabled) {
		return true
	}

	return false
}

// SetVrrpEnabled gets a reference to the given bool and assigns it to the VrrpEnabled field.
func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetVrrpEnabled(v bool) {
	o.VrrpEnabled = &v
}

func (o V1DevicesDeviceIdConfigPutRequestEdge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DevicesDeviceIdConfigPutRequestEdge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BgpEnabled) {
		toSerialize["bgpEnabled"] = o.BgpEnabled
	}
	if !IsNil(o.BgpInstance) {
		toSerialize["bgpInstance"] = o.BgpInstance
	}
	if !IsNil(o.Circuits) {
		toSerialize["circuits"] = o.Circuits
	}
	if !IsNil(o.DhcpServerEnabled) {
		toSerialize["dhcpServerEnabled"] = o.DhcpServerEnabled
	}
	if !IsNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !IsNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	if !IsNil(o.IpfixEnabled) {
		toSerialize["ipfixEnabled"] = o.IpfixEnabled
	}
	if !IsNil(o.IpfixExporters) {
		toSerialize["ipfixExporters"] = o.IpfixExporters
	}
	if !IsNil(o.LldpEnabled) {
		toSerialize["lldpEnabled"] = o.LldpEnabled
	}
	if !IsNil(o.LocalRouteTag) {
		toSerialize["localRouteTag"] = o.LocalRouteTag
	}
	if !IsNil(o.LocalWebServerPassword) {
		toSerialize["localWebServerPassword"] = o.LocalWebServerPassword
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.MaintenanceMode) {
		toSerialize["maintenanceMode"] = o.MaintenanceMode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NatPolicy) {
		toSerialize["natPolicy"] = o.NatPolicy
	}
	if !IsNil(o.Ospfv2Enabled) {
		toSerialize["ospfv2Enabled"] = o.Ospfv2Enabled
	}
	if !IsNil(o.Ospfv3Enabled) {
		toSerialize["ospfv3Enabled"] = o.Ospfv3Enabled
	}
	if !IsNil(o.PrefixSets) {
		toSerialize["prefixSets"] = o.PrefixSets
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !IsNil(o.RoutePolicies) {
		toSerialize["routePolicies"] = o.RoutePolicies
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.SiteToSiteVpn) {
		toSerialize["siteToSiteVpn"] = o.SiteToSiteVpn
	}
	if !IsNil(o.Snmp) {
		toSerialize["snmp"] = o.Snmp
	}
	if !IsNil(o.SnmpGlobalObject) {
		toSerialize["snmpGlobalObject"] = o.SnmpGlobalObject
	}
	if !IsNil(o.StaticRoutesEnabled) {
		toSerialize["staticRoutesEnabled"] = o.StaticRoutesEnabled
	}
	if !IsNil(o.TrafficPolicy) {
		toSerialize["trafficPolicy"] = o.TrafficPolicy
	}
	if !IsNil(o.VrrpEnabled) {
		toSerialize["vrrpEnabled"] = o.VrrpEnabled
	}
	return toSerialize, nil
}

type NullableV1DevicesDeviceIdConfigPutRequestEdge struct {
	value *V1DevicesDeviceIdConfigPutRequestEdge
	isSet bool
}

func (v NullableV1DevicesDeviceIdConfigPutRequestEdge) Get() *V1DevicesDeviceIdConfigPutRequestEdge {
	return v.value
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestEdge) Set(val *V1DevicesDeviceIdConfigPutRequestEdge) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DevicesDeviceIdConfigPutRequestEdge) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestEdge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DevicesDeviceIdConfigPutRequestEdge(val *V1DevicesDeviceIdConfigPutRequestEdge) *NullableV1DevicesDeviceIdConfigPutRequestEdge {
	return &NullableV1DevicesDeviceIdConfigPutRequestEdge{value: val, isSet: true}
}

func (v NullableV1DevicesDeviceIdConfigPutRequestEdge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DevicesDeviceIdConfigPutRequestEdge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


