---
title: "Steampipe Table: auth0_log - Query Auth0 Logs using SQL"
description: "Allows users to query Auth0 Logs, providing detailed insights into user activities, system events, and security-relevant incidents."
---

# Table: auth0_log - Query Auth0 Logs using SQL

Auth0 Logs is a resource within the Auth0 Identity Platform that records and stores user activities, system events, and security-relevant incidents. It is designed to give administrators detailed visibility into the behaviors and actions within their Auth0 environment. Auth0 Logs aids in monitoring, troubleshooting, and maintaining the health and security of Auth0 applications.

## Table Usage Guide

The `auth0_log` table provides insights into the logs within Auth0 Identity Platform. As a system administrator or security analyst, you can use this table to explore detailed log entries, including user activities, system events, and potential security incidents. This table is particularly useful for monitoring user behavior, troubleshooting issues, and enhancing the security posture of your Auth0 applications.

## Examples

### Failed login attempts
Identify instances where login attempts have failed to gain insights into potential security risks. This allows for a review of the associated IP addresses and user agents, enabling the detection and prevention of unauthorized access.

```sql+postgres
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

```sql+sqlite
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
Explore which authentication events are associated with a specific client ID. This can help in analyzing user behavior or troubleshooting issues related to a particular client application.

```sql+postgres
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

```sql+sqlite
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
Explore instances of account and IP blockings to understand potential security threats. This query helps in identifying suspicious activities by analyzing the patterns of blocked accounts and IP addresses.

```sql+postgres
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

```sql+sqlite
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
Analyze successful login patterns to understand the proportion of mobile versus non-mobile users. This can aid in tailoring user experiences based on device preference.

```sql+postgres
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

```sql+sqlite
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