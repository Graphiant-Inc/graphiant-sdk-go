# ManaV2BgpNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamilies** | Pointer to [**[]ManaV2BgpNeighborAddressFamily**](ManaV2BgpNeighborAddressFamily.md) |  | [optional] 
**AllowAsIn** | Pointer to **int32** |  | [optional] 
**AsOverride** | Pointer to **bool** |  | [optional] 
**Bfd** | Pointer to [**ManaV2BfdInstance**](ManaV2BfdInstance.md) |  | [optional] 
**BfdNeighbor** | Pointer to [**ManaV2BfdNeighbor**](ManaV2BfdNeighbor.md) |  | [optional] 
**BgpType** | Pointer to **string** |  | [optional] 
**DefaultOriginate** | Pointer to **string** | Set when default route needs to be advertised in BGP domain | [optional] 
**DynamicNeighborPeers** | Pointer to [**[]ManaV2BgpDynamicNeighborOperPeer**](ManaV2BgpDynamicNeighborOperPeer.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HoldTimer** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**KeepaliveTimer** | Pointer to **int32** |  | [optional] 
**LocalAddress** | Pointer to **string** |  | [optional] 
**LocalInterface** | Pointer to **string** |  | [optional] 
**MaxPrefix** | Pointer to **int32** | Maximum number of prefixes that can be received from neighbor | [optional] 
**Md5Password** | Pointer to **string** |  | [optional] 
**MultiHop** | Pointer to **int32** | Set when EBGP multi-hop functionality is enabled | [optional] 
**PeerAsn** | Pointer to **int32** |  | [optional] 
**RemoteAddress** | Pointer to **string** |  | [optional] 
**RemovePrivateAs** | Pointer to **bool** |  | [optional] 
**SendCommunity** | Pointer to **bool** | Flag for sending standard, extended, and large communities | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TimeSinceLastOperChange** | Pointer to [**GoogleProtobufDuration**](GoogleProtobufDuration.md) |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 

## Methods

### NewManaV2BgpNeighbor

`func NewManaV2BgpNeighbor() *ManaV2BgpNeighbor`

NewManaV2BgpNeighbor instantiates a new ManaV2BgpNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2BgpNeighborWithDefaults

`func NewManaV2BgpNeighborWithDefaults() *ManaV2BgpNeighbor`

NewManaV2BgpNeighborWithDefaults instantiates a new ManaV2BgpNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamilies

`func (o *ManaV2BgpNeighbor) GetAddressFamilies() []ManaV2BgpNeighborAddressFamily`

GetAddressFamilies returns the AddressFamilies field if non-nil, zero value otherwise.

### GetAddressFamiliesOk

`func (o *ManaV2BgpNeighbor) GetAddressFamiliesOk() (*[]ManaV2BgpNeighborAddressFamily, bool)`

GetAddressFamiliesOk returns a tuple with the AddressFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamilies

`func (o *ManaV2BgpNeighbor) SetAddressFamilies(v []ManaV2BgpNeighborAddressFamily)`

SetAddressFamilies sets AddressFamilies field to given value.

### HasAddressFamilies

`func (o *ManaV2BgpNeighbor) HasAddressFamilies() bool`

HasAddressFamilies returns a boolean if a field has been set.

### GetAllowAsIn

`func (o *ManaV2BgpNeighbor) GetAllowAsIn() int32`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *ManaV2BgpNeighbor) GetAllowAsInOk() (*int32, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *ManaV2BgpNeighbor) SetAllowAsIn(v int32)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *ManaV2BgpNeighbor) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetAsOverride

`func (o *ManaV2BgpNeighbor) GetAsOverride() bool`

GetAsOverride returns the AsOverride field if non-nil, zero value otherwise.

### GetAsOverrideOk

`func (o *ManaV2BgpNeighbor) GetAsOverrideOk() (*bool, bool)`

GetAsOverrideOk returns a tuple with the AsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOverride

`func (o *ManaV2BgpNeighbor) SetAsOverride(v bool)`

SetAsOverride sets AsOverride field to given value.

### HasAsOverride

`func (o *ManaV2BgpNeighbor) HasAsOverride() bool`

HasAsOverride returns a boolean if a field has been set.

### GetBfd

`func (o *ManaV2BgpNeighbor) GetBfd() ManaV2BfdInstance`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *ManaV2BgpNeighbor) GetBfdOk() (*ManaV2BfdInstance, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *ManaV2BgpNeighbor) SetBfd(v ManaV2BfdInstance)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *ManaV2BgpNeighbor) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetBfdNeighbor

`func (o *ManaV2BgpNeighbor) GetBfdNeighbor() ManaV2BfdNeighbor`

GetBfdNeighbor returns the BfdNeighbor field if non-nil, zero value otherwise.

### GetBfdNeighborOk

`func (o *ManaV2BgpNeighbor) GetBfdNeighborOk() (*ManaV2BfdNeighbor, bool)`

GetBfdNeighborOk returns a tuple with the BfdNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdNeighbor

`func (o *ManaV2BgpNeighbor) SetBfdNeighbor(v ManaV2BfdNeighbor)`

SetBfdNeighbor sets BfdNeighbor field to given value.

### HasBfdNeighbor

`func (o *ManaV2BgpNeighbor) HasBfdNeighbor() bool`

HasBfdNeighbor returns a boolean if a field has been set.

### GetBgpType

`func (o *ManaV2BgpNeighbor) GetBgpType() string`

GetBgpType returns the BgpType field if non-nil, zero value otherwise.

### GetBgpTypeOk

`func (o *ManaV2BgpNeighbor) GetBgpTypeOk() (*string, bool)`

GetBgpTypeOk returns a tuple with the BgpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpType

`func (o *ManaV2BgpNeighbor) SetBgpType(v string)`

SetBgpType sets BgpType field to given value.

### HasBgpType

`func (o *ManaV2BgpNeighbor) HasBgpType() bool`

HasBgpType returns a boolean if a field has been set.

### GetDefaultOriginate

`func (o *ManaV2BgpNeighbor) GetDefaultOriginate() string`

GetDefaultOriginate returns the DefaultOriginate field if non-nil, zero value otherwise.

### GetDefaultOriginateOk

`func (o *ManaV2BgpNeighbor) GetDefaultOriginateOk() (*string, bool)`

GetDefaultOriginateOk returns a tuple with the DefaultOriginate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOriginate

`func (o *ManaV2BgpNeighbor) SetDefaultOriginate(v string)`

SetDefaultOriginate sets DefaultOriginate field to given value.

### HasDefaultOriginate

`func (o *ManaV2BgpNeighbor) HasDefaultOriginate() bool`

HasDefaultOriginate returns a boolean if a field has been set.

### GetDynamicNeighborPeers

`func (o *ManaV2BgpNeighbor) GetDynamicNeighborPeers() []ManaV2BgpDynamicNeighborOperPeer`

GetDynamicNeighborPeers returns the DynamicNeighborPeers field if non-nil, zero value otherwise.

### GetDynamicNeighborPeersOk

`func (o *ManaV2BgpNeighbor) GetDynamicNeighborPeersOk() (*[]ManaV2BgpDynamicNeighborOperPeer, bool)`

GetDynamicNeighborPeersOk returns a tuple with the DynamicNeighborPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicNeighborPeers

`func (o *ManaV2BgpNeighbor) SetDynamicNeighborPeers(v []ManaV2BgpDynamicNeighborOperPeer)`

SetDynamicNeighborPeers sets DynamicNeighborPeers field to given value.

### HasDynamicNeighborPeers

`func (o *ManaV2BgpNeighbor) HasDynamicNeighborPeers() bool`

HasDynamicNeighborPeers returns a boolean if a field has been set.

### GetEnabled

`func (o *ManaV2BgpNeighbor) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ManaV2BgpNeighbor) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ManaV2BgpNeighbor) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ManaV2BgpNeighbor) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHoldTimer

`func (o *ManaV2BgpNeighbor) GetHoldTimer() int32`

GetHoldTimer returns the HoldTimer field if non-nil, zero value otherwise.

### GetHoldTimerOk

`func (o *ManaV2BgpNeighbor) GetHoldTimerOk() (*int32, bool)`

GetHoldTimerOk returns a tuple with the HoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTimer

`func (o *ManaV2BgpNeighbor) SetHoldTimer(v int32)`

SetHoldTimer sets HoldTimer field to given value.

### HasHoldTimer

`func (o *ManaV2BgpNeighbor) HasHoldTimer() bool`

HasHoldTimer returns a boolean if a field has been set.

### GetId

`func (o *ManaV2BgpNeighbor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2BgpNeighbor) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2BgpNeighbor) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ManaV2BgpNeighbor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeepaliveTimer

`func (o *ManaV2BgpNeighbor) GetKeepaliveTimer() int32`

GetKeepaliveTimer returns the KeepaliveTimer field if non-nil, zero value otherwise.

### GetKeepaliveTimerOk

`func (o *ManaV2BgpNeighbor) GetKeepaliveTimerOk() (*int32, bool)`

GetKeepaliveTimerOk returns a tuple with the KeepaliveTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveTimer

`func (o *ManaV2BgpNeighbor) SetKeepaliveTimer(v int32)`

SetKeepaliveTimer sets KeepaliveTimer field to given value.

### HasKeepaliveTimer

`func (o *ManaV2BgpNeighbor) HasKeepaliveTimer() bool`

HasKeepaliveTimer returns a boolean if a field has been set.

### GetLocalAddress

`func (o *ManaV2BgpNeighbor) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *ManaV2BgpNeighbor) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *ManaV2BgpNeighbor) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *ManaV2BgpNeighbor) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetLocalInterface

`func (o *ManaV2BgpNeighbor) GetLocalInterface() string`

GetLocalInterface returns the LocalInterface field if non-nil, zero value otherwise.

### GetLocalInterfaceOk

`func (o *ManaV2BgpNeighbor) GetLocalInterfaceOk() (*string, bool)`

GetLocalInterfaceOk returns a tuple with the LocalInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInterface

`func (o *ManaV2BgpNeighbor) SetLocalInterface(v string)`

SetLocalInterface sets LocalInterface field to given value.

### HasLocalInterface

`func (o *ManaV2BgpNeighbor) HasLocalInterface() bool`

HasLocalInterface returns a boolean if a field has been set.

### GetMaxPrefix

`func (o *ManaV2BgpNeighbor) GetMaxPrefix() int32`

GetMaxPrefix returns the MaxPrefix field if non-nil, zero value otherwise.

### GetMaxPrefixOk

`func (o *ManaV2BgpNeighbor) GetMaxPrefixOk() (*int32, bool)`

GetMaxPrefixOk returns a tuple with the MaxPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrefix

`func (o *ManaV2BgpNeighbor) SetMaxPrefix(v int32)`

SetMaxPrefix sets MaxPrefix field to given value.

### HasMaxPrefix

`func (o *ManaV2BgpNeighbor) HasMaxPrefix() bool`

HasMaxPrefix returns a boolean if a field has been set.

### GetMd5Password

`func (o *ManaV2BgpNeighbor) GetMd5Password() string`

GetMd5Password returns the Md5Password field if non-nil, zero value otherwise.

### GetMd5PasswordOk

`func (o *ManaV2BgpNeighbor) GetMd5PasswordOk() (*string, bool)`

GetMd5PasswordOk returns a tuple with the Md5Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Password

`func (o *ManaV2BgpNeighbor) SetMd5Password(v string)`

SetMd5Password sets Md5Password field to given value.

### HasMd5Password

`func (o *ManaV2BgpNeighbor) HasMd5Password() bool`

HasMd5Password returns a boolean if a field has been set.

### GetMultiHop

`func (o *ManaV2BgpNeighbor) GetMultiHop() int32`

GetMultiHop returns the MultiHop field if non-nil, zero value otherwise.

### GetMultiHopOk

`func (o *ManaV2BgpNeighbor) GetMultiHopOk() (*int32, bool)`

GetMultiHopOk returns a tuple with the MultiHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiHop

`func (o *ManaV2BgpNeighbor) SetMultiHop(v int32)`

SetMultiHop sets MultiHop field to given value.

### HasMultiHop

`func (o *ManaV2BgpNeighbor) HasMultiHop() bool`

HasMultiHop returns a boolean if a field has been set.

### GetPeerAsn

`func (o *ManaV2BgpNeighbor) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *ManaV2BgpNeighbor) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *ManaV2BgpNeighbor) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *ManaV2BgpNeighbor) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *ManaV2BgpNeighbor) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *ManaV2BgpNeighbor) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *ManaV2BgpNeighbor) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *ManaV2BgpNeighbor) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetRemovePrivateAs

`func (o *ManaV2BgpNeighbor) GetRemovePrivateAs() bool`

GetRemovePrivateAs returns the RemovePrivateAs field if non-nil, zero value otherwise.

### GetRemovePrivateAsOk

`func (o *ManaV2BgpNeighbor) GetRemovePrivateAsOk() (*bool, bool)`

GetRemovePrivateAsOk returns a tuple with the RemovePrivateAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePrivateAs

`func (o *ManaV2BgpNeighbor) SetRemovePrivateAs(v bool)`

SetRemovePrivateAs sets RemovePrivateAs field to given value.

### HasRemovePrivateAs

`func (o *ManaV2BgpNeighbor) HasRemovePrivateAs() bool`

HasRemovePrivateAs returns a boolean if a field has been set.

### GetSendCommunity

`func (o *ManaV2BgpNeighbor) GetSendCommunity() bool`

GetSendCommunity returns the SendCommunity field if non-nil, zero value otherwise.

### GetSendCommunityOk

`func (o *ManaV2BgpNeighbor) GetSendCommunityOk() (*bool, bool)`

GetSendCommunityOk returns a tuple with the SendCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCommunity

`func (o *ManaV2BgpNeighbor) SetSendCommunity(v bool)`

SetSendCommunity sets SendCommunity field to given value.

### HasSendCommunity

`func (o *ManaV2BgpNeighbor) HasSendCommunity() bool`

HasSendCommunity returns a boolean if a field has been set.

### GetState

`func (o *ManaV2BgpNeighbor) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManaV2BgpNeighbor) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManaV2BgpNeighbor) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ManaV2BgpNeighbor) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeSinceLastOperChange

`func (o *ManaV2BgpNeighbor) GetTimeSinceLastOperChange() GoogleProtobufDuration`

GetTimeSinceLastOperChange returns the TimeSinceLastOperChange field if non-nil, zero value otherwise.

### GetTimeSinceLastOperChangeOk

`func (o *ManaV2BgpNeighbor) GetTimeSinceLastOperChangeOk() (*GoogleProtobufDuration, bool)`

GetTimeSinceLastOperChangeOk returns a tuple with the TimeSinceLastOperChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSinceLastOperChange

`func (o *ManaV2BgpNeighbor) SetTimeSinceLastOperChange(v GoogleProtobufDuration)`

SetTimeSinceLastOperChange sets TimeSinceLastOperChange field to given value.

### HasTimeSinceLastOperChange

`func (o *ManaV2BgpNeighbor) HasTimeSinceLastOperChange() bool`

HasTimeSinceLastOperChange returns a boolean if a field has been set.

### GetUp

`func (o *ManaV2BgpNeighbor) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *ManaV2BgpNeighbor) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *ManaV2BgpNeighbor) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *ManaV2BgpNeighbor) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


