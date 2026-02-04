# AssuranceDnsProxyEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsproxyEntryId** | **string** | dns proxy table entry id (required) | 
**IpList** | **[]string** |  | 
**Name** | **string** | name of the entry  (required) | 
**NameText** | **string** | user defined name of the record (required) | 
**PortList** | **[]string** |  | 
**Protocol** | **string** | protocol to exclude (required) | 

## Methods

### NewAssuranceDnsProxyEntry

`func NewAssuranceDnsProxyEntry(dnsproxyEntryId string, ipList []string, name string, nameText string, portList []string, protocol string, ) *AssuranceDnsProxyEntry`

NewAssuranceDnsProxyEntry instantiates a new AssuranceDnsProxyEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssuranceDnsProxyEntryWithDefaults

`func NewAssuranceDnsProxyEntryWithDefaults() *AssuranceDnsProxyEntry`

NewAssuranceDnsProxyEntryWithDefaults instantiates a new AssuranceDnsProxyEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsproxyEntryId

`func (o *AssuranceDnsProxyEntry) GetDnsproxyEntryId() string`

GetDnsproxyEntryId returns the DnsproxyEntryId field if non-nil, zero value otherwise.

### GetDnsproxyEntryIdOk

`func (o *AssuranceDnsProxyEntry) GetDnsproxyEntryIdOk() (*string, bool)`

GetDnsproxyEntryIdOk returns a tuple with the DnsproxyEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsproxyEntryId

`func (o *AssuranceDnsProxyEntry) SetDnsproxyEntryId(v string)`

SetDnsproxyEntryId sets DnsproxyEntryId field to given value.


### GetIpList

`func (o *AssuranceDnsProxyEntry) GetIpList() []string`

GetIpList returns the IpList field if non-nil, zero value otherwise.

### GetIpListOk

`func (o *AssuranceDnsProxyEntry) GetIpListOk() (*[]string, bool)`

GetIpListOk returns a tuple with the IpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpList

`func (o *AssuranceDnsProxyEntry) SetIpList(v []string)`

SetIpList sets IpList field to given value.


### GetName

`func (o *AssuranceDnsProxyEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssuranceDnsProxyEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssuranceDnsProxyEntry) SetName(v string)`

SetName sets Name field to given value.


### GetNameText

`func (o *AssuranceDnsProxyEntry) GetNameText() string`

GetNameText returns the NameText field if non-nil, zero value otherwise.

### GetNameTextOk

`func (o *AssuranceDnsProxyEntry) GetNameTextOk() (*string, bool)`

GetNameTextOk returns a tuple with the NameText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameText

`func (o *AssuranceDnsProxyEntry) SetNameText(v string)`

SetNameText sets NameText field to given value.


### GetPortList

`func (o *AssuranceDnsProxyEntry) GetPortList() []string`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *AssuranceDnsProxyEntry) GetPortListOk() (*[]string, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *AssuranceDnsProxyEntry) SetPortList(v []string)`

SetPortList sets PortList field to given value.


### GetProtocol

`func (o *AssuranceDnsProxyEntry) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AssuranceDnsProxyEntry) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AssuranceDnsProxyEntry) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


