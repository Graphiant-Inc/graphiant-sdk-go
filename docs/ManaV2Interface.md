# ManaV2Interface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Circuit** | Pointer to **string** |  | [optional] 
**CircuitName** | Pointer to **string** |  | [optional] 
**ConfigUpdatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**ConfiguredMaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duplex** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**GatewayAddressV4** | Pointer to **string** |  | [optional] 
**GatewayAddressV6** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IpSec** | Pointer to [**ManaV2InterfaceIPsec**](ManaV2InterfaceIPsec.md) |  | [optional] 
**Ipv4** | Pointer to [**ManaV2InterfaceAddress**](ManaV2InterfaceAddress.md) |  | [optional] 
**Ipv6** | Pointer to [**ManaV2InterfaceAddress**](ManaV2InterfaceAddress.md) |  | [optional] 
**Ipv6Addresses** | Pointer to [**[]ManaV2InterfaceAddress**](ManaV2InterfaceAddress.md) |  | [optional] 
**LagInterface** | Pointer to [**ManaV2LagInterface**](ManaV2LagInterface.md) |  | [optional] 
**Lan** | Pointer to **string** |  | [optional] 
**LldpEnabled** | Pointer to **bool** |  | [optional] 
**Macsec** | Pointer to [**ManaV2InterfaceMaCsec**](ManaV2InterfaceMaCsec.md) |  | [optional] 
**MaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OperUpdatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**PhyAddress** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**SecurityZone** | Pointer to **string** |  | [optional] 
**SfpOpticalStrength** | Pointer to [**[]ManaV2InterfaceSfpOpticalStrength**](ManaV2InterfaceSfpOpticalStrength.md) |  | [optional] 
**SpeedMbps** | Pointer to **int32** |  | [optional] 
**Subinterfaces** | Pointer to [**[]ManaV2InterfaceVlan**](ManaV2InterfaceVlan.md) |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**TcpMssV4** | Pointer to **int32** |  | [optional] 
**TcpMssV6** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**VrfFunctionId** | Pointer to **int32** |  | [optional] 
**VrfName** | Pointer to **string** |  | [optional] 

## Methods

### NewManaV2Interface

`func NewManaV2Interface() *ManaV2Interface`

NewManaV2Interface instantiates a new ManaV2Interface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2InterfaceWithDefaults

`func NewManaV2InterfaceWithDefaults() *ManaV2Interface`

NewManaV2InterfaceWithDefaults instantiates a new ManaV2Interface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ManaV2Interface) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ManaV2Interface) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ManaV2Interface) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ManaV2Interface) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCircuit

`func (o *ManaV2Interface) GetCircuit() string`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *ManaV2Interface) GetCircuitOk() (*string, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *ManaV2Interface) SetCircuit(v string)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *ManaV2Interface) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetCircuitName

`func (o *ManaV2Interface) GetCircuitName() string`

GetCircuitName returns the CircuitName field if non-nil, zero value otherwise.

### GetCircuitNameOk

`func (o *ManaV2Interface) GetCircuitNameOk() (*string, bool)`

GetCircuitNameOk returns a tuple with the CircuitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitName

`func (o *ManaV2Interface) SetCircuitName(v string)`

SetCircuitName sets CircuitName field to given value.

### HasCircuitName

`func (o *ManaV2Interface) HasCircuitName() bool`

HasCircuitName returns a boolean if a field has been set.

### GetConfigUpdatedAt

`func (o *ManaV2Interface) GetConfigUpdatedAt() GoogleProtobufTimestamp`

GetConfigUpdatedAt returns the ConfigUpdatedAt field if non-nil, zero value otherwise.

### GetConfigUpdatedAtOk

`func (o *ManaV2Interface) GetConfigUpdatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetConfigUpdatedAtOk returns a tuple with the ConfigUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUpdatedAt

`func (o *ManaV2Interface) SetConfigUpdatedAt(v GoogleProtobufTimestamp)`

SetConfigUpdatedAt sets ConfigUpdatedAt field to given value.

### HasConfigUpdatedAt

`func (o *ManaV2Interface) HasConfigUpdatedAt() bool`

HasConfigUpdatedAt returns a boolean if a field has been set.

### GetConfiguredMaxTransmissionUnit

`func (o *ManaV2Interface) GetConfiguredMaxTransmissionUnit() int32`

GetConfiguredMaxTransmissionUnit returns the ConfiguredMaxTransmissionUnit field if non-nil, zero value otherwise.

### GetConfiguredMaxTransmissionUnitOk

`func (o *ManaV2Interface) GetConfiguredMaxTransmissionUnitOk() (*int32, bool)`

GetConfiguredMaxTransmissionUnitOk returns a tuple with the ConfiguredMaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredMaxTransmissionUnit

`func (o *ManaV2Interface) SetConfiguredMaxTransmissionUnit(v int32)`

SetConfiguredMaxTransmissionUnit sets ConfiguredMaxTransmissionUnit field to given value.

### HasConfiguredMaxTransmissionUnit

`func (o *ManaV2Interface) HasConfiguredMaxTransmissionUnit() bool`

HasConfiguredMaxTransmissionUnit returns a boolean if a field has been set.

### GetDescription

`func (o *ManaV2Interface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2Interface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2Interface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2Interface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuplex

`func (o *ManaV2Interface) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *ManaV2Interface) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *ManaV2Interface) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *ManaV2Interface) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEnabled

`func (o *ManaV2Interface) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ManaV2Interface) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ManaV2Interface) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ManaV2Interface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGatewayAddressV4

`func (o *ManaV2Interface) GetGatewayAddressV4() string`

GetGatewayAddressV4 returns the GatewayAddressV4 field if non-nil, zero value otherwise.

### GetGatewayAddressV4Ok

`func (o *ManaV2Interface) GetGatewayAddressV4Ok() (*string, bool)`

GetGatewayAddressV4Ok returns a tuple with the GatewayAddressV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddressV4

`func (o *ManaV2Interface) SetGatewayAddressV4(v string)`

SetGatewayAddressV4 sets GatewayAddressV4 field to given value.

### HasGatewayAddressV4

`func (o *ManaV2Interface) HasGatewayAddressV4() bool`

HasGatewayAddressV4 returns a boolean if a field has been set.

### GetGatewayAddressV6

`func (o *ManaV2Interface) GetGatewayAddressV6() string`

GetGatewayAddressV6 returns the GatewayAddressV6 field if non-nil, zero value otherwise.

### GetGatewayAddressV6Ok

`func (o *ManaV2Interface) GetGatewayAddressV6Ok() (*string, bool)`

GetGatewayAddressV6Ok returns a tuple with the GatewayAddressV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddressV6

`func (o *ManaV2Interface) SetGatewayAddressV6(v string)`

SetGatewayAddressV6 sets GatewayAddressV6 field to given value.

### HasGatewayAddressV6

`func (o *ManaV2Interface) HasGatewayAddressV6() bool`

HasGatewayAddressV6 returns a boolean if a field has been set.

### GetId

`func (o *ManaV2Interface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2Interface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2Interface) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ManaV2Interface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpSec

`func (o *ManaV2Interface) GetIpSec() ManaV2InterfaceIPsec`

GetIpSec returns the IpSec field if non-nil, zero value otherwise.

### GetIpSecOk

`func (o *ManaV2Interface) GetIpSecOk() (*ManaV2InterfaceIPsec, bool)`

GetIpSecOk returns a tuple with the IpSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSec

`func (o *ManaV2Interface) SetIpSec(v ManaV2InterfaceIPsec)`

SetIpSec sets IpSec field to given value.

### HasIpSec

`func (o *ManaV2Interface) HasIpSec() bool`

HasIpSec returns a boolean if a field has been set.

### GetIpv4

`func (o *ManaV2Interface) GetIpv4() ManaV2InterfaceAddress`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *ManaV2Interface) GetIpv4Ok() (*ManaV2InterfaceAddress, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *ManaV2Interface) SetIpv4(v ManaV2InterfaceAddress)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *ManaV2Interface) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *ManaV2Interface) GetIpv6() ManaV2InterfaceAddress`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *ManaV2Interface) GetIpv6Ok() (*ManaV2InterfaceAddress, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *ManaV2Interface) SetIpv6(v ManaV2InterfaceAddress)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *ManaV2Interface) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *ManaV2Interface) GetIpv6Addresses() []ManaV2InterfaceAddress`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *ManaV2Interface) GetIpv6AddressesOk() (*[]ManaV2InterfaceAddress, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *ManaV2Interface) SetIpv6Addresses(v []ManaV2InterfaceAddress)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *ManaV2Interface) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetLagInterface

`func (o *ManaV2Interface) GetLagInterface() ManaV2LagInterface`

GetLagInterface returns the LagInterface field if non-nil, zero value otherwise.

### GetLagInterfaceOk

`func (o *ManaV2Interface) GetLagInterfaceOk() (*ManaV2LagInterface, bool)`

GetLagInterfaceOk returns a tuple with the LagInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagInterface

`func (o *ManaV2Interface) SetLagInterface(v ManaV2LagInterface)`

SetLagInterface sets LagInterface field to given value.

### HasLagInterface

`func (o *ManaV2Interface) HasLagInterface() bool`

HasLagInterface returns a boolean if a field has been set.

### GetLan

`func (o *ManaV2Interface) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *ManaV2Interface) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *ManaV2Interface) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *ManaV2Interface) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLldpEnabled

`func (o *ManaV2Interface) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *ManaV2Interface) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *ManaV2Interface) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *ManaV2Interface) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.

### GetMacsec

`func (o *ManaV2Interface) GetMacsec() ManaV2InterfaceMaCsec`

GetMacsec returns the Macsec field if non-nil, zero value otherwise.

### GetMacsecOk

`func (o *ManaV2Interface) GetMacsecOk() (*ManaV2InterfaceMaCsec, bool)`

GetMacsecOk returns a tuple with the Macsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacsec

`func (o *ManaV2Interface) SetMacsec(v ManaV2InterfaceMaCsec)`

SetMacsec sets Macsec field to given value.

### HasMacsec

`func (o *ManaV2Interface) HasMacsec() bool`

HasMacsec returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *ManaV2Interface) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *ManaV2Interface) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *ManaV2Interface) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *ManaV2Interface) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetName

`func (o *ManaV2Interface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2Interface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2Interface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2Interface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperUpdatedAt

`func (o *ManaV2Interface) GetOperUpdatedAt() GoogleProtobufTimestamp`

GetOperUpdatedAt returns the OperUpdatedAt field if non-nil, zero value otherwise.

### GetOperUpdatedAtOk

`func (o *ManaV2Interface) GetOperUpdatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetOperUpdatedAtOk returns a tuple with the OperUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperUpdatedAt

`func (o *ManaV2Interface) SetOperUpdatedAt(v GoogleProtobufTimestamp)`

SetOperUpdatedAt sets OperUpdatedAt field to given value.

### HasOperUpdatedAt

`func (o *ManaV2Interface) HasOperUpdatedAt() bool`

HasOperUpdatedAt returns a boolean if a field has been set.

### GetPhyAddress

`func (o *ManaV2Interface) GetPhyAddress() string`

GetPhyAddress returns the PhyAddress field if non-nil, zero value otherwise.

### GetPhyAddressOk

`func (o *ManaV2Interface) GetPhyAddressOk() (*string, bool)`

GetPhyAddressOk returns a tuple with the PhyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhyAddress

`func (o *ManaV2Interface) SetPhyAddress(v string)`

SetPhyAddress sets PhyAddress field to given value.

### HasPhyAddress

`func (o *ManaV2Interface) HasPhyAddress() bool`

HasPhyAddress returns a boolean if a field has been set.

### GetProtocol

`func (o *ManaV2Interface) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ManaV2Interface) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ManaV2Interface) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ManaV2Interface) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSecurityZone

`func (o *ManaV2Interface) GetSecurityZone() string`

GetSecurityZone returns the SecurityZone field if non-nil, zero value otherwise.

### GetSecurityZoneOk

`func (o *ManaV2Interface) GetSecurityZoneOk() (*string, bool)`

GetSecurityZoneOk returns a tuple with the SecurityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityZone

`func (o *ManaV2Interface) SetSecurityZone(v string)`

SetSecurityZone sets SecurityZone field to given value.

### HasSecurityZone

`func (o *ManaV2Interface) HasSecurityZone() bool`

HasSecurityZone returns a boolean if a field has been set.

### GetSfpOpticalStrength

`func (o *ManaV2Interface) GetSfpOpticalStrength() []ManaV2InterfaceSfpOpticalStrength`

GetSfpOpticalStrength returns the SfpOpticalStrength field if non-nil, zero value otherwise.

### GetSfpOpticalStrengthOk

`func (o *ManaV2Interface) GetSfpOpticalStrengthOk() (*[]ManaV2InterfaceSfpOpticalStrength, bool)`

GetSfpOpticalStrengthOk returns a tuple with the SfpOpticalStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpOpticalStrength

`func (o *ManaV2Interface) SetSfpOpticalStrength(v []ManaV2InterfaceSfpOpticalStrength)`

SetSfpOpticalStrength sets SfpOpticalStrength field to given value.

### HasSfpOpticalStrength

`func (o *ManaV2Interface) HasSfpOpticalStrength() bool`

HasSfpOpticalStrength returns a boolean if a field has been set.

### GetSpeedMbps

`func (o *ManaV2Interface) GetSpeedMbps() int32`

GetSpeedMbps returns the SpeedMbps field if non-nil, zero value otherwise.

### GetSpeedMbpsOk

`func (o *ManaV2Interface) GetSpeedMbpsOk() (*int32, bool)`

GetSpeedMbpsOk returns a tuple with the SpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMbps

`func (o *ManaV2Interface) SetSpeedMbps(v int32)`

SetSpeedMbps sets SpeedMbps field to given value.

### HasSpeedMbps

`func (o *ManaV2Interface) HasSpeedMbps() bool`

HasSpeedMbps returns a boolean if a field has been set.

### GetSubinterfaces

`func (o *ManaV2Interface) GetSubinterfaces() []ManaV2InterfaceVlan`

GetSubinterfaces returns the Subinterfaces field if non-nil, zero value otherwise.

### GetSubinterfacesOk

`func (o *ManaV2Interface) GetSubinterfacesOk() (*[]ManaV2InterfaceVlan, bool)`

GetSubinterfacesOk returns a tuple with the Subinterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubinterfaces

`func (o *ManaV2Interface) SetSubinterfaces(v []ManaV2InterfaceVlan)`

SetSubinterfaces sets Subinterfaces field to given value.

### HasSubinterfaces

`func (o *ManaV2Interface) HasSubinterfaces() bool`

HasSubinterfaces returns a boolean if a field has been set.

### GetTcpMss

`func (o *ManaV2Interface) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *ManaV2Interface) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *ManaV2Interface) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *ManaV2Interface) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetTcpMssV4

`func (o *ManaV2Interface) GetTcpMssV4() int32`

GetTcpMssV4 returns the TcpMssV4 field if non-nil, zero value otherwise.

### GetTcpMssV4Ok

`func (o *ManaV2Interface) GetTcpMssV4Ok() (*int32, bool)`

GetTcpMssV4Ok returns a tuple with the TcpMssV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV4

`func (o *ManaV2Interface) SetTcpMssV4(v int32)`

SetTcpMssV4 sets TcpMssV4 field to given value.

### HasTcpMssV4

`func (o *ManaV2Interface) HasTcpMssV4() bool`

HasTcpMssV4 returns a boolean if a field has been set.

### GetTcpMssV6

`func (o *ManaV2Interface) GetTcpMssV6() int32`

GetTcpMssV6 returns the TcpMssV6 field if non-nil, zero value otherwise.

### GetTcpMssV6Ok

`func (o *ManaV2Interface) GetTcpMssV6Ok() (*int32, bool)`

GetTcpMssV6Ok returns a tuple with the TcpMssV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMssV6

`func (o *ManaV2Interface) SetTcpMssV6(v int32)`

SetTcpMssV6 sets TcpMssV6 field to given value.

### HasTcpMssV6

`func (o *ManaV2Interface) HasTcpMssV6() bool`

HasTcpMssV6 returns a boolean if a field has been set.

### GetType

`func (o *ManaV2Interface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManaV2Interface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManaV2Interface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManaV2Interface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUp

`func (o *ManaV2Interface) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *ManaV2Interface) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *ManaV2Interface) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *ManaV2Interface) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetVrfFunctionId

`func (o *ManaV2Interface) GetVrfFunctionId() int32`

GetVrfFunctionId returns the VrfFunctionId field if non-nil, zero value otherwise.

### GetVrfFunctionIdOk

`func (o *ManaV2Interface) GetVrfFunctionIdOk() (*int32, bool)`

GetVrfFunctionIdOk returns a tuple with the VrfFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfFunctionId

`func (o *ManaV2Interface) SetVrfFunctionId(v int32)`

SetVrfFunctionId sets VrfFunctionId field to given value.

### HasVrfFunctionId

`func (o *ManaV2Interface) HasVrfFunctionId() bool`

HasVrfFunctionId returns a boolean if a field has been set.

### GetVrfName

`func (o *ManaV2Interface) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *ManaV2Interface) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *ManaV2Interface) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *ManaV2Interface) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


