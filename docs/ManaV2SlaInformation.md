# ManaV2SlaInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCircuit** | Pointer to **string** | Backup circuit for the SLA information | [optional] 
**Class** | Pointer to **string** | Class for the SLA information | [optional] 
**PrimaryCircuit** | Pointer to **string** | Primary circuit for the SLA information | [optional] 

## Methods

### NewManaV2SlaInformation

`func NewManaV2SlaInformation() *ManaV2SlaInformation`

NewManaV2SlaInformation instantiates a new ManaV2SlaInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2SlaInformationWithDefaults

`func NewManaV2SlaInformationWithDefaults() *ManaV2SlaInformation`

NewManaV2SlaInformationWithDefaults instantiates a new ManaV2SlaInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupCircuit

`func (o *ManaV2SlaInformation) GetBackupCircuit() string`

GetBackupCircuit returns the BackupCircuit field if non-nil, zero value otherwise.

### GetBackupCircuitOk

`func (o *ManaV2SlaInformation) GetBackupCircuitOk() (*string, bool)`

GetBackupCircuitOk returns a tuple with the BackupCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCircuit

`func (o *ManaV2SlaInformation) SetBackupCircuit(v string)`

SetBackupCircuit sets BackupCircuit field to given value.

### HasBackupCircuit

`func (o *ManaV2SlaInformation) HasBackupCircuit() bool`

HasBackupCircuit returns a boolean if a field has been set.

### GetClass

`func (o *ManaV2SlaInformation) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ManaV2SlaInformation) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ManaV2SlaInformation) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *ManaV2SlaInformation) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetPrimaryCircuit

`func (o *ManaV2SlaInformation) GetPrimaryCircuit() string`

GetPrimaryCircuit returns the PrimaryCircuit field if non-nil, zero value otherwise.

### GetPrimaryCircuitOk

`func (o *ManaV2SlaInformation) GetPrimaryCircuitOk() (*string, bool)`

GetPrimaryCircuitOk returns a tuple with the PrimaryCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCircuit

`func (o *ManaV2SlaInformation) SetPrimaryCircuit(v string)`

SetPrimaryCircuit sets PrimaryCircuit field to given value.

### HasPrimaryCircuit

`func (o *ManaV2SlaInformation) HasPrimaryCircuit() bool`

HasPrimaryCircuit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


