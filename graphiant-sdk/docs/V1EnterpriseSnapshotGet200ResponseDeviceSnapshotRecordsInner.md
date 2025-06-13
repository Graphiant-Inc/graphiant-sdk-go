# V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 
**SnapshotCount** | Pointer to **int32** |  | [optional] 
**Snapshots** | Pointer to [**[]V1DeviceSnapshotGet200ResponseFirstSnapshot**](V1DeviceSnapshotGet200ResponseFirstSnapshot.md) |  | [optional] 
**Uptime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner

`func NewV1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner() *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner`

NewV1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner instantiates a new V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInnerWithDefaults

`func NewV1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInnerWithDefaults() *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner`

NewV1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInnerWithDefaults instantiates a new V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetHostname

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetRegion

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSite

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSnapshotCount

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetSnapshotCount() int32`

GetSnapshotCount returns the SnapshotCount field if non-nil, zero value otherwise.

### GetSnapshotCountOk

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetSnapshotCountOk() (*int32, bool)`

GetSnapshotCountOk returns a tuple with the SnapshotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCount

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) SetSnapshotCount(v int32)`

SetSnapshotCount sets SnapshotCount field to given value.

### HasSnapshotCount

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) HasSnapshotCount() bool`

HasSnapshotCount returns a boolean if a field has been set.

### GetSnapshots

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetSnapshots() []V1DeviceSnapshotGet200ResponseFirstSnapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetSnapshotsOk() (*[]V1DeviceSnapshotGet200ResponseFirstSnapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) SetSnapshots(v []V1DeviceSnapshotGet200ResponseFirstSnapshot)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetUptime

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetUptime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) GetUptimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) SetUptime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *V1EnterpriseSnapshotGet200ResponseDeviceSnapshotRecordsInner) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


