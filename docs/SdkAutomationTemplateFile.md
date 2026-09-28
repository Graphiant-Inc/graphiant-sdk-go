# SdkAutomationTemplateFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | File contents, verbatim from the collection | [optional] 
**Filename** | Pointer to **string** | Suggested download name (e.g. sample_interface_config.yaml) | [optional] 
**ModuleKey** | Pointer to **string** | Catalog module key; empty for the bundle-level playbook not scoped to a slot | [optional] 
**SourcePath** | Pointer to **string** | Path within the collection for provenance (e.g. playbooks/interface_management.yml) | [optional] 

## Methods

### NewSdkAutomationTemplateFile

`func NewSdkAutomationTemplateFile() *SdkAutomationTemplateFile`

NewSdkAutomationTemplateFile instantiates a new SdkAutomationTemplateFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkAutomationTemplateFileWithDefaults

`func NewSdkAutomationTemplateFileWithDefaults() *SdkAutomationTemplateFile`

NewSdkAutomationTemplateFileWithDefaults instantiates a new SdkAutomationTemplateFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SdkAutomationTemplateFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SdkAutomationTemplateFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SdkAutomationTemplateFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SdkAutomationTemplateFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFilename

`func (o *SdkAutomationTemplateFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SdkAutomationTemplateFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SdkAutomationTemplateFile) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SdkAutomationTemplateFile) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetModuleKey

`func (o *SdkAutomationTemplateFile) GetModuleKey() string`

GetModuleKey returns the ModuleKey field if non-nil, zero value otherwise.

### GetModuleKeyOk

`func (o *SdkAutomationTemplateFile) GetModuleKeyOk() (*string, bool)`

GetModuleKeyOk returns a tuple with the ModuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleKey

`func (o *SdkAutomationTemplateFile) SetModuleKey(v string)`

SetModuleKey sets ModuleKey field to given value.

### HasModuleKey

`func (o *SdkAutomationTemplateFile) HasModuleKey() bool`

HasModuleKey returns a boolean if a field has been set.

### GetSourcePath

`func (o *SdkAutomationTemplateFile) GetSourcePath() string`

GetSourcePath returns the SourcePath field if non-nil, zero value otherwise.

### GetSourcePathOk

`func (o *SdkAutomationTemplateFile) GetSourcePathOk() (*string, bool)`

GetSourcePathOk returns a tuple with the SourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePath

`func (o *SdkAutomationTemplateFile) SetSourcePath(v string)`

SetSourcePath sets SourcePath field to given value.

### HasSourcePath

`func (o *SdkAutomationTemplateFile) HasSourcePath() bool`

HasSourcePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


