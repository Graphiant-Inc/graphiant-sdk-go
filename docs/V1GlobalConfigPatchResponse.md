# V1GlobalConfigPatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalPrefixSets** | Pointer to **map[string]int64** |  | [optional] 
**IpfixExporters** | Pointer to **map[string]int64** |  | [optional] 
**Ntps** | Pointer to **map[string]int64** |  | [optional] 
**PrefixSets** | Pointer to **map[string]int64** |  | [optional] 
**RoutingPolicies** | Pointer to **map[string]int64** |  | [optional] 
**Snmps** | Pointer to **map[string]int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SyslogServers** | Pointer to **map[string]int64** |  | [optional] 
**TrafficPolicies** | Pointer to **map[string]int64** |  | [optional] 
**VpnProfiles** | Pointer to **map[string]int64** |  | [optional] 

## Methods

### NewV1GlobalConfigPatchResponse

`func NewV1GlobalConfigPatchResponse() *V1GlobalConfigPatchResponse`

NewV1GlobalConfigPatchResponse instantiates a new V1GlobalConfigPatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchResponseWithDefaults

`func NewV1GlobalConfigPatchResponseWithDefaults() *V1GlobalConfigPatchResponse`

NewV1GlobalConfigPatchResponseWithDefaults instantiates a new V1GlobalConfigPatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalPrefixSets

`func (o *V1GlobalConfigPatchResponse) GetGlobalPrefixSets() map[string]int64`

GetGlobalPrefixSets returns the GlobalPrefixSets field if non-nil, zero value otherwise.

### GetGlobalPrefixSetsOk

`func (o *V1GlobalConfigPatchResponse) GetGlobalPrefixSetsOk() (*map[string]int64, bool)`

GetGlobalPrefixSetsOk returns a tuple with the GlobalPrefixSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPrefixSets

`func (o *V1GlobalConfigPatchResponse) SetGlobalPrefixSets(v map[string]int64)`

SetGlobalPrefixSets sets GlobalPrefixSets field to given value.

### HasGlobalPrefixSets

`func (o *V1GlobalConfigPatchResponse) HasGlobalPrefixSets() bool`

HasGlobalPrefixSets returns a boolean if a field has been set.

### GetIpfixExporters

`func (o *V1GlobalConfigPatchResponse) GetIpfixExporters() map[string]int64`

GetIpfixExporters returns the IpfixExporters field if non-nil, zero value otherwise.

### GetIpfixExportersOk

`func (o *V1GlobalConfigPatchResponse) GetIpfixExportersOk() (*map[string]int64, bool)`

GetIpfixExportersOk returns a tuple with the IpfixExporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporters

`func (o *V1GlobalConfigPatchResponse) SetIpfixExporters(v map[string]int64)`

SetIpfixExporters sets IpfixExporters field to given value.

### HasIpfixExporters

`func (o *V1GlobalConfigPatchResponse) HasIpfixExporters() bool`

HasIpfixExporters returns a boolean if a field has been set.

### GetNtps

`func (o *V1GlobalConfigPatchResponse) GetNtps() map[string]int64`

GetNtps returns the Ntps field if non-nil, zero value otherwise.

### GetNtpsOk

`func (o *V1GlobalConfigPatchResponse) GetNtpsOk() (*map[string]int64, bool)`

GetNtpsOk returns a tuple with the Ntps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtps

`func (o *V1GlobalConfigPatchResponse) SetNtps(v map[string]int64)`

SetNtps sets Ntps field to given value.

### HasNtps

`func (o *V1GlobalConfigPatchResponse) HasNtps() bool`

HasNtps returns a boolean if a field has been set.

### GetPrefixSets

`func (o *V1GlobalConfigPatchResponse) GetPrefixSets() map[string]int64`

GetPrefixSets returns the PrefixSets field if non-nil, zero value otherwise.

### GetPrefixSetsOk

`func (o *V1GlobalConfigPatchResponse) GetPrefixSetsOk() (*map[string]int64, bool)`

GetPrefixSetsOk returns a tuple with the PrefixSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSets

`func (o *V1GlobalConfigPatchResponse) SetPrefixSets(v map[string]int64)`

SetPrefixSets sets PrefixSets field to given value.

### HasPrefixSets

`func (o *V1GlobalConfigPatchResponse) HasPrefixSets() bool`

HasPrefixSets returns a boolean if a field has been set.

### GetRoutingPolicies

`func (o *V1GlobalConfigPatchResponse) GetRoutingPolicies() map[string]int64`

GetRoutingPolicies returns the RoutingPolicies field if non-nil, zero value otherwise.

### GetRoutingPoliciesOk

`func (o *V1GlobalConfigPatchResponse) GetRoutingPoliciesOk() (*map[string]int64, bool)`

GetRoutingPoliciesOk returns a tuple with the RoutingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicies

`func (o *V1GlobalConfigPatchResponse) SetRoutingPolicies(v map[string]int64)`

SetRoutingPolicies sets RoutingPolicies field to given value.

### HasRoutingPolicies

`func (o *V1GlobalConfigPatchResponse) HasRoutingPolicies() bool`

HasRoutingPolicies returns a boolean if a field has been set.

### GetSnmps

`func (o *V1GlobalConfigPatchResponse) GetSnmps() map[string]int64`

GetSnmps returns the Snmps field if non-nil, zero value otherwise.

### GetSnmpsOk

`func (o *V1GlobalConfigPatchResponse) GetSnmpsOk() (*map[string]int64, bool)`

GetSnmpsOk returns a tuple with the Snmps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmps

`func (o *V1GlobalConfigPatchResponse) SetSnmps(v map[string]int64)`

SetSnmps sets Snmps field to given value.

### HasSnmps

`func (o *V1GlobalConfigPatchResponse) HasSnmps() bool`

HasSnmps returns a boolean if a field has been set.

### GetStatus

`func (o *V1GlobalConfigPatchResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1GlobalConfigPatchResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1GlobalConfigPatchResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1GlobalConfigPatchResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyslogServers

`func (o *V1GlobalConfigPatchResponse) GetSyslogServers() map[string]int64`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *V1GlobalConfigPatchResponse) GetSyslogServersOk() (*map[string]int64, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *V1GlobalConfigPatchResponse) SetSyslogServers(v map[string]int64)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *V1GlobalConfigPatchResponse) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.

### GetTrafficPolicies

`func (o *V1GlobalConfigPatchResponse) GetTrafficPolicies() map[string]int64`

GetTrafficPolicies returns the TrafficPolicies field if non-nil, zero value otherwise.

### GetTrafficPoliciesOk

`func (o *V1GlobalConfigPatchResponse) GetTrafficPoliciesOk() (*map[string]int64, bool)`

GetTrafficPoliciesOk returns a tuple with the TrafficPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPolicies

`func (o *V1GlobalConfigPatchResponse) SetTrafficPolicies(v map[string]int64)`

SetTrafficPolicies sets TrafficPolicies field to given value.

### HasTrafficPolicies

`func (o *V1GlobalConfigPatchResponse) HasTrafficPolicies() bool`

HasTrafficPolicies returns a boolean if a field has been set.

### GetVpnProfiles

`func (o *V1GlobalConfigPatchResponse) GetVpnProfiles() map[string]int64`

GetVpnProfiles returns the VpnProfiles field if non-nil, zero value otherwise.

### GetVpnProfilesOk

`func (o *V1GlobalConfigPatchResponse) GetVpnProfilesOk() (*map[string]int64, bool)`

GetVpnProfilesOk returns a tuple with the VpnProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfiles

`func (o *V1GlobalConfigPatchResponse) SetVpnProfiles(v map[string]int64)`

SetVpnProfiles sets VpnProfiles field to given value.

### HasVpnProfiles

`func (o *V1GlobalConfigPatchResponse) HasVpnProfiles() bool`

HasVpnProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


