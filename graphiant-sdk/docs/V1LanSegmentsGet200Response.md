# V1LanSegmentsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 
**Segments** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner.md) |  | [optional] 

## Methods

### NewV1LanSegmentsGet200Response

`func NewV1LanSegmentsGet200Response() *V1LanSegmentsGet200Response`

NewV1LanSegmentsGet200Response instantiates a new V1LanSegmentsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LanSegmentsGet200ResponseWithDefaults

`func NewV1LanSegmentsGet200ResponseWithDefaults() *V1LanSegmentsGet200Response`

NewV1LanSegmentsGet200ResponseWithDefaults instantiates a new V1LanSegmentsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1LanSegmentsGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1LanSegmentsGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1LanSegmentsGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1LanSegmentsGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetSegments

`func (o *V1LanSegmentsGet200Response) GetSegments() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *V1LanSegmentsGet200Response) GetSegmentsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *V1LanSegmentsGet200Response) SetSegments(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInner)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *V1LanSegmentsGet200Response) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


