# V1ExtranetsB2bPeeringProducerPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**ManaV2B2bExtranetPeeringServiceProducerPolicy**](ManaV2B2bExtranetPeeringServiceProducerPolicy.md) |  | 
**ServiceName** | Pointer to **string** |  | [optional] 
**Type** | **string** | Type of the service whether it is application or peering (required) | 

## Methods

### NewV1ExtranetsB2bPeeringProducerPostRequest

`func NewV1ExtranetsB2bPeeringProducerPostRequest(policy ManaV2B2bExtranetPeeringServiceProducerPolicy, type_ string, ) *V1ExtranetsB2bPeeringProducerPostRequest`

NewV1ExtranetsB2bPeeringProducerPostRequest instantiates a new V1ExtranetsB2bPeeringProducerPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsB2bPeeringProducerPostRequestWithDefaults

`func NewV1ExtranetsB2bPeeringProducerPostRequestWithDefaults() *V1ExtranetsB2bPeeringProducerPostRequest`

NewV1ExtranetsB2bPeeringProducerPostRequestWithDefaults instantiates a new V1ExtranetsB2bPeeringProducerPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *V1ExtranetsB2bPeeringProducerPostRequest) GetPolicy() ManaV2B2bExtranetPeeringServiceProducerPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *V1ExtranetsB2bPeeringProducerPostRequest) GetPolicyOk() (*ManaV2B2bExtranetPeeringServiceProducerPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *V1ExtranetsB2bPeeringProducerPostRequest) SetPolicy(v ManaV2B2bExtranetPeeringServiceProducerPolicy)`

SetPolicy sets Policy field to given value.


### GetServiceName

`func (o *V1ExtranetsB2bPeeringProducerPostRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1ExtranetsB2bPeeringProducerPostRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1ExtranetsB2bPeeringProducerPostRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *V1ExtranetsB2bPeeringProducerPostRequest) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetType

`func (o *V1ExtranetsB2bPeeringProducerPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ExtranetsB2bPeeringProducerPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ExtranetsB2bPeeringProducerPostRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


