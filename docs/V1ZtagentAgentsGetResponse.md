# V1ZtagentAgentsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]ConcealAgent**](ConcealAgent.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1ZtagentAgentsGetResponse

`func NewV1ZtagentAgentsGetResponse() *V1ZtagentAgentsGetResponse`

NewV1ZtagentAgentsGetResponse instantiates a new V1ZtagentAgentsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ZtagentAgentsGetResponseWithDefaults

`func NewV1ZtagentAgentsGetResponseWithDefaults() *V1ZtagentAgentsGetResponse`

NewV1ZtagentAgentsGetResponseWithDefaults instantiates a new V1ZtagentAgentsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *V1ZtagentAgentsGetResponse) GetAgents() []ConcealAgent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *V1ZtagentAgentsGetResponse) GetAgentsOk() (*[]ConcealAgent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *V1ZtagentAgentsGetResponse) SetAgents(v []ConcealAgent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *V1ZtagentAgentsGetResponse) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetCount

`func (o *V1ZtagentAgentsGetResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V1ZtagentAgentsGetResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V1ZtagentAgentsGetResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V1ZtagentAgentsGetResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


