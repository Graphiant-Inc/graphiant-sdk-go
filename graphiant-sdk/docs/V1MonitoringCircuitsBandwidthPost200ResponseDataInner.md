# V1MonitoringCircuitsBandwidthPost200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Carrier** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**DlBwKbpsSamples** | Pointer to [**[]V2MonitoringBfdPost200ResponseDataInnerSamplesInner**](V2MonitoringBfdPost200ResponseDataInnerSamplesInner.md) |  | [optional] 
**Selector** | Pointer to [**V1MonitoringCircuitsBandwidthPostRequestSelectorsInner**](V1MonitoringCircuitsBandwidthPostRequestSelectorsInner.md) |  | [optional] 
**UlBwKbpsSamples** | Pointer to [**[]V2MonitoringBfdPost200ResponseDataInnerSamplesInner**](V2MonitoringBfdPost200ResponseDataInnerSamplesInner.md) |  | [optional] 

## Methods

### NewV1MonitoringCircuitsBandwidthPost200ResponseDataInner

`func NewV1MonitoringCircuitsBandwidthPost200ResponseDataInner() *V1MonitoringCircuitsBandwidthPost200ResponseDataInner`

NewV1MonitoringCircuitsBandwidthPost200ResponseDataInner instantiates a new V1MonitoringCircuitsBandwidthPost200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MonitoringCircuitsBandwidthPost200ResponseDataInnerWithDefaults

`func NewV1MonitoringCircuitsBandwidthPost200ResponseDataInnerWithDefaults() *V1MonitoringCircuitsBandwidthPost200ResponseDataInner`

NewV1MonitoringCircuitsBandwidthPost200ResponseDataInnerWithDefaults instantiates a new V1MonitoringCircuitsBandwidthPost200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrier

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDlBwKbpsSamples

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) GetDlBwKbpsSamples() []V2MonitoringBfdPost200ResponseDataInnerSamplesInner`

GetDlBwKbpsSamples returns the DlBwKbpsSamples field if non-nil, zero value otherwise.

### GetDlBwKbpsSamplesOk

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) GetDlBwKbpsSamplesOk() (*[]V2MonitoringBfdPost200ResponseDataInnerSamplesInner, bool)`

GetDlBwKbpsSamplesOk returns a tuple with the DlBwKbpsSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlBwKbpsSamples

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) SetDlBwKbpsSamples(v []V2MonitoringBfdPost200ResponseDataInnerSamplesInner)`

SetDlBwKbpsSamples sets DlBwKbpsSamples field to given value.

### HasDlBwKbpsSamples

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) HasDlBwKbpsSamples() bool`

HasDlBwKbpsSamples returns a boolean if a field has been set.

### GetSelector

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) GetSelector() V1MonitoringCircuitsBandwidthPostRequestSelectorsInner`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) GetSelectorOk() (*V1MonitoringCircuitsBandwidthPostRequestSelectorsInner, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) SetSelector(v V1MonitoringCircuitsBandwidthPostRequestSelectorsInner)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetUlBwKbpsSamples

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) GetUlBwKbpsSamples() []V2MonitoringBfdPost200ResponseDataInnerSamplesInner`

GetUlBwKbpsSamples returns the UlBwKbpsSamples field if non-nil, zero value otherwise.

### GetUlBwKbpsSamplesOk

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) GetUlBwKbpsSamplesOk() (*[]V2MonitoringBfdPost200ResponseDataInnerSamplesInner, bool)`

GetUlBwKbpsSamplesOk returns a tuple with the UlBwKbpsSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlBwKbpsSamples

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) SetUlBwKbpsSamples(v []V2MonitoringBfdPost200ResponseDataInnerSamplesInner)`

SetUlBwKbpsSamples sets UlBwKbpsSamples field to given value.

### HasUlBwKbpsSamples

`func (o *V1MonitoringCircuitsBandwidthPost200ResponseDataInner) HasUlBwKbpsSamples() bool`

HasUlBwKbpsSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


