# V1EnterpriseContractPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractedCredits** | Pointer to **float32** | Amount of credits billed for a contract term or per month when no expiration is provided. For an MSP tenant, this is the total credits in the MSP billing pool (from which credits are allocated to managed enterprises). | [optional] 
**ExpirationDate** | Pointer to [**ManaV2TimePeriod**](ManaV2TimePeriod.md) |  | [optional] 

## Methods

### NewV1EnterpriseContractPutRequest

`func NewV1EnterpriseContractPutRequest() *V1EnterpriseContractPutRequest`

NewV1EnterpriseContractPutRequest instantiates a new V1EnterpriseContractPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EnterpriseContractPutRequestWithDefaults

`func NewV1EnterpriseContractPutRequestWithDefaults() *V1EnterpriseContractPutRequest`

NewV1EnterpriseContractPutRequestWithDefaults instantiates a new V1EnterpriseContractPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractedCredits

`func (o *V1EnterpriseContractPutRequest) GetContractedCredits() float32`

GetContractedCredits returns the ContractedCredits field if non-nil, zero value otherwise.

### GetContractedCreditsOk

`func (o *V1EnterpriseContractPutRequest) GetContractedCreditsOk() (*float32, bool)`

GetContractedCreditsOk returns a tuple with the ContractedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractedCredits

`func (o *V1EnterpriseContractPutRequest) SetContractedCredits(v float32)`

SetContractedCredits sets ContractedCredits field to given value.

### HasContractedCredits

`func (o *V1EnterpriseContractPutRequest) HasContractedCredits() bool`

HasContractedCredits returns a boolean if a field has been set.

### GetExpirationDate

`func (o *V1EnterpriseContractPutRequest) GetExpirationDate() ManaV2TimePeriod`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *V1EnterpriseContractPutRequest) GetExpirationDateOk() (*ManaV2TimePeriod, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *V1EnterpriseContractPutRequest) SetExpirationDate(v ManaV2TimePeriod)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *V1EnterpriseContractPutRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


