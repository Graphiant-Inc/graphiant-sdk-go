# ManaV2PublicVifConsumerPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerServiceId** | Pointer to **int64** |  | [optional] 
**LanSegments** | Pointer to [**[]ManaV2PublicVifConsumerLanSegmentResponse**](ManaV2PublicVifConsumerLanSegmentResponse.md) |  | [optional] 
**Sites** | Pointer to [**ManaV2SiteInformation**](ManaV2SiteInformation.md) |  | [optional] 

## Methods

### NewManaV2PublicVifConsumerPolicyResponse

`func NewManaV2PublicVifConsumerPolicyResponse() *ManaV2PublicVifConsumerPolicyResponse`

NewManaV2PublicVifConsumerPolicyResponse instantiates a new ManaV2PublicVifConsumerPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PublicVifConsumerPolicyResponseWithDefaults

`func NewManaV2PublicVifConsumerPolicyResponseWithDefaults() *ManaV2PublicVifConsumerPolicyResponse`

NewManaV2PublicVifConsumerPolicyResponseWithDefaults instantiates a new ManaV2PublicVifConsumerPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerServiceId

`func (o *ManaV2PublicVifConsumerPolicyResponse) GetConsumerServiceId() int64`

GetConsumerServiceId returns the ConsumerServiceId field if non-nil, zero value otherwise.

### GetConsumerServiceIdOk

`func (o *ManaV2PublicVifConsumerPolicyResponse) GetConsumerServiceIdOk() (*int64, bool)`

GetConsumerServiceIdOk returns a tuple with the ConsumerServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerServiceId

`func (o *ManaV2PublicVifConsumerPolicyResponse) SetConsumerServiceId(v int64)`

SetConsumerServiceId sets ConsumerServiceId field to given value.

### HasConsumerServiceId

`func (o *ManaV2PublicVifConsumerPolicyResponse) HasConsumerServiceId() bool`

HasConsumerServiceId returns a boolean if a field has been set.

### GetLanSegments

`func (o *ManaV2PublicVifConsumerPolicyResponse) GetLanSegments() []ManaV2PublicVifConsumerLanSegmentResponse`

GetLanSegments returns the LanSegments field if non-nil, zero value otherwise.

### GetLanSegmentsOk

`func (o *ManaV2PublicVifConsumerPolicyResponse) GetLanSegmentsOk() (*[]ManaV2PublicVifConsumerLanSegmentResponse, bool)`

GetLanSegmentsOk returns a tuple with the LanSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegments

`func (o *ManaV2PublicVifConsumerPolicyResponse) SetLanSegments(v []ManaV2PublicVifConsumerLanSegmentResponse)`

SetLanSegments sets LanSegments field to given value.

### HasLanSegments

`func (o *ManaV2PublicVifConsumerPolicyResponse) HasLanSegments() bool`

HasLanSegments returns a boolean if a field has been set.

### GetSites

`func (o *ManaV2PublicVifConsumerPolicyResponse) GetSites() ManaV2SiteInformation`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ManaV2PublicVifConsumerPolicyResponse) GetSitesOk() (*ManaV2SiteInformation, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ManaV2PublicVifConsumerPolicyResponse) SetSites(v ManaV2SiteInformation)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ManaV2PublicVifConsumerPolicyResponse) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


