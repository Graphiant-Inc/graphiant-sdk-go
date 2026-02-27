# V2DeviceDeviceIdTopologyPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotTime** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**TimeWindow** | Pointer to [**StatsmonV2TimeWindow**](StatsmonV2TimeWindow.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewV2DeviceDeviceIdTopologyPostRequest

`func NewV2DeviceDeviceIdTopologyPostRequest() *V2DeviceDeviceIdTopologyPostRequest`

NewV2DeviceDeviceIdTopologyPostRequest instantiates a new V2DeviceDeviceIdTopologyPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2DeviceDeviceIdTopologyPostRequestWithDefaults

`func NewV2DeviceDeviceIdTopologyPostRequestWithDefaults() *V2DeviceDeviceIdTopologyPostRequest`

NewV2DeviceDeviceIdTopologyPostRequestWithDefaults instantiates a new V2DeviceDeviceIdTopologyPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotTime

`func (o *V2DeviceDeviceIdTopologyPostRequest) GetSnapshotTime() GoogleProtobufTimestamp`

GetSnapshotTime returns the SnapshotTime field if non-nil, zero value otherwise.

### GetSnapshotTimeOk

`func (o *V2DeviceDeviceIdTopologyPostRequest) GetSnapshotTimeOk() (*GoogleProtobufTimestamp, bool)`

GetSnapshotTimeOk returns a tuple with the SnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTime

`func (o *V2DeviceDeviceIdTopologyPostRequest) SetSnapshotTime(v GoogleProtobufTimestamp)`

SetSnapshotTime sets SnapshotTime field to given value.

### HasSnapshotTime

`func (o *V2DeviceDeviceIdTopologyPostRequest) HasSnapshotTime() bool`

HasSnapshotTime returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2DeviceDeviceIdTopologyPostRequest) GetTimeWindow() StatsmonV2TimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2DeviceDeviceIdTopologyPostRequest) GetTimeWindowOk() (*StatsmonV2TimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2DeviceDeviceIdTopologyPostRequest) SetTimeWindow(v StatsmonV2TimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2DeviceDeviceIdTopologyPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetVersion

`func (o *V2DeviceDeviceIdTopologyPostRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V2DeviceDeviceIdTopologyPostRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V2DeviceDeviceIdTopologyPostRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V2DeviceDeviceIdTopologyPostRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


