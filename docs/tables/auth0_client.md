# Table: auth0_client

Client is an application or a sso integration.

## Examples

### Number of clients by type

```sql
select
  app_type,
  count(1)
from
  auth0_client
group by
  app_type;
```

### Token lifetime

```sql
select
  client_id,
  name,
  refresh_token ->> 'token_lifetime' as token_lifetime,
  refresh_token ->> 'idle_token_lifetime' as idle_token_lifetime
from
  auth0_client
order by
  name;
```

### Grant types of a client

```sql
select
  g as grant_types
from
  auth0_client c,
  jsonb_array_elements(grant_types) g
where
  client_id = 'Jh5ap2mN94TJmZZ1sVeVmtW9Fpaim190';
```
