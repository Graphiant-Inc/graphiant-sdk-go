# StatsmonExtranetPrefixHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to **string** | the health status of the prefix (Healthy, Impaired, Down) | [optional] 
**MissingPrefixes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStatsmonExtranetPrefixHealth

`func NewStatsmonExtranetPrefixHealth() *StatsmonExtranetPrefixHealth`

NewStatsmonExtranetPrefixHealth instantiates a new StatsmonExtranetPrefixHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsmonExtranetPrefixHealthWithDefaults

`func NewStatsmonExtranetPrefixHealthWithDefaults() *StatsmonExtranetPrefixHealth`

NewStatsmonExtranetPrefixHealthWithDefaults instantiates a new StatsmonExtranetPrefixHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *StatsmonExtranetPrefixHealth) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StatsmonExtranetPrefixHealth) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StatsmonExtranetPrefixHealth) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StatsmonExtranetPrefixHealth) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetMissingPrefixes

`func (o *StatsmonExtranetPrefixHealth) GetMissingPrefixes() []string`

GetMissingPrefixes returns the MissingPrefixes field if non-nil, zero value otherwise.

### GetMissingPrefixesOk

`func (o *StatsmonExtranetPrefixHealth) GetMissingPrefixesOk() (*[]string, bool)`

GetMissingPrefixesOk returns a tuple with the MissingPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingPrefixes

`func (o *StatsmonExtranetPrefixHealth) SetMissingPrefixes(v []string)`

SetMissingPrefixes sets MissingPrefixes field to given value.

### HasMissingPrefixes

`func (o *StatsmonExtranetPrefixHealth) HasMissingPrefixes() bool`

HasMissingPrefixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


