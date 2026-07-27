# ManaV2OspfAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Shared secret for OSPFv2 MD5 authentication on this interface | [optional] 
**KeyId** | Pointer to **int32** | Identifier of the OSPFv2 MD5 authentication key configured on this interface | [optional] 

## Methods

### NewManaV2OspfAuthentication

`func NewManaV2OspfAuthentication() *ManaV2OspfAuthentication`

NewManaV2OspfAuthentication instantiates a new ManaV2OspfAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2OspfAuthenticationWithDefaults

`func NewManaV2OspfAuthenticationWithDefaults() *ManaV2OspfAuthentication`

NewManaV2OspfAuthenticationWithDefaults instantiates a new ManaV2OspfAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ManaV2OspfAuthentication) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ManaV2OspfAuthentication) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ManaV2OspfAuthentication) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ManaV2OspfAuthentication) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeyId

`func (o *ManaV2OspfAuthentication) GetKeyId() int32`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ManaV2OspfAuthentication) GetKeyIdOk() (*int32, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ManaV2OspfAuthentication) SetKeyId(v int32)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ManaV2OspfAuthentication) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


