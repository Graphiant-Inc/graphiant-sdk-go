# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAggregations** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpAggregationsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpAggregationsInner.md) |  | [optional] 
**BgpMultipath** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpMultipath**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpMultipath.md) |  | [optional] 
**BgpNeighbors** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner.md) |  | [optional] 
**BgpRedistributions** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpRedistributions**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpRedistributions.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DhcpSubnets** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInner.md) |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**Function** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IpfixExporters** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpfixExportersInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpfixExportersInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NatRuleset** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to **[]string** |  | [optional] 
**Ospfv2Process** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2Process**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2Process.md) |  | [optional] 
**Ospfv3Process** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv3Process**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv3Process.md) |  | [optional] 
**OverlayFilters** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOverlayFilters**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOverlayFilters.md) |  | [optional] 
**Routable** | Pointer to **bool** |  | [optional] 
**RouteDistinguisher** | Pointer to **string** |  | [optional] 
**StaticRoutes** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerStaticRoutesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerStaticRoutesInner.md) |  | [optional] 
**SyslogTargets** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerSyslogTargetsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerSyslogTargetsInner.md) |  | [optional] 
**TrafficRuleset** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpAggregations

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetBgpAggregations() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpAggregationsInner`

GetBgpAggregations returns the BgpAggregations field if non-nil, zero value otherwise.

### GetBgpAggregationsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetBgpAggregationsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpAggregationsInner, bool)`

GetBgpAggregationsOk returns a tuple with the BgpAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAggregations

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetBgpAggregations(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpAggregationsInner)`

SetBgpAggregations sets BgpAggregations field to given value.

### HasBgpAggregations

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasBgpAggregations() bool`

HasBgpAggregations returns a boolean if a field has been set.

### GetBgpMultipath

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetBgpMultipath() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpMultipath`

GetBgpMultipath returns the BgpMultipath field if non-nil, zero value otherwise.

### GetBgpMultipathOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetBgpMultipathOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpMultipath, bool)`

GetBgpMultipathOk returns a tuple with the BgpMultipath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpMultipath

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetBgpMultipath(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpMultipath)`

SetBgpMultipath sets BgpMultipath field to given value.

### HasBgpMultipath

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasBgpMultipath() bool`

HasBgpMultipath returns a boolean if a field has been set.

### GetBgpNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetBgpNeighbors() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner`

GetBgpNeighbors returns the BgpNeighbors field if non-nil, zero value otherwise.

### GetBgpNeighborsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetBgpNeighborsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner, bool)`

GetBgpNeighborsOk returns a tuple with the BgpNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetBgpNeighbors(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner)`

SetBgpNeighbors sets BgpNeighbors field to given value.

### HasBgpNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasBgpNeighbors() bool`

HasBgpNeighbors returns a boolean if a field has been set.

### GetBgpRedistributions

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetBgpRedistributions() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpRedistributions`

GetBgpRedistributions returns the BgpRedistributions field if non-nil, zero value otherwise.

### GetBgpRedistributionsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetBgpRedistributionsOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpRedistributions, bool)`

GetBgpRedistributionsOk returns a tuple with the BgpRedistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpRedistributions

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetBgpRedistributions(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpRedistributions)`

SetBgpRedistributions sets BgpRedistributions field to given value.

### HasBgpRedistributions

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasBgpRedistributions() bool`

HasBgpRedistributions returns a boolean if a field has been set.

### GetDescription

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDhcpSubnets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetDhcpSubnets() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInner`

GetDhcpSubnets returns the DhcpSubnets field if non-nil, zero value otherwise.

### GetDhcpSubnetsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetDhcpSubnetsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInner, bool)`

GetDhcpSubnetsOk returns a tuple with the DhcpSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSubnets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetDhcpSubnets(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInner)`

SetDhcpSubnets sets DhcpSubnets field to given value.

### HasDhcpSubnets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasDhcpSubnets() bool`

HasDhcpSubnets returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetFunction

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpfixExporters

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetIpfixExporters() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpfixExportersInner`

GetIpfixExporters returns the IpfixExporters field if non-nil, zero value otherwise.

### GetIpfixExportersOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetIpfixExportersOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpfixExportersInner, bool)`

GetIpfixExportersOk returns a tuple with the IpfixExporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporters

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetIpfixExporters(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerIpfixExportersInner)`

SetIpfixExporters sets IpfixExporters field to given value.

### HasIpfixExporters

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasIpfixExporters() bool`

HasIpfixExporters returns a boolean if a field has been set.

### GetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNatRuleset

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetNatRuleset() string`

GetNatRuleset returns the NatRuleset field if non-nil, zero value otherwise.

### GetNatRulesetOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetNatRulesetOk() (*string, bool)`

GetNatRulesetOk returns a tuple with the NatRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatRuleset

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetNatRuleset(v string)`

SetNatRuleset sets NatRuleset field to given value.

### HasNatRuleset

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasNatRuleset() bool`

HasNatRuleset returns a boolean if a field has been set.

### GetNetworks

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOspfv2Process

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetOspfv2Process() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2Process`

GetOspfv2Process returns the Ospfv2Process field if non-nil, zero value otherwise.

### GetOspfv2ProcessOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetOspfv2ProcessOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2Process, bool)`

GetOspfv2ProcessOk returns a tuple with the Ospfv2Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfv2Process

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetOspfv2Process(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2Process)`

SetOspfv2Process sets Ospfv2Process field to given value.

### HasOspfv2Process

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasOspfv2Process() bool`

HasOspfv2Process returns a boolean if a field has been set.

### GetOspfv3Process

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetOspfv3Process() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv3Process`

GetOspfv3Process returns the Ospfv3Process field if non-nil, zero value otherwise.

### GetOspfv3ProcessOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetOspfv3ProcessOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv3Process, bool)`

GetOspfv3ProcessOk returns a tuple with the Ospfv3Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfv3Process

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetOspfv3Process(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv3Process)`

SetOspfv3Process sets Ospfv3Process field to given value.

### HasOspfv3Process

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasOspfv3Process() bool`

HasOspfv3Process returns a boolean if a field has been set.

### GetOverlayFilters

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetOverlayFilters() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOverlayFilters`

GetOverlayFilters returns the OverlayFilters field if non-nil, zero value otherwise.

### GetOverlayFiltersOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetOverlayFiltersOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOverlayFilters, bool)`

GetOverlayFiltersOk returns a tuple with the OverlayFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayFilters

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetOverlayFilters(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOverlayFilters)`

SetOverlayFilters sets OverlayFilters field to given value.

### HasOverlayFilters

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasOverlayFilters() bool`

HasOverlayFilters returns a boolean if a field has been set.

### GetRoutable

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetRoutable() bool`

GetRoutable returns the Routable field if non-nil, zero value otherwise.

### GetRoutableOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetRoutableOk() (*bool, bool)`

GetRoutableOk returns a tuple with the Routable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutable

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetRoutable(v bool)`

SetRoutable sets Routable field to given value.

### HasRoutable

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasRoutable() bool`

HasRoutable returns a boolean if a field has been set.

### GetRouteDistinguisher

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetRouteDistinguisher() string`

GetRouteDistinguisher returns the RouteDistinguisher field if non-nil, zero value otherwise.

### GetRouteDistinguisherOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetRouteDistinguisherOk() (*string, bool)`

GetRouteDistinguisherOk returns a tuple with the RouteDistinguisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDistinguisher

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetRouteDistinguisher(v string)`

SetRouteDistinguisher sets RouteDistinguisher field to given value.

### HasRouteDistinguisher

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasRouteDistinguisher() bool`

HasRouteDistinguisher returns a boolean if a field has been set.

### GetStaticRoutes

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetStaticRoutes() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerStaticRoutesInner`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetStaticRoutesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerStaticRoutesInner, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetStaticRoutes(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerStaticRoutesInner)`

SetStaticRoutes sets StaticRoutes field to given value.

### HasStaticRoutes

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasStaticRoutes() bool`

HasStaticRoutes returns a boolean if a field has been set.

### GetSyslogTargets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetSyslogTargets() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerSyslogTargetsInner`

GetSyslogTargets returns the SyslogTargets field if non-nil, zero value otherwise.

### GetSyslogTargetsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetSyslogTargetsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerSyslogTargetsInner, bool)`

GetSyslogTargetsOk returns a tuple with the SyslogTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogTargets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetSyslogTargets(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerSyslogTargetsInner)`

SetSyslogTargets sets SyslogTargets field to given value.

### HasSyslogTargets

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasSyslogTargets() bool`

HasSyslogTargets returns a boolean if a field has been set.

### GetTrafficRuleset

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetTrafficRuleset() string`

GetTrafficRuleset returns the TrafficRuleset field if non-nil, zero value otherwise.

### GetTrafficRulesetOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) GetTrafficRulesetOk() (*string, bool)`

GetTrafficRulesetOk returns a tuple with the TrafficRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRuleset

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) SetTrafficRuleset(v string)`

SetTrafficRuleset sets TrafficRuleset field to given value.

### HasTrafficRuleset

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner) HasTrafficRuleset() bool`

HasTrafficRuleset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


