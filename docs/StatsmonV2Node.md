# StatsmonV2Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitInfo** | Pointer to [**[]StatsmonV2NodeCircuitInfo**](StatsmonV2NodeCircuitInfo.md) |  | [optional] 
**Connections** | Pointer to [**[]StatsmonV2NodeConnection**](StatsmonV2NodeConnection.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeInfo** | Pointer to [**StatsmonV2NodeDeviceInfo**](StatsmonV2NodeDeviceInfo.md) |  | [optional] 
**PreferredRegionName** | Pointer to **string** |  | [optional] 
**Quality** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewStatsmonV2Node

`func NewStatsmonV2Node() *StatsmonV2Node`

NewStatsmonV2Node instantiates a new StatsmonV2Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsmonV2NodeWithDefaults

`func NewStatsmonV2NodeWithDefaults() *StatsmonV2Node`

NewStatsmonV2NodeWithDefaults instantiates a new StatsmonV2Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitInfo

`func (o *StatsmonV2Node) GetCircuitInfo() []StatsmonV2NodeCircuitInfo`

GetCircuitInfo returns the CircuitInfo field if non-nil, zero value otherwise.

### GetCircuitInfoOk

`func (o *StatsmonV2Node) GetCircuitInfoOk() (*[]StatsmonV2NodeCircuitInfo, bool)`

GetCircuitInfoOk returns a tuple with the CircuitInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitInfo

`func (o *StatsmonV2Node) SetCircuitInfo(v []StatsmonV2NodeCircuitInfo)`

SetCircuitInfo sets CircuitInfo field to given value.

### HasCircuitInfo

`func (o *StatsmonV2Node) HasCircuitInfo() bool`

HasCircuitInfo returns a boolean if a field has been set.

### GetConnections

`func (o *StatsmonV2Node) GetConnections() []StatsmonV2NodeConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *StatsmonV2Node) GetConnectionsOk() (*[]StatsmonV2NodeConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *StatsmonV2Node) SetConnections(v []StatsmonV2NodeConnection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *StatsmonV2Node) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetId

`func (o *StatsmonV2Node) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatsmonV2Node) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatsmonV2Node) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StatsmonV2Node) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StatsmonV2Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatsmonV2Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatsmonV2Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatsmonV2Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeInfo

`func (o *StatsmonV2Node) GetNodeInfo() StatsmonV2NodeDeviceInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *StatsmonV2Node) GetNodeInfoOk() (*StatsmonV2NodeDeviceInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *StatsmonV2Node) SetNodeInfo(v StatsmonV2NodeDeviceInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *StatsmonV2Node) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetPreferredRegionName

`func (o *StatsmonV2Node) GetPreferredRegionName() string`

GetPreferredRegionName returns the PreferredRegionName field if non-nil, zero value otherwise.

### GetPreferredRegionNameOk

`func (o *StatsmonV2Node) GetPreferredRegionNameOk() (*string, bool)`

GetPreferredRegionNameOk returns a tuple with the PreferredRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredRegionName

`func (o *StatsmonV2Node) SetPreferredRegionName(v string)`

SetPreferredRegionName sets PreferredRegionName field to given value.

### HasPreferredRegionName

`func (o *StatsmonV2Node) HasPreferredRegionName() bool`

HasPreferredRegionName returns a boolean if a field has been set.

### GetQuality

`func (o *StatsmonV2Node) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *StatsmonV2Node) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *StatsmonV2Node) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *StatsmonV2Node) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetType

`func (o *StatsmonV2Node) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatsmonV2Node) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatsmonV2Node) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StatsmonV2Node) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


