# V1DiagnosticSpeedtestPost200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgPingTime** | Pointer to **float64** |  | [optional] 
**DateTime** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**DownloadSpeed** | Pointer to **float64** |  | [optional] 
**Isp** | Pointer to **string** |  | [optional] 
**MaxPingTime** | Pointer to **float64** |  | [optional] 
**MinPingTime** | Pointer to **float64** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**ServerDetails** | Pointer to [**V1DiagnosticSpeedtestServersGet200ResponseServerInner**](V1DiagnosticSpeedtestServersGet200ResponseServerInner.md) |  | [optional] 
**UploadSpeed** | Pointer to **float64** |  | [optional] 

## Methods

### NewV1DiagnosticSpeedtestPost200ResponseResult

`func NewV1DiagnosticSpeedtestPost200ResponseResult() *V1DiagnosticSpeedtestPost200ResponseResult`

NewV1DiagnosticSpeedtestPost200ResponseResult instantiates a new V1DiagnosticSpeedtestPost200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DiagnosticSpeedtestPost200ResponseResultWithDefaults

`func NewV1DiagnosticSpeedtestPost200ResponseResultWithDefaults() *V1DiagnosticSpeedtestPost200ResponseResult`

NewV1DiagnosticSpeedtestPost200ResponseResultWithDefaults instantiates a new V1DiagnosticSpeedtestPost200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPingTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetAvgPingTime() float64`

GetAvgPingTime returns the AvgPingTime field if non-nil, zero value otherwise.

### GetAvgPingTimeOk

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetAvgPingTimeOk() (*float64, bool)`

GetAvgPingTimeOk returns a tuple with the AvgPingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPingTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) SetAvgPingTime(v float64)`

SetAvgPingTime sets AvgPingTime field to given value.

### HasAvgPingTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) HasAvgPingTime() bool`

HasAvgPingTime returns a boolean if a field has been set.

### GetDateTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetDateTime() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetDateTimeOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) SetDateTime(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetDownloadSpeed

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetDownloadSpeed() float64`

GetDownloadSpeed returns the DownloadSpeed field if non-nil, zero value otherwise.

### GetDownloadSpeedOk

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetDownloadSpeedOk() (*float64, bool)`

GetDownloadSpeedOk returns a tuple with the DownloadSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadSpeed

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) SetDownloadSpeed(v float64)`

SetDownloadSpeed sets DownloadSpeed field to given value.

### HasDownloadSpeed

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) HasDownloadSpeed() bool`

HasDownloadSpeed returns a boolean if a field has been set.

### GetIsp

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetMaxPingTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetMaxPingTime() float64`

GetMaxPingTime returns the MaxPingTime field if non-nil, zero value otherwise.

### GetMaxPingTimeOk

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetMaxPingTimeOk() (*float64, bool)`

GetMaxPingTimeOk returns a tuple with the MaxPingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPingTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) SetMaxPingTime(v float64)`

SetMaxPingTime sets MaxPingTime field to given value.

### HasMaxPingTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) HasMaxPingTime() bool`

HasMaxPingTime returns a boolean if a field has been set.

### GetMinPingTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetMinPingTime() float64`

GetMinPingTime returns the MinPingTime field if non-nil, zero value otherwise.

### GetMinPingTimeOk

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetMinPingTimeOk() (*float64, bool)`

GetMinPingTimeOk returns a tuple with the MinPingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPingTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) SetMinPingTime(v float64)`

SetMinPingTime sets MinPingTime field to given value.

### HasMinPingTime

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) HasMinPingTime() bool`

HasMinPingTime returns a boolean if a field has been set.

### GetResult

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetServerDetails

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetServerDetails() V1DiagnosticSpeedtestServersGet200ResponseServerInner`

GetServerDetails returns the ServerDetails field if non-nil, zero value otherwise.

### GetServerDetailsOk

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetServerDetailsOk() (*V1DiagnosticSpeedtestServersGet200ResponseServerInner, bool)`

GetServerDetailsOk returns a tuple with the ServerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDetails

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) SetServerDetails(v V1DiagnosticSpeedtestServersGet200ResponseServerInner)`

SetServerDetails sets ServerDetails field to given value.

### HasServerDetails

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) HasServerDetails() bool`

HasServerDetails returns a boolean if a field has been set.

### GetUploadSpeed

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetUploadSpeed() float64`

GetUploadSpeed returns the UploadSpeed field if non-nil, zero value otherwise.

### GetUploadSpeedOk

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) GetUploadSpeedOk() (*float64, bool)`

GetUploadSpeedOk returns a tuple with the UploadSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSpeed

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) SetUploadSpeed(v float64)`

SetUploadSpeed sets UploadSpeed field to given value.

### HasUploadSpeed

`func (o *V1DiagnosticSpeedtestPost200ResponseResult) HasUploadSpeed() bool`

HasUploadSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


