# V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreNeighbor** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor.md) |  | [optional] 
**CoreToCoreTunnel** | Pointer to **map[string]interface{}** |  | [optional] 
**Default** | Pointer to **map[string]interface{}** |  | [optional] 
**GatewayNeighbor** | Pointer to [**V1AccountMfaGet200Response**](V1AccountMfaGet200Response.md) |  | [optional] 
**Wan** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceTypeWan**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceTypeWan.md) |  | [optional] 
**WanManagement** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceTypeWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceTypeWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceTypeWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetCoreNeighbor() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor`

GetCoreNeighbor returns the CoreNeighbor field if non-nil, zero value otherwise.

### GetCoreNeighborOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetCoreNeighborOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor, bool)`

GetCoreNeighborOk returns a tuple with the CoreNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) SetCoreNeighbor(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceCoreNeighbor)`

SetCoreNeighbor sets CoreNeighbor field to given value.

### HasCoreNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) HasCoreNeighbor() bool`

HasCoreNeighbor returns a boolean if a field has been set.

### GetCoreToCoreTunnel

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetCoreToCoreTunnel() map[string]interface{}`

GetCoreToCoreTunnel returns the CoreToCoreTunnel field if non-nil, zero value otherwise.

### GetCoreToCoreTunnelOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetCoreToCoreTunnelOk() (*map[string]interface{}, bool)`

GetCoreToCoreTunnelOk returns a tuple with the CoreToCoreTunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreToCoreTunnel

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) SetCoreToCoreTunnel(v map[string]interface{})`

SetCoreToCoreTunnel sets CoreToCoreTunnel field to given value.

### HasCoreToCoreTunnel

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) HasCoreToCoreTunnel() bool`

HasCoreToCoreTunnel returns a boolean if a field has been set.

### GetDefault

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) SetDefault(v map[string]interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetGatewayNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetGatewayNeighbor() V1AccountMfaGet200Response`

GetGatewayNeighbor returns the GatewayNeighbor field if non-nil, zero value otherwise.

### GetGatewayNeighborOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetGatewayNeighborOk() (*V1AccountMfaGet200Response, bool)`

GetGatewayNeighborOk returns a tuple with the GatewayNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) SetGatewayNeighbor(v V1AccountMfaGet200Response)`

SetGatewayNeighbor sets GatewayNeighbor field to given value.

### HasGatewayNeighbor

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) HasGatewayNeighbor() bool`

HasGatewayNeighbor returns a boolean if a field has been set.

### GetWan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetWan() V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceTypeWan`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetWanOk() (*V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceTypeWan, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) SetWan(v V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceTypeWan)`

SetWan sets Wan field to given value.

### HasWan

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) HasWan() bool`

HasWan returns a boolean if a field has been set.

### GetWanManagement

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetWanManagement() map[string]interface{}`

GetWanManagement returns the WanManagement field if non-nil, zero value otherwise.

### GetWanManagementOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) GetWanManagementOk() (*map[string]interface{}, bool)`

GetWanManagementOk returns a tuple with the WanManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanManagement

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) SetWanManagement(v map[string]interface{})`

SetWanManagement sets WanManagement field to given value.

### HasWanManagement

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceInterfaceType) HasWanManagement() bool`

HasWanManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


