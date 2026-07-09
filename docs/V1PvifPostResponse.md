# V1PvifPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiseAllSites** | Pointer to **bool** | True when consumer scope is all symmetric sites (snapshotted on create and update); false when advertisement lists explicit sites or site lists | [optional] 
**Advertisement** | Pointer to [**ManaV2SiteInformation**](ManaV2SiteInformation.md) |  | [optional] 
**ConsumerLanSegments** | Pointer to [**map[string]ManaV2PublicVifGatewayConsumerLanDevices**](ManaV2PublicVifGatewayConsumerLanDevices.md) |  | [optional] 
**CoveringPrefixes** | Pointer to **[]string** |  | [optional] 
**GatewayBgpNeighbors** | Pointer to [**map[string]ManaV2BgpNeighbor**](ManaV2BgpNeighbor.md) |  | [optional] 
**Id** | Pointer to **int64** | Producer service id | [optional] 
**LanSegmentId** | Pointer to **int64** | Producer LAN segment (VRF) id | [optional] 
**NatPrefixStrategy** | Pointer to [**ManaV2PublicVifGatewayNatPrefixStrategy**](ManaV2PublicVifGatewayNatPrefixStrategy.md) |  | [optional] 
**RegionId** | Pointer to **int32** | Graphiant region for gateway appliances | [optional] 
**ServiceName** | Pointer to **string** | Service display name | [optional] 
**StorageProvider** | Pointer to **string** | Storage provider for gateway appliances | [optional] 

## Methods

### NewV1PvifPostResponse

`func NewV1PvifPostResponse() *V1PvifPostResponse`

NewV1PvifPostResponse instantiates a new V1PvifPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PvifPostResponseWithDefaults

`func NewV1PvifPostResponseWithDefaults() *V1PvifPostResponse`

NewV1PvifPostResponseWithDefaults instantiates a new V1PvifPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertiseAllSites

`func (o *V1PvifPostResponse) GetAdvertiseAllSites() bool`

GetAdvertiseAllSites returns the AdvertiseAllSites field if non-nil, zero value otherwise.

### GetAdvertiseAllSitesOk

`func (o *V1PvifPostResponse) GetAdvertiseAllSitesOk() (*bool, bool)`

GetAdvertiseAllSitesOk returns a tuple with the AdvertiseAllSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseAllSites

`func (o *V1PvifPostResponse) SetAdvertiseAllSites(v bool)`

SetAdvertiseAllSites sets AdvertiseAllSites field to given value.

### HasAdvertiseAllSites

`func (o *V1PvifPostResponse) HasAdvertiseAllSites() bool`

HasAdvertiseAllSites returns a boolean if a field has been set.

### GetAdvertisement

`func (o *V1PvifPostResponse) GetAdvertisement() ManaV2SiteInformation`

GetAdvertisement returns the Advertisement field if non-nil, zero value otherwise.

### GetAdvertisementOk

`func (o *V1PvifPostResponse) GetAdvertisementOk() (*ManaV2SiteInformation, bool)`

GetAdvertisementOk returns a tuple with the Advertisement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisement

`func (o *V1PvifPostResponse) SetAdvertisement(v ManaV2SiteInformation)`

SetAdvertisement sets Advertisement field to given value.

### HasAdvertisement

`func (o *V1PvifPostResponse) HasAdvertisement() bool`

HasAdvertisement returns a boolean if a field has been set.

### GetConsumerLanSegments

`func (o *V1PvifPostResponse) GetConsumerLanSegments() map[string]ManaV2PublicVifGatewayConsumerLanDevices`

GetConsumerLanSegments returns the ConsumerLanSegments field if non-nil, zero value otherwise.

### GetConsumerLanSegmentsOk

`func (o *V1PvifPostResponse) GetConsumerLanSegmentsOk() (*map[string]ManaV2PublicVifGatewayConsumerLanDevices, bool)`

GetConsumerLanSegmentsOk returns a tuple with the ConsumerLanSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerLanSegments

`func (o *V1PvifPostResponse) SetConsumerLanSegments(v map[string]ManaV2PublicVifGatewayConsumerLanDevices)`

SetConsumerLanSegments sets ConsumerLanSegments field to given value.

### HasConsumerLanSegments

`func (o *V1PvifPostResponse) HasConsumerLanSegments() bool`

HasConsumerLanSegments returns a boolean if a field has been set.

### GetCoveringPrefixes

`func (o *V1PvifPostResponse) GetCoveringPrefixes() []string`

GetCoveringPrefixes returns the CoveringPrefixes field if non-nil, zero value otherwise.

### GetCoveringPrefixesOk

`func (o *V1PvifPostResponse) GetCoveringPrefixesOk() (*[]string, bool)`

GetCoveringPrefixesOk returns a tuple with the CoveringPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveringPrefixes

`func (o *V1PvifPostResponse) SetCoveringPrefixes(v []string)`

SetCoveringPrefixes sets CoveringPrefixes field to given value.

### HasCoveringPrefixes

`func (o *V1PvifPostResponse) HasCoveringPrefixes() bool`

HasCoveringPrefixes returns a boolean if a field has been set.

### GetGatewayBgpNeighbors

`func (o *V1PvifPostResponse) GetGatewayBgpNeighbors() map[string]ManaV2BgpNeighbor`

GetGatewayBgpNeighbors returns the GatewayBgpNeighbors field if non-nil, zero value otherwise.

### GetGatewayBgpNeighborsOk

`func (o *V1PvifPostResponse) GetGatewayBgpNeighborsOk() (*map[string]ManaV2BgpNeighbor, bool)`

GetGatewayBgpNeighborsOk returns a tuple with the GatewayBgpNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayBgpNeighbors

`func (o *V1PvifPostResponse) SetGatewayBgpNeighbors(v map[string]ManaV2BgpNeighbor)`

SetGatewayBgpNeighbors sets GatewayBgpNeighbors field to given value.

### HasGatewayBgpNeighbors

`func (o *V1PvifPostResponse) HasGatewayBgpNeighbors() bool`

HasGatewayBgpNeighbors returns a boolean if a field has been set.

### GetId

`func (o *V1PvifPostResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1PvifPostResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1PvifPostResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1PvifPostResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanSegmentId

`func (o *V1PvifPostResponse) GetLanSegmentId() int64`

GetLanSegmentId returns the LanSegmentId field if non-nil, zero value otherwise.

### GetLanSegmentIdOk

`func (o *V1PvifPostResponse) GetLanSegmentIdOk() (*int64, bool)`

GetLanSegmentIdOk returns a tuple with the LanSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegmentId

`func (o *V1PvifPostResponse) SetLanSegmentId(v int64)`

SetLanSegmentId sets LanSegmentId field to given value.

### HasLanSegmentId

`func (o *V1PvifPostResponse) HasLanSegmentId() bool`

HasLanSegmentId returns a boolean if a field has been set.

### GetNatPrefixStrategy

`func (o *V1PvifPostResponse) GetNatPrefixStrategy() ManaV2PublicVifGatewayNatPrefixStrategy`

GetNatPrefixStrategy returns the NatPrefixStrategy field if non-nil, zero value otherwise.

### GetNatPrefixStrategyOk

`func (o *V1PvifPostResponse) GetNatPrefixStrategyOk() (*ManaV2PublicVifGatewayNatPrefixStrategy, bool)`

GetNatPrefixStrategyOk returns a tuple with the NatPrefixStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPrefixStrategy

`func (o *V1PvifPostResponse) SetNatPrefixStrategy(v ManaV2PublicVifGatewayNatPrefixStrategy)`

SetNatPrefixStrategy sets NatPrefixStrategy field to given value.

### HasNatPrefixStrategy

`func (o *V1PvifPostResponse) HasNatPrefixStrategy() bool`

HasNatPrefixStrategy returns a boolean if a field has been set.

### GetRegionId

`func (o *V1PvifPostResponse) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *V1PvifPostResponse) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *V1PvifPostResponse) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *V1PvifPostResponse) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetServiceName

`func (o *V1PvifPostResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1PvifPostResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1PvifPostResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *V1PvifPostResponse) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStorageProvider

`func (o *V1PvifPostResponse) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *V1PvifPostResponse) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *V1PvifPostResponse) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *V1PvifPostResponse) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


