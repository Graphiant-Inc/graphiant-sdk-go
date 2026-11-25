# V1MspRegionsExpectedInterRegionLatenciesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterRegionLatencyMatrix** | Pointer to [**ManaV2InterRegionLatencyMatrix**](ManaV2InterRegionLatencyMatrix.md) |  | [optional] 
**SourceRegionMissingDestinations** | Pointer to [**map[string]V1MspRegionsExpectedInterRegionLatenciesGetResponseMissingDestinationRegions**](V1MspRegionsExpectedInterRegionLatenciesGetResponseMissingDestinationRegions.md) |  | [optional] 

## Methods

### NewV1MspRegionsExpectedInterRegionLatenciesGetResponse

`func NewV1MspRegionsExpectedInterRegionLatenciesGetResponse() *V1MspRegionsExpectedInterRegionLatenciesGetResponse`

NewV1MspRegionsExpectedInterRegionLatenciesGetResponse instantiates a new V1MspRegionsExpectedInterRegionLatenciesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MspRegionsExpectedInterRegionLatenciesGetResponseWithDefaults

`func NewV1MspRegionsExpectedInterRegionLatenciesGetResponseWithDefaults() *V1MspRegionsExpectedInterRegionLatenciesGetResponse`

NewV1MspRegionsExpectedInterRegionLatenciesGetResponseWithDefaults instantiates a new V1MspRegionsExpectedInterRegionLatenciesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterRegionLatencyMatrix

`func (o *V1MspRegionsExpectedInterRegionLatenciesGetResponse) GetInterRegionLatencyMatrix() ManaV2InterRegionLatencyMatrix`

GetInterRegionLatencyMatrix returns the InterRegionLatencyMatrix field if non-nil, zero value otherwise.

### GetInterRegionLatencyMatrixOk

`func (o *V1MspRegionsExpectedInterRegionLatenciesGetResponse) GetInterRegionLatencyMatrixOk() (*ManaV2InterRegionLatencyMatrix, bool)`

GetInterRegionLatencyMatrixOk returns a tuple with the InterRegionLatencyMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterRegionLatencyMatrix

`func (o *V1MspRegionsExpectedInterRegionLatenciesGetResponse) SetInterRegionLatencyMatrix(v ManaV2InterRegionLatencyMatrix)`

SetInterRegionLatencyMatrix sets InterRegionLatencyMatrix field to given value.

### HasInterRegionLatencyMatrix

`func (o *V1MspRegionsExpectedInterRegionLatenciesGetResponse) HasInterRegionLatencyMatrix() bool`

HasInterRegionLatencyMatrix returns a boolean if a field has been set.

### GetSourceRegionMissingDestinations

`func (o *V1MspRegionsExpectedInterRegionLatenciesGetResponse) GetSourceRegionMissingDestinations() map[string]V1MspRegionsExpectedInterRegionLatenciesGetResponseMissingDestinationRegions`

GetSourceRegionMissingDestinations returns the SourceRegionMissingDestinations field if non-nil, zero value otherwise.

### GetSourceRegionMissingDestinationsOk

`func (o *V1MspRegionsExpectedInterRegionLatenciesGetResponse) GetSourceRegionMissingDestinationsOk() (*map[string]V1MspRegionsExpectedInterRegionLatenciesGetResponseMissingDestinationRegions, bool)`

GetSourceRegionMissingDestinationsOk returns a tuple with the SourceRegionMissingDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegionMissingDestinations

`func (o *V1MspRegionsExpectedInterRegionLatenciesGetResponse) SetSourceRegionMissingDestinations(v map[string]V1MspRegionsExpectedInterRegionLatenciesGetResponseMissingDestinationRegions)`

SetSourceRegionMissingDestinations sets SourceRegionMissingDestinations field to given value.

### HasSourceRegionMissingDestinations

`func (o *V1MspRegionsExpectedInterRegionLatenciesGetResponse) HasSourceRegionMissingDestinations() bool`

HasSourceRegionMissingDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


