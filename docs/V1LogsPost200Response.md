# V1LogsPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorRef** | Pointer to **string** |  | [optional] 
**Histogram** | Pointer to [**[]V1AuditLogsPost200ResponseHistogramInner**](V1AuditLogsPost200ResponseHistogramInner.md) |  | [optional] 
**Logs** | Pointer to [**[]V1LogsPost200ResponseLogsInner**](V1LogsPost200ResponseLogsInner.md) |  | [optional] 
**TotalLogs** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1LogsPost200Response

`func NewV1LogsPost200Response() *V1LogsPost200Response`

NewV1LogsPost200Response instantiates a new V1LogsPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LogsPost200ResponseWithDefaults

`func NewV1LogsPost200ResponseWithDefaults() *V1LogsPost200Response`

NewV1LogsPost200ResponseWithDefaults instantiates a new V1LogsPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorRef

`func (o *V1LogsPost200Response) GetCursorRef() string`

GetCursorRef returns the CursorRef field if non-nil, zero value otherwise.

### GetCursorRefOk

`func (o *V1LogsPost200Response) GetCursorRefOk() (*string, bool)`

GetCursorRefOk returns a tuple with the CursorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorRef

`func (o *V1LogsPost200Response) SetCursorRef(v string)`

SetCursorRef sets CursorRef field to given value.

### HasCursorRef

`func (o *V1LogsPost200Response) HasCursorRef() bool`

HasCursorRef returns a boolean if a field has been set.

### GetHistogram

`func (o *V1LogsPost200Response) GetHistogram() []V1AuditLogsPost200ResponseHistogramInner`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *V1LogsPost200Response) GetHistogramOk() (*[]V1AuditLogsPost200ResponseHistogramInner, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *V1LogsPost200Response) SetHistogram(v []V1AuditLogsPost200ResponseHistogramInner)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *V1LogsPost200Response) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetLogs

`func (o *V1LogsPost200Response) GetLogs() []V1LogsPost200ResponseLogsInner`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *V1LogsPost200Response) GetLogsOk() (*[]V1LogsPost200ResponseLogsInner, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *V1LogsPost200Response) SetLogs(v []V1LogsPost200ResponseLogsInner)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *V1LogsPost200Response) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetTotalLogs

`func (o *V1LogsPost200Response) GetTotalLogs() int64`

GetTotalLogs returns the TotalLogs field if non-nil, zero value otherwise.

### GetTotalLogsOk

`func (o *V1LogsPost200Response) GetTotalLogsOk() (*int64, bool)`

GetTotalLogsOk returns a tuple with the TotalLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogs

`func (o *V1LogsPost200Response) SetTotalLogs(v int64)`

SetTotalLogs sets TotalLogs field to given value.

### HasTotalLogs

`func (o *V1LogsPost200Response) HasTotalLogs() bool`

HasTotalLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


