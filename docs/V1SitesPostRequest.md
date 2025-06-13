# V1SitesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**Site** | Pointer to [**V1SitesPostRequestSite**](V1SitesPostRequestSite.md) |  | [optional] 

## Methods

### NewV1SitesPostRequest

`func NewV1SitesPostRequest() *V1SitesPostRequest`

NewV1SitesPostRequest instantiates a new V1SitesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SitesPostRequestWithDefaults

`func NewV1SitesPostRequestWithDefaults() *V1SitesPostRequest`

NewV1SitesPostRequestWithDefaults instantiates a new V1SitesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnterpriseId

`func (o *V1SitesPostRequest) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *V1SitesPostRequest) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *V1SitesPostRequest) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *V1SitesPostRequest) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetSite

`func (o *V1SitesPostRequest) GetSite() V1SitesPostRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *V1SitesPostRequest) GetSiteOk() (*V1SitesPostRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *V1SitesPostRequest) SetSite(v V1SitesPostRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *V1SitesPostRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


