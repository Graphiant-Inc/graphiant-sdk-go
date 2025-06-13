# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Devices** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSiteDevicesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSiteDevicesInner.md) |  | [optional] 
**EdgeCount** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Location** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**PolicyReferenceCount** | Pointer to **int32** |  | [optional] 
**PolicyTag** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSitePolicyTag**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSitePolicyTag.md) |  | [optional] 
**SegmentCount** | Pointer to **int32** |  | [optional] 
**SiteListReferenceCount** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSiteWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSiteWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSiteWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetCreatedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetCreatedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetCreatedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDevices

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetDevices() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSiteDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetDevicesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSiteDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetDevices(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSiteDevicesInner)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetEdgeCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetEdgeCount() int32`

GetEdgeCount returns the EdgeCount field if non-nil, zero value otherwise.

### GetEdgeCountOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetEdgeCountOk() (*int32, bool)`

GetEdgeCountOk returns a tuple with the EdgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetEdgeCount(v int32)`

SetEdgeCount sets EdgeCount field to given value.

### HasEdgeCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasEdgeCount() bool`

HasEdgeCount returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetLocation() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetLocationOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetLocation(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPolicyReferenceCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetPolicyReferenceCount() int32`

GetPolicyReferenceCount returns the PolicyReferenceCount field if non-nil, zero value otherwise.

### GetPolicyReferenceCountOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetPolicyReferenceCountOk() (*int32, bool)`

GetPolicyReferenceCountOk returns a tuple with the PolicyReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReferenceCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetPolicyReferenceCount(v int32)`

SetPolicyReferenceCount sets PolicyReferenceCount field to given value.

### HasPolicyReferenceCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasPolicyReferenceCount() bool`

HasPolicyReferenceCount returns a boolean if a field has been set.

### GetPolicyTag

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetPolicyTag() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSitePolicyTag`

GetPolicyTag returns the PolicyTag field if non-nil, zero value otherwise.

### GetPolicyTagOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetPolicyTagOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSitePolicyTag, bool)`

GetPolicyTagOk returns a tuple with the PolicyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTag

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetPolicyTag(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSitePolicyTag)`

SetPolicyTag sets PolicyTag field to given value.

### HasPolicyTag

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasPolicyTag() bool`

HasPolicyTag returns a boolean if a field has been set.

### GetSegmentCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetSegmentCount() int32`

GetSegmentCount returns the SegmentCount field if non-nil, zero value otherwise.

### GetSegmentCountOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetSegmentCountOk() (*int32, bool)`

GetSegmentCountOk returns a tuple with the SegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetSegmentCount(v int32)`

SetSegmentCount sets SegmentCount field to given value.

### HasSegmentCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasSegmentCount() bool`

HasSegmentCount returns a boolean if a field has been set.

### GetSiteListReferenceCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetSiteListReferenceCount() int32`

GetSiteListReferenceCount returns the SiteListReferenceCount field if non-nil, zero value otherwise.

### GetSiteListReferenceCountOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetSiteListReferenceCountOk() (*int32, bool)`

GetSiteListReferenceCountOk returns a tuple with the SiteListReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteListReferenceCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetSiteListReferenceCount(v int32)`

SetSiteListReferenceCount sets SiteListReferenceCount field to given value.

### HasSiteListReferenceCount

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasSiteListReferenceCount() bool`

HasSiteListReferenceCount returns a boolean if a field has been set.

### GetTags

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetUpdatedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) GetUpdatedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) SetUpdatedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


