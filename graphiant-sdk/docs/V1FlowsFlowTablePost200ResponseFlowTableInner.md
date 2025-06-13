# V1FlowsFlowTablePost200ResponseFlowTableInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestIp** | Pointer to **string** |  | [optional] 
**DestPort** | Pointer to **int32** |  | [optional] 
**DlCircuitName** | Pointer to **string** |  | [optional] 
**DlUsage** | Pointer to **float64** |  | [optional] 
**EgressLocalCoreRegion** | Pointer to **string** |  | [optional] 
**IngressLocalCoreRegion** | Pointer to **string** |  | [optional] 
**LanSegment** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**RemoteCoreRegion** | Pointer to **string** |  | [optional] 
**SlaClass** | Pointer to **string** |  | [optional] 
**SrcIp** | Pointer to **string** |  | [optional] 
**SrcPort** | Pointer to **int32** |  | [optional] 
**Ts** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**UlCircuitName** | Pointer to **string** |  | [optional] 
**UlUsage** | Pointer to **float64** |  | [optional] 

## Methods

### NewV1FlowsFlowTablePost200ResponseFlowTableInner

`func NewV1FlowsFlowTablePost200ResponseFlowTableInner() *V1FlowsFlowTablePost200ResponseFlowTableInner`

NewV1FlowsFlowTablePost200ResponseFlowTableInner instantiates a new V1FlowsFlowTablePost200ResponseFlowTableInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FlowsFlowTablePost200ResponseFlowTableInnerWithDefaults

`func NewV1FlowsFlowTablePost200ResponseFlowTableInnerWithDefaults() *V1FlowsFlowTablePost200ResponseFlowTableInner`

NewV1FlowsFlowTablePost200ResponseFlowTableInnerWithDefaults instantiates a new V1FlowsFlowTablePost200ResponseFlowTableInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestIp

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetDestIp() string`

GetDestIp returns the DestIp field if non-nil, zero value otherwise.

### GetDestIpOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetDestIpOk() (*string, bool)`

GetDestIpOk returns a tuple with the DestIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIp

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetDestIp(v string)`

SetDestIp sets DestIp field to given value.

### HasDestIp

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasDestIp() bool`

HasDestIp returns a boolean if a field has been set.

### GetDestPort

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetDestPort() int32`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetDestPortOk() (*int32, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetDestPort(v int32)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDlCircuitName

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetDlCircuitName() string`

GetDlCircuitName returns the DlCircuitName field if non-nil, zero value otherwise.

### GetDlCircuitNameOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetDlCircuitNameOk() (*string, bool)`

GetDlCircuitNameOk returns a tuple with the DlCircuitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlCircuitName

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetDlCircuitName(v string)`

SetDlCircuitName sets DlCircuitName field to given value.

### HasDlCircuitName

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasDlCircuitName() bool`

HasDlCircuitName returns a boolean if a field has been set.

### GetDlUsage

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetDlUsage() float64`

GetDlUsage returns the DlUsage field if non-nil, zero value otherwise.

### GetDlUsageOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetDlUsageOk() (*float64, bool)`

GetDlUsageOk returns a tuple with the DlUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlUsage

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetDlUsage(v float64)`

SetDlUsage sets DlUsage field to given value.

### HasDlUsage

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasDlUsage() bool`

HasDlUsage returns a boolean if a field has been set.

### GetEgressLocalCoreRegion

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetEgressLocalCoreRegion() string`

GetEgressLocalCoreRegion returns the EgressLocalCoreRegion field if non-nil, zero value otherwise.

### GetEgressLocalCoreRegionOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetEgressLocalCoreRegionOk() (*string, bool)`

GetEgressLocalCoreRegionOk returns a tuple with the EgressLocalCoreRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressLocalCoreRegion

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetEgressLocalCoreRegion(v string)`

SetEgressLocalCoreRegion sets EgressLocalCoreRegion field to given value.

### HasEgressLocalCoreRegion

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasEgressLocalCoreRegion() bool`

HasEgressLocalCoreRegion returns a boolean if a field has been set.

### GetIngressLocalCoreRegion

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetIngressLocalCoreRegion() string`

GetIngressLocalCoreRegion returns the IngressLocalCoreRegion field if non-nil, zero value otherwise.

### GetIngressLocalCoreRegionOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetIngressLocalCoreRegionOk() (*string, bool)`

GetIngressLocalCoreRegionOk returns a tuple with the IngressLocalCoreRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressLocalCoreRegion

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetIngressLocalCoreRegion(v string)`

SetIngressLocalCoreRegion sets IngressLocalCoreRegion field to given value.

### HasIngressLocalCoreRegion

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasIngressLocalCoreRegion() bool`

HasIngressLocalCoreRegion returns a boolean if a field has been set.

### GetLanSegment

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetLanSegment() string`

GetLanSegment returns the LanSegment field if non-nil, zero value otherwise.

### GetLanSegmentOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetLanSegmentOk() (*string, bool)`

GetLanSegmentOk returns a tuple with the LanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegment

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetLanSegment(v string)`

SetLanSegment sets LanSegment field to given value.

### HasLanSegment

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasLanSegment() bool`

HasLanSegment returns a boolean if a field has been set.

### GetProtocol

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemoteCoreRegion

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetRemoteCoreRegion() string`

GetRemoteCoreRegion returns the RemoteCoreRegion field if non-nil, zero value otherwise.

### GetRemoteCoreRegionOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetRemoteCoreRegionOk() (*string, bool)`

GetRemoteCoreRegionOk returns a tuple with the RemoteCoreRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCoreRegion

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetRemoteCoreRegion(v string)`

SetRemoteCoreRegion sets RemoteCoreRegion field to given value.

### HasRemoteCoreRegion

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasRemoteCoreRegion() bool`

HasRemoteCoreRegion returns a boolean if a field has been set.

### GetSlaClass

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetSlaClass() string`

GetSlaClass returns the SlaClass field if non-nil, zero value otherwise.

### GetSlaClassOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetSlaClassOk() (*string, bool)`

GetSlaClassOk returns a tuple with the SlaClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaClass

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetSlaClass(v string)`

SetSlaClass sets SlaClass field to given value.

### HasSlaClass

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasSlaClass() bool`

HasSlaClass returns a boolean if a field has been set.

### GetSrcIp

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetSrcIp() string`

GetSrcIp returns the SrcIp field if non-nil, zero value otherwise.

### GetSrcIpOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetSrcIpOk() (*string, bool)`

GetSrcIpOk returns a tuple with the SrcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIp

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetSrcIp(v string)`

SetSrcIp sets SrcIp field to given value.

### HasSrcIp

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasSrcIp() bool`

HasSrcIp returns a boolean if a field has been set.

### GetSrcPort

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetTs

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetTs sets Ts field to given value.

### HasTs

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetUlCircuitName

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetUlCircuitName() string`

GetUlCircuitName returns the UlCircuitName field if non-nil, zero value otherwise.

### GetUlCircuitNameOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetUlCircuitNameOk() (*string, bool)`

GetUlCircuitNameOk returns a tuple with the UlCircuitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlCircuitName

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetUlCircuitName(v string)`

SetUlCircuitName sets UlCircuitName field to given value.

### HasUlCircuitName

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasUlCircuitName() bool`

HasUlCircuitName returns a boolean if a field has been set.

### GetUlUsage

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetUlUsage() float64`

GetUlUsage returns the UlUsage field if non-nil, zero value otherwise.

### GetUlUsageOk

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) GetUlUsageOk() (*float64, bool)`

GetUlUsageOk returns a tuple with the UlUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlUsage

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) SetUlUsage(v float64)`

SetUlUsage sets UlUsage field to given value.

### HasUlUsage

`func (o *V1FlowsFlowTablePost200ResponseFlowTableInner) HasUlUsage() bool`

HasUlUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


