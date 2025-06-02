# V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue**](V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue.md) |  | [optional] 
**AutoIpv4** | Pointer to **bool** |  | [optional] 
**AutoIpv6** | Pointer to **bool** |  | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**LanSegment** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint

`func NewV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint() *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint`

NewV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointWithDefaults

`func NewV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointWithDefaults() *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint`

NewV1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointWithDefaults instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAddresses() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAddressesOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetAddresses(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpointAddressesValue)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAutoIpv4

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAutoIpv4() bool`

GetAutoIpv4 returns the AutoIpv4 field if non-nil, zero value otherwise.

### GetAutoIpv4Ok

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAutoIpv4Ok() (*bool, bool)`

GetAutoIpv4Ok returns a tuple with the AutoIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoIpv4

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetAutoIpv4(v bool)`

SetAutoIpv4 sets AutoIpv4 field to given value.

### HasAutoIpv4

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasAutoIpv4() bool`

HasAutoIpv4 returns a boolean if a field has been set.

### GetAutoIpv6

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAutoIpv6() bool`

GetAutoIpv6 returns the AutoIpv6 field if non-nil, zero value otherwise.

### GetAutoIpv6Ok

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetAutoIpv6Ok() (*bool, bool)`

GetAutoIpv6Ok returns a tuple with the AutoIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoIpv6

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetAutoIpv6(v bool)`

SetAutoIpv6 sets AutoIpv6 field to given value.

### HasAutoIpv6

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasAutoIpv6() bool`

HasAutoIpv6 returns a boolean if a field has been set.

### GetInterface

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetLanSegment

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetLanSegment() string`

GetLanSegment returns the LanSegment field if non-nil, zero value otherwise.

### GetLanSegmentOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetLanSegmentOk() (*string, bool)`

GetLanSegmentOk returns a tuple with the LanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanSegment

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetLanSegment(v string)`

SetLanSegment sets LanSegment field to given value.

### HasLanSegment

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasLanSegment() bool`

HasLanSegment returns a boolean if a field has been set.

### GetName

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValueEngineEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


