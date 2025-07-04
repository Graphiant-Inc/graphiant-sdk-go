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

// checks if the V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup{}

// V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup struct for V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup
type V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup struct {
	AcceptMode *bool `json:"acceptMode,omitempty"`
	AllowInterOperability *bool `json:"allowInterOperability,omitempty"`
	Description *string `json:"description,omitempty"`
	EffectivePriority *int32 `json:"effectivePriority,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	GroupMembers []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupGroupMembersInner `json:"groupMembers,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Preempt *bool `json:"preempt,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	State *string `json:"state,omitempty"`
	TrackedInterfaces []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner `json:"trackedInterfaces,omitempty"`
	VirtualIpAddress *string `json:"virtualIpAddress,omitempty"`
	VirtualMacAddress *string `json:"virtualMacAddress,omitempty"`
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup{}
	return &this
}

// NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupWithDefaults instantiates a new V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupWithDefaults() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup {
	this := V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup{}
	return &this
}

// GetAcceptMode returns the AcceptMode field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetAcceptMode() bool {
	if o == nil || IsNil(o.AcceptMode) {
		var ret bool
		return ret
	}
	return *o.AcceptMode
}

// GetAcceptModeOk returns a tuple with the AcceptMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetAcceptModeOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptMode) {
		return nil, false
	}
	return o.AcceptMode, true
}

// HasAcceptMode returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasAcceptMode() bool {
	if o != nil && !IsNil(o.AcceptMode) {
		return true
	}

	return false
}

// SetAcceptMode gets a reference to the given bool and assigns it to the AcceptMode field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetAcceptMode(v bool) {
	o.AcceptMode = &v
}

// GetAllowInterOperability returns the AllowInterOperability field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetAllowInterOperability() bool {
	if o == nil || IsNil(o.AllowInterOperability) {
		var ret bool
		return ret
	}
	return *o.AllowInterOperability
}

// GetAllowInterOperabilityOk returns a tuple with the AllowInterOperability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetAllowInterOperabilityOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowInterOperability) {
		return nil, false
	}
	return o.AllowInterOperability, true
}

// HasAllowInterOperability returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasAllowInterOperability() bool {
	if o != nil && !IsNil(o.AllowInterOperability) {
		return true
	}

	return false
}

// SetAllowInterOperability gets a reference to the given bool and assigns it to the AllowInterOperability field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetAllowInterOperability(v bool) {
	o.AllowInterOperability = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetDescription(v string) {
	o.Description = &v
}

// GetEffectivePriority returns the EffectivePriority field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetEffectivePriority() int32 {
	if o == nil || IsNil(o.EffectivePriority) {
		var ret int32
		return ret
	}
	return *o.EffectivePriority
}

// GetEffectivePriorityOk returns a tuple with the EffectivePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetEffectivePriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.EffectivePriority) {
		return nil, false
	}
	return o.EffectivePriority, true
}

// HasEffectivePriority returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasEffectivePriority() bool {
	if o != nil && !IsNil(o.EffectivePriority) {
		return true
	}

	return false
}

// SetEffectivePriority gets a reference to the given int32 and assigns it to the EffectivePriority field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetEffectivePriority(v int32) {
	o.EffectivePriority = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetGroupMembers returns the GroupMembers field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetGroupMembers() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupGroupMembersInner {
	if o == nil || IsNil(o.GroupMembers) {
		var ret []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupGroupMembersInner
		return ret
	}
	return o.GroupMembers
}

// GetGroupMembersOk returns a tuple with the GroupMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetGroupMembersOk() ([]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupGroupMembersInner, bool) {
	if o == nil || IsNil(o.GroupMembers) {
		return nil, false
	}
	return o.GroupMembers, true
}

// HasGroupMembers returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasGroupMembers() bool {
	if o != nil && !IsNil(o.GroupMembers) {
		return true
	}

	return false
}

// SetGroupMembers gets a reference to the given []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupGroupMembersInner and assigns it to the GroupMembers field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetGroupMembers(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupGroupMembersInner) {
	o.GroupMembers = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetName(v string) {
	o.Name = &v
}

// GetPreempt returns the Preempt field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetPreempt() bool {
	if o == nil || IsNil(o.Preempt) {
		var ret bool
		return ret
	}
	return *o.Preempt
}

// GetPreemptOk returns a tuple with the Preempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetPreemptOk() (*bool, bool) {
	if o == nil || IsNil(o.Preempt) {
		return nil, false
	}
	return o.Preempt, true
}

// HasPreempt returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasPreempt() bool {
	if o != nil && !IsNil(o.Preempt) {
		return true
	}

	return false
}

// SetPreempt gets a reference to the given bool and assigns it to the Preempt field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetPreempt(v bool) {
	o.Preempt = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetPriority(v int32) {
	o.Priority = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetState(v string) {
	o.State = &v
}

// GetTrackedInterfaces returns the TrackedInterfaces field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetTrackedInterfaces() []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner {
	if o == nil || IsNil(o.TrackedInterfaces) {
		var ret []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner
		return ret
	}
	return o.TrackedInterfaces
}

// GetTrackedInterfacesOk returns a tuple with the TrackedInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetTrackedInterfacesOk() ([]V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner, bool) {
	if o == nil || IsNil(o.TrackedInterfaces) {
		return nil, false
	}
	return o.TrackedInterfaces, true
}

// HasTrackedInterfaces returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasTrackedInterfaces() bool {
	if o != nil && !IsNil(o.TrackedInterfaces) {
		return true
	}

	return false
}

// SetTrackedInterfaces gets a reference to the given []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner and assigns it to the TrackedInterfaces field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetTrackedInterfaces(v []V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroupTrackedInterfacesInner) {
	o.TrackedInterfaces = v
}

// GetVirtualIpAddress returns the VirtualIpAddress field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetVirtualIpAddress() string {
	if o == nil || IsNil(o.VirtualIpAddress) {
		var ret string
		return ret
	}
	return *o.VirtualIpAddress
}

// GetVirtualIpAddressOk returns a tuple with the VirtualIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetVirtualIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualIpAddress) {
		return nil, false
	}
	return o.VirtualIpAddress, true
}

// HasVirtualIpAddress returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasVirtualIpAddress() bool {
	if o != nil && !IsNil(o.VirtualIpAddress) {
		return true
	}

	return false
}

// SetVirtualIpAddress gets a reference to the given string and assigns it to the VirtualIpAddress field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetVirtualIpAddress(v string) {
	o.VirtualIpAddress = &v
}

// GetVirtualMacAddress returns the VirtualMacAddress field value if set, zero value otherwise.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetVirtualMacAddress() string {
	if o == nil || IsNil(o.VirtualMacAddress) {
		var ret string
		return ret
	}
	return *o.VirtualMacAddress
}

// GetVirtualMacAddressOk returns a tuple with the VirtualMacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) GetVirtualMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualMacAddress) {
		return nil, false
	}
	return o.VirtualMacAddress, true
}

// HasVirtualMacAddress returns a boolean if a field has been set.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) HasVirtualMacAddress() bool {
	if o != nil && !IsNil(o.VirtualMacAddress) {
		return true
	}

	return false
}

// SetVirtualMacAddress gets a reference to the given string and assigns it to the VirtualMacAddress field.
func (o *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) SetVirtualMacAddress(v string) {
	o.VirtualMacAddress = &v
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptMode) {
		toSerialize["acceptMode"] = o.AcceptMode
	}
	if !IsNil(o.AllowInterOperability) {
		toSerialize["allowInterOperability"] = o.AllowInterOperability
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EffectivePriority) {
		toSerialize["effectivePriority"] = o.EffectivePriority
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.GroupMembers) {
		toSerialize["groupMembers"] = o.GroupMembers
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Preempt) {
		toSerialize["preempt"] = o.Preempt
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TrackedInterfaces) {
		toSerialize["trackedInterfaces"] = o.TrackedInterfaces
	}
	if !IsNil(o.VirtualIpAddress) {
		toSerialize["virtualIpAddress"] = o.VirtualIpAddress
	}
	if !IsNil(o.VirtualMacAddress) {
		toSerialize["virtualMacAddress"] = o.VirtualMacAddress
	}
	return toSerialize, nil
}

type NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup struct {
	value *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup
	isSet bool
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) Get() *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup {
	return v.value
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) Set(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup(val *V1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup {
	return &NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup{value: val, isSet: true}
}

func (v NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ExtranetsGet200ResponsePoliciesInnerBranchesExcludedDevicesInnerInterfacesInnerIpv4VrrpGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


