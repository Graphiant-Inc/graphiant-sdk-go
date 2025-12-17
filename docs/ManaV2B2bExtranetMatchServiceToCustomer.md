# ManaV2B2bExtranetMatchServiceToCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the service being subscribed by the customer (required) | 
**LanSegment** | Pointer to **int64** |  | [optional] 
**Nat** | [**[]ManaV2B2bNat**](ManaV2B2bNat.md) |  | 
**NumCustomers** | Pointer to **int32** | Number of customers subscribed to the service | [optional] 
**ServicePrefixes** | [**[]ManaV2B2bExtranetPrefixTag**](ManaV2B2bExtranetPrefixTag.md) |  | 

## Methods

### NewManaV2B2bExtranetMatchServiceToCustomer

`func NewManaV2B2bExtranetMatchServiceToCustomer(id int64, nat []ManaV2B2bNat, servicePrefixes []ManaV2B2bExtranetPrefixTag, ) *ManaV2B2bExtranetMatchServiceToCustomer`

NewManaV2B2bExtranetMatchServiceToCustomer instantiates a new ManaV2B2bExtranetMatchServiceToCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2B2bExtranetMatchServiceToCustomerWithDefaults

`func NewManaV2B2bExtranetMatchServiceToCustomerWithDefaults() *ManaV2B2bExtranetMatchServiceToCustomer`

NewManaV2B2bExtranetMatchServiceToCustomerWithDefaults instantiates a new ManaV2B2bExtranetMatchServiceToCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) SetId(v int64)`

SetId sets Id field to given value.


### GetLanSegment

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) GetLanSegment() int64`

GetLanSegment returns the LanSegment field if non-nil, zero value otherwise.

### GetLanSegmentOk

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) GetLanSegmentOk() (*int64, bool)`

GetLanSegmentOk returns a tuple with the LanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegment

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) SetLanSegment(v int64)`

SetLanSegment sets LanSegment field to given value.

### HasLanSegment

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) HasLanSegment() bool`

HasLanSegment returns a boolean if a field has been set.

### GetNat

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) GetNat() []ManaV2B2bNat`

GetNat returns the Nat field if non-nil, zero value otherwise.

### GetNatOk

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) GetNatOk() (*[]ManaV2B2bNat, bool)`

GetNatOk returns a tuple with the Nat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNat

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) SetNat(v []ManaV2B2bNat)`

SetNat sets Nat field to given value.


### GetNumCustomers

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) GetNumCustomers() int32`

GetNumCustomers returns the NumCustomers field if non-nil, zero value otherwise.

### GetNumCustomersOk

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) GetNumCustomersOk() (*int32, bool)`

GetNumCustomersOk returns a tuple with the NumCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCustomers

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) SetNumCustomers(v int32)`

SetNumCustomers sets NumCustomers field to given value.

### HasNumCustomers

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) HasNumCustomers() bool`

HasNumCustomers returns a boolean if a field has been set.

### GetServicePrefixes

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) GetServicePrefixes() []ManaV2B2bExtranetPrefixTag`

GetServicePrefixes returns the ServicePrefixes field if non-nil, zero value otherwise.

### GetServicePrefixesOk

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) GetServicePrefixesOk() (*[]ManaV2B2bExtranetPrefixTag, bool)`

GetServicePrefixesOk returns a tuple with the ServicePrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrefixes

`func (o *ManaV2B2bExtranetMatchServiceToCustomer) SetServicePrefixes(v []ManaV2B2bExtranetPrefixTag)`

SetServicePrefixes sets ServicePrefixes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


