# AssuranceAppIdRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedHosts** | Pointer to **int64** |  | [optional] 
**AffectedRegions** | Pointer to **int64** |  | [optional] 
**AffectedSites** | Pointer to **int64** |  | [optional] 
**AffectedVrfs** | Pointer to **int64** |  | [optional] 
**AppIdKey** | Pointer to **string** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**AppType** | Pointer to **string** |  | [optional] 
**BlockedPolicyList** | Pointer to **[]string** |  | [optional] 
**BlockedReasonList** | Pointer to **[]string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ClassficationField** | Pointer to **string** |  | [optional] 
**ClassificationField** | Pointer to **string** |  | [optional] 
**Clients** | Pointer to **[]string** |  | [optional] 
**ExchangeService** | Pointer to [**[]AssuranceExchangeServiceIdentifier**](AssuranceExchangeServiceIdentifier.md) |  | [optional] 
**FirstSeen** | Pointer to **int64** |  | [optional] 
**FlexAlgo** | Pointer to [**[]AssuranceFlexAlgoIdentifier**](AssuranceFlexAlgoIdentifier.md) |  | [optional] 
**FlowsAnalyzed** | Pointer to **int64** |  | [optional] 
**LastSeen** | Pointer to **int64** |  | [optional] 
**NewIpHint** | Pointer to **bool** |  | [optional] 
**Recommendation** | Pointer to **string** |  | [optional] 
**RegionList** | Pointer to **[]string** |  | [optional] 
**RiskStatus** | Pointer to **string** |  | [optional] 
**ServerIp** | Pointer to **string** |  | [optional] 
**ServerPort** | Pointer to **string** |  | [optional] 
**SiteList** | Pointer to **[]string** |  | [optional] 
**ThreatScore** | Pointer to **int64** |  | [optional] 
**VrfList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAssuranceAppIdRecord

`func NewAssuranceAppIdRecord() *AssuranceAppIdRecord`

NewAssuranceAppIdRecord instantiates a new AssuranceAppIdRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssuranceAppIdRecordWithDefaults

`func NewAssuranceAppIdRecordWithDefaults() *AssuranceAppIdRecord`

NewAssuranceAppIdRecordWithDefaults instantiates a new AssuranceAppIdRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedHosts

`func (o *AssuranceAppIdRecord) GetAffectedHosts() int64`

GetAffectedHosts returns the AffectedHosts field if non-nil, zero value otherwise.

### GetAffectedHostsOk

`func (o *AssuranceAppIdRecord) GetAffectedHostsOk() (*int64, bool)`

GetAffectedHostsOk returns a tuple with the AffectedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedHosts

`func (o *AssuranceAppIdRecord) SetAffectedHosts(v int64)`

SetAffectedHosts sets AffectedHosts field to given value.

### HasAffectedHosts

`func (o *AssuranceAppIdRecord) HasAffectedHosts() bool`

HasAffectedHosts returns a boolean if a field has been set.

### GetAffectedRegions

`func (o *AssuranceAppIdRecord) GetAffectedRegions() int64`

GetAffectedRegions returns the AffectedRegions field if non-nil, zero value otherwise.

### GetAffectedRegionsOk

`func (o *AssuranceAppIdRecord) GetAffectedRegionsOk() (*int64, bool)`

GetAffectedRegionsOk returns a tuple with the AffectedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedRegions

`func (o *AssuranceAppIdRecord) SetAffectedRegions(v int64)`

SetAffectedRegions sets AffectedRegions field to given value.

### HasAffectedRegions

`func (o *AssuranceAppIdRecord) HasAffectedRegions() bool`

HasAffectedRegions returns a boolean if a field has been set.

### GetAffectedSites

`func (o *AssuranceAppIdRecord) GetAffectedSites() int64`

GetAffectedSites returns the AffectedSites field if non-nil, zero value otherwise.

### GetAffectedSitesOk

`func (o *AssuranceAppIdRecord) GetAffectedSitesOk() (*int64, bool)`

GetAffectedSitesOk returns a tuple with the AffectedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedSites

`func (o *AssuranceAppIdRecord) SetAffectedSites(v int64)`

SetAffectedSites sets AffectedSites field to given value.

### HasAffectedSites

`func (o *AssuranceAppIdRecord) HasAffectedSites() bool`

HasAffectedSites returns a boolean if a field has been set.

### GetAffectedVrfs

`func (o *AssuranceAppIdRecord) GetAffectedVrfs() int64`

GetAffectedVrfs returns the AffectedVrfs field if non-nil, zero value otherwise.

### GetAffectedVrfsOk

`func (o *AssuranceAppIdRecord) GetAffectedVrfsOk() (*int64, bool)`

GetAffectedVrfsOk returns a tuple with the AffectedVrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVrfs

`func (o *AssuranceAppIdRecord) SetAffectedVrfs(v int64)`

SetAffectedVrfs sets AffectedVrfs field to given value.

### HasAffectedVrfs

`func (o *AssuranceAppIdRecord) HasAffectedVrfs() bool`

HasAffectedVrfs returns a boolean if a field has been set.

### GetAppIdKey

`func (o *AssuranceAppIdRecord) GetAppIdKey() string`

GetAppIdKey returns the AppIdKey field if non-nil, zero value otherwise.

### GetAppIdKeyOk

`func (o *AssuranceAppIdRecord) GetAppIdKeyOk() (*string, bool)`

GetAppIdKeyOk returns a tuple with the AppIdKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIdKey

`func (o *AssuranceAppIdRecord) SetAppIdKey(v string)`

SetAppIdKey sets AppIdKey field to given value.

### HasAppIdKey

`func (o *AssuranceAppIdRecord) HasAppIdKey() bool`

HasAppIdKey returns a boolean if a field has been set.

### GetAppName

`func (o *AssuranceAppIdRecord) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AssuranceAppIdRecord) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AssuranceAppIdRecord) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *AssuranceAppIdRecord) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppType

`func (o *AssuranceAppIdRecord) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *AssuranceAppIdRecord) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *AssuranceAppIdRecord) SetAppType(v string)`

SetAppType sets AppType field to given value.

### HasAppType

`func (o *AssuranceAppIdRecord) HasAppType() bool`

HasAppType returns a boolean if a field has been set.

### GetBlockedPolicyList

`func (o *AssuranceAppIdRecord) GetBlockedPolicyList() []string`

GetBlockedPolicyList returns the BlockedPolicyList field if non-nil, zero value otherwise.

### GetBlockedPolicyListOk

`func (o *AssuranceAppIdRecord) GetBlockedPolicyListOk() (*[]string, bool)`

GetBlockedPolicyListOk returns a tuple with the BlockedPolicyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPolicyList

`func (o *AssuranceAppIdRecord) SetBlockedPolicyList(v []string)`

SetBlockedPolicyList sets BlockedPolicyList field to given value.

### HasBlockedPolicyList

`func (o *AssuranceAppIdRecord) HasBlockedPolicyList() bool`

HasBlockedPolicyList returns a boolean if a field has been set.

### GetBlockedReasonList

`func (o *AssuranceAppIdRecord) GetBlockedReasonList() []string`

GetBlockedReasonList returns the BlockedReasonList field if non-nil, zero value otherwise.

### GetBlockedReasonListOk

`func (o *AssuranceAppIdRecord) GetBlockedReasonListOk() (*[]string, bool)`

GetBlockedReasonListOk returns a tuple with the BlockedReasonList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReasonList

`func (o *AssuranceAppIdRecord) SetBlockedReasonList(v []string)`

SetBlockedReasonList sets BlockedReasonList field to given value.

### HasBlockedReasonList

`func (o *AssuranceAppIdRecord) HasBlockedReasonList() bool`

HasBlockedReasonList returns a boolean if a field has been set.

### GetCategory

`func (o *AssuranceAppIdRecord) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AssuranceAppIdRecord) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AssuranceAppIdRecord) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AssuranceAppIdRecord) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClassficationField

`func (o *AssuranceAppIdRecord) GetClassficationField() string`

GetClassficationField returns the ClassficationField field if non-nil, zero value otherwise.

### GetClassficationFieldOk

`func (o *AssuranceAppIdRecord) GetClassficationFieldOk() (*string, bool)`

GetClassficationFieldOk returns a tuple with the ClassficationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassficationField

`func (o *AssuranceAppIdRecord) SetClassficationField(v string)`

SetClassficationField sets ClassficationField field to given value.

### HasClassficationField

`func (o *AssuranceAppIdRecord) HasClassficationField() bool`

HasClassficationField returns a boolean if a field has been set.

### GetClassificationField

`func (o *AssuranceAppIdRecord) GetClassificationField() string`

GetClassificationField returns the ClassificationField field if non-nil, zero value otherwise.

### GetClassificationFieldOk

`func (o *AssuranceAppIdRecord) GetClassificationFieldOk() (*string, bool)`

GetClassificationFieldOk returns a tuple with the ClassificationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationField

`func (o *AssuranceAppIdRecord) SetClassificationField(v string)`

SetClassificationField sets ClassificationField field to given value.

### HasClassificationField

`func (o *AssuranceAppIdRecord) HasClassificationField() bool`

HasClassificationField returns a boolean if a field has been set.

### GetClients

`func (o *AssuranceAppIdRecord) GetClients() []string`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *AssuranceAppIdRecord) GetClientsOk() (*[]string, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *AssuranceAppIdRecord) SetClients(v []string)`

SetClients sets Clients field to given value.

### HasClients

`func (o *AssuranceAppIdRecord) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetExchangeService

`func (o *AssuranceAppIdRecord) GetExchangeService() []AssuranceExchangeServiceIdentifier`

GetExchangeService returns the ExchangeService field if non-nil, zero value otherwise.

### GetExchangeServiceOk

`func (o *AssuranceAppIdRecord) GetExchangeServiceOk() (*[]AssuranceExchangeServiceIdentifier, bool)`

GetExchangeServiceOk returns a tuple with the ExchangeService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeService

`func (o *AssuranceAppIdRecord) SetExchangeService(v []AssuranceExchangeServiceIdentifier)`

SetExchangeService sets ExchangeService field to given value.

### HasExchangeService

`func (o *AssuranceAppIdRecord) HasExchangeService() bool`

HasExchangeService returns a boolean if a field has been set.

### GetFirstSeen

`func (o *AssuranceAppIdRecord) GetFirstSeen() int64`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *AssuranceAppIdRecord) GetFirstSeenOk() (*int64, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *AssuranceAppIdRecord) SetFirstSeen(v int64)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *AssuranceAppIdRecord) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetFlexAlgo

`func (o *AssuranceAppIdRecord) GetFlexAlgo() []AssuranceFlexAlgoIdentifier`

GetFlexAlgo returns the FlexAlgo field if non-nil, zero value otherwise.

### GetFlexAlgoOk

`func (o *AssuranceAppIdRecord) GetFlexAlgoOk() (*[]AssuranceFlexAlgoIdentifier, bool)`

GetFlexAlgoOk returns a tuple with the FlexAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAlgo

`func (o *AssuranceAppIdRecord) SetFlexAlgo(v []AssuranceFlexAlgoIdentifier)`

SetFlexAlgo sets FlexAlgo field to given value.

### HasFlexAlgo

`func (o *AssuranceAppIdRecord) HasFlexAlgo() bool`

HasFlexAlgo returns a boolean if a field has been set.

### GetFlowsAnalyzed

`func (o *AssuranceAppIdRecord) GetFlowsAnalyzed() int64`

GetFlowsAnalyzed returns the FlowsAnalyzed field if non-nil, zero value otherwise.

### GetFlowsAnalyzedOk

`func (o *AssuranceAppIdRecord) GetFlowsAnalyzedOk() (*int64, bool)`

GetFlowsAnalyzedOk returns a tuple with the FlowsAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowsAnalyzed

`func (o *AssuranceAppIdRecord) SetFlowsAnalyzed(v int64)`

SetFlowsAnalyzed sets FlowsAnalyzed field to given value.

### HasFlowsAnalyzed

`func (o *AssuranceAppIdRecord) HasFlowsAnalyzed() bool`

HasFlowsAnalyzed returns a boolean if a field has been set.

### GetLastSeen

`func (o *AssuranceAppIdRecord) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *AssuranceAppIdRecord) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *AssuranceAppIdRecord) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *AssuranceAppIdRecord) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetNewIpHint

`func (o *AssuranceAppIdRecord) GetNewIpHint() bool`

GetNewIpHint returns the NewIpHint field if non-nil, zero value otherwise.

### GetNewIpHintOk

`func (o *AssuranceAppIdRecord) GetNewIpHintOk() (*bool, bool)`

GetNewIpHintOk returns a tuple with the NewIpHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIpHint

`func (o *AssuranceAppIdRecord) SetNewIpHint(v bool)`

SetNewIpHint sets NewIpHint field to given value.

### HasNewIpHint

`func (o *AssuranceAppIdRecord) HasNewIpHint() bool`

HasNewIpHint returns a boolean if a field has been set.

### GetRecommendation

`func (o *AssuranceAppIdRecord) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *AssuranceAppIdRecord) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *AssuranceAppIdRecord) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *AssuranceAppIdRecord) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetRegionList

`func (o *AssuranceAppIdRecord) GetRegionList() []string`

GetRegionList returns the RegionList field if non-nil, zero value otherwise.

### GetRegionListOk

`func (o *AssuranceAppIdRecord) GetRegionListOk() (*[]string, bool)`

GetRegionListOk returns a tuple with the RegionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionList

`func (o *AssuranceAppIdRecord) SetRegionList(v []string)`

SetRegionList sets RegionList field to given value.

### HasRegionList

`func (o *AssuranceAppIdRecord) HasRegionList() bool`

HasRegionList returns a boolean if a field has been set.

### GetRiskStatus

`func (o *AssuranceAppIdRecord) GetRiskStatus() string`

GetRiskStatus returns the RiskStatus field if non-nil, zero value otherwise.

### GetRiskStatusOk

`func (o *AssuranceAppIdRecord) GetRiskStatusOk() (*string, bool)`

GetRiskStatusOk returns a tuple with the RiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskStatus

`func (o *AssuranceAppIdRecord) SetRiskStatus(v string)`

SetRiskStatus sets RiskStatus field to given value.

### HasRiskStatus

`func (o *AssuranceAppIdRecord) HasRiskStatus() bool`

HasRiskStatus returns a boolean if a field has been set.

### GetServerIp

`func (o *AssuranceAppIdRecord) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *AssuranceAppIdRecord) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *AssuranceAppIdRecord) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *AssuranceAppIdRecord) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetServerPort

`func (o *AssuranceAppIdRecord) GetServerPort() string`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AssuranceAppIdRecord) GetServerPortOk() (*string, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AssuranceAppIdRecord) SetServerPort(v string)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AssuranceAppIdRecord) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetSiteList

`func (o *AssuranceAppIdRecord) GetSiteList() []string`

GetSiteList returns the SiteList field if non-nil, zero value otherwise.

### GetSiteListOk

`func (o *AssuranceAppIdRecord) GetSiteListOk() (*[]string, bool)`

GetSiteListOk returns a tuple with the SiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteList

`func (o *AssuranceAppIdRecord) SetSiteList(v []string)`

SetSiteList sets SiteList field to given value.

### HasSiteList

`func (o *AssuranceAppIdRecord) HasSiteList() bool`

HasSiteList returns a boolean if a field has been set.

### GetThreatScore

`func (o *AssuranceAppIdRecord) GetThreatScore() int64`

GetThreatScore returns the ThreatScore field if non-nil, zero value otherwise.

### GetThreatScoreOk

`func (o *AssuranceAppIdRecord) GetThreatScoreOk() (*int64, bool)`

GetThreatScoreOk returns a tuple with the ThreatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatScore

`func (o *AssuranceAppIdRecord) SetThreatScore(v int64)`

SetThreatScore sets ThreatScore field to given value.

### HasThreatScore

`func (o *AssuranceAppIdRecord) HasThreatScore() bool`

HasThreatScore returns a boolean if a field has been set.

### GetVrfList

`func (o *AssuranceAppIdRecord) GetVrfList() []string`

GetVrfList returns the VrfList field if non-nil, zero value otherwise.

### GetVrfListOk

`func (o *AssuranceAppIdRecord) GetVrfListOk() (*[]string, bool)`

GetVrfListOk returns a tuple with the VrfList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfList

`func (o *AssuranceAppIdRecord) SetVrfList(v []string)`

SetVrfList sets VrfList field to given value.

### HasVrfList

`func (o *AssuranceAppIdRecord) HasVrfList() bool`

HasVrfList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


