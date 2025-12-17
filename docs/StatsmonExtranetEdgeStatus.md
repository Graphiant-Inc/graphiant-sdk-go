# StatsmonExtranetEdgeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**DisconnectedReason** | Pointer to **string** | the reason for the edge being disconnected | [optional] 
**Hostname** | Pointer to **string** | the hostname of the edge | [optional] 
**Id** | Pointer to **int64** | the id of the edge/device | [optional] 
**SiteName** | Pointer to **string** | the name of the site | [optional] 
**Status** | Pointer to **string** | the health status of the edge (Healthy, Impaired, Down) | [optional] 

## Methods

### NewStatsmonExtranetEdgeStatus

`func NewStatsmonExtranetEdgeStatus() *StatsmonExtranetEdgeStatus`

NewStatsmonExtranetEdgeStatus instantiates a new StatsmonExtranetEdgeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsmonExtranetEdgeStatusWithDefaults

`func NewStatsmonExtranetEdgeStatusWithDefaults() *StatsmonExtranetEdgeStatus`

NewStatsmonExtranetEdgeStatusWithDefaults instantiates a new StatsmonExtranetEdgeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *StatsmonExtranetEdgeStatus) GetCreatedAt() GoogleProtobufTimestamp`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StatsmonExtranetEdgeStatus) GetCreatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StatsmonExtranetEdgeStatus) SetCreatedAt(v GoogleProtobufTimestamp)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StatsmonExtranetEdgeStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisconnectedReason

`func (o *StatsmonExtranetEdgeStatus) GetDisconnectedReason() string`

GetDisconnectedReason returns the DisconnectedReason field if non-nil, zero value otherwise.

### GetDisconnectedReasonOk

`func (o *StatsmonExtranetEdgeStatus) GetDisconnectedReasonOk() (*string, bool)`

GetDisconnectedReasonOk returns a tuple with the DisconnectedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedReason

`func (o *StatsmonExtranetEdgeStatus) SetDisconnectedReason(v string)`

SetDisconnectedReason sets DisconnectedReason field to given value.

### HasDisconnectedReason

`func (o *StatsmonExtranetEdgeStatus) HasDisconnectedReason() bool`

HasDisconnectedReason returns a boolean if a field has been set.

### GetHostname

`func (o *StatsmonExtranetEdgeStatus) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *StatsmonExtranetEdgeStatus) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *StatsmonExtranetEdgeStatus) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *StatsmonExtranetEdgeStatus) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *StatsmonExtranetEdgeStatus) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatsmonExtranetEdgeStatus) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatsmonExtranetEdgeStatus) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StatsmonExtranetEdgeStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSiteName

`func (o *StatsmonExtranetEdgeStatus) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *StatsmonExtranetEdgeStatus) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *StatsmonExtranetEdgeStatus) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *StatsmonExtranetEdgeStatus) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetStatus

`func (o *StatsmonExtranetEdgeStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatsmonExtranetEdgeStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatsmonExtranetEdgeStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StatsmonExtranetEdgeStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


