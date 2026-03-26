# ManaV2BgpDynamicNeighborOperPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastOperStatusChange** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**LocalAddress** | Pointer to **string** | Local address used for this peer session, if known from oper | [optional] 
**OperStatus** | Pointer to **bool** | True when the BGP session to this peer is operationally up (e.g. established) | [optional] 
**PeerAsn** | Pointer to **int32** | Peer ASN from oper, if reported | [optional] 
**RemoteAddress** | Pointer to **string** | Peer address from device oper state (IPv4/IPv6; may include IPv6 zone id) | [optional] 
**State** | Pointer to **string** | BGP FSM state for this peer session | [optional] 

## Methods

### NewManaV2BgpDynamicNeighborOperPeer

`func NewManaV2BgpDynamicNeighborOperPeer() *ManaV2BgpDynamicNeighborOperPeer`

NewManaV2BgpDynamicNeighborOperPeer instantiates a new ManaV2BgpDynamicNeighborOperPeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2BgpDynamicNeighborOperPeerWithDefaults

`func NewManaV2BgpDynamicNeighborOperPeerWithDefaults() *ManaV2BgpDynamicNeighborOperPeer`

NewManaV2BgpDynamicNeighborOperPeerWithDefaults instantiates a new ManaV2BgpDynamicNeighborOperPeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastOperStatusChange

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetLastOperStatusChange() GoogleProtobufTimestamp`

GetLastOperStatusChange returns the LastOperStatusChange field if non-nil, zero value otherwise.

### GetLastOperStatusChangeOk

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetLastOperStatusChangeOk() (*GoogleProtobufTimestamp, bool)`

GetLastOperStatusChangeOk returns a tuple with the LastOperStatusChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOperStatusChange

`func (o *ManaV2BgpDynamicNeighborOperPeer) SetLastOperStatusChange(v GoogleProtobufTimestamp)`

SetLastOperStatusChange sets LastOperStatusChange field to given value.

### HasLastOperStatusChange

`func (o *ManaV2BgpDynamicNeighborOperPeer) HasLastOperStatusChange() bool`

HasLastOperStatusChange returns a boolean if a field has been set.

### GetLocalAddress

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *ManaV2BgpDynamicNeighborOperPeer) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *ManaV2BgpDynamicNeighborOperPeer) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetOperStatus

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetOperStatus() bool`

GetOperStatus returns the OperStatus field if non-nil, zero value otherwise.

### GetOperStatusOk

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetOperStatusOk() (*bool, bool)`

GetOperStatusOk returns a tuple with the OperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStatus

`func (o *ManaV2BgpDynamicNeighborOperPeer) SetOperStatus(v bool)`

SetOperStatus sets OperStatus field to given value.

### HasOperStatus

`func (o *ManaV2BgpDynamicNeighborOperPeer) HasOperStatus() bool`

HasOperStatus returns a boolean if a field has been set.

### GetPeerAsn

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *ManaV2BgpDynamicNeighborOperPeer) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *ManaV2BgpDynamicNeighborOperPeer) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *ManaV2BgpDynamicNeighborOperPeer) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *ManaV2BgpDynamicNeighborOperPeer) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetState

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManaV2BgpDynamicNeighborOperPeer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManaV2BgpDynamicNeighborOperPeer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ManaV2BgpDynamicNeighborOperPeer) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


