# ManaV2B2bExtranetApplicationInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEmail** | **string** | Admin email of the customer (required) | 
**ConsumerBurstSize** | **int32** | Maximum Burst size per site for the customer (required) | 
**ConsumerBwSite** | **int32** | Maximum Bandwidth allocation per site for the customer (required) | 
**EnterpriseId** | **int64** | Enterprise ID of the customer (required) | 
**MaximumSiteCount** | **int32** | Maximum number of sites for the customer (required) | 
**ServicePrefixes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewManaV2B2bExtranetApplicationInvite

`func NewManaV2B2bExtranetApplicationInvite(adminEmail string, consumerBurstSize int32, consumerBwSite int32, enterpriseId int64, maximumSiteCount int32, ) *ManaV2B2bExtranetApplicationInvite`

NewManaV2B2bExtranetApplicationInvite instantiates a new ManaV2B2bExtranetApplicationInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2B2bExtranetApplicationInviteWithDefaults

`func NewManaV2B2bExtranetApplicationInviteWithDefaults() *ManaV2B2bExtranetApplicationInvite`

NewManaV2B2bExtranetApplicationInviteWithDefaults instantiates a new ManaV2B2bExtranetApplicationInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEmail

`func (o *ManaV2B2bExtranetApplicationInvite) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *ManaV2B2bExtranetApplicationInvite) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *ManaV2B2bExtranetApplicationInvite) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.


### GetConsumerBurstSize

`func (o *ManaV2B2bExtranetApplicationInvite) GetConsumerBurstSize() int32`

GetConsumerBurstSize returns the ConsumerBurstSize field if non-nil, zero value otherwise.

### GetConsumerBurstSizeOk

`func (o *ManaV2B2bExtranetApplicationInvite) GetConsumerBurstSizeOk() (*int32, bool)`

GetConsumerBurstSizeOk returns a tuple with the ConsumerBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerBurstSize

`func (o *ManaV2B2bExtranetApplicationInvite) SetConsumerBurstSize(v int32)`

SetConsumerBurstSize sets ConsumerBurstSize field to given value.


### GetConsumerBwSite

`func (o *ManaV2B2bExtranetApplicationInvite) GetConsumerBwSite() int32`

GetConsumerBwSite returns the ConsumerBwSite field if non-nil, zero value otherwise.

### GetConsumerBwSiteOk

`func (o *ManaV2B2bExtranetApplicationInvite) GetConsumerBwSiteOk() (*int32, bool)`

GetConsumerBwSiteOk returns a tuple with the ConsumerBwSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerBwSite

`func (o *ManaV2B2bExtranetApplicationInvite) SetConsumerBwSite(v int32)`

SetConsumerBwSite sets ConsumerBwSite field to given value.


### GetEnterpriseId

`func (o *ManaV2B2bExtranetApplicationInvite) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *ManaV2B2bExtranetApplicationInvite) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *ManaV2B2bExtranetApplicationInvite) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.


### GetMaximumSiteCount

`func (o *ManaV2B2bExtranetApplicationInvite) GetMaximumSiteCount() int32`

GetMaximumSiteCount returns the MaximumSiteCount field if non-nil, zero value otherwise.

### GetMaximumSiteCountOk

`func (o *ManaV2B2bExtranetApplicationInvite) GetMaximumSiteCountOk() (*int32, bool)`

GetMaximumSiteCountOk returns a tuple with the MaximumSiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSiteCount

`func (o *ManaV2B2bExtranetApplicationInvite) SetMaximumSiteCount(v int32)`

SetMaximumSiteCount sets MaximumSiteCount field to given value.


### GetServicePrefixes

`func (o *ManaV2B2bExtranetApplicationInvite) GetServicePrefixes() []string`

GetServicePrefixes returns the ServicePrefixes field if non-nil, zero value otherwise.

### GetServicePrefixesOk

`func (o *ManaV2B2bExtranetApplicationInvite) GetServicePrefixesOk() (*[]string, bool)`

GetServicePrefixesOk returns a tuple with the ServicePrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrefixes

`func (o *ManaV2B2bExtranetApplicationInvite) SetServicePrefixes(v []string)`

SetServicePrefixes sets ServicePrefixes field to given value.

### HasServicePrefixes

`func (o *ManaV2B2bExtranetApplicationInvite) HasServicePrefixes() bool`

HasServicePrefixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


