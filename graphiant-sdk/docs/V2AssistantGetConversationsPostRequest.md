# V2AssistantGetConversationsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssistantGetConversationsPostRequest

`func NewV2AssistantGetConversationsPostRequest() *V2AssistantGetConversationsPostRequest`

NewV2AssistantGetConversationsPostRequest instantiates a new V2AssistantGetConversationsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssistantGetConversationsPostRequestWithDefaults

`func NewV2AssistantGetConversationsPostRequestWithDefaults() *V2AssistantGetConversationsPostRequest`

NewV2AssistantGetConversationsPostRequestWithDefaults instantiates a new V2AssistantGetConversationsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeWindow

`func (o *V2AssistantGetConversationsPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2AssistantGetConversationsPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2AssistantGetConversationsPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2AssistantGetConversationsPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetUserId

`func (o *V2AssistantGetConversationsPostRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V2AssistantGetConversationsPostRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V2AssistantGetConversationsPostRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V2AssistantGetConversationsPostRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


