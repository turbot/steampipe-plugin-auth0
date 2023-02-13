# Table: auth0_role

Role is a collection of permissions that you can apply to user.

## Examples

### Non-admin roles with 'all:*' permissions

```sql
select
  r.name as role_name,
  p.permission_name,
  p.description,
  p.resource_server_name
from
  auth0_role r
join
  auth0_role_permission p
on
  p.role_id = r.id
where
  r.name <> 'admin' and
  p.permission_name like 'all:%';
```

### List all permissions assigned to admin role

```sql
select
  p.permission_name,
  p.description,
  p.resource_server_name
from
  auth0_role r
join
  auth0_role_permission p
on
  p.role_id = r.id
where
  r.name = 'admin';
```

