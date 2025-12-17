# ManaV2B2bSiteInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BwAllocationSiteLists** | Pointer to **int32** | Total Bandwidth allocation for the service on these site lists | [optional] 
**BwAllocationSites** | Pointer to **int32** | Total Bandwidth allocation for the service on these sites | [optional] 
**PolicerSiteLists** | Pointer to [**ManaV2Policer**](ManaV2Policer.md) |  | [optional] 
**PolicerSites** | Pointer to [**ManaV2Policer**](ManaV2Policer.md) |  | [optional] 
**SiteLists** | Pointer to **[]int64** |  | [optional] 
**Sites** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewManaV2B2bSiteInformation

`func NewManaV2B2bSiteInformation() *ManaV2B2bSiteInformation`

NewManaV2B2bSiteInformation instantiates a new ManaV2B2bSiteInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2B2bSiteInformationWithDefaults

`func NewManaV2B2bSiteInformationWithDefaults() *ManaV2B2bSiteInformation`

NewManaV2B2bSiteInformationWithDefaults instantiates a new ManaV2B2bSiteInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBwAllocationSiteLists

`func (o *ManaV2B2bSiteInformation) GetBwAllocationSiteLists() int32`

GetBwAllocationSiteLists returns the BwAllocationSiteLists field if non-nil, zero value otherwise.

### GetBwAllocationSiteListsOk

`func (o *ManaV2B2bSiteInformation) GetBwAllocationSiteListsOk() (*int32, bool)`

GetBwAllocationSiteListsOk returns a tuple with the BwAllocationSiteLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwAllocationSiteLists

`func (o *ManaV2B2bSiteInformation) SetBwAllocationSiteLists(v int32)`

SetBwAllocationSiteLists sets BwAllocationSiteLists field to given value.

### HasBwAllocationSiteLists

`func (o *ManaV2B2bSiteInformation) HasBwAllocationSiteLists() bool`

HasBwAllocationSiteLists returns a boolean if a field has been set.

### GetBwAllocationSites

`func (o *ManaV2B2bSiteInformation) GetBwAllocationSites() int32`

GetBwAllocationSites returns the BwAllocationSites field if non-nil, zero value otherwise.

### GetBwAllocationSitesOk

`func (o *ManaV2B2bSiteInformation) GetBwAllocationSitesOk() (*int32, bool)`

GetBwAllocationSitesOk returns a tuple with the BwAllocationSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwAllocationSites

`func (o *ManaV2B2bSiteInformation) SetBwAllocationSites(v int32)`

SetBwAllocationSites sets BwAllocationSites field to given value.

### HasBwAllocationSites

`func (o *ManaV2B2bSiteInformation) HasBwAllocationSites() bool`

HasBwAllocationSites returns a boolean if a field has been set.

### GetPolicerSiteLists

`func (o *ManaV2B2bSiteInformation) GetPolicerSiteLists() ManaV2Policer`

GetPolicerSiteLists returns the PolicerSiteLists field if non-nil, zero value otherwise.

### GetPolicerSiteListsOk

`func (o *ManaV2B2bSiteInformation) GetPolicerSiteListsOk() (*ManaV2Policer, bool)`

GetPolicerSiteListsOk returns a tuple with the PolicerSiteLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicerSiteLists

`func (o *ManaV2B2bSiteInformation) SetPolicerSiteLists(v ManaV2Policer)`

SetPolicerSiteLists sets PolicerSiteLists field to given value.

### HasPolicerSiteLists

`func (o *ManaV2B2bSiteInformation) HasPolicerSiteLists() bool`

HasPolicerSiteLists returns a boolean if a field has been set.

### GetPolicerSites

`func (o *ManaV2B2bSiteInformation) GetPolicerSites() ManaV2Policer`

GetPolicerSites returns the PolicerSites field if non-nil, zero value otherwise.

### GetPolicerSitesOk

`func (o *ManaV2B2bSiteInformation) GetPolicerSitesOk() (*ManaV2Policer, bool)`

GetPolicerSitesOk returns a tuple with the PolicerSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicerSites

`func (o *ManaV2B2bSiteInformation) SetPolicerSites(v ManaV2Policer)`

SetPolicerSites sets PolicerSites field to given value.

### HasPolicerSites

`func (o *ManaV2B2bSiteInformation) HasPolicerSites() bool`

HasPolicerSites returns a boolean if a field has been set.

### GetSiteLists

`func (o *ManaV2B2bSiteInformation) GetSiteLists() []int64`

GetSiteLists returns the SiteLists field if non-nil, zero value otherwise.

### GetSiteListsOk

`func (o *ManaV2B2bSiteInformation) GetSiteListsOk() (*[]int64, bool)`

GetSiteListsOk returns a tuple with the SiteLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteLists

`func (o *ManaV2B2bSiteInformation) SetSiteLists(v []int64)`

SetSiteLists sets SiteLists field to given value.

### HasSiteLists

`func (o *ManaV2B2bSiteInformation) HasSiteLists() bool`

HasSiteLists returns a boolean if a field has been set.

### GetSites

`func (o *ManaV2B2bSiteInformation) GetSites() []int64`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ManaV2B2bSiteInformation) GetSitesOk() (*[]int64, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ManaV2B2bSiteInformation) SetSites(v []int64)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ManaV2B2bSiteInformation) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


