# V1ExtranetsB2bConsumerPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**[]V1ExtranetsB2bConsumerPostRequestPolicyInner**](V1ExtranetsB2bConsumerPostRequestPolicyInner.md) |  | [optional] 
**ProviderEnterpriseId** | Pointer to **int64** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**SiteInformation** | Pointer to [**[]V1ExtranetsB2bConsumerPostRequestSiteInformationInner**](V1ExtranetsB2bConsumerPostRequestSiteInformationInner.md) |  | [optional] 

## Methods

### NewV1ExtranetsB2bConsumerPostRequest

`func NewV1ExtranetsB2bConsumerPostRequest() *V1ExtranetsB2bConsumerPostRequest`

NewV1ExtranetsB2bConsumerPostRequest instantiates a new V1ExtranetsB2bConsumerPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsB2bConsumerPostRequestWithDefaults

`func NewV1ExtranetsB2bConsumerPostRequestWithDefaults() *V1ExtranetsB2bConsumerPostRequest`

NewV1ExtranetsB2bConsumerPostRequestWithDefaults instantiates a new V1ExtranetsB2bConsumerPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *V1ExtranetsB2bConsumerPostRequest) GetPolicy() []V1ExtranetsB2bConsumerPostRequestPolicyInner`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *V1ExtranetsB2bConsumerPostRequest) GetPolicyOk() (*[]V1ExtranetsB2bConsumerPostRequestPolicyInner, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *V1ExtranetsB2bConsumerPostRequest) SetPolicy(v []V1ExtranetsB2bConsumerPostRequestPolicyInner)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *V1ExtranetsB2bConsumerPostRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProviderEnterpriseId

`func (o *V1ExtranetsB2bConsumerPostRequest) GetProviderEnterpriseId() int64`

GetProviderEnterpriseId returns the ProviderEnterpriseId field if non-nil, zero value otherwise.

### GetProviderEnterpriseIdOk

`func (o *V1ExtranetsB2bConsumerPostRequest) GetProviderEnterpriseIdOk() (*int64, bool)`

GetProviderEnterpriseIdOk returns a tuple with the ProviderEnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderEnterpriseId

`func (o *V1ExtranetsB2bConsumerPostRequest) SetProviderEnterpriseId(v int64)`

SetProviderEnterpriseId sets ProviderEnterpriseId field to given value.

### HasProviderEnterpriseId

`func (o *V1ExtranetsB2bConsumerPostRequest) HasProviderEnterpriseId() bool`

HasProviderEnterpriseId returns a boolean if a field has been set.

### GetServiceName

`func (o *V1ExtranetsB2bConsumerPostRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1ExtranetsB2bConsumerPostRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1ExtranetsB2bConsumerPostRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *V1ExtranetsB2bConsumerPostRequest) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSiteInformation

`func (o *V1ExtranetsB2bConsumerPostRequest) GetSiteInformation() []V1ExtranetsB2bConsumerPostRequestSiteInformationInner`

GetSiteInformation returns the SiteInformation field if non-nil, zero value otherwise.

### GetSiteInformationOk

`func (o *V1ExtranetsB2bConsumerPostRequest) GetSiteInformationOk() (*[]V1ExtranetsB2bConsumerPostRequestSiteInformationInner, bool)`

GetSiteInformationOk returns a tuple with the SiteInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteInformation

`func (o *V1ExtranetsB2bConsumerPostRequest) SetSiteInformation(v []V1ExtranetsB2bConsumerPostRequestSiteInformationInner)`

SetSiteInformation sets SiteInformation field to given value.

### HasSiteInformation

`func (o *V1ExtranetsB2bConsumerPostRequest) HasSiteInformation() bool`

HasSiteInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


