# V1DeviceSnapshotPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**Golden** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DeviceSnapshotPostRequest

`func NewV1DeviceSnapshotPostRequest() *V1DeviceSnapshotPostRequest`

NewV1DeviceSnapshotPostRequest instantiates a new V1DeviceSnapshotPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeviceSnapshotPostRequestWithDefaults

`func NewV1DeviceSnapshotPostRequestWithDefaults() *V1DeviceSnapshotPostRequest`

NewV1DeviceSnapshotPostRequestWithDefaults instantiates a new V1DeviceSnapshotPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *V1DeviceSnapshotPostRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *V1DeviceSnapshotPostRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *V1DeviceSnapshotPostRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *V1DeviceSnapshotPostRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1DeviceSnapshotPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1DeviceSnapshotPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1DeviceSnapshotPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1DeviceSnapshotPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetGolden

`func (o *V1DeviceSnapshotPostRequest) GetGolden() bool`

GetGolden returns the Golden field if non-nil, zero value otherwise.

### GetGoldenOk

`func (o *V1DeviceSnapshotPostRequest) GetGoldenOk() (*bool, bool)`

GetGoldenOk returns a tuple with the Golden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGolden

`func (o *V1DeviceSnapshotPostRequest) SetGolden(v bool)`

SetGolden sets Golden field to given value.

### HasGolden

`func (o *V1DeviceSnapshotPostRequest) HasGolden() bool`

HasGolden returns a boolean if a field has been set.

### GetName

`func (o *V1DeviceSnapshotPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DeviceSnapshotPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DeviceSnapshotPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DeviceSnapshotPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


