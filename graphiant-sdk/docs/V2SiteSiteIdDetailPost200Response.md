# V2SiteSiteIdDetailPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to [**V2SiteSiteIdDetailPost200ResponseSite**](V2SiteSiteIdDetailPost200ResponseSite.md) |  | [optional] 
**Snapshots** | Pointer to [**[]V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV2SiteSiteIdDetailPost200Response

`func NewV2SiteSiteIdDetailPost200Response() *V2SiteSiteIdDetailPost200Response`

NewV2SiteSiteIdDetailPost200Response instantiates a new V2SiteSiteIdDetailPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2SiteSiteIdDetailPost200ResponseWithDefaults

`func NewV2SiteSiteIdDetailPost200ResponseWithDefaults() *V2SiteSiteIdDetailPost200Response`

NewV2SiteSiteIdDetailPost200ResponseWithDefaults instantiates a new V2SiteSiteIdDetailPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *V2SiteSiteIdDetailPost200Response) GetSite() V2SiteSiteIdDetailPost200ResponseSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *V2SiteSiteIdDetailPost200Response) GetSiteOk() (*V2SiteSiteIdDetailPost200ResponseSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *V2SiteSiteIdDetailPost200Response) SetSite(v V2SiteSiteIdDetailPost200ResponseSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *V2SiteSiteIdDetailPost200Response) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSnapshots

`func (o *V2SiteSiteIdDetailPost200Response) GetSnapshots() []V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *V2SiteSiteIdDetailPost200Response) GetSnapshotsOk() (*[]V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *V2SiteSiteIdDetailPost200Response) SetSnapshots(v []V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *V2SiteSiteIdDetailPost200Response) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


