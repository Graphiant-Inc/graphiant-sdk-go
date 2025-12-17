# V2ExtranetTotalUsagePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | the ID associated with an entity - consumer_id for consumer, and service_id for the producer/service (required) | 
**IsB2B** | **bool** | whether the entity is a b2b entity (true for b2b entity, false for local extranet entity) (required) | 
**IsProvider** | **bool** | whether the entity is a provider or consumer (required) | 
**TimeWindow** | [**StatsmonTimeWindow**](StatsmonTimeWindow.md) |  | 

## Methods

### NewV2ExtranetTotalUsagePostRequest

`func NewV2ExtranetTotalUsagePostRequest(id int64, isB2B bool, isProvider bool, timeWindow StatsmonTimeWindow, ) *V2ExtranetTotalUsagePostRequest`

NewV2ExtranetTotalUsagePostRequest instantiates a new V2ExtranetTotalUsagePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ExtranetTotalUsagePostRequestWithDefaults

`func NewV2ExtranetTotalUsagePostRequestWithDefaults() *V2ExtranetTotalUsagePostRequest`

NewV2ExtranetTotalUsagePostRequestWithDefaults instantiates a new V2ExtranetTotalUsagePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2ExtranetTotalUsagePostRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2ExtranetTotalUsagePostRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2ExtranetTotalUsagePostRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetIsB2B

`func (o *V2ExtranetTotalUsagePostRequest) GetIsB2B() bool`

GetIsB2B returns the IsB2B field if non-nil, zero value otherwise.

### GetIsB2BOk

`func (o *V2ExtranetTotalUsagePostRequest) GetIsB2BOk() (*bool, bool)`

GetIsB2BOk returns a tuple with the IsB2B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsB2B

`func (o *V2ExtranetTotalUsagePostRequest) SetIsB2B(v bool)`

SetIsB2B sets IsB2B field to given value.


### GetIsProvider

`func (o *V2ExtranetTotalUsagePostRequest) GetIsProvider() bool`

GetIsProvider returns the IsProvider field if non-nil, zero value otherwise.

### GetIsProviderOk

`func (o *V2ExtranetTotalUsagePostRequest) GetIsProviderOk() (*bool, bool)`

GetIsProviderOk returns a tuple with the IsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvider

`func (o *V2ExtranetTotalUsagePostRequest) SetIsProvider(v bool)`

SetIsProvider sets IsProvider field to given value.


### GetTimeWindow

`func (o *V2ExtranetTotalUsagePostRequest) GetTimeWindow() StatsmonTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2ExtranetTotalUsagePostRequest) GetTimeWindowOk() (*StatsmonTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2ExtranetTotalUsagePostRequest) SetTimeWindow(v StatsmonTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


