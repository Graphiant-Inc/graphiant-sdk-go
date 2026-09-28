# V1SdkAutomationPlaybookConfigsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleKey** | Pointer to **string** | Catalog key from playbook_catalog_playbook (e.g. system_bundle) | [optional] 
**Files** | Pointer to [**[]SdkAutomationModuleFile**](SdkAutomationModuleFile.md) |  | [optional] 
**ValidateOnly** | Pointer to **bool** | When true, validate without persisting (Validate Bundle); when false, create on success | [optional] 

## Methods

### NewV1SdkAutomationPlaybookConfigsPostRequest

`func NewV1SdkAutomationPlaybookConfigsPostRequest() *V1SdkAutomationPlaybookConfigsPostRequest`

NewV1SdkAutomationPlaybookConfigsPostRequest instantiates a new V1SdkAutomationPlaybookConfigsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SdkAutomationPlaybookConfigsPostRequestWithDefaults

`func NewV1SdkAutomationPlaybookConfigsPostRequestWithDefaults() *V1SdkAutomationPlaybookConfigsPostRequest`

NewV1SdkAutomationPlaybookConfigsPostRequestWithDefaults instantiates a new V1SdkAutomationPlaybookConfigsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleKey

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) GetBundleKey() string`

GetBundleKey returns the BundleKey field if non-nil, zero value otherwise.

### GetBundleKeyOk

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) GetBundleKeyOk() (*string, bool)`

GetBundleKeyOk returns a tuple with the BundleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleKey

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) SetBundleKey(v string)`

SetBundleKey sets BundleKey field to given value.

### HasBundleKey

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) HasBundleKey() bool`

HasBundleKey returns a boolean if a field has been set.

### GetFiles

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) GetFiles() []SdkAutomationModuleFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) GetFilesOk() (*[]SdkAutomationModuleFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) SetFiles(v []SdkAutomationModuleFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetValidateOnly

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) GetValidateOnly() bool`

GetValidateOnly returns the ValidateOnly field if non-nil, zero value otherwise.

### GetValidateOnlyOk

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) GetValidateOnlyOk() (*bool, bool)`

GetValidateOnlyOk returns a tuple with the ValidateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOnly

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) SetValidateOnly(v bool)`

SetValidateOnly sets ValidateOnly field to given value.

### HasValidateOnly

`func (o *V1SdkAutomationPlaybookConfigsPostRequest) HasValidateOnly() bool`

HasValidateOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


