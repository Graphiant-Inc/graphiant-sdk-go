/*
Graphiant APIs

**To use the APIs:**  1) Login using `/api/v1/auth/login`   2) Copy the value of \"token\" in the response   3) Click the \"Authorize\" button   4) In the \"Value\" text field enter: `Bearer <your token>`   5) Click \"Authorize\"   6) All requests are now authorized.  **Token valid for 2 hours. If expired:**   - Login again, click \"Authorize\", paste new token.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graphiant_sdk

import (
	"encoding/json"
)

// checks if the V1GlobalConfigPatchRequestSnmpsValueConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestSnmpsValueConfig{}

// V1GlobalConfigPatchRequestSnmpsValueConfig struct for V1GlobalConfigPatchRequestSnmpsValueConfig
type V1GlobalConfigPatchRequestSnmpsValueConfig struct {
	Communities *map[string]V1PortalApikeysPostRequest `json:"communities,omitempty"`
	EngineAuthenTraps *bool `json:"engineAuthenTraps,omitempty"`
	EngineEnabled *bool `json:"engineEnabled,omitempty"`
	EngineEndpoints *map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue `json:"engineEndpoints,omitempty"`
	EngineHighCpuTraps *bool `json:"engineHighCpuTraps,omitempty"`
	EngineHighMemoryTraps *bool `json:"engineHighMemoryTraps,omitempty"`
	EngineIdAdminOctets *string `json:"engineIdAdminOctets,omitempty"`
	EngineIdAdminText *string `json:"engineIdAdminText,omitempty"`
	EngineIdIpv4 *string `json:"engineIdIpv4,omitempty"`
	EngineIdIpv6 *string `json:"engineIdIpv6,omitempty"`
	EngineIdMac *string `json:"engineIdMac,omitempty"`
	EngineIdRaw *string `json:"engineIdRaw,omitempty"`
	EngineLocalAcessV4 *bool `json:"engineLocalAcessV4,omitempty"`
	EngineLocalAcessV6 *bool `json:"engineLocalAcessV6,omitempty"`
	EngineUserHints *bool `json:"engineUserHints,omitempty"`
	EngineUserValidation *bool `json:"engineUserValidation,omitempty"`
	GlobalId *int64 `json:"globalId,omitempty"`
	IsGlobalSync *bool `json:"isGlobalSync,omitempty"`
	Name *string `json:"name,omitempty"`
	NotifyFilterProfiles *map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue `json:"notifyFilterProfiles,omitempty"`
	Targets *map[string]V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue `json:"targets,omitempty"`
	UsmLocalUsers *map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue `json:"usmLocalUsers,omitempty"`
	UsmRemoteUsers *map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue `json:"usmRemoteUsers,omitempty"`
	V2cEnabled *bool `json:"v2cEnabled,omitempty"`
	V3Enabled *bool `json:"v3Enabled,omitempty"`
	VacmGroups *map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue `json:"vacmGroups,omitempty"`
	VacmViews *map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue `json:"vacmViews,omitempty"`
}

// NewV1GlobalConfigPatchRequestSnmpsValueConfig instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestSnmpsValueConfig() *V1GlobalConfigPatchRequestSnmpsValueConfig {
	this := V1GlobalConfigPatchRequestSnmpsValueConfig{}
	return &this
}

// NewV1GlobalConfigPatchRequestSnmpsValueConfigWithDefaults instantiates a new V1GlobalConfigPatchRequestSnmpsValueConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestSnmpsValueConfigWithDefaults() *V1GlobalConfigPatchRequestSnmpsValueConfig {
	this := V1GlobalConfigPatchRequestSnmpsValueConfig{}
	return &this
}

// GetCommunities returns the Communities field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetCommunities() map[string]V1PortalApikeysPostRequest {
	if o == nil || IsNil(o.Communities) {
		var ret map[string]V1PortalApikeysPostRequest
		return ret
	}
	return *o.Communities
}

// GetCommunitiesOk returns a tuple with the Communities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetCommunitiesOk() (*map[string]V1PortalApikeysPostRequest, bool) {
	if o == nil || IsNil(o.Communities) {
		return nil, false
	}
	return o.Communities, true
}

// HasCommunities returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasCommunities() bool {
	if o != nil && !IsNil(o.Communities) {
		return true
	}

	return false
}

// SetCommunities gets a reference to the given map[string]V1PortalApikeysPostRequest and assigns it to the Communities field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetCommunities(v map[string]V1PortalApikeysPostRequest) {
	o.Communities = &v
}

// GetEngineAuthenTraps returns the EngineAuthenTraps field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineAuthenTraps() bool {
	if o == nil || IsNil(o.EngineAuthenTraps) {
		var ret bool
		return ret
	}
	return *o.EngineAuthenTraps
}

// GetEngineAuthenTrapsOk returns a tuple with the EngineAuthenTraps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineAuthenTrapsOk() (*bool, bool) {
	if o == nil || IsNil(o.EngineAuthenTraps) {
		return nil, false
	}
	return o.EngineAuthenTraps, true
}

// HasEngineAuthenTraps returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineAuthenTraps() bool {
	if o != nil && !IsNil(o.EngineAuthenTraps) {
		return true
	}

	return false
}

// SetEngineAuthenTraps gets a reference to the given bool and assigns it to the EngineAuthenTraps field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineAuthenTraps(v bool) {
	o.EngineAuthenTraps = &v
}

// GetEngineEnabled returns the EngineEnabled field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineEnabled() bool {
	if o == nil || IsNil(o.EngineEnabled) {
		var ret bool
		return ret
	}
	return *o.EngineEnabled
}

// GetEngineEnabledOk returns a tuple with the EngineEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EngineEnabled) {
		return nil, false
	}
	return o.EngineEnabled, true
}

// HasEngineEnabled returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineEnabled() bool {
	if o != nil && !IsNil(o.EngineEnabled) {
		return true
	}

	return false
}

// SetEngineEnabled gets a reference to the given bool and assigns it to the EngineEnabled field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineEnabled(v bool) {
	o.EngineEnabled = &v
}

// GetEngineEndpoints returns the EngineEndpoints field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineEndpoints() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue {
	if o == nil || IsNil(o.EngineEndpoints) {
		var ret map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue
		return ret
	}
	return *o.EngineEndpoints
}

// GetEngineEndpointsOk returns a tuple with the EngineEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineEndpointsOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue, bool) {
	if o == nil || IsNil(o.EngineEndpoints) {
		return nil, false
	}
	return o.EngineEndpoints, true
}

// HasEngineEndpoints returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineEndpoints() bool {
	if o != nil && !IsNil(o.EngineEndpoints) {
		return true
	}

	return false
}

// SetEngineEndpoints gets a reference to the given map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue and assigns it to the EngineEndpoints field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineEndpoints(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigEngineEndpointsValue) {
	o.EngineEndpoints = &v
}

// GetEngineHighCpuTraps returns the EngineHighCpuTraps field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineHighCpuTraps() bool {
	if o == nil || IsNil(o.EngineHighCpuTraps) {
		var ret bool
		return ret
	}
	return *o.EngineHighCpuTraps
}

// GetEngineHighCpuTrapsOk returns a tuple with the EngineHighCpuTraps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineHighCpuTrapsOk() (*bool, bool) {
	if o == nil || IsNil(o.EngineHighCpuTraps) {
		return nil, false
	}
	return o.EngineHighCpuTraps, true
}

// HasEngineHighCpuTraps returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineHighCpuTraps() bool {
	if o != nil && !IsNil(o.EngineHighCpuTraps) {
		return true
	}

	return false
}

// SetEngineHighCpuTraps gets a reference to the given bool and assigns it to the EngineHighCpuTraps field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineHighCpuTraps(v bool) {
	o.EngineHighCpuTraps = &v
}

// GetEngineHighMemoryTraps returns the EngineHighMemoryTraps field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineHighMemoryTraps() bool {
	if o == nil || IsNil(o.EngineHighMemoryTraps) {
		var ret bool
		return ret
	}
	return *o.EngineHighMemoryTraps
}

// GetEngineHighMemoryTrapsOk returns a tuple with the EngineHighMemoryTraps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineHighMemoryTrapsOk() (*bool, bool) {
	if o == nil || IsNil(o.EngineHighMemoryTraps) {
		return nil, false
	}
	return o.EngineHighMemoryTraps, true
}

// HasEngineHighMemoryTraps returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineHighMemoryTraps() bool {
	if o != nil && !IsNil(o.EngineHighMemoryTraps) {
		return true
	}

	return false
}

// SetEngineHighMemoryTraps gets a reference to the given bool and assigns it to the EngineHighMemoryTraps field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineHighMemoryTraps(v bool) {
	o.EngineHighMemoryTraps = &v
}

// GetEngineIdAdminOctets returns the EngineIdAdminOctets field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdAdminOctets() string {
	if o == nil || IsNil(o.EngineIdAdminOctets) {
		var ret string
		return ret
	}
	return *o.EngineIdAdminOctets
}

// GetEngineIdAdminOctetsOk returns a tuple with the EngineIdAdminOctets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdAdminOctetsOk() (*string, bool) {
	if o == nil || IsNil(o.EngineIdAdminOctets) {
		return nil, false
	}
	return o.EngineIdAdminOctets, true
}

// HasEngineIdAdminOctets returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdAdminOctets() bool {
	if o != nil && !IsNil(o.EngineIdAdminOctets) {
		return true
	}

	return false
}

// SetEngineIdAdminOctets gets a reference to the given string and assigns it to the EngineIdAdminOctets field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdAdminOctets(v string) {
	o.EngineIdAdminOctets = &v
}

// GetEngineIdAdminText returns the EngineIdAdminText field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdAdminText() string {
	if o == nil || IsNil(o.EngineIdAdminText) {
		var ret string
		return ret
	}
	return *o.EngineIdAdminText
}

// GetEngineIdAdminTextOk returns a tuple with the EngineIdAdminText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdAdminTextOk() (*string, bool) {
	if o == nil || IsNil(o.EngineIdAdminText) {
		return nil, false
	}
	return o.EngineIdAdminText, true
}

// HasEngineIdAdminText returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdAdminText() bool {
	if o != nil && !IsNil(o.EngineIdAdminText) {
		return true
	}

	return false
}

// SetEngineIdAdminText gets a reference to the given string and assigns it to the EngineIdAdminText field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdAdminText(v string) {
	o.EngineIdAdminText = &v
}

// GetEngineIdIpv4 returns the EngineIdIpv4 field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdIpv4() string {
	if o == nil || IsNil(o.EngineIdIpv4) {
		var ret string
		return ret
	}
	return *o.EngineIdIpv4
}

// GetEngineIdIpv4Ok returns a tuple with the EngineIdIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdIpv4Ok() (*string, bool) {
	if o == nil || IsNil(o.EngineIdIpv4) {
		return nil, false
	}
	return o.EngineIdIpv4, true
}

// HasEngineIdIpv4 returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdIpv4() bool {
	if o != nil && !IsNil(o.EngineIdIpv4) {
		return true
	}

	return false
}

// SetEngineIdIpv4 gets a reference to the given string and assigns it to the EngineIdIpv4 field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdIpv4(v string) {
	o.EngineIdIpv4 = &v
}

// GetEngineIdIpv6 returns the EngineIdIpv6 field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdIpv6() string {
	if o == nil || IsNil(o.EngineIdIpv6) {
		var ret string
		return ret
	}
	return *o.EngineIdIpv6
}

// GetEngineIdIpv6Ok returns a tuple with the EngineIdIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdIpv6Ok() (*string, bool) {
	if o == nil || IsNil(o.EngineIdIpv6) {
		return nil, false
	}
	return o.EngineIdIpv6, true
}

// HasEngineIdIpv6 returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdIpv6() bool {
	if o != nil && !IsNil(o.EngineIdIpv6) {
		return true
	}

	return false
}

// SetEngineIdIpv6 gets a reference to the given string and assigns it to the EngineIdIpv6 field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdIpv6(v string) {
	o.EngineIdIpv6 = &v
}

// GetEngineIdMac returns the EngineIdMac field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdMac() string {
	if o == nil || IsNil(o.EngineIdMac) {
		var ret string
		return ret
	}
	return *o.EngineIdMac
}

// GetEngineIdMacOk returns a tuple with the EngineIdMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdMacOk() (*string, bool) {
	if o == nil || IsNil(o.EngineIdMac) {
		return nil, false
	}
	return o.EngineIdMac, true
}

// HasEngineIdMac returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdMac() bool {
	if o != nil && !IsNil(o.EngineIdMac) {
		return true
	}

	return false
}

// SetEngineIdMac gets a reference to the given string and assigns it to the EngineIdMac field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdMac(v string) {
	o.EngineIdMac = &v
}

// GetEngineIdRaw returns the EngineIdRaw field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdRaw() string {
	if o == nil || IsNil(o.EngineIdRaw) {
		var ret string
		return ret
	}
	return *o.EngineIdRaw
}

// GetEngineIdRawOk returns a tuple with the EngineIdRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineIdRawOk() (*string, bool) {
	if o == nil || IsNil(o.EngineIdRaw) {
		return nil, false
	}
	return o.EngineIdRaw, true
}

// HasEngineIdRaw returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineIdRaw() bool {
	if o != nil && !IsNil(o.EngineIdRaw) {
		return true
	}

	return false
}

// SetEngineIdRaw gets a reference to the given string and assigns it to the EngineIdRaw field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineIdRaw(v string) {
	o.EngineIdRaw = &v
}

// GetEngineLocalAcessV4 returns the EngineLocalAcessV4 field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineLocalAcessV4() bool {
	if o == nil || IsNil(o.EngineLocalAcessV4) {
		var ret bool
		return ret
	}
	return *o.EngineLocalAcessV4
}

// GetEngineLocalAcessV4Ok returns a tuple with the EngineLocalAcessV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineLocalAcessV4Ok() (*bool, bool) {
	if o == nil || IsNil(o.EngineLocalAcessV4) {
		return nil, false
	}
	return o.EngineLocalAcessV4, true
}

// HasEngineLocalAcessV4 returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineLocalAcessV4() bool {
	if o != nil && !IsNil(o.EngineLocalAcessV4) {
		return true
	}

	return false
}

// SetEngineLocalAcessV4 gets a reference to the given bool and assigns it to the EngineLocalAcessV4 field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineLocalAcessV4(v bool) {
	o.EngineLocalAcessV4 = &v
}

// GetEngineLocalAcessV6 returns the EngineLocalAcessV6 field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineLocalAcessV6() bool {
	if o == nil || IsNil(o.EngineLocalAcessV6) {
		var ret bool
		return ret
	}
	return *o.EngineLocalAcessV6
}

// GetEngineLocalAcessV6Ok returns a tuple with the EngineLocalAcessV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineLocalAcessV6Ok() (*bool, bool) {
	if o == nil || IsNil(o.EngineLocalAcessV6) {
		return nil, false
	}
	return o.EngineLocalAcessV6, true
}

// HasEngineLocalAcessV6 returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineLocalAcessV6() bool {
	if o != nil && !IsNil(o.EngineLocalAcessV6) {
		return true
	}

	return false
}

// SetEngineLocalAcessV6 gets a reference to the given bool and assigns it to the EngineLocalAcessV6 field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineLocalAcessV6(v bool) {
	o.EngineLocalAcessV6 = &v
}

// GetEngineUserHints returns the EngineUserHints field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineUserHints() bool {
	if o == nil || IsNil(o.EngineUserHints) {
		var ret bool
		return ret
	}
	return *o.EngineUserHints
}

// GetEngineUserHintsOk returns a tuple with the EngineUserHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineUserHintsOk() (*bool, bool) {
	if o == nil || IsNil(o.EngineUserHints) {
		return nil, false
	}
	return o.EngineUserHints, true
}

// HasEngineUserHints returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineUserHints() bool {
	if o != nil && !IsNil(o.EngineUserHints) {
		return true
	}

	return false
}

// SetEngineUserHints gets a reference to the given bool and assigns it to the EngineUserHints field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineUserHints(v bool) {
	o.EngineUserHints = &v
}

// GetEngineUserValidation returns the EngineUserValidation field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineUserValidation() bool {
	if o == nil || IsNil(o.EngineUserValidation) {
		var ret bool
		return ret
	}
	return *o.EngineUserValidation
}

// GetEngineUserValidationOk returns a tuple with the EngineUserValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetEngineUserValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.EngineUserValidation) {
		return nil, false
	}
	return o.EngineUserValidation, true
}

// HasEngineUserValidation returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasEngineUserValidation() bool {
	if o != nil && !IsNil(o.EngineUserValidation) {
		return true
	}

	return false
}

// SetEngineUserValidation gets a reference to the given bool and assigns it to the EngineUserValidation field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetEngineUserValidation(v bool) {
	o.EngineUserValidation = &v
}

// GetGlobalId returns the GlobalId field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetGlobalId() int64 {
	if o == nil || IsNil(o.GlobalId) {
		var ret int64
		return ret
	}
	return *o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetGlobalIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GlobalId) {
		return nil, false
	}
	return o.GlobalId, true
}

// HasGlobalId returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasGlobalId() bool {
	if o != nil && !IsNil(o.GlobalId) {
		return true
	}

	return false
}

// SetGlobalId gets a reference to the given int64 and assigns it to the GlobalId field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetGlobalId(v int64) {
	o.GlobalId = &v
}

// GetIsGlobalSync returns the IsGlobalSync field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetIsGlobalSync() bool {
	if o == nil || IsNil(o.IsGlobalSync) {
		var ret bool
		return ret
	}
	return *o.IsGlobalSync
}

// GetIsGlobalSyncOk returns a tuple with the IsGlobalSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetIsGlobalSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobalSync) {
		return nil, false
	}
	return o.IsGlobalSync, true
}

// HasIsGlobalSync returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasIsGlobalSync() bool {
	if o != nil && !IsNil(o.IsGlobalSync) {
		return true
	}

	return false
}

// SetIsGlobalSync gets a reference to the given bool and assigns it to the IsGlobalSync field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetIsGlobalSync(v bool) {
	o.IsGlobalSync = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetName(v string) {
	o.Name = &v
}

// GetNotifyFilterProfiles returns the NotifyFilterProfiles field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetNotifyFilterProfiles() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue {
	if o == nil || IsNil(o.NotifyFilterProfiles) {
		var ret map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue
		return ret
	}
	return *o.NotifyFilterProfiles
}

// GetNotifyFilterProfilesOk returns a tuple with the NotifyFilterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetNotifyFilterProfilesOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue, bool) {
	if o == nil || IsNil(o.NotifyFilterProfiles) {
		return nil, false
	}
	return o.NotifyFilterProfiles, true
}

// HasNotifyFilterProfiles returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasNotifyFilterProfiles() bool {
	if o != nil && !IsNil(o.NotifyFilterProfiles) {
		return true
	}

	return false
}

// SetNotifyFilterProfiles gets a reference to the given map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue and assigns it to the NotifyFilterProfiles field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetNotifyFilterProfiles(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigNotifyFilterProfilesValue) {
	o.NotifyFilterProfiles = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetTargets() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue {
	if o == nil || IsNil(o.Targets) {
		var ret map[string]V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetTargetsOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given map[string]V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue and assigns it to the Targets field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetTargets(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigTargetsValue) {
	o.Targets = &v
}

// GetUsmLocalUsers returns the UsmLocalUsers field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetUsmLocalUsers() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue {
	if o == nil || IsNil(o.UsmLocalUsers) {
		var ret map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue
		return ret
	}
	return *o.UsmLocalUsers
}

// GetUsmLocalUsersOk returns a tuple with the UsmLocalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetUsmLocalUsersOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue, bool) {
	if o == nil || IsNil(o.UsmLocalUsers) {
		return nil, false
	}
	return o.UsmLocalUsers, true
}

// HasUsmLocalUsers returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasUsmLocalUsers() bool {
	if o != nil && !IsNil(o.UsmLocalUsers) {
		return true
	}

	return false
}

// SetUsmLocalUsers gets a reference to the given map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue and assigns it to the UsmLocalUsers field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetUsmLocalUsers(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmLocalUsersValue) {
	o.UsmLocalUsers = &v
}

// GetUsmRemoteUsers returns the UsmRemoteUsers field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetUsmRemoteUsers() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue {
	if o == nil || IsNil(o.UsmRemoteUsers) {
		var ret map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue
		return ret
	}
	return *o.UsmRemoteUsers
}

// GetUsmRemoteUsersOk returns a tuple with the UsmRemoteUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetUsmRemoteUsersOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue, bool) {
	if o == nil || IsNil(o.UsmRemoteUsers) {
		return nil, false
	}
	return o.UsmRemoteUsers, true
}

// HasUsmRemoteUsers returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasUsmRemoteUsers() bool {
	if o != nil && !IsNil(o.UsmRemoteUsers) {
		return true
	}

	return false
}

// SetUsmRemoteUsers gets a reference to the given map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue and assigns it to the UsmRemoteUsers field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetUsmRemoteUsers(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigUsmRemoteUsersValue) {
	o.UsmRemoteUsers = &v
}

// GetV2cEnabled returns the V2cEnabled field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetV2cEnabled() bool {
	if o == nil || IsNil(o.V2cEnabled) {
		var ret bool
		return ret
	}
	return *o.V2cEnabled
}

// GetV2cEnabledOk returns a tuple with the V2cEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetV2cEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.V2cEnabled) {
		return nil, false
	}
	return o.V2cEnabled, true
}

// HasV2cEnabled returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasV2cEnabled() bool {
	if o != nil && !IsNil(o.V2cEnabled) {
		return true
	}

	return false
}

// SetV2cEnabled gets a reference to the given bool and assigns it to the V2cEnabled field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetV2cEnabled(v bool) {
	o.V2cEnabled = &v
}

// GetV3Enabled returns the V3Enabled field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetV3Enabled() bool {
	if o == nil || IsNil(o.V3Enabled) {
		var ret bool
		return ret
	}
	return *o.V3Enabled
}

// GetV3EnabledOk returns a tuple with the V3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetV3EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.V3Enabled) {
		return nil, false
	}
	return o.V3Enabled, true
}

// HasV3Enabled returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasV3Enabled() bool {
	if o != nil && !IsNil(o.V3Enabled) {
		return true
	}

	return false
}

// SetV3Enabled gets a reference to the given bool and assigns it to the V3Enabled field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetV3Enabled(v bool) {
	o.V3Enabled = &v
}

// GetVacmGroups returns the VacmGroups field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetVacmGroups() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue {
	if o == nil || IsNil(o.VacmGroups) {
		var ret map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue
		return ret
	}
	return *o.VacmGroups
}

// GetVacmGroupsOk returns a tuple with the VacmGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetVacmGroupsOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue, bool) {
	if o == nil || IsNil(o.VacmGroups) {
		return nil, false
	}
	return o.VacmGroups, true
}

// HasVacmGroups returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasVacmGroups() bool {
	if o != nil && !IsNil(o.VacmGroups) {
		return true
	}

	return false
}

// SetVacmGroups gets a reference to the given map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue and assigns it to the VacmGroups field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetVacmGroups(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmGroupsValue) {
	o.VacmGroups = &v
}

// GetVacmViews returns the VacmViews field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetVacmViews() map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue {
	if o == nil || IsNil(o.VacmViews) {
		var ret map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue
		return ret
	}
	return *o.VacmViews
}

// GetVacmViewsOk returns a tuple with the VacmViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) GetVacmViewsOk() (*map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue, bool) {
	if o == nil || IsNil(o.VacmViews) {
		return nil, false
	}
	return o.VacmViews, true
}

// HasVacmViews returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) HasVacmViews() bool {
	if o != nil && !IsNil(o.VacmViews) {
		return true
	}

	return false
}

// SetVacmViews gets a reference to the given map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue and assigns it to the VacmViews field.
func (o *V1GlobalConfigPatchRequestSnmpsValueConfig) SetVacmViews(v map[string]V1GlobalConfigPatchRequestSnmpsValueConfigVacmViewsValue) {
	o.VacmViews = &v
}

func (o V1GlobalConfigPatchRequestSnmpsValueConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestSnmpsValueConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Communities) {
		toSerialize["communities"] = o.Communities
	}
	if !IsNil(o.EngineAuthenTraps) {
		toSerialize["engineAuthenTraps"] = o.EngineAuthenTraps
	}
	if !IsNil(o.EngineEnabled) {
		toSerialize["engineEnabled"] = o.EngineEnabled
	}
	if !IsNil(o.EngineEndpoints) {
		toSerialize["engineEndpoints"] = o.EngineEndpoints
	}
	if !IsNil(o.EngineHighCpuTraps) {
		toSerialize["engineHighCpuTraps"] = o.EngineHighCpuTraps
	}
	if !IsNil(o.EngineHighMemoryTraps) {
		toSerialize["engineHighMemoryTraps"] = o.EngineHighMemoryTraps
	}
	if !IsNil(o.EngineIdAdminOctets) {
		toSerialize["engineIdAdminOctets"] = o.EngineIdAdminOctets
	}
	if !IsNil(o.EngineIdAdminText) {
		toSerialize["engineIdAdminText"] = o.EngineIdAdminText
	}
	if !IsNil(o.EngineIdIpv4) {
		toSerialize["engineIdIpv4"] = o.EngineIdIpv4
	}
	if !IsNil(o.EngineIdIpv6) {
		toSerialize["engineIdIpv6"] = o.EngineIdIpv6
	}
	if !IsNil(o.EngineIdMac) {
		toSerialize["engineIdMac"] = o.EngineIdMac
	}
	if !IsNil(o.EngineIdRaw) {
		toSerialize["engineIdRaw"] = o.EngineIdRaw
	}
	if !IsNil(o.EngineLocalAcessV4) {
		toSerialize["engineLocalAcessV4"] = o.EngineLocalAcessV4
	}
	if !IsNil(o.EngineLocalAcessV6) {
		toSerialize["engineLocalAcessV6"] = o.EngineLocalAcessV6
	}
	if !IsNil(o.EngineUserHints) {
		toSerialize["engineUserHints"] = o.EngineUserHints
	}
	if !IsNil(o.EngineUserValidation) {
		toSerialize["engineUserValidation"] = o.EngineUserValidation
	}
	if !IsNil(o.GlobalId) {
		toSerialize["globalId"] = o.GlobalId
	}
	if !IsNil(o.IsGlobalSync) {
		toSerialize["isGlobalSync"] = o.IsGlobalSync
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NotifyFilterProfiles) {
		toSerialize["notifyFilterProfiles"] = o.NotifyFilterProfiles
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	if !IsNil(o.UsmLocalUsers) {
		toSerialize["usmLocalUsers"] = o.UsmLocalUsers
	}
	if !IsNil(o.UsmRemoteUsers) {
		toSerialize["usmRemoteUsers"] = o.UsmRemoteUsers
	}
	if !IsNil(o.V2cEnabled) {
		toSerialize["v2cEnabled"] = o.V2cEnabled
	}
	if !IsNil(o.V3Enabled) {
		toSerialize["v3Enabled"] = o.V3Enabled
	}
	if !IsNil(o.VacmGroups) {
		toSerialize["vacmGroups"] = o.VacmGroups
	}
	if !IsNil(o.VacmViews) {
		toSerialize["vacmViews"] = o.VacmViews
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestSnmpsValueConfig struct {
	value *V1GlobalConfigPatchRequestSnmpsValueConfig
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfig) Get() *V1GlobalConfigPatchRequestSnmpsValueConfig {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfig) Set(val *V1GlobalConfigPatchRequestSnmpsValueConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestSnmpsValueConfig(val *V1GlobalConfigPatchRequestSnmpsValueConfig) *NullableV1GlobalConfigPatchRequestSnmpsValueConfig {
	return &NullableV1GlobalConfigPatchRequestSnmpsValueConfig{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestSnmpsValueConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestSnmpsValueConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


