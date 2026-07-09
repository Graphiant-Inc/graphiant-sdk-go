# V1EnterpriseAllocationGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumptionSummary** | Pointer to [**ManaV2BandwidthConsumptionSummary**](ManaV2BandwidthConsumptionSummary.md) |  | [optional] 
**ConversionHolders** | Pointer to [**map[string]ManaV2AllocationConversionHolder**](ManaV2AllocationConversionHolder.md) |  | [optional] 
**RegionalAllocations** | Pointer to [**[]ManaV2RegionalAllocation**](ManaV2RegionalAllocation.md) |  | [optional] 
**ZeroTrustSummary** | Pointer to [**ManaV2ZeroTrustConsumptionSummary**](ManaV2ZeroTrustConsumptionSummary.md) |  | [optional] 

## Methods

### NewV1EnterpriseAllocationGetResponse

`func NewV1EnterpriseAllocationGetResponse() *V1EnterpriseAllocationGetResponse`

NewV1EnterpriseAllocationGetResponse instantiates a new V1EnterpriseAllocationGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EnterpriseAllocationGetResponseWithDefaults

`func NewV1EnterpriseAllocationGetResponseWithDefaults() *V1EnterpriseAllocationGetResponse`

NewV1EnterpriseAllocationGetResponseWithDefaults instantiates a new V1EnterpriseAllocationGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumptionSummary

`func (o *V1EnterpriseAllocationGetResponse) GetConsumptionSummary() ManaV2BandwidthConsumptionSummary`

GetConsumptionSummary returns the ConsumptionSummary field if non-nil, zero value otherwise.

### GetConsumptionSummaryOk

`func (o *V1EnterpriseAllocationGetResponse) GetConsumptionSummaryOk() (*ManaV2BandwidthConsumptionSummary, bool)`

GetConsumptionSummaryOk returns a tuple with the ConsumptionSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionSummary

`func (o *V1EnterpriseAllocationGetResponse) SetConsumptionSummary(v ManaV2BandwidthConsumptionSummary)`

SetConsumptionSummary sets ConsumptionSummary field to given value.

### HasConsumptionSummary

`func (o *V1EnterpriseAllocationGetResponse) HasConsumptionSummary() bool`

HasConsumptionSummary returns a boolean if a field has been set.

### GetConversionHolders

`func (o *V1EnterpriseAllocationGetResponse) GetConversionHolders() map[string]ManaV2AllocationConversionHolder`

GetConversionHolders returns the ConversionHolders field if non-nil, zero value otherwise.

### GetConversionHoldersOk

`func (o *V1EnterpriseAllocationGetResponse) GetConversionHoldersOk() (*map[string]ManaV2AllocationConversionHolder, bool)`

GetConversionHoldersOk returns a tuple with the ConversionHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionHolders

`func (o *V1EnterpriseAllocationGetResponse) SetConversionHolders(v map[string]ManaV2AllocationConversionHolder)`

SetConversionHolders sets ConversionHolders field to given value.

### HasConversionHolders

`func (o *V1EnterpriseAllocationGetResponse) HasConversionHolders() bool`

HasConversionHolders returns a boolean if a field has been set.

### GetRegionalAllocations

`func (o *V1EnterpriseAllocationGetResponse) GetRegionalAllocations() []ManaV2RegionalAllocation`

GetRegionalAllocations returns the RegionalAllocations field if non-nil, zero value otherwise.

### GetRegionalAllocationsOk

`func (o *V1EnterpriseAllocationGetResponse) GetRegionalAllocationsOk() (*[]ManaV2RegionalAllocation, bool)`

GetRegionalAllocationsOk returns a tuple with the RegionalAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalAllocations

`func (o *V1EnterpriseAllocationGetResponse) SetRegionalAllocations(v []ManaV2RegionalAllocation)`

SetRegionalAllocations sets RegionalAllocations field to given value.

### HasRegionalAllocations

`func (o *V1EnterpriseAllocationGetResponse) HasRegionalAllocations() bool`

HasRegionalAllocations returns a boolean if a field has been set.

### GetZeroTrustSummary

`func (o *V1EnterpriseAllocationGetResponse) GetZeroTrustSummary() ManaV2ZeroTrustConsumptionSummary`

GetZeroTrustSummary returns the ZeroTrustSummary field if non-nil, zero value otherwise.

### GetZeroTrustSummaryOk

`func (o *V1EnterpriseAllocationGetResponse) GetZeroTrustSummaryOk() (*ManaV2ZeroTrustConsumptionSummary, bool)`

GetZeroTrustSummaryOk returns a tuple with the ZeroTrustSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroTrustSummary

`func (o *V1EnterpriseAllocationGetResponse) SetZeroTrustSummary(v ManaV2ZeroTrustConsumptionSummary)`

SetZeroTrustSummary sets ZeroTrustSummary field to given value.

### HasZeroTrustSummary

`func (o *V1EnterpriseAllocationGetResponse) HasZeroTrustSummary() bool`

HasZeroTrustSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


