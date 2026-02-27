# ManaV2InterfaceVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Circuit** | Pointer to **string** |  | [optional] 
**ConfigUpdatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duplex** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**GatewayAddressV4** | Pointer to **string** |  | [optional] 
**GatewayAddressV6** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ipv4** | Pointer to [**ManaV2InterfaceAddress**](ManaV2InterfaceAddress.md) |  | [optional] 
**Ipv6** | Pointer to [**ManaV2InterfaceAddress**](ManaV2InterfaceAddress.md) |  | [optional] 
**Ipv6Addresses** | Pointer to [**[]ManaV2InterfaceAddress**](ManaV2InterfaceAddress.md) |  | [optional] 
**Lan** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**MaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OperUpdatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**ParentMacAddress** | Pointer to **string** |  | [optional] 
**SecurityZone** | Pointer to **string** |  | [optional] 
**SpeedMbps** | Pointer to **int32** |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**TcpMssV4** | Pointer to **int32** |  | [optional] 
**TcpMssV6** | Pointer to **int32** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**Vlan** | Pointer to **int32** |  | [optional] 

## Methods

### NewManaV2InterfaceVlan

`func NewManaV2InterfaceVlan() *ManaV2InterfaceVlan`

NewManaV2InterfaceVlan instantiates a new ManaV2InterfaceVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2InterfaceVlanWithDefaults

`func NewManaV2InterfaceVlanWithDefaults() *ManaV2InterfaceVlan`

NewManaV2InterfaceVlanWithDefaults instantiates a new ManaV2InterfaceVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ManaV2InterfaceVlan) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ManaV2InterfaceVlan) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ManaV2InterfaceVlan) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ManaV2InterfaceVlan) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCircuit

`func (o *ManaV2InterfaceVlan) GetCircuit() string`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *ManaV2InterfaceVlan) GetCircuitOk() (*string, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *ManaV2InterfaceVlan) SetCircuit(v string)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *ManaV2InterfaceVlan) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetConfigUpdatedAt

`func (o *ManaV2InterfaceVlan) GetConfigUpdatedAt() GoogleProtobufTimestamp`

GetConfigUpdatedAt returns the ConfigUpdatedAt field if non-nil, zero value otherwise.

### GetConfigUpdatedAtOk

`func (o *ManaV2InterfaceVlan) GetConfigUpdatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetConfigUpdatedAtOk returns a tuple with the ConfigUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUpdatedAt

`func (o *ManaV2InterfaceVlan) SetConfigUpdatedAt(v GoogleProtobufTimestamp)`

SetConfigUpdatedAt sets ConfigUpdatedAt field to given value.

### HasConfigUpdatedAt

`func (o *ManaV2InterfaceVlan) HasConfigUpdatedAt() bool`

HasConfigUpdatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ManaV2InterfaceVlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2InterfaceVlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2InterfaceVlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2InterfaceVlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuplex

`func (o *ManaV2InterfaceVlan) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *ManaV2InterfaceVlan) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *ManaV2InterfaceVlan) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *ManaV2InterfaceVlan) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEnabled

`func (o *ManaV2InterfaceVlan) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ManaV2InterfaceVlan) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ManaV2InterfaceVlan) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ManaV2InterfaceVlan) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGatewayAddressV4

`func (o *ManaV2InterfaceVlan) GetGatewayAddressV4() string`

GetGatewayAddressV4 returns the GatewayAddressV4 field if non-nil, zero value otherwise.

### GetGatewayAddressV4Ok

`func (o *ManaV2InterfaceVlan) GetGatewayAddressV4Ok() (*string, bool)`

GetGatewayAddressV4Ok returns a tuple with the GatewayAddressV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddressV4

`func (o *ManaV2InterfaceVlan) SetGatewayAddressV4(v string)`

SetGatewayAddressV4 sets GatewayAddressV4 field to given value.

### HasGatewayAddressV4

`func (o *ManaV2InterfaceVlan) HasGatewayAddressV4() bool`

HasGatewayAddressV4 returns a boolean if a field has been set.

### GetGatewayAddressV6

`func (o *ManaV2InterfaceVlan) GetGatewayAddressV6() string`

GetGatewayAddressV6 returns the GatewayAddressV6 field if non-nil, zero value otherwise.

### GetGatewayAddressV6Ok

`func (o *ManaV2InterfaceVlan) GetGatewayAddressV6Ok() (*string, bool)`

GetGatewayAddressV6Ok returns a tuple with the GatewayAddressV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddressV6

`func (o *ManaV2InterfaceVlan) SetGatewayAddressV6(v string)`

SetGatewayAddressV6 sets GatewayAddressV6 field to given value.

### HasGatewayAddressV6

`func (o *ManaV2InterfaceVlan) HasGatewayAddressV6() bool`

HasGatewayAddressV6 returns a boolean if a field has been set.

### GetId

`func (o *ManaV2InterfaceVlan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2InterfaceVlan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2InterfaceVlan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ManaV2InterfaceVlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpv4

`func (o *ManaV2InterfaceVlan) GetIpv4() ManaV2InterfaceAddress`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *ManaV2InterfaceVlan) GetIpv4Ok() (*ManaV2InterfaceAddress, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *ManaV2InterfaceVlan) SetIpv4(v ManaV2InterfaceAddress)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *ManaV2InterfaceVlan) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *ManaV2InterfaceVlan) GetIpv6() ManaV2InterfaceAddress`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *ManaV2InterfaceVlan) GetIpv6Ok() (*ManaV2InterfaceAddress, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *ManaV2InterfaceVlan) SetIpv6(v ManaV2InterfaceAddress)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *ManaV2InterfaceVlan) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *ManaV2InterfaceVlan) GetIpv6Addresses() []ManaV2InterfaceAddress`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *ManaV2InterfaceVlan) GetIpv6AddressesOk() (*[]ManaV2InterfaceAddress, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *ManaV2InterfaceVlan) SetIpv6Addresses(v []ManaV2InterfaceAddress)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *ManaV2InterfaceVlan) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetLan

`func (o *ManaV2InterfaceVlan) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *ManaV2InterfaceVlan) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *ManaV2InterfaceVlan) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *ManaV2InterfaceVlan) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetMacAddress

`func (o *ManaV2InterfaceVlan) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ManaV2InterfaceVlan) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ManaV2InterfaceVlan) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ManaV2InterfaceVlan) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *ManaV2InterfaceVlan) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *ManaV2InterfaceVlan) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *ManaV2InterfaceVlan) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *ManaV2InterfaceVlan) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetName

`func (o *ManaV2InterfaceVlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2InterfaceVlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2InterfaceVlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2InterfaceVlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperUpdatedAt

`func (o *ManaV2InterfaceVlan) GetOperUpdatedAt() GoogleProtobufTimestamp`

GetOperUpdatedAt returns the OperUpdatedAt field if non-nil, zero value otherwise.

### GetOperUpdatedAtOk

`func (o *ManaV2InterfaceVlan) GetOperUpdatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetOperUpdatedAtOk returns a tuple with the OperUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperUpdatedAt

`func (o *ManaV2InterfaceVlan) SetOperUpdatedAt(v GoogleProtobufTimestamp)`

SetOperUpdatedAt sets OperUpdatedAt field to given value.

### HasOperUpdatedAt

`func (o *ManaV2InterfaceVlan) HasOperUpdatedAt() bool`

HasOperUpdatedAt returns a boolean if a field has been set.

### GetParentMacAddress

`func (o *ManaV2InterfaceVlan) GetParentMacAddress() string`

GetParentMacAddress returns the ParentMacAddress field if non-nil, zero value otherwise.

### GetParentMacAddressOk

`func (o *ManaV2InterfaceVlan) GetParentMacAddressOk() (*string, bool)`

GetParentMacAddressOk returns a tuple with the ParentMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentMacAddress

`func (o *ManaV2InterfaceVlan) SetParentMacAddress(v string)`

SetParentMacAddress sets ParentMacAddress field to given value.

### HasParentMacAddress

`func (o *ManaV2InterfaceVlan) HasParentMacAddress() bool`

HasParentMacAddress returns a boolean if a field has been set.

### GetSecurityZone

`func (o *ManaV2InterfaceVlan) GetSecurityZone() string`

GetSecurityZone returns the SecurityZone field if non-nil, zero value otherwise.

### GetSecurityZoneOk

`func (o *ManaV2InterfaceVlan) GetSecurityZoneOk() (*string, bool)`

GetSecurityZoneOk returns a tuple with the SecurityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityZone

`func (o *ManaV2InterfaceVlan) SetSecurityZone(v string)`

SetSecurityZone sets SecurityZone field to given value.

### HasSecurityZone

`func (o *ManaV2InterfaceVlan) HasSecurityZone() bool`

HasSecurityZone returns a boolean if a field has been set.

### GetSpeedMbps

`func (o *ManaV2InterfaceVlan) GetSpeedMbps() int32`

GetSpeedMbps returns the SpeedMbps field if non-nil, zero value otherwise.

### GetSpeedMbpsOk

`func (o *ManaV2InterfaceVlan) GetSpeedMbpsOk() (*int32, bool)`

GetSpeedMbpsOk returns a tuple with the SpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMbps

`func (o *ManaV2InterfaceVlan) SetSpeedMbps(v int32)`

SetSpeedMbps sets SpeedMbps field to given value.

### HasSpeedMbps

`func (o *ManaV2InterfaceVlan) HasSpeedMbps() bool`

HasSpeedMbps returns a boolean if a field has been set.

### GetTcpMss

`func (o *ManaV2InterfaceVlan) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *ManaV2InterfaceVlan) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *ManaV2InterfaceVlan) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *ManaV2InterfaceVlan) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetTcpMssV4

`func (o *ManaV2InterfaceVlan) GetTcpMssV4() int32`

GetTcpMssV4 returns the TcpMssV4 field if non-nil, zero value otherwise.

### GetTcpMssV4Ok

`func (o *ManaV2InterfaceVlan) GetTcpMssV4Ok() (*int32, bool)`

GetTcpMssV4Ok returns a tuple with the TcpMssV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV4

`func (o *ManaV2InterfaceVlan) SetTcpMssV4(v int32)`

SetTcpMssV4 sets TcpMssV4 field to given value.

### HasTcpMssV4

`func (o *ManaV2InterfaceVlan) HasTcpMssV4() bool`

HasTcpMssV4 returns a boolean if a field has been set.

### GetTcpMssV6

`func (o *ManaV2InterfaceVlan) GetTcpMssV6() int32`

GetTcpMssV6 returns the TcpMssV6 field if non-nil, zero value otherwise.

### GetTcpMssV6Ok

`func (o *ManaV2InterfaceVlan) GetTcpMssV6Ok() (*int32, bool)`

GetTcpMssV6Ok returns a tuple with the TcpMssV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV6

`func (o *ManaV2InterfaceVlan) SetTcpMssV6(v int32)`

SetTcpMssV6 sets TcpMssV6 field to given value.

### HasTcpMssV6

`func (o *ManaV2InterfaceVlan) HasTcpMssV6() bool`

HasTcpMssV6 returns a boolean if a field has been set.

### GetUp

`func (o *ManaV2InterfaceVlan) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *ManaV2InterfaceVlan) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *ManaV2InterfaceVlan) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *ManaV2InterfaceVlan) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetVlan

`func (o *ManaV2InterfaceVlan) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *ManaV2InterfaceVlan) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *ManaV2InterfaceVlan) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *ManaV2InterfaceVlan) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


