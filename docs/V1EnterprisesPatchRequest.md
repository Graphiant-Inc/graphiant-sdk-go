# V1EnterprisesPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEmail** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CreditLimit** | Pointer to **int32** |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**ImpersonationEnabled** | Pointer to **bool** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**PortalBanner** | Pointer to **string** |  | [optional] 
**SmallLogo** | Pointer to **string** |  | [optional] 
**TokenExpiry** | Pointer to [**V1EnterprisesPatchRequestTokenExpiry**](V1EnterprisesPatchRequestTokenExpiry.md) |  | [optional] 

## Methods

### NewV1EnterprisesPatchRequest

`func NewV1EnterprisesPatchRequest() *V1EnterprisesPatchRequest`

NewV1EnterprisesPatchRequest instantiates a new V1EnterprisesPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EnterprisesPatchRequestWithDefaults

`func NewV1EnterprisesPatchRequestWithDefaults() *V1EnterprisesPatchRequest`

NewV1EnterprisesPatchRequestWithDefaults instantiates a new V1EnterprisesPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEmail

`func (o *V1EnterprisesPatchRequest) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *V1EnterprisesPatchRequest) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *V1EnterprisesPatchRequest) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *V1EnterprisesPatchRequest) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetCloudProvider

`func (o *V1EnterprisesPatchRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *V1EnterprisesPatchRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *V1EnterprisesPatchRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *V1EnterprisesPatchRequest) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCompanyName

`func (o *V1EnterprisesPatchRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *V1EnterprisesPatchRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *V1EnterprisesPatchRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *V1EnterprisesPatchRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreditLimit

`func (o *V1EnterprisesPatchRequest) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *V1EnterprisesPatchRequest) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *V1EnterprisesPatchRequest) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *V1EnterprisesPatchRequest) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *V1EnterprisesPatchRequest) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *V1EnterprisesPatchRequest) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *V1EnterprisesPatchRequest) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *V1EnterprisesPatchRequest) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetImpersonationEnabled

`func (o *V1EnterprisesPatchRequest) GetImpersonationEnabled() bool`

GetImpersonationEnabled returns the ImpersonationEnabled field if non-nil, zero value otherwise.

### GetImpersonationEnabledOk

`func (o *V1EnterprisesPatchRequest) GetImpersonationEnabledOk() (*bool, bool)`

GetImpersonationEnabledOk returns a tuple with the ImpersonationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonationEnabled

`func (o *V1EnterprisesPatchRequest) SetImpersonationEnabled(v bool)`

SetImpersonationEnabled sets ImpersonationEnabled field to given value.

### HasImpersonationEnabled

`func (o *V1EnterprisesPatchRequest) HasImpersonationEnabled() bool`

HasImpersonationEnabled returns a boolean if a field has been set.

### GetLogo

`func (o *V1EnterprisesPatchRequest) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *V1EnterprisesPatchRequest) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *V1EnterprisesPatchRequest) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *V1EnterprisesPatchRequest) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPortalBanner

`func (o *V1EnterprisesPatchRequest) GetPortalBanner() string`

GetPortalBanner returns the PortalBanner field if non-nil, zero value otherwise.

### GetPortalBannerOk

`func (o *V1EnterprisesPatchRequest) GetPortalBannerOk() (*string, bool)`

GetPortalBannerOk returns a tuple with the PortalBanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalBanner

`func (o *V1EnterprisesPatchRequest) SetPortalBanner(v string)`

SetPortalBanner sets PortalBanner field to given value.

### HasPortalBanner

`func (o *V1EnterprisesPatchRequest) HasPortalBanner() bool`

HasPortalBanner returns a boolean if a field has been set.

### GetSmallLogo

`func (o *V1EnterprisesPatchRequest) GetSmallLogo() string`

GetSmallLogo returns the SmallLogo field if non-nil, zero value otherwise.

### GetSmallLogoOk

`func (o *V1EnterprisesPatchRequest) GetSmallLogoOk() (*string, bool)`

GetSmallLogoOk returns a tuple with the SmallLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallLogo

`func (o *V1EnterprisesPatchRequest) SetSmallLogo(v string)`

SetSmallLogo sets SmallLogo field to given value.

### HasSmallLogo

`func (o *V1EnterprisesPatchRequest) HasSmallLogo() bool`

HasSmallLogo returns a boolean if a field has been set.

### GetTokenExpiry

`func (o *V1EnterprisesPatchRequest) GetTokenExpiry() V1EnterprisesPatchRequestTokenExpiry`

GetTokenExpiry returns the TokenExpiry field if non-nil, zero value otherwise.

### GetTokenExpiryOk

`func (o *V1EnterprisesPatchRequest) GetTokenExpiryOk() (*V1EnterprisesPatchRequestTokenExpiry, bool)`

GetTokenExpiryOk returns a tuple with the TokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiry

`func (o *V1EnterprisesPatchRequest) SetTokenExpiry(v V1EnterprisesPatchRequestTokenExpiry)`

SetTokenExpiry sets TokenExpiry field to given value.

### HasTokenExpiry

`func (o *V1EnterprisesPatchRequest) HasTokenExpiry() bool`

HasTokenExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


