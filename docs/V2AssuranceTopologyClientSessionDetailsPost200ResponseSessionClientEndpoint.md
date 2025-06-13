# V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Circuits** | Pointer to **[]string** |  | [optional] 
**Edges** | Pointer to [**[]V2AssuranceFlowSummaryPost200ResponseClientEndpointEdgesInner**](V2AssuranceFlowSummaryPost200ResponseClientEndpointEdgesInner.md) |  | [optional] 
**IsGateway** | Pointer to **bool** |  | [optional] 
**Jitter** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter**](V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter.md) |  | [optional] 
**Latency** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter**](V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter.md) |  | [optional] 
**Loss** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter**](V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter.md) |  | [optional] 
**Site** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpointSite**](V2AssuranceFlowSummaryPost200ResponseClientEndpointSite.md) |  | [optional] 
**TotalDownlinkUsage** | Pointer to **int64** |  | [optional] 
**TotalUplinkUsage** | Pointer to **int64** |  | [optional] 

## Methods

### NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint

`func NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint() *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint`

NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint instantiates a new V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpointWithDefaults

`func NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpointWithDefaults() *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint`

NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpointWithDefaults instantiates a new V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuits

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetCircuits() []string`

GetCircuits returns the Circuits field if non-nil, zero value otherwise.

### GetCircuitsOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetCircuitsOk() (*[]string, bool)`

GetCircuitsOk returns a tuple with the Circuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuits

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) SetCircuits(v []string)`

SetCircuits sets Circuits field to given value.

### HasCircuits

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) HasCircuits() bool`

HasCircuits returns a boolean if a field has been set.

### GetEdges

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetEdges() []V2AssuranceFlowSummaryPost200ResponseClientEndpointEdgesInner`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetEdgesOk() (*[]V2AssuranceFlowSummaryPost200ResponseClientEndpointEdgesInner, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) SetEdges(v []V2AssuranceFlowSummaryPost200ResponseClientEndpointEdgesInner)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetIsGateway

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetIsGateway() bool`

GetIsGateway returns the IsGateway field if non-nil, zero value otherwise.

### GetIsGatewayOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetIsGatewayOk() (*bool, bool)`

GetIsGatewayOk returns a tuple with the IsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGateway

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) SetIsGateway(v bool)`

SetIsGateway sets IsGateway field to given value.

### HasIsGateway

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) HasIsGateway() bool`

HasIsGateway returns a boolean if a field has been set.

### GetJitter

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetJitter() V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetJitterOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) SetJitter(v V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### GetLatency

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetLatency() V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetLatencyOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) SetLatency(v V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLoss

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetLoss() V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetLossOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) SetLoss(v V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetSite

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetSite() V2AssuranceFlowSummaryPost200ResponseClientEndpointSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetSiteOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpointSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) SetSite(v V2AssuranceFlowSummaryPost200ResponseClientEndpointSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetTotalDownlinkUsage

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetTotalDownlinkUsage() int64`

GetTotalDownlinkUsage returns the TotalDownlinkUsage field if non-nil, zero value otherwise.

### GetTotalDownlinkUsageOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetTotalDownlinkUsageOk() (*int64, bool)`

GetTotalDownlinkUsageOk returns a tuple with the TotalDownlinkUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDownlinkUsage

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) SetTotalDownlinkUsage(v int64)`

SetTotalDownlinkUsage sets TotalDownlinkUsage field to given value.

### HasTotalDownlinkUsage

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) HasTotalDownlinkUsage() bool`

HasTotalDownlinkUsage returns a boolean if a field has been set.

### GetTotalUplinkUsage

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetTotalUplinkUsage() int64`

GetTotalUplinkUsage returns the TotalUplinkUsage field if non-nil, zero value otherwise.

### GetTotalUplinkUsageOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) GetTotalUplinkUsageOk() (*int64, bool)`

GetTotalUplinkUsageOk returns a tuple with the TotalUplinkUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUplinkUsage

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) SetTotalUplinkUsage(v int64)`

SetTotalUplinkUsage sets TotalUplinkUsage field to given value.

### HasTotalUplinkUsage

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint) HasTotalUplinkUsage() bool`

HasTotalUplinkUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


