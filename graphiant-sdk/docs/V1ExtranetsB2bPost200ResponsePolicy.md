# V1ExtranetsB2bPost200ResponsePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsName** | Pointer to **string** |  | [optional] 
**InboundSecurityRules** | Pointer to [**[]V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner**](V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner.md) |  | [optional] 
**Policy** | Pointer to [**V1ExtranetsB2bPostRequestPolicy**](V1ExtranetsB2bPostRequestPolicy.md) |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**TrafficRules** | Pointer to [**[]V1ExtranetsB2bConsumerPost200ResponsePolicyInnerTrafficRulesInner**](V1ExtranetsB2bConsumerPost200ResponsePolicyInnerTrafficRulesInner.md) |  | [optional] 

## Methods

### NewV1ExtranetsB2bPost200ResponsePolicy

`func NewV1ExtranetsB2bPost200ResponsePolicy() *V1ExtranetsB2bPost200ResponsePolicy`

NewV1ExtranetsB2bPost200ResponsePolicy instantiates a new V1ExtranetsB2bPost200ResponsePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsB2bPost200ResponsePolicyWithDefaults

`func NewV1ExtranetsB2bPost200ResponsePolicyWithDefaults() *V1ExtranetsB2bPost200ResponsePolicy`

NewV1ExtranetsB2bPost200ResponsePolicyWithDefaults instantiates a new V1ExtranetsB2bPost200ResponsePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *V1ExtranetsB2bPost200ResponsePolicy) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *V1ExtranetsB2bPost200ResponsePolicy) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *V1ExtranetsB2bPost200ResponsePolicy) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *V1ExtranetsB2bPost200ResponsePolicy) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetInboundSecurityRules

`func (o *V1ExtranetsB2bPost200ResponsePolicy) GetInboundSecurityRules() []V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner`

GetInboundSecurityRules returns the InboundSecurityRules field if non-nil, zero value otherwise.

### GetInboundSecurityRulesOk

`func (o *V1ExtranetsB2bPost200ResponsePolicy) GetInboundSecurityRulesOk() (*[]V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner, bool)`

GetInboundSecurityRulesOk returns a tuple with the InboundSecurityRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSecurityRules

`func (o *V1ExtranetsB2bPost200ResponsePolicy) SetInboundSecurityRules(v []V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner)`

SetInboundSecurityRules sets InboundSecurityRules field to given value.

### HasInboundSecurityRules

`func (o *V1ExtranetsB2bPost200ResponsePolicy) HasInboundSecurityRules() bool`

HasInboundSecurityRules returns a boolean if a field has been set.

### GetPolicy

`func (o *V1ExtranetsB2bPost200ResponsePolicy) GetPolicy() V1ExtranetsB2bPostRequestPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *V1ExtranetsB2bPost200ResponsePolicy) GetPolicyOk() (*V1ExtranetsB2bPostRequestPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *V1ExtranetsB2bPost200ResponsePolicy) SetPolicy(v V1ExtranetsB2bPostRequestPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *V1ExtranetsB2bPost200ResponsePolicy) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetServiceName

`func (o *V1ExtranetsB2bPost200ResponsePolicy) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1ExtranetsB2bPost200ResponsePolicy) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1ExtranetsB2bPost200ResponsePolicy) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *V1ExtranetsB2bPost200ResponsePolicy) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetTrafficRules

`func (o *V1ExtranetsB2bPost200ResponsePolicy) GetTrafficRules() []V1ExtranetsB2bConsumerPost200ResponsePolicyInnerTrafficRulesInner`

GetTrafficRules returns the TrafficRules field if non-nil, zero value otherwise.

### GetTrafficRulesOk

`func (o *V1ExtranetsB2bPost200ResponsePolicy) GetTrafficRulesOk() (*[]V1ExtranetsB2bConsumerPost200ResponsePolicyInnerTrafficRulesInner, bool)`

GetTrafficRulesOk returns a tuple with the TrafficRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRules

`func (o *V1ExtranetsB2bPost200ResponsePolicy) SetTrafficRules(v []V1ExtranetsB2bConsumerPost200ResponsePolicyInnerTrafficRulesInner)`

SetTrafficRules sets TrafficRules field to given value.

### HasTrafficRules

`func (o *V1ExtranetsB2bPost200ResponsePolicy) HasTrafficRules() bool`

HasTrafficRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


