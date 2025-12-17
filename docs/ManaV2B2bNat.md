# ManaV2B2bNat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutsideNatPrefix** | Pointer to **string** | Outside NAT prefix for the customer prefix | [optional] 
**Prefix** | **string** | Prefix imported into the service (required) | 

## Methods

### NewManaV2B2bNat

`func NewManaV2B2bNat(prefix string, ) *ManaV2B2bNat`

NewManaV2B2bNat instantiates a new ManaV2B2bNat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2B2bNatWithDefaults

`func NewManaV2B2bNatWithDefaults() *ManaV2B2bNat`

NewManaV2B2bNatWithDefaults instantiates a new ManaV2B2bNat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutsideNatPrefix

`func (o *ManaV2B2bNat) GetOutsideNatPrefix() string`

GetOutsideNatPrefix returns the OutsideNatPrefix field if non-nil, zero value otherwise.

### GetOutsideNatPrefixOk

`func (o *ManaV2B2bNat) GetOutsideNatPrefixOk() (*string, bool)`

GetOutsideNatPrefixOk returns a tuple with the OutsideNatPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideNatPrefix

`func (o *ManaV2B2bNat) SetOutsideNatPrefix(v string)`

SetOutsideNatPrefix sets OutsideNatPrefix field to given value.

### HasOutsideNatPrefix

`func (o *ManaV2B2bNat) HasOutsideNatPrefix() bool`

HasOutsideNatPrefix returns a boolean if a field has been set.

### GetPrefix

`func (o *ManaV2B2bNat) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ManaV2B2bNat) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ManaV2B2bNat) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


