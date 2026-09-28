# CommonUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BearerToken** | Pointer to **string** |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**Exp** | Pointer to **int64** |  | [optional] 
**OriginalEnterpriseId** | Pointer to **int64** |  | [optional] 
**Permissions** | Pointer to [**CommonPermissions**](CommonPermissions.md) |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewCommonUserInfo

`func NewCommonUserInfo() *CommonUserInfo`

NewCommonUserInfo instantiates a new CommonUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonUserInfoWithDefaults

`func NewCommonUserInfoWithDefaults() *CommonUserInfo`

NewCommonUserInfoWithDefaults instantiates a new CommonUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBearerToken

`func (o *CommonUserInfo) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *CommonUserInfo) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *CommonUserInfo) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *CommonUserInfo) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *CommonUserInfo) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *CommonUserInfo) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *CommonUserInfo) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *CommonUserInfo) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetExp

`func (o *CommonUserInfo) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *CommonUserInfo) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *CommonUserInfo) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *CommonUserInfo) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetOriginalEnterpriseId

`func (o *CommonUserInfo) GetOriginalEnterpriseId() int64`

GetOriginalEnterpriseId returns the OriginalEnterpriseId field if non-nil, zero value otherwise.

### GetOriginalEnterpriseIdOk

`func (o *CommonUserInfo) GetOriginalEnterpriseIdOk() (*int64, bool)`

GetOriginalEnterpriseIdOk returns a tuple with the OriginalEnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEnterpriseId

`func (o *CommonUserInfo) SetOriginalEnterpriseId(v int64)`

SetOriginalEnterpriseId sets OriginalEnterpriseId field to given value.

### HasOriginalEnterpriseId

`func (o *CommonUserInfo) HasOriginalEnterpriseId() bool`

HasOriginalEnterpriseId returns a boolean if a field has been set.

### GetPermissions

`func (o *CommonUserInfo) GetPermissions() CommonPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CommonUserInfo) GetPermissionsOk() (*CommonPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CommonUserInfo) SetPermissions(v CommonPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CommonUserInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTimeZone

`func (o *CommonUserInfo) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CommonUserInfo) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CommonUserInfo) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *CommonUserInfo) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUserId

`func (o *CommonUserInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CommonUserInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CommonUserInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CommonUserInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


