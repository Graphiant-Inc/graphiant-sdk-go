# V2MonitoringExtranetStatusDetailsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeStatuses** | Pointer to [**[]StatsmonExtranetEdgeStatus**](StatsmonExtranetEdgeStatus.md) |  | [optional] 
**Location** | Pointer to **string** | the location of the site | [optional] 
**Region** | Pointer to **string** | the region of the site | [optional] 
**SiteStatus** | Pointer to [**StatsmonExtranetSiteStatus**](StatsmonExtranetSiteStatus.md) |  | [optional] 

## Methods

### NewV2MonitoringExtranetStatusDetailsGetResponse

`func NewV2MonitoringExtranetStatusDetailsGetResponse() *V2MonitoringExtranetStatusDetailsGetResponse`

NewV2MonitoringExtranetStatusDetailsGetResponse instantiates a new V2MonitoringExtranetStatusDetailsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringExtranetStatusDetailsGetResponseWithDefaults

`func NewV2MonitoringExtranetStatusDetailsGetResponseWithDefaults() *V2MonitoringExtranetStatusDetailsGetResponse`

NewV2MonitoringExtranetStatusDetailsGetResponseWithDefaults instantiates a new V2MonitoringExtranetStatusDetailsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeStatuses

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) GetEdgeStatuses() []StatsmonExtranetEdgeStatus`

GetEdgeStatuses returns the EdgeStatuses field if non-nil, zero value otherwise.

### GetEdgeStatusesOk

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) GetEdgeStatusesOk() (*[]StatsmonExtranetEdgeStatus, bool)`

GetEdgeStatusesOk returns a tuple with the EdgeStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeStatuses

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) SetEdgeStatuses(v []StatsmonExtranetEdgeStatus)`

SetEdgeStatuses sets EdgeStatuses field to given value.

### HasEdgeStatuses

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) HasEdgeStatuses() bool`

HasEdgeStatuses returns a boolean if a field has been set.

### GetLocation

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRegion

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSiteStatus

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) GetSiteStatus() StatsmonExtranetSiteStatus`

GetSiteStatus returns the SiteStatus field if non-nil, zero value otherwise.

### GetSiteStatusOk

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) GetSiteStatusOk() (*StatsmonExtranetSiteStatus, bool)`

GetSiteStatusOk returns a tuple with the SiteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteStatus

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) SetSiteStatus(v StatsmonExtranetSiteStatus)`

SetSiteStatus sets SiteStatus field to given value.

### HasSiteStatus

`func (o *V2MonitoringExtranetStatusDetailsGetResponse) HasSiteStatus() bool`

HasSiteStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


