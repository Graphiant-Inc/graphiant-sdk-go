# V1PvifPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advertisement** | Pointer to [**ManaV2SiteInformation**](ManaV2SiteInformation.md) |  | [optional] 
**ConsumerLanSegments** | [**map[string]ManaV2PublicVifGatewayConsumerLanDevices**](ManaV2PublicVifGatewayConsumerLanDevices.md) |  | 
**CoveringPrefixes** | Pointer to **[]string** |  | [optional] 
**GatewayBgpNeighbors** | [**map[string]ManaV2BgpNeighborConfig**](ManaV2BgpNeighborConfig.md) |  | 
**LanSegmentId** | **int64** | Producer LAN segment (VRF) on gateway appliances (required) | 
**NatPrefixStrategy** | [**ManaV2PublicVifGatewayNatPrefixStrategy**](ManaV2PublicVifGatewayNatPrefixStrategy.md) |  | 
**RegionId** | **int32** | Graphiant region for gateway appliances (required) | 
**ServiceName** | **string** | Service display name (required) | 
**StorageProvider** | **string** | Storage provider; each gateway appliance must match region and provider (required) | 

## Methods

### NewV1PvifPostRequest

`func NewV1PvifPostRequest(consumerLanSegments map[string]ManaV2PublicVifGatewayConsumerLanDevices, gatewayBgpNeighbors map[string]ManaV2BgpNeighborConfig, lanSegmentId int64, natPrefixStrategy ManaV2PublicVifGatewayNatPrefixStrategy, regionId int32, serviceName string, storageProvider string, ) *V1PvifPostRequest`

NewV1PvifPostRequest instantiates a new V1PvifPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PvifPostRequestWithDefaults

`func NewV1PvifPostRequestWithDefaults() *V1PvifPostRequest`

NewV1PvifPostRequestWithDefaults instantiates a new V1PvifPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertisement

`func (o *V1PvifPostRequest) GetAdvertisement() ManaV2SiteInformation`

GetAdvertisement returns the Advertisement field if non-nil, zero value otherwise.

### GetAdvertisementOk

`func (o *V1PvifPostRequest) GetAdvertisementOk() (*ManaV2SiteInformation, bool)`

GetAdvertisementOk returns a tuple with the Advertisement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisement

`func (o *V1PvifPostRequest) SetAdvertisement(v ManaV2SiteInformation)`

SetAdvertisement sets Advertisement field to given value.

### HasAdvertisement

`func (o *V1PvifPostRequest) HasAdvertisement() bool`

HasAdvertisement returns a boolean if a field has been set.

### GetConsumerLanSegments

`func (o *V1PvifPostRequest) GetConsumerLanSegments() map[string]ManaV2PublicVifGatewayConsumerLanDevices`

GetConsumerLanSegments returns the ConsumerLanSegments field if non-nil, zero value otherwise.

### GetConsumerLanSegmentsOk

`func (o *V1PvifPostRequest) GetConsumerLanSegmentsOk() (*map[string]ManaV2PublicVifGatewayConsumerLanDevices, bool)`

GetConsumerLanSegmentsOk returns a tuple with the ConsumerLanSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerLanSegments

`func (o *V1PvifPostRequest) SetConsumerLanSegments(v map[string]ManaV2PublicVifGatewayConsumerLanDevices)`

SetConsumerLanSegments sets ConsumerLanSegments field to given value.


### GetCoveringPrefixes

`func (o *V1PvifPostRequest) GetCoveringPrefixes() []string`

GetCoveringPrefixes returns the CoveringPrefixes field if non-nil, zero value otherwise.

### GetCoveringPrefixesOk

`func (o *V1PvifPostRequest) GetCoveringPrefixesOk() (*[]string, bool)`

GetCoveringPrefixesOk returns a tuple with the CoveringPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveringPrefixes

`func (o *V1PvifPostRequest) SetCoveringPrefixes(v []string)`

SetCoveringPrefixes sets CoveringPrefixes field to given value.

### HasCoveringPrefixes

`func (o *V1PvifPostRequest) HasCoveringPrefixes() bool`

HasCoveringPrefixes returns a boolean if a field has been set.

### GetGatewayBgpNeighbors

`func (o *V1PvifPostRequest) GetGatewayBgpNeighbors() map[string]ManaV2BgpNeighborConfig`

GetGatewayBgpNeighbors returns the GatewayBgpNeighbors field if non-nil, zero value otherwise.

### GetGatewayBgpNeighborsOk

`func (o *V1PvifPostRequest) GetGatewayBgpNeighborsOk() (*map[string]ManaV2BgpNeighborConfig, bool)`

GetGatewayBgpNeighborsOk returns a tuple with the GatewayBgpNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayBgpNeighbors

`func (o *V1PvifPostRequest) SetGatewayBgpNeighbors(v map[string]ManaV2BgpNeighborConfig)`

SetGatewayBgpNeighbors sets GatewayBgpNeighbors field to given value.


### GetLanSegmentId

`func (o *V1PvifPostRequest) GetLanSegmentId() int64`

GetLanSegmentId returns the LanSegmentId field if non-nil, zero value otherwise.

### GetLanSegmentIdOk

`func (o *V1PvifPostRequest) GetLanSegmentIdOk() (*int64, bool)`

GetLanSegmentIdOk returns a tuple with the LanSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegmentId

`func (o *V1PvifPostRequest) SetLanSegmentId(v int64)`

SetLanSegmentId sets LanSegmentId field to given value.


### GetNatPrefixStrategy

`func (o *V1PvifPostRequest) GetNatPrefixStrategy() ManaV2PublicVifGatewayNatPrefixStrategy`

GetNatPrefixStrategy returns the NatPrefixStrategy field if non-nil, zero value otherwise.

### GetNatPrefixStrategyOk

`func (o *V1PvifPostRequest) GetNatPrefixStrategyOk() (*ManaV2PublicVifGatewayNatPrefixStrategy, bool)`

GetNatPrefixStrategyOk returns a tuple with the NatPrefixStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPrefixStrategy

`func (o *V1PvifPostRequest) SetNatPrefixStrategy(v ManaV2PublicVifGatewayNatPrefixStrategy)`

SetNatPrefixStrategy sets NatPrefixStrategy field to given value.


### GetRegionId

`func (o *V1PvifPostRequest) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *V1PvifPostRequest) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *V1PvifPostRequest) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.


### GetServiceName

`func (o *V1PvifPostRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1PvifPostRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1PvifPostRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStorageProvider

`func (o *V1PvifPostRequest) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *V1PvifPostRequest) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *V1PvifPostRequest) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


