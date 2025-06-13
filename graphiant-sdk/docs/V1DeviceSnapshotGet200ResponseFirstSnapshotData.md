# V1DeviceSnapshotGet200ResponseFirstSnapshotData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OSPFInstalledRouteCount** | Pointer to **int32** |  | [optional] 
**T2SessionCount** | Pointer to **int32** |  | [optional] 
**TWAMPSessionCount** | Pointer to **int32** |  | [optional] 
**BfdSessionCount** | Pointer to **int32** |  | [optional] 
**BgpNeighborIpList** | Pointer to **[]string** |  | [optional] 
**BgpSessionCount** | Pointer to **int32** |  | [optional] 
**DeviceRole** | Pointer to **string** |  | [optional] 
**FailedServicesCount** | Pointer to **int32** |  | [optional] 
**GraphnosVersion** | Pointer to **string** |  | [optional] 
**InstalledLabels** | Pointer to **int32** |  | [optional] 
**IpsecSessionCount** | Pointer to **int32** |  | [optional] 
**OngoingAlerts** | Pointer to **int32** |  | [optional] 
**OspfNeighborIpList** | Pointer to **[]string** |  | [optional] 
**OspfSessionCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1DeviceSnapshotGet200ResponseFirstSnapshotData

`func NewV1DeviceSnapshotGet200ResponseFirstSnapshotData() *V1DeviceSnapshotGet200ResponseFirstSnapshotData`

NewV1DeviceSnapshotGet200ResponseFirstSnapshotData instantiates a new V1DeviceSnapshotGet200ResponseFirstSnapshotData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeviceSnapshotGet200ResponseFirstSnapshotDataWithDefaults

`func NewV1DeviceSnapshotGet200ResponseFirstSnapshotDataWithDefaults() *V1DeviceSnapshotGet200ResponseFirstSnapshotData`

NewV1DeviceSnapshotGet200ResponseFirstSnapshotDataWithDefaults instantiates a new V1DeviceSnapshotGet200ResponseFirstSnapshotData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOSPFInstalledRouteCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetOSPFInstalledRouteCount() int32`

GetOSPFInstalledRouteCount returns the OSPFInstalledRouteCount field if non-nil, zero value otherwise.

### GetOSPFInstalledRouteCountOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetOSPFInstalledRouteCountOk() (*int32, bool)`

GetOSPFInstalledRouteCountOk returns a tuple with the OSPFInstalledRouteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSPFInstalledRouteCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetOSPFInstalledRouteCount(v int32)`

SetOSPFInstalledRouteCount sets OSPFInstalledRouteCount field to given value.

### HasOSPFInstalledRouteCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasOSPFInstalledRouteCount() bool`

HasOSPFInstalledRouteCount returns a boolean if a field has been set.

### GetT2SessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetT2SessionCount() int32`

GetT2SessionCount returns the T2SessionCount field if non-nil, zero value otherwise.

### GetT2SessionCountOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetT2SessionCountOk() (*int32, bool)`

GetT2SessionCountOk returns a tuple with the T2SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT2SessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetT2SessionCount(v int32)`

SetT2SessionCount sets T2SessionCount field to given value.

### HasT2SessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasT2SessionCount() bool`

HasT2SessionCount returns a boolean if a field has been set.

### GetTWAMPSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetTWAMPSessionCount() int32`

GetTWAMPSessionCount returns the TWAMPSessionCount field if non-nil, zero value otherwise.

### GetTWAMPSessionCountOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetTWAMPSessionCountOk() (*int32, bool)`

GetTWAMPSessionCountOk returns a tuple with the TWAMPSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTWAMPSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetTWAMPSessionCount(v int32)`

SetTWAMPSessionCount sets TWAMPSessionCount field to given value.

### HasTWAMPSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasTWAMPSessionCount() bool`

HasTWAMPSessionCount returns a boolean if a field has been set.

### GetBfdSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetBfdSessionCount() int32`

GetBfdSessionCount returns the BfdSessionCount field if non-nil, zero value otherwise.

### GetBfdSessionCountOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetBfdSessionCountOk() (*int32, bool)`

GetBfdSessionCountOk returns a tuple with the BfdSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetBfdSessionCount(v int32)`

SetBfdSessionCount sets BfdSessionCount field to given value.

### HasBfdSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasBfdSessionCount() bool`

HasBfdSessionCount returns a boolean if a field has been set.

### GetBgpNeighborIpList

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetBgpNeighborIpList() []string`

GetBgpNeighborIpList returns the BgpNeighborIpList field if non-nil, zero value otherwise.

### GetBgpNeighborIpListOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetBgpNeighborIpListOk() (*[]string, bool)`

GetBgpNeighborIpListOk returns a tuple with the BgpNeighborIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighborIpList

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetBgpNeighborIpList(v []string)`

SetBgpNeighborIpList sets BgpNeighborIpList field to given value.

### HasBgpNeighborIpList

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasBgpNeighborIpList() bool`

HasBgpNeighborIpList returns a boolean if a field has been set.

### GetBgpSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetBgpSessionCount() int32`

GetBgpSessionCount returns the BgpSessionCount field if non-nil, zero value otherwise.

### GetBgpSessionCountOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetBgpSessionCountOk() (*int32, bool)`

GetBgpSessionCountOk returns a tuple with the BgpSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetBgpSessionCount(v int32)`

SetBgpSessionCount sets BgpSessionCount field to given value.

### HasBgpSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasBgpSessionCount() bool`

HasBgpSessionCount returns a boolean if a field has been set.

### GetDeviceRole

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetDeviceRole() string`

GetDeviceRole returns the DeviceRole field if non-nil, zero value otherwise.

### GetDeviceRoleOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetDeviceRoleOk() (*string, bool)`

GetDeviceRoleOk returns a tuple with the DeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRole

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetDeviceRole(v string)`

SetDeviceRole sets DeviceRole field to given value.

### HasDeviceRole

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasDeviceRole() bool`

HasDeviceRole returns a boolean if a field has been set.

### GetFailedServicesCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetFailedServicesCount() int32`

GetFailedServicesCount returns the FailedServicesCount field if non-nil, zero value otherwise.

### GetFailedServicesCountOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetFailedServicesCountOk() (*int32, bool)`

GetFailedServicesCountOk returns a tuple with the FailedServicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedServicesCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetFailedServicesCount(v int32)`

SetFailedServicesCount sets FailedServicesCount field to given value.

### HasFailedServicesCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasFailedServicesCount() bool`

HasFailedServicesCount returns a boolean if a field has been set.

### GetGraphnosVersion

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetGraphnosVersion() string`

GetGraphnosVersion returns the GraphnosVersion field if non-nil, zero value otherwise.

### GetGraphnosVersionOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetGraphnosVersionOk() (*string, bool)`

GetGraphnosVersionOk returns a tuple with the GraphnosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphnosVersion

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetGraphnosVersion(v string)`

SetGraphnosVersion sets GraphnosVersion field to given value.

### HasGraphnosVersion

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasGraphnosVersion() bool`

HasGraphnosVersion returns a boolean if a field has been set.

### GetInstalledLabels

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetInstalledLabels() int32`

GetInstalledLabels returns the InstalledLabels field if non-nil, zero value otherwise.

### GetInstalledLabelsOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetInstalledLabelsOk() (*int32, bool)`

GetInstalledLabelsOk returns a tuple with the InstalledLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledLabels

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetInstalledLabels(v int32)`

SetInstalledLabels sets InstalledLabels field to given value.

### HasInstalledLabels

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasInstalledLabels() bool`

HasInstalledLabels returns a boolean if a field has been set.

### GetIpsecSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetIpsecSessionCount() int32`

GetIpsecSessionCount returns the IpsecSessionCount field if non-nil, zero value otherwise.

### GetIpsecSessionCountOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetIpsecSessionCountOk() (*int32, bool)`

GetIpsecSessionCountOk returns a tuple with the IpsecSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetIpsecSessionCount(v int32)`

SetIpsecSessionCount sets IpsecSessionCount field to given value.

### HasIpsecSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasIpsecSessionCount() bool`

HasIpsecSessionCount returns a boolean if a field has been set.

### GetOngoingAlerts

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetOngoingAlerts() int32`

GetOngoingAlerts returns the OngoingAlerts field if non-nil, zero value otherwise.

### GetOngoingAlertsOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetOngoingAlertsOk() (*int32, bool)`

GetOngoingAlertsOk returns a tuple with the OngoingAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoingAlerts

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetOngoingAlerts(v int32)`

SetOngoingAlerts sets OngoingAlerts field to given value.

### HasOngoingAlerts

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasOngoingAlerts() bool`

HasOngoingAlerts returns a boolean if a field has been set.

### GetOspfNeighborIpList

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetOspfNeighborIpList() []string`

GetOspfNeighborIpList returns the OspfNeighborIpList field if non-nil, zero value otherwise.

### GetOspfNeighborIpListOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetOspfNeighborIpListOk() (*[]string, bool)`

GetOspfNeighborIpListOk returns a tuple with the OspfNeighborIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfNeighborIpList

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetOspfNeighborIpList(v []string)`

SetOspfNeighborIpList sets OspfNeighborIpList field to given value.

### HasOspfNeighborIpList

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasOspfNeighborIpList() bool`

HasOspfNeighborIpList returns a boolean if a field has been set.

### GetOspfSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetOspfSessionCount() int32`

GetOspfSessionCount returns the OspfSessionCount field if non-nil, zero value otherwise.

### GetOspfSessionCountOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) GetOspfSessionCountOk() (*int32, bool)`

GetOspfSessionCountOk returns a tuple with the OspfSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) SetOspfSessionCount(v int32)`

SetOspfSessionCount sets OspfSessionCount field to given value.

### HasOspfSessionCount

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshotData) HasOspfSessionCount() bool`

HasOspfSessionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


