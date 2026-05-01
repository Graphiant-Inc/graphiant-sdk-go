# ManaV2GlobalContentFilterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LanNames** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Display name for this global content filter configuration. | [optional] 
**Rules** | Pointer to [**[]ManaV2GlobalContentFilterRule**](ManaV2GlobalContentFilterRule.md) |  | [optional] 
**SiteListId** | Pointer to **int64** | Site list whose members this content filter applies to; omit the oneof when no site scope is set. | [optional] 
**UseAllSites** | Pointer to **bool** | When true, the filter applies to all sites in the tenant (must be the constant true). | [optional] 

## Methods

### NewManaV2GlobalContentFilterConfig

`func NewManaV2GlobalContentFilterConfig() *ManaV2GlobalContentFilterConfig`

NewManaV2GlobalContentFilterConfig instantiates a new ManaV2GlobalContentFilterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2GlobalContentFilterConfigWithDefaults

`func NewManaV2GlobalContentFilterConfigWithDefaults() *ManaV2GlobalContentFilterConfig`

NewManaV2GlobalContentFilterConfigWithDefaults instantiates a new ManaV2GlobalContentFilterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanNames

`func (o *ManaV2GlobalContentFilterConfig) GetLanNames() []string`

GetLanNames returns the LanNames field if non-nil, zero value otherwise.

### GetLanNamesOk

`func (o *ManaV2GlobalContentFilterConfig) GetLanNamesOk() (*[]string, bool)`

GetLanNamesOk returns a tuple with the LanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNames

`func (o *ManaV2GlobalContentFilterConfig) SetLanNames(v []string)`

SetLanNames sets LanNames field to given value.

### HasLanNames

`func (o *ManaV2GlobalContentFilterConfig) HasLanNames() bool`

HasLanNames returns a boolean if a field has been set.

### GetName

`func (o *ManaV2GlobalContentFilterConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2GlobalContentFilterConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2GlobalContentFilterConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2GlobalContentFilterConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *ManaV2GlobalContentFilterConfig) GetRules() []ManaV2GlobalContentFilterRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ManaV2GlobalContentFilterConfig) GetRulesOk() (*[]ManaV2GlobalContentFilterRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ManaV2GlobalContentFilterConfig) SetRules(v []ManaV2GlobalContentFilterRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ManaV2GlobalContentFilterConfig) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSiteListId

`func (o *ManaV2GlobalContentFilterConfig) GetSiteListId() int64`

GetSiteListId returns the SiteListId field if non-nil, zero value otherwise.

### GetSiteListIdOk

`func (o *ManaV2GlobalContentFilterConfig) GetSiteListIdOk() (*int64, bool)`

GetSiteListIdOk returns a tuple with the SiteListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteListId

`func (o *ManaV2GlobalContentFilterConfig) SetSiteListId(v int64)`

SetSiteListId sets SiteListId field to given value.

### HasSiteListId

`func (o *ManaV2GlobalContentFilterConfig) HasSiteListId() bool`

HasSiteListId returns a boolean if a field has been set.

### GetUseAllSites

`func (o *ManaV2GlobalContentFilterConfig) GetUseAllSites() bool`

GetUseAllSites returns the UseAllSites field if non-nil, zero value otherwise.

### GetUseAllSitesOk

`func (o *ManaV2GlobalContentFilterConfig) GetUseAllSitesOk() (*bool, bool)`

GetUseAllSitesOk returns a tuple with the UseAllSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllSites

`func (o *ManaV2GlobalContentFilterConfig) SetUseAllSites(v bool)`

SetUseAllSites sets UseAllSites field to given value.

### HasUseAllSites

`func (o *ManaV2GlobalContentFilterConfig) HasUseAllSites() bool`

HasUseAllSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


