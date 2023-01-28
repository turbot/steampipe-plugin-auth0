# Table: auth0_tenant_settings

Tenant settings definitions.

## Examples


### SSO settings

```sql
select
  flags ->> 'enable_sso' as enable_sso,
  flags ->> 'allow_changing_enable_sso' as allow_changing_enable_sso
from
  auth0_tenant_settings
```

### Enabled locales

```sql
select
  l as enabled_locales
from
  auth0_tenant_settings t,
  jsonb_array_elements(t.enabled_locales) l
```

### Session and idle session lifetime settings

```sql
select
  session_lifetime,
  idle_session_lifetime
from
  auth0_tenant_settings
```
