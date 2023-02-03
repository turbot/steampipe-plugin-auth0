# Table: auth0_action

Actions are secure, tenant-specific, versioned functions written in Node.js that execute at certain points within the Auth0 platform.

## Examples


### Deprecated NodeJS 12 based code deployed 

```sql
select
  id,
  name,
  supported_triggers,
  updated_at
from
  auth0_action
where
  runtime = 'node12' and
  all_changes_deployed
```

### Post change password actions

```sql
select
  id,
  name,
  updated_at
from
  auth0_action
where
  supported_triggers -> 0 ->> 'id' = 'post-change-password'
```

### Action code by name

```sql
select
  code
from
  auth0_action
where
  name = 'send-notification'
```

### Deployed actions

```sql
select
  id,
  name,
  deployed_version ->> 'number' version,
  supported_triggers,
  updated_at
from
  auth0_action
where
  all_changes_deployed
```
