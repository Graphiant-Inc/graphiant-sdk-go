# ManaV2ExtranetConsumerNatRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutsideNatPrefix** | Pointer to **string** | Optional nat prefix associated with a service prefix with an empty string indicating no NATing | [optional] 
**ServicePrefix** | **string** | Service prefix for the NAT rule (required) | 

## Methods

### NewManaV2ExtranetConsumerNatRule

`func NewManaV2ExtranetConsumerNatRule(servicePrefix string, ) *ManaV2ExtranetConsumerNatRule`

NewManaV2ExtranetConsumerNatRule instantiates a new ManaV2ExtranetConsumerNatRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2ExtranetConsumerNatRuleWithDefaults

`func NewManaV2ExtranetConsumerNatRuleWithDefaults() *ManaV2ExtranetConsumerNatRule`

NewManaV2ExtranetConsumerNatRuleWithDefaults instantiates a new ManaV2ExtranetConsumerNatRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutsideNatPrefix

`func (o *ManaV2ExtranetConsumerNatRule) GetOutsideNatPrefix() string`

GetOutsideNatPrefix returns the OutsideNatPrefix field if non-nil, zero value otherwise.

### GetOutsideNatPrefixOk

`func (o *ManaV2ExtranetConsumerNatRule) GetOutsideNatPrefixOk() (*string, bool)`

GetOutsideNatPrefixOk returns a tuple with the OutsideNatPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideNatPrefix

`func (o *ManaV2ExtranetConsumerNatRule) SetOutsideNatPrefix(v string)`

SetOutsideNatPrefix sets OutsideNatPrefix field to given value.

### HasOutsideNatPrefix

`func (o *ManaV2ExtranetConsumerNatRule) HasOutsideNatPrefix() bool`

HasOutsideNatPrefix returns a boolean if a field has been set.

### GetServicePrefix

`func (o *ManaV2ExtranetConsumerNatRule) GetServicePrefix() string`

GetServicePrefix returns the ServicePrefix field if non-nil, zero value otherwise.

### GetServicePrefixOk

`func (o *ManaV2ExtranetConsumerNatRule) GetServicePrefixOk() (*string, bool)`

GetServicePrefixOk returns a tuple with the ServicePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrefix

`func (o *ManaV2ExtranetConsumerNatRule) SetServicePrefix(v string)`

SetServicePrefix sets ServicePrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


