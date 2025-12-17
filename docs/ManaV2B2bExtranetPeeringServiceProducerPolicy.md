# ManaV2B2bExtranetPeeringServiceProducerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description for the service | [optional] 
**GlobalObjectOps** | Pointer to [**map[string]ManaV2GlobalObjectServiceOps**](ManaV2GlobalObjectServiceOps.md) |  | [optional] 
**PrefixTags** | [**[]ManaV2B2bExtranetPrefixTag**](ManaV2B2bExtranetPrefixTag.md) |  | 
**ServiceLanSegment** | **int64** | LAN segment ID for the service (required) | 
**Site** | [**[]ManaV2B2bSiteInformation**](ManaV2B2bSiteInformation.md) |  | 
**Type** | **string** | Type of the service whether it is application or peering (required) | 

## Methods

### NewManaV2B2bExtranetPeeringServiceProducerPolicy

`func NewManaV2B2bExtranetPeeringServiceProducerPolicy(prefixTags []ManaV2B2bExtranetPrefixTag, serviceLanSegment int64, site []ManaV2B2bSiteInformation, type_ string, ) *ManaV2B2bExtranetPeeringServiceProducerPolicy`

NewManaV2B2bExtranetPeeringServiceProducerPolicy instantiates a new ManaV2B2bExtranetPeeringServiceProducerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2B2bExtranetPeeringServiceProducerPolicyWithDefaults

`func NewManaV2B2bExtranetPeeringServiceProducerPolicyWithDefaults() *ManaV2B2bExtranetPeeringServiceProducerPolicy`

NewManaV2B2bExtranetPeeringServiceProducerPolicyWithDefaults instantiates a new ManaV2B2bExtranetPeeringServiceProducerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGlobalObjectOps

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetGlobalObjectOps() map[string]ManaV2GlobalObjectServiceOps`

GetGlobalObjectOps returns the GlobalObjectOps field if non-nil, zero value otherwise.

### GetGlobalObjectOpsOk

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetGlobalObjectOpsOk() (*map[string]ManaV2GlobalObjectServiceOps, bool)`

GetGlobalObjectOpsOk returns a tuple with the GlobalObjectOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalObjectOps

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) SetGlobalObjectOps(v map[string]ManaV2GlobalObjectServiceOps)`

SetGlobalObjectOps sets GlobalObjectOps field to given value.

### HasGlobalObjectOps

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) HasGlobalObjectOps() bool`

HasGlobalObjectOps returns a boolean if a field has been set.

### GetPrefixTags

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetPrefixTags() []ManaV2B2bExtranetPrefixTag`

GetPrefixTags returns the PrefixTags field if non-nil, zero value otherwise.

### GetPrefixTagsOk

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetPrefixTagsOk() (*[]ManaV2B2bExtranetPrefixTag, bool)`

GetPrefixTagsOk returns a tuple with the PrefixTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixTags

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) SetPrefixTags(v []ManaV2B2bExtranetPrefixTag)`

SetPrefixTags sets PrefixTags field to given value.


### GetServiceLanSegment

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetServiceLanSegment() int64`

GetServiceLanSegment returns the ServiceLanSegment field if non-nil, zero value otherwise.

### GetServiceLanSegmentOk

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetServiceLanSegmentOk() (*int64, bool)`

GetServiceLanSegmentOk returns a tuple with the ServiceLanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLanSegment

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) SetServiceLanSegment(v int64)`

SetServiceLanSegment sets ServiceLanSegment field to given value.


### GetSite

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetSite() []ManaV2B2bSiteInformation`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetSiteOk() (*[]ManaV2B2bSiteInformation, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) SetSite(v []ManaV2B2bSiteInformation)`

SetSite sets Site field to given value.


### GetType

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManaV2B2bExtranetPeeringServiceProducerPolicy) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


