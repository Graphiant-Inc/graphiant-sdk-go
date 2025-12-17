# V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the customer (required) | 
**Service** | [**ManaV2B2bExtranetMatchServiceToCustomer**](ManaV2B2bExtranetMatchServiceToCustomer.md) |  | 

## Methods

### NewV1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest

`func NewV1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest(id int64, service ManaV2B2bExtranetMatchServiceToCustomer, ) *V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest`

NewV1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest instantiates a new V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsB2bPeeringMatchServiceToCustomerPostRequestWithDefaults

`func NewV1ExtranetsB2bPeeringMatchServiceToCustomerPostRequestWithDefaults() *V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest`

NewV1ExtranetsB2bPeeringMatchServiceToCustomerPostRequestWithDefaults instantiates a new V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetService

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest) GetService() ManaV2B2bExtranetMatchServiceToCustomer`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest) GetServiceOk() (*ManaV2B2bExtranetMatchServiceToCustomer, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerPostRequest) SetService(v ManaV2B2bExtranetMatchServiceToCustomer)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


