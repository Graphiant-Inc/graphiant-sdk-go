# SdkAutomationCatalogModuleSlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleKey** | Pointer to **string** | Owning catalog playbook key (e.g. system_bundle) | [optional] 
**Description** | Pointer to **string** | Short description of the module slot | [optional] 
**FileDistinguisher** | Pointer to **string** | Top-level YAML key from playbook_catalog_module the UI uses to map an uploaded file to this module (e.g. device_system, interfaces, circuits) | [optional] 
**IsRequired** | Pointer to **bool** | When true, create/validate must include this module_key | [optional] 
**ModuleKey** | Pointer to **string** | Stable catalog key from playbook_catalog_module (e.g. system_config_file); used as ModuleFile.module_key | [optional] 
**Name** | Pointer to **string** | Display name for the module slot | [optional] 

## Methods

### NewSdkAutomationCatalogModuleSlot

`func NewSdkAutomationCatalogModuleSlot() *SdkAutomationCatalogModuleSlot`

NewSdkAutomationCatalogModuleSlot instantiates a new SdkAutomationCatalogModuleSlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkAutomationCatalogModuleSlotWithDefaults

`func NewSdkAutomationCatalogModuleSlotWithDefaults() *SdkAutomationCatalogModuleSlot`

NewSdkAutomationCatalogModuleSlotWithDefaults instantiates a new SdkAutomationCatalogModuleSlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleKey

`func (o *SdkAutomationCatalogModuleSlot) GetBundleKey() string`

GetBundleKey returns the BundleKey field if non-nil, zero value otherwise.

### GetBundleKeyOk

`func (o *SdkAutomationCatalogModuleSlot) GetBundleKeyOk() (*string, bool)`

GetBundleKeyOk returns a tuple with the BundleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleKey

`func (o *SdkAutomationCatalogModuleSlot) SetBundleKey(v string)`

SetBundleKey sets BundleKey field to given value.

### HasBundleKey

`func (o *SdkAutomationCatalogModuleSlot) HasBundleKey() bool`

HasBundleKey returns a boolean if a field has been set.

### GetDescription

`func (o *SdkAutomationCatalogModuleSlot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdkAutomationCatalogModuleSlot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdkAutomationCatalogModuleSlot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdkAutomationCatalogModuleSlot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileDistinguisher

`func (o *SdkAutomationCatalogModuleSlot) GetFileDistinguisher() string`

GetFileDistinguisher returns the FileDistinguisher field if non-nil, zero value otherwise.

### GetFileDistinguisherOk

`func (o *SdkAutomationCatalogModuleSlot) GetFileDistinguisherOk() (*string, bool)`

GetFileDistinguisherOk returns a tuple with the FileDistinguisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDistinguisher

`func (o *SdkAutomationCatalogModuleSlot) SetFileDistinguisher(v string)`

SetFileDistinguisher sets FileDistinguisher field to given value.

### HasFileDistinguisher

`func (o *SdkAutomationCatalogModuleSlot) HasFileDistinguisher() bool`

HasFileDistinguisher returns a boolean if a field has been set.

### GetIsRequired

`func (o *SdkAutomationCatalogModuleSlot) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *SdkAutomationCatalogModuleSlot) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *SdkAutomationCatalogModuleSlot) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *SdkAutomationCatalogModuleSlot) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetModuleKey

`func (o *SdkAutomationCatalogModuleSlot) GetModuleKey() string`

GetModuleKey returns the ModuleKey field if non-nil, zero value otherwise.

### GetModuleKeyOk

`func (o *SdkAutomationCatalogModuleSlot) GetModuleKeyOk() (*string, bool)`

GetModuleKeyOk returns a tuple with the ModuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleKey

`func (o *SdkAutomationCatalogModuleSlot) SetModuleKey(v string)`

SetModuleKey sets ModuleKey field to given value.

### HasModuleKey

`func (o *SdkAutomationCatalogModuleSlot) HasModuleKey() bool`

HasModuleKey returns a boolean if a field has been set.

### GetName

`func (o *SdkAutomationCatalogModuleSlot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkAutomationCatalogModuleSlot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkAutomationCatalogModuleSlot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdkAutomationCatalogModuleSlot) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


