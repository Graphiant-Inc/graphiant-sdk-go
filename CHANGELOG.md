# Changelog

All notable changes to the Graphiant SDK Go will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [26.8.0] - 2026-08-26

### Added
- **API endpoints:**
  - `PUT /v1/adaptive-fec`
- **Models (new):**
  - `ManaV2AdaptiveFecConfiguration` (`adaptiveFecEnabled`, `slaClasses`)
  - `V1AdaptiveFecPutRequest`

### Changed
- Updated to API specification version 26.8.0
- **Version:**
  - Updated version constant to v26.8.0
  - Updated API documentation reference to `graphiant_api_docs_v26.8.0.json`
- **Documentation:** updated SDK generation examples in README to use `graphiant_api_docs_v26.8.0.json`; **SECURITY.md** supported-versions table updated for **26.8.0**
- **Models (updated):**
  - `AlertserviceZendeskDetails`: added required `zendeskClientId` and `zendeskClientSecret` (OAuth); `zendeskApiToken` and `zendeskEmail` are now optional and deprecated
  - `V1EnterpriseConfigurationGetResponse`: added `adaptiveFecConfig` field

### Removed
- **API endpoints:** none
- **Models:** none

## [26.7.0] - 2026-07-27

### Added
- **API endpoints (B2B Extranet — consumers):**
  - `GET /v1/extranet/b2b/consumers/{customerId}`
  - `DELETE /v1/extranet/b2b/consumers/{id}`
  - `PUT /v1/extranet/b2b/consumers/{id}`
  - `GET /v1/extranet/b2b/consumers/{id}/device-status`
  - `PUT /v1/extranet/b2b/consumers/{id}/prefixes`
- **API endpoints (B2B Extranet — customers):**
  - `POST /v1/extranet/b2b/customers`
  - `GET /v1/extranet/b2b/customers/summary`
  - `DELETE /v1/extranet/b2b/customers/{id}`
  - `PUT /v1/extranet/b2b/customers/{id}`
  - `GET /v1/extranet/b2b/customers/{id}/details`
  - `GET /v1/extranet/b2b/customers/{id}/matches/summary`
  - `POST /v1/extranet/b2b/customers/{id}/retry`
- **API endpoints (B2B Extranet — matches):**
  - `POST /v1/extranet/b2b/matches`
  - `POST /v1/extranet/b2b/matches/customers`
  - `PUT /v1/extranet/b2b/matches/pause`
  - `POST /v1/extranet/b2b/matches/review`
  - `DELETE /v1/extranet/b2b/matches/{matchId}`
  - `GET /v1/extranet/b2b/matches/{matchId}`
  - `PUT /v1/extranet/b2b/matches/{matchId}`
  - `GET /v1/extranet/b2b/matches/{matchId}/details`
  - `POST /v1/extranet/b2b/matches/{matchId}/consumer`
  - `POST /v1/extranet/b2b/matches/{matchId}/consumer/check`
  - `PUT /v1/extranet/b2b/matches/{matchId}/status`
- **API endpoints (B2B Extranet — producer):**
  - `POST /v1/extranet/b2b/producer`
  - `POST /v1/extranet/b2b/producer/review`
  - `DELETE /v1/extranet/b2b/producer/{id}`
  - `GET /v1/extranet/b2b/producer/{id}`
  - `PUT /v1/extranet/b2b/producer/{id}`
  - `GET /v1/extranet/b2b/producer/{id}/customers`
  - `GET /v1/extranet/b2b/producer/{id}/device-status`
  - `PUT /v1/extranet/b2b/producer/{id}/status`
- **API endpoints (B2B Extranet — monitoring & services):**
  - `GET /v1/extranet/b2b/services/summary`
  - `POST /v1/extranet/b2b/monitoring/peering-service/bandwidth-usage`
  - `POST /v1/extranet/b2b/monitoring/peering-service/consumers-usage-top`
  - `POST /v1/extranet/b2b/monitoring/peering-service/consumption-overview`
  - `POST /v1/extranet/b2b/monitoring/peering-service/service-customer-list`
  - `POST /v1/extranet/b2b/monitoring/peering-service/service-health`
  - `POST /v1/extranet/b2b/monitoring/peering-service/service-overtime-consumption`
  - `POST /v1/extranet/sites/usage-top`
- **Models (new):**
  - Added B2B extranet match types (`ManaV2B2bExtranetMatch`, `ManaV2B2bExtranetMatchConsumerDetails`, `ManaV2B2BExtranetMatchConsumerDetailsCustomer`, `ManaV2B2BExtranetMatchConsumerDetailsProducerPrefix`, `ManaV2B2BExtranetMatchConsumerDetailsService`)
  - Added extranet consumer/producer policy types (`ManaV2ExtranetServiceConsumerPolicy`, `ManaV2ExtranetServiceProducerPolicy`, `ManaV2ExtranetServiceProducerCustomer`, `ManaV2ExtranetServicePolicyResponse`)
  - Added extranet service summary types (`ManaV2ExtranetServiceSummary`, `ManaV2ExtranetServiceCustomerSummary`, `ManaV2ExtranetServiceCustomerMatchSummary`, `ManaV2ExtranetServiceCustomerInvite`)
  - Added extranet NAT translation types (`ManaV2ExtranetNatTranslationMode`, `ManaV2ExtranetNatTranslationCentralized`, `ManaV2ExtranetNatTranslationDecentralized`, `ManaV2ExtranetNatTranslationPeerToPeer`, `ManaV2ExtranetNatTranslationPeerToPeerPrefix`, `ManaV2ExtranetNatTranslationDevicePrefixes`)
  - Added `ManaV2ExtranetConsumerLanPrefixes`
  - Added OSPF authentication types (`ManaV2OspfAuthentication`, `ManaV2NullableOspfAuthenticationConfig`)
  - Added `V1ExtranetB2b*` request/response types for all new consumer, customer, match, producer, and services endpoints

### Changed
- Updated to API specification version 26.7.0
- **Version:**
  - Updated version constant to v26.7.0
  - Updated API documentation reference to `graphiant_api_docs_v26.7.0.json`
  - Removed superseded bundle `graphiant_api_docs_v26.6.0.json` from the repository
- **Documentation:** updated SDK generation examples in README to use `graphiant_api_docs_v26.7.0.json`; **SECURITY.md** supported-versions table updated for **26.7.0**
- **Models (updated):**
  - `AssuranceKpiMetric`: added `tokenUsage` field
  - `AssuranceUserDefinition`: added `tokensUsage` and `tokensUsageDays` fields
  - `ManaV2GatewayDetails`: added `ipsecGatewayPeers` field
  - `ManaV2OspfInterfaceConfig`: added `authentication` field
  - `V1InvitationEmailPostRequest`: added `serviceType` field

### Removed
- **API endpoints:** none
- **Models:** none

## [26.6.0] - 2026-07-08

### Added
- **API endpoints:**
  - `GET /v1/lan-segments/interfaces/public`
  - `POST /v1/pvif`
  - `GET /v1/pvif/summary`
  - `GET /v1/pvif/{id}`
  - `PUT /v1/pvif/{id}`
  - `DELETE /v1/pvif/{id}`
  - `GET /v1/pvif/{id}/details`
  - `GET /v1/regions/{regionId}/gateways`
  - `GET /v1/ztagent/agents`
  - `POST /v2/assurance/ai-adoption-summary`
  - `POST /v2/assurance/create-ai-adoption-approve-entry`
  - `DELETE /v2/assurance/delete-ai-adoption-approve-entry`
  - `POST /v2/assurance/get-app-names`
  - `GET /v2/assurance/read-ai-adoption-approve-entries`
  - `POST /v2/assurance/update-ai-adoption-approve-entry`
- **Models:**
  - Added PVIF types (`v1Pvif*`, `v1PvifSummaryGetResponse`, `v1PvifIdDetailsGetResponse`)
  - Added LAN segment public interface types (`v1LanSegmentsInterfacesPublicGetResponse`)
  - Added regional gateway types (`v1RegionsRegionIdGatewaysGetResponse` and nested gateway schemas)
  - Added ZTAgent agent listing types (`v1ZtagentAgentsGetResponse`)
  - Added assurance AI adoption types (`v2AssuranceAiAdoption*`, `v2AssuranceCreateAiAdoptionApproveEntry*`, `v2AssuranceReadAiAdoptionApproveEntriesGetResponse`, `v2AssuranceUpdateAiAdoptionApproveEntryPostRequest`)
  - Added assurance app name types (`v2AssuranceGetAppNamesPostRequest`, `v2AssuranceGetAppNamesPostResponse`)

### Changed
- Updated to API specification version 26.6.0
- **Version:**
  - Updated version constant to v26.6.0
  - Updated API documentation reference to `graphiant_api_docs_v26.6.0.json`
  - Removed superseded bundle `graphiant_api_docs_v26.5.0.json` from the repository
- **Documentation:** updated SDK generation examples in README and CONTRIBUTING to use `graphiant_api_docs_v26.6.0.json`; **SECURITY.md** supported-versions table for **26.6.0**
- **Tooling:** `scripts/generate.sh` now requires OpenAPI Generator `>= 7.23.0`
- **API endpoints (updated):** existing routes regenerated for revised schemas (see regenerated **`DefaultAPI`**); path additions are listed under **Added** above
- **Models:** updated (OpenAPI spec refresh; representative schema updates include `assuranceAppNameRecord`, `manaV2GuestConsumerSiteToSiteVpnConfig`, `manaV2SiteDeviceSummary`, `upgradeRollout`, `v1EnterpriseAllocationGetResponse`, `v1GatewaysReferenceConsumerGetResponse`)

### Removed
- **API endpoints:**
  - `GET /v1/extranet-public-vif`
  - `POST /v1/extranet-public-vif`
  - `POST /v1/extranet-public-vif/check`
  - `GET /v1/extranet-public-vif/{id}`
  - `PUT /v1/extranet-public-vif/{id}`
  - `DELETE /v1/extranet-public-vif/{id}`
- **Models:** removed extranet public VIF client methods and associated generated API surface (`v1ExtranetPublicVif*`); use the new PVIF endpoints instead

## [26.5.0] - 2026-06-10

### Added
- **API endpoints:**
  - `GET /v1/extranet-public-vif`
  - `POST /v1/extranet-public-vif`
  - `POST /v1/extranet-public-vif/check`
  - `GET /v1/extranet-public-vif/{id}`
  - `PUT /v1/extranet-public-vif/{id}`
  - `DELETE /v1/extranet-public-vif/{id}`
  - `GET /v1/msp/managed-enterprise-contract-info`
  - `GET /v1/sites/map/details`
- **Models:**
  - Added extranet public VIF types (`v1ExtranetPublicVif*`, `manaV2PublicVif*`, and nested consumer/producer policy, device, dynamic/fixed, and NAT schemas)
  - Added site map types (`v1SitesMapDetailsGetResponse`, `manaV2LanSegmentSitesMap`, `manaV2SiteInformation`, `manaV2SiteDeviceSummary`, `manaV2SiteLanSegmentDeviceBuckets`)
  - Added MSP managed enterprise contract types (`v1MspManagedEnterpriseContractInfoGetResponse`, `manaV2ManagedEnterpriseContractInfo`)
  - Added billing contract types (`commonBillingContract`, `commonBillingTimePeriod`)

### Changed
- Updated to API specification version 26.5.0
- **Version:**
  - Updated version constant to v26.5.0
  - Updated API documentation reference to `graphiant_api_docs_v26.5.0.json`
- **Documentation:** updated SDK generation example in README to use `graphiant_api_docs_v26.5.0.json` and `packageVersion=26.5.0`; **SECURITY.md** supported-versions table for **26.5.0**
- **API endpoints (updated):** existing routes regenerated for revised schemas (see regenerated **`DefaultAPI`**); path additions are listed under **Added** above
- **Models:** updated (OpenAPI spec refresh; representative schema updates include `assuranceClientSession`, `assuranceSite`, `manaV2ApplicationProfile`, `searchEdgeSummary`, `v1EnterpriseContractPutRequest`, `v1EnterprisesPatchRequest`, `v1EnterprisesPutRequest`, `v2AssuranceApplicationdetailsbynamePostRequest`, `v2AssuranceTopologyClientSessionsPostRequest`, `v2AssuranceTopologyClientSummariesPostRequest`, `v2AssuranceTopologyClientSummariesPostResponseSummary`, `v2AssuranceTopologyOverviewPostRequest`, `v2AssuranceTopologySiteSummariesPostRequest`)

### Removed
- **API endpoints:** none
- **Models:** none

## [26.4.0] - 2026-04-30

### Added
- **API endpoints:**
  - `GET /v1/global/content-filters`
  - `POST /v1/global/content-filters`
  - `DELETE /v1/global/content-filters/{globalContentFilterId}`
  - `GET /v1/global/content-filters/{globalContentFilterId}`
  - `PUT /v1/global/content-filters/{globalContentFilterId}`
  - `GET /v1/global/domain-categories`
  - `GET /v1/software/rollouts`
  - `POST /v1/software/rollouts`
  - `PUT /v1/software/rollouts`
  - `DELETE /v1/software/rollouts/{id}`
  - `GET /v1/software/rollouts/{id}`
  - `POST /v1/software/rollouts/schedule`
  - `GET /v1/ztagent/bindings`
  - `PUT /v1/ztagent/bindings`
  - `GET /v2/monitoring/macsec/{deviceId}/status`
- **Models:**
  - Added global content filter types (`v1GlobalContentFilters*`, `manaV2GlobalContentFilterConfig`, `manaV2GlobalContentFilterRule`, and nested row/rule/site/lan entry schemas)
  - Added domain category types (`v1GlobalDomainCategoriesGetResponse`, `manaV2DomainCategory`)
  - Added software rollout types (`v1SoftwareRollouts*`, `upgradeRollout`, `upgradeRolloutConfig`, `upgradeRolloutDevice`, `upgradeRecurringSchedule`, `upgradeMonthlyRecurrence`, `upgradeWeeklyRecurrence`, `upgradeYearlyRecurrence`)
  - Added ZTAgent binding types (`v1ZtagentBindingsGetResponse`, `v1ZtagentBindingsPutRequest`, `v1ZtagentBindingsPutResponse`)
  - Added MACsec monitoring types (`v2MonitoringMacsecDeviceIdStatusGetResponse` and nested MACsec status schemas)
  - Added `manaV2RegionCoordinates`

### Changed
- Updated to API specification version 26.4.0
- **Version:**
  - Updated version constant to v26.4.0
  - Updated API documentation reference to `graphiant_api_docs_v26.4.0.json`
  - Removed superseded bundle `graphiant_api_docs_v26.3.1.json` from the repository
- **Documentation:** updated SDK generation example in README to use `graphiant_api_docs_v26.4.0.json` and `packageVersion=26.4.0`; **SECURITY.md** supported-versions table for **26.4.0**
- **API endpoints (updated):** existing routes regenerated for revised schemas (see regenerated **`DefaultAPI`**); path additions are listed under **Added** above
- **Models:** updated (OpenAPI spec refresh; representative schema updates include `alertserviceNotificationBody`, `assuranceClientSession`, `assuranceRegion`, `upgradeUpgradeOccurrence`, `upgradeUpgradeSummary`, `v2AssuranceTopologyClientSummariesPostResponseSummary`, `manaV2Region`)

### Removed
- **API endpoints:** none
- **Models:** none

## [26.3.2] - 2026-03-27

### Added
- **`GRAPHIANT_ACCESS_TOKEN`** support aligned with **graphiant-sdk-python**: `AccessTokenFromEnv`, `AuthorizationBearerFromEnv`, `ConfigureHostFromEnv`, **`AuthorizationBearerFromEnvOrLogin`** / **`LoginBearerFromEnvCredentials`** (fallback to **`GRAPHIANT_USERNAME`** / **`GRAPHIANT_PASSWORD`** when token unset) in **`auth_env.go`** (optional **`GRAPHIANT_API_HOST`** / **`GRAPHIANT_HOST`** for API base URL).

### Changed
- **Version:** SDK **v26.3.2** (same OpenAPI bundle filename as 26.3.1: `graphiant_api_docs_v26.3.1.json`; no API spec change for this patch).

### Documentation
- README: **`GRAPHIANT_ACCESS_TOKEN`** quick start and environment-variable table; clarified **`packageVersion`** vs **`-i`** spec filename for regeneration; **SECURITY.md** supported-versions note for current patch.

## [26.3.1] - 2026-03-26

### Added
- **API endpoints:** none
- **Models:**
  - Added `manaV2BgpDynamicNeighborOperPeer`
  - Added `manaV2NullableMaCsecRekeyInterval`
  - Added `manaV2NullableMaCsecReplayProtectionWindowSize`

### Changed
- Updated to API specification version 26.3.1
- **Version:**
  - Updated version constant to v26.3.1
  - Updated API documentation reference to `graphiant_api_docs_v26.3.1.json`
- **Documentation:** updated SDK generation example in README to use `graphiant_api_docs_v26.3.1.json` and `packageVersion=26.3.1`
- **API endpoints (updated):**
  - `GET /v1/device/routing/ospfv2/statistics`
  - `GET /v1/device/routing/ospfv3/statistics`
  - `POST /v1/device/routing/rib/route-count`
  - `GET /v1/devices/inventory`
  - `DELETE /v1/devices/inventory/serial-num`
  - `GET /v1/devices/routing/vrf/protocol-route-count`
  - `GET /v1/devices/{deviceId}/circuits/vrf-associations`
  - `GET /v1/devices/{deviceId}/vrf/protocols`
  - `GET /v1/devices/{deviceId}/vrrp`
  - `GET /v1/diagnostic/speedtest-servers`
  - `GET /v1/edges-hardware/assigned`
  - `GET /v1/edges-summary`
  - `GET /v1/enterprises`
  - `GET /v1/enterprises/managed`
  - `GET /v1/extranets-b2b-peering/consumer/{customerId}/consumer-details`
  - `GET /v1/extranets-b2b-peering/match/service-to-customer/{id}`
  - `GET /v1/extranets-monitoring/lan-segments`
  - `GET /v1/extranets-monitoring/nat-usage`
  - `GET /v1/global/device-status`
  - `GET /v1/global/ipfix/device`
  - `GET /v1/global/ntps/device`
  - `GET /v1/global/ntps/site`
  - `GET /v1/global/site-status`
  - `GET /v1/global/snmps/device`
  - `GET /v1/global/syslogs/device`
  - `GET /v1/software/releases/download`
  - `GET /v1/users/{id}/enterprises`
  - `GET /v1/users/{id}/groups`
  - `GET /v1/users/{id}/groups/enterprises`
  - `GET /v1/users/{id}/groups/root`
  - `DELETE /v2/assurance/deleteclassifiedapplication`
  - `GET /v2/monitoring/extranet/edge-status`
  - `GET /v2/monitoring/extranet/service-status`
  - `GET /v2/monitoring/extranet/service-status/details`
  - `GET /v2/monitoring/extranet/site-status`
  - `GET /v2/monitoring/extranet/status-details`
- **Models:** updated (1390 schema changes in OpenAPI spec)

### Removed
- **API endpoints:** none
- **Models:**
  - Removed `healthcheckOdpStatusDetails`
  - Removed `healthcheckOnboardingStatusDetails`
  - Removed `healthcheckStatusDetails`
  - Removed `healthcheckT2StatusDetails`
  - Removed `v1HealthcheckDevicesGetResponse`

## [26.2.1] - 2026-02-26

### Added
- **Auth:**
  - New endpoint: `POST /v1/users/passwords/expire`
  - New models: `V1UsersPasswordsExpirePostRequest`, `V1UsersPasswordsExpirePostResponse`
- **Alerting:** new model `AlertserviceZendeskDetails`
- **IAM:** new model `IamFailedUser`
- **Device:** new model `ManaV2NullableGatewayConfig`
- Updated to API specification version 26.2.1

### Changed
- **Version:**
  - Updated version constant to v26.2.1
  - Updated API documentation reference to `graphiant_api_docs_v26.2.1.json`
- **Documentation:** updated SDK generation example in README to use `graphiant_api_docs_v26.2.1.json` and `packageVersion=26.2.1`
- **API endpoints (updated):**
  - `GET /v1/device/routing/bgp/nbrs/counters`
  - `GET /v1/diagnostic/gnmi-ping`
  - `GET /v1/diagnostic/speedtest-servers`
  - `GET /v1/enterprises/{enterpriseId}/device-status`
  - `GET /v1/extranets-monitoring/lan-segments`
  - `GET /v1/extranets-monitoring/nat-usage`
  - `GET /v1/global/device-status`
  - `GET /v1/global/ipfix/site`
  - `GET /v1/global/site-status`
  - `GET /v1/global/snmps/site`
  - `GET /v1/global/syslogs/site`
  - `GET /v1/software/running/details`
- **Models (updated):** `AlertserviceIntegration`, `AlertserviceIntegrationDetails`, `ManaV2Interface`, `ManaV2InterfaceIpConfig`, `ManaV2InterfaceVlan`, `StatsmonV2Node`, `StatsmonV2NodeCircuitInfo`, `V2DeviceDeviceIdTopologyPostRequest`

### Removed
- **API endpoints:**
  - **DELETE**
    - `DELETE /v1/enterprises/self`
    - `DELETE /v1/policy/prefix-sets/{id}`
    - `DELETE /v1/portal/apikeys`
    - `DELETE /v2/assistant/delete-conversation/{conversationId}`
    - `DELETE /v2/assistant/{conversationId}`
  - **GET**
    - `GET //v1/devices/oauth/redirect`
    - `GET /v1/alarm-history`
    - `GET /v1/alarms-events`
    - `GET /v1/alarms-list`
    - `GET /v1/device/routing/bgp/nbr/stats`
    - `GET /v1/device/routing/ospfv2/area/interfaceid`
    - `GET /v1/device/routing/ospfv3/area/interface/nbrid`
    - `GET /v1/device/routing/ospfv3/area/interfaceid`
    - `GET /v1/device/routing/vrf/bgp/graphiant-eiroute-count`
    - `GET /v1/devices/{deviceId}/edges`
    - `GET /v1/devices/{deviceId}/ndcache`
    - `GET /v1/devices/{deviceId}/policy/customapplications`
    - `GET /v1/devices/{deviceId}/versions/compare`
    - `GET /v1/devices/{deviceId}/vrf/bgp/as`
    - `GET /v1/event/device`
    - `GET /v1/event/enterprise`
    - `GET /v1/event/system`
    - `GET /v1/global/prefix-sets/device`
    - `GET /v1/global/prefix-sets/site`
    - `GET /v1/global/routing-policies/device`
    - `GET /v1/global/routing-policies/site`
    - `GET /v1/global/traffic-policies/device`
    - `GET /v1/global/traffic-policies/site`
    - `GET /v1/groups/{id}`
    - `GET /v1/portal/apikeys`
    - `GET /v1/portal/private/details`
    - `GET /v1/portal/private/inventory_details`
    - `GET /v1/software/release/notes`
    - `GET /v1/tt/{ttIdentity}/device-status`
    - `GET /v2/assurance/bucket-app-servers/all`
  - **PATCH**
    - `PATCH /v1/{id}/password/recover`
  - **POST**
    - `POST /v1/b2b-extranet-monitoring/filter`
    - `POST /v1/bwtracker/region/cloud/chart`
    - `POST /v1/bwtracker/region/cloud/csv`
    - `POST /v1/bwtracker/region/cloud/summary`
    - `POST /v1/devices/inventory/serial-num`
    - `POST /v1/diagnostic/ping-pause-resume`
    - `POST /v1/event/system/ack`
    - `POST /v1/extranet/sites-usage`
    - `POST /v1/global/config/site`
    - `POST /v1/monitoring/circuits/bandwidth`
    - `POST /v1/monitoring/circuits/incidents`
    - `POST /v1/monitoring/circuits/summary`
    - `POST /v1/monitoring/circuits/utilization`
    - `POST /v1/monitoring/circuits/visualization`
    - `POST /v1/policy/prefix-sets`
    - `POST /v1/portal/apikeys`
    - `POST /v1/portal/private`
    - `POST /v1/portal/private/register`
    - `POST /v1/portal/private/sync`
    - `POST /v2/assurance/endpoint-intel`
    - `POST /v2/assurance/flow-summary`
    - `POST /v2/assurance/topology-flows`
    - `POST /v2/assurance/version`
    - `POST /v2/monitoring/ospf`
    - `POST /v2/monitoring/queue`
    - `POST /v2/monitoring/system/generic`
    - `POST /v2/site/{siteId}/detail`
    - `POST /v2/version`
  - **PUT**
    - `PUT /v1/alarm-mute/{alarmId}`
    - `PUT /v1/policy/prefix-sets/{id}`
- **Models:** removed 136 component models (not exhaustively listed; includes alarms/assurance/events/statsmon/portal and related request/response types)

## [26.1.1] - 2026-02-01

### Added
- Updated to API specification version 26.1.1
- Enhanced API coverage with new endpoints and models from the latest Graphiant portal API

### Changed
- Updated version constant to v26.1.1
- SDK aligned with Graphiant portal API v26.1.1

## [25.12.1] - 2025-12-17

### Added
- Updated to API specification version 25.12.1
- Enhanced API coverage with new endpoints and models

### Changed
- Updated version constant to v25.12.1

## [25.11.1] - 2025-11-11

### Added
- Updated to API specification version 25.11.1
- Enhanced API coverage with new endpoints and models

### Changed
- **API Specification Optimization**: Major API specification update with significant improvements:
  - Reduced specification size from 9.8M to 1.5M (~85% reduction) through schema optimization and reuse
  - Enhanced documentation for better developer experience
  - Cleaner API names: Response APIs no longer include HTTP status codes
  - Reusable schemas: Common schemas share the same inner API names across different endpoints

### Migration Notes
- Response API names have been updated to remove HTTP status codes:
  - `Post200Response` → `PostResponse`
  - `Get200Response` → `GetResponse`
  - `Put202Response` → `PutResponse`
  - `Put204Response` → `PutResponse`
  - `Post201Response` → `PostResponse`
- See [Migration Guide](README.md#-migration-guide-upgrading-from-25102-to-25111) for detailed migration instructions

## [25.10.2] - 2025-10-15

### Fixed
- Hotfix release addressing critical issues (TE-4117)
- SDK upgrade to version 25.10.2

## [25.10.1] - 2025-10-08

### Added
- Updated to API specification version 25.10.1
- Enhanced API coverage with new endpoints and models (TE-4100)

### Changed
- Updated version constant to v25.10.1

## [25.9.2] - 2025-09-25

### Added
- Updated to API specification version 25.9.2
- Enhanced API coverage with new endpoints and models (TE-4067)

### Changed
- Updated version constant to v25.9.2

## [25.9.1] - 2025-09-23

### Added
- Updated to API specification version 25.9.1
- Enhanced API coverage with new endpoints and models (TE-4092, TE-4062)

### Changed
- Updated version constant to v25.9.1

## [25.8.1] - 2025-08-22

### Added
- **Device Configuration Helper**: Added `PollAndPutDeviceConfig` wrapper function to check device status and push configuration (TE-3040)
  - Function polls device status and executes configuration when device is ready
  - Simplifies device configuration workflow
- Updated to API specification version 25.8.1
- Enhanced API coverage with new endpoints and models

### Changed
- Updated README with improved documentation and examples (TE-3993, TE-3909)
- Updated version constant to v25.8.1

### Documentation
- Enhanced README with better examples and usage patterns
- Improved documentation structure and clarity

## [25.7.1] - TBD

### Added
- Updated to API specification version 25.7.1
- Enhanced API coverage with new endpoints and models

### Changed
- Updated version constant to v25.7.1

## [25.6.2] - 2025-06-13

### Added
- **CI/CD Improvements**: Enhanced CI/CD pipeline and build processes (TE-3814)
- Improved build and release automation

### Changed
- Cleaned up repository structure
- Removed unwanted files from repository

## [25.6.1] - 2025-06-13

### Added
- **Initial Release**: First public release of Graphiant SDK Go
- Complete API coverage for all Graphiant REST API endpoints
- Type-safe Go client generated from OpenAPI specification
- Built-in bearer token authentication
- Comprehensive device management capabilities
- Network operations support (circuit management, BGP configuration, routing)
- Real-time network monitoring and metrics collection
- Robust error handling with detailed error messages
- Extranet service configuration and monitoring
- Third-party integration capabilities
- Version constant accessible via `graphiant_sdk.Version`
- Example code demonstrating SDK usage
- Comprehensive documentation and README

### Documentation
- Initial README with installation instructions
- API reference documentation
- Usage examples
- License information

---

## Version History Summary

| Version | Release Date | Key Features |
|---------|--------------|--------------|
| 26.6.0 | 2026-07-08 | API v26.6.0; PVIF endpoints; OpenAPI Generator 7.23+ |
| 26.5.0 | 2026-06-10 | API v26.5.0; extranet public VIF endpoints |
| 26.4.0 | 2026-04-30 | API v26.4.0; content filters, software rollouts, MACsec |
| 26.3.2 | 2026-03-27 | Patch release v26.3.2; docs; same API bundle as 26.3.1 |
| 26.3.1 | 2026-03-26 | API v26.3.1 support, regenerated SDK/tests |
| 26.2.1 | 2026-02-26 | API v26.2.1 support, regenerated SDK/tests |
| 26.1.1 | 2026-02-01 | API v26.1.1 support |
| 25.12.1 | 2025-12-17 | API v25.12.1 support |
| 25.11.1 | 2025-11-11 | Major API optimization, schema reuse |
| 25.10.2 | 2025-10-15 | Hotfix release |
| 25.10.1 | 2025-10-08 | API v25.10.1 support |
| 25.9.2 | 2025-09-25 | API v25.9.2 support |
| 25.9.1 | 2025-09-23 | API v25.9.1 support |
| 25.8.1 | 2025-08-22 | PollAndPutDeviceConfig helper function |
| 25.7.1 | TBD | API v25.7.1 support |
| 25.6.2 | 2025-06-13 | CI/CD improvements |
| 25.6.1 | 2025-06-13 | Initial release |

---

## Types of Changes

- **Added** for new features
- **Changed** for changes in existing functionality
- **Deprecated** for soon-to-be removed features
- **Removed** for now removed features
- **Fixed** for any bug fixes
- **Security** for vulnerability fixes
