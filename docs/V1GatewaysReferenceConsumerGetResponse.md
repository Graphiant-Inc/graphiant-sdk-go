# V1GatewaysReferenceConsumerGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpsecGatewayDetails** | Pointer to [**ManaV2IPsecGatewayDetails**](ManaV2IPsecGatewayDetails.md) |  | [optional] 
**IpsecGatewayPeers** | Pointer to [**ManaV2IPsecGatewayPeersConfig**](ManaV2IPsecGatewayPeersConfig.md) |  | [optional] 

## Methods

### NewV1GatewaysReferenceConsumerGetResponse

`func NewV1GatewaysReferenceConsumerGetResponse() *V1GatewaysReferenceConsumerGetResponse`

NewV1GatewaysReferenceConsumerGetResponse instantiates a new V1GatewaysReferenceConsumerGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GatewaysReferenceConsumerGetResponseWithDefaults

`func NewV1GatewaysReferenceConsumerGetResponseWithDefaults() *V1GatewaysReferenceConsumerGetResponse`

NewV1GatewaysReferenceConsumerGetResponseWithDefaults instantiates a new V1GatewaysReferenceConsumerGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpsecGatewayDetails

`func (o *V1GatewaysReferenceConsumerGetResponse) GetIpsecGatewayDetails() ManaV2IPsecGatewayDetails`

GetIpsecGatewayDetails returns the IpsecGatewayDetails field if non-nil, zero value otherwise.

### GetIpsecGatewayDetailsOk

`func (o *V1GatewaysReferenceConsumerGetResponse) GetIpsecGatewayDetailsOk() (*ManaV2IPsecGatewayDetails, bool)`

GetIpsecGatewayDetailsOk returns a tuple with the IpsecGatewayDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecGatewayDetails

`func (o *V1GatewaysReferenceConsumerGetResponse) SetIpsecGatewayDetails(v ManaV2IPsecGatewayDetails)`

SetIpsecGatewayDetails sets IpsecGatewayDetails field to given value.

### HasIpsecGatewayDetails

`func (o *V1GatewaysReferenceConsumerGetResponse) HasIpsecGatewayDetails() bool`

HasIpsecGatewayDetails returns a boolean if a field has been set.

### GetIpsecGatewayPeers

`func (o *V1GatewaysReferenceConsumerGetResponse) GetIpsecGatewayPeers() ManaV2IPsecGatewayPeersConfig`

GetIpsecGatewayPeers returns the IpsecGatewayPeers field if non-nil, zero value otherwise.

### GetIpsecGatewayPeersOk

`func (o *V1GatewaysReferenceConsumerGetResponse) GetIpsecGatewayPeersOk() (*ManaV2IPsecGatewayPeersConfig, bool)`

GetIpsecGatewayPeersOk returns a tuple with the IpsecGatewayPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecGatewayPeers

`func (o *V1GatewaysReferenceConsumerGetResponse) SetIpsecGatewayPeers(v ManaV2IPsecGatewayPeersConfig)`

SetIpsecGatewayPeers sets IpsecGatewayPeers field to given value.

### HasIpsecGatewayPeers

`func (o *V1GatewaysReferenceConsumerGetResponse) HasIpsecGatewayPeers() bool`

HasIpsecGatewayPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


