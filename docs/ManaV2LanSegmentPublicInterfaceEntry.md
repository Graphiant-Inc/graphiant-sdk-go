# ManaV2LanSegmentPublicInterfaceEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | Pointer to **int64** | network.interface id (required) | [optional] 
**Ipv4Addresses** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Device interface name (BGP local_interface uses this) (required) | [optional] 
**StorageProvider** | Pointer to **string** | Interface storage provider (cloud provider for gateway LAN interfaces) | [optional] 

## Methods

### NewManaV2LanSegmentPublicInterfaceEntry

`func NewManaV2LanSegmentPublicInterfaceEntry() *ManaV2LanSegmentPublicInterfaceEntry`

NewManaV2LanSegmentPublicInterfaceEntry instantiates a new ManaV2LanSegmentPublicInterfaceEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2LanSegmentPublicInterfaceEntryWithDefaults

`func NewManaV2LanSegmentPublicInterfaceEntryWithDefaults() *ManaV2LanSegmentPublicInterfaceEntry`

NewManaV2LanSegmentPublicInterfaceEntryWithDefaults instantiates a new ManaV2LanSegmentPublicInterfaceEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *ManaV2LanSegmentPublicInterfaceEntry) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *ManaV2LanSegmentPublicInterfaceEntry) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *ManaV2LanSegmentPublicInterfaceEntry) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *ManaV2LanSegmentPublicInterfaceEntry) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *ManaV2LanSegmentPublicInterfaceEntry) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *ManaV2LanSegmentPublicInterfaceEntry) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *ManaV2LanSegmentPublicInterfaceEntry) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *ManaV2LanSegmentPublicInterfaceEntry) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetName

`func (o *ManaV2LanSegmentPublicInterfaceEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2LanSegmentPublicInterfaceEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2LanSegmentPublicInterfaceEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2LanSegmentPublicInterfaceEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageProvider

`func (o *ManaV2LanSegmentPublicInterfaceEntry) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ManaV2LanSegmentPublicInterfaceEntry) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ManaV2LanSegmentPublicInterfaceEntry) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ManaV2LanSegmentPublicInterfaceEntry) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


