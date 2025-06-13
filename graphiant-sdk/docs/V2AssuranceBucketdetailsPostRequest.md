# V2AssuranceBucketdetailsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketId** | Pointer to **string** |  | [optional] 
**TimeWindow** | Pointer to [**V2NotificationlistPostRequestTimeWindow**](V2NotificationlistPostRequestTimeWindow.md) |  | [optional] 
**UnclassifiedOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewV2AssuranceBucketdetailsPostRequest

`func NewV2AssuranceBucketdetailsPostRequest() *V2AssuranceBucketdetailsPostRequest`

NewV2AssuranceBucketdetailsPostRequest instantiates a new V2AssuranceBucketdetailsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceBucketdetailsPostRequestWithDefaults

`func NewV2AssuranceBucketdetailsPostRequestWithDefaults() *V2AssuranceBucketdetailsPostRequest`

NewV2AssuranceBucketdetailsPostRequestWithDefaults instantiates a new V2AssuranceBucketdetailsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketId

`func (o *V2AssuranceBucketdetailsPostRequest) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *V2AssuranceBucketdetailsPostRequest) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *V2AssuranceBucketdetailsPostRequest) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *V2AssuranceBucketdetailsPostRequest) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetTimeWindow

`func (o *V2AssuranceBucketdetailsPostRequest) GetTimeWindow() V2NotificationlistPostRequestTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2AssuranceBucketdetailsPostRequest) GetTimeWindowOk() (*V2NotificationlistPostRequestTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2AssuranceBucketdetailsPostRequest) SetTimeWindow(v V2NotificationlistPostRequestTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *V2AssuranceBucketdetailsPostRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetUnclassifiedOnly

`func (o *V2AssuranceBucketdetailsPostRequest) GetUnclassifiedOnly() bool`

GetUnclassifiedOnly returns the UnclassifiedOnly field if non-nil, zero value otherwise.

### GetUnclassifiedOnlyOk

`func (o *V2AssuranceBucketdetailsPostRequest) GetUnclassifiedOnlyOk() (*bool, bool)`

GetUnclassifiedOnlyOk returns a tuple with the UnclassifiedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclassifiedOnly

`func (o *V2AssuranceBucketdetailsPostRequest) SetUnclassifiedOnly(v bool)`

SetUnclassifiedOnly sets UnclassifiedOnly field to given value.

### HasUnclassifiedOnly

`func (o *V2AssuranceBucketdetailsPostRequest) HasUnclassifiedOnly() bool`

HasUnclassifiedOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


