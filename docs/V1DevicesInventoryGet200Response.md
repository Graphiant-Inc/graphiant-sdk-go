# V1DevicesInventoryGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inventory** | Pointer to [**[]V1DevicesInventoryGet200ResponseInventoryInner**](V1DevicesInventoryGet200ResponseInventoryInner.md) |  | [optional] 
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 

## Methods

### NewV1DevicesInventoryGet200Response

`func NewV1DevicesInventoryGet200Response() *V1DevicesInventoryGet200Response`

NewV1DevicesInventoryGet200Response instantiates a new V1DevicesInventoryGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesInventoryGet200ResponseWithDefaults

`func NewV1DevicesInventoryGet200ResponseWithDefaults() *V1DevicesInventoryGet200Response`

NewV1DevicesInventoryGet200ResponseWithDefaults instantiates a new V1DevicesInventoryGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventory

`func (o *V1DevicesInventoryGet200Response) GetInventory() []V1DevicesInventoryGet200ResponseInventoryInner`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *V1DevicesInventoryGet200Response) GetInventoryOk() (*[]V1DevicesInventoryGet200ResponseInventoryInner, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *V1DevicesInventoryGet200Response) SetInventory(v []V1DevicesInventoryGet200ResponseInventoryInner)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *V1DevicesInventoryGet200Response) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetPageInfo

`func (o *V1DevicesInventoryGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1DevicesInventoryGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1DevicesInventoryGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1DevicesInventoryGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


