# UpgradeInventoryKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** | Blob storage container this image was indexed from. Same model + version can exist in more than one container Empty for rows indexed before multi-container support. | [optional] 
**Model** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewUpgradeInventoryKey

`func NewUpgradeInventoryKey() *UpgradeInventoryKey`

NewUpgradeInventoryKey instantiates a new UpgradeInventoryKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeInventoryKeyWithDefaults

`func NewUpgradeInventoryKeyWithDefaults() *UpgradeInventoryKey`

NewUpgradeInventoryKeyWithDefaults instantiates a new UpgradeInventoryKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *UpgradeInventoryKey) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *UpgradeInventoryKey) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *UpgradeInventoryKey) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *UpgradeInventoryKey) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetModel

`func (o *UpgradeInventoryKey) GetModel() int32`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UpgradeInventoryKey) GetModelOk() (*int32, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UpgradeInventoryKey) SetModel(v int32)`

SetModel sets Model field to given value.

### HasModel

`func (o *UpgradeInventoryKey) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVersion

`func (o *UpgradeInventoryKey) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpgradeInventoryKey) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpgradeInventoryKey) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpgradeInventoryKey) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


