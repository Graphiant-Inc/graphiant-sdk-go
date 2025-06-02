# V1ExtranetsGet200ResponsePoliciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auto** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerAuto**](V1ExtranetsGet200ResponsePoliciesInnerAuto.md) |  | [optional] 
**Branches** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranches**](V1ExtranetsGet200ResponsePoliciesInnerBranches.md) |  | [optional] 
**CreatedAt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HostPrefixSet** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesPrefixSet**](V1ExtranetsGet200ResponsePoliciesInnerBranchesPrefixSet.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Manual** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerManual**](V1ExtranetsGet200ResponsePoliciesInnerManual.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SharedPrefixes** | Pointer to **[]string** |  | [optional] 
**SharedSegment** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner.md) |  | [optional] 
**Source** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranches**](V1ExtranetsGet200ResponsePoliciesInnerBranches.md) |  | [optional] 
**TargetSegments** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInner

`func NewV1ExtranetsGet200ResponsePoliciesInner() *V1ExtranetsGet200ResponsePoliciesInner`

NewV1ExtranetsGet200ResponsePoliciesInner instantiates a new V1ExtranetsGet200ResponsePoliciesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInner`

NewV1ExtranetsGet200ResponsePoliciesInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuto

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetAuto() V1ExtranetsGet200ResponsePoliciesInnerAuto`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetAutoOk() (*V1ExtranetsGet200ResponsePoliciesInnerAuto, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetAuto(v V1ExtranetsGet200ResponsePoliciesInnerAuto)`

SetAuto sets Auto field to given value.

### HasAuto

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasAuto() bool`

HasAuto returns a boolean if a field has been set.

### GetBranches

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetBranches() V1ExtranetsGet200ResponsePoliciesInnerBranches`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetBranchesOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranches, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetBranches(v V1ExtranetsGet200ResponsePoliciesInnerBranches)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetCreatedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetCreatedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetCreatedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostPrefixSet

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetHostPrefixSet() V1ExtranetsGet200ResponsePoliciesInnerBranchesPrefixSet`

GetHostPrefixSet returns the HostPrefixSet field if non-nil, zero value otherwise.

### GetHostPrefixSetOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetHostPrefixSetOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesPrefixSet, bool)`

GetHostPrefixSetOk returns a tuple with the HostPrefixSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPrefixSet

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetHostPrefixSet(v V1ExtranetsGet200ResponsePoliciesInnerBranchesPrefixSet)`

SetHostPrefixSet sets HostPrefixSet field to given value.

### HasHostPrefixSet

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasHostPrefixSet() bool`

HasHostPrefixSet returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManual

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetManual() V1ExtranetsGet200ResponsePoliciesInnerManual`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetManualOk() (*V1ExtranetsGet200ResponsePoliciesInnerManual, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetManual(v V1ExtranetsGet200ResponsePoliciesInnerManual)`

SetManual sets Manual field to given value.

### HasManual

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetName

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSharedPrefixes

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetSharedPrefixes() []string`

GetSharedPrefixes returns the SharedPrefixes field if non-nil, zero value otherwise.

### GetSharedPrefixesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetSharedPrefixesOk() (*[]string, bool)`

GetSharedPrefixesOk returns a tuple with the SharedPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPrefixes

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetSharedPrefixes(v []string)`

SetSharedPrefixes sets SharedPrefixes field to given value.

### HasSharedPrefixes

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasSharedPrefixes() bool`

HasSharedPrefixes returns a boolean if a field has been set.

### GetSharedSegment

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetSharedSegment() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner`

GetSharedSegment returns the SharedSegment field if non-nil, zero value otherwise.

### GetSharedSegmentOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetSharedSegmentOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner, bool)`

GetSharedSegmentOk returns a tuple with the SharedSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSegment

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetSharedSegment(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner)`

SetSharedSegment sets SharedSegment field to given value.

### HasSharedSegment

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasSharedSegment() bool`

HasSharedSegment returns a boolean if a field has been set.

### GetSource

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetSource() V1ExtranetsGet200ResponsePoliciesInnerBranches`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetSourceOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranches, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetSource(v V1ExtranetsGet200ResponsePoliciesInnerBranches)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTargetSegments

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetTargetSegments() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner`

GetTargetSegments returns the TargetSegments field if non-nil, zero value otherwise.

### GetTargetSegmentsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetTargetSegmentsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner, bool)`

GetTargetSegmentsOk returns a tuple with the TargetSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSegments

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetTargetSegments(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner)`

SetTargetSegments sets TargetSegments field to given value.

### HasTargetSegments

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasTargetSegments() bool`

HasTargetSegments returns a boolean if a field has been set.

### GetType

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetUpdatedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V1ExtranetsGet200ResponsePoliciesInner) GetUpdatedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInner) SetUpdatedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


