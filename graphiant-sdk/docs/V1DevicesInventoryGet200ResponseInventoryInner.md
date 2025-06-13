# V1DevicesInventoryGet200ResponseInventoryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedOn** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**CreatedOn** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**DeviceModel** | Pointer to **string** |  | [optional] 
**DeviceSerial** | Pointer to **string** |  | [optional] 
**EkCert** | Pointer to **string** |  | [optional] 
**EnterpriseId** | Pointer to **int64** |  | [optional] 
**EnterpriseName** | Pointer to **string** |  | [optional] 
**GekPub** | Pointer to **string** |  | [optional] 
**IsCore** | Pointer to **bool** |  | [optional] 
**IsNew** | Pointer to **bool** |  | [optional] 
**IsRequested** | Pointer to **bool** |  | [optional] 
**ParentEnterpriseId** | Pointer to **int64** |  | [optional] 
**ParentEnterpriseName** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**UseOauth** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DevicesInventoryGet200ResponseInventoryInner

`func NewV1DevicesInventoryGet200ResponseInventoryInner() *V1DevicesInventoryGet200ResponseInventoryInner`

NewV1DevicesInventoryGet200ResponseInventoryInner instantiates a new V1DevicesInventoryGet200ResponseInventoryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesInventoryGet200ResponseInventoryInnerWithDefaults

`func NewV1DevicesInventoryGet200ResponseInventoryInnerWithDefaults() *V1DevicesInventoryGet200ResponseInventoryInner`

NewV1DevicesInventoryGet200ResponseInventoryInnerWithDefaults instantiates a new V1DevicesInventoryGet200ResponseInventoryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedOn

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetAssignedOn() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetAssignedOn returns the AssignedOn field if non-nil, zero value otherwise.

### GetAssignedOnOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetAssignedOnOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetAssignedOnOk returns a tuple with the AssignedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedOn

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetAssignedOn(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetAssignedOn sets AssignedOn field to given value.

### HasAssignedOn

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasAssignedOn() bool`

HasAssignedOn returns a boolean if a field has been set.

### GetCreatedOn

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetCreatedOn() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetCreatedOnOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetCreatedOn(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDeviceModel

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetEkCert

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetEkCert() string`

GetEkCert returns the EkCert field if non-nil, zero value otherwise.

### GetEkCertOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetEkCertOk() (*string, bool)`

GetEkCertOk returns a tuple with the EkCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEkCert

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetEkCert(v string)`

SetEkCert sets EkCert field to given value.

### HasEkCert

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasEkCert() bool`

HasEkCert returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetEnterpriseId() int64`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetEnterpriseIdOk() (*int64, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetEnterpriseId(v int64)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetEnterpriseName

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetEnterpriseName() string`

GetEnterpriseName returns the EnterpriseName field if non-nil, zero value otherwise.

### GetEnterpriseNameOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetEnterpriseNameOk() (*string, bool)`

GetEnterpriseNameOk returns a tuple with the EnterpriseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseName

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetEnterpriseName(v string)`

SetEnterpriseName sets EnterpriseName field to given value.

### HasEnterpriseName

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasEnterpriseName() bool`

HasEnterpriseName returns a boolean if a field has been set.

### GetGekPub

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetGekPub() string`

GetGekPub returns the GekPub field if non-nil, zero value otherwise.

### GetGekPubOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetGekPubOk() (*string, bool)`

GetGekPubOk returns a tuple with the GekPub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGekPub

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetGekPub(v string)`

SetGekPub sets GekPub field to given value.

### HasGekPub

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasGekPub() bool`

HasGekPub returns a boolean if a field has been set.

### GetIsCore

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetIsCore() bool`

GetIsCore returns the IsCore field if non-nil, zero value otherwise.

### GetIsCoreOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetIsCoreOk() (*bool, bool)`

GetIsCoreOk returns a tuple with the IsCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCore

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetIsCore(v bool)`

SetIsCore sets IsCore field to given value.

### HasIsCore

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasIsCore() bool`

HasIsCore returns a boolean if a field has been set.

### GetIsNew

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.

### HasIsNew

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasIsNew() bool`

HasIsNew returns a boolean if a field has been set.

### GetIsRequested

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetIsRequested() bool`

GetIsRequested returns the IsRequested field if non-nil, zero value otherwise.

### GetIsRequestedOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetIsRequestedOk() (*bool, bool)`

GetIsRequestedOk returns a tuple with the IsRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequested

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetIsRequested(v bool)`

SetIsRequested sets IsRequested field to given value.

### HasIsRequested

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasIsRequested() bool`

HasIsRequested returns a boolean if a field has been set.

### GetParentEnterpriseId

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetParentEnterpriseId() int64`

GetParentEnterpriseId returns the ParentEnterpriseId field if non-nil, zero value otherwise.

### GetParentEnterpriseIdOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetParentEnterpriseIdOk() (*int64, bool)`

GetParentEnterpriseIdOk returns a tuple with the ParentEnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEnterpriseId

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetParentEnterpriseId(v int64)`

SetParentEnterpriseId sets ParentEnterpriseId field to given value.

### HasParentEnterpriseId

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasParentEnterpriseId() bool`

HasParentEnterpriseId returns a boolean if a field has been set.

### GetParentEnterpriseName

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetParentEnterpriseName() string`

GetParentEnterpriseName returns the ParentEnterpriseName field if non-nil, zero value otherwise.

### GetParentEnterpriseNameOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetParentEnterpriseNameOk() (*string, bool)`

GetParentEnterpriseNameOk returns a tuple with the ParentEnterpriseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEnterpriseName

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetParentEnterpriseName(v string)`

SetParentEnterpriseName sets ParentEnterpriseName field to given value.

### HasParentEnterpriseName

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasParentEnterpriseName() bool`

HasParentEnterpriseName returns a boolean if a field has been set.

### GetRole

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUseOauth

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetUseOauth() bool`

GetUseOauth returns the UseOauth field if non-nil, zero value otherwise.

### GetUseOauthOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetUseOauthOk() (*bool, bool)`

GetUseOauthOk returns a tuple with the UseOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOauth

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetUseOauth(v bool)`

SetUseOauth sets UseOauth field to given value.

### HasUseOauth

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasUseOauth() bool`

HasUseOauth returns a boolean if a field has been set.

### GetUuid

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V1DevicesInventoryGet200ResponseInventoryInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


