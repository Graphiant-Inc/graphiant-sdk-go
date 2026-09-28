# V1SdkAutomationPlaybookJobsJobIdRunPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipDryRun** | Pointer to **bool** | When true, skip dry-run and gate and deploy directly | [optional] 
**VerboseLogs** | Pointer to **bool** | When true, request ansible -vvv for the new run | [optional] 

## Methods

### NewV1SdkAutomationPlaybookJobsJobIdRunPostRequest

`func NewV1SdkAutomationPlaybookJobsJobIdRunPostRequest() *V1SdkAutomationPlaybookJobsJobIdRunPostRequest`

NewV1SdkAutomationPlaybookJobsJobIdRunPostRequest instantiates a new V1SdkAutomationPlaybookJobsJobIdRunPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SdkAutomationPlaybookJobsJobIdRunPostRequestWithDefaults

`func NewV1SdkAutomationPlaybookJobsJobIdRunPostRequestWithDefaults() *V1SdkAutomationPlaybookJobsJobIdRunPostRequest`

NewV1SdkAutomationPlaybookJobsJobIdRunPostRequestWithDefaults instantiates a new V1SdkAutomationPlaybookJobsJobIdRunPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipDryRun

`func (o *V1SdkAutomationPlaybookJobsJobIdRunPostRequest) GetSkipDryRun() bool`

GetSkipDryRun returns the SkipDryRun field if non-nil, zero value otherwise.

### GetSkipDryRunOk

`func (o *V1SdkAutomationPlaybookJobsJobIdRunPostRequest) GetSkipDryRunOk() (*bool, bool)`

GetSkipDryRunOk returns a tuple with the SkipDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDryRun

`func (o *V1SdkAutomationPlaybookJobsJobIdRunPostRequest) SetSkipDryRun(v bool)`

SetSkipDryRun sets SkipDryRun field to given value.

### HasSkipDryRun

`func (o *V1SdkAutomationPlaybookJobsJobIdRunPostRequest) HasSkipDryRun() bool`

HasSkipDryRun returns a boolean if a field has been set.

### GetVerboseLogs

`func (o *V1SdkAutomationPlaybookJobsJobIdRunPostRequest) GetVerboseLogs() bool`

GetVerboseLogs returns the VerboseLogs field if non-nil, zero value otherwise.

### GetVerboseLogsOk

`func (o *V1SdkAutomationPlaybookJobsJobIdRunPostRequest) GetVerboseLogsOk() (*bool, bool)`

GetVerboseLogsOk returns a tuple with the VerboseLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseLogs

`func (o *V1SdkAutomationPlaybookJobsJobIdRunPostRequest) SetVerboseLogs(v bool)`

SetVerboseLogs sets VerboseLogs field to given value.

### HasVerboseLogs

`func (o *V1SdkAutomationPlaybookJobsJobIdRunPostRequest) HasVerboseLogs() bool`

HasVerboseLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


