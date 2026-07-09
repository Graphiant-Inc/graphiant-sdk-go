# AssuranceApprovedAppEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | app name to approve (required) | 
**Domain** | Pointer to **string** | app name to approve | [optional] 
**Id** | Pointer to **string** | approved app entry identifier | [optional] 
**TagRequested** | **string** | resulting tag, approved or shadow (required) | 

## Methods

### NewAssuranceApprovedAppEntryRequest

`func NewAssuranceApprovedAppEntryRequest(appName string, tagRequested string, ) *AssuranceApprovedAppEntryRequest`

NewAssuranceApprovedAppEntryRequest instantiates a new AssuranceApprovedAppEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssuranceApprovedAppEntryRequestWithDefaults

`func NewAssuranceApprovedAppEntryRequestWithDefaults() *AssuranceApprovedAppEntryRequest`

NewAssuranceApprovedAppEntryRequestWithDefaults instantiates a new AssuranceApprovedAppEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *AssuranceApprovedAppEntryRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AssuranceApprovedAppEntryRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AssuranceApprovedAppEntryRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetDomain

`func (o *AssuranceApprovedAppEntryRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AssuranceApprovedAppEntryRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AssuranceApprovedAppEntryRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AssuranceApprovedAppEntryRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetId

`func (o *AssuranceApprovedAppEntryRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssuranceApprovedAppEntryRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssuranceApprovedAppEntryRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssuranceApprovedAppEntryRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTagRequested

`func (o *AssuranceApprovedAppEntryRequest) GetTagRequested() string`

GetTagRequested returns the TagRequested field if non-nil, zero value otherwise.

### GetTagRequestedOk

`func (o *AssuranceApprovedAppEntryRequest) GetTagRequestedOk() (*string, bool)`

GetTagRequestedOk returns a tuple with the TagRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagRequested

`func (o *AssuranceApprovedAppEntryRequest) SetTagRequested(v string)`

SetTagRequested sets TagRequested field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


