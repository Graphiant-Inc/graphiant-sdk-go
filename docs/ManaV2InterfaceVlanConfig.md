# ManaV2InterfaceVlanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **bool** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Circuit** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Ipv4** | Pointer to [**ManaV2InterfaceIpConfig**](ManaV2InterfaceIpConfig.md) |  | [optional] 
**Ipv6** | Pointer to [**ManaV2InterfaceIpConfig**](ManaV2InterfaceIpConfig.md) |  | [optional] 
**Lan** | Pointer to **string** |  | [optional] 
**LldpEnabled** | Pointer to **bool** |  | [optional] 
**MaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**PppoeInterfaces** | Pointer to [**map[string]ManaV2NullablePpPoEInterfaceConfig**](ManaV2NullablePpPoEInterfaceConfig.md) |  | [optional] 
**SecurityZone** | Pointer to **string** |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**TcpMssV4** | Pointer to **int32** |  | [optional] 
**TcpMssV6** | Pointer to **int32** |  | [optional] 
**V4TcpMss** | Pointer to [**ManaV2NullableTcpMssV4**](ManaV2NullableTcpMssV4.md) |  | [optional] 
**V6TcpMss** | Pointer to [**ManaV2NullableTcpMssV6**](ManaV2NullableTcpMssV6.md) |  | [optional] 
**Vlan** | Pointer to **int32** |  | [optional] 
**Vrrp** | Pointer to [**ManaV2NullableVrrpGroupConfig**](ManaV2NullableVrrpGroupConfig.md) |  | [optional] 

## Methods

### NewManaV2InterfaceVlanConfig

`func NewManaV2InterfaceVlanConfig() *ManaV2InterfaceVlanConfig`

NewManaV2InterfaceVlanConfig instantiates a new ManaV2InterfaceVlanConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2InterfaceVlanConfigWithDefaults

`func NewManaV2InterfaceVlanConfigWithDefaults() *ManaV2InterfaceVlanConfig`

NewManaV2InterfaceVlanConfigWithDefaults instantiates a new ManaV2InterfaceVlanConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *ManaV2InterfaceVlanConfig) GetAdminStatus() bool`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *ManaV2InterfaceVlanConfig) GetAdminStatusOk() (*bool, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *ManaV2InterfaceVlanConfig) SetAdminStatus(v bool)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *ManaV2InterfaceVlanConfig) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAlias

`func (o *ManaV2InterfaceVlanConfig) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ManaV2InterfaceVlanConfig) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ManaV2InterfaceVlanConfig) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ManaV2InterfaceVlanConfig) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCircuit

`func (o *ManaV2InterfaceVlanConfig) GetCircuit() string`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *ManaV2InterfaceVlanConfig) GetCircuitOk() (*string, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *ManaV2InterfaceVlanConfig) SetCircuit(v string)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *ManaV2InterfaceVlanConfig) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetDescription

`func (o *ManaV2InterfaceVlanConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2InterfaceVlanConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2InterfaceVlanConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2InterfaceVlanConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *ManaV2InterfaceVlanConfig) GetIpv4() ManaV2InterfaceIpConfig`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *ManaV2InterfaceVlanConfig) GetIpv4Ok() (*ManaV2InterfaceIpConfig, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *ManaV2InterfaceVlanConfig) SetIpv4(v ManaV2InterfaceIpConfig)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *ManaV2InterfaceVlanConfig) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *ManaV2InterfaceVlanConfig) GetIpv6() ManaV2InterfaceIpConfig`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *ManaV2InterfaceVlanConfig) GetIpv6Ok() (*ManaV2InterfaceIpConfig, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *ManaV2InterfaceVlanConfig) SetIpv6(v ManaV2InterfaceIpConfig)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *ManaV2InterfaceVlanConfig) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetLan

`func (o *ManaV2InterfaceVlanConfig) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *ManaV2InterfaceVlanConfig) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *ManaV2InterfaceVlanConfig) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *ManaV2InterfaceVlanConfig) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLldpEnabled

`func (o *ManaV2InterfaceVlanConfig) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *ManaV2InterfaceVlanConfig) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *ManaV2InterfaceVlanConfig) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *ManaV2InterfaceVlanConfig) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *ManaV2InterfaceVlanConfig) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *ManaV2InterfaceVlanConfig) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *ManaV2InterfaceVlanConfig) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *ManaV2InterfaceVlanConfig) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetPppoeInterfaces

`func (o *ManaV2InterfaceVlanConfig) GetPppoeInterfaces() map[string]ManaV2NullablePpPoEInterfaceConfig`

GetPppoeInterfaces returns the PppoeInterfaces field if non-nil, zero value otherwise.

### GetPppoeInterfacesOk

`func (o *ManaV2InterfaceVlanConfig) GetPppoeInterfacesOk() (*map[string]ManaV2NullablePpPoEInterfaceConfig, bool)`

GetPppoeInterfacesOk returns a tuple with the PppoeInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPppoeInterfaces

`func (o *ManaV2InterfaceVlanConfig) SetPppoeInterfaces(v map[string]ManaV2NullablePpPoEInterfaceConfig)`

SetPppoeInterfaces sets PppoeInterfaces field to given value.

### HasPppoeInterfaces

`func (o *ManaV2InterfaceVlanConfig) HasPppoeInterfaces() bool`

HasPppoeInterfaces returns a boolean if a field has been set.

### GetSecurityZone

`func (o *ManaV2InterfaceVlanConfig) GetSecurityZone() string`

GetSecurityZone returns the SecurityZone field if non-nil, zero value otherwise.

### GetSecurityZoneOk

`func (o *ManaV2InterfaceVlanConfig) GetSecurityZoneOk() (*string, bool)`

GetSecurityZoneOk returns a tuple with the SecurityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityZone

`func (o *ManaV2InterfaceVlanConfig) SetSecurityZone(v string)`

SetSecurityZone sets SecurityZone field to given value.

### HasSecurityZone

`func (o *ManaV2InterfaceVlanConfig) HasSecurityZone() bool`

HasSecurityZone returns a boolean if a field has been set.

### GetTcpMss

`func (o *ManaV2InterfaceVlanConfig) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *ManaV2InterfaceVlanConfig) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *ManaV2InterfaceVlanConfig) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *ManaV2InterfaceVlanConfig) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetTcpMssV4

`func (o *ManaV2InterfaceVlanConfig) GetTcpMssV4() int32`

GetTcpMssV4 returns the TcpMssV4 field if non-nil, zero value otherwise.

### GetTcpMssV4Ok

`func (o *ManaV2InterfaceVlanConfig) GetTcpMssV4Ok() (*int32, bool)`

GetTcpMssV4Ok returns a tuple with the TcpMssV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV4

`func (o *ManaV2InterfaceVlanConfig) SetTcpMssV4(v int32)`

SetTcpMssV4 sets TcpMssV4 field to given value.

### HasTcpMssV4

`func (o *ManaV2InterfaceVlanConfig) HasTcpMssV4() bool`

HasTcpMssV4 returns a boolean if a field has been set.

### GetTcpMssV6

`func (o *ManaV2InterfaceVlanConfig) GetTcpMssV6() int32`

GetTcpMssV6 returns the TcpMssV6 field if non-nil, zero value otherwise.

### GetTcpMssV6Ok

`func (o *ManaV2InterfaceVlanConfig) GetTcpMssV6Ok() (*int32, bool)`

GetTcpMssV6Ok returns a tuple with the TcpMssV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV6

`func (o *ManaV2InterfaceVlanConfig) SetTcpMssV6(v int32)`

SetTcpMssV6 sets TcpMssV6 field to given value.

### HasTcpMssV6

`func (o *ManaV2InterfaceVlanConfig) HasTcpMssV6() bool`

HasTcpMssV6 returns a boolean if a field has been set.

### GetV4TcpMss

`func (o *ManaV2InterfaceVlanConfig) GetV4TcpMss() ManaV2NullableTcpMssV4`

GetV4TcpMss returns the V4TcpMss field if non-nil, zero value otherwise.

### GetV4TcpMssOk

`func (o *ManaV2InterfaceVlanConfig) GetV4TcpMssOk() (*ManaV2NullableTcpMssV4, bool)`

GetV4TcpMssOk returns a tuple with the V4TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4TcpMss

`func (o *ManaV2InterfaceVlanConfig) SetV4TcpMss(v ManaV2NullableTcpMssV4)`

SetV4TcpMss sets V4TcpMss field to given value.

### HasV4TcpMss

`func (o *ManaV2InterfaceVlanConfig) HasV4TcpMss() bool`

HasV4TcpMss returns a boolean if a field has been set.

### GetV6TcpMss

`func (o *ManaV2InterfaceVlanConfig) GetV6TcpMss() ManaV2NullableTcpMssV6`

GetV6TcpMss returns the V6TcpMss field if non-nil, zero value otherwise.

### GetV6TcpMssOk

`func (o *ManaV2InterfaceVlanConfig) GetV6TcpMssOk() (*ManaV2NullableTcpMssV6, bool)`

GetV6TcpMssOk returns a tuple with the V6TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6TcpMss

`func (o *ManaV2InterfaceVlanConfig) SetV6TcpMss(v ManaV2NullableTcpMssV6)`

SetV6TcpMss sets V6TcpMss field to given value.

### HasV6TcpMss

`func (o *ManaV2InterfaceVlanConfig) HasV6TcpMss() bool`

HasV6TcpMss returns a boolean if a field has been set.

### GetVlan

`func (o *ManaV2InterfaceVlanConfig) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *ManaV2InterfaceVlanConfig) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *ManaV2InterfaceVlanConfig) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *ManaV2InterfaceVlanConfig) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVrrp

`func (o *ManaV2InterfaceVlanConfig) GetVrrp() ManaV2NullableVrrpGroupConfig`

GetVrrp returns the Vrrp field if non-nil, zero value otherwise.

### GetVrrpOk

`func (o *ManaV2InterfaceVlanConfig) GetVrrpOk() (*ManaV2NullableVrrpGroupConfig, bool)`

GetVrrpOk returns a tuple with the Vrrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrp

`func (o *ManaV2InterfaceVlanConfig) SetVrrp(v ManaV2NullableVrrpGroupConfig)`

SetVrrp sets Vrrp field to given value.

### HasVrrp

`func (o *ManaV2InterfaceVlanConfig) HasVrrp() bool`

HasVrrp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


