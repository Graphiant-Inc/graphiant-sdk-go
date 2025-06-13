# V1GatewaysSummaryGet200ResponseSummariesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**GatewayDeviceSummary** | Pointer to [**[]V1GatewaysSummaryGet200ResponseSummariesInnerGatewayDeviceSummaryInner**](V1GatewaysSummaryGet200ResponseSummariesInnerGatewayDeviceSummaryInner.md) |  | [optional] 
**GraphiantRegion** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SupportStatus** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1GatewaysSummaryGet200ResponseSummariesInner

`func NewV1GatewaysSummaryGet200ResponseSummariesInner() *V1GatewaysSummaryGet200ResponseSummariesInner`

NewV1GatewaysSummaryGet200ResponseSummariesInner instantiates a new V1GatewaysSummaryGet200ResponseSummariesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GatewaysSummaryGet200ResponseSummariesInnerWithDefaults

`func NewV1GatewaysSummaryGet200ResponseSummariesInnerWithDefaults() *V1GatewaysSummaryGet200ResponseSummariesInner`

NewV1GatewaysSummaryGet200ResponseSummariesInnerWithDefaults instantiates a new V1GatewaysSummaryGet200ResponseSummariesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetCreatedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetCreatedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) SetCreatedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGatewayDeviceSummary

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetGatewayDeviceSummary() []V1GatewaysSummaryGet200ResponseSummariesInnerGatewayDeviceSummaryInner`

GetGatewayDeviceSummary returns the GatewayDeviceSummary field if non-nil, zero value otherwise.

### GetGatewayDeviceSummaryOk

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetGatewayDeviceSummaryOk() (*[]V1GatewaysSummaryGet200ResponseSummariesInnerGatewayDeviceSummaryInner, bool)`

GetGatewayDeviceSummaryOk returns a tuple with the GatewayDeviceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDeviceSummary

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) SetGatewayDeviceSummary(v []V1GatewaysSummaryGet200ResponseSummariesInnerGatewayDeviceSummaryInner)`

SetGatewayDeviceSummary sets GatewayDeviceSummary field to given value.

### HasGatewayDeviceSummary

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) HasGatewayDeviceSummary() bool`

HasGatewayDeviceSummary returns a boolean if a field has been set.

### GetGraphiantRegion

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetGraphiantRegion() string`

GetGraphiantRegion returns the GraphiantRegion field if non-nil, zero value otherwise.

### GetGraphiantRegionOk

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetGraphiantRegionOk() (*string, bool)`

GetGraphiantRegionOk returns a tuple with the GraphiantRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphiantRegion

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) SetGraphiantRegion(v string)`

SetGraphiantRegion sets GraphiantRegion field to given value.

### HasGraphiantRegion

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) HasGraphiantRegion() bool`

HasGraphiantRegion returns a boolean if a field has been set.

### GetId

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpeed

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportStatus

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetSupportStatus() string`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### GetSupportStatusOk

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetSupportStatusOk() (*string, bool)`

GetSupportStatusOk returns a tuple with the SupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatus

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) SetSupportStatus(v string)`

SetSupportStatus sets SupportStatus field to given value.

### HasSupportStatus

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) HasSupportStatus() bool`

HasSupportStatus returns a boolean if a field has been set.

### GetType

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetUpdatedAt() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) GetUpdatedAtOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) SetUpdatedAt(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V1GatewaysSummaryGet200ResponseSummariesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


