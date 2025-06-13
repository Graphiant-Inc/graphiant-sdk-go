# V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **string** |  | [optional] 
**AdvertisementRcvd** | Pointer to **int64** |  | [optional] 
**AdvertisementSent** | Pointer to **int64** |  | [optional] 
**EffectivePriority** | Pointer to **int32** |  | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**IsOwner** | Pointer to **bool** |  | [optional] 
**MasterTransition** | Pointer to **int32** |  | [optional] 
**NewMasterReason** | Pointer to **string** |  | [optional] 
**PriorityZeroPktsRcvd** | Pointer to **int64** |  | [optional] 
**PriorityZeroPktsSent** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner

`func NewV1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner() *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner`

NewV1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner instantiates a new V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInnerWithDefaults

`func NewV1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInnerWithDefaults() *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner`

NewV1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInnerWithDefaults instantiates a new V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAdvertisementRcvd

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetAdvertisementRcvd() int64`

GetAdvertisementRcvd returns the AdvertisementRcvd field if non-nil, zero value otherwise.

### GetAdvertisementRcvdOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetAdvertisementRcvdOk() (*int64, bool)`

GetAdvertisementRcvdOk returns a tuple with the AdvertisementRcvd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisementRcvd

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetAdvertisementRcvd(v int64)`

SetAdvertisementRcvd sets AdvertisementRcvd field to given value.

### HasAdvertisementRcvd

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasAdvertisementRcvd() bool`

HasAdvertisementRcvd returns a boolean if a field has been set.

### GetAdvertisementSent

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetAdvertisementSent() int64`

GetAdvertisementSent returns the AdvertisementSent field if non-nil, zero value otherwise.

### GetAdvertisementSentOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetAdvertisementSentOk() (*int64, bool)`

GetAdvertisementSentOk returns a tuple with the AdvertisementSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisementSent

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetAdvertisementSent(v int64)`

SetAdvertisementSent sets AdvertisementSent field to given value.

### HasAdvertisementSent

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasAdvertisementSent() bool`

HasAdvertisementSent returns a boolean if a field has been set.

### GetEffectivePriority

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetEffectivePriority() int32`

GetEffectivePriority returns the EffectivePriority field if non-nil, zero value otherwise.

### GetEffectivePriorityOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetEffectivePriorityOk() (*int32, bool)`

GetEffectivePriorityOk returns a tuple with the EffectivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePriority

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetEffectivePriority(v int32)`

SetEffectivePriority sets EffectivePriority field to given value.

### HasEffectivePriority

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasEffectivePriority() bool`

HasEffectivePriority returns a boolean if a field has been set.

### GetGroupId

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIsOwner

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetMasterTransition

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetMasterTransition() int32`

GetMasterTransition returns the MasterTransition field if non-nil, zero value otherwise.

### GetMasterTransitionOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetMasterTransitionOk() (*int32, bool)`

GetMasterTransitionOk returns a tuple with the MasterTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterTransition

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetMasterTransition(v int32)`

SetMasterTransition sets MasterTransition field to given value.

### HasMasterTransition

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasMasterTransition() bool`

HasMasterTransition returns a boolean if a field has been set.

### GetNewMasterReason

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetNewMasterReason() string`

GetNewMasterReason returns the NewMasterReason field if non-nil, zero value otherwise.

### GetNewMasterReasonOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetNewMasterReasonOk() (*string, bool)`

GetNewMasterReasonOk returns a tuple with the NewMasterReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMasterReason

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetNewMasterReason(v string)`

SetNewMasterReason sets NewMasterReason field to given value.

### HasNewMasterReason

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasNewMasterReason() bool`

HasNewMasterReason returns a boolean if a field has been set.

### GetPriorityZeroPktsRcvd

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetPriorityZeroPktsRcvd() int64`

GetPriorityZeroPktsRcvd returns the PriorityZeroPktsRcvd field if non-nil, zero value otherwise.

### GetPriorityZeroPktsRcvdOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetPriorityZeroPktsRcvdOk() (*int64, bool)`

GetPriorityZeroPktsRcvdOk returns a tuple with the PriorityZeroPktsRcvd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityZeroPktsRcvd

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetPriorityZeroPktsRcvd(v int64)`

SetPriorityZeroPktsRcvd sets PriorityZeroPktsRcvd field to given value.

### HasPriorityZeroPktsRcvd

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasPriorityZeroPktsRcvd() bool`

HasPriorityZeroPktsRcvd returns a boolean if a field has been set.

### GetPriorityZeroPktsSent

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetPriorityZeroPktsSent() int64`

GetPriorityZeroPktsSent returns the PriorityZeroPktsSent field if non-nil, zero value otherwise.

### GetPriorityZeroPktsSentOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetPriorityZeroPktsSentOk() (*int64, bool)`

GetPriorityZeroPktsSentOk returns a tuple with the PriorityZeroPktsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityZeroPktsSent

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetPriorityZeroPktsSent(v int64)`

SetPriorityZeroPktsSent sets PriorityZeroPktsSent field to given value.

### HasPriorityZeroPktsSent

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasPriorityZeroPktsSent() bool`

HasPriorityZeroPktsSent returns a boolean if a field has been set.

### GetState

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V1DevicesDeviceIdVrrpGet200ResponseVrrpEntryInner) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


