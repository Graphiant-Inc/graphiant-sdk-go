# ManaV2ExtranetServiceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsPublisher** | Pointer to **bool** | True when this enterprise publishes the service; false when consuming a remote producer | [optional] 
**LanSegment** | Pointer to **int64** | Service LAN segment when is_publisher and VRF exists | [optional] 
**LastUpdated** | Pointer to [**GoogleProtobufTimestamp**](GoogleProtobufTimestamp.md) |  | [optional] 
**ServerPrefixes** | Pointer to **[]string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**Sites** | Pointer to **[]int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalCustomers** | Pointer to **int32** |  | [optional] 

## Methods

### NewManaV2ExtranetServiceSummary

`func NewManaV2ExtranetServiceSummary() *ManaV2ExtranetServiceSummary`

NewManaV2ExtranetServiceSummary instantiates a new ManaV2ExtranetServiceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2ExtranetServiceSummaryWithDefaults

`func NewManaV2ExtranetServiceSummaryWithDefaults() *ManaV2ExtranetServiceSummary`

NewManaV2ExtranetServiceSummaryWithDefaults instantiates a new ManaV2ExtranetServiceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ManaV2ExtranetServiceSummary) GetCreatedAt() GoogleProtobufTimestamp`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManaV2ExtranetServiceSummary) GetCreatedAtOk() (*GoogleProtobufTimestamp, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManaV2ExtranetServiceSummary) SetCreatedAt(v GoogleProtobufTimestamp)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ManaV2ExtranetServiceSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ManaV2ExtranetServiceSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManaV2ExtranetServiceSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManaV2ExtranetServiceSummary) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ManaV2ExtranetServiceSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPublisher

`func (o *ManaV2ExtranetServiceSummary) GetIsPublisher() bool`

GetIsPublisher returns the IsPublisher field if non-nil, zero value otherwise.

### GetIsPublisherOk

`func (o *ManaV2ExtranetServiceSummary) GetIsPublisherOk() (*bool, bool)`

GetIsPublisherOk returns a tuple with the IsPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublisher

`func (o *ManaV2ExtranetServiceSummary) SetIsPublisher(v bool)`

SetIsPublisher sets IsPublisher field to given value.

### HasIsPublisher

`func (o *ManaV2ExtranetServiceSummary) HasIsPublisher() bool`

HasIsPublisher returns a boolean if a field has been set.

### GetLanSegment

`func (o *ManaV2ExtranetServiceSummary) GetLanSegment() int64`

GetLanSegment returns the LanSegment field if non-nil, zero value otherwise.

### GetLanSegmentOk

`func (o *ManaV2ExtranetServiceSummary) GetLanSegmentOk() (*int64, bool)`

GetLanSegmentOk returns a tuple with the LanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegment

`func (o *ManaV2ExtranetServiceSummary) SetLanSegment(v int64)`

SetLanSegment sets LanSegment field to given value.

### HasLanSegment

`func (o *ManaV2ExtranetServiceSummary) HasLanSegment() bool`

HasLanSegment returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ManaV2ExtranetServiceSummary) GetLastUpdated() GoogleProtobufTimestamp`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ManaV2ExtranetServiceSummary) GetLastUpdatedOk() (*GoogleProtobufTimestamp, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ManaV2ExtranetServiceSummary) SetLastUpdated(v GoogleProtobufTimestamp)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ManaV2ExtranetServiceSummary) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetServerPrefixes

`func (o *ManaV2ExtranetServiceSummary) GetServerPrefixes() []string`

GetServerPrefixes returns the ServerPrefixes field if non-nil, zero value otherwise.

### GetServerPrefixesOk

`func (o *ManaV2ExtranetServiceSummary) GetServerPrefixesOk() (*[]string, bool)`

GetServerPrefixesOk returns a tuple with the ServerPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPrefixes

`func (o *ManaV2ExtranetServiceSummary) SetServerPrefixes(v []string)`

SetServerPrefixes sets ServerPrefixes field to given value.

### HasServerPrefixes

`func (o *ManaV2ExtranetServiceSummary) HasServerPrefixes() bool`

HasServerPrefixes returns a boolean if a field has been set.

### GetServiceName

`func (o *ManaV2ExtranetServiceSummary) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ManaV2ExtranetServiceSummary) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ManaV2ExtranetServiceSummary) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ManaV2ExtranetServiceSummary) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceType

`func (o *ManaV2ExtranetServiceSummary) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ManaV2ExtranetServiceSummary) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ManaV2ExtranetServiceSummary) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ManaV2ExtranetServiceSummary) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSites

`func (o *ManaV2ExtranetServiceSummary) GetSites() []int64`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ManaV2ExtranetServiceSummary) GetSitesOk() (*[]int64, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ManaV2ExtranetServiceSummary) SetSites(v []int64)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ManaV2ExtranetServiceSummary) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetStatus

`func (o *ManaV2ExtranetServiceSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManaV2ExtranetServiceSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManaV2ExtranetServiceSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManaV2ExtranetServiceSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalCustomers

`func (o *ManaV2ExtranetServiceSummary) GetTotalCustomers() int32`

GetTotalCustomers returns the TotalCustomers field if non-nil, zero value otherwise.

### GetTotalCustomersOk

`func (o *ManaV2ExtranetServiceSummary) GetTotalCustomersOk() (*int32, bool)`

GetTotalCustomersOk returns a tuple with the TotalCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCustomers

`func (o *ManaV2ExtranetServiceSummary) SetTotalCustomers(v int32)`

SetTotalCustomers sets TotalCustomers field to given value.

### HasTotalCustomers

`func (o *ManaV2ExtranetServiceSummary) HasTotalCustomers() bool`

HasTotalCustomers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


