# V1DevicesDeviceIdInterfacesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interfaces** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner.md) |  | [optional] 
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdInterfacesGet200Response

`func NewV1DevicesDeviceIdInterfacesGet200Response() *V1DevicesDeviceIdInterfacesGet200Response`

NewV1DevicesDeviceIdInterfacesGet200Response instantiates a new V1DevicesDeviceIdInterfacesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdInterfacesGet200ResponseWithDefaults

`func NewV1DevicesDeviceIdInterfacesGet200ResponseWithDefaults() *V1DevicesDeviceIdInterfacesGet200Response`

NewV1DevicesDeviceIdInterfacesGet200ResponseWithDefaults instantiates a new V1DevicesDeviceIdInterfacesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaces

`func (o *V1DevicesDeviceIdInterfacesGet200Response) GetInterfaces() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *V1DevicesDeviceIdInterfacesGet200Response) GetInterfacesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *V1DevicesDeviceIdInterfacesGet200Response) SetInterfaces(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *V1DevicesDeviceIdInterfacesGet200Response) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetPageInfo

`func (o *V1DevicesDeviceIdInterfacesGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1DevicesDeviceIdInterfacesGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1DevicesDeviceIdInterfacesGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1DevicesDeviceIdInterfacesGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


