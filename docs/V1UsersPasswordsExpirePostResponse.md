# V1UsersPasswordsExpirePostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedCount** | Pointer to **int32** |  | [optional] 
**FailedUsers** | Pointer to [**[]IamFailedUser**](IamFailedUser.md) |  | [optional] 
**SuccessCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1UsersPasswordsExpirePostResponse

`func NewV1UsersPasswordsExpirePostResponse() *V1UsersPasswordsExpirePostResponse`

NewV1UsersPasswordsExpirePostResponse instantiates a new V1UsersPasswordsExpirePostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UsersPasswordsExpirePostResponseWithDefaults

`func NewV1UsersPasswordsExpirePostResponseWithDefaults() *V1UsersPasswordsExpirePostResponse`

NewV1UsersPasswordsExpirePostResponseWithDefaults instantiates a new V1UsersPasswordsExpirePostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedCount

`func (o *V1UsersPasswordsExpirePostResponse) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *V1UsersPasswordsExpirePostResponse) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *V1UsersPasswordsExpirePostResponse) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *V1UsersPasswordsExpirePostResponse) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetFailedUsers

`func (o *V1UsersPasswordsExpirePostResponse) GetFailedUsers() []IamFailedUser`

GetFailedUsers returns the FailedUsers field if non-nil, zero value otherwise.

### GetFailedUsersOk

`func (o *V1UsersPasswordsExpirePostResponse) GetFailedUsersOk() (*[]IamFailedUser, bool)`

GetFailedUsersOk returns a tuple with the FailedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedUsers

`func (o *V1UsersPasswordsExpirePostResponse) SetFailedUsers(v []IamFailedUser)`

SetFailedUsers sets FailedUsers field to given value.

### HasFailedUsers

`func (o *V1UsersPasswordsExpirePostResponse) HasFailedUsers() bool`

HasFailedUsers returns a boolean if a field has been set.

### GetSuccessCount

`func (o *V1UsersPasswordsExpirePostResponse) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *V1UsersPasswordsExpirePostResponse) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *V1UsersPasswordsExpirePostResponse) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *V1UsersPasswordsExpirePostResponse) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


