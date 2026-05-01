# ManaV2DomainCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional longer text describing what sites or traffic belong in this category. | [optional] 
**Id** | Pointer to **int64** | Stable identifier for this domain category in the content-filter catalog. | [optional] 
**Name** | Pointer to **string** | Human-readable name of the domain category. | [optional] 
**Type** | Pointer to **string** | Optional category classifier or grouping hint from the catalog (e.g. vendor-specific type). | [optional] 

## Methods

### NewManaV2DomainCategory

`func NewManaV2DomainCategory() *ManaV2DomainCategory`

NewManaV2DomainCategory instantiates a new ManaV2DomainCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2DomainCategoryWithDefaults

`func NewManaV2DomainCategoryWithDefaults() *ManaV2DomainCategory`

NewManaV2DomainCategoryWithDefaults instantiates a new ManaV2DomainCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ManaV2DomainCategory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2DomainCategory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2DomainCategory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2DomainCategory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ManaV2DomainCategory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2DomainCategory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2DomainCategory) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ManaV2DomainCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ManaV2DomainCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2DomainCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2DomainCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2DomainCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ManaV2DomainCategory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManaV2DomainCategory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManaV2DomainCategory) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManaV2DomainCategory) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


