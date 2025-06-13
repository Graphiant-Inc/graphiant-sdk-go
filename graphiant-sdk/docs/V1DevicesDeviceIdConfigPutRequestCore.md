# V1DevicesDeviceIdConfigPutRequestCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpInstance** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreBgpInstance**](V1DevicesDeviceIdConfigPutRequestCoreBgpInstance.md) |  | [optional] 
**CoreVrf** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrf**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrf.md) |  | [optional] 
**Interfaces** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreInterfacesValue**](V1DevicesDeviceIdConfigPutRequestCoreInterfacesValue.md) |  | [optional] 
**IpfixExporters** | Pointer to [**map[string]V1GlobalConfigPatchRequestIpfixExportersValue**](V1GlobalConfigPatchRequestIpfixExportersValue.md) |  | [optional] 
**IspVrfs** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrf.md) |  | [optional] 
**MaintenanceMode** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PrefixSets** | Pointer to [**map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue**](V1GlobalConfigPatchRequestGlobalPrefixSetsValue.md) |  | [optional] 
**Prometheus** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCorePrometheus**](V1DevicesDeviceIdConfigPutRequestCorePrometheus.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**RoutePolicies** | Pointer to [**map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue**](V1GlobalConfigPatchRequestRoutingPoliciesValue.md) |  | [optional] 
**Site** | Pointer to [**V1SitesPostRequestSite**](V1SitesPostRequestSite.md) |  | [optional] 
**TrafficPolicy** | Pointer to [**V1GlobalConfigPatchRequestTrafficPolicies**](V1GlobalConfigPatchRequestTrafficPolicies.md) |  | [optional] 
**Vrfs** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrf.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCore

`func NewV1DevicesDeviceIdConfigPutRequestCore() *V1DevicesDeviceIdConfigPutRequestCore`

NewV1DevicesDeviceIdConfigPutRequestCore instantiates a new V1DevicesDeviceIdConfigPutRequestCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreWithDefaults() *V1DevicesDeviceIdConfigPutRequestCore`

NewV1DevicesDeviceIdConfigPutRequestCoreWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpInstance

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetBgpInstance() V1DevicesDeviceIdConfigPutRequestCoreBgpInstance`

GetBgpInstance returns the BgpInstance field if non-nil, zero value otherwise.

### GetBgpInstanceOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetBgpInstanceOk() (*V1DevicesDeviceIdConfigPutRequestCoreBgpInstance, bool)`

GetBgpInstanceOk returns a tuple with the BgpInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpInstance

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetBgpInstance(v V1DevicesDeviceIdConfigPutRequestCoreBgpInstance)`

SetBgpInstance sets BgpInstance field to given value.

### HasBgpInstance

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasBgpInstance() bool`

HasBgpInstance returns a boolean if a field has been set.

### GetCoreVrf

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetCoreVrf() V1DevicesDeviceIdConfigPutRequestCoreCoreVrf`

GetCoreVrf returns the CoreVrf field if non-nil, zero value otherwise.

### GetCoreVrfOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetCoreVrfOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrf, bool)`

GetCoreVrfOk returns a tuple with the CoreVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreVrf

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetCoreVrf(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrf)`

SetCoreVrf sets CoreVrf field to given value.

### HasCoreVrf

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasCoreVrf() bool`

HasCoreVrf returns a boolean if a field has been set.

### GetInterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetInterfaces() map[string]V1DevicesDeviceIdConfigPutRequestCoreInterfacesValue`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetInterfacesOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreInterfacesValue, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetInterfaces(v map[string]V1DevicesDeviceIdConfigPutRequestCoreInterfacesValue)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetIpfixExporters

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetIpfixExporters() map[string]V1GlobalConfigPatchRequestIpfixExportersValue`

GetIpfixExporters returns the IpfixExporters field if non-nil, zero value otherwise.

### GetIpfixExportersOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetIpfixExportersOk() (*map[string]V1GlobalConfigPatchRequestIpfixExportersValue, bool)`

GetIpfixExportersOk returns a tuple with the IpfixExporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporters

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetIpfixExporters(v map[string]V1GlobalConfigPatchRequestIpfixExportersValue)`

SetIpfixExporters sets IpfixExporters field to given value.

### HasIpfixExporters

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasIpfixExporters() bool`

HasIpfixExporters returns a boolean if a field has been set.

### GetIspVrfs

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetIspVrfs() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf`

GetIspVrfs returns the IspVrfs field if non-nil, zero value otherwise.

### GetIspVrfsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetIspVrfsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf, bool)`

GetIspVrfsOk returns a tuple with the IspVrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspVrfs

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetIspVrfs(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf)`

SetIspVrfs sets IspVrfs field to given value.

### HasIspVrfs

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasIspVrfs() bool`

HasIspVrfs returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetName

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefixSets

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetPrefixSets() map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue`

GetPrefixSets returns the PrefixSets field if non-nil, zero value otherwise.

### GetPrefixSetsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetPrefixSetsOk() (*map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue, bool)`

GetPrefixSetsOk returns a tuple with the PrefixSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSets

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetPrefixSets(v map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValue)`

SetPrefixSets sets PrefixSets field to given value.

### HasPrefixSets

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasPrefixSets() bool`

HasPrefixSets returns a boolean if a field has been set.

### GetPrometheus

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetPrometheus() V1DevicesDeviceIdConfigPutRequestCorePrometheus`

GetPrometheus returns the Prometheus field if non-nil, zero value otherwise.

### GetPrometheusOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetPrometheusOk() (*V1DevicesDeviceIdConfigPutRequestCorePrometheus, bool)`

GetPrometheusOk returns a tuple with the Prometheus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheus

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetPrometheus(v V1DevicesDeviceIdConfigPutRequestCorePrometheus)`

SetPrometheus sets Prometheus field to given value.

### HasPrometheus

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasPrometheus() bool`

HasPrometheus returns a boolean if a field has been set.

### GetRegion

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionName

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetRoutePolicies

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetRoutePolicies() map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue`

GetRoutePolicies returns the RoutePolicies field if non-nil, zero value otherwise.

### GetRoutePoliciesOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetRoutePoliciesOk() (*map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue, bool)`

GetRoutePoliciesOk returns a tuple with the RoutePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePolicies

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetRoutePolicies(v map[string]V1GlobalConfigPatchRequestRoutingPoliciesValue)`

SetRoutePolicies sets RoutePolicies field to given value.

### HasRoutePolicies

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasRoutePolicies() bool`

HasRoutePolicies returns a boolean if a field has been set.

### GetSite

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetSite() V1SitesPostRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetSiteOk() (*V1SitesPostRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetSite(v V1SitesPostRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetTrafficPolicy

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetTrafficPolicy() V1GlobalConfigPatchRequestTrafficPolicies`

GetTrafficPolicy returns the TrafficPolicy field if non-nil, zero value otherwise.

### GetTrafficPolicyOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetTrafficPolicyOk() (*V1GlobalConfigPatchRequestTrafficPolicies, bool)`

GetTrafficPolicyOk returns a tuple with the TrafficPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPolicy

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetTrafficPolicy(v V1GlobalConfigPatchRequestTrafficPolicies)`

SetTrafficPolicy sets TrafficPolicy field to given value.

### HasTrafficPolicy

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasTrafficPolicy() bool`

HasTrafficPolicy returns a boolean if a field has been set.

### GetVrfs

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetVrfs() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf`

GetVrfs returns the Vrfs field if non-nil, zero value otherwise.

### GetVrfsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCore) GetVrfsOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf, bool)`

GetVrfsOk returns a tuple with the Vrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfs

`func (o *V1DevicesDeviceIdConfigPutRequestCore) SetVrfs(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrf)`

SetVrfs sets Vrfs field to given value.

### HasVrfs

`func (o *V1DevicesDeviceIdConfigPutRequestCore) HasVrfs() bool`

HasVrfs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


