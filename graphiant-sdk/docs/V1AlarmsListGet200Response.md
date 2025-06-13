# V1AlarmsListGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alarms** | Pointer to [**[]V1AlarmsListGet200ResponseAlarmsInner**](V1AlarmsListGet200ResponseAlarmsInner.md) |  | [optional] 

## Methods

### NewV1AlarmsListGet200Response

`func NewV1AlarmsListGet200Response() *V1AlarmsListGet200Response`

NewV1AlarmsListGet200Response instantiates a new V1AlarmsListGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AlarmsListGet200ResponseWithDefaults

`func NewV1AlarmsListGet200ResponseWithDefaults() *V1AlarmsListGet200Response`

NewV1AlarmsListGet200ResponseWithDefaults instantiates a new V1AlarmsListGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarms

`func (o *V1AlarmsListGet200Response) GetAlarms() []V1AlarmsListGet200ResponseAlarmsInner`

GetAlarms returns the Alarms field if non-nil, zero value otherwise.

### GetAlarmsOk

`func (o *V1AlarmsListGet200Response) GetAlarmsOk() (*[]V1AlarmsListGet200ResponseAlarmsInner, bool)`

GetAlarmsOk returns a tuple with the Alarms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarms

`func (o *V1AlarmsListGet200Response) SetAlarms(v []V1AlarmsListGet200ResponseAlarmsInner)`

SetAlarms sets Alarms field to given value.

### HasAlarms

`func (o *V1AlarmsListGet200Response) HasAlarms() bool`

HasAlarms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


