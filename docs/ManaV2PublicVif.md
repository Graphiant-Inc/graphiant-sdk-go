# ManaV2PublicVif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoveringPrefixes** | [**ManaV2PublicVifDynamic**](ManaV2PublicVifDynamic.md) |  | 
**FixedPrefixes** | [**ManaV2PublicVifFixed**](ManaV2PublicVifFixed.md) |  | 
**Type** | **string** | Type of Public VIF dynamic/fixed (required) | 

## Methods

### NewManaV2PublicVif

`func NewManaV2PublicVif(coveringPrefixes ManaV2PublicVifDynamic, fixedPrefixes ManaV2PublicVifFixed, type_ string, ) *ManaV2PublicVif`

NewManaV2PublicVif instantiates a new ManaV2PublicVif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PublicVifWithDefaults

`func NewManaV2PublicVifWithDefaults() *ManaV2PublicVif`

NewManaV2PublicVifWithDefaults instantiates a new ManaV2PublicVif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoveringPrefixes

`func (o *ManaV2PublicVif) GetCoveringPrefixes() ManaV2PublicVifDynamic`

GetCoveringPrefixes returns the CoveringPrefixes field if non-nil, zero value otherwise.

### GetCoveringPrefixesOk

`func (o *ManaV2PublicVif) GetCoveringPrefixesOk() (*ManaV2PublicVifDynamic, bool)`

GetCoveringPrefixesOk returns a tuple with the CoveringPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveringPrefixes

`func (o *ManaV2PublicVif) SetCoveringPrefixes(v ManaV2PublicVifDynamic)`

SetCoveringPrefixes sets CoveringPrefixes field to given value.


### GetFixedPrefixes

`func (o *ManaV2PublicVif) GetFixedPrefixes() ManaV2PublicVifFixed`

GetFixedPrefixes returns the FixedPrefixes field if non-nil, zero value otherwise.

### GetFixedPrefixesOk

`func (o *ManaV2PublicVif) GetFixedPrefixesOk() (*ManaV2PublicVifFixed, bool)`

GetFixedPrefixesOk returns a tuple with the FixedPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedPrefixes

`func (o *ManaV2PublicVif) SetFixedPrefixes(v ManaV2PublicVifFixed)`

SetFixedPrefixes sets FixedPrefixes field to given value.


### GetType

`func (o *ManaV2PublicVif) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManaV2PublicVif) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManaV2PublicVif) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


