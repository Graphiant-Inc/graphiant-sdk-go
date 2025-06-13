# V1HealthcheckDevicesGet200ResponseDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpCoreState** | Pointer to **[]bool** |  | [optional] 
**BgpOdpState** | Pointer to **[]bool** |  | [optional] 
**CoreTunnelState** | Pointer to **[]bool** |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**GnmiState** | Pointer to **bool** |  | [optional] 
**IpsecCoreStatus** | Pointer to **string** |  | [optional] 
**IpsecOdpStatus** | Pointer to **string** |  | [optional] 
**OdpStatus** | Pointer to [**V1HealthcheckDevicesGet200ResponseDetailsInnerOdpStatus**](V1HealthcheckDevicesGet200ResponseDetailsInnerOdpStatus.md) |  | [optional] 
**OdpTunnelState** | Pointer to **[]bool** |  | [optional] 
**OnboardingStatus** | Pointer to [**V1HealthcheckDevicesGet200ResponseDetailsInnerOnboardingStatus**](V1HealthcheckDevicesGet200ResponseDetailsInnerOnboardingStatus.md) |  | [optional] 
**T2Status** | Pointer to [**V1HealthcheckDevicesGet200ResponseDetailsInnerOdpStatus**](V1HealthcheckDevicesGet200ResponseDetailsInnerOdpStatus.md) |  | [optional] 
**TopoGwState** | Pointer to **string** |  | [optional] 
**TtTunnelState** | Pointer to **[]bool** |  | [optional] 

## Methods

### NewV1HealthcheckDevicesGet200ResponseDetailsInner

`func NewV1HealthcheckDevicesGet200ResponseDetailsInner() *V1HealthcheckDevicesGet200ResponseDetailsInner`

NewV1HealthcheckDevicesGet200ResponseDetailsInner instantiates a new V1HealthcheckDevicesGet200ResponseDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HealthcheckDevicesGet200ResponseDetailsInnerWithDefaults

`func NewV1HealthcheckDevicesGet200ResponseDetailsInnerWithDefaults() *V1HealthcheckDevicesGet200ResponseDetailsInner`

NewV1HealthcheckDevicesGet200ResponseDetailsInnerWithDefaults instantiates a new V1HealthcheckDevicesGet200ResponseDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpCoreState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetBgpCoreState() []bool`

GetBgpCoreState returns the BgpCoreState field if non-nil, zero value otherwise.

### GetBgpCoreStateOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetBgpCoreStateOk() (*[]bool, bool)`

GetBgpCoreStateOk returns a tuple with the BgpCoreState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpCoreState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetBgpCoreState(v []bool)`

SetBgpCoreState sets BgpCoreState field to given value.

### HasBgpCoreState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasBgpCoreState() bool`

HasBgpCoreState returns a boolean if a field has been set.

### GetBgpOdpState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetBgpOdpState() []bool`

GetBgpOdpState returns the BgpOdpState field if non-nil, zero value otherwise.

### GetBgpOdpStateOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetBgpOdpStateOk() (*[]bool, bool)`

GetBgpOdpStateOk returns a tuple with the BgpOdpState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpOdpState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetBgpOdpState(v []bool)`

SetBgpOdpState sets BgpOdpState field to given value.

### HasBgpOdpState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasBgpOdpState() bool`

HasBgpOdpState returns a boolean if a field has been set.

### GetCoreTunnelState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetCoreTunnelState() []bool`

GetCoreTunnelState returns the CoreTunnelState field if non-nil, zero value otherwise.

### GetCoreTunnelStateOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetCoreTunnelStateOk() (*[]bool, bool)`

GetCoreTunnelStateOk returns a tuple with the CoreTunnelState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreTunnelState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetCoreTunnelState(v []bool)`

SetCoreTunnelState sets CoreTunnelState field to given value.

### HasCoreTunnelState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasCoreTunnelState() bool`

HasCoreTunnelState returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetGnmiState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetGnmiState() bool`

GetGnmiState returns the GnmiState field if non-nil, zero value otherwise.

### GetGnmiStateOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetGnmiStateOk() (*bool, bool)`

GetGnmiStateOk returns a tuple with the GnmiState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnmiState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetGnmiState(v bool)`

SetGnmiState sets GnmiState field to given value.

### HasGnmiState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasGnmiState() bool`

HasGnmiState returns a boolean if a field has been set.

### GetIpsecCoreStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetIpsecCoreStatus() string`

GetIpsecCoreStatus returns the IpsecCoreStatus field if non-nil, zero value otherwise.

### GetIpsecCoreStatusOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetIpsecCoreStatusOk() (*string, bool)`

GetIpsecCoreStatusOk returns a tuple with the IpsecCoreStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecCoreStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetIpsecCoreStatus(v string)`

SetIpsecCoreStatus sets IpsecCoreStatus field to given value.

### HasIpsecCoreStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasIpsecCoreStatus() bool`

HasIpsecCoreStatus returns a boolean if a field has been set.

### GetIpsecOdpStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetIpsecOdpStatus() string`

GetIpsecOdpStatus returns the IpsecOdpStatus field if non-nil, zero value otherwise.

### GetIpsecOdpStatusOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetIpsecOdpStatusOk() (*string, bool)`

GetIpsecOdpStatusOk returns a tuple with the IpsecOdpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecOdpStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetIpsecOdpStatus(v string)`

SetIpsecOdpStatus sets IpsecOdpStatus field to given value.

### HasIpsecOdpStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasIpsecOdpStatus() bool`

HasIpsecOdpStatus returns a boolean if a field has been set.

### GetOdpStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetOdpStatus() V1HealthcheckDevicesGet200ResponseDetailsInnerOdpStatus`

GetOdpStatus returns the OdpStatus field if non-nil, zero value otherwise.

### GetOdpStatusOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetOdpStatusOk() (*V1HealthcheckDevicesGet200ResponseDetailsInnerOdpStatus, bool)`

GetOdpStatusOk returns a tuple with the OdpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdpStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetOdpStatus(v V1HealthcheckDevicesGet200ResponseDetailsInnerOdpStatus)`

SetOdpStatus sets OdpStatus field to given value.

### HasOdpStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasOdpStatus() bool`

HasOdpStatus returns a boolean if a field has been set.

### GetOdpTunnelState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetOdpTunnelState() []bool`

GetOdpTunnelState returns the OdpTunnelState field if non-nil, zero value otherwise.

### GetOdpTunnelStateOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetOdpTunnelStateOk() (*[]bool, bool)`

GetOdpTunnelStateOk returns a tuple with the OdpTunnelState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdpTunnelState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetOdpTunnelState(v []bool)`

SetOdpTunnelState sets OdpTunnelState field to given value.

### HasOdpTunnelState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasOdpTunnelState() bool`

HasOdpTunnelState returns a boolean if a field has been set.

### GetOnboardingStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetOnboardingStatus() V1HealthcheckDevicesGet200ResponseDetailsInnerOnboardingStatus`

GetOnboardingStatus returns the OnboardingStatus field if non-nil, zero value otherwise.

### GetOnboardingStatusOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetOnboardingStatusOk() (*V1HealthcheckDevicesGet200ResponseDetailsInnerOnboardingStatus, bool)`

GetOnboardingStatusOk returns a tuple with the OnboardingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetOnboardingStatus(v V1HealthcheckDevicesGet200ResponseDetailsInnerOnboardingStatus)`

SetOnboardingStatus sets OnboardingStatus field to given value.

### HasOnboardingStatus

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasOnboardingStatus() bool`

HasOnboardingStatus returns a boolean if a field has been set.

### GetT2Status

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetT2Status() V1HealthcheckDevicesGet200ResponseDetailsInnerOdpStatus`

GetT2Status returns the T2Status field if non-nil, zero value otherwise.

### GetT2StatusOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetT2StatusOk() (*V1HealthcheckDevicesGet200ResponseDetailsInnerOdpStatus, bool)`

GetT2StatusOk returns a tuple with the T2Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT2Status

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetT2Status(v V1HealthcheckDevicesGet200ResponseDetailsInnerOdpStatus)`

SetT2Status sets T2Status field to given value.

### HasT2Status

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasT2Status() bool`

HasT2Status returns a boolean if a field has been set.

### GetTopoGwState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetTopoGwState() string`

GetTopoGwState returns the TopoGwState field if non-nil, zero value otherwise.

### GetTopoGwStateOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetTopoGwStateOk() (*string, bool)`

GetTopoGwStateOk returns a tuple with the TopoGwState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopoGwState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetTopoGwState(v string)`

SetTopoGwState sets TopoGwState field to given value.

### HasTopoGwState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasTopoGwState() bool`

HasTopoGwState returns a boolean if a field has been set.

### GetTtTunnelState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetTtTunnelState() []bool`

GetTtTunnelState returns the TtTunnelState field if non-nil, zero value otherwise.

### GetTtTunnelStateOk

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) GetTtTunnelStateOk() (*[]bool, bool)`

GetTtTunnelStateOk returns a tuple with the TtTunnelState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtTunnelState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) SetTtTunnelState(v []bool)`

SetTtTunnelState sets TtTunnelState field to given value.

### HasTtTunnelState

`func (o *V1HealthcheckDevicesGet200ResponseDetailsInner) HasTtTunnelState() bool`

HasTtTunnelState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


