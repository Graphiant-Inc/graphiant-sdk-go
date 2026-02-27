# ManaV2InterfaceIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**ManaV2NullableAddress**](ManaV2NullableAddress.md) |  | [optional] 
**Dhcp** | Pointer to [**ManaV2InterfaceDhcpConfig**](ManaV2InterfaceDhcpConfig.md) |  | [optional] 
**Gw** | Pointer to [**ManaV2NullableGatewayConfig**](ManaV2NullableGatewayConfig.md) |  | [optional] 
**Vrrp** | Pointer to [**ManaV2NullableVrrpGroupConfig**](ManaV2NullableVrrpGroupConfig.md) |  | [optional] 

## Methods

### NewManaV2InterfaceIpConfig

`func NewManaV2InterfaceIpConfig() *ManaV2InterfaceIpConfig`

NewManaV2InterfaceIpConfig instantiates a new ManaV2InterfaceIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2InterfaceIpConfigWithDefaults

`func NewManaV2InterfaceIpConfigWithDefaults() *ManaV2InterfaceIpConfig`

NewManaV2InterfaceIpConfigWithDefaults instantiates a new ManaV2InterfaceIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ManaV2InterfaceIpConfig) GetAddress() ManaV2NullableAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ManaV2InterfaceIpConfig) GetAddressOk() (*ManaV2NullableAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ManaV2InterfaceIpConfig) SetAddress(v ManaV2NullableAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ManaV2InterfaceIpConfig) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDhcp

`func (o *ManaV2InterfaceIpConfig) GetDhcp() ManaV2InterfaceDhcpConfig`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *ManaV2InterfaceIpConfig) GetDhcpOk() (*ManaV2InterfaceDhcpConfig, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *ManaV2InterfaceIpConfig) SetDhcp(v ManaV2InterfaceDhcpConfig)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *ManaV2InterfaceIpConfig) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetGw

`func (o *ManaV2InterfaceIpConfig) GetGw() ManaV2NullableGatewayConfig`

GetGw returns the Gw field if non-nil, zero value otherwise.

### GetGwOk

`func (o *ManaV2InterfaceIpConfig) GetGwOk() (*ManaV2NullableGatewayConfig, bool)`

GetGwOk returns a tuple with the Gw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGw

`func (o *ManaV2InterfaceIpConfig) SetGw(v ManaV2NullableGatewayConfig)`

SetGw sets Gw field to given value.

### HasGw

`func (o *ManaV2InterfaceIpConfig) HasGw() bool`

HasGw returns a boolean if a field has been set.

### GetVrrp

`func (o *ManaV2InterfaceIpConfig) GetVrrp() ManaV2NullableVrrpGroupConfig`

GetVrrp returns the Vrrp field if non-nil, zero value otherwise.

### GetVrrpOk

`func (o *ManaV2InterfaceIpConfig) GetVrrpOk() (*ManaV2NullableVrrpGroupConfig, bool)`

GetVrrpOk returns a tuple with the Vrrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrp

`func (o *ManaV2InterfaceIpConfig) SetVrrp(v ManaV2NullableVrrpGroupConfig)`

SetVrrp sets Vrrp field to given value.

### HasVrrp

`func (o *ManaV2InterfaceIpConfig) HasVrrp() bool`

HasVrrp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


