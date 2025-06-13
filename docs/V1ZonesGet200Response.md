# V1ZonesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 
**Zones** | Pointer to [**[]V1ZonesGet200ResponseZonesInner**](V1ZonesGet200ResponseZonesInner.md) |  | [optional] 

## Methods

### NewV1ZonesGet200Response

`func NewV1ZonesGet200Response() *V1ZonesGet200Response`

NewV1ZonesGet200Response instantiates a new V1ZonesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ZonesGet200ResponseWithDefaults

`func NewV1ZonesGet200ResponseWithDefaults() *V1ZonesGet200Response`

NewV1ZonesGet200ResponseWithDefaults instantiates a new V1ZonesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1ZonesGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1ZonesGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1ZonesGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1ZonesGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetZones

`func (o *V1ZonesGet200Response) GetZones() []V1ZonesGet200ResponseZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *V1ZonesGet200Response) GetZonesOk() (*[]V1ZonesGet200ResponseZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *V1ZonesGet200Response) SetZones(v []V1ZonesGet200ResponseZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *V1ZonesGet200Response) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


