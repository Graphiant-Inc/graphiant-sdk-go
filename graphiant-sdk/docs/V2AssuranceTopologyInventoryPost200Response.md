# V2AssuranceTopologyInventoryPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppNames** | Pointer to **[]string** |  | [optional] 
**ClientSites** | Pointer to [**[]V2AssuranceFlowSummaryPost200ResponseClientEndpointSite**](V2AssuranceFlowSummaryPost200ResponseClientEndpointSite.md) |  | [optional] 
**Regions** | Pointer to [**[]V2AssuranceTopologyInventoryPost200ResponseRegionsInner**](V2AssuranceTopologyInventoryPost200ResponseRegionsInner.md) |  | [optional] 
**ServerSites** | Pointer to [**[]V2AssuranceFlowSummaryPost200ResponseClientEndpointSite**](V2AssuranceFlowSummaryPost200ResponseClientEndpointSite.md) |  | [optional] 
**TopologyChangeTs** | Pointer to [**[]V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV2AssuranceTopologyInventoryPost200Response

`func NewV2AssuranceTopologyInventoryPost200Response() *V2AssuranceTopologyInventoryPost200Response`

NewV2AssuranceTopologyInventoryPost200Response instantiates a new V2AssuranceTopologyInventoryPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceTopologyInventoryPost200ResponseWithDefaults

`func NewV2AssuranceTopologyInventoryPost200ResponseWithDefaults() *V2AssuranceTopologyInventoryPost200Response`

NewV2AssuranceTopologyInventoryPost200ResponseWithDefaults instantiates a new V2AssuranceTopologyInventoryPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppNames

`func (o *V2AssuranceTopologyInventoryPost200Response) GetAppNames() []string`

GetAppNames returns the AppNames field if non-nil, zero value otherwise.

### GetAppNamesOk

`func (o *V2AssuranceTopologyInventoryPost200Response) GetAppNamesOk() (*[]string, bool)`

GetAppNamesOk returns a tuple with the AppNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppNames

`func (o *V2AssuranceTopologyInventoryPost200Response) SetAppNames(v []string)`

SetAppNames sets AppNames field to given value.

### HasAppNames

`func (o *V2AssuranceTopologyInventoryPost200Response) HasAppNames() bool`

HasAppNames returns a boolean if a field has been set.

### GetClientSites

`func (o *V2AssuranceTopologyInventoryPost200Response) GetClientSites() []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite`

GetClientSites returns the ClientSites field if non-nil, zero value otherwise.

### GetClientSitesOk

`func (o *V2AssuranceTopologyInventoryPost200Response) GetClientSitesOk() (*[]V2AssuranceFlowSummaryPost200ResponseClientEndpointSite, bool)`

GetClientSitesOk returns a tuple with the ClientSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSites

`func (o *V2AssuranceTopologyInventoryPost200Response) SetClientSites(v []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite)`

SetClientSites sets ClientSites field to given value.

### HasClientSites

`func (o *V2AssuranceTopologyInventoryPost200Response) HasClientSites() bool`

HasClientSites returns a boolean if a field has been set.

### GetRegions

`func (o *V2AssuranceTopologyInventoryPost200Response) GetRegions() []V2AssuranceTopologyInventoryPost200ResponseRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *V2AssuranceTopologyInventoryPost200Response) GetRegionsOk() (*[]V2AssuranceTopologyInventoryPost200ResponseRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *V2AssuranceTopologyInventoryPost200Response) SetRegions(v []V2AssuranceTopologyInventoryPost200ResponseRegionsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *V2AssuranceTopologyInventoryPost200Response) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetServerSites

`func (o *V2AssuranceTopologyInventoryPost200Response) GetServerSites() []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite`

GetServerSites returns the ServerSites field if non-nil, zero value otherwise.

### GetServerSitesOk

`func (o *V2AssuranceTopologyInventoryPost200Response) GetServerSitesOk() (*[]V2AssuranceFlowSummaryPost200ResponseClientEndpointSite, bool)`

GetServerSitesOk returns a tuple with the ServerSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSites

`func (o *V2AssuranceTopologyInventoryPost200Response) SetServerSites(v []V2AssuranceFlowSummaryPost200ResponseClientEndpointSite)`

SetServerSites sets ServerSites field to given value.

### HasServerSites

`func (o *V2AssuranceTopologyInventoryPost200Response) HasServerSites() bool`

HasServerSites returns a boolean if a field has been set.

### GetTopologyChangeTs

`func (o *V2AssuranceTopologyInventoryPost200Response) GetTopologyChangeTs() []V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetTopologyChangeTs returns the TopologyChangeTs field if non-nil, zero value otherwise.

### GetTopologyChangeTsOk

`func (o *V2AssuranceTopologyInventoryPost200Response) GetTopologyChangeTsOk() (*[]V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetTopologyChangeTsOk returns a tuple with the TopologyChangeTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyChangeTs

`func (o *V2AssuranceTopologyInventoryPost200Response) SetTopologyChangeTs(v []V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetTopologyChangeTs sets TopologyChangeTs field to given value.

### HasTopologyChangeTs

`func (o *V2AssuranceTopologyInventoryPost200Response) HasTopologyChangeTs() bool`

HasTopologyChangeTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


