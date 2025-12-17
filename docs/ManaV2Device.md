# ManaV2Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bgp** | Pointer to [**ManaV2BgpInstance**](ManaV2BgpInstance.md) |  | [optional] 
**BgpEnabled** | Pointer to **bool** |  | [optional] 
**Circuits** | Pointer to [**[]ManaV2Circuit**](ManaV2Circuit.md) |  | [optional] 
**ConfigUpdatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**CreatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**DhcpServerEnabled** | Pointer to **bool** |  | [optional] 
**Dns** | Pointer to [**ManaV2Dns**](ManaV2Dns.md) |  | [optional] 
**Gdi** | Pointer to **int32** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Interfaces** | Pointer to [**[]ManaV2Interface**](ManaV2Interface.md) |  | [optional] 
**InternalState** | Pointer to **string** |  | [optional] 
**IpfixEnabled** | Pointer to **bool** |  | [optional] 
**IpfixExporters** | Pointer to [**[]ManaV2IpfixExporter**](ManaV2IpfixExporter.md) |  | [optional] 
**IpsecTunnels** | Pointer to [**[]ManaV2SiteToSiteIPsec**](ManaV2SiteToSiteIPsec.md) |  | [optional] 
**LastBootedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**LldpEnabled** | Pointer to **bool** |  | [optional] 
**LocalRouteTag** | Pointer to [**ManaV2RouteTag**](ManaV2RouteTag.md) |  | [optional] 
**LocalWebServerPassword** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**ManaV2Location**](ManaV2Location.md) |  | [optional] 
**MaintenanceMode** | Pointer to **bool** |  | [optional] 
**NatPolicy** | Pointer to [**ManaV2NatPolicy**](ManaV2NatPolicy.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Ntp** | Pointer to [**ManaV2Ntp**](ManaV2Ntp.md) |  | [optional] 
**OperStaled** | Pointer to **bool** |  | [optional] 
**OperStaledAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**OperUpdatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**Ospfv2Enabled** | Pointer to **bool** |  | [optional] 
**Ospfv3Enabled** | Pointer to **bool** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**PrefixSets** | Pointer to [**[]ManaV2PrefixSet**](ManaV2PrefixSet.md) |  | [optional] 
**RebootReason** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**ManaV2Region**](ManaV2Region.md) |  | [optional] 
**RegionOverride** | Pointer to [**ManaV2Region**](ManaV2Region.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**RoutingPolicies** | Pointer to [**[]ManaV2RoutingPolicy**](ManaV2RoutingPolicy.md) |  | [optional] 
**Segments** | Pointer to [**[]ManaV2Vrf**](ManaV2Vrf.md) |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Site** | Pointer to [**ManaV2Site**](ManaV2Site.md) |  | [optional] 
**Snmp** | Pointer to [**ManaV2Snmp**](ManaV2Snmp.md) |  | [optional] 
**SoftwareVersion** | Pointer to **string** |  | [optional] 
**StaticRoutesEnabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TrafficPolicy** | Pointer to [**ManaV2ForwardingPolicy**](ManaV2ForwardingPolicy.md) |  | [optional] 
**Uptime** | Pointer to [**GoogleProtobufDuration**](GoogleProtobufDuration.md) |  | [optional] 
**VrrpEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewManaV2Device

`func NewManaV2Device() *ManaV2Device`

NewManaV2Device instantiates a new ManaV2Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2DeviceWithDefaults

`func NewManaV2DeviceWithDefaults() *ManaV2Device`

NewManaV2DeviceWithDefaults instantiates a new ManaV2Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgp

`func (o *ManaV2Device) GetBgp() ManaV2BgpInstance`

GetBgp returns the Bgp field if non-nil, zero value otherwise.

### GetBgpOk

`func (o *ManaV2Device) GetBgpOk() (*ManaV2BgpInstance, bool)`

GetBgpOk returns a tuple with the Bgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgp

`func (o *ManaV2Device) SetBgp(v ManaV2BgpInstance)`

SetBgp sets Bgp field to given value.

### HasBgp

`func (o *ManaV2Device) HasBgp() bool`

HasBgp returns a boolean if a field has been set.

### GetBgpEnabled

`func (o *ManaV2Device) GetBgpEnabled() bool`

GetBgpEnabled returns the BgpEnabled field if non-nil, zero value otherwise.

### GetBgpEnabledOk

`func (o *ManaV2Device) GetBgpEnabledOk() (*bool, bool)`

GetBgpEnabledOk returns a tuple with the BgpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpEnabled

`func (o *ManaV2Device) SetBgpEnabled(v bool)`

SetBgpEnabled sets BgpEnabled field to given value.

### HasBgpEnabled

`func (o *ManaV2Device) HasBgpEnabled() bool`

HasBgpEnabled returns a boolean if a field has been set.

### GetCircuits

`func (o *ManaV2Device) GetCircuits() []ManaV2Circuit`

GetCircuits returns the Circuits field if non-nil, zero value otherwise.

### GetCircuitsOk

`func (o *ManaV2Device) GetCircuitsOk() (*[]ManaV2Circuit, bool)`

GetCircuitsOk returns a tuple with the Circuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuits

`func (o *ManaV2Device) SetCircuits(v []ManaV2Circuit)`

SetCircuits sets Circuits field to given value.

### HasCircuits

`func (o *ManaV2Device) HasCircuits() bool`

HasCircuits returns a boolean if a field has been set.

### GetConfigUpdatedAt

`func (o *ManaV2Device) GetConfigUpdatedAt() GoogleProtobufTimestamp`

GetConfigUpdatedAt returns the ConfigUpdatedAt field if non-nil, zero value otherwise.

### GetConfigUpdatedAtOk

`func (o *ManaV2Device) GetConfigUpdatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetConfigUpdatedAtOk returns a tuple with the ConfigUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUpdatedAt

`func (o *ManaV2Device) SetConfigUpdatedAt(v GoogleProtobufTimestamp)`

SetConfigUpdatedAt sets ConfigUpdatedAt field to given value.

### HasConfigUpdatedAt

`func (o *ManaV2Device) HasConfigUpdatedAt() bool`

HasConfigUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ManaV2Device) GetCreatedAt() GoogleProtobufTimestamp`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManaV2Device) GetCreatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManaV2Device) SetCreatedAt(v GoogleProtobufTimestamp)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ManaV2Device) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDhcpServerEnabled

`func (o *ManaV2Device) GetDhcpServerEnabled() bool`

GetDhcpServerEnabled returns the DhcpServerEnabled field if non-nil, zero value otherwise.

### GetDhcpServerEnabledOk

`func (o *ManaV2Device) GetDhcpServerEnabledOk() (*bool, bool)`

GetDhcpServerEnabledOk returns a tuple with the DhcpServerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerEnabled

`func (o *ManaV2Device) SetDhcpServerEnabled(v bool)`

SetDhcpServerEnabled sets DhcpServerEnabled field to given value.

### HasDhcpServerEnabled

`func (o *ManaV2Device) HasDhcpServerEnabled() bool`

HasDhcpServerEnabled returns a boolean if a field has been set.

### GetDns

`func (o *ManaV2Device) GetDns() ManaV2Dns`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *ManaV2Device) GetDnsOk() (*ManaV2Dns, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *ManaV2Device) SetDns(v ManaV2Dns)`

SetDns sets Dns field to given value.

### HasDns

`func (o *ManaV2Device) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetGdi

`func (o *ManaV2Device) GetGdi() int32`

GetGdi returns the Gdi field if non-nil, zero value otherwise.

### GetGdiOk

`func (o *ManaV2Device) GetGdiOk() (*int32, bool)`

GetGdiOk returns a tuple with the Gdi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdi

`func (o *ManaV2Device) SetGdi(v int32)`

SetGdi sets Gdi field to given value.

### HasGdi

`func (o *ManaV2Device) HasGdi() bool`

HasGdi returns a boolean if a field has been set.

### GetHostname

`func (o *ManaV2Device) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ManaV2Device) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ManaV2Device) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ManaV2Device) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *ManaV2Device) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2Device) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2Device) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ManaV2Device) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaces

`func (o *ManaV2Device) GetInterfaces() []ManaV2Interface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ManaV2Device) GetInterfacesOk() (*[]ManaV2Interface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ManaV2Device) SetInterfaces(v []ManaV2Interface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ManaV2Device) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetInternalState

`func (o *ManaV2Device) GetInternalState() string`

GetInternalState returns the InternalState field if non-nil, zero value otherwise.

### GetInternalStateOk

`func (o *ManaV2Device) GetInternalStateOk() (*string, bool)`

GetInternalStateOk returns a tuple with the InternalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalState

`func (o *ManaV2Device) SetInternalState(v string)`

SetInternalState sets InternalState field to given value.

### HasInternalState

`func (o *ManaV2Device) HasInternalState() bool`

HasInternalState returns a boolean if a field has been set.

### GetIpfixEnabled

`func (o *ManaV2Device) GetIpfixEnabled() bool`

GetIpfixEnabled returns the IpfixEnabled field if non-nil, zero value otherwise.

### GetIpfixEnabledOk

`func (o *ManaV2Device) GetIpfixEnabledOk() (*bool, bool)`

GetIpfixEnabledOk returns a tuple with the IpfixEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixEnabled

`func (o *ManaV2Device) SetIpfixEnabled(v bool)`

SetIpfixEnabled sets IpfixEnabled field to given value.

### HasIpfixEnabled

`func (o *ManaV2Device) HasIpfixEnabled() bool`

HasIpfixEnabled returns a boolean if a field has been set.

### GetIpfixExporters

`func (o *ManaV2Device) GetIpfixExporters() []ManaV2IpfixExporter`

GetIpfixExporters returns the IpfixExporters field if non-nil, zero value otherwise.

### GetIpfixExportersOk

`func (o *ManaV2Device) GetIpfixExportersOk() (*[]ManaV2IpfixExporter, bool)`

GetIpfixExportersOk returns a tuple with the IpfixExporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporters

`func (o *ManaV2Device) SetIpfixExporters(v []ManaV2IpfixExporter)`

SetIpfixExporters sets IpfixExporters field to given value.

### HasIpfixExporters

`func (o *ManaV2Device) HasIpfixExporters() bool`

HasIpfixExporters returns a boolean if a field has been set.

### GetIpsecTunnels

`func (o *ManaV2Device) GetIpsecTunnels() []ManaV2SiteToSiteIPsec`

GetIpsecTunnels returns the IpsecTunnels field if non-nil, zero value otherwise.

### GetIpsecTunnelsOk

`func (o *ManaV2Device) GetIpsecTunnelsOk() (*[]ManaV2SiteToSiteIPsec, bool)`

GetIpsecTunnelsOk returns a tuple with the IpsecTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecTunnels

`func (o *ManaV2Device) SetIpsecTunnels(v []ManaV2SiteToSiteIPsec)`

SetIpsecTunnels sets IpsecTunnels field to given value.

### HasIpsecTunnels

`func (o *ManaV2Device) HasIpsecTunnels() bool`

HasIpsecTunnels returns a boolean if a field has been set.

### GetLastBootedAt

`func (o *ManaV2Device) GetLastBootedAt() GoogleProtobufTimestamp`

GetLastBootedAt returns the LastBootedAt field if non-nil, zero value otherwise.

### GetLastBootedAtOk

`func (o *ManaV2Device) GetLastBootedAtOk() (*GoogleProtobufTimestamp, bool)`

GetLastBootedAtOk returns a tuple with the LastBootedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBootedAt

`func (o *ManaV2Device) SetLastBootedAt(v GoogleProtobufTimestamp)`

SetLastBootedAt sets LastBootedAt field to given value.

### HasLastBootedAt

`func (o *ManaV2Device) HasLastBootedAt() bool`

HasLastBootedAt returns a boolean if a field has been set.

### GetLldpEnabled

`func (o *ManaV2Device) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *ManaV2Device) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *ManaV2Device) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *ManaV2Device) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.

### GetLocalRouteTag

`func (o *ManaV2Device) GetLocalRouteTag() ManaV2RouteTag`

GetLocalRouteTag returns the LocalRouteTag field if non-nil, zero value otherwise.

### GetLocalRouteTagOk

`func (o *ManaV2Device) GetLocalRouteTagOk() (*ManaV2RouteTag, bool)`

GetLocalRouteTagOk returns a tuple with the LocalRouteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRouteTag

`func (o *ManaV2Device) SetLocalRouteTag(v ManaV2RouteTag)`

SetLocalRouteTag sets LocalRouteTag field to given value.

### HasLocalRouteTag

`func (o *ManaV2Device) HasLocalRouteTag() bool`

HasLocalRouteTag returns a boolean if a field has been set.

### GetLocalWebServerPassword

`func (o *ManaV2Device) GetLocalWebServerPassword() string`

GetLocalWebServerPassword returns the LocalWebServerPassword field if non-nil, zero value otherwise.

### GetLocalWebServerPasswordOk

`func (o *ManaV2Device) GetLocalWebServerPasswordOk() (*string, bool)`

GetLocalWebServerPasswordOk returns a tuple with the LocalWebServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalWebServerPassword

`func (o *ManaV2Device) SetLocalWebServerPassword(v string)`

SetLocalWebServerPassword sets LocalWebServerPassword field to given value.

### HasLocalWebServerPassword

`func (o *ManaV2Device) HasLocalWebServerPassword() bool`

HasLocalWebServerPassword returns a boolean if a field has been set.

### GetLocation

`func (o *ManaV2Device) GetLocation() ManaV2Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ManaV2Device) GetLocationOk() (*ManaV2Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ManaV2Device) SetLocation(v ManaV2Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ManaV2Device) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *ManaV2Device) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *ManaV2Device) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *ManaV2Device) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *ManaV2Device) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetNatPolicy

`func (o *ManaV2Device) GetNatPolicy() ManaV2NatPolicy`

GetNatPolicy returns the NatPolicy field if non-nil, zero value otherwise.

### GetNatPolicyOk

`func (o *ManaV2Device) GetNatPolicyOk() (*ManaV2NatPolicy, bool)`

GetNatPolicyOk returns a tuple with the NatPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPolicy

`func (o *ManaV2Device) SetNatPolicy(v ManaV2NatPolicy)`

SetNatPolicy sets NatPolicy field to given value.

### HasNatPolicy

`func (o *ManaV2Device) HasNatPolicy() bool`

HasNatPolicy returns a boolean if a field has been set.

### GetNotes

`func (o *ManaV2Device) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ManaV2Device) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ManaV2Device) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ManaV2Device) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNtp

`func (o *ManaV2Device) GetNtp() ManaV2Ntp`

GetNtp returns the Ntp field if non-nil, zero value otherwise.

### GetNtpOk

`func (o *ManaV2Device) GetNtpOk() (*ManaV2Ntp, bool)`

GetNtpOk returns a tuple with the Ntp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtp

`func (o *ManaV2Device) SetNtp(v ManaV2Ntp)`

SetNtp sets Ntp field to given value.

### HasNtp

`func (o *ManaV2Device) HasNtp() bool`

HasNtp returns a boolean if a field has been set.

### GetOperStaled

`func (o *ManaV2Device) GetOperStaled() bool`

GetOperStaled returns the OperStaled field if non-nil, zero value otherwise.

### GetOperStaledOk

`func (o *ManaV2Device) GetOperStaledOk() (*bool, bool)`

GetOperStaledOk returns a tuple with the OperStaled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStaled

`func (o *ManaV2Device) SetOperStaled(v bool)`

SetOperStaled sets OperStaled field to given value.

### HasOperStaled

`func (o *ManaV2Device) HasOperStaled() bool`

HasOperStaled returns a boolean if a field has been set.

### GetOperStaledAt

`func (o *ManaV2Device) GetOperStaledAt() GoogleProtobufTimestamp`

GetOperStaledAt returns the OperStaledAt field if non-nil, zero value otherwise.

### GetOperStaledAtOk

`func (o *ManaV2Device) GetOperStaledAtOk() (*GoogleProtobufTimestamp, bool)`

GetOperStaledAtOk returns a tuple with the OperStaledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStaledAt

`func (o *ManaV2Device) SetOperStaledAt(v GoogleProtobufTimestamp)`

SetOperStaledAt sets OperStaledAt field to given value.

### HasOperStaledAt

`func (o *ManaV2Device) HasOperStaledAt() bool`

HasOperStaledAt returns a boolean if a field has been set.

### GetOperUpdatedAt

`func (o *ManaV2Device) GetOperUpdatedAt() GoogleProtobufTimestamp`

GetOperUpdatedAt returns the OperUpdatedAt field if non-nil, zero value otherwise.

### GetOperUpdatedAtOk

`func (o *ManaV2Device) GetOperUpdatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetOperUpdatedAtOk returns a tuple with the OperUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperUpdatedAt

`func (o *ManaV2Device) SetOperUpdatedAt(v GoogleProtobufTimestamp)`

SetOperUpdatedAt sets OperUpdatedAt field to given value.

### HasOperUpdatedAt

`func (o *ManaV2Device) HasOperUpdatedAt() bool`

HasOperUpdatedAt returns a boolean if a field has been set.

### GetOspfv2Enabled

`func (o *ManaV2Device) GetOspfv2Enabled() bool`

GetOspfv2Enabled returns the Ospfv2Enabled field if non-nil, zero value otherwise.

### GetOspfv2EnabledOk

`func (o *ManaV2Device) GetOspfv2EnabledOk() (*bool, bool)`

GetOspfv2EnabledOk returns a tuple with the Ospfv2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfv2Enabled

`func (o *ManaV2Device) SetOspfv2Enabled(v bool)`

SetOspfv2Enabled sets Ospfv2Enabled field to given value.

### HasOspfv2Enabled

`func (o *ManaV2Device) HasOspfv2Enabled() bool`

HasOspfv2Enabled returns a boolean if a field has been set.

### GetOspfv3Enabled

`func (o *ManaV2Device) GetOspfv3Enabled() bool`

GetOspfv3Enabled returns the Ospfv3Enabled field if non-nil, zero value otherwise.

### GetOspfv3EnabledOk

`func (o *ManaV2Device) GetOspfv3EnabledOk() (*bool, bool)`

GetOspfv3EnabledOk returns a tuple with the Ospfv3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfv3Enabled

`func (o *ManaV2Device) SetOspfv3Enabled(v bool)`

SetOspfv3Enabled sets Ospfv3Enabled field to given value.

### HasOspfv3Enabled

`func (o *ManaV2Device) HasOspfv3Enabled() bool`

HasOspfv3Enabled returns a boolean if a field has been set.

### GetPlatform

`func (o *ManaV2Device) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ManaV2Device) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ManaV2Device) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ManaV2Device) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPrefixSets

`func (o *ManaV2Device) GetPrefixSets() []ManaV2PrefixSet`

GetPrefixSets returns the PrefixSets field if non-nil, zero value otherwise.

### GetPrefixSetsOk

`func (o *ManaV2Device) GetPrefixSetsOk() (*[]ManaV2PrefixSet, bool)`

GetPrefixSetsOk returns a tuple with the PrefixSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSets

`func (o *ManaV2Device) SetPrefixSets(v []ManaV2PrefixSet)`

SetPrefixSets sets PrefixSets field to given value.

### HasPrefixSets

`func (o *ManaV2Device) HasPrefixSets() bool`

HasPrefixSets returns a boolean if a field has been set.

### GetRebootReason

`func (o *ManaV2Device) GetRebootReason() string`

GetRebootReason returns the RebootReason field if non-nil, zero value otherwise.

### GetRebootReasonOk

`func (o *ManaV2Device) GetRebootReasonOk() (*string, bool)`

GetRebootReasonOk returns a tuple with the RebootReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootReason

`func (o *ManaV2Device) SetRebootReason(v string)`

SetRebootReason sets RebootReason field to given value.

### HasRebootReason

`func (o *ManaV2Device) HasRebootReason() bool`

HasRebootReason returns a boolean if a field has been set.

### GetRegion

`func (o *ManaV2Device) GetRegion() ManaV2Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ManaV2Device) GetRegionOk() (*ManaV2Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ManaV2Device) SetRegion(v ManaV2Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ManaV2Device) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionOverride

`func (o *ManaV2Device) GetRegionOverride() ManaV2Region`

GetRegionOverride returns the RegionOverride field if non-nil, zero value otherwise.

### GetRegionOverrideOk

`func (o *ManaV2Device) GetRegionOverrideOk() (*ManaV2Region, bool)`

GetRegionOverrideOk returns a tuple with the RegionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionOverride

`func (o *ManaV2Device) SetRegionOverride(v ManaV2Region)`

SetRegionOverride sets RegionOverride field to given value.

### HasRegionOverride

`func (o *ManaV2Device) HasRegionOverride() bool`

HasRegionOverride returns a boolean if a field has been set.

### GetRole

`func (o *ManaV2Device) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ManaV2Device) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ManaV2Device) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ManaV2Device) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoutingPolicies

`func (o *ManaV2Device) GetRoutingPolicies() []ManaV2RoutingPolicy`

GetRoutingPolicies returns the RoutingPolicies field if non-nil, zero value otherwise.

### GetRoutingPoliciesOk

`func (o *ManaV2Device) GetRoutingPoliciesOk() (*[]ManaV2RoutingPolicy, bool)`

GetRoutingPoliciesOk returns a tuple with the RoutingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicies

`func (o *ManaV2Device) SetRoutingPolicies(v []ManaV2RoutingPolicy)`

SetRoutingPolicies sets RoutingPolicies field to given value.

### HasRoutingPolicies

`func (o *ManaV2Device) HasRoutingPolicies() bool`

HasRoutingPolicies returns a boolean if a field has been set.

### GetSegments

`func (o *ManaV2Device) GetSegments() []ManaV2Vrf`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *ManaV2Device) GetSegmentsOk() (*[]ManaV2Vrf, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *ManaV2Device) SetSegments(v []ManaV2Vrf)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *ManaV2Device) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ManaV2Device) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ManaV2Device) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ManaV2Device) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ManaV2Device) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSite

`func (o *ManaV2Device) GetSite() ManaV2Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ManaV2Device) GetSiteOk() (*ManaV2Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ManaV2Device) SetSite(v ManaV2Site)`

SetSite sets Site field to given value.

### HasSite

`func (o *ManaV2Device) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSnmp

`func (o *ManaV2Device) GetSnmp() ManaV2Snmp`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *ManaV2Device) GetSnmpOk() (*ManaV2Snmp, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *ManaV2Device) SetSnmp(v ManaV2Snmp)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *ManaV2Device) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *ManaV2Device) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *ManaV2Device) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *ManaV2Device) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *ManaV2Device) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetStaticRoutesEnabled

`func (o *ManaV2Device) GetStaticRoutesEnabled() bool`

GetStaticRoutesEnabled returns the StaticRoutesEnabled field if non-nil, zero value otherwise.

### GetStaticRoutesEnabledOk

`func (o *ManaV2Device) GetStaticRoutesEnabledOk() (*bool, bool)`

GetStaticRoutesEnabledOk returns a tuple with the StaticRoutesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutesEnabled

`func (o *ManaV2Device) SetStaticRoutesEnabled(v bool)`

SetStaticRoutesEnabled sets StaticRoutesEnabled field to given value.

### HasStaticRoutesEnabled

`func (o *ManaV2Device) HasStaticRoutesEnabled() bool`

HasStaticRoutesEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *ManaV2Device) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManaV2Device) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManaV2Device) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManaV2Device) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrafficPolicy

`func (o *ManaV2Device) GetTrafficPolicy() ManaV2ForwardingPolicy`

GetTrafficPolicy returns the TrafficPolicy field if non-nil, zero value otherwise.

### GetTrafficPolicyOk

`func (o *ManaV2Device) GetTrafficPolicyOk() (*ManaV2ForwardingPolicy, bool)`

GetTrafficPolicyOk returns a tuple with the TrafficPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPolicy

`func (o *ManaV2Device) SetTrafficPolicy(v ManaV2ForwardingPolicy)`

SetTrafficPolicy sets TrafficPolicy field to given value.

### HasTrafficPolicy

`func (o *ManaV2Device) HasTrafficPolicy() bool`

HasTrafficPolicy returns a boolean if a field has been set.

### GetUptime

`func (o *ManaV2Device) GetUptime() GoogleProtobufDuration`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ManaV2Device) GetUptimeOk() (*GoogleProtobufDuration, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ManaV2Device) SetUptime(v GoogleProtobufDuration)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ManaV2Device) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVrrpEnabled

`func (o *ManaV2Device) GetVrrpEnabled() bool`

GetVrrpEnabled returns the VrrpEnabled field if non-nil, zero value otherwise.

### GetVrrpEnabledOk

`func (o *ManaV2Device) GetVrrpEnabledOk() (*bool, bool)`

GetVrrpEnabledOk returns a tuple with the VrrpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpEnabled

`func (o *ManaV2Device) SetVrrpEnabled(v bool)`

SetVrrpEnabled sets VrrpEnabled field to given value.

### HasVrrpEnabled

`func (o *ManaV2Device) HasVrrpEnabled() bool`

HasVrrpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


