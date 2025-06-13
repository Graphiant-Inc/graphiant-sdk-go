# V1VersionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationMetadata** | Pointer to [**V1VersionPostRequestConfigurationMetadata**](V1VersionPostRequestConfigurationMetadata.md) |  | [optional] 
**DeviceId** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1VersionPostRequest

`func NewV1VersionPostRequest() *V1VersionPostRequest`

NewV1VersionPostRequest instantiates a new V1VersionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VersionPostRequestWithDefaults

`func NewV1VersionPostRequestWithDefaults() *V1VersionPostRequest`

NewV1VersionPostRequestWithDefaults instantiates a new V1VersionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationMetadata

`func (o *V1VersionPostRequest) GetConfigurationMetadata() V1VersionPostRequestConfigurationMetadata`

GetConfigurationMetadata returns the ConfigurationMetadata field if non-nil, zero value otherwise.

### GetConfigurationMetadataOk

`func (o *V1VersionPostRequest) GetConfigurationMetadataOk() (*V1VersionPostRequestConfigurationMetadata, bool)`

GetConfigurationMetadataOk returns a tuple with the ConfigurationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationMetadata

`func (o *V1VersionPostRequest) SetConfigurationMetadata(v V1VersionPostRequestConfigurationMetadata)`

SetConfigurationMetadata sets ConfigurationMetadata field to given value.

### HasConfigurationMetadata

`func (o *V1VersionPostRequest) HasConfigurationMetadata() bool`

HasConfigurationMetadata returns a boolean if a field has been set.

### GetDeviceId

`func (o *V1VersionPostRequest) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1VersionPostRequest) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1VersionPostRequest) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1VersionPostRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetVersion

`func (o *V1VersionPostRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1VersionPostRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1VersionPostRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1VersionPostRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


