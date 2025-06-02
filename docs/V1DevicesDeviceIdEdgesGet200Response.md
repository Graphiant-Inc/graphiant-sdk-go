# V1DevicesDeviceIdEdgesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInner.md) |  | [optional] 
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdEdgesGet200Response

`func NewV1DevicesDeviceIdEdgesGet200Response() *V1DevicesDeviceIdEdgesGet200Response`

NewV1DevicesDeviceIdEdgesGet200Response instantiates a new V1DevicesDeviceIdEdgesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdEdgesGet200ResponseWithDefaults

`func NewV1DevicesDeviceIdEdgesGet200ResponseWithDefaults() *V1DevicesDeviceIdEdgesGet200Response`

NewV1DevicesDeviceIdEdgesGet200ResponseWithDefaults instantiates a new V1DevicesDeviceIdEdgesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *V1DevicesDeviceIdEdgesGet200Response) GetDevices() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *V1DevicesDeviceIdEdgesGet200Response) GetDevicesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *V1DevicesDeviceIdEdgesGet200Response) SetDevices(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInner)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *V1DevicesDeviceIdEdgesGet200Response) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetPageInfo

`func (o *V1DevicesDeviceIdEdgesGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1DevicesDeviceIdEdgesGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1DevicesDeviceIdEdgesGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1DevicesDeviceIdEdgesGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


