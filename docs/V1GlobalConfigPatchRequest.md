# V1GlobalConfigPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalPrefixSets** | Pointer to [**map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue**](V1GlobalConfigPatchRequestGlobalPrefixSetsValue.md) |  | [optional] 
**IpfixExporters** | Pointer to [**map[string]V1GlobalConfigPatchRequestIpfixExportersValue**](V1GlobalConfigPatchRequestIpfixExportersValue.md) |  | [optional] 
**PrefixSets** | Pointer to [**map[string]V1GlobalConfigPatchRequestPrefixSetsValue**](V1GlobalConfigPatchRequestPrefixSetsValue.md) |  | [optional] 
**RoutingPolicies** | Pointer to [**map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue**](V1GlobalConfigPatchRequestRoutingPoliciesValue.md) |  | [optional] 
**Snmps** | Pointer to [**map[string]V1GlobalConfigPatchRequestSnmpsValue**](V1GlobalConfigPatchRequestSnmpsValue.md) |  | [optional] 
**SyslogServers** | Pointer to [**map[string]V1GlobalConfigPatchRequestSyslogServersValue**](V1GlobalConfigPatchRequestSyslogServersValue.md) |  | [optional] 
**TrafficPolicies** | Pointer to [**V1GlobalConfigPatchRequestTrafficPolicies**](V1GlobalConfigPatchRequestTrafficPolicies.md) |  | [optional] 
**VpnProfiles** | Pointer to [**map[string]V1GlobalConfigPatchRequestVpnProfilesValue**](V1GlobalConfigPatchRequestVpnProfilesValue.md) |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequest

`func NewV1GlobalConfigPatchRequest() *V1GlobalConfigPatchRequest`

NewV1GlobalConfigPatchRequest instantiates a new V1GlobalConfigPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestWithDefaults

`func NewV1GlobalConfigPatchRequestWithDefaults() *V1GlobalConfigPatchRequest`

NewV1GlobalConfigPatchRequestWithDefaults instantiates a new V1GlobalConfigPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalPrefixSets

`func (o *V1GlobalConfigPatchRequest) GetGlobalPrefixSets() map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue`

GetGlobalPrefixSets returns the GlobalPrefixSets field if non-nil, zero value otherwise.

### GetGlobalPrefixSetsOk

`func (o *V1GlobalConfigPatchRequest) GetGlobalPrefixSetsOk() (*map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue, bool)`

GetGlobalPrefixSetsOk returns a tuple with the GlobalPrefixSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPrefixSets

`func (o *V1GlobalConfigPatchRequest) SetGlobalPrefixSets(v map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue)`

SetGlobalPrefixSets sets GlobalPrefixSets field to given value.

### HasGlobalPrefixSets

`func (o *V1GlobalConfigPatchRequest) HasGlobalPrefixSets() bool`

HasGlobalPrefixSets returns a boolean if a field has been set.

### GetIpfixExporters

`func (o *V1GlobalConfigPatchRequest) GetIpfixExporters() map[string]V1GlobalConfigPatchRequestIpfixExportersValue`

GetIpfixExporters returns the IpfixExporters field if non-nil, zero value otherwise.

### GetIpfixExportersOk

`func (o *V1GlobalConfigPatchRequest) GetIpfixExportersOk() (*map[string]V1GlobalConfigPatchRequestIpfixExportersValue, bool)`

GetIpfixExportersOk returns a tuple with the IpfixExporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporters

`func (o *V1GlobalConfigPatchRequest) SetIpfixExporters(v map[string]V1GlobalConfigPatchRequestIpfixExportersValue)`

SetIpfixExporters sets IpfixExporters field to given value.

### HasIpfixExporters

`func (o *V1GlobalConfigPatchRequest) HasIpfixExporters() bool`

HasIpfixExporters returns a boolean if a field has been set.

### GetPrefixSets

`func (o *V1GlobalConfigPatchRequest) GetPrefixSets() map[string]V1GlobalConfigPatchRequestPrefixSetsValue`

GetPrefixSets returns the PrefixSets field if non-nil, zero value otherwise.

### GetPrefixSetsOk

`func (o *V1GlobalConfigPatchRequest) GetPrefixSetsOk() (*map[string]V1GlobalConfigPatchRequestPrefixSetsValue, bool)`

GetPrefixSetsOk returns a tuple with the PrefixSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSets

`func (o *V1GlobalConfigPatchRequest) SetPrefixSets(v map[string]V1GlobalConfigPatchRequestPrefixSetsValue)`

SetPrefixSets sets PrefixSets field to given value.

### HasPrefixSets

`func (o *V1GlobalConfigPatchRequest) HasPrefixSets() bool`

HasPrefixSets returns a boolean if a field has been set.

### GetRoutingPolicies

`func (o *V1GlobalConfigPatchRequest) GetRoutingPolicies() map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue`

GetRoutingPolicies returns the RoutingPolicies field if non-nil, zero value otherwise.

### GetRoutingPoliciesOk

`func (o *V1GlobalConfigPatchRequest) GetRoutingPoliciesOk() (*map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue, bool)`

GetRoutingPoliciesOk returns a tuple with the RoutingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicies

`func (o *V1GlobalConfigPatchRequest) SetRoutingPolicies(v map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue)`

SetRoutingPolicies sets RoutingPolicies field to given value.

### HasRoutingPolicies

`func (o *V1GlobalConfigPatchRequest) HasRoutingPolicies() bool`

HasRoutingPolicies returns a boolean if a field has been set.

### GetSnmps

`func (o *V1GlobalConfigPatchRequest) GetSnmps() map[string]V1GlobalConfigPatchRequestSnmpsValue`

GetSnmps returns the Snmps field if non-nil, zero value otherwise.

### GetSnmpsOk

`func (o *V1GlobalConfigPatchRequest) GetSnmpsOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValue, bool)`

GetSnmpsOk returns a tuple with the Snmps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmps

`func (o *V1GlobalConfigPatchRequest) SetSnmps(v map[string]V1GlobalConfigPatchRequestSnmpsValue)`

SetSnmps sets Snmps field to given value.

### HasSnmps

`func (o *V1GlobalConfigPatchRequest) HasSnmps() bool`

HasSnmps returns a boolean if a field has been set.

### GetSyslogServers

`func (o *V1GlobalConfigPatchRequest) GetSyslogServers() map[string]V1GlobalConfigPatchRequestSyslogServersValue`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *V1GlobalConfigPatchRequest) GetSyslogServersOk() (*map[string]V1GlobalConfigPatchRequestSyslogServersValue, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *V1GlobalConfigPatchRequest) SetSyslogServers(v map[string]V1GlobalConfigPatchRequestSyslogServersValue)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *V1GlobalConfigPatchRequest) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.

### GetTrafficPolicies

`func (o *V1GlobalConfigPatchRequest) GetTrafficPolicies() V1GlobalConfigPatchRequestTrafficPolicies`

GetTrafficPolicies returns the TrafficPolicies field if non-nil, zero value otherwise.

### GetTrafficPoliciesOk

`func (o *V1GlobalConfigPatchRequest) GetTrafficPoliciesOk() (*V1GlobalConfigPatchRequestTrafficPolicies, bool)`

GetTrafficPoliciesOk returns a tuple with the TrafficPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPolicies

`func (o *V1GlobalConfigPatchRequest) SetTrafficPolicies(v V1GlobalConfigPatchRequestTrafficPolicies)`

SetTrafficPolicies sets TrafficPolicies field to given value.

### HasTrafficPolicies

`func (o *V1GlobalConfigPatchRequest) HasTrafficPolicies() bool`

HasTrafficPolicies returns a boolean if a field has been set.

### GetVpnProfiles

`func (o *V1GlobalConfigPatchRequest) GetVpnProfiles() map[string]V1GlobalConfigPatchRequestVpnProfilesValue`

GetVpnProfiles returns the VpnProfiles field if non-nil, zero value otherwise.

### GetVpnProfilesOk

`func (o *V1GlobalConfigPatchRequest) GetVpnProfilesOk() (*map[string]V1GlobalConfigPatchRequestVpnProfilesValue, bool)`

GetVpnProfilesOk returns a tuple with the VpnProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfiles

`func (o *V1GlobalConfigPatchRequest) SetVpnProfiles(v map[string]V1GlobalConfigPatchRequestVpnProfilesValue)`

SetVpnProfiles sets VpnProfiles field to given value.

### HasVpnProfiles

`func (o *V1GlobalConfigPatchRequest) HasVpnProfiles() bool`

HasVpnProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


