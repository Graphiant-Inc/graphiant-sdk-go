# Changelog

All notable changes to the Graphiant SDK Go will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

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
