# V1DevicesDeviceIdConfigPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationMetadata** | Pointer to [**V1VersionPostRequestConfigurationMetadata**](V1VersionPostRequestConfigurationMetadata.md) |  | [optional] 
**Core** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCore**](V1DevicesDeviceIdConfigPutRequestCore.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Edge** | Pointer to [**V1DevicesDeviceIdConfigPutRequestEdge**](V1DevicesDeviceIdConfigPutRequestEdge.md) |  | [optional] 
**LocalWebServerPassword** | Pointer to **string** |  | [optional] 
**Replace** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequest

`func NewV1DevicesDeviceIdConfigPutRequest() *V1DevicesDeviceIdConfigPutRequest`

NewV1DevicesDeviceIdConfigPutRequest instantiates a new V1DevicesDeviceIdConfigPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestWithDefaults() *V1DevicesDeviceIdConfigPutRequest`

NewV1DevicesDeviceIdConfigPutRequestWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationMetadata

`func (o *V1DevicesDeviceIdConfigPutRequest) GetConfigurationMetadata() V1VersionPostRequestConfigurationMetadata`

GetConfigurationMetadata returns the ConfigurationMetadata field if non-nil, zero value otherwise.

### GetConfigurationMetadataOk

`func (o *V1DevicesDeviceIdConfigPutRequest) GetConfigurationMetadataOk() (*V1VersionPostRequestConfigurationMetadata, bool)`

GetConfigurationMetadataOk returns a tuple with the ConfigurationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationMetadata

`func (o *V1DevicesDeviceIdConfigPutRequest) SetConfigurationMetadata(v V1VersionPostRequestConfigurationMetadata)`

SetConfigurationMetadata sets ConfigurationMetadata field to given value.

### HasConfigurationMetadata

`func (o *V1DevicesDeviceIdConfigPutRequest) HasConfigurationMetadata() bool`

HasConfigurationMetadata returns a boolean if a field has been set.

### GetCore

`func (o *V1DevicesDeviceIdConfigPutRequest) GetCore() V1DevicesDeviceIdConfigPutRequestCore`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *V1DevicesDeviceIdConfigPutRequest) GetCoreOk() (*V1DevicesDeviceIdConfigPutRequestCore, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *V1DevicesDeviceIdConfigPutRequest) SetCore(v V1DevicesDeviceIdConfigPutRequestCore)`

SetCore sets Core field to given value.

### HasCore

`func (o *V1DevicesDeviceIdConfigPutRequest) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetDescription

`func (o *V1DevicesDeviceIdConfigPutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1DevicesDeviceIdConfigPutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1DevicesDeviceIdConfigPutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1DevicesDeviceIdConfigPutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEdge

`func (o *V1DevicesDeviceIdConfigPutRequest) GetEdge() V1DevicesDeviceIdConfigPutRequestEdge`

GetEdge returns the Edge field if non-nil, zero value otherwise.

### GetEdgeOk

`func (o *V1DevicesDeviceIdConfigPutRequest) GetEdgeOk() (*V1DevicesDeviceIdConfigPutRequestEdge, bool)`

GetEdgeOk returns a tuple with the Edge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdge

`func (o *V1DevicesDeviceIdConfigPutRequest) SetEdge(v V1DevicesDeviceIdConfigPutRequestEdge)`

SetEdge sets Edge field to given value.

### HasEdge

`func (o *V1DevicesDeviceIdConfigPutRequest) HasEdge() bool`

HasEdge returns a boolean if a field has been set.

### GetLocalWebServerPassword

`func (o *V1DevicesDeviceIdConfigPutRequest) GetLocalWebServerPassword() string`

GetLocalWebServerPassword returns the LocalWebServerPassword field if non-nil, zero value otherwise.

### GetLocalWebServerPasswordOk

`func (o *V1DevicesDeviceIdConfigPutRequest) GetLocalWebServerPasswordOk() (*string, bool)`

GetLocalWebServerPasswordOk returns a tuple with the LocalWebServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalWebServerPassword

`func (o *V1DevicesDeviceIdConfigPutRequest) SetLocalWebServerPassword(v string)`

SetLocalWebServerPassword sets LocalWebServerPassword field to given value.

### HasLocalWebServerPassword

`func (o *V1DevicesDeviceIdConfigPutRequest) HasLocalWebServerPassword() bool`

HasLocalWebServerPassword returns a boolean if a field has been set.

### GetReplace

`func (o *V1DevicesDeviceIdConfigPutRequest) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *V1DevicesDeviceIdConfigPutRequest) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *V1DevicesDeviceIdConfigPutRequest) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *V1DevicesDeviceIdConfigPutRequest) HasReplace() bool`

HasReplace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


