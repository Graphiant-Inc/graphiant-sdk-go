# ManaV2OspfInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**ManaV2OspfAuthentication**](ManaV2OspfAuthentication.md) |  | [optional] 
**Bfd** | Pointer to [**ManaV2BfdInstance**](ManaV2BfdInstance.md) |  | [optional] 
**BfdNeighbors** | Pointer to [**[]ManaV2BfdNeighbor**](ManaV2BfdNeighbor.md) |  | [optional] 
**Cost** | Pointer to **int32** |  | [optional] 
**DeadInterval** | Pointer to **int32** |  | [optional] 
**DeadIntervalValue** | Pointer to **int32** |  | [optional] 
**DrPriority** | Pointer to **int32** |  | [optional] 
**HelloInterval** | Pointer to **int32** |  | [optional] 
**HelloIntervalValue** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IfIndex** | Pointer to **int32** |  | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**MaxTransmissionUnit** | Pointer to **int32** |  | [optional] 
**MtuIgnore** | Pointer to **bool** |  | [optional] 
**PrefixSid** | Pointer to **int32** |  | [optional] 
**RetransmitInterval** | Pointer to **int32** |  | [optional] 
**RetransmitIntervalValue** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewManaV2OspfInterface

`func NewManaV2OspfInterface() *ManaV2OspfInterface`

NewManaV2OspfInterface instantiates a new ManaV2OspfInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2OspfInterfaceWithDefaults

`func NewManaV2OspfInterfaceWithDefaults() *ManaV2OspfInterface`

NewManaV2OspfInterfaceWithDefaults instantiates a new ManaV2OspfInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *ManaV2OspfInterface) GetAuthentication() ManaV2OspfAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ManaV2OspfInterface) GetAuthenticationOk() (*ManaV2OspfAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ManaV2OspfInterface) SetAuthentication(v ManaV2OspfAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *ManaV2OspfInterface) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetBfd

`func (o *ManaV2OspfInterface) GetBfd() ManaV2BfdInstance`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *ManaV2OspfInterface) GetBfdOk() (*ManaV2BfdInstance, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *ManaV2OspfInterface) SetBfd(v ManaV2BfdInstance)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *ManaV2OspfInterface) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetBfdNeighbors

`func (o *ManaV2OspfInterface) GetBfdNeighbors() []ManaV2BfdNeighbor`

GetBfdNeighbors returns the BfdNeighbors field if non-nil, zero value otherwise.

### GetBfdNeighborsOk

`func (o *ManaV2OspfInterface) GetBfdNeighborsOk() (*[]ManaV2BfdNeighbor, bool)`

GetBfdNeighborsOk returns a tuple with the BfdNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdNeighbors

`func (o *ManaV2OspfInterface) SetBfdNeighbors(v []ManaV2BfdNeighbor)`

SetBfdNeighbors sets BfdNeighbors field to given value.

### HasBfdNeighbors

`func (o *ManaV2OspfInterface) HasBfdNeighbors() bool`

HasBfdNeighbors returns a boolean if a field has been set.

### GetCost

`func (o *ManaV2OspfInterface) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ManaV2OspfInterface) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ManaV2OspfInterface) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ManaV2OspfInterface) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDeadInterval

`func (o *ManaV2OspfInterface) GetDeadInterval() int32`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *ManaV2OspfInterface) GetDeadIntervalOk() (*int32, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *ManaV2OspfInterface) SetDeadInterval(v int32)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *ManaV2OspfInterface) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetDeadIntervalValue

`func (o *ManaV2OspfInterface) GetDeadIntervalValue() int32`

GetDeadIntervalValue returns the DeadIntervalValue field if non-nil, zero value otherwise.

### GetDeadIntervalValueOk

`func (o *ManaV2OspfInterface) GetDeadIntervalValueOk() (*int32, bool)`

GetDeadIntervalValueOk returns a tuple with the DeadIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadIntervalValue

`func (o *ManaV2OspfInterface) SetDeadIntervalValue(v int32)`

SetDeadIntervalValue sets DeadIntervalValue field to given value.

### HasDeadIntervalValue

`func (o *ManaV2OspfInterface) HasDeadIntervalValue() bool`

HasDeadIntervalValue returns a boolean if a field has been set.

### GetDrPriority

`func (o *ManaV2OspfInterface) GetDrPriority() int32`

GetDrPriority returns the DrPriority field if non-nil, zero value otherwise.

### GetDrPriorityOk

`func (o *ManaV2OspfInterface) GetDrPriorityOk() (*int32, bool)`

GetDrPriorityOk returns a tuple with the DrPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrPriority

`func (o *ManaV2OspfInterface) SetDrPriority(v int32)`

SetDrPriority sets DrPriority field to given value.

### HasDrPriority

`func (o *ManaV2OspfInterface) HasDrPriority() bool`

HasDrPriority returns a boolean if a field has been set.

### GetHelloInterval

`func (o *ManaV2OspfInterface) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *ManaV2OspfInterface) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *ManaV2OspfInterface) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *ManaV2OspfInterface) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetHelloIntervalValue

`func (o *ManaV2OspfInterface) GetHelloIntervalValue() int32`

GetHelloIntervalValue returns the HelloIntervalValue field if non-nil, zero value otherwise.

### GetHelloIntervalValueOk

`func (o *ManaV2OspfInterface) GetHelloIntervalValueOk() (*int32, bool)`

GetHelloIntervalValueOk returns a tuple with the HelloIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloIntervalValue

`func (o *ManaV2OspfInterface) SetHelloIntervalValue(v int32)`

SetHelloIntervalValue sets HelloIntervalValue field to given value.

### HasHelloIntervalValue

`func (o *ManaV2OspfInterface) HasHelloIntervalValue() bool`

HasHelloIntervalValue returns a boolean if a field has been set.

### GetId

`func (o *ManaV2OspfInterface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2OspfInterface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2OspfInterface) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ManaV2OspfInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIfIndex

`func (o *ManaV2OspfInterface) GetIfIndex() int32`

GetIfIndex returns the IfIndex field if non-nil, zero value otherwise.

### GetIfIndexOk

`func (o *ManaV2OspfInterface) GetIfIndexOk() (*int32, bool)`

GetIfIndexOk returns a tuple with the IfIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfIndex

`func (o *ManaV2OspfInterface) SetIfIndex(v int32)`

SetIfIndex sets IfIndex field to given value.

### HasIfIndex

`func (o *ManaV2OspfInterface) HasIfIndex() bool`

HasIfIndex returns a boolean if a field has been set.

### GetInterface

`func (o *ManaV2OspfInterface) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *ManaV2OspfInterface) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *ManaV2OspfInterface) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *ManaV2OspfInterface) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *ManaV2OspfInterface) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *ManaV2OspfInterface) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *ManaV2OspfInterface) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *ManaV2OspfInterface) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetMtuIgnore

`func (o *ManaV2OspfInterface) GetMtuIgnore() bool`

GetMtuIgnore returns the MtuIgnore field if non-nil, zero value otherwise.

### GetMtuIgnoreOk

`func (o *ManaV2OspfInterface) GetMtuIgnoreOk() (*bool, bool)`

GetMtuIgnoreOk returns a tuple with the MtuIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtuIgnore

`func (o *ManaV2OspfInterface) SetMtuIgnore(v bool)`

SetMtuIgnore sets MtuIgnore field to given value.

### HasMtuIgnore

`func (o *ManaV2OspfInterface) HasMtuIgnore() bool`

HasMtuIgnore returns a boolean if a field has been set.

### GetPrefixSid

`func (o *ManaV2OspfInterface) GetPrefixSid() int32`

GetPrefixSid returns the PrefixSid field if non-nil, zero value otherwise.

### GetPrefixSidOk

`func (o *ManaV2OspfInterface) GetPrefixSidOk() (*int32, bool)`

GetPrefixSidOk returns a tuple with the PrefixSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSid

`func (o *ManaV2OspfInterface) SetPrefixSid(v int32)`

SetPrefixSid sets PrefixSid field to given value.

### HasPrefixSid

`func (o *ManaV2OspfInterface) HasPrefixSid() bool`

HasPrefixSid returns a boolean if a field has been set.

### GetRetransmitInterval

`func (o *ManaV2OspfInterface) GetRetransmitInterval() int32`

GetRetransmitInterval returns the RetransmitInterval field if non-nil, zero value otherwise.

### GetRetransmitIntervalOk

`func (o *ManaV2OspfInterface) GetRetransmitIntervalOk() (*int32, bool)`

GetRetransmitIntervalOk returns a tuple with the RetransmitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitInterval

`func (o *ManaV2OspfInterface) SetRetransmitInterval(v int32)`

SetRetransmitInterval sets RetransmitInterval field to given value.

### HasRetransmitInterval

`func (o *ManaV2OspfInterface) HasRetransmitInterval() bool`

HasRetransmitInterval returns a boolean if a field has been set.

### GetRetransmitIntervalValue

`func (o *ManaV2OspfInterface) GetRetransmitIntervalValue() int32`

GetRetransmitIntervalValue returns the RetransmitIntervalValue field if non-nil, zero value otherwise.

### GetRetransmitIntervalValueOk

`func (o *ManaV2OspfInterface) GetRetransmitIntervalValueOk() (*int32, bool)`

GetRetransmitIntervalValueOk returns a tuple with the RetransmitIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitIntervalValue

`func (o *ManaV2OspfInterface) SetRetransmitIntervalValue(v int32)`

SetRetransmitIntervalValue sets RetransmitIntervalValue field to given value.

### HasRetransmitIntervalValue

`func (o *ManaV2OspfInterface) HasRetransmitIntervalValue() bool`

HasRetransmitIntervalValue returns a boolean if a field has been set.

### GetType

`func (o *ManaV2OspfInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManaV2OspfInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManaV2OspfInterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManaV2OspfInterface) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


