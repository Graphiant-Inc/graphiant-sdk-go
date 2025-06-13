# V2AssistantAddToConversationPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | Pointer to **string** |  | [optional] 
**Question** | Pointer to [**V2AssistantAddToConversationPostRequestQuestion**](V2AssistantAddToConversationPostRequestQuestion.md) |  | [optional] 

## Methods

### NewV2AssistantAddToConversationPostRequest

`func NewV2AssistantAddToConversationPostRequest() *V2AssistantAddToConversationPostRequest`

NewV2AssistantAddToConversationPostRequest instantiates a new V2AssistantAddToConversationPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssistantAddToConversationPostRequestWithDefaults

`func NewV2AssistantAddToConversationPostRequestWithDefaults() *V2AssistantAddToConversationPostRequest`

NewV2AssistantAddToConversationPostRequestWithDefaults instantiates a new V2AssistantAddToConversationPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *V2AssistantAddToConversationPostRequest) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *V2AssistantAddToConversationPostRequest) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *V2AssistantAddToConversationPostRequest) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *V2AssistantAddToConversationPostRequest) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetQuestion

`func (o *V2AssistantAddToConversationPostRequest) GetQuestion() V2AssistantAddToConversationPostRequestQuestion`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *V2AssistantAddToConversationPostRequest) GetQuestionOk() (*V2AssistantAddToConversationPostRequestQuestion, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *V2AssistantAddToConversationPostRequest) SetQuestion(v V2AssistantAddToConversationPostRequestQuestion)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *V2AssistantAddToConversationPostRequest) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


