# V1FlowsFlowTablePost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorRef** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**FlowTable** | Pointer to [**[]V1FlowsFlowTablePost200ResponseFlowTableInner**](V1FlowsFlowTablePost200ResponseFlowTableInner.md) |  | [optional] 

## Methods

### NewV1FlowsFlowTablePost200Response

`func NewV1FlowsFlowTablePost200Response() *V1FlowsFlowTablePost200Response`

NewV1FlowsFlowTablePost200Response instantiates a new V1FlowsFlowTablePost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FlowsFlowTablePost200ResponseWithDefaults

`func NewV1FlowsFlowTablePost200ResponseWithDefaults() *V1FlowsFlowTablePost200Response`

NewV1FlowsFlowTablePost200ResponseWithDefaults instantiates a new V1FlowsFlowTablePost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorRef

`func (o *V1FlowsFlowTablePost200Response) GetCursorRef() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCursorRef returns the CursorRef field if non-nil, zero value otherwise.

### GetCursorRefOk

`func (o *V1FlowsFlowTablePost200Response) GetCursorRefOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCursorRefOk returns a tuple with the CursorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorRef

`func (o *V1FlowsFlowTablePost200Response) SetCursorRef(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCursorRef sets CursorRef field to given value.

### HasCursorRef

`func (o *V1FlowsFlowTablePost200Response) HasCursorRef() bool`

HasCursorRef returns a boolean if a field has been set.

### GetFlowTable

`func (o *V1FlowsFlowTablePost200Response) GetFlowTable() []V1FlowsFlowTablePost200ResponseFlowTableInner`

GetFlowTable returns the FlowTable field if non-nil, zero value otherwise.

### GetFlowTableOk

`func (o *V1FlowsFlowTablePost200Response) GetFlowTableOk() (*[]V1FlowsFlowTablePost200ResponseFlowTableInner, bool)`

GetFlowTableOk returns a tuple with the FlowTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowTable

`func (o *V1FlowsFlowTablePost200Response) SetFlowTable(v []V1FlowsFlowTablePost200ResponseFlowTableInner)`

SetFlowTable sets FlowTable field to given value.

### HasFlowTable

`func (o *V1FlowsFlowTablePost200Response) HasFlowTable() bool`

HasFlowTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


