# V2IntegrationPostRequestIntegrationBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**Details** | Pointer to [**V2IntegrationPostRequestIntegrationBodyDetails**](V2IntegrationPostRequestIntegrationBodyDetails.md) |  | [optional] 
**Enterprise** | Pointer to **int64** |  | [optional] 
**IntegrationType** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**NickName** | Pointer to **string** |  | [optional] 

## Methods

### NewV2IntegrationPostRequestIntegrationBody

`func NewV2IntegrationPostRequestIntegrationBody() *V2IntegrationPostRequestIntegrationBody`

NewV2IntegrationPostRequestIntegrationBody instantiates a new V2IntegrationPostRequestIntegrationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2IntegrationPostRequestIntegrationBodyWithDefaults

`func NewV2IntegrationPostRequestIntegrationBodyWithDefaults() *V2IntegrationPostRequestIntegrationBody`

NewV2IntegrationPostRequestIntegrationBodyWithDefaults instantiates a new V2IntegrationPostRequestIntegrationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *V2IntegrationPostRequestIntegrationBody) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *V2IntegrationPostRequestIntegrationBody) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *V2IntegrationPostRequestIntegrationBody) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *V2IntegrationPostRequestIntegrationBody) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *V2IntegrationPostRequestIntegrationBody) GetCreatedOn() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *V2IntegrationPostRequestIntegrationBody) GetCreatedOnOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *V2IntegrationPostRequestIntegrationBody) SetCreatedOn(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *V2IntegrationPostRequestIntegrationBody) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDetails

`func (o *V2IntegrationPostRequestIntegrationBody) GetDetails() V2IntegrationPostRequestIntegrationBodyDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *V2IntegrationPostRequestIntegrationBody) GetDetailsOk() (*V2IntegrationPostRequestIntegrationBodyDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *V2IntegrationPostRequestIntegrationBody) SetDetails(v V2IntegrationPostRequestIntegrationBodyDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *V2IntegrationPostRequestIntegrationBody) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEnterprise

`func (o *V2IntegrationPostRequestIntegrationBody) GetEnterprise() int64`

GetEnterprise returns the Enterprise field if non-nil, zero value otherwise.

### GetEnterpriseOk

`func (o *V2IntegrationPostRequestIntegrationBody) GetEnterpriseOk() (*int64, bool)`

GetEnterpriseOk returns a tuple with the Enterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprise

`func (o *V2IntegrationPostRequestIntegrationBody) SetEnterprise(v int64)`

SetEnterprise sets Enterprise field to given value.

### HasEnterprise

`func (o *V2IntegrationPostRequestIntegrationBody) HasEnterprise() bool`

HasEnterprise returns a boolean if a field has been set.

### GetIntegrationType

`func (o *V2IntegrationPostRequestIntegrationBody) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *V2IntegrationPostRequestIntegrationBody) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *V2IntegrationPostRequestIntegrationBody) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *V2IntegrationPostRequestIntegrationBody) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetIsActive

`func (o *V2IntegrationPostRequestIntegrationBody) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *V2IntegrationPostRequestIntegrationBody) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *V2IntegrationPostRequestIntegrationBody) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *V2IntegrationPostRequestIntegrationBody) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetNickName

`func (o *V2IntegrationPostRequestIntegrationBody) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *V2IntegrationPostRequestIntegrationBody) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *V2IntegrationPostRequestIntegrationBody) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *V2IntegrationPostRequestIntegrationBody) HasNickName() bool`

HasNickName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


