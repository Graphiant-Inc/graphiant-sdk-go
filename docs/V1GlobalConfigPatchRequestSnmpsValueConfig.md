# V1GlobalConfigPatchRequestSnmpsValueConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Communities** | Pointer to [**map[string]V1PortalApikeysPostRequest**](V1PortalApikeysPostRequest.md) |  | [optional] 
**EngineAuthenTraps** | Pointer to **bool** |  | [optional] 
**EngineEnabled** | Pointer to **bool** |  | [optional] 
**EngineEndpoints** | Pointer to [**map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue**](V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue.md) |  | [optional] 
**EngineHighCpuTraps** | Pointer to **bool** |  | [optional] 
**EngineHighMemoryTraps** | Pointer to **bool** |  | [optional] 
**EngineIdAdminOctets** | Pointer to **string** |  | [optional] 
**EngineIdAdminText** | Pointer to **string** |  | [optional] 
**EngineIdIpv4** | Pointer to **string** |  | [optional] 
**EngineIdIpv6** | Pointer to **string** |  | [optional] 
**EngineIdMac** | Pointer to **string** |  | [optional] 
**EngineIdRaw** | Pointer to **string** |  | [optional] 
**EngineLocalAcessV4** | Pointer to **bool** |  | [optional] 
**EngineLocalAcessV6** | Pointer to **bool** |  | [optional] 
**EngineUserHints** | Pointer to **bool** |  | [optional] 
**EngineUserValidation** | Pointer to **bool** |  | [optional] 
**GlobalId** | Pointer to **int64** |  | [optional] 
**IsGlobalSync** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotifyFilterProfiles** | Pointer to [**map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue**](V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue.md) |  | [optional] 
**Targets** | Pointer to [**map[string]V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue**](V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue.md) |  | [optional] 
**UsmLocalUsers** | Pointer to [**map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue**](V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue.md) |  | [optional] 
**UsmRemoteUsers** | Pointer to [**map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue**](V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue.md) |  | [optional] 
**V2cEnabled** | Pointer to **bool** |  | [optional] 
**V3Enabled** | Pointer to **bool** |  | [optional] 
**VacmGroups** | Pointer to [**map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue**](V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue.md) |  | [optional] 
**VacmViews** | Pointer to [**map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue**](V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue.md) |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestSnmpsValueConfig

`func NewV1GlobalConfigPatchRequestSnmpsValueConfig() *V1GlobalConfigPatchRequestSnmpsValueConfig`

NewV1GlobalConfigPatchRequestSnmpsValueConfig instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestSnmpsValueConfigWithDefaults

`func NewV1GlobalConfigPatchRequestSnmpsValueConfigWithDefaults() *V1GlobalConfigPatchRequestSnmpsValueConfig`

NewV1GlobalConfigPatchRequestSnmpsValueConfigWithDefaults instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunities

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetCommunities() map[string]V1PortalApikeysPostRequest`

GetCommunities returns the Communities field if non-nil, zero value otherwise.

### GetCommunitiesOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetCommunitiesOk() (*map[string]V1PortalApikeysPostRequest, bool)`

GetCommunitiesOk returns a tuple with the Communities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunities

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetCommunities(v map[string]V1PortalApikeysPostRequest)`

SetCommunities sets Communities field to given value.

### HasCommunities

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasCommunities() bool`

HasCommunities returns a boolean if a field has been set.

### GetEngineAuthenTraps

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineAuthenTraps() bool`

GetEngineAuthenTraps returns the EngineAuthenTraps field if non-nil, zero value otherwise.

### GetEngineAuthenTrapsOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineAuthenTrapsOk() (*bool, bool)`

GetEngineAuthenTrapsOk returns a tuple with the EngineAuthenTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineAuthenTraps

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineAuthenTraps(v bool)`

SetEngineAuthenTraps sets EngineAuthenTraps field to given value.

### HasEngineAuthenTraps

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineAuthenTraps() bool`

HasEngineAuthenTraps returns a boolean if a field has been set.

### GetEngineEnabled

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineEnabled() bool`

GetEngineEnabled returns the EngineEnabled field if non-nil, zero value otherwise.

### GetEngineEnabledOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineEnabledOk() (*bool, bool)`

GetEngineEnabledOk returns a tuple with the EngineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEnabled

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineEnabled(v bool)`

SetEngineEnabled sets EngineEnabled field to given value.

### HasEngineEnabled

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineEnabled() bool`

HasEngineEnabled returns a boolean if a field has been set.

### GetEngineEndpoints

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineEndpoints() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue`

GetEngineEndpoints returns the EngineEndpoints field if non-nil, zero value otherwise.

### GetEngineEndpointsOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineEndpointsOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue, bool)`

GetEngineEndpointsOk returns a tuple with the EngineEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEndpoints

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineEndpoints(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue)`

SetEngineEndpoints sets EngineEndpoints field to given value.

### HasEngineEndpoints

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineEndpoints() bool`

HasEngineEndpoints returns a boolean if a field has been set.

### GetEngineHighCpuTraps

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineHighCpuTraps() bool`

GetEngineHighCpuTraps returns the EngineHighCpuTraps field if non-nil, zero value otherwise.

### GetEngineHighCpuTrapsOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineHighCpuTrapsOk() (*bool, bool)`

GetEngineHighCpuTrapsOk returns a tuple with the EngineHighCpuTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHighCpuTraps

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineHighCpuTraps(v bool)`

SetEngineHighCpuTraps sets EngineHighCpuTraps field to given value.

### HasEngineHighCpuTraps

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineHighCpuTraps() bool`

HasEngineHighCpuTraps returns a boolean if a field has been set.

### GetEngineHighMemoryTraps

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineHighMemoryTraps() bool`

GetEngineHighMemoryTraps returns the EngineHighMemoryTraps field if non-nil, zero value otherwise.

### GetEngineHighMemoryTrapsOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineHighMemoryTrapsOk() (*bool, bool)`

GetEngineHighMemoryTrapsOk returns a tuple with the EngineHighMemoryTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHighMemoryTraps

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineHighMemoryTraps(v bool)`

SetEngineHighMemoryTraps sets EngineHighMemoryTraps field to given value.

### HasEngineHighMemoryTraps

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineHighMemoryTraps() bool`

HasEngineHighMemoryTraps returns a boolean if a field has been set.

### GetEngineIdAdminOctets

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdAdminOctets() string`

GetEngineIdAdminOctets returns the EngineIdAdminOctets field if non-nil, zero value otherwise.

### GetEngineIdAdminOctetsOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdAdminOctetsOk() (*string, bool)`

GetEngineIdAdminOctetsOk returns a tuple with the EngineIdAdminOctets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdAdminOctets

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdAdminOctets(v string)`

SetEngineIdAdminOctets sets EngineIdAdminOctets field to given value.

### HasEngineIdAdminOctets

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdAdminOctets() bool`

HasEngineIdAdminOctets returns a boolean if a field has been set.

### GetEngineIdAdminText

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdAdminText() string`

GetEngineIdAdminText returns the EngineIdAdminText field if non-nil, zero value otherwise.

### GetEngineIdAdminTextOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdAdminTextOk() (*string, bool)`

GetEngineIdAdminTextOk returns a tuple with the EngineIdAdminText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdAdminText

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdAdminText(v string)`

SetEngineIdAdminText sets EngineIdAdminText field to given value.

### HasEngineIdAdminText

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdAdminText() bool`

HasEngineIdAdminText returns a boolean if a field has been set.

### GetEngineIdIpv4

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdIpv4() string`

GetEngineIdIpv4 returns the EngineIdIpv4 field if non-nil, zero value otherwise.

### GetEngineIdIpv4Ok

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdIpv4Ok() (*string, bool)`

GetEngineIdIpv4Ok returns a tuple with the EngineIdIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdIpv4

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdIpv4(v string)`

SetEngineIdIpv4 sets EngineIdIpv4 field to given value.

### HasEngineIdIpv4

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdIpv4() bool`

HasEngineIdIpv4 returns a boolean if a field has been set.

### GetEngineIdIpv6

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdIpv6() string`

GetEngineIdIpv6 returns the EngineIdIpv6 field if non-nil, zero value otherwise.

### GetEngineIdIpv6Ok

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdIpv6Ok() (*string, bool)`

GetEngineIdIpv6Ok returns a tuple with the EngineIdIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdIpv6

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdIpv6(v string)`

SetEngineIdIpv6 sets EngineIdIpv6 field to given value.

### HasEngineIdIpv6

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdIpv6() bool`

HasEngineIdIpv6 returns a boolean if a field has been set.

### GetEngineIdMac

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdMac() string`

GetEngineIdMac returns the EngineIdMac field if non-nil, zero value otherwise.

### GetEngineIdMacOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdMacOk() (*string, bool)`

GetEngineIdMacOk returns a tuple with the EngineIdMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdMac

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdMac(v string)`

SetEngineIdMac sets EngineIdMac field to given value.

### HasEngineIdMac

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdMac() bool`

HasEngineIdMac returns a boolean if a field has been set.

### GetEngineIdRaw

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdRaw() string`

GetEngineIdRaw returns the EngineIdRaw field if non-nil, zero value otherwise.

### GetEngineIdRawOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdRawOk() (*string, bool)`

GetEngineIdRawOk returns a tuple with the EngineIdRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdRaw

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdRaw(v string)`

SetEngineIdRaw sets EngineIdRaw field to given value.

### HasEngineIdRaw

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdRaw() bool`

HasEngineIdRaw returns a boolean if a field has been set.

### GetEngineLocalAcessV4

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineLocalAcessV4() bool`

GetEngineLocalAcessV4 returns the EngineLocalAcessV4 field if non-nil, zero value otherwise.

### GetEngineLocalAcessV4Ok

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineLocalAcessV4Ok() (*bool, bool)`

GetEngineLocalAcessV4Ok returns a tuple with the EngineLocalAcessV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineLocalAcessV4

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineLocalAcessV4(v bool)`

SetEngineLocalAcessV4 sets EngineLocalAcessV4 field to given value.

### HasEngineLocalAcessV4

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineLocalAcessV4() bool`

HasEngineLocalAcessV4 returns a boolean if a field has been set.

### GetEngineLocalAcessV6

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineLocalAcessV6() bool`

GetEngineLocalAcessV6 returns the EngineLocalAcessV6 field if non-nil, zero value otherwise.

### GetEngineLocalAcessV6Ok

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineLocalAcessV6Ok() (*bool, bool)`

GetEngineLocalAcessV6Ok returns a tuple with the EngineLocalAcessV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineLocalAcessV6

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineLocalAcessV6(v bool)`

SetEngineLocalAcessV6 sets EngineLocalAcessV6 field to given value.

### HasEngineLocalAcessV6

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineLocalAcessV6() bool`

HasEngineLocalAcessV6 returns a boolean if a field has been set.

### GetEngineUserHints

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineUserHints() bool`

GetEngineUserHints returns the EngineUserHints field if non-nil, zero value otherwise.

### GetEngineUserHintsOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineUserHintsOk() (*bool, bool)`

GetEngineUserHintsOk returns a tuple with the EngineUserHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineUserHints

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineUserHints(v bool)`

SetEngineUserHints sets EngineUserHints field to given value.

### HasEngineUserHints

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineUserHints() bool`

HasEngineUserHints returns a boolean if a field has been set.

### GetEngineUserValidation

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineUserValidation() bool`

GetEngineUserValidation returns the EngineUserValidation field if non-nil, zero value otherwise.

### GetEngineUserValidationOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineUserValidationOk() (*bool, bool)`

GetEngineUserValidationOk returns a tuple with the EngineUserValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineUserValidation

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineUserValidation(v bool)`

SetEngineUserValidation sets EngineUserValidation field to given value.

### HasEngineUserValidation

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineUserValidation() bool`

HasEngineUserValidation returns a boolean if a field has been set.

### GetGlobalId

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetIsGlobalSync

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetIsGlobalSync() bool`

GetIsGlobalSync returns the IsGlobalSync field if non-nil, zero value otherwise.

### GetIsGlobalSyncOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetIsGlobalSyncOk() (*bool, bool)`

GetIsGlobalSyncOk returns a tuple with the IsGlobalSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalSync

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetIsGlobalSync(v bool)`

SetIsGlobalSync sets IsGlobalSync field to given value.

### HasIsGlobalSync

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasIsGlobalSync() bool`

HasIsGlobalSync returns a boolean if a field has been set.

### GetName

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyFilterProfiles

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetNotifyFilterProfiles() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue`

GetNotifyFilterProfiles returns the NotifyFilterProfiles field if non-nil, zero value otherwise.

### GetNotifyFilterProfilesOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetNotifyFilterProfilesOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue, bool)`

GetNotifyFilterProfilesOk returns a tuple with the NotifyFilterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFilterProfiles

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetNotifyFilterProfiles(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue)`

SetNotifyFilterProfiles sets NotifyFilterProfiles field to given value.

### HasNotifyFilterProfiles

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasNotifyFilterProfiles() bool`

HasNotifyFilterProfiles returns a boolean if a field has been set.

### GetTargets

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetTargets() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetTargetsOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetTargets(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetUsmLocalUsers

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetUsmLocalUsers() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue`

GetUsmLocalUsers returns the UsmLocalUsers field if non-nil, zero value otherwise.

### GetUsmLocalUsersOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetUsmLocalUsersOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue, bool)`

GetUsmLocalUsersOk returns a tuple with the UsmLocalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsmLocalUsers

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetUsmLocalUsers(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue)`

SetUsmLocalUsers sets UsmLocalUsers field to given value.

### HasUsmLocalUsers

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasUsmLocalUsers() bool`

HasUsmLocalUsers returns a boolean if a field has been set.

### GetUsmRemoteUsers

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetUsmRemoteUsers() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue`

GetUsmRemoteUsers returns the UsmRemoteUsers field if non-nil, zero value otherwise.

### GetUsmRemoteUsersOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetUsmRemoteUsersOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue, bool)`

GetUsmRemoteUsersOk returns a tuple with the UsmRemoteUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsmRemoteUsers

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetUsmRemoteUsers(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue)`

SetUsmRemoteUsers sets UsmRemoteUsers field to given value.

### HasUsmRemoteUsers

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasUsmRemoteUsers() bool`

HasUsmRemoteUsers returns a boolean if a field has been set.

### GetV2cEnabled

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetV2cEnabled() bool`

GetV2cEnabled returns the V2cEnabled field if non-nil, zero value otherwise.

### GetV2cEnabledOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetV2cEnabledOk() (*bool, bool)`

GetV2cEnabledOk returns a tuple with the V2cEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2cEnabled

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetV2cEnabled(v bool)`

SetV2cEnabled sets V2cEnabled field to given value.

### HasV2cEnabled

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasV2cEnabled() bool`

HasV2cEnabled returns a boolean if a field has been set.

### GetV3Enabled

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetV3Enabled() bool`

GetV3Enabled returns the V3Enabled field if non-nil, zero value otherwise.

### GetV3EnabledOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetV3EnabledOk() (*bool, bool)`

GetV3EnabledOk returns a tuple with the V3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Enabled

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetV3Enabled(v bool)`

SetV3Enabled sets V3Enabled field to given value.

### HasV3Enabled

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasV3Enabled() bool`

HasV3Enabled returns a boolean if a field has been set.

### GetVacmGroups

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetVacmGroups() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue`

GetVacmGroups returns the VacmGroups field if non-nil, zero value otherwise.

### GetVacmGroupsOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetVacmGroupsOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue, bool)`

GetVacmGroupsOk returns a tuple with the VacmGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacmGroups

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetVacmGroups(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue)`

SetVacmGroups sets VacmGroups field to given value.

### HasVacmGroups

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasVacmGroups() bool`

HasVacmGroups returns a boolean if a field has been set.

### GetVacmViews

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetVacmViews() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue`

GetVacmViews returns the VacmViews field if non-nil, zero value otherwise.

### GetVacmViewsOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetVacmViewsOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue, bool)`

GetVacmViewsOk returns a tuple with the VacmViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacmViews

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetVacmViews(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue)`

SetVacmViews sets VacmViews field to given value.

### HasVacmViews

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasVacmViews() bool`

HasVacmViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


