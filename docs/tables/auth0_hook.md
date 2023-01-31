# Table: auth0_hook

Hooks allow you to customize the behavior of Auth0 with Node.js code that is executed in selected extension points.

## Examples


### Enabled hook scripts

```sql
select
  id,
  name,
  dependencies,
  trigger_id,
  script
from
  auth0_hook
where
 enabled;
```

### Post change password script

```sql
select
  id,
  name,
  script,
  dependencies,
  enabled
from
  auth0_hook
where
 trigger_id = 'post-change-password';
```

### User registration scripts

```sql
select
  id,
  name,
  script,
  dependencies,
  enabled
from
  auth0_hook
where
 trigger_id in ('pre-user-registration', 'post-user-registration');
```
