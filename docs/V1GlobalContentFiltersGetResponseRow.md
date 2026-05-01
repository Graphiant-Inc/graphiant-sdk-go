# V1GlobalContentFiltersGetResponseRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**GlobalContentFilterId** | Pointer to **int64** | ID for the global content filter. | [optional] 
**GlobalContentFilterName** | Pointer to **string** | Given name of this global content filter. | [optional] 
**Lans** | Pointer to [**[]V1GlobalContentFiltersGetResponseRowLanEntry**](V1GlobalContentFiltersGetResponseRowLanEntry.md) |  | [optional] 
**Rules** | Pointer to [**[]V1GlobalContentFiltersGetResponseRowRuleEntry**](V1GlobalContentFiltersGetResponseRowRuleEntry.md) |  | [optional] 
**Sites** | Pointer to [**[]V1GlobalContentFiltersGetResponseRowSiteEntry**](V1GlobalContentFiltersGetResponseRowSiteEntry.md) |  | [optional] 
**UpdatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 

## Methods

### NewV1GlobalContentFiltersGetResponseRow

`func NewV1GlobalContentFiltersGetResponseRow() *V1GlobalContentFiltersGetResponseRow`

NewV1GlobalContentFiltersGetResponseRow instantiates a new V1GlobalContentFiltersGetResponseRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalContentFiltersGetResponseRowWithDefaults

`func NewV1GlobalContentFiltersGetResponseRowWithDefaults() *V1GlobalContentFiltersGetResponseRow`

NewV1GlobalContentFiltersGetResponseRowWithDefaults instantiates a new V1GlobalContentFiltersGetResponseRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *V1GlobalContentFiltersGetResponseRow) GetCreatedAt() GoogleProtobufTimestamp`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V1GlobalContentFiltersGetResponseRow) GetCreatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V1GlobalContentFiltersGetResponseRow) SetCreatedAt(v GoogleProtobufTimestamp)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V1GlobalContentFiltersGetResponseRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGlobalContentFilterId

`func (o *V1GlobalContentFiltersGetResponseRow) GetGlobalContentFilterId() int64`

GetGlobalContentFilterId returns the GlobalContentFilterId field if non-nil, zero value otherwise.

### GetGlobalContentFilterIdOk

`func (o *V1GlobalContentFiltersGetResponseRow) GetGlobalContentFilterIdOk() (*int64, bool)`

GetGlobalContentFilterIdOk returns a tuple with the GlobalContentFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalContentFilterId

`func (o *V1GlobalContentFiltersGetResponseRow) SetGlobalContentFilterId(v int64)`

SetGlobalContentFilterId sets GlobalContentFilterId field to given value.

### HasGlobalContentFilterId

`func (o *V1GlobalContentFiltersGetResponseRow) HasGlobalContentFilterId() bool`

HasGlobalContentFilterId returns a boolean if a field has been set.

### GetGlobalContentFilterName

`func (o *V1GlobalContentFiltersGetResponseRow) GetGlobalContentFilterName() string`

GetGlobalContentFilterName returns the GlobalContentFilterName field if non-nil, zero value otherwise.

### GetGlobalContentFilterNameOk

`func (o *V1GlobalContentFiltersGetResponseRow) GetGlobalContentFilterNameOk() (*string, bool)`

GetGlobalContentFilterNameOk returns a tuple with the GlobalContentFilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalContentFilterName

`func (o *V1GlobalContentFiltersGetResponseRow) SetGlobalContentFilterName(v string)`

SetGlobalContentFilterName sets GlobalContentFilterName field to given value.

### HasGlobalContentFilterName

`func (o *V1GlobalContentFiltersGetResponseRow) HasGlobalContentFilterName() bool`

HasGlobalContentFilterName returns a boolean if a field has been set.

### GetLans

`func (o *V1GlobalContentFiltersGetResponseRow) GetLans() []V1GlobalContentFiltersGetResponseRowLanEntry`

GetLans returns the Lans field if non-nil, zero value otherwise.

### GetLansOk

`func (o *V1GlobalContentFiltersGetResponseRow) GetLansOk() (*[]V1GlobalContentFiltersGetResponseRowLanEntry, bool)`

GetLansOk returns a tuple with the Lans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLans

`func (o *V1GlobalContentFiltersGetResponseRow) SetLans(v []V1GlobalContentFiltersGetResponseRowLanEntry)`

SetLans sets Lans field to given value.

### HasLans

`func (o *V1GlobalContentFiltersGetResponseRow) HasLans() bool`

HasLans returns a boolean if a field has been set.

### GetRules

`func (o *V1GlobalContentFiltersGetResponseRow) GetRules() []V1GlobalContentFiltersGetResponseRowRuleEntry`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *V1GlobalContentFiltersGetResponseRow) GetRulesOk() (*[]V1GlobalContentFiltersGetResponseRowRuleEntry, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *V1GlobalContentFiltersGetResponseRow) SetRules(v []V1GlobalContentFiltersGetResponseRowRuleEntry)`

SetRules sets Rules field to given value.

### HasRules

`func (o *V1GlobalContentFiltersGetResponseRow) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSites

`func (o *V1GlobalContentFiltersGetResponseRow) GetSites() []V1GlobalContentFiltersGetResponseRowSiteEntry`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *V1GlobalContentFiltersGetResponseRow) GetSitesOk() (*[]V1GlobalContentFiltersGetResponseRowSiteEntry, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *V1GlobalContentFiltersGetResponseRow) SetSites(v []V1GlobalContentFiltersGetResponseRowSiteEntry)`

SetSites sets Sites field to given value.

### HasSites

`func (o *V1GlobalContentFiltersGetResponseRow) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V1GlobalContentFiltersGetResponseRow) GetUpdatedAt() GoogleProtobufTimestamp`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V1GlobalContentFiltersGetResponseRow) GetUpdatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V1GlobalContentFiltersGetResponseRow) SetUpdatedAt(v GoogleProtobufTimestamp)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V1GlobalContentFiltersGetResponseRow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


