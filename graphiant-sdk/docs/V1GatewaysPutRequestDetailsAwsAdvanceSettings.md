# V1GatewaysPutRequestDetailsAwsAdvanceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **string** |  | [optional] 
**AllowedPrefixes** | Pointer to **[]string** |  | [optional] 
**AmazonBgpRouterIp** | Pointer to **string** |  | [optional] 
**BgpAuthKey** | Pointer to **string** |  | [optional] 
**CustomerBgpRouterIp** | Pointer to **string** |  | [optional] 
**IsJumbo** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1GatewaysPutRequestDetailsAwsAdvanceSettings

`func NewV1GatewaysPutRequestDetailsAwsAdvanceSettings() *V1GatewaysPutRequestDetailsAwsAdvanceSettings`

NewV1GatewaysPutRequestDetailsAwsAdvanceSettings instantiates a new V1GatewaysPutRequestDetailsAwsAdvanceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GatewaysPutRequestDetailsAwsAdvanceSettingsWithDefaults

`func NewV1GatewaysPutRequestDetailsAwsAdvanceSettingsWithDefaults() *V1GatewaysPutRequestDetailsAwsAdvanceSettings`

NewV1GatewaysPutRequestDetailsAwsAdvanceSettingsWithDefaults instantiates a new V1GatewaysPutRequestDetailsAwsAdvanceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAllowedPrefixes

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetAllowedPrefixes() []string`

GetAllowedPrefixes returns the AllowedPrefixes field if non-nil, zero value otherwise.

### GetAllowedPrefixesOk

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetAllowedPrefixesOk() (*[]string, bool)`

GetAllowedPrefixesOk returns a tuple with the AllowedPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPrefixes

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) SetAllowedPrefixes(v []string)`

SetAllowedPrefixes sets AllowedPrefixes field to given value.

### HasAllowedPrefixes

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) HasAllowedPrefixes() bool`

HasAllowedPrefixes returns a boolean if a field has been set.

### GetAmazonBgpRouterIp

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetAmazonBgpRouterIp() string`

GetAmazonBgpRouterIp returns the AmazonBgpRouterIp field if non-nil, zero value otherwise.

### GetAmazonBgpRouterIpOk

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetAmazonBgpRouterIpOk() (*string, bool)`

GetAmazonBgpRouterIpOk returns a tuple with the AmazonBgpRouterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonBgpRouterIp

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) SetAmazonBgpRouterIp(v string)`

SetAmazonBgpRouterIp sets AmazonBgpRouterIp field to given value.

### HasAmazonBgpRouterIp

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) HasAmazonBgpRouterIp() bool`

HasAmazonBgpRouterIp returns a boolean if a field has been set.

### GetBgpAuthKey

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetCustomerBgpRouterIp

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetCustomerBgpRouterIp() string`

GetCustomerBgpRouterIp returns the CustomerBgpRouterIp field if non-nil, zero value otherwise.

### GetCustomerBgpRouterIpOk

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetCustomerBgpRouterIpOk() (*string, bool)`

GetCustomerBgpRouterIpOk returns a tuple with the CustomerBgpRouterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBgpRouterIp

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) SetCustomerBgpRouterIp(v string)`

SetCustomerBgpRouterIp sets CustomerBgpRouterIp field to given value.

### HasCustomerBgpRouterIp

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) HasCustomerBgpRouterIp() bool`

HasCustomerBgpRouterIp returns a boolean if a field has been set.

### GetIsJumbo

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetIsJumbo() bool`

GetIsJumbo returns the IsJumbo field if non-nil, zero value otherwise.

### GetIsJumboOk

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) GetIsJumboOk() (*bool, bool)`

GetIsJumboOk returns a tuple with the IsJumbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJumbo

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) SetIsJumbo(v bool)`

SetIsJumbo sets IsJumbo field to given value.

### HasIsJumbo

`func (o *V1GatewaysPutRequestDetailsAwsAdvanceSettings) HasIsJumbo() bool`

HasIsJumbo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


