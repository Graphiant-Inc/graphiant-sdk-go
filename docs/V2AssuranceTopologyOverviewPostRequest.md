# V2AssuranceTopologyOverviewPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**AppServerKey** | Pointer to **string** |  | [optional] 
**BucketId** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**V2AssuranceTopologyClientSessionsPostRequestFilter**](V2AssuranceTopologyClientSessionsPostRequestFilter.md) |  | [optional] 
**FlexAlgoId** | Pointer to **int64** |  | [optional] 
**SliderTimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 
**TopologyTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**TopologyType** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssuranceTopologyOverviewPostRequest

`func NewV2AssuranceTopologyOverviewPostRequest() *V2AssuranceTopologyOverviewPostRequest`

NewV2AssuranceTopologyOverviewPostRequest instantiates a new V2AssuranceTopologyOverviewPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceTopologyOverviewPostRequestWithDefaults

`func NewV2AssuranceTopologyOverviewPostRequestWithDefaults() *V2AssuranceTopologyOverviewPostRequest`

NewV2AssuranceTopologyOverviewPostRequestWithDefaults instantiates a new V2AssuranceTopologyOverviewPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *V2AssuranceTopologyOverviewPostRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *V2AssuranceTopologyOverviewPostRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *V2AssuranceTopologyOverviewPostRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *V2AssuranceTopologyOverviewPostRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppServerKey

`func (o *V2AssuranceTopologyOverviewPostRequest) GetAppServerKey() string`

GetAppServerKey returns the AppServerKey field if non-nil, zero value otherwise.

### GetAppServerKeyOk

`func (o *V2AssuranceTopologyOverviewPostRequest) GetAppServerKeyOk() (*string, bool)`

GetAppServerKeyOk returns a tuple with the AppServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerKey

`func (o *V2AssuranceTopologyOverviewPostRequest) SetAppServerKey(v string)`

SetAppServerKey sets AppServerKey field to given value.

### HasAppServerKey

`func (o *V2AssuranceTopologyOverviewPostRequest) HasAppServerKey() bool`

HasAppServerKey returns a boolean if a field has been set.

### GetBucketId

`func (o *V2AssuranceTopologyOverviewPostRequest) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *V2AssuranceTopologyOverviewPostRequest) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *V2AssuranceTopologyOverviewPostRequest) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *V2AssuranceTopologyOverviewPostRequest) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetFilter

`func (o *V2AssuranceTopologyOverviewPostRequest) GetFilter() V2AssuranceTopologyClientSessionsPostRequestFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V2AssuranceTopologyOverviewPostRequest) GetFilterOk() (*V2AssuranceTopologyClientSessionsPostRequestFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V2AssuranceTopologyOverviewPostRequest) SetFilter(v V2AssuranceTopologyClientSessionsPostRequestFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *V2AssuranceTopologyOverviewPostRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFlexAlgoId

`func (o *V2AssuranceTopologyOverviewPostRequest) GetFlexAlgoId() int64`

GetFlexAlgoId returns the FlexAlgoId field if non-nil, zero value otherwise.

### GetFlexAlgoIdOk

`func (o *V2AssuranceTopologyOverviewPostRequest) GetFlexAlgoIdOk() (*int64, bool)`

GetFlexAlgoIdOk returns a tuple with the FlexAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAlgoId

`func (o *V2AssuranceTopologyOverviewPostRequest) SetFlexAlgoId(v int64)`

SetFlexAlgoId sets FlexAlgoId field to given value.

### HasFlexAlgoId

`func (o *V2AssuranceTopologyOverviewPostRequest) HasFlexAlgoId() bool`

HasFlexAlgoId returns a boolean if a field has been set.

### GetSliderTimeWindow

`func (o *V2AssuranceTopologyOverviewPostRequest) GetSliderTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetSliderTimeWindow returns the SliderTimeWindow field if non-nil, zero value otherwise.

### GetSliderTimeWindowOk

`func (o *V2AssuranceTopologyOverviewPostRequest) GetSliderTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetSliderTimeWindowOk returns a tuple with the SliderTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliderTimeWindow

`func (o *V2AssuranceTopologyOverviewPostRequest) SetSliderTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetSliderTimeWindow sets SliderTimeWindow field to given value.

### HasSliderTimeWindow

`func (o *V2AssuranceTopologyOverviewPostRequest) HasSliderTimeWindow() bool`

HasSliderTimeWindow returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2AssuranceTopologyOverviewPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2AssuranceTopologyOverviewPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2AssuranceTopologyOverviewPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2AssuranceTopologyOverviewPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetTopologyTs

`func (o *V2AssuranceTopologyOverviewPostRequest) GetTopologyTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetTopologyTs returns the TopologyTs field if non-nil, zero value otherwise.

### GetTopologyTsOk

`func (o *V2AssuranceTopologyOverviewPostRequest) GetTopologyTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetTopologyTsOk returns a tuple with the TopologyTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyTs

`func (o *V2AssuranceTopologyOverviewPostRequest) SetTopologyTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetTopologyTs sets TopologyTs field to given value.

### HasTopologyTs

`func (o *V2AssuranceTopologyOverviewPostRequest) HasTopologyTs() bool`

HasTopologyTs returns a boolean if a field has been set.

### GetTopologyType

`func (o *V2AssuranceTopologyOverviewPostRequest) GetTopologyType() string`

GetTopologyType returns the TopologyType field if non-nil, zero value otherwise.

### GetTopologyTypeOk

`func (o *V2AssuranceTopologyOverviewPostRequest) GetTopologyTypeOk() (*string, bool)`

GetTopologyTypeOk returns a tuple with the TopologyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyType

`func (o *V2AssuranceTopologyOverviewPostRequest) SetTopologyType(v string)`

SetTopologyType sets TopologyType field to given value.

### HasTopologyType

`func (o *V2AssuranceTopologyOverviewPostRequest) HasTopologyType() bool`

HasTopologyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


