# V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**DisconnectedReason** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner

`func NewV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner() *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner`

NewV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner instantiates a new V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInnerWithDefaults

`func NewV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInnerWithDefaults() *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner`

NewV2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInnerWithDefaults instantiates a new V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetCreatedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetCreatedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetCreatedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisconnectedReason

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetDisconnectedReason() string`

GetDisconnectedReason returns the DisconnectedReason field if non-nil, zero value otherwise.

### GetDisconnectedReasonOk

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetDisconnectedReasonOk() (*string, bool)`

GetDisconnectedReasonOk returns a tuple with the DisconnectedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedReason

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetDisconnectedReason(v string)`

SetDisconnectedReason sets DisconnectedReason field to given value.

### HasDisconnectedReason

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasDisconnectedReason() bool`

HasDisconnectedReason returns a boolean if a field has been set.

### GetHostname

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSiteName

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetStatus

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V2MonitoringExtranetEdgeStatusGet200ResponseEdgeStatusesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


