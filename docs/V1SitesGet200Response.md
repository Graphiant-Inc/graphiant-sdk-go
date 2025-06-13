# V1SitesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**V1NatEntriesDeviceIdGet200ResponsePageInfo**](V1NatEntriesDeviceIdGet200ResponsePageInfo.md) |  | [optional] 
**Sites** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite.md) |  | [optional] 

## Methods

### NewV1SitesGet200Response

`func NewV1SitesGet200Response() *V1SitesGet200Response`

NewV1SitesGet200Response instantiates a new V1SitesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SitesGet200ResponseWithDefaults

`func NewV1SitesGet200ResponseWithDefaults() *V1SitesGet200Response`

NewV1SitesGet200ResponseWithDefaults instantiates a new V1SitesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *V1SitesGet200Response) GetPageInfo() V1NatEntriesDeviceIdGet200ResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1SitesGet200Response) GetPageInfoOk() (*V1NatEntriesDeviceIdGet200ResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1SitesGet200Response) SetPageInfo(v V1NatEntriesDeviceIdGet200ResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1SitesGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetSites

`func (o *V1SitesGet200Response) GetSites() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *V1SitesGet200Response) GetSitesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *V1SitesGet200Response) SetSites(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSite)`

SetSites sets Sites field to given value.

### HasSites

`func (o *V1SitesGet200Response) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


