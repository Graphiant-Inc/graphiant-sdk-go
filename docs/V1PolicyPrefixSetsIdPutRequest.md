# V1PolicyPrefixSetsIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry**](V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PrefixSetEntries** | Pointer to [**map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValue**](V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValue.md) |  | [optional] 

## Methods

### NewV1PolicyPrefixSetsIdPutRequest

`func NewV1PolicyPrefixSetsIdPutRequest() *V1PolicyPrefixSetsIdPutRequest`

NewV1PolicyPrefixSetsIdPutRequest instantiates a new V1PolicyPrefixSetsIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PolicyPrefixSetsIdPutRequestWithDefaults

`func NewV1PolicyPrefixSetsIdPutRequestWithDefaults() *V1PolicyPrefixSetsIdPutRequest`

NewV1PolicyPrefixSetsIdPutRequestWithDefaults instantiates a new V1PolicyPrefixSetsIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1PolicyPrefixSetsIdPutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1PolicyPrefixSetsIdPutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1PolicyPrefixSetsIdPutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1PolicyPrefixSetsIdPutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *V1PolicyPrefixSetsIdPutRequest) GetEntries() map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V1PolicyPrefixSetsIdPutRequest) GetEntriesOk() (*map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V1PolicyPrefixSetsIdPutRequest) SetEntries(v map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValueEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *V1PolicyPrefixSetsIdPutRequest) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetName

`func (o *V1PolicyPrefixSetsIdPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1PolicyPrefixSetsIdPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1PolicyPrefixSetsIdPutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1PolicyPrefixSetsIdPutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefixSetEntries

`func (o *V1PolicyPrefixSetsIdPutRequest) GetPrefixSetEntries() map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValue`

GetPrefixSetEntries returns the PrefixSetEntries field if non-nil, zero value otherwise.

### GetPrefixSetEntriesOk

`func (o *V1PolicyPrefixSetsIdPutRequest) GetPrefixSetEntriesOk() (*map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValue, bool)`

GetPrefixSetEntriesOk returns a tuple with the PrefixSetEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSetEntries

`func (o *V1PolicyPrefixSetsIdPutRequest) SetPrefixSetEntries(v map[string]V1GlobalConfigPatchRequestGlobalPrefixSetsValuePrefixSetEntriesValue)`

SetPrefixSetEntries sets PrefixSetEntries field to given value.

### HasPrefixSetEntries

`func (o *V1PolicyPrefixSetsIdPutRequest) HasPrefixSetEntries() bool`

HasPrefixSetEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


