---
title: "Steampipe Table: auth0_connection - Query Auth0 Connections using SQL"
description: "Allows users to query Auth0 Connections, providing details about identity providers, status, options, and metadata associated with each connection."
---

# Table: auth0_connection - Query Auth0 Connections using SQL

Auth0 Connections are integrations with identity providers that allow users to authenticate with various credentials. These identity providers can be social like Google or Facebook, enterprise identity systems like Azure AD or LDAP, or a database of users stored directly within Auth0. Each connection represents a pipeline from an application, through Auth0, to the identity provider.

## Table Usage Guide

The `auth0_connection` table allows for deep insights into Auth0 Connections. As a security analyst, you can explore the details of each connection including the identity provider, status, options, and associated metadata. This table is particularly useful for auditing and ensuring the correct configuration of your identity providers.

## Examples

### Connections with MFA enabled
Explore which user connections have multi-factor authentication enabled. This is useful for assessing security measures and ensuring that extra layers of protection are in place.

```sql
select
  id,
  name,
  strategy
from
  auth0_connection
where
  options -> 'mfa' ->> 'active' = 'true';
```

### Connections with sign-up disabled
Explore which connections have not disabled the sign-up option, providing insights into potential areas of vulnerability or increased traffic. This can be useful for identifying and managing potential security risks or resource allocation.

```sql
select
  id,
  name,
  strategy
from
  auth0_connection
where
  options ->> 'disable_signup' = 'false';
```

### Connections with brute-force protection disabled
Identify instances where connections lack brute-force protection to enhance security measures and prevent unauthorized access.

```sql
select
  id,
  name,
  strategy
from
  auth0_connection
where
  options ->> 'brute_force_protection' = 'false';
```

### List password options for Auth0 database connections
Explore the password options for database connections in Auth0 to understand their policies, history, and additional security measures. This can aid in assessing the strength and security of your database connections.

```sql
select
  id,
  name,
  options ->> 'passwordPolicy' as "password_policy",
  options ->> 'password_dictionary' as "password_dictionary",
  options ->> 'password_history' as "password_history",
  options ->> 'password_no_personal_info' as "password_no_personal_info"
from
  auth0_connection
where
  strategy = 'auth0';
```

### List password complexity options for Auth0 database connections
Explore the password complexity options for your Auth0 database connections. This can help you assess and improve your system's security by understanding how complex the passwords need to be.

```sql
select
  id,
  name,
  options ->> 'password_complexity_options' as "password_complexity_options"
from
  auth0_connection
where
  strategy = 'auth0';
```

### Ensure connection's password policy requires a minimum length of 14 or greater
Determine the areas in which your connection's password policy meets or exceeds the recommended minimum length of 14 characters. This is important for enhancing the security of your connections by ensuring robust password policies.

```sql
select
  id,
  name,
  (options -> 'password_complexity_options' -> 'min_length')::integer >= 14 as "Min length > 14"
from
  auth0_connection
where
  strategy = 'auth0';
```