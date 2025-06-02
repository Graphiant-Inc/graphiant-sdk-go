# V2NotificationCreatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationBody** | Pointer to [**V2NotificationCreatePostRequestNotificationBody**](V2NotificationCreatePostRequestNotificationBody.md) |  | [optional] 
**RuleIdList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV2NotificationCreatePostRequest

`func NewV2NotificationCreatePostRequest() *V2NotificationCreatePostRequest`

NewV2NotificationCreatePostRequest instantiates a new V2NotificationCreatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2NotificationCreatePostRequestWithDefaults

`func NewV2NotificationCreatePostRequestWithDefaults() *V2NotificationCreatePostRequest`

NewV2NotificationCreatePostRequestWithDefaults instantiates a new V2NotificationCreatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationBody

`func (o *V2NotificationCreatePostRequest) GetNotificationBody() V2NotificationCreatePostRequestNotificationBody`

GetNotificationBody returns the NotificationBody field if non-nil, zero value otherwise.

### GetNotificationBodyOk

`func (o *V2NotificationCreatePostRequest) GetNotificationBodyOk() (*V2NotificationCreatePostRequestNotificationBody, bool)`

GetNotificationBodyOk returns a tuple with the NotificationBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationBody

`func (o *V2NotificationCreatePostRequest) SetNotificationBody(v V2NotificationCreatePostRequestNotificationBody)`

SetNotificationBody sets NotificationBody field to given value.

### HasNotificationBody

`func (o *V2NotificationCreatePostRequest) HasNotificationBody() bool`

HasNotificationBody returns a boolean if a field has been set.

### GetRuleIdList

`func (o *V2NotificationCreatePostRequest) GetRuleIdList() []string`

GetRuleIdList returns the RuleIdList field if non-nil, zero value otherwise.

### GetRuleIdListOk

`func (o *V2NotificationCreatePostRequest) GetRuleIdListOk() (*[]string, bool)`

GetRuleIdListOk returns a tuple with the RuleIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIdList

`func (o *V2NotificationCreatePostRequest) SetRuleIdList(v []string)`

SetRuleIdList sets RuleIdList field to given value.

### HasRuleIdList

`func (o *V2NotificationCreatePostRequest) HasRuleIdList() bool`

HasRuleIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


