# ManaV2MaCsecConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether MACsec is enabled or disabled (required) | 
**EncryptionEnforcementMode** | **string** | The encryption enforcement mode (required) | 
**GlobalSakConfiguration** | [**ManaV2NullableSakConfiguration**](ManaV2NullableSakConfiguration.md) |  | 
**KeyServerPriority** | Pointer to **int64** | The priority of the key server. Lower number means higher priority. | [optional] 
**PskConfigurations** | Pointer to [**[]ManaV2PskConfiguration**](ManaV2PskConfiguration.md) |  | [optional] 
**PskConfigurationsByNickname** | [**map[string]ManaV2NullablePskConfiguration**](ManaV2NullablePskConfiguration.md) |  | 
**SakConfigurations** | Pointer to [**[]ManaV2SakConfiguration**](ManaV2SakConfiguration.md) |  | [optional] 
**SakConfigurationsByLagMemberInterfaceId** | [**map[string]ManaV2NullableSakConfiguration**](ManaV2NullableSakConfiguration.md) |  | 
**SplitSakConfigByLagMember** | Pointer to **bool** | Whether to allow individual SAK configurations for each lag member | [optional] 
**TransparentVlan** | Pointer to **bool** | Whether transparent VLAN is enabled or disabled | [optional] 

## Methods

### NewManaV2MaCsecConfiguration

`func NewManaV2MaCsecConfiguration(enabled bool, encryptionEnforcementMode string, globalSakConfiguration ManaV2NullableSakConfiguration, pskConfigurationsByNickname map[string]ManaV2NullablePskConfiguration, sakConfigurationsByLagMemberInterfaceId map[string]ManaV2NullableSakConfiguration, ) *ManaV2MaCsecConfiguration`

NewManaV2MaCsecConfiguration instantiates a new ManaV2MaCsecConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2MaCsecConfigurationWithDefaults

`func NewManaV2MaCsecConfigurationWithDefaults() *ManaV2MaCsecConfiguration`

NewManaV2MaCsecConfigurationWithDefaults instantiates a new ManaV2MaCsecConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ManaV2MaCsecConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ManaV2MaCsecConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ManaV2MaCsecConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEncryptionEnforcementMode

`func (o *ManaV2MaCsecConfiguration) GetEncryptionEnforcementMode() string`

GetEncryptionEnforcementMode returns the EncryptionEnforcementMode field if non-nil, zero value otherwise.

### GetEncryptionEnforcementModeOk

`func (o *ManaV2MaCsecConfiguration) GetEncryptionEnforcementModeOk() (*string, bool)`

GetEncryptionEnforcementModeOk returns a tuple with the EncryptionEnforcementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnforcementMode

`func (o *ManaV2MaCsecConfiguration) SetEncryptionEnforcementMode(v string)`

SetEncryptionEnforcementMode sets EncryptionEnforcementMode field to given value.


### GetGlobalSakConfiguration

`func (o *ManaV2MaCsecConfiguration) GetGlobalSakConfiguration() ManaV2NullableSakConfiguration`

GetGlobalSakConfiguration returns the GlobalSakConfiguration field if non-nil, zero value otherwise.

### GetGlobalSakConfigurationOk

`func (o *ManaV2MaCsecConfiguration) GetGlobalSakConfigurationOk() (*ManaV2NullableSakConfiguration, bool)`

GetGlobalSakConfigurationOk returns a tuple with the GlobalSakConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSakConfiguration

`func (o *ManaV2MaCsecConfiguration) SetGlobalSakConfiguration(v ManaV2NullableSakConfiguration)`

SetGlobalSakConfiguration sets GlobalSakConfiguration field to given value.


### GetKeyServerPriority

`func (o *ManaV2MaCsecConfiguration) GetKeyServerPriority() int64`

GetKeyServerPriority returns the KeyServerPriority field if non-nil, zero value otherwise.

### GetKeyServerPriorityOk

`func (o *ManaV2MaCsecConfiguration) GetKeyServerPriorityOk() (*int64, bool)`

GetKeyServerPriorityOk returns a tuple with the KeyServerPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyServerPriority

`func (o *ManaV2MaCsecConfiguration) SetKeyServerPriority(v int64)`

SetKeyServerPriority sets KeyServerPriority field to given value.

### HasKeyServerPriority

`func (o *ManaV2MaCsecConfiguration) HasKeyServerPriority() bool`

HasKeyServerPriority returns a boolean if a field has been set.

### GetPskConfigurations

`func (o *ManaV2MaCsecConfiguration) GetPskConfigurations() []ManaV2PskConfiguration`

GetPskConfigurations returns the PskConfigurations field if non-nil, zero value otherwise.

### GetPskConfigurationsOk

`func (o *ManaV2MaCsecConfiguration) GetPskConfigurationsOk() (*[]ManaV2PskConfiguration, bool)`

GetPskConfigurationsOk returns a tuple with the PskConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskConfigurations

`func (o *ManaV2MaCsecConfiguration) SetPskConfigurations(v []ManaV2PskConfiguration)`

SetPskConfigurations sets PskConfigurations field to given value.

### HasPskConfigurations

`func (o *ManaV2MaCsecConfiguration) HasPskConfigurations() bool`

HasPskConfigurations returns a boolean if a field has been set.

### GetPskConfigurationsByNickname

`func (o *ManaV2MaCsecConfiguration) GetPskConfigurationsByNickname() map[string]ManaV2NullablePskConfiguration`

GetPskConfigurationsByNickname returns the PskConfigurationsByNickname field if non-nil, zero value otherwise.

### GetPskConfigurationsByNicknameOk

`func (o *ManaV2MaCsecConfiguration) GetPskConfigurationsByNicknameOk() (*map[string]ManaV2NullablePskConfiguration, bool)`

GetPskConfigurationsByNicknameOk returns a tuple with the PskConfigurationsByNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskConfigurationsByNickname

`func (o *ManaV2MaCsecConfiguration) SetPskConfigurationsByNickname(v map[string]ManaV2NullablePskConfiguration)`

SetPskConfigurationsByNickname sets PskConfigurationsByNickname field to given value.


### GetSakConfigurations

`func (o *ManaV2MaCsecConfiguration) GetSakConfigurations() []ManaV2SakConfiguration`

GetSakConfigurations returns the SakConfigurations field if non-nil, zero value otherwise.

### GetSakConfigurationsOk

`func (o *ManaV2MaCsecConfiguration) GetSakConfigurationsOk() (*[]ManaV2SakConfiguration, bool)`

GetSakConfigurationsOk returns a tuple with the SakConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSakConfigurations

`func (o *ManaV2MaCsecConfiguration) SetSakConfigurations(v []ManaV2SakConfiguration)`

SetSakConfigurations sets SakConfigurations field to given value.

### HasSakConfigurations

`func (o *ManaV2MaCsecConfiguration) HasSakConfigurations() bool`

HasSakConfigurations returns a boolean if a field has been set.

### GetSakConfigurationsByLagMemberInterfaceId

`func (o *ManaV2MaCsecConfiguration) GetSakConfigurationsByLagMemberInterfaceId() map[string]ManaV2NullableSakConfiguration`

GetSakConfigurationsByLagMemberInterfaceId returns the SakConfigurationsByLagMemberInterfaceId field if non-nil, zero value otherwise.

### GetSakConfigurationsByLagMemberInterfaceIdOk

`func (o *ManaV2MaCsecConfiguration) GetSakConfigurationsByLagMemberInterfaceIdOk() (*map[string]ManaV2NullableSakConfiguration, bool)`

GetSakConfigurationsByLagMemberInterfaceIdOk returns a tuple with the SakConfigurationsByLagMemberInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSakConfigurationsByLagMemberInterfaceId

`func (o *ManaV2MaCsecConfiguration) SetSakConfigurationsByLagMemberInterfaceId(v map[string]ManaV2NullableSakConfiguration)`

SetSakConfigurationsByLagMemberInterfaceId sets SakConfigurationsByLagMemberInterfaceId field to given value.


### GetSplitSakConfigByLagMember

`func (o *ManaV2MaCsecConfiguration) GetSplitSakConfigByLagMember() bool`

GetSplitSakConfigByLagMember returns the SplitSakConfigByLagMember field if non-nil, zero value otherwise.

### GetSplitSakConfigByLagMemberOk

`func (o *ManaV2MaCsecConfiguration) GetSplitSakConfigByLagMemberOk() (*bool, bool)`

GetSplitSakConfigByLagMemberOk returns a tuple with the SplitSakConfigByLagMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitSakConfigByLagMember

`func (o *ManaV2MaCsecConfiguration) SetSplitSakConfigByLagMember(v bool)`

SetSplitSakConfigByLagMember sets SplitSakConfigByLagMember field to given value.

### HasSplitSakConfigByLagMember

`func (o *ManaV2MaCsecConfiguration) HasSplitSakConfigByLagMember() bool`

HasSplitSakConfigByLagMember returns a boolean if a field has been set.

### GetTransparentVlan

`func (o *ManaV2MaCsecConfiguration) GetTransparentVlan() bool`

GetTransparentVlan returns the TransparentVlan field if non-nil, zero value otherwise.

### GetTransparentVlanOk

`func (o *ManaV2MaCsecConfiguration) GetTransparentVlanOk() (*bool, bool)`

GetTransparentVlanOk returns a tuple with the TransparentVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentVlan

`func (o *ManaV2MaCsecConfiguration) SetTransparentVlan(v bool)`

SetTransparentVlan sets TransparentVlan field to given value.

### HasTransparentVlan

`func (o *ManaV2MaCsecConfiguration) HasTransparentVlan() bool`

HasTransparentVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


