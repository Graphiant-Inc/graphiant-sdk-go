# ManaV2PublicVifDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerBurstSize** | **int32** | Maximum Burst size per device (required) | 
**ConsumerBwSite** | **int32** | Maximum Bandwidth allocation per device (required) | 
**NatPools** | **[]string** |  | 

## Methods

### NewManaV2PublicVifDevice

`func NewManaV2PublicVifDevice(consumerBurstSize int32, consumerBwSite int32, natPools []string, ) *ManaV2PublicVifDevice`

NewManaV2PublicVifDevice instantiates a new ManaV2PublicVifDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PublicVifDeviceWithDefaults

`func NewManaV2PublicVifDeviceWithDefaults() *ManaV2PublicVifDevice`

NewManaV2PublicVifDeviceWithDefaults instantiates a new ManaV2PublicVifDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerBurstSize

`func (o *ManaV2PublicVifDevice) GetConsumerBurstSize() int32`

GetConsumerBurstSize returns the ConsumerBurstSize field if non-nil, zero value otherwise.

### GetConsumerBurstSizeOk

`func (o *ManaV2PublicVifDevice) GetConsumerBurstSizeOk() (*int32, bool)`

GetConsumerBurstSizeOk returns a tuple with the ConsumerBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerBurstSize

`func (o *ManaV2PublicVifDevice) SetConsumerBurstSize(v int32)`

SetConsumerBurstSize sets ConsumerBurstSize field to given value.


### GetConsumerBwSite

`func (o *ManaV2PublicVifDevice) GetConsumerBwSite() int32`

GetConsumerBwSite returns the ConsumerBwSite field if non-nil, zero value otherwise.

### GetConsumerBwSiteOk

`func (o *ManaV2PublicVifDevice) GetConsumerBwSiteOk() (*int32, bool)`

GetConsumerBwSiteOk returns a tuple with the ConsumerBwSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerBwSite

`func (o *ManaV2PublicVifDevice) SetConsumerBwSite(v int32)`

SetConsumerBwSite sets ConsumerBwSite field to given value.


### GetNatPools

`func (o *ManaV2PublicVifDevice) GetNatPools() []string`

GetNatPools returns the NatPools field if non-nil, zero value otherwise.

### GetNatPoolsOk

`func (o *ManaV2PublicVifDevice) GetNatPoolsOk() (*[]string, bool)`

GetNatPoolsOk returns a tuple with the NatPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPools

`func (o *ManaV2PublicVifDevice) SetNatPools(v []string)`

SetNatPools sets NatPools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


