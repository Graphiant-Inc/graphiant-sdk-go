# V1ExtranetsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 
**Policies** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInner**](V1ExtranetsGet200ResponsePoliciesInner.md) |  | [optional] 

## Methods

### NewV1ExtranetsGet200Response

`func NewV1ExtranetsGet200Response() *V1ExtranetsGet200Response`

NewV1ExtranetsGet200Response instantiates a new V1ExtranetsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponseWithDefaults

`func NewV1ExtranetsGet200ResponseWithDefaults() *V1ExtranetsGet200Response`

NewV1ExtranetsGet200ResponseWithDefaults instantiates a new V1ExtranetsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1ExtranetsGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1ExtranetsGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1ExtranetsGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1ExtranetsGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetPolicies

`func (o *V1ExtranetsGet200Response) GetPolicies() []V1ExtranetsGet200ResponsePoliciesInner`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *V1ExtranetsGet200Response) GetPoliciesOk() (*[]V1ExtranetsGet200ResponsePoliciesInner, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *V1ExtranetsGet200Response) SetPolicies(v []V1ExtranetsGet200ResponsePoliciesInner)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *V1ExtranetsGet200Response) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


