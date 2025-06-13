# V1GlobalConfigPatchRequestIpfixExportersValueExporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationAddress** | Pointer to **string** |  | [optional] 
**DestinationPort** | Pointer to **int32** |  | [optional] 
**GlobalId** | Pointer to **int64** |  | [optional] 
**IsGlobalSync** | Pointer to **bool** |  | [optional] 
**MonitoredSegments** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SampleMode** | Pointer to **string** |  | [optional] 
**SampleRate** | Pointer to **int32** |  | [optional] 
**SourceInterfaceName** | Pointer to **string** |  | [optional] 
**VrfId** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestIpfixExportersValueExporter

`func NewV1GlobalConfigPatchRequestIpfixExportersValueExporter() *V1GlobalConfigPatchRequestIpfixExportersValueExporter`

NewV1GlobalConfigPatchRequestIpfixExportersValueExporter instantiates a new V1GlobalConfigPatchRequestIpfixExportersValueExporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestIpfixExportersValueExporterWithDefaults

`func NewV1GlobalConfigPatchRequestIpfixExportersValueExporterWithDefaults() *V1GlobalConfigPatchRequestIpfixExportersValueExporter`

NewV1GlobalConfigPatchRequestIpfixExportersValueExporterWithDefaults instantiates a new V1GlobalConfigPatchRequestIpfixExportersValueExporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationAddress

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetDestinationPort

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetDestinationPort() int32`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetDestinationPortOk() (*int32, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetDestinationPort(v int32)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetGlobalId

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetIsGlobalSync

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetIsGlobalSync() bool`

GetIsGlobalSync returns the IsGlobalSync field if non-nil, zero value otherwise.

### GetIsGlobalSyncOk

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetIsGlobalSyncOk() (*bool, bool)`

GetIsGlobalSyncOk returns a tuple with the IsGlobalSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalSync

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetIsGlobalSync(v bool)`

SetIsGlobalSync sets IsGlobalSync field to given value.

### HasIsGlobalSync

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasIsGlobalSync() bool`

HasIsGlobalSync returns a boolean if a field has been set.

### GetMonitoredSegments

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetMonitoredSegments() []string`

GetMonitoredSegments returns the MonitoredSegments field if non-nil, zero value otherwise.

### GetMonitoredSegmentsOk

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetMonitoredSegmentsOk() (*[]string, bool)`

GetMonitoredSegmentsOk returns a tuple with the MonitoredSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredSegments

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetMonitoredSegments(v []string)`

SetMonitoredSegments sets MonitoredSegments field to given value.

### HasMonitoredSegments

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasMonitoredSegments() bool`

HasMonitoredSegments returns a boolean if a field has been set.

### GetName

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSampleMode

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSampleMode() string`

GetSampleMode returns the SampleMode field if non-nil, zero value otherwise.

### GetSampleModeOk

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSampleModeOk() (*string, bool)`

GetSampleModeOk returns a tuple with the SampleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleMode

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetSampleMode(v string)`

SetSampleMode sets SampleMode field to given value.

### HasSampleMode

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasSampleMode() bool`

HasSampleMode returns a boolean if a field has been set.

### GetSampleRate

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSampleRate() int32`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSampleRateOk() (*int32, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetSampleRate(v int32)`

SetSampleRate sets SampleRate field to given value.

### HasSampleRate

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasSampleRate() bool`

HasSampleRate returns a boolean if a field has been set.

### GetSourceInterfaceName

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSourceInterfaceName() string`

GetSourceInterfaceName returns the SourceInterfaceName field if non-nil, zero value otherwise.

### GetSourceInterfaceNameOk

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetSourceInterfaceNameOk() (*string, bool)`

GetSourceInterfaceNameOk returns a tuple with the SourceInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterfaceName

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetSourceInterfaceName(v string)`

SetSourceInterfaceName sets SourceInterfaceName field to given value.

### HasSourceInterfaceName

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasSourceInterfaceName() bool`

HasSourceInterfaceName returns a boolean if a field has been set.

### GetVrfId

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetVrfId() int64`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) GetVrfIdOk() (*int64, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) SetVrfId(v int64)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *V1GlobalConfigPatchRequestIpfixExportersValueExporter) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


