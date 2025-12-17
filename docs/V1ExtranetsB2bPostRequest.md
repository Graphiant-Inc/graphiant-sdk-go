# V1ExtranetsB2bPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**ManaV2B2bExtranetProducerPolicy**](ManaV2B2bExtranetProducerPolicy.md) |  | 
**ServiceName** | **string** | Name of the service (required) | 
**Type** | **string** | Type of the service whether it is application or peering (required) | 

## Methods

### NewV1ExtranetsB2bPostRequest

`func NewV1ExtranetsB2bPostRequest(policy ManaV2B2bExtranetProducerPolicy, serviceName string, type_ string, ) *V1ExtranetsB2bPostRequest`

NewV1ExtranetsB2bPostRequest instantiates a new V1ExtranetsB2bPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsB2bPostRequestWithDefaults

`func NewV1ExtranetsB2bPostRequestWithDefaults() *V1ExtranetsB2bPostRequest`

NewV1ExtranetsB2bPostRequestWithDefaults instantiates a new V1ExtranetsB2bPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *V1ExtranetsB2bPostRequest) GetPolicy() ManaV2B2bExtranetProducerPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *V1ExtranetsB2bPostRequest) GetPolicyOk() (*ManaV2B2bExtranetProducerPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *V1ExtranetsB2bPostRequest) SetPolicy(v ManaV2B2bExtranetProducerPolicy)`

SetPolicy sets Policy field to given value.


### GetServiceName

`func (o *V1ExtranetsB2bPostRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1ExtranetsB2bPostRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1ExtranetsB2bPostRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetType

`func (o *V1ExtranetsB2bPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ExtranetsB2bPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ExtranetsB2bPostRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


