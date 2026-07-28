# ManaV2GatewayDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**ManaV2AwsGatewayDetails**](ManaV2AwsGatewayDetails.md) |  | [optional] 
**Azure** | Pointer to [**ManaV2AzureGatewayDetails**](ManaV2AzureGatewayDetails.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Gcp** | Pointer to [**ManaV2GcpGatewayDetails**](ManaV2GcpGatewayDetails.md) |  | [optional] 
**IpsecGateway** | Pointer to [**ManaV2IPsecGatewayDetails**](ManaV2IPsecGatewayDetails.md) |  | [optional] 
**IpsecGatewayPeers** | Pointer to [**ManaV2IPsecGatewayPeersConfig**](ManaV2IPsecGatewayPeersConfig.md) |  | [optional] 
**Oci** | Pointer to [**ManaV2OciGatewayDetails**](ManaV2OciGatewayDetails.md) |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**Speed** | Pointer to **string** |  | [optional] 
**VrfId** | Pointer to **int64** |  | [optional] 

## Methods

### NewManaV2GatewayDetails

`func NewManaV2GatewayDetails() *ManaV2GatewayDetails`

NewManaV2GatewayDetails instantiates a new ManaV2GatewayDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2GatewayDetailsWithDefaults

`func NewManaV2GatewayDetailsWithDefaults() *ManaV2GatewayDetails`

NewManaV2GatewayDetailsWithDefaults instantiates a new ManaV2GatewayDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *ManaV2GatewayDetails) GetAws() ManaV2AwsGatewayDetails`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *ManaV2GatewayDetails) GetAwsOk() (*ManaV2AwsGatewayDetails, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *ManaV2GatewayDetails) SetAws(v ManaV2AwsGatewayDetails)`

SetAws sets Aws field to given value.

### HasAws

`func (o *ManaV2GatewayDetails) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *ManaV2GatewayDetails) GetAzure() ManaV2AzureGatewayDetails`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *ManaV2GatewayDetails) GetAzureOk() (*ManaV2AzureGatewayDetails, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *ManaV2GatewayDetails) SetAzure(v ManaV2AzureGatewayDetails)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *ManaV2GatewayDetails) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetDescription

`func (o *ManaV2GatewayDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManaV2GatewayDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManaV2GatewayDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManaV2GatewayDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGcp

`func (o *ManaV2GatewayDetails) GetGcp() ManaV2GcpGatewayDetails`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *ManaV2GatewayDetails) GetGcpOk() (*ManaV2GcpGatewayDetails, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *ManaV2GatewayDetails) SetGcp(v ManaV2GcpGatewayDetails)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *ManaV2GatewayDetails) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### GetIpsecGateway

`func (o *ManaV2GatewayDetails) GetIpsecGateway() ManaV2IPsecGatewayDetails`

GetIpsecGateway returns the IpsecGateway field if non-nil, zero value otherwise.

### GetIpsecGatewayOk

`func (o *ManaV2GatewayDetails) GetIpsecGatewayOk() (*ManaV2IPsecGatewayDetails, bool)`

GetIpsecGatewayOk returns a tuple with the IpsecGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecGateway

`func (o *ManaV2GatewayDetails) SetIpsecGateway(v ManaV2IPsecGatewayDetails)`

SetIpsecGateway sets IpsecGateway field to given value.

### HasIpsecGateway

`func (o *ManaV2GatewayDetails) HasIpsecGateway() bool`

HasIpsecGateway returns a boolean if a field has been set.

### GetIpsecGatewayPeers

`func (o *ManaV2GatewayDetails) GetIpsecGatewayPeers() ManaV2IPsecGatewayPeersConfig`

GetIpsecGatewayPeers returns the IpsecGatewayPeers field if non-nil, zero value otherwise.

### GetIpsecGatewayPeersOk

`func (o *ManaV2GatewayDetails) GetIpsecGatewayPeersOk() (*ManaV2IPsecGatewayPeersConfig, bool)`

GetIpsecGatewayPeersOk returns a tuple with the IpsecGatewayPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecGatewayPeers

`func (o *ManaV2GatewayDetails) SetIpsecGatewayPeers(v ManaV2IPsecGatewayPeersConfig)`

SetIpsecGatewayPeers sets IpsecGatewayPeers field to given value.

### HasIpsecGatewayPeers

`func (o *ManaV2GatewayDetails) HasIpsecGatewayPeers() bool`

HasIpsecGatewayPeers returns a boolean if a field has been set.

### GetOci

`func (o *ManaV2GatewayDetails) GetOci() ManaV2OciGatewayDetails`

GetOci returns the Oci field if non-nil, zero value otherwise.

### GetOciOk

`func (o *ManaV2GatewayDetails) GetOciOk() (*ManaV2OciGatewayDetails, bool)`

GetOciOk returns a tuple with the Oci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOci

`func (o *ManaV2GatewayDetails) SetOci(v ManaV2OciGatewayDetails)`

SetOci sets Oci field to given value.

### HasOci

`func (o *ManaV2GatewayDetails) HasOci() bool`

HasOci returns a boolean if a field has been set.

### GetRegionId

`func (o *ManaV2GatewayDetails) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ManaV2GatewayDetails) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ManaV2GatewayDetails) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ManaV2GatewayDetails) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetSpeed

`func (o *ManaV2GatewayDetails) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ManaV2GatewayDetails) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ManaV2GatewayDetails) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ManaV2GatewayDetails) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetVrfId

`func (o *ManaV2GatewayDetails) GetVrfId() int64`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *ManaV2GatewayDetails) GetVrfIdOk() (*int64, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *ManaV2GatewayDetails) SetVrfId(v int64)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *ManaV2GatewayDetails) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


