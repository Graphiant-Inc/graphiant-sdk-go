# ManaV2B2bExtranetMatchConsumerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerId** | Pointer to **int64** |  | [optional] 
**ConsumerPrefixes** | Pointer to **[]string** |  | [optional] 
**Customer** | Pointer to [**ManaV2B2BExtranetMatchConsumerDetailsCustomer**](ManaV2B2BExtranetMatchConsumerDetailsCustomer.md) |  | [optional] 
**OldConsumerPrefixes** | Pointer to **[]string** |  | [optional] 
**OldServicePrefixes** | Pointer to [**[]ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix**](ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix.md) |  | [optional] 
**Service** | Pointer to [**ManaV2B2BExtranetMatchConsumerDetailsService**](ManaV2B2BExtranetMatchConsumerDetailsService.md) |  | [optional] 
**ServicePrefixes** | Pointer to [**[]ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix**](ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix.md) |  | [optional] 

## Methods

### NewManaV2B2bExtranetMatchConsumerDetails

`func NewManaV2B2bExtranetMatchConsumerDetails() *ManaV2B2bExtranetMatchConsumerDetails`

NewManaV2B2bExtranetMatchConsumerDetails instantiates a new ManaV2B2bExtranetMatchConsumerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2B2bExtranetMatchConsumerDetailsWithDefaults

`func NewManaV2B2bExtranetMatchConsumerDetailsWithDefaults() *ManaV2B2bExtranetMatchConsumerDetails`

NewManaV2B2bExtranetMatchConsumerDetailsWithDefaults instantiates a new ManaV2B2bExtranetMatchConsumerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerId

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetConsumerId() int64`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetConsumerIdOk() (*int64, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *ManaV2B2bExtranetMatchConsumerDetails) SetConsumerId(v int64)`

SetConsumerId sets ConsumerId field to given value.

### HasConsumerId

`func (o *ManaV2B2bExtranetMatchConsumerDetails) HasConsumerId() bool`

HasConsumerId returns a boolean if a field has been set.

### GetConsumerPrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetConsumerPrefixes() []string`

GetConsumerPrefixes returns the ConsumerPrefixes field if non-nil, zero value otherwise.

### GetConsumerPrefixesOk

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetConsumerPrefixesOk() (*[]string, bool)`

GetConsumerPrefixesOk returns a tuple with the ConsumerPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerPrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) SetConsumerPrefixes(v []string)`

SetConsumerPrefixes sets ConsumerPrefixes field to given value.

### HasConsumerPrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) HasConsumerPrefixes() bool`

HasConsumerPrefixes returns a boolean if a field has been set.

### GetCustomer

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetCustomer() ManaV2B2BExtranetMatchConsumerDetailsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetCustomerOk() (*ManaV2B2BExtranetMatchConsumerDetailsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *ManaV2B2bExtranetMatchConsumerDetails) SetCustomer(v ManaV2B2BExtranetMatchConsumerDetailsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *ManaV2B2bExtranetMatchConsumerDetails) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetOldConsumerPrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetOldConsumerPrefixes() []string`

GetOldConsumerPrefixes returns the OldConsumerPrefixes field if non-nil, zero value otherwise.

### GetOldConsumerPrefixesOk

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetOldConsumerPrefixesOk() (*[]string, bool)`

GetOldConsumerPrefixesOk returns a tuple with the OldConsumerPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldConsumerPrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) SetOldConsumerPrefixes(v []string)`

SetOldConsumerPrefixes sets OldConsumerPrefixes field to given value.

### HasOldConsumerPrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) HasOldConsumerPrefixes() bool`

HasOldConsumerPrefixes returns a boolean if a field has been set.

### GetOldServicePrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetOldServicePrefixes() []ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix`

GetOldServicePrefixes returns the OldServicePrefixes field if non-nil, zero value otherwise.

### GetOldServicePrefixesOk

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetOldServicePrefixesOk() (*[]ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix, bool)`

GetOldServicePrefixesOk returns a tuple with the OldServicePrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldServicePrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) SetOldServicePrefixes(v []ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix)`

SetOldServicePrefixes sets OldServicePrefixes field to given value.

### HasOldServicePrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) HasOldServicePrefixes() bool`

HasOldServicePrefixes returns a boolean if a field has been set.

### GetService

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetService() ManaV2B2BExtranetMatchConsumerDetailsService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetServiceOk() (*ManaV2B2BExtranetMatchConsumerDetailsService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ManaV2B2bExtranetMatchConsumerDetails) SetService(v ManaV2B2BExtranetMatchConsumerDetailsService)`

SetService sets Service field to given value.

### HasService

`func (o *ManaV2B2bExtranetMatchConsumerDetails) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServicePrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetServicePrefixes() []ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix`

GetServicePrefixes returns the ServicePrefixes field if non-nil, zero value otherwise.

### GetServicePrefixesOk

`func (o *ManaV2B2bExtranetMatchConsumerDetails) GetServicePrefixesOk() (*[]ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix, bool)`

GetServicePrefixesOk returns a tuple with the ServicePrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) SetServicePrefixes(v []ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix)`

SetServicePrefixes sets ServicePrefixes field to given value.

### HasServicePrefixes

`func (o *ManaV2B2bExtranetMatchConsumerDetails) HasServicePrefixes() bool`

HasServicePrefixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


