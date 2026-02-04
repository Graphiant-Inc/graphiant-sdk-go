# ManaV2InterfaceMaCsec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**EncryptionEnforcementMode** | Pointer to **string** |  | [optional] 
**KeyServerPriority** | Pointer to **int64** |  | [optional] 
**PskConfigurations** | Pointer to [**[]ManaV2PskConfiguration**](ManaV2PskConfiguration.md) |  | [optional] 
**SakConfigurations** | Pointer to [**[]ManaV2SakConfiguration**](ManaV2SakConfiguration.md) |  | [optional] 
**SplitSakConfigByLagMember** | Pointer to **bool** |  | [optional] 
**TransparentVlan** | Pointer to **bool** |  | [optional] 

## Methods

### NewManaV2InterfaceMaCsec

`func NewManaV2InterfaceMaCsec() *ManaV2InterfaceMaCsec`

NewManaV2InterfaceMaCsec instantiates a new ManaV2InterfaceMaCsec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2InterfaceMaCsecWithDefaults

`func NewManaV2InterfaceMaCsecWithDefaults() *ManaV2InterfaceMaCsec`

NewManaV2InterfaceMaCsecWithDefaults instantiates a new ManaV2InterfaceMaCsec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ManaV2InterfaceMaCsec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ManaV2InterfaceMaCsec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ManaV2InterfaceMaCsec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ManaV2InterfaceMaCsec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEncryptionEnforcementMode

`func (o *ManaV2InterfaceMaCsec) GetEncryptionEnforcementMode() string`

GetEncryptionEnforcementMode returns the EncryptionEnforcementMode field if non-nil, zero value otherwise.

### GetEncryptionEnforcementModeOk

`func (o *ManaV2InterfaceMaCsec) GetEncryptionEnforcementModeOk() (*string, bool)`

GetEncryptionEnforcementModeOk returns a tuple with the EncryptionEnforcementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnforcementMode

`func (o *ManaV2InterfaceMaCsec) SetEncryptionEnforcementMode(v string)`

SetEncryptionEnforcementMode sets EncryptionEnforcementMode field to given value.

### HasEncryptionEnforcementMode

`func (o *ManaV2InterfaceMaCsec) HasEncryptionEnforcementMode() bool`

HasEncryptionEnforcementMode returns a boolean if a field has been set.

### GetKeyServerPriority

`func (o *ManaV2InterfaceMaCsec) GetKeyServerPriority() int64`

GetKeyServerPriority returns the KeyServerPriority field if non-nil, zero value otherwise.

### GetKeyServerPriorityOk

`func (o *ManaV2InterfaceMaCsec) GetKeyServerPriorityOk() (*int64, bool)`

GetKeyServerPriorityOk returns a tuple with the KeyServerPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyServerPriority

`func (o *ManaV2InterfaceMaCsec) SetKeyServerPriority(v int64)`

SetKeyServerPriority sets KeyServerPriority field to given value.

### HasKeyServerPriority

`func (o *ManaV2InterfaceMaCsec) HasKeyServerPriority() bool`

HasKeyServerPriority returns a boolean if a field has been set.

### GetPskConfigurations

`func (o *ManaV2InterfaceMaCsec) GetPskConfigurations() []ManaV2PskConfiguration`

GetPskConfigurations returns the PskConfigurations field if non-nil, zero value otherwise.

### GetPskConfigurationsOk

`func (o *ManaV2InterfaceMaCsec) GetPskConfigurationsOk() (*[]ManaV2PskConfiguration, bool)`

GetPskConfigurationsOk returns a tuple with the PskConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskConfigurations

`func (o *ManaV2InterfaceMaCsec) SetPskConfigurations(v []ManaV2PskConfiguration)`

SetPskConfigurations sets PskConfigurations field to given value.

### HasPskConfigurations

`func (o *ManaV2InterfaceMaCsec) HasPskConfigurations() bool`

HasPskConfigurations returns a boolean if a field has been set.

### GetSakConfigurations

`func (o *ManaV2InterfaceMaCsec) GetSakConfigurations() []ManaV2SakConfiguration`

GetSakConfigurations returns the SakConfigurations field if non-nil, zero value otherwise.

### GetSakConfigurationsOk

`func (o *ManaV2InterfaceMaCsec) GetSakConfigurationsOk() (*[]ManaV2SakConfiguration, bool)`

GetSakConfigurationsOk returns a tuple with the SakConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSakConfigurations

`func (o *ManaV2InterfaceMaCsec) SetSakConfigurations(v []ManaV2SakConfiguration)`

SetSakConfigurations sets SakConfigurations field to given value.

### HasSakConfigurations

`func (o *ManaV2InterfaceMaCsec) HasSakConfigurations() bool`

HasSakConfigurations returns a boolean if a field has been set.

### GetSplitSakConfigByLagMember

`func (o *ManaV2InterfaceMaCsec) GetSplitSakConfigByLagMember() bool`

GetSplitSakConfigByLagMember returns the SplitSakConfigByLagMember field if non-nil, zero value otherwise.

### GetSplitSakConfigByLagMemberOk

`func (o *ManaV2InterfaceMaCsec) GetSplitSakConfigByLagMemberOk() (*bool, bool)`

GetSplitSakConfigByLagMemberOk returns a tuple with the SplitSakConfigByLagMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitSakConfigByLagMember

`func (o *ManaV2InterfaceMaCsec) SetSplitSakConfigByLagMember(v bool)`

SetSplitSakConfigByLagMember sets SplitSakConfigByLagMember field to given value.

### HasSplitSakConfigByLagMember

`func (o *ManaV2InterfaceMaCsec) HasSplitSakConfigByLagMember() bool`

HasSplitSakConfigByLagMember returns a boolean if a field has been set.

### GetTransparentVlan

`func (o *ManaV2InterfaceMaCsec) GetTransparentVlan() bool`

GetTransparentVlan returns the TransparentVlan field if non-nil, zero value otherwise.

### GetTransparentVlanOk

`func (o *ManaV2InterfaceMaCsec) GetTransparentVlanOk() (*bool, bool)`

GetTransparentVlanOk returns a tuple with the TransparentVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentVlan

`func (o *ManaV2InterfaceMaCsec) SetTransparentVlan(v bool)`

SetTransparentVlan sets TransparentVlan field to given value.

### HasTransparentVlan

`func (o *ManaV2InterfaceMaCsec) HasTransparentVlan() bool`

HasTransparentVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


