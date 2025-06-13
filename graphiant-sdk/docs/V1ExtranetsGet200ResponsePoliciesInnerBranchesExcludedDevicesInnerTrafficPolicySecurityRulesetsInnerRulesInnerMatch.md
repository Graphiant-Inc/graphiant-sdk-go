# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**DestinationNetwork** | Pointer to **string** |  | [optional] 
**DestinationPort** | Pointer to **int32** |  | [optional] 
**DestinationPortRange** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange**](V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange.md) |  | [optional] 
**DomainCategoryIds** | Pointer to **[]int64** |  | [optional] 
**Dscp** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscpMatch**](V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscpMatch.md) |  | [optional] 
**IcmpType** | Pointer to **int32** |  | [optional] 
**IpProtocol** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**SourceNetwork** | Pointer to **string** |  | [optional] 
**SourcePort** | Pointer to **int32** |  | [optional] 
**SourcePortRange** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange**](V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange.md) |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatchWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatchWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatchWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDestinationNetwork

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetDestinationNetwork() string`

GetDestinationNetwork returns the DestinationNetwork field if non-nil, zero value otherwise.

### GetDestinationNetworkOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetDestinationNetworkOk() (*string, bool)`

GetDestinationNetworkOk returns a tuple with the DestinationNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNetwork

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetDestinationNetwork(v string)`

SetDestinationNetwork sets DestinationNetwork field to given value.

### HasDestinationNetwork

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasDestinationNetwork() bool`

HasDestinationNetwork returns a boolean if a field has been set.

### GetDestinationPort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetDestinationPort() int32`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetDestinationPortOk() (*int32, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetDestinationPort(v int32)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetDestinationPortRange

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetDestinationPortRange() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetDestinationPortRangeOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetDestinationPortRange(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### GetDomainCategoryIds

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetDomainCategoryIds() []int64`

GetDomainCategoryIds returns the DomainCategoryIds field if non-nil, zero value otherwise.

### GetDomainCategoryIdsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetDomainCategoryIdsOk() (*[]int64, bool)`

GetDomainCategoryIdsOk returns a tuple with the DomainCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCategoryIds

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetDomainCategoryIds(v []int64)`

SetDomainCategoryIds sets DomainCategoryIds field to given value.

### HasDomainCategoryIds

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasDomainCategoryIds() bool`

HasDomainCategoryIds returns a boolean if a field has been set.

### GetDscp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetDscp() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscpMatch`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetDscpOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscpMatch, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetDscp(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscpMatch)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetIcmpType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.

### HasIcmpType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasIcmpType() bool`

HasIcmpType returns a boolean if a field has been set.

### GetIpProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetIpProtocol() string`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetIpProtocolOk() (*string, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetIpProtocol(v string)`

SetIpProtocol sets IpProtocol field to given value.

### HasIpProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasIpProtocol() bool`

HasIpProtocol returns a boolean if a field has been set.

### GetProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSourceNetwork

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetSourceNetwork() string`

GetSourceNetwork returns the SourceNetwork field if non-nil, zero value otherwise.

### GetSourceNetworkOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetSourceNetworkOk() (*string, bool)`

GetSourceNetworkOk returns a tuple with the SourceNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetwork

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetSourceNetwork(v string)`

SetSourceNetwork sets SourceNetwork field to given value.

### HasSourceNetwork

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasSourceNetwork() bool`

HasSourceNetwork returns a boolean if a field has been set.

### GetSourcePort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetSourcePort() int32`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetSourcePortOk() (*int32, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetSourcePort(v int32)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetSourcePortRange

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetSourcePortRange() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) GetSourcePortRangeOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) SetSourcePortRange(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerTrafficPolicySecurityRulesetsInnerRulesInnerMatch) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


