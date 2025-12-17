# V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | the id associated with an entity - consumer_id for consumer, and service_id for the producer/service (required) | 
**IsB2B** | **bool** | whether the entity is a b2b entity (true for b2b entity, false for local extranet entity) (required) | 
**IsProvider** | **bool** | whether the entity is a provider or consumer (required) | 
**SiteId** | Pointer to **int64** | a filter to get usage for a specific site | [optional] 
**SubscriptionName** | Pointer to **string** | Optional subscription name for filter | [optional] 
**TimeWindow** | [**StatsmonTimeWindow**](StatsmonTimeWindow.md) |  | 
**VrfId** | Pointer to **int64** | a filter to get usage for a specific vrf | [optional] 

## Methods

### NewV1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest

`func NewV1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest(id int64, isB2B bool, isProvider bool, timeWindow StatsmonTimeWindow, ) *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest`

NewV1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest instantiates a new V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequestWithDefaults

`func NewV1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequestWithDefaults() *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest`

NewV1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequestWithDefaults instantiates a new V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetIsB2B

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetIsB2B() bool`

GetIsB2B returns the IsB2B field if non-nil, zero value otherwise.

### GetIsB2BOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetIsB2BOk() (*bool, bool)`

GetIsB2BOk returns a tuple with the IsB2B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsB2B

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) SetIsB2B(v bool)`

SetIsB2B sets IsB2B field to given value.


### GetIsProvider

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetIsProvider() bool`

GetIsProvider returns the IsProvider field if non-nil, zero value otherwise.

### GetIsProviderOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetIsProviderOk() (*bool, bool)`

GetIsProviderOk returns a tuple with the IsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvider

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) SetIsProvider(v bool)`

SetIsProvider sets IsProvider field to given value.


### GetSiteId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSubscriptionName

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetTimeWindow() StatsmonTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetTimeWindowOk() (*StatsmonTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) SetTimeWindow(v StatsmonTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.


### GetVrfId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetVrfId() int64`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) GetVrfIdOk() (*int64, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) SetVrfId(v int64)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *V1ExtranetB2bMonitoringPeeringServiceServiceOvertimeConsumptionPostRequest) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


