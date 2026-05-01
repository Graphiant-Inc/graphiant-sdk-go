# V1GlobalContentFiltersGetResponseRowRuleEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainCategory** | Pointer to **string** | Name of the domain category whose traffic is blocked by this rule. | [optional] 
**ExceptionWildcards** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1GlobalContentFiltersGetResponseRowRuleEntry

`func NewV1GlobalContentFiltersGetResponseRowRuleEntry() *V1GlobalContentFiltersGetResponseRowRuleEntry`

NewV1GlobalContentFiltersGetResponseRowRuleEntry instantiates a new V1GlobalContentFiltersGetResponseRowRuleEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalContentFiltersGetResponseRowRuleEntryWithDefaults

`func NewV1GlobalContentFiltersGetResponseRowRuleEntryWithDefaults() *V1GlobalContentFiltersGetResponseRowRuleEntry`

NewV1GlobalContentFiltersGetResponseRowRuleEntryWithDefaults instantiates a new V1GlobalContentFiltersGetResponseRowRuleEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainCategory

`func (o *V1GlobalContentFiltersGetResponseRowRuleEntry) GetDomainCategory() string`

GetDomainCategory returns the DomainCategory field if non-nil, zero value otherwise.

### GetDomainCategoryOk

`func (o *V1GlobalContentFiltersGetResponseRowRuleEntry) GetDomainCategoryOk() (*string, bool)`

GetDomainCategoryOk returns a tuple with the DomainCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCategory

`func (o *V1GlobalContentFiltersGetResponseRowRuleEntry) SetDomainCategory(v string)`

SetDomainCategory sets DomainCategory field to given value.

### HasDomainCategory

`func (o *V1GlobalContentFiltersGetResponseRowRuleEntry) HasDomainCategory() bool`

HasDomainCategory returns a boolean if a field has been set.

### GetExceptionWildcards

`func (o *V1GlobalContentFiltersGetResponseRowRuleEntry) GetExceptionWildcards() []string`

GetExceptionWildcards returns the ExceptionWildcards field if non-nil, zero value otherwise.

### GetExceptionWildcardsOk

`func (o *V1GlobalContentFiltersGetResponseRowRuleEntry) GetExceptionWildcardsOk() (*[]string, bool)`

GetExceptionWildcardsOk returns a tuple with the ExceptionWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionWildcards

`func (o *V1GlobalContentFiltersGetResponseRowRuleEntry) SetExceptionWildcards(v []string)`

SetExceptionWildcards sets ExceptionWildcards field to given value.

### HasExceptionWildcards

`func (o *V1GlobalContentFiltersGetResponseRowRuleEntry) HasExceptionWildcards() bool`

HasExceptionWildcards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


