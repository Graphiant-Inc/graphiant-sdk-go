# V1SoftwareGcsreleaseUploadNotesPostRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**[]V1SoftwareGcsreleaseUploadNotesPostRequestDetailsCategoryInner**](V1SoftwareGcsreleaseUploadNotesPostRequestDetailsCategoryInner.md) |  | [optional] 
**Major** | Pointer to **bool** |  | [optional] 
**ReleaseTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1SoftwareGcsreleaseUploadNotesPostRequestDetails

`func NewV1SoftwareGcsreleaseUploadNotesPostRequestDetails() *V1SoftwareGcsreleaseUploadNotesPostRequestDetails`

NewV1SoftwareGcsreleaseUploadNotesPostRequestDetails instantiates a new V1SoftwareGcsreleaseUploadNotesPostRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SoftwareGcsreleaseUploadNotesPostRequestDetailsWithDefaults

`func NewV1SoftwareGcsreleaseUploadNotesPostRequestDetailsWithDefaults() *V1SoftwareGcsreleaseUploadNotesPostRequestDetails`

NewV1SoftwareGcsreleaseUploadNotesPostRequestDetailsWithDefaults instantiates a new V1SoftwareGcsreleaseUploadNotesPostRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) GetCategory() []V1SoftwareGcsreleaseUploadNotesPostRequestDetailsCategoryInner`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) GetCategoryOk() (*[]V1SoftwareGcsreleaseUploadNotesPostRequestDetailsCategoryInner, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) SetCategory(v []V1SoftwareGcsreleaseUploadNotesPostRequestDetailsCategoryInner)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetMajor

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) GetMajor() bool`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) GetMajorOk() (*bool, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) SetMajor(v bool)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetReleaseTs

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) GetReleaseTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetReleaseTs returns the ReleaseTs field if non-nil, zero value otherwise.

### GetReleaseTsOk

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) GetReleaseTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetReleaseTsOk returns a tuple with the ReleaseTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTs

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) SetReleaseTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetReleaseTs sets ReleaseTs field to given value.

### HasReleaseTs

`func (o *V1SoftwareGcsreleaseUploadNotesPostRequestDetails) HasReleaseTs() bool`

HasReleaseTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


