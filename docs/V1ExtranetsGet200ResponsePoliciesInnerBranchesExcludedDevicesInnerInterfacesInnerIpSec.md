# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiReplayWindowSize** | Pointer to **int32** |  | [optional] 
**DhGroup** | Pointer to **string** |  | [optional] 
**DpdInterval** | Pointer to **int32** |  | [optional] 
**EncryptionAlg** | Pointer to **string** |  | [optional] 
**Esn** | Pointer to **bool** |  | [optional] 
**EstablishedTime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**IkeIntegrity** | Pointer to **string** |  | [optional] 
**IpsecEncryptionAlg** | Pointer to **string** |  | [optional] 
**IpsecIntegrity** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LocalAddress** | Pointer to **string** |  | [optional] 
**LocalCircuit** | Pointer to **string** |  | [optional] 
**LocalIkePeerIdentity** | Pointer to **string** |  | [optional] 
**LocalIkesaSpi** | Pointer to **int64** |  | [optional] 
**LocalPort** | Pointer to **int32** |  | [optional] 
**NegotiatedAlgo** | Pointer to **string** |  | [optional] 
**OperState** | Pointer to **bool** |  | [optional] 
**PerfectForwardSecrecy** | Pointer to **string** |  | [optional] 
**PresharedKey** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**ReauthInterval** | Pointer to **int64** |  | [optional] 
**RekeyInterval** | Pointer to **int64** |  | [optional] 
**RemoteAddress** | Pointer to **string** |  | [optional] 
**RemoteIkePeerIdentity** | Pointer to **string** |  | [optional] 
**RemoteIkesaSpi** | Pointer to **int64** |  | [optional] 
**RemotePort** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSecWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSecWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSecWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiReplayWindowSize

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetAntiReplayWindowSize() int32`

GetAntiReplayWindowSize returns the AntiReplayWindowSize field if non-nil, zero value otherwise.

### GetAntiReplayWindowSizeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetAntiReplayWindowSizeOk() (*int32, bool)`

GetAntiReplayWindowSizeOk returns a tuple with the AntiReplayWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiReplayWindowSize

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetAntiReplayWindowSize(v int32)`

SetAntiReplayWindowSize sets AntiReplayWindowSize field to given value.

### HasAntiReplayWindowSize

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasAntiReplayWindowSize() bool`

HasAntiReplayWindowSize returns a boolean if a field has been set.

### GetDhGroup

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetDhGroup() string`

GetDhGroup returns the DhGroup field if non-nil, zero value otherwise.

### GetDhGroupOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetDhGroupOk() (*string, bool)`

GetDhGroupOk returns a tuple with the DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhGroup

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetDhGroup(v string)`

SetDhGroup sets DhGroup field to given value.

### HasDhGroup

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasDhGroup() bool`

HasDhGroup returns a boolean if a field has been set.

### GetDpdInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetDpdInterval() int32`

GetDpdInterval returns the DpdInterval field if non-nil, zero value otherwise.

### GetDpdIntervalOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetDpdIntervalOk() (*int32, bool)`

GetDpdIntervalOk returns a tuple with the DpdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetDpdInterval(v int32)`

SetDpdInterval sets DpdInterval field to given value.

### HasDpdInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasDpdInterval() bool`

HasDpdInterval returns a boolean if a field has been set.

### GetEncryptionAlg

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetEncryptionAlg() string`

GetEncryptionAlg returns the EncryptionAlg field if non-nil, zero value otherwise.

### GetEncryptionAlgOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetEncryptionAlgOk() (*string, bool)`

GetEncryptionAlgOk returns a tuple with the EncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlg

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetEncryptionAlg(v string)`

SetEncryptionAlg sets EncryptionAlg field to given value.

### HasEncryptionAlg

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasEncryptionAlg() bool`

HasEncryptionAlg returns a boolean if a field has been set.

### GetEsn

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetEsn() bool`

GetEsn returns the Esn field if non-nil, zero value otherwise.

### GetEsnOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetEsnOk() (*bool, bool)`

GetEsnOk returns a tuple with the Esn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsn

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetEsn(v bool)`

SetEsn sets Esn field to given value.

### HasEsn

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasEsn() bool`

HasEsn returns a boolean if a field has been set.

### GetEstablishedTime

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetEstablishedTime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetEstablishedTime returns the EstablishedTime field if non-nil, zero value otherwise.

### GetEstablishedTimeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetEstablishedTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetEstablishedTimeOk returns a tuple with the EstablishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablishedTime

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetEstablishedTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetEstablishedTime sets EstablishedTime field to given value.

### HasEstablishedTime

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasEstablishedTime() bool`

HasEstablishedTime returns a boolean if a field has been set.

### GetIkeIntegrity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetIkeIntegrity() string`

GetIkeIntegrity returns the IkeIntegrity field if non-nil, zero value otherwise.

### GetIkeIntegrityOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetIkeIntegrityOk() (*string, bool)`

GetIkeIntegrityOk returns a tuple with the IkeIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeIntegrity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetIkeIntegrity(v string)`

SetIkeIntegrity sets IkeIntegrity field to given value.

### HasIkeIntegrity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasIkeIntegrity() bool`

HasIkeIntegrity returns a boolean if a field has been set.

### GetIpsecEncryptionAlg

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetIpsecEncryptionAlg() string`

GetIpsecEncryptionAlg returns the IpsecEncryptionAlg field if non-nil, zero value otherwise.

### GetIpsecEncryptionAlgOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetIpsecEncryptionAlgOk() (*string, bool)`

GetIpsecEncryptionAlgOk returns a tuple with the IpsecEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecEncryptionAlg

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetIpsecEncryptionAlg(v string)`

SetIpsecEncryptionAlg sets IpsecEncryptionAlg field to given value.

### HasIpsecEncryptionAlg

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasIpsecEncryptionAlg() bool`

HasIpsecEncryptionAlg returns a boolean if a field has been set.

### GetIpsecIntegrity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetIpsecIntegrity() string`

GetIpsecIntegrity returns the IpsecIntegrity field if non-nil, zero value otherwise.

### GetIpsecIntegrityOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetIpsecIntegrityOk() (*string, bool)`

GetIpsecIntegrityOk returns a tuple with the IpsecIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecIntegrity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetIpsecIntegrity(v string)`

SetIpsecIntegrity sets IpsecIntegrity field to given value.

### HasIpsecIntegrity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasIpsecIntegrity() bool`

HasIpsecIntegrity returns a boolean if a field has been set.

### GetLabel

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLocalAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetLocalCircuit

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLocalCircuit() string`

GetLocalCircuit returns the LocalCircuit field if non-nil, zero value otherwise.

### GetLocalCircuitOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLocalCircuitOk() (*string, bool)`

GetLocalCircuitOk returns a tuple with the LocalCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCircuit

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetLocalCircuit(v string)`

SetLocalCircuit sets LocalCircuit field to given value.

### HasLocalCircuit

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasLocalCircuit() bool`

HasLocalCircuit returns a boolean if a field has been set.

### GetLocalIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLocalIkePeerIdentity() string`

GetLocalIkePeerIdentity returns the LocalIkePeerIdentity field if non-nil, zero value otherwise.

### GetLocalIkePeerIdentityOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLocalIkePeerIdentityOk() (*string, bool)`

GetLocalIkePeerIdentityOk returns a tuple with the LocalIkePeerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetLocalIkePeerIdentity(v string)`

SetLocalIkePeerIdentity sets LocalIkePeerIdentity field to given value.

### HasLocalIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasLocalIkePeerIdentity() bool`

HasLocalIkePeerIdentity returns a boolean if a field has been set.

### GetLocalIkesaSpi

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLocalIkesaSpi() int64`

GetLocalIkesaSpi returns the LocalIkesaSpi field if non-nil, zero value otherwise.

### GetLocalIkesaSpiOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLocalIkesaSpiOk() (*int64, bool)`

GetLocalIkesaSpiOk returns a tuple with the LocalIkesaSpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIkesaSpi

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetLocalIkesaSpi(v int64)`

SetLocalIkesaSpi sets LocalIkesaSpi field to given value.

### HasLocalIkesaSpi

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasLocalIkesaSpi() bool`

HasLocalIkesaSpi returns a boolean if a field has been set.

### GetLocalPort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### GetNegotiatedAlgo

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetNegotiatedAlgo() string`

GetNegotiatedAlgo returns the NegotiatedAlgo field if non-nil, zero value otherwise.

### GetNegotiatedAlgoOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetNegotiatedAlgoOk() (*string, bool)`

GetNegotiatedAlgoOk returns a tuple with the NegotiatedAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiatedAlgo

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetNegotiatedAlgo(v string)`

SetNegotiatedAlgo sets NegotiatedAlgo field to given value.

### HasNegotiatedAlgo

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasNegotiatedAlgo() bool`

HasNegotiatedAlgo returns a boolean if a field has been set.

### GetOperState

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetOperState() bool`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetOperStateOk() (*bool, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetOperState(v bool)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPerfectForwardSecrecy

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetPerfectForwardSecrecy() string`

GetPerfectForwardSecrecy returns the PerfectForwardSecrecy field if non-nil, zero value otherwise.

### GetPerfectForwardSecrecyOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetPerfectForwardSecrecyOk() (*string, bool)`

GetPerfectForwardSecrecyOk returns a tuple with the PerfectForwardSecrecy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectForwardSecrecy

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetPerfectForwardSecrecy(v string)`

SetPerfectForwardSecrecy sets PerfectForwardSecrecy field to given value.

### HasPerfectForwardSecrecy

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasPerfectForwardSecrecy() bool`

HasPerfectForwardSecrecy returns a boolean if a field has been set.

### GetPresharedKey

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetReauthInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetReauthInterval() int64`

GetReauthInterval returns the ReauthInterval field if non-nil, zero value otherwise.

### GetReauthIntervalOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetReauthIntervalOk() (*int64, bool)`

GetReauthIntervalOk returns a tuple with the ReauthInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetReauthInterval(v int64)`

SetReauthInterval sets ReauthInterval field to given value.

### HasReauthInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasReauthInterval() bool`

HasReauthInterval returns a boolean if a field has been set.

### GetRekeyInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetRekeyInterval() int64`

GetRekeyInterval returns the RekeyInterval field if non-nil, zero value otherwise.

### GetRekeyIntervalOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetRekeyIntervalOk() (*int64, bool)`

GetRekeyIntervalOk returns a tuple with the RekeyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetRekeyInterval(v int64)`

SetRekeyInterval sets RekeyInterval field to given value.

### HasRekeyInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasRekeyInterval() bool`

HasRekeyInterval returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetRemoteIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetRemoteIkePeerIdentity() string`

GetRemoteIkePeerIdentity returns the RemoteIkePeerIdentity field if non-nil, zero value otherwise.

### GetRemoteIkePeerIdentityOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetRemoteIkePeerIdentityOk() (*string, bool)`

GetRemoteIkePeerIdentityOk returns a tuple with the RemoteIkePeerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetRemoteIkePeerIdentity(v string)`

SetRemoteIkePeerIdentity sets RemoteIkePeerIdentity field to given value.

### HasRemoteIkePeerIdentity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasRemoteIkePeerIdentity() bool`

HasRemoteIkePeerIdentity returns a boolean if a field has been set.

### GetRemoteIkesaSpi

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetRemoteIkesaSpi() int64`

GetRemoteIkesaSpi returns the RemoteIkesaSpi field if non-nil, zero value otherwise.

### GetRemoteIkesaSpiOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetRemoteIkesaSpiOk() (*int64, bool)`

GetRemoteIkesaSpiOk returns a tuple with the RemoteIkesaSpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIkesaSpi

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetRemoteIkesaSpi(v int64)`

SetRemoteIkesaSpi sets RemoteIkesaSpi field to given value.

### HasRemoteIkesaSpi

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasRemoteIkesaSpi() bool`

HasRemoteIkesaSpi returns a boolean if a field has been set.

### GetRemotePort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpSec) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


