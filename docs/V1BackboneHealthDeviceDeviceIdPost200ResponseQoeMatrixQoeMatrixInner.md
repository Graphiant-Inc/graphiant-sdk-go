# V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Box** | Pointer to [**[]V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInnerBoxInner**](V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInnerBoxInner.md) |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**DeviceRegion** | Pointer to [**V2AssuranceTopologyInventoryPost200ResponseRegionsInner**](V2AssuranceTopologyInventoryPost200ResponseRegionsInner.md) |  | [optional] 
**PeerDeviceId** | Pointer to **int64** |  | [optional] 
**PeerDeviceRegion** | Pointer to [**V2AssuranceTopologyInventoryPost200ResponseRegionsInner**](V2AssuranceTopologyInventoryPost200ResponseRegionsInner.md) |  | [optional] 
**SessionName** | Pointer to **string** |  | [optional] 

## Methods

### NewV1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner

`func NewV1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner() *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner`

NewV1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInnerWithDefaults

`func NewV1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInnerWithDefaults() *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner`

NewV1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInnerWithDefaults instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBox

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetBox() []V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInnerBoxInner`

GetBox returns the Box field if non-nil, zero value otherwise.

### GetBoxOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetBoxOk() (*[]V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInnerBoxInner, bool)`

GetBoxOk returns a tuple with the Box field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBox

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) SetBox(v []V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInnerBoxInner)`

SetBox sets Box field to given value.

### HasBox

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) HasBox() bool`

HasBox returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceRegion

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetDeviceRegion() V2AssuranceTopologyInventoryPost200ResponseRegionsInner`

GetDeviceRegion returns the DeviceRegion field if non-nil, zero value otherwise.

### GetDeviceRegionOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetDeviceRegionOk() (*V2AssuranceTopologyInventoryPost200ResponseRegionsInner, bool)`

GetDeviceRegionOk returns a tuple with the DeviceRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegion

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) SetDeviceRegion(v V2AssuranceTopologyInventoryPost200ResponseRegionsInner)`

SetDeviceRegion sets DeviceRegion field to given value.

### HasDeviceRegion

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) HasDeviceRegion() bool`

HasDeviceRegion returns a boolean if a field has been set.

### GetPeerDeviceId

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetPeerDeviceId() int64`

GetPeerDeviceId returns the PeerDeviceId field if non-nil, zero value otherwise.

### GetPeerDeviceIdOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetPeerDeviceIdOk() (*int64, bool)`

GetPeerDeviceIdOk returns a tuple with the PeerDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDeviceId

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) SetPeerDeviceId(v int64)`

SetPeerDeviceId sets PeerDeviceId field to given value.

### HasPeerDeviceId

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) HasPeerDeviceId() bool`

HasPeerDeviceId returns a boolean if a field has been set.

### GetPeerDeviceRegion

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetPeerDeviceRegion() V2AssuranceTopologyInventoryPost200ResponseRegionsInner`

GetPeerDeviceRegion returns the PeerDeviceRegion field if non-nil, zero value otherwise.

### GetPeerDeviceRegionOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetPeerDeviceRegionOk() (*V2AssuranceTopologyInventoryPost200ResponseRegionsInner, bool)`

GetPeerDeviceRegionOk returns a tuple with the PeerDeviceRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDeviceRegion

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) SetPeerDeviceRegion(v V2AssuranceTopologyInventoryPost200ResponseRegionsInner)`

SetPeerDeviceRegion sets PeerDeviceRegion field to given value.

### HasPeerDeviceRegion

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) HasPeerDeviceRegion() bool`

HasPeerDeviceRegion returns a boolean if a field has been set.

### GetSessionName

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetSessionName() string`

GetSessionName returns the SessionName field if non-nil, zero value otherwise.

### GetSessionNameOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) GetSessionNameOk() (*string, bool)`

GetSessionNameOk returns a tuple with the SessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionName

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) SetSessionName(v string)`

SetSessionName sets SessionName field to given value.

### HasSessionName

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseQoeMatrixQoeMatrixInner) HasSessionName() bool`

HasSessionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


