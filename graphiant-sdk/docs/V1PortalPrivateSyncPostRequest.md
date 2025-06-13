# V1PortalPrivateSyncPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcsName** | Pointer to **string** |  | [optional] 
**Inventory** | Pointer to [**[]V1PortalPrivateSyncPostRequestInventoryInner**](V1PortalPrivateSyncPostRequestInventoryInner.md) |  | [optional] 
**IsFullSync** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1PortalPrivateSyncPostRequest

`func NewV1PortalPrivateSyncPostRequest() *V1PortalPrivateSyncPostRequest`

NewV1PortalPrivateSyncPostRequest instantiates a new V1PortalPrivateSyncPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PortalPrivateSyncPostRequestWithDefaults

`func NewV1PortalPrivateSyncPostRequestWithDefaults() *V1PortalPrivateSyncPostRequest`

NewV1PortalPrivateSyncPostRequestWithDefaults instantiates a new V1PortalPrivateSyncPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcsName

`func (o *V1PortalPrivateSyncPostRequest) GetGcsName() string`

GetGcsName returns the GcsName field if non-nil, zero value otherwise.

### GetGcsNameOk

`func (o *V1PortalPrivateSyncPostRequest) GetGcsNameOk() (*string, bool)`

GetGcsNameOk returns a tuple with the GcsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcsName

`func (o *V1PortalPrivateSyncPostRequest) SetGcsName(v string)`

SetGcsName sets GcsName field to given value.

### HasGcsName

`func (o *V1PortalPrivateSyncPostRequest) HasGcsName() bool`

HasGcsName returns a boolean if a field has been set.

### GetInventory

`func (o *V1PortalPrivateSyncPostRequest) GetInventory() []V1PortalPrivateSyncPostRequestInventoryInner`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *V1PortalPrivateSyncPostRequest) GetInventoryOk() (*[]V1PortalPrivateSyncPostRequestInventoryInner, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *V1PortalPrivateSyncPostRequest) SetInventory(v []V1PortalPrivateSyncPostRequestInventoryInner)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *V1PortalPrivateSyncPostRequest) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetIsFullSync

`func (o *V1PortalPrivateSyncPostRequest) GetIsFullSync() bool`

GetIsFullSync returns the IsFullSync field if non-nil, zero value otherwise.

### GetIsFullSyncOk

`func (o *V1PortalPrivateSyncPostRequest) GetIsFullSyncOk() (*bool, bool)`

GetIsFullSyncOk returns a tuple with the IsFullSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullSync

`func (o *V1PortalPrivateSyncPostRequest) SetIsFullSync(v bool)`

SetIsFullSync sets IsFullSync field to given value.

### HasIsFullSync

`func (o *V1PortalPrivateSyncPostRequest) HasIsFullSync() bool`

HasIsFullSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


