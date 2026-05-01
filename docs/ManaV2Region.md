# ManaV2Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinates** | Pointer to [**ManaV2RegionCoordinates**](ManaV2RegionCoordinates.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RegionIsoCode** | Pointer to **string** |  | [optional] 
**Unavailable** | Pointer to **bool** |  | [optional] 

## Methods

### NewManaV2Region

`func NewManaV2Region() *ManaV2Region`

NewManaV2Region instantiates a new ManaV2Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2RegionWithDefaults

`func NewManaV2RegionWithDefaults() *ManaV2Region`

NewManaV2RegionWithDefaults instantiates a new ManaV2Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinates

`func (o *ManaV2Region) GetCoordinates() ManaV2RegionCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *ManaV2Region) GetCoordinatesOk() (*ManaV2RegionCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *ManaV2Region) SetCoordinates(v ManaV2RegionCoordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *ManaV2Region) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetId

`func (o *ManaV2Region) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2Region) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2Region) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManaV2Region) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ManaV2Region) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2Region) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2Region) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2Region) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegionIsoCode

`func (o *ManaV2Region) GetRegionIsoCode() string`

GetRegionIsoCode returns the RegionIsoCode field if non-nil, zero value otherwise.

### GetRegionIsoCodeOk

`func (o *ManaV2Region) GetRegionIsoCodeOk() (*string, bool)`

GetRegionIsoCodeOk returns a tuple with the RegionIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIsoCode

`func (o *ManaV2Region) SetRegionIsoCode(v string)`

SetRegionIsoCode sets RegionIsoCode field to given value.

### HasRegionIsoCode

`func (o *ManaV2Region) HasRegionIsoCode() bool`

HasRegionIsoCode returns a boolean if a field has been set.

### GetUnavailable

`func (o *ManaV2Region) GetUnavailable() bool`

GetUnavailable returns the Unavailable field if non-nil, zero value otherwise.

### GetUnavailableOk

`func (o *ManaV2Region) GetUnavailableOk() (*bool, bool)`

GetUnavailableOk returns a tuple with the Unavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailable

`func (o *ManaV2Region) SetUnavailable(v bool)`

SetUnavailable sets Unavailable field to given value.

### HasUnavailable

`func (o *ManaV2Region) HasUnavailable() bool`

HasUnavailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


