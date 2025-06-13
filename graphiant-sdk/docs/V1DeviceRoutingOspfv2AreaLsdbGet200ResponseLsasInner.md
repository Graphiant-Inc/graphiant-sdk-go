# V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertisingRouter** | Pointer to **string** |  | [optional] 
**Age** | Pointer to **int32** |  | [optional] 
**AsexternalLsa** | Pointer to [**V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerAsexternalLsa**](V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerAsexternalLsa.md) |  | [optional] 
**Checksum** | Pointer to **int32** |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**LinkId** | Pointer to **string** |  | [optional] 
**NetworkLsa** | Pointer to [**V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa**](V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa.md) |  | [optional] 
**RouterLsa** | Pointer to [**V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerRouterLsa**](V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerRouterLsa.md) |  | [optional] 
**SequenceNumber** | Pointer to **string** |  | [optional] 
**SummaryLsa** | Pointer to [**V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerSummaryLsa**](V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerSummaryLsa.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner

`func NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner() *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner`

NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner instantiates a new V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerWithDefaults

`func NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerWithDefaults() *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner`

NewV1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerWithDefaults instantiates a new V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertisingRouter

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetAdvertisingRouter() string`

GetAdvertisingRouter returns the AdvertisingRouter field if non-nil, zero value otherwise.

### GetAdvertisingRouterOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetAdvertisingRouterOk() (*string, bool)`

GetAdvertisingRouterOk returns a tuple with the AdvertisingRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisingRouter

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetAdvertisingRouter(v string)`

SetAdvertisingRouter sets AdvertisingRouter field to given value.

### HasAdvertisingRouter

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasAdvertisingRouter() bool`

HasAdvertisingRouter returns a boolean if a field has been set.

### GetAge

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAsexternalLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetAsexternalLsa() V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerAsexternalLsa`

GetAsexternalLsa returns the AsexternalLsa field if non-nil, zero value otherwise.

### GetAsexternalLsaOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetAsexternalLsaOk() (*V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerAsexternalLsa, bool)`

GetAsexternalLsaOk returns a tuple with the AsexternalLsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsexternalLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetAsexternalLsa(v V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerAsexternalLsa)`

SetAsexternalLsa sets AsexternalLsa field to given value.

### HasAsexternalLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasAsexternalLsa() bool`

HasAsexternalLsa returns a boolean if a field has been set.

### GetChecksum

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetChecksum() int32`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetChecksumOk() (*int32, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetChecksum(v int32)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetLength

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetLinkId

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetNetworkLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetNetworkLsa() V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa`

GetNetworkLsa returns the NetworkLsa field if non-nil, zero value otherwise.

### GetNetworkLsaOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetNetworkLsaOk() (*V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa, bool)`

GetNetworkLsaOk returns a tuple with the NetworkLsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetNetworkLsa(v V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerNetworkLsa)`

SetNetworkLsa sets NetworkLsa field to given value.

### HasNetworkLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasNetworkLsa() bool`

HasNetworkLsa returns a boolean if a field has been set.

### GetRouterLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetRouterLsa() V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerRouterLsa`

GetRouterLsa returns the RouterLsa field if non-nil, zero value otherwise.

### GetRouterLsaOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetRouterLsaOk() (*V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerRouterLsa, bool)`

GetRouterLsaOk returns a tuple with the RouterLsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetRouterLsa(v V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerRouterLsa)`

SetRouterLsa sets RouterLsa field to given value.

### HasRouterLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasRouterLsa() bool`

HasRouterLsa returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetSequenceNumber() string`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetSequenceNumberOk() (*string, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetSequenceNumber(v string)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSummaryLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetSummaryLsa() V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerSummaryLsa`

GetSummaryLsa returns the SummaryLsa field if non-nil, zero value otherwise.

### GetSummaryLsaOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetSummaryLsaOk() (*V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerSummaryLsa, bool)`

GetSummaryLsaOk returns a tuple with the SummaryLsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetSummaryLsa(v V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInnerSummaryLsa)`

SetSummaryLsa sets SummaryLsa field to given value.

### HasSummaryLsa

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasSummaryLsa() bool`

HasSummaryLsa returns a boolean if a field has been set.

### GetType

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1DeviceRoutingOspfv2AreaLsdbGet200ResponseLsasInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


