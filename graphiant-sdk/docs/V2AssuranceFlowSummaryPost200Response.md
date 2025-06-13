# V2AssuranceFlowSummaryPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**ClientEndpoint** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpoint**](V2AssuranceFlowSummaryPost200ResponseClientEndpoint.md) |  | [optional] 
**ClientIp** | Pointer to **string** |  | [optional] 
**FirstSeenTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**LanSegment** | Pointer to **string** |  | [optional] 
**LastSeenTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**RiskStatus** | Pointer to **string** |  | [optional] 
**ServerEndpoint** | Pointer to [**V2AssuranceFlowSummaryPost200ResponseClientEndpoint**](V2AssuranceFlowSummaryPost200ResponseClientEndpoint.md) |  | [optional] 
**ServerIp** | Pointer to **string** |  | [optional] 
**ServerPort** | Pointer to **int32** |  | [optional] 
**SlaClass** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssuranceFlowSummaryPost200Response

`func NewV2AssuranceFlowSummaryPost200Response() *V2AssuranceFlowSummaryPost200Response`

NewV2AssuranceFlowSummaryPost200Response instantiates a new V2AssuranceFlowSummaryPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceFlowSummaryPost200ResponseWithDefaults

`func NewV2AssuranceFlowSummaryPost200ResponseWithDefaults() *V2AssuranceFlowSummaryPost200Response`

NewV2AssuranceFlowSummaryPost200ResponseWithDefaults instantiates a new V2AssuranceFlowSummaryPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *V2AssuranceFlowSummaryPost200Response) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *V2AssuranceFlowSummaryPost200Response) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *V2AssuranceFlowSummaryPost200Response) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetBucket

`func (o *V2AssuranceFlowSummaryPost200Response) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *V2AssuranceFlowSummaryPost200Response) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *V2AssuranceFlowSummaryPost200Response) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetClientEndpoint

`func (o *V2AssuranceFlowSummaryPost200Response) GetClientEndpoint() V2AssuranceFlowSummaryPost200ResponseClientEndpoint`

GetClientEndpoint returns the ClientEndpoint field if non-nil, zero value otherwise.

### GetClientEndpointOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetClientEndpointOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpoint, bool)`

GetClientEndpointOk returns a tuple with the ClientEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEndpoint

`func (o *V2AssuranceFlowSummaryPost200Response) SetClientEndpoint(v V2AssuranceFlowSummaryPost200ResponseClientEndpoint)`

SetClientEndpoint sets ClientEndpoint field to given value.

### HasClientEndpoint

`func (o *V2AssuranceFlowSummaryPost200Response) HasClientEndpoint() bool`

HasClientEndpoint returns a boolean if a field has been set.

### GetClientIp

`func (o *V2AssuranceFlowSummaryPost200Response) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *V2AssuranceFlowSummaryPost200Response) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *V2AssuranceFlowSummaryPost200Response) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetFirstSeenTs

`func (o *V2AssuranceFlowSummaryPost200Response) GetFirstSeenTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetFirstSeenTs returns the FirstSeenTs field if non-nil, zero value otherwise.

### GetFirstSeenTsOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetFirstSeenTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetFirstSeenTsOk returns a tuple with the FirstSeenTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTs

`func (o *V2AssuranceFlowSummaryPost200Response) SetFirstSeenTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetFirstSeenTs sets FirstSeenTs field to given value.

### HasFirstSeenTs

`func (o *V2AssuranceFlowSummaryPost200Response) HasFirstSeenTs() bool`

HasFirstSeenTs returns a boolean if a field has been set.

### GetLanSegment

`func (o *V2AssuranceFlowSummaryPost200Response) GetLanSegment() string`

GetLanSegment returns the LanSegment field if non-nil, zero value otherwise.

### GetLanSegmentOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetLanSegmentOk() (*string, bool)`

GetLanSegmentOk returns a tuple with the LanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegment

`func (o *V2AssuranceFlowSummaryPost200Response) SetLanSegment(v string)`

SetLanSegment sets LanSegment field to given value.

### HasLanSegment

`func (o *V2AssuranceFlowSummaryPost200Response) HasLanSegment() bool`

HasLanSegment returns a boolean if a field has been set.

### GetLastSeenTs

`func (o *V2AssuranceFlowSummaryPost200Response) GetLastSeenTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetLastSeenTs returns the LastSeenTs field if non-nil, zero value otherwise.

### GetLastSeenTsOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetLastSeenTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetLastSeenTsOk returns a tuple with the LastSeenTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTs

`func (o *V2AssuranceFlowSummaryPost200Response) SetLastSeenTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetLastSeenTs sets LastSeenTs field to given value.

### HasLastSeenTs

`func (o *V2AssuranceFlowSummaryPost200Response) HasLastSeenTs() bool`

HasLastSeenTs returns a boolean if a field has been set.

### GetRiskStatus

`func (o *V2AssuranceFlowSummaryPost200Response) GetRiskStatus() string`

GetRiskStatus returns the RiskStatus field if non-nil, zero value otherwise.

### GetRiskStatusOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetRiskStatusOk() (*string, bool)`

GetRiskStatusOk returns a tuple with the RiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskStatus

`func (o *V2AssuranceFlowSummaryPost200Response) SetRiskStatus(v string)`

SetRiskStatus sets RiskStatus field to given value.

### HasRiskStatus

`func (o *V2AssuranceFlowSummaryPost200Response) HasRiskStatus() bool`

HasRiskStatus returns a boolean if a field has been set.

### GetServerEndpoint

`func (o *V2AssuranceFlowSummaryPost200Response) GetServerEndpoint() V2AssuranceFlowSummaryPost200ResponseClientEndpoint`

GetServerEndpoint returns the ServerEndpoint field if non-nil, zero value otherwise.

### GetServerEndpointOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetServerEndpointOk() (*V2AssuranceFlowSummaryPost200ResponseClientEndpoint, bool)`

GetServerEndpointOk returns a tuple with the ServerEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerEndpoint

`func (o *V2AssuranceFlowSummaryPost200Response) SetServerEndpoint(v V2AssuranceFlowSummaryPost200ResponseClientEndpoint)`

SetServerEndpoint sets ServerEndpoint field to given value.

### HasServerEndpoint

`func (o *V2AssuranceFlowSummaryPost200Response) HasServerEndpoint() bool`

HasServerEndpoint returns a boolean if a field has been set.

### GetServerIp

`func (o *V2AssuranceFlowSummaryPost200Response) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *V2AssuranceFlowSummaryPost200Response) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *V2AssuranceFlowSummaryPost200Response) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetServerPort

`func (o *V2AssuranceFlowSummaryPost200Response) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *V2AssuranceFlowSummaryPost200Response) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *V2AssuranceFlowSummaryPost200Response) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetSlaClass

`func (o *V2AssuranceFlowSummaryPost200Response) GetSlaClass() string`

GetSlaClass returns the SlaClass field if non-nil, zero value otherwise.

### GetSlaClassOk

`func (o *V2AssuranceFlowSummaryPost200Response) GetSlaClassOk() (*string, bool)`

GetSlaClassOk returns a tuple with the SlaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaClass

`func (o *V2AssuranceFlowSummaryPost200Response) SetSlaClass(v string)`

SetSlaClass sets SlaClass field to given value.

### HasSlaClass

`func (o *V2AssuranceFlowSummaryPost200Response) HasSlaClass() bool`

HasSlaClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


