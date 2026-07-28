# ManaV2ExtranetNatTranslationPeerToPeerPrefix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutsideNatPrefix** | Pointer to **string** | Optional outside address presented for prefix on the far side of the attachment; omit for no NAT on that prefix | [optional] 
**Prefix** | Pointer to **string** | At match: customer export prefix. At consumer accept/update (peering): subscribed service prefix | [optional] 

## Methods

### NewManaV2ExtranetNatTranslationPeerToPeerPrefix

`func NewManaV2ExtranetNatTranslationPeerToPeerPrefix() *ManaV2ExtranetNatTranslationPeerToPeerPrefix`

NewManaV2ExtranetNatTranslationPeerToPeerPrefix instantiates a new ManaV2ExtranetNatTranslationPeerToPeerPrefix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2ExtranetNatTranslationPeerToPeerPrefixWithDefaults

`func NewManaV2ExtranetNatTranslationPeerToPeerPrefixWithDefaults() *ManaV2ExtranetNatTranslationPeerToPeerPrefix`

NewManaV2ExtranetNatTranslationPeerToPeerPrefixWithDefaults instantiates a new ManaV2ExtranetNatTranslationPeerToPeerPrefix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutsideNatPrefix

`func (o *ManaV2ExtranetNatTranslationPeerToPeerPrefix) GetOutsideNatPrefix() string`

GetOutsideNatPrefix returns the OutsideNatPrefix field if non-nil, zero value otherwise.

### GetOutsideNatPrefixOk

`func (o *ManaV2ExtranetNatTranslationPeerToPeerPrefix) GetOutsideNatPrefixOk() (*string, bool)`

GetOutsideNatPrefixOk returns a tuple with the OutsideNatPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideNatPrefix

`func (o *ManaV2ExtranetNatTranslationPeerToPeerPrefix) SetOutsideNatPrefix(v string)`

SetOutsideNatPrefix sets OutsideNatPrefix field to given value.

### HasOutsideNatPrefix

`func (o *ManaV2ExtranetNatTranslationPeerToPeerPrefix) HasOutsideNatPrefix() bool`

HasOutsideNatPrefix returns a boolean if a field has been set.

### GetPrefix

`func (o *ManaV2ExtranetNatTranslationPeerToPeerPrefix) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ManaV2ExtranetNatTranslationPeerToPeerPrefix) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ManaV2ExtranetNatTranslationPeerToPeerPrefix) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ManaV2ExtranetNatTranslationPeerToPeerPrefix) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


