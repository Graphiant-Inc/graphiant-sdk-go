# V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamilies** | Pointer to **[]string** |  | [optional] 
**AdminDistance** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance.md) |  | [optional] 
**Areas** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue.md) |  | [optional] 
**Auto** | Pointer to **bool** |  | [optional] 
**DefaultOriginate** | Pointer to **string** |  | [optional] 
**Manual** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Redistribution** | Pointer to [**map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process

`func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process`

NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process`

NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamilies

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAddressFamilies() []string`

GetAddressFamilies returns the AddressFamilies field if non-nil, zero value otherwise.

### GetAddressFamiliesOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAddressFamiliesOk() (*[]string, bool)`

GetAddressFamiliesOk returns a tuple with the AddressFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamilies

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetAddressFamilies(v []string)`

SetAddressFamilies sets AddressFamilies field to given value.

### HasAddressFamilies

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasAddressFamilies() bool`

HasAddressFamilies returns a boolean if a field has been set.

### GetAdminDistance

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAdminDistance() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance`

GetAdminDistance returns the AdminDistance field if non-nil, zero value otherwise.

### GetAdminDistanceOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAdminDistanceOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance, bool)`

GetAdminDistanceOk returns a tuple with the AdminDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDistance

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetAdminDistance(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAdminDistance)`

SetAdminDistance sets AdminDistance field to given value.

### HasAdminDistance

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasAdminDistance() bool`

HasAdminDistance returns a boolean if a field has been set.

### GetAreas

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAreas() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAreasOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetAreas(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessAreasValue)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetAuto

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAuto() bool`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetAutoOk() (*bool, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetAuto(v bool)`

SetAuto sets Auto field to given value.

### HasAuto

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasAuto() bool`

HasAuto returns a boolean if a field has been set.

### GetDefaultOriginate

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetDefaultOriginate() string`

GetDefaultOriginate returns the DefaultOriginate field if non-nil, zero value otherwise.

### GetDefaultOriginateOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetDefaultOriginateOk() (*string, bool)`

GetDefaultOriginateOk returns a tuple with the DefaultOriginate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOriginate

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetDefaultOriginate(v string)`

SetDefaultOriginate sets DefaultOriginate field to given value.

### HasDefaultOriginate

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasDefaultOriginate() bool`

HasDefaultOriginate returns a boolean if a field has been set.

### GetManual

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetManual() string`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetManualOk() (*string, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetManual(v string)`

SetManual sets Manual field to given value.

### HasManual

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRedistribution

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetRedistribution() map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue`

GetRedistribution returns the Redistribution field if non-nil, zero value otherwise.

### GetRedistributionOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) GetRedistributionOk() (*map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue, bool)`

GetRedistributionOk returns a tuple with the Redistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedistribution

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) SetRedistribution(v map[string]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2ProcessRedistributionValue)`

SetRedistribution sets Redistribution field to given value.

### HasRedistribution

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfOspfv2Process) HasRedistribution() bool`

HasRedistribution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


