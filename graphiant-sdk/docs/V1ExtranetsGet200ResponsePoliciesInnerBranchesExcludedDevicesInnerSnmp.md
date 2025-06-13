# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Communities** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpCommunitiesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpCommunitiesInner.md) |  | [optional] 
**EngineEnableAuthenTraps** | Pointer to **bool** |  | [optional] 
**EngineEnableHighMemoryTraps** | Pointer to **bool** |  | [optional] 
**EngineEnableHighCpuTraps** | Pointer to **bool** |  | [optional] 
**EngineEnableLocalAcessV4** | Pointer to **bool** |  | [optional] 
**EngineEnableLocalAcessV6** | Pointer to **bool** |  | [optional] 
**EngineEnableUserHints** | Pointer to **bool** |  | [optional] 
**EngineEnableUserValidation** | Pointer to **bool** |  | [optional] 
**EngineEnabled** | Pointer to **bool** |  | [optional] 
**EngineEndpoints** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpEngineEndpointsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpEngineEndpointsInner.md) |  | [optional] 
**EngineIdAdminOctets** | Pointer to **string** |  | [optional] 
**EngineIdAdminText** | Pointer to **string** |  | [optional] 
**EngineIdIpv4** | Pointer to **string** |  | [optional] 
**EngineIdIpv6** | Pointer to **string** |  | [optional] 
**EngineIdMac** | Pointer to **string** |  | [optional] 
**EngineIdRaw** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**GlobalId** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotifyFilterProfiles** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpNotifyFilterProfilesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpNotifyFilterProfilesInner.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpTargetsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpTargetsInner.md) |  | [optional] 
**UsmLocalUsers** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpUsmLocalUsersInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpUsmLocalUsersInner.md) |  | [optional] 
**UsmRemoteUsers** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpUsmRemoteUsersInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpUsmRemoteUsersInner.md) |  | [optional] 
**V2cEnabled** | Pointer to **bool** |  | [optional] 
**V3Enabled** | Pointer to **bool** |  | [optional] 
**VacmGroups** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner.md) |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunities

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetCommunities() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpCommunitiesInner`

GetCommunities returns the Communities field if non-nil, zero value otherwise.

### GetCommunitiesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetCommunitiesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpCommunitiesInner, bool)`

GetCommunitiesOk returns a tuple with the Communities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunities

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetCommunities(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpCommunitiesInner)`

SetCommunities sets Communities field to given value.

### HasCommunities

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasCommunities() bool`

HasCommunities returns a boolean if a field has been set.

### GetEngineEnableAuthenTraps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableAuthenTraps() bool`

GetEngineEnableAuthenTraps returns the EngineEnableAuthenTraps field if non-nil, zero value otherwise.

### GetEngineEnableAuthenTrapsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableAuthenTrapsOk() (*bool, bool)`

GetEngineEnableAuthenTrapsOk returns a tuple with the EngineEnableAuthenTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEnableAuthenTraps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineEnableAuthenTraps(v bool)`

SetEngineEnableAuthenTraps sets EngineEnableAuthenTraps field to given value.

### HasEngineEnableAuthenTraps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineEnableAuthenTraps() bool`

HasEngineEnableAuthenTraps returns a boolean if a field has been set.

### GetEngineEnableHighMemoryTraps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableHighMemoryTraps() bool`

GetEngineEnableHighMemoryTraps returns the EngineEnableHighMemoryTraps field if non-nil, zero value otherwise.

### GetEngineEnableHighMemoryTrapsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableHighMemoryTrapsOk() (*bool, bool)`

GetEngineEnableHighMemoryTrapsOk returns a tuple with the EngineEnableHighMemoryTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEnableHighMemoryTraps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineEnableHighMemoryTraps(v bool)`

SetEngineEnableHighMemoryTraps sets EngineEnableHighMemoryTraps field to given value.

### HasEngineEnableHighMemoryTraps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineEnableHighMemoryTraps() bool`

HasEngineEnableHighMemoryTraps returns a boolean if a field has been set.

### GetEngineEnableHighCpuTraps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableHighCpuTraps() bool`

GetEngineEnableHighCpuTraps returns the EngineEnableHighCpuTraps field if non-nil, zero value otherwise.

### GetEngineEnableHighCpuTrapsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableHighCpuTrapsOk() (*bool, bool)`

GetEngineEnableHighCpuTrapsOk returns a tuple with the EngineEnableHighCpuTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEnableHighCpuTraps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineEnableHighCpuTraps(v bool)`

SetEngineEnableHighCpuTraps sets EngineEnableHighCpuTraps field to given value.

### HasEngineEnableHighCpuTraps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineEnableHighCpuTraps() bool`

HasEngineEnableHighCpuTraps returns a boolean if a field has been set.

### GetEngineEnableLocalAcessV4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableLocalAcessV4() bool`

GetEngineEnableLocalAcessV4 returns the EngineEnableLocalAcessV4 field if non-nil, zero value otherwise.

### GetEngineEnableLocalAcessV4Ok

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableLocalAcessV4Ok() (*bool, bool)`

GetEngineEnableLocalAcessV4Ok returns a tuple with the EngineEnableLocalAcessV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEnableLocalAcessV4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineEnableLocalAcessV4(v bool)`

SetEngineEnableLocalAcessV4 sets EngineEnableLocalAcessV4 field to given value.

### HasEngineEnableLocalAcessV4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineEnableLocalAcessV4() bool`

HasEngineEnableLocalAcessV4 returns a boolean if a field has been set.

### GetEngineEnableLocalAcessV6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableLocalAcessV6() bool`

GetEngineEnableLocalAcessV6 returns the EngineEnableLocalAcessV6 field if non-nil, zero value otherwise.

### GetEngineEnableLocalAcessV6Ok

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableLocalAcessV6Ok() (*bool, bool)`

GetEngineEnableLocalAcessV6Ok returns a tuple with the EngineEnableLocalAcessV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEnableLocalAcessV6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineEnableLocalAcessV6(v bool)`

SetEngineEnableLocalAcessV6 sets EngineEnableLocalAcessV6 field to given value.

### HasEngineEnableLocalAcessV6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineEnableLocalAcessV6() bool`

HasEngineEnableLocalAcessV6 returns a boolean if a field has been set.

### GetEngineEnableUserHints

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableUserHints() bool`

GetEngineEnableUserHints returns the EngineEnableUserHints field if non-nil, zero value otherwise.

### GetEngineEnableUserHintsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableUserHintsOk() (*bool, bool)`

GetEngineEnableUserHintsOk returns a tuple with the EngineEnableUserHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEnableUserHints

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineEnableUserHints(v bool)`

SetEngineEnableUserHints sets EngineEnableUserHints field to given value.

### HasEngineEnableUserHints

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineEnableUserHints() bool`

HasEngineEnableUserHints returns a boolean if a field has been set.

### GetEngineEnableUserValidation

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableUserValidation() bool`

GetEngineEnableUserValidation returns the EngineEnableUserValidation field if non-nil, zero value otherwise.

### GetEngineEnableUserValidationOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnableUserValidationOk() (*bool, bool)`

GetEngineEnableUserValidationOk returns a tuple with the EngineEnableUserValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEnableUserValidation

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineEnableUserValidation(v bool)`

SetEngineEnableUserValidation sets EngineEnableUserValidation field to given value.

### HasEngineEnableUserValidation

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineEnableUserValidation() bool`

HasEngineEnableUserValidation returns a boolean if a field has been set.

### GetEngineEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnabled() bool`

GetEngineEnabled returns the EngineEnabled field if non-nil, zero value otherwise.

### GetEngineEnabledOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEnabledOk() (*bool, bool)`

GetEngineEnabledOk returns a tuple with the EngineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineEnabled(v bool)`

SetEngineEnabled sets EngineEnabled field to given value.

### HasEngineEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineEnabled() bool`

HasEngineEnabled returns a boolean if a field has been set.

### GetEngineEndpoints

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEndpoints() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpEngineEndpointsInner`

GetEngineEndpoints returns the EngineEndpoints field if non-nil, zero value otherwise.

### GetEngineEndpointsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineEndpointsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpEngineEndpointsInner, bool)`

GetEngineEndpointsOk returns a tuple with the EngineEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineEndpoints

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineEndpoints(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpEngineEndpointsInner)`

SetEngineEndpoints sets EngineEndpoints field to given value.

### HasEngineEndpoints

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineEndpoints() bool`

HasEngineEndpoints returns a boolean if a field has been set.

### GetEngineIdAdminOctets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdAdminOctets() string`

GetEngineIdAdminOctets returns the EngineIdAdminOctets field if non-nil, zero value otherwise.

### GetEngineIdAdminOctetsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdAdminOctetsOk() (*string, bool)`

GetEngineIdAdminOctetsOk returns a tuple with the EngineIdAdminOctets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdAdminOctets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineIdAdminOctets(v string)`

SetEngineIdAdminOctets sets EngineIdAdminOctets field to given value.

### HasEngineIdAdminOctets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineIdAdminOctets() bool`

HasEngineIdAdminOctets returns a boolean if a field has been set.

### GetEngineIdAdminText

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdAdminText() string`

GetEngineIdAdminText returns the EngineIdAdminText field if non-nil, zero value otherwise.

### GetEngineIdAdminTextOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdAdminTextOk() (*string, bool)`

GetEngineIdAdminTextOk returns a tuple with the EngineIdAdminText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdAdminText

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineIdAdminText(v string)`

SetEngineIdAdminText sets EngineIdAdminText field to given value.

### HasEngineIdAdminText

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineIdAdminText() bool`

HasEngineIdAdminText returns a boolean if a field has been set.

### GetEngineIdIpv4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdIpv4() string`

GetEngineIdIpv4 returns the EngineIdIpv4 field if non-nil, zero value otherwise.

### GetEngineIdIpv4Ok

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdIpv4Ok() (*string, bool)`

GetEngineIdIpv4Ok returns a tuple with the EngineIdIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdIpv4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineIdIpv4(v string)`

SetEngineIdIpv4 sets EngineIdIpv4 field to given value.

### HasEngineIdIpv4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineIdIpv4() bool`

HasEngineIdIpv4 returns a boolean if a field has been set.

### GetEngineIdIpv6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdIpv6() string`

GetEngineIdIpv6 returns the EngineIdIpv6 field if non-nil, zero value otherwise.

### GetEngineIdIpv6Ok

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdIpv6Ok() (*string, bool)`

GetEngineIdIpv6Ok returns a tuple with the EngineIdIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdIpv6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineIdIpv6(v string)`

SetEngineIdIpv6 sets EngineIdIpv6 field to given value.

### HasEngineIdIpv6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineIdIpv6() bool`

HasEngineIdIpv6 returns a boolean if a field has been set.

### GetEngineIdMac

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdMac() string`

GetEngineIdMac returns the EngineIdMac field if non-nil, zero value otherwise.

### GetEngineIdMacOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdMacOk() (*string, bool)`

GetEngineIdMacOk returns a tuple with the EngineIdMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdMac

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineIdMac(v string)`

SetEngineIdMac sets EngineIdMac field to given value.

### HasEngineIdMac

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineIdMac() bool`

HasEngineIdMac returns a boolean if a field has been set.

### GetEngineIdRaw

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdRaw() string`

GetEngineIdRaw returns the EngineIdRaw field if non-nil, zero value otherwise.

### GetEngineIdRawOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetEngineIdRawOk() (*string, bool)`

GetEngineIdRawOk returns a tuple with the EngineIdRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdRaw

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetEngineIdRaw(v string)`

SetEngineIdRaw sets EngineIdRaw field to given value.

### HasEngineIdRaw

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasEngineIdRaw() bool`

HasEngineIdRaw returns a boolean if a field has been set.

### GetErrorMessage

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetGlobalId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyFilterProfiles

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetNotifyFilterProfiles() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpNotifyFilterProfilesInner`

GetNotifyFilterProfiles returns the NotifyFilterProfiles field if non-nil, zero value otherwise.

### GetNotifyFilterProfilesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetNotifyFilterProfilesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpNotifyFilterProfilesInner, bool)`

GetNotifyFilterProfilesOk returns a tuple with the NotifyFilterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFilterProfiles

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetNotifyFilterProfiles(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpNotifyFilterProfilesInner)`

SetNotifyFilterProfiles sets NotifyFilterProfiles field to given value.

### HasNotifyFilterProfiles

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasNotifyFilterProfiles() bool`

HasNotifyFilterProfiles returns a boolean if a field has been set.

### GetStatus

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetTargets() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetTargetsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetTargets(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetUsmLocalUsers

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetUsmLocalUsers() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpUsmLocalUsersInner`

GetUsmLocalUsers returns the UsmLocalUsers field if non-nil, zero value otherwise.

### GetUsmLocalUsersOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetUsmLocalUsersOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpUsmLocalUsersInner, bool)`

GetUsmLocalUsersOk returns a tuple with the UsmLocalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsmLocalUsers

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetUsmLocalUsers(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpUsmLocalUsersInner)`

SetUsmLocalUsers sets UsmLocalUsers field to given value.

### HasUsmLocalUsers

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasUsmLocalUsers() bool`

HasUsmLocalUsers returns a boolean if a field has been set.

### GetUsmRemoteUsers

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetUsmRemoteUsers() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpUsmRemoteUsersInner`

GetUsmRemoteUsers returns the UsmRemoteUsers field if non-nil, zero value otherwise.

### GetUsmRemoteUsersOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetUsmRemoteUsersOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpUsmRemoteUsersInner, bool)`

GetUsmRemoteUsersOk returns a tuple with the UsmRemoteUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsmRemoteUsers

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetUsmRemoteUsers(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpUsmRemoteUsersInner)`

SetUsmRemoteUsers sets UsmRemoteUsers field to given value.

### HasUsmRemoteUsers

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasUsmRemoteUsers() bool`

HasUsmRemoteUsers returns a boolean if a field has been set.

### GetV2cEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetV2cEnabled() bool`

GetV2cEnabled returns the V2cEnabled field if non-nil, zero value otherwise.

### GetV2cEnabledOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetV2cEnabledOk() (*bool, bool)`

GetV2cEnabledOk returns a tuple with the V2cEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2cEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetV2cEnabled(v bool)`

SetV2cEnabled sets V2cEnabled field to given value.

### HasV2cEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasV2cEnabled() bool`

HasV2cEnabled returns a boolean if a field has been set.

### GetV3Enabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetV3Enabled() bool`

GetV3Enabled returns the V3Enabled field if non-nil, zero value otherwise.

### GetV3EnabledOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetV3EnabledOk() (*bool, bool)`

GetV3EnabledOk returns a tuple with the V3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Enabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetV3Enabled(v bool)`

SetV3Enabled sets V3Enabled field to given value.

### HasV3Enabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasV3Enabled() bool`

HasV3Enabled returns a boolean if a field has been set.

### GetVacmGroups

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetVacmGroups() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner`

GetVacmGroups returns the VacmGroups field if non-nil, zero value otherwise.

### GetVacmGroupsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) GetVacmGroupsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner, bool)`

GetVacmGroupsOk returns a tuple with the VacmGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacmGroups

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) SetVacmGroups(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmpVacmGroupsInner)`

SetVacmGroups sets VacmGroups field to given value.

### HasVacmGroups

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSnmp) HasVacmGroups() bool`

HasVacmGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


