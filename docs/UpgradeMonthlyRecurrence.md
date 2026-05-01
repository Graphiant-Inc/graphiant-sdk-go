# UpgradeMonthlyRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **int32** | Calendar day of month (1–31) for fixed-date monthly recurrence; optional if ordinal and weekday are set. | [optional] 
**Ordinal** | Pointer to **string** | For nth-weekday-of-month style recurrence; use together with weekday, or use date instead. | [optional] 
**Weekday** | Pointer to **string** | Weekday paired with ordinal for monthly nth-weekday recurrence; optional if date is set. | [optional] 

## Methods

### NewUpgradeMonthlyRecurrence

`func NewUpgradeMonthlyRecurrence() *UpgradeMonthlyRecurrence`

NewUpgradeMonthlyRecurrence instantiates a new UpgradeMonthlyRecurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeMonthlyRecurrenceWithDefaults

`func NewUpgradeMonthlyRecurrenceWithDefaults() *UpgradeMonthlyRecurrence`

NewUpgradeMonthlyRecurrenceWithDefaults instantiates a new UpgradeMonthlyRecurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *UpgradeMonthlyRecurrence) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UpgradeMonthlyRecurrence) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UpgradeMonthlyRecurrence) SetDate(v int32)`

SetDate sets Date field to given value.

### HasDate

`func (o *UpgradeMonthlyRecurrence) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetOrdinal

`func (o *UpgradeMonthlyRecurrence) GetOrdinal() string`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *UpgradeMonthlyRecurrence) GetOrdinalOk() (*string, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *UpgradeMonthlyRecurrence) SetOrdinal(v string)`

SetOrdinal sets Ordinal field to given value.

### HasOrdinal

`func (o *UpgradeMonthlyRecurrence) HasOrdinal() bool`

HasOrdinal returns a boolean if a field has been set.

### GetWeekday

`func (o *UpgradeMonthlyRecurrence) GetWeekday() string`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *UpgradeMonthlyRecurrence) GetWeekdayOk() (*string, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *UpgradeMonthlyRecurrence) SetWeekday(v string)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *UpgradeMonthlyRecurrence) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


