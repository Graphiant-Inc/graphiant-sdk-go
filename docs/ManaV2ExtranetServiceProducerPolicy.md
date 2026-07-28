# ManaV2ExtranetServiceProducerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**GlobalObjectDeviceSummaries** | Pointer to [**map[string]ManaV2GlobalObjectServiceSummaries**](ManaV2GlobalObjectServiceSummaries.md) |  | [optional] 
**GlobalObjectOps** | Pointer to [**map[string]ManaV2GlobalObjectServiceOps**](ManaV2GlobalObjectServiceOps.md) |  | [optional] 
**NatTranslationMode** | Pointer to [**ManaV2ExtranetNatTranslationMode**](ManaV2ExtranetNatTranslationMode.md) |  | [optional] 
**PrefixTags** | Pointer to [**[]ManaV2B2bExtranetPrefixTag**](ManaV2B2bExtranetPrefixTag.md) |  | [optional] 
**ServiceLanSegment** | Pointer to **int64** | LAN segment ID for the service | [optional] 
**Sites** | Pointer to [**[]ManaV2B2bSiteInformation**](ManaV2B2bSiteInformation.md) |  | [optional] 

## Methods

### NewManaV2ExtranetServiceProducerPolicy

`func NewManaV2ExtranetServiceProducerPolicy() *ManaV2ExtranetServiceProducerPolicy`

NewManaV2ExtranetServiceProducerPolicy instantiates a new ManaV2ExtranetServiceProducerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2ExtranetServiceProducerPolicyWithDefaults

`func NewManaV2ExtranetServiceProducerPolicyWithDefaults() *ManaV2ExtranetServiceProducerPolicy`

NewManaV2ExtranetServiceProducerPolicyWithDefaults instantiates a new ManaV2ExtranetServiceProducerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ManaV2ExtranetServiceProducerPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2ExtranetServiceProducerPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2ExtranetServiceProducerPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2ExtranetServiceProducerPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGlobalObjectDeviceSummaries

`func (o *ManaV2ExtranetServiceProducerPolicy) GetGlobalObjectDeviceSummaries() map[string]ManaV2GlobalObjectServiceSummaries`

GetGlobalObjectDeviceSummaries returns the GlobalObjectDeviceSummaries field if non-nil, zero value otherwise.

### GetGlobalObjectDeviceSummariesOk

`func (o *ManaV2ExtranetServiceProducerPolicy) GetGlobalObjectDeviceSummariesOk() (*map[string]ManaV2GlobalObjectServiceSummaries, bool)`

GetGlobalObjectDeviceSummariesOk returns a tuple with the GlobalObjectDeviceSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalObjectDeviceSummaries

`func (o *ManaV2ExtranetServiceProducerPolicy) SetGlobalObjectDeviceSummaries(v map[string]ManaV2GlobalObjectServiceSummaries)`

SetGlobalObjectDeviceSummaries sets GlobalObjectDeviceSummaries field to given value.

### HasGlobalObjectDeviceSummaries

`func (o *ManaV2ExtranetServiceProducerPolicy) HasGlobalObjectDeviceSummaries() bool`

HasGlobalObjectDeviceSummaries returns a boolean if a field has been set.

### GetGlobalObjectOps

`func (o *ManaV2ExtranetServiceProducerPolicy) GetGlobalObjectOps() map[string]ManaV2GlobalObjectServiceOps`

GetGlobalObjectOps returns the GlobalObjectOps field if non-nil, zero value otherwise.

### GetGlobalObjectOpsOk

`func (o *ManaV2ExtranetServiceProducerPolicy) GetGlobalObjectOpsOk() (*map[string]ManaV2GlobalObjectServiceOps, bool)`

GetGlobalObjectOpsOk returns a tuple with the GlobalObjectOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalObjectOps

`func (o *ManaV2ExtranetServiceProducerPolicy) SetGlobalObjectOps(v map[string]ManaV2GlobalObjectServiceOps)`

SetGlobalObjectOps sets GlobalObjectOps field to given value.

### HasGlobalObjectOps

`func (o *ManaV2ExtranetServiceProducerPolicy) HasGlobalObjectOps() bool`

HasGlobalObjectOps returns a boolean if a field has been set.

### GetNatTranslationMode

`func (o *ManaV2ExtranetServiceProducerPolicy) GetNatTranslationMode() ManaV2ExtranetNatTranslationMode`

GetNatTranslationMode returns the NatTranslationMode field if non-nil, zero value otherwise.

### GetNatTranslationModeOk

`func (o *ManaV2ExtranetServiceProducerPolicy) GetNatTranslationModeOk() (*ManaV2ExtranetNatTranslationMode, bool)`

GetNatTranslationModeOk returns a tuple with the NatTranslationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTranslationMode

`func (o *ManaV2ExtranetServiceProducerPolicy) SetNatTranslationMode(v ManaV2ExtranetNatTranslationMode)`

SetNatTranslationMode sets NatTranslationMode field to given value.

### HasNatTranslationMode

`func (o *ManaV2ExtranetServiceProducerPolicy) HasNatTranslationMode() bool`

HasNatTranslationMode returns a boolean if a field has been set.

### GetPrefixTags

`func (o *ManaV2ExtranetServiceProducerPolicy) GetPrefixTags() []ManaV2B2bExtranetPrefixTag`

GetPrefixTags returns the PrefixTags field if non-nil, zero value otherwise.

### GetPrefixTagsOk

`func (o *ManaV2ExtranetServiceProducerPolicy) GetPrefixTagsOk() (*[]ManaV2B2bExtranetPrefixTag, bool)`

GetPrefixTagsOk returns a tuple with the PrefixTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixTags

`func (o *ManaV2ExtranetServiceProducerPolicy) SetPrefixTags(v []ManaV2B2bExtranetPrefixTag)`

SetPrefixTags sets PrefixTags field to given value.

### HasPrefixTags

`func (o *ManaV2ExtranetServiceProducerPolicy) HasPrefixTags() bool`

HasPrefixTags returns a boolean if a field has been set.

### GetServiceLanSegment

`func (o *ManaV2ExtranetServiceProducerPolicy) GetServiceLanSegment() int64`

GetServiceLanSegment returns the ServiceLanSegment field if non-nil, zero value otherwise.

### GetServiceLanSegmentOk

`func (o *ManaV2ExtranetServiceProducerPolicy) GetServiceLanSegmentOk() (*int64, bool)`

GetServiceLanSegmentOk returns a tuple with the ServiceLanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLanSegment

`func (o *ManaV2ExtranetServiceProducerPolicy) SetServiceLanSegment(v int64)`

SetServiceLanSegment sets ServiceLanSegment field to given value.

### HasServiceLanSegment

`func (o *ManaV2ExtranetServiceProducerPolicy) HasServiceLanSegment() bool`

HasServiceLanSegment returns a boolean if a field has been set.

### GetSites

`func (o *ManaV2ExtranetServiceProducerPolicy) GetSites() []ManaV2B2bSiteInformation`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ManaV2ExtranetServiceProducerPolicy) GetSitesOk() (*[]ManaV2B2bSiteInformation, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ManaV2ExtranetServiceProducerPolicy) SetSites(v []ManaV2B2bSiteInformation)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ManaV2ExtranetServiceProducerPolicy) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


