# Table: auth0_log

Logs are all the events that occur in your tenants including user authentication and administrative actions such as adding and updating applications, connections, and rules.
Log type codes: https://auth0.com/docs/deploy-monitor/logs/log-event-type-codes

## Examples

### Failed login attempts

```sql
select
  date,
  description,
  ip,
  user_agent
from
  auth0_log
where
  type = 'f'
order by
  date desc;
```

### Logs filtered by client

```sql
select
  date,
  description,
  ip,
  is_mobile
from
  auth0_log
where
  client_id = 'FrSZDDFGUH0afar5LHmCji1khmPmst6R'
order by
  date desc;
```

### Account and IP blockings

```sql
select
  date,
  description,
  ip,
  is_mobile
from
  auth0_log
where
  type in ('limit_mu', 'limit_wc', 'limit_sul')
order by
  date desc;
```

### Number of mobile and non-mobile successful logins

```sql
select
  is_mobile,
  count(1)
from
  auth0_log
where
  type = 's'
group by
  is_mobile;
```
