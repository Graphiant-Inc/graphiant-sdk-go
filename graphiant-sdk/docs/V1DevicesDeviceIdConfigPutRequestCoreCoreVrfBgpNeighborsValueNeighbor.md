# V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamilies** | Pointer to [**map[string]V1GatewaysPutRequestDetailsIpsecGatewayRoutingBgpAddressFamiliesValue**](V1GatewaysPutRequestDetailsIpsecGatewayRoutingBgpAddressFamiliesValue.md) |  | [optional] 
**AllowAsIn** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborAllowAsIn**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborAllowAsIn.md) |  | [optional] 
**AsOverride** | Pointer to **bool** |  | [optional] 
**Bfd** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborBfd**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborBfd.md) |  | [optional] 
**DefaultOriginate** | Pointer to **string** |  | [optional] 
**EbgpMultihopTtl** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborEbgpMultihopTtl**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborEbgpMultihopTtl.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HoldTimer** | Pointer to **int32** |  | [optional] 
**HoldTimerValue** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborHoldTimerValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborHoldTimerValue.md) |  | [optional] 
**KeepaliveTimer** | Pointer to **int32** |  | [optional] 
**KeepaliveTimerValue** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborHoldTimerValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborHoldTimerValue.md) |  | [optional] 
**LocalAddress** | Pointer to **string** |  | [optional] 
**LocalInterface** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface.md) |  | [optional] 
**MaxPrefixValue** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborMaxPrefixValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborMaxPrefixValue.md) |  | [optional] 
**Md5Password** | Pointer to [**V1GatewaysPutRequestDetailsIpsecGatewayRoutingBgpMd5Password**](V1GatewaysPutRequestDetailsIpsecGatewayRoutingBgpMd5Password.md) |  | [optional] 
**PeerAsn** | Pointer to **int32** |  | [optional] 
**RemoteAddress** | Pointer to **string** |  | [optional] 
**RemovePrivateAs** | Pointer to **bool** |  | [optional] 
**SendCommunity** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor

`func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor`

NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor`

NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamilies

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetAddressFamilies() map[string]V1GatewaysPutRequestDetailsIpsecGatewayRoutingBgpAddressFamiliesValue`

GetAddressFamilies returns the AddressFamilies field if non-nil, zero value otherwise.

### GetAddressFamiliesOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetAddressFamiliesOk() (*map[string]V1GatewaysPutRequestDetailsIpsecGatewayRoutingBgpAddressFamiliesValue, bool)`

GetAddressFamiliesOk returns a tuple with the AddressFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamilies

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetAddressFamilies(v map[string]V1GatewaysPutRequestDetailsIpsecGatewayRoutingBgpAddressFamiliesValue)`

SetAddressFamilies sets AddressFamilies field to given value.

### HasAddressFamilies

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasAddressFamilies() bool`

HasAddressFamilies returns a boolean if a field has been set.

### GetAllowAsIn

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetAllowAsIn() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborAllowAsIn`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetAllowAsInOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborAllowAsIn, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetAllowAsIn(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborAllowAsIn)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetAsOverride

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetAsOverride() bool`

GetAsOverride returns the AsOverride field if non-nil, zero value otherwise.

### GetAsOverrideOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetAsOverrideOk() (*bool, bool)`

GetAsOverrideOk returns a tuple with the AsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOverride

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetAsOverride(v bool)`

SetAsOverride sets AsOverride field to given value.

### HasAsOverride

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasAsOverride() bool`

HasAsOverride returns a boolean if a field has been set.

### GetBfd

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetBfd() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborBfd`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetBfdOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborBfd, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetBfd(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborBfd)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetDefaultOriginate

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetDefaultOriginate() string`

GetDefaultOriginate returns the DefaultOriginate field if non-nil, zero value otherwise.

### GetDefaultOriginateOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetDefaultOriginateOk() (*string, bool)`

GetDefaultOriginateOk returns a tuple with the DefaultOriginate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOriginate

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetDefaultOriginate(v string)`

SetDefaultOriginate sets DefaultOriginate field to given value.

### HasDefaultOriginate

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasDefaultOriginate() bool`

HasDefaultOriginate returns a boolean if a field has been set.

### GetEbgpMultihopTtl

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetEbgpMultihopTtl() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborEbgpMultihopTtl`

GetEbgpMultihopTtl returns the EbgpMultihopTtl field if non-nil, zero value otherwise.

### GetEbgpMultihopTtlOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetEbgpMultihopTtlOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborEbgpMultihopTtl, bool)`

GetEbgpMultihopTtlOk returns a tuple with the EbgpMultihopTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpMultihopTtl

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetEbgpMultihopTtl(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborEbgpMultihopTtl)`

SetEbgpMultihopTtl sets EbgpMultihopTtl field to given value.

### HasEbgpMultihopTtl

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasEbgpMultihopTtl() bool`

HasEbgpMultihopTtl returns a boolean if a field has been set.

### GetEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHoldTimer

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetHoldTimer() int32`

GetHoldTimer returns the HoldTimer field if non-nil, zero value otherwise.

### GetHoldTimerOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetHoldTimerOk() (*int32, bool)`

GetHoldTimerOk returns a tuple with the HoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTimer

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetHoldTimer(v int32)`

SetHoldTimer sets HoldTimer field to given value.

### HasHoldTimer

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasHoldTimer() bool`

HasHoldTimer returns a boolean if a field has been set.

### GetHoldTimerValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetHoldTimerValue() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborHoldTimerValue`

GetHoldTimerValue returns the HoldTimerValue field if non-nil, zero value otherwise.

### GetHoldTimerValueOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetHoldTimerValueOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborHoldTimerValue, bool)`

GetHoldTimerValueOk returns a tuple with the HoldTimerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTimerValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetHoldTimerValue(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborHoldTimerValue)`

SetHoldTimerValue sets HoldTimerValue field to given value.

### HasHoldTimerValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasHoldTimerValue() bool`

HasHoldTimerValue returns a boolean if a field has been set.

### GetKeepaliveTimer

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetKeepaliveTimer() int32`

GetKeepaliveTimer returns the KeepaliveTimer field if non-nil, zero value otherwise.

### GetKeepaliveTimerOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetKeepaliveTimerOk() (*int32, bool)`

GetKeepaliveTimerOk returns a tuple with the KeepaliveTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveTimer

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetKeepaliveTimer(v int32)`

SetKeepaliveTimer sets KeepaliveTimer field to given value.

### HasKeepaliveTimer

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasKeepaliveTimer() bool`

HasKeepaliveTimer returns a boolean if a field has been set.

### GetKeepaliveTimerValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetKeepaliveTimerValue() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborHoldTimerValue`

GetKeepaliveTimerValue returns the KeepaliveTimerValue field if non-nil, zero value otherwise.

### GetKeepaliveTimerValueOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetKeepaliveTimerValueOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborHoldTimerValue, bool)`

GetKeepaliveTimerValueOk returns a tuple with the KeepaliveTimerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveTimerValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetKeepaliveTimerValue(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborHoldTimerValue)`

SetKeepaliveTimerValue sets KeepaliveTimerValue field to given value.

### HasKeepaliveTimerValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasKeepaliveTimerValue() bool`

HasKeepaliveTimerValue returns a boolean if a field has been set.

### GetLocalAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetLocalInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetLocalInterface() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface`

GetLocalInterface returns the LocalInterface field if non-nil, zero value otherwise.

### GetLocalInterfaceOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetLocalInterfaceOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface, bool)`

GetLocalInterfaceOk returns a tuple with the LocalInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetLocalInterface(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface)`

SetLocalInterface sets LocalInterface field to given value.

### HasLocalInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasLocalInterface() bool`

HasLocalInterface returns a boolean if a field has been set.

### GetMaxPrefixValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetMaxPrefixValue() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborMaxPrefixValue`

GetMaxPrefixValue returns the MaxPrefixValue field if non-nil, zero value otherwise.

### GetMaxPrefixValueOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetMaxPrefixValueOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborMaxPrefixValue, bool)`

GetMaxPrefixValueOk returns a tuple with the MaxPrefixValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrefixValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetMaxPrefixValue(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighborMaxPrefixValue)`

SetMaxPrefixValue sets MaxPrefixValue field to given value.

### HasMaxPrefixValue

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasMaxPrefixValue() bool`

HasMaxPrefixValue returns a boolean if a field has been set.

### GetMd5Password

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetMd5Password() V1GatewaysPutRequestDetailsIpsecGatewayRoutingBgpMd5Password`

GetMd5Password returns the Md5Password field if non-nil, zero value otherwise.

### GetMd5PasswordOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetMd5PasswordOk() (*V1GatewaysPutRequestDetailsIpsecGatewayRoutingBgpMd5Password, bool)`

GetMd5PasswordOk returns a tuple with the Md5Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Password

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetMd5Password(v V1GatewaysPutRequestDetailsIpsecGatewayRoutingBgpMd5Password)`

SetMd5Password sets Md5Password field to given value.

### HasMd5Password

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasMd5Password() bool`

HasMd5Password returns a boolean if a field has been set.

### GetPeerAsn

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetRemovePrivateAs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetRemovePrivateAs() bool`

GetRemovePrivateAs returns the RemovePrivateAs field if non-nil, zero value otherwise.

### GetRemovePrivateAsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetRemovePrivateAsOk() (*bool, bool)`

GetRemovePrivateAsOk returns a tuple with the RemovePrivateAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePrivateAs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetRemovePrivateAs(v bool)`

SetRemovePrivateAs sets RemovePrivateAs field to given value.

### HasRemovePrivateAs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasRemovePrivateAs() bool`

HasRemovePrivateAs returns a boolean if a field has been set.

### GetSendCommunity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetSendCommunity() bool`

GetSendCommunity returns the SendCommunity field if non-nil, zero value otherwise.

### GetSendCommunityOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) GetSendCommunityOk() (*bool, bool)`

GetSendCommunityOk returns a tuple with the SendCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCommunity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) SetSendCommunity(v bool)`

SetSendCommunity sets SendCommunity field to given value.

### HasSendCommunity

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfBgpNeighborsValueNeighbor) HasSendCommunity() bool`

HasSendCommunity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


