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

// checks if the V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch{}

// V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch struct for V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch
type V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch struct {
	Application *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchApplication `json:"application,omitempty"`
	ContentFilter *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchContentFilter `json:"contentFilter,omitempty"`
	DestinationNetwork *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationNetwork `json:"destinationNetwork,omitempty"`
	DestinationPort *int32 `json:"destinationPort,omitempty"`
	DestinationPortRange *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange `json:"destinationPortRange,omitempty"`
	Dscp *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscp `json:"dscp,omitempty"`
	IcmpType *int32 `json:"icmpType,omitempty"`
	IpProtocol *string `json:"ipProtocol,omitempty"`
	Protocol *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchProtocol `json:"protocol,omitempty"`
	SourceNetwork *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchSourceNetwork `json:"sourceNetwork,omitempty"`
	SourcePort *int32 `json:"sourcePort,omitempty"`
	SourcePortRange *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange `json:"sourcePortRange,omitempty"`
}

// NewV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch() *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch {
	this := V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch{}
	return &this
}

// NewV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchWithDefaults instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchWithDefaults() *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch {
	this := V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetApplication() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchApplication {
	if o == nil || IsNil(o.Application) {
		var ret V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchApplication
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetApplicationOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchApplication, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchApplication and assigns it to the Application field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetApplication(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchApplication) {
	o.Application = &v
}

// GetContentFilter returns the ContentFilter field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetContentFilter() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchContentFilter {
	if o == nil || IsNil(o.ContentFilter) {
		var ret V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchContentFilter
		return ret
	}
	return *o.ContentFilter
}

// GetContentFilterOk returns a tuple with the ContentFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetContentFilterOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchContentFilter, bool) {
	if o == nil || IsNil(o.ContentFilter) {
		return nil, false
	}
	return o.ContentFilter, true
}

// HasContentFilter returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasContentFilter() bool {
	if o != nil && !IsNil(o.ContentFilter) {
		return true
	}

	return false
}

// SetContentFilter gets a reference to the given V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchContentFilter and assigns it to the ContentFilter field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetContentFilter(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchContentFilter) {
	o.ContentFilter = &v
}

// GetDestinationNetwork returns the DestinationNetwork field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetDestinationNetwork() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationNetwork {
	if o == nil || IsNil(o.DestinationNetwork) {
		var ret V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationNetwork
		return ret
	}
	return *o.DestinationNetwork
}

// GetDestinationNetworkOk returns a tuple with the DestinationNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetDestinationNetworkOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationNetwork, bool) {
	if o == nil || IsNil(o.DestinationNetwork) {
		return nil, false
	}
	return o.DestinationNetwork, true
}

// HasDestinationNetwork returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasDestinationNetwork() bool {
	if o != nil && !IsNil(o.DestinationNetwork) {
		return true
	}

	return false
}

// SetDestinationNetwork gets a reference to the given V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationNetwork and assigns it to the DestinationNetwork field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetDestinationNetwork(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationNetwork) {
	o.DestinationNetwork = &v
}

// GetDestinationPort returns the DestinationPort field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetDestinationPort() int32 {
	if o == nil || IsNil(o.DestinationPort) {
		var ret int32
		return ret
	}
	return *o.DestinationPort
}

// GetDestinationPortOk returns a tuple with the DestinationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetDestinationPortOk() (*int32, bool) {
	if o == nil || IsNil(o.DestinationPort) {
		return nil, false
	}
	return o.DestinationPort, true
}

// HasDestinationPort returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasDestinationPort() bool {
	if o != nil && !IsNil(o.DestinationPort) {
		return true
	}

	return false
}

// SetDestinationPort gets a reference to the given int32 and assigns it to the DestinationPort field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetDestinationPort(v int32) {
	o.DestinationPort = &v
}

// GetDestinationPortRange returns the DestinationPortRange field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetDestinationPortRange() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange {
	if o == nil || IsNil(o.DestinationPortRange) {
		var ret V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange
		return ret
	}
	return *o.DestinationPortRange
}

// GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetDestinationPortRangeOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange, bool) {
	if o == nil || IsNil(o.DestinationPortRange) {
		return nil, false
	}
	return o.DestinationPortRange, true
}

// HasDestinationPortRange returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasDestinationPortRange() bool {
	if o != nil && !IsNil(o.DestinationPortRange) {
		return true
	}

	return false
}

// SetDestinationPortRange gets a reference to the given V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange and assigns it to the DestinationPortRange field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetDestinationPortRange(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange) {
	o.DestinationPortRange = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetDscp() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscp {
	if o == nil || IsNil(o.Dscp) {
		var ret V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscp
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetDscpOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscp, bool) {
	if o == nil || IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasDscp() bool {
	if o != nil && !IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscp and assigns it to the Dscp field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetDscp(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDscp) {
	o.Dscp = &v
}

// GetIcmpType returns the IcmpType field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetIcmpType() int32 {
	if o == nil || IsNil(o.IcmpType) {
		var ret int32
		return ret
	}
	return *o.IcmpType
}

// GetIcmpTypeOk returns a tuple with the IcmpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetIcmpTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.IcmpType) {
		return nil, false
	}
	return o.IcmpType, true
}

// HasIcmpType returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasIcmpType() bool {
	if o != nil && !IsNil(o.IcmpType) {
		return true
	}

	return false
}

// SetIcmpType gets a reference to the given int32 and assigns it to the IcmpType field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetIcmpType(v int32) {
	o.IcmpType = &v
}

// GetIpProtocol returns the IpProtocol field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetIpProtocol() string {
	if o == nil || IsNil(o.IpProtocol) {
		var ret string
		return ret
	}
	return *o.IpProtocol
}

// GetIpProtocolOk returns a tuple with the IpProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetIpProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.IpProtocol) {
		return nil, false
	}
	return o.IpProtocol, true
}

// HasIpProtocol returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasIpProtocol() bool {
	if o != nil && !IsNil(o.IpProtocol) {
		return true
	}

	return false
}

// SetIpProtocol gets a reference to the given string and assigns it to the IpProtocol field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetIpProtocol(v string) {
	o.IpProtocol = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetProtocol() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetProtocolOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchProtocol and assigns it to the Protocol field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetProtocol(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchProtocol) {
	o.Protocol = &v
}

// GetSourceNetwork returns the SourceNetwork field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetSourceNetwork() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchSourceNetwork {
	if o == nil || IsNil(o.SourceNetwork) {
		var ret V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchSourceNetwork
		return ret
	}
	return *o.SourceNetwork
}

// GetSourceNetworkOk returns a tuple with the SourceNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetSourceNetworkOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchSourceNetwork, bool) {
	if o == nil || IsNil(o.SourceNetwork) {
		return nil, false
	}
	return o.SourceNetwork, true
}

// HasSourceNetwork returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasSourceNetwork() bool {
	if o != nil && !IsNil(o.SourceNetwork) {
		return true
	}

	return false
}

// SetSourceNetwork gets a reference to the given V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchSourceNetwork and assigns it to the SourceNetwork field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetSourceNetwork(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchSourceNetwork) {
	o.SourceNetwork = &v
}

// GetSourcePort returns the SourcePort field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetSourcePort() int32 {
	if o == nil || IsNil(o.SourcePort) {
		var ret int32
		return ret
	}
	return *o.SourcePort
}

// GetSourcePortOk returns a tuple with the SourcePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetSourcePortOk() (*int32, bool) {
	if o == nil || IsNil(o.SourcePort) {
		return nil, false
	}
	return o.SourcePort, true
}

// HasSourcePort returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasSourcePort() bool {
	if o != nil && !IsNil(o.SourcePort) {
		return true
	}

	return false
}

// SetSourcePort gets a reference to the given int32 and assigns it to the SourcePort field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetSourcePort(v int32) {
	o.SourcePort = &v
}

// GetSourcePortRange returns the SourcePortRange field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetSourcePortRange() V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange {
	if o == nil || IsNil(o.SourcePortRange) {
		var ret V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange
		return ret
	}
	return *o.SourcePortRange
}

// GetSourcePortRangeOk returns a tuple with the SourcePortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) GetSourcePortRangeOk() (*V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange, bool) {
	if o == nil || IsNil(o.SourcePortRange) {
		return nil, false
	}
	return o.SourcePortRange, true
}

// HasSourcePortRange returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) HasSourcePortRange() bool {
	if o != nil && !IsNil(o.SourcePortRange) {
		return true
	}

	return false
}

// SetSourcePortRange gets a reference to the given V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange and assigns it to the SourcePortRange field.
func (o *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) SetSourcePortRange(v V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatchDestinationPortRange) {
	o.SourcePortRange = &v
}

func (o V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.ContentFilter) {
		toSerialize["contentFilter"] = o.ContentFilter
	}
	if !IsNil(o.DestinationNetwork) {
		toSerialize["destinationNetwork"] = o.DestinationNetwork
	}
	if !IsNil(o.DestinationPort) {
		toSerialize["destinationPort"] = o.DestinationPort
	}
	if !IsNil(o.DestinationPortRange) {
		toSerialize["destinationPortRange"] = o.DestinationPortRange
	}
	if !IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	if !IsNil(o.IcmpType) {
		toSerialize["icmpType"] = o.IcmpType
	}
	if !IsNil(o.IpProtocol) {
		toSerialize["ipProtocol"] = o.IpProtocol
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.SourceNetwork) {
		toSerialize["sourceNetwork"] = o.SourceNetwork
	}
	if !IsNil(o.SourcePort) {
		toSerialize["sourcePort"] = o.SourcePort
	}
	if !IsNil(o.SourcePortRange) {
		toSerialize["sourcePortRange"] = o.SourcePortRange
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch struct {
	value *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) Get() *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) Set(val *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch(val *V1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) *NullableV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch {
	return &NullableV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestTrafficPoliciesSecurityRulesetsValueRulesetRulesValueRuleMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


