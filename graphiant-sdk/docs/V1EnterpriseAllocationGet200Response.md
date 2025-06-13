# V1EnterpriseAllocationGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumptionSummary** | Pointer to [**V1EnterpriseAllocationGet200ResponseConsumptionSummary**](V1EnterpriseAllocationGet200ResponseConsumptionSummary.md) |  | [optional] 
**ConversionHolders** | Pointer to [**map[string]V1EnterpriseAllocationGet200ResponseConversionHoldersValue**](V1EnterpriseAllocationGet200ResponseConversionHoldersValue.md) |  | [optional] 
**RegionalAllocations** | Pointer to [**[]V1EnterpriseAllocationGet200ResponseRegionalAllocationsInner**](V1EnterpriseAllocationGet200ResponseRegionalAllocationsInner.md) |  | [optional] 

## Methods

### NewV1EnterpriseAllocationGet200Response

`func NewV1EnterpriseAllocationGet200Response() *V1EnterpriseAllocationGet200Response`

NewV1EnterpriseAllocationGet200Response instantiates a new V1EnterpriseAllocationGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EnterpriseAllocationGet200ResponseWithDefaults

`func NewV1EnterpriseAllocationGet200ResponseWithDefaults() *V1EnterpriseAllocationGet200Response`

NewV1EnterpriseAllocationGet200ResponseWithDefaults instantiates a new V1EnterpriseAllocationGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumptionSummary

`func (o *V1EnterpriseAllocationGet200Response) GetConsumptionSummary() V1EnterpriseAllocationGet200ResponseConsumptionSummary`

GetConsumptionSummary returns the ConsumptionSummary field if non-nil, zero value otherwise.

### GetConsumptionSummaryOk

`func (o *V1EnterpriseAllocationGet200Response) GetConsumptionSummaryOk() (*V1EnterpriseAllocationGet200ResponseConsumptionSummary, bool)`

GetConsumptionSummaryOk returns a tuple with the ConsumptionSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionSummary

`func (o *V1EnterpriseAllocationGet200Response) SetConsumptionSummary(v V1EnterpriseAllocationGet200ResponseConsumptionSummary)`

SetConsumptionSummary sets ConsumptionSummary field to given value.

### HasConsumptionSummary

`func (o *V1EnterpriseAllocationGet200Response) HasConsumptionSummary() bool`

HasConsumptionSummary returns a boolean if a field has been set.

### GetConversionHolders

`func (o *V1EnterpriseAllocationGet200Response) GetConversionHolders() map[string]V1EnterpriseAllocationGet200ResponseConversionHoldersValue`

GetConversionHolders returns the ConversionHolders field if non-nil, zero value otherwise.

### GetConversionHoldersOk

`func (o *V1EnterpriseAllocationGet200Response) GetConversionHoldersOk() (*map[string]V1EnterpriseAllocationGet200ResponseConversionHoldersValue, bool)`

GetConversionHoldersOk returns a tuple with the ConversionHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionHolders

`func (o *V1EnterpriseAllocationGet200Response) SetConversionHolders(v map[string]V1EnterpriseAllocationGet200ResponseConversionHoldersValue)`

SetConversionHolders sets ConversionHolders field to given value.

### HasConversionHolders

`func (o *V1EnterpriseAllocationGet200Response) HasConversionHolders() bool`

HasConversionHolders returns a boolean if a field has been set.

### GetRegionalAllocations

`func (o *V1EnterpriseAllocationGet200Response) GetRegionalAllocations() []V1EnterpriseAllocationGet200ResponseRegionalAllocationsInner`

GetRegionalAllocations returns the RegionalAllocations field if non-nil, zero value otherwise.

### GetRegionalAllocationsOk

`func (o *V1EnterpriseAllocationGet200Response) GetRegionalAllocationsOk() (*[]V1EnterpriseAllocationGet200ResponseRegionalAllocationsInner, bool)`

GetRegionalAllocationsOk returns a tuple with the RegionalAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalAllocations

`func (o *V1EnterpriseAllocationGet200Response) SetRegionalAllocations(v []V1EnterpriseAllocationGet200ResponseRegionalAllocationsInner)`

SetRegionalAllocations sets RegionalAllocations field to given value.

### HasRegionalAllocations

`func (o *V1EnterpriseAllocationGet200Response) HasRegionalAllocations() bool`

HasRegionalAllocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


