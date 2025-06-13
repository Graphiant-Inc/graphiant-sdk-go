# V2MonitoringCircuitsSummaryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV2MonitoringCircuitsSummaryPostRequest

`func NewV2MonitoringCircuitsSummaryPostRequest() *V2MonitoringCircuitsSummaryPostRequest`

NewV2MonitoringCircuitsSummaryPostRequest instantiates a new V2MonitoringCircuitsSummaryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringCircuitsSummaryPostRequestWithDefaults

`func NewV2MonitoringCircuitsSummaryPostRequestWithDefaults() *V2MonitoringCircuitsSummaryPostRequest`

NewV2MonitoringCircuitsSummaryPostRequestWithDefaults instantiates a new V2MonitoringCircuitsSummaryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V2MonitoringCircuitsSummaryPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V2MonitoringCircuitsSummaryPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V2MonitoringCircuitsSummaryPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V2MonitoringCircuitsSummaryPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2MonitoringCircuitsSummaryPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2MonitoringCircuitsSummaryPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2MonitoringCircuitsSummaryPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2MonitoringCircuitsSummaryPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


