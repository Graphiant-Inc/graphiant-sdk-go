# AlertserviceIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**AlertserviceIntegrationDetails**](AlertserviceIntegrationDetails.md) |  | [optional] 
**EnterpriseId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**LastScanned** | Pointer to **int64** | last scanned time for integrations (zendesk) | [optional] 
**NickName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAlertserviceIntegration

`func NewAlertserviceIntegration() *AlertserviceIntegration`

NewAlertserviceIntegration instantiates a new AlertserviceIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertserviceIntegrationWithDefaults

`func NewAlertserviceIntegrationWithDefaults() *AlertserviceIntegration`

NewAlertserviceIntegrationWithDefaults instantiates a new AlertserviceIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *AlertserviceIntegration) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AlertserviceIntegration) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AlertserviceIntegration) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AlertserviceIntegration) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *AlertserviceIntegration) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *AlertserviceIntegration) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *AlertserviceIntegration) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *AlertserviceIntegration) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDetails

`func (o *AlertserviceIntegration) GetDetails() AlertserviceIntegrationDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AlertserviceIntegration) GetDetailsOk() (*AlertserviceIntegrationDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AlertserviceIntegration) SetDetails(v AlertserviceIntegrationDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AlertserviceIntegration) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *AlertserviceIntegration) GetEnterpriseId() string`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *AlertserviceIntegration) GetEnterpriseIdOk() (*string, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *AlertserviceIntegration) SetEnterpriseId(v string)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *AlertserviceIntegration) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetId

`func (o *AlertserviceIntegration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertserviceIntegration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertserviceIntegration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AlertserviceIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *AlertserviceIntegration) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AlertserviceIntegration) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AlertserviceIntegration) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AlertserviceIntegration) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastScanned

`func (o *AlertserviceIntegration) GetLastScanned() int64`

GetLastScanned returns the LastScanned field if non-nil, zero value otherwise.

### GetLastScannedOk

`func (o *AlertserviceIntegration) GetLastScannedOk() (*int64, bool)`

GetLastScannedOk returns a tuple with the LastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScanned

`func (o *AlertserviceIntegration) SetLastScanned(v int64)`

SetLastScanned sets LastScanned field to given value.

### HasLastScanned

`func (o *AlertserviceIntegration) HasLastScanned() bool`

HasLastScanned returns a boolean if a field has been set.

### GetNickName

`func (o *AlertserviceIntegration) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *AlertserviceIntegration) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *AlertserviceIntegration) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *AlertserviceIntegration) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### GetType

`func (o *AlertserviceIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertserviceIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertserviceIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AlertserviceIntegration) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


