# V1GlobalConfigPatchRequestTrafficPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpiApplications** | Pointer to [**map[string]V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValue**](V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValue.md) |  | [optional] 
**NetworkLists** | Pointer to [**map[string]V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue**](V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue.md) |  | [optional] 
**PortLists** | Pointer to [**map[string]V1GlobalConfigPatchRequestTrafficPoliciesPortListsValue**](V1GlobalConfigPatchRequestTrafficPoliciesPortListsValue.md) |  | [optional] 
**SecurityRulesets** | Pointer to [**map[string]V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValue**](V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValue.md) |  | [optional] 
**TrafficRulesets** | Pointer to [**map[string]V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValue**](V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValue.md) |  | [optional] 
**ZoneFirewalls** | Pointer to [**map[string]V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValue**](V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValue.md) |  | [optional] 
**Zones** | Pointer to [**map[string]V1GlobalConfigPatchRequestTrafficPoliciesZonesValue**](V1GlobalConfigPatchRequestTrafficPoliciesZonesValue.md) |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestTrafficPolicies

`func NewV1GlobalConfigPatchRequestTrafficPolicies() *V1GlobalConfigPatchRequestTrafficPolicies`

NewV1GlobalConfigPatchRequestTrafficPolicies instantiates a new V1GlobalConfigPatchRequestTrafficPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestTrafficPoliciesWithDefaults

`func NewV1GlobalConfigPatchRequestTrafficPoliciesWithDefaults() *V1GlobalConfigPatchRequestTrafficPolicies`

NewV1GlobalConfigPatchRequestTrafficPoliciesWithDefaults instantiates a new V1GlobalConfigPatchRequestTrafficPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpiApplications

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetDpiApplications() map[string]V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValue`

GetDpiApplications returns the DpiApplications field if non-nil, zero value otherwise.

### GetDpiApplicationsOk

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetDpiApplicationsOk() (*map[string]V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValue, bool)`

GetDpiApplicationsOk returns a tuple with the DpiApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpiApplications

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) SetDpiApplications(v map[string]V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValue)`

SetDpiApplications sets DpiApplications field to given value.

### HasDpiApplications

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) HasDpiApplications() bool`

HasDpiApplications returns a boolean if a field has been set.

### GetNetworkLists

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetNetworkLists() map[string]V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue`

GetNetworkLists returns the NetworkLists field if non-nil, zero value otherwise.

### GetNetworkListsOk

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetNetworkListsOk() (*map[string]V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue, bool)`

GetNetworkListsOk returns a tuple with the NetworkLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLists

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) SetNetworkLists(v map[string]V1GlobalConfigPatchRequestTrafficPoliciesNetworkListsValue)`

SetNetworkLists sets NetworkLists field to given value.

### HasNetworkLists

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) HasNetworkLists() bool`

HasNetworkLists returns a boolean if a field has been set.

### GetPortLists

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetPortLists() map[string]V1GlobalConfigPatchRequestTrafficPoliciesPortListsValue`

GetPortLists returns the PortLists field if non-nil, zero value otherwise.

### GetPortListsOk

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetPortListsOk() (*map[string]V1GlobalConfigPatchRequestTrafficPoliciesPortListsValue, bool)`

GetPortListsOk returns a tuple with the PortLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLists

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) SetPortLists(v map[string]V1GlobalConfigPatchRequestTrafficPoliciesPortListsValue)`

SetPortLists sets PortLists field to given value.

### HasPortLists

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) HasPortLists() bool`

HasPortLists returns a boolean if a field has been set.

### GetSecurityRulesets

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetSecurityRulesets() map[string]V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValue`

GetSecurityRulesets returns the SecurityRulesets field if non-nil, zero value otherwise.

### GetSecurityRulesetsOk

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetSecurityRulesetsOk() (*map[string]V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValue, bool)`

GetSecurityRulesetsOk returns a tuple with the SecurityRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRulesets

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) SetSecurityRulesets(v map[string]V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValue)`

SetSecurityRulesets sets SecurityRulesets field to given value.

### HasSecurityRulesets

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) HasSecurityRulesets() bool`

HasSecurityRulesets returns a boolean if a field has been set.

### GetTrafficRulesets

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetTrafficRulesets() map[string]V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValue`

GetTrafficRulesets returns the TrafficRulesets field if non-nil, zero value otherwise.

### GetTrafficRulesetsOk

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetTrafficRulesetsOk() (*map[string]V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValue, bool)`

GetTrafficRulesetsOk returns a tuple with the TrafficRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRulesets

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) SetTrafficRulesets(v map[string]V1GlobalConfigPatchRequestTrafficPoliciesTrafficRulesetsValue)`

SetTrafficRulesets sets TrafficRulesets field to given value.

### HasTrafficRulesets

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) HasTrafficRulesets() bool`

HasTrafficRulesets returns a boolean if a field has been set.

### GetZoneFirewalls

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetZoneFirewalls() map[string]V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValue`

GetZoneFirewalls returns the ZoneFirewalls field if non-nil, zero value otherwise.

### GetZoneFirewallsOk

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetZoneFirewallsOk() (*map[string]V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValue, bool)`

GetZoneFirewallsOk returns a tuple with the ZoneFirewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFirewalls

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) SetZoneFirewalls(v map[string]V1GlobalConfigPatchRequestTrafficPoliciesZoneFirewallsValue)`

SetZoneFirewalls sets ZoneFirewalls field to given value.

### HasZoneFirewalls

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) HasZoneFirewalls() bool`

HasZoneFirewalls returns a boolean if a field has been set.

### GetZones

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetZones() map[string]V1GlobalConfigPatchRequestTrafficPoliciesZonesValue`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) GetZonesOk() (*map[string]V1GlobalConfigPatchRequestTrafficPoliciesZonesValue, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) SetZones(v map[string]V1GlobalConfigPatchRequestTrafficPoliciesZonesValue)`

SetZones sets Zones field to given value.

### HasZones

`func (o *V1GlobalConfigPatchRequestTrafficPolicies) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


