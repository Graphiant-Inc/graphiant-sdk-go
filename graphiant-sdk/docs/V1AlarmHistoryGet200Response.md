# V1AlarmHistoryGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**History** | Pointer to [**[]V1AlarmHistoryGet200ResponseHistoryInner**](V1AlarmHistoryGet200ResponseHistoryInner.md) |  | [optional] 

## Methods

### NewV1AlarmHistoryGet200Response

`func NewV1AlarmHistoryGet200Response() *V1AlarmHistoryGet200Response`

NewV1AlarmHistoryGet200Response instantiates a new V1AlarmHistoryGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AlarmHistoryGet200ResponseWithDefaults

`func NewV1AlarmHistoryGet200ResponseWithDefaults() *V1AlarmHistoryGet200Response`

NewV1AlarmHistoryGet200ResponseWithDefaults instantiates a new V1AlarmHistoryGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistory

`func (o *V1AlarmHistoryGet200Response) GetHistory() []V1AlarmHistoryGet200ResponseHistoryInner`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *V1AlarmHistoryGet200Response) GetHistoryOk() (*[]V1AlarmHistoryGet200ResponseHistoryInner, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *V1AlarmHistoryGet200Response) SetHistory(v []V1AlarmHistoryGet200ResponseHistoryInner)`

SetHistory sets History field to given value.

### HasHistory

`func (o *V1AlarmHistoryGet200Response) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


