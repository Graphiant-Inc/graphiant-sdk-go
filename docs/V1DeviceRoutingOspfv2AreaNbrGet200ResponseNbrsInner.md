# V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **int32** |  | [optional] 
**DeadTimer** | Pointer to **int32** |  | [optional] 
**LastStateChange** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**RouterId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner

`func NewV1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner() *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner`

NewV1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner instantiates a new V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInnerWithDefaults

`func NewV1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInnerWithDefaults() *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner`

NewV1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInnerWithDefaults instantiates a new V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCost

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDeadTimer

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetDeadTimer() int32`

GetDeadTimer returns the DeadTimer field if non-nil, zero value otherwise.

### GetDeadTimerOk

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetDeadTimerOk() (*int32, bool)`

GetDeadTimerOk returns a tuple with the DeadTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadTimer

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) SetDeadTimer(v int32)`

SetDeadTimer sets DeadTimer field to given value.

### HasDeadTimer

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) HasDeadTimer() bool`

HasDeadTimer returns a boolean if a field has been set.

### GetLastStateChange

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetLastStateChange() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetLastStateChange returns the LastStateChange field if non-nil, zero value otherwise.

### GetLastStateChangeOk

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetLastStateChangeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetLastStateChangeOk returns a tuple with the LastStateChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateChange

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) SetLastStateChange(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetLastStateChange sets LastStateChange field to given value.

### HasLastStateChange

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) HasLastStateChange() bool`

HasLastStateChange returns a boolean if a field has been set.

### GetRouterId

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetState

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V1DeviceRoutingOspfv2AreaNbrGet200ResponseNbrsInner) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


