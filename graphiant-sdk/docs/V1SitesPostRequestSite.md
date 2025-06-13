# V1SitesPostRequestSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalPrefixSetOps** | Pointer to **map[string]string** |  | [optional] 
**IpfixExporterOps** | Pointer to **map[string]string** |  | [optional] 
**IpfixExporterOpsV2** | Pointer to [**map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value**](V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value.md) |  | [optional] 
**Location** | Pointer to [**V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation**](V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**PrefixSetOps** | Pointer to **map[string]string** |  | [optional] 
**RouteTag** | Pointer to [**V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTagEntry**](V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTagEntry.md) |  | [optional] 
**RoutingPolicyOps** | Pointer to **map[string]string** |  | [optional] 
**SnmpOps** | Pointer to **map[string]string** |  | [optional] 
**SyslogServerOps** | Pointer to **map[string]string** |  | [optional] 
**SyslogServerOpsV2** | Pointer to [**map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value**](V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value.md) |  | [optional] 
**TrafficPolicyOps** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewV1SitesPostRequestSite

`func NewV1SitesPostRequestSite() *V1SitesPostRequestSite`

NewV1SitesPostRequestSite instantiates a new V1SitesPostRequestSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SitesPostRequestSiteWithDefaults

`func NewV1SitesPostRequestSiteWithDefaults() *V1SitesPostRequestSite`

NewV1SitesPostRequestSiteWithDefaults instantiates a new V1SitesPostRequestSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalPrefixSetOps

`func (o *V1SitesPostRequestSite) GetGlobalPrefixSetOps() map[string]string`

GetGlobalPrefixSetOps returns the GlobalPrefixSetOps field if non-nil, zero value otherwise.

### GetGlobalPrefixSetOpsOk

`func (o *V1SitesPostRequestSite) GetGlobalPrefixSetOpsOk() (*map[string]string, bool)`

GetGlobalPrefixSetOpsOk returns a tuple with the GlobalPrefixSetOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPrefixSetOps

`func (o *V1SitesPostRequestSite) SetGlobalPrefixSetOps(v map[string]string)`

SetGlobalPrefixSetOps sets GlobalPrefixSetOps field to given value.

### HasGlobalPrefixSetOps

`func (o *V1SitesPostRequestSite) HasGlobalPrefixSetOps() bool`

HasGlobalPrefixSetOps returns a boolean if a field has been set.

### GetIpfixExporterOps

`func (o *V1SitesPostRequestSite) GetIpfixExporterOps() map[string]string`

GetIpfixExporterOps returns the IpfixExporterOps field if non-nil, zero value otherwise.

### GetIpfixExporterOpsOk

`func (o *V1SitesPostRequestSite) GetIpfixExporterOpsOk() (*map[string]string, bool)`

GetIpfixExporterOpsOk returns a tuple with the IpfixExporterOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporterOps

`func (o *V1SitesPostRequestSite) SetIpfixExporterOps(v map[string]string)`

SetIpfixExporterOps sets IpfixExporterOps field to given value.

### HasIpfixExporterOps

`func (o *V1SitesPostRequestSite) HasIpfixExporterOps() bool`

HasIpfixExporterOps returns a boolean if a field has been set.

### GetIpfixExporterOpsV2

`func (o *V1SitesPostRequestSite) GetIpfixExporterOpsV2() map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value`

GetIpfixExporterOpsV2 returns the IpfixExporterOpsV2 field if non-nil, zero value otherwise.

### GetIpfixExporterOpsV2Ok

`func (o *V1SitesPostRequestSite) GetIpfixExporterOpsV2Ok() (*map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value, bool)`

GetIpfixExporterOpsV2Ok returns a tuple with the IpfixExporterOpsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporterOpsV2

`func (o *V1SitesPostRequestSite) SetIpfixExporterOpsV2(v map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value)`

SetIpfixExporterOpsV2 sets IpfixExporterOpsV2 field to given value.

### HasIpfixExporterOpsV2

`func (o *V1SitesPostRequestSite) HasIpfixExporterOpsV2() bool`

HasIpfixExporterOpsV2 returns a boolean if a field has been set.

### GetLocation

`func (o *V1SitesPostRequestSite) GetLocation() V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V1SitesPostRequestSite) GetLocationOk() (*V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V1SitesPostRequestSite) SetLocation(v V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V1SitesPostRequestSite) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *V1SitesPostRequestSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1SitesPostRequestSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1SitesPostRequestSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1SitesPostRequestSite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *V1SitesPostRequestSite) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *V1SitesPostRequestSite) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *V1SitesPostRequestSite) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *V1SitesPostRequestSite) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPrefixSetOps

`func (o *V1SitesPostRequestSite) GetPrefixSetOps() map[string]string`

GetPrefixSetOps returns the PrefixSetOps field if non-nil, zero value otherwise.

### GetPrefixSetOpsOk

`func (o *V1SitesPostRequestSite) GetPrefixSetOpsOk() (*map[string]string, bool)`

GetPrefixSetOpsOk returns a tuple with the PrefixSetOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSetOps

`func (o *V1SitesPostRequestSite) SetPrefixSetOps(v map[string]string)`

SetPrefixSetOps sets PrefixSetOps field to given value.

### HasPrefixSetOps

`func (o *V1SitesPostRequestSite) HasPrefixSetOps() bool`

HasPrefixSetOps returns a boolean if a field has been set.

### GetRouteTag

`func (o *V1SitesPostRequestSite) GetRouteTag() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTagEntry`

GetRouteTag returns the RouteTag field if non-nil, zero value otherwise.

### GetRouteTagOk

`func (o *V1SitesPostRequestSite) GetRouteTagOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTagEntry, bool)`

GetRouteTagOk returns a tuple with the RouteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTag

`func (o *V1SitesPostRequestSite) SetRouteTag(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTagEntry)`

SetRouteTag sets RouteTag field to given value.

### HasRouteTag

`func (o *V1SitesPostRequestSite) HasRouteTag() bool`

HasRouteTag returns a boolean if a field has been set.

### GetRoutingPolicyOps

`func (o *V1SitesPostRequestSite) GetRoutingPolicyOps() map[string]string`

GetRoutingPolicyOps returns the RoutingPolicyOps field if non-nil, zero value otherwise.

### GetRoutingPolicyOpsOk

`func (o *V1SitesPostRequestSite) GetRoutingPolicyOpsOk() (*map[string]string, bool)`

GetRoutingPolicyOpsOk returns a tuple with the RoutingPolicyOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicyOps

`func (o *V1SitesPostRequestSite) SetRoutingPolicyOps(v map[string]string)`

SetRoutingPolicyOps sets RoutingPolicyOps field to given value.

### HasRoutingPolicyOps

`func (o *V1SitesPostRequestSite) HasRoutingPolicyOps() bool`

HasRoutingPolicyOps returns a boolean if a field has been set.

### GetSnmpOps

`func (o *V1SitesPostRequestSite) GetSnmpOps() map[string]string`

GetSnmpOps returns the SnmpOps field if non-nil, zero value otherwise.

### GetSnmpOpsOk

`func (o *V1SitesPostRequestSite) GetSnmpOpsOk() (*map[string]string, bool)`

GetSnmpOpsOk returns a tuple with the SnmpOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpOps

`func (o *V1SitesPostRequestSite) SetSnmpOps(v map[string]string)`

SetSnmpOps sets SnmpOps field to given value.

### HasSnmpOps

`func (o *V1SitesPostRequestSite) HasSnmpOps() bool`

HasSnmpOps returns a boolean if a field has been set.

### GetSyslogServerOps

`func (o *V1SitesPostRequestSite) GetSyslogServerOps() map[string]string`

GetSyslogServerOps returns the SyslogServerOps field if non-nil, zero value otherwise.

### GetSyslogServerOpsOk

`func (o *V1SitesPostRequestSite) GetSyslogServerOpsOk() (*map[string]string, bool)`

GetSyslogServerOpsOk returns a tuple with the SyslogServerOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServerOps

`func (o *V1SitesPostRequestSite) SetSyslogServerOps(v map[string]string)`

SetSyslogServerOps sets SyslogServerOps field to given value.

### HasSyslogServerOps

`func (o *V1SitesPostRequestSite) HasSyslogServerOps() bool`

HasSyslogServerOps returns a boolean if a field has been set.

### GetSyslogServerOpsV2

`func (o *V1SitesPostRequestSite) GetSyslogServerOpsV2() map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value`

GetSyslogServerOpsV2 returns the SyslogServerOpsV2 field if non-nil, zero value otherwise.

### GetSyslogServerOpsV2Ok

`func (o *V1SitesPostRequestSite) GetSyslogServerOpsV2Ok() (*map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value, bool)`

GetSyslogServerOpsV2Ok returns a tuple with the SyslogServerOpsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServerOpsV2

`func (o *V1SitesPostRequestSite) SetSyslogServerOpsV2(v map[string]V1GlobalConfigSitePostRequestIpfixExporterOpsV2Value)`

SetSyslogServerOpsV2 sets SyslogServerOpsV2 field to given value.

### HasSyslogServerOpsV2

`func (o *V1SitesPostRequestSite) HasSyslogServerOpsV2() bool`

HasSyslogServerOpsV2 returns a boolean if a field has been set.

### GetTrafficPolicyOps

`func (o *V1SitesPostRequestSite) GetTrafficPolicyOps() map[string]string`

GetTrafficPolicyOps returns the TrafficPolicyOps field if non-nil, zero value otherwise.

### GetTrafficPolicyOpsOk

`func (o *V1SitesPostRequestSite) GetTrafficPolicyOpsOk() (*map[string]string, bool)`

GetTrafficPolicyOpsOk returns a tuple with the TrafficPolicyOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPolicyOps

`func (o *V1SitesPostRequestSite) SetTrafficPolicyOps(v map[string]string)`

SetTrafficPolicyOps sets TrafficPolicyOps field to given value.

### HasTrafficPolicyOps

`func (o *V1SitesPostRequestSite) HasTrafficPolicyOps() bool`

HasTrafficPolicyOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


