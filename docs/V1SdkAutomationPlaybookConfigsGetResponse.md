# V1SdkAutomationPlaybookConfigsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configs** | Pointer to [**[]SdkAutomationPlaybookConfigSummary**](SdkAutomationPlaybookConfigSummary.md) |  | [optional] 
**PageInfo** | Pointer to [**CommonPageInfo**](CommonPageInfo.md) |  | [optional] 

## Methods

### NewV1SdkAutomationPlaybookConfigsGetResponse

`func NewV1SdkAutomationPlaybookConfigsGetResponse() *V1SdkAutomationPlaybookConfigsGetResponse`

NewV1SdkAutomationPlaybookConfigsGetResponse instantiates a new V1SdkAutomationPlaybookConfigsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SdkAutomationPlaybookConfigsGetResponseWithDefaults

`func NewV1SdkAutomationPlaybookConfigsGetResponseWithDefaults() *V1SdkAutomationPlaybookConfigsGetResponse`

NewV1SdkAutomationPlaybookConfigsGetResponseWithDefaults instantiates a new V1SdkAutomationPlaybookConfigsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigs

`func (o *V1SdkAutomationPlaybookConfigsGetResponse) GetConfigs() []SdkAutomationPlaybookConfigSummary`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *V1SdkAutomationPlaybookConfigsGetResponse) GetConfigsOk() (*[]SdkAutomationPlaybookConfigSummary, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *V1SdkAutomationPlaybookConfigsGetResponse) SetConfigs(v []SdkAutomationPlaybookConfigSummary)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *V1SdkAutomationPlaybookConfigsGetResponse) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetPageInfo

`func (o *V1SdkAutomationPlaybookConfigsGetResponse) GetPageInfo() CommonPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1SdkAutomationPlaybookConfigsGetResponse) GetPageInfoOk() (*CommonPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1SdkAutomationPlaybookConfigsGetResponse) SetPageInfo(v CommonPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1SdkAutomationPlaybookConfigsGetResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


