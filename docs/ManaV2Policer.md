# ManaV2Policer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurstSize** | **int32** | Burst size for the policer (required) | 
**Bw** | **int32** | Bandwidth for the policer (required) | 

## Methods

### NewManaV2Policer

`func NewManaV2Policer(burstSize int32, bw int32, ) *ManaV2Policer`

NewManaV2Policer instantiates a new ManaV2Policer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PolicerWithDefaults

`func NewManaV2PolicerWithDefaults() *ManaV2Policer`

NewManaV2PolicerWithDefaults instantiates a new ManaV2Policer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurstSize

`func (o *ManaV2Policer) GetBurstSize() int32`

GetBurstSize returns the BurstSize field if non-nil, zero value otherwise.

### GetBurstSizeOk

`func (o *ManaV2Policer) GetBurstSizeOk() (*int32, bool)`

GetBurstSizeOk returns a tuple with the BurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstSize

`func (o *ManaV2Policer) SetBurstSize(v int32)`

SetBurstSize sets BurstSize field to given value.


### GetBw

`func (o *ManaV2Policer) GetBw() int32`

GetBw returns the Bw field if non-nil, zero value otherwise.

### GetBwOk

`func (o *ManaV2Policer) GetBwOk() (*int32, bool)`

GetBwOk returns a tuple with the Bw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBw

`func (o *ManaV2Policer) SetBw(v int32)`

SetBw sets Bw field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


