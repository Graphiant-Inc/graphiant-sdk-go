# V1NatEntriesDeviceIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatEntries** | Pointer to [**[]V1NatEntriesDeviceIdGet200ResponseNatEntriesInner**](V1NatEntriesDeviceIdGet200ResponseNatEntriesInner.md) |  | [optional] 
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 
**Ts** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1NatEntriesDeviceIdGet200Response

`func NewV1NatEntriesDeviceIdGet200Response() *V1NatEntriesDeviceIdGet200Response`

NewV1NatEntriesDeviceIdGet200Response instantiates a new V1NatEntriesDeviceIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NatEntriesDeviceIdGet200ResponseWithDefaults

`func NewV1NatEntriesDeviceIdGet200ResponseWithDefaults() *V1NatEntriesDeviceIdGet200Response`

NewV1NatEntriesDeviceIdGet200ResponseWithDefaults instantiates a new V1NatEntriesDeviceIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatEntries

`func (o *V1NatEntriesDeviceIdGet200Response) GetNatEntries() []V1NatEntriesDeviceIdGet200ResponseNatEntriesInner`

GetNatEntries returns the NatEntries field if non-nil, zero value otherwise.

### GetNatEntriesOk

`func (o *V1NatEntriesDeviceIdGet200Response) GetNatEntriesOk() (*[]V1NatEntriesDeviceIdGet200ResponseNatEntriesInner, bool)`

GetNatEntriesOk returns a tuple with the NatEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEntries

`func (o *V1NatEntriesDeviceIdGet200Response) SetNatEntries(v []V1NatEntriesDeviceIdGet200ResponseNatEntriesInner)`

SetNatEntries sets NatEntries field to given value.

### HasNatEntries

`func (o *V1NatEntriesDeviceIdGet200Response) HasNatEntries() bool`

HasNatEntries returns a boolean if a field has been set.

### GetPageInfo

`func (o *V1NatEntriesDeviceIdGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1NatEntriesDeviceIdGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1NatEntriesDeviceIdGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1NatEntriesDeviceIdGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetTs

`func (o *V1NatEntriesDeviceIdGet200Response) GetTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *V1NatEntriesDeviceIdGet200Response) GetTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *V1NatEntriesDeviceIdGet200Response) SetTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetTs sets Ts field to given value.

### HasTs

`func (o *V1NatEntriesDeviceIdGet200Response) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


