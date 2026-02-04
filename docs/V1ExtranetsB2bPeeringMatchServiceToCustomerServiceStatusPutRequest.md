# V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | match ID for the customer&#39;s service subscription (required) | 
**Status** | **string** | Customer’s service status: Paused or Active (required) | 

## Methods

### NewV1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest

`func NewV1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest(id int64, status string, ) *V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest`

NewV1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest instantiates a new V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequestWithDefaults

`func NewV1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequestWithDefaults() *V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest`

NewV1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequestWithDefaults instantiates a new V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetStatus

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ExtranetsB2bPeeringMatchServiceToCustomerServiceStatusPutRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


