# V1ExtranetsB2bPeeringConsumerMatchIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **int64** |  | [optional] 
**GlobalObjectOps** | Pointer to [**map[string]ManaV2GlobalObjectServiceOps**](ManaV2GlobalObjectServiceOps.md) |  | [optional] 
**Id** | **int64** | ID of the service which is being consumed by the customer (required) | 
**Nat** | [**[]ManaV2B2bNat**](ManaV2B2bNat.md) |  | 
**Policy** | [**[]ManaV2B2bExtranetPeeringServiceConsumerLanSegmentPolicy**](ManaV2B2bExtranetPeeringServiceConsumerLanSegmentPolicy.md) |  | 
**SiteInformation** | [**[]ManaV2B2bSiteInformation**](ManaV2B2bSiteInformation.md) |  | 
**SiteToSiteVpn** | Pointer to [**ManaV2GuestConsumerSiteToSiteVpnConfig**](ManaV2GuestConsumerSiteToSiteVpnConfig.md) |  | [optional] 

## Methods

### NewV1ExtranetsB2bPeeringConsumerMatchIdPostRequest

`func NewV1ExtranetsB2bPeeringConsumerMatchIdPostRequest(id int64, nat []ManaV2B2bNat, policy []ManaV2B2bExtranetPeeringServiceConsumerLanSegmentPolicy, siteInformation []ManaV2B2bSiteInformation, ) *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest`

NewV1ExtranetsB2bPeeringConsumerMatchIdPostRequest instantiates a new V1ExtranetsB2bPeeringConsumerMatchIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsB2bPeeringConsumerMatchIdPostRequestWithDefaults

`func NewV1ExtranetsB2bPeeringConsumerMatchIdPostRequestWithDefaults() *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest`

NewV1ExtranetsB2bPeeringConsumerMatchIdPostRequestWithDefaults instantiates a new V1ExtranetsB2bPeeringConsumerMatchIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetCustomerId() int64`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetCustomerIdOk() (*int64, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) SetCustomerId(v int64)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetGlobalObjectOps

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetGlobalObjectOps() map[string]ManaV2GlobalObjectServiceOps`

GetGlobalObjectOps returns the GlobalObjectOps field if non-nil, zero value otherwise.

### GetGlobalObjectOpsOk

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetGlobalObjectOpsOk() (*map[string]ManaV2GlobalObjectServiceOps, bool)`

GetGlobalObjectOpsOk returns a tuple with the GlobalObjectOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalObjectOps

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) SetGlobalObjectOps(v map[string]ManaV2GlobalObjectServiceOps)`

SetGlobalObjectOps sets GlobalObjectOps field to given value.

### HasGlobalObjectOps

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) HasGlobalObjectOps() bool`

HasGlobalObjectOps returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetNat

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetNat() []ManaV2B2bNat`

GetNat returns the Nat field if non-nil, zero value otherwise.

### GetNatOk

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetNatOk() (*[]ManaV2B2bNat, bool)`

GetNatOk returns a tuple with the Nat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNat

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) SetNat(v []ManaV2B2bNat)`

SetNat sets Nat field to given value.


### GetPolicy

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetPolicy() []ManaV2B2bExtranetPeeringServiceConsumerLanSegmentPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetPolicyOk() (*[]ManaV2B2bExtranetPeeringServiceConsumerLanSegmentPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) SetPolicy(v []ManaV2B2bExtranetPeeringServiceConsumerLanSegmentPolicy)`

SetPolicy sets Policy field to given value.


### GetSiteInformation

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetSiteInformation() []ManaV2B2bSiteInformation`

GetSiteInformation returns the SiteInformation field if non-nil, zero value otherwise.

### GetSiteInformationOk

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetSiteInformationOk() (*[]ManaV2B2bSiteInformation, bool)`

GetSiteInformationOk returns a tuple with the SiteInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteInformation

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) SetSiteInformation(v []ManaV2B2bSiteInformation)`

SetSiteInformation sets SiteInformation field to given value.


### GetSiteToSiteVpn

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetSiteToSiteVpn() ManaV2GuestConsumerSiteToSiteVpnConfig`

GetSiteToSiteVpn returns the SiteToSiteVpn field if non-nil, zero value otherwise.

### GetSiteToSiteVpnOk

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) GetSiteToSiteVpnOk() (*ManaV2GuestConsumerSiteToSiteVpnConfig, bool)`

GetSiteToSiteVpnOk returns a tuple with the SiteToSiteVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteToSiteVpn

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) SetSiteToSiteVpn(v ManaV2GuestConsumerSiteToSiteVpnConfig)`

SetSiteToSiteVpn sets SiteToSiteVpn field to given value.

### HasSiteToSiteVpn

`func (o *V1ExtranetsB2bPeeringConsumerMatchIdPostRequest) HasSiteToSiteVpn() bool`

HasSiteToSiteVpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


