# Changelog

All notable changes to the Graphiant SDK Go will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Support for API version 25.12.1

## [25.12.1] - TBD

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
| 25.12.1 | TBD | API v25.12.1 support |
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
