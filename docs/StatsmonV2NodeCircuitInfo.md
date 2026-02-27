# StatsmonV2NodeCircuitInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageDownlinkUtilization** | Pointer to **float64** |  | [optional] 
**AverageUplinkUtilization** | Pointer to **float64** |  | [optional] 
**CircuitCarrier** | Pointer to **string** |  | [optional] 
**CircuitName** | Pointer to **string** |  | [optional] 
**ConnectedRegionName** | Pointer to **string** |  | [optional] 
**CurrentDownlinkUtilization** | Pointer to **float64** |  | [optional] 
**CurrentUplinkUtilization** | Pointer to **float64** |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**Jitter** | Pointer to **int64** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LastResort** | Pointer to **bool** |  | [optional] 
**Latency** | Pointer to **int64** |  | [optional] 
**Loss** | Pointer to **float32** |  | [optional] 
**Qoe** | Pointer to **float32** |  | [optional] 
**Quality** | Pointer to **string** |  | [optional] 

## Methods

### NewStatsmonV2NodeCircuitInfo

`func NewStatsmonV2NodeCircuitInfo() *StatsmonV2NodeCircuitInfo`

NewStatsmonV2NodeCircuitInfo instantiates a new StatsmonV2NodeCircuitInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsmonV2NodeCircuitInfoWithDefaults

`func NewStatsmonV2NodeCircuitInfoWithDefaults() *StatsmonV2NodeCircuitInfo`

NewStatsmonV2NodeCircuitInfoWithDefaults instantiates a new StatsmonV2NodeCircuitInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageDownlinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) GetAverageDownlinkUtilization() float64`

GetAverageDownlinkUtilization returns the AverageDownlinkUtilization field if non-nil, zero value otherwise.

### GetAverageDownlinkUtilizationOk

`func (o *StatsmonV2NodeCircuitInfo) GetAverageDownlinkUtilizationOk() (*float64, bool)`

GetAverageDownlinkUtilizationOk returns a tuple with the AverageDownlinkUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDownlinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) SetAverageDownlinkUtilization(v float64)`

SetAverageDownlinkUtilization sets AverageDownlinkUtilization field to given value.

### HasAverageDownlinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) HasAverageDownlinkUtilization() bool`

HasAverageDownlinkUtilization returns a boolean if a field has been set.

### GetAverageUplinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) GetAverageUplinkUtilization() float64`

GetAverageUplinkUtilization returns the AverageUplinkUtilization field if non-nil, zero value otherwise.

### GetAverageUplinkUtilizationOk

`func (o *StatsmonV2NodeCircuitInfo) GetAverageUplinkUtilizationOk() (*float64, bool)`

GetAverageUplinkUtilizationOk returns a tuple with the AverageUplinkUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageUplinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) SetAverageUplinkUtilization(v float64)`

SetAverageUplinkUtilization sets AverageUplinkUtilization field to given value.

### HasAverageUplinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) HasAverageUplinkUtilization() bool`

HasAverageUplinkUtilization returns a boolean if a field has been set.

### GetCircuitCarrier

`func (o *StatsmonV2NodeCircuitInfo) GetCircuitCarrier() string`

GetCircuitCarrier returns the CircuitCarrier field if non-nil, zero value otherwise.

### GetCircuitCarrierOk

`func (o *StatsmonV2NodeCircuitInfo) GetCircuitCarrierOk() (*string, bool)`

GetCircuitCarrierOk returns a tuple with the CircuitCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitCarrier

`func (o *StatsmonV2NodeCircuitInfo) SetCircuitCarrier(v string)`

SetCircuitCarrier sets CircuitCarrier field to given value.

### HasCircuitCarrier

`func (o *StatsmonV2NodeCircuitInfo) HasCircuitCarrier() bool`

HasCircuitCarrier returns a boolean if a field has been set.

### GetCircuitName

`func (o *StatsmonV2NodeCircuitInfo) GetCircuitName() string`

GetCircuitName returns the CircuitName field if non-nil, zero value otherwise.

### GetCircuitNameOk

`func (o *StatsmonV2NodeCircuitInfo) GetCircuitNameOk() (*string, bool)`

GetCircuitNameOk returns a tuple with the CircuitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitName

`func (o *StatsmonV2NodeCircuitInfo) SetCircuitName(v string)`

SetCircuitName sets CircuitName field to given value.

### HasCircuitName

`func (o *StatsmonV2NodeCircuitInfo) HasCircuitName() bool`

HasCircuitName returns a boolean if a field has been set.

### GetConnectedRegionName

`func (o *StatsmonV2NodeCircuitInfo) GetConnectedRegionName() string`

GetConnectedRegionName returns the ConnectedRegionName field if non-nil, zero value otherwise.

### GetConnectedRegionNameOk

`func (o *StatsmonV2NodeCircuitInfo) GetConnectedRegionNameOk() (*string, bool)`

GetConnectedRegionNameOk returns a tuple with the ConnectedRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedRegionName

`func (o *StatsmonV2NodeCircuitInfo) SetConnectedRegionName(v string)`

SetConnectedRegionName sets ConnectedRegionName field to given value.

### HasConnectedRegionName

`func (o *StatsmonV2NodeCircuitInfo) HasConnectedRegionName() bool`

HasConnectedRegionName returns a boolean if a field has been set.

### GetCurrentDownlinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) GetCurrentDownlinkUtilization() float64`

GetCurrentDownlinkUtilization returns the CurrentDownlinkUtilization field if non-nil, zero value otherwise.

### GetCurrentDownlinkUtilizationOk

`func (o *StatsmonV2NodeCircuitInfo) GetCurrentDownlinkUtilizationOk() (*float64, bool)`

GetCurrentDownlinkUtilizationOk returns a tuple with the CurrentDownlinkUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDownlinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) SetCurrentDownlinkUtilization(v float64)`

SetCurrentDownlinkUtilization sets CurrentDownlinkUtilization field to given value.

### HasCurrentDownlinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) HasCurrentDownlinkUtilization() bool`

HasCurrentDownlinkUtilization returns a boolean if a field has been set.

### GetCurrentUplinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) GetCurrentUplinkUtilization() float64`

GetCurrentUplinkUtilization returns the CurrentUplinkUtilization field if non-nil, zero value otherwise.

### GetCurrentUplinkUtilizationOk

`func (o *StatsmonV2NodeCircuitInfo) GetCurrentUplinkUtilizationOk() (*float64, bool)`

GetCurrentUplinkUtilizationOk returns a tuple with the CurrentUplinkUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUplinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) SetCurrentUplinkUtilization(v float64)`

SetCurrentUplinkUtilization sets CurrentUplinkUtilization field to given value.

### HasCurrentUplinkUtilization

`func (o *StatsmonV2NodeCircuitInfo) HasCurrentUplinkUtilization() bool`

HasCurrentUplinkUtilization returns a boolean if a field has been set.

### GetDeviceId

`func (o *StatsmonV2NodeCircuitInfo) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *StatsmonV2NodeCircuitInfo) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *StatsmonV2NodeCircuitInfo) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *StatsmonV2NodeCircuitInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *StatsmonV2NodeCircuitInfo) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *StatsmonV2NodeCircuitInfo) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *StatsmonV2NodeCircuitInfo) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *StatsmonV2NodeCircuitInfo) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetJitter

`func (o *StatsmonV2NodeCircuitInfo) GetJitter() int64`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *StatsmonV2NodeCircuitInfo) GetJitterOk() (*int64, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *StatsmonV2NodeCircuitInfo) SetJitter(v int64)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *StatsmonV2NodeCircuitInfo) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### GetLabel

`func (o *StatsmonV2NodeCircuitInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StatsmonV2NodeCircuitInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StatsmonV2NodeCircuitInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *StatsmonV2NodeCircuitInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastResort

`func (o *StatsmonV2NodeCircuitInfo) GetLastResort() bool`

GetLastResort returns the LastResort field if non-nil, zero value otherwise.

### GetLastResortOk

`func (o *StatsmonV2NodeCircuitInfo) GetLastResortOk() (*bool, bool)`

GetLastResortOk returns a tuple with the LastResort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResort

`func (o *StatsmonV2NodeCircuitInfo) SetLastResort(v bool)`

SetLastResort sets LastResort field to given value.

### HasLastResort

`func (o *StatsmonV2NodeCircuitInfo) HasLastResort() bool`

HasLastResort returns a boolean if a field has been set.

### GetLatency

`func (o *StatsmonV2NodeCircuitInfo) GetLatency() int64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *StatsmonV2NodeCircuitInfo) GetLatencyOk() (*int64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *StatsmonV2NodeCircuitInfo) SetLatency(v int64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *StatsmonV2NodeCircuitInfo) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLoss

`func (o *StatsmonV2NodeCircuitInfo) GetLoss() float32`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *StatsmonV2NodeCircuitInfo) GetLossOk() (*float32, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *StatsmonV2NodeCircuitInfo) SetLoss(v float32)`

SetLoss sets Loss field to given value.

### HasLoss

`func (o *StatsmonV2NodeCircuitInfo) HasLoss() bool`

HasLoss returns a boolean if a field has been set.

### GetQoe

`func (o *StatsmonV2NodeCircuitInfo) GetQoe() float32`

GetQoe returns the Qoe field if non-nil, zero value otherwise.

### GetQoeOk

`func (o *StatsmonV2NodeCircuitInfo) GetQoeOk() (*float32, bool)`

GetQoeOk returns a tuple with the Qoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoe

`func (o *StatsmonV2NodeCircuitInfo) SetQoe(v float32)`

SetQoe sets Qoe field to given value.

### HasQoe

`func (o *StatsmonV2NodeCircuitInfo) HasQoe() bool`

HasQoe returns a boolean if a field has been set.

### GetQuality

`func (o *StatsmonV2NodeCircuitInfo) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *StatsmonV2NodeCircuitInfo) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *StatsmonV2NodeCircuitInfo) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *StatsmonV2NodeCircuitInfo) HasQuality() bool`

HasQuality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


