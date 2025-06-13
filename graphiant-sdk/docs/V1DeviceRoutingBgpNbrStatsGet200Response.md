# V1DeviceRoutingBgpNbrStatsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 
**Stats** | Pointer to [**V1DeviceRoutingBgpNbrStatsGet200ResponseStats**](V1DeviceRoutingBgpNbrStatsGet200ResponseStats.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DeviceRoutingBgpNbrStatsGet200Response

`func NewV1DeviceRoutingBgpNbrStatsGet200Response() *V1DeviceRoutingBgpNbrStatsGet200Response`

NewV1DeviceRoutingBgpNbrStatsGet200Response instantiates a new V1DeviceRoutingBgpNbrStatsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeviceRoutingBgpNbrStatsGet200ResponseWithDefaults

`func NewV1DeviceRoutingBgpNbrStatsGet200ResponseWithDefaults() *V1DeviceRoutingBgpNbrStatsGet200Response`

NewV1DeviceRoutingBgpNbrStatsGet200ResponseWithDefaults instantiates a new V1DeviceRoutingBgpNbrStatsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetStats

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetStats() V1DeviceRoutingBgpNbrStatsGet200ResponseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetStatsOk() (*V1DeviceRoutingBgpNbrStatsGet200ResponseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) SetStats(v V1DeviceRoutingBgpNbrStatsGet200ResponseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetToken

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *V1DeviceRoutingBgpNbrStatsGet200Response) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


