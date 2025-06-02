# V2AssuranceFlowSummaryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** |  | [optional] 
**ServerIp** | Pointer to **string** |  | [optional] 
**ServerPort** | Pointer to **int32** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV2AssuranceFlowSummaryPostRequest

`func NewV2AssuranceFlowSummaryPostRequest() *V2AssuranceFlowSummaryPostRequest`

NewV2AssuranceFlowSummaryPostRequest instantiates a new V2AssuranceFlowSummaryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceFlowSummaryPostRequestWithDefaults

`func NewV2AssuranceFlowSummaryPostRequestWithDefaults() *V2AssuranceFlowSummaryPostRequest`

NewV2AssuranceFlowSummaryPostRequestWithDefaults instantiates a new V2AssuranceFlowSummaryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *V2AssuranceFlowSummaryPostRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *V2AssuranceFlowSummaryPostRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *V2AssuranceFlowSummaryPostRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *V2AssuranceFlowSummaryPostRequest) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetServerIp

`func (o *V2AssuranceFlowSummaryPostRequest) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *V2AssuranceFlowSummaryPostRequest) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *V2AssuranceFlowSummaryPostRequest) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *V2AssuranceFlowSummaryPostRequest) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetServerPort

`func (o *V2AssuranceFlowSummaryPostRequest) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *V2AssuranceFlowSummaryPostRequest) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *V2AssuranceFlowSummaryPostRequest) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *V2AssuranceFlowSummaryPostRequest) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2AssuranceFlowSummaryPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2AssuranceFlowSummaryPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2AssuranceFlowSummaryPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2AssuranceFlowSummaryPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


