# StatsmonExtranetServerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | the IP address of the server | [optional] 
**Status** | Pointer to **string** | the health status of the server (Healthy, Impaired, Down) | [optional] 

## Methods

### NewStatsmonExtranetServerStatus

`func NewStatsmonExtranetServerStatus() *StatsmonExtranetServerStatus`

NewStatsmonExtranetServerStatus instantiates a new StatsmonExtranetServerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsmonExtranetServerStatusWithDefaults

`func NewStatsmonExtranetServerStatusWithDefaults() *StatsmonExtranetServerStatus`

NewStatsmonExtranetServerStatusWithDefaults instantiates a new StatsmonExtranetServerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *StatsmonExtranetServerStatus) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StatsmonExtranetServerStatus) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StatsmonExtranetServerStatus) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *StatsmonExtranetServerStatus) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStatus

`func (o *StatsmonExtranetServerStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatsmonExtranetServerStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatsmonExtranetServerStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StatsmonExtranetServerStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


