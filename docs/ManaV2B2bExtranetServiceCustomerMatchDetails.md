# ManaV2B2bExtranetServiceCustomerMatchDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerId** | Pointer to **int64** |  | [optional] 
**ConsumerPrefixes** | Pointer to **[]string** |  | [optional] 
**Customer** | Pointer to [**ManaV2B2BExtranetServiceCustomerMatchDetailsCustomer**](ManaV2B2BExtranetServiceCustomerMatchDetailsCustomer.md) |  | [optional] 
**OldConsumerPrefixes** | Pointer to **[]string** |  | [optional] 
**OldServicePrefixes** | Pointer to [**[]ManaV2B2BExtranetServiceCustomerMatchDetailsProducerPrefix**](ManaV2B2BExtranetServiceCustomerMatchDetailsProducerPrefix.md) |  | [optional] 
**Service** | Pointer to [**ManaV2B2BExtranetServiceCustomerMatchDetailsService**](ManaV2B2BExtranetServiceCustomerMatchDetailsService.md) |  | [optional] 
**ServicePrefixes** | Pointer to [**[]ManaV2B2BExtranetServiceCustomerMatchDetailsProducerPrefix**](ManaV2B2BExtranetServiceCustomerMatchDetailsProducerPrefix.md) |  | [optional] 

## Methods

### NewManaV2B2bExtranetServiceCustomerMatchDetails

`func NewManaV2B2bExtranetServiceCustomerMatchDetails() *ManaV2B2bExtranetServiceCustomerMatchDetails`

NewManaV2B2bExtranetServiceCustomerMatchDetails instantiates a new ManaV2B2bExtranetServiceCustomerMatchDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2B2bExtranetServiceCustomerMatchDetailsWithDefaults

`func NewManaV2B2bExtranetServiceCustomerMatchDetailsWithDefaults() *ManaV2B2bExtranetServiceCustomerMatchDetails`

NewManaV2B2bExtranetServiceCustomerMatchDetailsWithDefaults instantiates a new ManaV2B2bExtranetServiceCustomerMatchDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerId

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetConsumerId() int64`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetConsumerIdOk() (*int64, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) SetConsumerId(v int64)`

SetConsumerId sets ConsumerId field to given value.

### HasConsumerId

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) HasConsumerId() bool`

HasConsumerId returns a boolean if a field has been set.

### GetConsumerPrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetConsumerPrefixes() []string`

GetConsumerPrefixes returns the ConsumerPrefixes field if non-nil, zero value otherwise.

### GetConsumerPrefixesOk

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetConsumerPrefixesOk() (*[]string, bool)`

GetConsumerPrefixesOk returns a tuple with the ConsumerPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerPrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) SetConsumerPrefixes(v []string)`

SetConsumerPrefixes sets ConsumerPrefixes field to given value.

### HasConsumerPrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) HasConsumerPrefixes() bool`

HasConsumerPrefixes returns a boolean if a field has been set.

### GetCustomer

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetCustomer() ManaV2B2BExtranetServiceCustomerMatchDetailsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetCustomerOk() (*ManaV2B2BExtranetServiceCustomerMatchDetailsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) SetCustomer(v ManaV2B2BExtranetServiceCustomerMatchDetailsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetOldConsumerPrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetOldConsumerPrefixes() []string`

GetOldConsumerPrefixes returns the OldConsumerPrefixes field if non-nil, zero value otherwise.

### GetOldConsumerPrefixesOk

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetOldConsumerPrefixesOk() (*[]string, bool)`

GetOldConsumerPrefixesOk returns a tuple with the OldConsumerPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldConsumerPrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) SetOldConsumerPrefixes(v []string)`

SetOldConsumerPrefixes sets OldConsumerPrefixes field to given value.

### HasOldConsumerPrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) HasOldConsumerPrefixes() bool`

HasOldConsumerPrefixes returns a boolean if a field has been set.

### GetOldServicePrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetOldServicePrefixes() []ManaV2B2BExtranetServiceCustomerMatchDetailsProducerPrefix`

GetOldServicePrefixes returns the OldServicePrefixes field if non-nil, zero value otherwise.

### GetOldServicePrefixesOk

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetOldServicePrefixesOk() (*[]ManaV2B2BExtranetServiceCustomerMatchDetailsProducerPrefix, bool)`

GetOldServicePrefixesOk returns a tuple with the OldServicePrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldServicePrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) SetOldServicePrefixes(v []ManaV2B2BExtranetServiceCustomerMatchDetailsProducerPrefix)`

SetOldServicePrefixes sets OldServicePrefixes field to given value.

### HasOldServicePrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) HasOldServicePrefixes() bool`

HasOldServicePrefixes returns a boolean if a field has been set.

### GetService

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetService() ManaV2B2BExtranetServiceCustomerMatchDetailsService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetServiceOk() (*ManaV2B2BExtranetServiceCustomerMatchDetailsService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) SetService(v ManaV2B2BExtranetServiceCustomerMatchDetailsService)`

SetService sets Service field to given value.

### HasService

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServicePrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetServicePrefixes() []ManaV2B2BExtranetServiceCustomerMatchDetailsProducerPrefix`

GetServicePrefixes returns the ServicePrefixes field if non-nil, zero value otherwise.

### GetServicePrefixesOk

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) GetServicePrefixesOk() (*[]ManaV2B2BExtranetServiceCustomerMatchDetailsProducerPrefix, bool)`

GetServicePrefixesOk returns a tuple with the ServicePrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) SetServicePrefixes(v []ManaV2B2BExtranetServiceCustomerMatchDetailsProducerPrefix)`

SetServicePrefixes sets ServicePrefixes field to given value.

### HasServicePrefixes

`func (o *ManaV2B2bExtranetServiceCustomerMatchDetails) HasServicePrefixes() bool`

HasServicePrefixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


