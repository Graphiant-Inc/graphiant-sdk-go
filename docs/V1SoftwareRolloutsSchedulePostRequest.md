# V1SoftwareRolloutsSchedulePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedOnly** | Pointer to **bool** | When true, only devices previously marked failed are rescheduled. | [optional] 
**Id** | **int64** | Rollout identifier to schedule. (required) | 
**Ts** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 

## Methods

### NewV1SoftwareRolloutsSchedulePostRequest

`func NewV1SoftwareRolloutsSchedulePostRequest(id int64, ) *V1SoftwareRolloutsSchedulePostRequest`

NewV1SoftwareRolloutsSchedulePostRequest instantiates a new V1SoftwareRolloutsSchedulePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SoftwareRolloutsSchedulePostRequestWithDefaults

`func NewV1SoftwareRolloutsSchedulePostRequestWithDefaults() *V1SoftwareRolloutsSchedulePostRequest`

NewV1SoftwareRolloutsSchedulePostRequestWithDefaults instantiates a new V1SoftwareRolloutsSchedulePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedOnly

`func (o *V1SoftwareRolloutsSchedulePostRequest) GetFailedOnly() bool`

GetFailedOnly returns the FailedOnly field if non-nil, zero value otherwise.

### GetFailedOnlyOk

`func (o *V1SoftwareRolloutsSchedulePostRequest) GetFailedOnlyOk() (*bool, bool)`

GetFailedOnlyOk returns a tuple with the FailedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedOnly

`func (o *V1SoftwareRolloutsSchedulePostRequest) SetFailedOnly(v bool)`

SetFailedOnly sets FailedOnly field to given value.

### HasFailedOnly

`func (o *V1SoftwareRolloutsSchedulePostRequest) HasFailedOnly() bool`

HasFailedOnly returns a boolean if a field has been set.

### GetId

`func (o *V1SoftwareRolloutsSchedulePostRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1SoftwareRolloutsSchedulePostRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1SoftwareRolloutsSchedulePostRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetTs

`func (o *V1SoftwareRolloutsSchedulePostRequest) GetTs() GoogleProtobufTimestamp`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *V1SoftwareRolloutsSchedulePostRequest) GetTsOk() (*GoogleProtobufTimestamp, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *V1SoftwareRolloutsSchedulePostRequest) SetTs(v GoogleProtobufTimestamp)`

SetTs sets Ts field to given value.

### HasTs

`func (o *V1SoftwareRolloutsSchedulePostRequest) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


