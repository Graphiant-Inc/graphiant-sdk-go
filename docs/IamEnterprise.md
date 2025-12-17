# IamEnterprise

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptEula** | Pointer to **bool** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**AdminEmail** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Counts** | Pointer to [**IamCounts**](IamCounts.md) |  | [optional] 
**CreditLimit** | Pointer to **int32** |  | [optional] 
**Customers** | Pointer to [**map[string]IamCustomer**](IamCustomer.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**EulaAgreementDate** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**ImpersonationEnabled** | Pointer to **bool** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**MarketplaceId** | Pointer to **string** |  | [optional] 
**ParentCompanyName** | Pointer to **string** |  | [optional] 
**ParentEnterpriseId** | Pointer to **int64** |  | [optional] 
**PortalBanner** | Pointer to **string** |  | [optional] 
**ProxyTenantId** | Pointer to **int64** |  | [optional] 
**SmallLogo** | Pointer to **string** |  | [optional] 
**TokenExpiry** | Pointer to **string** |  | [optional] 

## Methods

### NewIamEnterprise

`func NewIamEnterprise() *IamEnterprise`

NewIamEnterprise instantiates a new IamEnterprise object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEnterpriseWithDefaults

`func NewIamEnterpriseWithDefaults() *IamEnterprise`

NewIamEnterpriseWithDefaults instantiates a new IamEnterprise object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptEula

`func (o *IamEnterprise) GetAcceptEula() bool`

GetAcceptEula returns the AcceptEula field if non-nil, zero value otherwise.

### GetAcceptEulaOk

`func (o *IamEnterprise) GetAcceptEulaOk() (*bool, bool)`

GetAcceptEulaOk returns a tuple with the AcceptEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptEula

`func (o *IamEnterprise) SetAcceptEula(v bool)`

SetAcceptEula sets AcceptEula field to given value.

### HasAcceptEula

`func (o *IamEnterprise) HasAcceptEula() bool`

HasAcceptEula returns a boolean if a field has been set.

### GetAccountType

`func (o *IamEnterprise) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *IamEnterprise) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *IamEnterprise) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *IamEnterprise) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAdminEmail

`func (o *IamEnterprise) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *IamEnterprise) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *IamEnterprise) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *IamEnterprise) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetCloudProvider

`func (o *IamEnterprise) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *IamEnterprise) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *IamEnterprise) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *IamEnterprise) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCompanyName

`func (o *IamEnterprise) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *IamEnterprise) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *IamEnterprise) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *IamEnterprise) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCounts

`func (o *IamEnterprise) GetCounts() IamCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *IamEnterprise) GetCountsOk() (*IamCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *IamEnterprise) SetCounts(v IamCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *IamEnterprise) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetCreditLimit

`func (o *IamEnterprise) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *IamEnterprise) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *IamEnterprise) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *IamEnterprise) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCustomers

`func (o *IamEnterprise) GetCustomers() map[string]IamCustomer`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *IamEnterprise) GetCustomersOk() (*map[string]IamCustomer, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *IamEnterprise) SetCustomers(v map[string]IamCustomer)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *IamEnterprise) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetDescription

`func (o *IamEnterprise) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamEnterprise) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamEnterprise) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamEnterprise) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *IamEnterprise) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *IamEnterprise) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *IamEnterprise) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *IamEnterprise) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetEulaAgreementDate

`func (o *IamEnterprise) GetEulaAgreementDate() GoogleProtobufTimestamp`

GetEulaAgreementDate returns the EulaAgreementDate field if non-nil, zero value otherwise.

### GetEulaAgreementDateOk

`func (o *IamEnterprise) GetEulaAgreementDateOk() (*GoogleProtobufTimestamp, bool)`

GetEulaAgreementDateOk returns a tuple with the EulaAgreementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAgreementDate

`func (o *IamEnterprise) SetEulaAgreementDate(v GoogleProtobufTimestamp)`

SetEulaAgreementDate sets EulaAgreementDate field to given value.

### HasEulaAgreementDate

`func (o *IamEnterprise) HasEulaAgreementDate() bool`

HasEulaAgreementDate returns a boolean if a field has been set.

### GetImpersonationEnabled

`func (o *IamEnterprise) GetImpersonationEnabled() bool`

GetImpersonationEnabled returns the ImpersonationEnabled field if non-nil, zero value otherwise.

### GetImpersonationEnabledOk

`func (o *IamEnterprise) GetImpersonationEnabledOk() (*bool, bool)`

GetImpersonationEnabledOk returns a tuple with the ImpersonationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonationEnabled

`func (o *IamEnterprise) SetImpersonationEnabled(v bool)`

SetImpersonationEnabled sets ImpersonationEnabled field to given value.

### HasImpersonationEnabled

`func (o *IamEnterprise) HasImpersonationEnabled() bool`

HasImpersonationEnabled returns a boolean if a field has been set.

### GetLogo

`func (o *IamEnterprise) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *IamEnterprise) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *IamEnterprise) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *IamEnterprise) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetMarketplaceId

`func (o *IamEnterprise) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *IamEnterprise) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *IamEnterprise) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *IamEnterprise) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetParentCompanyName

`func (o *IamEnterprise) GetParentCompanyName() string`

GetParentCompanyName returns the ParentCompanyName field if non-nil, zero value otherwise.

### GetParentCompanyNameOk

`func (o *IamEnterprise) GetParentCompanyNameOk() (*string, bool)`

GetParentCompanyNameOk returns a tuple with the ParentCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCompanyName

`func (o *IamEnterprise) SetParentCompanyName(v string)`

SetParentCompanyName sets ParentCompanyName field to given value.

### HasParentCompanyName

`func (o *IamEnterprise) HasParentCompanyName() bool`

HasParentCompanyName returns a boolean if a field has been set.

### GetParentEnterpriseId

`func (o *IamEnterprise) GetParentEnterpriseId() int64`

GetParentEnterpriseId returns the ParentEnterpriseId field if non-nil, zero value otherwise.

### GetParentEnterpriseIdOk

`func (o *IamEnterprise) GetParentEnterpriseIdOk() (*int64, bool)`

GetParentEnterpriseIdOk returns a tuple with the ParentEnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEnterpriseId

`func (o *IamEnterprise) SetParentEnterpriseId(v int64)`

SetParentEnterpriseId sets ParentEnterpriseId field to given value.

### HasParentEnterpriseId

`func (o *IamEnterprise) HasParentEnterpriseId() bool`

HasParentEnterpriseId returns a boolean if a field has been set.

### GetPortalBanner

`func (o *IamEnterprise) GetPortalBanner() string`

GetPortalBanner returns the PortalBanner field if non-nil, zero value otherwise.

### GetPortalBannerOk

`func (o *IamEnterprise) GetPortalBannerOk() (*string, bool)`

GetPortalBannerOk returns a tuple with the PortalBanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalBanner

`func (o *IamEnterprise) SetPortalBanner(v string)`

SetPortalBanner sets PortalBanner field to given value.

### HasPortalBanner

`func (o *IamEnterprise) HasPortalBanner() bool`

HasPortalBanner returns a boolean if a field has been set.

### GetProxyTenantId

`func (o *IamEnterprise) GetProxyTenantId() int64`

GetProxyTenantId returns the ProxyTenantId field if non-nil, zero value otherwise.

### GetProxyTenantIdOk

`func (o *IamEnterprise) GetProxyTenantIdOk() (*int64, bool)`

GetProxyTenantIdOk returns a tuple with the ProxyTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyTenantId

`func (o *IamEnterprise) SetProxyTenantId(v int64)`

SetProxyTenantId sets ProxyTenantId field to given value.

### HasProxyTenantId

`func (o *IamEnterprise) HasProxyTenantId() bool`

HasProxyTenantId returns a boolean if a field has been set.

### GetSmallLogo

`func (o *IamEnterprise) GetSmallLogo() string`

GetSmallLogo returns the SmallLogo field if non-nil, zero value otherwise.

### GetSmallLogoOk

`func (o *IamEnterprise) GetSmallLogoOk() (*string, bool)`

GetSmallLogoOk returns a tuple with the SmallLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallLogo

`func (o *IamEnterprise) SetSmallLogo(v string)`

SetSmallLogo sets SmallLogo field to given value.

### HasSmallLogo

`func (o *IamEnterprise) HasSmallLogo() bool`

HasSmallLogo returns a boolean if a field has been set.

### GetTokenExpiry

`func (o *IamEnterprise) GetTokenExpiry() string`

GetTokenExpiry returns the TokenExpiry field if non-nil, zero value otherwise.

### GetTokenExpiryOk

`func (o *IamEnterprise) GetTokenExpiryOk() (*string, bool)`

GetTokenExpiryOk returns a tuple with the TokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiry

`func (o *IamEnterprise) SetTokenExpiry(v string)`

SetTokenExpiry sets TokenExpiry field to given value.

### HasTokenExpiry

`func (o *IamEnterprise) HasTokenExpiry() bool`

HasTokenExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


