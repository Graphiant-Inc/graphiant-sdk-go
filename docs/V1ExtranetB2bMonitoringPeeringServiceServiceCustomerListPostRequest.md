# V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | service id (required) | 
**IsProvider** | **bool** | whether the entity is a provider or consumer (required) | 
**TimeWindow** | [**StatsmonTimeWindow**](StatsmonTimeWindow.md) |  | 

## Methods

### NewV1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest

`func NewV1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest(id int64, isProvider bool, timeWindow StatsmonTimeWindow, ) *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest`

NewV1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest instantiates a new V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequestWithDefaults

`func NewV1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequestWithDefaults() *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest`

NewV1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequestWithDefaults instantiates a new V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetIsProvider

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest) GetIsProvider() bool`

GetIsProvider returns the IsProvider field if non-nil, zero value otherwise.

### GetIsProviderOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest) GetIsProviderOk() (*bool, bool)`

GetIsProviderOk returns a tuple with the IsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvider

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest) SetIsProvider(v bool)`

SetIsProvider sets IsProvider field to given value.


### GetTimeWindow

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest) GetTimeWindow() StatsmonTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest) GetTimeWindowOk() (*StatsmonTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceCustomerListPostRequest) SetTimeWindow(v StatsmonTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


