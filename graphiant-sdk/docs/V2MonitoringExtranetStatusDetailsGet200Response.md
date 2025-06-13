# V2MonitoringExtranetStatusDetailsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeStatuses** | Pointer to [**[]V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner**](V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner.md) |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SiteStatus** | Pointer to [**V2MonitoringExtranetServiceStatusDetailsGet200ResponseStatusesInner**](V2MonitoringExtranetServiceStatusDetailsGet200ResponseStatusesInner.md) |  | [optional] 

## Methods

### NewV2MonitoringExtranetStatusDetailsGet200Response

`func NewV2MonitoringExtranetStatusDetailsGet200Response() *V2MonitoringExtranetStatusDetailsGet200Response`

NewV2MonitoringExtranetStatusDetailsGet200Response instantiates a new V2MonitoringExtranetStatusDetailsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringExtranetStatusDetailsGet200ResponseWithDefaults

`func NewV2MonitoringExtranetStatusDetailsGet200ResponseWithDefaults() *V2MonitoringExtranetStatusDetailsGet200Response`

NewV2MonitoringExtranetStatusDetailsGet200ResponseWithDefaults instantiates a new V2MonitoringExtranetStatusDetailsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeStatuses

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) GetEdgeStatuses() []V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner`

GetEdgeStatuses returns the EdgeStatuses field if non-nil, zero value otherwise.

### GetEdgeStatusesOk

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) GetEdgeStatusesOk() (*[]V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner, bool)`

GetEdgeStatusesOk returns a tuple with the EdgeStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeStatuses

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) SetEdgeStatuses(v []V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner)`

SetEdgeStatuses sets EdgeStatuses field to given value.

### HasEdgeStatuses

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) HasEdgeStatuses() bool`

HasEdgeStatuses returns a boolean if a field has been set.

### GetLocation

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRegion

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSiteStatus

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) GetSiteStatus() V2MonitoringExtranetServiceStatusDetailsGet200ResponseStatusesInner`

GetSiteStatus returns the SiteStatus field if non-nil, zero value otherwise.

### GetSiteStatusOk

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) GetSiteStatusOk() (*V2MonitoringExtranetServiceStatusDetailsGet200ResponseStatusesInner, bool)`

GetSiteStatusOk returns a tuple with the SiteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteStatus

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) SetSiteStatus(v V2MonitoringExtranetServiceStatusDetailsGet200ResponseStatusesInner)`

SetSiteStatus sets SiteStatus field to given value.

### HasSiteStatus

`func (o *V2MonitoringExtranetStatusDetailsGet200Response) HasSiteStatus() bool`

HasSiteStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


