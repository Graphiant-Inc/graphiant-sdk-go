# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bgp** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerBgp**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerBgp.md) |  | [optional] 
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
**Static** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerStatic**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerStatic.md) |  | [optional] 
**TcpMss** | Pointer to **int32** |  | [optional] 
**VpnProfile** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetBgp() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerBgp`

GetBgp returns the Bgp field if non-nil, zero value otherwise.

### GetBgpOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetBgpOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerBgp, bool)`

GetBgpOk returns a tuple with the Bgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetBgp(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerBgp)`

SetBgp sets Bgp field to given value.

### HasBgp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasBgp() bool`

HasBgp returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetIkeInitiator

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetIkeInitiator() bool`

GetIkeInitiator returns the IkeInitiator field if non-nil, zero value otherwise.

### GetIkeInitiatorOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetIkeInitiatorOk() (*bool, bool)`

GetIkeInitiatorOk returns a tuple with the IkeInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeInitiator

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetIkeInitiator(v bool)`

SetIkeInitiator sets IkeInitiator field to given value.

### HasIkeInitiator

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasIkeInitiator() bool`

HasIkeInitiator returns a boolean if a field has been set.

### GetIpsecLabel

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetIpsecLabel() string`

GetIpsecLabel returns the IpsecLabel field if non-nil, zero value otherwise.

### GetIpsecLabelOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetIpsecLabelOk() (*string, bool)`

GetIpsecLabelOk returns a tuple with the IpsecLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecLabel

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetIpsecLabel(v string)`

SetIpsecLabel sets IpsecLabel field to given value.

### HasIpsecLabel

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasIpsecLabel() bool`

HasIpsecLabel returns a boolean if a field has been set.

### GetLan

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetLan(v string)`

SetLan sets Lan field to given value.

### HasLan

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasLan() bool`

HasLan returns a boolean if a field has been set.

### GetLocalAddressV4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetLocalAddressV4() string`

GetLocalAddressV4 returns the LocalAddressV4 field if non-nil, zero value otherwise.

### GetLocalAddressV4Ok

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetLocalAddressV4Ok() (*string, bool)`

GetLocalAddressV4Ok returns a tuple with the LocalAddressV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddressV4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetLocalAddressV4(v string)`

SetLocalAddressV4 sets LocalAddressV4 field to given value.

### HasLocalAddressV4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasLocalAddressV4() bool`

HasLocalAddressV4 returns a boolean if a field has been set.

### GetLocalAddressV6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetLocalAddressV6() string`

GetLocalAddressV6 returns the LocalAddressV6 field if non-nil, zero value otherwise.

### GetLocalAddressV6Ok

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetLocalAddressV6Ok() (*string, bool)`

GetLocalAddressV6Ok returns a tuple with the LocalAddressV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddressV6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetLocalAddressV6(v string)`

SetLocalAddressV6 sets LocalAddressV6 field to given value.

### HasLocalAddressV6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasLocalAddressV6() bool`

HasLocalAddressV6 returns a boolean if a field has been set.

### GetLocalCircuit

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetLocalCircuit() string`

GetLocalCircuit returns the LocalCircuit field if non-nil, zero value otherwise.

### GetLocalCircuitOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetLocalCircuitOk() (*string, bool)`

GetLocalCircuitOk returns a tuple with the LocalCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCircuit

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetLocalCircuit(v string)`

SetLocalCircuit sets LocalCircuit field to given value.

### HasLocalCircuit

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasLocalCircuit() bool`

HasLocalCircuit returns a boolean if a field has been set.

### GetLocalIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetLocalIkePeerIdentity() string`

GetLocalIkePeerIdentity returns the LocalIkePeerIdentity field if non-nil, zero value otherwise.

### GetLocalIkePeerIdentityOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetLocalIkePeerIdentityOk() (*string, bool)`

GetLocalIkePeerIdentityOk returns a tuple with the LocalIkePeerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetLocalIkePeerIdentity(v string)`

SetLocalIkePeerIdentity sets LocalIkePeerIdentity field to given value.

### HasLocalIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasLocalIkePeerIdentity() bool`

HasLocalIkePeerIdentity returns a boolean if a field has been set.

### GetMtu

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPresharedKey

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetRemoteAddressV4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetRemoteAddressV4() string`

GetRemoteAddressV4 returns the RemoteAddressV4 field if non-nil, zero value otherwise.

### GetRemoteAddressV4Ok

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetRemoteAddressV4Ok() (*string, bool)`

GetRemoteAddressV4Ok returns a tuple with the RemoteAddressV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddressV4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetRemoteAddressV4(v string)`

SetRemoteAddressV4 sets RemoteAddressV4 field to given value.

### HasRemoteAddressV4

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasRemoteAddressV4() bool`

HasRemoteAddressV4 returns a boolean if a field has been set.

### GetRemoteAddressV6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetRemoteAddressV6() string`

GetRemoteAddressV6 returns the RemoteAddressV6 field if non-nil, zero value otherwise.

### GetRemoteAddressV6Ok

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetRemoteAddressV6Ok() (*string, bool)`

GetRemoteAddressV6Ok returns a tuple with the RemoteAddressV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddressV6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetRemoteAddressV6(v string)`

SetRemoteAddressV6 sets RemoteAddressV6 field to given value.

### HasRemoteAddressV6

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasRemoteAddressV6() bool`

HasRemoteAddressV6 returns a boolean if a field has been set.

### GetRemoteIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetRemoteIkePeerIdentity() string`

GetRemoteIkePeerIdentity returns the RemoteIkePeerIdentity field if non-nil, zero value otherwise.

### GetRemoteIkePeerIdentityOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetRemoteIkePeerIdentityOk() (*string, bool)`

GetRemoteIkePeerIdentityOk returns a tuple with the RemoteIkePeerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetRemoteIkePeerIdentity(v string)`

SetRemoteIkePeerIdentity sets RemoteIkePeerIdentity field to given value.

### HasRemoteIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasRemoteIkePeerIdentity() bool`

HasRemoteIkePeerIdentity returns a boolean if a field has been set.

### GetStatic

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetStatic() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerStatic`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetStaticOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerStatic, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetStatic(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInnerStatic)`

SetStatic sets Static field to given value.

### HasStatic

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasStatic() bool`

HasStatic returns a boolean if a field has been set.

### GetTcpMss

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetTcpMss() int32`

GetTcpMss returns the TcpMss field if non-nil, zero value otherwise.

### GetTcpMssOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetTcpMssOk() (*int32, bool)`

GetTcpMssOk returns a tuple with the TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMss

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetTcpMss(v int32)`

SetTcpMss sets TcpMss field to given value.

### HasTcpMss

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasTcpMss() bool`

HasTcpMss returns a boolean if a field has been set.

### GetVpnProfile

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetVpnProfile() string`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) GetVpnProfileOk() (*string, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) SetVpnProfile(v string)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpsecTunnelsInner) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


