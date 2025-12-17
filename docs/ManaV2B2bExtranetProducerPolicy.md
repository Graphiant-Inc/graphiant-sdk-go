# ManaV2B2bExtranetProducerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description for the service | [optional] 
**GlobalObjectDeviceSummaries** | Pointer to [**map[string]ManaV2GlobalObjectServiceSummaries**](ManaV2GlobalObjectServiceSummaries.md) |  | [optional] 
**GlobalObjectSummaries** | Pointer to [**map[string]ManaV2GlobalObjectServiceSummaries**](ManaV2GlobalObjectServiceSummaries.md) |  | [optional] 
**NatPools** | **[]string** |  | 
**PrefixTags** | Pointer to [**[]ManaV2B2bExtranetPrefixTag**](ManaV2B2bExtranetPrefixTag.md) |  | [optional] 
**Profiles** | Pointer to [**[]ManaV2ApplicationProfile**](ManaV2ApplicationProfile.md) |  | [optional] 
**ServiceLanSegment** | **int64** | LAN segment for the service (required) | 
**ServicePrefixes** | **[]string** |  | 
**Sites** | [**[]ManaV2B2bSiteInformation**](ManaV2B2bSiteInformation.md) |  | 
**Sla** | Pointer to [**ManaV2SlaInformation**](ManaV2SlaInformation.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | **string** | Type of the service whether it is application or peering (required) | 
**UnmatchedCustomers** | Pointer to **int32** |  | [optional] 

## Methods

### NewManaV2B2bExtranetProducerPolicy

`func NewManaV2B2bExtranetProducerPolicy(natPools []string, serviceLanSegment int64, servicePrefixes []string, sites []ManaV2B2bSiteInformation, type_ string, ) *ManaV2B2bExtranetProducerPolicy`

NewManaV2B2bExtranetProducerPolicy instantiates a new ManaV2B2bExtranetProducerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2B2bExtranetProducerPolicyWithDefaults

`func NewManaV2B2bExtranetProducerPolicyWithDefaults() *ManaV2B2bExtranetProducerPolicy`

NewManaV2B2bExtranetProducerPolicyWithDefaults instantiates a new ManaV2B2bExtranetProducerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ManaV2B2bExtranetProducerPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2B2bExtranetProducerPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2B2bExtranetProducerPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGlobalObjectDeviceSummaries

`func (o *ManaV2B2bExtranetProducerPolicy) GetGlobalObjectDeviceSummaries() map[string]ManaV2GlobalObjectServiceSummaries`

GetGlobalObjectDeviceSummaries returns the GlobalObjectDeviceSummaries field if non-nil, zero value otherwise.

### GetGlobalObjectDeviceSummariesOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetGlobalObjectDeviceSummariesOk() (*map[string]ManaV2GlobalObjectServiceSummaries, bool)`

GetGlobalObjectDeviceSummariesOk returns a tuple with the GlobalObjectDeviceSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalObjectDeviceSummaries

`func (o *ManaV2B2bExtranetProducerPolicy) SetGlobalObjectDeviceSummaries(v map[string]ManaV2GlobalObjectServiceSummaries)`

SetGlobalObjectDeviceSummaries sets GlobalObjectDeviceSummaries field to given value.

### HasGlobalObjectDeviceSummaries

`func (o *ManaV2B2bExtranetProducerPolicy) HasGlobalObjectDeviceSummaries() bool`

HasGlobalObjectDeviceSummaries returns a boolean if a field has been set.

### GetGlobalObjectSummaries

`func (o *ManaV2B2bExtranetProducerPolicy) GetGlobalObjectSummaries() map[string]ManaV2GlobalObjectServiceSummaries`

GetGlobalObjectSummaries returns the GlobalObjectSummaries field if non-nil, zero value otherwise.

### GetGlobalObjectSummariesOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetGlobalObjectSummariesOk() (*map[string]ManaV2GlobalObjectServiceSummaries, bool)`

GetGlobalObjectSummariesOk returns a tuple with the GlobalObjectSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalObjectSummaries

`func (o *ManaV2B2bExtranetProducerPolicy) SetGlobalObjectSummaries(v map[string]ManaV2GlobalObjectServiceSummaries)`

SetGlobalObjectSummaries sets GlobalObjectSummaries field to given value.

### HasGlobalObjectSummaries

`func (o *ManaV2B2bExtranetProducerPolicy) HasGlobalObjectSummaries() bool`

HasGlobalObjectSummaries returns a boolean if a field has been set.

### GetNatPools

`func (o *ManaV2B2bExtranetProducerPolicy) GetNatPools() []string`

GetNatPools returns the NatPools field if non-nil, zero value otherwise.

### GetNatPoolsOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetNatPoolsOk() (*[]string, bool)`

GetNatPoolsOk returns a tuple with the NatPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPools

`func (o *ManaV2B2bExtranetProducerPolicy) SetNatPools(v []string)`

SetNatPools sets NatPools field to given value.


### GetPrefixTags

`func (o *ManaV2B2bExtranetProducerPolicy) GetPrefixTags() []ManaV2B2bExtranetPrefixTag`

GetPrefixTags returns the PrefixTags field if non-nil, zero value otherwise.

### GetPrefixTagsOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetPrefixTagsOk() (*[]ManaV2B2bExtranetPrefixTag, bool)`

GetPrefixTagsOk returns a tuple with the PrefixTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixTags

`func (o *ManaV2B2bExtranetProducerPolicy) SetPrefixTags(v []ManaV2B2bExtranetPrefixTag)`

SetPrefixTags sets PrefixTags field to given value.

### HasPrefixTags

`func (o *ManaV2B2bExtranetProducerPolicy) HasPrefixTags() bool`

HasPrefixTags returns a boolean if a field has been set.

### GetProfiles

`func (o *ManaV2B2bExtranetProducerPolicy) GetProfiles() []ManaV2ApplicationProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetProfilesOk() (*[]ManaV2ApplicationProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *ManaV2B2bExtranetProducerPolicy) SetProfiles(v []ManaV2ApplicationProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *ManaV2B2bExtranetProducerPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetServiceLanSegment

`func (o *ManaV2B2bExtranetProducerPolicy) GetServiceLanSegment() int64`

GetServiceLanSegment returns the ServiceLanSegment field if non-nil, zero value otherwise.

### GetServiceLanSegmentOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetServiceLanSegmentOk() (*int64, bool)`

GetServiceLanSegmentOk returns a tuple with the ServiceLanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLanSegment

`func (o *ManaV2B2bExtranetProducerPolicy) SetServiceLanSegment(v int64)`

SetServiceLanSegment sets ServiceLanSegment field to given value.


### GetServicePrefixes

`func (o *ManaV2B2bExtranetProducerPolicy) GetServicePrefixes() []string`

GetServicePrefixes returns the ServicePrefixes field if non-nil, zero value otherwise.

### GetServicePrefixesOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetServicePrefixesOk() (*[]string, bool)`

GetServicePrefixesOk returns a tuple with the ServicePrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrefixes

`func (o *ManaV2B2bExtranetProducerPolicy) SetServicePrefixes(v []string)`

SetServicePrefixes sets ServicePrefixes field to given value.


### GetSites

`func (o *ManaV2B2bExtranetProducerPolicy) GetSites() []ManaV2B2bSiteInformation`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetSitesOk() (*[]ManaV2B2bSiteInformation, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ManaV2B2bExtranetProducerPolicy) SetSites(v []ManaV2B2bSiteInformation)`

SetSites sets Sites field to given value.


### GetSla

`func (o *ManaV2B2bExtranetProducerPolicy) GetSla() ManaV2SlaInformation`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetSlaOk() (*ManaV2SlaInformation, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *ManaV2B2bExtranetProducerPolicy) SetSla(v ManaV2SlaInformation)`

SetSla sets Sla field to given value.

### HasSla

`func (o *ManaV2B2bExtranetProducerPolicy) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetStatus

`func (o *ManaV2B2bExtranetProducerPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManaV2B2bExtranetProducerPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManaV2B2bExtranetProducerPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ManaV2B2bExtranetProducerPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManaV2B2bExtranetProducerPolicy) SetType(v string)`

SetType sets Type field to given value.


### GetUnmatchedCustomers

`func (o *ManaV2B2bExtranetProducerPolicy) GetUnmatchedCustomers() int32`

GetUnmatchedCustomers returns the UnmatchedCustomers field if non-nil, zero value otherwise.

### GetUnmatchedCustomersOk

`func (o *ManaV2B2bExtranetProducerPolicy) GetUnmatchedCustomersOk() (*int32, bool)`

GetUnmatchedCustomersOk returns a tuple with the UnmatchedCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedCustomers

`func (o *ManaV2B2bExtranetProducerPolicy) SetUnmatchedCustomers(v int32)`

SetUnmatchedCustomers sets UnmatchedCustomers field to given value.

### HasUnmatchedCustomers

`func (o *ManaV2B2bExtranetProducerPolicy) HasUnmatchedCustomers() bool`

HasUnmatchedCustomers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


