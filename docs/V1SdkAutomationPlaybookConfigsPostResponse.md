# V1SdkAutomationPlaybookConfigsPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **string** | Created config id when validation succeeds and validate_only is false; empty otherwise | [optional] 
**Errors** | Pointer to [**[]SdkAutomationValidationError**](SdkAutomationValidationError.md) |  | [optional] 

## Methods

### NewV1SdkAutomationPlaybookConfigsPostResponse

`func NewV1SdkAutomationPlaybookConfigsPostResponse() *V1SdkAutomationPlaybookConfigsPostResponse`

NewV1SdkAutomationPlaybookConfigsPostResponse instantiates a new V1SdkAutomationPlaybookConfigsPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SdkAutomationPlaybookConfigsPostResponseWithDefaults

`func NewV1SdkAutomationPlaybookConfigsPostResponseWithDefaults() *V1SdkAutomationPlaybookConfigsPostResponse`

NewV1SdkAutomationPlaybookConfigsPostResponseWithDefaults instantiates a new V1SdkAutomationPlaybookConfigsPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *V1SdkAutomationPlaybookConfigsPostResponse) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *V1SdkAutomationPlaybookConfigsPostResponse) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *V1SdkAutomationPlaybookConfigsPostResponse) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *V1SdkAutomationPlaybookConfigsPostResponse) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetErrors

`func (o *V1SdkAutomationPlaybookConfigsPostResponse) GetErrors() []SdkAutomationValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V1SdkAutomationPlaybookConfigsPostResponse) GetErrorsOk() (*[]SdkAutomationValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V1SdkAutomationPlaybookConfigsPostResponse) SetErrors(v []SdkAutomationValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V1SdkAutomationPlaybookConfigsPostResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


