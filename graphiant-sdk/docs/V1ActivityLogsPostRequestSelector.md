# V1ActivityLogsPostRequestSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | Pointer to **[]int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InProgress** | Pointer to **bool** |  | [optional] 
**JobEntity** | Pointer to [**V1ActivityLogsPostRequestSelectorJobEntity**](V1ActivityLogsPostRequestSelectorJobEntity.md) |  | [optional] 
**TargetIds** | Pointer to [**[]V1ActivityLogsPostRequestSelectorJobEntity**](V1ActivityLogsPostRequestSelectorJobEntity.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ActivityLogsPostRequestSelector

`func NewV1ActivityLogsPostRequestSelector() *V1ActivityLogsPostRequestSelector`

NewV1ActivityLogsPostRequestSelector instantiates a new V1ActivityLogsPostRequestSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ActivityLogsPostRequestSelectorWithDefaults

`func NewV1ActivityLogsPostRequestSelectorWithDefaults() *V1ActivityLogsPostRequestSelector`

NewV1ActivityLogsPostRequestSelectorWithDefaults instantiates a new V1ActivityLogsPostRequestSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *V1ActivityLogsPostRequestSelector) GetDeviceIds() []int64`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *V1ActivityLogsPostRequestSelector) GetDeviceIdsOk() (*[]int64, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *V1ActivityLogsPostRequestSelector) SetDeviceIds(v []int64)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *V1ActivityLogsPostRequestSelector) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetId

`func (o *V1ActivityLogsPostRequestSelector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ActivityLogsPostRequestSelector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ActivityLogsPostRequestSelector) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1ActivityLogsPostRequestSelector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInProgress

`func (o *V1ActivityLogsPostRequestSelector) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *V1ActivityLogsPostRequestSelector) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *V1ActivityLogsPostRequestSelector) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *V1ActivityLogsPostRequestSelector) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetJobEntity

`func (o *V1ActivityLogsPostRequestSelector) GetJobEntity() V1ActivityLogsPostRequestSelectorJobEntity`

GetJobEntity returns the JobEntity field if non-nil, zero value otherwise.

### GetJobEntityOk

`func (o *V1ActivityLogsPostRequestSelector) GetJobEntityOk() (*V1ActivityLogsPostRequestSelectorJobEntity, bool)`

GetJobEntityOk returns a tuple with the JobEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobEntity

`func (o *V1ActivityLogsPostRequestSelector) SetJobEntity(v V1ActivityLogsPostRequestSelectorJobEntity)`

SetJobEntity sets JobEntity field to given value.

### HasJobEntity

`func (o *V1ActivityLogsPostRequestSelector) HasJobEntity() bool`

HasJobEntity returns a boolean if a field has been set.

### GetTargetIds

`func (o *V1ActivityLogsPostRequestSelector) GetTargetIds() []V1ActivityLogsPostRequestSelectorJobEntity`

GetTargetIds returns the TargetIds field if non-nil, zero value otherwise.

### GetTargetIdsOk

`func (o *V1ActivityLogsPostRequestSelector) GetTargetIdsOk() (*[]V1ActivityLogsPostRequestSelectorJobEntity, bool)`

GetTargetIdsOk returns a tuple with the TargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIds

`func (o *V1ActivityLogsPostRequestSelector) SetTargetIds(v []V1ActivityLogsPostRequestSelectorJobEntity)`

SetTargetIds sets TargetIds field to given value.

### HasTargetIds

`func (o *V1ActivityLogsPostRequestSelector) HasTargetIds() bool`

HasTargetIds returns a boolean if a field has been set.

### GetType

`func (o *V1ActivityLogsPostRequestSelector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ActivityLogsPostRequestSelector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ActivityLogsPostRequestSelector) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1ActivityLogsPostRequestSelector) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


