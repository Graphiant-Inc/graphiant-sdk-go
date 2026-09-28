# V1SdkAutomationPlaybookConfigsConfigIdPutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]SdkAutomationValidationError**](SdkAutomationValidationError.md) |  | [optional] 
**IsValid** | Pointer to **bool** | True when validation succeeded and the update was applied | [optional] 

## Methods

### NewV1SdkAutomationPlaybookConfigsConfigIdPutResponse

`func NewV1SdkAutomationPlaybookConfigsConfigIdPutResponse() *V1SdkAutomationPlaybookConfigsConfigIdPutResponse`

NewV1SdkAutomationPlaybookConfigsConfigIdPutResponse instantiates a new V1SdkAutomationPlaybookConfigsConfigIdPutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SdkAutomationPlaybookConfigsConfigIdPutResponseWithDefaults

`func NewV1SdkAutomationPlaybookConfigsConfigIdPutResponseWithDefaults() *V1SdkAutomationPlaybookConfigsConfigIdPutResponse`

NewV1SdkAutomationPlaybookConfigsConfigIdPutResponseWithDefaults instantiates a new V1SdkAutomationPlaybookConfigsConfigIdPutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V1SdkAutomationPlaybookConfigsConfigIdPutResponse) GetErrors() []SdkAutomationValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V1SdkAutomationPlaybookConfigsConfigIdPutResponse) GetErrorsOk() (*[]SdkAutomationValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V1SdkAutomationPlaybookConfigsConfigIdPutResponse) SetErrors(v []SdkAutomationValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V1SdkAutomationPlaybookConfigsConfigIdPutResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetIsValid

`func (o *V1SdkAutomationPlaybookConfigsConfigIdPutResponse) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *V1SdkAutomationPlaybookConfigsConfigIdPutResponse) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *V1SdkAutomationPlaybookConfigsConfigIdPutResponse) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *V1SdkAutomationPlaybookConfigsConfigIdPutResponse) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


