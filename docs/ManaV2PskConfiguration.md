# ManaV2PskConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cak** | Pointer to **string** | The Connection Authentication Key (CAK) | [optional] 
**CakCryptographicAlgorithm** | Pointer to **string** | The cryptographic algorithm for the CAK (required) | [optional] 
**Ckn** | Pointer to **string** | The Connection Key Name (CKN) (required) | [optional] 
**Nickname** | Pointer to **string** | The nickname of the PSK (required) | [optional] 
**StartTime** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 

## Methods

### NewManaV2PskConfiguration

`func NewManaV2PskConfiguration() *ManaV2PskConfiguration`

NewManaV2PskConfiguration instantiates a new ManaV2PskConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PskConfigurationWithDefaults

`func NewManaV2PskConfigurationWithDefaults() *ManaV2PskConfiguration`

NewManaV2PskConfigurationWithDefaults instantiates a new ManaV2PskConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCak

`func (o *ManaV2PskConfiguration) GetCak() string`

GetCak returns the Cak field if non-nil, zero value otherwise.

### GetCakOk

`func (o *ManaV2PskConfiguration) GetCakOk() (*string, bool)`

GetCakOk returns a tuple with the Cak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCak

`func (o *ManaV2PskConfiguration) SetCak(v string)`

SetCak sets Cak field to given value.

### HasCak

`func (o *ManaV2PskConfiguration) HasCak() bool`

HasCak returns a boolean if a field has been set.

### GetCakCryptographicAlgorithm

`func (o *ManaV2PskConfiguration) GetCakCryptographicAlgorithm() string`

GetCakCryptographicAlgorithm returns the CakCryptographicAlgorithm field if non-nil, zero value otherwise.

### GetCakCryptographicAlgorithmOk

`func (o *ManaV2PskConfiguration) GetCakCryptographicAlgorithmOk() (*string, bool)`

GetCakCryptographicAlgorithmOk returns a tuple with the CakCryptographicAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCakCryptographicAlgorithm

`func (o *ManaV2PskConfiguration) SetCakCryptographicAlgorithm(v string)`

SetCakCryptographicAlgorithm sets CakCryptographicAlgorithm field to given value.

### HasCakCryptographicAlgorithm

`func (o *ManaV2PskConfiguration) HasCakCryptographicAlgorithm() bool`

HasCakCryptographicAlgorithm returns a boolean if a field has been set.

### GetCkn

`func (o *ManaV2PskConfiguration) GetCkn() string`

GetCkn returns the Ckn field if non-nil, zero value otherwise.

### GetCknOk

`func (o *ManaV2PskConfiguration) GetCknOk() (*string, bool)`

GetCknOk returns a tuple with the Ckn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCkn

`func (o *ManaV2PskConfiguration) SetCkn(v string)`

SetCkn sets Ckn field to given value.

### HasCkn

`func (o *ManaV2PskConfiguration) HasCkn() bool`

HasCkn returns a boolean if a field has been set.

### GetNickname

`func (o *ManaV2PskConfiguration) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *ManaV2PskConfiguration) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *ManaV2PskConfiguration) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *ManaV2PskConfiguration) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetStartTime

`func (o *ManaV2PskConfiguration) GetStartTime() GoogleProtobufTimestamp`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ManaV2PskConfiguration) GetStartTimeOk() (*GoogleProtobufTimestamp, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ManaV2PskConfiguration) SetStartTime(v GoogleProtobufTimestamp)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ManaV2PskConfiguration) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


