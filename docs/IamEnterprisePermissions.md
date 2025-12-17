# IamEnterprisePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupType** | Pointer to **string** |  (required) | [optional] 
**Permissions** | Pointer to [**CommonPermissions**](CommonPermissions.md) |  | [optional] 

## Methods

### NewIamEnterprisePermissions

`func NewIamEnterprisePermissions() *IamEnterprisePermissions`

NewIamEnterprisePermissions instantiates a new IamEnterprisePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEnterprisePermissionsWithDefaults

`func NewIamEnterprisePermissionsWithDefaults() *IamEnterprisePermissions`

NewIamEnterprisePermissionsWithDefaults instantiates a new IamEnterprisePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupType

`func (o *IamEnterprisePermissions) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *IamEnterprisePermissions) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *IamEnterprisePermissions) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *IamEnterprisePermissions) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetPermissions

`func (o *IamEnterprisePermissions) GetPermissions() CommonPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamEnterprisePermissions) GetPermissionsOk() (*CommonPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamEnterprisePermissions) SetPermissions(v CommonPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamEnterprisePermissions) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


