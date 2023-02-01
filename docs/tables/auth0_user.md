# Table: auth0_user

User represents an Auth0 user resource.

## Examples

### Users without MFA

```sql
select
  email,
  id,
  updated_at
from
  auth0_user
where
  multifactor is null;
```

### Admin users without MFA

```sql
select
  email,
  id,
  updated_at
from
  auth0_user,
  jsonb_array_elements(roles) r
where
  r->> 'name' = 'admin' and
  multifactor is null;
```

### Users with unverified email

```sql
select
  email,
  id,
  updated_at
from
  auth0_user
where
  not email_verified;
```

### Ranking of most used auth connections

```sql
select
  i ->> 'connection' as "connection",
  count(1)
from
  auth0_user u,
  jsonb_array_elements(u.identities) i
group by
  i ->> 'connection'
order by
  count desc;
```

### Users signed up through GitHub

```sql
select
  nickname,
  id,
  last_login
from
  auth0_user
where
  identities -> 0 ->> 'connection' = 'github';
```

### Granted permissions

```sql
select
  p ->> 'permission_name' as permission_name,
  p ->> 'description' as description,
  p ->> 'resource_server_name' as resource_server_name
from
  auth0_user,
  jsonb_array_elements(permissions) p
where
  email = 'select-joey@mail.com';
```

### Roles assigned to

```sql
select
  r ->> 'id' as id,
  r ->> 'name' as name,
  r ->> 'description' as description
from
  auth0_user,
  jsonb_array_elements(roles) r
where
  email = 'select-joey@mail.com';
```
