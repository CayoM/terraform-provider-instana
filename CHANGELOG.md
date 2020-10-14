# Changelog

## [v0.9.1](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.9.1) (2020-10-14)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.9.0...v0.9.1)

**Fixed bugs:**

- Breaking Change of EQUAL Condition Operator in Release 188 [\#69](https://github.com/gessnerfl/terraform-provider-instana/issues/69)

**Merged pull requests:**

- Bugfix/69 equal condition operator [\#70](https://github.com/gessnerfl/terraform-provider-instana/pull/70) ([gessnerfl](https://github.com/gessnerfl))

## [v0.9.0](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.9.0) (2020-08-27)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.8.2...v0.9.0)

**Fixed bugs:**

- Breaking change of matching\_operator in Instana API [\#48](https://github.com/gessnerfl/terraform-provider-instana/issues/48)
- Idempotence in "instana\_alerting\_config" only with priorities/event\_filter\_event\_types in alphabetical order [\#43](https://github.com/gessnerfl/terraform-provider-instana/issues/43)

**Closed issues:**

- Adjust user role to changes in REST API [\#65](https://github.com/gessnerfl/terraform-provider-instana/issues/65)
- Migrate to Github Actions [\#63](https://github.com/gessnerfl/terraform-provider-instana/issues/63)
- Add github pages based documentation [\#61](https://github.com/gessnerfl/terraform-provider-instana/issues/61)
- Allow Instana API Representation for Threshold Rule Matching Operator also in Terraform [\#54](https://github.com/gessnerfl/terraform-provider-instana/issues/54)
- Add boundary scope to application configuration [\#51](https://github.com/gessnerfl/terraform-provider-instana/issues/51)
- Add new operators to application config [\#50](https://github.com/gessnerfl/terraform-provider-instana/issues/50)
- Add Event Type `agent\_monitoring\_issue` [\#49](https://github.com/gessnerfl/terraform-provider-instana/issues/49)
- Add support for custom event configurations of rule type "Threshold Rule using a dynamic built-in metric by pattern" [\#47](https://github.com/gessnerfl/terraform-provider-instana/issues/47)

**Merged pull requests:**

- \#65: add new and remove deprecated fields of user role [\#66](https://github.com/gessnerfl/terraform-provider-instana/pull/66) ([gessnerfl](https://github.com/gessnerfl))
- Feature/63 GitHub actions [\#64](https://github.com/gessnerfl/terraform-provider-instana/pull/64) ([gessnerfl](https://github.com/gessnerfl))
- Create github pages based documentation [\#62](https://github.com/gessnerfl/terraform-provider-instana/pull/62) ([gessnerfl](https://github.com/gessnerfl))
- Bugfix/47 fix dynamic built in metrics [\#60](https://github.com/gessnerfl/terraform-provider-instana/pull/60) ([gessnerfl](https://github.com/gessnerfl))
- Feature/51 app config boundary scope [\#59](https://github.com/gessnerfl/terraform-provider-instana/pull/59) ([gessnerfl](https://github.com/gessnerfl))
- \#49 add new supported event type agent\_monitoring\_issue [\#57](https://github.com/gessnerfl/terraform-provider-instana/pull/57) ([gessnerfl](https://github.com/gessnerfl))
- \#50 Add additional comparison operators for application configurations [\#56](https://github.com/gessnerfl/terraform-provider-instana/pull/56) ([gessnerfl](https://github.com/gessnerfl))
- \#54: allow instana representation in addition to terraform for matchi… [\#55](https://github.com/gessnerfl/terraform-provider-instana/pull/55) ([gessnerfl](https://github.com/gessnerfl))
- Feature/47 custom events dynamic filter [\#53](https://github.com/gessnerfl/terraform-provider-instana/pull/53) ([gessnerfl](https://github.com/gessnerfl))
- Feature/48 fix api changes [\#52](https://github.com/gessnerfl/terraform-provider-instana/pull/52) ([gessnerfl](https://github.com/gessnerfl))

## [v0.8.2](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.8.2) (2020-03-02)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.8.1...v0.8.2)

**Fixed bugs:**

- Downstream integration Ids not supported anymore for Custom Event [\#44](https://github.com/gessnerfl/terraform-provider-instana/issues/44)

**Closed issues:**

- Consistently use testify assert [\#41](https://github.com/gessnerfl/terraform-provider-instana/issues/41)
- Allow up to 1024 rule-ids per instana\_alerting\_config [\#39](https://github.com/gessnerfl/terraform-provider-instana/issues/39)

**Merged pull requests:**

- Bugfix/43 order of event types [\#46](https://github.com/gessnerfl/terraform-provider-instana/pull/46) ([gessnerfl](https://github.com/gessnerfl))
- \#44: Remove downstream integration IDs from custom event specs [\#45](https://github.com/gessnerfl/terraform-provider-instana/pull/45) ([gessnerfl](https://github.com/gessnerfl))
- \#41: use testify instead of plain go if checks in tests [\#42](https://github.com/gessnerfl/terraform-provider-instana/pull/42) ([gessnerfl](https://github.com/gessnerfl))

## [v0.8.1](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.8.1) (2020-02-21)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.8.0...v0.8.1)

**Merged pull requests:**

- \#39 Allow 1024 rule-ids per instana\_alerting\_config [\#40](https://github.com/gessnerfl/terraform-provider-instana/pull/40) ([ppuschmann](https://github.com/ppuschmann))

## [v0.8.0](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.8.0) (2020-02-14)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.7.0...v0.8.0)

**Closed issues:**

- Migrate to new resource approach [\#35](https://github.com/gessnerfl/terraform-provider-instana/issues/35)
- REST Client should support retries [\#32](https://github.com/gessnerfl/terraform-provider-instana/issues/32)
- Alerting Configuration [\#30](https://github.com/gessnerfl/terraform-provider-instana/issues/30)
- Alerting Channel Configuration [\#29](https://github.com/gessnerfl/terraform-provider-instana/issues/29)

**Merged pull requests:**

- Feature/32 rest throttling [\#38](https://github.com/gessnerfl/terraform-provider-instana/pull/38) ([gessnerfl](https://github.com/gessnerfl))
- Feature/30 alerting configuration [\#37](https://github.com/gessnerfl/terraform-provider-instana/pull/37) ([gessnerfl](https://github.com/gessnerfl))
- Feature/35 new resource approach [\#36](https://github.com/gessnerfl/terraform-provider-instana/pull/36) ([gessnerfl](https://github.com/gessnerfl))
- Feature/29 altering channels [\#34](https://github.com/gessnerfl/terraform-provider-instana/pull/34) ([gessnerfl](https://github.com/gessnerfl))

## [v0.7.0](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.7.0) (2019-12-17)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.6.0...v0.7.0)

**Closed issues:**

- Add support for Entity Verification Rule Type [\#27](https://github.com/gessnerfl/terraform-provider-instana/issues/27)

**Merged pull requests:**

- Feature/27 entity verification events [\#28](https://github.com/gessnerfl/terraform-provider-instana/pull/28) ([gessnerfl](https://github.com/gessnerfl))

## [v0.6.0](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.6.0) (2019-12-16)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.5.0...v0.6.0)

**Fixed bugs:**

- Threshold rule support window and rule together [\#24](https://github.com/gessnerfl/terraform-provider-instana/issues/24)

**Closed issues:**

- Update to terraform 0.12.x [\#20](https://github.com/gessnerfl/terraform-provider-instana/issues/20)

**Merged pull requests:**

- \#20: Update project to terraform 0.12.x [\#26](https://github.com/gessnerfl/terraform-provider-instana/pull/26) ([gessnerfl](https://github.com/gessnerfl))
- \#24: Support rollup and window in threshold rule together [\#25](https://github.com/gessnerfl/terraform-provider-instana/pull/25) ([gessnerfl](https://github.com/gessnerfl))

## [v0.5.0](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.5.0) (2019-10-15)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.4.0...v0.5.0)

**Closed issues:**

- Support for label/name prefix and suffix [\#22](https://github.com/gessnerfl/terraform-provider-instana/issues/22)

**Merged pull requests:**

- \#22: migrate to customizable default name prefix and suffix [\#23](https://github.com/gessnerfl/terraform-provider-instana/pull/23) ([gessnerfl](https://github.com/gessnerfl))

## [v0.4.0](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.4.0) (2019-10-14)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.3.2...v0.4.0)

**Closed issues:**

- Add support to append terraform managed string [\#19](https://github.com/gessnerfl/terraform-provider-instana/issues/19)

**Merged pull requests:**

- Feature/19 append terraform managed string [\#21](https://github.com/gessnerfl/terraform-provider-instana/pull/21) ([gessnerfl](https://github.com/gessnerfl))

## [v0.3.2](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.3.2) (2019-06-19)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.3.1...v0.3.2)

**Fixed bugs:**

- Terraform provider should not have platform name in executable [\#17](https://github.com/gessnerfl/terraform-provider-instana/issues/17)

**Merged pull requests:**

- \#17: fix binary name [\#18](https://github.com/gessnerfl/terraform-provider-instana/pull/18) ([gessnerfl](https://github.com/gessnerfl))

## [v0.3.1](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.3.1) (2019-06-19)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.3.0...v0.3.1)

## [v0.3.0](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.3.0) (2019-06-19)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.2.2...v0.3.0)

**Closed issues:**

- Change release output [\#15](https://github.com/gessnerfl/terraform-provider-instana/issues/15)

**Merged pull requests:**

- Feature/15 release naming [\#16](https://github.com/gessnerfl/terraform-provider-instana/pull/16) ([gessnerfl](https://github.com/gessnerfl))

## [v0.2.2](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.2.2) (2019-05-07)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.2.1...v0.2.2)

**Closed issues:**

- Add support for dashes in tags in match\_specification [\#12](https://github.com/gessnerfl/terraform-provider-instana/issues/12)

**Merged pull requests:**

- Fix application configuration example in README [\#14](https://github.com/gessnerfl/terraform-provider-instana/pull/14) ([steinex](https://github.com/steinex))

## [v0.2.1](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.2.1) (2019-05-07)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.2.0...v0.2.1)

**Merged pull requests:**

- Feature/12 add support for dashes in identifiers [\#13](https://github.com/gessnerfl/terraform-provider-instana/pull/13) ([gessnerfl](https://github.com/gessnerfl))

## [v0.2.0](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.2.0) (2019-04-25)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/v0.1.0...v0.2.0)

**Closed issues:**

- Severity for Rules are not user friendly [\#8](https://github.com/gessnerfl/terraform-provider-instana/issues/8)
- Add support to manage events [\#7](https://github.com/gessnerfl/terraform-provider-instana/issues/7)
- Migrate to OpenAPI [\#4](https://github.com/gessnerfl/terraform-provider-instana/issues/4)
- Add support to manage groups  [\#3](https://github.com/gessnerfl/terraform-provider-instana/issues/3)
- Add support to create Application Perspectives [\#1](https://github.com/gessnerfl/terraform-provider-instana/issues/1)

**Merged pull requests:**

- Feature/7 events [\#11](https://github.com/gessnerfl/terraform-provider-instana/pull/11) ([gessnerfl](https://github.com/gessnerfl))
- Feature/1 application perspective [\#10](https://github.com/gessnerfl/terraform-provider-instana/pull/10) ([gessnerfl](https://github.com/gessnerfl))
- \#8 Change severity to a user friendly text instead of int codes [\#9](https://github.com/gessnerfl/terraform-provider-instana/pull/9) ([gessnerfl](https://github.com/gessnerfl))
- Feature/3 manage groups [\#6](https://github.com/gessnerfl/terraform-provider-instana/pull/6) ([gessnerfl](https://github.com/gessnerfl))
- Feature/4 migrate to open api [\#5](https://github.com/gessnerfl/terraform-provider-instana/pull/5) ([gessnerfl](https://github.com/gessnerfl))

## [v0.1.0](https://github.com/gessnerfl/terraform-provider-instana/tree/v0.1.0) (2019-03-14)

[Full Changelog](https://github.com/gessnerfl/terraform-provider-instana/compare/627e6874cfda8cf8e5d5793f016aaf60b5285e6f...v0.1.0)



\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*
