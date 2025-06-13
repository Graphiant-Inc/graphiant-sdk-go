# V1ActivityLogsPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorRef** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]V1ActivityLogsPost200ResponseDetailsInner**](V1ActivityLogsPost200ResponseDetailsInner.md) |  | [optional] 
**FilterEntities** | Pointer to [**map[string]V1ActivityLogsPost200ResponseFilterEntitiesValue**](V1ActivityLogsPost200ResponseFilterEntitiesValue.md) |  | [optional] 
**FilterJobTypes** | Pointer to **[]string** |  | [optional] 
**TotalLogs** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1ActivityLogsPost200Response

`func NewV1ActivityLogsPost200Response() *V1ActivityLogsPost200Response`

NewV1ActivityLogsPost200Response instantiates a new V1ActivityLogsPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ActivityLogsPost200ResponseWithDefaults

`func NewV1ActivityLogsPost200ResponseWithDefaults() *V1ActivityLogsPost200Response`

NewV1ActivityLogsPost200ResponseWithDefaults instantiates a new V1ActivityLogsPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorRef

`func (o *V1ActivityLogsPost200Response) GetCursorRef() string`

GetCursorRef returns the CursorRef field if non-nil, zero value otherwise.

### GetCursorRefOk

`func (o *V1ActivityLogsPost200Response) GetCursorRefOk() (*string, bool)`

GetCursorRefOk returns a tuple with the CursorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorRef

`func (o *V1ActivityLogsPost200Response) SetCursorRef(v string)`

SetCursorRef sets CursorRef field to given value.

### HasCursorRef

`func (o *V1ActivityLogsPost200Response) HasCursorRef() bool`

HasCursorRef returns a boolean if a field has been set.

### GetDetails

`func (o *V1ActivityLogsPost200Response) GetDetails() []V1ActivityLogsPost200ResponseDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *V1ActivityLogsPost200Response) GetDetailsOk() (*[]V1ActivityLogsPost200ResponseDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *V1ActivityLogsPost200Response) SetDetails(v []V1ActivityLogsPost200ResponseDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *V1ActivityLogsPost200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFilterEntities

`func (o *V1ActivityLogsPost200Response) GetFilterEntities() map[string]V1ActivityLogsPost200ResponseFilterEntitiesValue`

GetFilterEntities returns the FilterEntities field if non-nil, zero value otherwise.

### GetFilterEntitiesOk

`func (o *V1ActivityLogsPost200Response) GetFilterEntitiesOk() (*map[string]V1ActivityLogsPost200ResponseFilterEntitiesValue, bool)`

GetFilterEntitiesOk returns a tuple with the FilterEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterEntities

`func (o *V1ActivityLogsPost200Response) SetFilterEntities(v map[string]V1ActivityLogsPost200ResponseFilterEntitiesValue)`

SetFilterEntities sets FilterEntities field to given value.

### HasFilterEntities

`func (o *V1ActivityLogsPost200Response) HasFilterEntities() bool`

HasFilterEntities returns a boolean if a field has been set.

### GetFilterJobTypes

`func (o *V1ActivityLogsPost200Response) GetFilterJobTypes() []string`

GetFilterJobTypes returns the FilterJobTypes field if non-nil, zero value otherwise.

### GetFilterJobTypesOk

`func (o *V1ActivityLogsPost200Response) GetFilterJobTypesOk() (*[]string, bool)`

GetFilterJobTypesOk returns a tuple with the FilterJobTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterJobTypes

`func (o *V1ActivityLogsPost200Response) SetFilterJobTypes(v []string)`

SetFilterJobTypes sets FilterJobTypes field to given value.

### HasFilterJobTypes

`func (o *V1ActivityLogsPost200Response) HasFilterJobTypes() bool`

HasFilterJobTypes returns a boolean if a field has been set.

### GetTotalLogs

`func (o *V1ActivityLogsPost200Response) GetTotalLogs() int64`

GetTotalLogs returns the TotalLogs field if non-nil, zero value otherwise.

### GetTotalLogsOk

`func (o *V1ActivityLogsPost200Response) GetTotalLogsOk() (*int64, bool)`

GetTotalLogsOk returns a tuple with the TotalLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogs

`func (o *V1ActivityLogsPost200Response) SetTotalLogs(v int64)`

SetTotalLogs sets TotalLogs field to given value.

### HasTotalLogs

`func (o *V1ActivityLogsPost200Response) HasTotalLogs() bool`

HasTotalLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


