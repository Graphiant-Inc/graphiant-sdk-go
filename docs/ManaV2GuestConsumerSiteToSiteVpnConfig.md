# ManaV2GuestConsumerSiteToSiteVpnConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **[]string** |  | [optional] 
**IpsecGatewayDetails** | Pointer to [**ManaV2IPsecGatewayDetails**](ManaV2IPsecGatewayDetails.md) |  | [optional] 
**IpsecGatewayPeers** | Pointer to [**ManaV2IPsecGatewayPeersConfig**](ManaV2IPsecGatewayPeersConfig.md) |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewManaV2GuestConsumerSiteToSiteVpnConfig

`func NewManaV2GuestConsumerSiteToSiteVpnConfig() *ManaV2GuestConsumerSiteToSiteVpnConfig`

NewManaV2GuestConsumerSiteToSiteVpnConfig instantiates a new ManaV2GuestConsumerSiteToSiteVpnConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2GuestConsumerSiteToSiteVpnConfigWithDefaults

`func NewManaV2GuestConsumerSiteToSiteVpnConfigWithDefaults() *ManaV2GuestConsumerSiteToSiteVpnConfig`

NewManaV2GuestConsumerSiteToSiteVpnConfigWithDefaults instantiates a new ManaV2GuestConsumerSiteToSiteVpnConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetIpsecGatewayDetails

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) GetIpsecGatewayDetails() ManaV2IPsecGatewayDetails`

GetIpsecGatewayDetails returns the IpsecGatewayDetails field if non-nil, zero value otherwise.

### GetIpsecGatewayDetailsOk

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) GetIpsecGatewayDetailsOk() (*ManaV2IPsecGatewayDetails, bool)`

GetIpsecGatewayDetailsOk returns a tuple with the IpsecGatewayDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecGatewayDetails

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) SetIpsecGatewayDetails(v ManaV2IPsecGatewayDetails)`

SetIpsecGatewayDetails sets IpsecGatewayDetails field to given value.

### HasIpsecGatewayDetails

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) HasIpsecGatewayDetails() bool`

HasIpsecGatewayDetails returns a boolean if a field has been set.

### GetIpsecGatewayPeers

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) GetIpsecGatewayPeers() ManaV2IPsecGatewayPeersConfig`

GetIpsecGatewayPeers returns the IpsecGatewayPeers field if non-nil, zero value otherwise.

### GetIpsecGatewayPeersOk

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) GetIpsecGatewayPeersOk() (*ManaV2IPsecGatewayPeersConfig, bool)`

GetIpsecGatewayPeersOk returns a tuple with the IpsecGatewayPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecGatewayPeers

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) SetIpsecGatewayPeers(v ManaV2IPsecGatewayPeersConfig)`

SetIpsecGatewayPeers sets IpsecGatewayPeers field to given value.

### HasIpsecGatewayPeers

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) HasIpsecGatewayPeers() bool`

HasIpsecGatewayPeers returns a boolean if a field has been set.

### GetRegionId

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ManaV2GuestConsumerSiteToSiteVpnConfig) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


