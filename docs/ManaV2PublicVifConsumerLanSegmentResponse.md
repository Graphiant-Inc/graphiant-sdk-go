# ManaV2PublicVifConsumerLanSegmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPrefixes** | Pointer to **[]string** |  | [optional] 
**OutboundSecurityRules** | Pointer to [**[]ManaV2SecurityPolicyRule**](ManaV2SecurityPolicyRule.md) |  | [optional] 
**ServiceLanSegment** | Pointer to **int64** | LAN segment ID for the service | [optional] 
**TrafficRules** | Pointer to [**[]ManaV2TrafficPolicyRule**](ManaV2TrafficPolicyRule.md) |  | [optional] 

## Methods

### NewManaV2PublicVifConsumerLanSegmentResponse

`func NewManaV2PublicVifConsumerLanSegmentResponse() *ManaV2PublicVifConsumerLanSegmentResponse`

NewManaV2PublicVifConsumerLanSegmentResponse instantiates a new ManaV2PublicVifConsumerLanSegmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PublicVifConsumerLanSegmentResponseWithDefaults

`func NewManaV2PublicVifConsumerLanSegmentResponseWithDefaults() *ManaV2PublicVifConsumerLanSegmentResponse`

NewManaV2PublicVifConsumerLanSegmentResponseWithDefaults instantiates a new ManaV2PublicVifConsumerLanSegmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPrefixes

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) GetAllowedPrefixes() []string`

GetAllowedPrefixes returns the AllowedPrefixes field if non-nil, zero value otherwise.

### GetAllowedPrefixesOk

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) GetAllowedPrefixesOk() (*[]string, bool)`

GetAllowedPrefixesOk returns a tuple with the AllowedPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPrefixes

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) SetAllowedPrefixes(v []string)`

SetAllowedPrefixes sets AllowedPrefixes field to given value.

### HasAllowedPrefixes

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) HasAllowedPrefixes() bool`

HasAllowedPrefixes returns a boolean if a field has been set.

### GetOutboundSecurityRules

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) GetOutboundSecurityRules() []ManaV2SecurityPolicyRule`

GetOutboundSecurityRules returns the OutboundSecurityRules field if non-nil, zero value otherwise.

### GetOutboundSecurityRulesOk

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) GetOutboundSecurityRulesOk() (*[]ManaV2SecurityPolicyRule, bool)`

GetOutboundSecurityRulesOk returns a tuple with the OutboundSecurityRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSecurityRules

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) SetOutboundSecurityRules(v []ManaV2SecurityPolicyRule)`

SetOutboundSecurityRules sets OutboundSecurityRules field to given value.

### HasOutboundSecurityRules

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) HasOutboundSecurityRules() bool`

HasOutboundSecurityRules returns a boolean if a field has been set.

### GetServiceLanSegment

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) GetServiceLanSegment() int64`

GetServiceLanSegment returns the ServiceLanSegment field if non-nil, zero value otherwise.

### GetServiceLanSegmentOk

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) GetServiceLanSegmentOk() (*int64, bool)`

GetServiceLanSegmentOk returns a tuple with the ServiceLanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLanSegment

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) SetServiceLanSegment(v int64)`

SetServiceLanSegment sets ServiceLanSegment field to given value.

### HasServiceLanSegment

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) HasServiceLanSegment() bool`

HasServiceLanSegment returns a boolean if a field has been set.

### GetTrafficRules

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) GetTrafficRules() []ManaV2TrafficPolicyRule`

GetTrafficRules returns the TrafficRules field if non-nil, zero value otherwise.

### GetTrafficRulesOk

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) GetTrafficRulesOk() (*[]ManaV2TrafficPolicyRule, bool)`

GetTrafficRulesOk returns a tuple with the TrafficRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRules

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) SetTrafficRules(v []ManaV2TrafficPolicyRule)`

SetTrafficRules sets TrafficRules field to given value.

### HasTrafficRules

`func (o *ManaV2PublicVifConsumerLanSegmentResponse) HasTrafficRules() bool`

HasTrafficRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


