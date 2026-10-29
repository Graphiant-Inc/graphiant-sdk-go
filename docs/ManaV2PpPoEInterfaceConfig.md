# ManaV2PpPoEInterfaceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **bool** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**AuthenticationType** | Pointer to **string** |  | [optional] 
**Circuit** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LcpEchoFailure** | Pointer to **int32** |  | [optional] 
**LcpEchoInterval** | Pointer to **int32** |  | [optional] 
**MaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PppConnectTimeout** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**V4TcpMss** | Pointer to [**ManaV2NullableTcpMssV4**](ManaV2NullableTcpMssV4.md) |  | [optional] 
**V6TcpMss** | Pointer to [**ManaV2NullableTcpMssV6**](ManaV2NullableTcpMssV6.md) |  | [optional] 

## Methods

### NewManaV2PpPoEInterfaceConfig

`func NewManaV2PpPoEInterfaceConfig() *ManaV2PpPoEInterfaceConfig`

NewManaV2PpPoEInterfaceConfig instantiates a new ManaV2PpPoEInterfaceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PpPoEInterfaceConfigWithDefaults

`func NewManaV2PpPoEInterfaceConfigWithDefaults() *ManaV2PpPoEInterfaceConfig`

NewManaV2PpPoEInterfaceConfigWithDefaults instantiates a new ManaV2PpPoEInterfaceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *ManaV2PpPoEInterfaceConfig) GetAdminStatus() bool`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *ManaV2PpPoEInterfaceConfig) GetAdminStatusOk() (*bool, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *ManaV2PpPoEInterfaceConfig) SetAdminStatus(v bool)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *ManaV2PpPoEInterfaceConfig) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAlias

`func (o *ManaV2PpPoEInterfaceConfig) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ManaV2PpPoEInterfaceConfig) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ManaV2PpPoEInterfaceConfig) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ManaV2PpPoEInterfaceConfig) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *ManaV2PpPoEInterfaceConfig) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *ManaV2PpPoEInterfaceConfig) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *ManaV2PpPoEInterfaceConfig) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *ManaV2PpPoEInterfaceConfig) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetCircuit

`func (o *ManaV2PpPoEInterfaceConfig) GetCircuit() string`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *ManaV2PpPoEInterfaceConfig) GetCircuitOk() (*string, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *ManaV2PpPoEInterfaceConfig) SetCircuit(v string)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *ManaV2PpPoEInterfaceConfig) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetDescription

`func (o *ManaV2PpPoEInterfaceConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2PpPoEInterfaceConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2PpPoEInterfaceConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2PpPoEInterfaceConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLcpEchoFailure

`func (o *ManaV2PpPoEInterfaceConfig) GetLcpEchoFailure() int32`

GetLcpEchoFailure returns the LcpEchoFailure field if non-nil, zero value otherwise.

### GetLcpEchoFailureOk

`func (o *ManaV2PpPoEInterfaceConfig) GetLcpEchoFailureOk() (*int32, bool)`

GetLcpEchoFailureOk returns a tuple with the LcpEchoFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcpEchoFailure

`func (o *ManaV2PpPoEInterfaceConfig) SetLcpEchoFailure(v int32)`

SetLcpEchoFailure sets LcpEchoFailure field to given value.

### HasLcpEchoFailure

`func (o *ManaV2PpPoEInterfaceConfig) HasLcpEchoFailure() bool`

HasLcpEchoFailure returns a boolean if a field has been set.

### GetLcpEchoInterval

`func (o *ManaV2PpPoEInterfaceConfig) GetLcpEchoInterval() int32`

GetLcpEchoInterval returns the LcpEchoInterval field if non-nil, zero value otherwise.

### GetLcpEchoIntervalOk

`func (o *ManaV2PpPoEInterfaceConfig) GetLcpEchoIntervalOk() (*int32, bool)`

GetLcpEchoIntervalOk returns a tuple with the LcpEchoInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcpEchoInterval

`func (o *ManaV2PpPoEInterfaceConfig) SetLcpEchoInterval(v int32)`

SetLcpEchoInterval sets LcpEchoInterval field to given value.

### HasLcpEchoInterval

`func (o *ManaV2PpPoEInterfaceConfig) HasLcpEchoInterval() bool`

HasLcpEchoInterval returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *ManaV2PpPoEInterfaceConfig) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *ManaV2PpPoEInterfaceConfig) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *ManaV2PpPoEInterfaceConfig) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *ManaV2PpPoEInterfaceConfig) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetPassword

`func (o *ManaV2PpPoEInterfaceConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ManaV2PpPoEInterfaceConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ManaV2PpPoEInterfaceConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ManaV2PpPoEInterfaceConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPppConnectTimeout

`func (o *ManaV2PpPoEInterfaceConfig) GetPppConnectTimeout() int32`

GetPppConnectTimeout returns the PppConnectTimeout field if non-nil, zero value otherwise.

### GetPppConnectTimeoutOk

`func (o *ManaV2PpPoEInterfaceConfig) GetPppConnectTimeoutOk() (*int32, bool)`

GetPppConnectTimeoutOk returns a tuple with the PppConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPppConnectTimeout

`func (o *ManaV2PpPoEInterfaceConfig) SetPppConnectTimeout(v int32)`

SetPppConnectTimeout sets PppConnectTimeout field to given value.

### HasPppConnectTimeout

`func (o *ManaV2PpPoEInterfaceConfig) HasPppConnectTimeout() bool`

HasPppConnectTimeout returns a boolean if a field has been set.

### GetUsername

`func (o *ManaV2PpPoEInterfaceConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ManaV2PpPoEInterfaceConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ManaV2PpPoEInterfaceConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ManaV2PpPoEInterfaceConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetV4TcpMss

`func (o *ManaV2PpPoEInterfaceConfig) GetV4TcpMss() ManaV2NullableTcpMssV4`

GetV4TcpMss returns the V4TcpMss field if non-nil, zero value otherwise.

### GetV4TcpMssOk

`func (o *ManaV2PpPoEInterfaceConfig) GetV4TcpMssOk() (*ManaV2NullableTcpMssV4, bool)`

GetV4TcpMssOk returns a tuple with the V4TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4TcpMss

`func (o *ManaV2PpPoEInterfaceConfig) SetV4TcpMss(v ManaV2NullableTcpMssV4)`

SetV4TcpMss sets V4TcpMss field to given value.

### HasV4TcpMss

`func (o *ManaV2PpPoEInterfaceConfig) HasV4TcpMss() bool`

HasV4TcpMss returns a boolean if a field has been set.

### GetV6TcpMss

`func (o *ManaV2PpPoEInterfaceConfig) GetV6TcpMss() ManaV2NullableTcpMssV6`

GetV6TcpMss returns the V6TcpMss field if non-nil, zero value otherwise.

### GetV6TcpMssOk

`func (o *ManaV2PpPoEInterfaceConfig) GetV6TcpMssOk() (*ManaV2NullableTcpMssV6, bool)`

GetV6TcpMssOk returns a tuple with the V6TcpMss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6TcpMss

`func (o *ManaV2PpPoEInterfaceConfig) SetV6TcpMss(v ManaV2NullableTcpMssV6)`

SetV6TcpMss sets V6TcpMss field to given value.

### HasV6TcpMss

`func (o *ManaV2PpPoEInterfaceConfig) HasV6TcpMss() bool`

HasV6TcpMss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


