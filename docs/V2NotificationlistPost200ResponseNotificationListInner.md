# V2NotificationlistPost200ResponseNotificationListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertType** | Pointer to **string** |  | [optional] 
**MuteCount** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotificationBody** | Pointer to [**V2NotificationCreatePostRequestNotificationBody**](V2NotificationCreatePostRequestNotificationBody.md) |  | [optional] 
**NotificationId** | Pointer to **string** |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**TimesTriggered** | Pointer to **int64** |  | [optional] 

## Methods

### NewV2NotificationlistPost200ResponseNotificationListInner

`func NewV2NotificationlistPost200ResponseNotificationListInner() *V2NotificationlistPost200ResponseNotificationListInner`

NewV2NotificationlistPost200ResponseNotificationListInner instantiates a new V2NotificationlistPost200ResponseNotificationListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2NotificationlistPost200ResponseNotificationListInnerWithDefaults

`func NewV2NotificationlistPost200ResponseNotificationListInnerWithDefaults() *V2NotificationlistPost200ResponseNotificationListInner`

NewV2NotificationlistPost200ResponseNotificationListInnerWithDefaults instantiates a new V2NotificationlistPost200ResponseNotificationListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertType

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *V2NotificationlistPost200ResponseNotificationListInner) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *V2NotificationlistPost200ResponseNotificationListInner) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetMuteCount

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetMuteCount() int64`

GetMuteCount returns the MuteCount field if non-nil, zero value otherwise.

### GetMuteCountOk

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetMuteCountOk() (*int64, bool)`

GetMuteCountOk returns a tuple with the MuteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteCount

`func (o *V2NotificationlistPost200ResponseNotificationListInner) SetMuteCount(v int64)`

SetMuteCount sets MuteCount field to given value.

### HasMuteCount

`func (o *V2NotificationlistPost200ResponseNotificationListInner) HasMuteCount() bool`

HasMuteCount returns a boolean if a field has been set.

### GetName

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2NotificationlistPost200ResponseNotificationListInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2NotificationlistPost200ResponseNotificationListInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotificationBody

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetNotificationBody() V2NotificationCreatePostRequestNotificationBody`

GetNotificationBody returns the NotificationBody field if non-nil, zero value otherwise.

### GetNotificationBodyOk

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetNotificationBodyOk() (*V2NotificationCreatePostRequestNotificationBody, bool)`

GetNotificationBodyOk returns a tuple with the NotificationBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationBody

`func (o *V2NotificationlistPost200ResponseNotificationListInner) SetNotificationBody(v V2NotificationCreatePostRequestNotificationBody)`

SetNotificationBody sets NotificationBody field to given value.

### HasNotificationBody

`func (o *V2NotificationlistPost200ResponseNotificationListInner) HasNotificationBody() bool`

HasNotificationBody returns a boolean if a field has been set.

### GetNotificationId

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *V2NotificationlistPost200ResponseNotificationListInner) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *V2NotificationlistPost200ResponseNotificationListInner) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetRuleId

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *V2NotificationlistPost200ResponseNotificationListInner) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *V2NotificationlistPost200ResponseNotificationListInner) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetTimesTriggered

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetTimesTriggered() int64`

GetTimesTriggered returns the TimesTriggered field if non-nil, zero value otherwise.

### GetTimesTriggeredOk

`func (o *V2NotificationlistPost200ResponseNotificationListInner) GetTimesTriggeredOk() (*int64, bool)`

GetTimesTriggeredOk returns a tuple with the TimesTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesTriggered

`func (o *V2NotificationlistPost200ResponseNotificationListInner) SetTimesTriggered(v int64)`

SetTimesTriggered sets TimesTriggered field to given value.

### HasTimesTriggered

`func (o *V2NotificationlistPost200ResponseNotificationListInner) HasTimesTriggered() bool`

HasTimesTriggered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


