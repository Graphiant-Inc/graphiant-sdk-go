# V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BwusageRoleSummary** | Pointer to [**[]V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryBwusageRoleSummaryInner**](V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryBwusageRoleSummaryInner.md) |  | [optional] 
**BwusageTopRegions** | Pointer to [**[]V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryBwusageTopRegionsInner**](V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryBwusageTopRegionsInner.md) |  | [optional] 
**MinTime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**PercentChanged** | Pointer to **float64** |  | [optional] 
**ProviderCount** | Pointer to **int64** |  | [optional] 
**RegionCount** | Pointer to **int64** |  | [optional] 
**SiteCount** | Pointer to **int64** |  | [optional] 
**UsageKbps** | Pointer to **float64** |  | [optional] 

## Methods

### NewV1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary

`func NewV1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary() *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary`

NewV1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary instantiates a new V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryWithDefaults

`func NewV1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryWithDefaults() *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary`

NewV1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryWithDefaults instantiates a new V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBwusageRoleSummary

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetBwusageRoleSummary() []V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryBwusageRoleSummaryInner`

GetBwusageRoleSummary returns the BwusageRoleSummary field if non-nil, zero value otherwise.

### GetBwusageRoleSummaryOk

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetBwusageRoleSummaryOk() (*[]V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryBwusageRoleSummaryInner, bool)`

GetBwusageRoleSummaryOk returns a tuple with the BwusageRoleSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwusageRoleSummary

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) SetBwusageRoleSummary(v []V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryBwusageRoleSummaryInner)`

SetBwusageRoleSummary sets BwusageRoleSummary field to given value.

### HasBwusageRoleSummary

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) HasBwusageRoleSummary() bool`

HasBwusageRoleSummary returns a boolean if a field has been set.

### GetBwusageTopRegions

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetBwusageTopRegions() []V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryBwusageTopRegionsInner`

GetBwusageTopRegions returns the BwusageTopRegions field if non-nil, zero value otherwise.

### GetBwusageTopRegionsOk

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetBwusageTopRegionsOk() (*[]V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryBwusageTopRegionsInner, bool)`

GetBwusageTopRegionsOk returns a tuple with the BwusageTopRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwusageTopRegions

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) SetBwusageTopRegions(v []V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummaryBwusageTopRegionsInner)`

SetBwusageTopRegions sets BwusageTopRegions field to given value.

### HasBwusageTopRegions

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) HasBwusageTopRegions() bool`

HasBwusageTopRegions returns a boolean if a field has been set.

### GetMinTime

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetMinTime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetMinTime returns the MinTime field if non-nil, zero value otherwise.

### GetMinTimeOk

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetMinTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetMinTimeOk returns a tuple with the MinTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTime

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) SetMinTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetMinTime sets MinTime field to given value.

### HasMinTime

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) HasMinTime() bool`

HasMinTime returns a boolean if a field has been set.

### GetPercentChanged

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetPercentChanged() float64`

GetPercentChanged returns the PercentChanged field if non-nil, zero value otherwise.

### GetPercentChangedOk

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetPercentChangedOk() (*float64, bool)`

GetPercentChangedOk returns a tuple with the PercentChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChanged

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) SetPercentChanged(v float64)`

SetPercentChanged sets PercentChanged field to given value.

### HasPercentChanged

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) HasPercentChanged() bool`

HasPercentChanged returns a boolean if a field has been set.

### GetProviderCount

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetProviderCount() int64`

GetProviderCount returns the ProviderCount field if non-nil, zero value otherwise.

### GetProviderCountOk

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetProviderCountOk() (*int64, bool)`

GetProviderCountOk returns a tuple with the ProviderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCount

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) SetProviderCount(v int64)`

SetProviderCount sets ProviderCount field to given value.

### HasProviderCount

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) HasProviderCount() bool`

HasProviderCount returns a boolean if a field has been set.

### GetRegionCount

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetRegionCount() int64`

GetRegionCount returns the RegionCount field if non-nil, zero value otherwise.

### GetRegionCountOk

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetRegionCountOk() (*int64, bool)`

GetRegionCountOk returns a tuple with the RegionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCount

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) SetRegionCount(v int64)`

SetRegionCount sets RegionCount field to given value.

### HasRegionCount

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) HasRegionCount() bool`

HasRegionCount returns a boolean if a field has been set.

### GetSiteCount

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetSiteCount() int64`

GetSiteCount returns the SiteCount field if non-nil, zero value otherwise.

### GetSiteCountOk

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetSiteCountOk() (*int64, bool)`

GetSiteCountOk returns a tuple with the SiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCount

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) SetSiteCount(v int64)`

SetSiteCount sets SiteCount field to given value.

### HasSiteCount

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) HasSiteCount() bool`

HasSiteCount returns a boolean if a field has been set.

### GetUsageKbps

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetUsageKbps() float64`

GetUsageKbps returns the UsageKbps field if non-nil, zero value otherwise.

### GetUsageKbpsOk

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) GetUsageKbpsOk() (*float64, bool)`

GetUsageKbpsOk returns a tuple with the UsageKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageKbps

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) SetUsageKbps(v float64)`

SetUsageKbps sets UsageKbps field to given value.

### HasUsageKbps

`func (o *V1BwtrackerEnterpriseSummaryPost200ResponseBwusageSummary) HasUsageKbps() bool`

HasUsageKbps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


