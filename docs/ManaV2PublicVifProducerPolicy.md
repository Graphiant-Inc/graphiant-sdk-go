# ManaV2PublicVifProducerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**map[string]ManaV2PublicVifDevice**](ManaV2PublicVifDevice.md) |  | 
**Profiles** | Pointer to [**[]ManaV2ApplicationProfile**](ManaV2ApplicationProfile.md) |  | [optional] 
**ServiceLanSegment** | **int64** | LAN segment ID for the service (required) | 
**ServiceName** | Pointer to **string** | Public VIF service name (local_extranet_producer_service.name) | [optional] 
**Sites** | [**ManaV2SiteInformation**](ManaV2SiteInformation.md) |  | 
**Sla** | Pointer to [**ManaV2SlaInformation**](ManaV2SlaInformation.md) |  | [optional] 
**Vifs** | [**ManaV2PublicVif**](ManaV2PublicVif.md) |  | 

## Methods

### NewManaV2PublicVifProducerPolicy

`func NewManaV2PublicVifProducerPolicy(devices map[string]ManaV2PublicVifDevice, serviceLanSegment int64, sites ManaV2SiteInformation, vifs ManaV2PublicVif, ) *ManaV2PublicVifProducerPolicy`

NewManaV2PublicVifProducerPolicy instantiates a new ManaV2PublicVifProducerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManaV2PublicVifProducerPolicyWithDefaults

`func NewManaV2PublicVifProducerPolicyWithDefaults() *ManaV2PublicVifProducerPolicy`

NewManaV2PublicVifProducerPolicyWithDefaults instantiates a new ManaV2PublicVifProducerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *ManaV2PublicVifProducerPolicy) GetDevices() map[string]ManaV2PublicVifDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ManaV2PublicVifProducerPolicy) GetDevicesOk() (*map[string]ManaV2PublicVifDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ManaV2PublicVifProducerPolicy) SetDevices(v map[string]ManaV2PublicVifDevice)`

SetDevices sets Devices field to given value.


### GetProfiles

`func (o *ManaV2PublicVifProducerPolicy) GetProfiles() []ManaV2ApplicationProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *ManaV2PublicVifProducerPolicy) GetProfilesOk() (*[]ManaV2ApplicationProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *ManaV2PublicVifProducerPolicy) SetProfiles(v []ManaV2ApplicationProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *ManaV2PublicVifProducerPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetServiceLanSegment

`func (o *ManaV2PublicVifProducerPolicy) GetServiceLanSegment() int64`

GetServiceLanSegment returns the ServiceLanSegment field if non-nil, zero value otherwise.

### GetServiceLanSegmentOk

`func (o *ManaV2PublicVifProducerPolicy) GetServiceLanSegmentOk() (*int64, bool)`

GetServiceLanSegmentOk returns a tuple with the ServiceLanSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLanSegment

`func (o *ManaV2PublicVifProducerPolicy) SetServiceLanSegment(v int64)`

SetServiceLanSegment sets ServiceLanSegment field to given value.


### GetServiceName

`func (o *ManaV2PublicVifProducerPolicy) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ManaV2PublicVifProducerPolicy) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ManaV2PublicVifProducerPolicy) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ManaV2PublicVifProducerPolicy) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSites

`func (o *ManaV2PublicVifProducerPolicy) GetSites() ManaV2SiteInformation`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ManaV2PublicVifProducerPolicy) GetSitesOk() (*ManaV2SiteInformation, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ManaV2PublicVifProducerPolicy) SetSites(v ManaV2SiteInformation)`

SetSites sets Sites field to given value.


### GetSla

`func (o *ManaV2PublicVifProducerPolicy) GetSla() ManaV2SlaInformation`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *ManaV2PublicVifProducerPolicy) GetSlaOk() (*ManaV2SlaInformation, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *ManaV2PublicVifProducerPolicy) SetSla(v ManaV2SlaInformation)`

SetSla sets Sla field to given value.

### HasSla

`func (o *ManaV2PublicVifProducerPolicy) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetVifs

`func (o *ManaV2PublicVifProducerPolicy) GetVifs() ManaV2PublicVif`

GetVifs returns the Vifs field if non-nil, zero value otherwise.

### GetVifsOk

`func (o *ManaV2PublicVifProducerPolicy) GetVifsOk() (*ManaV2PublicVif, bool)`

GetVifsOk returns a tuple with the Vifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifs

`func (o *ManaV2PublicVifProducerPolicy) SetVifs(v ManaV2PublicVif)`

SetVifs sets Vifs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


