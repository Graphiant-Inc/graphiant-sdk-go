# V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationAddress** | Pointer to **string** |  | [optional] 
**IkeInitiator** | Pointer to **bool** |  | [optional] 
**IpsecLabel** | Pointer to **string** |  | [optional] 
**Lan** | Pointer to **string** |  | [optional] 
**LocalAddressV4** | Pointer to **string** |  | [optional] 
**LocalAddressV6** | Pointer to **string** |  | [optional] 
**LocalCircuit** | Pointer to **string** |  | [optional] 
**LocalIkePeerIdentity** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PresharedKey** | Pointer to **string** |  | [optional] 
**RemoteAddressV4** | Pointer to **string** |  | [optional] 
**RemoteAddressV6** | Pointer to **string** |  | [optional] 
**RemoteIkePeerIdentity** | Pointer to **string** |  | [optional] 
**Routing** | Pointer to [**V1GatewaysPutRequestDetailsIpsecGatewayRouting**](V1GatewaysPutRequestDetailsIpsecGatewayRouting.md) |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**VpnProfile** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn

`func NewV1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn() *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn`

NewV1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn instantiates a new V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpnWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpnWithDefaults() *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn`

NewV1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpnWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationAddress

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetIkeInitiator

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetIkeInitiator() bool`

GetIkeInitiator returns the IkeInitiator field if non-nil, zero value otherwise.

### GetIkeInitiatorOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetIkeInitiatorOk() (*bool, bool)`

GetIkeInitiatorOk returns a tuple with the IkeInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeInitiator

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetIkeInitiator(v bool)`

SetIkeInitiator sets IkeInitiator field to given value.

### HasIkeInitiator

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasIkeInitiator() bool`

HasIkeInitiator returns a boolean if a field has been set.

### GetIpsecLabel

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetIpsecLabel() string`

GetIpsecLabel returns the IpsecLabel field if non-nil, zero value otherwise.

### GetIpsecLabelOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetIpsecLabelOk() (*string, bool)`

GetIpsecLabelOk returns a tuple with the IpsecLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecLabel

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetIpsecLabel(v string)`

SetIpsecLabel sets IpsecLabel field to given value.

### HasIpsecLabel

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasIpsecLabel() bool`

HasIpsecLabel returns a boolean if a field has been set.

### GetLan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLocalAddressV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetLocalAddressV4() string`

GetLocalAddressV4 returns the LocalAddressV4 field if non-nil, zero value otherwise.

### GetLocalAddressV4Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetLocalAddressV4Ok() (*string, bool)`

GetLocalAddressV4Ok returns a tuple with the LocalAddressV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddressV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetLocalAddressV4(v string)`

SetLocalAddressV4 sets LocalAddressV4 field to given value.

### HasLocalAddressV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasLocalAddressV4() bool`

HasLocalAddressV4 returns a boolean if a field has been set.

### GetLocalAddressV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetLocalAddressV6() string`

GetLocalAddressV6 returns the LocalAddressV6 field if non-nil, zero value otherwise.

### GetLocalAddressV6Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetLocalAddressV6Ok() (*string, bool)`

GetLocalAddressV6Ok returns a tuple with the LocalAddressV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddressV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetLocalAddressV6(v string)`

SetLocalAddressV6 sets LocalAddressV6 field to given value.

### HasLocalAddressV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasLocalAddressV6() bool`

HasLocalAddressV6 returns a boolean if a field has been set.

### GetLocalCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetLocalCircuit() string`

GetLocalCircuit returns the LocalCircuit field if non-nil, zero value otherwise.

### GetLocalCircuitOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetLocalCircuitOk() (*string, bool)`

GetLocalCircuitOk returns a tuple with the LocalCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetLocalCircuit(v string)`

SetLocalCircuit sets LocalCircuit field to given value.

### HasLocalCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasLocalCircuit() bool`

HasLocalCircuit returns a boolean if a field has been set.

### GetLocalIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetLocalIkePeerIdentity() string`

GetLocalIkePeerIdentity returns the LocalIkePeerIdentity field if non-nil, zero value otherwise.

### GetLocalIkePeerIdentityOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetLocalIkePeerIdentityOk() (*string, bool)`

GetLocalIkePeerIdentityOk returns a tuple with the LocalIkePeerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetLocalIkePeerIdentity(v string)`

SetLocalIkePeerIdentity sets LocalIkePeerIdentity field to given value.

### HasLocalIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasLocalIkePeerIdentity() bool`

HasLocalIkePeerIdentity returns a boolean if a field has been set.

### GetMtu

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPresharedKey

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetRemoteAddressV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetRemoteAddressV4() string`

GetRemoteAddressV4 returns the RemoteAddressV4 field if non-nil, zero value otherwise.

### GetRemoteAddressV4Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetRemoteAddressV4Ok() (*string, bool)`

GetRemoteAddressV4Ok returns a tuple with the RemoteAddressV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddressV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetRemoteAddressV4(v string)`

SetRemoteAddressV4 sets RemoteAddressV4 field to given value.

### HasRemoteAddressV4

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasRemoteAddressV4() bool`

HasRemoteAddressV4 returns a boolean if a field has been set.

### GetRemoteAddressV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetRemoteAddressV6() string`

GetRemoteAddressV6 returns the RemoteAddressV6 field if non-nil, zero value otherwise.

### GetRemoteAddressV6Ok

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetRemoteAddressV6Ok() (*string, bool)`

GetRemoteAddressV6Ok returns a tuple with the RemoteAddressV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddressV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetRemoteAddressV6(v string)`

SetRemoteAddressV6 sets RemoteAddressV6 field to given value.

### HasRemoteAddressV6

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasRemoteAddressV6() bool`

HasRemoteAddressV6 returns a boolean if a field has been set.

### GetRemoteIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetRemoteIkePeerIdentity() string`

GetRemoteIkePeerIdentity returns the RemoteIkePeerIdentity field if non-nil, zero value otherwise.

### GetRemoteIkePeerIdentityOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetRemoteIkePeerIdentityOk() (*string, bool)`

GetRemoteIkePeerIdentityOk returns a tuple with the RemoteIkePeerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetRemoteIkePeerIdentity(v string)`

SetRemoteIkePeerIdentity sets RemoteIkePeerIdentity field to given value.

### HasRemoteIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasRemoteIkePeerIdentity() bool`

HasRemoteIkePeerIdentity returns a boolean if a field has been set.

### GetRouting

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetRouting() V1GatewaysPutRequestDetailsIpsecGatewayRouting`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetRoutingOk() (*V1GatewaysPutRequestDetailsIpsecGatewayRouting, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetRouting(v V1GatewaysPutRequestDetailsIpsecGatewayRouting)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetVpnProfile

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetVpnProfile() string`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) GetVpnProfileOk() (*string, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) SetVpnProfile(v string)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeSiteToSiteVpnValueSiteToSiteVpn) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


