# V1SitesSiteIdDevicesGet200ResponseDeviceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**MaintenanceMode** | Pointer to **bool** |  | [optional] 
**ManagementIp** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**SoftwareVersion** | Pointer to **string** |  | [optional] 
**StagingMode** | Pointer to **bool** |  | [optional] 
**Uptime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**VrrpInterface** | Pointer to **string** |  | [optional] 
**VrrpState** | Pointer to **string** |  | [optional] 

## Methods

### NewV1SitesSiteIdDevicesGet200ResponseDeviceInner

`func NewV1SitesSiteIdDevicesGet200ResponseDeviceInner() *V1SitesSiteIdDevicesGet200ResponseDeviceInner`

NewV1SitesSiteIdDevicesGet200ResponseDeviceInner instantiates a new V1SitesSiteIdDevicesGet200ResponseDeviceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SitesSiteIdDevicesGet200ResponseDeviceInnerWithDefaults

`func NewV1SitesSiteIdDevicesGet200ResponseDeviceInnerWithDefaults() *V1SitesSiteIdDevicesGet200ResponseDeviceInner`

NewV1SitesSiteIdDevicesGet200ResponseDeviceInnerWithDefaults instantiates a new V1SitesSiteIdDevicesGet200ResponseDeviceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetHostname

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLocation

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetManagementIp

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetManagementIp() string`

GetManagementIp returns the ManagementIp field if non-nil, zero value otherwise.

### GetManagementIpOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetManagementIpOk() (*string, bool)`

GetManagementIpOk returns a tuple with the ManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIp

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetManagementIp(v string)`

SetManagementIp sets ManagementIp field to given value.

### HasManagementIp

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasManagementIp() bool`

HasManagementIp returns a boolean if a field has been set.

### GetModel

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRole

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSerialNumber

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSiteId

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetStagingMode

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetStagingMode() bool`

GetStagingMode returns the StagingMode field if non-nil, zero value otherwise.

### GetStagingModeOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetStagingModeOk() (*bool, bool)`

GetStagingModeOk returns a tuple with the StagingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingMode

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetStagingMode(v bool)`

SetStagingMode sets StagingMode field to given value.

### HasStagingMode

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasStagingMode() bool`

HasStagingMode returns a boolean if a field has been set.

### GetUptime

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetUptime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetUptimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetUptime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVrrpInterface

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetVrrpInterface() string`

GetVrrpInterface returns the VrrpInterface field if non-nil, zero value otherwise.

### GetVrrpInterfaceOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetVrrpInterfaceOk() (*string, bool)`

GetVrrpInterfaceOk returns a tuple with the VrrpInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpInterface

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetVrrpInterface(v string)`

SetVrrpInterface sets VrrpInterface field to given value.

### HasVrrpInterface

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasVrrpInterface() bool`

HasVrrpInterface returns a boolean if a field has been set.

### GetVrrpState

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetVrrpState() string`

GetVrrpState returns the VrrpState field if non-nil, zero value otherwise.

### GetVrrpStateOk

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) GetVrrpStateOk() (*string, bool)`

GetVrrpStateOk returns a tuple with the VrrpState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpState

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) SetVrrpState(v string)`

SetVrrpState sets VrrpState field to given value.

### HasVrrpState

`func (o *V1SitesSiteIdDevicesGet200ResponseDeviceInner) HasVrrpState() bool`

HasVrrpState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


