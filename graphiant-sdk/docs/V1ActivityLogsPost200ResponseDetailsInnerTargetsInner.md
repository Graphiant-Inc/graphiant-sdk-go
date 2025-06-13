# V1ActivityLogsPost200ResponseDetailsInnerTargetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailedFailureReason** | Pointer to **string** |  | [optional] 
**EndTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Events** | Pointer to [**[]V1ActivityLogsPost200ResponseDetailsInnerTargetsInnerEventsInner**](V1ActivityLogsPost200ResponseDetailsInnerTargetsInnerEventsInner.md) |  | [optional] 
**FailureReason** | Pointer to **string** |  | [optional] 
**Ids** | Pointer to [**[]V1ActivityLogsPostRequestSelectorJobEntity**](V1ActivityLogsPostRequestSelectorJobEntity.md) |  | [optional] 
**StartTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ActivityLogsPost200ResponseDetailsInnerTargetsInner

`func NewV1ActivityLogsPost200ResponseDetailsInnerTargetsInner() *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner`

NewV1ActivityLogsPost200ResponseDetailsInnerTargetsInner instantiates a new V1ActivityLogsPost200ResponseDetailsInnerTargetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ActivityLogsPost200ResponseDetailsInnerTargetsInnerWithDefaults

`func NewV1ActivityLogsPost200ResponseDetailsInnerTargetsInnerWithDefaults() *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner`

NewV1ActivityLogsPost200ResponseDetailsInnerTargetsInnerWithDefaults instantiates a new V1ActivityLogsPost200ResponseDetailsInnerTargetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailedFailureReason

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetDetailedFailureReason() string`

GetDetailedFailureReason returns the DetailedFailureReason field if non-nil, zero value otherwise.

### GetDetailedFailureReasonOk

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetDetailedFailureReasonOk() (*string, bool)`

GetDetailedFailureReasonOk returns a tuple with the DetailedFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedFailureReason

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) SetDetailedFailureReason(v string)`

SetDetailedFailureReason sets DetailedFailureReason field to given value.

### HasDetailedFailureReason

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) HasDetailedFailureReason() bool`

HasDetailedFailureReason returns a boolean if a field has been set.

### GetEndTs

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetEndTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetEndTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) SetEndTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetEvents

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetEvents() []V1ActivityLogsPost200ResponseDetailsInnerTargetsInnerEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetEventsOk() (*[]V1ActivityLogsPost200ResponseDetailsInnerTargetsInnerEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) SetEvents(v []V1ActivityLogsPost200ResponseDetailsInnerTargetsInnerEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFailureReason

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetIds

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetIds() []V1ActivityLogsPostRequestSelectorJobEntity`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetIdsOk() (*[]V1ActivityLogsPostRequestSelectorJobEntity, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) SetIds(v []V1ActivityLogsPostRequestSelectorJobEntity)`

SetIds sets Ids field to given value.

### HasIds

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetStartTs

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetStartTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetStartTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) SetStartTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetStatus

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ActivityLogsPost200ResponseDetailsInnerTargetsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


