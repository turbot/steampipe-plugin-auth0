# Table: auth0_role

Role is a collection of permissions that you can apply to user.

## Examples


### Non-admin roles with 'all:*' permissions

```sql
select
  r.name as role_name,
  p ->> 'permission_name' as permission_name,
  p ->> 'description' as description,
  p ->> 'resource_server_name' as resource_server_name
from
  auth0_role r,
  jsonb_array_elements(permissions) p
where
  r.name <> 'admin' and
  p ->> 'permission_name' like 'all:%';
```

### Permissions of a role

```sql
select
  p ->> 'permission_name' as permission_name,
  p ->> 'description' as description,
  p ->> 'resource_server_name' as resource_server_name
from
  auth0_role r,
  jsonb_array_elements(permissions) p
where
  id = 'rol_AsrdJewEuFtKCWIe';
```

### Users of a role

```sql
select
  u ->> 'name' as name,
  u ->> 'email' as email,
  u ->> 'user_id' as user_id
from
  auth0_role r,
  jsonb_array_elements(users) u
where
  id = 'rol_AsrdJewEuFtKCWIe'
order by
  u ->> 'name';
```
