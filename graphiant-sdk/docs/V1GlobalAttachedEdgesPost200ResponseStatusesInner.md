# V1GlobalAttachedEdgesPost200ResponseStatusesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**InternalState** | Pointer to **string** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusSince** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1GlobalAttachedEdgesPost200ResponseStatusesInner

`func NewV1GlobalAttachedEdgesPost200ResponseStatusesInner() *V1GlobalAttachedEdgesPost200ResponseStatusesInner`

NewV1GlobalAttachedEdgesPost200ResponseStatusesInner instantiates a new V1GlobalAttachedEdgesPost200ResponseStatusesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalAttachedEdgesPost200ResponseStatusesInnerWithDefaults

`func NewV1GlobalAttachedEdgesPost200ResponseStatusesInnerWithDefaults() *V1GlobalAttachedEdgesPost200ResponseStatusesInner`

NewV1GlobalAttachedEdgesPost200ResponseStatusesInnerWithDefaults instantiates a new V1GlobalAttachedEdgesPost200ResponseStatusesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetHostname

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInternalState

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetInternalState() string`

GetInternalState returns the InternalState field if non-nil, zero value otherwise.

### GetInternalStateOk

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetInternalStateOk() (*string, bool)`

GetInternalStateOk returns a tuple with the InternalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalState

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) SetInternalState(v string)`

SetInternalState sets InternalState field to given value.

### HasInternalState

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) HasInternalState() bool`

HasInternalState returns a boolean if a field has been set.

### GetSiteName

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetStatus

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusSince

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetStatusSince() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetStatusSince returns the StatusSince field if non-nil, zero value otherwise.

### GetStatusSinceOk

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) GetStatusSinceOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetStatusSinceOk returns a tuple with the StatusSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSince

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) SetStatusSince(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetStatusSince sets StatusSince field to given value.

### HasStatusSince

`func (o *V1GlobalAttachedEdgesPost200ResponseStatusesInner) HasStatusSince() bool`

HasStatusSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


