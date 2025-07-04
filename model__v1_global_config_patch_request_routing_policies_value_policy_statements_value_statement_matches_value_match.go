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

// checks if the V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch{}

// V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch struct for V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch
type V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch struct {
	Community *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity `json:"community,omitempty"`
	PrefixSet *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet `json:"prefixSet,omitempty"`
	ProtocolRouteType *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType `json:"protocolRouteType,omitempty"`
	RouteTag *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag `json:"routeTag,omitempty"`
	Seq *int32 `json:"seq,omitempty"`
	SourceInterface *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface `json:"sourceInterface,omitempty"`
	SourceProtocol *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol `json:"sourceProtocol,omitempty"`
	Stale *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale `json:"stale,omitempty"`
}

// NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch {
	this := V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch{}
	return &this
}

// NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchWithDefaults instantiates a new V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchWithDefaults() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch {
	this := V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch{}
	return &this
}

// GetCommunity returns the Community field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetCommunity() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity {
	if o == nil || IsNil(o.Community) {
		var ret V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity
		return ret
	}
	return *o.Community
}

// GetCommunityOk returns a tuple with the Community field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetCommunityOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity, bool) {
	if o == nil || IsNil(o.Community) {
		return nil, false
	}
	return o.Community, true
}

// HasCommunity returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasCommunity() bool {
	if o != nil && !IsNil(o.Community) {
		return true
	}

	return false
}

// SetCommunity gets a reference to the given V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity and assigns it to the Community field.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetCommunity(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchCommunity) {
	o.Community = &v
}

// GetPrefixSet returns the PrefixSet field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetPrefixSet() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet {
	if o == nil || IsNil(o.PrefixSet) {
		var ret V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet
		return ret
	}
	return *o.PrefixSet
}

// GetPrefixSetOk returns a tuple with the PrefixSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetPrefixSetOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet, bool) {
	if o == nil || IsNil(o.PrefixSet) {
		return nil, false
	}
	return o.PrefixSet, true
}

// HasPrefixSet returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasPrefixSet() bool {
	if o != nil && !IsNil(o.PrefixSet) {
		return true
	}

	return false
}

// SetPrefixSet gets a reference to the given V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet and assigns it to the PrefixSet field.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetPrefixSet(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchPrefixSet) {
	o.PrefixSet = &v
}

// GetProtocolRouteType returns the ProtocolRouteType field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetProtocolRouteType() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType {
	if o == nil || IsNil(o.ProtocolRouteType) {
		var ret V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType
		return ret
	}
	return *o.ProtocolRouteType
}

// GetProtocolRouteTypeOk returns a tuple with the ProtocolRouteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetProtocolRouteTypeOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType, bool) {
	if o == nil || IsNil(o.ProtocolRouteType) {
		return nil, false
	}
	return o.ProtocolRouteType, true
}

// HasProtocolRouteType returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasProtocolRouteType() bool {
	if o != nil && !IsNil(o.ProtocolRouteType) {
		return true
	}

	return false
}

// SetProtocolRouteType gets a reference to the given V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType and assigns it to the ProtocolRouteType field.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetProtocolRouteType(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchProtocolRouteType) {
	o.ProtocolRouteType = &v
}

// GetRouteTag returns the RouteTag field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetRouteTag() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag {
	if o == nil || IsNil(o.RouteTag) {
		var ret V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag
		return ret
	}
	return *o.RouteTag
}

// GetRouteTagOk returns a tuple with the RouteTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetRouteTagOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag, bool) {
	if o == nil || IsNil(o.RouteTag) {
		return nil, false
	}
	return o.RouteTag, true
}

// HasRouteTag returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasRouteTag() bool {
	if o != nil && !IsNil(o.RouteTag) {
		return true
	}

	return false
}

// SetRouteTag gets a reference to the given V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag and assigns it to the RouteTag field.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetRouteTag(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchRouteTag) {
	o.RouteTag = &v
}

// GetSeq returns the Seq field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSeq() int32 {
	if o == nil || IsNil(o.Seq) {
		var ret int32
		return ret
	}
	return *o.Seq
}

// GetSeqOk returns a tuple with the Seq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSeqOk() (*int32, bool) {
	if o == nil || IsNil(o.Seq) {
		return nil, false
	}
	return o.Seq, true
}

// HasSeq returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasSeq() bool {
	if o != nil && !IsNil(o.Seq) {
		return true
	}

	return false
}

// SetSeq gets a reference to the given int32 and assigns it to the Seq field.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetSeq(v int32) {
	o.Seq = &v
}

// GetSourceInterface returns the SourceInterface field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSourceInterface() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface {
	if o == nil || IsNil(o.SourceInterface) {
		var ret V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface
		return ret
	}
	return *o.SourceInterface
}

// GetSourceInterfaceOk returns a tuple with the SourceInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSourceInterfaceOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface, bool) {
	if o == nil || IsNil(o.SourceInterface) {
		return nil, false
	}
	return o.SourceInterface, true
}

// HasSourceInterface returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasSourceInterface() bool {
	if o != nil && !IsNil(o.SourceInterface) {
		return true
	}

	return false
}

// SetSourceInterface gets a reference to the given V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface and assigns it to the SourceInterface field.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetSourceInterface(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceInterface) {
	o.SourceInterface = &v
}

// GetSourceProtocol returns the SourceProtocol field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSourceProtocol() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol {
	if o == nil || IsNil(o.SourceProtocol) {
		var ret V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol
		return ret
	}
	return *o.SourceProtocol
}

// GetSourceProtocolOk returns a tuple with the SourceProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetSourceProtocolOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol, bool) {
	if o == nil || IsNil(o.SourceProtocol) {
		return nil, false
	}
	return o.SourceProtocol, true
}

// HasSourceProtocol returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasSourceProtocol() bool {
	if o != nil && !IsNil(o.SourceProtocol) {
		return true
	}

	return false
}

// SetSourceProtocol gets a reference to the given V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol and assigns it to the SourceProtocol field.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetSourceProtocol(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchSourceProtocol) {
	o.SourceProtocol = &v
}

// GetStale returns the Stale field value if set, zero value otherwise.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetStale() V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale {
	if o == nil || IsNil(o.Stale) {
		var ret V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale
		return ret
	}
	return *o.Stale
}

// GetStaleOk returns a tuple with the Stale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) GetStaleOk() (*V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale, bool) {
	if o == nil || IsNil(o.Stale) {
		return nil, false
	}
	return o.Stale, true
}

// HasStale returns a boolean if a field has been set.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) HasStale() bool {
	if o != nil && !IsNil(o.Stale) {
		return true
	}

	return false
}

// SetStale gets a reference to the given V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale and assigns it to the Stale field.
func (o *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) SetStale(v V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatchStale) {
	o.Stale = &v
}

func (o V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Community) {
		toSerialize["community"] = o.Community
	}
	if !IsNil(o.PrefixSet) {
		toSerialize["prefixSet"] = o.PrefixSet
	}
	if !IsNil(o.ProtocolRouteType) {
		toSerialize["protocolRouteType"] = o.ProtocolRouteType
	}
	if !IsNil(o.RouteTag) {
		toSerialize["routeTag"] = o.RouteTag
	}
	if !IsNil(o.Seq) {
		toSerialize["seq"] = o.Seq
	}
	if !IsNil(o.SourceInterface) {
		toSerialize["sourceInterface"] = o.SourceInterface
	}
	if !IsNil(o.SourceProtocol) {
		toSerialize["sourceProtocol"] = o.SourceProtocol
	}
	if !IsNil(o.Stale) {
		toSerialize["stale"] = o.Stale
	}
	return toSerialize, nil
}

type NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch struct {
	value *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch
	isSet bool
}

func (v NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) Get() *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch {
	return v.value
}

func (v *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) Set(val *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch(val *V1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch {
	return &NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch{value: val, isSet: true}
}

func (v NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalConfigPatchRequestRoutingPoliciesValuePolicyStatementsValueStatementMatchesValueMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


