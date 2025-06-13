# V1AppsAppSummaryPost200ResponseAppSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppHealth** | Pointer to **map[string]int32** |  | [optional] 
**AppIncidents** | Pointer to [**V1AppsAppSummaryPost200ResponseAppSummaryAppIncidents**](V1AppsAppSummaryPost200ResponseAppSummaryAppIncidents.md) |  | [optional] 
**AppsOnDeviceCount** | Pointer to **int32** |  | [optional] 
**AverageQoe** | Pointer to **float64** |  | [optional] 
**CircuitsIncidents** | Pointer to [**[]V1AppsAppSummaryPost200ResponseAppSummaryCircuitsIncidentsInner**](V1AppsAppSummaryPost200ResponseAppSummaryCircuitsIncidentsInner.md) |  | [optional] 
**CircuitsIncidentsv2** | Pointer to [**[]V1AppsAppSummaryPost200ResponseAppSummaryCircuitsIncidentsv2Inner**](V1AppsAppSummaryPost200ResponseAppSummaryCircuitsIncidentsv2Inner.md) |  | [optional] 

## Methods

### NewV1AppsAppSummaryPost200ResponseAppSummary

`func NewV1AppsAppSummaryPost200ResponseAppSummary() *V1AppsAppSummaryPost200ResponseAppSummary`

NewV1AppsAppSummaryPost200ResponseAppSummary instantiates a new V1AppsAppSummaryPost200ResponseAppSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AppsAppSummaryPost200ResponseAppSummaryWithDefaults

`func NewV1AppsAppSummaryPost200ResponseAppSummaryWithDefaults() *V1AppsAppSummaryPost200ResponseAppSummary`

NewV1AppsAppSummaryPost200ResponseAppSummaryWithDefaults instantiates a new V1AppsAppSummaryPost200ResponseAppSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppHealth

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetAppHealth() map[string]int32`

GetAppHealth returns the AppHealth field if non-nil, zero value otherwise.

### GetAppHealthOk

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetAppHealthOk() (*map[string]int32, bool)`

GetAppHealthOk returns a tuple with the AppHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppHealth

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) SetAppHealth(v map[string]int32)`

SetAppHealth sets AppHealth field to given value.

### HasAppHealth

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) HasAppHealth() bool`

HasAppHealth returns a boolean if a field has been set.

### GetAppIncidents

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetAppIncidents() V1AppsAppSummaryPost200ResponseAppSummaryAppIncidents`

GetAppIncidents returns the AppIncidents field if non-nil, zero value otherwise.

### GetAppIncidentsOk

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetAppIncidentsOk() (*V1AppsAppSummaryPost200ResponseAppSummaryAppIncidents, bool)`

GetAppIncidentsOk returns a tuple with the AppIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIncidents

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) SetAppIncidents(v V1AppsAppSummaryPost200ResponseAppSummaryAppIncidents)`

SetAppIncidents sets AppIncidents field to given value.

### HasAppIncidents

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) HasAppIncidents() bool`

HasAppIncidents returns a boolean if a field has been set.

### GetAppsOnDeviceCount

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetAppsOnDeviceCount() int32`

GetAppsOnDeviceCount returns the AppsOnDeviceCount field if non-nil, zero value otherwise.

### GetAppsOnDeviceCountOk

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetAppsOnDeviceCountOk() (*int32, bool)`

GetAppsOnDeviceCountOk returns a tuple with the AppsOnDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsOnDeviceCount

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) SetAppsOnDeviceCount(v int32)`

SetAppsOnDeviceCount sets AppsOnDeviceCount field to given value.

### HasAppsOnDeviceCount

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) HasAppsOnDeviceCount() bool`

HasAppsOnDeviceCount returns a boolean if a field has been set.

### GetAverageQoe

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetAverageQoe() float64`

GetAverageQoe returns the AverageQoe field if non-nil, zero value otherwise.

### GetAverageQoeOk

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetAverageQoeOk() (*float64, bool)`

GetAverageQoeOk returns a tuple with the AverageQoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageQoe

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) SetAverageQoe(v float64)`

SetAverageQoe sets AverageQoe field to given value.

### HasAverageQoe

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) HasAverageQoe() bool`

HasAverageQoe returns a boolean if a field has been set.

### GetCircuitsIncidents

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetCircuitsIncidents() []V1AppsAppSummaryPost200ResponseAppSummaryCircuitsIncidentsInner`

GetCircuitsIncidents returns the CircuitsIncidents field if non-nil, zero value otherwise.

### GetCircuitsIncidentsOk

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetCircuitsIncidentsOk() (*[]V1AppsAppSummaryPost200ResponseAppSummaryCircuitsIncidentsInner, bool)`

GetCircuitsIncidentsOk returns a tuple with the CircuitsIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitsIncidents

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) SetCircuitsIncidents(v []V1AppsAppSummaryPost200ResponseAppSummaryCircuitsIncidentsInner)`

SetCircuitsIncidents sets CircuitsIncidents field to given value.

### HasCircuitsIncidents

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) HasCircuitsIncidents() bool`

HasCircuitsIncidents returns a boolean if a field has been set.

### GetCircuitsIncidentsv2

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetCircuitsIncidentsv2() []V1AppsAppSummaryPost200ResponseAppSummaryCircuitsIncidentsv2Inner`

GetCircuitsIncidentsv2 returns the CircuitsIncidentsv2 field if non-nil, zero value otherwise.

### GetCircuitsIncidentsv2Ok

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) GetCircuitsIncidentsv2Ok() (*[]V1AppsAppSummaryPost200ResponseAppSummaryCircuitsIncidentsv2Inner, bool)`

GetCircuitsIncidentsv2Ok returns a tuple with the CircuitsIncidentsv2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitsIncidentsv2

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) SetCircuitsIncidentsv2(v []V1AppsAppSummaryPost200ResponseAppSummaryCircuitsIncidentsv2Inner)`

SetCircuitsIncidentsv2 sets CircuitsIncidentsv2 field to given value.

### HasCircuitsIncidentsv2

`func (o *V1AppsAppSummaryPost200ResponseAppSummary) HasCircuitsIncidentsv2() bool`

HasCircuitsIncidentsv2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


