# V1TroubleshootingDeviceDeviceIdPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceMode** | Pointer to **bool** |  | [optional] 
**ColrActive** | Pointer to **bool** |  | [optional] 
**ControlPlane** | Pointer to [**V1TroubleshootingDeviceDeviceIdPost200ResponseControlPlane**](V1TroubleshootingDeviceDeviceIdPost200ResponseControlPlane.md) |  | [optional] 
**DataPlane** | Pointer to [**V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane**](V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane.md) |  | [optional] 
**Issues** | Pointer to [**[]V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner**](V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner.md) |  | [optional] 
**LifecycleStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SwVersion** | Pointer to **string** |  | [optional] 
**SwVersionV2** | Pointer to [**V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion**](V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion.md) |  | [optional] 
**SystemPlane** | Pointer to [**V1TroubleshootingDeviceDeviceIdPost200ResponseSystemPlane**](V1TroubleshootingDeviceDeviceIdPost200ResponseSystemPlane.md) |  | [optional] 
**UpSinceTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1TroubleshootingDeviceDeviceIdPost200Response

`func NewV1TroubleshootingDeviceDeviceIdPost200Response() *V1TroubleshootingDeviceDeviceIdPost200Response`

NewV1TroubleshootingDeviceDeviceIdPost200Response instantiates a new V1TroubleshootingDeviceDeviceIdPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TroubleshootingDeviceDeviceIdPost200ResponseWithDefaults

`func NewV1TroubleshootingDeviceDeviceIdPost200ResponseWithDefaults() *V1TroubleshootingDeviceDeviceIdPost200Response`

NewV1TroubleshootingDeviceDeviceIdPost200ResponseWithDefaults instantiates a new V1TroubleshootingDeviceDeviceIdPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceMode

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetColrActive

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetColrActive() bool`

GetColrActive returns the ColrActive field if non-nil, zero value otherwise.

### GetColrActiveOk

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetColrActiveOk() (*bool, bool)`

GetColrActiveOk returns a tuple with the ColrActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColrActive

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetColrActive(v bool)`

SetColrActive sets ColrActive field to given value.

### HasColrActive

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasColrActive() bool`

HasColrActive returns a boolean if a field has been set.

### GetControlPlane

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetControlPlane() V1TroubleshootingDeviceDeviceIdPost200ResponseControlPlane`

GetControlPlane returns the ControlPlane field if non-nil, zero value otherwise.

### GetControlPlaneOk

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetControlPlaneOk() (*V1TroubleshootingDeviceDeviceIdPost200ResponseControlPlane, bool)`

GetControlPlaneOk returns a tuple with the ControlPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlane

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetControlPlane(v V1TroubleshootingDeviceDeviceIdPost200ResponseControlPlane)`

SetControlPlane sets ControlPlane field to given value.

### HasControlPlane

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasControlPlane() bool`

HasControlPlane returns a boolean if a field has been set.

### GetDataPlane

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetDataPlane() V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane`

GetDataPlane returns the DataPlane field if non-nil, zero value otherwise.

### GetDataPlaneOk

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetDataPlaneOk() (*V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane, bool)`

GetDataPlaneOk returns a tuple with the DataPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPlane

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetDataPlane(v V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane)`

SetDataPlane sets DataPlane field to given value.

### HasDataPlane

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasDataPlane() bool`

HasDataPlane returns a boolean if a field has been set.

### GetIssues

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetIssues() []V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetIssuesOk() (*[]V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetIssues(v []V1BackboneHealthDeviceDeviceIdPost200ResponseIssuesInner)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetLifecycleStatus

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetLifecycleStatus() string`

GetLifecycleStatus returns the LifecycleStatus field if non-nil, zero value otherwise.

### GetLifecycleStatusOk

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetLifecycleStatusOk() (*string, bool)`

GetLifecycleStatusOk returns a tuple with the LifecycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStatus

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetLifecycleStatus(v string)`

SetLifecycleStatus sets LifecycleStatus field to given value.

### HasLifecycleStatus

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasLifecycleStatus() bool`

HasLifecycleStatus returns a boolean if a field has been set.

### GetStatus

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwVersion

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetSwVersion() string`

GetSwVersion returns the SwVersion field if non-nil, zero value otherwise.

### GetSwVersionOk

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetSwVersionOk() (*string, bool)`

GetSwVersionOk returns a tuple with the SwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwVersion

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetSwVersion(v string)`

SetSwVersion sets SwVersion field to given value.

### HasSwVersion

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasSwVersion() bool`

HasSwVersion returns a boolean if a field has been set.

### GetSwVersionV2

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetSwVersionV2() V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion`

GetSwVersionV2 returns the SwVersionV2 field if non-nil, zero value otherwise.

### GetSwVersionV2Ok

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetSwVersionV2Ok() (*V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion, bool)`

GetSwVersionV2Ok returns a tuple with the SwVersionV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwVersionV2

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetSwVersionV2(v V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInnerUpgradeSummaryRunningVersion)`

SetSwVersionV2 sets SwVersionV2 field to given value.

### HasSwVersionV2

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasSwVersionV2() bool`

HasSwVersionV2 returns a boolean if a field has been set.

### GetSystemPlane

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetSystemPlane() V1TroubleshootingDeviceDeviceIdPost200ResponseSystemPlane`

GetSystemPlane returns the SystemPlane field if non-nil, zero value otherwise.

### GetSystemPlaneOk

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetSystemPlaneOk() (*V1TroubleshootingDeviceDeviceIdPost200ResponseSystemPlane, bool)`

GetSystemPlaneOk returns a tuple with the SystemPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPlane

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetSystemPlane(v V1TroubleshootingDeviceDeviceIdPost200ResponseSystemPlane)`

SetSystemPlane sets SystemPlane field to given value.

### HasSystemPlane

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasSystemPlane() bool`

HasSystemPlane returns a boolean if a field has been set.

### GetUpSinceTs

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetUpSinceTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetUpSinceTs returns the UpSinceTs field if non-nil, zero value otherwise.

### GetUpSinceTsOk

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) GetUpSinceTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetUpSinceTsOk returns a tuple with the UpSinceTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSinceTs

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) SetUpSinceTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetUpSinceTs sets UpSinceTs field to given value.

### HasUpSinceTs

`func (o *V1TroubleshootingDeviceDeviceIdPost200Response) HasUpSinceTs() bool`

HasUpSinceTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


