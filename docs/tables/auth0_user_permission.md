# Table: auth0_user_permission

Permissions assigned to user, including permissions assigned through users roles.

## Examples

### All user granted permissions

```sql
select
  p.permission_name,
  p.description,
  p.resource_server_name
from
  auth0_user u
join
  auth0_user_permission p
on
  p.user_id = u.id
where
  email = 'select-joey@mail.com'
order by
  p.resource_server_name,
  p.permission_name;
```

### User granted permissions in a resource server

```sql
select
  u.email,
  p.permission_name,
  p.description
from
  auth0_user u
join
  auth0_user_permission p
on
  p.user_id = u.id
where
  u.email = 'select-joey@mail.com' and
  p.resource_server_name = 'novel-mutt';
```
