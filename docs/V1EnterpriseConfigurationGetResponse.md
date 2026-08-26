# V1EnterpriseConfigurationGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdaptiveFecConfig** | Pointer to [**ManaV2AdaptiveFecConfiguration**](ManaV2AdaptiveFecConfiguration.md) |  | [optional] 
**Configuration** | Pointer to [**ManaV2EnterpriseConfiguration**](ManaV2EnterpriseConfiguration.md) |  | [optional] 

## Methods

### NewV1EnterpriseConfigurationGetResponse

`func NewV1EnterpriseConfigurationGetResponse() *V1EnterpriseConfigurationGetResponse`

NewV1EnterpriseConfigurationGetResponse instantiates a new V1EnterpriseConfigurationGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EnterpriseConfigurationGetResponseWithDefaults

`func NewV1EnterpriseConfigurationGetResponseWithDefaults() *V1EnterpriseConfigurationGetResponse`

NewV1EnterpriseConfigurationGetResponseWithDefaults instantiates a new V1EnterpriseConfigurationGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptiveFecConfig

`func (o *V1EnterpriseConfigurationGetResponse) GetAdaptiveFecConfig() ManaV2AdaptiveFecConfiguration`

GetAdaptiveFecConfig returns the AdaptiveFecConfig field if non-nil, zero value otherwise.

### GetAdaptiveFecConfigOk

`func (o *V1EnterpriseConfigurationGetResponse) GetAdaptiveFecConfigOk() (*ManaV2AdaptiveFecConfiguration, bool)`

GetAdaptiveFecConfigOk returns a tuple with the AdaptiveFecConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveFecConfig

`func (o *V1EnterpriseConfigurationGetResponse) SetAdaptiveFecConfig(v ManaV2AdaptiveFecConfiguration)`

SetAdaptiveFecConfig sets AdaptiveFecConfig field to given value.

### HasAdaptiveFecConfig

`func (o *V1EnterpriseConfigurationGetResponse) HasAdaptiveFecConfig() bool`

HasAdaptiveFecConfig returns a boolean if a field has been set.

### GetConfiguration

`func (o *V1EnterpriseConfigurationGetResponse) GetConfiguration() ManaV2EnterpriseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *V1EnterpriseConfigurationGetResponse) GetConfigurationOk() (*ManaV2EnterpriseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *V1EnterpriseConfigurationGetResponse) SetConfiguration(v ManaV2EnterpriseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *V1EnterpriseConfigurationGetResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


