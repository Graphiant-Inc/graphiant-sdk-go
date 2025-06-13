# V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativeDistance** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAdministrativeDistance**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAdministrativeDistance.md) |  | [optional] 
**AspathPrepend** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAspathPrepend**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAspathPrepend.md) |  | [optional] 
**BgpSetNextHop** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionBgpSetNextHop**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionBgpSetNextHop.md) |  | [optional] 
**CallPolicy** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCallPolicy**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCallPolicy.md) |  | [optional] 
**Communities** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCommunities**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCommunities.md) |  | [optional] 
**LocalPref** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionLocalPref**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionLocalPref.md) |  | [optional] 
**Metric** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionMetric**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionMetric.md) |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int32** |  | [optional] 
**Weight** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionWeight**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionWeight.md) |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction

`func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction`

NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionWithDefaults

`func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionWithDefaults() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction`

NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionWithDefaults instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativeDistance

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetAdministrativeDistance() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAdministrativeDistance`

GetAdministrativeDistance returns the AdministrativeDistance field if non-nil, zero value otherwise.

### GetAdministrativeDistanceOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetAdministrativeDistanceOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAdministrativeDistance, bool)`

GetAdministrativeDistanceOk returns a tuple with the AdministrativeDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeDistance

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) SetAdministrativeDistance(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAdministrativeDistance)`

SetAdministrativeDistance sets AdministrativeDistance field to given value.

### HasAdministrativeDistance

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) HasAdministrativeDistance() bool`

HasAdministrativeDistance returns a boolean if a field has been set.

### GetAspathPrepend

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetAspathPrepend() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAspathPrepend`

GetAspathPrepend returns the AspathPrepend field if non-nil, zero value otherwise.

### GetAspathPrependOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetAspathPrependOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAspathPrepend, bool)`

GetAspathPrependOk returns a tuple with the AspathPrepend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspathPrepend

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) SetAspathPrepend(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAspathPrepend)`

SetAspathPrepend sets AspathPrepend field to given value.

### HasAspathPrepend

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) HasAspathPrepend() bool`

HasAspathPrepend returns a boolean if a field has been set.

### GetBgpSetNextHop

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetBgpSetNextHop() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionBgpSetNextHop`

GetBgpSetNextHop returns the BgpSetNextHop field if non-nil, zero value otherwise.

### GetBgpSetNextHopOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetBgpSetNextHopOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionBgpSetNextHop, bool)`

GetBgpSetNextHopOk returns a tuple with the BgpSetNextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpSetNextHop

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) SetBgpSetNextHop(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionBgpSetNextHop)`

SetBgpSetNextHop sets BgpSetNextHop field to given value.

### HasBgpSetNextHop

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) HasBgpSetNextHop() bool`

HasBgpSetNextHop returns a boolean if a field has been set.

### GetCallPolicy

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetCallPolicy() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCallPolicy`

GetCallPolicy returns the CallPolicy field if non-nil, zero value otherwise.

### GetCallPolicyOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetCallPolicyOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCallPolicy, bool)`

GetCallPolicyOk returns a tuple with the CallPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPolicy

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) SetCallPolicy(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCallPolicy)`

SetCallPolicy sets CallPolicy field to given value.

### HasCallPolicy

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) HasCallPolicy() bool`

HasCallPolicy returns a boolean if a field has been set.

### GetCommunities

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetCommunities() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCommunities`

GetCommunities returns the Communities field if non-nil, zero value otherwise.

### GetCommunitiesOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetCommunitiesOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCommunities, bool)`

GetCommunitiesOk returns a tuple with the Communities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunities

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) SetCommunities(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCommunities)`

SetCommunities sets Communities field to given value.

### HasCommunities

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) HasCommunities() bool`

HasCommunities returns a boolean if a field has been set.

### GetLocalPref

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetLocalPref() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionLocalPref`

GetLocalPref returns the LocalPref field if non-nil, zero value otherwise.

### GetLocalPrefOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetLocalPrefOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionLocalPref, bool)`

GetLocalPrefOk returns a tuple with the LocalPref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPref

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) SetLocalPref(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionLocalPref)`

SetLocalPref sets LocalPref field to given value.

### HasLocalPref

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) HasLocalPref() bool`

HasLocalPref returns a boolean if a field has been set.

### GetMetric

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetMetric() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetMetricOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) SetMetric(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetResult

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSeq

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetWeight

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetWeight() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionWeight`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) GetWeightOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionWeight, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) SetWeight(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionWeight)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueAction) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


