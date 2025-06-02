# V1NatUtilizationDeviceIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatUsage** | Pointer to [**V1NatUtilizationDeviceIdGet200ResponseNatUsage**](V1NatUtilizationDeviceIdGet200ResponseNatUsage.md) |  | [optional] 
**Ts** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 

## Methods

### NewV1NatUtilizationDeviceIdGet200Response

`func NewV1NatUtilizationDeviceIdGet200Response() *V1NatUtilizationDeviceIdGet200Response`

NewV1NatUtilizationDeviceIdGet200Response instantiates a new V1NatUtilizationDeviceIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NatUtilizationDeviceIdGet200ResponseWithDefaults

`func NewV1NatUtilizationDeviceIdGet200ResponseWithDefaults() *V1NatUtilizationDeviceIdGet200Response`

NewV1NatUtilizationDeviceIdGet200ResponseWithDefaults instantiates a new V1NatUtilizationDeviceIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatUsage

`func (o *V1NatUtilizationDeviceIdGet200Response) GetNatUsage() V1NatUtilizationDeviceIdGet200ResponseNatUsage`

GetNatUsage returns the NatUsage field if non-nil, zero value otherwise.

### GetNatUsageOk

`func (o *V1NatUtilizationDeviceIdGet200Response) GetNatUsageOk() (*V1NatUtilizationDeviceIdGet200ResponseNatUsage, bool)`

GetNatUsageOk returns a tuple with the NatUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatUsage

`func (o *V1NatUtilizationDeviceIdGet200Response) SetNatUsage(v V1NatUtilizationDeviceIdGet200ResponseNatUsage)`

SetNatUsage sets NatUsage field to given value.

### HasNatUsage

`func (o *V1NatUtilizationDeviceIdGet200Response) HasNatUsage() bool`

HasNatUsage returns a boolean if a field has been set.

### GetTs

`func (o *V1NatUtilizationDeviceIdGet200Response) GetTs() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *V1NatUtilizationDeviceIdGet200Response) GetTsOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *V1NatUtilizationDeviceIdGet200Response) SetTs(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetTs sets Ts field to given value.

### HasTs

`func (o *V1NatUtilizationDeviceIdGet200Response) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


