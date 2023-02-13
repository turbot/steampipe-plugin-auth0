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

### Connections with sign-up disabled

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

### Ensure connection password policy requires minimum length of 14 or greater 

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
