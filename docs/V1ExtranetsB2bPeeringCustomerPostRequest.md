# V1ExtranetsB2bPeeringCustomerPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invite** | [**ManaV2B2bExtranetPeeringServiceCustomerInvite**](ManaV2B2bExtranetPeeringServiceCustomerInvite.md) |  | 
**Name** | **string** | Name of the peering service customer (required) | 
**Type** | **string** | Type of the peerings servicecustomer whether it is a graphiant or non-graphiant (required) | 

## Methods

### NewV1ExtranetsB2bPeeringCustomerPostRequest

`func NewV1ExtranetsB2bPeeringCustomerPostRequest(invite ManaV2B2bExtranetPeeringServiceCustomerInvite, name string, type_ string, ) *V1ExtranetsB2bPeeringCustomerPostRequest`

NewV1ExtranetsB2bPeeringCustomerPostRequest instantiates a new V1ExtranetsB2bPeeringCustomerPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExtranetsB2bPeeringCustomerPostRequestWithDefaults

`func NewV1ExtranetsB2bPeeringCustomerPostRequestWithDefaults() *V1ExtranetsB2bPeeringCustomerPostRequest`

NewV1ExtranetsB2bPeeringCustomerPostRequestWithDefaults instantiates a new V1ExtranetsB2bPeeringCustomerPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvite

`func (o *V1ExtranetsB2bPeeringCustomerPostRequest) GetInvite() ManaV2B2bExtranetPeeringServiceCustomerInvite`

GetInvite returns the Invite field if non-nil, zero value otherwise.

### GetInviteOk

`func (o *V1ExtranetsB2bPeeringCustomerPostRequest) GetInviteOk() (*ManaV2B2bExtranetPeeringServiceCustomerInvite, bool)`

GetInviteOk returns a tuple with the Invite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvite

`func (o *V1ExtranetsB2bPeeringCustomerPostRequest) SetInvite(v ManaV2B2bExtranetPeeringServiceCustomerInvite)`

SetInvite sets Invite field to given value.


### GetName

`func (o *V1ExtranetsB2bPeeringCustomerPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ExtranetsB2bPeeringCustomerPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ExtranetsB2bPeeringCustomerPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *V1ExtranetsB2bPeeringCustomerPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ExtranetsB2bPeeringCustomerPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ExtranetsB2bPeeringCustomerPostRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


