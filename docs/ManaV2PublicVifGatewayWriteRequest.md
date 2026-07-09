# ManaV2PublicVifGatewayWriteRequest

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

### NewManaV2PublicVifGatewayWriteRequest

`func NewManaV2PublicVifGatewayWriteRequest(consumerLanSegments map[string]ManaV2PublicVifGatewayConsumerLanDevices, gatewayBgpNeighbors map[string]ManaV2BgpNeighborConfig, lanSegmentId int64, natPrefixStrategy ManaV2PublicVifGatewayNatPrefixStrategy, regionId int32, serviceName string, storageProvider string, ) *ManaV2PublicVifGatewayWriteRequest`

NewManaV2PublicVifGatewayWriteRequest instantiates a new ManaV2PublicVifGatewayWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PublicVifGatewayWriteRequestWithDefaults

`func NewManaV2PublicVifGatewayWriteRequestWithDefaults() *ManaV2PublicVifGatewayWriteRequest`

NewManaV2PublicVifGatewayWriteRequestWithDefaults instantiates a new ManaV2PublicVifGatewayWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertisement

`func (o *ManaV2PublicVifGatewayWriteRequest) GetAdvertisement() ManaV2SiteInformation`

GetAdvertisement returns the Advertisement field if non-nil, zero value otherwise.

### GetAdvertisementOk

`func (o *ManaV2PublicVifGatewayWriteRequest) GetAdvertisementOk() (*ManaV2SiteInformation, bool)`

GetAdvertisementOk returns a tuple with the Advertisement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisement

`func (o *ManaV2PublicVifGatewayWriteRequest) SetAdvertisement(v ManaV2SiteInformation)`

SetAdvertisement sets Advertisement field to given value.

### HasAdvertisement

`func (o *ManaV2PublicVifGatewayWriteRequest) HasAdvertisement() bool`

HasAdvertisement returns a boolean if a field has been set.

### GetConsumerLanSegments

`func (o *ManaV2PublicVifGatewayWriteRequest) GetConsumerLanSegments() map[string]ManaV2PublicVifGatewayConsumerLanDevices`

GetConsumerLanSegments returns the ConsumerLanSegments field if non-nil, zero value otherwise.

### GetConsumerLanSegmentsOk

`func (o *ManaV2PublicVifGatewayWriteRequest) GetConsumerLanSegmentsOk() (*map[string]ManaV2PublicVifGatewayConsumerLanDevices, bool)`

GetConsumerLanSegmentsOk returns a tuple with the ConsumerLanSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerLanSegments

`func (o *ManaV2PublicVifGatewayWriteRequest) SetConsumerLanSegments(v map[string]ManaV2PublicVifGatewayConsumerLanDevices)`

SetConsumerLanSegments sets ConsumerLanSegments field to given value.


### GetCoveringPrefixes

`func (o *ManaV2PublicVifGatewayWriteRequest) GetCoveringPrefixes() []string`

GetCoveringPrefixes returns the CoveringPrefixes field if non-nil, zero value otherwise.

### GetCoveringPrefixesOk

`func (o *ManaV2PublicVifGatewayWriteRequest) GetCoveringPrefixesOk() (*[]string, bool)`

GetCoveringPrefixesOk returns a tuple with the CoveringPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveringPrefixes

`func (o *ManaV2PublicVifGatewayWriteRequest) SetCoveringPrefixes(v []string)`

SetCoveringPrefixes sets CoveringPrefixes field to given value.

### HasCoveringPrefixes

`func (o *ManaV2PublicVifGatewayWriteRequest) HasCoveringPrefixes() bool`

HasCoveringPrefixes returns a boolean if a field has been set.

### GetGatewayBgpNeighbors

`func (o *ManaV2PublicVifGatewayWriteRequest) GetGatewayBgpNeighbors() map[string]ManaV2BgpNeighborConfig`

GetGatewayBgpNeighbors returns the GatewayBgpNeighbors field if non-nil, zero value otherwise.

### GetGatewayBgpNeighborsOk

`func (o *ManaV2PublicVifGatewayWriteRequest) GetGatewayBgpNeighborsOk() (*map[string]ManaV2BgpNeighborConfig, bool)`

GetGatewayBgpNeighborsOk returns a tuple with the GatewayBgpNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayBgpNeighbors

`func (o *ManaV2PublicVifGatewayWriteRequest) SetGatewayBgpNeighbors(v map[string]ManaV2BgpNeighborConfig)`

SetGatewayBgpNeighbors sets GatewayBgpNeighbors field to given value.


### GetLanSegmentId

`func (o *ManaV2PublicVifGatewayWriteRequest) GetLanSegmentId() int64`

GetLanSegmentId returns the LanSegmentId field if non-nil, zero value otherwise.

### GetLanSegmentIdOk

`func (o *ManaV2PublicVifGatewayWriteRequest) GetLanSegmentIdOk() (*int64, bool)`

GetLanSegmentIdOk returns a tuple with the LanSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegmentId

`func (o *ManaV2PublicVifGatewayWriteRequest) SetLanSegmentId(v int64)`

SetLanSegmentId sets LanSegmentId field to given value.


### GetNatPrefixStrategy

`func (o *ManaV2PublicVifGatewayWriteRequest) GetNatPrefixStrategy() ManaV2PublicVifGatewayNatPrefixStrategy`

GetNatPrefixStrategy returns the NatPrefixStrategy field if non-nil, zero value otherwise.

### GetNatPrefixStrategyOk

`func (o *ManaV2PublicVifGatewayWriteRequest) GetNatPrefixStrategyOk() (*ManaV2PublicVifGatewayNatPrefixStrategy, bool)`

GetNatPrefixStrategyOk returns a tuple with the NatPrefixStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPrefixStrategy

`func (o *ManaV2PublicVifGatewayWriteRequest) SetNatPrefixStrategy(v ManaV2PublicVifGatewayNatPrefixStrategy)`

SetNatPrefixStrategy sets NatPrefixStrategy field to given value.


### GetRegionId

`func (o *ManaV2PublicVifGatewayWriteRequest) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ManaV2PublicVifGatewayWriteRequest) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ManaV2PublicVifGatewayWriteRequest) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.


### GetServiceName

`func (o *ManaV2PublicVifGatewayWriteRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ManaV2PublicVifGatewayWriteRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ManaV2PublicVifGatewayWriteRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStorageProvider

`func (o *ManaV2PublicVifGatewayWriteRequest) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ManaV2PublicVifGatewayWriteRequest) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ManaV2PublicVifGatewayWriteRequest) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


