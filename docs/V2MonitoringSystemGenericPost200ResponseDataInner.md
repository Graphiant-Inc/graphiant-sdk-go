# V2MonitoringSystemGenericPost200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | Pointer to [**[]V2MonitoringBfdPost200ResponseDataInnerSamplesInner**](V2MonitoringBfdPost200ResponseDataInnerSamplesInner.md) |  | [optional] 
**Selector** | Pointer to [**V1AccountMfaGet200Response**](V1AccountMfaGet200Response.md) |  | [optional] 

## Methods

### NewV2MonitoringSystemGenericPost200ResponseDataInner

`func NewV2MonitoringSystemGenericPost200ResponseDataInner() *V2MonitoringSystemGenericPost200ResponseDataInner`

NewV2MonitoringSystemGenericPost200ResponseDataInner instantiates a new V2MonitoringSystemGenericPost200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MonitoringSystemGenericPost200ResponseDataInnerWithDefaults

`func NewV2MonitoringSystemGenericPost200ResponseDataInnerWithDefaults() *V2MonitoringSystemGenericPost200ResponseDataInner`

NewV2MonitoringSystemGenericPost200ResponseDataInnerWithDefaults instantiates a new V2MonitoringSystemGenericPost200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *V2MonitoringSystemGenericPost200ResponseDataInner) GetSamples() []V2MonitoringBfdPost200ResponseDataInnerSamplesInner`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *V2MonitoringSystemGenericPost200ResponseDataInner) GetSamplesOk() (*[]V2MonitoringBfdPost200ResponseDataInnerSamplesInner, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *V2MonitoringSystemGenericPost200ResponseDataInner) SetSamples(v []V2MonitoringBfdPost200ResponseDataInnerSamplesInner)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *V2MonitoringSystemGenericPost200ResponseDataInner) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSelector

`func (o *V2MonitoringSystemGenericPost200ResponseDataInner) GetSelector() V1AccountMfaGet200Response`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V2MonitoringSystemGenericPost200ResponseDataInner) GetSelectorOk() (*V1AccountMfaGet200Response, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V2MonitoringSystemGenericPost200ResponseDataInner) SetSelector(v V1AccountMfaGet200Response)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V2MonitoringSystemGenericPost200ResponseDataInner) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


