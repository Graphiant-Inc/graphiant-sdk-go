# V2AssuranceTopologyClientSummariesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**AppServerKey** | Pointer to **string** |  | [optional] 
**BucketId** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**V2AssuranceTopologyClientSessionsPostRequestFilter**](V2AssuranceTopologyClientSessionsPostRequestFilter.md) |  | [optional] 
**FlexAlgoId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 

## Methods

### NewV2AssuranceTopologyClientSummariesPostRequest

`func NewV2AssuranceTopologyClientSummariesPostRequest() *V2AssuranceTopologyClientSummariesPostRequest`

NewV2AssuranceTopologyClientSummariesPostRequest instantiates a new V2AssuranceTopologyClientSummariesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceTopologyClientSummariesPostRequestWithDefaults

`func NewV2AssuranceTopologyClientSummariesPostRequestWithDefaults() *V2AssuranceTopologyClientSummariesPostRequest`

NewV2AssuranceTopologyClientSummariesPostRequestWithDefaults instantiates a new V2AssuranceTopologyClientSummariesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *V2AssuranceTopologyClientSummariesPostRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *V2AssuranceTopologyClientSummariesPostRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppServerKey

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetAppServerKey() string`

GetAppServerKey returns the AppServerKey field if non-nil, zero value otherwise.

### GetAppServerKeyOk

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetAppServerKeyOk() (*string, bool)`

GetAppServerKeyOk returns a tuple with the AppServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerKey

`func (o *V2AssuranceTopologyClientSummariesPostRequest) SetAppServerKey(v string)`

SetAppServerKey sets AppServerKey field to given value.

### HasAppServerKey

`func (o *V2AssuranceTopologyClientSummariesPostRequest) HasAppServerKey() bool`

HasAppServerKey returns a boolean if a field has been set.

### GetBucketId

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *V2AssuranceTopologyClientSummariesPostRequest) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *V2AssuranceTopologyClientSummariesPostRequest) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetFilter

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetFilter() V2AssuranceTopologyClientSessionsPostRequestFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetFilterOk() (*V2AssuranceTopologyClientSessionsPostRequestFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V2AssuranceTopologyClientSummariesPostRequest) SetFilter(v V2AssuranceTopologyClientSessionsPostRequestFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *V2AssuranceTopologyClientSummariesPostRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFlexAlgoId

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetFlexAlgoId() int64`

GetFlexAlgoId returns the FlexAlgoId field if non-nil, zero value otherwise.

### GetFlexAlgoIdOk

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetFlexAlgoIdOk() (*int64, bool)`

GetFlexAlgoIdOk returns a tuple with the FlexAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAlgoId

`func (o *V2AssuranceTopologyClientSummariesPostRequest) SetFlexAlgoId(v int64)`

SetFlexAlgoId sets FlexAlgoId field to given value.

### HasFlexAlgoId

`func (o *V2AssuranceTopologyClientSummariesPostRequest) HasFlexAlgoId() bool`

HasFlexAlgoId returns a boolean if a field has been set.

### GetSiteId

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *V2AssuranceTopologyClientSummariesPostRequest) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *V2AssuranceTopologyClientSummariesPostRequest) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2AssuranceTopologyClientSummariesPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2AssuranceTopologyClientSummariesPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2AssuranceTopologyClientSummariesPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


