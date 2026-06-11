# V2AssuranceTopologySiteSummariesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**AppServerKey** | Pointer to **string** |  | [optional] 
**BucketId** | Pointer to **string** |  | [optional] 
**ExchangeServiceId** | Pointer to **int64** |  | [optional] 
**Filter** | Pointer to [**AssuranceTopologyFilter**](AssuranceTopologyFilter.md) |  | [optional] 
**FlexAlgoId** | Pointer to **int64** |  | [optional] 
**TimeWindow** | Pointer to [**AssuranceTimeWindow**](AssuranceTimeWindow.md) |  | [optional] 

## Methods

### NewV2AssuranceTopologySiteSummariesPostRequest

`func NewV2AssuranceTopologySiteSummariesPostRequest() *V2AssuranceTopologySiteSummariesPostRequest`

NewV2AssuranceTopologySiteSummariesPostRequest instantiates a new V2AssuranceTopologySiteSummariesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceTopologySiteSummariesPostRequestWithDefaults

`func NewV2AssuranceTopologySiteSummariesPostRequestWithDefaults() *V2AssuranceTopologySiteSummariesPostRequest`

NewV2AssuranceTopologySiteSummariesPostRequestWithDefaults instantiates a new V2AssuranceTopologySiteSummariesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *V2AssuranceTopologySiteSummariesPostRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *V2AssuranceTopologySiteSummariesPostRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppServerKey

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetAppServerKey() string`

GetAppServerKey returns the AppServerKey field if non-nil, zero value otherwise.

### GetAppServerKeyOk

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetAppServerKeyOk() (*string, bool)`

GetAppServerKeyOk returns a tuple with the AppServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerKey

`func (o *V2AssuranceTopologySiteSummariesPostRequest) SetAppServerKey(v string)`

SetAppServerKey sets AppServerKey field to given value.

### HasAppServerKey

`func (o *V2AssuranceTopologySiteSummariesPostRequest) HasAppServerKey() bool`

HasAppServerKey returns a boolean if a field has been set.

### GetBucketId

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *V2AssuranceTopologySiteSummariesPostRequest) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *V2AssuranceTopologySiteSummariesPostRequest) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetExchangeServiceId

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetExchangeServiceId() int64`

GetExchangeServiceId returns the ExchangeServiceId field if non-nil, zero value otherwise.

### GetExchangeServiceIdOk

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetExchangeServiceIdOk() (*int64, bool)`

GetExchangeServiceIdOk returns a tuple with the ExchangeServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeServiceId

`func (o *V2AssuranceTopologySiteSummariesPostRequest) SetExchangeServiceId(v int64)`

SetExchangeServiceId sets ExchangeServiceId field to given value.

### HasExchangeServiceId

`func (o *V2AssuranceTopologySiteSummariesPostRequest) HasExchangeServiceId() bool`

HasExchangeServiceId returns a boolean if a field has been set.

### GetFilter

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetFilter() AssuranceTopologyFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetFilterOk() (*AssuranceTopologyFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V2AssuranceTopologySiteSummariesPostRequest) SetFilter(v AssuranceTopologyFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *V2AssuranceTopologySiteSummariesPostRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFlexAlgoId

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetFlexAlgoId() int64`

GetFlexAlgoId returns the FlexAlgoId field if non-nil, zero value otherwise.

### GetFlexAlgoIdOk

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetFlexAlgoIdOk() (*int64, bool)`

GetFlexAlgoIdOk returns a tuple with the FlexAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAlgoId

`func (o *V2AssuranceTopologySiteSummariesPostRequest) SetFlexAlgoId(v int64)`

SetFlexAlgoId sets FlexAlgoId field to given value.

### HasFlexAlgoId

`func (o *V2AssuranceTopologySiteSummariesPostRequest) HasFlexAlgoId() bool`

HasFlexAlgoId returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetTimeWindow() AssuranceTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2AssuranceTopologySiteSummariesPostRequest) GetTimeWindowOk() (*AssuranceTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2AssuranceTopologySiteSummariesPostRequest) SetTimeWindow(v AssuranceTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2AssuranceTopologySiteSummariesPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


