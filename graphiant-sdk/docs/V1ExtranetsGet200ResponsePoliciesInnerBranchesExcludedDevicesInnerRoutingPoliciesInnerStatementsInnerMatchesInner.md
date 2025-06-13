# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Community** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**PrefixSet** | Pointer to **string** |  | [optional] 
**ProtocolRouteType** | Pointer to **string** |  | [optional] 
**RouteTag** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTagEntry**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTagEntry.md) |  | [optional] 
**Seq** | Pointer to **int32** |  | [optional] 
**SourceInterface** | Pointer to **string** |  | [optional] 
**SourceProtocol** | Pointer to **string** |  | [optional] 
**StalePurge** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInnerWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetCommunity() []string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetCommunityOk() (*[]string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) SetCommunity(v []string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrefixSet

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetPrefixSet() string`

GetPrefixSet returns the PrefixSet field if non-nil, zero value otherwise.

### GetPrefixSetOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetPrefixSetOk() (*string, bool)`

GetPrefixSetOk returns a tuple with the PrefixSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSet

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) SetPrefixSet(v string)`

SetPrefixSet sets PrefixSet field to given value.

### HasPrefixSet

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) HasPrefixSet() bool`

HasPrefixSet returns a boolean if a field has been set.

### GetProtocolRouteType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetProtocolRouteType() string`

GetProtocolRouteType returns the ProtocolRouteType field if non-nil, zero value otherwise.

### GetProtocolRouteTypeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetProtocolRouteTypeOk() (*string, bool)`

GetProtocolRouteTypeOk returns a tuple with the ProtocolRouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolRouteType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) SetProtocolRouteType(v string)`

SetProtocolRouteType sets ProtocolRouteType field to given value.

### HasProtocolRouteType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) HasProtocolRouteType() bool`

HasProtocolRouteType returns a boolean if a field has been set.

### GetRouteTag

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetRouteTag() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTagEntry`

GetRouteTag returns the RouteTag field if non-nil, zero value otherwise.

### GetRouteTagOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetRouteTagOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTagEntry, bool)`

GetRouteTagOk returns a tuple with the RouteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTag

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) SetRouteTag(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTagEntry)`

SetRouteTag sets RouteTag field to given value.

### HasRouteTag

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) HasRouteTag() bool`

HasRouteTag returns a boolean if a field has been set.

### GetSeq

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSourceInterface

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetSourceInterface() string`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetSourceInterfaceOk() (*string, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) SetSourceInterface(v string)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.

### GetSourceProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetSourceProtocol() string`

GetSourceProtocol returns the SourceProtocol field if non-nil, zero value otherwise.

### GetSourceProtocolOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetSourceProtocolOk() (*string, bool)`

GetSourceProtocolOk returns a tuple with the SourceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) SetSourceProtocol(v string)`

SetSourceProtocol sets SourceProtocol field to given value.

### HasSourceProtocol

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) HasSourceProtocol() bool`

HasSourceProtocol returns a boolean if a field has been set.

### GetStalePurge

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetStalePurge() bool`

GetStalePurge returns the StalePurge field if non-nil, zero value otherwise.

### GetStalePurgeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) GetStalePurgeOk() (*bool, bool)`

GetStalePurgeOk returns a tuple with the StalePurge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalePurge

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) SetStalePurge(v bool)`

SetStalePurge sets StalePurge field to given value.

### HasStalePurge

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerRoutingPoliciesInnerStatementsInnerMatchesInner) HasStalePurge() bool`

HasStalePurge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


