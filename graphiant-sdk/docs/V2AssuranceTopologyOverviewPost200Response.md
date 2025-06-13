# V2AssuranceTopologyOverviewPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumApplications** | Pointer to **int32** |  | [optional] 
**NumFlows** | Pointer to **int32** |  | [optional] 
**Topology** | Pointer to [**V2AssuranceTopologyOverviewPost200ResponseTopology**](V2AssuranceTopologyOverviewPost200ResponseTopology.md) |  | [optional] 
**TopologyChangeTs** | Pointer to [**[]V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**TrafficRegions** | Pointer to [**[]V2AssuranceTopologyOverviewPost200ResponseTrafficRegionsInner**](V2AssuranceTopologyOverviewPost200ResponseTrafficRegionsInner.md) |  | [optional] 

## Methods

### NewV2AssuranceTopologyOverviewPost200Response

`func NewV2AssuranceTopologyOverviewPost200Response() *V2AssuranceTopologyOverviewPost200Response`

NewV2AssuranceTopologyOverviewPost200Response instantiates a new V2AssuranceTopologyOverviewPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceTopologyOverviewPost200ResponseWithDefaults

`func NewV2AssuranceTopologyOverviewPost200ResponseWithDefaults() *V2AssuranceTopologyOverviewPost200Response`

NewV2AssuranceTopologyOverviewPost200ResponseWithDefaults instantiates a new V2AssuranceTopologyOverviewPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumApplications

`func (o *V2AssuranceTopologyOverviewPost200Response) GetNumApplications() int32`

GetNumApplications returns the NumApplications field if non-nil, zero value otherwise.

### GetNumApplicationsOk

`func (o *V2AssuranceTopologyOverviewPost200Response) GetNumApplicationsOk() (*int32, bool)`

GetNumApplicationsOk returns a tuple with the NumApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumApplications

`func (o *V2AssuranceTopologyOverviewPost200Response) SetNumApplications(v int32)`

SetNumApplications sets NumApplications field to given value.

### HasNumApplications

`func (o *V2AssuranceTopologyOverviewPost200Response) HasNumApplications() bool`

HasNumApplications returns a boolean if a field has been set.

### GetNumFlows

`func (o *V2AssuranceTopologyOverviewPost200Response) GetNumFlows() int32`

GetNumFlows returns the NumFlows field if non-nil, zero value otherwise.

### GetNumFlowsOk

`func (o *V2AssuranceTopologyOverviewPost200Response) GetNumFlowsOk() (*int32, bool)`

GetNumFlowsOk returns a tuple with the NumFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFlows

`func (o *V2AssuranceTopologyOverviewPost200Response) SetNumFlows(v int32)`

SetNumFlows sets NumFlows field to given value.

### HasNumFlows

`func (o *V2AssuranceTopologyOverviewPost200Response) HasNumFlows() bool`

HasNumFlows returns a boolean if a field has been set.

### GetTopology

`func (o *V2AssuranceTopologyOverviewPost200Response) GetTopology() V2AssuranceTopologyOverviewPost200ResponseTopology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *V2AssuranceTopologyOverviewPost200Response) GetTopologyOk() (*V2AssuranceTopologyOverviewPost200ResponseTopology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *V2AssuranceTopologyOverviewPost200Response) SetTopology(v V2AssuranceTopologyOverviewPost200ResponseTopology)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *V2AssuranceTopologyOverviewPost200Response) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetTopologyChangeTs

`func (o *V2AssuranceTopologyOverviewPost200Response) GetTopologyChangeTs() []V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetTopologyChangeTs returns the TopologyChangeTs field if non-nil, zero value otherwise.

### GetTopologyChangeTsOk

`func (o *V2AssuranceTopologyOverviewPost200Response) GetTopologyChangeTsOk() (*[]V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetTopologyChangeTsOk returns a tuple with the TopologyChangeTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyChangeTs

`func (o *V2AssuranceTopologyOverviewPost200Response) SetTopologyChangeTs(v []V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetTopologyChangeTs sets TopologyChangeTs field to given value.

### HasTopologyChangeTs

`func (o *V2AssuranceTopologyOverviewPost200Response) HasTopologyChangeTs() bool`

HasTopologyChangeTs returns a boolean if a field has been set.

### GetTrafficRegions

`func (o *V2AssuranceTopologyOverviewPost200Response) GetTrafficRegions() []V2AssuranceTopologyOverviewPost200ResponseTrafficRegionsInner`

GetTrafficRegions returns the TrafficRegions field if non-nil, zero value otherwise.

### GetTrafficRegionsOk

`func (o *V2AssuranceTopologyOverviewPost200Response) GetTrafficRegionsOk() (*[]V2AssuranceTopologyOverviewPost200ResponseTrafficRegionsInner, bool)`

GetTrafficRegionsOk returns a tuple with the TrafficRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRegions

`func (o *V2AssuranceTopologyOverviewPost200Response) SetTrafficRegions(v []V2AssuranceTopologyOverviewPost200ResponseTrafficRegionsInner)`

SetTrafficRegions sets TrafficRegions field to given value.

### HasTrafficRegions

`func (o *V2AssuranceTopologyOverviewPost200Response) HasTrafficRegions() bool`

HasTrafficRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


