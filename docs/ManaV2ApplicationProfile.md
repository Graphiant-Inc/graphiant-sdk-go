# ManaV2ApplicationProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortRange** | Pointer to **string** | Port Range | [optional] 
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

### GetPortRange

`func (o *ManaV2ApplicationProfile) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *ManaV2ApplicationProfile) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *ManaV2ApplicationProfile) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *ManaV2ApplicationProfile) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

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


