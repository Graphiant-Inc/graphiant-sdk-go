# V1DevicesBringupPost200ResponseSummariesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int64** |  | [optional] 
**DiscoveredLocation** | Pointer to **string** |  | [optional] 
**FirstAppearedOn** | Pointer to [**V1AlarmHistoryGet200ResponseHistoryInnerTime**](V1AlarmHistoryGet200ResponseHistoryInnerTime.md) |  | [optional] 
**IpDetected** | Pointer to **string** |  | [optional] 
**IsNew** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DevicesBringupPost200ResponseSummariesInner

`func NewV1DevicesBringupPost200ResponseSummariesInner() *V1DevicesBringupPost200ResponseSummariesInner`

NewV1DevicesBringupPost200ResponseSummariesInner instantiates a new V1DevicesBringupPost200ResponseSummariesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesBringupPost200ResponseSummariesInnerWithDefaults

`func NewV1DevicesBringupPost200ResponseSummariesInnerWithDefaults() *V1DevicesBringupPost200ResponseSummariesInner`

NewV1DevicesBringupPost200ResponseSummariesInnerWithDefaults instantiates a new V1DevicesBringupPost200ResponseSummariesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V1DevicesBringupPost200ResponseSummariesInner) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V1DevicesBringupPost200ResponseSummariesInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDiscoveredLocation

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetDiscoveredLocation() string`

GetDiscoveredLocation returns the DiscoveredLocation field if non-nil, zero value otherwise.

### GetDiscoveredLocationOk

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetDiscoveredLocationOk() (*string, bool)`

GetDiscoveredLocationOk returns a tuple with the DiscoveredLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredLocation

`func (o *V1DevicesBringupPost200ResponseSummariesInner) SetDiscoveredLocation(v string)`

SetDiscoveredLocation sets DiscoveredLocation field to given value.

### HasDiscoveredLocation

`func (o *V1DevicesBringupPost200ResponseSummariesInner) HasDiscoveredLocation() bool`

HasDiscoveredLocation returns a boolean if a field has been set.

### GetFirstAppearedOn

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetFirstAppearedOn() V1AlarmHistoryGet200ResponseHistoryInnerTime`

GetFirstAppearedOn returns the FirstAppearedOn field if non-nil, zero value otherwise.

### GetFirstAppearedOnOk

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetFirstAppearedOnOk() (*V1AlarmHistoryGet200ResponseHistoryInnerTime, bool)`

GetFirstAppearedOnOk returns a tuple with the FirstAppearedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAppearedOn

`func (o *V1DevicesBringupPost200ResponseSummariesInner) SetFirstAppearedOn(v V1AlarmHistoryGet200ResponseHistoryInnerTime)`

SetFirstAppearedOn sets FirstAppearedOn field to given value.

### HasFirstAppearedOn

`func (o *V1DevicesBringupPost200ResponseSummariesInner) HasFirstAppearedOn() bool`

HasFirstAppearedOn returns a boolean if a field has been set.

### GetIpDetected

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetIpDetected() string`

GetIpDetected returns the IpDetected field if non-nil, zero value otherwise.

### GetIpDetectedOk

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetIpDetectedOk() (*string, bool)`

GetIpDetectedOk returns a tuple with the IpDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDetected

`func (o *V1DevicesBringupPost200ResponseSummariesInner) SetIpDetected(v string)`

SetIpDetected sets IpDetected field to given value.

### HasIpDetected

`func (o *V1DevicesBringupPost200ResponseSummariesInner) HasIpDetected() bool`

HasIpDetected returns a boolean if a field has been set.

### GetIsNew

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *V1DevicesBringupPost200ResponseSummariesInner) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.

### HasIsNew

`func (o *V1DevicesBringupPost200ResponseSummariesInner) HasIsNew() bool`

HasIsNew returns a boolean if a field has been set.

### GetState

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V1DevicesBringupPost200ResponseSummariesInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V1DevicesBringupPost200ResponseSummariesInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1DevicesBringupPost200ResponseSummariesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1DevicesBringupPost200ResponseSummariesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1DevicesBringupPost200ResponseSummariesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


