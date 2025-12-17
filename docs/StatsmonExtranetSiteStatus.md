# StatsmonExtranetSiteStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | the id of the site | [optional] 
**Name** | Pointer to **string** | the name of the site | [optional] 
**Status** | Pointer to **string** | the status of the site (Healthy, Impaired, Down) | [optional] 
**Statuses** | Pointer to [**[]StatsmonExtranetServerStatus**](StatsmonExtranetServerStatus.md) |  | [optional] 

## Methods

### NewStatsmonExtranetSiteStatus

`func NewStatsmonExtranetSiteStatus() *StatsmonExtranetSiteStatus`

NewStatsmonExtranetSiteStatus instantiates a new StatsmonExtranetSiteStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsmonExtranetSiteStatusWithDefaults

`func NewStatsmonExtranetSiteStatusWithDefaults() *StatsmonExtranetSiteStatus`

NewStatsmonExtranetSiteStatusWithDefaults instantiates a new StatsmonExtranetSiteStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StatsmonExtranetSiteStatus) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatsmonExtranetSiteStatus) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatsmonExtranetSiteStatus) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StatsmonExtranetSiteStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StatsmonExtranetSiteStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatsmonExtranetSiteStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatsmonExtranetSiteStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatsmonExtranetSiteStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *StatsmonExtranetSiteStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatsmonExtranetSiteStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatsmonExtranetSiteStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StatsmonExtranetSiteStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatuses

`func (o *StatsmonExtranetSiteStatus) GetStatuses() []StatsmonExtranetServerStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *StatsmonExtranetSiteStatus) GetStatusesOk() (*[]StatsmonExtranetServerStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *StatsmonExtranetSiteStatus) SetStatuses(v []StatsmonExtranetServerStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *StatsmonExtranetSiteStatus) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


