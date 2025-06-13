# V1FlowsFlowTablePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **int32** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**CursorRef** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**IsDia** | Pointer to **bool** |  | [optional] 
**NumFlowRecords** | Pointer to **int32** |  | [optional] 
**Selector** | Pointer to [**V1FlowsFlowTablePostRequestSelector**](V1FlowsFlowTablePostRequestSelector.md) |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV1FlowsFlowTablePostRequest

`func NewV1FlowsFlowTablePostRequest() *V1FlowsFlowTablePostRequest`

NewV1FlowsFlowTablePostRequest instantiates a new V1FlowsFlowTablePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FlowsFlowTablePostRequestWithDefaults

`func NewV1FlowsFlowTablePostRequestWithDefaults() *V1FlowsFlowTablePostRequest`

NewV1FlowsFlowTablePostRequestWithDefaults instantiates a new V1FlowsFlowTablePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *V1FlowsFlowTablePostRequest) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *V1FlowsFlowTablePostRequest) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *V1FlowsFlowTablePostRequest) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *V1FlowsFlowTablePostRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppName

`func (o *V1FlowsFlowTablePostRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *V1FlowsFlowTablePostRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *V1FlowsFlowTablePostRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *V1FlowsFlowTablePostRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetCursorRef

`func (o *V1FlowsFlowTablePostRequest) GetCursorRef() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCursorRef returns the CursorRef field if non-nil, zero value otherwise.

### GetCursorRefOk

`func (o *V1FlowsFlowTablePostRequest) GetCursorRefOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCursorRefOk returns a tuple with the CursorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorRef

`func (o *V1FlowsFlowTablePostRequest) SetCursorRef(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCursorRef sets CursorRef field to given value.

### HasCursorRef

`func (o *V1FlowsFlowTablePostRequest) HasCursorRef() bool`

HasCursorRef returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1FlowsFlowTablePostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1FlowsFlowTablePostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1FlowsFlowTablePostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1FlowsFlowTablePostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetIsDia

`func (o *V1FlowsFlowTablePostRequest) GetIsDia() bool`

GetIsDia returns the IsDia field if non-nil, zero value otherwise.

### GetIsDiaOk

`func (o *V1FlowsFlowTablePostRequest) GetIsDiaOk() (*bool, bool)`

GetIsDiaOk returns a tuple with the IsDia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDia

`func (o *V1FlowsFlowTablePostRequest) SetIsDia(v bool)`

SetIsDia sets IsDia field to given value.

### HasIsDia

`func (o *V1FlowsFlowTablePostRequest) HasIsDia() bool`

HasIsDia returns a boolean if a field has been set.

### GetNumFlowRecords

`func (o *V1FlowsFlowTablePostRequest) GetNumFlowRecords() int32`

GetNumFlowRecords returns the NumFlowRecords field if non-nil, zero value otherwise.

### GetNumFlowRecordsOk

`func (o *V1FlowsFlowTablePostRequest) GetNumFlowRecordsOk() (*int32, bool)`

GetNumFlowRecordsOk returns a tuple with the NumFlowRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFlowRecords

`func (o *V1FlowsFlowTablePostRequest) SetNumFlowRecords(v int32)`

SetNumFlowRecords sets NumFlowRecords field to given value.

### HasNumFlowRecords

`func (o *V1FlowsFlowTablePostRequest) HasNumFlowRecords() bool`

HasNumFlowRecords returns a boolean if a field has been set.

### GetSelector

`func (o *V1FlowsFlowTablePostRequest) GetSelector() V1FlowsFlowTablePostRequestSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1FlowsFlowTablePostRequest) GetSelectorOk() (*V1FlowsFlowTablePostRequestSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1FlowsFlowTablePostRequest) SetSelector(v V1FlowsFlowTablePostRequestSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1FlowsFlowTablePostRequest) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V1FlowsFlowTablePostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V1FlowsFlowTablePostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V1FlowsFlowTablePostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V1FlowsFlowTablePostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


