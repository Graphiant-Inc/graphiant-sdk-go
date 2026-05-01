# UpgradeRolloutDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** | Device identifier. (required) | [optional] 
**Hostname** | Pointer to **string** | Device hostname for display. | [optional] 

## Methods

### NewUpgradeRolloutDevice

`func NewUpgradeRolloutDevice() *UpgradeRolloutDevice`

NewUpgradeRolloutDevice instantiates a new UpgradeRolloutDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeRolloutDeviceWithDefaults

`func NewUpgradeRolloutDeviceWithDefaults() *UpgradeRolloutDevice`

NewUpgradeRolloutDeviceWithDefaults instantiates a new UpgradeRolloutDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *UpgradeRolloutDevice) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *UpgradeRolloutDevice) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *UpgradeRolloutDevice) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *UpgradeRolloutDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetHostname

`func (o *UpgradeRolloutDevice) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpgradeRolloutDevice) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpgradeRolloutDevice) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *UpgradeRolloutDevice) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


