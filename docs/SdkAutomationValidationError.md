# SdkAutomationValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human-readable validation error detail | [optional] 
**ModuleKey** | Pointer to **string** | Catalog module key when the error is scoped to a module; unset otherwise | [optional] 
**Type** | Pointer to **string** | Error category (syntax or schema) | [optional] 

## Methods

### NewSdkAutomationValidationError

`func NewSdkAutomationValidationError() *SdkAutomationValidationError`

NewSdkAutomationValidationError instantiates a new SdkAutomationValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkAutomationValidationErrorWithDefaults

`func NewSdkAutomationValidationErrorWithDefaults() *SdkAutomationValidationError`

NewSdkAutomationValidationErrorWithDefaults instantiates a new SdkAutomationValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SdkAutomationValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SdkAutomationValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SdkAutomationValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SdkAutomationValidationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModuleKey

`func (o *SdkAutomationValidationError) GetModuleKey() string`

GetModuleKey returns the ModuleKey field if non-nil, zero value otherwise.

### GetModuleKeyOk

`func (o *SdkAutomationValidationError) GetModuleKeyOk() (*string, bool)`

GetModuleKeyOk returns a tuple with the ModuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleKey

`func (o *SdkAutomationValidationError) SetModuleKey(v string)`

SetModuleKey sets ModuleKey field to given value.

### HasModuleKey

`func (o *SdkAutomationValidationError) HasModuleKey() bool`

HasModuleKey returns a boolean if a field has been set.

### GetType

`func (o *SdkAutomationValidationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SdkAutomationValidationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SdkAutomationValidationError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SdkAutomationValidationError) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


