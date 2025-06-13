# V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaId** | Pointer to **string** |  | [optional] 
**Bfd** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd.md) |  | [optional] 
**BfdNeighbors** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Interfaces** | Pointer to [**[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerWithDefaults

`func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner`

NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.

### HasAreaId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) HasAreaId() bool`

HasAreaId returns a boolean if a field has been set.

### GetBfd

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetBfd() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetBfdOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) SetBfd(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfd)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetBfdNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetBfdNeighbors() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor`

GetBfdNeighbors returns the BfdNeighbors field if non-nil, zero value otherwise.

### GetBfdNeighborsOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetBfdNeighborsOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor, bool)`

GetBfdNeighborsOk returns a tuple with the BfdNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) SetBfdNeighbors(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerCircuitsInnerBgpNeighborsInnerBfdNeighbor)`

SetBfdNeighbors sets BfdNeighbors field to given value.

### HasBfdNeighbors

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) HasBfdNeighbors() bool`

HasBfdNeighbors returns a boolean if a field has been set.

### GetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaces

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetInterfaces() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetInterfacesOk() (*[]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) SetInterfaces(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerSegmentsInnerOspfv2ProcessAreasInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


