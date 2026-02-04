# V1AuthLoginPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** | User email address (returned for MFA users) | [optional] 
**MfaType** | Pointer to **string** | MFA type (returned for MFA users) | [optional] 
**StateToken** | Pointer to **string** | State token for MFA verification (returned for MFA users) | [optional] 
**Status** | Pointer to **string** | Authentication status (returned for MFA users) | [optional] 

## Methods

### NewV1AuthLoginPostResponse

`func NewV1AuthLoginPostResponse() *V1AuthLoginPostResponse`

NewV1AuthLoginPostResponse instantiates a new V1AuthLoginPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AuthLoginPostResponseWithDefaults

`func NewV1AuthLoginPostResponseWithDefaults() *V1AuthLoginPostResponse`

NewV1AuthLoginPostResponseWithDefaults instantiates a new V1AuthLoginPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *V1AuthLoginPostResponse) GetAuth() bool`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *V1AuthLoginPostResponse) GetAuthOk() (*bool, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *V1AuthLoginPostResponse) SetAuth(v bool)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *V1AuthLoginPostResponse) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetToken

`func (o *V1AuthLoginPostResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *V1AuthLoginPostResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *V1AuthLoginPostResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *V1AuthLoginPostResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAccountType

`func (o *V1AuthLoginPostResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *V1AuthLoginPostResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *V1AuthLoginPostResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *V1AuthLoginPostResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetEmail

`func (o *V1AuthLoginPostResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1AuthLoginPostResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1AuthLoginPostResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *V1AuthLoginPostResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMfaType

`func (o *V1AuthLoginPostResponse) GetMfaType() string`

GetMfaType returns the MfaType field if non-nil, zero value otherwise.

### GetMfaTypeOk

`func (o *V1AuthLoginPostResponse) GetMfaTypeOk() (*string, bool)`

GetMfaTypeOk returns a tuple with the MfaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaType

`func (o *V1AuthLoginPostResponse) SetMfaType(v string)`

SetMfaType sets MfaType field to given value.

### HasMfaType

`func (o *V1AuthLoginPostResponse) HasMfaType() bool`

HasMfaType returns a boolean if a field has been set.

### GetStateToken

`func (o *V1AuthLoginPostResponse) GetStateToken() string`

GetStateToken returns the StateToken field if non-nil, zero value otherwise.

### GetStateTokenOk

`func (o *V1AuthLoginPostResponse) GetStateTokenOk() (*string, bool)`

GetStateTokenOk returns a tuple with the StateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateToken

`func (o *V1AuthLoginPostResponse) SetStateToken(v string)`

SetStateToken sets StateToken field to given value.

### HasStateToken

`func (o *V1AuthLoginPostResponse) HasStateToken() bool`

HasStateToken returns a boolean if a field has been set.

### GetStatus

`func (o *V1AuthLoginPostResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1AuthLoginPostResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1AuthLoginPostResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1AuthLoginPostResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


