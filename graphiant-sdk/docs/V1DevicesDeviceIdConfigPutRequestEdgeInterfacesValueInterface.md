# V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **bool** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Circuit** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Ipsec** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec.md) |  | [optional] 
**Ipv4** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw.md) |  | [optional] 
**Ipv6** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw.md) |  | [optional] 
**Lan** | Pointer to **string** |  | [optional] 
**LldpEnabled** | Pointer to **bool** |  | [optional] 
**Loopback** | Pointer to **bool** |  | [optional] 
**MaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**SecurityZone** | Pointer to **string** |  | [optional] 
**Subinterfaces** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValue**](V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValue.md) |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**TcpMssV4** | Pointer to **int32** |  | [optional] 
**TcpMssV6** | Pointer to **int32** |  | [optional] 
**V4TcpMss** | Pointer to [**V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV4TcpMss**](V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV4TcpMss.md) |  | [optional] 
**V6TcpMss** | Pointer to [**V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV6TcpMss**](V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV6TcpMss.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface

`func NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface() *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface`

NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface instantiates a new V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceWithDefaults() *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface`

NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetAdminStatus() bool`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetAdminStatusOk() (*bool, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetAdminStatus(v bool)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAlias

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetCircuit() string`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetCircuitOk() (*string, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetCircuit(v string)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpsec

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetIpsec() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetIpsecOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetIpsec(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.

### GetIpv4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetIpv4() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetIpv4Ok() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetIpv4(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetIpv6() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetIpv6Ok() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetIpv6(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetLan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.

### GetLoopback

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetLoopback() bool`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetLoopbackOk() (*bool, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetLoopback(v bool)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetSecurityZone

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetSecurityZone() string`

GetSecurityZone returns the SecurityZone field if non-nil, zero value otherwise.

### GetSecurityZoneOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetSecurityZoneOk() (*string, bool)`

GetSecurityZoneOk returns a tuple with the SecurityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityZone

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetSecurityZone(v string)`

SetSecurityZone sets SecurityZone field to given value.

### HasSecurityZone

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasSecurityZone() bool`

HasSecurityZone returns a boolean if a field has been set.

### GetSubinterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetSubinterfaces() map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValue`

GetSubinterfaces returns the Subinterfaces field if non-nil, zero value otherwise.

### GetSubinterfacesOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetSubinterfacesOk() (*map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValue, bool)`

GetSubinterfacesOk returns a tuple with the Subinterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubinterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetSubinterfaces(v map[string]V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValue)`

SetSubinterfaces sets Subinterfaces field to given value.

### HasSubinterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasSubinterfaces() bool`

HasSubinterfaces returns a boolean if a field has been set.

### GetTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetTcpMssV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetTcpMssV4() int32`

GetTcpMssV4 returns the TcpMssV4 field if non-nil, zero value otherwise.

### GetTcpMssV4Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetTcpMssV4Ok() (*int32, bool)`

GetTcpMssV4Ok returns a tuple with the TcpMssV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetTcpMssV4(v int32)`

SetTcpMssV4 sets TcpMssV4 field to given value.

### HasTcpMssV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasTcpMssV4() bool`

HasTcpMssV4 returns a boolean if a field has been set.

### GetTcpMssV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetTcpMssV6() int32`

GetTcpMssV6 returns the TcpMssV6 field if non-nil, zero value otherwise.

### GetTcpMssV6Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetTcpMssV6Ok() (*int32, bool)`

GetTcpMssV6Ok returns a tuple with the TcpMssV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetTcpMssV6(v int32)`

SetTcpMssV6 sets TcpMssV6 field to given value.

### HasTcpMssV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasTcpMssV6() bool`

HasTcpMssV6 returns a boolean if a field has been set.

### GetV4TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetV4TcpMss() V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV4TcpMss`

GetV4TcpMss returns the V4TcpMss field if non-nil, zero value otherwise.

### GetV4TcpMssOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetV4TcpMssOk() (*V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV4TcpMss, bool)`

GetV4TcpMssOk returns a tuple with the V4TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetV4TcpMss(v V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV4TcpMss)`

SetV4TcpMss sets V4TcpMss field to given value.

### HasV4TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasV4TcpMss() bool`

HasV4TcpMss returns a boolean if a field has been set.

### GetV6TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetV6TcpMss() V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV6TcpMss`

GetV6TcpMss returns the V6TcpMss field if non-nil, zero value otherwise.

### GetV6TcpMssOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) GetV6TcpMssOk() (*V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV6TcpMss, bool)`

GetV6TcpMssOk returns a tuple with the V6TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) SetV6TcpMss(v V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV6TcpMss)`

SetV6TcpMss sets V6TcpMss field to given value.

### HasV6TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterface) HasV6TcpMss() bool`

HasV6TcpMss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


