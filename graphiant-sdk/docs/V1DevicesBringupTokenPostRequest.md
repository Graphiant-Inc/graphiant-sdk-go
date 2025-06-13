# V1DevicesBringupTokenPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**ValidTillTs** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**ValiditySec** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1DevicesBringupTokenPostRequest

`func NewV1DevicesBringupTokenPostRequest() *V1DevicesBringupTokenPostRequest`

NewV1DevicesBringupTokenPostRequest instantiates a new V1DevicesBringupTokenPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesBringupTokenPostRequestWithDefaults

`func NewV1DevicesBringupTokenPostRequestWithDefaults() *V1DevicesBringupTokenPostRequest`

NewV1DevicesBringupTokenPostRequestWithDefaults instantiates a new V1DevicesBringupTokenPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *V1DevicesBringupTokenPostRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *V1DevicesBringupTokenPostRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *V1DevicesBringupTokenPostRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *V1DevicesBringupTokenPostRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetValidTillTs

`func (o *V1DevicesBringupTokenPostRequest) GetValidTillTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetValidTillTs returns the ValidTillTs field if non-nil, zero value otherwise.

### GetValidTillTsOk

`func (o *V1DevicesBringupTokenPostRequest) GetValidTillTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetValidTillTsOk returns a tuple with the ValidTillTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTillTs

`func (o *V1DevicesBringupTokenPostRequest) SetValidTillTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetValidTillTs sets ValidTillTs field to given value.

### HasValidTillTs

`func (o *V1DevicesBringupTokenPostRequest) HasValidTillTs() bool`

HasValidTillTs returns a boolean if a field has been set.

### GetValiditySec

`func (o *V1DevicesBringupTokenPostRequest) GetValiditySec() int64`

GetValiditySec returns the ValiditySec field if non-nil, zero value otherwise.

### GetValiditySecOk

`func (o *V1DevicesBringupTokenPostRequest) GetValiditySecOk() (*int64, bool)`

GetValiditySecOk returns a tuple with the ValiditySec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValiditySec

`func (o *V1DevicesBringupTokenPostRequest) SetValiditySec(v int64)`

SetValiditySec sets ValiditySec field to given value.

### HasValiditySec

`func (o *V1DevicesBringupTokenPostRequest) HasValiditySec() bool`

HasValiditySec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


