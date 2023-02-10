# Table: auth0_role_permission

Permissions assigned to role.

## Examples

### All the permissions assigned to a role

```sql
select
  permission_name,
  description,
  resource_server_name
from
  auth0_role_permission
where
  role_id = 'rol_VkaG05dncCpNN3oI'
order by
  resource_server_name,
  permission_name;
```

### List roles with assigned permission of a resource server

```sql
select
  r.name
from
  auth0_role r
join
  auth0_role_permission p
on
  p.role_id = r.id
where
  p.resource_server_name = 'novel-mutt'
group by
  r.name
order by
  r.name;
```
