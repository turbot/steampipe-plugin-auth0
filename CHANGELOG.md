## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#33](https://github.com/turbot/steampipe-plugin-auth0/pull/33))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#33](https://github.com/turbot/steampipe-plugin-auth0/pull/33))

## v0.3.0 [2023-12-13]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/install/steampipe.sh), as a [Postgres FDW](https://steampipe.io/install/postgres.sh), as a [SQLite extension](https://steampipe.io/install/sqlite.sh) and as a standalone [exporter](https://steampipe.io/install/export.sh). ([#26](https://github.com/turbot/steampipe-plugin-auth0/pull/26))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#26](https://github.com/turbot/steampipe-plugin-auth0/pull/26))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-auth0/blob/main/docs/LICENSE). ([#26](https://github.com/turbot/steampipe-plugin-auth0/pull/26))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server enacapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#27](https://github.com/turbot/steampipe-plugin-auth0/pull/27))

## v0.2.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#16](https://github.com/turbot/steampipe-plugin-auth0/pull/16))

## v0.2.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#13](https://github.com/turbot/steampipe-plugin-auth0/pull/13))
- Recompiled plugin with Go version `1.21`. ([#13](https://github.com/turbot/steampipe-plugin-auth0/pull/13))

## v0.1.0 [2023-04-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which adds go-getter support to dynamic tables. ([#7](https://github.com/turbot/steampipe-plugin-auth0/pull/7))

## v0.0.2 [2023-02-25]

_Bug fixes_

- Fixed the brand color. ([#5](https://github.com/turbot/steampipe-plugin-auth0/pull/5))

## v0.0.1 [2023-02-17]

_What's new?_

- New tables added

  - [auth0_action](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_action)
  - [auth0_client](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_client)
  - [auth0_connection](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_connection)
  - [auth0_hook](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_hook)
  - [auth0_log](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_log)
  - [auth0_organization](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_organization)
  - [auth0_role](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_role)
  - [auth0_role_assigned_user](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_role_assigned_user)
  - [auth0_role_permission](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_role_permission)
  - [auth0_signing_key](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_signing_key)
  - [auth0_stats_daily](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_stats_daily)
  - [auth0_tenant_settings](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_tenant_settings)
  - [auth0_user](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_user)
  - [auth0_user_assigned_role](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_user_assigned_role)
  - [auth0_user_permission](https://hub.steampipe.io/plugins/turbot/auth0/tables/auth0_user_permission)
