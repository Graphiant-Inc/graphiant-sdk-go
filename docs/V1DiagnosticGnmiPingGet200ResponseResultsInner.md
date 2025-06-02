# V1DiagnosticGnmiPingGet200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Rtt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**TtDeviceId** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1DiagnosticGnmiPingGet200ResponseResultsInner

`func NewV1DiagnosticGnmiPingGet200ResponseResultsInner() *V1DiagnosticGnmiPingGet200ResponseResultsInner`

NewV1DiagnosticGnmiPingGet200ResponseResultsInner instantiates a new V1DiagnosticGnmiPingGet200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DiagnosticGnmiPingGet200ResponseResultsInnerWithDefaults

`func NewV1DiagnosticGnmiPingGet200ResponseResultsInnerWithDefaults() *V1DiagnosticGnmiPingGet200ResponseResultsInner`

NewV1DiagnosticGnmiPingGet200ResponseResultsInnerWithDefaults instantiates a new V1DiagnosticGnmiPingGet200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetError

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRtt

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) GetRtt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetRtt returns the Rtt field if non-nil, zero value otherwise.

### GetRttOk

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) GetRttOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetRttOk returns a tuple with the Rtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtt

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) SetRtt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetRtt sets Rtt field to given value.

### HasRtt

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) HasRtt() bool`

HasRtt returns a boolean if a field has been set.

### GetTtDeviceId

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) GetTtDeviceId() int64`

GetTtDeviceId returns the TtDeviceId field if non-nil, zero value otherwise.

### GetTtDeviceIdOk

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) GetTtDeviceIdOk() (*int64, bool)`

GetTtDeviceIdOk returns a tuple with the TtDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtDeviceId

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) SetTtDeviceId(v int64)`

SetTtDeviceId sets TtDeviceId field to given value.

### HasTtDeviceId

`func (o *V1DiagnosticGnmiPingGet200ResponseResultsInner) HasTtDeviceId() bool`

HasTtDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


