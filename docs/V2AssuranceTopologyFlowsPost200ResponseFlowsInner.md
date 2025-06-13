# V2AssuranceTopologyFlowsPost200ResponseFlowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**ClientIp** | Pointer to **string** |  | [optional] 
**ClientSite** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpointSite**](V2AssuranceFlowSummaryPost200ResponseClientEndpointSite.md) |  | [optional] 
**FirstSeenTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**LanSegment** | Pointer to **string** |  | [optional] 
**LastSeenTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**ServerIp** | Pointer to **string** |  | [optional] 
**ServerPort** | Pointer to **int32** |  | [optional] 
**ServerSite** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpointSite**](V2AssuranceFlowSummaryPost200ResponseClientEndpointSite.md) |  | [optional] 

## Methods

### NewV2AssuranceTopologyFlowsPost200ResponseFlowsInner

`func NewV2AssuranceTopologyFlowsPost200ResponseFlowsInner() *V2AssuranceTopologyFlowsPost200ResponseFlowsInner`

NewV2AssuranceTopologyFlowsPost200ResponseFlowsInner instantiates a new V2AssuranceTopologyFlowsPost200ResponseFlowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceTopologyFlowsPost200ResponseFlowsInnerWithDefaults

`func NewV2AssuranceTopologyFlowsPost200ResponseFlowsInnerWithDefaults() *V2AssuranceTopologyFlowsPost200ResponseFlowsInner`

NewV2AssuranceTopologyFlowsPost200ResponseFlowsInnerWithDefaults instantiates a new V2AssuranceTopologyFlowsPost200ResponseFlowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetClientIp

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientSite

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetClientSite() V2AssuranceFlowSummaryPost200ResponseClientEndpointSite`

GetClientSite returns the ClientSite field if non-nil, zero value otherwise.

### GetClientSiteOk

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetClientSiteOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpointSite, bool)`

GetClientSiteOk returns a tuple with the ClientSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSite

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) SetClientSite(v V2AssuranceFlowSummaryPost200ResponseClientEndpointSite)`

SetClientSite sets ClientSite field to given value.

### HasClientSite

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) HasClientSite() bool`

HasClientSite returns a boolean if a field has been set.

### GetFirstSeenTs

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetFirstSeenTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetFirstSeenTs returns the FirstSeenTs field if non-nil, zero value otherwise.

### GetFirstSeenTsOk

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetFirstSeenTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetFirstSeenTsOk returns a tuple with the FirstSeenTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTs

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) SetFirstSeenTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetFirstSeenTs sets FirstSeenTs field to given value.

### HasFirstSeenTs

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) HasFirstSeenTs() bool`

HasFirstSeenTs returns a boolean if a field has been set.

### GetLanSegment

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetLanSegment() string`

GetLanSegment returns the LanSegment field if non-nil, zero value otherwise.

### GetLanSegmentOk

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetLanSegmentOk() (*string, bool)`

GetLanSegmentOk returns a tuple with the LanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegment

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) SetLanSegment(v string)`

SetLanSegment sets LanSegment field to given value.

### HasLanSegment

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) HasLanSegment() bool`

HasLanSegment returns a boolean if a field has been set.

### GetLastSeenTs

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetLastSeenTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetLastSeenTs returns the LastSeenTs field if non-nil, zero value otherwise.

### GetLastSeenTsOk

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetLastSeenTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetLastSeenTsOk returns a tuple with the LastSeenTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTs

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) SetLastSeenTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetLastSeenTs sets LastSeenTs field to given value.

### HasLastSeenTs

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) HasLastSeenTs() bool`

HasLastSeenTs returns a boolean if a field has been set.

### GetServerIp

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetServerPort

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetServerSite

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetServerSite() V2AssuranceFlowSummaryPost200ResponseClientEndpointSite`

GetServerSite returns the ServerSite field if non-nil, zero value otherwise.

### GetServerSiteOk

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) GetServerSiteOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpointSite, bool)`

GetServerSiteOk returns a tuple with the ServerSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSite

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) SetServerSite(v V2AssuranceFlowSummaryPost200ResponseClientEndpointSite)`

SetServerSite sets ServerSite field to given value.

### HasServerSite

`func (o *V2AssuranceTopologyFlowsPost200ResponseFlowsInner) HasServerSite() bool`

HasServerSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


