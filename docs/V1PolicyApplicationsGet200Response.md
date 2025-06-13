# V1PolicyApplicationsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]V1PolicyApplicationsGet200ResponseApplicationsInner**](V1PolicyApplicationsGet200ResponseApplicationsInner.md) |  | [optional] 
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 

## Methods

### NewV1PolicyApplicationsGet200Response

`func NewV1PolicyApplicationsGet200Response() *V1PolicyApplicationsGet200Response`

NewV1PolicyApplicationsGet200Response instantiates a new V1PolicyApplicationsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PolicyApplicationsGet200ResponseWithDefaults

`func NewV1PolicyApplicationsGet200ResponseWithDefaults() *V1PolicyApplicationsGet200Response`

NewV1PolicyApplicationsGet200ResponseWithDefaults instantiates a new V1PolicyApplicationsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *V1PolicyApplicationsGet200Response) GetApplications() []V1PolicyApplicationsGet200ResponseApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *V1PolicyApplicationsGet200Response) GetApplicationsOk() (*[]V1PolicyApplicationsGet200ResponseApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *V1PolicyApplicationsGet200Response) SetApplications(v []V1PolicyApplicationsGet200ResponseApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *V1PolicyApplicationsGet200Response) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetPageInfo

`func (o *V1PolicyApplicationsGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1PolicyApplicationsGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1PolicyApplicationsGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1PolicyApplicationsGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


