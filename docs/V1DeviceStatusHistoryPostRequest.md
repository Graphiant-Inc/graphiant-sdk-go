# V1DeviceStatusHistoryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | Pointer to **[]int64** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DeviceStatusHistoryPostRequest

`func NewV1DeviceStatusHistoryPostRequest() *V1DeviceStatusHistoryPostRequest`

NewV1DeviceStatusHistoryPostRequest instantiates a new V1DeviceStatusHistoryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeviceStatusHistoryPostRequestWithDefaults

`func NewV1DeviceStatusHistoryPostRequestWithDefaults() *V1DeviceStatusHistoryPostRequest`

NewV1DeviceStatusHistoryPostRequestWithDefaults instantiates a new V1DeviceStatusHistoryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *V1DeviceStatusHistoryPostRequest) GetDeviceIds() []int64`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *V1DeviceStatusHistoryPostRequest) GetDeviceIdsOk() (*[]int64, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *V1DeviceStatusHistoryPostRequest) SetDeviceIds(v []int64)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *V1DeviceStatusHistoryPostRequest) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetRole

`func (o *V1DeviceStatusHistoryPostRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *V1DeviceStatusHistoryPostRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *V1DeviceStatusHistoryPostRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *V1DeviceStatusHistoryPostRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


