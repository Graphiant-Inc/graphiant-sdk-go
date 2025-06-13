# V1GlobalConfigSitePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalPrefixSetOps** | Pointer to **map[string]string** |  | [optional] 
**IpfixExporterOps** | Pointer to **map[string]string** |  | [optional] 
**IpfixExporterOpsV2** | Pointer to [**map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value**](V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value.md) |  | [optional] 
**PrefixSetOps** | Pointer to **map[string]string** |  | [optional] 
**RoutingPolicyOps** | Pointer to **map[string]string** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**SnmpOps** | Pointer to **map[string]string** |  | [optional] 
**SyslogServerOps** | Pointer to **map[string]string** |  | [optional] 
**SyslogServerOpsV2** | Pointer to [**map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value**](V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value.md) |  | [optional] 
**TrafficPolicyOps** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewV1GlobalConfigSitePostRequest

`func NewV1GlobalConfigSitePostRequest() *V1GlobalConfigSitePostRequest`

NewV1GlobalConfigSitePostRequest instantiates a new V1GlobalConfigSitePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigSitePostRequestWithDefaults

`func NewV1GlobalConfigSitePostRequestWithDefaults() *V1GlobalConfigSitePostRequest`

NewV1GlobalConfigSitePostRequestWithDefaults instantiates a new V1GlobalConfigSitePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalPrefixSetOps

`func (o *V1GlobalConfigSitePostRequest) GetGlobalPrefixSetOps() map[string]string`

GetGlobalPrefixSetOps returns the GlobalPrefixSetOps field if non-nil, zero value otherwise.

### GetGlobalPrefixSetOpsOk

`func (o *V1GlobalConfigSitePostRequest) GetGlobalPrefixSetOpsOk() (*map[string]string, bool)`

GetGlobalPrefixSetOpsOk returns a tuple with the GlobalPrefixSetOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPrefixSetOps

`func (o *V1GlobalConfigSitePostRequest) SetGlobalPrefixSetOps(v map[string]string)`

SetGlobalPrefixSetOps sets GlobalPrefixSetOps field to given value.

### HasGlobalPrefixSetOps

`func (o *V1GlobalConfigSitePostRequest) HasGlobalPrefixSetOps() bool`

HasGlobalPrefixSetOps returns a boolean if a field has been set.

### GetIpfixExporterOps

`func (o *V1GlobalConfigSitePostRequest) GetIpfixExporterOps() map[string]string`

GetIpfixExporterOps returns the IpfixExporterOps field if non-nil, zero value otherwise.

### GetIpfixExporterOpsOk

`func (o *V1GlobalConfigSitePostRequest) GetIpfixExporterOpsOk() (*map[string]string, bool)`

GetIpfixExporterOpsOk returns a tuple with the IpfixExporterOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporterOps

`func (o *V1GlobalConfigSitePostRequest) SetIpfixExporterOps(v map[string]string)`

SetIpfixExporterOps sets IpfixExporterOps field to given value.

### HasIpfixExporterOps

`func (o *V1GlobalConfigSitePostRequest) HasIpfixExporterOps() bool`

HasIpfixExporterOps returns a boolean if a field has been set.

### GetIpfixExporterOpsV2

`func (o *V1GlobalConfigSitePostRequest) GetIpfixExporterOpsV2() map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value`

GetIpfixExporterOpsV2 returns the IpfixExporterOpsV2 field if non-nil, zero value otherwise.

### GetIpfixExporterOpsV2Ok

`func (o *V1GlobalConfigSitePostRequest) GetIpfixExporterOpsV2Ok() (*map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value, bool)`

GetIpfixExporterOpsV2Ok returns a tuple with the IpfixExporterOpsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporterOpsV2

`func (o *V1GlobalConfigSitePostRequest) SetIpfixExporterOpsV2(v map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value)`

SetIpfixExporterOpsV2 sets IpfixExporterOpsV2 field to given value.

### HasIpfixExporterOpsV2

`func (o *V1GlobalConfigSitePostRequest) HasIpfixExporterOpsV2() bool`

HasIpfixExporterOpsV2 returns a boolean if a field has been set.

### GetPrefixSetOps

`func (o *V1GlobalConfigSitePostRequest) GetPrefixSetOps() map[string]string`

GetPrefixSetOps returns the PrefixSetOps field if non-nil, zero value otherwise.

### GetPrefixSetOpsOk

`func (o *V1GlobalConfigSitePostRequest) GetPrefixSetOpsOk() (*map[string]string, bool)`

GetPrefixSetOpsOk returns a tuple with the PrefixSetOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSetOps

`func (o *V1GlobalConfigSitePostRequest) SetPrefixSetOps(v map[string]string)`

SetPrefixSetOps sets PrefixSetOps field to given value.

### HasPrefixSetOps

`func (o *V1GlobalConfigSitePostRequest) HasPrefixSetOps() bool`

HasPrefixSetOps returns a boolean if a field has been set.

### GetRoutingPolicyOps

`func (o *V1GlobalConfigSitePostRequest) GetRoutingPolicyOps() map[string]string`

GetRoutingPolicyOps returns the RoutingPolicyOps field if non-nil, zero value otherwise.

### GetRoutingPolicyOpsOk

`func (o *V1GlobalConfigSitePostRequest) GetRoutingPolicyOpsOk() (*map[string]string, bool)`

GetRoutingPolicyOpsOk returns a tuple with the RoutingPolicyOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicyOps

`func (o *V1GlobalConfigSitePostRequest) SetRoutingPolicyOps(v map[string]string)`

SetRoutingPolicyOps sets RoutingPolicyOps field to given value.

### HasRoutingPolicyOps

`func (o *V1GlobalConfigSitePostRequest) HasRoutingPolicyOps() bool`

HasRoutingPolicyOps returns a boolean if a field has been set.

### GetSiteId

`func (o *V1GlobalConfigSitePostRequest) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *V1GlobalConfigSitePostRequest) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *V1GlobalConfigSitePostRequest) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *V1GlobalConfigSitePostRequest) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSnmpOps

`func (o *V1GlobalConfigSitePostRequest) GetSnmpOps() map[string]string`

GetSnmpOps returns the SnmpOps field if non-nil, zero value otherwise.

### GetSnmpOpsOk

`func (o *V1GlobalConfigSitePostRequest) GetSnmpOpsOk() (*map[string]string, bool)`

GetSnmpOpsOk returns a tuple with the SnmpOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpOps

`func (o *V1GlobalConfigSitePostRequest) SetSnmpOps(v map[string]string)`

SetSnmpOps sets SnmpOps field to given value.

### HasSnmpOps

`func (o *V1GlobalConfigSitePostRequest) HasSnmpOps() bool`

HasSnmpOps returns a boolean if a field has been set.

### GetSyslogServerOps

`func (o *V1GlobalConfigSitePostRequest) GetSyslogServerOps() map[string]string`

GetSyslogServerOps returns the SyslogServerOps field if non-nil, zero value otherwise.

### GetSyslogServerOpsOk

`func (o *V1GlobalConfigSitePostRequest) GetSyslogServerOpsOk() (*map[string]string, bool)`

GetSyslogServerOpsOk returns a tuple with the SyslogServerOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServerOps

`func (o *V1GlobalConfigSitePostRequest) SetSyslogServerOps(v map[string]string)`

SetSyslogServerOps sets SyslogServerOps field to given value.

### HasSyslogServerOps

`func (o *V1GlobalConfigSitePostRequest) HasSyslogServerOps() bool`

HasSyslogServerOps returns a boolean if a field has been set.

### GetSyslogServerOpsV2

`func (o *V1GlobalConfigSitePostRequest) GetSyslogServerOpsV2() map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value`

GetSyslogServerOpsV2 returns the SyslogServerOpsV2 field if non-nil, zero value otherwise.

### GetSyslogServerOpsV2Ok

`func (o *V1GlobalConfigSitePostRequest) GetSyslogServerOpsV2Ok() (*map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value, bool)`

GetSyslogServerOpsV2Ok returns a tuple with the SyslogServerOpsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServerOpsV2

`func (o *V1GlobalConfigSitePostRequest) SetSyslogServerOpsV2(v map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value)`

SetSyslogServerOpsV2 sets SyslogServerOpsV2 field to given value.

### HasSyslogServerOpsV2

`func (o *V1GlobalConfigSitePostRequest) HasSyslogServerOpsV2() bool`

HasSyslogServerOpsV2 returns a boolean if a field has been set.

### GetTrafficPolicyOps

`func (o *V1GlobalConfigSitePostRequest) GetTrafficPolicyOps() map[string]string`

GetTrafficPolicyOps returns the TrafficPolicyOps field if non-nil, zero value otherwise.

### GetTrafficPolicyOpsOk

`func (o *V1GlobalConfigSitePostRequest) GetTrafficPolicyOpsOk() (*map[string]string, bool)`

GetTrafficPolicyOpsOk returns a tuple with the TrafficPolicyOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPolicyOps

`func (o *V1GlobalConfigSitePostRequest) SetTrafficPolicyOps(v map[string]string)`

SetTrafficPolicyOps sets TrafficPolicyOps field to given value.

### HasTrafficPolicyOps

`func (o *V1GlobalConfigSitePostRequest) HasTrafficPolicyOps() bool`

HasTrafficPolicyOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


