# Table: auth0_stats_daily

Daily stats for an Auth0 Tenant.

## Examples

### Top 5 days with the highest number of sign-ups

```sql
select
  sign_ups,
  date
from
  auth0_stats_daily
order by
  sign_ups desc
limit 5;
```

### Last 12 months average number of logins

```sql
select
  avg(logins) as average_logins
from
  auth0_stats_daily
where
  date > current_date - interval '12' month;
```

### Days there were leaked passwords

```sql
select
  leaked_passwords,
  date
from
  auth0_stats_daily
where
  leaked_passwords != 0
order by
  leaked_passwords desc;
```
