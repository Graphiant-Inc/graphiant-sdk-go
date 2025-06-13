# V1DeviceRoutingOspfv2AreaLsdbGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lsas** | Pointer to [**[]V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner**](V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner.md) |  | [optional] 
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DeviceRoutingOspfv2AreaLsdbGet200Response

`func NewV1DeviceRoutingOspfv2AreaLsdbGet200Response() *V1DeviceRoutingOspfv2AreaLsdbGet200Response`

NewV1DeviceRoutingOspfv2AreaLsdbGet200Response instantiates a new V1DeviceRoutingOspfv2AreaLsdbGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseWithDefaults

`func NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseWithDefaults() *V1DeviceRoutingOspfv2AreaLsdbGet200Response`

NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseWithDefaults instantiates a new V1DeviceRoutingOspfv2AreaLsdbGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLsas

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) GetLsas() []V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner`

GetLsas returns the Lsas field if non-nil, zero value otherwise.

### GetLsasOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) GetLsasOk() (*[]V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner, bool)`

GetLsasOk returns a tuple with the Lsas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLsas

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) SetLsas(v []V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner)`

SetLsas sets Lsas field to given value.

### HasLsas

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) HasLsas() bool`

HasLsas returns a boolean if a field has been set.

### GetPageInfo

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetToken

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200Response) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


