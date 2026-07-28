# ManaV2ExtranetServiceProducerCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEmails** | Pointer to **[]string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**MatchId** | Pointer to **int64** |  | [optional] 
**MatchedServices** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 

## Methods

### NewManaV2ExtranetServiceProducerCustomer

`func NewManaV2ExtranetServiceProducerCustomer() *ManaV2ExtranetServiceProducerCustomer`

NewManaV2ExtranetServiceProducerCustomer instantiates a new ManaV2ExtranetServiceProducerCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2ExtranetServiceProducerCustomerWithDefaults

`func NewManaV2ExtranetServiceProducerCustomerWithDefaults() *ManaV2ExtranetServiceProducerCustomer`

NewManaV2ExtranetServiceProducerCustomerWithDefaults instantiates a new ManaV2ExtranetServiceProducerCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEmails

`func (o *ManaV2ExtranetServiceProducerCustomer) GetAdminEmails() []string`

GetAdminEmails returns the AdminEmails field if non-nil, zero value otherwise.

### GetAdminEmailsOk

`func (o *ManaV2ExtranetServiceProducerCustomer) GetAdminEmailsOk() (*[]string, bool)`

GetAdminEmailsOk returns a tuple with the AdminEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmails

`func (o *ManaV2ExtranetServiceProducerCustomer) SetAdminEmails(v []string)`

SetAdminEmails sets AdminEmails field to given value.

### HasAdminEmails

`func (o *ManaV2ExtranetServiceProducerCustomer) HasAdminEmails() bool`

HasAdminEmails returns a boolean if a field has been set.

### GetCustomerId

`func (o *ManaV2ExtranetServiceProducerCustomer) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ManaV2ExtranetServiceProducerCustomer) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ManaV2ExtranetServiceProducerCustomer) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ManaV2ExtranetServiceProducerCustomer) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetMatchId

`func (o *ManaV2ExtranetServiceProducerCustomer) GetMatchId() int64`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *ManaV2ExtranetServiceProducerCustomer) GetMatchIdOk() (*int64, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *ManaV2ExtranetServiceProducerCustomer) SetMatchId(v int64)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *ManaV2ExtranetServiceProducerCustomer) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetMatchedServices

`func (o *ManaV2ExtranetServiceProducerCustomer) GetMatchedServices() int32`

GetMatchedServices returns the MatchedServices field if non-nil, zero value otherwise.

### GetMatchedServicesOk

`func (o *ManaV2ExtranetServiceProducerCustomer) GetMatchedServicesOk() (*int32, bool)`

GetMatchedServicesOk returns a tuple with the MatchedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedServices

`func (o *ManaV2ExtranetServiceProducerCustomer) SetMatchedServices(v int32)`

SetMatchedServices sets MatchedServices field to given value.

### HasMatchedServices

`func (o *ManaV2ExtranetServiceProducerCustomer) HasMatchedServices() bool`

HasMatchedServices returns a boolean if a field has been set.

### GetName

`func (o *ManaV2ExtranetServiceProducerCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2ExtranetServiceProducerCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2ExtranetServiceProducerCustomer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2ExtranetServiceProducerCustomer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ManaV2ExtranetServiceProducerCustomer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManaV2ExtranetServiceProducerCustomer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManaV2ExtranetServiceProducerCustomer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManaV2ExtranetServiceProducerCustomer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ManaV2ExtranetServiceProducerCustomer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManaV2ExtranetServiceProducerCustomer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManaV2ExtranetServiceProducerCustomer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManaV2ExtranetServiceProducerCustomer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ManaV2ExtranetServiceProducerCustomer) GetUpdatedAt() GoogleProtobufTimestamp`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManaV2ExtranetServiceProducerCustomer) GetUpdatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManaV2ExtranetServiceProducerCustomer) SetUpdatedAt(v GoogleProtobufTimestamp)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ManaV2ExtranetServiceProducerCustomer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


