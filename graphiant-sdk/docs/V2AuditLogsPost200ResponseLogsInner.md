# V2AuditLogsPost200ResponseLogsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ActivityId** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to [**V1ActivityLogsPostRequestSelectorJobEntity**](V1ActivityLogsPostRequestSelectorJobEntity.md) |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**V1ActivityLogsPostRequestSelectorJobEntity**](V1ActivityLogsPostRequestSelectorJobEntity.md) |  | [optional] 
**Ts** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AuditLogsPost200ResponseLogsInner

`func NewV2AuditLogsPost200ResponseLogsInner() *V2AuditLogsPost200ResponseLogsInner`

NewV2AuditLogsPost200ResponseLogsInner instantiates a new V2AuditLogsPost200ResponseLogsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AuditLogsPost200ResponseLogsInnerWithDefaults

`func NewV2AuditLogsPost200ResponseLogsInnerWithDefaults() *V2AuditLogsPost200ResponseLogsInner`

NewV2AuditLogsPost200ResponseLogsInnerWithDefaults instantiates a new V2AuditLogsPost200ResponseLogsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *V2AuditLogsPost200ResponseLogsInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V2AuditLogsPost200ResponseLogsInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V2AuditLogsPost200ResponseLogsInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *V2AuditLogsPost200ResponseLogsInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActivityId

`func (o *V2AuditLogsPost200ResponseLogsInner) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *V2AuditLogsPost200ResponseLogsInner) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *V2AuditLogsPost200ResponseLogsInner) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *V2AuditLogsPost200ResponseLogsInner) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetCategory

`func (o *V2AuditLogsPost200ResponseLogsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *V2AuditLogsPost200ResponseLogsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *V2AuditLogsPost200ResponseLogsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *V2AuditLogsPost200ResponseLogsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEntity

`func (o *V2AuditLogsPost200ResponseLogsInner) GetEntity() V1ActivityLogsPostRequestSelectorJobEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *V2AuditLogsPost200ResponseLogsInner) GetEntityOk() (*V1ActivityLogsPostRequestSelectorJobEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *V2AuditLogsPost200ResponseLogsInner) SetEntity(v V1ActivityLogsPostRequestSelectorJobEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *V2AuditLogsPost200ResponseLogsInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetJobType

`func (o *V2AuditLogsPost200ResponseLogsInner) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *V2AuditLogsPost200ResponseLogsInner) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *V2AuditLogsPost200ResponseLogsInner) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *V2AuditLogsPost200ResponseLogsInner) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetTarget

`func (o *V2AuditLogsPost200ResponseLogsInner) GetTarget() V1ActivityLogsPostRequestSelectorJobEntity`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *V2AuditLogsPost200ResponseLogsInner) GetTargetOk() (*V1ActivityLogsPostRequestSelectorJobEntity, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *V2AuditLogsPost200ResponseLogsInner) SetTarget(v V1ActivityLogsPostRequestSelectorJobEntity)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *V2AuditLogsPost200ResponseLogsInner) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTs

`func (o *V2AuditLogsPost200ResponseLogsInner) GetTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *V2AuditLogsPost200ResponseLogsInner) GetTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *V2AuditLogsPost200ResponseLogsInner) SetTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetTs sets Ts field to given value.

### HasTs

`func (o *V2AuditLogsPost200ResponseLogsInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetUser

`func (o *V2AuditLogsPost200ResponseLogsInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V2AuditLogsPost200ResponseLogsInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V2AuditLogsPost200ResponseLogsInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V2AuditLogsPost200ResponseLogsInner) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


