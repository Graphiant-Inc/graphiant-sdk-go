# V1BackboneHealthOverviewPost200ResponseDevicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlStatus** | Pointer to **string** |  | [optional] 
**DataStatus** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceRole** | Pointer to **string** |  | [optional] 
**OverallStatus** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**V2AssuranceTopologyInventoryPost200ResponseRegionsInner**](V2AssuranceTopologyInventoryPost200ResponseRegionsInner.md) |  | [optional] 
**SelectedStatus** | Pointer to **string** |  | [optional] 
**SystemStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewV1BackboneHealthOverviewPost200ResponseDevicesInner

`func NewV1BackboneHealthOverviewPost200ResponseDevicesInner() *V1BackboneHealthOverviewPost200ResponseDevicesInner`

NewV1BackboneHealthOverviewPost200ResponseDevicesInner instantiates a new V1BackboneHealthOverviewPost200ResponseDevicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BackboneHealthOverviewPost200ResponseDevicesInnerWithDefaults

`func NewV1BackboneHealthOverviewPost200ResponseDevicesInnerWithDefaults() *V1BackboneHealthOverviewPost200ResponseDevicesInner`

NewV1BackboneHealthOverviewPost200ResponseDevicesInnerWithDefaults instantiates a new V1BackboneHealthOverviewPost200ResponseDevicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetControlStatus() string`

GetControlStatus returns the ControlStatus field if non-nil, zero value otherwise.

### GetControlStatusOk

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetControlStatusOk() (*string, bool)`

GetControlStatusOk returns a tuple with the ControlStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetControlStatus(v string)`

SetControlStatus sets ControlStatus field to given value.

### HasControlStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasControlStatus() bool`

HasControlStatus returns a boolean if a field has been set.

### GetDataStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDataStatus() string`

GetDataStatus returns the DataStatus field if non-nil, zero value otherwise.

### GetDataStatusOk

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDataStatusOk() (*string, bool)`

GetDataStatusOk returns a tuple with the DataStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetDataStatus(v string)`

SetDataStatus sets DataStatus field to given value.

### HasDataStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasDataStatus() bool`

HasDataStatus returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceRole

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceRole() string`

GetDeviceRole returns the DeviceRole field if non-nil, zero value otherwise.

### GetDeviceRoleOk

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetDeviceRoleOk() (*string, bool)`

GetDeviceRoleOk returns a tuple with the DeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRole

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetDeviceRole(v string)`

SetDeviceRole sets DeviceRole field to given value.

### HasDeviceRole

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasDeviceRole() bool`

HasDeviceRole returns a boolean if a field has been set.

### GetOverallStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### GetRegion

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetRegion() V2AssuranceTopologyInventoryPost200ResponseRegionsInner`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetRegionOk() (*V2AssuranceTopologyInventoryPost200ResponseRegionsInner, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetRegion(v V2AssuranceTopologyInventoryPost200ResponseRegionsInner)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSelectedStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetSelectedStatus() string`

GetSelectedStatus returns the SelectedStatus field if non-nil, zero value otherwise.

### GetSelectedStatusOk

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetSelectedStatusOk() (*string, bool)`

GetSelectedStatusOk returns a tuple with the SelectedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetSelectedStatus(v string)`

SetSelectedStatus sets SelectedStatus field to given value.

### HasSelectedStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasSelectedStatus() bool`

HasSelectedStatus returns a boolean if a field has been set.

### GetSystemStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetSystemStatus() string`

GetSystemStatus returns the SystemStatus field if non-nil, zero value otherwise.

### GetSystemStatusOk

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) GetSystemStatusOk() (*string, bool)`

GetSystemStatusOk returns a tuple with the SystemStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) SetSystemStatus(v string)`

SetSystemStatus sets SystemStatus field to given value.

### HasSystemStatus

`func (o *V1BackboneHealthOverviewPost200ResponseDevicesInner) HasSystemStatus() bool`

HasSystemStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


