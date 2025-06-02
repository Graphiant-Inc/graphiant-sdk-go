# V2AssuranceTopologyClientSessionDetailsPost200ResponseSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**ClientEndpoint** | Pointer to [**V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint**](V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint.md) |  | [optional] 
**ClientFlexAlgo** | Pointer to **string** |  | [optional] 
**ClientIp** | Pointer to **string** |  | [optional] 
**ClientLinks** | Pointer to [**[]V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientLinksInner**](V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientLinksInner.md) |  | [optional] 
**FirstSeenTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**LanSegment** | Pointer to **[]string** |  | [optional] 
**LastSeenTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**LocalDiaLinks** | Pointer to [**[]V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionLocalDiaLinksInner**](V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionLocalDiaLinksInner.md) |  | [optional] 
**PopLinks** | Pointer to [**[]V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionPopLinksInner**](V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionPopLinksInner.md) |  | [optional] 
**RemoteDiaLinks** | Pointer to [**[]V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionLocalDiaLinksInner**](V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionLocalDiaLinksInner.md) |  | [optional] 
**RiskStatus** | Pointer to **string** |  | [optional] 
**ServerEndpoint** | Pointer to [**V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint**](V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint.md) |  | [optional] 
**ServerFlexAlgos** | Pointer to **[]string** |  | [optional] 
**ServerIp** | Pointer to **string** |  | [optional] 
**ServerLinks** | Pointer to [**[]V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientLinksInner**](V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientLinksInner.md) |  | [optional] 
**ServerPort** | Pointer to **int32** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**Sla** | Pointer to **string** |  | [optional] 
**SlaClass** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSession

`func NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSession() *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession`

NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSession instantiates a new V2AssuranceTopologyClientSessionDetailsPost200ResponseSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSessionWithDefaults

`func NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSessionWithDefaults() *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession`

NewV2AssuranceTopologyClientSessionDetailsPost200ResponseSessionWithDefaults instantiates a new V2AssuranceTopologyClientSessionDetailsPost200ResponseSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetBucket

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetClientEndpoint

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetClientEndpoint() V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint`

GetClientEndpoint returns the ClientEndpoint field if non-nil, zero value otherwise.

### GetClientEndpointOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetClientEndpointOk() (*V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint, bool)`

GetClientEndpointOk returns a tuple with the ClientEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEndpoint

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetClientEndpoint(v V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint)`

SetClientEndpoint sets ClientEndpoint field to given value.

### HasClientEndpoint

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasClientEndpoint() bool`

HasClientEndpoint returns a boolean if a field has been set.

### GetClientFlexAlgo

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetClientFlexAlgo() string`

GetClientFlexAlgo returns the ClientFlexAlgo field if non-nil, zero value otherwise.

### GetClientFlexAlgoOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetClientFlexAlgoOk() (*string, bool)`

GetClientFlexAlgoOk returns a tuple with the ClientFlexAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFlexAlgo

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetClientFlexAlgo(v string)`

SetClientFlexAlgo sets ClientFlexAlgo field to given value.

### HasClientFlexAlgo

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasClientFlexAlgo() bool`

HasClientFlexAlgo returns a boolean if a field has been set.

### GetClientIp

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetClientLinks() []V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientLinksInner`

GetClientLinks returns the ClientLinks field if non-nil, zero value otherwise.

### GetClientLinksOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetClientLinksOk() (*[]V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientLinksInner, bool)`

GetClientLinksOk returns a tuple with the ClientLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetClientLinks(v []V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientLinksInner)`

SetClientLinks sets ClientLinks field to given value.

### HasClientLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasClientLinks() bool`

HasClientLinks returns a boolean if a field has been set.

### GetFirstSeenTs

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetFirstSeenTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetFirstSeenTs returns the FirstSeenTs field if non-nil, zero value otherwise.

### GetFirstSeenTsOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetFirstSeenTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetFirstSeenTsOk returns a tuple with the FirstSeenTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTs

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetFirstSeenTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetFirstSeenTs sets FirstSeenTs field to given value.

### HasFirstSeenTs

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasFirstSeenTs() bool`

HasFirstSeenTs returns a boolean if a field has been set.

### GetLanSegment

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetLanSegment() []string`

GetLanSegment returns the LanSegment field if non-nil, zero value otherwise.

### GetLanSegmentOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetLanSegmentOk() (*[]string, bool)`

GetLanSegmentOk returns a tuple with the LanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegment

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetLanSegment(v []string)`

SetLanSegment sets LanSegment field to given value.

### HasLanSegment

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasLanSegment() bool`

HasLanSegment returns a boolean if a field has been set.

### GetLastSeenTs

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetLastSeenTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetLastSeenTs returns the LastSeenTs field if non-nil, zero value otherwise.

### GetLastSeenTsOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetLastSeenTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetLastSeenTsOk returns a tuple with the LastSeenTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTs

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetLastSeenTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetLastSeenTs sets LastSeenTs field to given value.

### HasLastSeenTs

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasLastSeenTs() bool`

HasLastSeenTs returns a boolean if a field has been set.

### GetLocalDiaLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetLocalDiaLinks() []V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionLocalDiaLinksInner`

GetLocalDiaLinks returns the LocalDiaLinks field if non-nil, zero value otherwise.

### GetLocalDiaLinksOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetLocalDiaLinksOk() (*[]V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionLocalDiaLinksInner, bool)`

GetLocalDiaLinksOk returns a tuple with the LocalDiaLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDiaLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetLocalDiaLinks(v []V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionLocalDiaLinksInner)`

SetLocalDiaLinks sets LocalDiaLinks field to given value.

### HasLocalDiaLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasLocalDiaLinks() bool`

HasLocalDiaLinks returns a boolean if a field has been set.

### GetPopLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetPopLinks() []V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionPopLinksInner`

GetPopLinks returns the PopLinks field if non-nil, zero value otherwise.

### GetPopLinksOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetPopLinksOk() (*[]V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionPopLinksInner, bool)`

GetPopLinksOk returns a tuple with the PopLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetPopLinks(v []V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionPopLinksInner)`

SetPopLinks sets PopLinks field to given value.

### HasPopLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasPopLinks() bool`

HasPopLinks returns a boolean if a field has been set.

### GetRemoteDiaLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetRemoteDiaLinks() []V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionLocalDiaLinksInner`

GetRemoteDiaLinks returns the RemoteDiaLinks field if non-nil, zero value otherwise.

### GetRemoteDiaLinksOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetRemoteDiaLinksOk() (*[]V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionLocalDiaLinksInner, bool)`

GetRemoteDiaLinksOk returns a tuple with the RemoteDiaLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDiaLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetRemoteDiaLinks(v []V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionLocalDiaLinksInner)`

SetRemoteDiaLinks sets RemoteDiaLinks field to given value.

### HasRemoteDiaLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasRemoteDiaLinks() bool`

HasRemoteDiaLinks returns a boolean if a field has been set.

### GetRiskStatus

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetRiskStatus() string`

GetRiskStatus returns the RiskStatus field if non-nil, zero value otherwise.

### GetRiskStatusOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetRiskStatusOk() (*string, bool)`

GetRiskStatusOk returns a tuple with the RiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskStatus

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetRiskStatus(v string)`

SetRiskStatus sets RiskStatus field to given value.

### HasRiskStatus

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasRiskStatus() bool`

HasRiskStatus returns a boolean if a field has been set.

### GetServerEndpoint

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetServerEndpoint() V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint`

GetServerEndpoint returns the ServerEndpoint field if non-nil, zero value otherwise.

### GetServerEndpointOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetServerEndpointOk() (*V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint, bool)`

GetServerEndpointOk returns a tuple with the ServerEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerEndpoint

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetServerEndpoint(v V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientEndpoint)`

SetServerEndpoint sets ServerEndpoint field to given value.

### HasServerEndpoint

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasServerEndpoint() bool`

HasServerEndpoint returns a boolean if a field has been set.

### GetServerFlexAlgos

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetServerFlexAlgos() []string`

GetServerFlexAlgos returns the ServerFlexAlgos field if non-nil, zero value otherwise.

### GetServerFlexAlgosOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetServerFlexAlgosOk() (*[]string, bool)`

GetServerFlexAlgosOk returns a tuple with the ServerFlexAlgos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFlexAlgos

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetServerFlexAlgos(v []string)`

SetServerFlexAlgos sets ServerFlexAlgos field to given value.

### HasServerFlexAlgos

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasServerFlexAlgos() bool`

HasServerFlexAlgos returns a boolean if a field has been set.

### GetServerIp

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetServerLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetServerLinks() []V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientLinksInner`

GetServerLinks returns the ServerLinks field if non-nil, zero value otherwise.

### GetServerLinksOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetServerLinksOk() (*[]V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientLinksInner, bool)`

GetServerLinksOk returns a tuple with the ServerLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetServerLinks(v []V2AssuranceTopologyClientSessionDetailsPost200ResponseSessionClientLinksInner)`

SetServerLinks sets ServerLinks field to given value.

### HasServerLinks

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasServerLinks() bool`

HasServerLinks returns a boolean if a field has been set.

### GetServerPort

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetSessionId

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSla

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetSla() string`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetSlaOk() (*string, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetSla(v string)`

SetSla sets Sla field to given value.

### HasSla

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetSlaClass

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetSlaClass() string`

GetSlaClass returns the SlaClass field if non-nil, zero value otherwise.

### GetSlaClassOk

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) GetSlaClassOk() (*string, bool)`

GetSlaClassOk returns a tuple with the SlaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaClass

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) SetSlaClass(v string)`

SetSlaClass sets SlaClass field to given value.

### HasSlaClass

`func (o *V2AssuranceTopologyClientSessionDetailsPost200ResponseSession) HasSlaClass() bool`

HasSlaClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


