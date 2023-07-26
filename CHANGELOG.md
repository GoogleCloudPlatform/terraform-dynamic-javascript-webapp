# Changelog

All notable changes to this project will be documented in this file.

The format is based on
[Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).
This changelog is generated automatically based on [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

## [1.3.0](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/compare/v1.2.0...v1.3.0) (2023-07-26)


### Features

* define verify for int tests ([#64](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/64)) ([2c6bfcc](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/2c6bfcce44b33026707c487b5e9b36ad930a467c))
* enable integration tests ([#61](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/61)) ([a75938c](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/a75938cb34c4b732a8b6ec5be65b951e72321fb4))
* **testing:** add verification tests for this blueprint ([#67](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/67)) ([3b5aacf](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/3b5aacf46559134e0fae37ecbf384f22c7adf518))


### Bug Fixes

* correct fileset() globbing when this module is not the root. ([#68](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/68)) ([9ecf36a](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/9ecf36a18dd72378687c3ed473ca749236c1c345))
* enable uniform bucket level access ([#51](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/51)) ([4a4cf2d](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/4a4cf2dd953a9b5063bfd35c5e0dd4ad252ffd6d))
* pinning google provider &lt; 4.75.0 ([#69](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/69)) ([cf01bae](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/cf01baef81fbbb12aa3814029ab2f228c174bc36))
* remove db dependency on LB time delay ([#65](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/65)) ([3f71c49](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/3f71c491f1159e3fb5437beafcce938e2f7877e2))

## [1.2.0](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/compare/v1.1.6...v1.2.0) (2023-06-22)


### Features

* adds metadata generation and fixes lint issues ([#43](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/43)) ([bcae2a1](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/bcae2a13a12d49d517410eadb358459706b8d7db))


### Bug Fixes

* add delay to iam service account resource ([#47](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/47)) ([889f21c](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/889f21ccb117966a4f47b05befb8e0d2cb6b926a))
* neos toc url ([#46](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/46)) ([8a4a24f](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/8a4a24f85d785c2bfbcc5e19c8d07460e4944a08))
* update neos toc url ([#50](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/issues/50)) ([1ab6c7f](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/1ab6c7f816d0aaffe6e0fd695185e5751b348c3d))

## [1.1.6](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/compare/v1.1.5...v1.1.6) (2023-04-26)


### Bug Fixes

* adding in random suffix to service account to handle multiple deployment ([932b1a8](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/932b1a89ff5394d37a936d592f6ff0b20bf3c789))

## [1.1.5](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/compare/v1.1.4...v1.1.5) (2023-04-26)


### Bug Fixes

* Update outputs.tf to stay at solutions deployments page ([e8d8b5b](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/e8d8b5b5352d72b92cd8c8033c422c455e3e8552))

## [1.1.2](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/compare/v1.1.1...v1.1.2) (2023-04-21)


### Bug Fixes

* shortening account_id name for service account to fit regex max length ([18b0b19](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/18b0b192346789ebd311b791a63aa0c9b24e8edc))

## [1.1.1](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/compare/v1.1.0...v1.1.1) (2023-04-20)


### Bug Fixes

* buffing up startup probe time for cloud run v2 resource ([c340e2b](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/c340e2bba55151c129074875d91354ec10437c82))

## [1.1.0](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/compare/v1.0.1...v1.1.0) (2023-04-19)


### Features

* adding maintenance items ([0de8607](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/0de8607c038988aa4d3f334700262d3c8556016d))


### Bug Fixes

* yaml file type fix ([f8a997c](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/f8a997c4db05aa538a9c28173111d354f1b33863))

## [1.0.1](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/compare/v1.0.0...v1.0.1) (2023-04-18)


### Bug Fixes

* flip firestore_enabled ([1c8b545](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/1c8b545276c978384812654240a6fcad739a235d))
* remove beta resource, add in missing env for cloud run container ([a7b52c6](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/a7b52c638513631ed2a40bd895aa3525f4918de1))

## 1.0.0 (2023-04-12)


### Features

* adding initial solutions base ([09be428](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/09be428619fe3cdecff81117897fcde2e781c987))


### Bug Fixes

* add release-please bot ([5e6ad95](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/5e6ad95868f5c15576cb21e50a31e3cd88d5f39a))
* make sure that output generates neos_toc_url ([d12787c](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/commit/d12787c718b5068583201eb1afd0b76bf7e3d791))

## [0.1.0](https://github.com/terraform-google-modules/terraform-google-dynamic-javascript-webapp/releases/tag/v0.1.0) - 20XX-YY-ZZ

### Features

- Initial release

[0.1.0]: https://github.com/terraform-google-modules/terraform-google-dynamic-javascript-webapp/releases/tag/v0.1.0
