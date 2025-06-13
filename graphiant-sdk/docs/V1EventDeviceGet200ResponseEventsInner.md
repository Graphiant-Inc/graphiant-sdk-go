# V1EventDeviceGet200ResponseEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**EventData** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **int64** |  | [optional] 
**EventTime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1EventDeviceGet200ResponseEventsInner

`func NewV1EventDeviceGet200ResponseEventsInner() *V1EventDeviceGet200ResponseEventsInner`

NewV1EventDeviceGet200ResponseEventsInner instantiates a new V1EventDeviceGet200ResponseEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EventDeviceGet200ResponseEventsInnerWithDefaults

`func NewV1EventDeviceGet200ResponseEventsInnerWithDefaults() *V1EventDeviceGet200ResponseEventsInner`

NewV1EventDeviceGet200ResponseEventsInnerWithDefaults instantiates a new V1EventDeviceGet200ResponseEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *V1EventDeviceGet200ResponseEventsInner) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *V1EventDeviceGet200ResponseEventsInner) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *V1EventDeviceGet200ResponseEventsInner) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *V1EventDeviceGet200ResponseEventsInner) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1EventDeviceGet200ResponseEventsInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1EventDeviceGet200ResponseEventsInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1EventDeviceGet200ResponseEventsInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1EventDeviceGet200ResponseEventsInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *V1EventDeviceGet200ResponseEventsInner) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *V1EventDeviceGet200ResponseEventsInner) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *V1EventDeviceGet200ResponseEventsInner) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *V1EventDeviceGet200ResponseEventsInner) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetEventData

`func (o *V1EventDeviceGet200ResponseEventsInner) GetEventData() string`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *V1EventDeviceGet200ResponseEventsInner) GetEventDataOk() (*string, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *V1EventDeviceGet200ResponseEventsInner) SetEventData(v string)`

SetEventData sets EventData field to given value.

### HasEventData

`func (o *V1EventDeviceGet200ResponseEventsInner) HasEventData() bool`

HasEventData returns a boolean if a field has been set.

### GetEventId

`func (o *V1EventDeviceGet200ResponseEventsInner) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *V1EventDeviceGet200ResponseEventsInner) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *V1EventDeviceGet200ResponseEventsInner) SetEventId(v int64)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *V1EventDeviceGet200ResponseEventsInner) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *V1EventDeviceGet200ResponseEventsInner) GetEventTime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *V1EventDeviceGet200ResponseEventsInner) GetEventTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *V1EventDeviceGet200ResponseEventsInner) SetEventTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *V1EventDeviceGet200ResponseEventsInner) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetHostname

`func (o *V1EventDeviceGet200ResponseEventsInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V1EventDeviceGet200ResponseEventsInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V1EventDeviceGet200ResponseEventsInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V1EventDeviceGet200ResponseEventsInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetSeverity

`func (o *V1EventDeviceGet200ResponseEventsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *V1EventDeviceGet200ResponseEventsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *V1EventDeviceGet200ResponseEventsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *V1EventDeviceGet200ResponseEventsInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetType

`func (o *V1EventDeviceGet200ResponseEventsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1EventDeviceGet200ResponseEventsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1EventDeviceGet200ResponseEventsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1EventDeviceGet200ResponseEventsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


