# V1DiagnosticPingPauseResumePost204ResponseResultPingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgLoss** | Pointer to **float32** |  | [optional] 
**AvgTime** | Pointer to **float64** |  | [optional] 
**CompletedTime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**MaxTime** | Pointer to **float64** |  | [optional] 
**MinTime** | Pointer to **float64** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DiagnosticPingPauseResumePost204ResponseResultPingResult

`func NewV1DiagnosticPingPauseResumePost204ResponseResultPingResult() *V1DiagnosticPingPauseResumePost204ResponseResultPingResult`

NewV1DiagnosticPingPauseResumePost204ResponseResultPingResult instantiates a new V1DiagnosticPingPauseResumePost204ResponseResultPingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DiagnosticPingPauseResumePost204ResponseResultPingResultWithDefaults

`func NewV1DiagnosticPingPauseResumePost204ResponseResultPingResultWithDefaults() *V1DiagnosticPingPauseResumePost204ResponseResultPingResult`

NewV1DiagnosticPingPauseResumePost204ResponseResultPingResultWithDefaults instantiates a new V1DiagnosticPingPauseResumePost204ResponseResultPingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgLoss

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetAvgLoss() float32`

GetAvgLoss returns the AvgLoss field if non-nil, zero value otherwise.

### GetAvgLossOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetAvgLossOk() (*float32, bool)`

GetAvgLossOk returns a tuple with the AvgLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLoss

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) SetAvgLoss(v float32)`

SetAvgLoss sets AvgLoss field to given value.

### HasAvgLoss

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) HasAvgLoss() bool`

HasAvgLoss returns a boolean if a field has been set.

### GetAvgTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetAvgTime() float64`

GetAvgTime returns the AvgTime field if non-nil, zero value otherwise.

### GetAvgTimeOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetAvgTimeOk() (*float64, bool)`

GetAvgTimeOk returns a tuple with the AvgTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) SetAvgTime(v float64)`

SetAvgTime sets AvgTime field to given value.

### HasAvgTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) HasAvgTime() bool`

HasAvgTime returns a boolean if a field has been set.

### GetCompletedTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetCompletedTime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCompletedTime returns the CompletedTime field if non-nil, zero value otherwise.

### GetCompletedTimeOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetCompletedTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCompletedTimeOk returns a tuple with the CompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) SetCompletedTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCompletedTime sets CompletedTime field to given value.

### HasCompletedTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) HasCompletedTime() bool`

HasCompletedTime returns a boolean if a field has been set.

### GetMaxTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetMaxTime() float64`

GetMaxTime returns the MaxTime field if non-nil, zero value otherwise.

### GetMaxTimeOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetMaxTimeOk() (*float64, bool)`

GetMaxTimeOk returns a tuple with the MaxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) SetMaxTime(v float64)`

SetMaxTime sets MaxTime field to given value.

### HasMaxTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) HasMaxTime() bool`

HasMaxTime returns a boolean if a field has been set.

### GetMinTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetMinTime() float64`

GetMinTime returns the MinTime field if non-nil, zero value otherwise.

### GetMinTimeOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetMinTimeOk() (*float64, bool)`

GetMinTimeOk returns a tuple with the MinTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) SetMinTime(v float64)`

SetMinTime sets MinTime field to given value.

### HasMinTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) HasMinTime() bool`

HasMinTime returns a boolean if a field has been set.

### GetResult

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultPingResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


