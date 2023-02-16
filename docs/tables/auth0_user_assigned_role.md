# Table: auth0_user

List of roles assigned to user.

## Examples

### Admin users without MFA

```sql
select
  u.email,
  u.id,
  u.updated_at
from
  auth0_user u
  join
    auth0_user_assigned_role r
    on r.user_id = u.id
where
  r.name = 'admin'
  and u.multifactor is null;
```

### Roles a user is assigned to

```sql
select
  r.role_id,
  r.name,
  r.description
from
  auth0_user u
  join
    auth0_user_assigned_role r
    on r.user_id = u.id
where
  email = 'select-joey@mail.com';
```
