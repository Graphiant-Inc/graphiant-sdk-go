# V1AlarmsListGet200ResponseAlarmsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedBy** | Pointer to **string** |  | [optional] 
**AlarmId** | Pointer to **int64** |  | [optional] 
**AlarmTypeId** | Pointer to **string** |  | [optional] 
**AlarmTypeQualifier** | Pointer to **string** |  | [optional] 
**AltComponent** | Pointer to **string** |  | [optional] 
**BootId** | Pointer to **string** |  | [optional] 
**Component** | Pointer to **string** |  | [optional] 
**Created** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsCleared** | Pointer to **bool** |  | [optional] 
**IsResolved** | Pointer to **bool** |  | [optional] 
**LastChanged** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**LastRaised** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**PerceivedSeverity** | Pointer to **string** |  | [optional] 
**ResolvedBy** | Pointer to **string** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**Where** | Pointer to **string** |  | [optional] 

## Methods

### NewV1AlarmsListGet200ResponseAlarmsInner

`func NewV1AlarmsListGet200ResponseAlarmsInner() *V1AlarmsListGet200ResponseAlarmsInner`

NewV1AlarmsListGet200ResponseAlarmsInner instantiates a new V1AlarmsListGet200ResponseAlarmsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AlarmsListGet200ResponseAlarmsInnerWithDefaults

`func NewV1AlarmsListGet200ResponseAlarmsInnerWithDefaults() *V1AlarmsListGet200ResponseAlarmsInner`

NewV1AlarmsListGet200ResponseAlarmsInnerWithDefaults instantiates a new V1AlarmsListGet200ResponseAlarmsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedBy

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetAcknowledgedBy() string`

GetAcknowledgedBy returns the AcknowledgedBy field if non-nil, zero value otherwise.

### GetAcknowledgedByOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetAcknowledgedByOk() (*string, bool)`

GetAcknowledgedByOk returns a tuple with the AcknowledgedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedBy

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetAcknowledgedBy(v string)`

SetAcknowledgedBy sets AcknowledgedBy field to given value.

### HasAcknowledgedBy

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasAcknowledgedBy() bool`

HasAcknowledgedBy returns a boolean if a field has been set.

### GetAlarmId

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetAlarmId() int64`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetAlarmIdOk() (*int64, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetAlarmId(v int64)`

SetAlarmId sets AlarmId field to given value.

### HasAlarmId

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasAlarmId() bool`

HasAlarmId returns a boolean if a field has been set.

### GetAlarmTypeId

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetAlarmTypeId() string`

GetAlarmTypeId returns the AlarmTypeId field if non-nil, zero value otherwise.

### GetAlarmTypeIdOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetAlarmTypeIdOk() (*string, bool)`

GetAlarmTypeIdOk returns a tuple with the AlarmTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmTypeId

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetAlarmTypeId(v string)`

SetAlarmTypeId sets AlarmTypeId field to given value.

### HasAlarmTypeId

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasAlarmTypeId() bool`

HasAlarmTypeId returns a boolean if a field has been set.

### GetAlarmTypeQualifier

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetAlarmTypeQualifier() string`

GetAlarmTypeQualifier returns the AlarmTypeQualifier field if non-nil, zero value otherwise.

### GetAlarmTypeQualifierOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetAlarmTypeQualifierOk() (*string, bool)`

GetAlarmTypeQualifierOk returns a tuple with the AlarmTypeQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmTypeQualifier

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetAlarmTypeQualifier(v string)`

SetAlarmTypeQualifier sets AlarmTypeQualifier field to given value.

### HasAlarmTypeQualifier

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasAlarmTypeQualifier() bool`

HasAlarmTypeQualifier returns a boolean if a field has been set.

### GetAltComponent

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetAltComponent() string`

GetAltComponent returns the AltComponent field if non-nil, zero value otherwise.

### GetAltComponentOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetAltComponentOk() (*string, bool)`

GetAltComponentOk returns a tuple with the AltComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltComponent

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetAltComponent(v string)`

SetAltComponent sets AltComponent field to given value.

### HasAltComponent

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasAltComponent() bool`

HasAltComponent returns a boolean if a field has been set.

### GetBootId

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetBootId() string`

GetBootId returns the BootId field if non-nil, zero value otherwise.

### GetBootIdOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetBootIdOk() (*string, bool)`

GetBootIdOk returns a tuple with the BootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootId

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetBootId(v string)`

SetBootId sets BootId field to given value.

### HasBootId

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasBootId() bool`

HasBootId returns a boolean if a field has been set.

### GetComponent

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCreated

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetCreated() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetCreatedOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetCreated(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsCleared

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetIsCleared() bool`

GetIsCleared returns the IsCleared field if non-nil, zero value otherwise.

### GetIsClearedOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetIsClearedOk() (*bool, bool)`

GetIsClearedOk returns a tuple with the IsCleared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCleared

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetIsCleared(v bool)`

SetIsCleared sets IsCleared field to given value.

### HasIsCleared

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasIsCleared() bool`

HasIsCleared returns a boolean if a field has been set.

### GetIsResolved

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetIsResolved() bool`

GetIsResolved returns the IsResolved field if non-nil, zero value otherwise.

### GetIsResolvedOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetIsResolvedOk() (*bool, bool)`

GetIsResolvedOk returns a tuple with the IsResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResolved

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetIsResolved(v bool)`

SetIsResolved sets IsResolved field to given value.

### HasIsResolved

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasIsResolved() bool`

HasIsResolved returns a boolean if a field has been set.

### GetLastChanged

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetLastChanged() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetLastChanged returns the LastChanged field if non-nil, zero value otherwise.

### GetLastChangedOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetLastChangedOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetLastChangedOk returns a tuple with the LastChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChanged

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetLastChanged(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetLastChanged sets LastChanged field to given value.

### HasLastChanged

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasLastChanged() bool`

HasLastChanged returns a boolean if a field has been set.

### GetLastRaised

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetLastRaised() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetLastRaised returns the LastRaised field if non-nil, zero value otherwise.

### GetLastRaisedOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetLastRaisedOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetLastRaisedOk returns a tuple with the LastRaised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRaised

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetLastRaised(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetLastRaised sets LastRaised field to given value.

### HasLastRaised

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasLastRaised() bool`

HasLastRaised returns a boolean if a field has been set.

### GetPerceivedSeverity

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetPerceivedSeverity() string`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetPerceivedSeverityOk() (*string, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetPerceivedSeverity(v string)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.

### HasPerceivedSeverity

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasPerceivedSeverity() bool`

HasPerceivedSeverity returns a boolean if a field has been set.

### GetResolvedBy

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetResolvedBy() string`

GetResolvedBy returns the ResolvedBy field if non-nil, zero value otherwise.

### GetResolvedByOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetResolvedByOk() (*string, bool)`

GetResolvedByOk returns a tuple with the ResolvedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedBy

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetResolvedBy(v string)`

SetResolvedBy sets ResolvedBy field to given value.

### HasResolvedBy

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasResolvedBy() bool`

HasResolvedBy returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetWhere

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetWhere() string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *V1AlarmsListGet200ResponseAlarmsInner) GetWhereOk() (*string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *V1AlarmsListGet200ResponseAlarmsInner) SetWhere(v string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *V1AlarmsListGet200ResponseAlarmsInner) HasWhere() bool`

HasWhere returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


