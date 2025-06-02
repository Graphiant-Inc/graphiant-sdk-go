# V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**LastDiscoveredTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**LastUpgradeTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**RunningVersion** | Pointer to [**V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion**](V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion.md) |  | [optional] 
**Schedule** | Pointer to [**V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummarySchedule**](V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummarySchedule.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary

`func NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary() *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary`

NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary instantiates a new V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryWithDefaults

`func NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryWithDefaults() *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary`

NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryWithDefaults instantiates a new V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetLastDiscoveredTs

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetLastDiscoveredTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetLastDiscoveredTs returns the LastDiscoveredTs field if non-nil, zero value otherwise.

### GetLastDiscoveredTsOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetLastDiscoveredTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetLastDiscoveredTsOk returns a tuple with the LastDiscoveredTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiscoveredTs

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) SetLastDiscoveredTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetLastDiscoveredTs sets LastDiscoveredTs field to given value.

### HasLastDiscoveredTs

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) HasLastDiscoveredTs() bool`

HasLastDiscoveredTs returns a boolean if a field has been set.

### GetLastUpgradeTs

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetLastUpgradeTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetLastUpgradeTs returns the LastUpgradeTs field if non-nil, zero value otherwise.

### GetLastUpgradeTsOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetLastUpgradeTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetLastUpgradeTsOk returns a tuple with the LastUpgradeTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgradeTs

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) SetLastUpgradeTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetLastUpgradeTs sets LastUpgradeTs field to given value.

### HasLastUpgradeTs

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) HasLastUpgradeTs() bool`

HasLastUpgradeTs returns a boolean if a field has been set.

### GetRunningVersion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetRunningVersion() V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion`

GetRunningVersion returns the RunningVersion field if non-nil, zero value otherwise.

### GetRunningVersionOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetRunningVersionOk() (*V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion, bool)`

GetRunningVersionOk returns a tuple with the RunningVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningVersion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) SetRunningVersion(v V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion)`

SetRunningVersion sets RunningVersion field to given value.

### HasRunningVersion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) HasRunningVersion() bool`

HasRunningVersion returns a boolean if a field has been set.

### GetSchedule

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetSchedule() V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummarySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetScheduleOk() (*V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummarySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) SetSchedule(v V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummarySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStatus

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


