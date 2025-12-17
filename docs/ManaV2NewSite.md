# ManaV2NewSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalPrefixSetOps** | Pointer to **map[string]string** |  | [optional] 
**IpfixExporterOps** | Pointer to **map[string]string** |  | [optional] 
**IpfixExporterOpsV2** | Pointer to [**map[string]ManaV2GlobalObjectOperationConfig**](ManaV2GlobalObjectOperationConfig.md) |  | [optional] 
**Location** | Pointer to [**ManaV2Location**](ManaV2Location.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**NtpOps** | Pointer to **map[string]string** |  | [optional] 
**PrefixSetOps** | Pointer to **map[string]string** |  | [optional] 
**RouteTag** | Pointer to [**ManaV2RouteTag**](ManaV2RouteTag.md) |  | [optional] 
**RoutingPolicyOps** | Pointer to **map[string]string** |  | [optional] 
**SnmpOps** | Pointer to **map[string]string** |  | [optional] 
**SyslogServerOps** | Pointer to **map[string]string** |  | [optional] 
**SyslogServerOpsV2** | Pointer to [**map[string]ManaV2GlobalObjectOperationConfig**](ManaV2GlobalObjectOperationConfig.md) |  | [optional] 
**TrafficPolicyOps** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManaV2NewSite

`func NewManaV2NewSite() *ManaV2NewSite`

NewManaV2NewSite instantiates a new ManaV2NewSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2NewSiteWithDefaults

`func NewManaV2NewSiteWithDefaults() *ManaV2NewSite`

NewManaV2NewSiteWithDefaults instantiates a new ManaV2NewSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalPrefixSetOps

`func (o *ManaV2NewSite) GetGlobalPrefixSetOps() map[string]string`

GetGlobalPrefixSetOps returns the GlobalPrefixSetOps field if non-nil, zero value otherwise.

### GetGlobalPrefixSetOpsOk

`func (o *ManaV2NewSite) GetGlobalPrefixSetOpsOk() (*map[string]string, bool)`

GetGlobalPrefixSetOpsOk returns a tuple with the GlobalPrefixSetOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPrefixSetOps

`func (o *ManaV2NewSite) SetGlobalPrefixSetOps(v map[string]string)`

SetGlobalPrefixSetOps sets GlobalPrefixSetOps field to given value.

### HasGlobalPrefixSetOps

`func (o *ManaV2NewSite) HasGlobalPrefixSetOps() bool`

HasGlobalPrefixSetOps returns a boolean if a field has been set.

### GetIpfixExporterOps

`func (o *ManaV2NewSite) GetIpfixExporterOps() map[string]string`

GetIpfixExporterOps returns the IpfixExporterOps field if non-nil, zero value otherwise.

### GetIpfixExporterOpsOk

`func (o *ManaV2NewSite) GetIpfixExporterOpsOk() (*map[string]string, bool)`

GetIpfixExporterOpsOk returns a tuple with the IpfixExporterOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporterOps

`func (o *ManaV2NewSite) SetIpfixExporterOps(v map[string]string)`

SetIpfixExporterOps sets IpfixExporterOps field to given value.

### HasIpfixExporterOps

`func (o *ManaV2NewSite) HasIpfixExporterOps() bool`

HasIpfixExporterOps returns a boolean if a field has been set.

### GetIpfixExporterOpsV2

`func (o *ManaV2NewSite) GetIpfixExporterOpsV2() map[string]ManaV2GlobalObjectOperationConfig`

GetIpfixExporterOpsV2 returns the IpfixExporterOpsV2 field if non-nil, zero value otherwise.

### GetIpfixExporterOpsV2Ok

`func (o *ManaV2NewSite) GetIpfixExporterOpsV2Ok() (*map[string]ManaV2GlobalObjectOperationConfig, bool)`

GetIpfixExporterOpsV2Ok returns a tuple with the IpfixExporterOpsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfixExporterOpsV2

`func (o *ManaV2NewSite) SetIpfixExporterOpsV2(v map[string]ManaV2GlobalObjectOperationConfig)`

SetIpfixExporterOpsV2 sets IpfixExporterOpsV2 field to given value.

### HasIpfixExporterOpsV2

`func (o *ManaV2NewSite) HasIpfixExporterOpsV2() bool`

HasIpfixExporterOpsV2 returns a boolean if a field has been set.

### GetLocation

`func (o *ManaV2NewSite) GetLocation() ManaV2Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ManaV2NewSite) GetLocationOk() (*ManaV2Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ManaV2NewSite) SetLocation(v ManaV2Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ManaV2NewSite) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *ManaV2NewSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManaV2NewSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManaV2NewSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManaV2NewSite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *ManaV2NewSite) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ManaV2NewSite) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ManaV2NewSite) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ManaV2NewSite) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNtpOps

`func (o *ManaV2NewSite) GetNtpOps() map[string]string`

GetNtpOps returns the NtpOps field if non-nil, zero value otherwise.

### GetNtpOpsOk

`func (o *ManaV2NewSite) GetNtpOpsOk() (*map[string]string, bool)`

GetNtpOpsOk returns a tuple with the NtpOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpOps

`func (o *ManaV2NewSite) SetNtpOps(v map[string]string)`

SetNtpOps sets NtpOps field to given value.

### HasNtpOps

`func (o *ManaV2NewSite) HasNtpOps() bool`

HasNtpOps returns a boolean if a field has been set.

### GetPrefixSetOps

`func (o *ManaV2NewSite) GetPrefixSetOps() map[string]string`

GetPrefixSetOps returns the PrefixSetOps field if non-nil, zero value otherwise.

### GetPrefixSetOpsOk

`func (o *ManaV2NewSite) GetPrefixSetOpsOk() (*map[string]string, bool)`

GetPrefixSetOpsOk returns a tuple with the PrefixSetOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSetOps

`func (o *ManaV2NewSite) SetPrefixSetOps(v map[string]string)`

SetPrefixSetOps sets PrefixSetOps field to given value.

### HasPrefixSetOps

`func (o *ManaV2NewSite) HasPrefixSetOps() bool`

HasPrefixSetOps returns a boolean if a field has been set.

### GetRouteTag

`func (o *ManaV2NewSite) GetRouteTag() ManaV2RouteTag`

GetRouteTag returns the RouteTag field if non-nil, zero value otherwise.

### GetRouteTagOk

`func (o *ManaV2NewSite) GetRouteTagOk() (*ManaV2RouteTag, bool)`

GetRouteTagOk returns a tuple with the RouteTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTag

`func (o *ManaV2NewSite) SetRouteTag(v ManaV2RouteTag)`

SetRouteTag sets RouteTag field to given value.

### HasRouteTag

`func (o *ManaV2NewSite) HasRouteTag() bool`

HasRouteTag returns a boolean if a field has been set.

### GetRoutingPolicyOps

`func (o *ManaV2NewSite) GetRoutingPolicyOps() map[string]string`

GetRoutingPolicyOps returns the RoutingPolicyOps field if non-nil, zero value otherwise.

### GetRoutingPolicyOpsOk

`func (o *ManaV2NewSite) GetRoutingPolicyOpsOk() (*map[string]string, bool)`

GetRoutingPolicyOpsOk returns a tuple with the RoutingPolicyOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicyOps

`func (o *ManaV2NewSite) SetRoutingPolicyOps(v map[string]string)`

SetRoutingPolicyOps sets RoutingPolicyOps field to given value.

### HasRoutingPolicyOps

`func (o *ManaV2NewSite) HasRoutingPolicyOps() bool`

HasRoutingPolicyOps returns a boolean if a field has been set.

### GetSnmpOps

`func (o *ManaV2NewSite) GetSnmpOps() map[string]string`

GetSnmpOps returns the SnmpOps field if non-nil, zero value otherwise.

### GetSnmpOpsOk

`func (o *ManaV2NewSite) GetSnmpOpsOk() (*map[string]string, bool)`

GetSnmpOpsOk returns a tuple with the SnmpOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpOps

`func (o *ManaV2NewSite) SetSnmpOps(v map[string]string)`

SetSnmpOps sets SnmpOps field to given value.

### HasSnmpOps

`func (o *ManaV2NewSite) HasSnmpOps() bool`

HasSnmpOps returns a boolean if a field has been set.

### GetSyslogServerOps

`func (o *ManaV2NewSite) GetSyslogServerOps() map[string]string`

GetSyslogServerOps returns the SyslogServerOps field if non-nil, zero value otherwise.

### GetSyslogServerOpsOk

`func (o *ManaV2NewSite) GetSyslogServerOpsOk() (*map[string]string, bool)`

GetSyslogServerOpsOk returns a tuple with the SyslogServerOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServerOps

`func (o *ManaV2NewSite) SetSyslogServerOps(v map[string]string)`

SetSyslogServerOps sets SyslogServerOps field to given value.

### HasSyslogServerOps

`func (o *ManaV2NewSite) HasSyslogServerOps() bool`

HasSyslogServerOps returns a boolean if a field has been set.

### GetSyslogServerOpsV2

`func (o *ManaV2NewSite) GetSyslogServerOpsV2() map[string]ManaV2GlobalObjectOperationConfig`

GetSyslogServerOpsV2 returns the SyslogServerOpsV2 field if non-nil, zero value otherwise.

### GetSyslogServerOpsV2Ok

`func (o *ManaV2NewSite) GetSyslogServerOpsV2Ok() (*map[string]ManaV2GlobalObjectOperationConfig, bool)`

GetSyslogServerOpsV2Ok returns a tuple with the SyslogServerOpsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServerOpsV2

`func (o *ManaV2NewSite) SetSyslogServerOpsV2(v map[string]ManaV2GlobalObjectOperationConfig)`

SetSyslogServerOpsV2 sets SyslogServerOpsV2 field to given value.

### HasSyslogServerOpsV2

`func (o *ManaV2NewSite) HasSyslogServerOpsV2() bool`

HasSyslogServerOpsV2 returns a boolean if a field has been set.

### GetTrafficPolicyOps

`func (o *ManaV2NewSite) GetTrafficPolicyOps() map[string]string`

GetTrafficPolicyOps returns the TrafficPolicyOps field if non-nil, zero value otherwise.

### GetTrafficPolicyOpsOk

`func (o *ManaV2NewSite) GetTrafficPolicyOpsOk() (*map[string]string, bool)`

GetTrafficPolicyOpsOk returns a tuple with the TrafficPolicyOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPolicyOps

`func (o *ManaV2NewSite) SetTrafficPolicyOps(v map[string]string)`

SetTrafficPolicyOps sets TrafficPolicyOps field to given value.

### HasTrafficPolicyOps

`func (o *ManaV2NewSite) HasTrafficPolicyOps() bool`

HasTrafficPolicyOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


