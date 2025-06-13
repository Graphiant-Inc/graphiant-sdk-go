# V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Community** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity.md) |  | [optional] 
**PrefixSet** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet.md) |  | [optional] 
**ProtocolRouteType** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType.md) |  | [optional] 
**RouteTag** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag.md) |  | [optional] 
**Seq** | Pointer to **int32** |  | [optional] 
**SourceInterface** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface.md) |  | [optional] 
**SourceProtocol** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol.md) |  | [optional] 
**Stale** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale.md) |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch

`func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch`

NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchWithDefaults

`func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchWithDefaults() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch`

NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchWithDefaults instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunity

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetCommunity() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetCommunityOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetCommunity(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetPrefixSet

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetPrefixSet() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet`

GetPrefixSet returns the PrefixSet field if non-nil, zero value otherwise.

### GetPrefixSetOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetPrefixSetOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet, bool)`

GetPrefixSetOk returns a tuple with the PrefixSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSet

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetPrefixSet(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet)`

SetPrefixSet sets PrefixSet field to given value.

### HasPrefixSet

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasPrefixSet() bool`

HasPrefixSet returns a boolean if a field has been set.

### GetProtocolRouteType

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetProtocolRouteType() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType`

GetProtocolRouteType returns the ProtocolRouteType field if non-nil, zero value otherwise.

### GetProtocolRouteTypeOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetProtocolRouteTypeOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType, bool)`

GetProtocolRouteTypeOk returns a tuple with the ProtocolRouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolRouteType

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetProtocolRouteType(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType)`

SetProtocolRouteType sets ProtocolRouteType field to given value.

### HasProtocolRouteType

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasProtocolRouteType() bool`

HasProtocolRouteType returns a boolean if a field has been set.

### GetRouteTag

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetRouteTag() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag`

GetRouteTag returns the RouteTag field if non-nil, zero value otherwise.

### GetRouteTagOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetRouteTagOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag, bool)`

GetRouteTagOk returns a tuple with the RouteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTag

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetRouteTag(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag)`

SetRouteTag sets RouteTag field to given value.

### HasRouteTag

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasRouteTag() bool`

HasRouteTag returns a boolean if a field has been set.

### GetSeq

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSourceInterface

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSourceInterface() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSourceInterfaceOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetSourceInterface(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.

### GetSourceProtocol

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSourceProtocol() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol`

GetSourceProtocol returns the SourceProtocol field if non-nil, zero value otherwise.

### GetSourceProtocolOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSourceProtocolOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol, bool)`

GetSourceProtocolOk returns a tuple with the SourceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProtocol

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetSourceProtocol(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol)`

SetSourceProtocol sets SourceProtocol field to given value.

### HasSourceProtocol

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasSourceProtocol() bool`

HasSourceProtocol returns a boolean if a field has been set.

### GetStale

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetStale() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetStaleOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetStale(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale)`

SetStale sets Stale field to given value.

### HasStale

`func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


