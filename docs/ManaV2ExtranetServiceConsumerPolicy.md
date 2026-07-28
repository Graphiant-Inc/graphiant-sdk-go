# ManaV2ExtranetServiceConsumerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerLanSegments** | [**map[string]ManaV2ExtranetConsumerLanPrefixes**](ManaV2ExtranetConsumerLanPrefixes.md) |  | 
**GlobalObjectOps** | Pointer to [**map[string]ManaV2GlobalObjectServiceOps**](ManaV2GlobalObjectServiceOps.md) |  | [optional] 
**NatTranslationMode** | Pointer to [**ManaV2ExtranetNatTranslationMode**](ManaV2ExtranetNatTranslationMode.md) |  | [optional] 
**SiteToSiteVpn** | Pointer to [**ManaV2GuestConsumerSiteToSiteVpnConfig**](ManaV2GuestConsumerSiteToSiteVpnConfig.md) |  | [optional] 
**Sites** | Pointer to [**[]ManaV2B2bSiteInformation**](ManaV2B2bSiteInformation.md) |  | [optional] 

## Methods

### NewManaV2ExtranetServiceConsumerPolicy

`func NewManaV2ExtranetServiceConsumerPolicy(consumerLanSegments map[string]ManaV2ExtranetConsumerLanPrefixes, ) *ManaV2ExtranetServiceConsumerPolicy`

NewManaV2ExtranetServiceConsumerPolicy instantiates a new ManaV2ExtranetServiceConsumerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2ExtranetServiceConsumerPolicyWithDefaults

`func NewManaV2ExtranetServiceConsumerPolicyWithDefaults() *ManaV2ExtranetServiceConsumerPolicy`

NewManaV2ExtranetServiceConsumerPolicyWithDefaults instantiates a new ManaV2ExtranetServiceConsumerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerLanSegments

`func (o *ManaV2ExtranetServiceConsumerPolicy) GetConsumerLanSegments() map[string]ManaV2ExtranetConsumerLanPrefixes`

GetConsumerLanSegments returns the ConsumerLanSegments field if non-nil, zero value otherwise.

### GetConsumerLanSegmentsOk

`func (o *ManaV2ExtranetServiceConsumerPolicy) GetConsumerLanSegmentsOk() (*map[string]ManaV2ExtranetConsumerLanPrefixes, bool)`

GetConsumerLanSegmentsOk returns a tuple with the ConsumerLanSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerLanSegments

`func (o *ManaV2ExtranetServiceConsumerPolicy) SetConsumerLanSegments(v map[string]ManaV2ExtranetConsumerLanPrefixes)`

SetConsumerLanSegments sets ConsumerLanSegments field to given value.


### GetGlobalObjectOps

`func (o *ManaV2ExtranetServiceConsumerPolicy) GetGlobalObjectOps() map[string]ManaV2GlobalObjectServiceOps`

GetGlobalObjectOps returns the GlobalObjectOps field if non-nil, zero value otherwise.

### GetGlobalObjectOpsOk

`func (o *ManaV2ExtranetServiceConsumerPolicy) GetGlobalObjectOpsOk() (*map[string]ManaV2GlobalObjectServiceOps, bool)`

GetGlobalObjectOpsOk returns a tuple with the GlobalObjectOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalObjectOps

`func (o *ManaV2ExtranetServiceConsumerPolicy) SetGlobalObjectOps(v map[string]ManaV2GlobalObjectServiceOps)`

SetGlobalObjectOps sets GlobalObjectOps field to given value.

### HasGlobalObjectOps

`func (o *ManaV2ExtranetServiceConsumerPolicy) HasGlobalObjectOps() bool`

HasGlobalObjectOps returns a boolean if a field has been set.

### GetNatTranslationMode

`func (o *ManaV2ExtranetServiceConsumerPolicy) GetNatTranslationMode() ManaV2ExtranetNatTranslationMode`

GetNatTranslationMode returns the NatTranslationMode field if non-nil, zero value otherwise.

### GetNatTranslationModeOk

`func (o *ManaV2ExtranetServiceConsumerPolicy) GetNatTranslationModeOk() (*ManaV2ExtranetNatTranslationMode, bool)`

GetNatTranslationModeOk returns a tuple with the NatTranslationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTranslationMode

`func (o *ManaV2ExtranetServiceConsumerPolicy) SetNatTranslationMode(v ManaV2ExtranetNatTranslationMode)`

SetNatTranslationMode sets NatTranslationMode field to given value.

### HasNatTranslationMode

`func (o *ManaV2ExtranetServiceConsumerPolicy) HasNatTranslationMode() bool`

HasNatTranslationMode returns a boolean if a field has been set.

### GetSiteToSiteVpn

`func (o *ManaV2ExtranetServiceConsumerPolicy) GetSiteToSiteVpn() ManaV2GuestConsumerSiteToSiteVpnConfig`

GetSiteToSiteVpn returns the SiteToSiteVpn field if non-nil, zero value otherwise.

### GetSiteToSiteVpnOk

`func (o *ManaV2ExtranetServiceConsumerPolicy) GetSiteToSiteVpnOk() (*ManaV2GuestConsumerSiteToSiteVpnConfig, bool)`

GetSiteToSiteVpnOk returns a tuple with the SiteToSiteVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteToSiteVpn

`func (o *ManaV2ExtranetServiceConsumerPolicy) SetSiteToSiteVpn(v ManaV2GuestConsumerSiteToSiteVpnConfig)`

SetSiteToSiteVpn sets SiteToSiteVpn field to given value.

### HasSiteToSiteVpn

`func (o *ManaV2ExtranetServiceConsumerPolicy) HasSiteToSiteVpn() bool`

HasSiteToSiteVpn returns a boolean if a field has been set.

### GetSites

`func (o *ManaV2ExtranetServiceConsumerPolicy) GetSites() []ManaV2B2bSiteInformation`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ManaV2ExtranetServiceConsumerPolicy) GetSitesOk() (*[]ManaV2B2bSiteInformation, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ManaV2ExtranetServiceConsumerPolicy) SetSites(v []ManaV2B2bSiteInformation)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ManaV2ExtranetServiceConsumerPolicy) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


