# V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownTransitions** | Pointer to [**[]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner**](V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner.md) |  | [optional] 
**SessionSlas** | Pointer to [**[]V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlaneSessionSlasInner**](V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlaneSessionSlasInner.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewV1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane

`func NewV1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane() *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane`

NewV1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BackboneHealthDeviceDeviceIdPost200ResponseDataPlaneWithDefaults

`func NewV1BackboneHealthDeviceDeviceIdPost200ResponseDataPlaneWithDefaults() *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane`

NewV1BackboneHealthDeviceDeviceIdPost200ResponseDataPlaneWithDefaults instantiates a new V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownTransitions

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) GetDownTransitions() []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner`

GetDownTransitions returns the DownTransitions field if non-nil, zero value otherwise.

### GetDownTransitionsOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) GetDownTransitionsOk() (*[]V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner, bool)`

GetDownTransitionsOk returns a tuple with the DownTransitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownTransitions

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) SetDownTransitions(v []V1BackboneHealthDeviceDeviceIdPost200ResponseControlPlaneControlTransitionsTransitionsInner)`

SetDownTransitions sets DownTransitions field to given value.

### HasDownTransitions

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) HasDownTransitions() bool`

HasDownTransitions returns a boolean if a field has been set.

### GetSessionSlas

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) GetSessionSlas() []V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlaneSessionSlasInner`

GetSessionSlas returns the SessionSlas field if non-nil, zero value otherwise.

### GetSessionSlasOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) GetSessionSlasOk() (*[]V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlaneSessionSlasInner, bool)`

GetSessionSlasOk returns a tuple with the SessionSlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSlas

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) SetSessionSlas(v []V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlaneSessionSlasInner)`

SetSessionSlas sets SessionSlas field to given value.

### HasSessionSlas

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) HasSessionSlas() bool`

HasSessionSlas returns a boolean if a field has been set.

### GetStatus

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1BackboneHealthDeviceDeviceIdPost200ResponseDataPlane) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


