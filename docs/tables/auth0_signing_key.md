# Table: auth0_signing_key

Signing keys are used to sign ID tokens, access tokens, SAML assertions, and WS-Fed assertions sent to your application or API.

## Examples

### For how long has current signing key been available

```sql
select
  current_date - current_since as current_for
from
  auth0_signing_key
where
  current;
```

### Next signing key

```sql
select
  kid,
  fingerprint,
  thumbprint
from
  auth0_signing_key
where
  next;
```

### Previous signing key

```sql
select
  kid,
  fingerprint,
  thumbprint,
  current_since,
  current_until
from
  auth0_signing_key
where
  previous;
```

### Average time for which the previous signing keys were available

```sql
select
  avg(current_until - current_since) as average_duration
from
  auth0_signing_key
where
  previous;
```

### Revoked signing keys

```sql
select
  kid,
  fingerprint,
  thumbprint,
  revoked_at
from
  auth0_signing_key
where
  revoked
order by
  revoked_at desc;
```
