# V2ChildalertlistPost200ResponseAlertListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedList** | Pointer to **[]string** |  | [optional] 
**AcknowledgementReason** | Pointer to **string** |  | [optional] 
**AlertBody** | Pointer to **string** |  | [optional] 
**AlertId** | Pointer to **string** |  | [optional] 
**AllowListed** | Pointer to **bool** |  | [optional] 
**ChildrenAlertList** | Pointer to [**V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertList**](V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertList.md) |  | [optional] 
**DescendantPresent** | Pointer to **bool** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**EnterpriseId** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**MuteListed** | Pointer to **bool** |  | [optional] 
**NotificationCreated** | Pointer to **bool** |  | [optional] 
**Occurrences** | Pointer to **int64** |  | [optional] 
**PeerDeviceId** | Pointer to **string** |  | [optional] 
**PeerInterfaceName** | Pointer to **string** |  | [optional] 
**PeerName** | Pointer to **string** |  | [optional] 
**Plane** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Recommendation** | Pointer to **string** |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TroubleshootingDisabledReason** | Pointer to **string** |  | [optional] 
**TroubleshootingEnabled** | Pointer to **bool** |  | [optional] 
**TunnelInterfaceName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV2ChildalertlistPost200ResponseAlertListInner

`func NewV2ChildalertlistPost200ResponseAlertListInner() *V2ChildalertlistPost200ResponseAlertListInner`

NewV2ChildalertlistPost200ResponseAlertListInner instantiates a new V2ChildalertlistPost200ResponseAlertListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ChildalertlistPost200ResponseAlertListInnerWithDefaults

`func NewV2ChildalertlistPost200ResponseAlertListInnerWithDefaults() *V2ChildalertlistPost200ResponseAlertListInner`

NewV2ChildalertlistPost200ResponseAlertListInnerWithDefaults instantiates a new V2ChildalertlistPost200ResponseAlertListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedList

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetAcknowledgedList() []string`

GetAcknowledgedList returns the AcknowledgedList field if non-nil, zero value otherwise.

### GetAcknowledgedListOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetAcknowledgedListOk() (*[]string, bool)`

GetAcknowledgedListOk returns a tuple with the AcknowledgedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedList

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetAcknowledgedList(v []string)`

SetAcknowledgedList sets AcknowledgedList field to given value.

### HasAcknowledgedList

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasAcknowledgedList() bool`

HasAcknowledgedList returns a boolean if a field has been set.

### GetAcknowledgementReason

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetAcknowledgementReason() string`

GetAcknowledgementReason returns the AcknowledgementReason field if non-nil, zero value otherwise.

### GetAcknowledgementReasonOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetAcknowledgementReasonOk() (*string, bool)`

GetAcknowledgementReasonOk returns a tuple with the AcknowledgementReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementReason

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetAcknowledgementReason(v string)`

SetAcknowledgementReason sets AcknowledgementReason field to given value.

### HasAcknowledgementReason

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasAcknowledgementReason() bool`

HasAcknowledgementReason returns a boolean if a field has been set.

### GetAlertBody

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetAlertBody() string`

GetAlertBody returns the AlertBody field if non-nil, zero value otherwise.

### GetAlertBodyOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetAlertBodyOk() (*string, bool)`

GetAlertBodyOk returns a tuple with the AlertBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertBody

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetAlertBody(v string)`

SetAlertBody sets AlertBody field to given value.

### HasAlertBody

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasAlertBody() bool`

HasAlertBody returns a boolean if a field has been set.

### GetAlertId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetAllowListed

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetAllowListed() bool`

GetAllowListed returns the AllowListed field if non-nil, zero value otherwise.

### GetAllowListedOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetAllowListedOk() (*bool, bool)`

GetAllowListedOk returns a tuple with the AllowListed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowListed

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetAllowListed(v bool)`

SetAllowListed sets AllowListed field to given value.

### HasAllowListed

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasAllowListed() bool`

HasAllowListed returns a boolean if a field has been set.

### GetChildrenAlertList

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetChildrenAlertList() V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertList`

GetChildrenAlertList returns the ChildrenAlertList field if non-nil, zero value otherwise.

### GetChildrenAlertListOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetChildrenAlertListOk() (*V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertList, bool)`

GetChildrenAlertListOk returns a tuple with the ChildrenAlertList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenAlertList

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetChildrenAlertList(v V2ChildalertlistPost200ResponseAlertListInnerChildrenAlertList)`

SetChildrenAlertList sets ChildrenAlertList field to given value.

### HasChildrenAlertList

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasChildrenAlertList() bool`

HasChildrenAlertList returns a boolean if a field has been set.

### GetDescendantPresent

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetDescendantPresent() bool`

GetDescendantPresent returns the DescendantPresent field if non-nil, zero value otherwise.

### GetDescendantPresentOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetDescendantPresentOk() (*bool, bool)`

GetDescendantPresentOk returns a tuple with the DescendantPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescendantPresent

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetDescendantPresent(v bool)`

SetDescendantPresent sets DescendantPresent field to given value.

### HasDescendantPresent

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasDescendantPresent() bool`

HasDescendantPresent returns a boolean if a field has been set.

### GetDeviceId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEndTime

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEnterpriseId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetEnterpriseId() string`

GetEnterpriseId returns the EnterpriseId field if non-nil, zero value otherwise.

### GetEnterpriseIdOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetEnterpriseIdOk() (*string, bool)`

GetEnterpriseIdOk returns a tuple with the EnterpriseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetEnterpriseId(v string)`

SetEnterpriseId sets EnterpriseId field to given value.

### HasEnterpriseId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasEnterpriseId() bool`

HasEnterpriseId returns a boolean if a field has been set.

### GetEntity

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetInterfaceName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetMuteListed

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetMuteListed() bool`

GetMuteListed returns the MuteListed field if non-nil, zero value otherwise.

### GetMuteListedOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetMuteListedOk() (*bool, bool)`

GetMuteListedOk returns a tuple with the MuteListed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteListed

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetMuteListed(v bool)`

SetMuteListed sets MuteListed field to given value.

### HasMuteListed

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasMuteListed() bool`

HasMuteListed returns a boolean if a field has been set.

### GetNotificationCreated

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetNotificationCreated() bool`

GetNotificationCreated returns the NotificationCreated field if non-nil, zero value otherwise.

### GetNotificationCreatedOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetNotificationCreatedOk() (*bool, bool)`

GetNotificationCreatedOk returns a tuple with the NotificationCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCreated

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetNotificationCreated(v bool)`

SetNotificationCreated sets NotificationCreated field to given value.

### HasNotificationCreated

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasNotificationCreated() bool`

HasNotificationCreated returns a boolean if a field has been set.

### GetOccurrences

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetOccurrences() int64`

GetOccurrences returns the Occurrences field if non-nil, zero value otherwise.

### GetOccurrencesOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetOccurrencesOk() (*int64, bool)`

GetOccurrencesOk returns a tuple with the Occurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrences

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetOccurrences(v int64)`

SetOccurrences sets Occurrences field to given value.

### HasOccurrences

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasOccurrences() bool`

HasOccurrences returns a boolean if a field has been set.

### GetPeerDeviceId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetPeerDeviceId() string`

GetPeerDeviceId returns the PeerDeviceId field if non-nil, zero value otherwise.

### GetPeerDeviceIdOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetPeerDeviceIdOk() (*string, bool)`

GetPeerDeviceIdOk returns a tuple with the PeerDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDeviceId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetPeerDeviceId(v string)`

SetPeerDeviceId sets PeerDeviceId field to given value.

### HasPeerDeviceId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasPeerDeviceId() bool`

HasPeerDeviceId returns a boolean if a field has been set.

### GetPeerInterfaceName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetPeerInterfaceName() string`

GetPeerInterfaceName returns the PeerInterfaceName field if non-nil, zero value otherwise.

### GetPeerInterfaceNameOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetPeerInterfaceNameOk() (*string, bool)`

GetPeerInterfaceNameOk returns a tuple with the PeerInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterfaceName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetPeerInterfaceName(v string)`

SetPeerInterfaceName sets PeerInterfaceName field to given value.

### HasPeerInterfaceName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasPeerInterfaceName() bool`

HasPeerInterfaceName returns a boolean if a field has been set.

### GetPeerName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetPeerName() string`

GetPeerName returns the PeerName field if non-nil, zero value otherwise.

### GetPeerNameOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetPeerNameOk() (*string, bool)`

GetPeerNameOk returns a tuple with the PeerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetPeerName(v string)`

SetPeerName sets PeerName field to given value.

### HasPeerName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasPeerName() bool`

HasPeerName returns a boolean if a field has been set.

### GetPlane

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetPlane() string`

GetPlane returns the Plane field if non-nil, zero value otherwise.

### GetPlaneOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetPlaneOk() (*string, bool)`

GetPlaneOk returns a tuple with the Plane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlane

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetPlane(v string)`

SetPlane sets Plane field to given value.

### HasPlane

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasPlane() bool`

HasPlane returns a boolean if a field has been set.

### GetReason

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRecommendation

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetRuleId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSeverity

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSiteId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStartTime

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTroubleshootingDisabledReason

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetTroubleshootingDisabledReason() string`

GetTroubleshootingDisabledReason returns the TroubleshootingDisabledReason field if non-nil, zero value otherwise.

### GetTroubleshootingDisabledReasonOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetTroubleshootingDisabledReasonOk() (*string, bool)`

GetTroubleshootingDisabledReasonOk returns a tuple with the TroubleshootingDisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTroubleshootingDisabledReason

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetTroubleshootingDisabledReason(v string)`

SetTroubleshootingDisabledReason sets TroubleshootingDisabledReason field to given value.

### HasTroubleshootingDisabledReason

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasTroubleshootingDisabledReason() bool`

HasTroubleshootingDisabledReason returns a boolean if a field has been set.

### GetTroubleshootingEnabled

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetTroubleshootingEnabled() bool`

GetTroubleshootingEnabled returns the TroubleshootingEnabled field if non-nil, zero value otherwise.

### GetTroubleshootingEnabledOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetTroubleshootingEnabledOk() (*bool, bool)`

GetTroubleshootingEnabledOk returns a tuple with the TroubleshootingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTroubleshootingEnabled

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetTroubleshootingEnabled(v bool)`

SetTroubleshootingEnabled sets TroubleshootingEnabled field to given value.

### HasTroubleshootingEnabled

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasTroubleshootingEnabled() bool`

HasTroubleshootingEnabled returns a boolean if a field has been set.

### GetTunnelInterfaceName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetTunnelInterfaceName() string`

GetTunnelInterfaceName returns the TunnelInterfaceName field if non-nil, zero value otherwise.

### GetTunnelInterfaceNameOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetTunnelInterfaceNameOk() (*string, bool)`

GetTunnelInterfaceNameOk returns a tuple with the TunnelInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelInterfaceName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetTunnelInterfaceName(v string)`

SetTunnelInterfaceName sets TunnelInterfaceName field to given value.

### HasTunnelInterfaceName

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasTunnelInterfaceName() bool`

HasTunnelInterfaceName returns a boolean if a field has been set.

### GetType

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2ChildalertlistPost200ResponseAlertListInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2ChildalertlistPost200ResponseAlertListInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V2ChildalertlistPost200ResponseAlertListInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


