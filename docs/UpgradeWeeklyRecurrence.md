# UpgradeWeeklyRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **int32** | Which occurrence of the weekday in the month applies for weekly-style recurrence (1–52, aligned with schedule validation). (required) | 
**Weekday** | **string** | Day of week for the weekly recurrence. (required) | 

## Methods

### NewUpgradeWeeklyRecurrence

`func NewUpgradeWeeklyRecurrence(interval int32, weekday string, ) *UpgradeWeeklyRecurrence`

NewUpgradeWeeklyRecurrence instantiates a new UpgradeWeeklyRecurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeWeeklyRecurrenceWithDefaults

`func NewUpgradeWeeklyRecurrenceWithDefaults() *UpgradeWeeklyRecurrence`

NewUpgradeWeeklyRecurrenceWithDefaults instantiates a new UpgradeWeeklyRecurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *UpgradeWeeklyRecurrence) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UpgradeWeeklyRecurrence) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UpgradeWeeklyRecurrence) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetWeekday

`func (o *UpgradeWeeklyRecurrence) GetWeekday() string`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *UpgradeWeeklyRecurrence) GetWeekdayOk() (*string, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *UpgradeWeeklyRecurrence) SetWeekday(v string)`

SetWeekday sets Weekday field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


