# V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptMode** | Pointer to **bool** |  | [optional] 
**AllowInterOperability** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Preempt** | Pointer to **bool** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**TrackedInterfaces** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner.md) |  | [optional] 
**VirtualIp** | Pointer to **string** |  | [optional] 
**VirtualRouterId** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroupWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroupWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup`

NewV1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroupWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptMode

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetAcceptMode() bool`

GetAcceptMode returns the AcceptMode field if non-nil, zero value otherwise.

### GetAcceptModeOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetAcceptModeOk() (*bool, bool)`

GetAcceptModeOk returns a tuple with the AcceptMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptMode

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) SetAcceptMode(v bool)`

SetAcceptMode sets AcceptMode field to given value.

### HasAcceptMode

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) HasAcceptMode() bool`

HasAcceptMode returns a boolean if a field has been set.

### GetAllowInterOperability

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetAllowInterOperability() bool`

GetAllowInterOperability returns the AllowInterOperability field if non-nil, zero value otherwise.

### GetAllowInterOperabilityOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetAllowInterOperabilityOk() (*bool, bool)`

GetAllowInterOperabilityOk returns a tuple with the AllowInterOperability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInterOperability

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) SetAllowInterOperability(v bool)`

SetAllowInterOperability sets AllowInterOperability field to given value.

### HasAllowInterOperability

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) HasAllowInterOperability() bool`

HasAllowInterOperability returns a boolean if a field has been set.

### GetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPreempt

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetPreempt() bool`

GetPreempt returns the Preempt field if non-nil, zero value otherwise.

### GetPreemptOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetPreemptOk() (*bool, bool)`

GetPreemptOk returns a tuple with the Preempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreempt

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) SetPreempt(v bool)`

SetPreempt sets Preempt field to given value.

### HasPreempt

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) HasPreempt() bool`

HasPreempt returns a boolean if a field has been set.

### GetPriority

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTrackedInterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetTrackedInterfaces() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner`

GetTrackedInterfaces returns the TrackedInterfaces field if non-nil, zero value otherwise.

### GetTrackedInterfacesOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetTrackedInterfacesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner, bool)`

GetTrackedInterfacesOk returns a tuple with the TrackedInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedInterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) SetTrackedInterfaces(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner)`

SetTrackedInterfaces sets TrackedInterfaces field to given value.

### HasTrackedInterfaces

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) HasTrackedInterfaces() bool`

HasTrackedInterfaces returns a boolean if a field has been set.

### GetVirtualIp

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetVirtualIp() string`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetVirtualIpOk() (*string, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) SetVirtualIp(v string)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### GetVirtualRouterId

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetVirtualRouterId() int32`

GetVirtualRouterId returns the VirtualRouterId field if non-nil, zero value otherwise.

### GetVirtualRouterIdOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) GetVirtualRouterIdOk() (*int32, bool)`

GetVirtualRouterIdOk returns a tuple with the VirtualRouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouterId

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) SetVirtualRouterId(v int32)`

SetVirtualRouterId sets VirtualRouterId field to given value.

### HasVirtualRouterId

`func (o *V1DevicesDeviceIdConfigPutRequestCoreInterfacesValueInterfaceGwGwVrrpGroup) HasVirtualRouterId() bool`

HasVirtualRouterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


