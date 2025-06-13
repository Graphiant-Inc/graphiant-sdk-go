# V2AssuranceTopologyInventoryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketId** | Pointer to **string** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 
**TopologyType** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssuranceTopologyInventoryPostRequest

`func NewV2AssuranceTopologyInventoryPostRequest() *V2AssuranceTopologyInventoryPostRequest`

NewV2AssuranceTopologyInventoryPostRequest instantiates a new V2AssuranceTopologyInventoryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceTopologyInventoryPostRequestWithDefaults

`func NewV2AssuranceTopologyInventoryPostRequestWithDefaults() *V2AssuranceTopologyInventoryPostRequest`

NewV2AssuranceTopologyInventoryPostRequestWithDefaults instantiates a new V2AssuranceTopologyInventoryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketId

`func (o *V2AssuranceTopologyInventoryPostRequest) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *V2AssuranceTopologyInventoryPostRequest) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *V2AssuranceTopologyInventoryPostRequest) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *V2AssuranceTopologyInventoryPostRequest) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2AssuranceTopologyInventoryPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2AssuranceTopologyInventoryPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2AssuranceTopologyInventoryPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2AssuranceTopologyInventoryPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetTopologyType

`func (o *V2AssuranceTopologyInventoryPostRequest) GetTopologyType() string`

GetTopologyType returns the TopologyType field if non-nil, zero value otherwise.

### GetTopologyTypeOk

`func (o *V2AssuranceTopologyInventoryPostRequest) GetTopologyTypeOk() (*string, bool)`

GetTopologyTypeOk returns a tuple with the TopologyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyType

`func (o *V2AssuranceTopologyInventoryPostRequest) SetTopologyType(v string)`

SetTopologyType sets TopologyType field to given value.

### HasTopologyType

`func (o *V2AssuranceTopologyInventoryPostRequest) HasTopologyType() bool`

HasTopologyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


