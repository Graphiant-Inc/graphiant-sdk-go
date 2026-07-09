# AssuranceAppNameRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedHosts** | Pointer to **int64** |  | [optional] 
**AffectedRegions** | Pointer to **int64** |  | [optional] 
**AffectedSites** | Pointer to **int64** |  | [optional] 
**AffectedVrfs** | Pointer to **int64** |  | [optional] 
**AppAiTag** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **int64** |  | [optional] 
**AppIdRecords** | Pointer to [**[]AssuranceAppIdRecord**](AssuranceAppIdRecord.md) |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**AppType** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**DaClassified** | Pointer to **bool** |  | [optional] 
**ExchangeService** | Pointer to [**[]AssuranceExchangeServiceIdentifier**](AssuranceExchangeServiceIdentifier.md) |  | [optional] 
**FlexAlgo** | Pointer to [**[]AssuranceFlexAlgoIdentifier**](AssuranceFlexAlgoIdentifier.md) |  | [optional] 
**FlowsAnalyzed** | Pointer to **int64** |  | [optional] 
**Recommendation** | Pointer to **string** |  | [optional] 
**RiskStatus** | Pointer to **string** |  | [optional] 
**ThreatScore** | Pointer to **int64** |  | [optional] 

## Methods

### NewAssuranceAppNameRecord

`func NewAssuranceAppNameRecord() *AssuranceAppNameRecord`

NewAssuranceAppNameRecord instantiates a new AssuranceAppNameRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssuranceAppNameRecordWithDefaults

`func NewAssuranceAppNameRecordWithDefaults() *AssuranceAppNameRecord`

NewAssuranceAppNameRecordWithDefaults instantiates a new AssuranceAppNameRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedHosts

`func (o *AssuranceAppNameRecord) GetAffectedHosts() int64`

GetAffectedHosts returns the AffectedHosts field if non-nil, zero value otherwise.

### GetAffectedHostsOk

`func (o *AssuranceAppNameRecord) GetAffectedHostsOk() (*int64, bool)`

GetAffectedHostsOk returns a tuple with the AffectedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedHosts

`func (o *AssuranceAppNameRecord) SetAffectedHosts(v int64)`

SetAffectedHosts sets AffectedHosts field to given value.

### HasAffectedHosts

`func (o *AssuranceAppNameRecord) HasAffectedHosts() bool`

HasAffectedHosts returns a boolean if a field has been set.

### GetAffectedRegions

`func (o *AssuranceAppNameRecord) GetAffectedRegions() int64`

GetAffectedRegions returns the AffectedRegions field if non-nil, zero value otherwise.

### GetAffectedRegionsOk

`func (o *AssuranceAppNameRecord) GetAffectedRegionsOk() (*int64, bool)`

GetAffectedRegionsOk returns a tuple with the AffectedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedRegions

`func (o *AssuranceAppNameRecord) SetAffectedRegions(v int64)`

SetAffectedRegions sets AffectedRegions field to given value.

### HasAffectedRegions

`func (o *AssuranceAppNameRecord) HasAffectedRegions() bool`

HasAffectedRegions returns a boolean if a field has been set.

### GetAffectedSites

`func (o *AssuranceAppNameRecord) GetAffectedSites() int64`

GetAffectedSites returns the AffectedSites field if non-nil, zero value otherwise.

### GetAffectedSitesOk

`func (o *AssuranceAppNameRecord) GetAffectedSitesOk() (*int64, bool)`

GetAffectedSitesOk returns a tuple with the AffectedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedSites

`func (o *AssuranceAppNameRecord) SetAffectedSites(v int64)`

SetAffectedSites sets AffectedSites field to given value.

### HasAffectedSites

`func (o *AssuranceAppNameRecord) HasAffectedSites() bool`

HasAffectedSites returns a boolean if a field has been set.

### GetAffectedVrfs

`func (o *AssuranceAppNameRecord) GetAffectedVrfs() int64`

GetAffectedVrfs returns the AffectedVrfs field if non-nil, zero value otherwise.

### GetAffectedVrfsOk

`func (o *AssuranceAppNameRecord) GetAffectedVrfsOk() (*int64, bool)`

GetAffectedVrfsOk returns a tuple with the AffectedVrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVrfs

`func (o *AssuranceAppNameRecord) SetAffectedVrfs(v int64)`

SetAffectedVrfs sets AffectedVrfs field to given value.

### HasAffectedVrfs

`func (o *AssuranceAppNameRecord) HasAffectedVrfs() bool`

HasAffectedVrfs returns a boolean if a field has been set.

### GetAppAiTag

`func (o *AssuranceAppNameRecord) GetAppAiTag() string`

GetAppAiTag returns the AppAiTag field if non-nil, zero value otherwise.

### GetAppAiTagOk

`func (o *AssuranceAppNameRecord) GetAppAiTagOk() (*string, bool)`

GetAppAiTagOk returns a tuple with the AppAiTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAiTag

`func (o *AssuranceAppNameRecord) SetAppAiTag(v string)`

SetAppAiTag sets AppAiTag field to given value.

### HasAppAiTag

`func (o *AssuranceAppNameRecord) HasAppAiTag() bool`

HasAppAiTag returns a boolean if a field has been set.

### GetAppId

`func (o *AssuranceAppNameRecord) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AssuranceAppNameRecord) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AssuranceAppNameRecord) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AssuranceAppNameRecord) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppIdRecords

`func (o *AssuranceAppNameRecord) GetAppIdRecords() []AssuranceAppIdRecord`

GetAppIdRecords returns the AppIdRecords field if non-nil, zero value otherwise.

### GetAppIdRecordsOk

`func (o *AssuranceAppNameRecord) GetAppIdRecordsOk() (*[]AssuranceAppIdRecord, bool)`

GetAppIdRecordsOk returns a tuple with the AppIdRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIdRecords

`func (o *AssuranceAppNameRecord) SetAppIdRecords(v []AssuranceAppIdRecord)`

SetAppIdRecords sets AppIdRecords field to given value.

### HasAppIdRecords

`func (o *AssuranceAppNameRecord) HasAppIdRecords() bool`

HasAppIdRecords returns a boolean if a field has been set.

### GetAppName

`func (o *AssuranceAppNameRecord) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AssuranceAppNameRecord) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AssuranceAppNameRecord) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *AssuranceAppNameRecord) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppType

`func (o *AssuranceAppNameRecord) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *AssuranceAppNameRecord) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *AssuranceAppNameRecord) SetAppType(v string)`

SetAppType sets AppType field to given value.

### HasAppType

`func (o *AssuranceAppNameRecord) HasAppType() bool`

HasAppType returns a boolean if a field has been set.

### GetCategory

`func (o *AssuranceAppNameRecord) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AssuranceAppNameRecord) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AssuranceAppNameRecord) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AssuranceAppNameRecord) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDaClassified

`func (o *AssuranceAppNameRecord) GetDaClassified() bool`

GetDaClassified returns the DaClassified field if non-nil, zero value otherwise.

### GetDaClassifiedOk

`func (o *AssuranceAppNameRecord) GetDaClassifiedOk() (*bool, bool)`

GetDaClassifiedOk returns a tuple with the DaClassified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaClassified

`func (o *AssuranceAppNameRecord) SetDaClassified(v bool)`

SetDaClassified sets DaClassified field to given value.

### HasDaClassified

`func (o *AssuranceAppNameRecord) HasDaClassified() bool`

HasDaClassified returns a boolean if a field has been set.

### GetExchangeService

`func (o *AssuranceAppNameRecord) GetExchangeService() []AssuranceExchangeServiceIdentifier`

GetExchangeService returns the ExchangeService field if non-nil, zero value otherwise.

### GetExchangeServiceOk

`func (o *AssuranceAppNameRecord) GetExchangeServiceOk() (*[]AssuranceExchangeServiceIdentifier, bool)`

GetExchangeServiceOk returns a tuple with the ExchangeService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeService

`func (o *AssuranceAppNameRecord) SetExchangeService(v []AssuranceExchangeServiceIdentifier)`

SetExchangeService sets ExchangeService field to given value.

### HasExchangeService

`func (o *AssuranceAppNameRecord) HasExchangeService() bool`

HasExchangeService returns a boolean if a field has been set.

### GetFlexAlgo

`func (o *AssuranceAppNameRecord) GetFlexAlgo() []AssuranceFlexAlgoIdentifier`

GetFlexAlgo returns the FlexAlgo field if non-nil, zero value otherwise.

### GetFlexAlgoOk

`func (o *AssuranceAppNameRecord) GetFlexAlgoOk() (*[]AssuranceFlexAlgoIdentifier, bool)`

GetFlexAlgoOk returns a tuple with the FlexAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAlgo

`func (o *AssuranceAppNameRecord) SetFlexAlgo(v []AssuranceFlexAlgoIdentifier)`

SetFlexAlgo sets FlexAlgo field to given value.

### HasFlexAlgo

`func (o *AssuranceAppNameRecord) HasFlexAlgo() bool`

HasFlexAlgo returns a boolean if a field has been set.

### GetFlowsAnalyzed

`func (o *AssuranceAppNameRecord) GetFlowsAnalyzed() int64`

GetFlowsAnalyzed returns the FlowsAnalyzed field if non-nil, zero value otherwise.

### GetFlowsAnalyzedOk

`func (o *AssuranceAppNameRecord) GetFlowsAnalyzedOk() (*int64, bool)`

GetFlowsAnalyzedOk returns a tuple with the FlowsAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowsAnalyzed

`func (o *AssuranceAppNameRecord) SetFlowsAnalyzed(v int64)`

SetFlowsAnalyzed sets FlowsAnalyzed field to given value.

### HasFlowsAnalyzed

`func (o *AssuranceAppNameRecord) HasFlowsAnalyzed() bool`

HasFlowsAnalyzed returns a boolean if a field has been set.

### GetRecommendation

`func (o *AssuranceAppNameRecord) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *AssuranceAppNameRecord) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *AssuranceAppNameRecord) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *AssuranceAppNameRecord) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetRiskStatus

`func (o *AssuranceAppNameRecord) GetRiskStatus() string`

GetRiskStatus returns the RiskStatus field if non-nil, zero value otherwise.

### GetRiskStatusOk

`func (o *AssuranceAppNameRecord) GetRiskStatusOk() (*string, bool)`

GetRiskStatusOk returns a tuple with the RiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskStatus

`func (o *AssuranceAppNameRecord) SetRiskStatus(v string)`

SetRiskStatus sets RiskStatus field to given value.

### HasRiskStatus

`func (o *AssuranceAppNameRecord) HasRiskStatus() bool`

HasRiskStatus returns a boolean if a field has been set.

### GetThreatScore

`func (o *AssuranceAppNameRecord) GetThreatScore() int64`

GetThreatScore returns the ThreatScore field if non-nil, zero value otherwise.

### GetThreatScoreOk

`func (o *AssuranceAppNameRecord) GetThreatScoreOk() (*int64, bool)`

GetThreatScoreOk returns a tuple with the ThreatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatScore

`func (o *AssuranceAppNameRecord) SetThreatScore(v int64)`

SetThreatScore sets ThreatScore field to given value.

### HasThreatScore

`func (o *AssuranceAppNameRecord) HasThreatScore() bool`

HasThreatScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


