# V1ExtranetPublicVifCheckPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerPolicy** | [**ManaV2PublicVifConsumerPolicy**](ManaV2PublicVifConsumerPolicy.md) |  | 
**ProducerPolicy** | [**ManaV2PublicVifProducerPolicy**](ManaV2PublicVifProducerPolicy.md) |  | 
**ServiceName** | Pointer to **string** |  | [optional] 
**Type** | **string** | Type of the service whether it is application or peering (required) | 

## Methods

### NewV1ExtranetPublicVifCheckPostRequest

`func NewV1ExtranetPublicVifCheckPostRequest(consumerPolicy ManaV2PublicVifConsumerPolicy, producerPolicy ManaV2PublicVifProducerPolicy, type_ string, ) *V1ExtranetPublicVifCheckPostRequest`

NewV1ExtranetPublicVifCheckPostRequest instantiates a new V1ExtranetPublicVifCheckPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetPublicVifCheckPostRequestWithDefaults

`func NewV1ExtranetPublicVifCheckPostRequestWithDefaults() *V1ExtranetPublicVifCheckPostRequest`

NewV1ExtranetPublicVifCheckPostRequestWithDefaults instantiates a new V1ExtranetPublicVifCheckPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerPolicy

`func (o *V1ExtranetPublicVifCheckPostRequest) GetConsumerPolicy() ManaV2PublicVifConsumerPolicy`

GetConsumerPolicy returns the ConsumerPolicy field if non-nil, zero value otherwise.

### GetConsumerPolicyOk

`func (o *V1ExtranetPublicVifCheckPostRequest) GetConsumerPolicyOk() (*ManaV2PublicVifConsumerPolicy, bool)`

GetConsumerPolicyOk returns a tuple with the ConsumerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerPolicy

`func (o *V1ExtranetPublicVifCheckPostRequest) SetConsumerPolicy(v ManaV2PublicVifConsumerPolicy)`

SetConsumerPolicy sets ConsumerPolicy field to given value.


### GetProducerPolicy

`func (o *V1ExtranetPublicVifCheckPostRequest) GetProducerPolicy() ManaV2PublicVifProducerPolicy`

GetProducerPolicy returns the ProducerPolicy field if non-nil, zero value otherwise.

### GetProducerPolicyOk

`func (o *V1ExtranetPublicVifCheckPostRequest) GetProducerPolicyOk() (*ManaV2PublicVifProducerPolicy, bool)`

GetProducerPolicyOk returns a tuple with the ProducerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerPolicy

`func (o *V1ExtranetPublicVifCheckPostRequest) SetProducerPolicy(v ManaV2PublicVifProducerPolicy)`

SetProducerPolicy sets ProducerPolicy field to given value.


### GetServiceName

`func (o *V1ExtranetPublicVifCheckPostRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1ExtranetPublicVifCheckPostRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1ExtranetPublicVifCheckPostRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *V1ExtranetPublicVifCheckPostRequest) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetType

`func (o *V1ExtranetPublicVifCheckPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ExtranetPublicVifCheckPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ExtranetPublicVifCheckPostRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


