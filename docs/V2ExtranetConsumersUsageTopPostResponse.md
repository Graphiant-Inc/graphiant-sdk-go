# V2ExtranetConsumersUsageTopPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopConsumers** | Pointer to [**[]IpfixEntityUsage**](IpfixEntityUsage.md) |  | [optional] 
**TotalCustomers** | Pointer to **int64** | total number of customers | [optional] 
**TotalUsage** | Pointer to **float64** | total service usage in kilo bytes | [optional] 

## Methods

### NewV2ExtranetConsumersUsageTopPostResponse

`func NewV2ExtranetConsumersUsageTopPostResponse() *V2ExtranetConsumersUsageTopPostResponse`

NewV2ExtranetConsumersUsageTopPostResponse instantiates a new V2ExtranetConsumersUsageTopPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ExtranetConsumersUsageTopPostResponseWithDefaults

`func NewV2ExtranetConsumersUsageTopPostResponseWithDefaults() *V2ExtranetConsumersUsageTopPostResponse`

NewV2ExtranetConsumersUsageTopPostResponseWithDefaults instantiates a new V2ExtranetConsumersUsageTopPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopConsumers

`func (o *V2ExtranetConsumersUsageTopPostResponse) GetTopConsumers() []IpfixEntityUsage`

GetTopConsumers returns the TopConsumers field if non-nil, zero value otherwise.

### GetTopConsumersOk

`func (o *V2ExtranetConsumersUsageTopPostResponse) GetTopConsumersOk() (*[]IpfixEntityUsage, bool)`

GetTopConsumersOk returns a tuple with the TopConsumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopConsumers

`func (o *V2ExtranetConsumersUsageTopPostResponse) SetTopConsumers(v []IpfixEntityUsage)`

SetTopConsumers sets TopConsumers field to given value.

### HasTopConsumers

`func (o *V2ExtranetConsumersUsageTopPostResponse) HasTopConsumers() bool`

HasTopConsumers returns a boolean if a field has been set.

### GetTotalCustomers

`func (o *V2ExtranetConsumersUsageTopPostResponse) GetTotalCustomers() int64`

GetTotalCustomers returns the TotalCustomers field if non-nil, zero value otherwise.

### GetTotalCustomersOk

`func (o *V2ExtranetConsumersUsageTopPostResponse) GetTotalCustomersOk() (*int64, bool)`

GetTotalCustomersOk returns a tuple with the TotalCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCustomers

`func (o *V2ExtranetConsumersUsageTopPostResponse) SetTotalCustomers(v int64)`

SetTotalCustomers sets TotalCustomers field to given value.

### HasTotalCustomers

`func (o *V2ExtranetConsumersUsageTopPostResponse) HasTotalCustomers() bool`

HasTotalCustomers returns a boolean if a field has been set.

### GetTotalUsage

`func (o *V2ExtranetConsumersUsageTopPostResponse) GetTotalUsage() float64`

GetTotalUsage returns the TotalUsage field if non-nil, zero value otherwise.

### GetTotalUsageOk

`func (o *V2ExtranetConsumersUsageTopPostResponse) GetTotalUsageOk() (*float64, bool)`

GetTotalUsageOk returns a tuple with the TotalUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsage

`func (o *V2ExtranetConsumersUsageTopPostResponse) SetTotalUsage(v float64)`

SetTotalUsage sets TotalUsage field to given value.

### HasTotalUsage

`func (o *V2ExtranetConsumersUsageTopPostResponse) HasTotalUsage() bool`

HasTotalUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


