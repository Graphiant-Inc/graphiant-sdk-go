# V2AuditLogsPostRequestSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** |  | [optional] 
**Entities** | Pointer to [**[]V1ActivityLogsPostRequestSelectorJobEntity**](V1ActivityLogsPostRequestSelectorJobEntity.md) |  | [optional] 
**JobTypes** | Pointer to **[]string** |  | [optional] 
**Targets** | Pointer to [**[]V1ActivityLogsPostRequestSelectorJobEntity**](V1ActivityLogsPostRequestSelectorJobEntity.md) |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV2AuditLogsPostRequestSelector

`func NewV2AuditLogsPostRequestSelector() *V2AuditLogsPostRequestSelector`

NewV2AuditLogsPostRequestSelector instantiates a new V2AuditLogsPostRequestSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AuditLogsPostRequestSelectorWithDefaults

`func NewV2AuditLogsPostRequestSelectorWithDefaults() *V2AuditLogsPostRequestSelector`

NewV2AuditLogsPostRequestSelectorWithDefaults instantiates a new V2AuditLogsPostRequestSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *V2AuditLogsPostRequestSelector) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *V2AuditLogsPostRequestSelector) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *V2AuditLogsPostRequestSelector) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *V2AuditLogsPostRequestSelector) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetEntities

`func (o *V2AuditLogsPostRequestSelector) GetEntities() []V1ActivityLogsPostRequestSelectorJobEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *V2AuditLogsPostRequestSelector) GetEntitiesOk() (*[]V1ActivityLogsPostRequestSelectorJobEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *V2AuditLogsPostRequestSelector) SetEntities(v []V1ActivityLogsPostRequestSelectorJobEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *V2AuditLogsPostRequestSelector) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetJobTypes

`func (o *V2AuditLogsPostRequestSelector) GetJobTypes() []string`

GetJobTypes returns the JobTypes field if non-nil, zero value otherwise.

### GetJobTypesOk

`func (o *V2AuditLogsPostRequestSelector) GetJobTypesOk() (*[]string, bool)`

GetJobTypesOk returns a tuple with the JobTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypes

`func (o *V2AuditLogsPostRequestSelector) SetJobTypes(v []string)`

SetJobTypes sets JobTypes field to given value.

### HasJobTypes

`func (o *V2AuditLogsPostRequestSelector) HasJobTypes() bool`

HasJobTypes returns a boolean if a field has been set.

### GetTargets

`func (o *V2AuditLogsPostRequestSelector) GetTargets() []V1ActivityLogsPostRequestSelectorJobEntity`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *V2AuditLogsPostRequestSelector) GetTargetsOk() (*[]V1ActivityLogsPostRequestSelectorJobEntity, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *V2AuditLogsPostRequestSelector) SetTargets(v []V1ActivityLogsPostRequestSelectorJobEntity)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *V2AuditLogsPostRequestSelector) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetUsers

`func (o *V2AuditLogsPostRequestSelector) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V2AuditLogsPostRequestSelector) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V2AuditLogsPostRequestSelector) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V2AuditLogsPostRequestSelector) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


