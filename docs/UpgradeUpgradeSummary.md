# UpgradeUpgradeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**EndOfLife** | Pointer to **bool** |  | [optional] 
**LastDiscoveredTs** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**LastRunningVersion** | Pointer to [**UpgradeSwVersion**](UpgradeSwVersion.md) |  | [optional] 
**LastUpgradeTs** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**ReadyForActivationVersion** | Pointer to [**UpgradeSwVersion**](UpgradeSwVersion.md) |  | [optional] 
**RolloutId** | Pointer to **int64** |  | [optional] 
**RolloutName** | Pointer to **string** |  | [optional] 
**RunningVersion** | Pointer to [**UpgradeSwVersion**](UpgradeSwVersion.md) |  | [optional] 
**Schedule** | Pointer to [**UpgradeSchedule**](UpgradeSchedule.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewUpgradeUpgradeSummary

`func NewUpgradeUpgradeSummary() *UpgradeUpgradeSummary`

NewUpgradeUpgradeSummary instantiates a new UpgradeUpgradeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeUpgradeSummaryWithDefaults

`func NewUpgradeUpgradeSummaryWithDefaults() *UpgradeUpgradeSummary`

NewUpgradeUpgradeSummaryWithDefaults instantiates a new UpgradeUpgradeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *UpgradeUpgradeSummary) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *UpgradeUpgradeSummary) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *UpgradeUpgradeSummary) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *UpgradeUpgradeSummary) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEndOfLife

`func (o *UpgradeUpgradeSummary) GetEndOfLife() bool`

GetEndOfLife returns the EndOfLife field if non-nil, zero value otherwise.

### GetEndOfLifeOk

`func (o *UpgradeUpgradeSummary) GetEndOfLifeOk() (*bool, bool)`

GetEndOfLifeOk returns a tuple with the EndOfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLife

`func (o *UpgradeUpgradeSummary) SetEndOfLife(v bool)`

SetEndOfLife sets EndOfLife field to given value.

### HasEndOfLife

`func (o *UpgradeUpgradeSummary) HasEndOfLife() bool`

HasEndOfLife returns a boolean if a field has been set.

### GetLastDiscoveredTs

`func (o *UpgradeUpgradeSummary) GetLastDiscoveredTs() GoogleProtobufTimestamp`

GetLastDiscoveredTs returns the LastDiscoveredTs field if non-nil, zero value otherwise.

### GetLastDiscoveredTsOk

`func (o *UpgradeUpgradeSummary) GetLastDiscoveredTsOk() (*GoogleProtobufTimestamp, bool)`

GetLastDiscoveredTsOk returns a tuple with the LastDiscoveredTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiscoveredTs

`func (o *UpgradeUpgradeSummary) SetLastDiscoveredTs(v GoogleProtobufTimestamp)`

SetLastDiscoveredTs sets LastDiscoveredTs field to given value.

### HasLastDiscoveredTs

`func (o *UpgradeUpgradeSummary) HasLastDiscoveredTs() bool`

HasLastDiscoveredTs returns a boolean if a field has been set.

### GetLastRunningVersion

`func (o *UpgradeUpgradeSummary) GetLastRunningVersion() UpgradeSwVersion`

GetLastRunningVersion returns the LastRunningVersion field if non-nil, zero value otherwise.

### GetLastRunningVersionOk

`func (o *UpgradeUpgradeSummary) GetLastRunningVersionOk() (*UpgradeSwVersion, bool)`

GetLastRunningVersionOk returns a tuple with the LastRunningVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunningVersion

`func (o *UpgradeUpgradeSummary) SetLastRunningVersion(v UpgradeSwVersion)`

SetLastRunningVersion sets LastRunningVersion field to given value.

### HasLastRunningVersion

`func (o *UpgradeUpgradeSummary) HasLastRunningVersion() bool`

HasLastRunningVersion returns a boolean if a field has been set.

### GetLastUpgradeTs

`func (o *UpgradeUpgradeSummary) GetLastUpgradeTs() GoogleProtobufTimestamp`

GetLastUpgradeTs returns the LastUpgradeTs field if non-nil, zero value otherwise.

### GetLastUpgradeTsOk

`func (o *UpgradeUpgradeSummary) GetLastUpgradeTsOk() (*GoogleProtobufTimestamp, bool)`

GetLastUpgradeTsOk returns a tuple with the LastUpgradeTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgradeTs

`func (o *UpgradeUpgradeSummary) SetLastUpgradeTs(v GoogleProtobufTimestamp)`

SetLastUpgradeTs sets LastUpgradeTs field to given value.

### HasLastUpgradeTs

`func (o *UpgradeUpgradeSummary) HasLastUpgradeTs() bool`

HasLastUpgradeTs returns a boolean if a field has been set.

### GetReadyForActivationVersion

`func (o *UpgradeUpgradeSummary) GetReadyForActivationVersion() UpgradeSwVersion`

GetReadyForActivationVersion returns the ReadyForActivationVersion field if non-nil, zero value otherwise.

### GetReadyForActivationVersionOk

`func (o *UpgradeUpgradeSummary) GetReadyForActivationVersionOk() (*UpgradeSwVersion, bool)`

GetReadyForActivationVersionOk returns a tuple with the ReadyForActivationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyForActivationVersion

`func (o *UpgradeUpgradeSummary) SetReadyForActivationVersion(v UpgradeSwVersion)`

SetReadyForActivationVersion sets ReadyForActivationVersion field to given value.

### HasReadyForActivationVersion

`func (o *UpgradeUpgradeSummary) HasReadyForActivationVersion() bool`

HasReadyForActivationVersion returns a boolean if a field has been set.

### GetRolloutId

`func (o *UpgradeUpgradeSummary) GetRolloutId() int64`

GetRolloutId returns the RolloutId field if non-nil, zero value otherwise.

### GetRolloutIdOk

`func (o *UpgradeUpgradeSummary) GetRolloutIdOk() (*int64, bool)`

GetRolloutIdOk returns a tuple with the RolloutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutId

`func (o *UpgradeUpgradeSummary) SetRolloutId(v int64)`

SetRolloutId sets RolloutId field to given value.

### HasRolloutId

`func (o *UpgradeUpgradeSummary) HasRolloutId() bool`

HasRolloutId returns a boolean if a field has been set.

### GetRolloutName

`func (o *UpgradeUpgradeSummary) GetRolloutName() string`

GetRolloutName returns the RolloutName field if non-nil, zero value otherwise.

### GetRolloutNameOk

`func (o *UpgradeUpgradeSummary) GetRolloutNameOk() (*string, bool)`

GetRolloutNameOk returns a tuple with the RolloutName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutName

`func (o *UpgradeUpgradeSummary) SetRolloutName(v string)`

SetRolloutName sets RolloutName field to given value.

### HasRolloutName

`func (o *UpgradeUpgradeSummary) HasRolloutName() bool`

HasRolloutName returns a boolean if a field has been set.

### GetRunningVersion

`func (o *UpgradeUpgradeSummary) GetRunningVersion() UpgradeSwVersion`

GetRunningVersion returns the RunningVersion field if non-nil, zero value otherwise.

### GetRunningVersionOk

`func (o *UpgradeUpgradeSummary) GetRunningVersionOk() (*UpgradeSwVersion, bool)`

GetRunningVersionOk returns a tuple with the RunningVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningVersion

`func (o *UpgradeUpgradeSummary) SetRunningVersion(v UpgradeSwVersion)`

SetRunningVersion sets RunningVersion field to given value.

### HasRunningVersion

`func (o *UpgradeUpgradeSummary) HasRunningVersion() bool`

HasRunningVersion returns a boolean if a field has been set.

### GetSchedule

`func (o *UpgradeUpgradeSummary) GetSchedule() UpgradeSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpgradeUpgradeSummary) GetScheduleOk() (*UpgradeSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpgradeUpgradeSummary) SetSchedule(v UpgradeSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpgradeUpgradeSummary) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStatus

`func (o *UpgradeUpgradeSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpgradeUpgradeSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpgradeUpgradeSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpgradeUpgradeSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


