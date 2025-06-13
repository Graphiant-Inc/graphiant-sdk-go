# V1GatewaysPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**V1GatewaysPutRequestDetails**](V1GatewaysPutRequestDetails.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1GatewaysPutRequest

`func NewV1GatewaysPutRequest() *V1GatewaysPutRequest`

NewV1GatewaysPutRequest instantiates a new V1GatewaysPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GatewaysPutRequestWithDefaults

`func NewV1GatewaysPutRequestWithDefaults() *V1GatewaysPutRequest`

NewV1GatewaysPutRequestWithDefaults instantiates a new V1GatewaysPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *V1GatewaysPutRequest) GetDetails() V1GatewaysPutRequestDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *V1GatewaysPutRequest) GetDetailsOk() (*V1GatewaysPutRequestDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *V1GatewaysPutRequest) SetDetails(v V1GatewaysPutRequestDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *V1GatewaysPutRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *V1GatewaysPutRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1GatewaysPutRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1GatewaysPutRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1GatewaysPutRequest) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


