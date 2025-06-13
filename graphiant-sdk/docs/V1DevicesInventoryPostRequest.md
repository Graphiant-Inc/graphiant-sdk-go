# V1DevicesInventoryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HardwareInventoryList** | Pointer to [**[]V1DevicesInventoryGet200ResponseInventoryInner**](V1DevicesInventoryGet200ResponseInventoryInner.md) |  | [optional] 

## Methods

### NewV1DevicesInventoryPostRequest

`func NewV1DevicesInventoryPostRequest() *V1DevicesInventoryPostRequest`

NewV1DevicesInventoryPostRequest instantiates a new V1DevicesInventoryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesInventoryPostRequestWithDefaults

`func NewV1DevicesInventoryPostRequestWithDefaults() *V1DevicesInventoryPostRequest`

NewV1DevicesInventoryPostRequestWithDefaults instantiates a new V1DevicesInventoryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHardwareInventoryList

`func (o *V1DevicesInventoryPostRequest) GetHardwareInventoryList() []V1DevicesInventoryGet200ResponseInventoryInner`

GetHardwareInventoryList returns the HardwareInventoryList field if non-nil, zero value otherwise.

### GetHardwareInventoryListOk

`func (o *V1DevicesInventoryPostRequest) GetHardwareInventoryListOk() (*[]V1DevicesInventoryGet200ResponseInventoryInner, bool)`

GetHardwareInventoryListOk returns a tuple with the HardwareInventoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareInventoryList

`func (o *V1DevicesInventoryPostRequest) SetHardwareInventoryList(v []V1DevicesInventoryGet200ResponseInventoryInner)`

SetHardwareInventoryList sets HardwareInventoryList field to given value.

### HasHardwareInventoryList

`func (o *V1DevicesInventoryPostRequest) HasHardwareInventoryList() bool`

HasHardwareInventoryList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


