# ManaV2PublicVifSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | id of this Public VIF | [optional] 
**ServiceName** | Pointer to **string** | name of this Public VIF | [optional] 
**UpdatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**UserName** | Pointer to **string** | creator of this Public VIF | [optional] 

## Methods

### NewManaV2PublicVifSummary

`func NewManaV2PublicVifSummary() *ManaV2PublicVifSummary`

NewManaV2PublicVifSummary instantiates a new ManaV2PublicVifSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PublicVifSummaryWithDefaults

`func NewManaV2PublicVifSummaryWithDefaults() *ManaV2PublicVifSummary`

NewManaV2PublicVifSummaryWithDefaults instantiates a new ManaV2PublicVifSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManaV2PublicVifSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2PublicVifSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2PublicVifSummary) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ManaV2PublicVifSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServiceName

`func (o *ManaV2PublicVifSummary) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ManaV2PublicVifSummary) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ManaV2PublicVifSummary) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ManaV2PublicVifSummary) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ManaV2PublicVifSummary) GetUpdatedAt() GoogleProtobufTimestamp`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManaV2PublicVifSummary) GetUpdatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManaV2PublicVifSummary) SetUpdatedAt(v GoogleProtobufTimestamp)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ManaV2PublicVifSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserName

`func (o *ManaV2PublicVifSummary) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ManaV2PublicVifSummary) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ManaV2PublicVifSummary) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ManaV2PublicVifSummary) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


