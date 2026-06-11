# ManaV2PublicVifConsumerLanSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPrefixes** | **[]string** |  | 
**ServiceLanSegment** | **int64** | LAN segment ID for the service (required) | 

## Methods

### NewManaV2PublicVifConsumerLanSegment

`func NewManaV2PublicVifConsumerLanSegment(allowedPrefixes []string, serviceLanSegment int64, ) *ManaV2PublicVifConsumerLanSegment`

NewManaV2PublicVifConsumerLanSegment instantiates a new ManaV2PublicVifConsumerLanSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PublicVifConsumerLanSegmentWithDefaults

`func NewManaV2PublicVifConsumerLanSegmentWithDefaults() *ManaV2PublicVifConsumerLanSegment`

NewManaV2PublicVifConsumerLanSegmentWithDefaults instantiates a new ManaV2PublicVifConsumerLanSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPrefixes

`func (o *ManaV2PublicVifConsumerLanSegment) GetAllowedPrefixes() []string`

GetAllowedPrefixes returns the AllowedPrefixes field if non-nil, zero value otherwise.

### GetAllowedPrefixesOk

`func (o *ManaV2PublicVifConsumerLanSegment) GetAllowedPrefixesOk() (*[]string, bool)`

GetAllowedPrefixesOk returns a tuple with the AllowedPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPrefixes

`func (o *ManaV2PublicVifConsumerLanSegment) SetAllowedPrefixes(v []string)`

SetAllowedPrefixes sets AllowedPrefixes field to given value.


### GetServiceLanSegment

`func (o *ManaV2PublicVifConsumerLanSegment) GetServiceLanSegment() int64`

GetServiceLanSegment returns the ServiceLanSegment field if non-nil, zero value otherwise.

### GetServiceLanSegmentOk

`func (o *ManaV2PublicVifConsumerLanSegment) GetServiceLanSegmentOk() (*int64, bool)`

GetServiceLanSegmentOk returns a tuple with the ServiceLanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLanSegment

`func (o *ManaV2PublicVifConsumerLanSegment) SetServiceLanSegment(v int64)`

SetServiceLanSegment sets ServiceLanSegment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


