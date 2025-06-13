# V1GroupsPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**GroupType** | Pointer to **string** |  | [optional] 
**ManagesEnterprises** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**V1GroupsGet200ResponseGroupsInnerPermissions**](V1GroupsGet200ResponseGroupsInnerPermissions.md) |  | [optional] 
**TimeWindowEnd** | Pointer to **int64** |  | [optional] 
**TimeWindowStart** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1GroupsPutRequest

`func NewV1GroupsPutRequest() *V1GroupsPutRequest`

NewV1GroupsPutRequest instantiates a new V1GroupsPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GroupsPutRequestWithDefaults

`func NewV1GroupsPutRequestWithDefaults() *V1GroupsPutRequest`

NewV1GroupsPutRequestWithDefaults instantiates a new V1GroupsPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1GroupsPutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1GroupsPutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1GroupsPutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1GroupsPutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupId

`func (o *V1GroupsPutRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V1GroupsPutRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V1GroupsPutRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V1GroupsPutRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupType

`func (o *V1GroupsPutRequest) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *V1GroupsPutRequest) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *V1GroupsPutRequest) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *V1GroupsPutRequest) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetManagesEnterprises

`func (o *V1GroupsPutRequest) GetManagesEnterprises() bool`

GetManagesEnterprises returns the ManagesEnterprises field if non-nil, zero value otherwise.

### GetManagesEnterprisesOk

`func (o *V1GroupsPutRequest) GetManagesEnterprisesOk() (*bool, bool)`

GetManagesEnterprisesOk returns a tuple with the ManagesEnterprises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagesEnterprises

`func (o *V1GroupsPutRequest) SetManagesEnterprises(v bool)`

SetManagesEnterprises sets ManagesEnterprises field to given value.

### HasManagesEnterprises

`func (o *V1GroupsPutRequest) HasManagesEnterprises() bool`

HasManagesEnterprises returns a boolean if a field has been set.

### GetName

`func (o *V1GroupsPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GroupsPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GroupsPutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GroupsPutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *V1GroupsPutRequest) GetPermissions() V1GroupsGet200ResponseGroupsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *V1GroupsPutRequest) GetPermissionsOk() (*V1GroupsGet200ResponseGroupsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *V1GroupsPutRequest) SetPermissions(v V1GroupsGet200ResponseGroupsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *V1GroupsPutRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTimeWindowEnd

`func (o *V1GroupsPutRequest) GetTimeWindowEnd() int64`

GetTimeWindowEnd returns the TimeWindowEnd field if non-nil, zero value otherwise.

### GetTimeWindowEndOk

`func (o *V1GroupsPutRequest) GetTimeWindowEndOk() (*int64, bool)`

GetTimeWindowEndOk returns a tuple with the TimeWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowEnd

`func (o *V1GroupsPutRequest) SetTimeWindowEnd(v int64)`

SetTimeWindowEnd sets TimeWindowEnd field to given value.

### HasTimeWindowEnd

`func (o *V1GroupsPutRequest) HasTimeWindowEnd() bool`

HasTimeWindowEnd returns a boolean if a field has been set.

### GetTimeWindowStart

`func (o *V1GroupsPutRequest) GetTimeWindowStart() int64`

GetTimeWindowStart returns the TimeWindowStart field if non-nil, zero value otherwise.

### GetTimeWindowStartOk

`func (o *V1GroupsPutRequest) GetTimeWindowStartOk() (*int64, bool)`

GetTimeWindowStartOk returns a tuple with the TimeWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowStart

`func (o *V1GroupsPutRequest) SetTimeWindowStart(v int64)`

SetTimeWindowStart sets TimeWindowStart field to given value.

### HasTimeWindowStart

`func (o *V1GroupsPutRequest) HasTimeWindowStart() bool`

HasTimeWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


