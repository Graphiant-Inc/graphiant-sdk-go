# V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAggregations** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpAggregationsValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpAggregationsValue.md) |  | [optional] 
**BgpMultipath** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfEbgpMultipath**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfEbgpMultipath.md) |  | [optional] 
**BgpNeighbors** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue.md) |  | [optional] 
**BgpRedistribution** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue.md) |  | [optional] 
**Carrier** | Pointer to **string** |  | [optional] 
**CircuitType** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DiaEnabled** | Pointer to **bool** |  | [optional] 
**DropMechanism** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LastResort** | Pointer to **bool** |  | [optional] 
**LinkDownSpeedMbps** | Pointer to **int32** |  | [optional] 
**LinkUpSpeedMbps** | Pointer to **int32** |  | [optional] 
**Loopback** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PatAddresses** | Pointer to [**V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValuePatAddresses**](V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValuePatAddresses.md) |  | [optional] 
**QosProfile** | Pointer to **string** |  | [optional] 
**QosProfileType** | Pointer to **string** |  | [optional] 
**StaticRoutes** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValue.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue

`func NewV1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue() *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue`

NewV1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue instantiates a new V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestEdgeCircuitsValueWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestEdgeCircuitsValueWithDefaults() *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue`

NewV1DevicesDeviceIdConfigPutRequestEdgeCircuitsValueWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpAggregations

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetBgpAggregations() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpAggregationsValue`

GetBgpAggregations returns the BgpAggregations field if non-nil, zero value otherwise.

### GetBgpAggregationsOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetBgpAggregationsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpAggregationsValue, bool)`

GetBgpAggregationsOk returns a tuple with the BgpAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAggregations

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetBgpAggregations(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpAggregationsValue)`

SetBgpAggregations sets BgpAggregations field to given value.

### HasBgpAggregations

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasBgpAggregations() bool`

HasBgpAggregations returns a boolean if a field has been set.

### GetBgpMultipath

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetBgpMultipath() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfEbgpMultipath`

GetBgpMultipath returns the BgpMultipath field if non-nil, zero value otherwise.

### GetBgpMultipathOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetBgpMultipathOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfEbgpMultipath, bool)`

GetBgpMultipathOk returns a tuple with the BgpMultipath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpMultipath

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetBgpMultipath(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfEbgpMultipath)`

SetBgpMultipath sets BgpMultipath field to given value.

### HasBgpMultipath

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasBgpMultipath() bool`

HasBgpMultipath returns a boolean if a field has been set.

### GetBgpNeighbors

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetBgpNeighbors() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue`

GetBgpNeighbors returns the BgpNeighbors field if non-nil, zero value otherwise.

### GetBgpNeighborsOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetBgpNeighborsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue, bool)`

GetBgpNeighborsOk returns a tuple with the BgpNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighbors

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetBgpNeighbors(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValue)`

SetBgpNeighbors sets BgpNeighbors field to given value.

### HasBgpNeighbors

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasBgpNeighbors() bool`

HasBgpNeighbors returns a boolean if a field has been set.

### GetBgpRedistribution

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetBgpRedistribution() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue`

GetBgpRedistribution returns the BgpRedistribution field if non-nil, zero value otherwise.

### GetBgpRedistributionOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetBgpRedistributionOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue, bool)`

GetBgpRedistributionOk returns a tuple with the BgpRedistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpRedistribution

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetBgpRedistribution(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpRedistributionValue)`

SetBgpRedistribution sets BgpRedistribution field to given value.

### HasBgpRedistribution

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasBgpRedistribution() bool`

HasBgpRedistribution returns a boolean if a field has been set.

### GetCarrier

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCircuitType

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetCircuitType() string`

GetCircuitType returns the CircuitType field if non-nil, zero value otherwise.

### GetCircuitTypeOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetCircuitTypeOk() (*string, bool)`

GetCircuitTypeOk returns a tuple with the CircuitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitType

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetCircuitType(v string)`

SetCircuitType sets CircuitType field to given value.

### HasCircuitType

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasCircuitType() bool`

HasCircuitType returns a boolean if a field has been set.

### GetConnectionType

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiaEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetDiaEnabled() bool`

GetDiaEnabled returns the DiaEnabled field if non-nil, zero value otherwise.

### GetDiaEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetDiaEnabledOk() (*bool, bool)`

GetDiaEnabledOk returns a tuple with the DiaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiaEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetDiaEnabled(v bool)`

SetDiaEnabled sets DiaEnabled field to given value.

### HasDiaEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasDiaEnabled() bool`

HasDiaEnabled returns a boolean if a field has been set.

### GetDropMechanism

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetDropMechanism() string`

GetDropMechanism returns the DropMechanism field if non-nil, zero value otherwise.

### GetDropMechanismOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetDropMechanismOk() (*string, bool)`

GetDropMechanismOk returns a tuple with the DropMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropMechanism

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetDropMechanism(v string)`

SetDropMechanism sets DropMechanism field to given value.

### HasDropMechanism

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasDropMechanism() bool`

HasDropMechanism returns a boolean if a field has been set.

### GetLabel

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastResort

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetLastResort() bool`

GetLastResort returns the LastResort field if non-nil, zero value otherwise.

### GetLastResortOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetLastResortOk() (*bool, bool)`

GetLastResortOk returns a tuple with the LastResort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResort

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetLastResort(v bool)`

SetLastResort sets LastResort field to given value.

### HasLastResort

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasLastResort() bool`

HasLastResort returns a boolean if a field has been set.

### GetLinkDownSpeedMbps

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetLinkDownSpeedMbps() int32`

GetLinkDownSpeedMbps returns the LinkDownSpeedMbps field if non-nil, zero value otherwise.

### GetLinkDownSpeedMbpsOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetLinkDownSpeedMbpsOk() (*int32, bool)`

GetLinkDownSpeedMbpsOk returns a tuple with the LinkDownSpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDownSpeedMbps

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetLinkDownSpeedMbps(v int32)`

SetLinkDownSpeedMbps sets LinkDownSpeedMbps field to given value.

### HasLinkDownSpeedMbps

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasLinkDownSpeedMbps() bool`

HasLinkDownSpeedMbps returns a boolean if a field has been set.

### GetLinkUpSpeedMbps

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetLinkUpSpeedMbps() int32`

GetLinkUpSpeedMbps returns the LinkUpSpeedMbps field if non-nil, zero value otherwise.

### GetLinkUpSpeedMbpsOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetLinkUpSpeedMbpsOk() (*int32, bool)`

GetLinkUpSpeedMbpsOk returns a tuple with the LinkUpSpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUpSpeedMbps

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetLinkUpSpeedMbps(v int32)`

SetLinkUpSpeedMbps sets LinkUpSpeedMbps field to given value.

### HasLinkUpSpeedMbps

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasLinkUpSpeedMbps() bool`

HasLinkUpSpeedMbps returns a boolean if a field has been set.

### GetLoopback

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetLoopback() bool`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetLoopbackOk() (*bool, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetLoopback(v bool)`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetName

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPatAddresses

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetPatAddresses() V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValuePatAddresses`

GetPatAddresses returns the PatAddresses field if non-nil, zero value otherwise.

### GetPatAddressesOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetPatAddressesOk() (*V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValuePatAddresses, bool)`

GetPatAddressesOk returns a tuple with the PatAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatAddresses

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetPatAddresses(v V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValuePatAddresses)`

SetPatAddresses sets PatAddresses field to given value.

### HasPatAddresses

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasPatAddresses() bool`

HasPatAddresses returns a boolean if a field has been set.

### GetQosProfile

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetQosProfile() string`

GetQosProfile returns the QosProfile field if non-nil, zero value otherwise.

### GetQosProfileOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetQosProfileOk() (*string, bool)`

GetQosProfileOk returns a tuple with the QosProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosProfile

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetQosProfile(v string)`

SetQosProfile sets QosProfile field to given value.

### HasQosProfile

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasQosProfile() bool`

HasQosProfile returns a boolean if a field has been set.

### GetQosProfileType

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetQosProfileType() string`

GetQosProfileType returns the QosProfileType field if non-nil, zero value otherwise.

### GetQosProfileTypeOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetQosProfileTypeOk() (*string, bool)`

GetQosProfileTypeOk returns a tuple with the QosProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosProfileType

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetQosProfileType(v string)`

SetQosProfileType sets QosProfileType field to given value.

### HasQosProfileType

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasQosProfileType() bool`

HasQosProfileType returns a boolean if a field has been set.

### GetStaticRoutes

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetStaticRoutes() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValue`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) GetStaticRoutesOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValue, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) SetStaticRoutes(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValue)`

SetStaticRoutes sets StaticRoutes field to given value.

### HasStaticRoutes

`func (o *V1DevicesDeviceIdConfigPutRequestEdgeCircuitsValue) HasStaticRoutes() bool`

HasStaticRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


