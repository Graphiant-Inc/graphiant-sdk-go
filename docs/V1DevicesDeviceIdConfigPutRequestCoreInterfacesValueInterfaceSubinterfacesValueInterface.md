# V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreNeighbor** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor.md) |  | [optional] 
**Default** | Pointer to **map[string]interface{}** |  | [optional] 
**GatewayNeighbor** | Pointer to [**V1AccountMfaGet200Response**](V1AccountMfaGet200Response.md) |  | [optional] 
**Interface** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface.md) |  | [optional] 
**InterfaceType** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType.md) |  | [optional] 
**VlanId** | Pointer to **int32** |  | [optional] 
**Wan** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetCoreNeighbor() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor`

GetCoreNeighbor returns the CoreNeighbor field if non-nil, zero value otherwise.

### GetCoreNeighborOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetCoreNeighborOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor, bool)`

GetCoreNeighborOk returns a tuple with the CoreNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetCoreNeighbor(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor)`

SetCoreNeighbor sets CoreNeighbor field to given value.

### HasCoreNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasCoreNeighbor() bool`

HasCoreNeighbor returns a boolean if a field has been set.

### GetDefault

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetDefault(v map[string]interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetGatewayNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetGatewayNeighbor() V1AccountMfaGet200Response`

GetGatewayNeighbor returns the GatewayNeighbor field if non-nil, zero value otherwise.

### GetGatewayNeighborOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetGatewayNeighborOk() (*V1AccountMfaGet200Response, bool)`

GetGatewayNeighborOk returns a tuple with the GatewayNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetGatewayNeighbor(v V1AccountMfaGet200Response)`

SetGatewayNeighbor sets GatewayNeighbor field to given value.

### HasGatewayNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasGatewayNeighbor() bool`

HasGatewayNeighbor returns a boolean if a field has been set.

### GetInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetInterface() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetInterfaceOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetInterface(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetInterfaceType

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetInterfaceType() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetInterfaceTypeOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetInterfaceType(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceInterfaceType)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetVlanId

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetWan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetWan() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) GetWanOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) SetWan(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterfaceWan)`

SetWan sets Wan field to given value.

### HasWan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceSubinterfacesValueInterface) HasWan() bool`

HasWan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


