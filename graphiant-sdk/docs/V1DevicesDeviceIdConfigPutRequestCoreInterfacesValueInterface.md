# V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **bool** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Circuit** | Pointer to **string** |  | [optional] 
**CoreNeighbor** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor.md) |  | [optional] 
**CoreToCoreTunnel** | Pointer to **map[string]interface{}** |  | [optional] 
**CreateLinkLocalAddress** | Pointer to **bool** |  | [optional] 
**Default** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duplex** | Pointer to **string** |  | [optional] 
**Dynamic** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighborCostDynamic**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighborCostDynamic.md) |  | [optional] 
**FlexAlgos** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos.md) |  | [optional] 
**GatewayNeighbor** | Pointer to [**V1AccountMfaGet200Response**](V1AccountMfaGet200Response.md) |  | [optional] 
**Gw** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGw**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGw.md) |  | [optional] 
**InterfaceType** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType.md) |  | [optional] 
**Ipsec** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec.md) |  | [optional] 
**Ipv4** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw.md) |  | [optional] 
**Ipv6** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw.md) |  | [optional] 
**Lan** | Pointer to **string** |  | [optional] 
**LldpEnabled** | Pointer to **bool** |  | [optional] 
**Loopback** | Pointer to **bool** |  | [optional] 
**MaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**MplsEnabled** | Pointer to **bool** |  | [optional] 
**OspfCost** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighborCost**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighborCost.md) |  | [optional] 
**OspfInterface** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterface**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterface.md) |  | [optional] 
**PeerDeviceId** | Pointer to **int64** |  | [optional] 
**PeerHostname** | Pointer to **string** |  | [optional] 
**SecurityZone** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int64** |  | [optional] 
**Static** | Pointer to **int32** |  | [optional] 
**Subinterfaces** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue.md) |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**TcpMssV4** | Pointer to **int32** |  | [optional] 
**TcpMssV6** | Pointer to **int32** |  | [optional] 
**TunnelInterface** | Pointer to **string** |  | [optional] 
**TunnelUnderlay** | Pointer to **string** |  | [optional] 
**Wan** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceWan**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceWan.md) |  | [optional] 
**WanManagement** | Pointer to **map[string]interface{}** |  | [optional] 
**XTalkFilter** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetAdminStatus() bool`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetAdminStatusOk() (*bool, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetAdminStatus(v bool)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAlias

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetCircuit() string`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetCircuitOk() (*string, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetCircuit(v string)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetCoreNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetCoreNeighbor() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor`

GetCoreNeighbor returns the CoreNeighbor field if non-nil, zero value otherwise.

### GetCoreNeighborOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetCoreNeighborOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor, bool)`

GetCoreNeighborOk returns a tuple with the CoreNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetCoreNeighbor(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor)`

SetCoreNeighbor sets CoreNeighbor field to given value.

### HasCoreNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasCoreNeighbor() bool`

HasCoreNeighbor returns a boolean if a field has been set.

### GetCoreToCoreTunnel

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetCoreToCoreTunnel() map[string]interface{}`

GetCoreToCoreTunnel returns the CoreToCoreTunnel field if non-nil, zero value otherwise.

### GetCoreToCoreTunnelOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetCoreToCoreTunnelOk() (*map[string]interface{}, bool)`

GetCoreToCoreTunnelOk returns a tuple with the CoreToCoreTunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreToCoreTunnel

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetCoreToCoreTunnel(v map[string]interface{})`

SetCoreToCoreTunnel sets CoreToCoreTunnel field to given value.

### HasCoreToCoreTunnel

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasCoreToCoreTunnel() bool`

HasCoreToCoreTunnel returns a boolean if a field has been set.

### GetCreateLinkLocalAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetCreateLinkLocalAddress() bool`

GetCreateLinkLocalAddress returns the CreateLinkLocalAddress field if non-nil, zero value otherwise.

### GetCreateLinkLocalAddressOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetCreateLinkLocalAddressOk() (*bool, bool)`

GetCreateLinkLocalAddressOk returns a tuple with the CreateLinkLocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateLinkLocalAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetCreateLinkLocalAddress(v bool)`

SetCreateLinkLocalAddress sets CreateLinkLocalAddress field to given value.

### HasCreateLinkLocalAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasCreateLinkLocalAddress() bool`

HasCreateLinkLocalAddress returns a boolean if a field has been set.

### GetDefault

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetDefault(v map[string]interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuplex

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetDynamic

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetDynamic() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighborCostDynamic`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetDynamicOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighborCostDynamic, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetDynamic(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighborCostDynamic)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetFlexAlgos

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetFlexAlgos() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos`

GetFlexAlgos returns the FlexAlgos field if non-nil, zero value otherwise.

### GetFlexAlgosOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetFlexAlgosOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos, bool)`

GetFlexAlgosOk returns a tuple with the FlexAlgos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAlgos

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetFlexAlgos(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceFlexAlgos)`

SetFlexAlgos sets FlexAlgos field to given value.

### HasFlexAlgos

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasFlexAlgos() bool`

HasFlexAlgos returns a boolean if a field has been set.

### GetGatewayNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetGatewayNeighbor() V1AccountMfaGet200Response`

GetGatewayNeighbor returns the GatewayNeighbor field if non-nil, zero value otherwise.

### GetGatewayNeighborOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetGatewayNeighborOk() (*V1AccountMfaGet200Response, bool)`

GetGatewayNeighborOk returns a tuple with the GatewayNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetGatewayNeighbor(v V1AccountMfaGet200Response)`

SetGatewayNeighbor sets GatewayNeighbor field to given value.

### HasGatewayNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasGatewayNeighbor() bool`

HasGatewayNeighbor returns a boolean if a field has been set.

### GetGw

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetGw() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGw`

GetGw returns the Gw field if non-nil, zero value otherwise.

### GetGwOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetGwOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGw, bool)`

GetGwOk returns a tuple with the Gw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGw

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetGw(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGw)`

SetGw sets Gw field to given value.

### HasGw

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasGw() bool`

HasGw returns a boolean if a field has been set.

### GetInterfaceType

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetInterfaceType() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetInterfaceTypeOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetInterfaceType(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetIpsec

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetIpsec() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetIpsecOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetIpsec(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.

### GetIpv4

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetIpv4() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetIpv4Ok() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetIpv4(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetIpv6() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetIpv6Ok() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetIpv6(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetLan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.

### GetLoopback

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetLoopback() bool`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetLoopbackOk() (*bool, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetLoopback(v bool)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetMplsEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetMplsEnabled() bool`

GetMplsEnabled returns the MplsEnabled field if non-nil, zero value otherwise.

### GetMplsEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetMplsEnabledOk() (*bool, bool)`

GetMplsEnabledOk returns a tuple with the MplsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMplsEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetMplsEnabled(v bool)`

SetMplsEnabled sets MplsEnabled field to given value.

### HasMplsEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasMplsEnabled() bool`

HasMplsEnabled returns a boolean if a field has been set.

### GetOspfCost

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetOspfCost() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighborCost`

GetOspfCost returns the OspfCost field if non-nil, zero value otherwise.

### GetOspfCostOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetOspfCostOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighborCost, bool)`

GetOspfCostOk returns a tuple with the OspfCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfCost

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetOspfCost(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighborCost)`

SetOspfCost sets OspfCost field to given value.

### HasOspfCost

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasOspfCost() bool`

HasOspfCost returns a boolean if a field has been set.

### GetOspfInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetOspfInterface() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterface`

GetOspfInterface returns the OspfInterface field if non-nil, zero value otherwise.

### GetOspfInterfaceOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetOspfInterfaceOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterface, bool)`

GetOspfInterfaceOk returns a tuple with the OspfInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetOspfInterface(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterface)`

SetOspfInterface sets OspfInterface field to given value.

### HasOspfInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasOspfInterface() bool`

HasOspfInterface returns a boolean if a field has been set.

### GetPeerDeviceId

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetPeerDeviceId() int64`

GetPeerDeviceId returns the PeerDeviceId field if non-nil, zero value otherwise.

### GetPeerDeviceIdOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetPeerDeviceIdOk() (*int64, bool)`

GetPeerDeviceIdOk returns a tuple with the PeerDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDeviceId

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetPeerDeviceId(v int64)`

SetPeerDeviceId sets PeerDeviceId field to given value.

### HasPeerDeviceId

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasPeerDeviceId() bool`

HasPeerDeviceId returns a boolean if a field has been set.

### GetPeerHostname

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetPeerHostname() string`

GetPeerHostname returns the PeerHostname field if non-nil, zero value otherwise.

### GetPeerHostnameOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetPeerHostnameOk() (*string, bool)`

GetPeerHostnameOk returns a tuple with the PeerHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerHostname

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetPeerHostname(v string)`

SetPeerHostname sets PeerHostname field to given value.

### HasPeerHostname

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasPeerHostname() bool`

HasPeerHostname returns a boolean if a field has been set.

### GetSecurityZone

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetSecurityZone() string`

GetSecurityZone returns the SecurityZone field if non-nil, zero value otherwise.

### GetSecurityZoneOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetSecurityZoneOk() (*string, bool)`

GetSecurityZoneOk returns a tuple with the SecurityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityZone

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetSecurityZone(v string)`

SetSecurityZone sets SecurityZone field to given value.

### HasSecurityZone

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasSecurityZone() bool`

HasSecurityZone returns a boolean if a field has been set.

### GetSpeed

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatic

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetStatic() int32`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetStaticOk() (*int32, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetStatic(v int32)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasStatic() bool`

HasStatic returns a boolean if a field has been set.

### GetSubinterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetSubinterfaces() map[string]V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue`

GetSubinterfaces returns the Subinterfaces field if non-nil, zero value otherwise.

### GetSubinterfacesOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetSubinterfacesOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue, bool)`

GetSubinterfacesOk returns a tuple with the Subinterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubinterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetSubinterfaces(v map[string]V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValue)`

SetSubinterfaces sets Subinterfaces field to given value.

### HasSubinterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasSubinterfaces() bool`

HasSubinterfaces returns a boolean if a field has been set.

### GetTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetTcpMssV4

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetTcpMssV4() int32`

GetTcpMssV4 returns the TcpMssV4 field if non-nil, zero value otherwise.

### GetTcpMssV4Ok

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetTcpMssV4Ok() (*int32, bool)`

GetTcpMssV4Ok returns a tuple with the TcpMssV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV4

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetTcpMssV4(v int32)`

SetTcpMssV4 sets TcpMssV4 field to given value.

### HasTcpMssV4

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasTcpMssV4() bool`

HasTcpMssV4 returns a boolean if a field has been set.

### GetTcpMssV6

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetTcpMssV6() int32`

GetTcpMssV6 returns the TcpMssV6 field if non-nil, zero value otherwise.

### GetTcpMssV6Ok

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetTcpMssV6Ok() (*int32, bool)`

GetTcpMssV6Ok returns a tuple with the TcpMssV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV6

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetTcpMssV6(v int32)`

SetTcpMssV6 sets TcpMssV6 field to given value.

### HasTcpMssV6

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasTcpMssV6() bool`

HasTcpMssV6 returns a boolean if a field has been set.

### GetTunnelInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetTunnelInterface() string`

GetTunnelInterface returns the TunnelInterface field if non-nil, zero value otherwise.

### GetTunnelInterfaceOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetTunnelInterfaceOk() (*string, bool)`

GetTunnelInterfaceOk returns a tuple with the TunnelInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetTunnelInterface(v string)`

SetTunnelInterface sets TunnelInterface field to given value.

### HasTunnelInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasTunnelInterface() bool`

HasTunnelInterface returns a boolean if a field has been set.

### GetTunnelUnderlay

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetTunnelUnderlay() string`

GetTunnelUnderlay returns the TunnelUnderlay field if non-nil, zero value otherwise.

### GetTunnelUnderlayOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetTunnelUnderlayOk() (*string, bool)`

GetTunnelUnderlayOk returns a tuple with the TunnelUnderlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelUnderlay

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetTunnelUnderlay(v string)`

SetTunnelUnderlay sets TunnelUnderlay field to given value.

### HasTunnelUnderlay

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasTunnelUnderlay() bool`

HasTunnelUnderlay returns a boolean if a field has been set.

### GetWan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetWan() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceWan`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetWanOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceWan, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetWan(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceWan)`

SetWan sets Wan field to given value.

### HasWan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasWan() bool`

HasWan returns a boolean if a field has been set.

### GetWanManagement

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetWanManagement() map[string]interface{}`

GetWanManagement returns the WanManagement field if non-nil, zero value otherwise.

### GetWanManagementOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetWanManagementOk() (*map[string]interface{}, bool)`

GetWanManagementOk returns a tuple with the WanManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanManagement

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetWanManagement(v map[string]interface{})`

SetWanManagement sets WanManagement field to given value.

### HasWanManagement

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasWanManagement() bool`

HasWanManagement returns a boolean if a field has been set.

### GetXTalkFilter

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetXTalkFilter() bool`

GetXTalkFilter returns the XTalkFilter field if non-nil, zero value otherwise.

### GetXTalkFilterOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) GetXTalkFilterOk() (*bool, bool)`

GetXTalkFilterOk returns a tuple with the XTalkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXTalkFilter

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) SetXTalkFilter(v bool)`

SetXTalkFilter sets XTalkFilter field to given value.

### HasXTalkFilter

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterface) HasXTalkFilter() bool`

HasXTalkFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


