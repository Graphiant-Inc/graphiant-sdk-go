# V1DevicesDeviceIdConfigPutRequestCoreCoreVrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAggregations** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpAggregationsValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpAggregationsValue.md) |  | [optional] 
**BgpNeighbors** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue.md) |  | [optional] 
**BgpRedistribution** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue.md) |  | [optional] 
**DhcpSubnets** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValue.md) |  | [optional] 
**EbgpMultipath** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfEbgpMultipath**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfEbgpMultipath.md) |  | [optional] 
**IpfixExporters** | Pointer to [**map[string]V1GlobalConfigPatchRequestIpfixExportersValue**](V1GlobalConfigPatchRequestIpfixExportersValue.md) |  | [optional] 
**NatRuleset** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfNatRuleset**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfNatRuleset.md) |  | [optional] 
**Networks** | Pointer to **[]string** |  | [optional] 
**Ospfv2** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2.md) |  | [optional] 
**Ospfv3** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2.md) |  | [optional] 
**OverlayFilters** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOverlayFilters**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOverlayFilters.md) |  | [optional] 
**StaticRoutes** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValue.md) |  | [optional] 
**SyslogTargets** | Pointer to [**map[string]V1GlobalConfigPatchRequestSyslogServersValue**](V1GlobalConfigPatchRequestSyslogServersValue.md) |  | [optional] 
**TrafficRuleset** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfNatRuleset**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfNatRuleset.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrf

`func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrf() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf`

NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrf instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf`

NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpAggregations

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetBgpAggregations() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpAggregationsValue`

GetBgpAggregations returns the BgpAggregations field if non-nil, zero value otherwise.

### GetBgpAggregationsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetBgpAggregationsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpAggregationsValue, bool)`

GetBgpAggregationsOk returns a tuple with the BgpAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAggregations

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetBgpAggregations(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpAggregationsValue)`

SetBgpAggregations sets BgpAggregations field to given value.

### HasBgpAggregations

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasBgpAggregations() bool`

HasBgpAggregations returns a boolean if a field has been set.

### GetBgpNeighbors

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetBgpNeighbors() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue`

GetBgpNeighbors returns the BgpNeighbors field if non-nil, zero value otherwise.

### GetBgpNeighborsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetBgpNeighborsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue, bool)`

GetBgpNeighborsOk returns a tuple with the BgpNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighbors

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetBgpNeighbors(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue)`

SetBgpNeighbors sets BgpNeighbors field to given value.

### HasBgpNeighbors

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasBgpNeighbors() bool`

HasBgpNeighbors returns a boolean if a field has been set.

### GetBgpRedistribution

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetBgpRedistribution() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue`

GetBgpRedistribution returns the BgpRedistribution field if non-nil, zero value otherwise.

### GetBgpRedistributionOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetBgpRedistributionOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue, bool)`

GetBgpRedistributionOk returns a tuple with the BgpRedistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpRedistribution

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetBgpRedistribution(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue)`

SetBgpRedistribution sets BgpRedistribution field to given value.

### HasBgpRedistribution

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasBgpRedistribution() bool`

HasBgpRedistribution returns a boolean if a field has been set.

### GetDhcpSubnets

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetDhcpSubnets() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValue`

GetDhcpSubnets returns the DhcpSubnets field if non-nil, zero value otherwise.

### GetDhcpSubnetsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetDhcpSubnetsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValue, bool)`

GetDhcpSubnetsOk returns a tuple with the DhcpSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSubnets

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetDhcpSubnets(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValue)`

SetDhcpSubnets sets DhcpSubnets field to given value.

### HasDhcpSubnets

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasDhcpSubnets() bool`

HasDhcpSubnets returns a boolean if a field has been set.

### GetEbgpMultipath

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetEbgpMultipath() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfEbgpMultipath`

GetEbgpMultipath returns the EbgpMultipath field if non-nil, zero value otherwise.

### GetEbgpMultipathOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetEbgpMultipathOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfEbgpMultipath, bool)`

GetEbgpMultipathOk returns a tuple with the EbgpMultipath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpMultipath

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetEbgpMultipath(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfEbgpMultipath)`

SetEbgpMultipath sets EbgpMultipath field to given value.

### HasEbgpMultipath

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasEbgpMultipath() bool`

HasEbgpMultipath returns a boolean if a field has been set.

### GetIpfixExporters

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetIpfixExporters() map[string]V1GlobalConfigPatchRequestIpfixExportersValue`

GetIpfixExporters returns the IpfixExporters field if non-nil, zero value otherwise.

### GetIpfixExportersOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetIpfixExportersOk() (*map[string]V1GlobalConfigPatchRequestIpfixExportersValue, bool)`

GetIpfixExportersOk returns a tuple with the IpfixExporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporters

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetIpfixExporters(v map[string]V1GlobalConfigPatchRequestIpfixExportersValue)`

SetIpfixExporters sets IpfixExporters field to given value.

### HasIpfixExporters

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasIpfixExporters() bool`

HasIpfixExporters returns a boolean if a field has been set.

### GetNatRuleset

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetNatRuleset() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfNatRuleset`

GetNatRuleset returns the NatRuleset field if non-nil, zero value otherwise.

### GetNatRulesetOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetNatRulesetOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfNatRuleset, bool)`

GetNatRulesetOk returns a tuple with the NatRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatRuleset

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetNatRuleset(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfNatRuleset)`

SetNatRuleset sets NatRuleset field to given value.

### HasNatRuleset

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasNatRuleset() bool`

HasNatRuleset returns a boolean if a field has been set.

### GetNetworks

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOspfv2

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetOspfv2() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2`

GetOspfv2 returns the Ospfv2 field if non-nil, zero value otherwise.

### GetOspfv2Ok

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetOspfv2Ok() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2, bool)`

GetOspfv2Ok returns a tuple with the Ospfv2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfv2

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetOspfv2(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2)`

SetOspfv2 sets Ospfv2 field to given value.

### HasOspfv2

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasOspfv2() bool`

HasOspfv2 returns a boolean if a field has been set.

### GetOspfv3

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetOspfv3() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2`

GetOspfv3 returns the Ospfv3 field if non-nil, zero value otherwise.

### GetOspfv3Ok

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetOspfv3Ok() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2, bool)`

GetOspfv3Ok returns a tuple with the Ospfv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfv3

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetOspfv3(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2)`

SetOspfv3 sets Ospfv3 field to given value.

### HasOspfv3

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasOspfv3() bool`

HasOspfv3 returns a boolean if a field has been set.

### GetOverlayFilters

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetOverlayFilters() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOverlayFilters`

GetOverlayFilters returns the OverlayFilters field if non-nil, zero value otherwise.

### GetOverlayFiltersOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetOverlayFiltersOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOverlayFilters, bool)`

GetOverlayFiltersOk returns a tuple with the OverlayFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayFilters

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetOverlayFilters(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOverlayFilters)`

SetOverlayFilters sets OverlayFilters field to given value.

### HasOverlayFilters

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasOverlayFilters() bool`

HasOverlayFilters returns a boolean if a field has been set.

### GetStaticRoutes

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetStaticRoutes() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValue`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetStaticRoutesOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValue, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetStaticRoutes(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValue)`

SetStaticRoutes sets StaticRoutes field to given value.

### HasStaticRoutes

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasStaticRoutes() bool`

HasStaticRoutes returns a boolean if a field has been set.

### GetSyslogTargets

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetSyslogTargets() map[string]V1GlobalConfigPatchRequestSyslogServersValue`

GetSyslogTargets returns the SyslogTargets field if non-nil, zero value otherwise.

### GetSyslogTargetsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetSyslogTargetsOk() (*map[string]V1GlobalConfigPatchRequestSyslogServersValue, bool)`

GetSyslogTargetsOk returns a tuple with the SyslogTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogTargets

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetSyslogTargets(v map[string]V1GlobalConfigPatchRequestSyslogServersValue)`

SetSyslogTargets sets SyslogTargets field to given value.

### HasSyslogTargets

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasSyslogTargets() bool`

HasSyslogTargets returns a boolean if a field has been set.

### GetTrafficRuleset

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetTrafficRuleset() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfNatRuleset`

GetTrafficRuleset returns the TrafficRuleset field if non-nil, zero value otherwise.

### GetTrafficRulesetOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) GetTrafficRulesetOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfNatRuleset, bool)`

GetTrafficRulesetOk returns a tuple with the TrafficRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRuleset

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) SetTrafficRuleset(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfNatRuleset)`

SetTrafficRuleset sets TrafficRuleset field to given value.

### HasTrafficRuleset

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrf) HasTrafficRuleset() bool`

HasTrafficRuleset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


