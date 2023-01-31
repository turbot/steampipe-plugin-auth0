# Table: auth0_organization

Organization is used to allow B2B customers to better manage their partners and customers, and to customize the ways that end-users access their applications.

## Examples


### List of my organizations

```sql
select
  name,
  display_name,
  branding ->> 'logo_url' as logo_url,
  metadata
from
  auth0_organization
order by
  name;
```

### Filter my organizations by metadata tags

```sql
select
  name,
  display_name,
  branding ->> 'logo_url' as logo_url,
  metadata
from
  auth0_organization
where
  metadata ->> 'cost_id' = 'e42345'
order by
  name;
```
