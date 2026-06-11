# ManaV2PublicVifConsumerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LanSegments** | [**[]ManaV2PublicVifConsumerLanSegment**](ManaV2PublicVifConsumerLanSegment.md) |  | 
**Sites** | [**ManaV2SiteInformation**](ManaV2SiteInformation.md) |  | 

## Methods

### NewManaV2PublicVifConsumerPolicy

`func NewManaV2PublicVifConsumerPolicy(lanSegments []ManaV2PublicVifConsumerLanSegment, sites ManaV2SiteInformation, ) *ManaV2PublicVifConsumerPolicy`

NewManaV2PublicVifConsumerPolicy instantiates a new ManaV2PublicVifConsumerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PublicVifConsumerPolicyWithDefaults

`func NewManaV2PublicVifConsumerPolicyWithDefaults() *ManaV2PublicVifConsumerPolicy`

NewManaV2PublicVifConsumerPolicyWithDefaults instantiates a new ManaV2PublicVifConsumerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanSegments

`func (o *ManaV2PublicVifConsumerPolicy) GetLanSegments() []ManaV2PublicVifConsumerLanSegment`

GetLanSegments returns the LanSegments field if non-nil, zero value otherwise.

### GetLanSegmentsOk

`func (o *ManaV2PublicVifConsumerPolicy) GetLanSegmentsOk() (*[]ManaV2PublicVifConsumerLanSegment, bool)`

GetLanSegmentsOk returns a tuple with the LanSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegments

`func (o *ManaV2PublicVifConsumerPolicy) SetLanSegments(v []ManaV2PublicVifConsumerLanSegment)`

SetLanSegments sets LanSegments field to given value.


### GetSites

`func (o *ManaV2PublicVifConsumerPolicy) GetSites() ManaV2SiteInformation`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ManaV2PublicVifConsumerPolicy) GetSitesOk() (*ManaV2SiteInformation, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ManaV2PublicVifConsumerPolicy) SetSites(v ManaV2SiteInformation)`

SetSites sets Sites field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


