# V1SdkAutomationPlaybookTemplatesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionVersion** | Pointer to **string** | graphiant.naas collection version the files were read from | [optional] 
**Files** | Pointer to [**[]SdkAutomationTemplateFile**](SdkAutomationTemplateFile.md) |  | [optional] 

## Methods

### NewV1SdkAutomationPlaybookTemplatesGetResponse

`func NewV1SdkAutomationPlaybookTemplatesGetResponse() *V1SdkAutomationPlaybookTemplatesGetResponse`

NewV1SdkAutomationPlaybookTemplatesGetResponse instantiates a new V1SdkAutomationPlaybookTemplatesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SdkAutomationPlaybookTemplatesGetResponseWithDefaults

`func NewV1SdkAutomationPlaybookTemplatesGetResponseWithDefaults() *V1SdkAutomationPlaybookTemplatesGetResponse`

NewV1SdkAutomationPlaybookTemplatesGetResponseWithDefaults instantiates a new V1SdkAutomationPlaybookTemplatesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionVersion

`func (o *V1SdkAutomationPlaybookTemplatesGetResponse) GetCollectionVersion() string`

GetCollectionVersion returns the CollectionVersion field if non-nil, zero value otherwise.

### GetCollectionVersionOk

`func (o *V1SdkAutomationPlaybookTemplatesGetResponse) GetCollectionVersionOk() (*string, bool)`

GetCollectionVersionOk returns a tuple with the CollectionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionVersion

`func (o *V1SdkAutomationPlaybookTemplatesGetResponse) SetCollectionVersion(v string)`

SetCollectionVersion sets CollectionVersion field to given value.

### HasCollectionVersion

`func (o *V1SdkAutomationPlaybookTemplatesGetResponse) HasCollectionVersion() bool`

HasCollectionVersion returns a boolean if a field has been set.

### GetFiles

`func (o *V1SdkAutomationPlaybookTemplatesGetResponse) GetFiles() []SdkAutomationTemplateFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *V1SdkAutomationPlaybookTemplatesGetResponse) GetFilesOk() (*[]SdkAutomationTemplateFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *V1SdkAutomationPlaybookTemplatesGetResponse) SetFiles(v []SdkAutomationTemplateFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *V1SdkAutomationPlaybookTemplatesGetResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


