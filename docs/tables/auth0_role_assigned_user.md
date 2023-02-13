# Table: auth0_role_assigned_user

List of users assigned to role.

## Examples

### List users assigned to a role

```sql
select
  u.name,
  u.email,
  u.user_id
from
  auth0_role r
join
  auth0_role_assigned_user u
on
  u.role_id = r.id
where
  r.name = 'operator'
order by
  u.name;
```

### Admin users with unverified email

```sql
select
  u.email,
  u.id,
  u.updated_at
from
  auth0_role r
join
  auth0_role_assigned_user ru
on
  ru.role_id = r.id
join
  auth0_user u
on
  u.id = ru.user_id
where
  r.name = 'admin' and
  not u.email_verified;
```
