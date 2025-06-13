# V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DestinationNetwork** | Pointer to **string** |  | [optional] 
**DestinationNetworkList** | Pointer to **string** |  | [optional] 
**DestinationNetworks** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks**](V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks.md) |  | [optional] 
**DestinationPort** | Pointer to **int32** |  | [optional] 
**DestinationPortList** | Pointer to **string** |  | [optional] 
**DestinationPorts** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationPorts**](V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationPorts.md) |  | [optional] 
**IcmpType** | Pointer to **int32** |  | [optional] 
**IpProtocol** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SourceNetwork** | Pointer to **string** |  | [optional] 
**SourceNetworkList** | Pointer to **string** |  | [optional] 
**SourceNetworks** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks**](V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks.md) |  | [optional] 
**SourcePort** | Pointer to **int32** |  | [optional] 
**SourcePortList** | Pointer to **string** |  | [optional] 
**SourcePorts** | Pointer to [**V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationPorts**](V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationPorts.md) |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication

`func NewV1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication() *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication`

NewV1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationWithDefaults

`func NewV1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationWithDefaults() *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication`

NewV1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationWithDefaults instantiates a new V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinationNetwork

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationNetwork() string`

GetDestinationNetwork returns the DestinationNetwork field if non-nil, zero value otherwise.

### GetDestinationNetworkOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationNetworkOk() (*string, bool)`

GetDestinationNetworkOk returns a tuple with the DestinationNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNetwork

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetDestinationNetwork(v string)`

SetDestinationNetwork sets DestinationNetwork field to given value.

### HasDestinationNetwork

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasDestinationNetwork() bool`

HasDestinationNetwork returns a boolean if a field has been set.

### GetDestinationNetworkList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationNetworkList() string`

GetDestinationNetworkList returns the DestinationNetworkList field if non-nil, zero value otherwise.

### GetDestinationNetworkListOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationNetworkListOk() (*string, bool)`

GetDestinationNetworkListOk returns a tuple with the DestinationNetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNetworkList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetDestinationNetworkList(v string)`

SetDestinationNetworkList sets DestinationNetworkList field to given value.

### HasDestinationNetworkList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasDestinationNetworkList() bool`

HasDestinationNetworkList returns a boolean if a field has been set.

### GetDestinationNetworks

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationNetworks() V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks`

GetDestinationNetworks returns the DestinationNetworks field if non-nil, zero value otherwise.

### GetDestinationNetworksOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationNetworksOk() (*V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks, bool)`

GetDestinationNetworksOk returns a tuple with the DestinationNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNetworks

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetDestinationNetworks(v V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks)`

SetDestinationNetworks sets DestinationNetworks field to given value.

### HasDestinationNetworks

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasDestinationNetworks() bool`

HasDestinationNetworks returns a boolean if a field has been set.

### GetDestinationPort

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationPort() int32`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationPortOk() (*int32, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetDestinationPort(v int32)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetDestinationPortList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationPortList() string`

GetDestinationPortList returns the DestinationPortList field if non-nil, zero value otherwise.

### GetDestinationPortListOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationPortListOk() (*string, bool)`

GetDestinationPortListOk returns a tuple with the DestinationPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetDestinationPortList(v string)`

SetDestinationPortList sets DestinationPortList field to given value.

### HasDestinationPortList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasDestinationPortList() bool`

HasDestinationPortList returns a boolean if a field has been set.

### GetDestinationPorts

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationPorts() V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationPorts`

GetDestinationPorts returns the DestinationPorts field if non-nil, zero value otherwise.

### GetDestinationPortsOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetDestinationPortsOk() (*V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationPorts, bool)`

GetDestinationPortsOk returns a tuple with the DestinationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPorts

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetDestinationPorts(v V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationPorts)`

SetDestinationPorts sets DestinationPorts field to given value.

### HasDestinationPorts

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasDestinationPorts() bool`

HasDestinationPorts returns a boolean if a field has been set.

### GetIcmpType

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.

### HasIcmpType

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasIcmpType() bool`

HasIcmpType returns a boolean if a field has been set.

### GetIpProtocol

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetIpProtocol() string`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetIpProtocolOk() (*string, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpProtocol

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetIpProtocol(v string)`

SetIpProtocol sets IpProtocol field to given value.

### HasIpProtocol

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasIpProtocol() bool`

HasIpProtocol returns a boolean if a field has been set.

### GetName

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceNetwork

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourceNetwork() string`

GetSourceNetwork returns the SourceNetwork field if non-nil, zero value otherwise.

### GetSourceNetworkOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourceNetworkOk() (*string, bool)`

GetSourceNetworkOk returns a tuple with the SourceNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetwork

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetSourceNetwork(v string)`

SetSourceNetwork sets SourceNetwork field to given value.

### HasSourceNetwork

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasSourceNetwork() bool`

HasSourceNetwork returns a boolean if a field has been set.

### GetSourceNetworkList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourceNetworkList() string`

GetSourceNetworkList returns the SourceNetworkList field if non-nil, zero value otherwise.

### GetSourceNetworkListOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourceNetworkListOk() (*string, bool)`

GetSourceNetworkListOk returns a tuple with the SourceNetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetworkList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetSourceNetworkList(v string)`

SetSourceNetworkList sets SourceNetworkList field to given value.

### HasSourceNetworkList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasSourceNetworkList() bool`

HasSourceNetworkList returns a boolean if a field has been set.

### GetSourceNetworks

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourceNetworks() V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks`

GetSourceNetworks returns the SourceNetworks field if non-nil, zero value otherwise.

### GetSourceNetworksOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourceNetworksOk() (*V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks, bool)`

GetSourceNetworksOk returns a tuple with the SourceNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetworks

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetSourceNetworks(v V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationNetworks)`

SetSourceNetworks sets SourceNetworks field to given value.

### HasSourceNetworks

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasSourceNetworks() bool`

HasSourceNetworks returns a boolean if a field has been set.

### GetSourcePort

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourcePort() int32`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourcePortOk() (*int32, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetSourcePort(v int32)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetSourcePortList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourcePortList() string`

GetSourcePortList returns the SourcePortList field if non-nil, zero value otherwise.

### GetSourcePortListOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourcePortListOk() (*string, bool)`

GetSourcePortListOk returns a tuple with the SourcePortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetSourcePortList(v string)`

SetSourcePortList sets SourcePortList field to given value.

### HasSourcePortList

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasSourcePortList() bool`

HasSourcePortList returns a boolean if a field has been set.

### GetSourcePorts

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourcePorts() V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationPorts`

GetSourcePorts returns the SourcePorts field if non-nil, zero value otherwise.

### GetSourcePortsOk

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) GetSourcePortsOk() (*V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationPorts, bool)`

GetSourcePortsOk returns a tuple with the SourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePorts

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) SetSourcePorts(v V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplicationDestinationPorts)`

SetSourcePorts sets SourcePorts field to given value.

### HasSourcePorts

`func (o *V1GlobalConfigPatchRequestTrafficPoliciesDpiApplicationsValueApplication) HasSourcePorts() bool`

HasSourcePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


