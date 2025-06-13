# V2AssistantAddToConversationPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | Pointer to **string** |  | [optional] 
**DataframeDictionary** | Pointer to [**[]V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner**](V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner.md) |  | [optional] 
**OriginalQuestion** | Pointer to [**V2AssistantAddToConversationPostRequestQuestion**](V2AssistantAddToConversationPostRequestQuestion.md) |  | [optional] 
**ResponseId** | Pointer to **string** |  | [optional] 
**ResponseLanguage** | Pointer to **string** |  | [optional] 
**ResponseText** | Pointer to **string** |  | [optional] 
**ResponseTimestamp** | Pointer to **int64** |  | [optional] 
**ResponseType** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssistantAddToConversationPost200Response

`func NewV2AssistantAddToConversationPost200Response() *V2AssistantAddToConversationPost200Response`

NewV2AssistantAddToConversationPost200Response instantiates a new V2AssistantAddToConversationPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssistantAddToConversationPost200ResponseWithDefaults

`func NewV2AssistantAddToConversationPost200ResponseWithDefaults() *V2AssistantAddToConversationPost200Response`

NewV2AssistantAddToConversationPost200ResponseWithDefaults instantiates a new V2AssistantAddToConversationPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *V2AssistantAddToConversationPost200Response) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *V2AssistantAddToConversationPost200Response) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *V2AssistantAddToConversationPost200Response) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *V2AssistantAddToConversationPost200Response) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetDataframeDictionary

`func (o *V2AssistantAddToConversationPost200Response) GetDataframeDictionary() []V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner`

GetDataframeDictionary returns the DataframeDictionary field if non-nil, zero value otherwise.

### GetDataframeDictionaryOk

`func (o *V2AssistantAddToConversationPost200Response) GetDataframeDictionaryOk() (*[]V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner, bool)`

GetDataframeDictionaryOk returns a tuple with the DataframeDictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataframeDictionary

`func (o *V2AssistantAddToConversationPost200Response) SetDataframeDictionary(v []V2AssistantAddToConversationPost200ResponseDataframeDictionaryInner)`

SetDataframeDictionary sets DataframeDictionary field to given value.

### HasDataframeDictionary

`func (o *V2AssistantAddToConversationPost200Response) HasDataframeDictionary() bool`

HasDataframeDictionary returns a boolean if a field has been set.

### GetOriginalQuestion

`func (o *V2AssistantAddToConversationPost200Response) GetOriginalQuestion() V2AssistantAddToConversationPostRequestQuestion`

GetOriginalQuestion returns the OriginalQuestion field if non-nil, zero value otherwise.

### GetOriginalQuestionOk

`func (o *V2AssistantAddToConversationPost200Response) GetOriginalQuestionOk() (*V2AssistantAddToConversationPostRequestQuestion, bool)`

GetOriginalQuestionOk returns a tuple with the OriginalQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalQuestion

`func (o *V2AssistantAddToConversationPost200Response) SetOriginalQuestion(v V2AssistantAddToConversationPostRequestQuestion)`

SetOriginalQuestion sets OriginalQuestion field to given value.

### HasOriginalQuestion

`func (o *V2AssistantAddToConversationPost200Response) HasOriginalQuestion() bool`

HasOriginalQuestion returns a boolean if a field has been set.

### GetResponseId

`func (o *V2AssistantAddToConversationPost200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *V2AssistantAddToConversationPost200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *V2AssistantAddToConversationPost200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *V2AssistantAddToConversationPost200Response) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetResponseLanguage

`func (o *V2AssistantAddToConversationPost200Response) GetResponseLanguage() string`

GetResponseLanguage returns the ResponseLanguage field if non-nil, zero value otherwise.

### GetResponseLanguageOk

`func (o *V2AssistantAddToConversationPost200Response) GetResponseLanguageOk() (*string, bool)`

GetResponseLanguageOk returns a tuple with the ResponseLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLanguage

`func (o *V2AssistantAddToConversationPost200Response) SetResponseLanguage(v string)`

SetResponseLanguage sets ResponseLanguage field to given value.

### HasResponseLanguage

`func (o *V2AssistantAddToConversationPost200Response) HasResponseLanguage() bool`

HasResponseLanguage returns a boolean if a field has been set.

### GetResponseText

`func (o *V2AssistantAddToConversationPost200Response) GetResponseText() string`

GetResponseText returns the ResponseText field if non-nil, zero value otherwise.

### GetResponseTextOk

`func (o *V2AssistantAddToConversationPost200Response) GetResponseTextOk() (*string, bool)`

GetResponseTextOk returns a tuple with the ResponseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseText

`func (o *V2AssistantAddToConversationPost200Response) SetResponseText(v string)`

SetResponseText sets ResponseText field to given value.

### HasResponseText

`func (o *V2AssistantAddToConversationPost200Response) HasResponseText() bool`

HasResponseText returns a boolean if a field has been set.

### GetResponseTimestamp

`func (o *V2AssistantAddToConversationPost200Response) GetResponseTimestamp() int64`

GetResponseTimestamp returns the ResponseTimestamp field if non-nil, zero value otherwise.

### GetResponseTimestampOk

`func (o *V2AssistantAddToConversationPost200Response) GetResponseTimestampOk() (*int64, bool)`

GetResponseTimestampOk returns a tuple with the ResponseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimestamp

`func (o *V2AssistantAddToConversationPost200Response) SetResponseTimestamp(v int64)`

SetResponseTimestamp sets ResponseTimestamp field to given value.

### HasResponseTimestamp

`func (o *V2AssistantAddToConversationPost200Response) HasResponseTimestamp() bool`

HasResponseTimestamp returns a boolean if a field has been set.

### GetResponseType

`func (o *V2AssistantAddToConversationPost200Response) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *V2AssistantAddToConversationPost200Response) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *V2AssistantAddToConversationPost200Response) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *V2AssistantAddToConversationPost200Response) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


