# V1DevicesUpgradeSchedulePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**DeviceIds** | Pointer to **[]int64** |  | [optional] 
**Ts** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Version** | Pointer to [**V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion**](V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion.md) |  | [optional] 

## Methods

### NewV1DevicesUpgradeSchedulePutRequest

`func NewV1DevicesUpgradeSchedulePutRequest() *V1DevicesUpgradeSchedulePutRequest`

NewV1DevicesUpgradeSchedulePutRequest instantiates a new V1DevicesUpgradeSchedulePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesUpgradeSchedulePutRequestWithDefaults

`func NewV1DevicesUpgradeSchedulePutRequestWithDefaults() *V1DevicesUpgradeSchedulePutRequest`

NewV1DevicesUpgradeSchedulePutRequestWithDefaults instantiates a new V1DevicesUpgradeSchedulePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *V1DevicesUpgradeSchedulePutRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V1DevicesUpgradeSchedulePutRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V1DevicesUpgradeSchedulePutRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *V1DevicesUpgradeSchedulePutRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDeviceIds

`func (o *V1DevicesUpgradeSchedulePutRequest) GetDeviceIds() []int64`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *V1DevicesUpgradeSchedulePutRequest) GetDeviceIdsOk() (*[]int64, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *V1DevicesUpgradeSchedulePutRequest) SetDeviceIds(v []int64)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *V1DevicesUpgradeSchedulePutRequest) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetTs

`func (o *V1DevicesUpgradeSchedulePutRequest) GetTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *V1DevicesUpgradeSchedulePutRequest) GetTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *V1DevicesUpgradeSchedulePutRequest) SetTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetTs sets Ts field to given value.

### HasTs

`func (o *V1DevicesUpgradeSchedulePutRequest) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetVersion

`func (o *V1DevicesUpgradeSchedulePutRequest) GetVersion() V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1DevicesUpgradeSchedulePutRequest) GetVersionOk() (*V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1DevicesUpgradeSchedulePutRequest) SetVersion(v V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1DevicesUpgradeSchedulePutRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


