# CommonBillingContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractedCredits** | Pointer to **float32** | Number of credits agreed upon for the entirety of the current contract | [optional] 
**ExpirationDate** | Pointer to [**CommonBillingTimePeriod**](CommonBillingTimePeriod.md) |  | [optional] 

## Methods

### NewCommonBillingContract

`func NewCommonBillingContract() *CommonBillingContract`

NewCommonBillingContract instantiates a new CommonBillingContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonBillingContractWithDefaults

`func NewCommonBillingContractWithDefaults() *CommonBillingContract`

NewCommonBillingContractWithDefaults instantiates a new CommonBillingContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractedCredits

`func (o *CommonBillingContract) GetContractedCredits() float32`

GetContractedCredits returns the ContractedCredits field if non-nil, zero value otherwise.

### GetContractedCreditsOk

`func (o *CommonBillingContract) GetContractedCreditsOk() (*float32, bool)`

GetContractedCreditsOk returns a tuple with the ContractedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractedCredits

`func (o *CommonBillingContract) SetContractedCredits(v float32)`

SetContractedCredits sets ContractedCredits field to given value.

### HasContractedCredits

`func (o *CommonBillingContract) HasContractedCredits() bool`

HasContractedCredits returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CommonBillingContract) GetExpirationDate() CommonBillingTimePeriod`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CommonBillingContract) GetExpirationDateOk() (*CommonBillingTimePeriod, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CommonBillingContract) SetExpirationDate(v CommonBillingTimePeriod)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CommonBillingContract) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


