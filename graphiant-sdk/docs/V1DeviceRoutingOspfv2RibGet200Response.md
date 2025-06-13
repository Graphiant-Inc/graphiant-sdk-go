# V1DeviceRoutingOspfv2RibGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 
**Routes** | Pointer to [**[]V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner**](V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DeviceRoutingOspfv2RibGet200Response

`func NewV1DeviceRoutingOspfv2RibGet200Response() *V1DeviceRoutingOspfv2RibGet200Response`

NewV1DeviceRoutingOspfv2RibGet200Response instantiates a new V1DeviceRoutingOspfv2RibGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeviceRoutingOspfv2RibGet200ResponseWithDefaults

`func NewV1DeviceRoutingOspfv2RibGet200ResponseWithDefaults() *V1DeviceRoutingOspfv2RibGet200Response`

NewV1DeviceRoutingOspfv2RibGet200ResponseWithDefaults instantiates a new V1DeviceRoutingOspfv2RibGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1DeviceRoutingOspfv2RibGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1DeviceRoutingOspfv2RibGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1DeviceRoutingOspfv2RibGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1DeviceRoutingOspfv2RibGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetRoutes

`func (o *V1DeviceRoutingOspfv2RibGet200Response) GetRoutes() []V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *V1DeviceRoutingOspfv2RibGet200Response) GetRoutesOk() (*[]V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *V1DeviceRoutingOspfv2RibGet200Response) SetRoutes(v []V1DeviceRoutingOspfv2RibGet200ResponseRoutesInner)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *V1DeviceRoutingOspfv2RibGet200Response) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetToken

`func (o *V1DeviceRoutingOspfv2RibGet200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *V1DeviceRoutingOspfv2RibGet200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *V1DeviceRoutingOspfv2RibGet200Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *V1DeviceRoutingOspfv2RibGet200Response) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


