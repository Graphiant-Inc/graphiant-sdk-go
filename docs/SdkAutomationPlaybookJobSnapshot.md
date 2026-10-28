# SdkAutomationPlaybookJobSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleKey** | Pointer to **string** | Catalog bundle key frozen at job creation | [optional] 
**ConfigId** | Pointer to **string** | Config id that owned this run when the snapshot was taken | [optional] 
**Description** | Pointer to **string** | Config notes frozen at job creation; empty for older jobs | [optional] 
**Files** | Pointer to [**[]SdkAutomationModuleFile**](SdkAutomationModuleFile.md) |  | [optional] 
**JobId** | Pointer to **string** | Owning job id | [optional] 
**Name** | Pointer to **string** | Config display name frozen at job creation; empty for older jobs | [optional] 
**Playbook** | Pointer to **string** | Playbook filename frozen at job creation (e.g. network_bundle.yml); empty for older jobs | [optional] 
**Tags** | Pointer to **string** | Ansible --tags frozen at job creation (union of present modules&#39; associated_tags); empty for older jobs or when no tags were applied | [optional] 

## Methods

### NewSdkAutomationPlaybookJobSnapshot

`func NewSdkAutomationPlaybookJobSnapshot() *SdkAutomationPlaybookJobSnapshot`

NewSdkAutomationPlaybookJobSnapshot instantiates a new SdkAutomationPlaybookJobSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkAutomationPlaybookJobSnapshotWithDefaults

`func NewSdkAutomationPlaybookJobSnapshotWithDefaults() *SdkAutomationPlaybookJobSnapshot`

NewSdkAutomationPlaybookJobSnapshotWithDefaults instantiates a new SdkAutomationPlaybookJobSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleKey

`func (o *SdkAutomationPlaybookJobSnapshot) GetBundleKey() string`

GetBundleKey returns the BundleKey field if non-nil, zero value otherwise.

### GetBundleKeyOk

`func (o *SdkAutomationPlaybookJobSnapshot) GetBundleKeyOk() (*string, bool)`

GetBundleKeyOk returns a tuple with the BundleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleKey

`func (o *SdkAutomationPlaybookJobSnapshot) SetBundleKey(v string)`

SetBundleKey sets BundleKey field to given value.

### HasBundleKey

`func (o *SdkAutomationPlaybookJobSnapshot) HasBundleKey() bool`

HasBundleKey returns a boolean if a field has been set.

### GetConfigId

`func (o *SdkAutomationPlaybookJobSnapshot) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *SdkAutomationPlaybookJobSnapshot) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *SdkAutomationPlaybookJobSnapshot) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *SdkAutomationPlaybookJobSnapshot) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetDescription

`func (o *SdkAutomationPlaybookJobSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdkAutomationPlaybookJobSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdkAutomationPlaybookJobSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdkAutomationPlaybookJobSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFiles

`func (o *SdkAutomationPlaybookJobSnapshot) GetFiles() []SdkAutomationModuleFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SdkAutomationPlaybookJobSnapshot) GetFilesOk() (*[]SdkAutomationModuleFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SdkAutomationPlaybookJobSnapshot) SetFiles(v []SdkAutomationModuleFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SdkAutomationPlaybookJobSnapshot) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetJobId

`func (o *SdkAutomationPlaybookJobSnapshot) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SdkAutomationPlaybookJobSnapshot) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SdkAutomationPlaybookJobSnapshot) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *SdkAutomationPlaybookJobSnapshot) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetName

`func (o *SdkAutomationPlaybookJobSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkAutomationPlaybookJobSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkAutomationPlaybookJobSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdkAutomationPlaybookJobSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaybook

`func (o *SdkAutomationPlaybookJobSnapshot) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *SdkAutomationPlaybookJobSnapshot) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *SdkAutomationPlaybookJobSnapshot) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *SdkAutomationPlaybookJobSnapshot) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.

### GetTags

`func (o *SdkAutomationPlaybookJobSnapshot) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SdkAutomationPlaybookJobSnapshot) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SdkAutomationPlaybookJobSnapshot) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SdkAutomationPlaybookJobSnapshot) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


