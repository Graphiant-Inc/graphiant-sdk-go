# IamGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**EnterpriseIds** | Pointer to **[]int64** |  | [optional] 
**EnterprisePermissions** | Pointer to [**map[string]IamEnterprisePermissions**](IamEnterprisePermissions.md) |  | [optional] 
**GroupType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**CommonPermissions**](CommonPermissions.md) |  | [optional] 
**TimeWindowEnd** | Pointer to **int64** |  | [optional] 
**TimeWindowStart** | Pointer to **int64** |  | [optional] 

## Methods

### NewIamGroup

`func NewIamGroup() *IamGroup`

NewIamGroup instantiates a new IamGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamGroupWithDefaults

`func NewIamGroupWithDefaults() *IamGroup`

NewIamGroupWithDefaults instantiates a new IamGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IamGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnterpriseIds

`func (o *IamGroup) GetEnterpriseIds() []int64`

GetEnterpriseIds returns the EnterpriseIds field if non-nil, zero value otherwise.

### GetEnterpriseIdsOk

`func (o *IamGroup) GetEnterpriseIdsOk() (*[]int64, bool)`

GetEnterpriseIdsOk returns a tuple with the EnterpriseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseIds

`func (o *IamGroup) SetEnterpriseIds(v []int64)`

SetEnterpriseIds sets EnterpriseIds field to given value.

### HasEnterpriseIds

`func (o *IamGroup) HasEnterpriseIds() bool`

HasEnterpriseIds returns a boolean if a field has been set.

### GetEnterprisePermissions

`func (o *IamGroup) GetEnterprisePermissions() map[string]IamEnterprisePermissions`

GetEnterprisePermissions returns the EnterprisePermissions field if non-nil, zero value otherwise.

### GetEnterprisePermissionsOk

`func (o *IamGroup) GetEnterprisePermissionsOk() (*map[string]IamEnterprisePermissions, bool)`

GetEnterprisePermissionsOk returns a tuple with the EnterprisePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprisePermissions

`func (o *IamGroup) SetEnterprisePermissions(v map[string]IamEnterprisePermissions)`

SetEnterprisePermissions sets EnterprisePermissions field to given value.

### HasEnterprisePermissions

`func (o *IamGroup) HasEnterprisePermissions() bool`

HasEnterprisePermissions returns a boolean if a field has been set.

### GetGroupType

`func (o *IamGroup) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *IamGroup) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *IamGroup) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *IamGroup) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetId

`func (o *IamGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IamGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *IamGroup) GetPermissions() CommonPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamGroup) GetPermissionsOk() (*CommonPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamGroup) SetPermissions(v CommonPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamGroup) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTimeWindowEnd

`func (o *IamGroup) GetTimeWindowEnd() int64`

GetTimeWindowEnd returns the TimeWindowEnd field if non-nil, zero value otherwise.

### GetTimeWindowEndOk

`func (o *IamGroup) GetTimeWindowEndOk() (*int64, bool)`

GetTimeWindowEndOk returns a tuple with the TimeWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowEnd

`func (o *IamGroup) SetTimeWindowEnd(v int64)`

SetTimeWindowEnd sets TimeWindowEnd field to given value.

### HasTimeWindowEnd

`func (o *IamGroup) HasTimeWindowEnd() bool`

HasTimeWindowEnd returns a boolean if a field has been set.

### GetTimeWindowStart

`func (o *IamGroup) GetTimeWindowStart() int64`

GetTimeWindowStart returns the TimeWindowStart field if non-nil, zero value otherwise.

### GetTimeWindowStartOk

`func (o *IamGroup) GetTimeWindowStartOk() (*int64, bool)`

GetTimeWindowStartOk returns a tuple with the TimeWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowStart

`func (o *IamGroup) SetTimeWindowStart(v int64)`

SetTimeWindowStart sets TimeWindowStart field to given value.

### HasTimeWindowStart

`func (o *IamGroup) HasTimeWindowStart() bool`

HasTimeWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


