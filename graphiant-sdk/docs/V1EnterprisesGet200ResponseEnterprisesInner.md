# V1EnterprisesGet200ResponseEnterprisesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptEula** | Pointer to **bool** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**AdminEmail** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Counts** | Pointer to [**V1EnterprisesGet200ResponseEnterprisesInnerCounts**](V1EnterprisesGet200ResponseEnterprisesInnerCounts.md) |  | [optional] 
**CreditLimit** | Pointer to **int32** |  | [optional] 
**Customers** | Pointer to [**map[string]V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue**](V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue.md) |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**EulaAgreementDate** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**ImpersonationEnabled** | Pointer to **bool** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**ParentCompanyName** | Pointer to **string** |  | [optional] 
**ParentEnterpriseId** | Pointer to **int64** |  | [optional] 
**PortalBanner** | Pointer to **string** |  | [optional] 
**SmallLogo** | Pointer to **string** |  | [optional] 
**TokenExpiry** | Pointer to **string** |  | [optional] 

## Methods

### NewV1EnterprisesGet200ResponseEnterprisesInner

`func NewV1EnterprisesGet200ResponseEnterprisesInner() *V1EnterprisesGet200ResponseEnterprisesInner`

NewV1EnterprisesGet200ResponseEnterprisesInner instantiates a new V1EnterprisesGet200ResponseEnterprisesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EnterprisesGet200ResponseEnterprisesInnerWithDefaults

`func NewV1EnterprisesGet200ResponseEnterprisesInnerWithDefaults() *V1EnterprisesGet200ResponseEnterprisesInner`

NewV1EnterprisesGet200ResponseEnterprisesInnerWithDefaults instantiates a new V1EnterprisesGet200ResponseEnterprisesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptEula

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetAcceptEula() bool`

GetAcceptEula returns the AcceptEula field if non-nil, zero value otherwise.

### GetAcceptEulaOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetAcceptEulaOk() (*bool, bool)`

GetAcceptEulaOk returns a tuple with the AcceptEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptEula

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetAcceptEula(v bool)`

SetAcceptEula sets AcceptEula field to given value.

### HasAcceptEula

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasAcceptEula() bool`

HasAcceptEula returns a boolean if a field has been set.

### GetAccountType

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAdminEmail

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetCloudProvider

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCompanyName

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCounts

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetCounts() V1EnterprisesGet200ResponseEnterprisesInnerCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetCountsOk() (*V1EnterprisesGet200ResponseEnterprisesInnerCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetCounts(v V1EnterprisesGet200ResponseEnterprisesInnerCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetCreditLimit

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCustomers

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetCustomers() map[string]V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetCustomersOk() (*map[string]V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetCustomers(v map[string]V1EnterprisesGet200ResponseEnterprisesInnerCustomersValue)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetEulaAgreementDate

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetEulaAgreementDate() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetEulaAgreementDate returns the EulaAgreementDate field if non-nil, zero value otherwise.

### GetEulaAgreementDateOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetEulaAgreementDateOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetEulaAgreementDateOk returns a tuple with the EulaAgreementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAgreementDate

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetEulaAgreementDate(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetEulaAgreementDate sets EulaAgreementDate field to given value.

### HasEulaAgreementDate

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasEulaAgreementDate() bool`

HasEulaAgreementDate returns a boolean if a field has been set.

### GetImpersonationEnabled

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetImpersonationEnabled() bool`

GetImpersonationEnabled returns the ImpersonationEnabled field if non-nil, zero value otherwise.

### GetImpersonationEnabledOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetImpersonationEnabledOk() (*bool, bool)`

GetImpersonationEnabledOk returns a tuple with the ImpersonationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonationEnabled

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetImpersonationEnabled(v bool)`

SetImpersonationEnabled sets ImpersonationEnabled field to given value.

### HasImpersonationEnabled

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasImpersonationEnabled() bool`

HasImpersonationEnabled returns a boolean if a field has been set.

### GetLogo

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetParentCompanyName

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetParentCompanyName() string`

GetParentCompanyName returns the ParentCompanyName field if non-nil, zero value otherwise.

### GetParentCompanyNameOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetParentCompanyNameOk() (*string, bool)`

GetParentCompanyNameOk returns a tuple with the ParentCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCompanyName

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetParentCompanyName(v string)`

SetParentCompanyName sets ParentCompanyName field to given value.

### HasParentCompanyName

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasParentCompanyName() bool`

HasParentCompanyName returns a boolean if a field has been set.

### GetParentEnterpriseId

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetParentEnterpriseId() int64`

GetParentEnterpriseId returns the ParentEnterpriseId field if non-nil, zero value otherwise.

### GetParentEnterpriseIdOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetParentEnterpriseIdOk() (*int64, bool)`

GetParentEnterpriseIdOk returns a tuple with the ParentEnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEnterpriseId

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetParentEnterpriseId(v int64)`

SetParentEnterpriseId sets ParentEnterpriseId field to given value.

### HasParentEnterpriseId

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasParentEnterpriseId() bool`

HasParentEnterpriseId returns a boolean if a field has been set.

### GetPortalBanner

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetPortalBanner() string`

GetPortalBanner returns the PortalBanner field if non-nil, zero value otherwise.

### GetPortalBannerOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetPortalBannerOk() (*string, bool)`

GetPortalBannerOk returns a tuple with the PortalBanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalBanner

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetPortalBanner(v string)`

SetPortalBanner sets PortalBanner field to given value.

### HasPortalBanner

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasPortalBanner() bool`

HasPortalBanner returns a boolean if a field has been set.

### GetSmallLogo

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetSmallLogo() string`

GetSmallLogo returns the SmallLogo field if non-nil, zero value otherwise.

### GetSmallLogoOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetSmallLogoOk() (*string, bool)`

GetSmallLogoOk returns a tuple with the SmallLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallLogo

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetSmallLogo(v string)`

SetSmallLogo sets SmallLogo field to given value.

### HasSmallLogo

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasSmallLogo() bool`

HasSmallLogo returns a boolean if a field has been set.

### GetTokenExpiry

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetTokenExpiry() string`

GetTokenExpiry returns the TokenExpiry field if non-nil, zero value otherwise.

### GetTokenExpiryOk

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) GetTokenExpiryOk() (*string, bool)`

GetTokenExpiryOk returns a tuple with the TokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiry

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) SetTokenExpiry(v string)`

SetTokenExpiry sets TokenExpiry field to given value.

### HasTokenExpiry

`func (o *V1EnterprisesGet200ResponseEnterprisesInner) HasTokenExpiry() bool`

HasTokenExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


