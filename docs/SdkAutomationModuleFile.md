# SdkAutomationModuleFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | YAML file contents | [optional] 
**Filename** | Pointer to **string** | Original upload filename (informational; not used for routing) | [optional] 
**ModuleKey** | Pointer to **string** | Catalog key from playbook_catalog_module (e.g. system_config_file) | [optional] 

## Methods

### NewSdkAutomationModuleFile

`func NewSdkAutomationModuleFile() *SdkAutomationModuleFile`

NewSdkAutomationModuleFile instantiates a new SdkAutomationModuleFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkAutomationModuleFileWithDefaults

`func NewSdkAutomationModuleFileWithDefaults() *SdkAutomationModuleFile`

NewSdkAutomationModuleFileWithDefaults instantiates a new SdkAutomationModuleFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SdkAutomationModuleFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SdkAutomationModuleFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SdkAutomationModuleFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SdkAutomationModuleFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFilename

`func (o *SdkAutomationModuleFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SdkAutomationModuleFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SdkAutomationModuleFile) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SdkAutomationModuleFile) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetModuleKey

`func (o *SdkAutomationModuleFile) GetModuleKey() string`

GetModuleKey returns the ModuleKey field if non-nil, zero value otherwise.

### GetModuleKeyOk

`func (o *SdkAutomationModuleFile) GetModuleKeyOk() (*string, bool)`

GetModuleKeyOk returns a tuple with the ModuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleKey

`func (o *SdkAutomationModuleFile) SetModuleKey(v string)`

SetModuleKey sets ModuleKey field to given value.

### HasModuleKey

`func (o *SdkAutomationModuleFile) HasModuleKey() bool`

HasModuleKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


