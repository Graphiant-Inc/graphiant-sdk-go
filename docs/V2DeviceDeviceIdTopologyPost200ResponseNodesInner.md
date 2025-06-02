# V2DeviceDeviceIdTopologyPost200ResponseNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitInfo** | Pointer to [**[]V2DeviceDeviceIdTopologyPost200ResponseNodesInnerCircuitInfoInner**](V2DeviceDeviceIdTopologyPost200ResponseNodesInnerCircuitInfoInner.md) |  | [optional] 
**Connections** | Pointer to [**[]V2DeviceDeviceIdTopologyPost200ResponseEdgesInnerConnectionsInner**](V2DeviceDeviceIdTopologyPost200ResponseEdgesInnerConnectionsInner.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeInfo** | Pointer to [**V2DeviceDeviceIdTopologyPost200ResponseNodesInnerNodeInfo**](V2DeviceDeviceIdTopologyPost200ResponseNodesInnerNodeInfo.md) |  | [optional] 
**Quality** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV2DeviceDeviceIdTopologyPost200ResponseNodesInner

`func NewV2DeviceDeviceIdTopologyPost200ResponseNodesInner() *V2DeviceDeviceIdTopologyPost200ResponseNodesInner`

NewV2DeviceDeviceIdTopologyPost200ResponseNodesInner instantiates a new V2DeviceDeviceIdTopologyPost200ResponseNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2DeviceDeviceIdTopologyPost200ResponseNodesInnerWithDefaults

`func NewV2DeviceDeviceIdTopologyPost200ResponseNodesInnerWithDefaults() *V2DeviceDeviceIdTopologyPost200ResponseNodesInner`

NewV2DeviceDeviceIdTopologyPost200ResponseNodesInnerWithDefaults instantiates a new V2DeviceDeviceIdTopologyPost200ResponseNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitInfo

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetCircuitInfo() []V2DeviceDeviceIdTopologyPost200ResponseNodesInnerCircuitInfoInner`

GetCircuitInfo returns the CircuitInfo field if non-nil, zero value otherwise.

### GetCircuitInfoOk

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetCircuitInfoOk() (*[]V2DeviceDeviceIdTopologyPost200ResponseNodesInnerCircuitInfoInner, bool)`

GetCircuitInfoOk returns a tuple with the CircuitInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitInfo

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) SetCircuitInfo(v []V2DeviceDeviceIdTopologyPost200ResponseNodesInnerCircuitInfoInner)`

SetCircuitInfo sets CircuitInfo field to given value.

### HasCircuitInfo

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) HasCircuitInfo() bool`

HasCircuitInfo returns a boolean if a field has been set.

### GetConnections

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetConnections() []V2DeviceDeviceIdTopologyPost200ResponseEdgesInnerConnectionsInner`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetConnectionsOk() (*[]V2DeviceDeviceIdTopologyPost200ResponseEdgesInnerConnectionsInner, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) SetConnections(v []V2DeviceDeviceIdTopologyPost200ResponseEdgesInnerConnectionsInner)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetId

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeInfo

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetNodeInfo() V2DeviceDeviceIdTopologyPost200ResponseNodesInnerNodeInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetNodeInfoOk() (*V2DeviceDeviceIdTopologyPost200ResponseNodesInnerNodeInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) SetNodeInfo(v V2DeviceDeviceIdTopologyPost200ResponseNodesInnerNodeInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetQuality

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetType

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V2DeviceDeviceIdTopologyPost200ResponseNodesInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


