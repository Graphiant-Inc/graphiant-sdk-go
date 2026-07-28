# V1ExtranetB2bCustomersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invite** | [**ManaV2ExtranetServiceCustomerInvite**](ManaV2ExtranetServiceCustomerInvite.md) |  | 
**Name** | **string** | Partner display name (required) | 
**Type** | **string** | Graphiant peer vs guest (non-Graphiant) partner (required) | 

## Methods

### NewV1ExtranetB2bCustomersPostRequest

`func NewV1ExtranetB2bCustomersPostRequest(invite ManaV2ExtranetServiceCustomerInvite, name string, type_ string, ) *V1ExtranetB2bCustomersPostRequest`

NewV1ExtranetB2bCustomersPostRequest instantiates a new V1ExtranetB2bCustomersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetB2bCustomersPostRequestWithDefaults

`func NewV1ExtranetB2bCustomersPostRequestWithDefaults() *V1ExtranetB2bCustomersPostRequest`

NewV1ExtranetB2bCustomersPostRequestWithDefaults instantiates a new V1ExtranetB2bCustomersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvite

`func (o *V1ExtranetB2bCustomersPostRequest) GetInvite() ManaV2ExtranetServiceCustomerInvite`

GetInvite returns the Invite field if non-nil, zero value otherwise.

### GetInviteOk

`func (o *V1ExtranetB2bCustomersPostRequest) GetInviteOk() (*ManaV2ExtranetServiceCustomerInvite, bool)`

GetInviteOk returns a tuple with the Invite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvite

`func (o *V1ExtranetB2bCustomersPostRequest) SetInvite(v ManaV2ExtranetServiceCustomerInvite)`

SetInvite sets Invite field to given value.


### GetName

`func (o *V1ExtranetB2bCustomersPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ExtranetB2bCustomersPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ExtranetB2bCustomersPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *V1ExtranetB2bCustomersPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ExtranetB2bCustomersPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ExtranetB2bCustomersPostRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


