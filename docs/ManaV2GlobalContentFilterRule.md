# ManaV2GlobalContentFilterRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainCategoryId** | Pointer to **int64** | ID of the category whose traffic will be blocked by the content filter. | [optional] 
**ExceptionWildcards** | Pointer to **[]string** |  | [optional] 

## Methods

### NewManaV2GlobalContentFilterRule

`func NewManaV2GlobalContentFilterRule() *ManaV2GlobalContentFilterRule`

NewManaV2GlobalContentFilterRule instantiates a new ManaV2GlobalContentFilterRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2GlobalContentFilterRuleWithDefaults

`func NewManaV2GlobalContentFilterRuleWithDefaults() *ManaV2GlobalContentFilterRule`

NewManaV2GlobalContentFilterRuleWithDefaults instantiates a new ManaV2GlobalContentFilterRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainCategoryId

`func (o *ManaV2GlobalContentFilterRule) GetDomainCategoryId() int64`

GetDomainCategoryId returns the DomainCategoryId field if non-nil, zero value otherwise.

### GetDomainCategoryIdOk

`func (o *ManaV2GlobalContentFilterRule) GetDomainCategoryIdOk() (*int64, bool)`

GetDomainCategoryIdOk returns a tuple with the DomainCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCategoryId

`func (o *ManaV2GlobalContentFilterRule) SetDomainCategoryId(v int64)`

SetDomainCategoryId sets DomainCategoryId field to given value.

### HasDomainCategoryId

`func (o *ManaV2GlobalContentFilterRule) HasDomainCategoryId() bool`

HasDomainCategoryId returns a boolean if a field has been set.

### GetExceptionWildcards

`func (o *ManaV2GlobalContentFilterRule) GetExceptionWildcards() []string`

GetExceptionWildcards returns the ExceptionWildcards field if non-nil, zero value otherwise.

### GetExceptionWildcardsOk

`func (o *ManaV2GlobalContentFilterRule) GetExceptionWildcardsOk() (*[]string, bool)`

GetExceptionWildcardsOk returns a tuple with the ExceptionWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionWildcards

`func (o *ManaV2GlobalContentFilterRule) SetExceptionWildcards(v []string)`

SetExceptionWildcards sets ExceptionWildcards field to given value.

### HasExceptionWildcards

`func (o *ManaV2GlobalContentFilterRule) HasExceptionWildcards() bool`

HasExceptionWildcards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


