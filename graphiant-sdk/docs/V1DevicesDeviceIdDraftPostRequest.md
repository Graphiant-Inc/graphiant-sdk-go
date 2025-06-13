# V1DevicesDeviceIdDraftPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseVersion** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Draft** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInner.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdDraftPostRequest

`func NewV1DevicesDeviceIdDraftPostRequest() *V1DevicesDeviceIdDraftPostRequest`

NewV1DevicesDeviceIdDraftPostRequest instantiates a new V1DevicesDeviceIdDraftPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdDraftPostRequestWithDefaults

`func NewV1DevicesDeviceIdDraftPostRequestWithDefaults() *V1DevicesDeviceIdDraftPostRequest`

NewV1DevicesDeviceIdDraftPostRequestWithDefaults instantiates a new V1DevicesDeviceIdDraftPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseVersion

`func (o *V1DevicesDeviceIdDraftPostRequest) GetBaseVersion() int32`

GetBaseVersion returns the BaseVersion field if non-nil, zero value otherwise.

### GetBaseVersionOk

`func (o *V1DevicesDeviceIdDraftPostRequest) GetBaseVersionOk() (*int32, bool)`

GetBaseVersionOk returns a tuple with the BaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVersion

`func (o *V1DevicesDeviceIdDraftPostRequest) SetBaseVersion(v int32)`

SetBaseVersion sets BaseVersion field to given value.

### HasBaseVersion

`func (o *V1DevicesDeviceIdDraftPostRequest) HasBaseVersion() bool`

HasBaseVersion returns a boolean if a field has been set.

### GetDescription

`func (o *V1DevicesDeviceIdDraftPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1DevicesDeviceIdDraftPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1DevicesDeviceIdDraftPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1DevicesDeviceIdDraftPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDraft

`func (o *V1DevicesDeviceIdDraftPostRequest) GetDraft() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInner`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *V1DevicesDeviceIdDraftPostRequest) GetDraftOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInner, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *V1DevicesDeviceIdDraftPostRequest) SetDraft(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInner)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *V1DevicesDeviceIdDraftPostRequest) HasDraft() bool`

HasDraft returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


