# V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCircuit** | Pointer to [**V1PortalApikeysPostRequest**](V1PortalApikeysPostRequest.md) |  | [optional] 
**BackupCircuitLabel** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionBackupCircuitLabel**](V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionBackupCircuitLabel.md) |  | [optional] 
**Egress** | Pointer to **string** |  | [optional] 
**Logging** | Pointer to **bool** |  | [optional] 
**PrimaryCircuit** | Pointer to [**V1PortalApikeysPostRequest**](V1PortalApikeysPostRequest.md) |  | [optional] 
**PrimaryCircuitLabel** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionBackupCircuitLabel**](V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionBackupCircuitLabel.md) |  | [optional] 
**Remark** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionRemark**](V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionRemark.md) |  | [optional] 
**SetSlaClass** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionSetSlaClass**](V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionSetSlaClass.md) |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction

`func NewV1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction() *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction`

NewV1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionWithDefaults

`func NewV1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionWithDefaults() *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction`

NewV1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionWithDefaults instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupCircuit

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetBackupCircuit() V1PortalApikeysPostRequest`

GetBackupCircuit returns the BackupCircuit field if non-nil, zero value otherwise.

### GetBackupCircuitOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetBackupCircuitOk() (*V1PortalApikeysPostRequest, bool)`

GetBackupCircuitOk returns a tuple with the BackupCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCircuit

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) SetBackupCircuit(v V1PortalApikeysPostRequest)`

SetBackupCircuit sets BackupCircuit field to given value.

### HasBackupCircuit

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) HasBackupCircuit() bool`

HasBackupCircuit returns a boolean if a field has been set.

### GetBackupCircuitLabel

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetBackupCircuitLabel() V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionBackupCircuitLabel`

GetBackupCircuitLabel returns the BackupCircuitLabel field if non-nil, zero value otherwise.

### GetBackupCircuitLabelOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetBackupCircuitLabelOk() (*V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionBackupCircuitLabel, bool)`

GetBackupCircuitLabelOk returns a tuple with the BackupCircuitLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCircuitLabel

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) SetBackupCircuitLabel(v V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionBackupCircuitLabel)`

SetBackupCircuitLabel sets BackupCircuitLabel field to given value.

### HasBackupCircuitLabel

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) HasBackupCircuitLabel() bool`

HasBackupCircuitLabel returns a boolean if a field has been set.

### GetEgress

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetEgress() string`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetEgressOk() (*string, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) SetEgress(v string)`

SetEgress sets Egress field to given value.

### HasEgress

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) HasEgress() bool`

HasEgress returns a boolean if a field has been set.

### GetLogging

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetLogging() bool`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetLoggingOk() (*bool, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) SetLogging(v bool)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetPrimaryCircuit

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetPrimaryCircuit() V1PortalApikeysPostRequest`

GetPrimaryCircuit returns the PrimaryCircuit field if non-nil, zero value otherwise.

### GetPrimaryCircuitOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetPrimaryCircuitOk() (*V1PortalApikeysPostRequest, bool)`

GetPrimaryCircuitOk returns a tuple with the PrimaryCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCircuit

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) SetPrimaryCircuit(v V1PortalApikeysPostRequest)`

SetPrimaryCircuit sets PrimaryCircuit field to given value.

### HasPrimaryCircuit

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) HasPrimaryCircuit() bool`

HasPrimaryCircuit returns a boolean if a field has been set.

### GetPrimaryCircuitLabel

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetPrimaryCircuitLabel() V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionBackupCircuitLabel`

GetPrimaryCircuitLabel returns the PrimaryCircuitLabel field if non-nil, zero value otherwise.

### GetPrimaryCircuitLabelOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetPrimaryCircuitLabelOk() (*V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionBackupCircuitLabel, bool)`

GetPrimaryCircuitLabelOk returns a tuple with the PrimaryCircuitLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCircuitLabel

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) SetPrimaryCircuitLabel(v V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionBackupCircuitLabel)`

SetPrimaryCircuitLabel sets PrimaryCircuitLabel field to given value.

### HasPrimaryCircuitLabel

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) HasPrimaryCircuitLabel() bool`

HasPrimaryCircuitLabel returns a boolean if a field has been set.

### GetRemark

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetRemark() V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionRemark`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetRemarkOk() (*V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionRemark, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) SetRemark(v V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionRemark)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetSetSlaClass

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetSetSlaClass() V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionSetSlaClass`

GetSetSlaClass returns the SetSlaClass field if non-nil, zero value otherwise.

### GetSetSlaClassOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) GetSetSlaClassOk() (*V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionSetSlaClass, bool)`

GetSetSlaClassOk returns a tuple with the SetSlaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSlaClass

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) SetSetSlaClass(v V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleActionSetSlaClass)`

SetSetSlaClass sets SetSlaClass field to given value.

### HasSetSlaClass

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValueRulesetRulesValueRuleAction) HasSetSlaClass() bool`

HasSetSlaClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


