# UpgradeRecurringSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monthly** | Pointer to [**UpgradeMonthlyRecurrence**](UpgradeMonthlyRecurrence.md) |  | [optional] 
**StartsAtTs** | [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | 
**Weekly** | Pointer to [**UpgradeWeeklyRecurrence**](UpgradeWeeklyRecurrence.md) |  | [optional] 
**Yearly** | Pointer to [**UpgradeYearlyRecurrence**](UpgradeYearlyRecurrence.md) |  | [optional] 

## Methods

### NewUpgradeRecurringSchedule

`func NewUpgradeRecurringSchedule(startsAtTs GoogleProtobufTimestamp, ) *UpgradeRecurringSchedule`

NewUpgradeRecurringSchedule instantiates a new UpgradeRecurringSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeRecurringScheduleWithDefaults

`func NewUpgradeRecurringScheduleWithDefaults() *UpgradeRecurringSchedule`

NewUpgradeRecurringScheduleWithDefaults instantiates a new UpgradeRecurringSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonthly

`func (o *UpgradeRecurringSchedule) GetMonthly() UpgradeMonthlyRecurrence`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *UpgradeRecurringSchedule) GetMonthlyOk() (*UpgradeMonthlyRecurrence, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *UpgradeRecurringSchedule) SetMonthly(v UpgradeMonthlyRecurrence)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *UpgradeRecurringSchedule) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.

### GetStartsAtTs

`func (o *UpgradeRecurringSchedule) GetStartsAtTs() GoogleProtobufTimestamp`

GetStartsAtTs returns the StartsAtTs field if non-nil, zero value otherwise.

### GetStartsAtTsOk

`func (o *UpgradeRecurringSchedule) GetStartsAtTsOk() (*GoogleProtobufTimestamp, bool)`

GetStartsAtTsOk returns a tuple with the StartsAtTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAtTs

`func (o *UpgradeRecurringSchedule) SetStartsAtTs(v GoogleProtobufTimestamp)`

SetStartsAtTs sets StartsAtTs field to given value.


### GetWeekly

`func (o *UpgradeRecurringSchedule) GetWeekly() UpgradeWeeklyRecurrence`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *UpgradeRecurringSchedule) GetWeeklyOk() (*UpgradeWeeklyRecurrence, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *UpgradeRecurringSchedule) SetWeekly(v UpgradeWeeklyRecurrence)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *UpgradeRecurringSchedule) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### GetYearly

`func (o *UpgradeRecurringSchedule) GetYearly() UpgradeYearlyRecurrence`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *UpgradeRecurringSchedule) GetYearlyOk() (*UpgradeYearlyRecurrence, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *UpgradeRecurringSchedule) SetYearly(v UpgradeYearlyRecurrence)`

SetYearly sets Yearly field to given value.

### HasYearly

`func (o *UpgradeRecurringSchedule) HasYearly() bool`

HasYearly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


