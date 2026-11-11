# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed
- `WriteJSON` marshals before writing headers/status; preserves caller Content-Type.


## [0.1.0] - 2026-11-11

### Added

- Initial HTTP JSON writing helpers for `onyx-response` envelopes.
- Explicit status-code parameters; no domain-error mapping policy.

[Unreleased]: https://github.com/MahdiPezeshkian/onyx-http/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/MahdiPezeshkian/onyx-http/releases/tag/v0.1.0
