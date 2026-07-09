# V2AssuranceAiAdoptionSummaryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeWindow** | [**AssuranceTimeWindow**](AssuranceTimeWindow.md) |  | 
**UserListSize** | **int64** | list size of user list in where widget (required) | 

## Methods

### NewV2AssuranceAiAdoptionSummaryPostRequest

`func NewV2AssuranceAiAdoptionSummaryPostRequest(timeWindow AssuranceTimeWindow, userListSize int64, ) *V2AssuranceAiAdoptionSummaryPostRequest`

NewV2AssuranceAiAdoptionSummaryPostRequest instantiates a new V2AssuranceAiAdoptionSummaryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssuranceAiAdoptionSummaryPostRequestWithDefaults

`func NewV2AssuranceAiAdoptionSummaryPostRequestWithDefaults() *V2AssuranceAiAdoptionSummaryPostRequest`

NewV2AssuranceAiAdoptionSummaryPostRequestWithDefaults instantiates a new V2AssuranceAiAdoptionSummaryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeWindow

`func (o *V2AssuranceAiAdoptionSummaryPostRequest) GetTimeWindow() AssuranceTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *V2AssuranceAiAdoptionSummaryPostRequest) GetTimeWindowOk() (*AssuranceTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *V2AssuranceAiAdoptionSummaryPostRequest) SetTimeWindow(v AssuranceTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.


### GetUserListSize

`func (o *V2AssuranceAiAdoptionSummaryPostRequest) GetUserListSize() int64`

GetUserListSize returns the UserListSize field if non-nil, zero value otherwise.

### GetUserListSizeOk

`func (o *V2AssuranceAiAdoptionSummaryPostRequest) GetUserListSizeOk() (*int64, bool)`

GetUserListSizeOk returns a tuple with the UserListSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListSize

`func (o *V2AssuranceAiAdoptionSummaryPostRequest) SetUserListSize(v int64)`

SetUserListSize sets UserListSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


