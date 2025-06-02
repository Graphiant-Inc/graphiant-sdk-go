# V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiReplayWSize** | Pointer to **int32** |  | [optional] 
**EstablishedTime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**LocalCircuit** | Pointer to **string** |  | [optional] 
**LocalInterface** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner.md) |  | [optional] 
**LocalPort** | Pointer to **int32** |  | [optional] 
**LocalSpi** | Pointer to **int64** |  | [optional] 
**NegotiatedAlgorithms** | Pointer to **string** |  | [optional] 
**OperState** | Pointer to **bool** |  | [optional] 
**PeerAddress** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**RekeyTime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**RemotePort** | Pointer to **int32** |  | [optional] 
**RemoteSpi** | Pointer to **int64** |  | [optional] 
**SessionId** | Pointer to **int64** |  | [optional] 
**SourceAddress** | Pointer to **string** |  | [optional] 
**TunnelInterface** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection

`func NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection() *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection`

NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection instantiates a new V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnectionWithDefaults

`func NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnectionWithDefaults() *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection`

NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnectionWithDefaults instantiates a new V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiReplayWSize

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetAntiReplayWSize() int32`

GetAntiReplayWSize returns the AntiReplayWSize field if non-nil, zero value otherwise.

### GetAntiReplayWSizeOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetAntiReplayWSizeOk() (*int32, bool)`

GetAntiReplayWSizeOk returns a tuple with the AntiReplayWSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiReplayWSize

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetAntiReplayWSize(v int32)`

SetAntiReplayWSize sets AntiReplayWSize field to given value.

### HasAntiReplayWSize

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasAntiReplayWSize() bool`

HasAntiReplayWSize returns a boolean if a field has been set.

### GetEstablishedTime

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetEstablishedTime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetEstablishedTime returns the EstablishedTime field if non-nil, zero value otherwise.

### GetEstablishedTimeOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetEstablishedTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetEstablishedTimeOk returns a tuple with the EstablishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablishedTime

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetEstablishedTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetEstablishedTime sets EstablishedTime field to given value.

### HasEstablishedTime

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasEstablishedTime() bool`

HasEstablishedTime returns a boolean if a field has been set.

### GetLocalCircuit

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetLocalCircuit() string`

GetLocalCircuit returns the LocalCircuit field if non-nil, zero value otherwise.

### GetLocalCircuitOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetLocalCircuitOk() (*string, bool)`

GetLocalCircuitOk returns a tuple with the LocalCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCircuit

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetLocalCircuit(v string)`

SetLocalCircuit sets LocalCircuit field to given value.

### HasLocalCircuit

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasLocalCircuit() bool`

HasLocalCircuit returns a boolean if a field has been set.

### GetLocalInterface

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetLocalInterface() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner`

GetLocalInterface returns the LocalInterface field if non-nil, zero value otherwise.

### GetLocalInterfaceOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetLocalInterfaceOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner, bool)`

GetLocalInterfaceOk returns a tuple with the LocalInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInterface

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetLocalInterface(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner)`

SetLocalInterface sets LocalInterface field to given value.

### HasLocalInterface

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasLocalInterface() bool`

HasLocalInterface returns a boolean if a field has been set.

### GetLocalPort

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### GetLocalSpi

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetLocalSpi() int64`

GetLocalSpi returns the LocalSpi field if non-nil, zero value otherwise.

### GetLocalSpiOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetLocalSpiOk() (*int64, bool)`

GetLocalSpiOk returns a tuple with the LocalSpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSpi

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetLocalSpi(v int64)`

SetLocalSpi sets LocalSpi field to given value.

### HasLocalSpi

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasLocalSpi() bool`

HasLocalSpi returns a boolean if a field has been set.

### GetNegotiatedAlgorithms

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetNegotiatedAlgorithms() string`

GetNegotiatedAlgorithms returns the NegotiatedAlgorithms field if non-nil, zero value otherwise.

### GetNegotiatedAlgorithmsOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetNegotiatedAlgorithmsOk() (*string, bool)`

GetNegotiatedAlgorithmsOk returns a tuple with the NegotiatedAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiatedAlgorithms

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetNegotiatedAlgorithms(v string)`

SetNegotiatedAlgorithms sets NegotiatedAlgorithms field to given value.

### HasNegotiatedAlgorithms

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasNegotiatedAlgorithms() bool`

HasNegotiatedAlgorithms returns a boolean if a field has been set.

### GetOperState

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetOperState() bool`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetOperStateOk() (*bool, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetOperState(v bool)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPeerAddress

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetPeerAddress() string`

GetPeerAddress returns the PeerAddress field if non-nil, zero value otherwise.

### GetPeerAddressOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetPeerAddressOk() (*string, bool)`

GetPeerAddressOk returns a tuple with the PeerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAddress

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetPeerAddress(v string)`

SetPeerAddress sets PeerAddress field to given value.

### HasPeerAddress

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasPeerAddress() bool`

HasPeerAddress returns a boolean if a field has been set.

### GetProtocol

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRekeyTime

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetRekeyTime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetRekeyTime returns the RekeyTime field if non-nil, zero value otherwise.

### GetRekeyTimeOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetRekeyTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetRekeyTimeOk returns a tuple with the RekeyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyTime

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetRekeyTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetRekeyTime sets RekeyTime field to given value.

### HasRekeyTime

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasRekeyTime() bool`

HasRekeyTime returns a boolean if a field has been set.

### GetRemotePort

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetRemoteSpi

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetRemoteSpi() int64`

GetRemoteSpi returns the RemoteSpi field if non-nil, zero value otherwise.

### GetRemoteSpiOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetRemoteSpiOk() (*int64, bool)`

GetRemoteSpiOk returns a tuple with the RemoteSpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSpi

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetRemoteSpi(v int64)`

SetRemoteSpi sets RemoteSpi field to given value.

### HasRemoteSpi

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasRemoteSpi() bool`

HasRemoteSpi returns a boolean if a field has been set.

### GetSessionId

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetSessionId() int64`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetSessionIdOk() (*int64, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetSessionId(v int64)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSourceAddress

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### GetTunnelInterface

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetTunnelInterface() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner`

GetTunnelInterface returns the TunnelInterface field if non-nil, zero value otherwise.

### GetTunnelInterfaceOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) GetTunnelInterfaceOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner, bool)`

GetTunnelInterfaceOk returns a tuple with the TunnelInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelInterface

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) SetTunnelInterface(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner)`

SetTunnelInterface sets TunnelInterface field to given value.

### HasTunnelInterface

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection) HasTunnelInterface() bool`

HasTunnelInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


