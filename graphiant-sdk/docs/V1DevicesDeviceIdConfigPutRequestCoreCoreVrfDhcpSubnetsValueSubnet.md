# V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultLeaseTimeSecs** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**DomainNameServer** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInnerNameservers**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInnerNameservers.md) |  | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**IpGateway** | Pointer to **string** |  | [optional] 
**IpPrefix** | Pointer to **string** |  | [optional] 
**IpRanges** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInnerRangesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInnerRangesInner.md) |  | [optional] 
**IpRangesV2** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetIpRangesV2**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetIpRangesV2.md) |  | [optional] 
**MaxLeaseTimeSecs** | Pointer to **int32** |  | [optional] 
**MinLeaseTimeSecs** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StaticLeases** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetStaticLeasesValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetStaticLeasesValue.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet

`func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet`

NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet`

NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultLeaseTimeSecs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetDefaultLeaseTimeSecs() int32`

GetDefaultLeaseTimeSecs returns the DefaultLeaseTimeSecs field if non-nil, zero value otherwise.

### GetDefaultLeaseTimeSecsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetDefaultLeaseTimeSecsOk() (*int32, bool)`

GetDefaultLeaseTimeSecsOk returns a tuple with the DefaultLeaseTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLeaseTimeSecs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetDefaultLeaseTimeSecs(v int32)`

SetDefaultLeaseTimeSecs sets DefaultLeaseTimeSecs field to given value.

### HasDefaultLeaseTimeSecs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasDefaultLeaseTimeSecs() bool`

HasDefaultLeaseTimeSecs returns a boolean if a field has been set.

### GetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomainName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServer

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetDomainNameServer() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInnerNameservers`

GetDomainNameServer returns the DomainNameServer field if non-nil, zero value otherwise.

### GetDomainNameServerOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetDomainNameServerOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInnerNameservers, bool)`

GetDomainNameServerOk returns a tuple with the DomainNameServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServer

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetDomainNameServer(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInnerNameservers)`

SetDomainNameServer sets DomainNameServer field to given value.

### HasDomainNameServer

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasDomainNameServer() bool`

HasDomainNameServer returns a boolean if a field has been set.

### GetInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetIpGateway

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetIpGateway() string`

GetIpGateway returns the IpGateway field if non-nil, zero value otherwise.

### GetIpGatewayOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetIpGatewayOk() (*string, bool)`

GetIpGatewayOk returns a tuple with the IpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpGateway

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetIpGateway(v string)`

SetIpGateway sets IpGateway field to given value.

### HasIpGateway

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasIpGateway() bool`

HasIpGateway returns a boolean if a field has been set.

### GetIpPrefix

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetIpPrefix() string`

GetIpPrefix returns the IpPrefix field if non-nil, zero value otherwise.

### GetIpPrefixOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetIpPrefixOk() (*string, bool)`

GetIpPrefixOk returns a tuple with the IpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPrefix

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetIpPrefix(v string)`

SetIpPrefix sets IpPrefix field to given value.

### HasIpPrefix

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasIpPrefix() bool`

HasIpPrefix returns a boolean if a field has been set.

### GetIpRanges

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetIpRanges() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInnerRangesInner`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetIpRangesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInnerRangesInner, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetIpRanges(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerDhcpSubnetsInnerRangesInner)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetIpRangesV2

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetIpRangesV2() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetIpRangesV2`

GetIpRangesV2 returns the IpRangesV2 field if non-nil, zero value otherwise.

### GetIpRangesV2Ok

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetIpRangesV2Ok() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetIpRangesV2, bool)`

GetIpRangesV2Ok returns a tuple with the IpRangesV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangesV2

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetIpRangesV2(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetIpRangesV2)`

SetIpRangesV2 sets IpRangesV2 field to given value.

### HasIpRangesV2

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasIpRangesV2() bool`

HasIpRangesV2 returns a boolean if a field has been set.

### GetMaxLeaseTimeSecs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetMaxLeaseTimeSecs() int32`

GetMaxLeaseTimeSecs returns the MaxLeaseTimeSecs field if non-nil, zero value otherwise.

### GetMaxLeaseTimeSecsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetMaxLeaseTimeSecsOk() (*int32, bool)`

GetMaxLeaseTimeSecsOk returns a tuple with the MaxLeaseTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLeaseTimeSecs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetMaxLeaseTimeSecs(v int32)`

SetMaxLeaseTimeSecs sets MaxLeaseTimeSecs field to given value.

### HasMaxLeaseTimeSecs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasMaxLeaseTimeSecs() bool`

HasMaxLeaseTimeSecs returns a boolean if a field has been set.

### GetMinLeaseTimeSecs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetMinLeaseTimeSecs() int32`

GetMinLeaseTimeSecs returns the MinLeaseTimeSecs field if non-nil, zero value otherwise.

### GetMinLeaseTimeSecsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetMinLeaseTimeSecsOk() (*int32, bool)`

GetMinLeaseTimeSecsOk returns a tuple with the MinLeaseTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLeaseTimeSecs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetMinLeaseTimeSecs(v int32)`

SetMinLeaseTimeSecs sets MinLeaseTimeSecs field to given value.

### HasMinLeaseTimeSecs

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasMinLeaseTimeSecs() bool`

HasMinLeaseTimeSecs returns a boolean if a field has been set.

### GetName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStaticLeases

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetStaticLeases() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetStaticLeasesValue`

GetStaticLeases returns the StaticLeases field if non-nil, zero value otherwise.

### GetStaticLeasesOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) GetStaticLeasesOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetStaticLeasesValue, bool)`

GetStaticLeasesOk returns a tuple with the StaticLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticLeases

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) SetStaticLeases(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnetStaticLeasesValue)`

SetStaticLeases sets StaticLeases field to given value.

### HasStaticLeases

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfDhcpSubnetsValueSubnet) HasStaticLeases() bool`

HasStaticLeases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


