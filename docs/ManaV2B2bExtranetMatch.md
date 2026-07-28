# ManaV2B2bExtranetMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerPrefixes** | Pointer to **[]string** |  | [optional] 
**LanSegment** | Pointer to **int64** |  | [optional] 
**NatTranslationMode** | Pointer to [**ManaV2ExtranetNatTranslationMode**](ManaV2ExtranetNatTranslationMode.md) |  | [optional] 
**NumCustomers** | Pointer to **int32** | Number of customers subscribed to the service | [optional] 
**ServiceId** | Pointer to **int64** | Producer service id | [optional] 
**ServicePrefixes** | Pointer to [**[]ManaV2B2bExtranetPrefixTag**](ManaV2B2bExtranetPrefixTag.md) |  | [optional] 

## Methods

### NewManaV2B2bExtranetMatch

`func NewManaV2B2bExtranetMatch() *ManaV2B2bExtranetMatch`

NewManaV2B2bExtranetMatch instantiates a new ManaV2B2bExtranetMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2B2bExtranetMatchWithDefaults

`func NewManaV2B2bExtranetMatchWithDefaults() *ManaV2B2bExtranetMatch`

NewManaV2B2bExtranetMatchWithDefaults instantiates a new ManaV2B2bExtranetMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerPrefixes

`func (o *ManaV2B2bExtranetMatch) GetConsumerPrefixes() []string`

GetConsumerPrefixes returns the ConsumerPrefixes field if non-nil, zero value otherwise.

### GetConsumerPrefixesOk

`func (o *ManaV2B2bExtranetMatch) GetConsumerPrefixesOk() (*[]string, bool)`

GetConsumerPrefixesOk returns a tuple with the ConsumerPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerPrefixes

`func (o *ManaV2B2bExtranetMatch) SetConsumerPrefixes(v []string)`

SetConsumerPrefixes sets ConsumerPrefixes field to given value.

### HasConsumerPrefixes

`func (o *ManaV2B2bExtranetMatch) HasConsumerPrefixes() bool`

HasConsumerPrefixes returns a boolean if a field has been set.

### GetLanSegment

`func (o *ManaV2B2bExtranetMatch) GetLanSegment() int64`

GetLanSegment returns the LanSegment field if non-nil, zero value otherwise.

### GetLanSegmentOk

`func (o *ManaV2B2bExtranetMatch) GetLanSegmentOk() (*int64, bool)`

GetLanSegmentOk returns a tuple with the LanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegment

`func (o *ManaV2B2bExtranetMatch) SetLanSegment(v int64)`

SetLanSegment sets LanSegment field to given value.

### HasLanSegment

`func (o *ManaV2B2bExtranetMatch) HasLanSegment() bool`

HasLanSegment returns a boolean if a field has been set.

### GetNatTranslationMode

`func (o *ManaV2B2bExtranetMatch) GetNatTranslationMode() ManaV2ExtranetNatTranslationMode`

GetNatTranslationMode returns the NatTranslationMode field if non-nil, zero value otherwise.

### GetNatTranslationModeOk

`func (o *ManaV2B2bExtranetMatch) GetNatTranslationModeOk() (*ManaV2ExtranetNatTranslationMode, bool)`

GetNatTranslationModeOk returns a tuple with the NatTranslationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTranslationMode

`func (o *ManaV2B2bExtranetMatch) SetNatTranslationMode(v ManaV2ExtranetNatTranslationMode)`

SetNatTranslationMode sets NatTranslationMode field to given value.

### HasNatTranslationMode

`func (o *ManaV2B2bExtranetMatch) HasNatTranslationMode() bool`

HasNatTranslationMode returns a boolean if a field has been set.

### GetNumCustomers

`func (o *ManaV2B2bExtranetMatch) GetNumCustomers() int32`

GetNumCustomers returns the NumCustomers field if non-nil, zero value otherwise.

### GetNumCustomersOk

`func (o *ManaV2B2bExtranetMatch) GetNumCustomersOk() (*int32, bool)`

GetNumCustomersOk returns a tuple with the NumCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCustomers

`func (o *ManaV2B2bExtranetMatch) SetNumCustomers(v int32)`

SetNumCustomers sets NumCustomers field to given value.

### HasNumCustomers

`func (o *ManaV2B2bExtranetMatch) HasNumCustomers() bool`

HasNumCustomers returns a boolean if a field has been set.

### GetServiceId

`func (o *ManaV2B2bExtranetMatch) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ManaV2B2bExtranetMatch) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ManaV2B2bExtranetMatch) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ManaV2B2bExtranetMatch) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServicePrefixes

`func (o *ManaV2B2bExtranetMatch) GetServicePrefixes() []ManaV2B2bExtranetPrefixTag`

GetServicePrefixes returns the ServicePrefixes field if non-nil, zero value otherwise.

### GetServicePrefixesOk

`func (o *ManaV2B2bExtranetMatch) GetServicePrefixesOk() (*[]ManaV2B2bExtranetPrefixTag, bool)`

GetServicePrefixesOk returns a tuple with the ServicePrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrefixes

`func (o *ManaV2B2bExtranetMatch) SetServicePrefixes(v []ManaV2B2bExtranetPrefixTag)`

SetServicePrefixes sets ServicePrefixes field to given value.

### HasServicePrefixes

`func (o *ManaV2B2bExtranetMatch) HasServicePrefixes() bool`

HasServicePrefixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


