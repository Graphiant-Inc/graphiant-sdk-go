# V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerAppsInner**](V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerAppsInner.md) |  | [optional] 
**AssuranceId** | Pointer to **int64** |  | [optional] 
**AssuranceName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**FlexAlgo** | Pointer to **string** |  | [optional] 
**Lans** | Pointer to [**[]V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerLansInner**](V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerLansInner.md) |  | [optional] 
**Sites** | Pointer to [**[]V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerSitesInner**](V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerSitesInner.md) |  | [optional] 
**UpdatedAt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1DataAssuranceAssurancesGlobalGet200ResponseRowsInner

`func NewV1DataAssuranceAssurancesGlobalGet200ResponseRowsInner() *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner`

NewV1DataAssuranceAssurancesGlobalGet200ResponseRowsInner instantiates a new V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerWithDefaults

`func NewV1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerWithDefaults() *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner`

NewV1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerWithDefaults instantiates a new V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetApps() []V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerAppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetAppsOk() (*[]V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerAppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) SetApps(v []V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerAppsInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAssuranceId

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetAssuranceId() int64`

GetAssuranceId returns the AssuranceId field if non-nil, zero value otherwise.

### GetAssuranceIdOk

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetAssuranceIdOk() (*int64, bool)`

GetAssuranceIdOk returns a tuple with the AssuranceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceId

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) SetAssuranceId(v int64)`

SetAssuranceId sets AssuranceId field to given value.

### HasAssuranceId

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) HasAssuranceId() bool`

HasAssuranceId returns a boolean if a field has been set.

### GetAssuranceName

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetAssuranceName() string`

GetAssuranceName returns the AssuranceName field if non-nil, zero value otherwise.

### GetAssuranceNameOk

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetAssuranceNameOk() (*string, bool)`

GetAssuranceNameOk returns a tuple with the AssuranceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceName

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) SetAssuranceName(v string)`

SetAssuranceName sets AssuranceName field to given value.

### HasAssuranceName

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) HasAssuranceName() bool`

HasAssuranceName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetCreatedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetCreatedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) SetCreatedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFlexAlgo

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetFlexAlgo() string`

GetFlexAlgo returns the FlexAlgo field if non-nil, zero value otherwise.

### GetFlexAlgoOk

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetFlexAlgoOk() (*string, bool)`

GetFlexAlgoOk returns a tuple with the FlexAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAlgo

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) SetFlexAlgo(v string)`

SetFlexAlgo sets FlexAlgo field to given value.

### HasFlexAlgo

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) HasFlexAlgo() bool`

HasFlexAlgo returns a boolean if a field has been set.

### GetLans

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetLans() []V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerLansInner`

GetLans returns the Lans field if non-nil, zero value otherwise.

### GetLansOk

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetLansOk() (*[]V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerLansInner, bool)`

GetLansOk returns a tuple with the Lans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLans

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) SetLans(v []V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerLansInner)`

SetLans sets Lans field to given value.

### HasLans

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) HasLans() bool`

HasLans returns a boolean if a field has been set.

### GetSites

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetSites() []V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetSitesOk() (*[]V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) SetSites(v []V1DataAssuranceAssurancesGlobalGet200ResponseRowsInnerSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetUpdatedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) GetUpdatedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) SetUpdatedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V1DataAssuranceAssurancesGlobalGet200ResponseRowsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


