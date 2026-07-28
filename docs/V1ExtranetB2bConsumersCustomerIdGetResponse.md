# V1ExtranetB2bConsumersCustomerIdGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalObjectDeviceSummaries** | Pointer to [**map[string]ManaV2GlobalObjectServiceSummaries**](ManaV2GlobalObjectServiceSummaries.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IpsecTunnelConfig** | Pointer to [**[]V1ExtranetB2bConsumersCustomerIdGetResponseIpsecVpnTunnelConfig**](V1ExtranetB2bConsumersCustomerIdGetResponseIpsecVpnTunnelConfig.md) |  | [optional] 
**MatchDetails** | Pointer to [**ManaV2B2bExtranetMatchConsumerDetails**](ManaV2B2bExtranetMatchConsumerDetails.md) |  | [optional] 
**MatchId** | Pointer to **int64** |  | [optional] 
**PeerType** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to [**ManaV2ExtranetServiceConsumerPolicy**](ManaV2ExtranetServiceConsumerPolicy.md) |  | [optional] 
**SiteToSiteVpn** | Pointer to [**ManaV2GuestConsumerSiteToSiteVpnConfig**](ManaV2GuestConsumerSiteToSiteVpnConfig.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ExtranetB2bConsumersCustomerIdGetResponse

`func NewV1ExtranetB2bConsumersCustomerIdGetResponse() *V1ExtranetB2bConsumersCustomerIdGetResponse`

NewV1ExtranetB2bConsumersCustomerIdGetResponse instantiates a new V1ExtranetB2bConsumersCustomerIdGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetB2bConsumersCustomerIdGetResponseWithDefaults

`func NewV1ExtranetB2bConsumersCustomerIdGetResponseWithDefaults() *V1ExtranetB2bConsumersCustomerIdGetResponse`

NewV1ExtranetB2bConsumersCustomerIdGetResponseWithDefaults instantiates a new V1ExtranetB2bConsumersCustomerIdGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalObjectDeviceSummaries

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetGlobalObjectDeviceSummaries() map[string]ManaV2GlobalObjectServiceSummaries`

GetGlobalObjectDeviceSummaries returns the GlobalObjectDeviceSummaries field if non-nil, zero value otherwise.

### GetGlobalObjectDeviceSummariesOk

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetGlobalObjectDeviceSummariesOk() (*map[string]ManaV2GlobalObjectServiceSummaries, bool)`

GetGlobalObjectDeviceSummariesOk returns a tuple with the GlobalObjectDeviceSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalObjectDeviceSummaries

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) SetGlobalObjectDeviceSummaries(v map[string]ManaV2GlobalObjectServiceSummaries)`

SetGlobalObjectDeviceSummaries sets GlobalObjectDeviceSummaries field to given value.

### HasGlobalObjectDeviceSummaries

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) HasGlobalObjectDeviceSummaries() bool`

HasGlobalObjectDeviceSummaries returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpsecTunnelConfig

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetIpsecTunnelConfig() []V1ExtranetB2bConsumersCustomerIdGetResponseIpsecVpnTunnelConfig`

GetIpsecTunnelConfig returns the IpsecTunnelConfig field if non-nil, zero value otherwise.

### GetIpsecTunnelConfigOk

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetIpsecTunnelConfigOk() (*[]V1ExtranetB2bConsumersCustomerIdGetResponseIpsecVpnTunnelConfig, bool)`

GetIpsecTunnelConfigOk returns a tuple with the IpsecTunnelConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecTunnelConfig

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) SetIpsecTunnelConfig(v []V1ExtranetB2bConsumersCustomerIdGetResponseIpsecVpnTunnelConfig)`

SetIpsecTunnelConfig sets IpsecTunnelConfig field to given value.

### HasIpsecTunnelConfig

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) HasIpsecTunnelConfig() bool`

HasIpsecTunnelConfig returns a boolean if a field has been set.

### GetMatchDetails

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetMatchDetails() ManaV2B2bExtranetMatchConsumerDetails`

GetMatchDetails returns the MatchDetails field if non-nil, zero value otherwise.

### GetMatchDetailsOk

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetMatchDetailsOk() (*ManaV2B2bExtranetMatchConsumerDetails, bool)`

GetMatchDetailsOk returns a tuple with the MatchDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchDetails

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) SetMatchDetails(v ManaV2B2bExtranetMatchConsumerDetails)`

SetMatchDetails sets MatchDetails field to given value.

### HasMatchDetails

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) HasMatchDetails() bool`

HasMatchDetails returns a boolean if a field has been set.

### GetMatchId

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetMatchId() int64`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetMatchIdOk() (*int64, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) SetMatchId(v int64)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetPeerType

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetPeerType() string`

GetPeerType returns the PeerType field if non-nil, zero value otherwise.

### GetPeerTypeOk

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetPeerTypeOk() (*string, bool)`

GetPeerTypeOk returns a tuple with the PeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerType

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) SetPeerType(v string)`

SetPeerType sets PeerType field to given value.

### HasPeerType

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) HasPeerType() bool`

HasPeerType returns a boolean if a field has been set.

### GetPolicy

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetPolicy() ManaV2ExtranetServiceConsumerPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetPolicyOk() (*ManaV2ExtranetServiceConsumerPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) SetPolicy(v ManaV2ExtranetServiceConsumerPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSiteToSiteVpn

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetSiteToSiteVpn() ManaV2GuestConsumerSiteToSiteVpnConfig`

GetSiteToSiteVpn returns the SiteToSiteVpn field if non-nil, zero value otherwise.

### GetSiteToSiteVpnOk

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetSiteToSiteVpnOk() (*ManaV2GuestConsumerSiteToSiteVpnConfig, bool)`

GetSiteToSiteVpnOk returns a tuple with the SiteToSiteVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteToSiteVpn

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) SetSiteToSiteVpn(v ManaV2GuestConsumerSiteToSiteVpnConfig)`

SetSiteToSiteVpn sets SiteToSiteVpn field to given value.

### HasSiteToSiteVpn

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) HasSiteToSiteVpn() bool`

HasSiteToSiteVpn returns a boolean if a field has been set.

### GetStatus

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ExtranetB2bConsumersCustomerIdGetResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


