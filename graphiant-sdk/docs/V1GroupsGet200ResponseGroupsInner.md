# V1GroupsGet200ResponseGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**EnterpriseIds** | Pointer to **[]int64** |  | [optional] 
**GroupType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**V1GroupsGet200ResponseGroupsInnerPermissions**](V1GroupsGet200ResponseGroupsInnerPermissions.md) |  | [optional] 
**TimeWindowEnd** | Pointer to **int64** |  | [optional] 
**TimeWindowStart** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1GroupsGet200ResponseGroupsInner

`func NewV1GroupsGet200ResponseGroupsInner() *V1GroupsGet200ResponseGroupsInner`

NewV1GroupsGet200ResponseGroupsInner instantiates a new V1GroupsGet200ResponseGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GroupsGet200ResponseGroupsInnerWithDefaults

`func NewV1GroupsGet200ResponseGroupsInnerWithDefaults() *V1GroupsGet200ResponseGroupsInner`

NewV1GroupsGet200ResponseGroupsInnerWithDefaults instantiates a new V1GroupsGet200ResponseGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1GroupsGet200ResponseGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1GroupsGet200ResponseGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1GroupsGet200ResponseGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1GroupsGet200ResponseGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnterpriseIds

`func (o *V1GroupsGet200ResponseGroupsInner) GetEnterpriseIds() []int64`

GetEnterpriseIds returns the EnterpriseIds field if non-nil, zero value otherwise.

### GetEnterpriseIdsOk

`func (o *V1GroupsGet200ResponseGroupsInner) GetEnterpriseIdsOk() (*[]int64, bool)`

GetEnterpriseIdsOk returns a tuple with the EnterpriseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseIds

`func (o *V1GroupsGet200ResponseGroupsInner) SetEnterpriseIds(v []int64)`

SetEnterpriseIds sets EnterpriseIds field to given value.

### HasEnterpriseIds

`func (o *V1GroupsGet200ResponseGroupsInner) HasEnterpriseIds() bool`

HasEnterpriseIds returns a boolean if a field has been set.

### GetGroupType

`func (o *V1GroupsGet200ResponseGroupsInner) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *V1GroupsGet200ResponseGroupsInner) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *V1GroupsGet200ResponseGroupsInner) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *V1GroupsGet200ResponseGroupsInner) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetId

`func (o *V1GroupsGet200ResponseGroupsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1GroupsGet200ResponseGroupsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1GroupsGet200ResponseGroupsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1GroupsGet200ResponseGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1GroupsGet200ResponseGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GroupsGet200ResponseGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GroupsGet200ResponseGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GroupsGet200ResponseGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *V1GroupsGet200ResponseGroupsInner) GetPermissions() V1GroupsGet200ResponseGroupsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *V1GroupsGet200ResponseGroupsInner) GetPermissionsOk() (*V1GroupsGet200ResponseGroupsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *V1GroupsGet200ResponseGroupsInner) SetPermissions(v V1GroupsGet200ResponseGroupsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *V1GroupsGet200ResponseGroupsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTimeWindowEnd

`func (o *V1GroupsGet200ResponseGroupsInner) GetTimeWindowEnd() int64`

GetTimeWindowEnd returns the TimeWindowEnd field if non-nil, zero value otherwise.

### GetTimeWindowEndOk

`func (o *V1GroupsGet200ResponseGroupsInner) GetTimeWindowEndOk() (*int64, bool)`

GetTimeWindowEndOk returns a tuple with the TimeWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowEnd

`func (o *V1GroupsGet200ResponseGroupsInner) SetTimeWindowEnd(v int64)`

SetTimeWindowEnd sets TimeWindowEnd field to given value.

### HasTimeWindowEnd

`func (o *V1GroupsGet200ResponseGroupsInner) HasTimeWindowEnd() bool`

HasTimeWindowEnd returns a boolean if a field has been set.

### GetTimeWindowStart

`func (o *V1GroupsGet200ResponseGroupsInner) GetTimeWindowStart() int64`

GetTimeWindowStart returns the TimeWindowStart field if non-nil, zero value otherwise.

### GetTimeWindowStartOk

`func (o *V1GroupsGet200ResponseGroupsInner) GetTimeWindowStartOk() (*int64, bool)`

GetTimeWindowStartOk returns a tuple with the TimeWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowStart

`func (o *V1GroupsGet200ResponseGroupsInner) SetTimeWindowStart(v int64)`

SetTimeWindowStart sets TimeWindowStart field to given value.

### HasTimeWindowStart

`func (o *V1GroupsGet200ResponseGroupsInner) HasTimeWindowStart() bool`

HasTimeWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


