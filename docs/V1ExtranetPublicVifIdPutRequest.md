# V1ExtranetPublicVifIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerPolicy** | Pointer to [**ManaV2PublicVifConsumerPolicy**](ManaV2PublicVifConsumerPolicy.md) |  | [optional] 
**ProducerPolicy** | [**ManaV2PublicVifProducerPolicy**](ManaV2PublicVifProducerPolicy.md) |  | 

## Methods

### NewV1ExtranetPublicVifIdPutRequest

`func NewV1ExtranetPublicVifIdPutRequest(producerPolicy ManaV2PublicVifProducerPolicy, ) *V1ExtranetPublicVifIdPutRequest`

NewV1ExtranetPublicVifIdPutRequest instantiates a new V1ExtranetPublicVifIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetPublicVifIdPutRequestWithDefaults

`func NewV1ExtranetPublicVifIdPutRequestWithDefaults() *V1ExtranetPublicVifIdPutRequest`

NewV1ExtranetPublicVifIdPutRequestWithDefaults instantiates a new V1ExtranetPublicVifIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerPolicy

`func (o *V1ExtranetPublicVifIdPutRequest) GetConsumerPolicy() ManaV2PublicVifConsumerPolicy`

GetConsumerPolicy returns the ConsumerPolicy field if non-nil, zero value otherwise.

### GetConsumerPolicyOk

`func (o *V1ExtranetPublicVifIdPutRequest) GetConsumerPolicyOk() (*ManaV2PublicVifConsumerPolicy, bool)`

GetConsumerPolicyOk returns a tuple with the ConsumerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerPolicy

`func (o *V1ExtranetPublicVifIdPutRequest) SetConsumerPolicy(v ManaV2PublicVifConsumerPolicy)`

SetConsumerPolicy sets ConsumerPolicy field to given value.

### HasConsumerPolicy

`func (o *V1ExtranetPublicVifIdPutRequest) HasConsumerPolicy() bool`

HasConsumerPolicy returns a boolean if a field has been set.

### GetProducerPolicy

`func (o *V1ExtranetPublicVifIdPutRequest) GetProducerPolicy() ManaV2PublicVifProducerPolicy`

GetProducerPolicy returns the ProducerPolicy field if non-nil, zero value otherwise.

### GetProducerPolicyOk

`func (o *V1ExtranetPublicVifIdPutRequest) GetProducerPolicyOk() (*ManaV2PublicVifProducerPolicy, bool)`

GetProducerPolicyOk returns a tuple with the ProducerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerPolicy

`func (o *V1ExtranetPublicVifIdPutRequest) SetProducerPolicy(v ManaV2PublicVifProducerPolicy)`

SetProducerPolicy sets ProducerPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


