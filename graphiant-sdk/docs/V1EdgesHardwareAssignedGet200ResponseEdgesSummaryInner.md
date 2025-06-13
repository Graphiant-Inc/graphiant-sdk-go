# V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedOn** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**DiscoveredLocation** | Pointer to **string** |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**EnterpriseName** | Pointer to **string** |  | [optional] 
**FirstAppearedOn** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**IpDetected** | Pointer to **string** |  | [optional] 
**IsHardware** | Pointer to **bool** |  | [optional] 
**IsNew** | Pointer to **bool** |  | [optional] 
**IsRequested** | Pointer to **bool** |  | [optional] 
**LastBootedAt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Location** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation.md) |  | [optional] 
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
**UpgradeSummary** | Pointer to [**V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary**](V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary.md) |  | [optional] 

## Methods

### NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner

`func NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner() *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner`

NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner instantiates a new V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerWithDefaults

`func NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerWithDefaults() *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner`

NewV1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerWithDefaults instantiates a new V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedOn

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetAssignedOn() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetAssignedOn returns the AssignedOn field if non-nil, zero value otherwise.

### GetAssignedOnOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetAssignedOnOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetAssignedOnOk returns a tuple with the AssignedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedOn

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetAssignedOn(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetAssignedOn sets AssignedOn field to given value.

### HasAssignedOn

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasAssignedOn() bool`

HasAssignedOn returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDiscoveredLocation

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetDiscoveredLocation() string`

GetDiscoveredLocation returns the DiscoveredLocation field if non-nil, zero value otherwise.

### GetDiscoveredLocationOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetDiscoveredLocationOk() (*string, bool)`

GetDiscoveredLocationOk returns a tuple with the DiscoveredLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredLocation

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetDiscoveredLocation(v string)`

SetDiscoveredLocation sets DiscoveredLocation field to given value.

### HasDiscoveredLocation

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasDiscoveredLocation() bool`

HasDiscoveredLocation returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetEnterpriseName

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetEnterpriseName() string`

GetEnterpriseName returns the EnterpriseName field if non-nil, zero value otherwise.

### GetEnterpriseNameOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetEnterpriseNameOk() (*string, bool)`

GetEnterpriseNameOk returns a tuple with the EnterpriseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseName

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetEnterpriseName(v string)`

SetEnterpriseName sets EnterpriseName field to given value.

### HasEnterpriseName

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasEnterpriseName() bool`

HasEnterpriseName returns a boolean if a field has been set.

### GetFirstAppearedOn

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetFirstAppearedOn() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetFirstAppearedOn returns the FirstAppearedOn field if non-nil, zero value otherwise.

### GetFirstAppearedOnOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetFirstAppearedOnOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetFirstAppearedOnOk returns a tuple with the FirstAppearedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAppearedOn

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetFirstAppearedOn(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetFirstAppearedOn sets FirstAppearedOn field to given value.

### HasFirstAppearedOn

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasFirstAppearedOn() bool`

HasFirstAppearedOn returns a boolean if a field has been set.

### GetHostname

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpDetected

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetIpDetected() string`

GetIpDetected returns the IpDetected field if non-nil, zero value otherwise.

### GetIpDetectedOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetIpDetectedOk() (*string, bool)`

GetIpDetectedOk returns a tuple with the IpDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDetected

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetIpDetected(v string)`

SetIpDetected sets IpDetected field to given value.

### HasIpDetected

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasIpDetected() bool`

HasIpDetected returns a boolean if a field has been set.

### GetIsHardware

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetIsHardware() bool`

GetIsHardware returns the IsHardware field if non-nil, zero value otherwise.

### GetIsHardwareOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetIsHardwareOk() (*bool, bool)`

GetIsHardwareOk returns a tuple with the IsHardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardware

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetIsHardware(v bool)`

SetIsHardware sets IsHardware field to given value.

### HasIsHardware

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasIsHardware() bool`

HasIsHardware returns a boolean if a field has been set.

### GetIsNew

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.

### HasIsNew

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasIsNew() bool`

HasIsNew returns a boolean if a field has been set.

### GetIsRequested

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetIsRequested() bool`

GetIsRequested returns the IsRequested field if non-nil, zero value otherwise.

### GetIsRequestedOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetIsRequestedOk() (*bool, bool)`

GetIsRequestedOk returns a tuple with the IsRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequested

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetIsRequested(v bool)`

SetIsRequested sets IsRequested field to given value.

### HasIsRequested

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasIsRequested() bool`

HasIsRequested returns a boolean if a field has been set.

### GetLastBootedAt

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetLastBootedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetLastBootedAt returns the LastBootedAt field if non-nil, zero value otherwise.

### GetLastBootedAtOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetLastBootedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetLastBootedAtOk returns a tuple with the LastBootedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBootedAt

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetLastBootedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetLastBootedAt sets LastBootedAt field to given value.

### HasLastBootedAt

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasLastBootedAt() bool`

HasLastBootedAt returns a boolean if a field has been set.

### GetLocation

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetLocation() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetLocationOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetLocation(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetModel

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOverrideRegion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetOverrideRegion() string`

GetOverrideRegion returns the OverrideRegion field if non-nil, zero value otherwise.

### GetOverrideRegionOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetOverrideRegionOk() (*string, bool)`

GetOverrideRegionOk returns a tuple with the OverrideRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRegion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetOverrideRegion(v string)`

SetOverrideRegion sets OverrideRegion field to given value.

### HasOverrideRegion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasOverrideRegion() bool`

HasOverrideRegion returns a boolean if a field has been set.

### GetParentEnterpriseName

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetParentEnterpriseName() string`

GetParentEnterpriseName returns the ParentEnterpriseName field if non-nil, zero value otherwise.

### GetParentEnterpriseNameOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetParentEnterpriseNameOk() (*string, bool)`

GetParentEnterpriseNameOk returns a tuple with the ParentEnterpriseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEnterpriseName

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetParentEnterpriseName(v string)`

SetParentEnterpriseName sets ParentEnterpriseName field to given value.

### HasParentEnterpriseName

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasParentEnterpriseName() bool`

HasParentEnterpriseName returns a boolean if a field has been set.

### GetPortalStatus

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetPortalStatus() string`

GetPortalStatus returns the PortalStatus field if non-nil, zero value otherwise.

### GetPortalStatusOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetPortalStatusOk() (*string, bool)`

GetPortalStatusOk returns a tuple with the PortalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalStatus

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetPortalStatus(v string)`

SetPortalStatus sets PortalStatus field to given value.

### HasPortalStatus

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasPortalStatus() bool`

HasPortalStatus returns a boolean if a field has been set.

### GetRegion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRole

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSerialNum

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetSerialNum() string`

GetSerialNum returns the SerialNum field if non-nil, zero value otherwise.

### GetSerialNumOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetSerialNumOk() (*string, bool)`

GetSerialNumOk returns a tuple with the SerialNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNum

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetSerialNum(v string)`

SetSerialNum sets SerialNum field to given value.

### HasSerialNum

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasSerialNum() bool`

HasSerialNum returns a boolean if a field has been set.

### GetSite

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSiteId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStale

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetStatus

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwName

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetSwName() string`

GetSwName returns the SwName field if non-nil, zero value otherwise.

### GetSwNameOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetSwNameOk() (*string, bool)`

GetSwNameOk returns a tuple with the SwName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwName

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetSwName(v string)`

SetSwName sets SwName field to given value.

### HasSwName

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasSwName() bool`

HasSwName returns a boolean if a field has been set.

### GetSwVersion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetSwVersion() string`

GetSwVersion returns the SwVersion field if non-nil, zero value otherwise.

### GetSwVersionOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetSwVersionOk() (*string, bool)`

GetSwVersionOk returns a tuple with the SwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwVersion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetSwVersion(v string)`

SetSwVersion sets SwVersion field to given value.

### HasSwVersion

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasSwVersion() bool`

HasSwVersion returns a boolean if a field has been set.

### GetTtConnCount

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetTtConnCount() int32`

GetTtConnCount returns the TtConnCount field if non-nil, zero value otherwise.

### GetTtConnCountOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetTtConnCountOk() (*int32, bool)`

GetTtConnCountOk returns a tuple with the TtConnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtConnCount

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetTtConnCount(v int32)`

SetTtConnCount sets TtConnCount field to given value.

### HasTtConnCount

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasTtConnCount() bool`

HasTtConnCount returns a boolean if a field has been set.

### GetUpgradeSummary

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetUpgradeSummary() V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary`

GetUpgradeSummary returns the UpgradeSummary field if non-nil, zero value otherwise.

### GetUpgradeSummaryOk

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) GetUpgradeSummaryOk() (*V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary, bool)`

GetUpgradeSummaryOk returns a tuple with the UpgradeSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeSummary

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) SetUpgradeSummary(v V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummary)`

SetUpgradeSummary sets UpgradeSummary field to given value.

### HasUpgradeSummary

`func (o *V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner) HasUpgradeSummary() bool`

HasUpgradeSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


