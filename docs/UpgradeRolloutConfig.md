# UpgradeRolloutConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Upgrade action to perform (e.g. install+activate, install only, activate, auto-upgrade). (required) | 
**Description** | Pointer to **string** | Optional longer description of the rollout. | [optional] 
**DeviceIds** | Pointer to **[]int64** |  | [optional] 
**Name** | **string** | Human-readable rollout name unique within the enterprise. (required) | 
**Release** | **string** | Target software release for devices in this rollout. (required) | 
**Schedule** | Pointer to [**UpgradeRecurringSchedule**](UpgradeRecurringSchedule.md) |  | [optional] 

## Methods

### NewUpgradeRolloutConfig

`func NewUpgradeRolloutConfig(action string, name string, release string, ) *UpgradeRolloutConfig`

NewUpgradeRolloutConfig instantiates a new UpgradeRolloutConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeRolloutConfigWithDefaults

`func NewUpgradeRolloutConfigWithDefaults() *UpgradeRolloutConfig`

NewUpgradeRolloutConfigWithDefaults instantiates a new UpgradeRolloutConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpgradeRolloutConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpgradeRolloutConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpgradeRolloutConfig) SetAction(v string)`

SetAction sets Action field to given value.


### GetDescription

`func (o *UpgradeRolloutConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpgradeRolloutConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpgradeRolloutConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpgradeRolloutConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceIds

`func (o *UpgradeRolloutConfig) GetDeviceIds() []int64`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *UpgradeRolloutConfig) GetDeviceIdsOk() (*[]int64, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *UpgradeRolloutConfig) SetDeviceIds(v []int64)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *UpgradeRolloutConfig) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetName

`func (o *UpgradeRolloutConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpgradeRolloutConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpgradeRolloutConfig) SetName(v string)`

SetName sets Name field to given value.


### GetRelease

`func (o *UpgradeRolloutConfig) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *UpgradeRolloutConfig) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *UpgradeRolloutConfig) SetRelease(v string)`

SetRelease sets Release field to given value.


### GetSchedule

`func (o *UpgradeRolloutConfig) GetSchedule() UpgradeRecurringSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpgradeRolloutConfig) GetScheduleOk() (*UpgradeRecurringSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpgradeRolloutConfig) SetSchedule(v UpgradeRecurringSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpgradeRolloutConfig) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


