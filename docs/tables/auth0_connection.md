# Table: auth0_connection

Connection is the relationship between Auth0 and a source of users.

## Examples


### Connections with MFA enabled

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

### Connections with sign up disabled

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

### Connections with brutal force protection

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

### Password settings

```sql
select
  id,
  name,
  options ->> 'passwordPolicy' as "password_policy",
  options ->> 'password_complexity_options' as "password_complexity_options",
  options ->> 'password_dictionary' as "password_dictionary",
  options ->> 'password_history' as "password_history",
  options ->> 'password_no_personal_info' as "password_no_personal_info"
from
  auth0_connection
where
  strategy = 'auth0';
```
