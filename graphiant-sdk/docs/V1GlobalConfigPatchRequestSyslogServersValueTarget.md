# V1GlobalConfigPatchRequestSyslogServersValueTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**GlobalId** | Pointer to **int64** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**IsGlobalSync** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to **string** |  | [optional] 
**VrfId** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1GlobalConfigPatchRequestSyslogServersValueTarget

`func NewV1GlobalConfigPatchRequestSyslogServersValueTarget() *V1GlobalConfigPatchRequestSyslogServersValueTarget`

NewV1GlobalConfigPatchRequestSyslogServersValueTarget instantiates a new V1GlobalConfigPatchRequestSyslogServersValueTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GlobalConfigPatchRequestSyslogServersValueTargetWithDefaults

`func NewV1GlobalConfigPatchRequestSyslogServersValueTargetWithDefaults() *V1GlobalConfigPatchRequestSyslogServersValueTarget`

NewV1GlobalConfigPatchRequestSyslogServersValueTargetWithDefaults instantiates a new V1GlobalConfigPatchRequestSyslogServersValueTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGlobalId

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetHost

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInterfaceName

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIsGlobalSync

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetIsGlobalSync() bool`

GetIsGlobalSync returns the IsGlobalSync field if non-nil, zero value otherwise.

### GetIsGlobalSyncOk

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetIsGlobalSyncOk() (*bool, bool)`

GetIsGlobalSyncOk returns a tuple with the IsGlobalSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalSync

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) SetIsGlobalSync(v bool)`

SetIsGlobalSync sets IsGlobalSync field to given value.

### HasIsGlobalSync

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) HasIsGlobalSync() bool`

HasIsGlobalSync returns a boolean if a field has been set.

### GetName

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSeverity

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTransport

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetVrfId

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetVrfId() int64`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) GetVrfIdOk() (*int64, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) SetVrfId(v int64)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *V1GlobalConfigPatchRequestSyslogServersValueTarget) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


