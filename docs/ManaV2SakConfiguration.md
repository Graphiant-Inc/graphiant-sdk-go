# ManaV2SakConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CipherSuite** | Pointer to **string** | The cipher suite (required) | [optional] 
**LagMemberInterfaceId** | Pointer to **int64** | The interface ID (required for when each lag member has a different MACsec configuration - when split_sak_config_by_lag_member is true) (required) | [optional] 
**RekeyInterval** | Pointer to **int64** | The rekey interval in seconds. 0 means disabled | [optional] 
**ReplayProtectionWindowSize** | Pointer to **int64** | The replay protection window size in seconds. 0 means disabled | [optional] 

## Methods

### NewManaV2SakConfiguration

`func NewManaV2SakConfiguration() *ManaV2SakConfiguration`

NewManaV2SakConfiguration instantiates a new ManaV2SakConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2SakConfigurationWithDefaults

`func NewManaV2SakConfigurationWithDefaults() *ManaV2SakConfiguration`

NewManaV2SakConfigurationWithDefaults instantiates a new ManaV2SakConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCipherSuite

`func (o *ManaV2SakConfiguration) GetCipherSuite() string`

GetCipherSuite returns the CipherSuite field if non-nil, zero value otherwise.

### GetCipherSuiteOk

`func (o *ManaV2SakConfiguration) GetCipherSuiteOk() (*string, bool)`

GetCipherSuiteOk returns a tuple with the CipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuite

`func (o *ManaV2SakConfiguration) SetCipherSuite(v string)`

SetCipherSuite sets CipherSuite field to given value.

### HasCipherSuite

`func (o *ManaV2SakConfiguration) HasCipherSuite() bool`

HasCipherSuite returns a boolean if a field has been set.

### GetLagMemberInterfaceId

`func (o *ManaV2SakConfiguration) GetLagMemberInterfaceId() int64`

GetLagMemberInterfaceId returns the LagMemberInterfaceId field if non-nil, zero value otherwise.

### GetLagMemberInterfaceIdOk

`func (o *ManaV2SakConfiguration) GetLagMemberInterfaceIdOk() (*int64, bool)`

GetLagMemberInterfaceIdOk returns a tuple with the LagMemberInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagMemberInterfaceId

`func (o *ManaV2SakConfiguration) SetLagMemberInterfaceId(v int64)`

SetLagMemberInterfaceId sets LagMemberInterfaceId field to given value.

### HasLagMemberInterfaceId

`func (o *ManaV2SakConfiguration) HasLagMemberInterfaceId() bool`

HasLagMemberInterfaceId returns a boolean if a field has been set.

### GetRekeyInterval

`func (o *ManaV2SakConfiguration) GetRekeyInterval() int64`

GetRekeyInterval returns the RekeyInterval field if non-nil, zero value otherwise.

### GetRekeyIntervalOk

`func (o *ManaV2SakConfiguration) GetRekeyIntervalOk() (*int64, bool)`

GetRekeyIntervalOk returns a tuple with the RekeyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyInterval

`func (o *ManaV2SakConfiguration) SetRekeyInterval(v int64)`

SetRekeyInterval sets RekeyInterval field to given value.

### HasRekeyInterval

`func (o *ManaV2SakConfiguration) HasRekeyInterval() bool`

HasRekeyInterval returns a boolean if a field has been set.

### GetReplayProtectionWindowSize

`func (o *ManaV2SakConfiguration) GetReplayProtectionWindowSize() int64`

GetReplayProtectionWindowSize returns the ReplayProtectionWindowSize field if non-nil, zero value otherwise.

### GetReplayProtectionWindowSizeOk

`func (o *ManaV2SakConfiguration) GetReplayProtectionWindowSizeOk() (*int64, bool)`

GetReplayProtectionWindowSizeOk returns a tuple with the ReplayProtectionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayProtectionWindowSize

`func (o *ManaV2SakConfiguration) SetReplayProtectionWindowSize(v int64)`

SetReplayProtectionWindowSize sets ReplayProtectionWindowSize field to given value.

### HasReplayProtectionWindowSize

`func (o *ManaV2SakConfiguration) HasReplayProtectionWindowSize() bool`

HasReplayProtectionWindowSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


