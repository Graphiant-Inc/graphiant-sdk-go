# V1GatewaysPutRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**V1GatewaysPutRequestDetailsAws**](V1GatewaysPutRequestDetailsAws.md) |  | [optional] 
**Azure** | Pointer to [**V1GatewaysPutRequestDetailsAzure**](V1GatewaysPutRequestDetailsAzure.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Gcp** | Pointer to [**V1GatewaysPutRequestDetailsGcp**](V1GatewaysPutRequestDetailsGcp.md) |  | [optional] 
**IpsecGateway** | Pointer to [**V1GatewaysPutRequestDetailsIpsecGateway**](V1GatewaysPutRequestDetailsIpsecGateway.md) |  | [optional] 
**Oci** | Pointer to [**V1GatewaysPutRequestDetailsOci**](V1GatewaysPutRequestDetailsOci.md) |  | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] 
**Speed** | Pointer to **string** |  | [optional] 
**VrfId** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1GatewaysPutRequestDetails

`func NewV1GatewaysPutRequestDetails() *V1GatewaysPutRequestDetails`

NewV1GatewaysPutRequestDetails instantiates a new V1GatewaysPutRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GatewaysPutRequestDetailsWithDefaults

`func NewV1GatewaysPutRequestDetailsWithDefaults() *V1GatewaysPutRequestDetails`

NewV1GatewaysPutRequestDetailsWithDefaults instantiates a new V1GatewaysPutRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *V1GatewaysPutRequestDetails) GetAws() V1GatewaysPutRequestDetailsAws`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *V1GatewaysPutRequestDetails) GetAwsOk() (*V1GatewaysPutRequestDetailsAws, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *V1GatewaysPutRequestDetails) SetAws(v V1GatewaysPutRequestDetailsAws)`

SetAws sets Aws field to given value.

### HasAws

`func (o *V1GatewaysPutRequestDetails) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *V1GatewaysPutRequestDetails) GetAzure() V1GatewaysPutRequestDetailsAzure`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *V1GatewaysPutRequestDetails) GetAzureOk() (*V1GatewaysPutRequestDetailsAzure, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *V1GatewaysPutRequestDetails) SetAzure(v V1GatewaysPutRequestDetailsAzure)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *V1GatewaysPutRequestDetails) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetDescription

`func (o *V1GatewaysPutRequestDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1GatewaysPutRequestDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1GatewaysPutRequestDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1GatewaysPutRequestDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGcp

`func (o *V1GatewaysPutRequestDetails) GetGcp() V1GatewaysPutRequestDetailsGcp`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *V1GatewaysPutRequestDetails) GetGcpOk() (*V1GatewaysPutRequestDetailsGcp, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *V1GatewaysPutRequestDetails) SetGcp(v V1GatewaysPutRequestDetailsGcp)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *V1GatewaysPutRequestDetails) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### GetIpsecGateway

`func (o *V1GatewaysPutRequestDetails) GetIpsecGateway() V1GatewaysPutRequestDetailsIpsecGateway`

GetIpsecGateway returns the IpsecGateway field if non-nil, zero value otherwise.

### GetIpsecGatewayOk

`func (o *V1GatewaysPutRequestDetails) GetIpsecGatewayOk() (*V1GatewaysPutRequestDetailsIpsecGateway, bool)`

GetIpsecGatewayOk returns a tuple with the IpsecGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecGateway

`func (o *V1GatewaysPutRequestDetails) SetIpsecGateway(v V1GatewaysPutRequestDetailsIpsecGateway)`

SetIpsecGateway sets IpsecGateway field to given value.

### HasIpsecGateway

`func (o *V1GatewaysPutRequestDetails) HasIpsecGateway() bool`

HasIpsecGateway returns a boolean if a field has been set.

### GetOci

`func (o *V1GatewaysPutRequestDetails) GetOci() V1GatewaysPutRequestDetailsOci`

GetOci returns the Oci field if non-nil, zero value otherwise.

### GetOciOk

`func (o *V1GatewaysPutRequestDetails) GetOciOk() (*V1GatewaysPutRequestDetailsOci, bool)`

GetOciOk returns a tuple with the Oci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOci

`func (o *V1GatewaysPutRequestDetails) SetOci(v V1GatewaysPutRequestDetailsOci)`

SetOci sets Oci field to given value.

### HasOci

`func (o *V1GatewaysPutRequestDetails) HasOci() bool`

HasOci returns a boolean if a field has been set.

### GetRegionId

`func (o *V1GatewaysPutRequestDetails) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *V1GatewaysPutRequestDetails) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *V1GatewaysPutRequestDetails) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *V1GatewaysPutRequestDetails) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetSpeed

`func (o *V1GatewaysPutRequestDetails) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *V1GatewaysPutRequestDetails) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *V1GatewaysPutRequestDetails) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *V1GatewaysPutRequestDetails) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetVrfId

`func (o *V1GatewaysPutRequestDetails) GetVrfId() int64`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *V1GatewaysPutRequestDetails) GetVrfIdOk() (*int64, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *V1GatewaysPutRequestDetails) SetVrfId(v int64)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *V1GatewaysPutRequestDetails) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


