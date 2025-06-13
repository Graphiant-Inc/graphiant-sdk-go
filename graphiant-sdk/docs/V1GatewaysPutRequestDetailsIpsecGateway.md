# V1GatewaysPutRequestDetailsIpsecGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationAddress** | Pointer to **string** |  | [optional] 
**IkeInitiator** | Pointer to **bool** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RemoteIkePeerIdentity** | Pointer to **string** |  | [optional] 
**Routing** | Pointer to [**V1GatewaysPutRequestDetailsIpsecGatewayRouting**](V1GatewaysPutRequestDetailsIpsecGatewayRouting.md) |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**Tunnel1** | Pointer to [**V1GatewaysPutRequestDetailsIpsecGatewayTunnel1**](V1GatewaysPutRequestDetailsIpsecGatewayTunnel1.md) |  | [optional] 
**Tunnel2** | Pointer to [**V1GatewaysPutRequestDetailsIpsecGatewayTunnel1**](V1GatewaysPutRequestDetailsIpsecGatewayTunnel1.md) |  | [optional] 
**VpnProfile** | Pointer to **string** |  | [optional] 

## Methods

### NewV1GatewaysPutRequestDetailsIpsecGateway

`func NewV1GatewaysPutRequestDetailsIpsecGateway() *V1GatewaysPutRequestDetailsIpsecGateway`

NewV1GatewaysPutRequestDetailsIpsecGateway instantiates a new V1GatewaysPutRequestDetailsIpsecGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GatewaysPutRequestDetailsIpsecGatewayWithDefaults

`func NewV1GatewaysPutRequestDetailsIpsecGatewayWithDefaults() *V1GatewaysPutRequestDetailsIpsecGateway`

NewV1GatewaysPutRequestDetailsIpsecGatewayWithDefaults instantiates a new V1GatewaysPutRequestDetailsIpsecGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationAddress

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetIkeInitiator

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetIkeInitiator() bool`

GetIkeInitiator returns the IkeInitiator field if non-nil, zero value otherwise.

### GetIkeInitiatorOk

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetIkeInitiatorOk() (*bool, bool)`

GetIkeInitiatorOk returns a tuple with the IkeInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeInitiator

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) SetIkeInitiator(v bool)`

SetIkeInitiator sets IkeInitiator field to given value.

### HasIkeInitiator

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) HasIkeInitiator() bool`

HasIkeInitiator returns a boolean if a field has been set.

### GetMtu

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemoteIkePeerIdentity

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetRemoteIkePeerIdentity() string`

GetRemoteIkePeerIdentity returns the RemoteIkePeerIdentity field if non-nil, zero value otherwise.

### GetRemoteIkePeerIdentityOk

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetRemoteIkePeerIdentityOk() (*string, bool)`

GetRemoteIkePeerIdentityOk returns a tuple with the RemoteIkePeerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIkePeerIdentity

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) SetRemoteIkePeerIdentity(v string)`

SetRemoteIkePeerIdentity sets RemoteIkePeerIdentity field to given value.

### HasRemoteIkePeerIdentity

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) HasRemoteIkePeerIdentity() bool`

HasRemoteIkePeerIdentity returns a boolean if a field has been set.

### GetRouting

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetRouting() V1GatewaysPutRequestDetailsIpsecGatewayRouting`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetRoutingOk() (*V1GatewaysPutRequestDetailsIpsecGatewayRouting, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) SetRouting(v V1GatewaysPutRequestDetailsIpsecGatewayRouting)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetTcpMss

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetTunnel1

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetTunnel1() V1GatewaysPutRequestDetailsIpsecGatewayTunnel1`

GetTunnel1 returns the Tunnel1 field if non-nil, zero value otherwise.

### GetTunnel1Ok

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetTunnel1Ok() (*V1GatewaysPutRequestDetailsIpsecGatewayTunnel1, bool)`

GetTunnel1Ok returns a tuple with the Tunnel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel1

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) SetTunnel1(v V1GatewaysPutRequestDetailsIpsecGatewayTunnel1)`

SetTunnel1 sets Tunnel1 field to given value.

### HasTunnel1

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) HasTunnel1() bool`

HasTunnel1 returns a boolean if a field has been set.

### GetTunnel2

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetTunnel2() V1GatewaysPutRequestDetailsIpsecGatewayTunnel1`

GetTunnel2 returns the Tunnel2 field if non-nil, zero value otherwise.

### GetTunnel2Ok

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetTunnel2Ok() (*V1GatewaysPutRequestDetailsIpsecGatewayTunnel1, bool)`

GetTunnel2Ok returns a tuple with the Tunnel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel2

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) SetTunnel2(v V1GatewaysPutRequestDetailsIpsecGatewayTunnel1)`

SetTunnel2 sets Tunnel2 field to given value.

### HasTunnel2

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) HasTunnel2() bool`

HasTunnel2 returns a boolean if a field has been set.

### GetVpnProfile

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetVpnProfile() string`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) GetVpnProfileOk() (*string, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) SetVpnProfile(v string)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *V1GatewaysPutRequestDetailsIpsecGateway) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


