# AssuranceApprovedAppEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveUsers** | Pointer to **int64** | number of active users (required) | [optional] 
**Category** | Pointer to **string** | approved AI tool category (required) | [optional] 
**DataSentKbps** | Pointer to **float64** | data sent in kbps (required) | [optional] 
**Id** | Pointer to **string** | approved app entry identifier (required) | [optional] 
**Tool** | Pointer to **string** | approved AI tool name (required) | [optional] 

## Methods

### NewAssuranceApprovedAppEntry

`func NewAssuranceApprovedAppEntry() *AssuranceApprovedAppEntry`

NewAssuranceApprovedAppEntry instantiates a new AssuranceApprovedAppEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssuranceApprovedAppEntryWithDefaults

`func NewAssuranceApprovedAppEntryWithDefaults() *AssuranceApprovedAppEntry`

NewAssuranceApprovedAppEntryWithDefaults instantiates a new AssuranceApprovedAppEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveUsers

`func (o *AssuranceApprovedAppEntry) GetActiveUsers() int64`

GetActiveUsers returns the ActiveUsers field if non-nil, zero value otherwise.

### GetActiveUsersOk

`func (o *AssuranceApprovedAppEntry) GetActiveUsersOk() (*int64, bool)`

GetActiveUsersOk returns a tuple with the ActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUsers

`func (o *AssuranceApprovedAppEntry) SetActiveUsers(v int64)`

SetActiveUsers sets ActiveUsers field to given value.

### HasActiveUsers

`func (o *AssuranceApprovedAppEntry) HasActiveUsers() bool`

HasActiveUsers returns a boolean if a field has been set.

### GetCategory

`func (o *AssuranceApprovedAppEntry) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AssuranceApprovedAppEntry) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AssuranceApprovedAppEntry) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AssuranceApprovedAppEntry) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDataSentKbps

`func (o *AssuranceApprovedAppEntry) GetDataSentKbps() float64`

GetDataSentKbps returns the DataSentKbps field if non-nil, zero value otherwise.

### GetDataSentKbpsOk

`func (o *AssuranceApprovedAppEntry) GetDataSentKbpsOk() (*float64, bool)`

GetDataSentKbpsOk returns a tuple with the DataSentKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSentKbps

`func (o *AssuranceApprovedAppEntry) SetDataSentKbps(v float64)`

SetDataSentKbps sets DataSentKbps field to given value.

### HasDataSentKbps

`func (o *AssuranceApprovedAppEntry) HasDataSentKbps() bool`

HasDataSentKbps returns a boolean if a field has been set.

### GetId

`func (o *AssuranceApprovedAppEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssuranceApprovedAppEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssuranceApprovedAppEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssuranceApprovedAppEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTool

`func (o *AssuranceApprovedAppEntry) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *AssuranceApprovedAppEntry) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *AssuranceApprovedAppEntry) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *AssuranceApprovedAppEntry) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


