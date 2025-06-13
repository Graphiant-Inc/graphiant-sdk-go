# V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selectors** | Pointer to [**[]V1MonitoringCircuitsBandwidthPostRequestSelectorsInner**](V1MonitoringCircuitsBandwidthPostRequestSelectorsInner.md) |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV2MonitoringSiteCircuitsBandwidthSiteIdPostRequest

`func NewV2MonitoringSiteCircuitsBandwidthSiteIdPostRequest() *V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest`

NewV2MonitoringSiteCircuitsBandwidthSiteIdPostRequest instantiates a new V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringSiteCircuitsBandwidthSiteIdPostRequestWithDefaults

`func NewV2MonitoringSiteCircuitsBandwidthSiteIdPostRequestWithDefaults() *V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest`

NewV2MonitoringSiteCircuitsBandwidthSiteIdPostRequestWithDefaults instantiates a new V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectors

`func (o *V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest) GetSelectors() []V1MonitoringCircuitsBandwidthPostRequestSelectorsInner`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest) GetSelectorsOk() (*[]V1MonitoringCircuitsBandwidthPostRequestSelectorsInner, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest) SetSelectors(v []V1MonitoringCircuitsBandwidthPostRequestSelectorsInner)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2MonitoringSiteCircuitsBandwidthSiteIdPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


