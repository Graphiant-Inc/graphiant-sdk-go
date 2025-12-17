# IpfixConnectionMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to **map[string]int64** |  | [optional] 
**ConnectionsV2** | Pointer to **map[string]float64** |  | [optional] 
**Name** | Pointer to **string** | the name of the connection | [optional] 

## Methods

### NewIpfixConnectionMap

`func NewIpfixConnectionMap() *IpfixConnectionMap`

NewIpfixConnectionMap instantiates a new IpfixConnectionMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpfixConnectionMapWithDefaults

`func NewIpfixConnectionMapWithDefaults() *IpfixConnectionMap`

NewIpfixConnectionMapWithDefaults instantiates a new IpfixConnectionMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *IpfixConnectionMap) GetConnections() map[string]int64`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *IpfixConnectionMap) GetConnectionsOk() (*map[string]int64, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *IpfixConnectionMap) SetConnections(v map[string]int64)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *IpfixConnectionMap) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetConnectionsV2

`func (o *IpfixConnectionMap) GetConnectionsV2() map[string]float64`

GetConnectionsV2 returns the ConnectionsV2 field if non-nil, zero value otherwise.

### GetConnectionsV2Ok

`func (o *IpfixConnectionMap) GetConnectionsV2Ok() (*map[string]float64, bool)`

GetConnectionsV2Ok returns a tuple with the ConnectionsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsV2

`func (o *IpfixConnectionMap) SetConnectionsV2(v map[string]float64)`

SetConnectionsV2 sets ConnectionsV2 field to given value.

### HasConnectionsV2

`func (o *IpfixConnectionMap) HasConnectionsV2() bool`

HasConnectionsV2 returns a boolean if a field has been set.

### GetName

`func (o *IpfixConnectionMap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpfixConnectionMap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpfixConnectionMap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpfixConnectionMap) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


