# ManaV2IPsecGatewayRemotePeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationAddress** | Pointer to **string** |  | [optional] 
**IkeInitiator** | Pointer to **bool** | When true, Graphiant initiates IKE for this peer | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | Optional display name or label for this peer; used when generating tunnel names | [optional] 
**RemoteIkePeerIdentity** | Pointer to **string** | IKE identity of the remote peer | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**Tunnel1** | Pointer to [**ManaV2IPsecGatewayTunnelDetails**](ManaV2IPsecGatewayTunnelDetails.md) |  | [optional] 
**Tunnel2** | Pointer to [**ManaV2IPsecGatewayTunnelDetails**](ManaV2IPsecGatewayTunnelDetails.md) |  | [optional] 
**VpnProfile** | Pointer to **string** | Enterprise IPsec VPN profile name for this peer | [optional] 

## Methods

### NewManaV2IPsecGatewayRemotePeer

`func NewManaV2IPsecGatewayRemotePeer() *ManaV2IPsecGatewayRemotePeer`

NewManaV2IPsecGatewayRemotePeer instantiates a new ManaV2IPsecGatewayRemotePeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2IPsecGatewayRemotePeerWithDefaults

`func NewManaV2IPsecGatewayRemotePeerWithDefaults() *ManaV2IPsecGatewayRemotePeer`

NewManaV2IPsecGatewayRemotePeerWithDefaults instantiates a new ManaV2IPsecGatewayRemotePeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationAddress

`func (o *ManaV2IPsecGatewayRemotePeer) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *ManaV2IPsecGatewayRemotePeer) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *ManaV2IPsecGatewayRemotePeer) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *ManaV2IPsecGatewayRemotePeer) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetIkeInitiator

`func (o *ManaV2IPsecGatewayRemotePeer) GetIkeInitiator() bool`

GetIkeInitiator returns the IkeInitiator field if non-nil, zero value otherwise.

### GetIkeInitiatorOk

`func (o *ManaV2IPsecGatewayRemotePeer) GetIkeInitiatorOk() (*bool, bool)`

GetIkeInitiatorOk returns a tuple with the IkeInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeInitiator

`func (o *ManaV2IPsecGatewayRemotePeer) SetIkeInitiator(v bool)`

SetIkeInitiator sets IkeInitiator field to given value.

### HasIkeInitiator

`func (o *ManaV2IPsecGatewayRemotePeer) HasIkeInitiator() bool`

HasIkeInitiator returns a boolean if a field has been set.

### GetMtu

`func (o *ManaV2IPsecGatewayRemotePeer) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *ManaV2IPsecGatewayRemotePeer) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *ManaV2IPsecGatewayRemotePeer) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *ManaV2IPsecGatewayRemotePeer) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *ManaV2IPsecGatewayRemotePeer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2IPsecGatewayRemotePeer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2IPsecGatewayRemotePeer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2IPsecGatewayRemotePeer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemoteIkePeerIdentity

`func (o *ManaV2IPsecGatewayRemotePeer) GetRemoteIkePeerIdentity() string`

GetRemoteIkePeerIdentity returns the RemoteIkePeerIdentity field if non-nil, zero value otherwise.

### GetRemoteIkePeerIdentityOk

`func (o *ManaV2IPsecGatewayRemotePeer) GetRemoteIkePeerIdentityOk() (*string, bool)`

GetRemoteIkePeerIdentityOk returns a tuple with the RemoteIkePeerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIkePeerIdentity

`func (o *ManaV2IPsecGatewayRemotePeer) SetRemoteIkePeerIdentity(v string)`

SetRemoteIkePeerIdentity sets RemoteIkePeerIdentity field to given value.

### HasRemoteIkePeerIdentity

`func (o *ManaV2IPsecGatewayRemotePeer) HasRemoteIkePeerIdentity() bool`

HasRemoteIkePeerIdentity returns a boolean if a field has been set.

### GetTcpMss

`func (o *ManaV2IPsecGatewayRemotePeer) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *ManaV2IPsecGatewayRemotePeer) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *ManaV2IPsecGatewayRemotePeer) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *ManaV2IPsecGatewayRemotePeer) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetTunnel1

`func (o *ManaV2IPsecGatewayRemotePeer) GetTunnel1() ManaV2IPsecGatewayTunnelDetails`

GetTunnel1 returns the Tunnel1 field if non-nil, zero value otherwise.

### GetTunnel1Ok

`func (o *ManaV2IPsecGatewayRemotePeer) GetTunnel1Ok() (*ManaV2IPsecGatewayTunnelDetails, bool)`

GetTunnel1Ok returns a tuple with the Tunnel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel1

`func (o *ManaV2IPsecGatewayRemotePeer) SetTunnel1(v ManaV2IPsecGatewayTunnelDetails)`

SetTunnel1 sets Tunnel1 field to given value.

### HasTunnel1

`func (o *ManaV2IPsecGatewayRemotePeer) HasTunnel1() bool`

HasTunnel1 returns a boolean if a field has been set.

### GetTunnel2

`func (o *ManaV2IPsecGatewayRemotePeer) GetTunnel2() ManaV2IPsecGatewayTunnelDetails`

GetTunnel2 returns the Tunnel2 field if non-nil, zero value otherwise.

### GetTunnel2Ok

`func (o *ManaV2IPsecGatewayRemotePeer) GetTunnel2Ok() (*ManaV2IPsecGatewayTunnelDetails, bool)`

GetTunnel2Ok returns a tuple with the Tunnel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel2

`func (o *ManaV2IPsecGatewayRemotePeer) SetTunnel2(v ManaV2IPsecGatewayTunnelDetails)`

SetTunnel2 sets Tunnel2 field to given value.

### HasTunnel2

`func (o *ManaV2IPsecGatewayRemotePeer) HasTunnel2() bool`

HasTunnel2 returns a boolean if a field has been set.

### GetVpnProfile

`func (o *ManaV2IPsecGatewayRemotePeer) GetVpnProfile() string`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *ManaV2IPsecGatewayRemotePeer) GetVpnProfileOk() (*string, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *ManaV2IPsecGatewayRemotePeer) SetVpnProfile(v string)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *ManaV2IPsecGatewayRemotePeer) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


