# V2NotificationCreatePostRequestNotificationBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Frequency** | Pointer to **int64** |  | [optional] 
**MessageBody** | Pointer to **string** |  | [optional] 
**NotificationName** | Pointer to **string** |  | [optional] 
**OpsgenieList** | Pointer to **[]string** |  | [optional] 
**OpsrampList** | Pointer to **[]string** |  | [optional] 
**RecipientList** | Pointer to **[]string** |  | [optional] 
**TeamsList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV2NotificationCreatePostRequestNotificationBody

`func NewV2NotificationCreatePostRequestNotificationBody() *V2NotificationCreatePostRequestNotificationBody`

NewV2NotificationCreatePostRequestNotificationBody instantiates a new V2NotificationCreatePostRequestNotificationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2NotificationCreatePostRequestNotificationBodyWithDefaults

`func NewV2NotificationCreatePostRequestNotificationBodyWithDefaults() *V2NotificationCreatePostRequestNotificationBody`

NewV2NotificationCreatePostRequestNotificationBodyWithDefaults instantiates a new V2NotificationCreatePostRequestNotificationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V2NotificationCreatePostRequestNotificationBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2NotificationCreatePostRequestNotificationBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2NotificationCreatePostRequestNotificationBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2NotificationCreatePostRequestNotificationBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *V2NotificationCreatePostRequestNotificationBody) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *V2NotificationCreatePostRequestNotificationBody) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *V2NotificationCreatePostRequestNotificationBody) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *V2NotificationCreatePostRequestNotificationBody) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnabled

`func (o *V2NotificationCreatePostRequestNotificationBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V2NotificationCreatePostRequestNotificationBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V2NotificationCreatePostRequestNotificationBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V2NotificationCreatePostRequestNotificationBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFrequency

`func (o *V2NotificationCreatePostRequestNotificationBody) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *V2NotificationCreatePostRequestNotificationBody) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *V2NotificationCreatePostRequestNotificationBody) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *V2NotificationCreatePostRequestNotificationBody) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetMessageBody

`func (o *V2NotificationCreatePostRequestNotificationBody) GetMessageBody() string`

GetMessageBody returns the MessageBody field if non-nil, zero value otherwise.

### GetMessageBodyOk

`func (o *V2NotificationCreatePostRequestNotificationBody) GetMessageBodyOk() (*string, bool)`

GetMessageBodyOk returns a tuple with the MessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBody

`func (o *V2NotificationCreatePostRequestNotificationBody) SetMessageBody(v string)`

SetMessageBody sets MessageBody field to given value.

### HasMessageBody

`func (o *V2NotificationCreatePostRequestNotificationBody) HasMessageBody() bool`

HasMessageBody returns a boolean if a field has been set.

### GetNotificationName

`func (o *V2NotificationCreatePostRequestNotificationBody) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *V2NotificationCreatePostRequestNotificationBody) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *V2NotificationCreatePostRequestNotificationBody) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.

### HasNotificationName

`func (o *V2NotificationCreatePostRequestNotificationBody) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### GetOpsgenieList

`func (o *V2NotificationCreatePostRequestNotificationBody) GetOpsgenieList() []string`

GetOpsgenieList returns the OpsgenieList field if non-nil, zero value otherwise.

### GetOpsgenieListOk

`func (o *V2NotificationCreatePostRequestNotificationBody) GetOpsgenieListOk() (*[]string, bool)`

GetOpsgenieListOk returns a tuple with the OpsgenieList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieList

`func (o *V2NotificationCreatePostRequestNotificationBody) SetOpsgenieList(v []string)`

SetOpsgenieList sets OpsgenieList field to given value.

### HasOpsgenieList

`func (o *V2NotificationCreatePostRequestNotificationBody) HasOpsgenieList() bool`

HasOpsgenieList returns a boolean if a field has been set.

### GetOpsrampList

`func (o *V2NotificationCreatePostRequestNotificationBody) GetOpsrampList() []string`

GetOpsrampList returns the OpsrampList field if non-nil, zero value otherwise.

### GetOpsrampListOk

`func (o *V2NotificationCreatePostRequestNotificationBody) GetOpsrampListOk() (*[]string, bool)`

GetOpsrampListOk returns a tuple with the OpsrampList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsrampList

`func (o *V2NotificationCreatePostRequestNotificationBody) SetOpsrampList(v []string)`

SetOpsrampList sets OpsrampList field to given value.

### HasOpsrampList

`func (o *V2NotificationCreatePostRequestNotificationBody) HasOpsrampList() bool`

HasOpsrampList returns a boolean if a field has been set.

### GetRecipientList

`func (o *V2NotificationCreatePostRequestNotificationBody) GetRecipientList() []string`

GetRecipientList returns the RecipientList field if non-nil, zero value otherwise.

### GetRecipientListOk

`func (o *V2NotificationCreatePostRequestNotificationBody) GetRecipientListOk() (*[]string, bool)`

GetRecipientListOk returns a tuple with the RecipientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientList

`func (o *V2NotificationCreatePostRequestNotificationBody) SetRecipientList(v []string)`

SetRecipientList sets RecipientList field to given value.

### HasRecipientList

`func (o *V2NotificationCreatePostRequestNotificationBody) HasRecipientList() bool`

HasRecipientList returns a boolean if a field has been set.

### GetTeamsList

`func (o *V2NotificationCreatePostRequestNotificationBody) GetTeamsList() []string`

GetTeamsList returns the TeamsList field if non-nil, zero value otherwise.

### GetTeamsListOk

`func (o *V2NotificationCreatePostRequestNotificationBody) GetTeamsListOk() (*[]string, bool)`

GetTeamsListOk returns a tuple with the TeamsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsList

`func (o *V2NotificationCreatePostRequestNotificationBody) SetTeamsList(v []string)`

SetTeamsList sets TeamsList field to given value.

### HasTeamsList

`func (o *V2NotificationCreatePostRequestNotificationBody) HasTeamsList() bool`

HasTeamsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


