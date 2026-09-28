# SdkAutomationPlaybookConfigSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **string** | Config id | [optional] 
**CreatedByUserId** | Pointer to **string** | IAM user id of the config creator; UI resolves the display name | [optional] 
**JobStatus** | Pointer to **string** | Latest run status; unset when the config has no run yet | [optional] 
**Name** | Pointer to **string** | Display name | [optional] 
**RunId** | Pointer to **string** | Latest job id for this config; empty if never run | [optional] 
**RunNumber** | Pointer to **int32** | Per-config run counter; 0 if never run | [optional] 
**StartedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**StartedByUserId** | Pointer to **string** | IAM user id of who launched the latest run | [optional] 
**Status** | Pointer to **string** | Config lifecycle status | [optional] 
**UpdatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 

## Methods

### NewSdkAutomationPlaybookConfigSummary

`func NewSdkAutomationPlaybookConfigSummary() *SdkAutomationPlaybookConfigSummary`

NewSdkAutomationPlaybookConfigSummary instantiates a new SdkAutomationPlaybookConfigSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkAutomationPlaybookConfigSummaryWithDefaults

`func NewSdkAutomationPlaybookConfigSummaryWithDefaults() *SdkAutomationPlaybookConfigSummary`

NewSdkAutomationPlaybookConfigSummaryWithDefaults instantiates a new SdkAutomationPlaybookConfigSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *SdkAutomationPlaybookConfigSummary) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *SdkAutomationPlaybookConfigSummary) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *SdkAutomationPlaybookConfigSummary) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *SdkAutomationPlaybookConfigSummary) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *SdkAutomationPlaybookConfigSummary) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *SdkAutomationPlaybookConfigSummary) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *SdkAutomationPlaybookConfigSummary) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *SdkAutomationPlaybookConfigSummary) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetJobStatus

`func (o *SdkAutomationPlaybookConfigSummary) GetJobStatus() string`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *SdkAutomationPlaybookConfigSummary) GetJobStatusOk() (*string, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *SdkAutomationPlaybookConfigSummary) SetJobStatus(v string)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *SdkAutomationPlaybookConfigSummary) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetName

`func (o *SdkAutomationPlaybookConfigSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkAutomationPlaybookConfigSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkAutomationPlaybookConfigSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdkAutomationPlaybookConfigSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRunId

`func (o *SdkAutomationPlaybookConfigSummary) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *SdkAutomationPlaybookConfigSummary) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *SdkAutomationPlaybookConfigSummary) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *SdkAutomationPlaybookConfigSummary) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetRunNumber

`func (o *SdkAutomationPlaybookConfigSummary) GetRunNumber() int32`

GetRunNumber returns the RunNumber field if non-nil, zero value otherwise.

### GetRunNumberOk

`func (o *SdkAutomationPlaybookConfigSummary) GetRunNumberOk() (*int32, bool)`

GetRunNumberOk returns a tuple with the RunNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNumber

`func (o *SdkAutomationPlaybookConfigSummary) SetRunNumber(v int32)`

SetRunNumber sets RunNumber field to given value.

### HasRunNumber

`func (o *SdkAutomationPlaybookConfigSummary) HasRunNumber() bool`

HasRunNumber returns a boolean if a field has been set.

### GetStartedAt

`func (o *SdkAutomationPlaybookConfigSummary) GetStartedAt() GoogleProtobufTimestamp`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SdkAutomationPlaybookConfigSummary) GetStartedAtOk() (*GoogleProtobufTimestamp, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SdkAutomationPlaybookConfigSummary) SetStartedAt(v GoogleProtobufTimestamp)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SdkAutomationPlaybookConfigSummary) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStartedByUserId

`func (o *SdkAutomationPlaybookConfigSummary) GetStartedByUserId() string`

GetStartedByUserId returns the StartedByUserId field if non-nil, zero value otherwise.

### GetStartedByUserIdOk

`func (o *SdkAutomationPlaybookConfigSummary) GetStartedByUserIdOk() (*string, bool)`

GetStartedByUserIdOk returns a tuple with the StartedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedByUserId

`func (o *SdkAutomationPlaybookConfigSummary) SetStartedByUserId(v string)`

SetStartedByUserId sets StartedByUserId field to given value.

### HasStartedByUserId

`func (o *SdkAutomationPlaybookConfigSummary) HasStartedByUserId() bool`

HasStartedByUserId returns a boolean if a field has been set.

### GetStatus

`func (o *SdkAutomationPlaybookConfigSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SdkAutomationPlaybookConfigSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SdkAutomationPlaybookConfigSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SdkAutomationPlaybookConfigSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SdkAutomationPlaybookConfigSummary) GetUpdatedAt() GoogleProtobufTimestamp`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SdkAutomationPlaybookConfigSummary) GetUpdatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SdkAutomationPlaybookConfigSummary) SetUpdatedAt(v GoogleProtobufTimestamp)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SdkAutomationPlaybookConfigSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


