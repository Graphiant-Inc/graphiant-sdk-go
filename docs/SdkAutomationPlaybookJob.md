# SdkAutomationPlaybookJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionVersion** | Pointer to **string** | graphiant.naas collection version used for this run | [optional] 
**ConfigDescription** | Pointer to **string** | Config description / notes frozen at job creation (from config_snapshot); empty for older jobs | [optional] 
**ConfigId** | Pointer to **string** | Owning playbook config id | [optional] 
**ConfigName** | Pointer to **string** | Config display name frozen at job creation (from config_snapshot); empty for older jobs | [optional] 
**DeployEndedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**DeployStartedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**DryRunEndedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**DryRunStartedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**FailedPhase** | Pointer to **string** | When status is FAILED, which phase failed; unset when not failed | [optional] 
**JobId** | Pointer to **string** | Unique id of this playbook job run | [optional] 
**LogsAvailable** | Pointer to **bool** | True when ansible logs are available for this run | [optional] 
**PostDeployCheckStartedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**RunDurationMs** | Pointer to **int64** | Wall-clock duration in milliseconds from dry_run_started_at to run_ended_at; unset until run_ended_at is set | [optional] 
**RunEndedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**RunNumber** | Pointer to **int32** | Per-config run counter (1-based for started runs) | [optional] 
**RunnerJobId** | Pointer to **string** | External ansible runner job id, if the backend assigned one | [optional] 
**SdkVersion** | Pointer to **string** | SDK version used for this run | [optional] 
**StartedByUserId** | Pointer to **string** | IAM UserInfo.user_id of who started this run; UI resolves the display name | [optional] 
**Status** | Pointer to **string** | Gated pipeline state (dry-run / deploy / post-deploy-check) | [optional] 
**VerboseLogs** | Pointer to **bool** | Whether ansible -vvv was requested for the latest phase of this run | [optional] 

## Methods

### NewSdkAutomationPlaybookJob

`func NewSdkAutomationPlaybookJob() *SdkAutomationPlaybookJob`

NewSdkAutomationPlaybookJob instantiates a new SdkAutomationPlaybookJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkAutomationPlaybookJobWithDefaults

`func NewSdkAutomationPlaybookJobWithDefaults() *SdkAutomationPlaybookJob`

NewSdkAutomationPlaybookJobWithDefaults instantiates a new SdkAutomationPlaybookJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionVersion

`func (o *SdkAutomationPlaybookJob) GetCollectionVersion() string`

GetCollectionVersion returns the CollectionVersion field if non-nil, zero value otherwise.

### GetCollectionVersionOk

`func (o *SdkAutomationPlaybookJob) GetCollectionVersionOk() (*string, bool)`

GetCollectionVersionOk returns a tuple with the CollectionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionVersion

`func (o *SdkAutomationPlaybookJob) SetCollectionVersion(v string)`

SetCollectionVersion sets CollectionVersion field to given value.

### HasCollectionVersion

`func (o *SdkAutomationPlaybookJob) HasCollectionVersion() bool`

HasCollectionVersion returns a boolean if a field has been set.

### GetConfigDescription

`func (o *SdkAutomationPlaybookJob) GetConfigDescription() string`

GetConfigDescription returns the ConfigDescription field if non-nil, zero value otherwise.

### GetConfigDescriptionOk

`func (o *SdkAutomationPlaybookJob) GetConfigDescriptionOk() (*string, bool)`

GetConfigDescriptionOk returns a tuple with the ConfigDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDescription

`func (o *SdkAutomationPlaybookJob) SetConfigDescription(v string)`

SetConfigDescription sets ConfigDescription field to given value.

### HasConfigDescription

`func (o *SdkAutomationPlaybookJob) HasConfigDescription() bool`

HasConfigDescription returns a boolean if a field has been set.

### GetConfigId

`func (o *SdkAutomationPlaybookJob) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *SdkAutomationPlaybookJob) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *SdkAutomationPlaybookJob) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *SdkAutomationPlaybookJob) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetConfigName

`func (o *SdkAutomationPlaybookJob) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *SdkAutomationPlaybookJob) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *SdkAutomationPlaybookJob) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *SdkAutomationPlaybookJob) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetDeployEndedAt

`func (o *SdkAutomationPlaybookJob) GetDeployEndedAt() GoogleProtobufTimestamp`

GetDeployEndedAt returns the DeployEndedAt field if non-nil, zero value otherwise.

### GetDeployEndedAtOk

`func (o *SdkAutomationPlaybookJob) GetDeployEndedAtOk() (*GoogleProtobufTimestamp, bool)`

GetDeployEndedAtOk returns a tuple with the DeployEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployEndedAt

`func (o *SdkAutomationPlaybookJob) SetDeployEndedAt(v GoogleProtobufTimestamp)`

SetDeployEndedAt sets DeployEndedAt field to given value.

### HasDeployEndedAt

`func (o *SdkAutomationPlaybookJob) HasDeployEndedAt() bool`

HasDeployEndedAt returns a boolean if a field has been set.

### GetDeployStartedAt

`func (o *SdkAutomationPlaybookJob) GetDeployStartedAt() GoogleProtobufTimestamp`

GetDeployStartedAt returns the DeployStartedAt field if non-nil, zero value otherwise.

### GetDeployStartedAtOk

`func (o *SdkAutomationPlaybookJob) GetDeployStartedAtOk() (*GoogleProtobufTimestamp, bool)`

GetDeployStartedAtOk returns a tuple with the DeployStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStartedAt

`func (o *SdkAutomationPlaybookJob) SetDeployStartedAt(v GoogleProtobufTimestamp)`

SetDeployStartedAt sets DeployStartedAt field to given value.

### HasDeployStartedAt

`func (o *SdkAutomationPlaybookJob) HasDeployStartedAt() bool`

HasDeployStartedAt returns a boolean if a field has been set.

### GetDryRunEndedAt

`func (o *SdkAutomationPlaybookJob) GetDryRunEndedAt() GoogleProtobufTimestamp`

GetDryRunEndedAt returns the DryRunEndedAt field if non-nil, zero value otherwise.

### GetDryRunEndedAtOk

`func (o *SdkAutomationPlaybookJob) GetDryRunEndedAtOk() (*GoogleProtobufTimestamp, bool)`

GetDryRunEndedAtOk returns a tuple with the DryRunEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRunEndedAt

`func (o *SdkAutomationPlaybookJob) SetDryRunEndedAt(v GoogleProtobufTimestamp)`

SetDryRunEndedAt sets DryRunEndedAt field to given value.

### HasDryRunEndedAt

`func (o *SdkAutomationPlaybookJob) HasDryRunEndedAt() bool`

HasDryRunEndedAt returns a boolean if a field has been set.

### GetDryRunStartedAt

`func (o *SdkAutomationPlaybookJob) GetDryRunStartedAt() GoogleProtobufTimestamp`

GetDryRunStartedAt returns the DryRunStartedAt field if non-nil, zero value otherwise.

### GetDryRunStartedAtOk

`func (o *SdkAutomationPlaybookJob) GetDryRunStartedAtOk() (*GoogleProtobufTimestamp, bool)`

GetDryRunStartedAtOk returns a tuple with the DryRunStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRunStartedAt

`func (o *SdkAutomationPlaybookJob) SetDryRunStartedAt(v GoogleProtobufTimestamp)`

SetDryRunStartedAt sets DryRunStartedAt field to given value.

### HasDryRunStartedAt

`func (o *SdkAutomationPlaybookJob) HasDryRunStartedAt() bool`

HasDryRunStartedAt returns a boolean if a field has been set.

### GetFailedPhase

`func (o *SdkAutomationPlaybookJob) GetFailedPhase() string`

GetFailedPhase returns the FailedPhase field if non-nil, zero value otherwise.

### GetFailedPhaseOk

`func (o *SdkAutomationPlaybookJob) GetFailedPhaseOk() (*string, bool)`

GetFailedPhaseOk returns a tuple with the FailedPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedPhase

`func (o *SdkAutomationPlaybookJob) SetFailedPhase(v string)`

SetFailedPhase sets FailedPhase field to given value.

### HasFailedPhase

`func (o *SdkAutomationPlaybookJob) HasFailedPhase() bool`

HasFailedPhase returns a boolean if a field has been set.

### GetJobId

`func (o *SdkAutomationPlaybookJob) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SdkAutomationPlaybookJob) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SdkAutomationPlaybookJob) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *SdkAutomationPlaybookJob) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetLogsAvailable

`func (o *SdkAutomationPlaybookJob) GetLogsAvailable() bool`

GetLogsAvailable returns the LogsAvailable field if non-nil, zero value otherwise.

### GetLogsAvailableOk

`func (o *SdkAutomationPlaybookJob) GetLogsAvailableOk() (*bool, bool)`

GetLogsAvailableOk returns a tuple with the LogsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsAvailable

`func (o *SdkAutomationPlaybookJob) SetLogsAvailable(v bool)`

SetLogsAvailable sets LogsAvailable field to given value.

### HasLogsAvailable

`func (o *SdkAutomationPlaybookJob) HasLogsAvailable() bool`

HasLogsAvailable returns a boolean if a field has been set.

### GetPostDeployCheckStartedAt

`func (o *SdkAutomationPlaybookJob) GetPostDeployCheckStartedAt() GoogleProtobufTimestamp`

GetPostDeployCheckStartedAt returns the PostDeployCheckStartedAt field if non-nil, zero value otherwise.

### GetPostDeployCheckStartedAtOk

`func (o *SdkAutomationPlaybookJob) GetPostDeployCheckStartedAtOk() (*GoogleProtobufTimestamp, bool)`

GetPostDeployCheckStartedAtOk returns a tuple with the PostDeployCheckStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostDeployCheckStartedAt

`func (o *SdkAutomationPlaybookJob) SetPostDeployCheckStartedAt(v GoogleProtobufTimestamp)`

SetPostDeployCheckStartedAt sets PostDeployCheckStartedAt field to given value.

### HasPostDeployCheckStartedAt

`func (o *SdkAutomationPlaybookJob) HasPostDeployCheckStartedAt() bool`

HasPostDeployCheckStartedAt returns a boolean if a field has been set.

### GetRunDurationMs

`func (o *SdkAutomationPlaybookJob) GetRunDurationMs() int64`

GetRunDurationMs returns the RunDurationMs field if non-nil, zero value otherwise.

### GetRunDurationMsOk

`func (o *SdkAutomationPlaybookJob) GetRunDurationMsOk() (*int64, bool)`

GetRunDurationMsOk returns a tuple with the RunDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunDurationMs

`func (o *SdkAutomationPlaybookJob) SetRunDurationMs(v int64)`

SetRunDurationMs sets RunDurationMs field to given value.

### HasRunDurationMs

`func (o *SdkAutomationPlaybookJob) HasRunDurationMs() bool`

HasRunDurationMs returns a boolean if a field has been set.

### GetRunEndedAt

`func (o *SdkAutomationPlaybookJob) GetRunEndedAt() GoogleProtobufTimestamp`

GetRunEndedAt returns the RunEndedAt field if non-nil, zero value otherwise.

### GetRunEndedAtOk

`func (o *SdkAutomationPlaybookJob) GetRunEndedAtOk() (*GoogleProtobufTimestamp, bool)`

GetRunEndedAtOk returns a tuple with the RunEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunEndedAt

`func (o *SdkAutomationPlaybookJob) SetRunEndedAt(v GoogleProtobufTimestamp)`

SetRunEndedAt sets RunEndedAt field to given value.

### HasRunEndedAt

`func (o *SdkAutomationPlaybookJob) HasRunEndedAt() bool`

HasRunEndedAt returns a boolean if a field has been set.

### GetRunNumber

`func (o *SdkAutomationPlaybookJob) GetRunNumber() int32`

GetRunNumber returns the RunNumber field if non-nil, zero value otherwise.

### GetRunNumberOk

`func (o *SdkAutomationPlaybookJob) GetRunNumberOk() (*int32, bool)`

GetRunNumberOk returns a tuple with the RunNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNumber

`func (o *SdkAutomationPlaybookJob) SetRunNumber(v int32)`

SetRunNumber sets RunNumber field to given value.

### HasRunNumber

`func (o *SdkAutomationPlaybookJob) HasRunNumber() bool`

HasRunNumber returns a boolean if a field has been set.

### GetRunnerJobId

`func (o *SdkAutomationPlaybookJob) GetRunnerJobId() string`

GetRunnerJobId returns the RunnerJobId field if non-nil, zero value otherwise.

### GetRunnerJobIdOk

`func (o *SdkAutomationPlaybookJob) GetRunnerJobIdOk() (*string, bool)`

GetRunnerJobIdOk returns a tuple with the RunnerJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerJobId

`func (o *SdkAutomationPlaybookJob) SetRunnerJobId(v string)`

SetRunnerJobId sets RunnerJobId field to given value.

### HasRunnerJobId

`func (o *SdkAutomationPlaybookJob) HasRunnerJobId() bool`

HasRunnerJobId returns a boolean if a field has been set.

### GetSdkVersion

`func (o *SdkAutomationPlaybookJob) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *SdkAutomationPlaybookJob) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *SdkAutomationPlaybookJob) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *SdkAutomationPlaybookJob) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.

### GetStartedByUserId

`func (o *SdkAutomationPlaybookJob) GetStartedByUserId() string`

GetStartedByUserId returns the StartedByUserId field if non-nil, zero value otherwise.

### GetStartedByUserIdOk

`func (o *SdkAutomationPlaybookJob) GetStartedByUserIdOk() (*string, bool)`

GetStartedByUserIdOk returns a tuple with the StartedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedByUserId

`func (o *SdkAutomationPlaybookJob) SetStartedByUserId(v string)`

SetStartedByUserId sets StartedByUserId field to given value.

### HasStartedByUserId

`func (o *SdkAutomationPlaybookJob) HasStartedByUserId() bool`

HasStartedByUserId returns a boolean if a field has been set.

### GetStatus

`func (o *SdkAutomationPlaybookJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SdkAutomationPlaybookJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SdkAutomationPlaybookJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SdkAutomationPlaybookJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerboseLogs

`func (o *SdkAutomationPlaybookJob) GetVerboseLogs() bool`

GetVerboseLogs returns the VerboseLogs field if non-nil, zero value otherwise.

### GetVerboseLogsOk

`func (o *SdkAutomationPlaybookJob) GetVerboseLogsOk() (*bool, bool)`

GetVerboseLogsOk returns a tuple with the VerboseLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseLogs

`func (o *SdkAutomationPlaybookJob) SetVerboseLogs(v bool)`

SetVerboseLogs sets VerboseLogs field to given value.

### HasVerboseLogs

`func (o *SdkAutomationPlaybookJob) HasVerboseLogs() bool`

HasVerboseLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


