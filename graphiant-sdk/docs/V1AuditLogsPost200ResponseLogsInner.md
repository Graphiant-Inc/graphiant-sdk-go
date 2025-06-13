# V1AuditLogsPost200ResponseLogsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **string** |  | [optional] 
**Actor** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**End** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**FailedTargetResults** | Pointer to [**[]V1AuditLogsPost200ResponseLogsInnerFailedTargetResultsInner**](V1AuditLogsPost200ResponseLogsInnerFailedTargetResultsInner.md) |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Start** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**[]V1AuditLogsPost200ResponseLogsInnerFailedTargetResultsInner**](V1AuditLogsPost200ResponseLogsInnerFailedTargetResultsInner.md) |  | [optional] 

## Methods

### NewV1AuditLogsPost200ResponseLogsInner

`func NewV1AuditLogsPost200ResponseLogsInner() *V1AuditLogsPost200ResponseLogsInner`

NewV1AuditLogsPost200ResponseLogsInner instantiates a new V1AuditLogsPost200ResponseLogsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AuditLogsPost200ResponseLogsInnerWithDefaults

`func NewV1AuditLogsPost200ResponseLogsInnerWithDefaults() *V1AuditLogsPost200ResponseLogsInner`

NewV1AuditLogsPost200ResponseLogsInnerWithDefaults instantiates a new V1AuditLogsPost200ResponseLogsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *V1AuditLogsPost200ResponseLogsInner) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *V1AuditLogsPost200ResponseLogsInner) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *V1AuditLogsPost200ResponseLogsInner) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetActor

`func (o *V1AuditLogsPost200ResponseLogsInner) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *V1AuditLogsPost200ResponseLogsInner) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *V1AuditLogsPost200ResponseLogsInner) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCategory

`func (o *V1AuditLogsPost200ResponseLogsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *V1AuditLogsPost200ResponseLogsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *V1AuditLogsPost200ResponseLogsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnd

`func (o *V1AuditLogsPost200ResponseLogsInner) GetEnd() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetEndOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V1AuditLogsPost200ResponseLogsInner) SetEnd(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V1AuditLogsPost200ResponseLogsInner) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFailedTargetResults

`func (o *V1AuditLogsPost200ResponseLogsInner) GetFailedTargetResults() []V1AuditLogsPost200ResponseLogsInnerFailedTargetResultsInner`

GetFailedTargetResults returns the FailedTargetResults field if non-nil, zero value otherwise.

### GetFailedTargetResultsOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetFailedTargetResultsOk() (*[]V1AuditLogsPost200ResponseLogsInnerFailedTargetResultsInner, bool)`

GetFailedTargetResultsOk returns a tuple with the FailedTargetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTargetResults

`func (o *V1AuditLogsPost200ResponseLogsInner) SetFailedTargetResults(v []V1AuditLogsPost200ResponseLogsInnerFailedTargetResultsInner)`

SetFailedTargetResults sets FailedTargetResults field to given value.

### HasFailedTargetResults

`func (o *V1AuditLogsPost200ResponseLogsInner) HasFailedTargetResults() bool`

HasFailedTargetResults returns a boolean if a field has been set.

### GetInfo

`func (o *V1AuditLogsPost200ResponseLogsInner) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V1AuditLogsPost200ResponseLogsInner) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V1AuditLogsPost200ResponseLogsInner) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetReason

`func (o *V1AuditLogsPost200ResponseLogsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1AuditLogsPost200ResponseLogsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1AuditLogsPost200ResponseLogsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetService

`func (o *V1AuditLogsPost200ResponseLogsInner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1AuditLogsPost200ResponseLogsInner) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *V1AuditLogsPost200ResponseLogsInner) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStart

`func (o *V1AuditLogsPost200ResponseLogsInner) GetStart() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetStartOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V1AuditLogsPost200ResponseLogsInner) SetStart(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetStart sets Start field to given value.

### HasStart

`func (o *V1AuditLogsPost200ResponseLogsInner) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *V1AuditLogsPost200ResponseLogsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1AuditLogsPost200ResponseLogsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1AuditLogsPost200ResponseLogsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargets

`func (o *V1AuditLogsPost200ResponseLogsInner) GetTargets() []V1AuditLogsPost200ResponseLogsInnerFailedTargetResultsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *V1AuditLogsPost200ResponseLogsInner) GetTargetsOk() (*[]V1AuditLogsPost200ResponseLogsInnerFailedTargetResultsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *V1AuditLogsPost200ResponseLogsInner) SetTargets(v []V1AuditLogsPost200ResponseLogsInnerFailedTargetResultsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *V1AuditLogsPost200ResponseLogsInner) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


