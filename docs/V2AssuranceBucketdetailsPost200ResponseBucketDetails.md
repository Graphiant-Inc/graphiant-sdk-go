# V2AssuranceBucketdetailsPost200ResponseBucketDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppCountThreatHigh** | Pointer to **int64** |  | [optional] 
**AppCountThreatLow** | Pointer to **int64** |  | [optional] 
**AppCountThreatMedium** | Pointer to **int64** |  | [optional] 
**AppIdRecords** | Pointer to [**[]V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord**](V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord.md) |  | [optional] 
**AppNameRecords** | Pointer to [**[]V2AssuranceApplicationdetailsbynamePost200ResponseAppNameRecord**](V2AssuranceApplicationdetailsbynamePost200ResponseAppNameRecord.md) |  | [optional] 
**BucketNameToDisplay** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayIpPort** | Pointer to **bool** |  | [optional] 
**FlowCount** | Pointer to **int64** |  | [optional] 
**NewIpHint** | Pointer to **bool** |  | [optional] 
**Recommendation** | Pointer to **string** |  | [optional] 
**TrendValueList** | Pointer to [**[]V2AssuranceBucketdetailsPost200ResponseBucketDetailsTrendValueListInner**](V2AssuranceBucketdetailsPost200ResponseBucketDetailsTrendValueListInner.md) |  | [optional] 
**UniqueAppCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewV2AssuranceBucketdetailsPost200ResponseBucketDetails

`func NewV2AssuranceBucketdetailsPost200ResponseBucketDetails() *V2AssuranceBucketdetailsPost200ResponseBucketDetails`

NewV2AssuranceBucketdetailsPost200ResponseBucketDetails instantiates a new V2AssuranceBucketdetailsPost200ResponseBucketDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceBucketdetailsPost200ResponseBucketDetailsWithDefaults

`func NewV2AssuranceBucketdetailsPost200ResponseBucketDetailsWithDefaults() *V2AssuranceBucketdetailsPost200ResponseBucketDetails`

NewV2AssuranceBucketdetailsPost200ResponseBucketDetailsWithDefaults instantiates a new V2AssuranceBucketdetailsPost200ResponseBucketDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppCountThreatHigh

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetAppCountThreatHigh() int64`

GetAppCountThreatHigh returns the AppCountThreatHigh field if non-nil, zero value otherwise.

### GetAppCountThreatHighOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetAppCountThreatHighOk() (*int64, bool)`

GetAppCountThreatHighOk returns a tuple with the AppCountThreatHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCountThreatHigh

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetAppCountThreatHigh(v int64)`

SetAppCountThreatHigh sets AppCountThreatHigh field to given value.

### HasAppCountThreatHigh

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasAppCountThreatHigh() bool`

HasAppCountThreatHigh returns a boolean if a field has been set.

### GetAppCountThreatLow

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetAppCountThreatLow() int64`

GetAppCountThreatLow returns the AppCountThreatLow field if non-nil, zero value otherwise.

### GetAppCountThreatLowOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetAppCountThreatLowOk() (*int64, bool)`

GetAppCountThreatLowOk returns a tuple with the AppCountThreatLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCountThreatLow

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetAppCountThreatLow(v int64)`

SetAppCountThreatLow sets AppCountThreatLow field to given value.

### HasAppCountThreatLow

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasAppCountThreatLow() bool`

HasAppCountThreatLow returns a boolean if a field has been set.

### GetAppCountThreatMedium

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetAppCountThreatMedium() int64`

GetAppCountThreatMedium returns the AppCountThreatMedium field if non-nil, zero value otherwise.

### GetAppCountThreatMediumOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetAppCountThreatMediumOk() (*int64, bool)`

GetAppCountThreatMediumOk returns a tuple with the AppCountThreatMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCountThreatMedium

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetAppCountThreatMedium(v int64)`

SetAppCountThreatMedium sets AppCountThreatMedium field to given value.

### HasAppCountThreatMedium

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasAppCountThreatMedium() bool`

HasAppCountThreatMedium returns a boolean if a field has been set.

### GetAppIdRecords

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetAppIdRecords() []V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord`

GetAppIdRecords returns the AppIdRecords field if non-nil, zero value otherwise.

### GetAppIdRecordsOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetAppIdRecordsOk() (*[]V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord, bool)`

GetAppIdRecordsOk returns a tuple with the AppIdRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIdRecords

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetAppIdRecords(v []V2AssuranceApplicationdetailsbynamePost200ResponseAppIdRecord)`

SetAppIdRecords sets AppIdRecords field to given value.

### HasAppIdRecords

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasAppIdRecords() bool`

HasAppIdRecords returns a boolean if a field has been set.

### GetAppNameRecords

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetAppNameRecords() []V2AssuranceApplicationdetailsbynamePost200ResponseAppNameRecord`

GetAppNameRecords returns the AppNameRecords field if non-nil, zero value otherwise.

### GetAppNameRecordsOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetAppNameRecordsOk() (*[]V2AssuranceApplicationdetailsbynamePost200ResponseAppNameRecord, bool)`

GetAppNameRecordsOk returns a tuple with the AppNameRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppNameRecords

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetAppNameRecords(v []V2AssuranceApplicationdetailsbynamePost200ResponseAppNameRecord)`

SetAppNameRecords sets AppNameRecords field to given value.

### HasAppNameRecords

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasAppNameRecords() bool`

HasAppNameRecords returns a boolean if a field has been set.

### GetBucketNameToDisplay

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetBucketNameToDisplay() string`

GetBucketNameToDisplay returns the BucketNameToDisplay field if non-nil, zero value otherwise.

### GetBucketNameToDisplayOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetBucketNameToDisplayOk() (*string, bool)`

GetBucketNameToDisplayOk returns a tuple with the BucketNameToDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketNameToDisplay

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetBucketNameToDisplay(v string)`

SetBucketNameToDisplay sets BucketNameToDisplay field to given value.

### HasBucketNameToDisplay

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasBucketNameToDisplay() bool`

HasBucketNameToDisplay returns a boolean if a field has been set.

### GetDescription

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayIpPort

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetDisplayIpPort() bool`

GetDisplayIpPort returns the DisplayIpPort field if non-nil, zero value otherwise.

### GetDisplayIpPortOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetDisplayIpPortOk() (*bool, bool)`

GetDisplayIpPortOk returns a tuple with the DisplayIpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIpPort

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetDisplayIpPort(v bool)`

SetDisplayIpPort sets DisplayIpPort field to given value.

### HasDisplayIpPort

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasDisplayIpPort() bool`

HasDisplayIpPort returns a boolean if a field has been set.

### GetFlowCount

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetFlowCount() int64`

GetFlowCount returns the FlowCount field if non-nil, zero value otherwise.

### GetFlowCountOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetFlowCountOk() (*int64, bool)`

GetFlowCountOk returns a tuple with the FlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowCount

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetFlowCount(v int64)`

SetFlowCount sets FlowCount field to given value.

### HasFlowCount

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasFlowCount() bool`

HasFlowCount returns a boolean if a field has been set.

### GetNewIpHint

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetNewIpHint() bool`

GetNewIpHint returns the NewIpHint field if non-nil, zero value otherwise.

### GetNewIpHintOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetNewIpHintOk() (*bool, bool)`

GetNewIpHintOk returns a tuple with the NewIpHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIpHint

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetNewIpHint(v bool)`

SetNewIpHint sets NewIpHint field to given value.

### HasNewIpHint

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasNewIpHint() bool`

HasNewIpHint returns a boolean if a field has been set.

### GetRecommendation

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetTrendValueList

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetTrendValueList() []V2AssuranceBucketdetailsPost200ResponseBucketDetailsTrendValueListInner`

GetTrendValueList returns the TrendValueList field if non-nil, zero value otherwise.

### GetTrendValueListOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetTrendValueListOk() (*[]V2AssuranceBucketdetailsPost200ResponseBucketDetailsTrendValueListInner, bool)`

GetTrendValueListOk returns a tuple with the TrendValueList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendValueList

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetTrendValueList(v []V2AssuranceBucketdetailsPost200ResponseBucketDetailsTrendValueListInner)`

SetTrendValueList sets TrendValueList field to given value.

### HasTrendValueList

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasTrendValueList() bool`

HasTrendValueList returns a boolean if a field has been set.

### GetUniqueAppCount

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetUniqueAppCount() int64`

GetUniqueAppCount returns the UniqueAppCount field if non-nil, zero value otherwise.

### GetUniqueAppCountOk

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) GetUniqueAppCountOk() (*int64, bool)`

GetUniqueAppCountOk returns a tuple with the UniqueAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueAppCount

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) SetUniqueAppCount(v int64)`

SetUniqueAppCount sets UniqueAppCount field to given value.

### HasUniqueAppCount

`func (o *V2AssuranceBucketdetailsPost200ResponseBucketDetails) HasUniqueAppCount() bool`

HasUniqueAppCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


