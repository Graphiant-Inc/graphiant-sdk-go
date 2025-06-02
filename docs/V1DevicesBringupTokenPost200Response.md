# V1DevicesBringupTokenPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**ValidTillTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1DevicesBringupTokenPost200Response

`func NewV1DevicesBringupTokenPost200Response() *V1DevicesBringupTokenPost200Response`

NewV1DevicesBringupTokenPost200Response instantiates a new V1DevicesBringupTokenPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesBringupTokenPost200ResponseWithDefaults

`func NewV1DevicesBringupTokenPost200ResponseWithDefaults() *V1DevicesBringupTokenPost200Response`

NewV1DevicesBringupTokenPost200ResponseWithDefaults instantiates a new V1DevicesBringupTokenPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *V1DevicesBringupTokenPost200Response) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *V1DevicesBringupTokenPost200Response) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *V1DevicesBringupTokenPost200Response) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *V1DevicesBringupTokenPost200Response) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetValidTillTs

`func (o *V1DevicesBringupTokenPost200Response) GetValidTillTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetValidTillTs returns the ValidTillTs field if non-nil, zero value otherwise.

### GetValidTillTsOk

`func (o *V1DevicesBringupTokenPost200Response) GetValidTillTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetValidTillTsOk returns a tuple with the ValidTillTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTillTs

`func (o *V1DevicesBringupTokenPost200Response) SetValidTillTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetValidTillTs sets ValidTillTs field to given value.

### HasValidTillTs

`func (o *V1DevicesBringupTokenPost200Response) HasValidTillTs() bool`

HasValidTillTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


