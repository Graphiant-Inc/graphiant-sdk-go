# V2ExtranetSitesUsagePostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BwAllocation** | Pointer to **int32** | the bandwidth allocation for the entity in kbps | [optional] 
**DlStats** | Pointer to [**[]IpfixStats**](IpfixStats.md) |  | [optional] 
**UlStats** | Pointer to [**[]IpfixStats**](IpfixStats.md) |  | [optional] 

## Methods

### NewV2ExtranetSitesUsagePostResponse

`func NewV2ExtranetSitesUsagePostResponse() *V2ExtranetSitesUsagePostResponse`

NewV2ExtranetSitesUsagePostResponse instantiates a new V2ExtranetSitesUsagePostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ExtranetSitesUsagePostResponseWithDefaults

`func NewV2ExtranetSitesUsagePostResponseWithDefaults() *V2ExtranetSitesUsagePostResponse`

NewV2ExtranetSitesUsagePostResponseWithDefaults instantiates a new V2ExtranetSitesUsagePostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBwAllocation

`func (o *V2ExtranetSitesUsagePostResponse) GetBwAllocation() int32`

GetBwAllocation returns the BwAllocation field if non-nil, zero value otherwise.

### GetBwAllocationOk

`func (o *V2ExtranetSitesUsagePostResponse) GetBwAllocationOk() (*int32, bool)`

GetBwAllocationOk returns a tuple with the BwAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwAllocation

`func (o *V2ExtranetSitesUsagePostResponse) SetBwAllocation(v int32)`

SetBwAllocation sets BwAllocation field to given value.

### HasBwAllocation

`func (o *V2ExtranetSitesUsagePostResponse) HasBwAllocation() bool`

HasBwAllocation returns a boolean if a field has been set.

### GetDlStats

`func (o *V2ExtranetSitesUsagePostResponse) GetDlStats() []IpfixStats`

GetDlStats returns the DlStats field if non-nil, zero value otherwise.

### GetDlStatsOk

`func (o *V2ExtranetSitesUsagePostResponse) GetDlStatsOk() (*[]IpfixStats, bool)`

GetDlStatsOk returns a tuple with the DlStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlStats

`func (o *V2ExtranetSitesUsagePostResponse) SetDlStats(v []IpfixStats)`

SetDlStats sets DlStats field to given value.

### HasDlStats

`func (o *V2ExtranetSitesUsagePostResponse) HasDlStats() bool`

HasDlStats returns a boolean if a field has been set.

### GetUlStats

`func (o *V2ExtranetSitesUsagePostResponse) GetUlStats() []IpfixStats`

GetUlStats returns the UlStats field if non-nil, zero value otherwise.

### GetUlStatsOk

`func (o *V2ExtranetSitesUsagePostResponse) GetUlStatsOk() (*[]IpfixStats, bool)`

GetUlStatsOk returns a tuple with the UlStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlStats

`func (o *V2ExtranetSitesUsagePostResponse) SetUlStats(v []IpfixStats)`

SetUlStats sets UlStats field to given value.

### HasUlStats

`func (o *V2ExtranetSitesUsagePostResponse) HasUlStats() bool`

HasUlStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


