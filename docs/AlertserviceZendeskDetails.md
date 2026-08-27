# AlertserviceZendeskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZendeskApiToken** | Pointer to **string** | zendesk api token (deprecated, use zendesk_client_secret) | [optional] 
**ZendeskAssigneeId** | **string** | zendesk assignee id (required) | 
**ZendeskBaseUrl** | **string** | zendesk base url (required) | 
**ZendeskClientId** | **string** | zendesk oauth client id (required) | 
**ZendeskClientSecret** | **string** | zendesk oauth client secret (required) | 
**ZendeskEmail** | Pointer to **string** | zendesk email (deprecated, use zendesk_client_id) | [optional] 

## Methods

### NewAlertserviceZendeskDetails

`func NewAlertserviceZendeskDetails(zendeskAssigneeId string, zendeskBaseUrl string, zendeskClientId string, zendeskClientSecret string, ) *AlertserviceZendeskDetails`

NewAlertserviceZendeskDetails instantiates a new AlertserviceZendeskDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertserviceZendeskDetailsWithDefaults

`func NewAlertserviceZendeskDetailsWithDefaults() *AlertserviceZendeskDetails`

NewAlertserviceZendeskDetailsWithDefaults instantiates a new AlertserviceZendeskDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZendeskApiToken

`func (o *AlertserviceZendeskDetails) GetZendeskApiToken() string`

GetZendeskApiToken returns the ZendeskApiToken field if non-nil, zero value otherwise.

### GetZendeskApiTokenOk

`func (o *AlertserviceZendeskDetails) GetZendeskApiTokenOk() (*string, bool)`

GetZendeskApiTokenOk returns a tuple with the ZendeskApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskApiToken

`func (o *AlertserviceZendeskDetails) SetZendeskApiToken(v string)`

SetZendeskApiToken sets ZendeskApiToken field to given value.

### HasZendeskApiToken

`func (o *AlertserviceZendeskDetails) HasZendeskApiToken() bool`

HasZendeskApiToken returns a boolean if a field has been set.

### GetZendeskAssigneeId

`func (o *AlertserviceZendeskDetails) GetZendeskAssigneeId() string`

GetZendeskAssigneeId returns the ZendeskAssigneeId field if non-nil, zero value otherwise.

### GetZendeskAssigneeIdOk

`func (o *AlertserviceZendeskDetails) GetZendeskAssigneeIdOk() (*string, bool)`

GetZendeskAssigneeIdOk returns a tuple with the ZendeskAssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskAssigneeId

`func (o *AlertserviceZendeskDetails) SetZendeskAssigneeId(v string)`

SetZendeskAssigneeId sets ZendeskAssigneeId field to given value.


### GetZendeskBaseUrl

`func (o *AlertserviceZendeskDetails) GetZendeskBaseUrl() string`

GetZendeskBaseUrl returns the ZendeskBaseUrl field if non-nil, zero value otherwise.

### GetZendeskBaseUrlOk

`func (o *AlertserviceZendeskDetails) GetZendeskBaseUrlOk() (*string, bool)`

GetZendeskBaseUrlOk returns a tuple with the ZendeskBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskBaseUrl

`func (o *AlertserviceZendeskDetails) SetZendeskBaseUrl(v string)`

SetZendeskBaseUrl sets ZendeskBaseUrl field to given value.


### GetZendeskClientId

`func (o *AlertserviceZendeskDetails) GetZendeskClientId() string`

GetZendeskClientId returns the ZendeskClientId field if non-nil, zero value otherwise.

### GetZendeskClientIdOk

`func (o *AlertserviceZendeskDetails) GetZendeskClientIdOk() (*string, bool)`

GetZendeskClientIdOk returns a tuple with the ZendeskClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskClientId

`func (o *AlertserviceZendeskDetails) SetZendeskClientId(v string)`

SetZendeskClientId sets ZendeskClientId field to given value.


### GetZendeskClientSecret

`func (o *AlertserviceZendeskDetails) GetZendeskClientSecret() string`

GetZendeskClientSecret returns the ZendeskClientSecret field if non-nil, zero value otherwise.

### GetZendeskClientSecretOk

`func (o *AlertserviceZendeskDetails) GetZendeskClientSecretOk() (*string, bool)`

GetZendeskClientSecretOk returns a tuple with the ZendeskClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskClientSecret

`func (o *AlertserviceZendeskDetails) SetZendeskClientSecret(v string)`

SetZendeskClientSecret sets ZendeskClientSecret field to given value.


### GetZendeskEmail

`func (o *AlertserviceZendeskDetails) GetZendeskEmail() string`

GetZendeskEmail returns the ZendeskEmail field if non-nil, zero value otherwise.

### GetZendeskEmailOk

`func (o *AlertserviceZendeskDetails) GetZendeskEmailOk() (*string, bool)`

GetZendeskEmailOk returns a tuple with the ZendeskEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskEmail

`func (o *AlertserviceZendeskDetails) SetZendeskEmail(v string)`

SetZendeskEmail sets ZendeskEmail field to given value.

### HasZendeskEmail

`func (o *AlertserviceZendeskDetails) HasZendeskEmail() bool`

HasZendeskEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


