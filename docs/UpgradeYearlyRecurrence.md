# UpgradeYearlyRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **int32** | Calendar day (1–31) for fixed month+date yearly recurrence. | [optional] 
**Month** | **int32** | Month of year (1–12) for yearly recurrence. (required) | 
**Ordinal** | Pointer to **string** | For nth-weekday-in-month yearly recurrence; use with weekday, or use month + date. | [optional] 
**Weekday** | Pointer to **string** | Weekday paired with ordinal for yearly recurrence. | [optional] 

## Methods

### NewUpgradeYearlyRecurrence

`func NewUpgradeYearlyRecurrence(month int32, ) *UpgradeYearlyRecurrence`

NewUpgradeYearlyRecurrence instantiates a new UpgradeYearlyRecurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeYearlyRecurrenceWithDefaults

`func NewUpgradeYearlyRecurrenceWithDefaults() *UpgradeYearlyRecurrence`

NewUpgradeYearlyRecurrenceWithDefaults instantiates a new UpgradeYearlyRecurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *UpgradeYearlyRecurrence) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UpgradeYearlyRecurrence) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UpgradeYearlyRecurrence) SetDate(v int32)`

SetDate sets Date field to given value.

### HasDate

`func (o *UpgradeYearlyRecurrence) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMonth

`func (o *UpgradeYearlyRecurrence) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *UpgradeYearlyRecurrence) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *UpgradeYearlyRecurrence) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetOrdinal

`func (o *UpgradeYearlyRecurrence) GetOrdinal() string`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *UpgradeYearlyRecurrence) GetOrdinalOk() (*string, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *UpgradeYearlyRecurrence) SetOrdinal(v string)`

SetOrdinal sets Ordinal field to given value.

### HasOrdinal

`func (o *UpgradeYearlyRecurrence) HasOrdinal() bool`

HasOrdinal returns a boolean if a field has been set.

### GetWeekday

`func (o *UpgradeYearlyRecurrence) GetWeekday() string`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *UpgradeYearlyRecurrence) GetWeekdayOk() (*string, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *UpgradeYearlyRecurrence) SetWeekday(v string)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *UpgradeYearlyRecurrence) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


