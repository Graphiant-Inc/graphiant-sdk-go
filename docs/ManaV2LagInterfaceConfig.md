# ManaV2LagInterfaceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **bool** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Ipv4** | Pointer to [**ManaV2InterfaceIpConfig**](ManaV2InterfaceIpConfig.md) |  | [optional] 
**Ipv6** | Pointer to [**ManaV2InterfaceIpConfig**](ManaV2InterfaceIpConfig.md) |  | [optional] 
**Lacp** | Pointer to [**ManaV2LacpConfig**](ManaV2LacpConfig.md) |  | [optional] 
**LagMembers** | Pointer to [**map[string]ManaV2NullableLagMemberInterface**](ManaV2NullableLagMemberInterface.md) |  | [optional] 
**Macsec** | Pointer to [**ManaV2NullableMaCsecConfiguration**](ManaV2NullableMaCsecConfiguration.md) |  | [optional] 
**MinimumMembers** | Pointer to **int32** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Segment** | Pointer to **string** |  | [optional] 
**Subinterfaces** | Pointer to [**map[string]ManaV2NullableInterfaceLagvlanConfig**](ManaV2NullableInterfaceLagvlanConfig.md) |  | [optional] 

## Methods

### NewManaV2LagInterfaceConfig

`func NewManaV2LagInterfaceConfig() *ManaV2LagInterfaceConfig`

NewManaV2LagInterfaceConfig instantiates a new ManaV2LagInterfaceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2LagInterfaceConfigWithDefaults

`func NewManaV2LagInterfaceConfigWithDefaults() *ManaV2LagInterfaceConfig`

NewManaV2LagInterfaceConfigWithDefaults instantiates a new ManaV2LagInterfaceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *ManaV2LagInterfaceConfig) GetAdminStatus() bool`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *ManaV2LagInterfaceConfig) GetAdminStatusOk() (*bool, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *ManaV2LagInterfaceConfig) SetAdminStatus(v bool)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *ManaV2LagInterfaceConfig) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAlias

`func (o *ManaV2LagInterfaceConfig) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ManaV2LagInterfaceConfig) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ManaV2LagInterfaceConfig) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ManaV2LagInterfaceConfig) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDescription

`func (o *ManaV2LagInterfaceConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2LagInterfaceConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2LagInterfaceConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2LagInterfaceConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *ManaV2LagInterfaceConfig) GetIpv4() ManaV2InterfaceIpConfig`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *ManaV2LagInterfaceConfig) GetIpv4Ok() (*ManaV2InterfaceIpConfig, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *ManaV2LagInterfaceConfig) SetIpv4(v ManaV2InterfaceIpConfig)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *ManaV2LagInterfaceConfig) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *ManaV2LagInterfaceConfig) GetIpv6() ManaV2InterfaceIpConfig`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *ManaV2LagInterfaceConfig) GetIpv6Ok() (*ManaV2InterfaceIpConfig, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *ManaV2LagInterfaceConfig) SetIpv6(v ManaV2InterfaceIpConfig)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *ManaV2LagInterfaceConfig) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetLacp

`func (o *ManaV2LagInterfaceConfig) GetLacp() ManaV2LacpConfig`

GetLacp returns the Lacp field if non-nil, zero value otherwise.

### GetLacpOk

`func (o *ManaV2LagInterfaceConfig) GetLacpOk() (*ManaV2LacpConfig, bool)`

GetLacpOk returns a tuple with the Lacp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLacp

`func (o *ManaV2LagInterfaceConfig) SetLacp(v ManaV2LacpConfig)`

SetLacp sets Lacp field to given value.

### HasLacp

`func (o *ManaV2LagInterfaceConfig) HasLacp() bool`

HasLacp returns a boolean if a field has been set.

### GetLagMembers

`func (o *ManaV2LagInterfaceConfig) GetLagMembers() map[string]ManaV2NullableLagMemberInterface`

GetLagMembers returns the LagMembers field if non-nil, zero value otherwise.

### GetLagMembersOk

`func (o *ManaV2LagInterfaceConfig) GetLagMembersOk() (*map[string]ManaV2NullableLagMemberInterface, bool)`

GetLagMembersOk returns a tuple with the LagMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagMembers

`func (o *ManaV2LagInterfaceConfig) SetLagMembers(v map[string]ManaV2NullableLagMemberInterface)`

SetLagMembers sets LagMembers field to given value.

### HasLagMembers

`func (o *ManaV2LagInterfaceConfig) HasLagMembers() bool`

HasLagMembers returns a boolean if a field has been set.

### GetMacsec

`func (o *ManaV2LagInterfaceConfig) GetMacsec() ManaV2NullableMaCsecConfiguration`

GetMacsec returns the Macsec field if non-nil, zero value otherwise.

### GetMacsecOk

`func (o *ManaV2LagInterfaceConfig) GetMacsecOk() (*ManaV2NullableMaCsecConfiguration, bool)`

GetMacsecOk returns a tuple with the Macsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacsec

`func (o *ManaV2LagInterfaceConfig) SetMacsec(v ManaV2NullableMaCsecConfiguration)`

SetMacsec sets Macsec field to given value.

### HasMacsec

`func (o *ManaV2LagInterfaceConfig) HasMacsec() bool`

HasMacsec returns a boolean if a field has been set.

### GetMinimumMembers

`func (o *ManaV2LagInterfaceConfig) GetMinimumMembers() int32`

GetMinimumMembers returns the MinimumMembers field if non-nil, zero value otherwise.

### GetMinimumMembersOk

`func (o *ManaV2LagInterfaceConfig) GetMinimumMembersOk() (*int32, bool)`

GetMinimumMembersOk returns a tuple with the MinimumMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumMembers

`func (o *ManaV2LagInterfaceConfig) SetMinimumMembers(v int32)`

SetMinimumMembers sets MinimumMembers field to given value.

### HasMinimumMembers

`func (o *ManaV2LagInterfaceConfig) HasMinimumMembers() bool`

HasMinimumMembers returns a boolean if a field has been set.

### GetMtu

`func (o *ManaV2LagInterfaceConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *ManaV2LagInterfaceConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *ManaV2LagInterfaceConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *ManaV2LagInterfaceConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetSegment

`func (o *ManaV2LagInterfaceConfig) GetSegment() string`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *ManaV2LagInterfaceConfig) GetSegmentOk() (*string, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *ManaV2LagInterfaceConfig) SetSegment(v string)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *ManaV2LagInterfaceConfig) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetSubinterfaces

`func (o *ManaV2LagInterfaceConfig) GetSubinterfaces() map[string]ManaV2NullableInterfaceLagvlanConfig`

GetSubinterfaces returns the Subinterfaces field if non-nil, zero value otherwise.

### GetSubinterfacesOk

`func (o *ManaV2LagInterfaceConfig) GetSubinterfacesOk() (*map[string]ManaV2NullableInterfaceLagvlanConfig, bool)`

GetSubinterfacesOk returns a tuple with the Subinterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubinterfaces

`func (o *ManaV2LagInterfaceConfig) SetSubinterfaces(v map[string]ManaV2NullableInterfaceLagvlanConfig)`

SetSubinterfaces sets Subinterfaces field to given value.

### HasSubinterfaces

`func (o *ManaV2LagInterfaceConfig) HasSubinterfaces() bool`

HasSubinterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


