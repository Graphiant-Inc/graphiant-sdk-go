# IpfixEntityUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | id of the entity | [optional] 
**Name** | Pointer to **string** | name of the entity | [optional] 
**Usage** | Pointer to **float64** | usage in kbps | [optional] 

## Methods

### NewIpfixEntityUsage

`func NewIpfixEntityUsage() *IpfixEntityUsage`

NewIpfixEntityUsage instantiates a new IpfixEntityUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpfixEntityUsageWithDefaults

`func NewIpfixEntityUsageWithDefaults() *IpfixEntityUsage`

NewIpfixEntityUsageWithDefaults instantiates a new IpfixEntityUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpfixEntityUsage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpfixEntityUsage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpfixEntityUsage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IpfixEntityUsage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpfixEntityUsage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpfixEntityUsage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpfixEntityUsage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpfixEntityUsage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsage

`func (o *IpfixEntityUsage) GetUsage() float64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *IpfixEntityUsage) GetUsageOk() (*float64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *IpfixEntityUsage) SetUsage(v float64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *IpfixEntityUsage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


