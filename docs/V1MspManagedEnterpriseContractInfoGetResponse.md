# V1MspManagedEnterpriseContractInfoGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enterprises** | Pointer to [**[]ManaV2ManagedEnterpriseContractInfo**](ManaV2ManagedEnterpriseContractInfo.md) |  | [optional] 
**MspConsumedCredits** | Pointer to **float32** | All credits consumed by the MSP over the entirety of its contracts | [optional] 
**MspContract** | Pointer to [**CommonBillingContract**](CommonBillingContract.md) |  | [optional] 

## Methods

### NewV1MspManagedEnterpriseContractInfoGetResponse

`func NewV1MspManagedEnterpriseContractInfoGetResponse() *V1MspManagedEnterpriseContractInfoGetResponse`

NewV1MspManagedEnterpriseContractInfoGetResponse instantiates a new V1MspManagedEnterpriseContractInfoGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MspManagedEnterpriseContractInfoGetResponseWithDefaults

`func NewV1MspManagedEnterpriseContractInfoGetResponseWithDefaults() *V1MspManagedEnterpriseContractInfoGetResponse`

NewV1MspManagedEnterpriseContractInfoGetResponseWithDefaults instantiates a new V1MspManagedEnterpriseContractInfoGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnterprises

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) GetEnterprises() []ManaV2ManagedEnterpriseContractInfo`

GetEnterprises returns the Enterprises field if non-nil, zero value otherwise.

### GetEnterprisesOk

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) GetEnterprisesOk() (*[]ManaV2ManagedEnterpriseContractInfo, bool)`

GetEnterprisesOk returns a tuple with the Enterprises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprises

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) SetEnterprises(v []ManaV2ManagedEnterpriseContractInfo)`

SetEnterprises sets Enterprises field to given value.

### HasEnterprises

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) HasEnterprises() bool`

HasEnterprises returns a boolean if a field has been set.

### GetMspConsumedCredits

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) GetMspConsumedCredits() float32`

GetMspConsumedCredits returns the MspConsumedCredits field if non-nil, zero value otherwise.

### GetMspConsumedCreditsOk

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) GetMspConsumedCreditsOk() (*float32, bool)`

GetMspConsumedCreditsOk returns a tuple with the MspConsumedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspConsumedCredits

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) SetMspConsumedCredits(v float32)`

SetMspConsumedCredits sets MspConsumedCredits field to given value.

### HasMspConsumedCredits

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) HasMspConsumedCredits() bool`

HasMspConsumedCredits returns a boolean if a field has been set.

### GetMspContract

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) GetMspContract() CommonBillingContract`

GetMspContract returns the MspContract field if non-nil, zero value otherwise.

### GetMspContractOk

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) GetMspContractOk() (*CommonBillingContract, bool)`

GetMspContractOk returns a tuple with the MspContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspContract

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) SetMspContract(v CommonBillingContract)`

SetMspContract sets MspContract field to given value.

### HasMspContract

`func (o *V1MspManagedEnterpriseContractInfoGetResponse) HasMspContract() bool`

HasMspContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


