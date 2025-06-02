# V1GroupsIdPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**GroupType** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**V1GroupsGet200ResponseGroupsInnerPermissions**](V1GroupsGet200ResponseGroupsInnerPermissions.md) |  | [optional] 
**TimeWindowEnd** | Pointer to **int64** |  | [optional] 
**TimeWindowStart** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1GroupsIdPatchRequest

`func NewV1GroupsIdPatchRequest() *V1GroupsIdPatchRequest`

NewV1GroupsIdPatchRequest instantiates a new V1GroupsIdPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GroupsIdPatchRequestWithDefaults

`func NewV1GroupsIdPatchRequestWithDefaults() *V1GroupsIdPatchRequest`

NewV1GroupsIdPatchRequestWithDefaults instantiates a new V1GroupsIdPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1GroupsIdPatchRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1GroupsIdPatchRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1GroupsIdPatchRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1GroupsIdPatchRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *V1GroupsIdPatchRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *V1GroupsIdPatchRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *V1GroupsIdPatchRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *V1GroupsIdPatchRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGroupType

`func (o *V1GroupsIdPatchRequest) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *V1GroupsIdPatchRequest) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *V1GroupsIdPatchRequest) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *V1GroupsIdPatchRequest) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetPermissions

`func (o *V1GroupsIdPatchRequest) GetPermissions() V1GroupsGet200ResponseGroupsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *V1GroupsIdPatchRequest) GetPermissionsOk() (*V1GroupsGet200ResponseGroupsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *V1GroupsIdPatchRequest) SetPermissions(v V1GroupsGet200ResponseGroupsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *V1GroupsIdPatchRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTimeWindowEnd

`func (o *V1GroupsIdPatchRequest) GetTimeWindowEnd() int64`

GetTimeWindowEnd returns the TimeWindowEnd field if non-nil, zero value otherwise.

### GetTimeWindowEndOk

`func (o *V1GroupsIdPatchRequest) GetTimeWindowEndOk() (*int64, bool)`

GetTimeWindowEndOk returns a tuple with the TimeWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowEnd

`func (o *V1GroupsIdPatchRequest) SetTimeWindowEnd(v int64)`

SetTimeWindowEnd sets TimeWindowEnd field to given value.

### HasTimeWindowEnd

`func (o *V1GroupsIdPatchRequest) HasTimeWindowEnd() bool`

HasTimeWindowEnd returns a boolean if a field has been set.

### GetTimeWindowStart

`func (o *V1GroupsIdPatchRequest) GetTimeWindowStart() int64`

GetTimeWindowStart returns the TimeWindowStart field if non-nil, zero value otherwise.

### GetTimeWindowStartOk

`func (o *V1GroupsIdPatchRequest) GetTimeWindowStartOk() (*int64, bool)`

GetTimeWindowStartOk returns a tuple with the TimeWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowStart

`func (o *V1GroupsIdPatchRequest) SetTimeWindowStart(v int64)`

SetTimeWindowStart sets TimeWindowStart field to given value.

### HasTimeWindowStart

`func (o *V1GroupsIdPatchRequest) HasTimeWindowStart() bool`

HasTimeWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


