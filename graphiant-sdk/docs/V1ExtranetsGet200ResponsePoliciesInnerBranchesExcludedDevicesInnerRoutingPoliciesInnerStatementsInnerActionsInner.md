# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativeDistance** | Pointer to **int32** |  | [optional] 
**AspathPrepend** | Pointer to **int32** |  | [optional] 
**BgpSetNextHop** | Pointer to **string** |  | [optional] 
**CallPolicy** | Pointer to **string** |  | [optional] 
**Community** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCommunitiesCommunity**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCommunitiesCommunity.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LocalPref** | Pointer to **int32** |  | [optional] 
**MetricAbsolute** | Pointer to **int32** |  | [optional] 
**MetricModifier** | Pointer to **int32** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInnerWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativeDistance

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetAdministrativeDistance() int32`

GetAdministrativeDistance returns the AdministrativeDistance field if non-nil, zero value otherwise.

### GetAdministrativeDistanceOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetAdministrativeDistanceOk() (*int32, bool)`

GetAdministrativeDistanceOk returns a tuple with the AdministrativeDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeDistance

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetAdministrativeDistance(v int32)`

SetAdministrativeDistance sets AdministrativeDistance field to given value.

### HasAdministrativeDistance

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasAdministrativeDistance() bool`

HasAdministrativeDistance returns a boolean if a field has been set.

### GetAspathPrepend

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetAspathPrepend() int32`

GetAspathPrepend returns the AspathPrepend field if non-nil, zero value otherwise.

### GetAspathPrependOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetAspathPrependOk() (*int32, bool)`

GetAspathPrependOk returns a tuple with the AspathPrepend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspathPrepend

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetAspathPrepend(v int32)`

SetAspathPrepend sets AspathPrepend field to given value.

### HasAspathPrepend

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasAspathPrepend() bool`

HasAspathPrepend returns a boolean if a field has been set.

### GetBgpSetNextHop

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetBgpSetNextHop() string`

GetBgpSetNextHop returns the BgpSetNextHop field if non-nil, zero value otherwise.

### GetBgpSetNextHopOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetBgpSetNextHopOk() (*string, bool)`

GetBgpSetNextHopOk returns a tuple with the BgpSetNextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpSetNextHop

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetBgpSetNextHop(v string)`

SetBgpSetNextHop sets BgpSetNextHop field to given value.

### HasBgpSetNextHop

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasBgpSetNextHop() bool`

HasBgpSetNextHop returns a boolean if a field has been set.

### GetCallPolicy

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetCallPolicy() string`

GetCallPolicy returns the CallPolicy field if non-nil, zero value otherwise.

### GetCallPolicyOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetCallPolicyOk() (*string, bool)`

GetCallPolicyOk returns a tuple with the CallPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPolicy

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetCallPolicy(v string)`

SetCallPolicy sets CallPolicy field to given value.

### HasCallPolicy

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasCallPolicy() bool`

HasCallPolicy returns a boolean if a field has been set.

### GetCommunity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetCommunity() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCommunitiesCommunity`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetCommunityOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCommunitiesCommunity, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetCommunity(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionCommunitiesCommunity)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalPref

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetLocalPref() int32`

GetLocalPref returns the LocalPref field if non-nil, zero value otherwise.

### GetLocalPrefOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetLocalPrefOk() (*int32, bool)`

GetLocalPrefOk returns a tuple with the LocalPref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPref

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetLocalPref(v int32)`

SetLocalPref sets LocalPref field to given value.

### HasLocalPref

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasLocalPref() bool`

HasLocalPref returns a boolean if a field has been set.

### GetMetricAbsolute

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetMetricAbsolute() int32`

GetMetricAbsolute returns the MetricAbsolute field if non-nil, zero value otherwise.

### GetMetricAbsoluteOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetMetricAbsoluteOk() (*int32, bool)`

GetMetricAbsoluteOk returns a tuple with the MetricAbsolute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricAbsolute

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetMetricAbsolute(v int32)`

SetMetricAbsolute sets MetricAbsolute field to given value.

### HasMetricAbsolute

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasMetricAbsolute() bool`

HasMetricAbsolute returns a boolean if a field has been set.

### GetMetricModifier

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetMetricModifier() int32`

GetMetricModifier returns the MetricModifier field if non-nil, zero value otherwise.

### GetMetricModifierOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetMetricModifierOk() (*int32, bool)`

GetMetricModifierOk returns a tuple with the MetricModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricModifier

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetMetricModifier(v int32)`

SetMetricModifier sets MetricModifier field to given value.

### HasMetricModifier

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasMetricModifier() bool`

HasMetricModifier returns a boolean if a field has been set.

### GetResult

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSeq

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetWeight

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerActionsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


