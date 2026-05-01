# UpgradeRollout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]UpgradeRolloutDevice**](UpgradeRolloutDevice.md) |  | [optional] 
**HasFailed** | Pointer to **bool** | True if any device in the rollout has a failed upgrade state. | [optional] 
**Id** | Pointer to **int64** | Server-assigned rollout identifier. | [optional] 
**LastRunTs** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**NextRunTs** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**NumDevices** | Pointer to **int32** | Count of devices associated with the rollout. | [optional] 
**RolloutConfig** | Pointer to [**UpgradeRolloutConfig**](UpgradeRolloutConfig.md) |  | [optional] 

## Methods

### NewUpgradeRollout

`func NewUpgradeRollout() *UpgradeRollout`

NewUpgradeRollout instantiates a new UpgradeRollout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeRolloutWithDefaults

`func NewUpgradeRolloutWithDefaults() *UpgradeRollout`

NewUpgradeRolloutWithDefaults instantiates a new UpgradeRollout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *UpgradeRollout) GetDevices() []UpgradeRolloutDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *UpgradeRollout) GetDevicesOk() (*[]UpgradeRolloutDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *UpgradeRollout) SetDevices(v []UpgradeRolloutDevice)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *UpgradeRollout) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetHasFailed

`func (o *UpgradeRollout) GetHasFailed() bool`

GetHasFailed returns the HasFailed field if non-nil, zero value otherwise.

### GetHasFailedOk

`func (o *UpgradeRollout) GetHasFailedOk() (*bool, bool)`

GetHasFailedOk returns a tuple with the HasFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFailed

`func (o *UpgradeRollout) SetHasFailed(v bool)`

SetHasFailed sets HasFailed field to given value.

### HasHasFailed

`func (o *UpgradeRollout) HasHasFailed() bool`

HasHasFailed returns a boolean if a field has been set.

### GetId

`func (o *UpgradeRollout) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpgradeRollout) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpgradeRollout) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpgradeRollout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastRunTs

`func (o *UpgradeRollout) GetLastRunTs() GoogleProtobufTimestamp`

GetLastRunTs returns the LastRunTs field if non-nil, zero value otherwise.

### GetLastRunTsOk

`func (o *UpgradeRollout) GetLastRunTsOk() (*GoogleProtobufTimestamp, bool)`

GetLastRunTsOk returns a tuple with the LastRunTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunTs

`func (o *UpgradeRollout) SetLastRunTs(v GoogleProtobufTimestamp)`

SetLastRunTs sets LastRunTs field to given value.

### HasLastRunTs

`func (o *UpgradeRollout) HasLastRunTs() bool`

HasLastRunTs returns a boolean if a field has been set.

### GetNextRunTs

`func (o *UpgradeRollout) GetNextRunTs() GoogleProtobufTimestamp`

GetNextRunTs returns the NextRunTs field if non-nil, zero value otherwise.

### GetNextRunTsOk

`func (o *UpgradeRollout) GetNextRunTsOk() (*GoogleProtobufTimestamp, bool)`

GetNextRunTsOk returns a tuple with the NextRunTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunTs

`func (o *UpgradeRollout) SetNextRunTs(v GoogleProtobufTimestamp)`

SetNextRunTs sets NextRunTs field to given value.

### HasNextRunTs

`func (o *UpgradeRollout) HasNextRunTs() bool`

HasNextRunTs returns a boolean if a field has been set.

### GetNumDevices

`func (o *UpgradeRollout) GetNumDevices() int32`

GetNumDevices returns the NumDevices field if non-nil, zero value otherwise.

### GetNumDevicesOk

`func (o *UpgradeRollout) GetNumDevicesOk() (*int32, bool)`

GetNumDevicesOk returns a tuple with the NumDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDevices

`func (o *UpgradeRollout) SetNumDevices(v int32)`

SetNumDevices sets NumDevices field to given value.

### HasNumDevices

`func (o *UpgradeRollout) HasNumDevices() bool`

HasNumDevices returns a boolean if a field has been set.

### GetRolloutConfig

`func (o *UpgradeRollout) GetRolloutConfig() UpgradeRolloutConfig`

GetRolloutConfig returns the RolloutConfig field if non-nil, zero value otherwise.

### GetRolloutConfigOk

`func (o *UpgradeRollout) GetRolloutConfigOk() (*UpgradeRolloutConfig, bool)`

GetRolloutConfigOk returns a tuple with the RolloutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutConfig

`func (o *UpgradeRollout) SetRolloutConfig(v UpgradeRolloutConfig)`

SetRolloutConfig sets RolloutConfig field to given value.

### HasRolloutConfig

`func (o *UpgradeRollout) HasRolloutConfig() bool`

HasRolloutConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


