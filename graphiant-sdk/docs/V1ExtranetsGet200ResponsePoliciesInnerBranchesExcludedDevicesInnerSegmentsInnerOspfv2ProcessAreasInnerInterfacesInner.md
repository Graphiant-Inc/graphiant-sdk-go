# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bfd** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd.md) |  | [optional] 
**BfdNeighbors** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor.md) |  | [optional] 
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

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInnerWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBfd

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetBfd() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetBfdOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetBfd(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetBfdNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetBfdNeighbors() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor`

GetBfdNeighbors returns the BfdNeighbors field if non-nil, zero value otherwise.

### GetBfdNeighborsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetBfdNeighborsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor, bool)`

GetBfdNeighborsOk returns a tuple with the BfdNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetBfdNeighbors(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor)`

SetBfdNeighbors sets BfdNeighbors field to given value.

### HasBfdNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasBfdNeighbors() bool`

HasBfdNeighbors returns a boolean if a field has been set.

### GetCost

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDeadInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetDeadInterval() int32`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetDeadIntervalOk() (*int32, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetDeadInterval(v int32)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetDeadIntervalValue

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetDeadIntervalValue() int32`

GetDeadIntervalValue returns the DeadIntervalValue field if non-nil, zero value otherwise.

### GetDeadIntervalValueOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetDeadIntervalValueOk() (*int32, bool)`

GetDeadIntervalValueOk returns a tuple with the DeadIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadIntervalValue

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetDeadIntervalValue(v int32)`

SetDeadIntervalValue sets DeadIntervalValue field to given value.

### HasDeadIntervalValue

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasDeadIntervalValue() bool`

HasDeadIntervalValue returns a boolean if a field has been set.

### GetDrPriority

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetDrPriority() int32`

GetDrPriority returns the DrPriority field if non-nil, zero value otherwise.

### GetDrPriorityOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetDrPriorityOk() (*int32, bool)`

GetDrPriorityOk returns a tuple with the DrPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrPriority

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetDrPriority(v int32)`

SetDrPriority sets DrPriority field to given value.

### HasDrPriority

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasDrPriority() bool`

HasDrPriority returns a boolean if a field has been set.

### GetHelloInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetHelloIntervalValue

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetHelloIntervalValue() int32`

GetHelloIntervalValue returns the HelloIntervalValue field if non-nil, zero value otherwise.

### GetHelloIntervalValueOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetHelloIntervalValueOk() (*int32, bool)`

GetHelloIntervalValueOk returns a tuple with the HelloIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloIntervalValue

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetHelloIntervalValue(v int32)`

SetHelloIntervalValue sets HelloIntervalValue field to given value.

### HasHelloIntervalValue

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasHelloIntervalValue() bool`

HasHelloIntervalValue returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIfIndex

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetIfIndex() int32`

GetIfIndex returns the IfIndex field if non-nil, zero value otherwise.

### GetIfIndexOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetIfIndexOk() (*int32, bool)`

GetIfIndexOk returns a tuple with the IfIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfIndex

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetIfIndex(v int32)`

SetIfIndex sets IfIndex field to given value.

### HasIfIndex

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasIfIndex() bool`

HasIfIndex returns a boolean if a field has been set.

### GetInterface

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetMaxTransmissionUnit

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetMaxTransmissionUnit() int32`

GetMaxTransmissionUnit returns the MaxTransmissionUnit field if non-nil, zero value otherwise.

### GetMaxTransmissionUnitOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetMaxTransmissionUnitOk() (*int32, bool)`

GetMaxTransmissionUnitOk returns a tuple with the MaxTransmissionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransmissionUnit

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetMaxTransmissionUnit(v int32)`

SetMaxTransmissionUnit sets MaxTransmissionUnit field to given value.

### HasMaxTransmissionUnit

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasMaxTransmissionUnit() bool`

HasMaxTransmissionUnit returns a boolean if a field has been set.

### GetMtuIgnore

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetMtuIgnore() bool`

GetMtuIgnore returns the MtuIgnore field if non-nil, zero value otherwise.

### GetMtuIgnoreOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetMtuIgnoreOk() (*bool, bool)`

GetMtuIgnoreOk returns a tuple with the MtuIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtuIgnore

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetMtuIgnore(v bool)`

SetMtuIgnore sets MtuIgnore field to given value.

### HasMtuIgnore

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasMtuIgnore() bool`

HasMtuIgnore returns a boolean if a field has been set.

### GetPrefixSid

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetPrefixSid() int32`

GetPrefixSid returns the PrefixSid field if non-nil, zero value otherwise.

### GetPrefixSidOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetPrefixSidOk() (*int32, bool)`

GetPrefixSidOk returns a tuple with the PrefixSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSid

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetPrefixSid(v int32)`

SetPrefixSid sets PrefixSid field to given value.

### HasPrefixSid

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasPrefixSid() bool`

HasPrefixSid returns a boolean if a field has been set.

### GetRetransmitInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetRetransmitInterval() int32`

GetRetransmitInterval returns the RetransmitInterval field if non-nil, zero value otherwise.

### GetRetransmitIntervalOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetRetransmitIntervalOk() (*int32, bool)`

GetRetransmitIntervalOk returns a tuple with the RetransmitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetRetransmitInterval(v int32)`

SetRetransmitInterval sets RetransmitInterval field to given value.

### HasRetransmitInterval

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasRetransmitInterval() bool`

HasRetransmitInterval returns a boolean if a field has been set.

### GetRetransmitIntervalValue

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetRetransmitIntervalValue() int32`

GetRetransmitIntervalValue returns the RetransmitIntervalValue field if non-nil, zero value otherwise.

### GetRetransmitIntervalValueOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetRetransmitIntervalValueOk() (*int32, bool)`

GetRetransmitIntervalValueOk returns a tuple with the RetransmitIntervalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitIntervalValue

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetRetransmitIntervalValue(v int32)`

SetRetransmitIntervalValue sets RetransmitIntervalValue field to given value.

### HasRetransmitIntervalValue

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasRetransmitIntervalValue() bool`

HasRetransmitIntervalValue returns a boolean if a field has been set.

### GetType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


