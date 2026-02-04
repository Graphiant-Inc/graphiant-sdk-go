# ManaV2InterfaceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **bool** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Circuit** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duplex** | Pointer to **string** |  | [optional] 
**Ipsec** | Pointer to [**ManaV2InterfaceIPsecConfig**](ManaV2InterfaceIPsecConfig.md) |  | [optional] 
**Ipv4** | Pointer to [**ManaV2InterfaceIpConfig**](ManaV2InterfaceIpConfig.md) |  | [optional] 
**Ipv6** | Pointer to [**ManaV2InterfaceIpConfig**](ManaV2InterfaceIpConfig.md) |  | [optional] 
**Lan** | Pointer to **string** |  | [optional] 
**LldpEnabled** | Pointer to **bool** |  | [optional] 
**Loopback** | Pointer to **bool** |  | [optional] 
**Macsec** | Pointer to [**ManaV2NullableMaCsecConfiguration**](ManaV2NullableMaCsecConfiguration.md) |  | [optional] 
**MaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**SecurityZone** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int64** |  | [optional] 
**Subinterfaces** | Pointer to [**map[string]ManaV2NullableInterfaceVlanConfig**](ManaV2NullableInterfaceVlanConfig.md) |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**TcpMssV4** | Pointer to **int32** |  | [optional] 
**TcpMssV6** | Pointer to **int32** |  | [optional] 
**V4TcpMss** | Pointer to [**ManaV2NullableTcpMssV4**](ManaV2NullableTcpMssV4.md) |  | [optional] 
**V6TcpMss** | Pointer to [**ManaV2NullableTcpMssV6**](ManaV2NullableTcpMssV6.md) |  | [optional] 

## Methods

### NewManaV2InterfaceConfig

`func NewManaV2InterfaceConfig() *ManaV2InterfaceConfig`

NewManaV2InterfaceConfig instantiates a new ManaV2InterfaceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2InterfaceConfigWithDefaults

`func NewManaV2InterfaceConfigWithDefaults() *ManaV2InterfaceConfig`

NewManaV2InterfaceConfigWithDefaults instantiates a new ManaV2InterfaceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *ManaV2InterfaceConfig) GetAdminStatus() bool`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *ManaV2InterfaceConfig) GetAdminStatusOk() (*bool, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *ManaV2InterfaceConfig) SetAdminStatus(v bool)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *ManaV2InterfaceConfig) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAlias

`func (o *ManaV2InterfaceConfig) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ManaV2InterfaceConfig) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ManaV2InterfaceConfig) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ManaV2InterfaceConfig) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCircuit

`func (o *ManaV2InterfaceConfig) GetCircuit() string`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *ManaV2InterfaceConfig) GetCircuitOk() (*string, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *ManaV2InterfaceConfig) SetCircuit(v string)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *ManaV2InterfaceConfig) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetDescription

`func (o *ManaV2InterfaceConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2InterfaceConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2InterfaceConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2InterfaceConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuplex

`func (o *ManaV2InterfaceConfig) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *ManaV2InterfaceConfig) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *ManaV2InterfaceConfig) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *ManaV2InterfaceConfig) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetIpsec

`func (o *ManaV2InterfaceConfig) GetIpsec() ManaV2InterfaceIPsecConfig`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *ManaV2InterfaceConfig) GetIpsecOk() (*ManaV2InterfaceIPsecConfig, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *ManaV2InterfaceConfig) SetIpsec(v ManaV2InterfaceIPsecConfig)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *ManaV2InterfaceConfig) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.

### GetIpv4

`func (o *ManaV2InterfaceConfig) GetIpv4() ManaV2InterfaceIpConfig`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *ManaV2InterfaceConfig) GetIpv4Ok() (*ManaV2InterfaceIpConfig, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *ManaV2InterfaceConfig) SetIpv4(v ManaV2InterfaceIpConfig)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *ManaV2InterfaceConfig) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *ManaV2InterfaceConfig) GetIpv6() ManaV2InterfaceIpConfig`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *ManaV2InterfaceConfig) GetIpv6Ok() (*ManaV2InterfaceIpConfig, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *ManaV2InterfaceConfig) SetIpv6(v ManaV2InterfaceIpConfig)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *ManaV2InterfaceConfig) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetLan

`func (o *ManaV2InterfaceConfig) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *ManaV2InterfaceConfig) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *ManaV2InterfaceConfig) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *ManaV2InterfaceConfig) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLldpEnabled

`func (o *ManaV2InterfaceConfig) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *ManaV2InterfaceConfig) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *ManaV2InterfaceConfig) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *ManaV2InterfaceConfig) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.

### GetLoopback

`func (o *ManaV2InterfaceConfig) GetLoopback() bool`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *ManaV2InterfaceConfig) GetLoopbackOk() (*bool, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *ManaV2InterfaceConfig) SetLoopback(v bool)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *ManaV2InterfaceConfig) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetMacsec

`func (o *ManaV2InterfaceConfig) GetMacsec() ManaV2NullableMaCsecConfiguration`

GetMacsec returns the Macsec field if non-nil, zero value otherwise.

### GetMacsecOk

`func (o *ManaV2InterfaceConfig) GetMacsecOk() (*ManaV2NullableMaCsecConfiguration, bool)`

GetMacsecOk returns a tuple with the Macsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacsec

`func (o *ManaV2InterfaceConfig) SetMacsec(v ManaV2NullableMaCsecConfiguration)`

SetMacsec sets Macsec field to given value.

### HasMacsec

`func (o *ManaV2InterfaceConfig) HasMacsec() bool`

HasMacsec returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *ManaV2InterfaceConfig) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *ManaV2InterfaceConfig) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *ManaV2InterfaceConfig) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *ManaV2InterfaceConfig) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetSecurityZone

`func (o *ManaV2InterfaceConfig) GetSecurityZone() string`

GetSecurityZone returns the SecurityZone field if non-nil, zero value otherwise.

### GetSecurityZoneOk

`func (o *ManaV2InterfaceConfig) GetSecurityZoneOk() (*string, bool)`

GetSecurityZoneOk returns a tuple with the SecurityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityZone

`func (o *ManaV2InterfaceConfig) SetSecurityZone(v string)`

SetSecurityZone sets SecurityZone field to given value.

### HasSecurityZone

`func (o *ManaV2InterfaceConfig) HasSecurityZone() bool`

HasSecurityZone returns a boolean if a field has been set.

### GetSpeed

`func (o *ManaV2InterfaceConfig) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ManaV2InterfaceConfig) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ManaV2InterfaceConfig) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ManaV2InterfaceConfig) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSubinterfaces

`func (o *ManaV2InterfaceConfig) GetSubinterfaces() map[string]ManaV2NullableInterfaceVlanConfig`

GetSubinterfaces returns the Subinterfaces field if non-nil, zero value otherwise.

### GetSubinterfacesOk

`func (o *ManaV2InterfaceConfig) GetSubinterfacesOk() (*map[string]ManaV2NullableInterfaceVlanConfig, bool)`

GetSubinterfacesOk returns a tuple with the Subinterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubinterfaces

`func (o *ManaV2InterfaceConfig) SetSubinterfaces(v map[string]ManaV2NullableInterfaceVlanConfig)`

SetSubinterfaces sets Subinterfaces field to given value.

### HasSubinterfaces

`func (o *ManaV2InterfaceConfig) HasSubinterfaces() bool`

HasSubinterfaces returns a boolean if a field has been set.

### GetTcpMss

`func (o *ManaV2InterfaceConfig) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *ManaV2InterfaceConfig) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *ManaV2InterfaceConfig) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *ManaV2InterfaceConfig) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetTcpMssV4

`func (o *ManaV2InterfaceConfig) GetTcpMssV4() int32`

GetTcpMssV4 returns the TcpMssV4 field if non-nil, zero value otherwise.

### GetTcpMssV4Ok

`func (o *ManaV2InterfaceConfig) GetTcpMssV4Ok() (*int32, bool)`

GetTcpMssV4Ok returns a tuple with the TcpMssV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV4

`func (o *ManaV2InterfaceConfig) SetTcpMssV4(v int32)`

SetTcpMssV4 sets TcpMssV4 field to given value.

### HasTcpMssV4

`func (o *ManaV2InterfaceConfig) HasTcpMssV4() bool`

HasTcpMssV4 returns a boolean if a field has been set.

### GetTcpMssV6

`func (o *ManaV2InterfaceConfig) GetTcpMssV6() int32`

GetTcpMssV6 returns the TcpMssV6 field if non-nil, zero value otherwise.

### GetTcpMssV6Ok

`func (o *ManaV2InterfaceConfig) GetTcpMssV6Ok() (*int32, bool)`

GetTcpMssV6Ok returns a tuple with the TcpMssV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV6

`func (o *ManaV2InterfaceConfig) SetTcpMssV6(v int32)`

SetTcpMssV6 sets TcpMssV6 field to given value.

### HasTcpMssV6

`func (o *ManaV2InterfaceConfig) HasTcpMssV6() bool`

HasTcpMssV6 returns a boolean if a field has been set.

### GetV4TcpMss

`func (o *ManaV2InterfaceConfig) GetV4TcpMss() ManaV2NullableTcpMssV4`

GetV4TcpMss returns the V4TcpMss field if non-nil, zero value otherwise.

### GetV4TcpMssOk

`func (o *ManaV2InterfaceConfig) GetV4TcpMssOk() (*ManaV2NullableTcpMssV4, bool)`

GetV4TcpMssOk returns a tuple with the V4TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4TcpMss

`func (o *ManaV2InterfaceConfig) SetV4TcpMss(v ManaV2NullableTcpMssV4)`

SetV4TcpMss sets V4TcpMss field to given value.

### HasV4TcpMss

`func (o *ManaV2InterfaceConfig) HasV4TcpMss() bool`

HasV4TcpMss returns a boolean if a field has been set.

### GetV6TcpMss

`func (o *ManaV2InterfaceConfig) GetV6TcpMss() ManaV2NullableTcpMssV6`

GetV6TcpMss returns the V6TcpMss field if non-nil, zero value otherwise.

### GetV6TcpMssOk

`func (o *ManaV2InterfaceConfig) GetV6TcpMssOk() (*ManaV2NullableTcpMssV6, bool)`

GetV6TcpMssOk returns a tuple with the V6TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6TcpMss

`func (o *ManaV2InterfaceConfig) SetV6TcpMss(v ManaV2NullableTcpMssV6)`

SetV6TcpMss sets V6TcpMss field to given value.

### HasV6TcpMss

`func (o *ManaV2InterfaceConfig) HasV6TcpMss() bool`

HasV6TcpMss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


