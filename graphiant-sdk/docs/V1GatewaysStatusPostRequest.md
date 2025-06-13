# V1GatewaysStatusPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceInfo** | Pointer to [**[]V1GatewaysStatusPostRequestDeviceInfoInner**](V1GatewaysStatusPostRequestDeviceInfoInner.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SupportStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewV1GatewaysStatusPostRequest

`func NewV1GatewaysStatusPostRequest() *V1GatewaysStatusPostRequest`

NewV1GatewaysStatusPostRequest instantiates a new V1GatewaysStatusPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GatewaysStatusPostRequestWithDefaults

`func NewV1GatewaysStatusPostRequestWithDefaults() *V1GatewaysStatusPostRequest`

NewV1GatewaysStatusPostRequestWithDefaults instantiates a new V1GatewaysStatusPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceInfo

`func (o *V1GatewaysStatusPostRequest) GetDeviceInfo() []V1GatewaysStatusPostRequestDeviceInfoInner`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *V1GatewaysStatusPostRequest) GetDeviceInfoOk() (*[]V1GatewaysStatusPostRequestDeviceInfoInner, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *V1GatewaysStatusPostRequest) SetDeviceInfo(v []V1GatewaysStatusPostRequestDeviceInfoInner)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *V1GatewaysStatusPostRequest) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetId

`func (o *V1GatewaysStatusPostRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1GatewaysStatusPostRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1GatewaysStatusPostRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *V1GatewaysStatusPostRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *V1GatewaysStatusPostRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1GatewaysStatusPostRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1GatewaysStatusPostRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1GatewaysStatusPostRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportStatus

`func (o *V1GatewaysStatusPostRequest) GetSupportStatus() string`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### GetSupportStatusOk

`func (o *V1GatewaysStatusPostRequest) GetSupportStatusOk() (*string, bool)`

GetSupportStatusOk returns a tuple with the SupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatus

`func (o *V1GatewaysStatusPostRequest) SetSupportStatus(v string)`

SetSupportStatus sets SupportStatus field to given value.

### HasSupportStatus

`func (o *V1GatewaysStatusPostRequest) HasSupportStatus() bool`

HasSupportStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


