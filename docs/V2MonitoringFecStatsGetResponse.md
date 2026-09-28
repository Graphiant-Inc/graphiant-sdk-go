# V2MonitoringFecStatsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentRepairLevel** | Pointer to **float32** |  | [optional] 
**EffectiveQoe** | Pointer to **float32** |  | [optional] 
**RxRepairLevelHistogram** | Pointer to [**StatsmonV2FecRxRepairLevel**](StatsmonV2FecRxRepairLevel.md) |  | [optional] 
**RxStats** | Pointer to [**StatsmonV2FecRxStats**](StatsmonV2FecRxStats.md) |  | [optional] 
**TxRepairLevelHistogram** | Pointer to [**StatsmonV2FecTxRepairLevel**](StatsmonV2FecTxRepairLevel.md) |  | [optional] 
**TxStats** | Pointer to [**StatsmonV2FecTxStats**](StatsmonV2FecTxStats.md) |  | [optional] 
**UnrepairableRate** | Pointer to **float32** |  | [optional] 

## Methods

### NewV2MonitoringFecStatsGetResponse

`func NewV2MonitoringFecStatsGetResponse() *V2MonitoringFecStatsGetResponse`

NewV2MonitoringFecStatsGetResponse instantiates a new V2MonitoringFecStatsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringFecStatsGetResponseWithDefaults

`func NewV2MonitoringFecStatsGetResponseWithDefaults() *V2MonitoringFecStatsGetResponse`

NewV2MonitoringFecStatsGetResponseWithDefaults instantiates a new V2MonitoringFecStatsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentRepairLevel

`func (o *V2MonitoringFecStatsGetResponse) GetCurrentRepairLevel() float32`

GetCurrentRepairLevel returns the CurrentRepairLevel field if non-nil, zero value otherwise.

### GetCurrentRepairLevelOk

`func (o *V2MonitoringFecStatsGetResponse) GetCurrentRepairLevelOk() (*float32, bool)`

GetCurrentRepairLevelOk returns a tuple with the CurrentRepairLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRepairLevel

`func (o *V2MonitoringFecStatsGetResponse) SetCurrentRepairLevel(v float32)`

SetCurrentRepairLevel sets CurrentRepairLevel field to given value.

### HasCurrentRepairLevel

`func (o *V2MonitoringFecStatsGetResponse) HasCurrentRepairLevel() bool`

HasCurrentRepairLevel returns a boolean if a field has been set.

### GetEffectiveQoe

`func (o *V2MonitoringFecStatsGetResponse) GetEffectiveQoe() float32`

GetEffectiveQoe returns the EffectiveQoe field if non-nil, zero value otherwise.

### GetEffectiveQoeOk

`func (o *V2MonitoringFecStatsGetResponse) GetEffectiveQoeOk() (*float32, bool)`

GetEffectiveQoeOk returns a tuple with the EffectiveQoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveQoe

`func (o *V2MonitoringFecStatsGetResponse) SetEffectiveQoe(v float32)`

SetEffectiveQoe sets EffectiveQoe field to given value.

### HasEffectiveQoe

`func (o *V2MonitoringFecStatsGetResponse) HasEffectiveQoe() bool`

HasEffectiveQoe returns a boolean if a field has been set.

### GetRxRepairLevelHistogram

`func (o *V2MonitoringFecStatsGetResponse) GetRxRepairLevelHistogram() StatsmonV2FecRxRepairLevel`

GetRxRepairLevelHistogram returns the RxRepairLevelHistogram field if non-nil, zero value otherwise.

### GetRxRepairLevelHistogramOk

`func (o *V2MonitoringFecStatsGetResponse) GetRxRepairLevelHistogramOk() (*StatsmonV2FecRxRepairLevel, bool)`

GetRxRepairLevelHistogramOk returns a tuple with the RxRepairLevelHistogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRepairLevelHistogram

`func (o *V2MonitoringFecStatsGetResponse) SetRxRepairLevelHistogram(v StatsmonV2FecRxRepairLevel)`

SetRxRepairLevelHistogram sets RxRepairLevelHistogram field to given value.

### HasRxRepairLevelHistogram

`func (o *V2MonitoringFecStatsGetResponse) HasRxRepairLevelHistogram() bool`

HasRxRepairLevelHistogram returns a boolean if a field has been set.

### GetRxStats

`func (o *V2MonitoringFecStatsGetResponse) GetRxStats() StatsmonV2FecRxStats`

GetRxStats returns the RxStats field if non-nil, zero value otherwise.

### GetRxStatsOk

`func (o *V2MonitoringFecStatsGetResponse) GetRxStatsOk() (*StatsmonV2FecRxStats, bool)`

GetRxStatsOk returns a tuple with the RxStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxStats

`func (o *V2MonitoringFecStatsGetResponse) SetRxStats(v StatsmonV2FecRxStats)`

SetRxStats sets RxStats field to given value.

### HasRxStats

`func (o *V2MonitoringFecStatsGetResponse) HasRxStats() bool`

HasRxStats returns a boolean if a field has been set.

### GetTxRepairLevelHistogram

`func (o *V2MonitoringFecStatsGetResponse) GetTxRepairLevelHistogram() StatsmonV2FecTxRepairLevel`

GetTxRepairLevelHistogram returns the TxRepairLevelHistogram field if non-nil, zero value otherwise.

### GetTxRepairLevelHistogramOk

`func (o *V2MonitoringFecStatsGetResponse) GetTxRepairLevelHistogramOk() (*StatsmonV2FecTxRepairLevel, bool)`

GetTxRepairLevelHistogramOk returns a tuple with the TxRepairLevelHistogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRepairLevelHistogram

`func (o *V2MonitoringFecStatsGetResponse) SetTxRepairLevelHistogram(v StatsmonV2FecTxRepairLevel)`

SetTxRepairLevelHistogram sets TxRepairLevelHistogram field to given value.

### HasTxRepairLevelHistogram

`func (o *V2MonitoringFecStatsGetResponse) HasTxRepairLevelHistogram() bool`

HasTxRepairLevelHistogram returns a boolean if a field has been set.

### GetTxStats

`func (o *V2MonitoringFecStatsGetResponse) GetTxStats() StatsmonV2FecTxStats`

GetTxStats returns the TxStats field if non-nil, zero value otherwise.

### GetTxStatsOk

`func (o *V2MonitoringFecStatsGetResponse) GetTxStatsOk() (*StatsmonV2FecTxStats, bool)`

GetTxStatsOk returns a tuple with the TxStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStats

`func (o *V2MonitoringFecStatsGetResponse) SetTxStats(v StatsmonV2FecTxStats)`

SetTxStats sets TxStats field to given value.

### HasTxStats

`func (o *V2MonitoringFecStatsGetResponse) HasTxStats() bool`

HasTxStats returns a boolean if a field has been set.

### GetUnrepairableRate

`func (o *V2MonitoringFecStatsGetResponse) GetUnrepairableRate() float32`

GetUnrepairableRate returns the UnrepairableRate field if non-nil, zero value otherwise.

### GetUnrepairableRateOk

`func (o *V2MonitoringFecStatsGetResponse) GetUnrepairableRateOk() (*float32, bool)`

GetUnrepairableRateOk returns a tuple with the UnrepairableRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrepairableRate

`func (o *V2MonitoringFecStatsGetResponse) SetUnrepairableRate(v float32)`

SetUnrepairableRate sets UnrepairableRate field to given value.

### HasUnrepairableRate

`func (o *V2MonitoringFecStatsGetResponse) HasUnrepairableRate() bool`

HasUnrepairableRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


