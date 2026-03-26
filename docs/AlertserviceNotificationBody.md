# AlertserviceNotificationBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of notification | [optional] 
**Duration** | Pointer to **string** | Time interval for notification (required) | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable notification (required) | [optional] 
**Frequency** | Pointer to **int64** | Frequency of notifying a continuing alert (required) | [optional] 
**MessageBody** | Pointer to **string** | Message body to prepend to actual message | [optional] 
**NotificationName** | Pointer to **string** | Name of the notification record (required) | [optional] 
**OpsgenieList** | Pointer to **[]string** |  | [optional] 
**OpsrampList** | Pointer to **[]string** |  | [optional] 
**PagerdutyList** | Pointer to **[]string** |  | [optional] 
**RecipientList** | Pointer to **[]string** |  | [optional] 
**TeamsList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAlertserviceNotificationBody

`func NewAlertserviceNotificationBody() *AlertserviceNotificationBody`

NewAlertserviceNotificationBody instantiates a new AlertserviceNotificationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertserviceNotificationBodyWithDefaults

`func NewAlertserviceNotificationBodyWithDefaults() *AlertserviceNotificationBody`

NewAlertserviceNotificationBodyWithDefaults instantiates a new AlertserviceNotificationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AlertserviceNotificationBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertserviceNotificationBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertserviceNotificationBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertserviceNotificationBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *AlertserviceNotificationBody) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AlertserviceNotificationBody) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AlertserviceNotificationBody) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AlertserviceNotificationBody) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnabled

`func (o *AlertserviceNotificationBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertserviceNotificationBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertserviceNotificationBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlertserviceNotificationBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFrequency

`func (o *AlertserviceNotificationBody) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AlertserviceNotificationBody) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AlertserviceNotificationBody) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *AlertserviceNotificationBody) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetMessageBody

`func (o *AlertserviceNotificationBody) GetMessageBody() string`

GetMessageBody returns the MessageBody field if non-nil, zero value otherwise.

### GetMessageBodyOk

`func (o *AlertserviceNotificationBody) GetMessageBodyOk() (*string, bool)`

GetMessageBodyOk returns a tuple with the MessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBody

`func (o *AlertserviceNotificationBody) SetMessageBody(v string)`

SetMessageBody sets MessageBody field to given value.

### HasMessageBody

`func (o *AlertserviceNotificationBody) HasMessageBody() bool`

HasMessageBody returns a boolean if a field has been set.

### GetNotificationName

`func (o *AlertserviceNotificationBody) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *AlertserviceNotificationBody) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *AlertserviceNotificationBody) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.

### HasNotificationName

`func (o *AlertserviceNotificationBody) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### GetOpsgenieList

`func (o *AlertserviceNotificationBody) GetOpsgenieList() []string`

GetOpsgenieList returns the OpsgenieList field if non-nil, zero value otherwise.

### GetOpsgenieListOk

`func (o *AlertserviceNotificationBody) GetOpsgenieListOk() (*[]string, bool)`

GetOpsgenieListOk returns a tuple with the OpsgenieList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieList

`func (o *AlertserviceNotificationBody) SetOpsgenieList(v []string)`

SetOpsgenieList sets OpsgenieList field to given value.

### HasOpsgenieList

`func (o *AlertserviceNotificationBody) HasOpsgenieList() bool`

HasOpsgenieList returns a boolean if a field has been set.

### GetOpsrampList

`func (o *AlertserviceNotificationBody) GetOpsrampList() []string`

GetOpsrampList returns the OpsrampList field if non-nil, zero value otherwise.

### GetOpsrampListOk

`func (o *AlertserviceNotificationBody) GetOpsrampListOk() (*[]string, bool)`

GetOpsrampListOk returns a tuple with the OpsrampList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsrampList

`func (o *AlertserviceNotificationBody) SetOpsrampList(v []string)`

SetOpsrampList sets OpsrampList field to given value.

### HasOpsrampList

`func (o *AlertserviceNotificationBody) HasOpsrampList() bool`

HasOpsrampList returns a boolean if a field has been set.

### GetPagerdutyList

`func (o *AlertserviceNotificationBody) GetPagerdutyList() []string`

GetPagerdutyList returns the PagerdutyList field if non-nil, zero value otherwise.

### GetPagerdutyListOk

`func (o *AlertserviceNotificationBody) GetPagerdutyListOk() (*[]string, bool)`

GetPagerdutyListOk returns a tuple with the PagerdutyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerdutyList

`func (o *AlertserviceNotificationBody) SetPagerdutyList(v []string)`

SetPagerdutyList sets PagerdutyList field to given value.

### HasPagerdutyList

`func (o *AlertserviceNotificationBody) HasPagerdutyList() bool`

HasPagerdutyList returns a boolean if a field has been set.

### GetRecipientList

`func (o *AlertserviceNotificationBody) GetRecipientList() []string`

GetRecipientList returns the RecipientList field if non-nil, zero value otherwise.

### GetRecipientListOk

`func (o *AlertserviceNotificationBody) GetRecipientListOk() (*[]string, bool)`

GetRecipientListOk returns a tuple with the RecipientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientList

`func (o *AlertserviceNotificationBody) SetRecipientList(v []string)`

SetRecipientList sets RecipientList field to given value.

### HasRecipientList

`func (o *AlertserviceNotificationBody) HasRecipientList() bool`

HasRecipientList returns a boolean if a field has been set.

### GetTeamsList

`func (o *AlertserviceNotificationBody) GetTeamsList() []string`

GetTeamsList returns the TeamsList field if non-nil, zero value otherwise.

### GetTeamsListOk

`func (o *AlertserviceNotificationBody) GetTeamsListOk() (*[]string, bool)`

GetTeamsListOk returns a tuple with the TeamsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsList

`func (o *AlertserviceNotificationBody) SetTeamsList(v []string)`

SetTeamsList sets TeamsList field to given value.

### HasTeamsList

`func (o *AlertserviceNotificationBody) HasTeamsList() bool`

HasTeamsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


