# V1ActivityLogsPost200ResponseDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]V1ActivityLogsPostRequestSelectorJobEntity**](V1ActivityLogsPostRequestSelectorJobEntity.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**DisableAutoTimeout** | Pointer to **bool** |  | [optional] 
**EndTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InitiatorType** | Pointer to **string** |  | [optional] 
**JobEntities** | Pointer to [**[]V1ActivityLogsPostRequestSelectorJobEntity**](V1ActivityLogsPostRequestSelectorJobEntity.md) |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**OriginalEnterpriseId** | Pointer to **int64** |  | [optional] 
**StartTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**[]V1ActivityLogsPost200ResponseDetailsInnerTargetsInner**](V1ActivityLogsPost200ResponseDetailsInnerTargetsInner.md) |  | [optional] 
**TraceSessionId** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ActivityLogsPost200ResponseDetailsInner

`func NewV1ActivityLogsPost200ResponseDetailsInner() *V1ActivityLogsPost200ResponseDetailsInner`

NewV1ActivityLogsPost200ResponseDetailsInner instantiates a new V1ActivityLogsPost200ResponseDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ActivityLogsPost200ResponseDetailsInnerWithDefaults

`func NewV1ActivityLogsPost200ResponseDetailsInnerWithDefaults() *V1ActivityLogsPost200ResponseDetailsInner`

NewV1ActivityLogsPost200ResponseDetailsInnerWithDefaults instantiates a new V1ActivityLogsPost200ResponseDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAttributes

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetAttributes() []V1ActivityLogsPostRequestSelectorJobEntity`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetAttributesOk() (*[]V1ActivityLogsPostRequestSelectorJobEntity, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetAttributes(v []V1ActivityLogsPostRequestSelectorJobEntity)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCategory

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDisableAutoTimeout

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetDisableAutoTimeout() bool`

GetDisableAutoTimeout returns the DisableAutoTimeout field if non-nil, zero value otherwise.

### GetDisableAutoTimeoutOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetDisableAutoTimeoutOk() (*bool, bool)`

GetDisableAutoTimeoutOk returns a tuple with the DisableAutoTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoTimeout

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetDisableAutoTimeout(v bool)`

SetDisableAutoTimeout sets DisableAutoTimeout field to given value.

### HasDisableAutoTimeout

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasDisableAutoTimeout() bool`

HasDisableAutoTimeout returns a boolean if a field has been set.

### GetEndTs

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetEndTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetEndTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetEndTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitiatorType

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetInitiatorType() string`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetInitiatorTypeOk() (*string, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetInitiatorType(v string)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

### GetJobEntities

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetJobEntities() []V1ActivityLogsPostRequestSelectorJobEntity`

GetJobEntities returns the JobEntities field if non-nil, zero value otherwise.

### GetJobEntitiesOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetJobEntitiesOk() (*[]V1ActivityLogsPostRequestSelectorJobEntity, bool)`

GetJobEntitiesOk returns a tuple with the JobEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobEntities

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetJobEntities(v []V1ActivityLogsPostRequestSelectorJobEntity)`

SetJobEntities sets JobEntities field to given value.

### HasJobEntities

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasJobEntities() bool`

HasJobEntities returns a boolean if a field has been set.

### GetJobType

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetOriginalEnterpriseId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetOriginalEnterpriseId() int64`

GetOriginalEnterpriseId returns the OriginalEnterpriseId field if non-nil, zero value otherwise.

### GetOriginalEnterpriseIdOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetOriginalEnterpriseIdOk() (*int64, bool)`

GetOriginalEnterpriseIdOk returns a tuple with the OriginalEnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEnterpriseId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetOriginalEnterpriseId(v int64)`

SetOriginalEnterpriseId sets OriginalEnterpriseId field to given value.

### HasOriginalEnterpriseId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasOriginalEnterpriseId() bool`

HasOriginalEnterpriseId returns a boolean if a field has been set.

### GetStartTs

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetStartTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetStartTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetStartTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetStatus

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargets

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetTargets() []V1ActivityLogsPost200ResponseDetailsInnerTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetTargetsOk() (*[]V1ActivityLogsPost200ResponseDetailsInnerTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetTargets(v []V1ActivityLogsPost200ResponseDetailsInnerTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetTraceSessionId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetTraceSessionId() string`

GetTraceSessionId returns the TraceSessionId field if non-nil, zero value otherwise.

### GetTraceSessionIdOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetTraceSessionIdOk() (*string, bool)`

GetTraceSessionIdOk returns a tuple with the TraceSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceSessionId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetTraceSessionId(v string)`

SetTraceSessionId sets TraceSessionId field to given value.

### HasTraceSessionId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasTraceSessionId() bool`

HasTraceSessionId returns a boolean if a field has been set.

### GetUsage

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUser

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V1ActivityLogsPost200ResponseDetailsInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V1ActivityLogsPost200ResponseDetailsInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


