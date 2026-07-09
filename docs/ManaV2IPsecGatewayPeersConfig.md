# ManaV2IPsecGatewayPeersConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the IPsec gateway service | [optional] 
**RemotePeers** | Pointer to [**[]ManaV2IPsecGatewayRemotePeer**](ManaV2IPsecGatewayRemotePeer.md) |  | [optional] 
**Routing** | Pointer to [**ManaV2IpsecRoutingConfig**](ManaV2IpsecRoutingConfig.md) |  | [optional] 

## Methods

### NewManaV2IPsecGatewayPeersConfig

`func NewManaV2IPsecGatewayPeersConfig() *ManaV2IPsecGatewayPeersConfig`

NewManaV2IPsecGatewayPeersConfig instantiates a new ManaV2IPsecGatewayPeersConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2IPsecGatewayPeersConfigWithDefaults

`func NewManaV2IPsecGatewayPeersConfigWithDefaults() *ManaV2IPsecGatewayPeersConfig`

NewManaV2IPsecGatewayPeersConfigWithDefaults instantiates a new ManaV2IPsecGatewayPeersConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManaV2IPsecGatewayPeersConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2IPsecGatewayPeersConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2IPsecGatewayPeersConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2IPsecGatewayPeersConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemotePeers

`func (o *ManaV2IPsecGatewayPeersConfig) GetRemotePeers() []ManaV2IPsecGatewayRemotePeer`

GetRemotePeers returns the RemotePeers field if non-nil, zero value otherwise.

### GetRemotePeersOk

`func (o *ManaV2IPsecGatewayPeersConfig) GetRemotePeersOk() (*[]ManaV2IPsecGatewayRemotePeer, bool)`

GetRemotePeersOk returns a tuple with the RemotePeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePeers

`func (o *ManaV2IPsecGatewayPeersConfig) SetRemotePeers(v []ManaV2IPsecGatewayRemotePeer)`

SetRemotePeers sets RemotePeers field to given value.

### HasRemotePeers

`func (o *ManaV2IPsecGatewayPeersConfig) HasRemotePeers() bool`

HasRemotePeers returns a boolean if a field has been set.

### GetRouting

`func (o *ManaV2IPsecGatewayPeersConfig) GetRouting() ManaV2IpsecRoutingConfig`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *ManaV2IPsecGatewayPeersConfig) GetRoutingOk() (*ManaV2IpsecRoutingConfig, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *ManaV2IPsecGatewayPeersConfig) SetRouting(v ManaV2IpsecRoutingConfig)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *ManaV2IPsecGatewayPeersConfig) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


