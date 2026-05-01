# AssuranceClientSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**ClientEndpoint** | Pointer to [**AssuranceClientSessionEndpointDetails**](AssuranceClientSessionEndpointDetails.md) |  | [optional] 
**ClientFlexAlgo** | Pointer to **string** |  | [optional] 
**ClientIp** | Pointer to **string** |  | [optional] 
**ClientLinks** | Pointer to [**[]AssuranceClientSessionEndpointLink**](AssuranceClientSessionEndpointLink.md) |  | [optional] 
**ClientUsername** | Pointer to **string** |  | [optional] 
**FirstSeenTs** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**LanSegment** | Pointer to **[]string** |  | [optional] 
**LastSeenTs** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**LocalDiaLinks** | Pointer to [**[]AssuranceClientSessionDiaLink**](AssuranceClientSessionDiaLink.md) |  | [optional] 
**PopLinks** | Pointer to [**[]AssuranceClientSessionPopLink**](AssuranceClientSessionPopLink.md) |  | [optional] 
**RemoteDiaLinks** | Pointer to [**[]AssuranceClientSessionDiaLink**](AssuranceClientSessionDiaLink.md) |  | [optional] 
**RiskStatus** | Pointer to **string** |  | [optional] 
**ServerEndpoint** | Pointer to [**AssuranceClientSessionEndpointDetails**](AssuranceClientSessionEndpointDetails.md) |  | [optional] 
**ServerFlexAlgos** | Pointer to **[]string** |  | [optional] 
**ServerIp** | Pointer to **string** |  | [optional] 
**ServerLinks** | Pointer to [**[]AssuranceClientSessionEndpointLink**](AssuranceClientSessionEndpointLink.md) |  | [optional] 
**ServerPort** | Pointer to **int32** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**Sla** | Pointer to **string** |  | [optional] 
**SlaClass** | Pointer to **string** |  | [optional] 

## Methods

### NewAssuranceClientSession

`func NewAssuranceClientSession() *AssuranceClientSession`

NewAssuranceClientSession instantiates a new AssuranceClientSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssuranceClientSessionWithDefaults

`func NewAssuranceClientSessionWithDefaults() *AssuranceClientSession`

NewAssuranceClientSessionWithDefaults instantiates a new AssuranceClientSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *AssuranceClientSession) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AssuranceClientSession) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AssuranceClientSession) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *AssuranceClientSession) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetBucket

`func (o *AssuranceClientSession) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AssuranceClientSession) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AssuranceClientSession) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *AssuranceClientSession) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetClientEndpoint

`func (o *AssuranceClientSession) GetClientEndpoint() AssuranceClientSessionEndpointDetails`

GetClientEndpoint returns the ClientEndpoint field if non-nil, zero value otherwise.

### GetClientEndpointOk

`func (o *AssuranceClientSession) GetClientEndpointOk() (*AssuranceClientSessionEndpointDetails, bool)`

GetClientEndpointOk returns a tuple with the ClientEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEndpoint

`func (o *AssuranceClientSession) SetClientEndpoint(v AssuranceClientSessionEndpointDetails)`

SetClientEndpoint sets ClientEndpoint field to given value.

### HasClientEndpoint

`func (o *AssuranceClientSession) HasClientEndpoint() bool`

HasClientEndpoint returns a boolean if a field has been set.

### GetClientFlexAlgo

`func (o *AssuranceClientSession) GetClientFlexAlgo() string`

GetClientFlexAlgo returns the ClientFlexAlgo field if non-nil, zero value otherwise.

### GetClientFlexAlgoOk

`func (o *AssuranceClientSession) GetClientFlexAlgoOk() (*string, bool)`

GetClientFlexAlgoOk returns a tuple with the ClientFlexAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFlexAlgo

`func (o *AssuranceClientSession) SetClientFlexAlgo(v string)`

SetClientFlexAlgo sets ClientFlexAlgo field to given value.

### HasClientFlexAlgo

`func (o *AssuranceClientSession) HasClientFlexAlgo() bool`

HasClientFlexAlgo returns a boolean if a field has been set.

### GetClientIp

`func (o *AssuranceClientSession) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *AssuranceClientSession) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *AssuranceClientSession) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *AssuranceClientSession) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientLinks

`func (o *AssuranceClientSession) GetClientLinks() []AssuranceClientSessionEndpointLink`

GetClientLinks returns the ClientLinks field if non-nil, zero value otherwise.

### GetClientLinksOk

`func (o *AssuranceClientSession) GetClientLinksOk() (*[]AssuranceClientSessionEndpointLink, bool)`

GetClientLinksOk returns a tuple with the ClientLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientLinks

`func (o *AssuranceClientSession) SetClientLinks(v []AssuranceClientSessionEndpointLink)`

SetClientLinks sets ClientLinks field to given value.

### HasClientLinks

`func (o *AssuranceClientSession) HasClientLinks() bool`

HasClientLinks returns a boolean if a field has been set.

### GetClientUsername

`func (o *AssuranceClientSession) GetClientUsername() string`

GetClientUsername returns the ClientUsername field if non-nil, zero value otherwise.

### GetClientUsernameOk

`func (o *AssuranceClientSession) GetClientUsernameOk() (*string, bool)`

GetClientUsernameOk returns a tuple with the ClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUsername

`func (o *AssuranceClientSession) SetClientUsername(v string)`

SetClientUsername sets ClientUsername field to given value.

### HasClientUsername

`func (o *AssuranceClientSession) HasClientUsername() bool`

HasClientUsername returns a boolean if a field has been set.

### GetFirstSeenTs

`func (o *AssuranceClientSession) GetFirstSeenTs() GoogleProtobufTimestamp`

GetFirstSeenTs returns the FirstSeenTs field if non-nil, zero value otherwise.

### GetFirstSeenTsOk

`func (o *AssuranceClientSession) GetFirstSeenTsOk() (*GoogleProtobufTimestamp, bool)`

GetFirstSeenTsOk returns a tuple with the FirstSeenTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTs

`func (o *AssuranceClientSession) SetFirstSeenTs(v GoogleProtobufTimestamp)`

SetFirstSeenTs sets FirstSeenTs field to given value.

### HasFirstSeenTs

`func (o *AssuranceClientSession) HasFirstSeenTs() bool`

HasFirstSeenTs returns a boolean if a field has been set.

### GetLanSegment

`func (o *AssuranceClientSession) GetLanSegment() []string`

GetLanSegment returns the LanSegment field if non-nil, zero value otherwise.

### GetLanSegmentOk

`func (o *AssuranceClientSession) GetLanSegmentOk() (*[]string, bool)`

GetLanSegmentOk returns a tuple with the LanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegment

`func (o *AssuranceClientSession) SetLanSegment(v []string)`

SetLanSegment sets LanSegment field to given value.

### HasLanSegment

`func (o *AssuranceClientSession) HasLanSegment() bool`

HasLanSegment returns a boolean if a field has been set.

### GetLastSeenTs

`func (o *AssuranceClientSession) GetLastSeenTs() GoogleProtobufTimestamp`

GetLastSeenTs returns the LastSeenTs field if non-nil, zero value otherwise.

### GetLastSeenTsOk

`func (o *AssuranceClientSession) GetLastSeenTsOk() (*GoogleProtobufTimestamp, bool)`

GetLastSeenTsOk returns a tuple with the LastSeenTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTs

`func (o *AssuranceClientSession) SetLastSeenTs(v GoogleProtobufTimestamp)`

SetLastSeenTs sets LastSeenTs field to given value.

### HasLastSeenTs

`func (o *AssuranceClientSession) HasLastSeenTs() bool`

HasLastSeenTs returns a boolean if a field has been set.

### GetLocalDiaLinks

`func (o *AssuranceClientSession) GetLocalDiaLinks() []AssuranceClientSessionDiaLink`

GetLocalDiaLinks returns the LocalDiaLinks field if non-nil, zero value otherwise.

### GetLocalDiaLinksOk

`func (o *AssuranceClientSession) GetLocalDiaLinksOk() (*[]AssuranceClientSessionDiaLink, bool)`

GetLocalDiaLinksOk returns a tuple with the LocalDiaLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDiaLinks

`func (o *AssuranceClientSession) SetLocalDiaLinks(v []AssuranceClientSessionDiaLink)`

SetLocalDiaLinks sets LocalDiaLinks field to given value.

### HasLocalDiaLinks

`func (o *AssuranceClientSession) HasLocalDiaLinks() bool`

HasLocalDiaLinks returns a boolean if a field has been set.

### GetPopLinks

`func (o *AssuranceClientSession) GetPopLinks() []AssuranceClientSessionPopLink`

GetPopLinks returns the PopLinks field if non-nil, zero value otherwise.

### GetPopLinksOk

`func (o *AssuranceClientSession) GetPopLinksOk() (*[]AssuranceClientSessionPopLink, bool)`

GetPopLinksOk returns a tuple with the PopLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopLinks

`func (o *AssuranceClientSession) SetPopLinks(v []AssuranceClientSessionPopLink)`

SetPopLinks sets PopLinks field to given value.

### HasPopLinks

`func (o *AssuranceClientSession) HasPopLinks() bool`

HasPopLinks returns a boolean if a field has been set.

### GetRemoteDiaLinks

`func (o *AssuranceClientSession) GetRemoteDiaLinks() []AssuranceClientSessionDiaLink`

GetRemoteDiaLinks returns the RemoteDiaLinks field if non-nil, zero value otherwise.

### GetRemoteDiaLinksOk

`func (o *AssuranceClientSession) GetRemoteDiaLinksOk() (*[]AssuranceClientSessionDiaLink, bool)`

GetRemoteDiaLinksOk returns a tuple with the RemoteDiaLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDiaLinks

`func (o *AssuranceClientSession) SetRemoteDiaLinks(v []AssuranceClientSessionDiaLink)`

SetRemoteDiaLinks sets RemoteDiaLinks field to given value.

### HasRemoteDiaLinks

`func (o *AssuranceClientSession) HasRemoteDiaLinks() bool`

HasRemoteDiaLinks returns a boolean if a field has been set.

### GetRiskStatus

`func (o *AssuranceClientSession) GetRiskStatus() string`

GetRiskStatus returns the RiskStatus field if non-nil, zero value otherwise.

### GetRiskStatusOk

`func (o *AssuranceClientSession) GetRiskStatusOk() (*string, bool)`

GetRiskStatusOk returns a tuple with the RiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskStatus

`func (o *AssuranceClientSession) SetRiskStatus(v string)`

SetRiskStatus sets RiskStatus field to given value.

### HasRiskStatus

`func (o *AssuranceClientSession) HasRiskStatus() bool`

HasRiskStatus returns a boolean if a field has been set.

### GetServerEndpoint

`func (o *AssuranceClientSession) GetServerEndpoint() AssuranceClientSessionEndpointDetails`

GetServerEndpoint returns the ServerEndpoint field if non-nil, zero value otherwise.

### GetServerEndpointOk

`func (o *AssuranceClientSession) GetServerEndpointOk() (*AssuranceClientSessionEndpointDetails, bool)`

GetServerEndpointOk returns a tuple with the ServerEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerEndpoint

`func (o *AssuranceClientSession) SetServerEndpoint(v AssuranceClientSessionEndpointDetails)`

SetServerEndpoint sets ServerEndpoint field to given value.

### HasServerEndpoint

`func (o *AssuranceClientSession) HasServerEndpoint() bool`

HasServerEndpoint returns a boolean if a field has been set.

### GetServerFlexAlgos

`func (o *AssuranceClientSession) GetServerFlexAlgos() []string`

GetServerFlexAlgos returns the ServerFlexAlgos field if non-nil, zero value otherwise.

### GetServerFlexAlgosOk

`func (o *AssuranceClientSession) GetServerFlexAlgosOk() (*[]string, bool)`

GetServerFlexAlgosOk returns a tuple with the ServerFlexAlgos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFlexAlgos

`func (o *AssuranceClientSession) SetServerFlexAlgos(v []string)`

SetServerFlexAlgos sets ServerFlexAlgos field to given value.

### HasServerFlexAlgos

`func (o *AssuranceClientSession) HasServerFlexAlgos() bool`

HasServerFlexAlgos returns a boolean if a field has been set.

### GetServerIp

`func (o *AssuranceClientSession) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *AssuranceClientSession) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *AssuranceClientSession) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *AssuranceClientSession) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetServerLinks

`func (o *AssuranceClientSession) GetServerLinks() []AssuranceClientSessionEndpointLink`

GetServerLinks returns the ServerLinks field if non-nil, zero value otherwise.

### GetServerLinksOk

`func (o *AssuranceClientSession) GetServerLinksOk() (*[]AssuranceClientSessionEndpointLink, bool)`

GetServerLinksOk returns a tuple with the ServerLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLinks

`func (o *AssuranceClientSession) SetServerLinks(v []AssuranceClientSessionEndpointLink)`

SetServerLinks sets ServerLinks field to given value.

### HasServerLinks

`func (o *AssuranceClientSession) HasServerLinks() bool`

HasServerLinks returns a boolean if a field has been set.

### GetServerPort

`func (o *AssuranceClientSession) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AssuranceClientSession) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AssuranceClientSession) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AssuranceClientSession) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetSessionId

`func (o *AssuranceClientSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AssuranceClientSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AssuranceClientSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *AssuranceClientSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSla

`func (o *AssuranceClientSession) GetSla() string`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *AssuranceClientSession) GetSlaOk() (*string, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *AssuranceClientSession) SetSla(v string)`

SetSla sets Sla field to given value.

### HasSla

`func (o *AssuranceClientSession) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetSlaClass

`func (o *AssuranceClientSession) GetSlaClass() string`

GetSlaClass returns the SlaClass field if non-nil, zero value otherwise.

### GetSlaClassOk

`func (o *AssuranceClientSession) GetSlaClassOk() (*string, bool)`

GetSlaClassOk returns a tuple with the SlaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaClass

`func (o *AssuranceClientSession) SetSlaClass(v string)`

SetSlaClass sets SlaClass field to given value.

### HasSlaClass

`func (o *AssuranceClientSession) HasSlaClass() bool`

HasSlaClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


