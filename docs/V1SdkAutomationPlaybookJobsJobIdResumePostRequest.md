# V1SdkAutomationPlaybookJobsJobIdResumePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDeploy** | Pointer to **bool** | When true, resume straight into deploy; otherwise resume as a dry-run | [optional] 
**VerboseLogs** | Pointer to **bool** | When true, request ansible -vvv for the resumed phase | [optional] 

## Methods

### NewV1SdkAutomationPlaybookJobsJobIdResumePostRequest

`func NewV1SdkAutomationPlaybookJobsJobIdResumePostRequest() *V1SdkAutomationPlaybookJobsJobIdResumePostRequest`

NewV1SdkAutomationPlaybookJobsJobIdResumePostRequest instantiates a new V1SdkAutomationPlaybookJobsJobIdResumePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SdkAutomationPlaybookJobsJobIdResumePostRequestWithDefaults

`func NewV1SdkAutomationPlaybookJobsJobIdResumePostRequestWithDefaults() *V1SdkAutomationPlaybookJobsJobIdResumePostRequest`

NewV1SdkAutomationPlaybookJobsJobIdResumePostRequestWithDefaults instantiates a new V1SdkAutomationPlaybookJobsJobIdResumePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDeploy

`func (o *V1SdkAutomationPlaybookJobsJobIdResumePostRequest) GetIsDeploy() bool`

GetIsDeploy returns the IsDeploy field if non-nil, zero value otherwise.

### GetIsDeployOk

`func (o *V1SdkAutomationPlaybookJobsJobIdResumePostRequest) GetIsDeployOk() (*bool, bool)`

GetIsDeployOk returns a tuple with the IsDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeploy

`func (o *V1SdkAutomationPlaybookJobsJobIdResumePostRequest) SetIsDeploy(v bool)`

SetIsDeploy sets IsDeploy field to given value.

### HasIsDeploy

`func (o *V1SdkAutomationPlaybookJobsJobIdResumePostRequest) HasIsDeploy() bool`

HasIsDeploy returns a boolean if a field has been set.

### GetVerboseLogs

`func (o *V1SdkAutomationPlaybookJobsJobIdResumePostRequest) GetVerboseLogs() bool`

GetVerboseLogs returns the VerboseLogs field if non-nil, zero value otherwise.

### GetVerboseLogsOk

`func (o *V1SdkAutomationPlaybookJobsJobIdResumePostRequest) GetVerboseLogsOk() (*bool, bool)`

GetVerboseLogsOk returns a tuple with the VerboseLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseLogs

`func (o *V1SdkAutomationPlaybookJobsJobIdResumePostRequest) SetVerboseLogs(v bool)`

SetVerboseLogs sets VerboseLogs field to given value.

### HasVerboseLogs

`func (o *V1SdkAutomationPlaybookJobsJobIdResumePostRequest) HasVerboseLogs() bool`

HasVerboseLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


