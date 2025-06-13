# V1DeviceSnapshotGet200ResponseFirstSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**V1DeviceSnapshotGet200ResponseFirstSnapshotAuthor**](V1DeviceSnapshotGet200ResponseFirstSnapshotAuthor.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Data** | Pointer to [**V1DeviceSnapshotGet200ResponseFirstSnapshotData**](V1DeviceSnapshotGet200ResponseFirstSnapshotData.md) |  | [optional] 
**Golden** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DeviceSnapshotGet200ResponseFirstSnapshot

`func NewV1DeviceSnapshotGet200ResponseFirstSnapshot() *V1DeviceSnapshotGet200ResponseFirstSnapshot`

NewV1DeviceSnapshotGet200ResponseFirstSnapshot instantiates a new V1DeviceSnapshotGet200ResponseFirstSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeviceSnapshotGet200ResponseFirstSnapshotWithDefaults

`func NewV1DeviceSnapshotGet200ResponseFirstSnapshotWithDefaults() *V1DeviceSnapshotGet200ResponseFirstSnapshot`

NewV1DeviceSnapshotGet200ResponseFirstSnapshotWithDefaults instantiates a new V1DeviceSnapshotGet200ResponseFirstSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetAuthor() V1DeviceSnapshotGet200ResponseFirstSnapshotAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetAuthorOk() (*V1DeviceSnapshotGet200ResponseFirstSnapshotAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) SetAuthor(v V1DeviceSnapshotGet200ResponseFirstSnapshotAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCategory

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedOn

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetCreatedOn() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetCreatedOnOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) SetCreatedOn(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetData

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetData() V1DeviceSnapshotGet200ResponseFirstSnapshotData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetDataOk() (*V1DeviceSnapshotGet200ResponseFirstSnapshotData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) SetData(v V1DeviceSnapshotGet200ResponseFirstSnapshotData)`

SetData sets Data field to given value.

### HasData

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) HasData() bool`

HasData returns a boolean if a field has been set.

### GetGolden

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetGolden() bool`

GetGolden returns the Golden field if non-nil, zero value otherwise.

### GetGoldenOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetGoldenOk() (*bool, bool)`

GetGoldenOk returns a tuple with the Golden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGolden

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) SetGolden(v bool)`

SetGolden sets Golden field to given value.

### HasGolden

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) HasGolden() bool`

HasGolden returns a boolean if a field has been set.

### GetId

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DeviceSnapshotGet200ResponseFirstSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


