# V1ExtranetsPostRequestPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auto** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerAuto**](V1ExtranetsGet200ResponsePoliciesInnerAuto.md) |  | [optional] 
**Branches** | Pointer to [**V1ExtranetsPostRequestPolicyBranches**](V1ExtranetsPostRequestPolicyBranches.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HostPrefixSet** | Pointer to [**V1ExtranetsPostRequestPolicyBranchesPrefixSet**](V1ExtranetsPostRequestPolicyBranchesPrefixSet.md) |  | [optional] 
**Manual** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerManual**](V1ExtranetsGet200ResponsePoliciesInnerManual.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SharedPrefixes** | Pointer to **[]string** |  | [optional] 
**SharedSegment** | Pointer to **int64** |  | [optional] 
**Source** | Pointer to [**V1ExtranetsPostRequestPolicyBranches**](V1ExtranetsPostRequestPolicyBranches.md) |  | [optional] 
**TargetSegments** | Pointer to **[]int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ExtranetsPostRequestPolicy

`func NewV1ExtranetsPostRequestPolicy() *V1ExtranetsPostRequestPolicy`

NewV1ExtranetsPostRequestPolicy instantiates a new V1ExtranetsPostRequestPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsPostRequestPolicyWithDefaults

`func NewV1ExtranetsPostRequestPolicyWithDefaults() *V1ExtranetsPostRequestPolicy`

NewV1ExtranetsPostRequestPolicyWithDefaults instantiates a new V1ExtranetsPostRequestPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuto

`func (o *V1ExtranetsPostRequestPolicy) GetAuto() V1ExtranetsGet200ResponsePoliciesInnerAuto`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *V1ExtranetsPostRequestPolicy) GetAutoOk() (*V1ExtranetsGet200ResponsePoliciesInnerAuto, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *V1ExtranetsPostRequestPolicy) SetAuto(v V1ExtranetsGet200ResponsePoliciesInnerAuto)`

SetAuto sets Auto field to given value.

### HasAuto

`func (o *V1ExtranetsPostRequestPolicy) HasAuto() bool`

HasAuto returns a boolean if a field has been set.

### GetBranches

`func (o *V1ExtranetsPostRequestPolicy) GetBranches() V1ExtranetsPostRequestPolicyBranches`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *V1ExtranetsPostRequestPolicy) GetBranchesOk() (*V1ExtranetsPostRequestPolicyBranches, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *V1ExtranetsPostRequestPolicy) SetBranches(v V1ExtranetsPostRequestPolicyBranches)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *V1ExtranetsPostRequestPolicy) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetDescription

`func (o *V1ExtranetsPostRequestPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1ExtranetsPostRequestPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1ExtranetsPostRequestPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1ExtranetsPostRequestPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostPrefixSet

`func (o *V1ExtranetsPostRequestPolicy) GetHostPrefixSet() V1ExtranetsPostRequestPolicyBranchesPrefixSet`

GetHostPrefixSet returns the HostPrefixSet field if non-nil, zero value otherwise.

### GetHostPrefixSetOk

`func (o *V1ExtranetsPostRequestPolicy) GetHostPrefixSetOk() (*V1ExtranetsPostRequestPolicyBranchesPrefixSet, bool)`

GetHostPrefixSetOk returns a tuple with the HostPrefixSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPrefixSet

`func (o *V1ExtranetsPostRequestPolicy) SetHostPrefixSet(v V1ExtranetsPostRequestPolicyBranchesPrefixSet)`

SetHostPrefixSet sets HostPrefixSet field to given value.

### HasHostPrefixSet

`func (o *V1ExtranetsPostRequestPolicy) HasHostPrefixSet() bool`

HasHostPrefixSet returns a boolean if a field has been set.

### GetManual

`func (o *V1ExtranetsPostRequestPolicy) GetManual() V1ExtranetsGet200ResponsePoliciesInnerManual`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *V1ExtranetsPostRequestPolicy) GetManualOk() (*V1ExtranetsGet200ResponsePoliciesInnerManual, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *V1ExtranetsPostRequestPolicy) SetManual(v V1ExtranetsGet200ResponsePoliciesInnerManual)`

SetManual sets Manual field to given value.

### HasManual

`func (o *V1ExtranetsPostRequestPolicy) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetName

`func (o *V1ExtranetsPostRequestPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ExtranetsPostRequestPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ExtranetsPostRequestPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ExtranetsPostRequestPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSharedPrefixes

`func (o *V1ExtranetsPostRequestPolicy) GetSharedPrefixes() []string`

GetSharedPrefixes returns the SharedPrefixes field if non-nil, zero value otherwise.

### GetSharedPrefixesOk

`func (o *V1ExtranetsPostRequestPolicy) GetSharedPrefixesOk() (*[]string, bool)`

GetSharedPrefixesOk returns a tuple with the SharedPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPrefixes

`func (o *V1ExtranetsPostRequestPolicy) SetSharedPrefixes(v []string)`

SetSharedPrefixes sets SharedPrefixes field to given value.

### HasSharedPrefixes

`func (o *V1ExtranetsPostRequestPolicy) HasSharedPrefixes() bool`

HasSharedPrefixes returns a boolean if a field has been set.

### GetSharedSegment

`func (o *V1ExtranetsPostRequestPolicy) GetSharedSegment() int64`

GetSharedSegment returns the SharedSegment field if non-nil, zero value otherwise.

### GetSharedSegmentOk

`func (o *V1ExtranetsPostRequestPolicy) GetSharedSegmentOk() (*int64, bool)`

GetSharedSegmentOk returns a tuple with the SharedSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSegment

`func (o *V1ExtranetsPostRequestPolicy) SetSharedSegment(v int64)`

SetSharedSegment sets SharedSegment field to given value.

### HasSharedSegment

`func (o *V1ExtranetsPostRequestPolicy) HasSharedSegment() bool`

HasSharedSegment returns a boolean if a field has been set.

### GetSource

`func (o *V1ExtranetsPostRequestPolicy) GetSource() V1ExtranetsPostRequestPolicyBranches`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1ExtranetsPostRequestPolicy) GetSourceOk() (*V1ExtranetsPostRequestPolicyBranches, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1ExtranetsPostRequestPolicy) SetSource(v V1ExtranetsPostRequestPolicyBranches)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1ExtranetsPostRequestPolicy) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTargetSegments

`func (o *V1ExtranetsPostRequestPolicy) GetTargetSegments() []int64`

GetTargetSegments returns the TargetSegments field if non-nil, zero value otherwise.

### GetTargetSegmentsOk

`func (o *V1ExtranetsPostRequestPolicy) GetTargetSegmentsOk() (*[]int64, bool)`

GetTargetSegmentsOk returns a tuple with the TargetSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSegments

`func (o *V1ExtranetsPostRequestPolicy) SetTargetSegments(v []int64)`

SetTargetSegments sets TargetSegments field to given value.

### HasTargetSegments

`func (o *V1ExtranetsPostRequestPolicy) HasTargetSegments() bool`

HasTargetSegments returns a boolean if a field has been set.

### GetType

`func (o *V1ExtranetsPostRequestPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ExtranetsPostRequestPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ExtranetsPostRequestPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1ExtranetsPostRequestPolicy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


