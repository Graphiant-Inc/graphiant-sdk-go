# ManaV2PublicVifProducerPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**map[string]ManaV2PublicVifDevice**](ManaV2PublicVifDevice.md) |  | [optional] 
**InboundSecurityRules** | Pointer to [**[]ManaV2SecurityPolicyRule**](ManaV2SecurityPolicyRule.md) |  | [optional] 
**Profiles** | Pointer to [**[]ManaV2ApplicationProfile**](ManaV2ApplicationProfile.md) |  | [optional] 
**ServiceLanSegment** | Pointer to **int64** | LAN segment ID for the service | [optional] 
**ServiceName** | Pointer to **string** | Public VIF service name (local_extranet_producer_service.name) | [optional] 
**Sites** | Pointer to [**ManaV2SiteInformation**](ManaV2SiteInformation.md) |  | [optional] 
**Sla** | Pointer to [**ManaV2SlaInformation**](ManaV2SlaInformation.md) |  | [optional] 
**TrafficRules** | Pointer to [**[]ManaV2TrafficPolicyRule**](ManaV2TrafficPolicyRule.md) |  | [optional] 
**Vifs** | Pointer to [**ManaV2PublicVif**](ManaV2PublicVif.md) |  | [optional] 

## Methods

### NewManaV2PublicVifProducerPolicyResponse

`func NewManaV2PublicVifProducerPolicyResponse() *ManaV2PublicVifProducerPolicyResponse`

NewManaV2PublicVifProducerPolicyResponse instantiates a new ManaV2PublicVifProducerPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PublicVifProducerPolicyResponseWithDefaults

`func NewManaV2PublicVifProducerPolicyResponseWithDefaults() *ManaV2PublicVifProducerPolicyResponse`

NewManaV2PublicVifProducerPolicyResponseWithDefaults instantiates a new ManaV2PublicVifProducerPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *ManaV2PublicVifProducerPolicyResponse) GetDevices() map[string]ManaV2PublicVifDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ManaV2PublicVifProducerPolicyResponse) GetDevicesOk() (*map[string]ManaV2PublicVifDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ManaV2PublicVifProducerPolicyResponse) SetDevices(v map[string]ManaV2PublicVifDevice)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *ManaV2PublicVifProducerPolicyResponse) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetInboundSecurityRules

`func (o *ManaV2PublicVifProducerPolicyResponse) GetInboundSecurityRules() []ManaV2SecurityPolicyRule`

GetInboundSecurityRules returns the InboundSecurityRules field if non-nil, zero value otherwise.

### GetInboundSecurityRulesOk

`func (o *ManaV2PublicVifProducerPolicyResponse) GetInboundSecurityRulesOk() (*[]ManaV2SecurityPolicyRule, bool)`

GetInboundSecurityRulesOk returns a tuple with the InboundSecurityRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSecurityRules

`func (o *ManaV2PublicVifProducerPolicyResponse) SetInboundSecurityRules(v []ManaV2SecurityPolicyRule)`

SetInboundSecurityRules sets InboundSecurityRules field to given value.

### HasInboundSecurityRules

`func (o *ManaV2PublicVifProducerPolicyResponse) HasInboundSecurityRules() bool`

HasInboundSecurityRules returns a boolean if a field has been set.

### GetProfiles

`func (o *ManaV2PublicVifProducerPolicyResponse) GetProfiles() []ManaV2ApplicationProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *ManaV2PublicVifProducerPolicyResponse) GetProfilesOk() (*[]ManaV2ApplicationProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *ManaV2PublicVifProducerPolicyResponse) SetProfiles(v []ManaV2ApplicationProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *ManaV2PublicVifProducerPolicyResponse) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetServiceLanSegment

`func (o *ManaV2PublicVifProducerPolicyResponse) GetServiceLanSegment() int64`

GetServiceLanSegment returns the ServiceLanSegment field if non-nil, zero value otherwise.

### GetServiceLanSegmentOk

`func (o *ManaV2PublicVifProducerPolicyResponse) GetServiceLanSegmentOk() (*int64, bool)`

GetServiceLanSegmentOk returns a tuple with the ServiceLanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLanSegment

`func (o *ManaV2PublicVifProducerPolicyResponse) SetServiceLanSegment(v int64)`

SetServiceLanSegment sets ServiceLanSegment field to given value.

### HasServiceLanSegment

`func (o *ManaV2PublicVifProducerPolicyResponse) HasServiceLanSegment() bool`

HasServiceLanSegment returns a boolean if a field has been set.

### GetServiceName

`func (o *ManaV2PublicVifProducerPolicyResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ManaV2PublicVifProducerPolicyResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ManaV2PublicVifProducerPolicyResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ManaV2PublicVifProducerPolicyResponse) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSites

`func (o *ManaV2PublicVifProducerPolicyResponse) GetSites() ManaV2SiteInformation`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ManaV2PublicVifProducerPolicyResponse) GetSitesOk() (*ManaV2SiteInformation, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ManaV2PublicVifProducerPolicyResponse) SetSites(v ManaV2SiteInformation)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ManaV2PublicVifProducerPolicyResponse) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetSla

`func (o *ManaV2PublicVifProducerPolicyResponse) GetSla() ManaV2SlaInformation`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *ManaV2PublicVifProducerPolicyResponse) GetSlaOk() (*ManaV2SlaInformation, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *ManaV2PublicVifProducerPolicyResponse) SetSla(v ManaV2SlaInformation)`

SetSla sets Sla field to given value.

### HasSla

`func (o *ManaV2PublicVifProducerPolicyResponse) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetTrafficRules

`func (o *ManaV2PublicVifProducerPolicyResponse) GetTrafficRules() []ManaV2TrafficPolicyRule`

GetTrafficRules returns the TrafficRules field if non-nil, zero value otherwise.

### GetTrafficRulesOk

`func (o *ManaV2PublicVifProducerPolicyResponse) GetTrafficRulesOk() (*[]ManaV2TrafficPolicyRule, bool)`

GetTrafficRulesOk returns a tuple with the TrafficRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRules

`func (o *ManaV2PublicVifProducerPolicyResponse) SetTrafficRules(v []ManaV2TrafficPolicyRule)`

SetTrafficRules sets TrafficRules field to given value.

### HasTrafficRules

`func (o *ManaV2PublicVifProducerPolicyResponse) HasTrafficRules() bool`

HasTrafficRules returns a boolean if a field has been set.

### GetVifs

`func (o *ManaV2PublicVifProducerPolicyResponse) GetVifs() ManaV2PublicVif`

GetVifs returns the Vifs field if non-nil, zero value otherwise.

### GetVifsOk

`func (o *ManaV2PublicVifProducerPolicyResponse) GetVifsOk() (*ManaV2PublicVif, bool)`

GetVifsOk returns a tuple with the Vifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifs

`func (o *ManaV2PublicVifProducerPolicyResponse) SetVifs(v ManaV2PublicVif)`

SetVifs sets Vifs field to given value.

### HasVifs

`func (o *ManaV2PublicVifProducerPolicyResponse) HasVifs() bool`

HasVifs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


