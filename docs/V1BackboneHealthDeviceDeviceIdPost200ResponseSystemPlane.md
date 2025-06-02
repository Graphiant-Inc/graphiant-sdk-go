# V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**[]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner**](V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner.md) |  | [optional] 
**Crashes** | Pointer to [**[]V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneCrashesInner**](V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneCrashesInner.md) |  | [optional] 
**Disk** | Pointer to [**[]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner**](V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner.md) |  | [optional] 
**LastCrash** | Pointer to [**V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneLastCrash**](V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneLastCrash.md) |  | [optional] 
**MaintenanceWindows** | Pointer to [**[]V2NotificationlistPostRequest**](V2NotificationlistPostRequest.md) |  | [optional] 
**Memory** | Pointer to [**[]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner**](V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner.md) |  | [optional] 
**Overheating** | Pointer to [**[]V2NotificationlistPostRequest**](V2NotificationlistPostRequest.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TemperatureSeries** | Pointer to [**[]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner**](V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner.md) |  | [optional] 

## Methods

### NewV1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane

`func NewV1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane() *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane`

NewV1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneWithDefaults

`func NewV1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneWithDefaults() *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane`

NewV1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneWithDefaults instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetCpu() []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetCpuOk() (*[]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) SetCpu(v []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCrashes

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetCrashes() []V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneCrashesInner`

GetCrashes returns the Crashes field if non-nil, zero value otherwise.

### GetCrashesOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetCrashesOk() (*[]V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneCrashesInner, bool)`

GetCrashesOk returns a tuple with the Crashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrashes

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) SetCrashes(v []V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneCrashesInner)`

SetCrashes sets Crashes field to given value.

### HasCrashes

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) HasCrashes() bool`

HasCrashes returns a boolean if a field has been set.

### GetDisk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetDisk() []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetDiskOk() (*[]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) SetDisk(v []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetLastCrash

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetLastCrash() V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneLastCrash`

GetLastCrash returns the LastCrash field if non-nil, zero value otherwise.

### GetLastCrashOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetLastCrashOk() (*V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneLastCrash, bool)`

GetLastCrashOk returns a tuple with the LastCrash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCrash

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) SetLastCrash(v V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlaneLastCrash)`

SetLastCrash sets LastCrash field to given value.

### HasLastCrash

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) HasLastCrash() bool`

HasLastCrash returns a boolean if a field has been set.

### GetMaintenanceWindows

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetMaintenanceWindows() []V2NotificationlistPostRequest`

GetMaintenanceWindows returns the MaintenanceWindows field if non-nil, zero value otherwise.

### GetMaintenanceWindowsOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetMaintenanceWindowsOk() (*[]V2NotificationlistPostRequest, bool)`

GetMaintenanceWindowsOk returns a tuple with the MaintenanceWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindows

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) SetMaintenanceWindows(v []V2NotificationlistPostRequest)`

SetMaintenanceWindows sets MaintenanceWindows field to given value.

### HasMaintenanceWindows

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) HasMaintenanceWindows() bool`

HasMaintenanceWindows returns a boolean if a field has been set.

### GetMemory

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetMemory() []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetMemoryOk() (*[]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) SetMemory(v []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetOverheating

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetOverheating() []V2NotificationlistPostRequest`

GetOverheating returns the Overheating field if non-nil, zero value otherwise.

### GetOverheatingOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetOverheatingOk() (*[]V2NotificationlistPostRequest, bool)`

GetOverheatingOk returns a tuple with the Overheating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverheating

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) SetOverheating(v []V2NotificationlistPostRequest)`

SetOverheating sets Overheating field to given value.

### HasOverheating

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) HasOverheating() bool`

HasOverheating returns a boolean if a field has been set.

### GetStatus

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemperatureSeries

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetTemperatureSeries() []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner`

GetTemperatureSeries returns the TemperatureSeries field if non-nil, zero value otherwise.

### GetTemperatureSeriesOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) GetTemperatureSeriesOk() (*[]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner, bool)`

GetTemperatureSeriesOk returns a tuple with the TemperatureSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureSeries

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) SetTemperatureSeries(v []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInnerTransitionsInner)`

SetTemperatureSeries sets TemperatureSeries field to given value.

### HasTemperatureSeries

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseSystemPlane) HasTemperatureSeries() bool`

HasTemperatureSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


