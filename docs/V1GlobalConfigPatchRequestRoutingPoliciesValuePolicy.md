# V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachPoint** | Pointer to **string** |  | [optional] 
**DefaultAction** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GlobalId** | Pointer to **int64** |  | [optional] 
**IsGlobalSync** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Statements** | Pointer to [**map[string]V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValue**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValue.md) |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicy

`func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicy() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy`

NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicy instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyWithDefaults

`func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyWithDefaults() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy`

NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyWithDefaults instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachPoint

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetAttachPoint() string`

GetAttachPoint returns the AttachPoint field if non-nil, zero value otherwise.

### GetAttachPointOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetAttachPointOk() (*string, bool)`

GetAttachPointOk returns a tuple with the AttachPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachPoint

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) SetAttachPoint(v string)`

SetAttachPoint sets AttachPoint field to given value.

### HasAttachPoint

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) HasAttachPoint() bool`

HasAttachPoint returns a boolean if a field has been set.

### GetDefaultAction

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetDescription

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGlobalId

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetIsGlobalSync

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetIsGlobalSync() bool`

GetIsGlobalSync returns the IsGlobalSync field if non-nil, zero value otherwise.

### GetIsGlobalSyncOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetIsGlobalSyncOk() (*bool, bool)`

GetIsGlobalSyncOk returns a tuple with the IsGlobalSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalSync

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) SetIsGlobalSync(v bool)`

SetIsGlobalSync sets IsGlobalSync field to given value.

### HasIsGlobalSync

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) HasIsGlobalSync() bool`

HasIsGlobalSync returns a boolean if a field has been set.

### GetName

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatements

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetStatements() map[string]V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValue`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) GetStatementsOk() (*map[string]V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValue, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) SetStatements(v map[string]V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValue)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicy) HasStatements() bool`

HasStatements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


