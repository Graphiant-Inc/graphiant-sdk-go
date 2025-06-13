# V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bfd** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborBfd**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborBfd.md) |  | [optional] 
**Cost** | Pointer to **int32** |  | [optional] 
**DeadIntervalValue** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceDeadIntervalValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceDeadIntervalValue.md) |  | [optional] 
**DrPriority** | Pointer to **int32** |  | [optional] 
**HelloIntervalValue** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceHelloIntervalValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceHelloIntervalValue.md) |  | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**MtuIgnore** | Pointer to **bool** |  | [optional] 
**PrefixSid** | Pointer to **int32** |  | [optional] 
**RetransmitIntervalValue** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceRetransmitIntervalValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceRetransmitIntervalValue.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterfaceWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterfaceWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterfaceWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBfd

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetBfd() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborBfd`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetBfdOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborBfd, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetBfd(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborBfd)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetCost

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDeadIntervalValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetDeadIntervalValue() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceDeadIntervalValue`

GetDeadIntervalValue returns the DeadIntervalValue field if non-nil, zero value otherwise.

### GetDeadIntervalValueOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetDeadIntervalValueOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceDeadIntervalValue, bool)`

GetDeadIntervalValueOk returns a tuple with the DeadIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadIntervalValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetDeadIntervalValue(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceDeadIntervalValue)`

SetDeadIntervalValue sets DeadIntervalValue field to given value.

### HasDeadIntervalValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasDeadIntervalValue() bool`

HasDeadIntervalValue returns a boolean if a field has been set.

### GetDrPriority

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetDrPriority() int32`

GetDrPriority returns the DrPriority field if non-nil, zero value otherwise.

### GetDrPriorityOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetDrPriorityOk() (*int32, bool)`

GetDrPriorityOk returns a tuple with the DrPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrPriority

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetDrPriority(v int32)`

SetDrPriority sets DrPriority field to given value.

### HasDrPriority

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasDrPriority() bool`

HasDrPriority returns a boolean if a field has been set.

### GetHelloIntervalValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetHelloIntervalValue() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceHelloIntervalValue`

GetHelloIntervalValue returns the HelloIntervalValue field if non-nil, zero value otherwise.

### GetHelloIntervalValueOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetHelloIntervalValueOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceHelloIntervalValue, bool)`

GetHelloIntervalValueOk returns a tuple with the HelloIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloIntervalValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetHelloIntervalValue(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceHelloIntervalValue)`

SetHelloIntervalValue sets HelloIntervalValue field to given value.

### HasHelloIntervalValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasHelloIntervalValue() bool`

HasHelloIntervalValue returns a boolean if a field has been set.

### GetInterfaceName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetMtu

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetMtuIgnore

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetMtuIgnore() bool`

GetMtuIgnore returns the MtuIgnore field if non-nil, zero value otherwise.

### GetMtuIgnoreOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetMtuIgnoreOk() (*bool, bool)`

GetMtuIgnoreOk returns a tuple with the MtuIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtuIgnore

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetMtuIgnore(v bool)`

SetMtuIgnore sets MtuIgnore field to given value.

### HasMtuIgnore

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasMtuIgnore() bool`

HasMtuIgnore returns a boolean if a field has been set.

### GetPrefixSid

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetPrefixSid() int32`

GetPrefixSid returns the PrefixSid field if non-nil, zero value otherwise.

### GetPrefixSidOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetPrefixSidOk() (*int32, bool)`

GetPrefixSidOk returns a tuple with the PrefixSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSid

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetPrefixSid(v int32)`

SetPrefixSid sets PrefixSid field to given value.

### HasPrefixSid

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasPrefixSid() bool`

HasPrefixSid returns a boolean if a field has been set.

### GetRetransmitIntervalValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetRetransmitIntervalValue() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceRetransmitIntervalValue`

GetRetransmitIntervalValue returns the RetransmitIntervalValue field if non-nil, zero value otherwise.

### GetRetransmitIntervalValueOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetRetransmitIntervalValueOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceRetransmitIntervalValue, bool)`

GetRetransmitIntervalValueOk returns a tuple with the RetransmitIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitIntervalValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetRetransmitIntervalValue(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValueAreaInterfacesValueInterfaceRetransmitIntervalValue)`

SetRetransmitIntervalValue sets RetransmitIntervalValue field to given value.

### HasRetransmitIntervalValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasRetransmitIntervalValue() bool`

HasRetransmitIntervalValue returns a boolean if a field has been set.

### GetType

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceOspfInterfaceInterface) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


