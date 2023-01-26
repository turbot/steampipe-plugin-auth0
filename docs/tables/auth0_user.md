# Table: auth0_user

User represents an Auth0 user resource.

## Examples

### Users with unverified email

```sql
select
  email,
  id,
  name,
  given_name
from
  auth0_user
where
  not email_verified
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
  count desc
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
  identities -> 0 ->> 'connection' = 'github'
```
