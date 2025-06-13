# V1DevicesBringupPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | Pointer to **[]int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DevicesBringupPutRequest

`func NewV1DevicesBringupPutRequest() *V1DevicesBringupPutRequest`

NewV1DevicesBringupPutRequest instantiates a new V1DevicesBringupPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesBringupPutRequestWithDefaults

`func NewV1DevicesBringupPutRequestWithDefaults() *V1DevicesBringupPutRequest`

NewV1DevicesBringupPutRequestWithDefaults instantiates a new V1DevicesBringupPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *V1DevicesBringupPutRequest) GetDeviceIds() []int64`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *V1DevicesBringupPutRequest) GetDeviceIdsOk() (*[]int64, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *V1DevicesBringupPutRequest) SetDeviceIds(v []int64)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *V1DevicesBringupPutRequest) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetStatus

`func (o *V1DevicesBringupPutRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1DevicesBringupPutRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1DevicesBringupPutRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1DevicesBringupPutRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


