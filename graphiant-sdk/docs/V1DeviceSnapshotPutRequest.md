# V1DeviceSnapshotPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Golden** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SnapshotId** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1DeviceSnapshotPutRequest

`func NewV1DeviceSnapshotPutRequest() *V1DeviceSnapshotPutRequest`

NewV1DeviceSnapshotPutRequest instantiates a new V1DeviceSnapshotPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeviceSnapshotPutRequestWithDefaults

`func NewV1DeviceSnapshotPutRequestWithDefaults() *V1DeviceSnapshotPutRequest`

NewV1DeviceSnapshotPutRequestWithDefaults instantiates a new V1DeviceSnapshotPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGolden

`func (o *V1DeviceSnapshotPutRequest) GetGolden() bool`

GetGolden returns the Golden field if non-nil, zero value otherwise.

### GetGoldenOk

`func (o *V1DeviceSnapshotPutRequest) GetGoldenOk() (*bool, bool)`

GetGoldenOk returns a tuple with the Golden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGolden

`func (o *V1DeviceSnapshotPutRequest) SetGolden(v bool)`

SetGolden sets Golden field to given value.

### HasGolden

`func (o *V1DeviceSnapshotPutRequest) HasGolden() bool`

HasGolden returns a boolean if a field has been set.

### GetName

`func (o *V1DeviceSnapshotPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DeviceSnapshotPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DeviceSnapshotPutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DeviceSnapshotPutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSnapshotId

`func (o *V1DeviceSnapshotPutRequest) GetSnapshotId() int64`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *V1DeviceSnapshotPutRequest) GetSnapshotIdOk() (*int64, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *V1DeviceSnapshotPutRequest) SetSnapshotId(v int64)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *V1DeviceSnapshotPutRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


