# V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiReplayWindowSize** | Pointer to **int32** |  | [optional] 
**DpdInterval** | Pointer to **int32** |  | [optional] 
**Esn** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IkeDhGroup** | Pointer to **string** |  | [optional] 
**IkeEncryptionAlg** | Pointer to **string** |  | [optional] 
**IkeIntegrity** | Pointer to **string** |  | [optional] 
**IpsecEncryptionAlg** | Pointer to **string** |  | [optional] 
**IpsecIntegrity** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PerfectForwardSecrecy** | Pointer to **string** |  | [optional] 
**ReauthInterval** | Pointer to **int32** |  | [optional] 
**RekeyInterval** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestVpnProfilesValueVpnProfile

`func NewV1GlobalConfigPatchRequestVpnProfilesValueVpnProfile() *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile`

NewV1GlobalConfigPatchRequestVpnProfilesValueVpnProfile instantiates a new V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestVpnProfilesValueVpnProfileWithDefaults

`func NewV1GlobalConfigPatchRequestVpnProfilesValueVpnProfileWithDefaults() *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile`

NewV1GlobalConfigPatchRequestVpnProfilesValueVpnProfileWithDefaults instantiates a new V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntiReplayWindowSize

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetAntiReplayWindowSize() int32`

GetAntiReplayWindowSize returns the AntiReplayWindowSize field if non-nil, zero value otherwise.

### GetAntiReplayWindowSizeOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetAntiReplayWindowSizeOk() (*int32, bool)`

GetAntiReplayWindowSizeOk returns a tuple with the AntiReplayWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiReplayWindowSize

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetAntiReplayWindowSize(v int32)`

SetAntiReplayWindowSize sets AntiReplayWindowSize field to given value.

### HasAntiReplayWindowSize

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasAntiReplayWindowSize() bool`

HasAntiReplayWindowSize returns a boolean if a field has been set.

### GetDpdInterval

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetDpdInterval() int32`

GetDpdInterval returns the DpdInterval field if non-nil, zero value otherwise.

### GetDpdIntervalOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetDpdIntervalOk() (*int32, bool)`

GetDpdIntervalOk returns a tuple with the DpdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdInterval

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetDpdInterval(v int32)`

SetDpdInterval sets DpdInterval field to given value.

### HasDpdInterval

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasDpdInterval() bool`

HasDpdInterval returns a boolean if a field has been set.

### GetEsn

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetEsn() bool`

GetEsn returns the Esn field if non-nil, zero value otherwise.

### GetEsnOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetEsnOk() (*bool, bool)`

GetEsnOk returns a tuple with the Esn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsn

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetEsn(v bool)`

SetEsn sets Esn field to given value.

### HasEsn

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasEsn() bool`

HasEsn returns a boolean if a field has been set.

### GetId

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIkeDhGroup

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIkeDhGroup() string`

GetIkeDhGroup returns the IkeDhGroup field if non-nil, zero value otherwise.

### GetIkeDhGroupOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIkeDhGroupOk() (*string, bool)`

GetIkeDhGroupOk returns a tuple with the IkeDhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeDhGroup

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetIkeDhGroup(v string)`

SetIkeDhGroup sets IkeDhGroup field to given value.

### HasIkeDhGroup

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasIkeDhGroup() bool`

HasIkeDhGroup returns a boolean if a field has been set.

### GetIkeEncryptionAlg

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIkeEncryptionAlg() string`

GetIkeEncryptionAlg returns the IkeEncryptionAlg field if non-nil, zero value otherwise.

### GetIkeEncryptionAlgOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIkeEncryptionAlgOk() (*string, bool)`

GetIkeEncryptionAlgOk returns a tuple with the IkeEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeEncryptionAlg

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetIkeEncryptionAlg(v string)`

SetIkeEncryptionAlg sets IkeEncryptionAlg field to given value.

### HasIkeEncryptionAlg

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasIkeEncryptionAlg() bool`

HasIkeEncryptionAlg returns a boolean if a field has been set.

### GetIkeIntegrity

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIkeIntegrity() string`

GetIkeIntegrity returns the IkeIntegrity field if non-nil, zero value otherwise.

### GetIkeIntegrityOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIkeIntegrityOk() (*string, bool)`

GetIkeIntegrityOk returns a tuple with the IkeIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeIntegrity

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetIkeIntegrity(v string)`

SetIkeIntegrity sets IkeIntegrity field to given value.

### HasIkeIntegrity

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasIkeIntegrity() bool`

HasIkeIntegrity returns a boolean if a field has been set.

### GetIpsecEncryptionAlg

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIpsecEncryptionAlg() string`

GetIpsecEncryptionAlg returns the IpsecEncryptionAlg field if non-nil, zero value otherwise.

### GetIpsecEncryptionAlgOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIpsecEncryptionAlgOk() (*string, bool)`

GetIpsecEncryptionAlgOk returns a tuple with the IpsecEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecEncryptionAlg

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetIpsecEncryptionAlg(v string)`

SetIpsecEncryptionAlg sets IpsecEncryptionAlg field to given value.

### HasIpsecEncryptionAlg

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasIpsecEncryptionAlg() bool`

HasIpsecEncryptionAlg returns a boolean if a field has been set.

### GetIpsecIntegrity

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIpsecIntegrity() string`

GetIpsecIntegrity returns the IpsecIntegrity field if non-nil, zero value otherwise.

### GetIpsecIntegrityOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetIpsecIntegrityOk() (*string, bool)`

GetIpsecIntegrityOk returns a tuple with the IpsecIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecIntegrity

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetIpsecIntegrity(v string)`

SetIpsecIntegrity sets IpsecIntegrity field to given value.

### HasIpsecIntegrity

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasIpsecIntegrity() bool`

HasIpsecIntegrity returns a boolean if a field has been set.

### GetName

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPerfectForwardSecrecy

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetPerfectForwardSecrecy() string`

GetPerfectForwardSecrecy returns the PerfectForwardSecrecy field if non-nil, zero value otherwise.

### GetPerfectForwardSecrecyOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetPerfectForwardSecrecyOk() (*string, bool)`

GetPerfectForwardSecrecyOk returns a tuple with the PerfectForwardSecrecy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectForwardSecrecy

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetPerfectForwardSecrecy(v string)`

SetPerfectForwardSecrecy sets PerfectForwardSecrecy field to given value.

### HasPerfectForwardSecrecy

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasPerfectForwardSecrecy() bool`

HasPerfectForwardSecrecy returns a boolean if a field has been set.

### GetReauthInterval

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetReauthInterval() int32`

GetReauthInterval returns the ReauthInterval field if non-nil, zero value otherwise.

### GetReauthIntervalOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetReauthIntervalOk() (*int32, bool)`

GetReauthIntervalOk returns a tuple with the ReauthInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthInterval

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetReauthInterval(v int32)`

SetReauthInterval sets ReauthInterval field to given value.

### HasReauthInterval

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasReauthInterval() bool`

HasReauthInterval returns a boolean if a field has been set.

### GetRekeyInterval

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetRekeyInterval() int32`

GetRekeyInterval returns the RekeyInterval field if non-nil, zero value otherwise.

### GetRekeyIntervalOk

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) GetRekeyIntervalOk() (*int32, bool)`

GetRekeyIntervalOk returns a tuple with the RekeyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyInterval

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) SetRekeyInterval(v int32)`

SetRekeyInterval sets RekeyInterval field to given value.

### HasRekeyInterval

`func (o *V1GlobalConfigPatchRequestVpnProfilesValueVpnProfile) HasRekeyInterval() bool`

HasRekeyInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


