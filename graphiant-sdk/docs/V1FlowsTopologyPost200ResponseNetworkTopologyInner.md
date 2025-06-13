# V1FlowsTopologyPost200ResponseNetworkTopologyInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitStatus** | Pointer to **map[string]int32** |  | [optional] 
**Delta** | Pointer to [**V1FlowsTopologyPost200ResponseNetworkTopologyInnerDelta**](V1FlowsTopologyPost200ResponseNetworkTopologyInnerDelta.md) |  | [optional] 
**Edges** | Pointer to [**[]V1FlowsTopologyPost200ResponseNetworkTopologyInnerDeltaEdgesAddedInner**](V1FlowsTopologyPost200ResponseNetworkTopologyInnerDeltaEdgesAddedInner.md) |  | [optional] 
**Flows** | Pointer to **int32** |  | [optional] 
**Nodes** | Pointer to [**[]V1FlowsTopologyPost200ResponseNetworkTopologyInnerDeltaNodesAddedInner**](V1FlowsTopologyPost200ResponseNetworkTopologyInnerDeltaNodesAddedInner.md) |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV1FlowsTopologyPost200ResponseNetworkTopologyInner

`func NewV1FlowsTopologyPost200ResponseNetworkTopologyInner() *V1FlowsTopologyPost200ResponseNetworkTopologyInner`

NewV1FlowsTopologyPost200ResponseNetworkTopologyInner instantiates a new V1FlowsTopologyPost200ResponseNetworkTopologyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FlowsTopologyPost200ResponseNetworkTopologyInnerWithDefaults

`func NewV1FlowsTopologyPost200ResponseNetworkTopologyInnerWithDefaults() *V1FlowsTopologyPost200ResponseNetworkTopologyInner`

NewV1FlowsTopologyPost200ResponseNetworkTopologyInnerWithDefaults instantiates a new V1FlowsTopologyPost200ResponseNetworkTopologyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitStatus

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetCircuitStatus() map[string]int32`

GetCircuitStatus returns the CircuitStatus field if non-nil, zero value otherwise.

### GetCircuitStatusOk

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetCircuitStatusOk() (*map[string]int32, bool)`

GetCircuitStatusOk returns a tuple with the CircuitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitStatus

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) SetCircuitStatus(v map[string]int32)`

SetCircuitStatus sets CircuitStatus field to given value.

### HasCircuitStatus

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) HasCircuitStatus() bool`

HasCircuitStatus returns a boolean if a field has been set.

### GetDelta

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetDelta() V1FlowsTopologyPost200ResponseNetworkTopologyInnerDelta`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetDeltaOk() (*V1FlowsTopologyPost200ResponseNetworkTopologyInnerDelta, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) SetDelta(v V1FlowsTopologyPost200ResponseNetworkTopologyInnerDelta)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetEdges

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetEdges() []V1FlowsTopologyPost200ResponseNetworkTopologyInnerDeltaEdgesAddedInner`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetEdgesOk() (*[]V1FlowsTopologyPost200ResponseNetworkTopologyInnerDeltaEdgesAddedInner, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) SetEdges(v []V1FlowsTopologyPost200ResponseNetworkTopologyInnerDeltaEdgesAddedInner)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetFlows

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetFlows() int32`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetFlowsOk() (*int32, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) SetFlows(v int32)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) HasFlows() bool`

HasFlows returns a boolean if a field has been set.

### GetNodes

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetNodes() []V1FlowsTopologyPost200ResponseNetworkTopologyInnerDeltaNodesAddedInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetNodesOk() (*[]V1FlowsTopologyPost200ResponseNetworkTopologyInnerDeltaNodesAddedInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) SetNodes(v []V1FlowsTopologyPost200ResponseNetworkTopologyInnerDeltaNodesAddedInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V1FlowsTopologyPost200ResponseNetworkTopologyInner) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


