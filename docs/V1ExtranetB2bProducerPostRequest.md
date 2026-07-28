# V1ExtranetB2bProducerPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**ManaV2ExtranetServiceProducerPolicy**](ManaV2ExtranetServiceProducerPolicy.md) |  | 
**ServiceName** | **string** |  (required) | 
**ServiceType** | **string** | Branded extranet service type (peering_service, client_to_server, …) (required) | 

## Methods

### NewV1ExtranetB2bProducerPostRequest

`func NewV1ExtranetB2bProducerPostRequest(policy ManaV2ExtranetServiceProducerPolicy, serviceName string, serviceType string, ) *V1ExtranetB2bProducerPostRequest`

NewV1ExtranetB2bProducerPostRequest instantiates a new V1ExtranetB2bProducerPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetB2bProducerPostRequestWithDefaults

`func NewV1ExtranetB2bProducerPostRequestWithDefaults() *V1ExtranetB2bProducerPostRequest`

NewV1ExtranetB2bProducerPostRequestWithDefaults instantiates a new V1ExtranetB2bProducerPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *V1ExtranetB2bProducerPostRequest) GetPolicy() ManaV2ExtranetServiceProducerPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *V1ExtranetB2bProducerPostRequest) GetPolicyOk() (*ManaV2ExtranetServiceProducerPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *V1ExtranetB2bProducerPostRequest) SetPolicy(v ManaV2ExtranetServiceProducerPolicy)`

SetPolicy sets Policy field to given value.


### GetServiceName

`func (o *V1ExtranetB2bProducerPostRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1ExtranetB2bProducerPostRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1ExtranetB2bProducerPostRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServiceType

`func (o *V1ExtranetB2bProducerPostRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *V1ExtranetB2bProducerPostRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *V1ExtranetB2bProducerPostRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


