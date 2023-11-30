---
title: "Steampipe Table: auth0_tenant_settings - Query Auth0 Tenant Settings using SQL"
description: "Allows users to query Tenant Settings in Auth0, providing insights into tenant configuration details and settings."
---

# Table: auth0_tenant_settings - Query Auth0 Tenant Settings using SQL

Auth0 is a flexible, drop-in solution to add authentication and authorization services to your applications. Your software can use Auth0 to authenticate and authorize users with a wide variety of identity providers, including social, enterprise, and username/password databases. Tenant settings in Auth0 provide configuration details and settings for each tenant, including enabled features, default settings, and customization options.

## Table Usage Guide

The `auth0_tenant_settings` table provides insights into tenant settings within Auth0. As a security analyst or an application developer, explore tenant-specific details through this table, including enabled features, default settings, and other customization options. Utilize it to uncover information about each tenant's configuration, such as settings related to user registration, login, and identity providers.

## Examples

### SSO settings
Analyze the settings to understand if Single Sign-On (SSO) is enabled and if changes to these settings are permitted, helping to ensure secure access management.

```sql
select
  flags ->> 'enable_sso' as enable_sso,
  flags ->> 'allow_changing_enable_sso' as allow_changing_enable_sso
from
  auth0_tenant_settings;
```

### Enabled locales
Explore which locales have been enabled in your Auth0 tenant settings. This can help in understanding the geographical distribution of your user base.

```sql
select
  l as enabled_locales
from
  auth0_tenant_settings t,
  jsonb_array_elements(t.enabled_locales) l;
```

### Session and idle session lifetime settings
Analyze the settings to understand the duration of active and idle sessions within your Auth0 tenant. This can help optimize user experience by managing session lengths and idle times.

```sql
select
  session_lifetime,
  idle_session_lifetime
from
  auth0_tenant_settings;
```