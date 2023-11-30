---
title: "Steampipe Table: auth0_stats_daily - Query Auth0 Daily Statistics using SQL"
description: "Allows users to query daily statistics in Auth0, specifically the daily log event count, providing insights into application usage patterns and potential security incidents."
---

# Table: auth0_stats_daily - Query Auth0 Daily Statistics using SQL

Auth0 is an identity management platform for application builders and developers. It provides a universal authentication & authorization platform for web, mobile and legacy applications that focus on identity types like SSO, MFA, user management, and more. Auth0 is designed to make it easy for developers to manage user identities and create secure, seamless experiences for their users.

## Table Usage Guide

The `auth0_stats_daily` table provides insights into daily statistics within Auth0 Identity Management. As a security analyst, explore daily log event details through this table, including log event count, log event types, and associated metadata. Utilize it to uncover information about application usage patterns, such as peak usage times, most frequently used features, and potential security incidents.

## Examples

### Top 5 days with the highest number of sign-ups
Discover the dates that had the most user sign-ups, which can be useful for identifying trends or the impact of specific marketing campaigns.

```sql
select
  sign_ups,
  date
from
  auth0_stats_daily
order by
  sign_ups desc limit 5;
```

### Last 12 months average number of logins
Explore the average frequency of user logins over the past year to understand user engagement and activity trends. This can help in identifying patterns and informing strategies for user retention and engagement.

```sql
select
  avg(logins) as average_logins
from
  auth0_stats_daily
where
  date > current_date - interval '12' month;
```

### Days when the passwords were leaked
Identify instances where passwords have been compromised, allowing you to understand the severity and frequency of such security incidents.

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