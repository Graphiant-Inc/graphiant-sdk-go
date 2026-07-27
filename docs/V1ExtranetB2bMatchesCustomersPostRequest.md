# V1ExtranetB2bMatchesCustomersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | Pointer to [**[]ManaV2ExtranetServiceProducerCustomer**](ManaV2ExtranetServiceProducerCustomer.md) |  | [optional] 
**ServiceId** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1ExtranetB2bMatchesCustomersPostRequest

`func NewV1ExtranetB2bMatchesCustomersPostRequest() *V1ExtranetB2bMatchesCustomersPostRequest`

NewV1ExtranetB2bMatchesCustomersPostRequest instantiates a new V1ExtranetB2bMatchesCustomersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetB2bMatchesCustomersPostRequestWithDefaults

`func NewV1ExtranetB2bMatchesCustomersPostRequestWithDefaults() *V1ExtranetB2bMatchesCustomersPostRequest`

NewV1ExtranetB2bMatchesCustomersPostRequestWithDefaults instantiates a new V1ExtranetB2bMatchesCustomersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomers

`func (o *V1ExtranetB2bMatchesCustomersPostRequest) GetCustomers() []ManaV2ExtranetServiceProducerCustomer`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *V1ExtranetB2bMatchesCustomersPostRequest) GetCustomersOk() (*[]ManaV2ExtranetServiceProducerCustomer, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *V1ExtranetB2bMatchesCustomersPostRequest) SetCustomers(v []ManaV2ExtranetServiceProducerCustomer)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *V1ExtranetB2bMatchesCustomersPostRequest) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetServiceId

`func (o *V1ExtranetB2bMatchesCustomersPostRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *V1ExtranetB2bMatchesCustomersPostRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *V1ExtranetB2bMatchesCustomersPostRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *V1ExtranetB2bMatchesCustomersPostRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


