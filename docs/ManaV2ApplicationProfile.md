# ManaV2ApplicationProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | **[]int32** |  | 
**Protocol** | **int32** | Protocol for the application profile (required) | 

## Methods

### NewManaV2ApplicationProfile

`func NewManaV2ApplicationProfile(ports []int32, protocol int32, ) *ManaV2ApplicationProfile`

NewManaV2ApplicationProfile instantiates a new ManaV2ApplicationProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2ApplicationProfileWithDefaults

`func NewManaV2ApplicationProfileWithDefaults() *ManaV2ApplicationProfile`

NewManaV2ApplicationProfileWithDefaults instantiates a new ManaV2ApplicationProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *ManaV2ApplicationProfile) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ManaV2ApplicationProfile) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ManaV2ApplicationProfile) SetPorts(v []int32)`

SetPorts sets Ports field to given value.


### GetProtocol

`func (o *ManaV2ApplicationProfile) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ManaV2ApplicationProfile) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ManaV2ApplicationProfile) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


