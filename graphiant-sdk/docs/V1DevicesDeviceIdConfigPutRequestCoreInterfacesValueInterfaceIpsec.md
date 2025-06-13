# V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiReplayWindowSize** | Pointer to **int32** |  | [optional] 
**DhGroup** | Pointer to **string** |  | [optional] 
**DpdInterval** | Pointer to **int32** |  | [optional] 
**EncryptionAlg** | Pointer to **string** |  | [optional] 
**Esn** | Pointer to **bool** |  | [optional] 
**IkeIntegrity** | Pointer to **string** |  | [optional] 
**Initiator** | Pointer to **bool** |  | [optional] 
**IpsecEncryptionAlg** | Pointer to **string** |  | [optional] 
**IpsecIntegrity** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LocalAddress** | Pointer to **string** |  | [optional] 
**LocalCircuit** | Pointer to **string** |  | [optional] 
**LocalIkePeerIdentity** | Pointer to **string** |  | [optional] 
**PerfectForwardSecrecy** | Pointer to **string** |  | [optional] 
**PresharedKey** | Pointer to **string** |  | [optional] 
**ReauthInterval** | Pointer to **int32** |  | [optional] 
**RekeyInterval** | Pointer to **int32** |  | [optional] 
**RemoteAddress** | Pointer to **string** |  | [optional] 
**RemoteIkePeerIdentity** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsecWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsecWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsecWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiReplayWindowSize

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetAntiReplayWindowSize() int32`

GetAntiReplayWindowSize returns the AntiReplayWindowSize field if non-nil, zero value otherwise.

### GetAntiReplayWindowSizeOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetAntiReplayWindowSizeOk() (*int32, bool)`

GetAntiReplayWindowSizeOk returns a tuple with the AntiReplayWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiReplayWindowSize

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetAntiReplayWindowSize(v int32)`

SetAntiReplayWindowSize sets AntiReplayWindowSize field to given value.

### HasAntiReplayWindowSize

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasAntiReplayWindowSize() bool`

HasAntiReplayWindowSize returns a boolean if a field has been set.

### GetDhGroup

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetDhGroup() string`

GetDhGroup returns the DhGroup field if non-nil, zero value otherwise.

### GetDhGroupOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetDhGroupOk() (*string, bool)`

GetDhGroupOk returns a tuple with the DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhGroup

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetDhGroup(v string)`

SetDhGroup sets DhGroup field to given value.

### HasDhGroup

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasDhGroup() bool`

HasDhGroup returns a boolean if a field has been set.

### GetDpdInterval

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetDpdInterval() int32`

GetDpdInterval returns the DpdInterval field if non-nil, zero value otherwise.

### GetDpdIntervalOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetDpdIntervalOk() (*int32, bool)`

GetDpdIntervalOk returns a tuple with the DpdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdInterval

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetDpdInterval(v int32)`

SetDpdInterval sets DpdInterval field to given value.

### HasDpdInterval

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasDpdInterval() bool`

HasDpdInterval returns a boolean if a field has been set.

### GetEncryptionAlg

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetEncryptionAlg() string`

GetEncryptionAlg returns the EncryptionAlg field if non-nil, zero value otherwise.

### GetEncryptionAlgOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetEncryptionAlgOk() (*string, bool)`

GetEncryptionAlgOk returns a tuple with the EncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlg

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetEncryptionAlg(v string)`

SetEncryptionAlg sets EncryptionAlg field to given value.

### HasEncryptionAlg

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasEncryptionAlg() bool`

HasEncryptionAlg returns a boolean if a field has been set.

### GetEsn

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetEsn() bool`

GetEsn returns the Esn field if non-nil, zero value otherwise.

### GetEsnOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetEsnOk() (*bool, bool)`

GetEsnOk returns a tuple with the Esn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsn

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetEsn(v bool)`

SetEsn sets Esn field to given value.

### HasEsn

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasEsn() bool`

HasEsn returns a boolean if a field has been set.

### GetIkeIntegrity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetIkeIntegrity() string`

GetIkeIntegrity returns the IkeIntegrity field if non-nil, zero value otherwise.

### GetIkeIntegrityOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetIkeIntegrityOk() (*string, bool)`

GetIkeIntegrityOk returns a tuple with the IkeIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeIntegrity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetIkeIntegrity(v string)`

SetIkeIntegrity sets IkeIntegrity field to given value.

### HasIkeIntegrity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasIkeIntegrity() bool`

HasIkeIntegrity returns a boolean if a field has been set.

### GetInitiator

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetInitiator() bool`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetInitiatorOk() (*bool, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetInitiator(v bool)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetIpsecEncryptionAlg

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetIpsecEncryptionAlg() string`

GetIpsecEncryptionAlg returns the IpsecEncryptionAlg field if non-nil, zero value otherwise.

### GetIpsecEncryptionAlgOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetIpsecEncryptionAlgOk() (*string, bool)`

GetIpsecEncryptionAlgOk returns a tuple with the IpsecEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecEncryptionAlg

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetIpsecEncryptionAlg(v string)`

SetIpsecEncryptionAlg sets IpsecEncryptionAlg field to given value.

### HasIpsecEncryptionAlg

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasIpsecEncryptionAlg() bool`

HasIpsecEncryptionAlg returns a boolean if a field has been set.

### GetIpsecIntegrity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetIpsecIntegrity() string`

GetIpsecIntegrity returns the IpsecIntegrity field if non-nil, zero value otherwise.

### GetIpsecIntegrityOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetIpsecIntegrityOk() (*string, bool)`

GetIpsecIntegrityOk returns a tuple with the IpsecIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecIntegrity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetIpsecIntegrity(v string)`

SetIpsecIntegrity sets IpsecIntegrity field to given value.

### HasIpsecIntegrity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasIpsecIntegrity() bool`

HasIpsecIntegrity returns a boolean if a field has been set.

### GetLabel

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLocalAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetLocalCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetLocalCircuit() string`

GetLocalCircuit returns the LocalCircuit field if non-nil, zero value otherwise.

### GetLocalCircuitOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetLocalCircuitOk() (*string, bool)`

GetLocalCircuitOk returns a tuple with the LocalCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetLocalCircuit(v string)`

SetLocalCircuit sets LocalCircuit field to given value.

### HasLocalCircuit

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasLocalCircuit() bool`

HasLocalCircuit returns a boolean if a field has been set.

### GetLocalIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetLocalIkePeerIdentity() string`

GetLocalIkePeerIdentity returns the LocalIkePeerIdentity field if non-nil, zero value otherwise.

### GetLocalIkePeerIdentityOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetLocalIkePeerIdentityOk() (*string, bool)`

GetLocalIkePeerIdentityOk returns a tuple with the LocalIkePeerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetLocalIkePeerIdentity(v string)`

SetLocalIkePeerIdentity sets LocalIkePeerIdentity field to given value.

### HasLocalIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasLocalIkePeerIdentity() bool`

HasLocalIkePeerIdentity returns a boolean if a field has been set.

### GetPerfectForwardSecrecy

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetPerfectForwardSecrecy() string`

GetPerfectForwardSecrecy returns the PerfectForwardSecrecy field if non-nil, zero value otherwise.

### GetPerfectForwardSecrecyOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetPerfectForwardSecrecyOk() (*string, bool)`

GetPerfectForwardSecrecyOk returns a tuple with the PerfectForwardSecrecy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectForwardSecrecy

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetPerfectForwardSecrecy(v string)`

SetPerfectForwardSecrecy sets PerfectForwardSecrecy field to given value.

### HasPerfectForwardSecrecy

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasPerfectForwardSecrecy() bool`

HasPerfectForwardSecrecy returns a boolean if a field has been set.

### GetPresharedKey

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetReauthInterval

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetReauthInterval() int32`

GetReauthInterval returns the ReauthInterval field if non-nil, zero value otherwise.

### GetReauthIntervalOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetReauthIntervalOk() (*int32, bool)`

GetReauthIntervalOk returns a tuple with the ReauthInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthInterval

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetReauthInterval(v int32)`

SetReauthInterval sets ReauthInterval field to given value.

### HasReauthInterval

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasReauthInterval() bool`

HasReauthInterval returns a boolean if a field has been set.

### GetRekeyInterval

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetRekeyInterval() int32`

GetRekeyInterval returns the RekeyInterval field if non-nil, zero value otherwise.

### GetRekeyIntervalOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetRekeyIntervalOk() (*int32, bool)`

GetRekeyIntervalOk returns a tuple with the RekeyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyInterval

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetRekeyInterval(v int32)`

SetRekeyInterval sets RekeyInterval field to given value.

### HasRekeyInterval

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasRekeyInterval() bool`

HasRekeyInterval returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetRemoteIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetRemoteIkePeerIdentity() string`

GetRemoteIkePeerIdentity returns the RemoteIkePeerIdentity field if non-nil, zero value otherwise.

### GetRemoteIkePeerIdentityOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) GetRemoteIkePeerIdentityOk() (*string, bool)`

GetRemoteIkePeerIdentityOk returns a tuple with the RemoteIkePeerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) SetRemoteIkePeerIdentity(v string)`

SetRemoteIkePeerIdentity sets RemoteIkePeerIdentity field to given value.

### HasRemoteIkePeerIdentity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceIpsec) HasRemoteIkePeerIdentity() bool`

HasRemoteIkePeerIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


