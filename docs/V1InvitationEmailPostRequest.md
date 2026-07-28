# V1InvitationEmailPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEmail** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **int64** |  | [optional] 
**CustomerName** | **string** |  (required) | 
**IsGraphiant** | Pointer to **bool** |  | [optional] 
**MatchId** | **int64** |  (required) | 
**ServiceId** | **int64** |  (required) | 
**ServiceName** | **string** |  (required) | 
**ServiceType** | Pointer to **string** | Extranet service type URL segment (e.g. peering_service, client_to_server) | [optional] 

## Methods

### NewV1InvitationEmailPostRequest

`func NewV1InvitationEmailPostRequest(customerName string, matchId int64, serviceId int64, serviceName string, ) *V1InvitationEmailPostRequest`

NewV1InvitationEmailPostRequest instantiates a new V1InvitationEmailPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InvitationEmailPostRequestWithDefaults

`func NewV1InvitationEmailPostRequestWithDefaults() *V1InvitationEmailPostRequest`

NewV1InvitationEmailPostRequestWithDefaults instantiates a new V1InvitationEmailPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEmail

`func (o *V1InvitationEmailPostRequest) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *V1InvitationEmailPostRequest) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *V1InvitationEmailPostRequest) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *V1InvitationEmailPostRequest) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetCustomerId

`func (o *V1InvitationEmailPostRequest) GetCustomerId() int64`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *V1InvitationEmailPostRequest) GetCustomerIdOk() (*int64, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *V1InvitationEmailPostRequest) SetCustomerId(v int64)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *V1InvitationEmailPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerName

`func (o *V1InvitationEmailPostRequest) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *V1InvitationEmailPostRequest) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *V1InvitationEmailPostRequest) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### GetIsGraphiant

`func (o *V1InvitationEmailPostRequest) GetIsGraphiant() bool`

GetIsGraphiant returns the IsGraphiant field if non-nil, zero value otherwise.

### GetIsGraphiantOk

`func (o *V1InvitationEmailPostRequest) GetIsGraphiantOk() (*bool, bool)`

GetIsGraphiantOk returns a tuple with the IsGraphiant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGraphiant

`func (o *V1InvitationEmailPostRequest) SetIsGraphiant(v bool)`

SetIsGraphiant sets IsGraphiant field to given value.

### HasIsGraphiant

`func (o *V1InvitationEmailPostRequest) HasIsGraphiant() bool`

HasIsGraphiant returns a boolean if a field has been set.

### GetMatchId

`func (o *V1InvitationEmailPostRequest) GetMatchId() int64`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *V1InvitationEmailPostRequest) GetMatchIdOk() (*int64, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *V1InvitationEmailPostRequest) SetMatchId(v int64)`

SetMatchId sets MatchId field to given value.


### GetServiceId

`func (o *V1InvitationEmailPostRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *V1InvitationEmailPostRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *V1InvitationEmailPostRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetServiceName

`func (o *V1InvitationEmailPostRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1InvitationEmailPostRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1InvitationEmailPostRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServiceType

`func (o *V1InvitationEmailPostRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *V1InvitationEmailPostRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *V1InvitationEmailPostRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *V1InvitationEmailPostRequest) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


