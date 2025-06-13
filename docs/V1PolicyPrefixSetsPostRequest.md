# V1PolicyPrefixSetsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**map[string]V1PolicyPrefixSetsPostRequestEntriesValue**](V1PolicyPrefixSetsPostRequestEntriesValue.md) |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PrefixSetEntries** | Pointer to [**map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry**](V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry.md) |  | [optional] 

## Methods

### NewV1PolicyPrefixSetsPostRequest

`func NewV1PolicyPrefixSetsPostRequest() *V1PolicyPrefixSetsPostRequest`

NewV1PolicyPrefixSetsPostRequest instantiates a new V1PolicyPrefixSetsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PolicyPrefixSetsPostRequestWithDefaults

`func NewV1PolicyPrefixSetsPostRequestWithDefaults() *V1PolicyPrefixSetsPostRequest`

NewV1PolicyPrefixSetsPostRequestWithDefaults instantiates a new V1PolicyPrefixSetsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1PolicyPrefixSetsPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1PolicyPrefixSetsPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1PolicyPrefixSetsPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1PolicyPrefixSetsPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *V1PolicyPrefixSetsPostRequest) GetEntries() map[string]V1PolicyPrefixSetsPostRequestEntriesValue`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V1PolicyPrefixSetsPostRequest) GetEntriesOk() (*map[string]V1PolicyPrefixSetsPostRequestEntriesValue, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V1PolicyPrefixSetsPostRequest) SetEntries(v map[string]V1PolicyPrefixSetsPostRequestEntriesValue)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *V1PolicyPrefixSetsPostRequest) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetMode

`func (o *V1PolicyPrefixSetsPostRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *V1PolicyPrefixSetsPostRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *V1PolicyPrefixSetsPostRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *V1PolicyPrefixSetsPostRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *V1PolicyPrefixSetsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1PolicyPrefixSetsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1PolicyPrefixSetsPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1PolicyPrefixSetsPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefixSetEntries

`func (o *V1PolicyPrefixSetsPostRequest) GetPrefixSetEntries() map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry`

GetPrefixSetEntries returns the PrefixSetEntries field if non-nil, zero value otherwise.

### GetPrefixSetEntriesOk

`func (o *V1PolicyPrefixSetsPostRequest) GetPrefixSetEntriesOk() (*map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry, bool)`

GetPrefixSetEntriesOk returns a tuple with the PrefixSetEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSetEntries

`func (o *V1PolicyPrefixSetsPostRequest) SetPrefixSetEntries(v map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry)`

SetPrefixSetEntries sets PrefixSetEntries field to given value.

### HasPrefixSetEntries

`func (o *V1PolicyPrefixSetsPostRequest) HasPrefixSetEntries() bool`

HasPrefixSetEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


