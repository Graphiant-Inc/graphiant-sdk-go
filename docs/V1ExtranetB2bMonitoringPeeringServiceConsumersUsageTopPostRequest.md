# V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | the id of a producer/service (required) | 
**IsB2B** | **bool** | whether the entity is a b2b entity (true for b2b entity, false for local extranet entity) (required) | 
**IsProvider** | **bool** | whether the entity is a provider or consumer (required) | 
**TimeWindow** | [**StatsmonTimeWindow**](StatsmonTimeWindow.md) |  | 

## Methods

### NewV1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest

`func NewV1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest(id int64, isB2B bool, isProvider bool, timeWindow StatsmonTimeWindow, ) *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest`

NewV1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest instantiates a new V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequestWithDefaults

`func NewV1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequestWithDefaults() *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest`

NewV1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequestWithDefaults instantiates a new V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetIsB2B

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) GetIsB2B() bool`

GetIsB2B returns the IsB2B field if non-nil, zero value otherwise.

### GetIsB2BOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) GetIsB2BOk() (*bool, bool)`

GetIsB2BOk returns a tuple with the IsB2B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsB2B

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) SetIsB2B(v bool)`

SetIsB2B sets IsB2B field to given value.


### GetIsProvider

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) GetIsProvider() bool`

GetIsProvider returns the IsProvider field if non-nil, zero value otherwise.

### GetIsProviderOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) GetIsProviderOk() (*bool, bool)`

GetIsProviderOk returns a tuple with the IsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvider

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) SetIsProvider(v bool)`

SetIsProvider sets IsProvider field to given value.


### GetTimeWindow

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) GetTimeWindow() StatsmonTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) GetTimeWindowOk() (*StatsmonTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V1ExtranetB2bMonitoringPeeringServiceConsumersUsageTopPostRequest) SetTimeWindow(v StatsmonTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


