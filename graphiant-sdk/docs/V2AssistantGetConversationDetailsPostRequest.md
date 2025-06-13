# V2AssistantGetConversationDetailsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | Pointer to **string** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssistantGetConversationDetailsPostRequest

`func NewV2AssistantGetConversationDetailsPostRequest() *V2AssistantGetConversationDetailsPostRequest`

NewV2AssistantGetConversationDetailsPostRequest instantiates a new V2AssistantGetConversationDetailsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssistantGetConversationDetailsPostRequestWithDefaults

`func NewV2AssistantGetConversationDetailsPostRequestWithDefaults() *V2AssistantGetConversationDetailsPostRequest`

NewV2AssistantGetConversationDetailsPostRequestWithDefaults instantiates a new V2AssistantGetConversationDetailsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *V2AssistantGetConversationDetailsPostRequest) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *V2AssistantGetConversationDetailsPostRequest) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *V2AssistantGetConversationDetailsPostRequest) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *V2AssistantGetConversationDetailsPostRequest) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2AssistantGetConversationDetailsPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2AssistantGetConversationDetailsPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2AssistantGetConversationDetailsPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2AssistantGetConversationDetailsPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetUserId

`func (o *V2AssistantGetConversationDetailsPostRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V2AssistantGetConversationDetailsPostRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V2AssistantGetConversationDetailsPostRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V2AssistantGetConversationDetailsPostRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


