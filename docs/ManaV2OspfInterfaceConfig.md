# ManaV2OspfInterfaceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**ManaV2NullableOspfAuthenticationConfig**](ManaV2NullableOspfAuthenticationConfig.md) |  | [optional] 
**Bfd** | Pointer to [**ManaV2NullableBfdInstanceConfig**](ManaV2NullableBfdInstanceConfig.md) |  | [optional] 
**Cost** | Pointer to **int32** |  | [optional] 
**DeadIntervalValue** | Pointer to [**ManaV2NullableOspfDeadIntervalValue**](ManaV2NullableOspfDeadIntervalValue.md) |  | [optional] 
**DrPriority** | Pointer to **int32** |  | [optional] 
**HelloIntervalValue** | Pointer to [**ManaV2NullableOspfHelloIntervalValue**](ManaV2NullableOspfHelloIntervalValue.md) |  | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**MtuIgnore** | Pointer to **bool** |  | [optional] 
**PrefixSid** | Pointer to **int32** |  | [optional] 
**RetransmitIntervalValue** | Pointer to [**ManaV2NullableOspfRetransmitIntervalValue**](ManaV2NullableOspfRetransmitIntervalValue.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewManaV2OspfInterfaceConfig

`func NewManaV2OspfInterfaceConfig() *ManaV2OspfInterfaceConfig`

NewManaV2OspfInterfaceConfig instantiates a new ManaV2OspfInterfaceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2OspfInterfaceConfigWithDefaults

`func NewManaV2OspfInterfaceConfigWithDefaults() *ManaV2OspfInterfaceConfig`

NewManaV2OspfInterfaceConfigWithDefaults instantiates a new ManaV2OspfInterfaceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *ManaV2OspfInterfaceConfig) GetAuthentication() ManaV2NullableOspfAuthenticationConfig`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ManaV2OspfInterfaceConfig) GetAuthenticationOk() (*ManaV2NullableOspfAuthenticationConfig, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ManaV2OspfInterfaceConfig) SetAuthentication(v ManaV2NullableOspfAuthenticationConfig)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *ManaV2OspfInterfaceConfig) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetBfd

`func (o *ManaV2OspfInterfaceConfig) GetBfd() ManaV2NullableBfdInstanceConfig`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *ManaV2OspfInterfaceConfig) GetBfdOk() (*ManaV2NullableBfdInstanceConfig, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *ManaV2OspfInterfaceConfig) SetBfd(v ManaV2NullableBfdInstanceConfig)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *ManaV2OspfInterfaceConfig) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetCost

`func (o *ManaV2OspfInterfaceConfig) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ManaV2OspfInterfaceConfig) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ManaV2OspfInterfaceConfig) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ManaV2OspfInterfaceConfig) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDeadIntervalValue

`func (o *ManaV2OspfInterfaceConfig) GetDeadIntervalValue() ManaV2NullableOspfDeadIntervalValue`

GetDeadIntervalValue returns the DeadIntervalValue field if non-nil, zero value otherwise.

### GetDeadIntervalValueOk

`func (o *ManaV2OspfInterfaceConfig) GetDeadIntervalValueOk() (*ManaV2NullableOspfDeadIntervalValue, bool)`

GetDeadIntervalValueOk returns a tuple with the DeadIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadIntervalValue

`func (o *ManaV2OspfInterfaceConfig) SetDeadIntervalValue(v ManaV2NullableOspfDeadIntervalValue)`

SetDeadIntervalValue sets DeadIntervalValue field to given value.

### HasDeadIntervalValue

`func (o *ManaV2OspfInterfaceConfig) HasDeadIntervalValue() bool`

HasDeadIntervalValue returns a boolean if a field has been set.

### GetDrPriority

`func (o *ManaV2OspfInterfaceConfig) GetDrPriority() int32`

GetDrPriority returns the DrPriority field if non-nil, zero value otherwise.

### GetDrPriorityOk

`func (o *ManaV2OspfInterfaceConfig) GetDrPriorityOk() (*int32, bool)`

GetDrPriorityOk returns a tuple with the DrPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrPriority

`func (o *ManaV2OspfInterfaceConfig) SetDrPriority(v int32)`

SetDrPriority sets DrPriority field to given value.

### HasDrPriority

`func (o *ManaV2OspfInterfaceConfig) HasDrPriority() bool`

HasDrPriority returns a boolean if a field has been set.

### GetHelloIntervalValue

`func (o *ManaV2OspfInterfaceConfig) GetHelloIntervalValue() ManaV2NullableOspfHelloIntervalValue`

GetHelloIntervalValue returns the HelloIntervalValue field if non-nil, zero value otherwise.

### GetHelloIntervalValueOk

`func (o *ManaV2OspfInterfaceConfig) GetHelloIntervalValueOk() (*ManaV2NullableOspfHelloIntervalValue, bool)`

GetHelloIntervalValueOk returns a tuple with the HelloIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloIntervalValue

`func (o *ManaV2OspfInterfaceConfig) SetHelloIntervalValue(v ManaV2NullableOspfHelloIntervalValue)`

SetHelloIntervalValue sets HelloIntervalValue field to given value.

### HasHelloIntervalValue

`func (o *ManaV2OspfInterfaceConfig) HasHelloIntervalValue() bool`

HasHelloIntervalValue returns a boolean if a field has been set.

### GetInterfaceName

`func (o *ManaV2OspfInterfaceConfig) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *ManaV2OspfInterfaceConfig) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *ManaV2OspfInterfaceConfig) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *ManaV2OspfInterfaceConfig) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetMtu

`func (o *ManaV2OspfInterfaceConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *ManaV2OspfInterfaceConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *ManaV2OspfInterfaceConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *ManaV2OspfInterfaceConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetMtuIgnore

`func (o *ManaV2OspfInterfaceConfig) GetMtuIgnore() bool`

GetMtuIgnore returns the MtuIgnore field if non-nil, zero value otherwise.

### GetMtuIgnoreOk

`func (o *ManaV2OspfInterfaceConfig) GetMtuIgnoreOk() (*bool, bool)`

GetMtuIgnoreOk returns a tuple with the MtuIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtuIgnore

`func (o *ManaV2OspfInterfaceConfig) SetMtuIgnore(v bool)`

SetMtuIgnore sets MtuIgnore field to given value.

### HasMtuIgnore

`func (o *ManaV2OspfInterfaceConfig) HasMtuIgnore() bool`

HasMtuIgnore returns a boolean if a field has been set.

### GetPrefixSid

`func (o *ManaV2OspfInterfaceConfig) GetPrefixSid() int32`

GetPrefixSid returns the PrefixSid field if non-nil, zero value otherwise.

### GetPrefixSidOk

`func (o *ManaV2OspfInterfaceConfig) GetPrefixSidOk() (*int32, bool)`

GetPrefixSidOk returns a tuple with the PrefixSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSid

`func (o *ManaV2OspfInterfaceConfig) SetPrefixSid(v int32)`

SetPrefixSid sets PrefixSid field to given value.

### HasPrefixSid

`func (o *ManaV2OspfInterfaceConfig) HasPrefixSid() bool`

HasPrefixSid returns a boolean if a field has been set.

### GetRetransmitIntervalValue

`func (o *ManaV2OspfInterfaceConfig) GetRetransmitIntervalValue() ManaV2NullableOspfRetransmitIntervalValue`

GetRetransmitIntervalValue returns the RetransmitIntervalValue field if non-nil, zero value otherwise.

### GetRetransmitIntervalValueOk

`func (o *ManaV2OspfInterfaceConfig) GetRetransmitIntervalValueOk() (*ManaV2NullableOspfRetransmitIntervalValue, bool)`

GetRetransmitIntervalValueOk returns a tuple with the RetransmitIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitIntervalValue

`func (o *ManaV2OspfInterfaceConfig) SetRetransmitIntervalValue(v ManaV2NullableOspfRetransmitIntervalValue)`

SetRetransmitIntervalValue sets RetransmitIntervalValue field to given value.

### HasRetransmitIntervalValue

`func (o *ManaV2OspfInterfaceConfig) HasRetransmitIntervalValue() bool`

HasRetransmitIntervalValue returns a boolean if a field has been set.

### GetType

`func (o *ManaV2OspfInterfaceConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManaV2OspfInterfaceConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManaV2OspfInterfaceConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManaV2OspfInterfaceConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


