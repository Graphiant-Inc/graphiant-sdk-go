# AssuranceUserDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSent** | Pointer to **float64** | data sent by the user (required) | [optional] 
**Managed** | Pointer to **bool** | whether the user is managed (required) | [optional] 
**SessionsDay** | Pointer to **float64** | daily sessions for the user (required) | [optional] 
**User** | Pointer to **string** | user identifier (required) | [optional] 
**Vrf** | Pointer to **string** | VRF associated with the user (required) | [optional] 

## Methods

### NewAssuranceUserDefinition

`func NewAssuranceUserDefinition() *AssuranceUserDefinition`

NewAssuranceUserDefinition instantiates a new AssuranceUserDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssuranceUserDefinitionWithDefaults

`func NewAssuranceUserDefinitionWithDefaults() *AssuranceUserDefinition`

NewAssuranceUserDefinitionWithDefaults instantiates a new AssuranceUserDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSent

`func (o *AssuranceUserDefinition) GetDataSent() float64`

GetDataSent returns the DataSent field if non-nil, zero value otherwise.

### GetDataSentOk

`func (o *AssuranceUserDefinition) GetDataSentOk() (*float64, bool)`

GetDataSentOk returns a tuple with the DataSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSent

`func (o *AssuranceUserDefinition) SetDataSent(v float64)`

SetDataSent sets DataSent field to given value.

### HasDataSent

`func (o *AssuranceUserDefinition) HasDataSent() bool`

HasDataSent returns a boolean if a field has been set.

### GetManaged

`func (o *AssuranceUserDefinition) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *AssuranceUserDefinition) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *AssuranceUserDefinition) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *AssuranceUserDefinition) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetSessionsDay

`func (o *AssuranceUserDefinition) GetSessionsDay() float64`

GetSessionsDay returns the SessionsDay field if non-nil, zero value otherwise.

### GetSessionsDayOk

`func (o *AssuranceUserDefinition) GetSessionsDayOk() (*float64, bool)`

GetSessionsDayOk returns a tuple with the SessionsDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsDay

`func (o *AssuranceUserDefinition) SetSessionsDay(v float64)`

SetSessionsDay sets SessionsDay field to given value.

### HasSessionsDay

`func (o *AssuranceUserDefinition) HasSessionsDay() bool`

HasSessionsDay returns a boolean if a field has been set.

### GetUser

`func (o *AssuranceUserDefinition) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AssuranceUserDefinition) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AssuranceUserDefinition) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AssuranceUserDefinition) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVrf

`func (o *AssuranceUserDefinition) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *AssuranceUserDefinition) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *AssuranceUserDefinition) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *AssuranceUserDefinition) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


