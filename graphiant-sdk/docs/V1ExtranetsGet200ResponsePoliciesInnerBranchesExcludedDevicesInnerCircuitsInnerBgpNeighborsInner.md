# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamilies** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerAddressFamiliesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerAddressFamiliesInner.md) |  | [optional] 
**AllowAsIn** | Pointer to **int32** |  | [optional] 
**AsOverride** | Pointer to **bool** |  | [optional] 
**Bfd** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd.md) |  | [optional] 
**BfdNeighbor** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor.md) |  | [optional] 
**BgpType** | Pointer to **string** |  | [optional] 
**DefaultOriginate** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HoldTimer** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**KeepaliveTimer** | Pointer to **int32** |  | [optional] 
**LocalAddress** | Pointer to **string** |  | [optional] 
**LocalInterface** | Pointer to **string** |  | [optional] 
**MaxPrefix** | Pointer to **int32** |  | [optional] 
**Md5Password** | Pointer to **string** |  | [optional] 
**MultiHop** | Pointer to **int32** |  | [optional] 
**PeerAsn** | Pointer to **int32** |  | [optional] 
**RemoteAddress** | Pointer to **string** |  | [optional] 
**RemovePrivateAs** | Pointer to **bool** |  | [optional] 
**SendCommunity** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TimeSinceLastOperChange** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamilies

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetAddressFamilies() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerAddressFamiliesInner`

GetAddressFamilies returns the AddressFamilies field if non-nil, zero value otherwise.

### GetAddressFamiliesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetAddressFamiliesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerAddressFamiliesInner, bool)`

GetAddressFamiliesOk returns a tuple with the AddressFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamilies

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetAddressFamilies(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerAddressFamiliesInner)`

SetAddressFamilies sets AddressFamilies field to given value.

### HasAddressFamilies

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasAddressFamilies() bool`

HasAddressFamilies returns a boolean if a field has been set.

### GetAllowAsIn

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetAllowAsIn() int32`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetAllowAsInOk() (*int32, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetAllowAsIn(v int32)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetAsOverride

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetAsOverride() bool`

GetAsOverride returns the AsOverride field if non-nil, zero value otherwise.

### GetAsOverrideOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetAsOverrideOk() (*bool, bool)`

GetAsOverrideOk returns a tuple with the AsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOverride

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetAsOverride(v bool)`

SetAsOverride sets AsOverride field to given value.

### HasAsOverride

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasAsOverride() bool`

HasAsOverride returns a boolean if a field has been set.

### GetBfd

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetBfd() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetBfdOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetBfd(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetBfdNeighbor

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetBfdNeighbor() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor`

GetBfdNeighbor returns the BfdNeighbor field if non-nil, zero value otherwise.

### GetBfdNeighborOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetBfdNeighborOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor, bool)`

GetBfdNeighborOk returns a tuple with the BfdNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdNeighbor

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetBfdNeighbor(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor)`

SetBfdNeighbor sets BfdNeighbor field to given value.

### HasBfdNeighbor

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasBfdNeighbor() bool`

HasBfdNeighbor returns a boolean if a field has been set.

### GetBgpType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetBgpType() string`

GetBgpType returns the BgpType field if non-nil, zero value otherwise.

### GetBgpTypeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetBgpTypeOk() (*string, bool)`

GetBgpTypeOk returns a tuple with the BgpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetBgpType(v string)`

SetBgpType sets BgpType field to given value.

### HasBgpType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasBgpType() bool`

HasBgpType returns a boolean if a field has been set.

### GetDefaultOriginate

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetDefaultOriginate() string`

GetDefaultOriginate returns the DefaultOriginate field if non-nil, zero value otherwise.

### GetDefaultOriginateOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetDefaultOriginateOk() (*string, bool)`

GetDefaultOriginateOk returns a tuple with the DefaultOriginate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOriginate

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetDefaultOriginate(v string)`

SetDefaultOriginate sets DefaultOriginate field to given value.

### HasDefaultOriginate

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasDefaultOriginate() bool`

HasDefaultOriginate returns a boolean if a field has been set.

### GetEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHoldTimer

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetHoldTimer() int32`

GetHoldTimer returns the HoldTimer field if non-nil, zero value otherwise.

### GetHoldTimerOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetHoldTimerOk() (*int32, bool)`

GetHoldTimerOk returns a tuple with the HoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTimer

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetHoldTimer(v int32)`

SetHoldTimer sets HoldTimer field to given value.

### HasHoldTimer

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasHoldTimer() bool`

HasHoldTimer returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeepaliveTimer

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetKeepaliveTimer() int32`

GetKeepaliveTimer returns the KeepaliveTimer field if non-nil, zero value otherwise.

### GetKeepaliveTimerOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetKeepaliveTimerOk() (*int32, bool)`

GetKeepaliveTimerOk returns a tuple with the KeepaliveTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveTimer

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetKeepaliveTimer(v int32)`

SetKeepaliveTimer sets KeepaliveTimer field to given value.

### HasKeepaliveTimer

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasKeepaliveTimer() bool`

HasKeepaliveTimer returns a boolean if a field has been set.

### GetLocalAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetLocalInterface

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetLocalInterface() string`

GetLocalInterface returns the LocalInterface field if non-nil, zero value otherwise.

### GetLocalInterfaceOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetLocalInterfaceOk() (*string, bool)`

GetLocalInterfaceOk returns a tuple with the LocalInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInterface

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetLocalInterface(v string)`

SetLocalInterface sets LocalInterface field to given value.

### HasLocalInterface

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasLocalInterface() bool`

HasLocalInterface returns a boolean if a field has been set.

### GetMaxPrefix

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetMaxPrefix() int32`

GetMaxPrefix returns the MaxPrefix field if non-nil, zero value otherwise.

### GetMaxPrefixOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetMaxPrefixOk() (*int32, bool)`

GetMaxPrefixOk returns a tuple with the MaxPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrefix

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetMaxPrefix(v int32)`

SetMaxPrefix sets MaxPrefix field to given value.

### HasMaxPrefix

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasMaxPrefix() bool`

HasMaxPrefix returns a boolean if a field has been set.

### GetMd5Password

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetMd5Password() string`

GetMd5Password returns the Md5Password field if non-nil, zero value otherwise.

### GetMd5PasswordOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetMd5PasswordOk() (*string, bool)`

GetMd5PasswordOk returns a tuple with the Md5Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Password

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetMd5Password(v string)`

SetMd5Password sets Md5Password field to given value.

### HasMd5Password

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasMd5Password() bool`

HasMd5Password returns a boolean if a field has been set.

### GetMultiHop

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetMultiHop() int32`

GetMultiHop returns the MultiHop field if non-nil, zero value otherwise.

### GetMultiHopOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetMultiHopOk() (*int32, bool)`

GetMultiHopOk returns a tuple with the MultiHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiHop

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetMultiHop(v int32)`

SetMultiHop sets MultiHop field to given value.

### HasMultiHop

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasMultiHop() bool`

HasMultiHop returns a boolean if a field has been set.

### GetPeerAsn

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetRemovePrivateAs

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetRemovePrivateAs() bool`

GetRemovePrivateAs returns the RemovePrivateAs field if non-nil, zero value otherwise.

### GetRemovePrivateAsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetRemovePrivateAsOk() (*bool, bool)`

GetRemovePrivateAsOk returns a tuple with the RemovePrivateAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePrivateAs

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetRemovePrivateAs(v bool)`

SetRemovePrivateAs sets RemovePrivateAs field to given value.

### HasRemovePrivateAs

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasRemovePrivateAs() bool`

HasRemovePrivateAs returns a boolean if a field has been set.

### GetSendCommunity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetSendCommunity() bool`

GetSendCommunity returns the SendCommunity field if non-nil, zero value otherwise.

### GetSendCommunityOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetSendCommunityOk() (*bool, bool)`

GetSendCommunityOk returns a tuple with the SendCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCommunity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetSendCommunity(v bool)`

SetSendCommunity sets SendCommunity field to given value.

### HasSendCommunity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasSendCommunity() bool`

HasSendCommunity returns a boolean if a field has been set.

### GetState

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeSinceLastOperChange

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetTimeSinceLastOperChange() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetTimeSinceLastOperChange returns the TimeSinceLastOperChange field if non-nil, zero value otherwise.

### GetTimeSinceLastOperChangeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetTimeSinceLastOperChangeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetTimeSinceLastOperChangeOk returns a tuple with the TimeSinceLastOperChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSinceLastOperChange

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetTimeSinceLastOperChange(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetTimeSinceLastOperChange sets TimeSinceLastOperChange field to given value.

### HasTimeSinceLastOperChange

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasTimeSinceLastOperChange() bool`

HasTimeSinceLastOperChange returns a boolean if a field has been set.

### GetUp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInner) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


