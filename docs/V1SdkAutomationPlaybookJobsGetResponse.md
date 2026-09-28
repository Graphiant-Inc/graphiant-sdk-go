# V1SdkAutomationPlaybookJobsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]SdkAutomationPlaybookJob**](SdkAutomationPlaybookJob.md) |  | [optional] 
**PageInfo** | Pointer to [**CommonPageInfo**](CommonPageInfo.md) |  | [optional] 

## Methods

### NewV1SdkAutomationPlaybookJobsGetResponse

`func NewV1SdkAutomationPlaybookJobsGetResponse() *V1SdkAutomationPlaybookJobsGetResponse`

NewV1SdkAutomationPlaybookJobsGetResponse instantiates a new V1SdkAutomationPlaybookJobsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SdkAutomationPlaybookJobsGetResponseWithDefaults

`func NewV1SdkAutomationPlaybookJobsGetResponseWithDefaults() *V1SdkAutomationPlaybookJobsGetResponse`

NewV1SdkAutomationPlaybookJobsGetResponseWithDefaults instantiates a new V1SdkAutomationPlaybookJobsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *V1SdkAutomationPlaybookJobsGetResponse) GetJobs() []SdkAutomationPlaybookJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V1SdkAutomationPlaybookJobsGetResponse) GetJobsOk() (*[]SdkAutomationPlaybookJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V1SdkAutomationPlaybookJobsGetResponse) SetJobs(v []SdkAutomationPlaybookJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V1SdkAutomationPlaybookJobsGetResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetPageInfo

`func (o *V1SdkAutomationPlaybookJobsGetResponse) GetPageInfo() CommonPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *V1SdkAutomationPlaybookJobsGetResponse) GetPageInfoOk() (*CommonPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *V1SdkAutomationPlaybookJobsGetResponse) SetPageInfo(v CommonPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *V1SdkAutomationPlaybookJobsGetResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


