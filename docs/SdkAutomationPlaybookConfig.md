# SdkAutomationPlaybookConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleKey** | Pointer to **string** | Catalog key from playbook_catalog_playbook (e.g. system_bundle) | [optional] 
**ConfigId** | Pointer to **string** | Config id | [optional] 
**CreatedByUserId** | Pointer to **string** | IAM user id of the config creator; UI resolves the display name | [optional] 
**Description** | Pointer to **string** | Description / deployment notes | [optional] 
**Files** | Pointer to [**[]SdkAutomationModuleFile**](SdkAutomationModuleFile.md) |  | [optional] 
**Name** | Pointer to **string** | Display name | [optional] 
**Status** | Pointer to **string** | Config lifecycle status (draft / staged / paused) | [optional] 

## Methods

### NewSdkAutomationPlaybookConfig

`func NewSdkAutomationPlaybookConfig() *SdkAutomationPlaybookConfig`

NewSdkAutomationPlaybookConfig instantiates a new SdkAutomationPlaybookConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkAutomationPlaybookConfigWithDefaults

`func NewSdkAutomationPlaybookConfigWithDefaults() *SdkAutomationPlaybookConfig`

NewSdkAutomationPlaybookConfigWithDefaults instantiates a new SdkAutomationPlaybookConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleKey

`func (o *SdkAutomationPlaybookConfig) GetBundleKey() string`

GetBundleKey returns the BundleKey field if non-nil, zero value otherwise.

### GetBundleKeyOk

`func (o *SdkAutomationPlaybookConfig) GetBundleKeyOk() (*string, bool)`

GetBundleKeyOk returns a tuple with the BundleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleKey

`func (o *SdkAutomationPlaybookConfig) SetBundleKey(v string)`

SetBundleKey sets BundleKey field to given value.

### HasBundleKey

`func (o *SdkAutomationPlaybookConfig) HasBundleKey() bool`

HasBundleKey returns a boolean if a field has been set.

### GetConfigId

`func (o *SdkAutomationPlaybookConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *SdkAutomationPlaybookConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *SdkAutomationPlaybookConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *SdkAutomationPlaybookConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *SdkAutomationPlaybookConfig) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *SdkAutomationPlaybookConfig) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *SdkAutomationPlaybookConfig) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *SdkAutomationPlaybookConfig) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDescription

`func (o *SdkAutomationPlaybookConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdkAutomationPlaybookConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdkAutomationPlaybookConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdkAutomationPlaybookConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFiles

`func (o *SdkAutomationPlaybookConfig) GetFiles() []SdkAutomationModuleFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SdkAutomationPlaybookConfig) GetFilesOk() (*[]SdkAutomationModuleFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SdkAutomationPlaybookConfig) SetFiles(v []SdkAutomationModuleFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SdkAutomationPlaybookConfig) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetName

`func (o *SdkAutomationPlaybookConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkAutomationPlaybookConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkAutomationPlaybookConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdkAutomationPlaybookConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *SdkAutomationPlaybookConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SdkAutomationPlaybookConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SdkAutomationPlaybookConfig) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SdkAutomationPlaybookConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


