# V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | the id of the service (required) | 
**IsProvider** | **bool** | whether the entity is a provider/producer or consumer (required) | 

## Methods

### NewV1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest

`func NewV1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest(id int64, isProvider bool, ) *V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest`

NewV1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest instantiates a new V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequestWithDefaults

`func NewV1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequestWithDefaults() *V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest`

NewV1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequestWithDefaults instantiates a new V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetIsProvider

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest) GetIsProvider() bool`

GetIsProvider returns the IsProvider field if non-nil, zero value otherwise.

### GetIsProviderOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest) GetIsProviderOk() (*bool, bool)`

GetIsProviderOk returns a tuple with the IsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvider

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceHealthPostRequest) SetIsProvider(v bool)`

SetIsProvider sets IsProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


