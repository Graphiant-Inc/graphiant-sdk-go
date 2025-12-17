# ManaV2EdgeDeviceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpEnabled** | Pointer to **bool** |  | [optional] 
**BgpInstance** | Pointer to [**ManaV2BgpInstanceConfig**](ManaV2BgpInstanceConfig.md) |  | [optional] 
**Circuits** | Pointer to [**map[string]ManaV2CircuitConfig**](ManaV2CircuitConfig.md) |  | [optional] 
**DhcpServerEnabled** | Pointer to **bool** |  | [optional] 
**Dns** | Pointer to [**ManaV2NullableDnsConfig**](ManaV2NullableDnsConfig.md) |  | [optional] 
**Interfaces** | Pointer to [**map[string]ManaV2NullableInterfaceConfig**](ManaV2NullableInterfaceConfig.md) |  | [optional] 
**IpfixEnabled** | Pointer to **bool** |  | [optional] 
**IpfixExporters** | Pointer to [**map[string]ManaV2NullableIpfixExporterConfig**](ManaV2NullableIpfixExporterConfig.md) |  | [optional] 
**LagInterfaces** | Pointer to [**map[string]ManaV2NullableLagInterfaceConfig**](ManaV2NullableLagInterfaceConfig.md) |  | [optional] 
**LldpEnabled** | Pointer to **bool** |  | [optional] 
**LocalRouteTag** | Pointer to [**ManaV2NullableRouteTagSet**](ManaV2NullableRouteTagSet.md) |  | [optional] 
**LocalWebServerPassword** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**ManaV2Location**](ManaV2Location.md) |  | [optional] 
**MaintenanceMode** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NatPolicy** | Pointer to [**ManaV2NatPolicyConfig**](ManaV2NatPolicyConfig.md) |  | [optional] 
**NtpGlobalObject** | Pointer to [**map[string]ManaV2NullableNtpConfig**](ManaV2NullableNtpConfig.md) |  | [optional] 
**Ospfv2Enabled** | Pointer to **bool** |  | [optional] 
**Ospfv3Enabled** | Pointer to **bool** |  | [optional] 
**PrefixSets** | Pointer to [**map[string]ManaV2NullablePrefixSetConfig**](ManaV2NullablePrefixSetConfig.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**RoutePolicies** | Pointer to [**map[string]ManaV2NullableRoutingPolicyConfig**](ManaV2NullableRoutingPolicyConfig.md) |  | [optional] 
**Segments** | Pointer to [**map[string]ManaV2VrfConfig**](ManaV2VrfConfig.md) |  | [optional] 
**Site** | Pointer to [**ManaV2NewSite**](ManaV2NewSite.md) |  | [optional] 
**SiteToSiteVpn** | Pointer to [**map[string]ManaV2NullableIPsecTunnelConfig**](ManaV2NullableIPsecTunnelConfig.md) |  | [optional] 
**Snmp** | Pointer to [**ManaV2NullableSnmpConfig**](ManaV2NullableSnmpConfig.md) |  | [optional] 
**SnmpGlobalObject** | Pointer to [**map[string]ManaV2NullableSnmpConfig**](ManaV2NullableSnmpConfig.md) |  | [optional] 
**StaticRoutesEnabled** | Pointer to **bool** |  | [optional] 
**TrafficPolicy** | Pointer to [**ManaV2ForwardingPolicyConfig**](ManaV2ForwardingPolicyConfig.md) |  | [optional] 
**VrrpEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewManaV2EdgeDeviceConfig

`func NewManaV2EdgeDeviceConfig() *ManaV2EdgeDeviceConfig`

NewManaV2EdgeDeviceConfig instantiates a new ManaV2EdgeDeviceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2EdgeDeviceConfigWithDefaults

`func NewManaV2EdgeDeviceConfigWithDefaults() *ManaV2EdgeDeviceConfig`

NewManaV2EdgeDeviceConfigWithDefaults instantiates a new ManaV2EdgeDeviceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpEnabled

`func (o *ManaV2EdgeDeviceConfig) GetBgpEnabled() bool`

GetBgpEnabled returns the BgpEnabled field if non-nil, zero value otherwise.

### GetBgpEnabledOk

`func (o *ManaV2EdgeDeviceConfig) GetBgpEnabledOk() (*bool, bool)`

GetBgpEnabledOk returns a tuple with the BgpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpEnabled

`func (o *ManaV2EdgeDeviceConfig) SetBgpEnabled(v bool)`

SetBgpEnabled sets BgpEnabled field to given value.

### HasBgpEnabled

`func (o *ManaV2EdgeDeviceConfig) HasBgpEnabled() bool`

HasBgpEnabled returns a boolean if a field has been set.

### GetBgpInstance

`func (o *ManaV2EdgeDeviceConfig) GetBgpInstance() ManaV2BgpInstanceConfig`

GetBgpInstance returns the BgpInstance field if non-nil, zero value otherwise.

### GetBgpInstanceOk

`func (o *ManaV2EdgeDeviceConfig) GetBgpInstanceOk() (*ManaV2BgpInstanceConfig, bool)`

GetBgpInstanceOk returns a tuple with the BgpInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpInstance

`func (o *ManaV2EdgeDeviceConfig) SetBgpInstance(v ManaV2BgpInstanceConfig)`

SetBgpInstance sets BgpInstance field to given value.

### HasBgpInstance

`func (o *ManaV2EdgeDeviceConfig) HasBgpInstance() bool`

HasBgpInstance returns a boolean if a field has been set.

### GetCircuits

`func (o *ManaV2EdgeDeviceConfig) GetCircuits() map[string]ManaV2CircuitConfig`

GetCircuits returns the Circuits field if non-nil, zero value otherwise.

### GetCircuitsOk

`func (o *ManaV2EdgeDeviceConfig) GetCircuitsOk() (*map[string]ManaV2CircuitConfig, bool)`

GetCircuitsOk returns a tuple with the Circuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuits

`func (o *ManaV2EdgeDeviceConfig) SetCircuits(v map[string]ManaV2CircuitConfig)`

SetCircuits sets Circuits field to given value.

### HasCircuits

`func (o *ManaV2EdgeDeviceConfig) HasCircuits() bool`

HasCircuits returns a boolean if a field has been set.

### GetDhcpServerEnabled

`func (o *ManaV2EdgeDeviceConfig) GetDhcpServerEnabled() bool`

GetDhcpServerEnabled returns the DhcpServerEnabled field if non-nil, zero value otherwise.

### GetDhcpServerEnabledOk

`func (o *ManaV2EdgeDeviceConfig) GetDhcpServerEnabledOk() (*bool, bool)`

GetDhcpServerEnabledOk returns a tuple with the DhcpServerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerEnabled

`func (o *ManaV2EdgeDeviceConfig) SetDhcpServerEnabled(v bool)`

SetDhcpServerEnabled sets DhcpServerEnabled field to given value.

### HasDhcpServerEnabled

`func (o *ManaV2EdgeDeviceConfig) HasDhcpServerEnabled() bool`

HasDhcpServerEnabled returns a boolean if a field has been set.

### GetDns

`func (o *ManaV2EdgeDeviceConfig) GetDns() ManaV2NullableDnsConfig`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *ManaV2EdgeDeviceConfig) GetDnsOk() (*ManaV2NullableDnsConfig, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *ManaV2EdgeDeviceConfig) SetDns(v ManaV2NullableDnsConfig)`

SetDns sets Dns field to given value.

### HasDns

`func (o *ManaV2EdgeDeviceConfig) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetInterfaces

`func (o *ManaV2EdgeDeviceConfig) GetInterfaces() map[string]ManaV2NullableInterfaceConfig`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ManaV2EdgeDeviceConfig) GetInterfacesOk() (*map[string]ManaV2NullableInterfaceConfig, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ManaV2EdgeDeviceConfig) SetInterfaces(v map[string]ManaV2NullableInterfaceConfig)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ManaV2EdgeDeviceConfig) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetIpfixEnabled

`func (o *ManaV2EdgeDeviceConfig) GetIpfixEnabled() bool`

GetIpfixEnabled returns the IpfixEnabled field if non-nil, zero value otherwise.

### GetIpfixEnabledOk

`func (o *ManaV2EdgeDeviceConfig) GetIpfixEnabledOk() (*bool, bool)`

GetIpfixEnabledOk returns a tuple with the IpfixEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixEnabled

`func (o *ManaV2EdgeDeviceConfig) SetIpfixEnabled(v bool)`

SetIpfixEnabled sets IpfixEnabled field to given value.

### HasIpfixEnabled

`func (o *ManaV2EdgeDeviceConfig) HasIpfixEnabled() bool`

HasIpfixEnabled returns a boolean if a field has been set.

### GetIpfixExporters

`func (o *ManaV2EdgeDeviceConfig) GetIpfixExporters() map[string]ManaV2NullableIpfixExporterConfig`

GetIpfixExporters returns the IpfixExporters field if non-nil, zero value otherwise.

### GetIpfixExportersOk

`func (o *ManaV2EdgeDeviceConfig) GetIpfixExportersOk() (*map[string]ManaV2NullableIpfixExporterConfig, bool)`

GetIpfixExportersOk returns a tuple with the IpfixExporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporters

`func (o *ManaV2EdgeDeviceConfig) SetIpfixExporters(v map[string]ManaV2NullableIpfixExporterConfig)`

SetIpfixExporters sets IpfixExporters field to given value.

### HasIpfixExporters

`func (o *ManaV2EdgeDeviceConfig) HasIpfixExporters() bool`

HasIpfixExporters returns a boolean if a field has been set.

### GetLagInterfaces

`func (o *ManaV2EdgeDeviceConfig) GetLagInterfaces() map[string]ManaV2NullableLagInterfaceConfig`

GetLagInterfaces returns the LagInterfaces field if non-nil, zero value otherwise.

### GetLagInterfacesOk

`func (o *ManaV2EdgeDeviceConfig) GetLagInterfacesOk() (*map[string]ManaV2NullableLagInterfaceConfig, bool)`

GetLagInterfacesOk returns a tuple with the LagInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagInterfaces

`func (o *ManaV2EdgeDeviceConfig) SetLagInterfaces(v map[string]ManaV2NullableLagInterfaceConfig)`

SetLagInterfaces sets LagInterfaces field to given value.

### HasLagInterfaces

`func (o *ManaV2EdgeDeviceConfig) HasLagInterfaces() bool`

HasLagInterfaces returns a boolean if a field has been set.

### GetLldpEnabled

`func (o *ManaV2EdgeDeviceConfig) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *ManaV2EdgeDeviceConfig) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *ManaV2EdgeDeviceConfig) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *ManaV2EdgeDeviceConfig) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.

### GetLocalRouteTag

`func (o *ManaV2EdgeDeviceConfig) GetLocalRouteTag() ManaV2NullableRouteTagSet`

GetLocalRouteTag returns the LocalRouteTag field if non-nil, zero value otherwise.

### GetLocalRouteTagOk

`func (o *ManaV2EdgeDeviceConfig) GetLocalRouteTagOk() (*ManaV2NullableRouteTagSet, bool)`

GetLocalRouteTagOk returns a tuple with the LocalRouteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRouteTag

`func (o *ManaV2EdgeDeviceConfig) SetLocalRouteTag(v ManaV2NullableRouteTagSet)`

SetLocalRouteTag sets LocalRouteTag field to given value.

### HasLocalRouteTag

`func (o *ManaV2EdgeDeviceConfig) HasLocalRouteTag() bool`

HasLocalRouteTag returns a boolean if a field has been set.

### GetLocalWebServerPassword

`func (o *ManaV2EdgeDeviceConfig) GetLocalWebServerPassword() string`

GetLocalWebServerPassword returns the LocalWebServerPassword field if non-nil, zero value otherwise.

### GetLocalWebServerPasswordOk

`func (o *ManaV2EdgeDeviceConfig) GetLocalWebServerPasswordOk() (*string, bool)`

GetLocalWebServerPasswordOk returns a tuple with the LocalWebServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalWebServerPassword

`func (o *ManaV2EdgeDeviceConfig) SetLocalWebServerPassword(v string)`

SetLocalWebServerPassword sets LocalWebServerPassword field to given value.

### HasLocalWebServerPassword

`func (o *ManaV2EdgeDeviceConfig) HasLocalWebServerPassword() bool`

HasLocalWebServerPassword returns a boolean if a field has been set.

### GetLocation

`func (o *ManaV2EdgeDeviceConfig) GetLocation() ManaV2Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ManaV2EdgeDeviceConfig) GetLocationOk() (*ManaV2Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ManaV2EdgeDeviceConfig) SetLocation(v ManaV2Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ManaV2EdgeDeviceConfig) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *ManaV2EdgeDeviceConfig) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *ManaV2EdgeDeviceConfig) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *ManaV2EdgeDeviceConfig) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *ManaV2EdgeDeviceConfig) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetName

`func (o *ManaV2EdgeDeviceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2EdgeDeviceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2EdgeDeviceConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2EdgeDeviceConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNatPolicy

`func (o *ManaV2EdgeDeviceConfig) GetNatPolicy() ManaV2NatPolicyConfig`

GetNatPolicy returns the NatPolicy field if non-nil, zero value otherwise.

### GetNatPolicyOk

`func (o *ManaV2EdgeDeviceConfig) GetNatPolicyOk() (*ManaV2NatPolicyConfig, bool)`

GetNatPolicyOk returns a tuple with the NatPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPolicy

`func (o *ManaV2EdgeDeviceConfig) SetNatPolicy(v ManaV2NatPolicyConfig)`

SetNatPolicy sets NatPolicy field to given value.

### HasNatPolicy

`func (o *ManaV2EdgeDeviceConfig) HasNatPolicy() bool`

HasNatPolicy returns a boolean if a field has been set.

### GetNtpGlobalObject

`func (o *ManaV2EdgeDeviceConfig) GetNtpGlobalObject() map[string]ManaV2NullableNtpConfig`

GetNtpGlobalObject returns the NtpGlobalObject field if non-nil, zero value otherwise.

### GetNtpGlobalObjectOk

`func (o *ManaV2EdgeDeviceConfig) GetNtpGlobalObjectOk() (*map[string]ManaV2NullableNtpConfig, bool)`

GetNtpGlobalObjectOk returns a tuple with the NtpGlobalObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpGlobalObject

`func (o *ManaV2EdgeDeviceConfig) SetNtpGlobalObject(v map[string]ManaV2NullableNtpConfig)`

SetNtpGlobalObject sets NtpGlobalObject field to given value.

### HasNtpGlobalObject

`func (o *ManaV2EdgeDeviceConfig) HasNtpGlobalObject() bool`

HasNtpGlobalObject returns a boolean if a field has been set.

### GetOspfv2Enabled

`func (o *ManaV2EdgeDeviceConfig) GetOspfv2Enabled() bool`

GetOspfv2Enabled returns the Ospfv2Enabled field if non-nil, zero value otherwise.

### GetOspfv2EnabledOk

`func (o *ManaV2EdgeDeviceConfig) GetOspfv2EnabledOk() (*bool, bool)`

GetOspfv2EnabledOk returns a tuple with the Ospfv2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfv2Enabled

`func (o *ManaV2EdgeDeviceConfig) SetOspfv2Enabled(v bool)`

SetOspfv2Enabled sets Ospfv2Enabled field to given value.

### HasOspfv2Enabled

`func (o *ManaV2EdgeDeviceConfig) HasOspfv2Enabled() bool`

HasOspfv2Enabled returns a boolean if a field has been set.

### GetOspfv3Enabled

`func (o *ManaV2EdgeDeviceConfig) GetOspfv3Enabled() bool`

GetOspfv3Enabled returns the Ospfv3Enabled field if non-nil, zero value otherwise.

### GetOspfv3EnabledOk

`func (o *ManaV2EdgeDeviceConfig) GetOspfv3EnabledOk() (*bool, bool)`

GetOspfv3EnabledOk returns a tuple with the Ospfv3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfv3Enabled

`func (o *ManaV2EdgeDeviceConfig) SetOspfv3Enabled(v bool)`

SetOspfv3Enabled sets Ospfv3Enabled field to given value.

### HasOspfv3Enabled

`func (o *ManaV2EdgeDeviceConfig) HasOspfv3Enabled() bool`

HasOspfv3Enabled returns a boolean if a field has been set.

### GetPrefixSets

`func (o *ManaV2EdgeDeviceConfig) GetPrefixSets() map[string]ManaV2NullablePrefixSetConfig`

GetPrefixSets returns the PrefixSets field if non-nil, zero value otherwise.

### GetPrefixSetsOk

`func (o *ManaV2EdgeDeviceConfig) GetPrefixSetsOk() (*map[string]ManaV2NullablePrefixSetConfig, bool)`

GetPrefixSetsOk returns a tuple with the PrefixSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSets

`func (o *ManaV2EdgeDeviceConfig) SetPrefixSets(v map[string]ManaV2NullablePrefixSetConfig)`

SetPrefixSets sets PrefixSets field to given value.

### HasPrefixSets

`func (o *ManaV2EdgeDeviceConfig) HasPrefixSets() bool`

HasPrefixSets returns a boolean if a field has been set.

### GetRegion

`func (o *ManaV2EdgeDeviceConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ManaV2EdgeDeviceConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ManaV2EdgeDeviceConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ManaV2EdgeDeviceConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionName

`func (o *ManaV2EdgeDeviceConfig) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ManaV2EdgeDeviceConfig) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ManaV2EdgeDeviceConfig) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ManaV2EdgeDeviceConfig) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetRoutePolicies

`func (o *ManaV2EdgeDeviceConfig) GetRoutePolicies() map[string]ManaV2NullableRoutingPolicyConfig`

GetRoutePolicies returns the RoutePolicies field if non-nil, zero value otherwise.

### GetRoutePoliciesOk

`func (o *ManaV2EdgeDeviceConfig) GetRoutePoliciesOk() (*map[string]ManaV2NullableRoutingPolicyConfig, bool)`

GetRoutePoliciesOk returns a tuple with the RoutePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePolicies

`func (o *ManaV2EdgeDeviceConfig) SetRoutePolicies(v map[string]ManaV2NullableRoutingPolicyConfig)`

SetRoutePolicies sets RoutePolicies field to given value.

### HasRoutePolicies

`func (o *ManaV2EdgeDeviceConfig) HasRoutePolicies() bool`

HasRoutePolicies returns a boolean if a field has been set.

### GetSegments

`func (o *ManaV2EdgeDeviceConfig) GetSegments() map[string]ManaV2VrfConfig`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *ManaV2EdgeDeviceConfig) GetSegmentsOk() (*map[string]ManaV2VrfConfig, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *ManaV2EdgeDeviceConfig) SetSegments(v map[string]ManaV2VrfConfig)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *ManaV2EdgeDeviceConfig) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetSite

`func (o *ManaV2EdgeDeviceConfig) GetSite() ManaV2NewSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ManaV2EdgeDeviceConfig) GetSiteOk() (*ManaV2NewSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ManaV2EdgeDeviceConfig) SetSite(v ManaV2NewSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *ManaV2EdgeDeviceConfig) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSiteToSiteVpn

`func (o *ManaV2EdgeDeviceConfig) GetSiteToSiteVpn() map[string]ManaV2NullableIPsecTunnelConfig`

GetSiteToSiteVpn returns the SiteToSiteVpn field if non-nil, zero value otherwise.

### GetSiteToSiteVpnOk

`func (o *ManaV2EdgeDeviceConfig) GetSiteToSiteVpnOk() (*map[string]ManaV2NullableIPsecTunnelConfig, bool)`

GetSiteToSiteVpnOk returns a tuple with the SiteToSiteVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteToSiteVpn

`func (o *ManaV2EdgeDeviceConfig) SetSiteToSiteVpn(v map[string]ManaV2NullableIPsecTunnelConfig)`

SetSiteToSiteVpn sets SiteToSiteVpn field to given value.

### HasSiteToSiteVpn

`func (o *ManaV2EdgeDeviceConfig) HasSiteToSiteVpn() bool`

HasSiteToSiteVpn returns a boolean if a field has been set.

### GetSnmp

`func (o *ManaV2EdgeDeviceConfig) GetSnmp() ManaV2NullableSnmpConfig`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *ManaV2EdgeDeviceConfig) GetSnmpOk() (*ManaV2NullableSnmpConfig, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *ManaV2EdgeDeviceConfig) SetSnmp(v ManaV2NullableSnmpConfig)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *ManaV2EdgeDeviceConfig) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetSnmpGlobalObject

`func (o *ManaV2EdgeDeviceConfig) GetSnmpGlobalObject() map[string]ManaV2NullableSnmpConfig`

GetSnmpGlobalObject returns the SnmpGlobalObject field if non-nil, zero value otherwise.

### GetSnmpGlobalObjectOk

`func (o *ManaV2EdgeDeviceConfig) GetSnmpGlobalObjectOk() (*map[string]ManaV2NullableSnmpConfig, bool)`

GetSnmpGlobalObjectOk returns a tuple with the SnmpGlobalObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpGlobalObject

`func (o *ManaV2EdgeDeviceConfig) SetSnmpGlobalObject(v map[string]ManaV2NullableSnmpConfig)`

SetSnmpGlobalObject sets SnmpGlobalObject field to given value.

### HasSnmpGlobalObject

`func (o *ManaV2EdgeDeviceConfig) HasSnmpGlobalObject() bool`

HasSnmpGlobalObject returns a boolean if a field has been set.

### GetStaticRoutesEnabled

`func (o *ManaV2EdgeDeviceConfig) GetStaticRoutesEnabled() bool`

GetStaticRoutesEnabled returns the StaticRoutesEnabled field if non-nil, zero value otherwise.

### GetStaticRoutesEnabledOk

`func (o *ManaV2EdgeDeviceConfig) GetStaticRoutesEnabledOk() (*bool, bool)`

GetStaticRoutesEnabledOk returns a tuple with the StaticRoutesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutesEnabled

`func (o *ManaV2EdgeDeviceConfig) SetStaticRoutesEnabled(v bool)`

SetStaticRoutesEnabled sets StaticRoutesEnabled field to given value.

### HasStaticRoutesEnabled

`func (o *ManaV2EdgeDeviceConfig) HasStaticRoutesEnabled() bool`

HasStaticRoutesEnabled returns a boolean if a field has been set.

### GetTrafficPolicy

`func (o *ManaV2EdgeDeviceConfig) GetTrafficPolicy() ManaV2ForwardingPolicyConfig`

GetTrafficPolicy returns the TrafficPolicy field if non-nil, zero value otherwise.

### GetTrafficPolicyOk

`func (o *ManaV2EdgeDeviceConfig) GetTrafficPolicyOk() (*ManaV2ForwardingPolicyConfig, bool)`

GetTrafficPolicyOk returns a tuple with the TrafficPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPolicy

`func (o *ManaV2EdgeDeviceConfig) SetTrafficPolicy(v ManaV2ForwardingPolicyConfig)`

SetTrafficPolicy sets TrafficPolicy field to given value.

### HasTrafficPolicy

`func (o *ManaV2EdgeDeviceConfig) HasTrafficPolicy() bool`

HasTrafficPolicy returns a boolean if a field has been set.

### GetVrrpEnabled

`func (o *ManaV2EdgeDeviceConfig) GetVrrpEnabled() bool`

GetVrrpEnabled returns the VrrpEnabled field if non-nil, zero value otherwise.

### GetVrrpEnabledOk

`func (o *ManaV2EdgeDeviceConfig) GetVrrpEnabledOk() (*bool, bool)`

GetVrrpEnabledOk returns a tuple with the VrrpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpEnabled

`func (o *ManaV2EdgeDeviceConfig) SetVrrpEnabled(v bool)`

SetVrrpEnabled sets VrrpEnabled field to given value.

### HasVrrpEnabled

`func (o *ManaV2EdgeDeviceConfig) HasVrrpEnabled() bool`

HasVrrpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


