# V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpConnection** | Pointer to [**V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerBgpConnection**](V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerBgpConnection.md) |  | [optional] 
**ConnectionQuality** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**Gdi** | Pointer to **int32** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**IpsecConnection** | Pointer to [**V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection**](V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**WanAddresses** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner

`func NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner() *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner`

NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner instantiates a new V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerWithDefaults

`func NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerWithDefaults() *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner`

NewV1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerWithDefaults instantiates a new V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpConnection

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetBgpConnection() V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerBgpConnection`

GetBgpConnection returns the BgpConnection field if non-nil, zero value otherwise.

### GetBgpConnectionOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetBgpConnectionOk() (*V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerBgpConnection, bool)`

GetBgpConnectionOk returns a tuple with the BgpConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConnection

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) SetBgpConnection(v V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerBgpConnection)`

SetBgpConnection sets BgpConnection field to given value.

### HasBgpConnection

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) HasBgpConnection() bool`

HasBgpConnection returns a boolean if a field has been set.

### GetConnectionQuality

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetConnectionQuality() string`

GetConnectionQuality returns the ConnectionQuality field if non-nil, zero value otherwise.

### GetConnectionQualityOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetConnectionQualityOk() (*string, bool)`

GetConnectionQualityOk returns a tuple with the ConnectionQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionQuality

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) SetConnectionQuality(v string)`

SetConnectionQuality sets ConnectionQuality field to given value.

### HasConnectionQuality

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) HasConnectionQuality() bool`

HasConnectionQuality returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetGdi

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetGdi() int32`

GetGdi returns the Gdi field if non-nil, zero value otherwise.

### GetGdiOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetGdiOk() (*int32, bool)`

GetGdiOk returns a tuple with the Gdi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdi

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) SetGdi(v int32)`

SetGdi sets Gdi field to given value.

### HasGdi

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) HasGdi() bool`

HasGdi returns a boolean if a field has been set.

### GetHostname

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpsecConnection

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetIpsecConnection() V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection`

GetIpsecConnection returns the IpsecConnection field if non-nil, zero value otherwise.

### GetIpsecConnectionOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetIpsecConnectionOk() (*V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection, bool)`

GetIpsecConnectionOk returns a tuple with the IpsecConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecConnection

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) SetIpsecConnection(v V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInnerIpsecConnection)`

SetIpsecConnection sets IpsecConnection field to given value.

### HasIpsecConnection

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) HasIpsecConnection() bool`

HasIpsecConnection returns a boolean if a field has been set.

### GetState

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWanAddresses

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetWanAddresses() []string`

GetWanAddresses returns the WanAddresses field if non-nil, zero value otherwise.

### GetWanAddressesOk

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) GetWanAddressesOk() (*[]string, bool)`

GetWanAddressesOk returns a tuple with the WanAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanAddresses

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) SetWanAddresses(v []string)`

SetWanAddresses sets WanAddresses field to given value.

### HasWanAddresses

`func (o *V1DevicesDeviceIdSlicePeersGet200ResponseSlicesInnerPeersInner) HasWanAddresses() bool`

HasWanAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


