# V1EdgesSummaryGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgesSummary** | Pointer to [**[]V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner**](V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner.md) |  | [optional] 
**IsNewCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1EdgesSummaryGet200Response

`func NewV1EdgesSummaryGet200Response() *V1EdgesSummaryGet200Response`

NewV1EdgesSummaryGet200Response instantiates a new V1EdgesSummaryGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EdgesSummaryGet200ResponseWithDefaults

`func NewV1EdgesSummaryGet200ResponseWithDefaults() *V1EdgesSummaryGet200Response`

NewV1EdgesSummaryGet200ResponseWithDefaults instantiates a new V1EdgesSummaryGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgesSummary

`func (o *V1EdgesSummaryGet200Response) GetEdgesSummary() []V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner`

GetEdgesSummary returns the EdgesSummary field if non-nil, zero value otherwise.

### GetEdgesSummaryOk

`func (o *V1EdgesSummaryGet200Response) GetEdgesSummaryOk() (*[]V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner, bool)`

GetEdgesSummaryOk returns a tuple with the EdgesSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgesSummary

`func (o *V1EdgesSummaryGet200Response) SetEdgesSummary(v []V1EdgesHardwareAssignedGet200ResponseEdgesSummaryInner)`

SetEdgesSummary sets EdgesSummary field to given value.

### HasEdgesSummary

`func (o *V1EdgesSummaryGet200Response) HasEdgesSummary() bool`

HasEdgesSummary returns a boolean if a field has been set.

### GetIsNewCount

`func (o *V1EdgesSummaryGet200Response) GetIsNewCount() int32`

GetIsNewCount returns the IsNewCount field if non-nil, zero value otherwise.

### GetIsNewCountOk

`func (o *V1EdgesSummaryGet200Response) GetIsNewCountOk() (*int32, bool)`

GetIsNewCountOk returns a tuple with the IsNewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewCount

`func (o *V1EdgesSummaryGet200Response) SetIsNewCount(v int32)`

SetIsNewCount sets IsNewCount field to given value.

### HasIsNewCount

`func (o *V1EdgesSummaryGet200Response) HasIsNewCount() bool`

HasIsNewCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


