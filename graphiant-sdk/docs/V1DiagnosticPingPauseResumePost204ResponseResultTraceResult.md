# V1DiagnosticPingPauseResumePost204ResponseResultTraceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hops** | Pointer to [**[]V1DiagnosticPingPauseResumePost204ResponseResultTraceResultHopsInner**](V1DiagnosticPingPauseResumePost204ResponseResultTraceResultHopsInner.md) |  | [optional] 
**MaxLatency** | Pointer to **float64** |  | [optional] 
**MaxLatencyHost** | Pointer to **string** |  | [optional] 
**MaxPathMtu** | Pointer to **int32** |  | [optional] 
**MinPathMtu** | Pointer to **int32** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**StoppedTime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1DiagnosticPingPauseResumePost204ResponseResultTraceResult

`func NewV1DiagnosticPingPauseResumePost204ResponseResultTraceResult() *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult`

NewV1DiagnosticPingPauseResumePost204ResponseResultTraceResult instantiates a new V1DiagnosticPingPauseResumePost204ResponseResultTraceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DiagnosticPingPauseResumePost204ResponseResultTraceResultWithDefaults

`func NewV1DiagnosticPingPauseResumePost204ResponseResultTraceResultWithDefaults() *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult`

NewV1DiagnosticPingPauseResumePost204ResponseResultTraceResultWithDefaults instantiates a new V1DiagnosticPingPauseResumePost204ResponseResultTraceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHops

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetHops() []V1DiagnosticPingPauseResumePost204ResponseResultTraceResultHopsInner`

GetHops returns the Hops field if non-nil, zero value otherwise.

### GetHopsOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetHopsOk() (*[]V1DiagnosticPingPauseResumePost204ResponseResultTraceResultHopsInner, bool)`

GetHopsOk returns a tuple with the Hops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHops

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) SetHops(v []V1DiagnosticPingPauseResumePost204ResponseResultTraceResultHopsInner)`

SetHops sets Hops field to given value.

### HasHops

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) HasHops() bool`

HasHops returns a boolean if a field has been set.

### GetMaxLatency

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetMaxLatency() float64`

GetMaxLatency returns the MaxLatency field if non-nil, zero value otherwise.

### GetMaxLatencyOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetMaxLatencyOk() (*float64, bool)`

GetMaxLatencyOk returns a tuple with the MaxLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLatency

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) SetMaxLatency(v float64)`

SetMaxLatency sets MaxLatency field to given value.

### HasMaxLatency

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) HasMaxLatency() bool`

HasMaxLatency returns a boolean if a field has been set.

### GetMaxLatencyHost

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetMaxLatencyHost() string`

GetMaxLatencyHost returns the MaxLatencyHost field if non-nil, zero value otherwise.

### GetMaxLatencyHostOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetMaxLatencyHostOk() (*string, bool)`

GetMaxLatencyHostOk returns a tuple with the MaxLatencyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLatencyHost

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) SetMaxLatencyHost(v string)`

SetMaxLatencyHost sets MaxLatencyHost field to given value.

### HasMaxLatencyHost

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) HasMaxLatencyHost() bool`

HasMaxLatencyHost returns a boolean if a field has been set.

### GetMaxPathMtu

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetMaxPathMtu() int32`

GetMaxPathMtu returns the MaxPathMtu field if non-nil, zero value otherwise.

### GetMaxPathMtuOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetMaxPathMtuOk() (*int32, bool)`

GetMaxPathMtuOk returns a tuple with the MaxPathMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPathMtu

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) SetMaxPathMtu(v int32)`

SetMaxPathMtu sets MaxPathMtu field to given value.

### HasMaxPathMtu

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) HasMaxPathMtu() bool`

HasMaxPathMtu returns a boolean if a field has been set.

### GetMinPathMtu

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetMinPathMtu() int32`

GetMinPathMtu returns the MinPathMtu field if non-nil, zero value otherwise.

### GetMinPathMtuOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetMinPathMtuOk() (*int32, bool)`

GetMinPathMtuOk returns a tuple with the MinPathMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPathMtu

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) SetMinPathMtu(v int32)`

SetMinPathMtu sets MinPathMtu field to given value.

### HasMinPathMtu

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) HasMinPathMtu() bool`

HasMinPathMtu returns a boolean if a field has been set.

### GetResult

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStoppedTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetStoppedTime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetStoppedTime returns the StoppedTime field if non-nil, zero value otherwise.

### GetStoppedTimeOk

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) GetStoppedTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetStoppedTimeOk returns a tuple with the StoppedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) SetStoppedTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetStoppedTime sets StoppedTime field to given value.

### HasStoppedTime

`func (o *V1DiagnosticPingPauseResumePost204ResponseResultTraceResult) HasStoppedTime() bool`

HasStoppedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


