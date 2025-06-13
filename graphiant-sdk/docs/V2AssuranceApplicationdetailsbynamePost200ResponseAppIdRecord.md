# V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord

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
**BlockedReasonList** | Pointer to **[]string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ClassficationField** | Pointer to **string** |  | [optional] 
**ClassificationField** | Pointer to **string** |  | [optional] 
**FirstSeen** | Pointer to **int64** |  | [optional] 
**FlexAlgo** | Pointer to [**[]V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecordFlexAlgoInner**](V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecordFlexAlgoInner.md) |  | [optional] 
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

### NewV2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord

`func NewV2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord() *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord`

NewV2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord instantiates a new V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecordWithDefaults

`func NewV2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecordWithDefaults() *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord`

NewV2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecordWithDefaults instantiates a new V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedHosts

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAffectedHosts() int64`

GetAffectedHosts returns the AffectedHosts field if non-nil, zero value otherwise.

### GetAffectedHostsOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAffectedHostsOk() (*int64, bool)`

GetAffectedHostsOk returns a tuple with the AffectedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedHosts

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetAffectedHosts(v int64)`

SetAffectedHosts sets AffectedHosts field to given value.

### HasAffectedHosts

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasAffectedHosts() bool`

HasAffectedHosts returns a boolean if a field has been set.

### GetAffectedRegions

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAffectedRegions() int64`

GetAffectedRegions returns the AffectedRegions field if non-nil, zero value otherwise.

### GetAffectedRegionsOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAffectedRegionsOk() (*int64, bool)`

GetAffectedRegionsOk returns a tuple with the AffectedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedRegions

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetAffectedRegions(v int64)`

SetAffectedRegions sets AffectedRegions field to given value.

### HasAffectedRegions

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasAffectedRegions() bool`

HasAffectedRegions returns a boolean if a field has been set.

### GetAffectedSites

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAffectedSites() int64`

GetAffectedSites returns the AffectedSites field if non-nil, zero value otherwise.

### GetAffectedSitesOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAffectedSitesOk() (*int64, bool)`

GetAffectedSitesOk returns a tuple with the AffectedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedSites

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetAffectedSites(v int64)`

SetAffectedSites sets AffectedSites field to given value.

### HasAffectedSites

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasAffectedSites() bool`

HasAffectedSites returns a boolean if a field has been set.

### GetAffectedVrfs

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAffectedVrfs() int64`

GetAffectedVrfs returns the AffectedVrfs field if non-nil, zero value otherwise.

### GetAffectedVrfsOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAffectedVrfsOk() (*int64, bool)`

GetAffectedVrfsOk returns a tuple with the AffectedVrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVrfs

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetAffectedVrfs(v int64)`

SetAffectedVrfs sets AffectedVrfs field to given value.

### HasAffectedVrfs

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasAffectedVrfs() bool`

HasAffectedVrfs returns a boolean if a field has been set.

### GetAppIdKey

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAppIdKey() string`

GetAppIdKey returns the AppIdKey field if non-nil, zero value otherwise.

### GetAppIdKeyOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAppIdKeyOk() (*string, bool)`

GetAppIdKeyOk returns a tuple with the AppIdKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIdKey

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetAppIdKey(v string)`

SetAppIdKey sets AppIdKey field to given value.

### HasAppIdKey

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasAppIdKey() bool`

HasAppIdKey returns a boolean if a field has been set.

### GetAppName

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppType

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetAppType(v string)`

SetAppType sets AppType field to given value.

### HasAppType

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasAppType() bool`

HasAppType returns a boolean if a field has been set.

### GetBlockedReasonList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetBlockedReasonList() []string`

GetBlockedReasonList returns the BlockedReasonList field if non-nil, zero value otherwise.

### GetBlockedReasonListOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetBlockedReasonListOk() (*[]string, bool)`

GetBlockedReasonListOk returns a tuple with the BlockedReasonList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReasonList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetBlockedReasonList(v []string)`

SetBlockedReasonList sets BlockedReasonList field to given value.

### HasBlockedReasonList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasBlockedReasonList() bool`

HasBlockedReasonList returns a boolean if a field has been set.

### GetCategory

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClassficationField

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetClassficationField() string`

GetClassficationField returns the ClassficationField field if non-nil, zero value otherwise.

### GetClassficationFieldOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetClassficationFieldOk() (*string, bool)`

GetClassficationFieldOk returns a tuple with the ClassficationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassficationField

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetClassficationField(v string)`

SetClassficationField sets ClassficationField field to given value.

### HasClassficationField

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasClassficationField() bool`

HasClassficationField returns a boolean if a field has been set.

### GetClassificationField

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetClassificationField() string`

GetClassificationField returns the ClassificationField field if non-nil, zero value otherwise.

### GetClassificationFieldOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetClassificationFieldOk() (*string, bool)`

GetClassificationFieldOk returns a tuple with the ClassificationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationField

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetClassificationField(v string)`

SetClassificationField sets ClassificationField field to given value.

### HasClassificationField

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasClassificationField() bool`

HasClassificationField returns a boolean if a field has been set.

### GetFirstSeen

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetFirstSeen() int64`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetFirstSeenOk() (*int64, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetFirstSeen(v int64)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetFlexAlgo

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetFlexAlgo() []V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecordFlexAlgoInner`

GetFlexAlgo returns the FlexAlgo field if non-nil, zero value otherwise.

### GetFlexAlgoOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetFlexAlgoOk() (*[]V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecordFlexAlgoInner, bool)`

GetFlexAlgoOk returns a tuple with the FlexAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAlgo

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetFlexAlgo(v []V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecordFlexAlgoInner)`

SetFlexAlgo sets FlexAlgo field to given value.

### HasFlexAlgo

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasFlexAlgo() bool`

HasFlexAlgo returns a boolean if a field has been set.

### GetFlowsAnalyzed

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetFlowsAnalyzed() int64`

GetFlowsAnalyzed returns the FlowsAnalyzed field if non-nil, zero value otherwise.

### GetFlowsAnalyzedOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetFlowsAnalyzedOk() (*int64, bool)`

GetFlowsAnalyzedOk returns a tuple with the FlowsAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowsAnalyzed

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetFlowsAnalyzed(v int64)`

SetFlowsAnalyzed sets FlowsAnalyzed field to given value.

### HasFlowsAnalyzed

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasFlowsAnalyzed() bool`

HasFlowsAnalyzed returns a boolean if a field has been set.

### GetLastSeen

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetNewIpHint

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetNewIpHint() bool`

GetNewIpHint returns the NewIpHint field if non-nil, zero value otherwise.

### GetNewIpHintOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetNewIpHintOk() (*bool, bool)`

GetNewIpHintOk returns a tuple with the NewIpHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIpHint

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetNewIpHint(v bool)`

SetNewIpHint sets NewIpHint field to given value.

### HasNewIpHint

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasNewIpHint() bool`

HasNewIpHint returns a boolean if a field has been set.

### GetRecommendation

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetRegionList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetRegionList() []string`

GetRegionList returns the RegionList field if non-nil, zero value otherwise.

### GetRegionListOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetRegionListOk() (*[]string, bool)`

GetRegionListOk returns a tuple with the RegionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetRegionList(v []string)`

SetRegionList sets RegionList field to given value.

### HasRegionList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasRegionList() bool`

HasRegionList returns a boolean if a field has been set.

### GetRiskStatus

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetRiskStatus() string`

GetRiskStatus returns the RiskStatus field if non-nil, zero value otherwise.

### GetRiskStatusOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetRiskStatusOk() (*string, bool)`

GetRiskStatusOk returns a tuple with the RiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskStatus

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetRiskStatus(v string)`

SetRiskStatus sets RiskStatus field to given value.

### HasRiskStatus

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasRiskStatus() bool`

HasRiskStatus returns a boolean if a field has been set.

### GetServerIp

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetServerPort

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetServerPort() string`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetServerPortOk() (*string, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetServerPort(v string)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetSiteList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetSiteList() []string`

GetSiteList returns the SiteList field if non-nil, zero value otherwise.

### GetSiteListOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetSiteListOk() (*[]string, bool)`

GetSiteListOk returns a tuple with the SiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetSiteList(v []string)`

SetSiteList sets SiteList field to given value.

### HasSiteList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasSiteList() bool`

HasSiteList returns a boolean if a field has been set.

### GetThreatScore

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetThreatScore() int64`

GetThreatScore returns the ThreatScore field if non-nil, zero value otherwise.

### GetThreatScoreOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetThreatScoreOk() (*int64, bool)`

GetThreatScoreOk returns a tuple with the ThreatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatScore

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetThreatScore(v int64)`

SetThreatScore sets ThreatScore field to given value.

### HasThreatScore

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasThreatScore() bool`

HasThreatScore returns a boolean if a field has been set.

### GetVrfList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetVrfList() []string`

GetVrfList returns the VrfList field if non-nil, zero value otherwise.

### GetVrfListOk

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) GetVrfListOk() (*[]string, bool)`

GetVrfListOk returns a tuple with the VrfList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) SetVrfList(v []string)`

SetVrfList sets VrfList field to given value.

### HasVrfList

`func (o *V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord) HasVrfList() bool`

HasVrfList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


