# V1EdgesHardwareUnassignedGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inventory** | Pointer to [**[]V1DevicesInventoryGet200ResponseInventoryInner**](V1DevicesInventoryGet200ResponseInventoryInner.md) |  | [optional] 
**IsNewCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1EdgesHardwareUnassignedGet200Response

`func NewV1EdgesHardwareUnassignedGet200Response() *V1EdgesHardwareUnassignedGet200Response`

NewV1EdgesHardwareUnassignedGet200Response instantiates a new V1EdgesHardwareUnassignedGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EdgesHardwareUnassignedGet200ResponseWithDefaults

`func NewV1EdgesHardwareUnassignedGet200ResponseWithDefaults() *V1EdgesHardwareUnassignedGet200Response`

NewV1EdgesHardwareUnassignedGet200ResponseWithDefaults instantiates a new V1EdgesHardwareUnassignedGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventory

`func (o *V1EdgesHardwareUnassignedGet200Response) GetInventory() []V1DevicesInventoryGet200ResponseInventoryInner`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *V1EdgesHardwareUnassignedGet200Response) GetInventoryOk() (*[]V1DevicesInventoryGet200ResponseInventoryInner, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *V1EdgesHardwareUnassignedGet200Response) SetInventory(v []V1DevicesInventoryGet200ResponseInventoryInner)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *V1EdgesHardwareUnassignedGet200Response) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetIsNewCount

`func (o *V1EdgesHardwareUnassignedGet200Response) GetIsNewCount() int32`

GetIsNewCount returns the IsNewCount field if non-nil, zero value otherwise.

### GetIsNewCountOk

`func (o *V1EdgesHardwareUnassignedGet200Response) GetIsNewCountOk() (*int32, bool)`

GetIsNewCountOk returns a tuple with the IsNewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewCount

`func (o *V1EdgesHardwareUnassignedGet200Response) SetIsNewCount(v int32)`

SetIsNewCount sets IsNewCount field to given value.

### HasIsNewCount

`func (o *V1EdgesHardwareUnassignedGet200Response) HasIsNewCount() bool`

HasIsNewCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


