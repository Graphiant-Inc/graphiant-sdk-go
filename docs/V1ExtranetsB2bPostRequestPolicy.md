# V1ExtranetsB2bPostRequestPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatPools** | Pointer to **[]string** |  | [optional] 
**Profiles** | Pointer to [**[]V1ExtranetsB2bPostRequestPolicyProfilesInner**](V1ExtranetsB2bPostRequestPolicyProfilesInner.md) |  | [optional] 
**ServiceLanSegment** | Pointer to **int64** |  | [optional] 
**ServicePrefixes** | Pointer to **[]string** |  | [optional] 
**Sites** | Pointer to [**[]V1ExtranetsB2bConsumerPostRequestSiteInformationInner**](V1ExtranetsB2bConsumerPostRequestSiteInformationInner.md) |  | [optional] 
**Sla** | Pointer to [**V1ExtranetsB2bPostRequestPolicySla**](V1ExtranetsB2bPostRequestPolicySla.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ExtranetsB2bPostRequestPolicy

`func NewV1ExtranetsB2bPostRequestPolicy() *V1ExtranetsB2bPostRequestPolicy`

NewV1ExtranetsB2bPostRequestPolicy instantiates a new V1ExtranetsB2bPostRequestPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsB2bPostRequestPolicyWithDefaults

`func NewV1ExtranetsB2bPostRequestPolicyWithDefaults() *V1ExtranetsB2bPostRequestPolicy`

NewV1ExtranetsB2bPostRequestPolicyWithDefaults instantiates a new V1ExtranetsB2bPostRequestPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatPools

`func (o *V1ExtranetsB2bPostRequestPolicy) GetNatPools() []string`

GetNatPools returns the NatPools field if non-nil, zero value otherwise.

### GetNatPoolsOk

`func (o *V1ExtranetsB2bPostRequestPolicy) GetNatPoolsOk() (*[]string, bool)`

GetNatPoolsOk returns a tuple with the NatPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPools

`func (o *V1ExtranetsB2bPostRequestPolicy) SetNatPools(v []string)`

SetNatPools sets NatPools field to given value.

### HasNatPools

`func (o *V1ExtranetsB2bPostRequestPolicy) HasNatPools() bool`

HasNatPools returns a boolean if a field has been set.

### GetProfiles

`func (o *V1ExtranetsB2bPostRequestPolicy) GetProfiles() []V1ExtranetsB2bPostRequestPolicyProfilesInner`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *V1ExtranetsB2bPostRequestPolicy) GetProfilesOk() (*[]V1ExtranetsB2bPostRequestPolicyProfilesInner, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *V1ExtranetsB2bPostRequestPolicy) SetProfiles(v []V1ExtranetsB2bPostRequestPolicyProfilesInner)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *V1ExtranetsB2bPostRequestPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetServiceLanSegment

`func (o *V1ExtranetsB2bPostRequestPolicy) GetServiceLanSegment() int64`

GetServiceLanSegment returns the ServiceLanSegment field if non-nil, zero value otherwise.

### GetServiceLanSegmentOk

`func (o *V1ExtranetsB2bPostRequestPolicy) GetServiceLanSegmentOk() (*int64, bool)`

GetServiceLanSegmentOk returns a tuple with the ServiceLanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLanSegment

`func (o *V1ExtranetsB2bPostRequestPolicy) SetServiceLanSegment(v int64)`

SetServiceLanSegment sets ServiceLanSegment field to given value.

### HasServiceLanSegment

`func (o *V1ExtranetsB2bPostRequestPolicy) HasServiceLanSegment() bool`

HasServiceLanSegment returns a boolean if a field has been set.

### GetServicePrefixes

`func (o *V1ExtranetsB2bPostRequestPolicy) GetServicePrefixes() []string`

GetServicePrefixes returns the ServicePrefixes field if non-nil, zero value otherwise.

### GetServicePrefixesOk

`func (o *V1ExtranetsB2bPostRequestPolicy) GetServicePrefixesOk() (*[]string, bool)`

GetServicePrefixesOk returns a tuple with the ServicePrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrefixes

`func (o *V1ExtranetsB2bPostRequestPolicy) SetServicePrefixes(v []string)`

SetServicePrefixes sets ServicePrefixes field to given value.

### HasServicePrefixes

`func (o *V1ExtranetsB2bPostRequestPolicy) HasServicePrefixes() bool`

HasServicePrefixes returns a boolean if a field has been set.

### GetSites

`func (o *V1ExtranetsB2bPostRequestPolicy) GetSites() []V1ExtranetsB2bConsumerPostRequestSiteInformationInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *V1ExtranetsB2bPostRequestPolicy) GetSitesOk() (*[]V1ExtranetsB2bConsumerPostRequestSiteInformationInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *V1ExtranetsB2bPostRequestPolicy) SetSites(v []V1ExtranetsB2bConsumerPostRequestSiteInformationInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *V1ExtranetsB2bPostRequestPolicy) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetSla

`func (o *V1ExtranetsB2bPostRequestPolicy) GetSla() V1ExtranetsB2bPostRequestPolicySla`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *V1ExtranetsB2bPostRequestPolicy) GetSlaOk() (*V1ExtranetsB2bPostRequestPolicySla, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *V1ExtranetsB2bPostRequestPolicy) SetSla(v V1ExtranetsB2bPostRequestPolicySla)`

SetSla sets Sla field to given value.

### HasSla

`func (o *V1ExtranetsB2bPostRequestPolicy) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetType

`func (o *V1ExtranetsB2bPostRequestPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ExtranetsB2bPostRequestPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ExtranetsB2bPostRequestPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1ExtranetsB2bPostRequestPolicy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


