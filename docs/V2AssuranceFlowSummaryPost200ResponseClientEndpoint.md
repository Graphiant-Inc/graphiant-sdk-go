# V2AssuranceFlowSummaryPost200ResponseClientEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Circuits** | Pointer to **[]string** |  | [optional] 
**Edges** | Pointer to [**[]V2AssuranceFlowSummaryPost200ResponseClientEndpointEdgesInner**](V2AssuranceFlowSummaryPost200ResponseClientEndpointEdgesInner.md) |  | [optional] 
**Jitter** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter**](V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter.md) |  | [optional] 
**Latency** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter**](V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter.md) |  | [optional] 
**Loss** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter**](V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter.md) |  | [optional] 
**Site** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpointSite**](V2AssuranceFlowSummaryPost200ResponseClientEndpointSite.md) |  | [optional] 
**TotalDownlinkUsage** | Pointer to **int64** |  | [optional] 
**TotalUplinkUsage** | Pointer to **int64** |  | [optional] 

## Methods

### NewV2AssuranceFlowSummaryPost200ResponseClientEndpoint

`func NewV2AssuranceFlowSummaryPost200ResponseClientEndpoint() *V2AssuranceFlowSummaryPost200ResponseClientEndpoint`

NewV2AssuranceFlowSummaryPost200ResponseClientEndpoint instantiates a new V2AssuranceFlowSummaryPost200ResponseClientEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceFlowSummaryPost200ResponseClientEndpointWithDefaults

`func NewV2AssuranceFlowSummaryPost200ResponseClientEndpointWithDefaults() *V2AssuranceFlowSummaryPost200ResponseClientEndpoint`

NewV2AssuranceFlowSummaryPost200ResponseClientEndpointWithDefaults instantiates a new V2AssuranceFlowSummaryPost200ResponseClientEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuits

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetCircuits() []string`

GetCircuits returns the Circuits field if non-nil, zero value otherwise.

### GetCircuitsOk

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetCircuitsOk() (*[]string, bool)`

GetCircuitsOk returns a tuple with the Circuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuits

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) SetCircuits(v []string)`

SetCircuits sets Circuits field to given value.

### HasCircuits

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) HasCircuits() bool`

HasCircuits returns a boolean if a field has been set.

### GetEdges

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetEdges() []V2AssuranceFlowSummaryPost200ResponseClientEndpointEdgesInner`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetEdgesOk() (*[]V2AssuranceFlowSummaryPost200ResponseClientEndpointEdgesInner, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) SetEdges(v []V2AssuranceFlowSummaryPost200ResponseClientEndpointEdgesInner)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetJitter

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetJitter() V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetJitterOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) SetJitter(v V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### GetLatency

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetLatency() V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetLatencyOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) SetLatency(v V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLoss

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetLoss() V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetLossOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) SetLoss(v V2AssuranceFlowSummaryPost200ResponseClientEndpointJitter)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetSite

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetSite() V2AssuranceFlowSummaryPost200ResponseClientEndpointSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetSiteOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpointSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) SetSite(v V2AssuranceFlowSummaryPost200ResponseClientEndpointSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetTotalDownlinkUsage

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetTotalDownlinkUsage() int64`

GetTotalDownlinkUsage returns the TotalDownlinkUsage field if non-nil, zero value otherwise.

### GetTotalDownlinkUsageOk

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetTotalDownlinkUsageOk() (*int64, bool)`

GetTotalDownlinkUsageOk returns a tuple with the TotalDownlinkUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDownlinkUsage

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) SetTotalDownlinkUsage(v int64)`

SetTotalDownlinkUsage sets TotalDownlinkUsage field to given value.

### HasTotalDownlinkUsage

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) HasTotalDownlinkUsage() bool`

HasTotalDownlinkUsage returns a boolean if a field has been set.

### GetTotalUplinkUsage

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetTotalUplinkUsage() int64`

GetTotalUplinkUsage returns the TotalUplinkUsage field if non-nil, zero value otherwise.

### GetTotalUplinkUsageOk

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) GetTotalUplinkUsageOk() (*int64, bool)`

GetTotalUplinkUsageOk returns a tuple with the TotalUplinkUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUplinkUsage

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) SetTotalUplinkUsage(v int64)`

SetTotalUplinkUsage sets TotalUplinkUsage field to given value.

### HasTotalUplinkUsage

`func (o *V2AssuranceFlowSummaryPost200ResponseClientEndpoint) HasTotalUplinkUsage() bool`

HasTotalUplinkUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


