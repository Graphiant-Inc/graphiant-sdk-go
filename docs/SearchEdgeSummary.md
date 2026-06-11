# SearchEdgeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedOn** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**CanaryMode** | Pointer to **string** |  | [optional] 
**ConnectedRegions** | Pointer to **[]string** |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**DiscoveredLocation** | Pointer to **string** |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**EnterpriseName** | Pointer to **string** |  | [optional] 
**FirstAppearedOn** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**IpDetected** | Pointer to **string** |  | [optional] 
**IsHardware** | Pointer to **bool** |  | [optional] 
**IsNew** | Pointer to **bool** |  | [optional] 
**IsRequested** | Pointer to **bool** |  | [optional] 
**IsRequestedStatus** | Pointer to **string** |  | [optional] 
**LastBootedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**Location** | Pointer to [**ManaV2Location**](ManaV2Location.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**OverrideRegion** | Pointer to **string** |  | [optional] 
**ParentEnterpriseName** | Pointer to **string** |  | [optional] 
**PortalStatus** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**SerialNum** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**Stale** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SwName** | Pointer to **string** |  | [optional] 
**SwVersion** | Pointer to **string** |  | [optional] 
**TtConnCount** | Pointer to **int32** |  | [optional] 
**UpgradeSummary** | Pointer to [**UpgradeUpgradeSummary**](UpgradeUpgradeSummary.md) |  | [optional] 

## Methods

### NewSearchEdgeSummary

`func NewSearchEdgeSummary() *SearchEdgeSummary`

NewSearchEdgeSummary instantiates a new SearchEdgeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEdgeSummaryWithDefaults

`func NewSearchEdgeSummaryWithDefaults() *SearchEdgeSummary`

NewSearchEdgeSummaryWithDefaults instantiates a new SearchEdgeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedOn

`func (o *SearchEdgeSummary) GetAssignedOn() GoogleProtobufTimestamp`

GetAssignedOn returns the AssignedOn field if non-nil, zero value otherwise.

### GetAssignedOnOk

`func (o *SearchEdgeSummary) GetAssignedOnOk() (*GoogleProtobufTimestamp, bool)`

GetAssignedOnOk returns a tuple with the AssignedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedOn

`func (o *SearchEdgeSummary) SetAssignedOn(v GoogleProtobufTimestamp)`

SetAssignedOn sets AssignedOn field to given value.

### HasAssignedOn

`func (o *SearchEdgeSummary) HasAssignedOn() bool`

HasAssignedOn returns a boolean if a field has been set.

### GetCanaryMode

`func (o *SearchEdgeSummary) GetCanaryMode() string`

GetCanaryMode returns the CanaryMode field if non-nil, zero value otherwise.

### GetCanaryModeOk

`func (o *SearchEdgeSummary) GetCanaryModeOk() (*string, bool)`

GetCanaryModeOk returns a tuple with the CanaryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanaryMode

`func (o *SearchEdgeSummary) SetCanaryMode(v string)`

SetCanaryMode sets CanaryMode field to given value.

### HasCanaryMode

`func (o *SearchEdgeSummary) HasCanaryMode() bool`

HasCanaryMode returns a boolean if a field has been set.

### GetConnectedRegions

`func (o *SearchEdgeSummary) GetConnectedRegions() []string`

GetConnectedRegions returns the ConnectedRegions field if non-nil, zero value otherwise.

### GetConnectedRegionsOk

`func (o *SearchEdgeSummary) GetConnectedRegionsOk() (*[]string, bool)`

GetConnectedRegionsOk returns a tuple with the ConnectedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedRegions

`func (o *SearchEdgeSummary) SetConnectedRegions(v []string)`

SetConnectedRegions sets ConnectedRegions field to given value.

### HasConnectedRegions

`func (o *SearchEdgeSummary) HasConnectedRegions() bool`

HasConnectedRegions returns a boolean if a field has been set.

### GetDeviceId

`func (o *SearchEdgeSummary) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SearchEdgeSummary) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SearchEdgeSummary) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SearchEdgeSummary) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDiscoveredLocation

`func (o *SearchEdgeSummary) GetDiscoveredLocation() string`

GetDiscoveredLocation returns the DiscoveredLocation field if non-nil, zero value otherwise.

### GetDiscoveredLocationOk

`func (o *SearchEdgeSummary) GetDiscoveredLocationOk() (*string, bool)`

GetDiscoveredLocationOk returns a tuple with the DiscoveredLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredLocation

`func (o *SearchEdgeSummary) SetDiscoveredLocation(v string)`

SetDiscoveredLocation sets DiscoveredLocation field to given value.

### HasDiscoveredLocation

`func (o *SearchEdgeSummary) HasDiscoveredLocation() bool`

HasDiscoveredLocation returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *SearchEdgeSummary) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *SearchEdgeSummary) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *SearchEdgeSummary) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *SearchEdgeSummary) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetEnterpriseName

`func (o *SearchEdgeSummary) GetEnterpriseName() string`

GetEnterpriseName returns the EnterpriseName field if non-nil, zero value otherwise.

### GetEnterpriseNameOk

`func (o *SearchEdgeSummary) GetEnterpriseNameOk() (*string, bool)`

GetEnterpriseNameOk returns a tuple with the EnterpriseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseName

`func (o *SearchEdgeSummary) SetEnterpriseName(v string)`

SetEnterpriseName sets EnterpriseName field to given value.

### HasEnterpriseName

`func (o *SearchEdgeSummary) HasEnterpriseName() bool`

HasEnterpriseName returns a boolean if a field has been set.

### GetFirstAppearedOn

`func (o *SearchEdgeSummary) GetFirstAppearedOn() GoogleProtobufTimestamp`

GetFirstAppearedOn returns the FirstAppearedOn field if non-nil, zero value otherwise.

### GetFirstAppearedOnOk

`func (o *SearchEdgeSummary) GetFirstAppearedOnOk() (*GoogleProtobufTimestamp, bool)`

GetFirstAppearedOnOk returns a tuple with the FirstAppearedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAppearedOn

`func (o *SearchEdgeSummary) SetFirstAppearedOn(v GoogleProtobufTimestamp)`

SetFirstAppearedOn sets FirstAppearedOn field to given value.

### HasFirstAppearedOn

`func (o *SearchEdgeSummary) HasFirstAppearedOn() bool`

HasFirstAppearedOn returns a boolean if a field has been set.

### GetHostname

`func (o *SearchEdgeSummary) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SearchEdgeSummary) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SearchEdgeSummary) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SearchEdgeSummary) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpDetected

`func (o *SearchEdgeSummary) GetIpDetected() string`

GetIpDetected returns the IpDetected field if non-nil, zero value otherwise.

### GetIpDetectedOk

`func (o *SearchEdgeSummary) GetIpDetectedOk() (*string, bool)`

GetIpDetectedOk returns a tuple with the IpDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDetected

`func (o *SearchEdgeSummary) SetIpDetected(v string)`

SetIpDetected sets IpDetected field to given value.

### HasIpDetected

`func (o *SearchEdgeSummary) HasIpDetected() bool`

HasIpDetected returns a boolean if a field has been set.

### GetIsHardware

`func (o *SearchEdgeSummary) GetIsHardware() bool`

GetIsHardware returns the IsHardware field if non-nil, zero value otherwise.

### GetIsHardwareOk

`func (o *SearchEdgeSummary) GetIsHardwareOk() (*bool, bool)`

GetIsHardwareOk returns a tuple with the IsHardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardware

`func (o *SearchEdgeSummary) SetIsHardware(v bool)`

SetIsHardware sets IsHardware field to given value.

### HasIsHardware

`func (o *SearchEdgeSummary) HasIsHardware() bool`

HasIsHardware returns a boolean if a field has been set.

### GetIsNew

`func (o *SearchEdgeSummary) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *SearchEdgeSummary) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *SearchEdgeSummary) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.

### HasIsNew

`func (o *SearchEdgeSummary) HasIsNew() bool`

HasIsNew returns a boolean if a field has been set.

### GetIsRequested

`func (o *SearchEdgeSummary) GetIsRequested() bool`

GetIsRequested returns the IsRequested field if non-nil, zero value otherwise.

### GetIsRequestedOk

`func (o *SearchEdgeSummary) GetIsRequestedOk() (*bool, bool)`

GetIsRequestedOk returns a tuple with the IsRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequested

`func (o *SearchEdgeSummary) SetIsRequested(v bool)`

SetIsRequested sets IsRequested field to given value.

### HasIsRequested

`func (o *SearchEdgeSummary) HasIsRequested() bool`

HasIsRequested returns a boolean if a field has been set.

### GetIsRequestedStatus

`func (o *SearchEdgeSummary) GetIsRequestedStatus() string`

GetIsRequestedStatus returns the IsRequestedStatus field if non-nil, zero value otherwise.

### GetIsRequestedStatusOk

`func (o *SearchEdgeSummary) GetIsRequestedStatusOk() (*string, bool)`

GetIsRequestedStatusOk returns a tuple with the IsRequestedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequestedStatus

`func (o *SearchEdgeSummary) SetIsRequestedStatus(v string)`

SetIsRequestedStatus sets IsRequestedStatus field to given value.

### HasIsRequestedStatus

`func (o *SearchEdgeSummary) HasIsRequestedStatus() bool`

HasIsRequestedStatus returns a boolean if a field has been set.

### GetLastBootedAt

`func (o *SearchEdgeSummary) GetLastBootedAt() GoogleProtobufTimestamp`

GetLastBootedAt returns the LastBootedAt field if non-nil, zero value otherwise.

### GetLastBootedAtOk

`func (o *SearchEdgeSummary) GetLastBootedAtOk() (*GoogleProtobufTimestamp, bool)`

GetLastBootedAtOk returns a tuple with the LastBootedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBootedAt

`func (o *SearchEdgeSummary) SetLastBootedAt(v GoogleProtobufTimestamp)`

SetLastBootedAt sets LastBootedAt field to given value.

### HasLastBootedAt

`func (o *SearchEdgeSummary) HasLastBootedAt() bool`

HasLastBootedAt returns a boolean if a field has been set.

### GetLocation

`func (o *SearchEdgeSummary) GetLocation() ManaV2Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SearchEdgeSummary) GetLocationOk() (*ManaV2Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SearchEdgeSummary) SetLocation(v ManaV2Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SearchEdgeSummary) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetModel

`func (o *SearchEdgeSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SearchEdgeSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SearchEdgeSummary) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SearchEdgeSummary) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOverrideRegion

`func (o *SearchEdgeSummary) GetOverrideRegion() string`

GetOverrideRegion returns the OverrideRegion field if non-nil, zero value otherwise.

### GetOverrideRegionOk

`func (o *SearchEdgeSummary) GetOverrideRegionOk() (*string, bool)`

GetOverrideRegionOk returns a tuple with the OverrideRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRegion

`func (o *SearchEdgeSummary) SetOverrideRegion(v string)`

SetOverrideRegion sets OverrideRegion field to given value.

### HasOverrideRegion

`func (o *SearchEdgeSummary) HasOverrideRegion() bool`

HasOverrideRegion returns a boolean if a field has been set.

### GetParentEnterpriseName

`func (o *SearchEdgeSummary) GetParentEnterpriseName() string`

GetParentEnterpriseName returns the ParentEnterpriseName field if non-nil, zero value otherwise.

### GetParentEnterpriseNameOk

`func (o *SearchEdgeSummary) GetParentEnterpriseNameOk() (*string, bool)`

GetParentEnterpriseNameOk returns a tuple with the ParentEnterpriseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEnterpriseName

`func (o *SearchEdgeSummary) SetParentEnterpriseName(v string)`

SetParentEnterpriseName sets ParentEnterpriseName field to given value.

### HasParentEnterpriseName

`func (o *SearchEdgeSummary) HasParentEnterpriseName() bool`

HasParentEnterpriseName returns a boolean if a field has been set.

### GetPortalStatus

`func (o *SearchEdgeSummary) GetPortalStatus() string`

GetPortalStatus returns the PortalStatus field if non-nil, zero value otherwise.

### GetPortalStatusOk

`func (o *SearchEdgeSummary) GetPortalStatusOk() (*string, bool)`

GetPortalStatusOk returns a tuple with the PortalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalStatus

`func (o *SearchEdgeSummary) SetPortalStatus(v string)`

SetPortalStatus sets PortalStatus field to given value.

### HasPortalStatus

`func (o *SearchEdgeSummary) HasPortalStatus() bool`

HasPortalStatus returns a boolean if a field has been set.

### GetRegion

`func (o *SearchEdgeSummary) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SearchEdgeSummary) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SearchEdgeSummary) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SearchEdgeSummary) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRole

`func (o *SearchEdgeSummary) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SearchEdgeSummary) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SearchEdgeSummary) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SearchEdgeSummary) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSerialNum

`func (o *SearchEdgeSummary) GetSerialNum() string`

GetSerialNum returns the SerialNum field if non-nil, zero value otherwise.

### GetSerialNumOk

`func (o *SearchEdgeSummary) GetSerialNumOk() (*string, bool)`

GetSerialNumOk returns a tuple with the SerialNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNum

`func (o *SearchEdgeSummary) SetSerialNum(v string)`

SetSerialNum sets SerialNum field to given value.

### HasSerialNum

`func (o *SearchEdgeSummary) HasSerialNum() bool`

HasSerialNum returns a boolean if a field has been set.

### GetSite

`func (o *SearchEdgeSummary) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *SearchEdgeSummary) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *SearchEdgeSummary) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *SearchEdgeSummary) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSiteId

`func (o *SearchEdgeSummary) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SearchEdgeSummary) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SearchEdgeSummary) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SearchEdgeSummary) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStale

`func (o *SearchEdgeSummary) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *SearchEdgeSummary) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *SearchEdgeSummary) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *SearchEdgeSummary) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetStatus

`func (o *SearchEdgeSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchEdgeSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchEdgeSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchEdgeSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwName

`func (o *SearchEdgeSummary) GetSwName() string`

GetSwName returns the SwName field if non-nil, zero value otherwise.

### GetSwNameOk

`func (o *SearchEdgeSummary) GetSwNameOk() (*string, bool)`

GetSwNameOk returns a tuple with the SwName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwName

`func (o *SearchEdgeSummary) SetSwName(v string)`

SetSwName sets SwName field to given value.

### HasSwName

`func (o *SearchEdgeSummary) HasSwName() bool`

HasSwName returns a boolean if a field has been set.

### GetSwVersion

`func (o *SearchEdgeSummary) GetSwVersion() string`

GetSwVersion returns the SwVersion field if non-nil, zero value otherwise.

### GetSwVersionOk

`func (o *SearchEdgeSummary) GetSwVersionOk() (*string, bool)`

GetSwVersionOk returns a tuple with the SwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwVersion

`func (o *SearchEdgeSummary) SetSwVersion(v string)`

SetSwVersion sets SwVersion field to given value.

### HasSwVersion

`func (o *SearchEdgeSummary) HasSwVersion() bool`

HasSwVersion returns a boolean if a field has been set.

### GetTtConnCount

`func (o *SearchEdgeSummary) GetTtConnCount() int32`

GetTtConnCount returns the TtConnCount field if non-nil, zero value otherwise.

### GetTtConnCountOk

`func (o *SearchEdgeSummary) GetTtConnCountOk() (*int32, bool)`

GetTtConnCountOk returns a tuple with the TtConnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtConnCount

`func (o *SearchEdgeSummary) SetTtConnCount(v int32)`

SetTtConnCount sets TtConnCount field to given value.

### HasTtConnCount

`func (o *SearchEdgeSummary) HasTtConnCount() bool`

HasTtConnCount returns a boolean if a field has been set.

### GetUpgradeSummary

`func (o *SearchEdgeSummary) GetUpgradeSummary() UpgradeUpgradeSummary`

GetUpgradeSummary returns the UpgradeSummary field if non-nil, zero value otherwise.

### GetUpgradeSummaryOk

`func (o *SearchEdgeSummary) GetUpgradeSummaryOk() (*UpgradeUpgradeSummary, bool)`

GetUpgradeSummaryOk returns a tuple with the UpgradeSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeSummary

`func (o *SearchEdgeSummary) SetUpgradeSummary(v UpgradeUpgradeSummary)`

SetUpgradeSummary sets UpgradeSummary field to given value.

### HasUpgradeSummary

`func (o *SearchEdgeSummary) HasUpgradeSummary() bool`

HasUpgradeSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


