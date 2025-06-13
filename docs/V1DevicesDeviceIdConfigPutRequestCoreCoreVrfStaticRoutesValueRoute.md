# V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminDistance** | Pointer to **int32** |  | [optional] 
**AdministrativeDistance** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAdministrativeDistance**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAdministrativeDistance.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DestinationPrefix** | Pointer to **string** |  | [optional] 
**IpVersion** | Pointer to **int32** |  | [optional] 
**NextHop** | Pointer to [**V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteNextHop**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteNextHop.md) |  | [optional] 
**NextHops** | Pointer to [**[]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteNextHop**](V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteNextHop.md) |  | [optional] 

## Methods

### NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute

`func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute`

NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteWithDefaults

`func NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteWithDefaults() *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute`

NewV1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteWithDefaults instantiates a new V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminDistance

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetAdminDistance() int32`

GetAdminDistance returns the AdminDistance field if non-nil, zero value otherwise.

### GetAdminDistanceOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetAdminDistanceOk() (*int32, bool)`

GetAdminDistanceOk returns a tuple with the AdminDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDistance

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) SetAdminDistance(v int32)`

SetAdminDistance sets AdminDistance field to given value.

### HasAdminDistance

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) HasAdminDistance() bool`

HasAdminDistance returns a boolean if a field has been set.

### GetAdministrativeDistance

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetAdministrativeDistance() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAdministrativeDistance`

GetAdministrativeDistance returns the AdministrativeDistance field if non-nil, zero value otherwise.

### GetAdministrativeDistanceOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetAdministrativeDistanceOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAdministrativeDistance, bool)`

GetAdministrativeDistanceOk returns a tuple with the AdministrativeDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeDistance

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) SetAdministrativeDistance(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementActionsValueActionAdministrativeDistance)`

SetAdministrativeDistance sets AdministrativeDistance field to given value.

### HasAdministrativeDistance

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) HasAdministrativeDistance() bool`

HasAdministrativeDistance returns a boolean if a field has been set.

### GetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinationPrefix

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetDestinationPrefix() string`

GetDestinationPrefix returns the DestinationPrefix field if non-nil, zero value otherwise.

### GetDestinationPrefixOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetDestinationPrefixOk() (*string, bool)`

GetDestinationPrefixOk returns a tuple with the DestinationPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPrefix

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) SetDestinationPrefix(v string)`

SetDestinationPrefix sets DestinationPrefix field to given value.

### HasDestinationPrefix

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) HasDestinationPrefix() bool`

HasDestinationPrefix returns a boolean if a field has been set.

### GetIpVersion

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetNextHop

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetNextHop() V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteNextHop`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetNextHopOk() (*V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteNextHop, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) SetNextHop(v V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteNextHop)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetNextHops

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetNextHops() []V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteNextHop`

GetNextHops returns the NextHops field if non-nil, zero value otherwise.

### GetNextHopsOk

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) GetNextHopsOk() (*[]V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteNextHop, bool)`

GetNextHopsOk returns a tuple with the NextHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHops

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) SetNextHops(v []V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRouteNextHop)`

SetNextHops sets NextHops field to given value.

### HasNextHops

`func (o *V1DevicesDeviceIdConfigPutRequestCoreCoreVrfStaticRoutesValueRoute) HasNextHops() bool`

HasNextHops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


