# V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**IpLists** | Pointer to **[]string** |  | [optional] 
**IpPrefixes** | Pointer to **[]string** |  | [optional] 
**IpProtocol** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PortRanges** | Pointer to [**[]V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange**](V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewV1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig

`func NewV1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig() *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig`

NewV1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig instantiates a new V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalAppsCustomGet200ResponseEntriesInnerAppConfigWithDefaults

`func NewV1GlobalAppsCustomGet200ResponseEntriesInnerAppConfigWithDefaults() *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig`

NewV1GlobalAppsCustomGet200ResponseEntriesInnerAppConfigWithDefaults instantiates a new V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpLists

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetIpLists() []string`

GetIpLists returns the IpLists field if non-nil, zero value otherwise.

### GetIpListsOk

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetIpListsOk() (*[]string, bool)`

GetIpListsOk returns a tuple with the IpLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLists

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) SetIpLists(v []string)`

SetIpLists sets IpLists field to given value.

### HasIpLists

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) HasIpLists() bool`

HasIpLists returns a boolean if a field has been set.

### GetIpPrefixes

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetIpPrefixes() []string`

GetIpPrefixes returns the IpPrefixes field if non-nil, zero value otherwise.

### GetIpPrefixesOk

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetIpPrefixesOk() (*[]string, bool)`

GetIpPrefixesOk returns a tuple with the IpPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPrefixes

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) SetIpPrefixes(v []string)`

SetIpPrefixes sets IpPrefixes field to given value.

### HasIpPrefixes

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) HasIpPrefixes() bool`

HasIpPrefixes returns a boolean if a field has been set.

### GetIpProtocol

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetIpProtocol() string`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetIpProtocolOk() (*string, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpProtocol

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) SetIpProtocol(v string)`

SetIpProtocol sets IpProtocol field to given value.

### HasIpProtocol

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) HasIpProtocol() bool`

HasIpProtocol returns a boolean if a field has been set.

### GetName

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortRanges

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetPortRanges() []V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange`

GetPortRanges returns the PortRanges field if non-nil, zero value otherwise.

### GetPortRangesOk

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetPortRangesOk() (*[]V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange, bool)`

GetPortRangesOk returns a tuple with the PortRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRanges

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) SetPortRanges(v []V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange)`

SetPortRanges sets PortRanges field to given value.

### HasPortRanges

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) HasPortRanges() bool`

HasPortRanges returns a boolean if a field has been set.

### GetUrl

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V1GlobalAppsCustomGet200ResponseEntriesInnerAppConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


