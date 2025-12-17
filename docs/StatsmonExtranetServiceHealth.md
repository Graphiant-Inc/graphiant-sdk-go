# StatsmonExtranetServiceHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerName** | Pointer to **string** | the name of the customer | [optional] 
**CustomerPrefixHealth** | Pointer to [**StatsmonExtranetPrefixHealth**](StatsmonExtranetPrefixHealth.md) |  | [optional] 
**OverallHealth** | Pointer to **string** | the overall connectivity status of the service (Healthy, Impaired, Down) | [optional] 
**ProducerPrefixHealth** | Pointer to [**StatsmonExtranetPrefixHealth**](StatsmonExtranetPrefixHealth.md) |  | [optional] 

## Methods

### NewStatsmonExtranetServiceHealth

`func NewStatsmonExtranetServiceHealth() *StatsmonExtranetServiceHealth`

NewStatsmonExtranetServiceHealth instantiates a new StatsmonExtranetServiceHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsmonExtranetServiceHealthWithDefaults

`func NewStatsmonExtranetServiceHealthWithDefaults() *StatsmonExtranetServiceHealth`

NewStatsmonExtranetServiceHealthWithDefaults instantiates a new StatsmonExtranetServiceHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerName

`func (o *StatsmonExtranetServiceHealth) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *StatsmonExtranetServiceHealth) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *StatsmonExtranetServiceHealth) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *StatsmonExtranetServiceHealth) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerPrefixHealth

`func (o *StatsmonExtranetServiceHealth) GetCustomerPrefixHealth() StatsmonExtranetPrefixHealth`

GetCustomerPrefixHealth returns the CustomerPrefixHealth field if non-nil, zero value otherwise.

### GetCustomerPrefixHealthOk

`func (o *StatsmonExtranetServiceHealth) GetCustomerPrefixHealthOk() (*StatsmonExtranetPrefixHealth, bool)`

GetCustomerPrefixHealthOk returns a tuple with the CustomerPrefixHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPrefixHealth

`func (o *StatsmonExtranetServiceHealth) SetCustomerPrefixHealth(v StatsmonExtranetPrefixHealth)`

SetCustomerPrefixHealth sets CustomerPrefixHealth field to given value.

### HasCustomerPrefixHealth

`func (o *StatsmonExtranetServiceHealth) HasCustomerPrefixHealth() bool`

HasCustomerPrefixHealth returns a boolean if a field has been set.

### GetOverallHealth

`func (o *StatsmonExtranetServiceHealth) GetOverallHealth() string`

GetOverallHealth returns the OverallHealth field if non-nil, zero value otherwise.

### GetOverallHealthOk

`func (o *StatsmonExtranetServiceHealth) GetOverallHealthOk() (*string, bool)`

GetOverallHealthOk returns a tuple with the OverallHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallHealth

`func (o *StatsmonExtranetServiceHealth) SetOverallHealth(v string)`

SetOverallHealth sets OverallHealth field to given value.

### HasOverallHealth

`func (o *StatsmonExtranetServiceHealth) HasOverallHealth() bool`

HasOverallHealth returns a boolean if a field has been set.

### GetProducerPrefixHealth

`func (o *StatsmonExtranetServiceHealth) GetProducerPrefixHealth() StatsmonExtranetPrefixHealth`

GetProducerPrefixHealth returns the ProducerPrefixHealth field if non-nil, zero value otherwise.

### GetProducerPrefixHealthOk

`func (o *StatsmonExtranetServiceHealth) GetProducerPrefixHealthOk() (*StatsmonExtranetPrefixHealth, bool)`

GetProducerPrefixHealthOk returns a tuple with the ProducerPrefixHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerPrefixHealth

`func (o *StatsmonExtranetServiceHealth) SetProducerPrefixHealth(v StatsmonExtranetPrefixHealth)`

SetProducerPrefixHealth sets ProducerPrefixHealth field to given value.

### HasProducerPrefixHealth

`func (o *StatsmonExtranetServiceHealth) HasProducerPrefixHealth() bool`

HasProducerPrefixHealth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


