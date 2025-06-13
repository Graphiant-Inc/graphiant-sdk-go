# V2AssuranceTopologyFlowsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 
**TopologyId** | Pointer to **int32** |  | [optional] 
**TopologyType** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssuranceTopologyFlowsPostRequest

`func NewV2AssuranceTopologyFlowsPostRequest() *V2AssuranceTopologyFlowsPostRequest`

NewV2AssuranceTopologyFlowsPostRequest instantiates a new V2AssuranceTopologyFlowsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceTopologyFlowsPostRequestWithDefaults

`func NewV2AssuranceTopologyFlowsPostRequestWithDefaults() *V2AssuranceTopologyFlowsPostRequest`

NewV2AssuranceTopologyFlowsPostRequestWithDefaults instantiates a new V2AssuranceTopologyFlowsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *V2AssuranceTopologyFlowsPostRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *V2AssuranceTopologyFlowsPostRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *V2AssuranceTopologyFlowsPostRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *V2AssuranceTopologyFlowsPostRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2AssuranceTopologyFlowsPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2AssuranceTopologyFlowsPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2AssuranceTopologyFlowsPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2AssuranceTopologyFlowsPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetTopologyId

`func (o *V2AssuranceTopologyFlowsPostRequest) GetTopologyId() int32`

GetTopologyId returns the TopologyId field if non-nil, zero value otherwise.

### GetTopologyIdOk

`func (o *V2AssuranceTopologyFlowsPostRequest) GetTopologyIdOk() (*int32, bool)`

GetTopologyIdOk returns a tuple with the TopologyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyId

`func (o *V2AssuranceTopologyFlowsPostRequest) SetTopologyId(v int32)`

SetTopologyId sets TopologyId field to given value.

### HasTopologyId

`func (o *V2AssuranceTopologyFlowsPostRequest) HasTopologyId() bool`

HasTopologyId returns a boolean if a field has been set.

### GetTopologyType

`func (o *V2AssuranceTopologyFlowsPostRequest) GetTopologyType() string`

GetTopologyType returns the TopologyType field if non-nil, zero value otherwise.

### GetTopologyTypeOk

`func (o *V2AssuranceTopologyFlowsPostRequest) GetTopologyTypeOk() (*string, bool)`

GetTopologyTypeOk returns a tuple with the TopologyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyType

`func (o *V2AssuranceTopologyFlowsPostRequest) SetTopologyType(v string)`

SetTopologyType sets TopologyType field to given value.

### HasTopologyType

`func (o *V2AssuranceTopologyFlowsPostRequest) HasTopologyType() bool`

HasTopologyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


