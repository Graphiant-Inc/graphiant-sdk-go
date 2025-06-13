# V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **bool** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Circuit** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Ipv4** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw.md) |  | [optional] 
**Ipv6** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw.md) |  | [optional] 
**Lan** | Pointer to **string** |  | [optional] 
**LldpEnabled** | Pointer to **bool** |  | [optional] 
**MaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**SecurityZone** | Pointer to **string** |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**TcpMssV4** | Pointer to **int32** |  | [optional] 
**TcpMssV6** | Pointer to **int32** |  | [optional] 
**V4TcpMss** | Pointer to [**V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV4TcpMss**](V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV4TcpMss.md) |  | [optional] 
**V6TcpMss** | Pointer to [**V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV6TcpMss**](V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV6TcpMss.md) |  | [optional] 
**Vlan** | Pointer to **int32** |  | [optional] 
**Vrrp** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrp**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrp.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface

`func NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface() *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface`

NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface instantiates a new V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceWithDefaults() *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface`

NewV1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetAdminStatus() bool`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetAdminStatusOk() (*bool, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetAdminStatus(v bool)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAlias

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetCircuit() string`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetCircuitOk() (*string, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetCircuit(v string)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetIpv4() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetIpv4Ok() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetIpv4(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetIpv6() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetIpv6Ok() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetIpv6(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGw)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetLan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetSecurityZone

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetSecurityZone() string`

GetSecurityZone returns the SecurityZone field if non-nil, zero value otherwise.

### GetSecurityZoneOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetSecurityZoneOk() (*string, bool)`

GetSecurityZoneOk returns a tuple with the SecurityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityZone

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetSecurityZone(v string)`

SetSecurityZone sets SecurityZone field to given value.

### HasSecurityZone

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasSecurityZone() bool`

HasSecurityZone returns a boolean if a field has been set.

### GetTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetTcpMssV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetTcpMssV4() int32`

GetTcpMssV4 returns the TcpMssV4 field if non-nil, zero value otherwise.

### GetTcpMssV4Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetTcpMssV4Ok() (*int32, bool)`

GetTcpMssV4Ok returns a tuple with the TcpMssV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetTcpMssV4(v int32)`

SetTcpMssV4 sets TcpMssV4 field to given value.

### HasTcpMssV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasTcpMssV4() bool`

HasTcpMssV4 returns a boolean if a field has been set.

### GetTcpMssV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetTcpMssV6() int32`

GetTcpMssV6 returns the TcpMssV6 field if non-nil, zero value otherwise.

### GetTcpMssV6Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetTcpMssV6Ok() (*int32, bool)`

GetTcpMssV6Ok returns a tuple with the TcpMssV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetTcpMssV6(v int32)`

SetTcpMssV6 sets TcpMssV6 field to given value.

### HasTcpMssV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasTcpMssV6() bool`

HasTcpMssV6 returns a boolean if a field has been set.

### GetV4TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetV4TcpMss() V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV4TcpMss`

GetV4TcpMss returns the V4TcpMss field if non-nil, zero value otherwise.

### GetV4TcpMssOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetV4TcpMssOk() (*V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV4TcpMss, bool)`

GetV4TcpMssOk returns a tuple with the V4TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetV4TcpMss(v V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV4TcpMss)`

SetV4TcpMss sets V4TcpMss field to given value.

### HasV4TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasV4TcpMss() bool`

HasV4TcpMss returns a boolean if a field has been set.

### GetV6TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetV6TcpMss() V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV6TcpMss`

GetV6TcpMss returns the V6TcpMss field if non-nil, zero value otherwise.

### GetV6TcpMssOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetV6TcpMssOk() (*V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV6TcpMss, bool)`

GetV6TcpMssOk returns a tuple with the V6TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetV6TcpMss(v V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterfaceV6TcpMss)`

SetV6TcpMss sets V6TcpMss field to given value.

### HasV6TcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasV6TcpMss() bool`

HasV6TcpMss returns a boolean if a field has been set.

### GetVlan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVrrp

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetVrrp() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrp`

GetVrrp returns the Vrrp field if non-nil, zero value otherwise.

### GetVrrpOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) GetVrrpOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrp, bool)`

GetVrrpOk returns a tuple with the Vrrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrp

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) SetVrrp(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrp)`

SetVrrp sets Vrrp field to given value.

### HasVrrp

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeInterfacesValueInterfaceSubinterfacesValueInterface) HasVrrp() bool`

HasVrrp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


