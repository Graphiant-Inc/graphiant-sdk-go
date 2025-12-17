# V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BwAllocation** | Pointer to **int32** | the bandwidth allocation for the entity in kbps | [optional] 
**DlStats** | Pointer to [**[]IpfixStats**](IpfixStats.md) |  | [optional] 
**UlStats** | Pointer to [**[]IpfixStats**](IpfixStats.md) |  | [optional] 

## Methods

### NewV1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse

`func NewV1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse() *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse`

NewV1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse instantiates a new V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponseWithDefaults

`func NewV1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponseWithDefaults() *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse`

NewV1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponseWithDefaults instantiates a new V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBwAllocation

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) GetBwAllocation() int32`

GetBwAllocation returns the BwAllocation field if non-nil, zero value otherwise.

### GetBwAllocationOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) GetBwAllocationOk() (*int32, bool)`

GetBwAllocationOk returns a tuple with the BwAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwAllocation

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) SetBwAllocation(v int32)`

SetBwAllocation sets BwAllocation field to given value.

### HasBwAllocation

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) HasBwAllocation() bool`

HasBwAllocation returns a boolean if a field has been set.

### GetDlStats

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) GetDlStats() []IpfixStats`

GetDlStats returns the DlStats field if non-nil, zero value otherwise.

### GetDlStatsOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) GetDlStatsOk() (*[]IpfixStats, bool)`

GetDlStatsOk returns a tuple with the DlStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlStats

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) SetDlStats(v []IpfixStats)`

SetDlStats sets DlStats field to given value.

### HasDlStats

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) HasDlStats() bool`

HasDlStats returns a boolean if a field has been set.

### GetUlStats

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) GetUlStats() []IpfixStats`

GetUlStats returns the UlStats field if non-nil, zero value otherwise.

### GetUlStatsOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) GetUlStatsOk() (*[]IpfixStats, bool)`

GetUlStatsOk returns a tuple with the UlStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlStats

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) SetUlStats(v []IpfixStats)`

SetUlStats sets UlStats field to given value.

### HasUlStats

`func (o *V1ExtranetB2bMonitoringPeeringServiceBandwidthUsagePostResponse) HasUlStats() bool`

HasUlStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


