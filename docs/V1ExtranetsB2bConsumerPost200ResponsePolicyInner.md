# V1ExtranetsB2bConsumerPost200ResponsePolicyInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerLanSegment** | Pointer to **int64** |  | [optional] 
**InboundSecurityRules** | Pointer to [**[]V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner**](V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner.md) |  | [optional] 
**OutboundSecurityRules** | Pointer to [**[]V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner**](V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner.md) |  | [optional] 
**RestrictedPrefixes** | Pointer to **[]string** |  | [optional] 
**Rules** | Pointer to [**[]V1ExtranetsB2bConsumerPostRequestPolicyInnerRulesInner**](V1ExtranetsB2bConsumerPostRequestPolicyInnerRulesInner.md) |  | [optional] 
**TrafficRules** | Pointer to [**[]V1ExtranetsB2bConsumerPost200ResponsePolicyInnerTrafficRulesInner**](V1ExtranetsB2bConsumerPost200ResponsePolicyInnerTrafficRulesInner.md) |  | [optional] 

## Methods

### NewV1ExtranetsB2bConsumerPost200ResponsePolicyInner

`func NewV1ExtranetsB2bConsumerPost200ResponsePolicyInner() *V1ExtranetsB2bConsumerPost200ResponsePolicyInner`

NewV1ExtranetsB2bConsumerPost200ResponsePolicyInner instantiates a new V1ExtranetsB2bConsumerPost200ResponsePolicyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsB2bConsumerPost200ResponsePolicyInnerWithDefaults

`func NewV1ExtranetsB2bConsumerPost200ResponsePolicyInnerWithDefaults() *V1ExtranetsB2bConsumerPost200ResponsePolicyInner`

NewV1ExtranetsB2bConsumerPost200ResponsePolicyInnerWithDefaults instantiates a new V1ExtranetsB2bConsumerPost200ResponsePolicyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerLanSegment

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetConsumerLanSegment() int64`

GetConsumerLanSegment returns the ConsumerLanSegment field if non-nil, zero value otherwise.

### GetConsumerLanSegmentOk

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetConsumerLanSegmentOk() (*int64, bool)`

GetConsumerLanSegmentOk returns a tuple with the ConsumerLanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerLanSegment

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) SetConsumerLanSegment(v int64)`

SetConsumerLanSegment sets ConsumerLanSegment field to given value.

### HasConsumerLanSegment

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) HasConsumerLanSegment() bool`

HasConsumerLanSegment returns a boolean if a field has been set.

### GetInboundSecurityRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetInboundSecurityRules() []V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner`

GetInboundSecurityRules returns the InboundSecurityRules field if non-nil, zero value otherwise.

### GetInboundSecurityRulesOk

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetInboundSecurityRulesOk() (*[]V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner, bool)`

GetInboundSecurityRulesOk returns a tuple with the InboundSecurityRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSecurityRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) SetInboundSecurityRules(v []V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner)`

SetInboundSecurityRules sets InboundSecurityRules field to given value.

### HasInboundSecurityRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) HasInboundSecurityRules() bool`

HasInboundSecurityRules returns a boolean if a field has been set.

### GetOutboundSecurityRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetOutboundSecurityRules() []V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner`

GetOutboundSecurityRules returns the OutboundSecurityRules field if non-nil, zero value otherwise.

### GetOutboundSecurityRulesOk

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetOutboundSecurityRulesOk() (*[]V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner, bool)`

GetOutboundSecurityRulesOk returns a tuple with the OutboundSecurityRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSecurityRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) SetOutboundSecurityRules(v []V1ExtranetsB2bConsumerPost200ResponsePolicyInnerInboundSecurityRulesInner)`

SetOutboundSecurityRules sets OutboundSecurityRules field to given value.

### HasOutboundSecurityRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) HasOutboundSecurityRules() bool`

HasOutboundSecurityRules returns a boolean if a field has been set.

### GetRestrictedPrefixes

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetRestrictedPrefixes() []string`

GetRestrictedPrefixes returns the RestrictedPrefixes field if non-nil, zero value otherwise.

### GetRestrictedPrefixesOk

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetRestrictedPrefixesOk() (*[]string, bool)`

GetRestrictedPrefixesOk returns a tuple with the RestrictedPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedPrefixes

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) SetRestrictedPrefixes(v []string)`

SetRestrictedPrefixes sets RestrictedPrefixes field to given value.

### HasRestrictedPrefixes

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) HasRestrictedPrefixes() bool`

HasRestrictedPrefixes returns a boolean if a field has been set.

### GetRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetRules() []V1ExtranetsB2bConsumerPostRequestPolicyInnerRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetRulesOk() (*[]V1ExtranetsB2bConsumerPostRequestPolicyInnerRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) SetRules(v []V1ExtranetsB2bConsumerPostRequestPolicyInnerRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTrafficRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetTrafficRules() []V1ExtranetsB2bConsumerPost200ResponsePolicyInnerTrafficRulesInner`

GetTrafficRules returns the TrafficRules field if non-nil, zero value otherwise.

### GetTrafficRulesOk

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) GetTrafficRulesOk() (*[]V1ExtranetsB2bConsumerPost200ResponsePolicyInnerTrafficRulesInner, bool)`

GetTrafficRulesOk returns a tuple with the TrafficRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) SetTrafficRules(v []V1ExtranetsB2bConsumerPost200ResponsePolicyInnerTrafficRulesInner)`

SetTrafficRules sets TrafficRules field to given value.

### HasTrafficRules

`func (o *V1ExtranetsB2bConsumerPost200ResponsePolicyInner) HasTrafficRules() bool`

HasTrafficRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


