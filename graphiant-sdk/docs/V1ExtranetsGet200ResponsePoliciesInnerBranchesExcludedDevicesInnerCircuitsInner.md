# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAggregations** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpAggregationsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpAggregationsInner.md) |  | [optional] 
**BgpMultipath** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpMultipath**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpMultipath.md) |  | [optional] 
**BgpNeighbors** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner.md) |  | [optional] 
**BgpRedistributions** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpRedistributions**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpRedistributions.md) |  | [optional] 
**Carrier** | Pointer to **string** |  | [optional] 
**CircuitType** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**CoreLogicalInterfacesV2** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerCoreLogicalInterfacesV2Inner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerCoreLogicalInterfacesV2Inner.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DiaEnabled** | Pointer to **bool** |  | [optional] 
**DiaSnmpIndex** | Pointer to **int64** |  | [optional] 
**DiscoveredPublicIp** | Pointer to **string** |  | [optional] 
**DropMechanism** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LastResort** | Pointer to **bool** |  | [optional] 
**LinkDownSpeedMbps** | Pointer to **int32** |  | [optional] 
**LinkUpSpeedMbps** | Pointer to **int32** |  | [optional] 
**Loopback** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PatAddresses** | Pointer to **[]string** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerProfile**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerProfile.md) |  | [optional] 
**QosProfile** | Pointer to **string** |  | [optional] 
**QosProfileType** | Pointer to **string** |  | [optional] 
**SnmpIndex** | Pointer to **int64** |  | [optional] 
**StaticRoutes** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerStaticRoutesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerStaticRoutesInner.md) |  | [optional] 
**WanInterfaceV2** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerCoreLogicalInterfacesV2Inner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerCoreLogicalInterfacesV2Inner.md) |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpAggregations

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetBgpAggregations() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpAggregationsInner`

GetBgpAggregations returns the BgpAggregations field if non-nil, zero value otherwise.

### GetBgpAggregationsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetBgpAggregationsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpAggregationsInner, bool)`

GetBgpAggregationsOk returns a tuple with the BgpAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAggregations

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetBgpAggregations(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpAggregationsInner)`

SetBgpAggregations sets BgpAggregations field to given value.

### HasBgpAggregations

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasBgpAggregations() bool`

HasBgpAggregations returns a boolean if a field has been set.

### GetBgpMultipath

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetBgpMultipath() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpMultipath`

GetBgpMultipath returns the BgpMultipath field if non-nil, zero value otherwise.

### GetBgpMultipathOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetBgpMultipathOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpMultipath, bool)`

GetBgpMultipathOk returns a tuple with the BgpMultipath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpMultipath

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetBgpMultipath(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpMultipath)`

SetBgpMultipath sets BgpMultipath field to given value.

### HasBgpMultipath

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasBgpMultipath() bool`

HasBgpMultipath returns a boolean if a field has been set.

### GetBgpNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetBgpNeighbors() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner`

GetBgpNeighbors returns the BgpNeighbors field if non-nil, zero value otherwise.

### GetBgpNeighborsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetBgpNeighborsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner, bool)`

GetBgpNeighborsOk returns a tuple with the BgpNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetBgpNeighbors(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner)`

SetBgpNeighbors sets BgpNeighbors field to given value.

### HasBgpNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasBgpNeighbors() bool`

HasBgpNeighbors returns a boolean if a field has been set.

### GetBgpRedistributions

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetBgpRedistributions() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpRedistributions`

GetBgpRedistributions returns the BgpRedistributions field if non-nil, zero value otherwise.

### GetBgpRedistributionsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetBgpRedistributionsOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpRedistributions, bool)`

GetBgpRedistributionsOk returns a tuple with the BgpRedistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpRedistributions

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetBgpRedistributions(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpRedistributions)`

SetBgpRedistributions sets BgpRedistributions field to given value.

### HasBgpRedistributions

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasBgpRedistributions() bool`

HasBgpRedistributions returns a boolean if a field has been set.

### GetCarrier

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCircuitType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetCircuitType() string`

GetCircuitType returns the CircuitType field if non-nil, zero value otherwise.

### GetCircuitTypeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetCircuitTypeOk() (*string, bool)`

GetCircuitTypeOk returns a tuple with the CircuitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetCircuitType(v string)`

SetCircuitType sets CircuitType field to given value.

### HasCircuitType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasCircuitType() bool`

HasCircuitType returns a boolean if a field has been set.

### GetConnectionType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetCoreLogicalInterfacesV2

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetCoreLogicalInterfacesV2() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerCoreLogicalInterfacesV2Inner`

GetCoreLogicalInterfacesV2 returns the CoreLogicalInterfacesV2 field if non-nil, zero value otherwise.

### GetCoreLogicalInterfacesV2Ok

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetCoreLogicalInterfacesV2Ok() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerCoreLogicalInterfacesV2Inner, bool)`

GetCoreLogicalInterfacesV2Ok returns a tuple with the CoreLogicalInterfacesV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreLogicalInterfacesV2

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetCoreLogicalInterfacesV2(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerCoreLogicalInterfacesV2Inner)`

SetCoreLogicalInterfacesV2 sets CoreLogicalInterfacesV2 field to given value.

### HasCoreLogicalInterfacesV2

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasCoreLogicalInterfacesV2() bool`

HasCoreLogicalInterfacesV2 returns a boolean if a field has been set.

### GetDescription

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiaEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetDiaEnabled() bool`

GetDiaEnabled returns the DiaEnabled field if non-nil, zero value otherwise.

### GetDiaEnabledOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetDiaEnabledOk() (*bool, bool)`

GetDiaEnabledOk returns a tuple with the DiaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiaEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetDiaEnabled(v bool)`

SetDiaEnabled sets DiaEnabled field to given value.

### HasDiaEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasDiaEnabled() bool`

HasDiaEnabled returns a boolean if a field has been set.

### GetDiaSnmpIndex

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetDiaSnmpIndex() int64`

GetDiaSnmpIndex returns the DiaSnmpIndex field if non-nil, zero value otherwise.

### GetDiaSnmpIndexOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetDiaSnmpIndexOk() (*int64, bool)`

GetDiaSnmpIndexOk returns a tuple with the DiaSnmpIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiaSnmpIndex

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetDiaSnmpIndex(v int64)`

SetDiaSnmpIndex sets DiaSnmpIndex field to given value.

### HasDiaSnmpIndex

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasDiaSnmpIndex() bool`

HasDiaSnmpIndex returns a boolean if a field has been set.

### GetDiscoveredPublicIp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetDiscoveredPublicIp() string`

GetDiscoveredPublicIp returns the DiscoveredPublicIp field if non-nil, zero value otherwise.

### GetDiscoveredPublicIpOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetDiscoveredPublicIpOk() (*string, bool)`

GetDiscoveredPublicIpOk returns a tuple with the DiscoveredPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredPublicIp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetDiscoveredPublicIp(v string)`

SetDiscoveredPublicIp sets DiscoveredPublicIp field to given value.

### HasDiscoveredPublicIp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasDiscoveredPublicIp() bool`

HasDiscoveredPublicIp returns a boolean if a field has been set.

### GetDropMechanism

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetDropMechanism() string`

GetDropMechanism returns the DropMechanism field if non-nil, zero value otherwise.

### GetDropMechanismOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetDropMechanismOk() (*string, bool)`

GetDropMechanismOk returns a tuple with the DropMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropMechanism

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetDropMechanism(v string)`

SetDropMechanism sets DropMechanism field to given value.

### HasDropMechanism

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasDropMechanism() bool`

HasDropMechanism returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetLabel

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastResort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetLastResort() bool`

GetLastResort returns the LastResort field if non-nil, zero value otherwise.

### GetLastResortOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetLastResortOk() (*bool, bool)`

GetLastResortOk returns a tuple with the LastResort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetLastResort(v bool)`

SetLastResort sets LastResort field to given value.

### HasLastResort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasLastResort() bool`

HasLastResort returns a boolean if a field has been set.

### GetLinkDownSpeedMbps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetLinkDownSpeedMbps() int32`

GetLinkDownSpeedMbps returns the LinkDownSpeedMbps field if non-nil, zero value otherwise.

### GetLinkDownSpeedMbpsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetLinkDownSpeedMbpsOk() (*int32, bool)`

GetLinkDownSpeedMbpsOk returns a tuple with the LinkDownSpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDownSpeedMbps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetLinkDownSpeedMbps(v int32)`

SetLinkDownSpeedMbps sets LinkDownSpeedMbps field to given value.

### HasLinkDownSpeedMbps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasLinkDownSpeedMbps() bool`

HasLinkDownSpeedMbps returns a boolean if a field has been set.

### GetLinkUpSpeedMbps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetLinkUpSpeedMbps() int32`

GetLinkUpSpeedMbps returns the LinkUpSpeedMbps field if non-nil, zero value otherwise.

### GetLinkUpSpeedMbpsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetLinkUpSpeedMbpsOk() (*int32, bool)`

GetLinkUpSpeedMbpsOk returns a tuple with the LinkUpSpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUpSpeedMbps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetLinkUpSpeedMbps(v int32)`

SetLinkUpSpeedMbps sets LinkUpSpeedMbps field to given value.

### HasLinkUpSpeedMbps

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasLinkUpSpeedMbps() bool`

HasLinkUpSpeedMbps returns a boolean if a field has been set.

### GetLoopback

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetLoopback() bool`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetLoopbackOk() (*bool, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetLoopback(v bool)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPatAddresses

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetPatAddresses() []string`

GetPatAddresses returns the PatAddresses field if non-nil, zero value otherwise.

### GetPatAddressesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetPatAddressesOk() (*[]string, bool)`

GetPatAddressesOk returns a tuple with the PatAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatAddresses

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetPatAddresses(v []string)`

SetPatAddresses sets PatAddresses field to given value.

### HasPatAddresses

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasPatAddresses() bool`

HasPatAddresses returns a boolean if a field has been set.

### GetPrivateIp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProfile

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetProfile() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetProfileOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetProfile(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQosProfile

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetQosProfile() string`

GetQosProfile returns the QosProfile field if non-nil, zero value otherwise.

### GetQosProfileOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetQosProfileOk() (*string, bool)`

GetQosProfileOk returns a tuple with the QosProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosProfile

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetQosProfile(v string)`

SetQosProfile sets QosProfile field to given value.

### HasQosProfile

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasQosProfile() bool`

HasQosProfile returns a boolean if a field has been set.

### GetQosProfileType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetQosProfileType() string`

GetQosProfileType returns the QosProfileType field if non-nil, zero value otherwise.

### GetQosProfileTypeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetQosProfileTypeOk() (*string, bool)`

GetQosProfileTypeOk returns a tuple with the QosProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosProfileType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetQosProfileType(v string)`

SetQosProfileType sets QosProfileType field to given value.

### HasQosProfileType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasQosProfileType() bool`

HasQosProfileType returns a boolean if a field has been set.

### GetSnmpIndex

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetSnmpIndex() int64`

GetSnmpIndex returns the SnmpIndex field if non-nil, zero value otherwise.

### GetSnmpIndexOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetSnmpIndexOk() (*int64, bool)`

GetSnmpIndexOk returns a tuple with the SnmpIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpIndex

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetSnmpIndex(v int64)`

SetSnmpIndex sets SnmpIndex field to given value.

### HasSnmpIndex

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasSnmpIndex() bool`

HasSnmpIndex returns a boolean if a field has been set.

### GetStaticRoutes

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetStaticRoutes() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerStaticRoutesInner`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetStaticRoutesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerStaticRoutesInner, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetStaticRoutes(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerStaticRoutesInner)`

SetStaticRoutes sets StaticRoutes field to given value.

### HasStaticRoutes

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasStaticRoutes() bool`

HasStaticRoutes returns a boolean if a field has been set.

### GetWanInterfaceV2

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetWanInterfaceV2() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerCoreLogicalInterfacesV2Inner`

GetWanInterfaceV2 returns the WanInterfaceV2 field if non-nil, zero value otherwise.

### GetWanInterfaceV2Ok

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) GetWanInterfaceV2Ok() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerCoreLogicalInterfacesV2Inner, bool)`

GetWanInterfaceV2Ok returns a tuple with the WanInterfaceV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanInterfaceV2

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) SetWanInterfaceV2(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerCoreLogicalInterfacesV2Inner)`

SetWanInterfaceV2 sets WanInterfaceV2 field to given value.

### HasWanInterfaceV2

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInner) HasWanInterfaceV2() bool`

HasWanInterfaceV2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


