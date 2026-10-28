# ManaV2InterfaceCoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **bool** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Circuit** | Pointer to **string** |  | [optional] 
**CoreNeighbor** | Pointer to [**ManaV2InterfaceCoreToCorePeerConfig**](ManaV2InterfaceCoreToCorePeerConfig.md) |  | [optional] 
**CoreToCoreTunnel** | Pointer to **map[string]interface{}** |  | [optional] 
**CreateLinkLocalAddress** | Pointer to **bool** |  | [optional] 
**Default** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duplex** | Pointer to **string** |  | [optional] 
**Dynamic** | Pointer to [**ManaV2LatencyBandwidth**](ManaV2LatencyBandwidth.md) |  | [optional] 
**FlexAlgos** | Pointer to [**ManaV2InterfaceCoreFlexAlgoConfig**](ManaV2InterfaceCoreFlexAlgoConfig.md) |  | [optional] 
**GatewayNeighbor** | Pointer to [**ManaV2InterfaceCoreToGatewayPeerConfig**](ManaV2InterfaceCoreToGatewayPeerConfig.md) |  | [optional] 
**Gw** | Pointer to [**ManaV2NullableInterfaceIpConfig**](ManaV2NullableInterfaceIpConfig.md) |  | [optional] 
**InterfaceType** | Pointer to [**ManaV2interfaceConfigType**](ManaV2interfaceConfigType.md) |  | [optional] 
**Ipsec** | Pointer to [**ManaV2InterfaceIPsecConfig**](ManaV2InterfaceIPsecConfig.md) |  | [optional] 
**Ipv4** | Pointer to [**ManaV2InterfaceIpConfig**](ManaV2InterfaceIpConfig.md) |  | [optional] 
**Ipv6** | Pointer to [**ManaV2InterfaceIpConfig**](ManaV2InterfaceIpConfig.md) |  | [optional] 
**IsMaintenanceWan** | Pointer to **bool** |  | [optional] 
**Lan** | Pointer to **string** |  | [optional] 
**LldpEnabled** | Pointer to **bool** |  | [optional] 
**Loopback** | Pointer to **bool** |  | [optional] 
**MaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**MplsEnabled** | Pointer to **bool** |  | [optional] 
**OspfCost** | Pointer to [**ManaV2CoreLinkCost**](ManaV2CoreLinkCost.md) |  | [optional] 
**OspfInterface** | Pointer to [**ManaV2NullableOspfInterfaceConfig**](ManaV2NullableOspfInterfaceConfig.md) |  | [optional] 
**PeerDeviceId** | Pointer to **int64** |  | [optional] 
**PeerHostname** | Pointer to **string** |  | [optional] 
**SecurityZone** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int64** |  | [optional] 
**Static** | Pointer to **int32** |  | [optional] 
**Subinterfaces** | Pointer to [**map[string]ManaV2NullableCoreInterfaceVlanConfig**](ManaV2NullableCoreInterfaceVlanConfig.md) |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**TcpMssV4** | Pointer to **int32** |  | [optional] 
**TcpMssV6** | Pointer to **int32** |  | [optional] 
**TunnelInterface** | Pointer to **string** |  | [optional] 
**TunnelUnderlay** | Pointer to **string** |  | [optional] 
**Wan** | Pointer to [**ManaV2InterfaceWanConfig**](ManaV2InterfaceWanConfig.md) |  | [optional] 
**WanManagement** | Pointer to **map[string]interface{}** |  | [optional] 
**XTalkFilter** | Pointer to **bool** |  | [optional] 

## Methods

### NewManaV2InterfaceCoreConfig

`func NewManaV2InterfaceCoreConfig() *ManaV2InterfaceCoreConfig`

NewManaV2InterfaceCoreConfig instantiates a new ManaV2InterfaceCoreConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2InterfaceCoreConfigWithDefaults

`func NewManaV2InterfaceCoreConfigWithDefaults() *ManaV2InterfaceCoreConfig`

NewManaV2InterfaceCoreConfigWithDefaults instantiates a new ManaV2InterfaceCoreConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *ManaV2InterfaceCoreConfig) GetAdminStatus() bool`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *ManaV2InterfaceCoreConfig) GetAdminStatusOk() (*bool, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *ManaV2InterfaceCoreConfig) SetAdminStatus(v bool)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *ManaV2InterfaceCoreConfig) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAlias

`func (o *ManaV2InterfaceCoreConfig) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ManaV2InterfaceCoreConfig) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ManaV2InterfaceCoreConfig) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ManaV2InterfaceCoreConfig) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCircuit

`func (o *ManaV2InterfaceCoreConfig) GetCircuit() string`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *ManaV2InterfaceCoreConfig) GetCircuitOk() (*string, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *ManaV2InterfaceCoreConfig) SetCircuit(v string)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *ManaV2InterfaceCoreConfig) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetCoreNeighbor

`func (o *ManaV2InterfaceCoreConfig) GetCoreNeighbor() ManaV2InterfaceCoreToCorePeerConfig`

GetCoreNeighbor returns the CoreNeighbor field if non-nil, zero value otherwise.

### GetCoreNeighborOk

`func (o *ManaV2InterfaceCoreConfig) GetCoreNeighborOk() (*ManaV2InterfaceCoreToCorePeerConfig, bool)`

GetCoreNeighborOk returns a tuple with the CoreNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreNeighbor

`func (o *ManaV2InterfaceCoreConfig) SetCoreNeighbor(v ManaV2InterfaceCoreToCorePeerConfig)`

SetCoreNeighbor sets CoreNeighbor field to given value.

### HasCoreNeighbor

`func (o *ManaV2InterfaceCoreConfig) HasCoreNeighbor() bool`

HasCoreNeighbor returns a boolean if a field has been set.

### GetCoreToCoreTunnel

`func (o *ManaV2InterfaceCoreConfig) GetCoreToCoreTunnel() map[string]interface{}`

GetCoreToCoreTunnel returns the CoreToCoreTunnel field if non-nil, zero value otherwise.

### GetCoreToCoreTunnelOk

`func (o *ManaV2InterfaceCoreConfig) GetCoreToCoreTunnelOk() (*map[string]interface{}, bool)`

GetCoreToCoreTunnelOk returns a tuple with the CoreToCoreTunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreToCoreTunnel

`func (o *ManaV2InterfaceCoreConfig) SetCoreToCoreTunnel(v map[string]interface{})`

SetCoreToCoreTunnel sets CoreToCoreTunnel field to given value.

### HasCoreToCoreTunnel

`func (o *ManaV2InterfaceCoreConfig) HasCoreToCoreTunnel() bool`

HasCoreToCoreTunnel returns a boolean if a field has been set.

### GetCreateLinkLocalAddress

`func (o *ManaV2InterfaceCoreConfig) GetCreateLinkLocalAddress() bool`

GetCreateLinkLocalAddress returns the CreateLinkLocalAddress field if non-nil, zero value otherwise.

### GetCreateLinkLocalAddressOk

`func (o *ManaV2InterfaceCoreConfig) GetCreateLinkLocalAddressOk() (*bool, bool)`

GetCreateLinkLocalAddressOk returns a tuple with the CreateLinkLocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateLinkLocalAddress

`func (o *ManaV2InterfaceCoreConfig) SetCreateLinkLocalAddress(v bool)`

SetCreateLinkLocalAddress sets CreateLinkLocalAddress field to given value.

### HasCreateLinkLocalAddress

`func (o *ManaV2InterfaceCoreConfig) HasCreateLinkLocalAddress() bool`

HasCreateLinkLocalAddress returns a boolean if a field has been set.

### GetDefault

`func (o *ManaV2InterfaceCoreConfig) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ManaV2InterfaceCoreConfig) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ManaV2InterfaceCoreConfig) SetDefault(v map[string]interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ManaV2InterfaceCoreConfig) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *ManaV2InterfaceCoreConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2InterfaceCoreConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2InterfaceCoreConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2InterfaceCoreConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuplex

`func (o *ManaV2InterfaceCoreConfig) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *ManaV2InterfaceCoreConfig) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *ManaV2InterfaceCoreConfig) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *ManaV2InterfaceCoreConfig) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetDynamic

`func (o *ManaV2InterfaceCoreConfig) GetDynamic() ManaV2LatencyBandwidth`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *ManaV2InterfaceCoreConfig) GetDynamicOk() (*ManaV2LatencyBandwidth, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *ManaV2InterfaceCoreConfig) SetDynamic(v ManaV2LatencyBandwidth)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *ManaV2InterfaceCoreConfig) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetFlexAlgos

`func (o *ManaV2InterfaceCoreConfig) GetFlexAlgos() ManaV2InterfaceCoreFlexAlgoConfig`

GetFlexAlgos returns the FlexAlgos field if non-nil, zero value otherwise.

### GetFlexAlgosOk

`func (o *ManaV2InterfaceCoreConfig) GetFlexAlgosOk() (*ManaV2InterfaceCoreFlexAlgoConfig, bool)`

GetFlexAlgosOk returns a tuple with the FlexAlgos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAlgos

`func (o *ManaV2InterfaceCoreConfig) SetFlexAlgos(v ManaV2InterfaceCoreFlexAlgoConfig)`

SetFlexAlgos sets FlexAlgos field to given value.

### HasFlexAlgos

`func (o *ManaV2InterfaceCoreConfig) HasFlexAlgos() bool`

HasFlexAlgos returns a boolean if a field has been set.

### GetGatewayNeighbor

`func (o *ManaV2InterfaceCoreConfig) GetGatewayNeighbor() ManaV2InterfaceCoreToGatewayPeerConfig`

GetGatewayNeighbor returns the GatewayNeighbor field if non-nil, zero value otherwise.

### GetGatewayNeighborOk

`func (o *ManaV2InterfaceCoreConfig) GetGatewayNeighborOk() (*ManaV2InterfaceCoreToGatewayPeerConfig, bool)`

GetGatewayNeighborOk returns a tuple with the GatewayNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNeighbor

`func (o *ManaV2InterfaceCoreConfig) SetGatewayNeighbor(v ManaV2InterfaceCoreToGatewayPeerConfig)`

SetGatewayNeighbor sets GatewayNeighbor field to given value.

### HasGatewayNeighbor

`func (o *ManaV2InterfaceCoreConfig) HasGatewayNeighbor() bool`

HasGatewayNeighbor returns a boolean if a field has been set.

### GetGw

`func (o *ManaV2InterfaceCoreConfig) GetGw() ManaV2NullableInterfaceIpConfig`

GetGw returns the Gw field if non-nil, zero value otherwise.

### GetGwOk

`func (o *ManaV2InterfaceCoreConfig) GetGwOk() (*ManaV2NullableInterfaceIpConfig, bool)`

GetGwOk returns a tuple with the Gw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGw

`func (o *ManaV2InterfaceCoreConfig) SetGw(v ManaV2NullableInterfaceIpConfig)`

SetGw sets Gw field to given value.

### HasGw

`func (o *ManaV2InterfaceCoreConfig) HasGw() bool`

HasGw returns a boolean if a field has been set.

### GetInterfaceType

`func (o *ManaV2InterfaceCoreConfig) GetInterfaceType() ManaV2interfaceConfigType`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *ManaV2InterfaceCoreConfig) GetInterfaceTypeOk() (*ManaV2interfaceConfigType, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *ManaV2InterfaceCoreConfig) SetInterfaceType(v ManaV2interfaceConfigType)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *ManaV2InterfaceCoreConfig) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetIpsec

`func (o *ManaV2InterfaceCoreConfig) GetIpsec() ManaV2InterfaceIPsecConfig`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *ManaV2InterfaceCoreConfig) GetIpsecOk() (*ManaV2InterfaceIPsecConfig, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *ManaV2InterfaceCoreConfig) SetIpsec(v ManaV2InterfaceIPsecConfig)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *ManaV2InterfaceCoreConfig) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.

### GetIpv4

`func (o *ManaV2InterfaceCoreConfig) GetIpv4() ManaV2InterfaceIpConfig`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *ManaV2InterfaceCoreConfig) GetIpv4Ok() (*ManaV2InterfaceIpConfig, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *ManaV2InterfaceCoreConfig) SetIpv4(v ManaV2InterfaceIpConfig)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *ManaV2InterfaceCoreConfig) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *ManaV2InterfaceCoreConfig) GetIpv6() ManaV2InterfaceIpConfig`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *ManaV2InterfaceCoreConfig) GetIpv6Ok() (*ManaV2InterfaceIpConfig, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *ManaV2InterfaceCoreConfig) SetIpv6(v ManaV2InterfaceIpConfig)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *ManaV2InterfaceCoreConfig) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetIsMaintenanceWan

`func (o *ManaV2InterfaceCoreConfig) GetIsMaintenanceWan() bool`

GetIsMaintenanceWan returns the IsMaintenanceWan field if non-nil, zero value otherwise.

### GetIsMaintenanceWanOk

`func (o *ManaV2InterfaceCoreConfig) GetIsMaintenanceWanOk() (*bool, bool)`

GetIsMaintenanceWanOk returns a tuple with the IsMaintenanceWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenanceWan

`func (o *ManaV2InterfaceCoreConfig) SetIsMaintenanceWan(v bool)`

SetIsMaintenanceWan sets IsMaintenanceWan field to given value.

### HasIsMaintenanceWan

`func (o *ManaV2InterfaceCoreConfig) HasIsMaintenanceWan() bool`

HasIsMaintenanceWan returns a boolean if a field has been set.

### GetLan

`func (o *ManaV2InterfaceCoreConfig) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *ManaV2InterfaceCoreConfig) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *ManaV2InterfaceCoreConfig) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *ManaV2InterfaceCoreConfig) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLldpEnabled

`func (o *ManaV2InterfaceCoreConfig) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *ManaV2InterfaceCoreConfig) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *ManaV2InterfaceCoreConfig) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *ManaV2InterfaceCoreConfig) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.

### GetLoopback

`func (o *ManaV2InterfaceCoreConfig) GetLoopback() bool`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *ManaV2InterfaceCoreConfig) GetLoopbackOk() (*bool, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *ManaV2InterfaceCoreConfig) SetLoopback(v bool)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *ManaV2InterfaceCoreConfig) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *ManaV2InterfaceCoreConfig) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *ManaV2InterfaceCoreConfig) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *ManaV2InterfaceCoreConfig) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *ManaV2InterfaceCoreConfig) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetMplsEnabled

`func (o *ManaV2InterfaceCoreConfig) GetMplsEnabled() bool`

GetMplsEnabled returns the MplsEnabled field if non-nil, zero value otherwise.

### GetMplsEnabledOk

`func (o *ManaV2InterfaceCoreConfig) GetMplsEnabledOk() (*bool, bool)`

GetMplsEnabledOk returns a tuple with the MplsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMplsEnabled

`func (o *ManaV2InterfaceCoreConfig) SetMplsEnabled(v bool)`

SetMplsEnabled sets MplsEnabled field to given value.

### HasMplsEnabled

`func (o *ManaV2InterfaceCoreConfig) HasMplsEnabled() bool`

HasMplsEnabled returns a boolean if a field has been set.

### GetOspfCost

`func (o *ManaV2InterfaceCoreConfig) GetOspfCost() ManaV2CoreLinkCost`

GetOspfCost returns the OspfCost field if non-nil, zero value otherwise.

### GetOspfCostOk

`func (o *ManaV2InterfaceCoreConfig) GetOspfCostOk() (*ManaV2CoreLinkCost, bool)`

GetOspfCostOk returns a tuple with the OspfCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfCost

`func (o *ManaV2InterfaceCoreConfig) SetOspfCost(v ManaV2CoreLinkCost)`

SetOspfCost sets OspfCost field to given value.

### HasOspfCost

`func (o *ManaV2InterfaceCoreConfig) HasOspfCost() bool`

HasOspfCost returns a boolean if a field has been set.

### GetOspfInterface

`func (o *ManaV2InterfaceCoreConfig) GetOspfInterface() ManaV2NullableOspfInterfaceConfig`

GetOspfInterface returns the OspfInterface field if non-nil, zero value otherwise.

### GetOspfInterfaceOk

`func (o *ManaV2InterfaceCoreConfig) GetOspfInterfaceOk() (*ManaV2NullableOspfInterfaceConfig, bool)`

GetOspfInterfaceOk returns a tuple with the OspfInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfInterface

`func (o *ManaV2InterfaceCoreConfig) SetOspfInterface(v ManaV2NullableOspfInterfaceConfig)`

SetOspfInterface sets OspfInterface field to given value.

### HasOspfInterface

`func (o *ManaV2InterfaceCoreConfig) HasOspfInterface() bool`

HasOspfInterface returns a boolean if a field has been set.

### GetPeerDeviceId

`func (o *ManaV2InterfaceCoreConfig) GetPeerDeviceId() int64`

GetPeerDeviceId returns the PeerDeviceId field if non-nil, zero value otherwise.

### GetPeerDeviceIdOk

`func (o *ManaV2InterfaceCoreConfig) GetPeerDeviceIdOk() (*int64, bool)`

GetPeerDeviceIdOk returns a tuple with the PeerDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDeviceId

`func (o *ManaV2InterfaceCoreConfig) SetPeerDeviceId(v int64)`

SetPeerDeviceId sets PeerDeviceId field to given value.

### HasPeerDeviceId

`func (o *ManaV2InterfaceCoreConfig) HasPeerDeviceId() bool`

HasPeerDeviceId returns a boolean if a field has been set.

### GetPeerHostname

`func (o *ManaV2InterfaceCoreConfig) GetPeerHostname() string`

GetPeerHostname returns the PeerHostname field if non-nil, zero value otherwise.

### GetPeerHostnameOk

`func (o *ManaV2InterfaceCoreConfig) GetPeerHostnameOk() (*string, bool)`

GetPeerHostnameOk returns a tuple with the PeerHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerHostname

`func (o *ManaV2InterfaceCoreConfig) SetPeerHostname(v string)`

SetPeerHostname sets PeerHostname field to given value.

### HasPeerHostname

`func (o *ManaV2InterfaceCoreConfig) HasPeerHostname() bool`

HasPeerHostname returns a boolean if a field has been set.

### GetSecurityZone

`func (o *ManaV2InterfaceCoreConfig) GetSecurityZone() string`

GetSecurityZone returns the SecurityZone field if non-nil, zero value otherwise.

### GetSecurityZoneOk

`func (o *ManaV2InterfaceCoreConfig) GetSecurityZoneOk() (*string, bool)`

GetSecurityZoneOk returns a tuple with the SecurityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityZone

`func (o *ManaV2InterfaceCoreConfig) SetSecurityZone(v string)`

SetSecurityZone sets SecurityZone field to given value.

### HasSecurityZone

`func (o *ManaV2InterfaceCoreConfig) HasSecurityZone() bool`

HasSecurityZone returns a boolean if a field has been set.

### GetSpeed

`func (o *ManaV2InterfaceCoreConfig) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ManaV2InterfaceCoreConfig) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ManaV2InterfaceCoreConfig) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ManaV2InterfaceCoreConfig) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatic

`func (o *ManaV2InterfaceCoreConfig) GetStatic() int32`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *ManaV2InterfaceCoreConfig) GetStaticOk() (*int32, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *ManaV2InterfaceCoreConfig) SetStatic(v int32)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *ManaV2InterfaceCoreConfig) HasStatic() bool`

HasStatic returns a boolean if a field has been set.

### GetSubinterfaces

`func (o *ManaV2InterfaceCoreConfig) GetSubinterfaces() map[string]ManaV2NullableCoreInterfaceVlanConfig`

GetSubinterfaces returns the Subinterfaces field if non-nil, zero value otherwise.

### GetSubinterfacesOk

`func (o *ManaV2InterfaceCoreConfig) GetSubinterfacesOk() (*map[string]ManaV2NullableCoreInterfaceVlanConfig, bool)`

GetSubinterfacesOk returns a tuple with the Subinterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubinterfaces

`func (o *ManaV2InterfaceCoreConfig) SetSubinterfaces(v map[string]ManaV2NullableCoreInterfaceVlanConfig)`

SetSubinterfaces sets Subinterfaces field to given value.

### HasSubinterfaces

`func (o *ManaV2InterfaceCoreConfig) HasSubinterfaces() bool`

HasSubinterfaces returns a boolean if a field has been set.

### GetTcpMss

`func (o *ManaV2InterfaceCoreConfig) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *ManaV2InterfaceCoreConfig) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *ManaV2InterfaceCoreConfig) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *ManaV2InterfaceCoreConfig) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetTcpMssV4

`func (o *ManaV2InterfaceCoreConfig) GetTcpMssV4() int32`

GetTcpMssV4 returns the TcpMssV4 field if non-nil, zero value otherwise.

### GetTcpMssV4Ok

`func (o *ManaV2InterfaceCoreConfig) GetTcpMssV4Ok() (*int32, bool)`

GetTcpMssV4Ok returns a tuple with the TcpMssV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV4

`func (o *ManaV2InterfaceCoreConfig) SetTcpMssV4(v int32)`

SetTcpMssV4 sets TcpMssV4 field to given value.

### HasTcpMssV4

`func (o *ManaV2InterfaceCoreConfig) HasTcpMssV4() bool`

HasTcpMssV4 returns a boolean if a field has been set.

### GetTcpMssV6

`func (o *ManaV2InterfaceCoreConfig) GetTcpMssV6() int32`

GetTcpMssV6 returns the TcpMssV6 field if non-nil, zero value otherwise.

### GetTcpMssV6Ok

`func (o *ManaV2InterfaceCoreConfig) GetTcpMssV6Ok() (*int32, bool)`

GetTcpMssV6Ok returns a tuple with the TcpMssV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV6

`func (o *ManaV2InterfaceCoreConfig) SetTcpMssV6(v int32)`

SetTcpMssV6 sets TcpMssV6 field to given value.

### HasTcpMssV6

`func (o *ManaV2InterfaceCoreConfig) HasTcpMssV6() bool`

HasTcpMssV6 returns a boolean if a field has been set.

### GetTunnelInterface

`func (o *ManaV2InterfaceCoreConfig) GetTunnelInterface() string`

GetTunnelInterface returns the TunnelInterface field if non-nil, zero value otherwise.

### GetTunnelInterfaceOk

`func (o *ManaV2InterfaceCoreConfig) GetTunnelInterfaceOk() (*string, bool)`

GetTunnelInterfaceOk returns a tuple with the TunnelInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelInterface

`func (o *ManaV2InterfaceCoreConfig) SetTunnelInterface(v string)`

SetTunnelInterface sets TunnelInterface field to given value.

### HasTunnelInterface

`func (o *ManaV2InterfaceCoreConfig) HasTunnelInterface() bool`

HasTunnelInterface returns a boolean if a field has been set.

### GetTunnelUnderlay

`func (o *ManaV2InterfaceCoreConfig) GetTunnelUnderlay() string`

GetTunnelUnderlay returns the TunnelUnderlay field if non-nil, zero value otherwise.

### GetTunnelUnderlayOk

`func (o *ManaV2InterfaceCoreConfig) GetTunnelUnderlayOk() (*string, bool)`

GetTunnelUnderlayOk returns a tuple with the TunnelUnderlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelUnderlay

`func (o *ManaV2InterfaceCoreConfig) SetTunnelUnderlay(v string)`

SetTunnelUnderlay sets TunnelUnderlay field to given value.

### HasTunnelUnderlay

`func (o *ManaV2InterfaceCoreConfig) HasTunnelUnderlay() bool`

HasTunnelUnderlay returns a boolean if a field has been set.

### GetWan

`func (o *ManaV2InterfaceCoreConfig) GetWan() ManaV2InterfaceWanConfig`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *ManaV2InterfaceCoreConfig) GetWanOk() (*ManaV2InterfaceWanConfig, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *ManaV2InterfaceCoreConfig) SetWan(v ManaV2InterfaceWanConfig)`

SetWan sets Wan field to given value.

### HasWan

`func (o *ManaV2InterfaceCoreConfig) HasWan() bool`

HasWan returns a boolean if a field has been set.

### GetWanManagement

`func (o *ManaV2InterfaceCoreConfig) GetWanManagement() map[string]interface{}`

GetWanManagement returns the WanManagement field if non-nil, zero value otherwise.

### GetWanManagementOk

`func (o *ManaV2InterfaceCoreConfig) GetWanManagementOk() (*map[string]interface{}, bool)`

GetWanManagementOk returns a tuple with the WanManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanManagement

`func (o *ManaV2InterfaceCoreConfig) SetWanManagement(v map[string]interface{})`

SetWanManagement sets WanManagement field to given value.

### HasWanManagement

`func (o *ManaV2InterfaceCoreConfig) HasWanManagement() bool`

HasWanManagement returns a boolean if a field has been set.

### GetXTalkFilter

`func (o *ManaV2InterfaceCoreConfig) GetXTalkFilter() bool`

GetXTalkFilter returns the XTalkFilter field if non-nil, zero value otherwise.

### GetXTalkFilterOk

`func (o *ManaV2InterfaceCoreConfig) GetXTalkFilterOk() (*bool, bool)`

GetXTalkFilterOk returns a tuple with the XTalkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXTalkFilter

`func (o *ManaV2InterfaceCoreConfig) SetXTalkFilter(v bool)`

SetXTalkFilter sets XTalkFilter field to given value.

### HasXTalkFilter

`func (o *ManaV2InterfaceCoreConfig) HasXTalkFilter() bool`

HasXTalkFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


