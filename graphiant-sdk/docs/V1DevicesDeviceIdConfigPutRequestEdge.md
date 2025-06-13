# V1DevicesDeviceIdConfigPutRequestEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpEnabled** | Pointer to **bool** |  | [optional] 
**BgpInstance** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreBgpInstance**](V1DevicesDeviceIdConfigPutRequestCoreBgpInstance.md) |  | [optional] 
**Circuits** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue**](V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue.md) |  | [optional] 
**DhcpServerEnabled** | Pointer to **bool** |  | [optional] 
**Dns** | Pointer to [**V1DevicesDeviceIdConfigPutRequestEdgeDns**](V1DevicesDeviceIdConfigPutRequestEdgeDns.md) |  | [optional] 
**Interfaces** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue**](V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue.md) |  | [optional] 
**IpfixEnabled** | Pointer to **bool** |  | [optional] 
**IpfixExporters** | Pointer to [**map[string]V1GlobalConfigPatchRequestIpfixExportersValue**](V1GlobalConfigPatchRequestIpfixExportersValue.md) |  | [optional] 
**LldpEnabled** | Pointer to **bool** |  | [optional] 
**LocalRouteTag** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag.md) |  | [optional] 
**LocalWebServerPassword** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation.md) |  | [optional] 
**MaintenanceMode** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NatPolicy** | Pointer to [**V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy**](V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy.md) |  | [optional] 
**Ospfv2Enabled** | Pointer to **bool** |  | [optional] 
**Ospfv3Enabled** | Pointer to **bool** |  | [optional] 
**PrefixSets** | Pointer to [**map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue**](V1GlobalConfigPatchRequestGlobalPrefixSetsValue.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**RoutePolicies** | Pointer to [**map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue**](V1GlobalConfigPatchRequestRoutingPoliciesValue.md) |  | [optional] 
**Segments** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrf.md) |  | [optional] 
**Site** | Pointer to [**V1SitesPostRequestSite**](V1SitesPostRequestSite.md) |  | [optional] 
**SiteToSiteVpn** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue**](V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue.md) |  | [optional] 
**Snmp** | Pointer to [**V1GlobalConfigPatchRequestSnmpsValue**](V1GlobalConfigPatchRequestSnmpsValue.md) |  | [optional] 
**SnmpGlobalObject** | Pointer to [**map[string]V1GlobalConfigPatchRequestSnmpsValue**](V1GlobalConfigPatchRequestSnmpsValue.md) |  | [optional] 
**StaticRoutesEnabled** | Pointer to **bool** |  | [optional] 
**TrafficPolicy** | Pointer to [**V1GlobalConfigPatchRequestTrafficPolicies**](V1GlobalConfigPatchRequestTrafficPolicies.md) |  | [optional] 
**VrrpEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestEdge

`func NewV1DevicesDeviceIdConfigPutRequestEdge() *V1DevicesDeviceIdConfigPutRequestEdge`

NewV1DevicesDeviceIdConfigPutRequestEdge instantiates a new V1DevicesDeviceIdConfigPutRequestEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestEdgeWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestEdgeWithDefaults() *V1DevicesDeviceIdConfigPutRequestEdge`

NewV1DevicesDeviceIdConfigPutRequestEdgeWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetBgpEnabled() bool`

GetBgpEnabled returns the BgpEnabled field if non-nil, zero value otherwise.

### GetBgpEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetBgpEnabledOk() (*bool, bool)`

GetBgpEnabledOk returns a tuple with the BgpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetBgpEnabled(v bool)`

SetBgpEnabled sets BgpEnabled field to given value.

### HasBgpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasBgpEnabled() bool`

HasBgpEnabled returns a boolean if a field has been set.

### GetBgpInstance

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetBgpInstance() V1DevicesDeviceIdConfigPutRequestCoreBgpInstance`

GetBgpInstance returns the BgpInstance field if non-nil, zero value otherwise.

### GetBgpInstanceOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetBgpInstanceOk() (*V1DevicesDeviceIdConfigPutRequestCoreBgpInstance, bool)`

GetBgpInstanceOk returns a tuple with the BgpInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpInstance

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetBgpInstance(v V1DevicesDeviceIdConfigPutRequestCoreBgpInstance)`

SetBgpInstance sets BgpInstance field to given value.

### HasBgpInstance

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasBgpInstance() bool`

HasBgpInstance returns a boolean if a field has been set.

### GetCircuits

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetCircuits() map[string]V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue`

GetCircuits returns the Circuits field if non-nil, zero value otherwise.

### GetCircuitsOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetCircuitsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue, bool)`

GetCircuitsOk returns a tuple with the Circuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuits

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetCircuits(v map[string]V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue)`

SetCircuits sets Circuits field to given value.

### HasCircuits

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasCircuits() bool`

HasCircuits returns a boolean if a field has been set.

### GetDhcpServerEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetDhcpServerEnabled() bool`

GetDhcpServerEnabled returns the DhcpServerEnabled field if non-nil, zero value otherwise.

### GetDhcpServerEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetDhcpServerEnabledOk() (*bool, bool)`

GetDhcpServerEnabledOk returns a tuple with the DhcpServerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetDhcpServerEnabled(v bool)`

SetDhcpServerEnabled sets DhcpServerEnabled field to given value.

### HasDhcpServerEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasDhcpServerEnabled() bool`

HasDhcpServerEnabled returns a boolean if a field has been set.

### GetDns

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetDns() V1DevicesDeviceIdConfigPutRequestEdgeDns`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetDnsOk() (*V1DevicesDeviceIdConfigPutRequestEdgeDns, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetDns(v V1DevicesDeviceIdConfigPutRequestEdgeDns)`

SetDns sets Dns field to given value.

### HasDns

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetInterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetInterfaces() map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetInterfacesOk() (*map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetInterfaces(v map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValue)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetIpfixEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetIpfixEnabled() bool`

GetIpfixEnabled returns the IpfixEnabled field if non-nil, zero value otherwise.

### GetIpfixEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetIpfixEnabledOk() (*bool, bool)`

GetIpfixEnabledOk returns a tuple with the IpfixEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetIpfixEnabled(v bool)`

SetIpfixEnabled sets IpfixEnabled field to given value.

### HasIpfixEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasIpfixEnabled() bool`

HasIpfixEnabled returns a boolean if a field has been set.

### GetIpfixExporters

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetIpfixExporters() map[string]V1GlobalConfigPatchRequestIpfixExportersValue`

GetIpfixExporters returns the IpfixExporters field if non-nil, zero value otherwise.

### GetIpfixExportersOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetIpfixExportersOk() (*map[string]V1GlobalConfigPatchRequestIpfixExportersValue, bool)`

GetIpfixExportersOk returns a tuple with the IpfixExporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporters

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetIpfixExporters(v map[string]V1GlobalConfigPatchRequestIpfixExportersValue)`

SetIpfixExporters sets IpfixExporters field to given value.

### HasIpfixExporters

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasIpfixExporters() bool`

HasIpfixExporters returns a boolean if a field has been set.

### GetLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.

### GetLocalRouteTag

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocalRouteTag() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag`

GetLocalRouteTag returns the LocalRouteTag field if non-nil, zero value otherwise.

### GetLocalRouteTagOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocalRouteTagOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag, bool)`

GetLocalRouteTagOk returns a tuple with the LocalRouteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRouteTag

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetLocalRouteTag(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag)`

SetLocalRouteTag sets LocalRouteTag field to given value.

### HasLocalRouteTag

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasLocalRouteTag() bool`

HasLocalRouteTag returns a boolean if a field has been set.

### GetLocalWebServerPassword

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocalWebServerPassword() string`

GetLocalWebServerPassword returns the LocalWebServerPassword field if non-nil, zero value otherwise.

### GetLocalWebServerPasswordOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocalWebServerPasswordOk() (*string, bool)`

GetLocalWebServerPasswordOk returns a tuple with the LocalWebServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalWebServerPassword

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetLocalWebServerPassword(v string)`

SetLocalWebServerPassword sets LocalWebServerPassword field to given value.

### HasLocalWebServerPassword

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasLocalWebServerPassword() bool`

HasLocalWebServerPassword returns a boolean if a field has been set.

### GetLocation

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocation() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetLocationOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetLocation(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetName

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNatPolicy

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetNatPolicy() V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy`

GetNatPolicy returns the NatPolicy field if non-nil, zero value otherwise.

### GetNatPolicyOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetNatPolicyOk() (*V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy, bool)`

GetNatPolicyOk returns a tuple with the NatPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPolicy

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetNatPolicy(v V1DevicesDeviceIdConfigPutRequestEdgeNatPolicy)`

SetNatPolicy sets NatPolicy field to given value.

### HasNatPolicy

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasNatPolicy() bool`

HasNatPolicy returns a boolean if a field has been set.

### GetOspfv2Enabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetOspfv2Enabled() bool`

GetOspfv2Enabled returns the Ospfv2Enabled field if non-nil, zero value otherwise.

### GetOspfv2EnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetOspfv2EnabledOk() (*bool, bool)`

GetOspfv2EnabledOk returns a tuple with the Ospfv2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfv2Enabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetOspfv2Enabled(v bool)`

SetOspfv2Enabled sets Ospfv2Enabled field to given value.

### HasOspfv2Enabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasOspfv2Enabled() bool`

HasOspfv2Enabled returns a boolean if a field has been set.

### GetOspfv3Enabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetOspfv3Enabled() bool`

GetOspfv3Enabled returns the Ospfv3Enabled field if non-nil, zero value otherwise.

### GetOspfv3EnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetOspfv3EnabledOk() (*bool, bool)`

GetOspfv3EnabledOk returns a tuple with the Ospfv3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfv3Enabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetOspfv3Enabled(v bool)`

SetOspfv3Enabled sets Ospfv3Enabled field to given value.

### HasOspfv3Enabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasOspfv3Enabled() bool`

HasOspfv3Enabled returns a boolean if a field has been set.

### GetPrefixSets

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetPrefixSets() map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue`

GetPrefixSets returns the PrefixSets field if non-nil, zero value otherwise.

### GetPrefixSetsOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetPrefixSetsOk() (*map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue, bool)`

GetPrefixSetsOk returns a tuple with the PrefixSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSets

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetPrefixSets(v map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue)`

SetPrefixSets sets PrefixSets field to given value.

### HasPrefixSets

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasPrefixSets() bool`

HasPrefixSets returns a boolean if a field has been set.

### GetRegion

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionName

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetRoutePolicies

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRoutePolicies() map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue`

GetRoutePolicies returns the RoutePolicies field if non-nil, zero value otherwise.

### GetRoutePoliciesOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetRoutePoliciesOk() (*map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue, bool)`

GetRoutePoliciesOk returns a tuple with the RoutePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePolicies

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetRoutePolicies(v map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue)`

SetRoutePolicies sets RoutePolicies field to given value.

### HasRoutePolicies

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasRoutePolicies() bool`

HasRoutePolicies returns a boolean if a field has been set.

### GetSegments

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSegments() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSegmentsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetSegments(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetSite

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSite() V1SitesPostRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSiteOk() (*V1SitesPostRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetSite(v V1SitesPostRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSiteToSiteVpn

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSiteToSiteVpn() map[string]V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue`

GetSiteToSiteVpn returns the SiteToSiteVpn field if non-nil, zero value otherwise.

### GetSiteToSiteVpnOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSiteToSiteVpnOk() (*map[string]V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue, bool)`

GetSiteToSiteVpnOk returns a tuple with the SiteToSiteVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteToSiteVpn

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetSiteToSiteVpn(v map[string]V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValue)`

SetSiteToSiteVpn sets SiteToSiteVpn field to given value.

### HasSiteToSiteVpn

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasSiteToSiteVpn() bool`

HasSiteToSiteVpn returns a boolean if a field has been set.

### GetSnmp

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSnmp() V1GlobalConfigPatchRequestSnmpsValue`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSnmpOk() (*V1GlobalConfigPatchRequestSnmpsValue, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetSnmp(v V1GlobalConfigPatchRequestSnmpsValue)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetSnmpGlobalObject

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSnmpGlobalObject() map[string]V1GlobalConfigPatchRequestSnmpsValue`

GetSnmpGlobalObject returns the SnmpGlobalObject field if non-nil, zero value otherwise.

### GetSnmpGlobalObjectOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetSnmpGlobalObjectOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValue, bool)`

GetSnmpGlobalObjectOk returns a tuple with the SnmpGlobalObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpGlobalObject

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetSnmpGlobalObject(v map[string]V1GlobalConfigPatchRequestSnmpsValue)`

SetSnmpGlobalObject sets SnmpGlobalObject field to given value.

### HasSnmpGlobalObject

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasSnmpGlobalObject() bool`

HasSnmpGlobalObject returns a boolean if a field has been set.

### GetStaticRoutesEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetStaticRoutesEnabled() bool`

GetStaticRoutesEnabled returns the StaticRoutesEnabled field if non-nil, zero value otherwise.

### GetStaticRoutesEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetStaticRoutesEnabledOk() (*bool, bool)`

GetStaticRoutesEnabledOk returns a tuple with the StaticRoutesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutesEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetStaticRoutesEnabled(v bool)`

SetStaticRoutesEnabled sets StaticRoutesEnabled field to given value.

### HasStaticRoutesEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasStaticRoutesEnabled() bool`

HasStaticRoutesEnabled returns a boolean if a field has been set.

### GetTrafficPolicy

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetTrafficPolicy() V1GlobalConfigPatchRequestTrafficPolicies`

GetTrafficPolicy returns the TrafficPolicy field if non-nil, zero value otherwise.

### GetTrafficPolicyOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetTrafficPolicyOk() (*V1GlobalConfigPatchRequestTrafficPolicies, bool)`

GetTrafficPolicyOk returns a tuple with the TrafficPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPolicy

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetTrafficPolicy(v V1GlobalConfigPatchRequestTrafficPolicies)`

SetTrafficPolicy sets TrafficPolicy field to given value.

### HasTrafficPolicy

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasTrafficPolicy() bool`

HasTrafficPolicy returns a boolean if a field has been set.

### GetVrrpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetVrrpEnabled() bool`

GetVrrpEnabled returns the VrrpEnabled field if non-nil, zero value otherwise.

### GetVrrpEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) GetVrrpEnabledOk() (*bool, bool)`

GetVrrpEnabledOk returns a tuple with the VrrpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) SetVrrpEnabled(v bool)`

SetVrrpEnabled sets VrrpEnabled field to given value.

### HasVrrpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdge) HasVrrpEnabled() bool`

HasVrrpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


