# V2NotificationUpdatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationBody** | Pointer to [**V2NotificationCreatePostRequestNotificationBody**](V2NotificationCreatePostRequestNotificationBody.md) |  | [optional] 
**NotificationIdList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV2NotificationUpdatePostRequest

`func NewV2NotificationUpdatePostRequest() *V2NotificationUpdatePostRequest`

NewV2NotificationUpdatePostRequest instantiates a new V2NotificationUpdatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2NotificationUpdatePostRequestWithDefaults

`func NewV2NotificationUpdatePostRequestWithDefaults() *V2NotificationUpdatePostRequest`

NewV2NotificationUpdatePostRequestWithDefaults instantiates a new V2NotificationUpdatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationBody

`func (o *V2NotificationUpdatePostRequest) GetNotificationBody() V2NotificationCreatePostRequestNotificationBody`

GetNotificationBody returns the NotificationBody field if non-nil, zero value otherwise.

### GetNotificationBodyOk

`func (o *V2NotificationUpdatePostRequest) GetNotificationBodyOk() (*V2NotificationCreatePostRequestNotificationBody, bool)`

GetNotificationBodyOk returns a tuple with the NotificationBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationBody

`func (o *V2NotificationUpdatePostRequest) SetNotificationBody(v V2NotificationCreatePostRequestNotificationBody)`

SetNotificationBody sets NotificationBody field to given value.

### HasNotificationBody

`func (o *V2NotificationUpdatePostRequest) HasNotificationBody() bool`

HasNotificationBody returns a boolean if a field has been set.

### GetNotificationIdList

`func (o *V2NotificationUpdatePostRequest) GetNotificationIdList() []string`

GetNotificationIdList returns the NotificationIdList field if non-nil, zero value otherwise.

### GetNotificationIdListOk

`func (o *V2NotificationUpdatePostRequest) GetNotificationIdListOk() (*[]string, bool)`

GetNotificationIdListOk returns a tuple with the NotificationIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationIdList

`func (o *V2NotificationUpdatePostRequest) SetNotificationIdList(v []string)`

SetNotificationIdList sets NotificationIdList field to given value.

### HasNotificationIdList

`func (o *V2NotificationUpdatePostRequest) HasNotificationIdList() bool`

HasNotificationIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


