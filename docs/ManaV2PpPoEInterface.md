# ManaV2PpPoEInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to **string** |  | [optional] 
**DnsServers** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LastIpcpCloseReason** | Pointer to **string** |  | [optional] 
**LastIpv6CpCloseReason** | Pointer to **string** |  | [optional] 
**LastLcpCloseReason** | Pointer to **string** |  | [optional] 
**LcpEchoFailure** | Pointer to **int32** |  | [optional] 
**LcpEchoInterval** | Pointer to **int32** |  | [optional] 
**MaximumReceiveUnit** | Pointer to **int32** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PeerIpv4Address** | Pointer to **string** |  | [optional] 
**PeerIpv6Address** | Pointer to **string** |  | [optional] 
**PppConnectTimeout** | Pointer to **int32** |  | [optional] 
**ServerMacAddress** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewManaV2PpPoEInterface

`func NewManaV2PpPoEInterface() *ManaV2PpPoEInterface`

NewManaV2PpPoEInterface instantiates a new ManaV2PpPoEInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PpPoEInterfaceWithDefaults

`func NewManaV2PpPoEInterfaceWithDefaults() *ManaV2PpPoEInterface`

NewManaV2PpPoEInterfaceWithDefaults instantiates a new ManaV2PpPoEInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *ManaV2PpPoEInterface) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *ManaV2PpPoEInterface) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *ManaV2PpPoEInterface) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *ManaV2PpPoEInterface) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetDnsServers

`func (o *ManaV2PpPoEInterface) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *ManaV2PpPoEInterface) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *ManaV2PpPoEInterface) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *ManaV2PpPoEInterface) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetId

`func (o *ManaV2PpPoEInterface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2PpPoEInterface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2PpPoEInterface) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ManaV2PpPoEInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastIpcpCloseReason

`func (o *ManaV2PpPoEInterface) GetLastIpcpCloseReason() string`

GetLastIpcpCloseReason returns the LastIpcpCloseReason field if non-nil, zero value otherwise.

### GetLastIpcpCloseReasonOk

`func (o *ManaV2PpPoEInterface) GetLastIpcpCloseReasonOk() (*string, bool)`

GetLastIpcpCloseReasonOk returns a tuple with the LastIpcpCloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIpcpCloseReason

`func (o *ManaV2PpPoEInterface) SetLastIpcpCloseReason(v string)`

SetLastIpcpCloseReason sets LastIpcpCloseReason field to given value.

### HasLastIpcpCloseReason

`func (o *ManaV2PpPoEInterface) HasLastIpcpCloseReason() bool`

HasLastIpcpCloseReason returns a boolean if a field has been set.

### GetLastIpv6CpCloseReason

`func (o *ManaV2PpPoEInterface) GetLastIpv6CpCloseReason() string`

GetLastIpv6CpCloseReason returns the LastIpv6CpCloseReason field if non-nil, zero value otherwise.

### GetLastIpv6CpCloseReasonOk

`func (o *ManaV2PpPoEInterface) GetLastIpv6CpCloseReasonOk() (*string, bool)`

GetLastIpv6CpCloseReasonOk returns a tuple with the LastIpv6CpCloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIpv6CpCloseReason

`func (o *ManaV2PpPoEInterface) SetLastIpv6CpCloseReason(v string)`

SetLastIpv6CpCloseReason sets LastIpv6CpCloseReason field to given value.

### HasLastIpv6CpCloseReason

`func (o *ManaV2PpPoEInterface) HasLastIpv6CpCloseReason() bool`

HasLastIpv6CpCloseReason returns a boolean if a field has been set.

### GetLastLcpCloseReason

`func (o *ManaV2PpPoEInterface) GetLastLcpCloseReason() string`

GetLastLcpCloseReason returns the LastLcpCloseReason field if non-nil, zero value otherwise.

### GetLastLcpCloseReasonOk

`func (o *ManaV2PpPoEInterface) GetLastLcpCloseReasonOk() (*string, bool)`

GetLastLcpCloseReasonOk returns a tuple with the LastLcpCloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLcpCloseReason

`func (o *ManaV2PpPoEInterface) SetLastLcpCloseReason(v string)`

SetLastLcpCloseReason sets LastLcpCloseReason field to given value.

### HasLastLcpCloseReason

`func (o *ManaV2PpPoEInterface) HasLastLcpCloseReason() bool`

HasLastLcpCloseReason returns a boolean if a field has been set.

### GetLcpEchoFailure

`func (o *ManaV2PpPoEInterface) GetLcpEchoFailure() int32`

GetLcpEchoFailure returns the LcpEchoFailure field if non-nil, zero value otherwise.

### GetLcpEchoFailureOk

`func (o *ManaV2PpPoEInterface) GetLcpEchoFailureOk() (*int32, bool)`

GetLcpEchoFailureOk returns a tuple with the LcpEchoFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcpEchoFailure

`func (o *ManaV2PpPoEInterface) SetLcpEchoFailure(v int32)`

SetLcpEchoFailure sets LcpEchoFailure field to given value.

### HasLcpEchoFailure

`func (o *ManaV2PpPoEInterface) HasLcpEchoFailure() bool`

HasLcpEchoFailure returns a boolean if a field has been set.

### GetLcpEchoInterval

`func (o *ManaV2PpPoEInterface) GetLcpEchoInterval() int32`

GetLcpEchoInterval returns the LcpEchoInterval field if non-nil, zero value otherwise.

### GetLcpEchoIntervalOk

`func (o *ManaV2PpPoEInterface) GetLcpEchoIntervalOk() (*int32, bool)`

GetLcpEchoIntervalOk returns a tuple with the LcpEchoInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcpEchoInterval

`func (o *ManaV2PpPoEInterface) SetLcpEchoInterval(v int32)`

SetLcpEchoInterval sets LcpEchoInterval field to given value.

### HasLcpEchoInterval

`func (o *ManaV2PpPoEInterface) HasLcpEchoInterval() bool`

HasLcpEchoInterval returns a boolean if a field has been set.

### GetMaximumReceiveUnit

`func (o *ManaV2PpPoEInterface) GetMaximumReceiveUnit() int32`

GetMaximumReceiveUnit returns the MaximumReceiveUnit field if non-nil, zero value otherwise.

### GetMaximumReceiveUnitOk

`func (o *ManaV2PpPoEInterface) GetMaximumReceiveUnitOk() (*int32, bool)`

GetMaximumReceiveUnitOk returns a tuple with the MaximumReceiveUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumReceiveUnit

`func (o *ManaV2PpPoEInterface) SetMaximumReceiveUnit(v int32)`

SetMaximumReceiveUnit sets MaximumReceiveUnit field to given value.

### HasMaximumReceiveUnit

`func (o *ManaV2PpPoEInterface) HasMaximumReceiveUnit() bool`

HasMaximumReceiveUnit returns a boolean if a field has been set.

### GetPassword

`func (o *ManaV2PpPoEInterface) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ManaV2PpPoEInterface) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ManaV2PpPoEInterface) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ManaV2PpPoEInterface) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPeerIpv4Address

`func (o *ManaV2PpPoEInterface) GetPeerIpv4Address() string`

GetPeerIpv4Address returns the PeerIpv4Address field if non-nil, zero value otherwise.

### GetPeerIpv4AddressOk

`func (o *ManaV2PpPoEInterface) GetPeerIpv4AddressOk() (*string, bool)`

GetPeerIpv4AddressOk returns a tuple with the PeerIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIpv4Address

`func (o *ManaV2PpPoEInterface) SetPeerIpv4Address(v string)`

SetPeerIpv4Address sets PeerIpv4Address field to given value.

### HasPeerIpv4Address

`func (o *ManaV2PpPoEInterface) HasPeerIpv4Address() bool`

HasPeerIpv4Address returns a boolean if a field has been set.

### GetPeerIpv6Address

`func (o *ManaV2PpPoEInterface) GetPeerIpv6Address() string`

GetPeerIpv6Address returns the PeerIpv6Address field if non-nil, zero value otherwise.

### GetPeerIpv6AddressOk

`func (o *ManaV2PpPoEInterface) GetPeerIpv6AddressOk() (*string, bool)`

GetPeerIpv6AddressOk returns a tuple with the PeerIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIpv6Address

`func (o *ManaV2PpPoEInterface) SetPeerIpv6Address(v string)`

SetPeerIpv6Address sets PeerIpv6Address field to given value.

### HasPeerIpv6Address

`func (o *ManaV2PpPoEInterface) HasPeerIpv6Address() bool`

HasPeerIpv6Address returns a boolean if a field has been set.

### GetPppConnectTimeout

`func (o *ManaV2PpPoEInterface) GetPppConnectTimeout() int32`

GetPppConnectTimeout returns the PppConnectTimeout field if non-nil, zero value otherwise.

### GetPppConnectTimeoutOk

`func (o *ManaV2PpPoEInterface) GetPppConnectTimeoutOk() (*int32, bool)`

GetPppConnectTimeoutOk returns a tuple with the PppConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPppConnectTimeout

`func (o *ManaV2PpPoEInterface) SetPppConnectTimeout(v int32)`

SetPppConnectTimeout sets PppConnectTimeout field to given value.

### HasPppConnectTimeout

`func (o *ManaV2PpPoEInterface) HasPppConnectTimeout() bool`

HasPppConnectTimeout returns a boolean if a field has been set.

### GetServerMacAddress

`func (o *ManaV2PpPoEInterface) GetServerMacAddress() string`

GetServerMacAddress returns the ServerMacAddress field if non-nil, zero value otherwise.

### GetServerMacAddressOk

`func (o *ManaV2PpPoEInterface) GetServerMacAddressOk() (*string, bool)`

GetServerMacAddressOk returns a tuple with the ServerMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMacAddress

`func (o *ManaV2PpPoEInterface) SetServerMacAddress(v string)`

SetServerMacAddress sets ServerMacAddress field to given value.

### HasServerMacAddress

`func (o *ManaV2PpPoEInterface) HasServerMacAddress() bool`

HasServerMacAddress returns a boolean if a field has been set.

### GetServerName

`func (o *ManaV2PpPoEInterface) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ManaV2PpPoEInterface) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ManaV2PpPoEInterface) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ManaV2PpPoEInterface) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetSessionId

`func (o *ManaV2PpPoEInterface) GetSessionId() int32`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ManaV2PpPoEInterface) GetSessionIdOk() (*int32, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ManaV2PpPoEInterface) SetSessionId(v int32)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ManaV2PpPoEInterface) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetState

`func (o *ManaV2PpPoEInterface) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManaV2PpPoEInterface) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManaV2PpPoEInterface) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ManaV2PpPoEInterface) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUsername

`func (o *ManaV2PpPoEInterface) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ManaV2PpPoEInterface) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ManaV2PpPoEInterface) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ManaV2PpPoEInterface) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


